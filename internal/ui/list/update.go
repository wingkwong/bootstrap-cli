package list

import (
	"bytes"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type installFinishedMsg struct {
	err error
	out bytes.Buffer
}

func runInstall() tea.Cmd {
	c := exec.Command("npx", "create-react-app", "my-app")
	var out bytes.Buffer
	c.Stdout = &out
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return installFinishedMsg{err, out}
	})
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, w := lipgloss.NewStyle().GetFrameSize()
		vh = msg.Height - h
		vw = msg.Width - w
		b.SetSize(vw, vh)
		return b, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, selectListItemKey), key.Matches(msg, selectListItemByNextKey):
			item, ok := b.list.SelectedItem().(Item)
			if ok {
				if b.frameworkTypeChoice == "" {
					b.frameworkTypeChoice = item.title
				} else if b.frameworkChoice == "" {
					b.frameworkChoice = item.title
					return b, runInstall()
				}
			}
		}
	case installFinishedMsg:
		b.output = msg.out.Bytes()
		if msg.err != nil {
			// b.err = msg.err
			return b, tea.Quit
		}
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)
	return b, cmd
}
