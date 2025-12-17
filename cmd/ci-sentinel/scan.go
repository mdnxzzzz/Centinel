package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/medina/ci-sentinel/internal/scanners/docker"
	"github.com/medina/ci-sentinel/internal/scanners/secrets"
	"github.com/medina/ci-sentinel/pkg/core"
	"github.com/medina/ci-sentinel/pkg/models"
	"github.com/spf13/cobra"
)

var (
	// Styles
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Bold(true).
			Padding(0, 1)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0"))

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00ADD8")).
			Bold(true)
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [target]",
	Short: "Scan a directory for security issues",
	Long: `Scans the provided directory path for:
- CI/CD Misconfigurations (GitHub Actions, GitLab CI)
- Container Security (Dockerfile, Docker Compose)
- Secrets and Hardcoded Credentials`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := "."
		if len(args) > 0 {
			target = args[0]
		}

		profile, _ := cmd.Flags().GetString("profile")

		// Initialization
		eng := core.New()

		// Register Scanners
		eng.RegisterScanner(secrets.New())
		eng.RegisterScanner(docker.New())

		// UI Header
		fmt.Println(headerStyle.Render("CI-SENTINEL v0.1.0"))
		fmt.Printf("Target:  %s\n", target)
		fmt.Printf("Profile: %s\n", profile)
		fmt.Println(strings.Repeat("─", 60))
		fmt.Println()

		fmt.Println(infoStyle.Render("[*] Scanning..."))

		// Execute Scan
		issues, err := eng.Run(context.Background(), target)
		if err != nil {
			fmt.Printf("Error during scan: %v\n", err)
			os.Exit(1)
		}

		// Display Results
		if len(issues) == 0 {
			fmt.Println(titleStyle.Render("SCAN COMPLETE"))
			fmt.Println("No issues found.")
			fmt.Println(strings.Repeat("─", 60))
			return
		}

		for _, i := range issues {
			severityColor := lipgloss.Color("#00FF00") // Low/Info
			switch i.Severity {
			case models.SeverityCritical:
				severityColor = lipgloss.Color("#FF0000")
			case models.SeverityHigh:
				severityColor = lipgloss.Color("#FF5500")
			case models.SeverityMedium:
				severityColor = lipgloss.Color("#FFFF00")
			}

			severityStyle := lipgloss.NewStyle().Foreground(severityColor).Bold(true)

			fmt.Printf("[%s] %s\n", severityStyle.Render(string(i.Severity)), i.Title)
			fmt.Printf("  File: %s:%d\n", i.File, i.Line)
			fmt.Printf("  Snippet: %s\n", i.Snippet)
			fmt.Printf("  Fix: %s\n", i.Remediation)
			fmt.Println()
		}

		fmt.Println(strings.Repeat("─", 60))
		fmt.Printf("Total Issues: %d\n", len(issues))
		fmt.Println()
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Render("made by @mdnxzzzz with love ❤️"))
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Local flags
	scanCmd.Flags().String("fail-on", "high", "Exit with error if severity >= threshold")
}
