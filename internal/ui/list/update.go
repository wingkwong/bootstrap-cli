package list

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
)

type installFinishedMsg struct {
	err error
	out bytes.Buffer
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case installFinishedMsg:
		b.installOutput = msg.out.Bytes()
		if msg.err != nil {
			b.installError = msg.err
			// return b, tea.Quit
		} else {
			return b, tea.Quit
		}
		return b, nil
	case tea.WindowSizeMsg:
		h, w := lipgloss.NewStyle().GetFrameSize()
		vh = msg.Height - h
		vw = msg.Width - w
		b.SetSize(vw, vh)
		return b, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, selectListItemKey):
			if b.state == navigationState {
				item, ok := b.navigationList.SelectedItem().(Item)
				if ok {
					b.frameworkType = item.title
					b.state = templateState
				}
			} else if b.state == templateState {
				var item Item
				var ok bool
				if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
					item, ok = b.frontendTemplateList.SelectedItem().(Item)
				} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
					item, ok = b.backendTemplateList.SelectedItem().(Item)
				}

				if ok {
					b.state = installState
					var args = strings.Split(item.commandArgs, " ")
					c := exec.Command(item.command, args...)
					var out bytes.Buffer
					c.Stdout = &out
					cmds = append(cmds, tea.ExecProcess(c, func(err error) tea.Msg {
						return installFinishedMsg{err, out}
					}))
				}
			}
		}
	}

	b.navigationList, cmd = b.navigationList.Update(msg)
	cmds = append(cmds, cmd)

	b.frontendTemplateList, cmd = b.frontendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	b.backendTemplateList, cmd = b.backendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
