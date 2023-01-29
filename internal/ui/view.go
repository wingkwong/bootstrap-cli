package ui

func (b Bubble) View() string {
	switch b.state {
	case frontendTemplateListState:
		return b.frontendTemplateList.View()
	case backendTemplateListState:
		return b.backendTemplateList.View()
	default:
		return b.navigationList.View()
	}
}
