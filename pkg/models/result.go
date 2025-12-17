package models

type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityHigh     Severity = "HIGH"
	SeverityMedium   Severity = "MEDIUM"
	SeverityLow      Severity = "LOW"
	SeverityInfo     Severity = "INFO"
)

// Issue represents a single security finding
type Issue struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Severity    Severity `json:"severity"`
	File        string   `json:"file"`
	Line        int      `json:"line"`
	Col         int      `json:"col"`
	Snippet     string   `json:"snippet"`
	Remediation string   `json:"remediation"`
}

// CheckResult contains all issues found by a specific scanner
type CheckResult struct {
	ScannerName string  `json:"scanner_name"`
	Issues      []Issue `json:"issues"`
}
