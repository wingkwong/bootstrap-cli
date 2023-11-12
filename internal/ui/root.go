package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Execute() {
	p := tea.NewProgram(New())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running bootstrap-cli:", err)
		os.Exit(1)
	}
}
