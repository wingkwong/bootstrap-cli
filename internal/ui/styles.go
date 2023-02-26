package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	style                 = lipgloss.NewStyle()
	headerAppNameStyle    = style.Background(lipgloss.Color("167")).Bold(true).Margin(0, 0).Padding(1, 1)
	headerUrlStyle        = style.MarginLeft(1).Foreground(lipgloss.Color("167")).Background(lipgloss.Color("#FFFFFF")).Margin(0, 0).Padding(1, 1)
	headerAppDescStyle    = style.Foreground(lipgloss.Color("243")).Padding(1, 1, 0, 1)
	bubbleStyle           = style.Margin(1, 1)
	titleStyle            = style.Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1).Align(lipgloss.Center)
	frontendTitleStyle    = style.Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	backendTitleStyle     = style.Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	dockerTitleStyle      = style.Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	delegateStyle         = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#ADD8E6")).BorderLeftForeground(lipgloss.Color("#FFFFE0")).Bold(true)
	frontendDelegateStyle = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#ADD8E6")).BorderLeftForeground(lipgloss.Color("#FFFFE0")).Bold(true)
	backendDelegateStyle  = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#FFA500")).BorderLeftForeground(lipgloss.Color("#FFA500")).Bold(true)
	dockerDelegateStyle   = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#FDFD96")).BorderLeftForeground(lipgloss.Color("#FDFD96")).Bold(true)
	quitTextStyle         = style.Margin(1, 0, 2, 4)
	spinnerStyle          = style.Foreground(lipgloss.Color("205"))
	listHeight            = 20
)
