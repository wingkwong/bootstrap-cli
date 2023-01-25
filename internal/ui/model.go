package ui

import "github.com/wingkwong/bootstrap-cli/internal/ui/list"

type sessionState int

const (
	idleState sessionState = iota
	showFrontendTemplateList
	showBackendTemplateList
)

type Bubble struct {
	navigationList       list.Bubble
	frontendTemplateList list.Bubble
	backendTemplateList  list.Bubble
	state                sessionState
}

func New() Bubble {
	// navigationList
	navigationListModel := list.New()
	// frontendTemplateList
	frontendTemplateListModel := list.New()
	// backendTemplateList
	backendTemplateListModel := list.New()

	return Bubble{
		navigationList:       navigationListModel,
		frontendTemplateList: frontendTemplateListModel,
		backendTemplateList:  backendTemplateListModel,
	}
}
