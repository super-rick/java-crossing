# php-order-service(PHP 实现)

> 锚点:用 Slim 4 实现与 Go/Python/Node 相同的订单服务契约,展示 PHP「请求级生命周期」的独特形态。
>
> 验证环境:PHP 8.4.24 本地实测 / CI: Ubuntu 24.04 + PHP 8.4 + composer。

## 接口契约(与其他语言实现一致)

| 方法 | 路径 | 说明 | 成功 | 失败 |
|---|---|---|---|---|
| POST | /orders | 创建订单 `{"item":"coffee","quantity":2}` | 201 order | 400 参数错 / 409 库存不足 |
| GET | /orders/{id} | 查询订单 | 200 order | 404 |
| POST | /orders/{id}/pay | 幂等支付 | 200 {"paid":true} | 404 / 409 |
| GET | /healthz | 健康检查 | 200 | — |
| GET | /metrics | 指标文本 | 200 | — |

## 运行

```bash
composer install          # 安装依赖(生成 vendor/ 与 composer.lock)
php -S localhost:8080 -t public public/index.php   # 开发服务器
vendor/bin/phpunit        # 5 个测试
```

## ⚠️ 生产说明(本示例最核心的教学点)

**传统 PHP 是请求级生命周期:每个请求一个独立进程,Store 每次请求重建——跨请求的状态(订单/库存)不保留。**

实测行为(黑盒验证):

```
POST /orders          → 201 {"id":1,...}      # 本请求内有效
GET  /orders/1        → 404 order not found   # 新请求 = 新进程,状态已销毁!
```

**因此生产环境的正确形态:**
1. 订单/库存/幂等标志 → 存数据库(用事务 + 唯一约束/行锁实现并发扣减,如 `UPDATE inventory SET stock = stock - ? WHERE item = ? AND stock >= ?`)
2. 缓存/会话 → Redis
3. 高并发场景 → Swoole/Hyperf 协程常驻(进程内才有「并发编程」)

本示例的内存 Store 仅用于演示「单请求内的业务逻辑 + 幂等语义」(单元测试覆盖);跨请求持久化见 08-部署运维与 07-生产框架模块。

## 设计要点(对照其他语言)

| 点 | 本实现 | Java/Go/Python/Node 对应 |
|---|---|---|
| 生命周期 | **请求级进程(无共享)** | 常驻应用(线程/goroutine/协程) |
| 并发扣减 | 进程内无并发(架构规避);生产靠 DB 原子操作 | 锁/原子操作(语言级) |
| HTTP | Slim 4(PSR-7 标准) | Spring Boot / Gin / FastAPI / Express |
| JSON | json_encode/decode(标准库) | Jackson / encoding/json / pydantic |
| 测试 | PHPUnit | JUnit / go test / pytest / Jest |
| 依赖 | composer(vendor/) | Maven / go mod / pip / npm |

**踩坑记录(真实教学素材):**
- PSR-4 自动加载**不加载函数** → 函数文件要在 composer.json `autoload.files` 注册。
- Slim 4 的 `withJson()`/`write()` 是框架扩展方法,slim/psr7 标准实现没有 → 用 PSR-7 标准写法(`getBody()->write` + `withHeader`)。
- 应用函数需配 PSR-17 实现包(slim/psr7),否则 `AppFactory::create()` 报错。

## 目录结构

```
php-order-service/
├── composer.json      # slim/slim + phpunit + autoload(files 注册函数)
├── composer.lock      # 版本锁定(必须提交)
├── public/index.php   # 前端控制器入口
├── src/Store.php      # 内存存储(单请求内)
├── src/app.php        # createApp() 工厂 + 路由(PSR-7 写法)
├── tests/StoreTest.php# PHPUnit 5 个测试
├── Dockerfile         # php:8.4-fpm + composer
└── README.md
```
