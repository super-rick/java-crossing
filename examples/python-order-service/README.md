# python-order-service(Python 实现)

> 锚点:用 FastAPI 实现与 Go/Java 相同的订单服务契约,展示「类型注解 = 校验 + 文档」的 Python 生产形态。
>
> 验证环境:Python 3.11.15 本地实测 / CI: Ubuntu 24.04 + Python 3.13。

## 接口契约(与其他语言实现一致)

| 方法 | 路径 | 说明 | 成功 | 失败 |
|---|---|---|---|---|
| POST | /orders | 创建订单 `{"item":"coffee","quantity":2}` | 201 order | 422 校验失败 / 409 库存不足 |
| GET | /orders/{id} | 查询订单 | 200 order | 404 |
| POST | /orders/{id}/pay | **幂等支付(并发安全)** | 200 {"paid":true} | 404 / 409 |
| GET | /healthz | 健康检查 | 200 | — |
| GET | /metrics | Prometheus 文本指标 | 200 | — |

## 运行

```bash
pip install -r requirements.txt        # 或 uv sync
uvicorn app:app --port 8080            # 开发
gunicorn app:app -w 4 -k uvicorn.workers.UvicornWorker -b 0.0.0.0:8080   # 生产
```

## 测试

```bash
pytest -v          # 5 个测试(含 100 线程并发幂等)
```

核心测试 `test_concurrent_pay_idempotent`:100 个线程并发支付同一订单,断言**库存只扣减一次**,全部幂等成功。

## 设计要点(对照 Java/Go)

| 点 | 本实现 | Java 对应 | Go 对应 |
|---|---|---|---|
| 存储 | 内存 dict + threading.Lock | ConcurrentHashMap | sync.Mutex + map |
| 并发扣减 | `with lock:` 临界区 | synchronized | Mutex.Lock |
| 参数校验 | pydantic(类型注解) | Bean Validation | validator 库 |
| OpenAPI | 自动生成(/docs) | springdoc | swaggo |
| HTTP | FastAPI | Spring Boot | net/http / Gin |
| 测试 | pytest + TestClient | MockMvc + JUnit | httptest |
| 部署 | uvicorn/gunicorn + 解释器镜像 | JRE 镜像 | scratch 二进制 |

## 目录结构

```
python-order-service/
├── app.py             # FastAPI 入口(路由/校验/异常映射)
├── store.py           # 内存存储 + threading.Lock
├── test_app.py        # API + 并发幂等测试
├── requirements.txt   # 依赖(fastapi/uvicorn/pytest/httpx)
├── Dockerfile         # python:3.13-slim 镜像
└── README.md
```

## 与 Go 版对比要点(学习价值)

- 同样的「Mutex/Lock + 幂等标志」并发方案,Python 用 `with lock:` 自动释放(Go 要 `defer mu.Unlock()`)。
- pydantic 让校验+文档零成本,这是 Python 相对 Go/Java 的生产效率差异点。
- 进程模型不同:Go 单进程内 goroutine;Python 多 worker 进程(uvicorn/gunicorn)。
