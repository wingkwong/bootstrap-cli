package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Bubble) Init() tea.Cmd {
	return m.navigationList.Init()
}
