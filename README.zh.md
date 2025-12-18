# WellAll 基础设施（L1）Monorepo

[![English](https://img.shields.io/badge/Language-English-blue)](README.md)
![Status](https://img.shields.io/badge/Status-Active-2ea44f)
![License](https://img.shields.io/badge/License-Apache--2.0-blue)
![Layer](https://img.shields.io/badge/Layer-L1%20Infrastructure-0ea5e9)
![Spec](https://img.shields.io/badge/Spec-Health%20JSON-005bbb)


官网：https://www.wellally.tech/

## 概览
本仓库汇聚 WellAll 生态的 **基础设施级（L1）规范与 Schema**，作为健康数据结构化与跨系统互操作的权威参考来源（source of truth）。

## 权威性与适用范围
- 字段命名、数据类型、验证规则与兼容性原则以 `infrastructure/specs/health-json-spec` 为最高约束。
- 本仓库内的 Schema 面向上层工具/产品提供稳定的数据契约。
- `archive/` 用于保留历史 L2/L5 工具与实验项目，不属于当前基础设施契约的一部分。

## 仓库结构
- `infrastructure/schemas/health` — 个人健康数据总 Schema（顶层资源与基础类型）。
- `infrastructure/schemas/lab-report` — 生化/检验指标、单位与参考区间标准化。
- `infrastructure/schemas/imaging-report` — 影像报告结构化 Schema（CT/超声/MRI/X 光/PET-CT）。
- `infrastructure/schemas/medication` — 药物记录与剂量/频次表达。
- `infrastructure/schemas/family-health` — 家庭健康树与家族史数据结构。
- `infrastructure/specs/health-json-spec` — 健康 JSON 规范白皮书。
- `archive/` — 历史工具与原型项目。

## 使用指南
- 各模块 README 说明设计范围与状态；建议在模块内补齐 JSON Schema 与示例数据。
- 统一遵循 `health-json-spec` 的字段命名、验证规则与兼容性原则。
- 建议为每个 Schema 引入版本号（如 `v0.1.0`）并维护 `CHANGELOG.md`。
- 示例校验：`python infrastructure/scripts/validate_examples.py`（依赖 `jsonschema`）。

## 版本与治理
- 建议采用语义化版本（SemVer），并在变更说明中明确破坏性影响与迁移方式。
- 重大结构调整请附迁移指南与示例数据，确保可落地执行。

## 贡献
- 欢迎提交 Issue / PR，优先聚焦字段定义、命名一致性与验证规则完善。
- 变更 Schema 时请附示例数据与向后兼容性说明。

## 许可证
Apache 2.0，见 `LICENSE`。
