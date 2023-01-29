package list

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.SetWidth(msg.Width)
		return b, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, selectListItemKey):
			item, ok := b.list.SelectedItem().(Item)
			if ok {
				println(item.title)
				b.choice = item.title
			}
		}
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)
	return b, cmd
}
