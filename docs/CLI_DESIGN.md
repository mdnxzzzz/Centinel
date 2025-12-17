# CI-Sentinel CLI Design

## Design Philosophy
The CLI is the primary interface for users. It must be Professional, Clear, and Actionable.
We use **Lipgloss** for styling to provide subtle, accessible coloring that highlights severity without being garish.

## Global Flags
- `--config, -c`: Path to config file (default: `.ci-sentinel.yaml`)
- `--verbose, -v`: Enable debug logging
- `--no-color`: Disable all ANSI color output (for logs/pipelines)
- `--profile`: Security profile to use (`owasp`, `cis`, `enterprise`) [default: `enterprise`]

## Commands

### 1. `scan`
The main command. Analyzes the current directory or a specific target.

**Usage:**
```bash
ci-sentinel scan [path] [flags]
```

**Flags:**
- `--fail-on [severity]`: Exit code 1 if severity >= threshold (low, medium, high, critical)
- `--ignore-paths`: List of glob patterns to ignore
- `--output`: Export results to file (json, yaml)

**Example Output:**

```text
  CI-SENTINEL v1.0.0
  Target: /home/user/projects/backend-api
  Profile: Enterprise Grade
  ────────────────────────────────────────────────────────────

  [HIGH]     Hardcoded AWS Secret Key
  file:      deploy/k8s/secrets.yaml:12
  recommend: Use external secrets manager (Vault/AWS Secrets Manager)

  [MEDIUM]   Docker Image without tag pinning
  file:      Dockerfile:1
  found:     FROM golang:latest
  recommend: Use mutable tags e.g. golang:1.21-alpine

  [LOW]      Missing Security Headers in Nginx config
  file:      nginx.conf:24

  ────────────────────────────────────────────────────────────
  SUMMARY
  Total Issues: 3   (1 High, 1 Medium, 1 Low)
  Security Score: 85/100 [B+]
  
  Run 'ci-sentinel fix' to apply auto-remediation.
```

### 2. `fix`
Auto-remediation of detected issues.

**Usage:**
```bash
ci-sentinel fix [flags]
```

**Flags:**
- `--dry-run`: Show what would be changed without modifying files (default: true)
- `--apply`: Actually apply the changes
- `--backup`: Create .bak files (default: true)

**Example Flow:**
```text
> ci-sentinel fix --dry-run

[+] Dockerfile:1
- FROM golang:latest
+ FROM golang:1.21.5 # Fixed version based on current latest

[+] deploy/k8s/secrets.yaml:12
- params: "AKIA..."
+ params: "${AWS_ACCESS_KEY}"

Total changes proposed: 2.
Run with --apply to execute.
```

### 3. `report`
Generates comprehensive reports for stakeholders.

**Usage:**
```bash
ci-sentinel report --format html --out report.html
```

### 4. `version`
Displays version, build date, and commit hash.

## Exit Codes
- `0`: Success (No issues found or issues below threshold)
- `1`: Security Failure (Issues found >= threshold)
- `2`: Runtime Error (Config missing, permission denied)
