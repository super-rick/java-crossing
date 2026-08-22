## 变更说明

- **模块**:`languages/{lang}/` 哪个模块,或 `examples/xxx-order-service` 哪部分
- **变更类型**:内容新增 / 示例代码 / CI / 修复
- **验证环境**:语言版本 / OS(CI 会自动复验,但请说明本地实测结果)

## 自检清单

- [ ] 对照表保持「左列 Java,右列本语言」
- [ ] 每模块头部含「验证环境」字段(`check_structure.py` 会校验)
- [ ] 代码示例已本地实际运行,输出与文档一致
- [ ] ⚠️ Java 思维陷阱块 ≥1 个(语言内容 PR)
- [ ] `python .github/scripts/check_structure.py` 本地通过
- [ ] examples/ 改动已通过对应语言测试(go test / pytest / npm test / phpunit)

## 说明

修改 examples/ 会触发 `examples-test` CI 矩阵(四语言编译+测试);只改文档触发 `docs-check`(链接+结构校验)。请先本地跑通再提交,减少 CI 往返。
