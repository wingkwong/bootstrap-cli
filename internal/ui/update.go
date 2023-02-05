package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	b.l, cmd = b.l.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Padding(1, 2).GetFrameSize()
		b.l.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		}
	}

	return b, tea.Batch(cmds...)
}
