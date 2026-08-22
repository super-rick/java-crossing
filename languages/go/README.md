# Go · Java 开发者迁移指南(入口)

> 锚点:Go 在 Java 生态中的对应定位——「不需要 JVM 的、为云原生而生的系统语言」。
>
> 验证环境:Go 1.26.7 / Windows 11(git-bash)/ CI: Ubuntu 24.04 (go 1.24)

## 为什么 Java 开发者要学 Go

- 云原生/K8s 生态的第一语言,中间件、网关、代理几乎都是 Go 写的。
- 单静态二进制 + 毫秒启动 + 低内存:部署形态与 Java 形成最鲜明对照,学完立刻理解「Java 的 JVM 成本在哪」。
- 并发模型(goroutine/channel)与 Java 线程/锁差异清晰,是四门语言中**对照价值最大**的。

## 工具链

| 项 | Java | Go |
|---|---|---|
| SDK | JDK | Go 工具链(go.exe,含编译器/格式化/测试) |
| 版本管理 | jenv/sdkman | go1.26.7(官方安装包;多版本用 g) |
| 构建/依赖 | Maven/Gradle | go mod(内置) |
| IDE | IDEA | GoLand / VS Code(Go 插件) |
| 格式化 | spotless/checkstyle | **gofmt(内置,强制风格统一)** |
| 调试 | IDEA debug | delve(dlv) |
| 常用命令 | mvn test/package | go build / go test / go vet / go run |

## 常用命令速查

```bash
# 初始化模块(≈ mvn 的项目初始化)
go mod init example.com/myapp

# 构建(≈ mvn package,产出单个二进制)
go build -o myapp .

# 测试(≈ mvn test)
go test ./...

# 静态检查(≈ spotbugs/checkstyle)
go vet ./...

# 格式化(Go 强制,写任何代码前跑它)
gofmt -w .

# 安装第三方包(≈ 往 pom.xml 加依赖)
go get github.com/gin-gonic/gin
```

## 与 Java 生态对应表

| 概念 | Java | Go |
|---|---|---|
| 依赖管理 | pom.xml + Maven Central | go.mod + proxy.golang.org |
| Web 框架 | Spring Boot | Gin / Echo / Fiber |
| ORM | MyBatis / JPA | GORM / sqlx |
| 测试 | JUnit + Mockito | go test(内置)+ 表驱动 |
| 日志 | SLF4J + Logback | log/slog / zap |
| 配置 | application.yml | viper / env |
| 监控 | Actuator / Micrometer | prometheus/client_golang + pprof |
| 消息客户端 | Kafka/Redis 客户端 | 官方/社区客户端 |
| 线程 | Thread / 线程池 | goroutine(无需池) |
| 异步编排 | CompletableFuture | channel + select |

## 学习路径

1. 读 `01-语言特性` 与 `02-基础语法`,建立语法映射(1–2 天)
2. 精读 `03-数据结构` + `04-并发编程`(这是 Go 的灵魂,2 天)
3. `05-IO与网络`、`06-测试`(1 天)
4. `07-生产框架`、`08-部署运维`、`09-监控`(2 天,重点看「JVM 调优 → Go 对应物」)
5. `10-常见误区` 通读,避免 Java 思维翻车
6. 对照 `examples/go-order-service/` 实战(1–2 天)

## 模块导航

| 模块 | 内容 |
|---|---|
| 01-语言特性 | 范式/类型系统/编译/内存管理/接口 |
| 02-基础语法 | 变量/控制流/函数/结构体/错误处理/指针 |
| 03-数据结构 | slice/map/struct 与 Java 集合框架对应 |
| 04-并发编程 | goroutine/channel/Mutex 与线程/锁对应 |
| 05-IO与网络 | 标准库 IO、net/http、序列化 |
| 06-测试 | go test、表驱动、覆盖率 |
| 07-生产框架 | Gin/GORM/配置/日志生态 |
| 08-部署运维 | 静态二进制/镜像/调优 |
| 09-监控 | pprof/expvar/Prometheus |
| 10-常见误区 | Java 思维移植的坑 |
| 11-学习路径 | 7 天上手路线 + 书单 |

## 权威资源

- 官方文档:go.dev/doc(中文版 go.dev/doc/install 有教程)
- 《The Go Programming Language》(K&R 作者之一 Kernighan 合著)
- 官方 Wiki「Go for Java Programmers」
