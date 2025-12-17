package plugin

import (
	"context"

	"github.com/medina/ci-sentinel/pkg/models"
)

// Scanner is the interface that all security scanners must implement
type Scanner interface {
	// Name returns the unique name of the scanner
	Name() string

	// Run executes the scan on the given root path
	Run(ctx context.Context, rootPath string) ([]models.Issue, error)
}
