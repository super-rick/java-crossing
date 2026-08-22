# 测试:API 集成(TestClient ≈ MockMvc)+ 并发幂等(threading 压测)
import threading

from fastapi.testclient import TestClient

from app import app, store

client = TestClient(app)


def test_create_get_pay():
    r = client.post("/orders", json={"item": "coffee", "quantity": 2})
    assert r.status_code == 201
    oid = r.json()["id"]

    assert client.get(f"/orders/{oid}").status_code == 200

    r = client.post(f"/orders/{oid}/pay")
    assert r.status_code == 200
    assert r.json() == {"paid": True}
    assert store.inventory("coffee") == 98


def test_concurrent_pay_idempotent():
    # 100 个并发支付同一订单:只扣一次库存,全部幂等成功
    r = client.post("/orders", json={"item": "coffee", "quantity": 5})
    oid = r.json()["id"]
    before = store.inventory("coffee")
    errors = []

    def pay():
        try:
            rr = client.post(f"/orders/{oid}/pay")
            assert rr.status_code == 200, rr.text
        except Exception as e:  # noqa: BLE001
            errors.append(e)

    threads = [threading.Thread(target=pay) for _ in range(100)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()

    assert not errors, errors[:3]
    assert store.inventory("coffee") == before - 5  # 只扣一次


def test_insufficient_stock():
    r = client.post("/orders", json={"item": "coffee", "quantity": 9999})
    assert r.status_code == 409


def test_not_found():
    assert client.get("/orders/999").status_code == 404
    assert client.post("/orders/999/pay").status_code == 404


def test_validation():
    # pydantic 自动校验(≈ Bean Validation):quantity 必须 > 0
    r = client.post("/orders", json={"item": "coffee", "quantity": 0})
    assert r.status_code == 422
