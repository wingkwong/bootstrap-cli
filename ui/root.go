package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	list "github.com/wingkwong/bootstrap-cli/ui/list"
	tabs "github.com/wingkwong/bootstrap-cli/ui/tabs"
)

var (
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

type model struct {
	Tabs tea.Model
	List tea.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.Tabs.Init(),
		m.List.Init(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			// send the quit command
			cmds = append(cmds, tea.Quit)
		}
	}

	// handle updates for m.Tabs
	m.Tabs, cmd = m.Tabs.Update(msg)
	cmds = append(cmds, cmd)

	// handle updates for m.List
	m.List, cmd = m.List.Update(msg)
	cmds = append(cmds, cmd)

	// return updated model and any new commands
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	doc := strings.Builder{}
	doc.WriteString(m.Tabs.View())
	doc.WriteString("\n")
	doc.WriteString(m.List.View())
	return docStyle.Render(doc.String())
}

func New() tea.Model {
	tabs := tabs.New()
	list := list.New()
	return model{Tabs: tabs, List: list}
}
