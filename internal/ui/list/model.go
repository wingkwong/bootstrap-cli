package list

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	_common "github.com/wingkwong/bootstrap-cli/internal/common"
)

var (
	vh int
	vw int
)

type sessionState int

const (
	navigationState sessionState = iota
	templateState
	installState
)

type Bubble struct {
	navigationList       list.Model
	frontendTemplateList list.Model
	backendTemplateList  list.Model
	frameworkType        string
	framework            string
	output               []byte
	installError         error
	state                sessionState
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) GetFrameworkType() string {
	return b.frameworkType
}

func (b Bubble) GetFramework() string {
	return b.framework
}

func (b Bubble) GetSelecteNavigationItem() Item {
	item, ok := b.navigationList.SelectedItem().(Item)
	if ok {
		return item
	}
	return Item{}
}

func (b *Bubble) SetSize(width, height int) {
	b.navigationList.SetSize(
		width,
		height,
	)
	b.frontendTemplateList.SetSize(
		width,
		height,
	)
	b.backendTemplateList.SetSize(
		width,
		height,
	)
}

func New() Bubble {
	const defaultWidth = 20
	var navigationList list.Model
	var frontendTemplateList list.Model
	var backendTemplateList list.Model
	var items []list.Item

	// navigation
	items = []list.Item{
		Item{title: _common.FRONTEND_FRAMEWORKS, desc: _common.FRONTEND_FRAMEWORKS_DESC},
		Item{title: _common.BACKEND_FRAMEWORKS, desc: _common.BACKEND_FRAMEWORKS_DESC},
		Item{title: _common.KUBERNETES_FRAMEWORKS, desc: _common.KUBERNETES_FRAMEWORKS_DESC},
		Item{title: _common.DOCKER_FRAMEWORKS, desc: _common.DOCKER_FRAMEWORKS_DESC},
	}
	navigationList = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	navigationList.Title = "Select frameworks"
	navigationList.Styles.Title = titleStyle
	navigationList.SetShowStatusBar(false)
	navigationList.SetFilteringEnabled(true)
	navigationList.Styles.PaginationStyle = paginationStyle
	navigationList.Styles.HelpStyle = helpStyle

	// frontend
	items = []list.Item{
		Item{title: _common.FRONTEND_VUE, desc: _common.FRONTEND_VUE_DESC, command: _common.FRONTEND_VUE_CMD},
		Item{title: _common.FRONTEND_VUE_TS, desc: _common.FRONTEND_VUE_TS_DESC, command: _common.FRONTEND_VUE_TS_CMD},
		Item{title: _common.FRONTEND_REACT, desc: _common.FRONTEND_REACT_DESC, command: _common.FRONTEND_REACT_CMD},
		Item{title: _common.FRONTEND_REACT_TS, desc: _common.FRONTEND_REACT_TS_DESC, command: _common.FRONTEND_REACT_TS_CMD},
		Item{title: _common.FRONTEND_NEXT, desc: _common.FRONTEND_NEXT_DESC, command: _common.FRONTEND_NEXT_CMD},
		Item{title: _common.FRONTEND_NEXT_TS, desc: _common.FRONTEND_NEXT_TS_DESC, command: _common.FRONTEND_NEXT_TS_CMD},
		Item{title: _common.FRONTEND_VANILLA, desc: _common.FRONTEND_VANILLA_DESC, command: _common.FRONTEND_VANILLA_CMD},
		Item{title: _common.FRONTEND_VANILLA_TS, desc: _common.FRONTEND_VANILLA_TS_DESC, command: _common.FRONTEND_VANILLA_TS_CMD},
		Item{title: _common.FRONTEND_GATSBY, desc: _common.FRONTEND_GATSBY_DESC, command: _common.FRONTEND_GATSBY_CMD},
		Item{title: _common.FRONTEND_GATSBY_TS, desc: _common.FRONTEND_GATSBY_TS_DESC, command: _common.FRONTEND_GATSBY_TS_CMD},
	}
	frontendTemplateList = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	frontendTemplateList.Title = "Here's the available templates for Frontend Frameworks."

	frontendTemplateList.Styles.Title = titleStyle
	frontendTemplateList.SetShowStatusBar(false)
	frontendTemplateList.SetFilteringEnabled(true)
	frontendTemplateList.Styles.PaginationStyle = paginationStyle
	frontendTemplateList.Styles.HelpStyle = helpStyle

	// backend
	items = []list.Item{
		Item{title: "express", desc: "Generate Express.js App Template", command: ""}}
	backendTemplateList = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	backendTemplateList.Title = "Here's the available Templates for Backend Frameworks."

	backendTemplateList.Styles.Title = titleStyle
	backendTemplateList.SetShowStatusBar(false)
	backendTemplateList.SetFilteringEnabled(true)
	backendTemplateList.Styles.PaginationStyle = paginationStyle
	backendTemplateList.Styles.HelpStyle = helpStyle

	return Bubble{navigationList: navigationList, frontendTemplateList: frontendTemplateList, backendTemplateList: backendTemplateList}
}
