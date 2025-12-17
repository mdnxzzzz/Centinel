package core

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/medina/ci-sentinel/pkg/models"
	"github.com/medina/ci-sentinel/pkg/plugin"
)

type Engine struct {
	scanners []plugin.Scanner
}

// New returns a new engine instance
func New() *Engine {
	return &Engine{
		scanners: []plugin.Scanner{},
	}
}

// RegisterScanner adds a scanner to the engine
func (e *Engine) RegisterScanner(s plugin.Scanner) {
	e.scanners = append(e.scanners, s)
}

// Run executes all registered scanners on the target path
func (e *Engine) Run(ctx context.Context, targetPath string) ([]models.Issue, error) {
	var allIssues []models.Issue
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTime := time.Now()

	for _, s := range e.scanners {
		wg.Add(1)
		go func(s plugin.Scanner) {
			defer wg.Done()

			// In a real implementation we would have proper logging here
			// fmt.Printf("Running scanner: %s\n", s.Name())

			issues, err := s.Run(ctx, targetPath)
			if err != nil {
				fmt.Printf("Error running scanner %s: %v\n", s.Name(), err)
				return
			}

			mu.Lock()
			allIssues = append(allIssues, issues...)
			mu.Unlock()
		}(s)
	}

	wg.Wait()
	_ = startTime // Use logging later

	return allIssues, nil
}
