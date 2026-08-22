# Node.js · Java 开发者迁移指南(入口)

> 锚点:Node.js 在 Java 生态中的对应定位——「单线程事件循环 + 异步优先的 JS 后端运行时,适合 BFF/实时交互」。
>
> 验证环境:Node 22.23.0 / npm 10.9.8(本地与 CI 均为 Node 22 LTS 线)。

## 为什么 Java 开发者要学 Node.js

- **BFF(Backend for Frontend)与实时交互的首选**:前后端同语言、SSE/WebSocket 天然适配。
- **异步心智模型的必修课**:单线程事件循环是 Java 线程模型的最佳对照样本——理解后对「并发」认知上一个台阶。
- **生态**:npm 是最大软件仓库(300 万+包),全栈(React/Vue 生态)共用。

## 工具链

| 项 | Java | Node.js |
|---|---|---|
| 运行时 | JVM | V8 引擎(Node.js) |
| 版本管理 | jenv/sdkman | nvm / fnm |
| 包管理 | Maven/Gradle | **npm / pnpm / yarn** |
| 模块体系 | Maven 坐标 | package.json + node_modules |
| IDE | IDEA | VS Code(JS/TS 支持最佳) |
| 格式化 | checkstyle/spotless | Prettier / ESLint |
| 类型 | 编译期强类型 | TS(TypeScript)可选 |
| 测试 | JUnit | Jest / Vitest / node:test |

## 常用命令速查

```bash
# 初始化项目(生成 package.json,≈ pom.xml)
npm init -y

# 装依赖(≈ 往 pom.xml 加依赖)
npm install express          # 生产依赖
npm install -D jest          # 开发依赖(测试/构建工具)

# 按 lock 精确安装(CI 必用,≈ 锁定版本)
npm ci

# 运行脚本(≈ mvn 的插件目标)
npm start / npm test / npm run build

# 全局装工具
npm install -g pnpm
```

## 与 Java 生态对应表

| 概念 | Java | Node.js |
|---|---|---|
| Web 框架 | Spring Boot | **Express**(极简)/ Fastify(高性能)/ NestJS(结构化) |
| ORM | MyBatis / JPA | Prisma / TypeORM / Sequelize |
| 测试 | JUnit + Mockito | **Jest**(全家桶)/ Vitest / node:test |
| 日志 | SLF4J + Logback | pino / winston |
| 配置 | application.yml | dotenv / 环境变量 |
| 监控 | Actuator/Micrometer | prom-client |
| 线程 | Thread / 线程池 | worker_threads(特殊场景) |
| 异步 | CompletableFuture | **Promise / async-await(语言原生)** |
| 集合 | Collection 框架 | Array / Map / Set |
| JSON | Jackson | JSON 原生(语言内建) |
| 构建 | Maven/Gradle | 无编译(TS 项目用 tsc/bundler) |

## 学习路径

1. `01-语言特性` + `02-基础语法`:JS 动态类型、异步语法(1-2 天)
2. `03-数据结构`:Array/Map/Set 与 Java 集合对应(1 天)
3. `04-并发编程`:**Event Loop 心智模型**(2 天,最关键——这是 Node 与 Java 最大差异)
4. `05-IO与网络` + `06-测试`(1 天)
5. `07-生产框架`(Express/Fastify/NestJS)+ `08-部署运维`(2 天)
6. `10-常见误区` 通读 → 对照 `examples/nodejs-order-service/`(1-2 天)

## 模块导航

| 模块 | 内容 |
|---|---|
| 01-语言特性 | 单线程 Event Loop/V8/动态类型/ESM 与 CJS |
| 02-基础语法 | let-const/箭头/解构/async-await/类/模块 |
| 03-数据结构 | Array/Map/Set/Object 与 Java 集合对应 |
| 04-并发编程 | Event Loop 详解/worker_threads/异步并发 |
| 05-IO与网络 | fs/stream/http/JSON |
| 06-测试 | Jest/Vitest/node:test、异步测试 |
| 07-生产框架 | Express/Fastify/NestJS、Prisma |
| 08-部署运维 | node 镜像/pm2/cluster/多实例 |
| 09-监控 | prom-client/pino/--trace-gc |
| 10-常见误区 | Java 思维移植的坑 |
| 11-学习路径 | 7 天上手路线 + 书单 |

## 权威资源

- 官方文档:nodejs.org/docs/latest/api
- 《深入浅出 Node.js》(朴灵,中文经典)
- JavaScript 基础:developer.mozilla.org/zh-CN/docs/Web/JavaScript
