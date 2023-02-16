package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return b, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			b.cursorMode++
			if b.cursorMode > textinput.CursorHide {
				b.cursorMode = textinput.CursorBlink
			}
			cmds := make([]tea.Cmd, len(b.inputs))
			for i := range b.inputs {
				cmds[i] = b.inputs[i].SetCursorMode(b.cursorMode)
			}
			return b, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && b.focusIndex == len(b.inputs) {
				// return b, tea.Quit
				b.finished = true
				return b, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				b.focusIndex--
			} else {
				b.focusIndex++
			}

			if b.focusIndex > len(b.inputs) {
				b.focusIndex = 0
			} else if b.focusIndex < 0 {
				b.focusIndex = len(b.inputs)
			}

			cmds := make([]tea.Cmd, len(b.inputs))
			for i := 0; i <= len(b.inputs)-1; i++ {
				if i == b.focusIndex {
					// Set focused state
					cmds[i] = b.inputs[i].Focus()
					b.inputs[i].PromptStyle = focusedStyle
					b.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				b.inputs[i].Blur()
				b.inputs[i].PromptStyle = noStyle
				b.inputs[i].TextStyle = noStyle
			}

			return b, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := b.UpdateInputs(msg)

	return b, cmd
}

func (b *Bubble) UpdateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(b.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range b.inputs {
		b.inputs[i], cmds[i] = b.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
