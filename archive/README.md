# Archive

[![中文](https://img.shields.io/badge/Language-中文-red)](README.zh.md)

Historical L2/L3/L4/L5 utilities, AI helpers, and compliance prototypes. These are **not** part of the current L1 contract but remain available for reference or extraction into standalone repositories when needed.

## Catalog & Typical Use
- **wellally-lab-parser** — OCR lab slips → structured JSON; intake pipelines for lab data.
- **wellally-healthkit-mapper** — Map Apple HealthKit to WellAll schemas; consumer/BYOD data sync.
- **wellally-pdf-medical-parser** — Parse medical PDFs to structured fields; legacy migration.
- **wellally-medical-timeline** — Build patient event timelines; clinician/patient longitudinal views.
- **wellally-unit-normalizer** — Normalize clinical units (mg/dL ↔ mmol/L); pre-analytics standardization.
- **wellally-anomaly-flagger** — Data-quality anomaly checks; ETL quality gates and monitoring.
- **wellally-trend-detector** — Non-diagnostic trend analytics; wellness and monitoring dashboards.
- **wellally-data-correlation** — Correlation exploration between metrics and behaviors; hypothesis surfacing.
- **wellally-report-structurer-ai** — Unstructured → structured field extraction; registry/report automation.
- **wellally-fhir-lite** — Lightweight FHIR → WellAll mapping; PoC integrations with FHIR APIs.
- **wellally-consent-model** — Consent lifecycle model; fine-grained data access governance.
- **wellally-health-audit-log** — Tamper-resistant access logging; compliance evidence.
- **wellally-health-data-anonymizer** — De-identification/anonymization toolkit; privacy-safe sharing.
- **wellally-radiation-dose-calc** — CT dose accumulation calculator; safety tracking.

## How to Reuse
- Treat these as reference implementations; harden and version before production.
- When extracting a project, align schemas/fields with the active L1 specs and add tests.
- Update licenses and dependency baselines during extraction.
