package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wingkwong/bootstrap-cli/internal/ui"
)

func Execute() {
	p := tea.NewProgram(ui.New())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running bootstrap-cli:", err)
		os.Exit(1)
	}
}
