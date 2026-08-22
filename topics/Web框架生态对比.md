# Web 框架生态对比

> 验证环境:纯文档,无依赖;框架版本以官方当前稳定版为准(2026-08)。

**一句话:** Spring Boot 是「全家桶 + 注解 + 依赖注入」范式;Go 生态是「小而美组合」;Python 分「全栈(Django)vs API 优先(FastAPI)」;Node 分「极简(Express)vs 结构化(NestJS)」;PHP 是「Laravel 全家桶 vs 协程 Hyperf」。

## 框架对照表

| 维度 | Java | Go | Python | Node.js | PHP |
|---|---|---|---|---|---|
| 全栈代表 | Spring Boot | (无全家桶,组合) | Django | NestJS | Laravel |
| 轻量/API | Spring WebFlux / JAX-RS | Gin / Echo / Fiber | FastAPI / Flask | Express / Fastify | Slim / Lumen |
| 高性能 | Netty 系 | Fiber / Gin | FastAPI(异步) | Fastify | Hyperf(Swoole) |
| 依赖注入 | Spring 容器(注解) | 手动 / wire / fx | 手动 / FastAPI Depends | NestJS DI / tsyringe | Laravel 容器 / PHP-DI |
| 路由 | @RequestMapping 注解 | 方法注册 | 装饰器 / 声明式 | 链式注册 | 路由文件 |
| 中间件 | Filter / Interceptor | 中间件链 | 中间件(ASGI/WSGI) | 中间件链 | 中间件 |
| 参数校验 | Bean Validation | validator 库 | pydantic(自动) | zod / class-validator | 手动 / Laravel Validation |
| OpenAPI 文档 | springdoc | swaggo | FastAPI 自动生成 | @nestjs/swagger | swagger-php |
| 配置管理 | application.yml + 环境 | viper / env | pydantic-settings | dotenv + config 库 | .env + config/ |
| 内置 Web 服务器 | Tomcat(内嵌) | net/http(标准库) | uvicorn / gunicorn | Node 原生 + 框架 | PHP-FPM(或 Swoole) |
| 迁移成本(Java 视角) | —(已知) | 低 | 中 | 中 | 低 |

## ORM 对照表

| 维度 | Java | Go | Python | Node.js | PHP |
|---|---|---|---|---|---|
| 主流 ORM | MyBatis / JPA | GORM / sqlx | SQLAlchemy | Prisma / TypeORM | Eloquent / Doctrine |
| 原生 SQL 派 | MyBatis(手写 SQL) | sqlx / sqlc | SQLAlchemy Core | knex | Doctrine DBAL |
| 自动迁移 | Flyway / Liquibase | golang-migrate | Alembic | Prisma Migrate | Laravel Migrations |
| 连接池 | HikariCP | database/sql 内置池 | SQLAlchemy 池 | 各驱动自带 | PDO 持久连接 |

## 关键差异

### 1. Go:没有「全家桶」,组合是常态
- Spring Boot 一个依赖搞定 Web+DB+配置+监控;Go 是「选型自由」:Gin(Web)+ GORM(ORM)+ viper(配置)+ prometheus(监控)自己拼。
- 心智转换:从「框架替你决定」到「你自己决定」——好处是依赖少、二进制小;代价是选型成本。
- 好消息:组合后的惯用栈高度趋同(见 README 语言覆盖表),照抄主流组合即可。

### 2. Python:FastAPI 的自动校验与文档
- FastAPI 基于类型注解:pydantic 自动完成参数校验 + 自动生成 OpenAPI 文档——Java 里 springdoc + Bean Validation 的组合,它开箱即得。
- Django 是「全栈魔法」:Admin 后台、ORM、模板内置,适合内容型/管理型系统;API 优先选 FastAPI。

### 3. Node.js:NestJS 是「TypeScript 版 Spring」
- 模块化 + 依赖注入 + 装饰器,Java 开发者几乎零学习成本。
- 对比:Express 极简自由(类似 Go 组合思路);NestJS 结构化(类似 Spring Boot)。
- Fastify 性能优于 Express,新项目常用 Fastify 或 NestJS+Fastify 适配器。

### 4. PHP:Laravel 与 Hyperf
- Laravel:全家桶 + 优雅语法,传统 PHP 首选;PHP 8 后性能提升明显。
- Hyperf:Swoole 协程常驻,追求高并发时用,心智接近 Go/Java。
- 心智转换:Java 常驻服务 → PHP 传统请求级,框架本身差异反而小。

## 选型建议

| 场景 | 推荐 |
|---|---|
| 微服务/REST API | Spring Boot / NestJS / FastAPI / Gin 同梯队,按团队选 |
| 高吞吐网关/中间件 | Go(Gin/Fiber)+ Java(Netty) |
| 内容型/后台管理 | Django / Laravel 全家桶最省事 |
| 快速原型 | FastAPI / Express |
| 协程高并发 PHP | Hyperf(需 Swoole) |

## 常见误区

- ❌ 在 Go 里找 Spring Boot 式全家桶:组合栈是 Go 文化,找「全家桶」会一直找不到。
- ❌ 以为 FastAPI 只能异步:它同时支持同步接口(自动丢线程池),迁移平滑。
- ❌ Node 新项目无脑 Express:Fastify 性能更好;结构化需求选 NestJS。
- ❌ PHP 新项目用原生 + 手写路由:Laravel/Slim 的组件式引入即可,不必全家桶。
