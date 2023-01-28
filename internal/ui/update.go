package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	_common "github.com/wingkwong/bootstrap-cli/internal/common"
)

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	selectedItem := b.navigationList.GetSelectedItem()

	b.navigationList, cmd = b.navigationList.Update(msg)
	cmds = append(cmds, cmd)

	switch selectedItem.Title() {
	case _common.FRONTEND_FRAMEWORKS:
		b.state = frontendTemplateListState
		// show frontend
	case _common.BACKEND_FRAMEWORKS:
		b.state = backendTemplateListState
	case _common.KUBERNETES_FRAMEWORKS:
		// show k8s
	case _common.DOCKER_FRAMEWORKS:
		// show docker
	}

	b.frontendTemplateList, cmd = b.frontendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	b.backendTemplateList, cmd = b.backendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
