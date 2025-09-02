package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	primaryColor   = lipgloss.Color("39")  // Blue
	secondaryColor = lipgloss.Color("208") // Orange
	successColor   = lipgloss.Color("46")  // Green
	errorColor     = lipgloss.Color("196") // Red
	mutedColor     = lipgloss.Color("241") // Gray
	draftedColor   = lipgloss.Color("243") // Dark gray

	// Base styles
	baseStyle = lipgloss.NewStyle().
		Padding(1, 2)

	// Header style
	headerStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primaryColor)

	// Table styles
	tableHeaderStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Padding(0, 1).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(mutedColor)

	tableRowStyle = lipgloss.NewStyle().
		Padding(0, 1)

	selectedRowStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(primaryColor).
		Padding(0, 1)

	draftedRowStyle = lipgloss.NewStyle().
		Foreground(draftedColor).
		Strikethrough(true).
		Padding(0, 1)

	// Search styles
	searchStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primaryColor).
		Padding(0, 1).
		Margin(1, 0)

	// Footer style
	footerStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		Padding(1, 2)

	// Position tab styles
	tabStyle = lipgloss.NewStyle().
		Padding(0, 2).
		Margin(0, 1)

	activeTabStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(primaryColor).
		Padding(0, 2).
		Margin(0, 1)
)