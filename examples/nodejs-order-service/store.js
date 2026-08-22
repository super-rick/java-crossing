// 订单存储:内存实现
// 注意:Node 单线程,JS 语句天然原子(无数据竞争,无需锁)
// 幂等性靠 paid 标志;多实例部署时状态需外部化(Redis/DB)——见 README
class Store {
  constructor() {
    this.orders = new Map();
    this.inventory = new Map([["coffee", 100]]);
    this.nextId = 1;
  }

  createOrder(item, quantity) {
    if ((this.inventory.get(item) ?? 0) < quantity) {
      const err = new Error("insufficient stock");
      err.status = 409;
      throw err;
    }
    const order = { id: this.nextId++, item, quantity, paid: false };
    this.orders.set(order.id, order);
    return order;
  }

  getOrder(id) {
    const order = this.orders.get(id);
    if (!order) {
      const err = new Error("order not found");
      err.status = 404;
      throw err;
    }
    return order;
  }

  payOrder(id) {
    const order = this.getOrder(id);
    if (order.paid) return; // 幂等:已支付直接成功
    if ((this.inventory.get(order.item) ?? 0) < order.quantity) {
      const err = new Error("insufficient stock");
      err.status = 409;
      throw err;
    }
    this.inventory.set(order.item, this.inventory.get(order.item) - order.quantity);
    order.paid = true;
  }

  inventoryOf(item) {
    return this.inventory.get(item) ?? 0;
  }

  totalOrders() {
    return this.nextId - 1;
  }
}

module.exports = { Store };
