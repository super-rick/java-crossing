# Python · Java 开发者迁移指南(入口)

> 锚点:Python 在 Java 生态中的对应定位——「动态类型的脚本语言,却是 AI/数据领域的统治级后端语言」。
>
> 验证环境:Python 3.11.15(本地)/ CI: Ubuntu 24.04 + Python 3.13。

## 为什么 Java 开发者要学 Python

- **AI/LLM/数据生态的入口**:机器学习、数据分析、AI 应用后端几乎全是 Python,不懂 Python 就无法进入 AI 时代的技术栈。
- **开发效率**:动态类型 + 解释执行 + 丰富的标准库,原型/工具/脚本类任务开发速度是 Java 的 2-3 倍。
- **心智对照价值**:动态类型 vs Java 强类型、GIL vs 线程、解释 vs 编译——是「类型系统与运行时」认知的极好对照样本。

## 工具链

| 项 | Java | Python |
|---|---|---|
| 运行时 | JVM + JDK | CPython 解释器(python.exe) |
| 版本管理 | jenv/sdkman | pyenv / uv python install |
| 包管理 | Maven/Gradle | **uv**(现代,推荐)/ pip |
| 依赖隔离 | .m2 全局 | **venv(每项目必建)** |
| IDE | IDEA | PyCharm / VS Code(Python 插件) |
| 格式化 | checkstyle | **ruff format**(≈ gofmt) |
| 类型检查 | 编译期 | mypy / pyright(可选) |
| 测试 | JUnit | pytest |

## 常用命令速查

```bash
# 建虚拟环境(≈ 每个项目的独立 classpath)
python -m venv .venv
# Windows: .venv\Scripts\activate   macOS/Linux: source .venv/bin/activate

# 装依赖(≈ 往 pom.xml 加依赖)
pip install requests
# 现代替代(更快):uv pip install requests

# 导出/锁定依赖(≈ 提交 pom.xml)
pip freeze > requirements.txt
# 或现代:uv lock(uv.lock,可复现)

# 运行
python main.py

# 测试
python -m pytest

# 格式化 + 检查
ruff format . && ruff check .
```

## 与 Java 生态对应表

| 概念 | Java | Python |
|---|---|---|
| Web 框架 | Spring Boot | **FastAPI**(API)/ Django(全栈) |
| ORM | MyBatis / JPA | SQLAlchemy / Django ORM |
| 测试 | JUnit + Mockito | pytest + unittest.mock |
| 日志 | SLF4J + Logback | logging(标准库) |
| 配置 | application.yml | pydantic-settings / python-dotenv |
| 监控 | Actuator/Micrometer | prometheus_client |
| JSON | Jackson | json(标准库)/ pydantic |
| 线程 | Thread / 线程池 | threading(受 GIL) |
| 异步 | CompletableFuture | asyncio |
| 集合 | Collection 框架 | list / dict / set / tuple |
| 进程 | JVM 进程 | multiprocessing |

## 学习路径

1. `01-语言特性` + `02-基础语法`:动态类型、缩进、异常(1-2 天)
2. `03-数据结构`:list/dict/set 与 Java 集合对应(1 天)
3. `04-并发编程`:**三套模型选型**(threading/asyncio/multiprocessing,2 天,最关键)
4. `05-IO与网络` + `06-测试`(1 天)
5. `07-生产框架`(FastAPI)+ `08-部署运维`(2 天)
6. `10-常见误区` 通读 → 对照 `examples/python-order-service/`(1-2 天)

## 模块导航

| 模块 | 内容 |
|---|---|
| 01-语言特性 | 动态类型/解释执行/鸭子类型/GIL/引用计数 |
| 02-基础语法 | 缩进/变量/控制流/函数/类/异常/推导式 |
| 03-数据结构 | list/dict/set/tuple 与 Java 集合对应 |
| 04-并发编程 | threading(GIL)/multiprocessing/asyncio 三套选型 |
| 05-IO与网络 | 文件/JSON/HTTP 客户端/异步 IO |
| 06-测试 | pytest/fixture/参数化/mock |
| 07-生产框架 | FastAPI/Django/SQLAlchemy/pydantic |
| 08-部署运维 | venv/uvicorn-gunicorn/解释器镜像 |
| 09-监控 | prometheus_client/logging/py-spy |
| 10-常见误区 | Java 思维移植的坑 |
| 11-学习路径 | 7 天上手路线 + 书单 |

## 权威资源

- 官方文档:docs.python.org
- 《流畅的 Python》(Fluent Python,2nd)——进阶必读
- Python 3 官方教程:docs.python.org/3/tutorial
