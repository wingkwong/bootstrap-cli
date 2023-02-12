package list

import (
	"bytes"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	progress "github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
)

type installFinishedMsg struct {
	err error
	out bytes.Buffer
}

type tickMsg time.Time

const (
	padding  = 2
	maxWidth = 80
)

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tickMsg:
		if b.installProgress.Percent() == 1.0 {
			return b, nil
		}
		cmd := b.installProgress.IncrPercent(0.25)
		return b, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := b.installProgress.Update(msg)
		b.installProgress = progressModel.(progress.Model)
		return b, cmd
	case installFinishedMsg:
		b.installOutput = msg.out.Bytes()
		if msg.err != nil {
			b.installError = msg.err
			return b, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, w := lipgloss.NewStyle().GetFrameSize()
		vh = msg.Height - h
		vw = msg.Width - w
		b.SetSize(vw, vh)
		// installProgress
		b.installProgress.Width = msg.Width - padding*2 - 4
		if b.installProgress.Width > maxWidth {
			b.installProgress.Width = maxWidth
		}
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
					b.framework = item.title
					var args = strings.Split(item.commandArgs, " ")
					c := exec.Command(item.command, args...)
					var out bytes.Buffer
					c.Stdout = &out
					return b, tea.ExecProcess(c, func(err error) tea.Msg {
						return installFinishedMsg{err, out}
					})
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
