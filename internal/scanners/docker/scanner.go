package docker

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/medina/ci-sentinel/pkg/models"
)

type Scanner struct{}

func New() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Name() string {
	return "ContainerSecurity"
}

func (s *Scanner) Run(ctx context.Context, rootPath string) ([]models.Issue, error) {
	var issues []models.Issue

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			return nil
		}

		if info.Name() == "Dockerfile" || strings.HasSuffix(info.Name(), ".Dockerfile") {
			return scanDockerfile(path, &issues)
		}
		return nil
	})

	return issues, err
}

func scanDockerfile(path string, issues *[]models.Issue) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	hasUserInstruction := false

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Check 1: Unpinned Latest Tag
		if strings.HasPrefix(line, "FROM") {
			if strings.Contains(line, ":latest") || !strings.Contains(line, ":") {
				*issues = append(*issues, models.Issue{
					ID:          "CNT-001",
					Title:       "Base Image uses :latest tag",
					Description: "Using :latest can lead to non-reproducible builds and unexpected breaking changes.",
					Severity:    models.SeverityMedium,
					File:        path,
					Line:        lineNum,
					Snippet:     line,
					Remediation: "Pin the version explicitly (e.g. node:18-alpine)",
				})
			}
		}

		// Check 2: Check for USER instruction
		if strings.HasPrefix(line, "USER") {
			hasUserInstruction = true
		}
	}

	if !hasUserInstruction {
		*issues = append(*issues, models.Issue{
			ID:          "CNT-002",
			Title:       "Container running as Root",
			Description: "No USER instruction found. The container will likely run as root, which exposes the host to potential privilege escalation attacks.",
			Severity:    models.SeverityHigh,
			File:        path,
			Line:        1, // Attribution to whole file
			Snippet:     "(Missing USER instruction)",
			Remediation: "Add 'USER nonroot' or similar before the ENTRYPOINT/CMD.",
		})
	}

	return nil
}
