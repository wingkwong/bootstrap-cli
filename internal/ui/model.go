package ui

import (
	_common "github.com/wingkwong/bootstrap-cli/internal/common"
	"github.com/wingkwong/bootstrap-cli/internal/ui/list"
)

type sessionState int

const (
	idleState sessionState = iota
	navigationListState
	frontendTemplateListState
	backendTemplateListState
)

type Bubble struct {
	navigationList       list.Bubble
	frontendTemplateList list.Bubble
	backendTemplateList  list.Bubble
	state                sessionState
}

func New() Bubble {
	// navigationList
	navigationListModel := list.New(_common.NAVIGATION_TEMPLATE_LIST)
	// frontendTemplateList
	frontendTemplateListModel := list.New(_common.FRONTEND_TEMPLATE_LIST)
	// backendTemplateList
	backendTemplateListModel := list.New(_common.BACKEND_TEMPLATE_LIST)

	return Bubble{
		navigationList:       navigationListModel,
		frontendTemplateList: frontendTemplateListModel,
		backendTemplateList:  backendTemplateListModel,
	}
}
