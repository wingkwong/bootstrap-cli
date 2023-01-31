package list

import "github.com/charmbracelet/lipgloss"

func (b Bubble) View() string {
	if b.frameworkChoice != "" {
		return b.frameworkChoice
	}
	return bubbleStyle.Render(lipgloss.JoinVertical(lipgloss.Top, b.list.View()))
}
