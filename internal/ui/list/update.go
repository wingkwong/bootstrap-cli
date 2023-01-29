package list

import tea "github.com/charmbracelet/bubbletea"

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.SetWidth(msg.Width)
		return b, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return b, tea.Quit
		case "enter":
			item, ok := b.list.SelectedItem().(Item)
			if ok {
				println(item.title)
				b.choice = item.title
			}
			return b, nil
		}
	}

	var cmd tea.Cmd
	b.list, cmd = b.list.Update(msg)

	return b, cmd
}
