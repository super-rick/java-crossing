# 贡献指南

## 本仓库是什么 / 不是什么

- 是:以 Java 为锚点的**对照手册**,面向 ≥1 年 Java 经验的开发者。
- 不是:单语言入门教程的堆叠。请勿提交「XX 语言基础教程」式内容,除非它同时给出与 Java 的对应关系。

## 内容要求(每个语言模块页面)

每个语言目录下的 12 个模块页面必须满足:

1. **头部锚点**:`> 锚点:Java 的 X 对应本语言 Y`,一句话定位本页。
2. **验证环境**:头部必须含「验证环境:语言版本 / OS / 依赖版本」字段(CI 校验,缺失即失败)。
3. **对照表**:主体为对照表(左列 Java,右列本语言),表内代码片段必须可运行。
4. **Java 思维陷阱**:每模块 ≥1 个「⚠️ Java 思维陷阱」提示块(引用格式)。
5. **权威链接**:底部附官方文档/权威链接。
6. **实际验证**:页面中所有代码示例必须实际运行过,并注明运行输出或测试结果。**不提交未验证代码。**

## 代码示例要求

- 版本基线:以各语言目录 README 标注的 CI 版本为准,不写死未经验证的版本号。
- 并发/正确性相关的示例必须带可运行测试。
- examples/ 下的服务实现必须满足统一接口契约(见 examples/README.md),CI 全绿。

## 新增语言流程

1. 复制 `languages/go/` 目录结构(12 个文件)。
2. 按 01→11 顺序逐篇替换内容,锚点保持「Java X → 本语言 Y」。
3. 在 `examples/` 新增 `<语言>-order-service` 实现,并在 `.github/workflows/examples-test.yml` 矩阵加一行。
4. 更新根 README 的「语言覆盖」表。
5. 语言目录命名与允许列表见 `.github/scripts/check_structure.py`。

## PR 流程

1. fork + 分支,提交信息遵循 Conventional Commits(`feat:` / `fix:` / `docs:` / `ci:` / `chore:`)。
2. 必须通过 `docs-check`(链接检查 + 结构校验);修改 examples/ 必须通过 `examples-test`。
3. PR 描述中说明:改动模块、验证环境、本地运行结果。

## 结构校验脚本

`.github/scripts/check_structure.py` 本地可运行:

```bash
python .github/scripts/check_structure.py
```

Errors 会令 CI 失败;Warnings 提示未完成阶段的内容(不阻塞)。
