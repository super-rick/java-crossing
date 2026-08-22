# examples · 同一业务契约的多语言实现

> 验证环境:各子目录 README 标注的版本(Go 1.26 / Python 3.11 / Node 22 / PHP 8.4),CI 矩阵自动编译测试。

**核心价值:同一份订单服务契约,四种语言实现,直接 diff 学习「同样的需求,不同的语言心智」。**

## 业务契约

订单服务(贴近电商/支付领域的真实形态):

| 方法 | 路径 | 说明 | 成功 | 失败 |
|---|---|---|---|---|
| POST | /orders | 创建订单 `{"item":"coffee","quantity":2}` | 201 order | 400 参数错 / 409 库存不足 |
| GET | /orders/{id} | 查询订单 | 200 order | 404 |
| POST | /orders/{id}/pay | **幂等支付(并发安全)** | 200 {"paid":true} | 404 / 409 库存不足 |
| GET | /healthz | 健康检查 | 200 | — |
| GET | /metrics | Prometheus 指标 | 200 text | — |

**核心难点**:`POST /orders/{id}/pay` 的并发幂等——100 个并发支付请求,库存只能扣一次,其余全部幂等成功。

## 四种实现对比

| 维度 | Go | Python | Node.js | PHP |
|---|---|---|---|---|
| 目录 | go-order-service | python-order-service | nodejs-order-service | php-order-service |
| HTTP 框架 | net/http(标准库) | FastAPI | Express | Slim 4 |
| 数据校验 | 手动 | **pydantic(类型注解自动)** | 手动(生产用 zod) | 手动 |
| 并发模型 | goroutine + Mutex | threading.Lock(GIL) | **单线程(无锁原子)** | 请求级进程(无共享) |
| 幂等方案 | Mutex 临界区 + paid 标志 | with lock + paid 标志 | paid 标志(天然原子) | paid 标志(进程内) |
| 测试 | go test -race + httptest | pytest + TestClient | Jest + supertest | PHPUnit |
| 并发测试 | 100 goroutine | 100 线程 | 100 并发请求 | 100 次重复调用 |
| 部署形态 | scratch 静态二进制 | uvicorn/gunicorn | node 镜像 + pm2 | php-fpm + nginx |
| 关键教学点 | 标准库即生产级 | 注解=校验+文档 | 单线程无锁并发 | **请求级生命周期(跨请求无状态)** |

## 语言差异要点(学习顺序建议)

1. **Go**:并发扣减 = `Mutex.Lock` + defer 解锁——最直接的「锁」教学。
2. **Python**:`with lock:` 自动释放,语义与 Go 相同但语法更简洁;注意 GIL 场景选型。
3. **Node.js**:单线程下**无需锁**,`paid` 标志天然原子——理解「事件循环」后看这里会心一笑。
4. **PHP**:每个请求独立进程,Store 每次重建——**跨请求查询订单返回 404** 是刻意保留的真实行为,生产必须用数据库/Redis 外部化(见其 README)。

## 统一运行方式

```bash
# 每个子目录内:
# Go
go run .                    # :8080,测试: go test -race ./...
# Python
uvicorn app:app --port 8080 # 测试: pytest
# Node
npm ci && npm start         # 测试: npm test
# PHP
composer install && php -S localhost:8080 -t public public/index.php  # 测试: vendor/bin/phpunit
```

CI 矩阵(`.github/workflows/examples-test.yml`)自动编译 + 测试四种实现,任何改动不绿即拒绝合并。

## 新增语言指引

见 `docs/LANGUAGE_TEMPLATE.md` 与根目录 `CONTRIBUTING.md`:复制任一实现目录,保持契约不变,替换框架与测试即可。
