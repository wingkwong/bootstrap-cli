package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	bubbleStyle     = lipgloss.NewStyle().Margin(1, 1)
	titleStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#0074D9")).Padding(0, 1)
	delegateStyle   = list.NewDefaultDelegate().Styles.SelectedTitle.Foreground(lipgloss.Color("#ADD8E6")).BorderLeftForeground(lipgloss.Color("#FFFFE0"))
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	spinnerStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	listHeight      = 20
)
