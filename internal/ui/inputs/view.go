package inputs

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) View() string {
	var wrapper = lipgloss.NewStyle().Padding(0, 1)
	var sb strings.Builder

	for i := range b.Inputs {
		sb.WriteString(b.Inputs[i].View())
		if i < len(b.Inputs)-1 {
			sb.WriteRune('\n')
		}
	}

	button := &blurredButton
	if b.FocusIndex == len(b.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&sb, "\n\n%s\n\n", *button)

	return wrapper.Render(sb.String())
}
