# GLOSSARY · 术语对照表

> 验证环境:纯文档,无依赖。

用途:跨语言沟通、读别的语言文档时遇到术语卡壳、把 Java 知识「翻译」到目标语言。

## 核心概念对照

| 概念 | Java | Go | Python | Node.js | PHP |
|---|---|---|---|---|---|
| 包管理工具 | Maven / Gradle | go mod | pip / poetry | npm / pnpm | composer |
| 构建工具 | Maven / Gradle | go build | setuptools / hatch | (打包工具链) | (无) |
| 集合/容器 | Collection 框架 | slice + map | list / dict / set | Array / Map / Set | array(兼具 list/map) |
| 线程 | Thread | goroutine | threading.Thread | worker_threads | 进程(传统)/ Fiber |
| 锁 | synchronized / Lock | sync.Mutex | threading.Lock | (单线程,少用) | (进程隔离,少用) |
| 并发容器 | ConcurrentHashMap | sync.Map | queue.Queue | (单线程) | (进程隔离) |
| 线程池 | ThreadPoolExecutor | goroutine(无池概念) | ThreadPoolExecutor | worker pool(手动) | (无) |
| 异步编排 | CompletableFuture | channel + select | asyncio.Task | Promise / async-await | Swoole 协程 |
| 回调地狱解法 | (CompletableFuture) | goroutine 顺序写 | async/await | async/await | async/await(PHP8) |
| 接口 | interface | interface | ABC / Protocol | interface(TS) | interface(PHP8) |
| 继承 | extends | 组合(无继承) | class 继承 | class 继承(原型) | class 继承 |
| 泛型 | 泛型(类型擦除) | 泛型(1.18+) | 泛型(3.12+) | 泛型(TS) | 泛型(无) |
| 异常 | try/catch/throw | panic/recover | try/except/raise | try/catch/throw | try/catch/throw |
| 空值 | null / Optional | nil(无 Optional) | None | null / undefined | null |
| 日志 | SLF4J + Logback | log / slog | logging | pino / winston | Monolog |
| 配置 | application.yml + @Value | viper / env | pydantic-settings | dotenv / env | .env |
| ORM | MyBatis / JPA | GORM | SQLAlchemy | Prisma / TypeORM | Eloquent / Doctrine |
| 连接池 | HikariCP | database/sql 池 | SQLAlchemy 池 | 连接池库 | PDO 持久连接 |
| 测试框架 | JUnit | go test | pytest | Jest / Vitest | PHPUnit |
| Mock | Mockito | (接口 + 手动 mock) | unittest.mock | jest.mock / vi.mock | PHPUnit mock |
| 依赖注入 | Spring 容器 | (手动 / wire) | (手动 / dependency-injector) | NestJS DI / tsyringe | Laravel 容器 / PHP-DI |
| 消息队列客户端 | Kafka / RocketMQ 客户端 | 官方/社区客户端 | 各客户端 | 各客户端 | 各客户端 |
| 内存缓存 | Caffeine / Redis 客户端 | 内存 + Redis | 内存 + Redis | 内存 + Redis | 内存(进程内)+ Redis |
| 部署单元 | JAR/WAR | 静态二进制 | 源码/whl | 源码/打包 | 源码 |
| 容器镜像 | JRE + JAR | scratch/distroless + 二进制 | python 镜像 + 依赖 | node 镜像 | php-fpm 镜像 |
| 进程守护 | systemd / supervisor | 容器自启 | gunicorn / supervisor | pm2 / systemd | php-fpm |
| 监控指标 | Micrometer + Prometheus | prometheus/client_golang | prometheus_client | prom-client | (第三方) |
| 链路追踪 | SkyWalking / Zipkin | OpenTelemetry | OpenTelemetry | OpenTelemetry | OpenTelemetry |
| 性能剖析 | JFR / jstack / MAT | pprof | cProfile / py-spy | --prof-process / clinic | Xdebug / Blackfire |

## 常见缩写对照

| 缩写 | Java 语境 | 其他语言语境 |
|---|---|---|
| JVM | Java 虚拟机 | Go:无(单二进制);Node:V8;Python:解释器 |
| GC | 分代垃圾回收 | Go:并发 GC;Python:引用计数为主;PHP:请求结束即释放 |
| DI | Spring 依赖注入 | 各框架容器(见上表) |
| CI/CD | GitHub Actions 等 | 通用,语言无关 |
| API | 接口 | 通用,语言无关 |
| SDK | 开发工具包 | 通用 |

## 使用方式

- 查单个术语:直接搜本表。
- 学新语言前:通读「核心概念对照」两遍,建立对应关系索引。
- 发现缺失术语:欢迎补充(见 CONTRIBUTING.md)。
