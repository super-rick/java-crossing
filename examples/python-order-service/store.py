# 订单存储:内存实现,并发安全(threading.Lock ≈ synchronized)
import threading


class Store:
    def __init__(self):
        self._lock = threading.Lock()
        self._orders = {}
        self._inventory = {"coffee": 100}
        self._next_id = 1

    def create_order(self, item, quantity):
        with self._lock:  # ≈ synchronized 块,with 自动释放
            if self._inventory.get(item, 0) < quantity:
                raise ValueError("insufficient stock")
            oid = self._next_id
            self._next_id += 1
            self._orders[oid] = {
                "id": oid, "item": item, "quantity": quantity, "paid": False,
            }
            return dict(self._orders[oid])

    def get_order(self, oid):
        with self._lock:
            if oid not in self._orders:
                raise KeyError("order not found")
            return dict(self._orders[oid])

    def pay_order(self, oid):
        with self._lock:
            if oid not in self._orders:
                raise KeyError("order not found")
            o = self._orders[oid]
            if o["paid"]:
                return  # 幂等:已支付直接成功
            if self._inventory.get(o["item"], 0) < o["quantity"]:
                raise ValueError("insufficient stock")
            self._inventory[o["item"]] -= o["quantity"]
            o["paid"] = True

    def inventory(self, item):
        with self._lock:
            return self._inventory.get(item, 0)

    def total_orders(self):
        with self._lock:
            return self._next_id - 1
