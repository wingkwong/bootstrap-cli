package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_common "github.com/wingkwong/bootstrap-cli/internal/common"
	"github.com/wingkwong/bootstrap-cli/internal/ui/list"
)

func (b *Bubble) switchList(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	choice := b.navigationList.GetChoice()

	switch choice {
	case _common.FRONTEND_FRAMEWORKS:
		b.state = frontendTemplateListState
	case _common.BACKEND_FRAMEWORKS:
		b.state = backendTemplateListState
	case _common.KUBERNETES_FRAMEWORKS:
		// show k8s
	case _common.DOCKER_FRAMEWORKS:
		// show docker
	default:
		b.state = idleState
	}
	return cmds
}

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	b.navigationList, cmd = b.navigationList.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Padding(1, 2).GetFrameSize()
		b.navigationList.SetSize(msg.Width-h, msg.Height-v)
		b.frontendTemplateList.SetSize(msg.Width-h, msg.Height-v)
		b.backendTemplateList.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		case key.Matches(msg, b.keys.Exit):
			// TODO
		case key.Matches(msg, b.keys.SelectListItem):
			cmds = append(cmds, tea.Batch(b.switchList(msg)...))
		case key.Matches(msg, b.keys.Back):
			if b.state != idleState {
				b.state = idleState
				b.navigationList = list.New(_common.NAVIGATION_TEMPLATE_LIST)
				b.frontendTemplateList = list.New(_common.FRONTEND_TEMPLATE_LIST)
				b.backendTemplateList = list.New(_common.BACKEND_TEMPLATE_LIST)
			}
		}
	}

	b.frontendTemplateList, cmd = b.frontendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	b.backendTemplateList, cmd = b.backendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
