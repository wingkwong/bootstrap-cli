package list

import "github.com/charmbracelet/bubbles/key"

var (
	selectListItemKey       = key.NewBinding(key.WithKeys("enter"))
	selectListItemByNextKey = key.NewBinding(key.WithKeys("right"))
)
