# CI-Sentinel Reporting Design

## Goals
- **Auditability**: Provide a permanent record of the security state at a point in time.
- **Interoperability**: JSON output for integration with other Dashboards/SIEMs.
- **Clarity**: HTML reports for non-technical or management stakeholders.

## 1. JSON Schema (`report.json`)
The machine-readable format.

```json
{
  "meta": {
    "tool": "ci-sentinel",
    "version": "1.0.0",
    "timestamp": "2023-10-27T10:00:00Z",
    "scan_duration_ms": 150
  },
  "target": {
    "path": "/src/project",
    "branch": "main",
    "commit": "a1b2c3d4"
  },
  "summary": {
    "score": 85,
    "total_issues": 3,
    "breakdown": {
      "critical": 0,
      "high": 1,
      "medium": 1,
      "low": 1
    }
  },
  "findings": [
    {
      "id": "CIS-DOCKER-001",
      "severity": "HIGH",
      "title": "Secrets in Environment Variables",
      "description": "Hardcoded secret found in yaml values.",
      "location": {
        "file": "config.yaml",
        "line": 45
      },
      "remediation": "Move to secret manager"
    }
  ]
}
```

## 2. HTML Report (`report.html`)
A self-contained, single-file HTML report.

### Structure
1.  **Executive Summary (Hero Section)**
    *   Big "Score Card" (A/B/C/D/F).
    *   Pie chart of issue severity.
    *   Scan metadata (time, target).

2.  **Top Risks (High Priority)**
    *   Detailed breakdown of Critical/High issues.
    *   Why they matter.
    *   How to fix them immediately.

3.  **Full Findings Table**
    *   Sortable/Filterable table (JavaScript included in the single file).
    *   Columns: Severity, Rule, Location, Description.

4.  **Recommendations / Roadmap**
    *   General advice based on the profile (e.g., "Enable Branch Protection", "Use Signed Commits").

### Technology
- **Go Templates**: `html/template`
- **Embedded CSS**: Simple, printable stylesheet.
- **Embedded JS**: Minimal vanilla JS for sorting/filtering the table.
