# PHP · Java 开发者迁移指南(入口)

> 锚点:PHP 在 Java 生态中的对应定位——「请求级生命周期 + 多进程模型的 Web 服务端,存量系统占比高」。
>
> 验证环境:PHP 8.4.24(本地 winget)/ CI: Ubuntu 24.04 + PHP 8.4。

## 为什么 Java 开发者要了解 PHP

- **存量系统现实**:大量 Web 系统(内容站/传统业务)是 PHP 写的,迁移/维护/接手的必备知识。
- **生命周期模型对照价值最大**:PHP 的「请求级进程」与 Java 的「常驻 JVM」是两种极端——理解后对「应用状态管理」的认知更完整。
- **低成本部署**:无需编译、无需守护进程(FPM 管理),LNMP 一键起服务。

## 工具链

| 项 | Java | PHP |
|---|---|---|
| 运行时 | JVM | Zend 引擎(PHP-FPM 进程池) |
| 版本管理 | jenv/sdkman | phpenv / 系统包管理 |
| 包管理 | Maven/Gradle | **composer** |
| 依赖目录 | .m2(全局) | **vendor/(随项目)** |
| IDE | IDEA | PhpStorm / VS Code |
| 格式化 | checkstyle | PHP-CS-Fixer / Pint |
| 静态检查 | spotbugs | PHPStan / Psalm(强推) |
| 测试 | JUnit | PHPUnit |

## 常用命令速查

```bash
# 初始化项目(生成 composer.json,≈ pom.xml)
composer init

# 装依赖(≈ 往 pom.xml 加依赖)
composer require slim/slim:"4.*"
composer require --dev phpunit/phpunit

# 按 lock 精确安装(CI/生产)
composer install --no-dev

# 运行内置服务器(开发,≈ mvn spring-boot:run)
php -S localhost:8080

# 测试
vendor/bin/phpunit
```

## 与 Java 生态对应表

| 概念 | Java | PHP |
|---|---|---|
| Web 框架 | Spring Boot | **Laravel**(全家桶)/ Slim(微)/ Hyperf(协程) |
| ORM | MyBatis / JPA | Eloquent(Laravel)/ Doctrine |
| 测试 | JUnit + Mockito | **PHPUnit** |
| 日志 | SLF4J + Logback | Monolog |
| 配置 | application.yml | .env + config/ 目录 |
| 依赖注入 | Spring 容器 | Laravel 容器 / PHP-DI |
| 模板 | Thymeleaf | Blade(Laravel) |
| 进程模型 | 常驻 JVM 多线程 | **FPM 多进程,请求级生命周期** |
| 并发 | 线程 + 锁 | 进程隔离(无共享);Swoole 协程 |
| 部署 | JAR + JRE 镜像 | PHP-FPM + nginx |

## 学习路径

1. `01-语言特性`(生命周期模型!)+ `02-基础语法`(1-2 天)
2. `03-数据结构`:array 三合一(1 天)
3. `04-并发编程`:FPM 进程模型 vs Java 线程(1 天,概念为主)
4. `05-IO与网络` + `06-测试`(1 天)
5. `07-生产框架`(Laravel/Slim)+ `08-部署运维`(FPM 调优,2 天)
6. `10-常见误区` 通读 → 对照 `examples/php-order-service/`(1-2 天)

## 模块导航

| 模块 | 内容 |
|---|---|
| 01-语言特性 | 请求级生命周期/弱类型/解释执行/PHP8 新特性 |
| 02-基础语法 | 变量$ / 控制流 / 函数 / 类 / 命名空间 / 异常 |
| 03-数据结构 | array 三合一(list/map/set) |
| 04-并发编程 | FPM 进程模型/Swoole 协程/无共享内存 |
| 05-IO与网络 | 文件/JSON/HTTP(fopen/curl) |
| 06-测试 | PHPUnit/数据提供者 |
| 07-生产框架 | Laravel/Slim/Hyperf/Eloquent |
| 08-部署运维 | FPM+nginx/进程池调优/容器化 |
| 09-监控 | FPM 状态页/日志/指标 |
| 10-常见误区 | Java 思维移植的坑 |
| 11-学习路径 | 7 天上手路线 + 书单 |

## 权威资源

- 官方文档:php.net / php.net/manual/zh
- 《PHP 与 MySQL Web 开发》(经典入门)
- Laravel 官方文档:laravel.com/docs
