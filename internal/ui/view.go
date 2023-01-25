package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Bubble) View() string {
	navigationList := m.navigationList.View()
	templateList := ""

	switch m.state {
	case frontendTemplateListState:
		templateList = m.frontendTemplateList.View()
	case backendTemplateListState:
		templateList = m.backendTemplateList.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, navigationList, templateList),
		// m.statusbar.View(),
	)
}
