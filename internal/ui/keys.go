package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit           key.Binding
	Exit           key.Binding
	SelectListItem key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),
		Exit: key.NewBinding(
			key.WithKeys("ctrl+q"),
		),
		SelectListItem: key.NewBinding(
			key.WithKeys("enter"),
		),
	}
}
