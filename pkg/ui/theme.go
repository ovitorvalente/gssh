package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Header and title
	styleHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#24292F")).
			Padding(0, 2).
			MarginBottom(0)

	styleSeparator = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#475569"))

	// Steps
	styleStepNum = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64748B")).
			Bold(true)

	styleOk = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#22C55E")).
		Bold(true)

	styleSkip = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8"))

	styleWarn = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F59E0B"))

	styleFail = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EF4444")).
			Bold(true)

	// Content
	styleMessage = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E2E8F0"))

	styleKey = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#38BDF8"))

	styleLink = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#38BDF8")).
			Underline(true)

	styleCommand = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A78BFA"))
)
