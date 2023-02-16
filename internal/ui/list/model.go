package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	_constants "github.com/wingkwong/bootstrap-cli/internal/constants"
	_templates "github.com/wingkwong/bootstrap-cli/internal/templates"
	_inputs "github.com/wingkwong/bootstrap-cli/internal/ui/inputs"
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
	navigationList       list.Model
	frontendTemplateList list.Model
	backendTemplateList  list.Model
	dockerTemplateList   list.Model
	dockerTemplateInputs _inputs.Bubble
	frameworkType        string
	framework            string
	installOutput        []byte
	installError         error
	isInstalling         bool
	spinner              spinner.Model
	state                sessionState
}

func (b Bubble) Init() tea.Cmd {
	return b.spinner.Tick
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
	b.dockerTemplateList.SetSize(
		width,
		height,
	)
}

func New() Bubble {
	const defaultWidth = 20
	var navigationList list.Model
	var frontendTemplateList list.Model
	var backendTemplateList list.Model
	var dockerTemplateList list.Model
	var dockerTemplateInputs _inputs.Bubble
	var items []list.Item

	// navigation
	items = []list.Item{
		Item{title: _constants.FRONTEND_FRAMEWORKS, desc: _constants.FRONTEND_FRAMEWORKS_DESC},
		Item{title: _constants.BACKEND_FRAMEWORKS, desc: _constants.BACKEND_FRAMEWORKS_DESC},
		Item{title: _constants.KUBERNETES_FRAMEWORKS, desc: _constants.KUBERNETES_FRAMEWORKS_DESC},
		Item{title: _constants.DOCKER_FRAMEWORKS, desc: _constants.DOCKER_FRAMEWORKS_DESC},
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
			title:       "🔵 " + v.Title,
			name:        v.Title,
			desc:        v.Desc,
			command:     v.Command,
			commandArgs: v.CommandArgs,
		})
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
			title:       "🟠 " + v.Title,
			name:        v.Title,
			desc:        v.Desc,
			command:     v.Command,
			commandArgs: v.CommandArgs,
		})
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

	// docker

	items = []list.Item{}
	for _, v := range _templates.DOCKER_TEMPLATES {
		items = append(items, Item{
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

	dockerTemplateInputs = _inputs.NewMSSQLInputModel()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return Bubble{
		navigationList:       navigationList,
		frontendTemplateList: frontendTemplateList,
		backendTemplateList:  backendTemplateList,
		dockerTemplateList:   dockerTemplateList,
		dockerTemplateInputs: dockerTemplateInputs,
		spinner:              s,
		isInstalling:         false,
	}
}
