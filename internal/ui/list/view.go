package list

import (
	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) View() string {
	if b.frameworkChoice != "" {
		// return string(b.output)
		print(b.output)
		return ""
	}

	return bubbleStyle.Render(lipgloss.JoinVertical(lipgloss.Top, b.list.View()))
}
