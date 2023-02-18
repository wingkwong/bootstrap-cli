package list

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
	_inputs "github.com/wingkwong/bootstrap-cli/internal/ui/inputs"
)

type installFinishedMsg struct {
	err error
	out bytes.Buffer
}

func (b Bubble) getTemplateList() list.Model {
	if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
		return b.frontendTemplateList
	} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
		return b.backendTemplateList
	} else if b.frameworkType == _constants.DOCKER_FRAMEWORKS {
		return b.dockerTemplateList
	}
	return list.Model{}
}

func (b Bubble) getTemplateInputs(id int) _inputs.Bubble {
	if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
		return b.frontendTemplateInputs[id]
	} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
		return b.backendTemplateInputs[id]
	} else if b.frameworkType == _constants.DOCKER_FRAMEWORKS {
		return b.dockerTemplateInputs[id]
	}
	return _inputs.Bubble{}
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case installFinishedMsg:
		b.installOutput = msg.out.Bytes()
		b.isInstalling = false
		if msg.err != nil {
			b.installError = msg.err
			return b, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, w := lipgloss.NewStyle().GetFrameSize()
		vh = msg.Height - h
		vw = msg.Width - w
		b.SetSize(vw, vh)
		return b, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, selectListItemKey):
			var templateList list.Model
			var item Item
			var ok bool
			if b.state == navigationState {
				item, ok := b.navigationList.SelectedItem().(Item)
				if ok {
					b.frameworkType = item.title
					b.state = templateState
					// TODO: refactor
					// point to the first item
					b.frontendTemplateList.ResetSelected()
					b.backendTemplateList.ResetSelected()
					b.dockerTemplateList.ResetSelected()
				}
			} else if b.state == templateState {
				templateList = b.getTemplateList()
				item, ok = templateList.SelectedItem().(Item)
				if ok {
					b.selectedInputs = b.getTemplateInputs(item.id)
					b.state = inputState
				}
			} else if b.state == inputState {
				if b.selectedInputs.IsFinished() {
					// TODO: get inputs val
					templateList = b.getTemplateList()
					item, ok = templateList.SelectedItem().(Item)
					if ok {
						b.state = installState
						b.framework = item.name
						b.isInstalling = true
						var args = strings.Split(item.commandArgs, " ")
						c := exec.Command(item.command, args...)
						var out bytes.Buffer
						c.Stdout = &out
						return b, tea.ExecProcess(c, func(err error) tea.Msg {
							return installFinishedMsg{err, out}
						})
					}
				}
			} else {
				return b, tea.Quit
			}
		}
	default:
		var cmd tea.Cmd
		b.spinner, cmd = b.spinner.Update(msg)
		return b, cmd
	}

	b.navigationList, cmd = b.navigationList.Update(msg)
	cmds = append(cmds, cmd)

	b.frontendTemplateList, cmd = b.frontendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	b.backendTemplateList, cmd = b.backendTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	b.dockerTemplateList, cmd = b.dockerTemplateList.Update(msg)
	cmds = append(cmds, cmd)

	b.selectedInputs, cmd = b.selectedInputs.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
