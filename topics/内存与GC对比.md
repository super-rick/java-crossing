# 内存与 GC 对比

> 验证环境:纯文档,无依赖;版本基线 JDK 21 / Go 1.24 / Python 3.13 / Node 22 / PHP 8.4。

**一句话:** Java 的 GC 是「可调优的分代回收器家族」(G1/ZGC),Go 追求「低延迟全自动」,Python 是「引用计数 + 循环回收」,V8 是「分代 + 并行标记」,传统 PHP 干脆不 GC——请求结束整块内存回收。

## 内存管理模型对比

| 维度 | Java | Go | Python | Node.js(V8) | PHP |
|---|---|---|---|---|---|
| 管理方式 | 分代 GC | 并发标记清除 GC | 引用计数 + 分代循环 GC | 分代 GC(新生代/老生代) | 引用计数(请求级) |
| 回收器 | G1 / ZGC / Shenandoah | 运行时内置(并发) | CPython 自带 | Scavenger + Mark-Compact | Zend 引擎 |
| 调优入口 | 大量 JVM 参数(-Xmx、-XX:G1…) | GOGC / GOMEMLIMIT | sys.setswitchinterval 等(有限) | --max-old-space-size | memory_limit(请求上限) |
| 暂停时间 | G1 可预期;ZGC 亚毫秒 | 极短(并发,1ms 级) | 引用计数即时;循环回收有暂停 | 老生代标记暂停(可观察) | 无 GC 暂停(请求结束释放) |
| 内存上限 | 堆(可设)+ 堆外 | 进程内存(自动) | 进程内存 | 堆(可设上限) | 单请求上限 |
| 泄漏风险点 | 未关闭资源/静态集合 | 全局引用/gotcha 闭包 | 全局缓存/循环引用+__del__ | 闭包引用/未清理监听器 | 传统模式无(请求级);Swoole 常驻有 |

## 关键差异

### 1. Java 的「调优文化」在别处不适用
- JVM 调优(-Xms/-Xmx/GC 选择/GC 日志)是一整套工程实践。
- Go:默认参数 + `GOMEMLIMIT`(1.19+)基本够用,官方建议「不要手动调 GC」。
- Node:`--max-old-space-size` 设堆上限防止 OOM,其余交给 V8。
- Python:基本没有 GC 调优概念,问题通常靠「换数据结构/减少对象分配」解决。

> ⚠️ Java 思维陷阱:把 JVM 的 -Xmx 思维带到 Go —— Go 内存由运行时管理,设 GOMEMLIMIT 即可,瞎调 GOGC 反而增加延迟。

### 2. Python 引用计数 vs 分代 GC
- 大部分对象**即时回收**(引用计数),所以 Python 内存释放是确定的。
- 循环引用(A→B→A)由分代回收器兜底——延迟回收,可能积压。
- 迁移注意:Java 程序员习惯「GC 帮我兜底一切」;Python 里长期持有的引用(全局 dict、缓存)就是泄漏。

### 3. PHP 的请求级内存(传统模式)
- 每个请求独立进程,请求结束 → 进程回收 → 内存清零 → 天然无泄漏、无 GC 调优。
- 代价:跨请求无法复用(连接、缓存、JIT 编译结果在 Swoole 之前都是每请求重建)。
- Swoole/Hyperf 常驻后,Java 的泄漏经验才适用(连接池、静态变量、闭包引用)。

### 4. Node.js 的闭包与监听器泄漏
- V8 的 GC 与 Java 类似,但 JS 的「闭包捕获 + 未移除的事件监听器」是主要泄漏源。
- 心智转换:Java 的「静态集合持有对象」≈ Node 的「模块级闭包 + 全局监听器」。

## 调优对照表

| 目标 | Java | Go | Python | Node | PHP |
|---|---|---|---|---|---|
| 内存上限 | -Xmx + 堆外限制 | GOMEMLIMIT | 进程限制(cgroup) | --max-old-space-size | memory_limit |
| 降低暂停 | 换 GC 器(G1→ZGC) | 默认已并发 | — | 减少老生代压力(少大对象) | — |
| 定位泄漏 | MAT / JFR / jmap | pprof heap | tracemalloc / objgraph | --inspect heap snapshot | Xdebug / Swoole 分析器 |
| 排查 OOM | OOM Killer / 堆 dump | runtime/metrics / pprof | 进程 RSS 监控 | --heapsnapshot | memory_limit 触顶日志 |

## 常见误区

- ❌ 认为所有 GC 语言都能「无限分配,GC 兜底」:Go/Python/Node 长时间引用不释放同样是泄漏。
- ❌ Python 大量创建对象不释放还怪 GC 慢:先查引用是否被长期持有。
- ❌ 传统 PHP 不需要关注内存:单请求内大数组/循环拼接照样 OOM(memory_limit)。
- ❌ 以为 GC 参数越多越好:Go 官方明确「GOGC 默认值对绝大多数程序最优」。
