package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Bubble) View() string {
	navigationList := m.navigationList.View()
	templateList := m.frontendTemplateList.View()

	switch m.state {
	case showBackendTemplateList:
		templateList = m.backendTemplateList.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, navigationList, templateList),
		// m.statusbar.View(),
	)
}
