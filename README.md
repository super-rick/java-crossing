# java-crossing · Java 开发者多语言迁移指南

以 Java 为锚点,帮助 Java 后端开发者快速上手 **Go / Python / Node.js / PHP**(二期扩展 Rust / Kotlin / C#)。

> 对比驱动 · 示例可跑 · CI 验证 —— 不写「入门教程」,只写「Java 开发者视角的对照手册」。

[![docs-check](https://github.com/super-rick/java-crossing/actions/workflows/docs-check.yml/badge.svg)](https://github.com/super-rick/java-crossing/actions/workflows/docs-check.yml)
[![examples-test](https://github.com/super-rick/java-crossing/actions/workflows/examples-test.yml/badge.svg)](https://github.com/super-rick/java-crossing/actions/workflows/examples-test.yml)

## 为什么有这个仓库

Java 开发者学新语言,障碍从来不是语法,而是**心智模型**:并发模型、内存管理、构建/依赖体系、部署形态、监控手段,每个都要重新建立对应关系。本仓库把这种对应关系系统性整理出来——每个知识点按「Java 是什么 → 本语言对应什么 → 相同/不同 → 迁移陷阱」展开。

## 仓库结构

| 目录 | 内容 | 状态 |
|---|---|---|
| `docs/` | 总览、学习路径总图、语言全景对比、每语言内容模板、术语对照 | 建设中 |
| `topics/` | 跨语言主题对比总表:并发 / IO / 内存与 GC / 包管理 / 测试 / Web 框架 / 部署 / 监控 | 建设中 |
| `languages/{go,python,nodejs,php}/` | 各语言 12 模块对照内容(语言特性/语法/数据结构/并发/IO/测试/生产框架/部署运维/监控/误区/学习路径) | Go 先行 |
| `examples/` | 同一业务契约(订单服务)的多语言实现,CI 保证可跑 | 建设中 |

## 语言覆盖

| 语言 | 内容状态 | 示例 | CI |
|---|---|---|---|
| Go | 规划中 | — | — |
| Python | 规划中 | — | — |
| Node.js | 规划中 | — | — |
| PHP | 规划中 | — | — |

## 贡献

见 [CONTRIBUTING.md](CONTRIBUTING.md)。核心要求:对照表驱动、示例必须实际运行过并注明「验证环境」。

## License

[MIT](LICENSE)
