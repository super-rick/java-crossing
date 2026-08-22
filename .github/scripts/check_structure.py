#!/usr/bin/env python3
"""校验仓库结构完整性。

规则(Errors 令 CI 失败):
1. 根文件 README.md / LICENSE / CONTRIBUTING.md 必须存在。
2. languages/ 下每个已登记语言目录必须 12 个模块文件齐全,且每个文件含「验证环境」字段。
   未登记目录名 → Error。
3. languages/ / docs/ / topics/ 尚未创建时给出 Warning(阶段推进提示,不阻塞)。

用法: python .github/scripts/check_structure.py
"""
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]

REQUIRED_ROOT = ["README.md", "LICENSE", "CONTRIBUTING.md"]

LANG_MODULES = [
    "README.md",
    "01-语言特性.md", "02-基础语法.md", "03-数据结构.md", "04-并发编程.md",
    "05-IO与网络.md", "06-测试.md", "07-生产框架.md", "08-部署运维.md",
    "09-监控.md", "10-常见误区.md", "11-学习路径.md",
]

ALLOWED_LANGS = {"go", "python", "nodejs", "php"}

REQUIRED_DOCS = [
    "00-总览.md", "01-学习路径总图.md", "02-语言全景对比.md",
    "LANGUAGE_TEMPLATE.md", "GLOSSARY.md",
]

REQUIRED_TOPICS = [
    "并发模型对比.md", "IO模型对比.md", "内存与GC对比.md",
    "包管理与构建对比.md", "测试体系对比.md", "Web框架生态对比.md",
    "部署形态对比.md", "监控体系对比.md",
]


def main() -> int:
    errors: list[str] = []
    warnings: list[str] = []

    for f in REQUIRED_ROOT:
        if not (ROOT / f).exists():
            errors.append(f"缺失根文件: {f}")

    langs_dir = ROOT / "languages"
    if langs_dir.exists():
        for d in sorted(p for p in langs_dir.iterdir() if p.is_dir()):
            if d.name not in ALLOWED_LANGS:
                errors.append(f"未登记的语言目录: {d.name} (允许: {sorted(ALLOWED_LANGS)})")
                continue
            for m in LANG_MODULES:
                p = d / m
                if not p.exists():
                    errors.append(f"缺失模块文件: {d.name}/{m}")
                elif "验证环境" not in p.read_text(encoding="utf-8"):
                    errors.append(f"{d.name}/{m} 缺少「验证环境」字段")
    else:
        warnings.append("languages/ 尚未创建(Phase 1+ 预期)")

    docs_dir = ROOT / "docs"
    if docs_dir.exists():
        for f in REQUIRED_DOCS:
            if not (docs_dir / f).exists():
                warnings.append(f"docs/{f} 未创建(Phase 1 内容)")
    else:
        warnings.append("docs/ 尚未创建(Phase 1 内容)")

    topics_dir = ROOT / "topics"
    if topics_dir.exists():
        for f in REQUIRED_TOPICS:
            if not (topics_dir / f).exists():
                warnings.append(f"topics/{f} 未创建(Phase 1 内容)")
    else:
        warnings.append("topics/ 尚未创建(Phase 1 内容)")

    examples_dir = ROOT / "examples"
    if examples_dir.exists():
        if not (examples_dir / "README.md").exists():
            errors.append("examples/README.md 缺失(契约总览)")
        for d in sorted(p for p in examples_dir.iterdir() if p.is_dir()):
            if not (d / "README.md").exists():
                warnings.append(f"examples/{d.name}/README.md 缺失")

    print("== Errors ==")
    for e in errors:
        print(" -", e)
    print("== Warnings ==")
    for w in warnings:
        print(" -", w)

    if errors:
        print(f"FAIL: {len(errors)} 个错误")
        return 1
    print(f"OK: 结构校验通过({len(warnings)} 个警告, 均为后续阶段内容)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
