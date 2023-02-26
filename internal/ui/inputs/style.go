package inputs

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("167"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	submitButton  = lipgloss.NewStyle().Width(10).Padding(1, 0).Align(lipgloss.Center)
	focusedButton = submitButton.Foreground(lipgloss.Color("255")).Background(lipgloss.Color("167")).Render("Submit")
	blurredButton = submitButton.Foreground(lipgloss.Color("167")).Background(lipgloss.Color("255")).Render("Submit")
)
