package secrets

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/medina/ci-sentinel/pkg/models"
)

type Scanner struct{}

func New() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Name() string {
	return "SecretScanner"
}

var secretPatterns = map[string]*regexp.Regexp{
	"AWS Access Key":      regexp.MustCompile(`(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}`),
	"Generic Private Key": regexp.MustCompile(`-----BEGIN (RSA|DSA|EC|OPENSSH|PGP) PRIVATE KEY-----`),
	"Generic Token":       regexp.MustCompile(`(api_key|access_token|secret_key)\s*(=|:)\s*['"][a-zA-Z0-9_=-]{16,}['"]`),
	"GitHub Token":        regexp.MustCompile(`ghp_[a-zA-Z0-9]{36}`),
}

func (s *Scanner) Run(ctx context.Context, rootPath string) ([]models.Issue, error) {
	var issues []models.Issue

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip unreadable
		}
		if info.IsDir() {
			// Skip hidden dirs
			if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			return nil
		}

		// Simple exclusion for non-text files based on extension if needed
		// For now we scan everything that looks text-like or has no extension

		return scanFile(path, &issues)
	})

	return issues, err
}

func scanFile(path string, issues *[]models.Issue) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		for name, pattern := range secretPatterns {
			if pattern.MatchString(line) {
				*issues = append(*issues, models.Issue{
					ID:          "SEC-001",
					Title:       "Potential Secret Found: " + name,
					Description: "A pattern matching a known secret format was detected.",
					Severity:    models.SeverityHigh,
					File:        path,
					Line:        lineNum,
					Snippet:     strings.TrimSpace(line),
					Remediation: "Remove the secret and use environment variables.",
				})
			}
		}
	}
	return nil
}
