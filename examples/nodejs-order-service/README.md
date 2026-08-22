# nodejs-order-service(Node.js 实现)

> 锚点:用 Express 实现与 Go/Python 相同的订单服务契约,展示「单线程异步 + 无锁原子」的 Node 生产形态。
>
> 验证环境:Node 22.23.0 本地实测 / CI: Ubuntu 24.04 + Node 22。

## 接口契约(与其他语言实现一致)

| 方法 | 路径 | 说明 | 成功 | 失败 |
|---|---|---|---|---|
| POST | /orders | 创建订单 `{"item":"coffee","quantity":2}` | 201 order | 400 参数错 / 409 库存不足 |
| GET | /orders/:id | 查询订单 | 200 order | 404 |
| POST | /orders/:id/pay | **幂等支付(并发安全)** | 200 {"paid":true} | 404 / 409 |
| GET | /healthz | 健康检查 | 200 | — |
| GET | /metrics | Prometheus 文本指标 | 200 | — |

## 运行

```bash
npm ci              # 按 lock 精确安装(≈ 版本钉死)
npm start           # 默认 :8080,可用 PORT 环境变量覆盖
npm test            # Jest 5 个测试(含 100 并发幂等)
```

## 设计要点(对照 Java/Go/Python)

| 点 | 本实现 | Java 对应 | Go/Python 对应 |
|---|---|---|---|
| 并发模型 | **单线程 + 异步**(无锁) | 线程 + synchronized | Mutex / threading.Lock |
| 幂等 | paid 标志(天然原子) | 状态机 + 锁 | 同左 |
| HTTP | Express | Spring Boot | Gin / FastAPI |
| 测试 | Jest + supertest | MockMvc + JUnit | httptest / pytest |
| 校验 | 手动(生产用 zod/schema) | Bean Validation | pydantic / validator |
| 部署 | node:22-alpine 镜像 + pm2 多实例 | JRE 镜像 | scratch / python 镜像 |

## 关键差异(学习价值)

1. **并发扣减无需锁**:Node 单线程,`payOrder` 的 JS 语句天然原子——这是与 Go/Python/Java 最大的对照点。
2. **多实例必须外部化状态**:单进程内存 store 只在单实例演示;pm2 -i max 或多副本部署后,库存/幂等标志必须放 Redis/DB(见 `08-部署运维.md`)。
3. **错误处理**:Express 4 的 async handler 抛错需包装(本项目 handler 为同步,天然安全);Fastify/NestJS 原生支持 async。

## 目录结构

```
nodejs-order-service/
├── package.json      # 依赖与脚本(express/jest/supertest)
├── package-lock.json # 版本锁定(必须提交)
├── server.js         # 入口(独立于 app,便于测试)
├── app.js            # Express 应用与路由
├── store.js          # 内存存储(单线程原子 + 幂等标志)
├── app.test.js       # Jest + supertest 测试
├── Dockerfile        # node:22-alpine 镜像
└── README.md
```
