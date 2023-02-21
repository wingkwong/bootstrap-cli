package inputs

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	if !b.Active {
		return b, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && b.FocusIndex == len(b.Inputs) {
				// return b, tea.Quit
				b.Finished = true
				return b, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				b.FocusIndex--
			} else {
				b.FocusIndex++
			}

			if b.FocusIndex > len(b.Inputs) {
				b.FocusIndex = 0
			} else if b.FocusIndex < 0 {
				b.FocusIndex = len(b.Inputs)
			}

			cmds := make([]tea.Cmd, len(b.Inputs))
			for i := 0; i <= len(b.Inputs)-1; i++ {
				if i == b.FocusIndex {
					// Set focused state
					cmds[i] = b.Inputs[i].Focus()
					b.Inputs[i].PromptStyle = focusedStyle
					b.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				b.Inputs[i].Blur()
				b.Inputs[i].PromptStyle = noStyle
				b.Inputs[i].TextStyle = noStyle
			}

			return b, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	return b, b.UpdateInputs(msg)
}

func (b *Bubble) UpdateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(b.Inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range b.Inputs {
		b.Inputs[i], cmds[i] = b.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
