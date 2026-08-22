// 测试:Jest + supertest(≈ MockMvc)+ 并发幂等
const request = require("supertest");
const { app, store } = require("./app");

test("create/get/pay 流程", async () => {
  const created = await request(app)
    .post("/orders")
    .send({ item: "coffee", quantity: 2 })
    .expect(201);
  const oid = created.body.id;

  await request(app).get(`/orders/${oid}`).expect(200);

  const paid = await request(app).post(`/orders/${oid}/pay`).expect(200);
  expect(paid.body).toEqual({ paid: true });
  expect(store.inventoryOf("coffee")).toBe(98);
});

test("并发幂等:100 并发支付只扣一次", async () => {
  const created = await request(app)
    .post("/orders")
    .send({ item: "coffee", quantity: 5 })
    .expect(201);
  const oid = created.body.id;
  const before = store.inventoryOf("coffee");

  // 单线程下同步 handler 天然原子;并发 = 异步请求交错
  const results = await Promise.all(
    Array.from({ length: 100 }, () => request(app).post(`/orders/${oid}/pay`).expect(200))
  );
  expect(results.length).toBe(100);
  expect(store.inventoryOf("coffee")).toBe(before - 5); // 只扣一次
});

test("库存不足 409", async () => {
  await request(app)
    .post("/orders")
    .send({ item: "coffee", quantity: 9999 })
    .expect(409);
});

test("不存在 404", async () => {
  await request(app).get("/orders/999").expect(404);
  await request(app).post("/orders/999/pay").expect(404);
});

test("参数校验 400", async () => {
  await request(app).post("/orders").send({ item: "coffee", quantity: 0 }).expect(400);
});
