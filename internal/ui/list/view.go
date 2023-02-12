package list

import (
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
)

func (b Bubble) View() string {
	var view string
	if b.state == navigationState {
		view = b.navigationList.View()
	} else if b.state == templateState {
		if b.frameworkType == _constants.FRONTEND_FRAMEWORKS {
			view = b.frontendTemplateList.View()
		} else if b.frameworkType == _constants.BACKEND_FRAMEWORKS {
			view = b.backendTemplateList.View()
		}
	} else if b.state == installState {
		if b.installError != nil {
			view = "Error: " + b.installError.Error() + "\n"
		} else if b.installProgress.Percent() == 1.0 {
			view = b.installProgress.View() + "\n\n" + string(b.installOutput)
		} else {
			view = b.installProgress.View() + "\n\n" + "Installing ..."
		}
	}
	return bubbleStyle.Render(view)
}
