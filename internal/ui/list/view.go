package list

import (
	"github.com/charmbracelet/lipgloss"
	_common "github.com/wingkwong/bootstrap-cli/internal/common"
)

func (b Bubble) View() string {
	var view string
	if b.state == navigationState {
		view = b.navigationList.View()
	} else if b.state == templateState {
		if b.frameworkType == _common.FRONTEND_FRAMEWORKS {
			view = b.frontendTemplateList.View()
		} else if b.frameworkType == _common.BACKEND_FRAMEWORKS {
			view = b.backendTemplateList.View()
		}
	} else if b.state == installState {
		view = string(b.output)
	}
	return bubbleStyle.Render(lipgloss.JoinVertical(lipgloss.Top, view))
}
