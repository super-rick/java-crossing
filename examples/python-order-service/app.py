# 订单服务入口:FastAPI(类型注解即校验 + 自动 OpenAPI 文档)
# 契约(与其他语言实现一致):
#   POST /orders            创建订单 -> 201 {order} / 409 库存不足
#   GET  /orders/{id}       查询订单 -> 200 {order} / 404
#   POST /orders/{id}/pay   幂等支付(并发安全) -> 200 {"paid":true} / 404 / 409
#   GET  /healthz           健康检查 -> 200
#   GET  /metrics           Prometheus 文本指标 -> 200
from fastapi import FastAPI, HTTPException
from fastapi.responses import Response
from pydantic import BaseModel, Field

from store import Store

app = FastAPI(title="order-service (python)")
store = Store()


class CreateOrderReq(BaseModel):
    item: str = Field(min_length=1)
    quantity: int = Field(gt=0)


@app.post("/orders", status_code=201)
def create_order(req: CreateOrderReq):
    try:
        return store.create_order(req.item, req.quantity)
    except ValueError as e:
        raise HTTPException(409, str(e))


@app.get("/orders/{oid}")
def get_order(oid: int):
    try:
        return store.get_order(oid)
    except KeyError as e:
        raise HTTPException(404, str(e))


@app.post("/orders/{oid}/pay")
def pay_order(oid: int):
    try:
        store.pay_order(oid)
    except KeyError as e:
        raise HTTPException(404, str(e))
    except ValueError as e:
        raise HTTPException(409, str(e))
    return {"paid": True}


@app.get("/healthz")
def healthz():
    return {"status": "ok"}


@app.get("/metrics")
def metrics():
    # 极简 Prometheus 文本(完整接入用 prometheus_client,见 09-监控模块)
    return Response(
        content=f"orders_total {store.total_orders()}\n",
        media_type="text/plain",
    )
