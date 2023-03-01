package ui

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
	_inputs "github.com/wingkwong/bootstrap-cli/internal/ui/inputs"
	_list "github.com/wingkwong/bootstrap-cli/internal/ui/list"
)

type installFinishedMsg struct {
	err error
	out bytes.Buffer
}

func (b *Bubble) getTemplateList() _list.Bubble {
	if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
		return b.frontendTemplateList
	} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
		return b.backendTemplateList
	} else if b.frameworkType == _constants.DOCKER_FRAMEWORKS {
		return b.dockerTemplateList
	}
	return b.navigationList
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

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		vh = msg.Height - h - 20
		vw = msg.Width - w
		b.resizeAllBubbles(vw, vh)
		b.width = vw
	case tea.KeyMsg:
		templateList := b.getTemplateList()
		switch {
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		case key.Matches(msg, b.keys.Exit):
			if !templateList.IsFiltering() && !b.isInputting {
				return b, tea.Quit
			}
		case key.Matches(msg, b.keys.SelectListItemKey):
			if templateList.IsFiltering() {
				return b, nil
			}
			var item Item
			var ok bool
			if b.state == navigationState {
				item, ok := b.navigationList.List.SelectedItem().(Item)
				if ok {
					b.frameworkType = item.title
					b.state = templateState
				}
			} else if b.state == templateState {
				item, ok = templateList.List.SelectedItem().(Item)
				if ok {
					b.isInputting = true
					b.selectedInputs = b.getTemplateInputs(item.id)
					b.framework = item.name
					b.state = inputState
					b.deactivateAllBubbles()
					b.selectedInputs.SetActive(true)
					return b, cmd
				}
			} else if b.state == inputState {
				if b.selectedInputs.IsFinished() {
					b.isInputting = false
					item, ok = templateList.List.SelectedItem().(Item)
					if ok {
						var cmdArgs = item.commandArgs
						// set app name
						if b.selectedInputs.Inputs[0].Value() != "" {
							cmdArgs = strings.Replace(
								item.commandArgs,
								fmt.Sprintf("my-%s-app", item.name),
								b.selectedInputs.Inputs[0].Value(),
								1)
						}

						b.state = installState
						b.isInstalling = true
						var args = strings.Split(cmdArgs, " ")
						c := exec.Command(item.command, args...)

						// set directory
						if b.selectedInputs.Inputs[1].Value() != "" {
							c.Dir = b.selectedInputs.Inputs[1].Value()
						}

						var out bytes.Buffer
						c.Stdout = &out
						return b, tea.ExecProcess(c, func(err error) tea.Msg {
							return installFinishedMsg{err, out}
						})
					}
				}
			}
		}
	default:
		var cmd tea.Cmd
		b.spinner, cmd = b.spinner.Update(msg)
		return b, cmd
	}

	if b.state == navigationState {
		b.deactivateAllBubbles()
		b.navigationList.SetActive(true)
		b.navigationList.List, cmd = b.navigationList.List.Update(msg)
		cmds = append(cmds, cmd)
	} else if b.state == templateState {
		b.deactivateAllBubbles()
		if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
			b.frontendTemplateList.List, cmd = b.frontendTemplateList.List.Update(msg)
			cmds = append(cmds, cmd)
		} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
			b.backendTemplateList.List, cmd = b.backendTemplateList.List.Update(msg)
			cmds = append(cmds, cmd)
		} else if b.frameworkType == _constants.DOCKER_FRAMEWORKS {
			b.dockerTemplateList.List, cmd = b.dockerTemplateList.List.Update(msg)
			cmds = append(cmds, cmd)
		}
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
