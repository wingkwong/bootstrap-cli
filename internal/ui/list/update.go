package list

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, w := lipgloss.NewStyle().GetFrameSize()
		vh = msg.Height - h
		vw = msg.Width - w
		b.SetSize(vw, vh)
		return b, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, selectListItemKey):
		case key.Matches(msg, selectListItemByNextKey):
			item, ok := b.list.SelectedItem().(Item)
			if ok {
				b.choice = item.title
			}
		}
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)
	return b, cmd
}
