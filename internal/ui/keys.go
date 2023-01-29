package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit           key.Binding
	Exit           key.Binding
	SelectListItem key.Binding
	Back           key.Binding
	Next           key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),
		Exit: key.NewBinding(
			key.WithKeys("q"),
		),
		SelectListItem: key.NewBinding(
			key.WithKeys("enter"),
		),
		Back: key.NewBinding(
			key.WithKeys("left"),
		),
		Next: key.NewBinding(
			key.WithKeys("right"),
		),
	}
}
