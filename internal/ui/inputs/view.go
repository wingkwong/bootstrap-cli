package inputs

import (
	"fmt"
	"strings"
)

func (b Bubble) View() string {
	var sb strings.Builder

	for i := range b.inputs {
		sb.WriteString(b.inputs[i].View())
		if i < len(b.inputs)-1 {
			sb.WriteRune('\n')
		}
	}

	button := &blurredButton
	if b.focusIndex == len(b.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&sb, "\n\n%s\n\n", *button)

	return sb.String()
}
