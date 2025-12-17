# CI-Sentinel Roadmap

## Phase 1: MVP (Month 1)
**Goal:** Functional CLI capable of scanning CI workflows and basic Dockerfiles.
- [ ] **Core Engine**: Basic plugin loading and execution.
- [ ] **Scanners**: 
    - GitHub Actions (Secrets, Untagged actions)
    - Dockerfile (User root, :latest tags)
- [ ] **CLI**: `scan` command with JSON output.
- [ ] **Rules**: Hardcoded set of ~10 critical rules.

## Phase 2: Remediation & Depth (Month 2)
**Goal:** Interactive fixing and deeper analysis.
- [ ] **Auto-Fix**: `fix` command implementation (Regex based replacement).
- [ ] **Scanners**:
    - GitLab CI support.
    - Generic Secret scanner (High entropy + patterns).
- [ ] **Reporting**: HTML Report writer.
- [ ] **Config**: Support for `.ci-sentinel.yaml` and ignore logic.

## Phase 3: Enterprise & Compliance (Month 3)
**Goal:** Readiness for organizational rollout.
- [ ] **Profiles**: Implant `OWASP` and `CIS` profiles.
- [ ] **Scoring**: Implement weighted scoring algorithm 0-100.
- [ ] **Performance**: Parallel scanning for monorepos.
- [ ] **Integration**: GitHub Action / GitLab Component wrappers for easy CI usage.

## Phase 4: Expansion (Month 4+)
- [ ] **Jenkins**: Support for Declarative Jenkinsfiles.
- [ ] **Kubernetes**: Helm chart scanning.
- [ ] **Policy**: Open Policy Agent (Rego) integration for custom rules.
- [ ] **IDE**: VS Code Extension wrapping the CLI.
