# Archive

[![中文](https://img.shields.io/badge/Language-中文-red)](README.zh.md)

Historical L2/L3/L4/L5 utilities, AI helpers, and compliance prototypes. These are **not** part of the current L1 contract but remain available for reference or extraction into standalone repositories when needed.

## 🎯 Implementation Status

Run `python PROJECT_STATUS.py` to see detailed implementation status.

**Quick Summary:**
- ✅ **2 Complete** - Fully implemented with examples
- 🚧 **1 In Progress** - Core implementation done
- 📝 **11 Planned** - Structure ready, implementation pending
- 📦 **14 Total** - All with Python + pyproject.toml structure

## 📦 Projects Catalog

### ✅ Complete Implementations

#### **wellally-lab-parser**
OCR lab slips → structured JSON using GLM-4V-Flash + LangChain
- 🔬 Parse lab reports from images
- 📊 Extract structured data (tests, values, units, ranges)
- 🗺️  Map to WellAlly LabReport schema
- ✅ Built-in validation
- 🌐 Support for Chinese & English reports

**Tech Stack:** LangChain, Zhipu AI GLM-4V-Flash (free), WellAlly schemas

#### **wellally-healthkit-mapper**
Map Apple HealthKit exports to WellAlly schemas
- 📱 Parse HealthKit XML exports
- 💓 Map vital signs (HR, BP, temp, O2)
- 🧪 Map lab results (glucose, HbA1c)
- ⚖️ Map body measurements (weight, height, BMI)
- 🏃 Map workouts and activities
- 🏷️  LOINC code mapping

**Tech Stack:** Python XML parsing, WellAlly schemas, LOINC/UCUM standards

### 🚧 In Progress

#### **wellally-unit-normalizer**
Normalize clinical units (mg/dL ↔ mmol/L) for pre-analytics standardization
- 🔢 UCUM unit conversions
- 🩸 Lab-specific conversions (glucose, cholesterol)
- 🌡️  Temperature conversions
- 📏 Mass, volume, length conversions

### 📝 Planned Implementations

#### **wellally-pdf-medical-parser**
Parse medical PDFs to structured fields for legacy migration
- 📄 PDF text extraction
- 🏥 Medical report parsing
- 🤖 AI-powered structure extraction
- 🌐 Multi-language support

#### **wellally-medical-timeline**
Build patient event timelines for clinician/patient longitudinal views
- 📅 Chronological event ordering
- 🔄 Multi-source data aggregation
- 🏷️  Event categorization
- 📊 Timeline visualization data

#### **wellally-anomaly-flagger**
Data-quality anomaly checks for ETL quality gates and monitoring
- 📈 Statistical anomaly detection
- ✅ Range validation
- 🔍 Duplicate detection
- ⚠️  Missing data flagging

#### **wellally-trend-detector**
Non-diagnostic trend analytics for wellness and monitoring dashboards
- 📉 Time series analysis
- ↗️  Trend direction detection
- 🔄 Seasonal pattern recognition
- 🚨 Anomaly detection

#### **wellally-data-correlation**
Correlation exploration between metrics and behaviors for hypothesis surfacing
- 📊 Pearson/Spearman correlation
- 🔢 Multi-variate analysis
- ⏱️  Lag correlation
- 💡 Hypothesis generation

#### **wellally-report-structurer-ai**
Unstructured → structured field extraction for registry/report automation
- 🤖 NLP-based extraction
- 🏷️  Entity recognition
- 🔗 Relationship extraction
- 📑 Multi-format support

#### **wellally-fhir-lite**
Lightweight FHIR → WellAlly mapping for PoC integrations with FHIR APIs
- 🔄 FHIR R4 resource mapping
- 📊 Observation → LabResult
- 👤 Patient → Person
- ⚡ Minimal FHIR subset

#### **wellally-consent-model**
Consent lifecycle model for fine-grained data access governance
- ✅ Consent capture
- 🔐 Consent verification
- 📜 Audit trail
- 🇪🇺 GDPR compliance

#### **wellally-health-audit-log**
Tamper-resistant access logging for compliance evidence
- 📝 Access logging
- ⛓️  Blockchain-style integrity
- 🔍 Query audit trail
- 📊 Compliance reporting

#### **wellally-health-data-anonymizer**
De-identification/anonymization toolkit for privacy-safe sharing
- 🔒 PII removal
- 📅 Date shifting
- 🔢 K-anonymity
- ⚠️  Re-identification risk analysis

#### **wellally-radiation-dose-calc**
CT dose accumulation calculator for safety tracking
- ☢️  DLP calculation
- 📈 Cumulative dose tracking
- 👶 Age-adjusted risk
- 🫁 Organ-specific doses

## 🏗️  Project Structure

Each project follows a consistent structure:

```
wellally-{project-name}/
├── language/
│   ├── python/              # Python implementation
│   │   ├── pyproject.toml   # Dependencies & config
│   │   ├── README.md        # Documentation
│   │   ├── examples.py      # Usage examples
│   │   └── wellally_{name}/ # Package code
│   │       ├── __init__.py
│   │       ├── *.py         # Core modules
│   │       └── py.typed     # Type hints
│   ├── typescript/          # Future TS implementation
│   ├── go/                  # Future Go implementation
│   └── rust/                # Future Rust implementation
└── README.md                # Project overview
```

## 🚀 Quick Start

### Install a Project

```bash
# Navigate to project
cd wellally-lab-parser/language/python

# Install in development mode
pip install -e .

# Run examples
python examples.py
```

### Use in Your Code

```python
# Lab Parser
from wellally_lab_parser import LabReportParser
parser = LabReportParser()
result = parser.parse_image("lab_report.jpg")

# HealthKit Mapper
from wellally_healthkit_mapper import HealthKitMapper
mapper = HealthKitMapper()
mapper.load_export("export.xml")
lab_reports = mapper.map_lab_results("patient-123")
```

## 🔧 Development

### Add a New Feature

1. Navigate to the project directory
2. Edit the relevant Python module
3. Add tests
4. Update documentation
5. Run examples to verify

### Create a New Project

Use existing projects as templates:
- Copy structure from `wellally-lab-parser` or `wellally-healthkit-mapper`
- Update `pyproject.toml` with project-specific dependencies
- Implement core logic in package modules
- Add examples and documentation

## 📚 Dependencies

All projects depend on:
- **wellally** (>= 0.1.0) - Core schemas from `/language/python`
- Python >= 3.8

Project-specific dependencies are listed in each `pyproject.toml`.

## 🤝 Contributing

These are reference implementations. To contribute:

1. Choose a "Planned" project from the list
2. Implement core functionality in Python
3. Add comprehensive examples
4. Document all features in README
5. Ensure compatibility with WellAlly schemas

## ⚖️  License

MIT License - see individual project LICENSE files

## 🔗 Links

- [WellAlly Platform](https://www.wellally.tech/)
- [WellAlly Schemas](../language/python/)
- [LOINC Database](https://loinc.org/)
- [FHIR Specification](https://hl7.org/fhir/)
- [UCUM Units](https://ucum.org/)

## How to Reuse
- Treat these as reference implementations; harden and version before production.
- When extracting a project, align schemas/fields with the active L1 specs and add tests.
- Update licenses and dependency baselines during extraction.
