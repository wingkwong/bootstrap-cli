package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
)

func (b Bubble) View() string {
	// header
	headerAppName := headerAppNameStyle.Render(_constants.APP_NAME)
	headerDesc := headerAppDescStyle.Render(_constants.APP_DESC)
	mxWidth := lipgloss.Width(headerDesc)
	if b.width > mxWidth {
		mxWidth = b.width
	}
	wrapper := lipgloss.NewStyle().Width(mxWidth)
	headerUrl := headerUrlStyle.Copy().Width(b.width - lipgloss.Width(headerAppName) - 5).Align(lipgloss.Right).Render(_constants.APP_REPO_URL)
	var view = style.Render(lipgloss.JoinVertical(lipgloss.Top,
		wrapper.Render(style.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				headerAppName,
				headerUrl))),
		wrapper.Render(headerDesc)))

	view += "\n\n"
	// content
	if b.state == navigationState {
		view += b.navigationList.View()
	} else if b.state == templateState {
		if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
			view += b.frontendTemplateList.View()
		} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
			view += b.backendTemplateList.View()
		} else if b.frameworkType == _constants.DOCKER_FRAMEWORKS {
			view += b.dockerTemplateList.View()
		}
	} else if b.state == installState {
		if b.installError != nil {
			view += "Error: " + b.installError.Error() + "\n"
		} else if b.isInstalling {
			view += fmt.Sprintf("%s Installing ... ", b.spinner.View())
		} else if b.installOutput != nil {
			view += fmt.Sprintf("%s \n 🚀 %s %s", b.installOutput, b.framework, "has been installed. Press `Enter` to quit. ")
		}
	} else if b.state == inputState {
		view += fmt.Sprintf("Please enter the following info before installing %s.\n\n", b.framework)
		view += b.selectedInputs.View()
	}
	return bubbleStyle.Render(view)
}
