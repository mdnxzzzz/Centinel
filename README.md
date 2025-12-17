# CI-Sentinel

<div align="center">
  <h3>🛡️ CI-Sentinel 🛡️</h3>
  <p><strong>Enterprise-Grade Defensive Security for CI/CD Pipelines</strong></p>

  [![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8.svg?style=flat-square)](https://golang.org)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
  [![Security](https://img.shields.io/badge/security-defensive-green.svg?style=flat-square)](SECURITY.md)
  [![Compliance](https://img.shields.io/badge/compliance-ready-purple.svg?style=flat-square)](docs/COMPLIANCE.md)
</div>

---

## 📖 Overview

**CI-Sentinel** is a static analysis tool designed to audit, detect, and remediate security risks within Build Pipelines (GitHub Actions, GitLab CI, Jenkins) and Deployment Artifacts (Dockerfile, Kubernetes manifests).

Built with a **Zero Trust** philosophy for the supply chain, it empowers DevSecOps teams to prevent hardcoded secrets, insecure dependencies, and misconfigurations before they reach production.

> ⚠️ **Disclaimer**: This tool is strictly **DEFENSIVE**. It is designed for auditing and hardening purposes only. It does not contain any offensive capabilities or payloads.

## ✨ Key Features

- **🔍 Pipeline Analysis**: Detects untagged actions, insecure permissions, and risky script injections in GitHub Actions, GitLab CI, and Jenkins.
- **📦 Artifact Scanning**: Audits `Dockerfiles` and `docker-compose.yml` for best practices (CIS Benchmarks).
- **🛠️ Self-Healing (Auto-Fix)**: Safe, dry-run capable auto-remediation for common issues (e.g., pinning versions, replacing secrets).
- **📊 Enterprise Reporting**: Generates HTML executive summaries and JSON artifacts for compliance audits.
- **🛡️ Security Profiles**: Built-in profiles for `OWASP`, `CIS`, and `Enterprise` standards.

## 🚀 Getting Started

### Installation

```bash
# via Go install
go install github.com/medina/ci-sentinel/cmd/ci-sentinel@latest

# via Docker
docker run -v $(pwd):/src ci-sentinel:latest scan /src
```

### Basic Usage

**Scan a project:**
```bash
ci-sentinel scan ./my-project
```

**Fix issues automatically (Dry Run):**
```bash
ci-sentinel fix ./my-project --dry-run
```

**Generate Compliance Report:**
```bash
ci-sentinel report --format html --out audit_report.html
```

## 📐 Architecture

CI-Sentinel follows a modular Hexagonal Architecture:
- **Core**: Orchestrates lifecycle and concurrency.
- **Scanners**: Pluggable modules for different target types (CI, Containers, Configs).
- **Rules Engine**: Policy definitions handling `OWASP` and `CIS` logic.

See [ARCHITECTURE.md](docs/ARCHITECTURE.md) for deep dive.

## 🤝 Contributing

We welcome contributions from the community! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## ⚖️ License

Distributed under the MIT License. See `LICENSE` for more information.

---
<div align="center">
  <p><i>"Security is a process, not a product."</i></p>
</div>
