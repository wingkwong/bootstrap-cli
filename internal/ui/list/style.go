package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	bubbleStyle     = lipgloss.NewStyle()
	titleStyle      = lipgloss.NewStyle()
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	listHeight      = 20
)
