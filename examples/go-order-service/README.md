# go-order-service(Go 实现)

> 锚点:用 Go(标准库)实现与 Java 版相同的订单服务契约,展示「标准库即生产级」。
>
> 验证环境:Go 1.26.7 本地实测 / CI: Ubuntu 24.04 + go 1.24(`go vet && go test -race`)。

## 接口契约(与其他语言实现一致)

| 方法 | 路径 | 说明 | 成功 | 失败 |
|---|---|---|---|---|
| POST | /orders | 创建订单 `{"item":"coffee","quantity":2}` | 201 order | 400 参数错 / 409 库存不足 |
| GET | /orders/{id} | 查询订单 | 200 order | 404 |
| POST | /orders/{id}/pay | **幂等支付(并发安全)** | 200 {"paid":true} | 404 / 409 库存不足 |
| GET | /healthz | 健康检查 | 200 ok | — |
| GET | /metrics | Prometheus 指标 | 200 text | — |

## 运行

```bash
go run .                # 默认 :8080,可用 PORT 环境变量覆盖
curl http://localhost:8080/healthz
```

## 测试(含并发幂等)

```bash
go vet ./...
go test -race ./... -v   # 竞态检测 + 6 个测试
```

核心测试 `TestConcurrentPayIdempotent`:100 个 goroutine 并发支付同一订单,断言**库存只扣减一次**(95),全部调用幂等成功。

## 设计要点(对照 Java)

| 点 | 本实现 | Java 对应 |
|---|---|---|
| 存储 | 内存 map + sync.Mutex | ConcurrentHashMap / 数据库事务 |
| 并发扣减 | Mutex 临界区(原子性) | synchronized / 乐观锁 |
| 幂等 | Paid 标志 + 重复调用返回成功 | 状态机 + 唯一约束 |
| HTTP | net/http 标准库 | Spring Boot |
| 测试 | httptest + 表驱动 | MockMvc + JUnit |
| 部署 | Dockerfile → scratch 镜像(约 10MB) | JRE 镜像(200MB+) |
| 监控 | /metrics 手写 + pprof | Actuator |

## 目录结构

```
go-order-service/
├── go.mod          # 模块定义(零第三方依赖,纯标准库)
├── main.go         # HTTP 入口与路由
├── main_test.go    # API 集成测试(≈ MockMvc)
├── store.go        # 内存存储 + 并发安全
├── store_test.go   # 并发幂等/库存测试(核心)
├── Dockerfile      # 多阶段构建 → scratch
└── README.md
```
