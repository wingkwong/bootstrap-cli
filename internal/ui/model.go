package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
	_templates "github.com/wingkwong/bootstrap-cli/internal/templates"
	_inputs "github.com/wingkwong/bootstrap-cli/internal/ui/inputs"
	_list "github.com/wingkwong/bootstrap-cli/internal/ui/list"
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
	inputState
)

type Bubble struct {
	navigationList         _list.Bubble
	frontendTemplateList   _list.Bubble
	backendTemplateList    _list.Bubble
	dockerTemplateList     _list.Bubble
	frontendTemplateInputs []_inputs.Bubble
	backendTemplateInputs  []_inputs.Bubble
	dockerTemplateInputs   []_inputs.Bubble
	selectedInputs         _inputs.Bubble
	frameworkType          string
	framework              string
	installOutput          []byte
	installError           error
	isInstalling           bool
	spinner                spinner.Model
	state                  sessionState
}

func (b Bubble) GetFrameworkType() string {
	return b.frameworkType
}

func (b Bubble) GetFramework() string {
	return b.framework
}

func (b Bubble) GetSelecteNavigationItem() Item {
	item, ok := b.navigationList.List.SelectedItem().(Item)
	if ok {
		return item
	}
	return Item{}
}

func New() Bubble {
	const defaultWidth = 40
	var navigationList list.Model
	var frontendTemplateList list.Model
	var backendTemplateList list.Model
	var dockerTemplateList list.Model
	var frontendTemplateInputs []_inputs.Bubble
	var backendTemplateInputs []_inputs.Bubble
	var dockerTemplateInputs []_inputs.Bubble
	var items []list.Item

	// navigation
	items = []list.Item{
		Item{title: _constants.FRONTEND_FRAMEWORKS, desc: _constants.FRONTEND_FRAMEWORKS_DESC},
		Item{title: _constants.BACKEND_FRAMEWORKS, desc: _constants.BACKEND_FRAMEWORKS_DESC},
		// TODO: hide at this moment
		// Item{title: _constants.KUBERNETES_FRAMEWORKS, desc: _constants.KUBERNETES_FRAMEWORKS_DESC},
		// Item{title: _constants.DOCKER_FRAMEWORKS, desc: _constants.DOCKER_FRAMEWORKS_DESC},
	}
	listDelegate := list.NewDefaultDelegate()
	listDelegate.Styles.SelectedTitle = delegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()

	navigationList = list.New(items, listDelegate, defaultWidth, listHeight)
	navigationList.Title = _constants.NAVIGATION_TEMPLATE_LIST_TITLE
	navigationList.Styles.Title = titleStyle
	navigationList.SetShowStatusBar(false)
	navigationList.SetFilteringEnabled(true)
	navigationList.Styles.PaginationStyle = paginationStyle
	navigationList.Styles.HelpStyle = helpStyle

	// frontend
	items = []list.Item{}
	for _, v := range _templates.FRONTEND_TEMPLATES {
		items = append(items, Item{
			id:          v.Id,
			title:       "🔵 " + v.Title,
			name:        v.Title,
			desc:        v.Desc,
			command:     v.Command,
			commandArgs: v.CommandArgs,
		})
		frontendTemplateInputs = append(frontendTemplateInputs, _inputs.NewViteInputModel())
	}

	listDelegate.Styles.SelectedTitle = frontendDelegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()
	frontendTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	frontendTemplateList.Title = _constants.FRONTEND_TEMPLATE_LIST_TITLE

	frontendTemplateList.Styles.Title = frontendTitleStyle
	frontendTemplateList.SetShowStatusBar(false)
	frontendTemplateList.SetFilteringEnabled(true)
	frontendTemplateList.Styles.PaginationStyle = paginationStyle
	frontendTemplateList.Styles.HelpStyle = helpStyle

	// backend
	items = []list.Item{}
	for _, v := range _templates.BACKEND_TEMPLATES {
		items = append(items, Item{
			id:          v.Id,
			title:       "🟠 " + v.Title,
			name:        v.Title,
			desc:        v.Desc,
			command:     v.Command,
			commandArgs: v.CommandArgs,
		})
		backendTemplateInputs = append(backendTemplateInputs, _inputs.NewViteInputModel())
	}
	listDelegate.Styles.SelectedTitle = backendDelegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()
	backendTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	backendTemplateList.Title = _constants.BACKEND_TEMPLATE_LIST_TITLE

	backendTemplateList.Styles.Title = backendTitleStyle
	backendTemplateList.SetShowStatusBar(false)
	backendTemplateList.SetFilteringEnabled(true)
	backendTemplateList.Styles.PaginationStyle = paginationStyle
	backendTemplateList.Styles.HelpStyle = helpStyle

	// docker (TODO)
	items = []list.Item{}
	for _, v := range _templates.DOCKER_TEMPLATES {
		items = append(items, Item{
			id:          v.Id,
			title:       "🟡 " + v.Title,
			name:        v.Title,
			desc:        v.Desc,
			command:     v.Command,
			commandArgs: v.CommandArgs,
		})
	}
	listDelegate.Styles.SelectedTitle = dockerDelegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()
	dockerTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	dockerTemplateList.Title = _constants.DOCKER_TEMPLATE_LIST_TITLE

	dockerTemplateList.Styles.Title = dockerTitleStyle
	dockerTemplateList.SetShowStatusBar(false)
	dockerTemplateList.SetFilteringEnabled(true)
	dockerTemplateList.Styles.PaginationStyle = paginationStyle
	dockerTemplateList.Styles.HelpStyle = helpStyle

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return Bubble{
		navigationList:         _list.New(navigationList, false),
		frontendTemplateList:   _list.New(frontendTemplateList, false),
		backendTemplateList:    _list.New(backendTemplateList, false),
		dockerTemplateList:     _list.New(dockerTemplateList, false),
		frontendTemplateInputs: frontendTemplateInputs,
		backendTemplateInputs:  backendTemplateInputs,
		dockerTemplateInputs:   dockerTemplateInputs,
		spinner:                s,
		isInstalling:           false,
	}
}
