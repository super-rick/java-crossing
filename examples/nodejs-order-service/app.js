// 订单服务入口:Express
// 契约(与其他语言实现一致):
//   POST /orders            创建订单 -> 201 {order} / 400 / 409
//   GET  /orders/:id        查询订单 -> 200 {order} / 404
//   POST /orders/:id/pay    幂等支付 -> 200 {"paid":true} / 404 / 409
//   GET  /healthz           健康检查 -> 200
//   GET  /metrics           Prometheus 文本指标 -> 200
const express = require("express");
const { Store } = require("./store");

const app = express();
const store = new Store();

app.use(express.json());

// POST /orders 创建订单
app.post("/orders", (req, res) => {
  const { item, quantity } = req.body ?? {};
  if (!item || !Number.isInteger(quantity) || quantity <= 0) {
    return res.status(400).json({ error: "invalid request" });
  }
  try {
    const order = store.createOrder(item, quantity);
    res.status(201).json(order);
  } catch (e) {
    res.status(e.status ?? 500).json({ error: e.message });
  }
});

// GET /orders/:id 查询
app.get("/orders/:id", (req, res) => {
  const id = Number(req.params.id);
  try {
    res.json(store.getOrder(id));
  } catch (e) {
    res.status(e.status ?? 500).json({ error: e.message });
  }
});

// POST /orders/:id/pay 幂等支付
app.post("/orders/:id/pay", (req, res) => {
  const id = Number(req.params.id);
  try {
    store.payOrder(id);
    res.json({ paid: true });
  } catch (e) {
    res.status(e.status ?? 500).json({ error: e.message });
  }
});

app.get("/healthz", (req, res) => res.json({ status: "ok" }));

app.get("/metrics", (req, res) => {
  res.type("text/plain").send(`orders_total ${store.totalOrders()}\n`);
});

module.exports = { app, store };
