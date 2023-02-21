package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (b Bubble) Init() tea.Cmd {
	return b.spinner.Tick
}
