package list

import "github.com/charmbracelet/lipgloss"

func (b Bubble) View() string {
	return bubbleStyle.Render(lipgloss.JoinVertical(lipgloss.Top, b.list.View()))
}
