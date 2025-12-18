# WellAll Infrastructure (L1) Monorepo

[![中文](https://img.shields.io/badge/Language-中文-red)](README.zh.md)
![Status](https://img.shields.io/badge/Status-Active-2ea44f)
![License](https://img.shields.io/badge/License-Apache--2.0-blue)
![Layer](https://img.shields.io/badge/Layer-L1%20Infrastructure-0ea5e9)
![Spec](https://img.shields.io/badge/Spec-Health%20JSON-005bbb)


Official site: https://www.wellally.tech/

## Overview
This monorepo hosts WellAll’s **Layer-1 (infrastructure) schemas and specifications**, acting as the source of truth for structured health data across the ecosystem.

## Authority & Scope
- All naming, validation, and compatibility rules are anchored by `infrastructure/specs/health-json-spec`.
- Schemas here are the canonical contracts consumed by upper-layer tools and products.
- `archive/` keeps historical L2/L5 experiments and utilities; they are not part of the current infra contract.

## Repository Layout
- `infrastructure/schemas/health` — master personal health schema.
- `infrastructure/schemas/lab-report` — lab metrics, units, and reference ranges.
- `infrastructure/schemas/imaging-report` — structured imaging report schema (CT/US/MRI/X-ray/PET-CT).
- `infrastructure/schemas/medication` — medication records and dosing.
- `infrastructure/schemas/family-health` — family health tree and history.
- `infrastructure/specs/health-json-spec` — the Health JSON specification.
- `archive/` — historical L2/L5 tools and prototypes.

## Usage
- Each module ships its own README for scope and status; add JSON Schemas and examples beside it.
- Follow `health-json-spec` for naming, validation, and compatibility rules.
- Prefer semantic versions (e.g., `v0.1.0`) and keep a `CHANGELOG.md` per schema.
- Validate examples: `python infrastructure/scripts/validate_examples.py` (requires `jsonschema`).

## Versioning & Governance
- Use SemVer and call out breaking changes with migration notes and sample payloads.
- Significant structural changes should include backward-compatibility guidance.

## Contributing
- Issues/PRs welcome—focus on field definitions, naming consistency, and validation rules.
- When altering schemas, include sample data and compatibility notes.

## License
Apache 2.0. See `LICENSE`.
