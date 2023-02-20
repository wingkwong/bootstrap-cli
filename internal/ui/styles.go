package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	bubbleStyle           = lipgloss.NewStyle().Margin(1, 1)
	titleStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#0074D9")).Padding(0, 1)
	frontendTitleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#0074D9")).Padding(0, 1)
	backendTitleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#0074D9")).Padding(0, 1)
	dockerTitleStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#0074D9")).Padding(0, 1)
	delegateStyle         = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#ADD8E6")).BorderLeftForeground(lipgloss.Color("#FFFFE0"))
	frontendDelegateStyle = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#ADD8E6")).BorderLeftForeground(lipgloss.Color("#FFFFE0"))
	backendDelegateStyle  = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#FFA500")).BorderLeftForeground(lipgloss.Color("#FFA500"))
	dockerDelegateStyle   = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#FDFD96")).BorderLeftForeground(lipgloss.Color("#FDFD96"))
	quitTextStyle         = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	spinnerStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	listHeight            = 20
)
