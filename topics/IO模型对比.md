# IO 模型对比

> 验证环境:纯文档,无依赖;版本基线 Go 1.24 / Python 3.13 / Node 22 / PHP 8.4。

**一句话:** Java 从 BIO 走到 NIO(Selector)再到 Netty(Reactor);Go 把 epoll 封装成「阻塞式写法」;Node 把异步非阻塞作为语言默认;Python 的 asyncio 与 Java NIO 同源(事件循环);PHP 传统阻塞,FPM 多进程扛并发。

## IO 模型总览

| 模型 | Java | Go | Python | Node.js | PHP |
|---|---|---|---|---|---|
| 同步阻塞(BIO) | java.io(老代码) | —(默认不暴露) | 默认文件/网络 IO | —(默认异步) | 传统默认 |
| 多路复用封装 | NIO Selector / Netty | 运行时内置(epoll/kqueue) | asyncio 事件循环 | libuv 事件循环 | Swoole 事件循环 |
| 异步文件 IO | 响应式库 / CompletableFuture | os 包(线程池实现) | asyncio.to_thread | fs(线程池) | Swoole 协程 + 线程 |
| 用户感知写法 | 回调/阻塞 → 响应式 | **同步写法,底层异步** | async/await | async/await 或回调 | 同步(FPM) / async(Swoole) |

## 关键差异

### 1. Go:给你「同步的皮,异步的里」
- 开发者写 `conn.Read(buf)` 是阻塞式代码,底层由运行时用多路复用 + goroutine 调度实现高并发。
- 心智转换:Java 里纠结「用 BIO 还是 NIO、要不要 Netty」,Go 里不用选——标准库 `net/http` 已经是最优实践。

### 2. Node.js:IO 默认异步
- 事件循环 + libuv 线程池:`fs` 文件操作、DNS、部分加密等走线程池;网络 IO 走系统多路复用。
- 心智转换:Java 的 `Files.readAllBytes()` 同步写法,Node 里必须 `await fs.promises.readFile()`;同步版本 `readFileSync` 只允许在启动阶段用。

### 3. Python:asyncio 与 Java NIO 同构
- 事件循环 + 协程,和 Netty 的 Reactor 模式同源——理解 Java NIO 的人学 asyncio 很快。
- 区别:Python 的同步库(requests、pandas 读文件)是阻塞的,混用会卡事件循环,要用 `asyncio.to_thread` 或纯异步库(httpx、aiofiles)。

### 4. PHP:阻塞是默认,FPM 进程池是并发手段
- 一个 FPM 进程处理一个请求,阻塞 IO 时进程挂着——并发靠「进程数量」而不是「IO 非阻塞」。
- 连接数上限 = 进程数 × 每进程连接数;调优思路完全不同(JVM 调线程池 → PHP 调 FPM 进程数)。
- Swoole/Hyperf:协程化后 IO 模型才接近 Go/Node。

### 5. 文件 IO vs 网络 IO
| 维度 | Java | Go | Python | Node | PHP |
|---|---|---|---|---|---|
| 网络 IO 默认 | NIO(Netty 生态) | 非阻塞多路复用 | asyncio 或线程 | libuv 非阻塞 | 阻塞(FPM 多进程) |
| 文件 IO | 阻塞为主(响应式可异步) | 内部线程池异步 | 同步;asyncio.to_thread | fs 线程池异步 | 同步 |
| 大文件流式 | InputStream 包装 | io.Reader/Writer | 文件迭代器 | stream.Readable | fopen 流 |

## 选型建议

- 需要极致网络吞吐:Go / Node / Java+Netty 都在同一梯队(多路复用),选型看团队与生态。
- 简单 CRUD + 阻塞 IO 可接受:PHP / 传统 Java BIO 足够,别过度设计。
- 混合负载(文件 + 网络 + CPU):Java(线程模型最全能)或 Go;Node 需小心 CPU 密集阻塞。

## 常见误区

- ❌ 认为「非阻塞 = 异步写法」:Go 非阻塞但写法同步;Java BIO 阻塞但可配线程池达到并发。
- ❌ Node 里用同步文件 API 处理请求内 IO:阻塞事件循环,拖垮所有并发请求。
- ❌ Python asyncio 里调同步 requests:事件循环被阻塞,协程全部卡住。
- ❌ 按 Java 思路调 PHP 的「线程池」:PHP 没有,调的是 FPM 的 `pm.max_children`。
