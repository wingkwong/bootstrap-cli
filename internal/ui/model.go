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
	isInputting            bool
	spinner                spinner.Model
	state                  sessionState
	keys                   KeyMap
	width                  int
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

func (b *Bubble) deactivateAllBubbles() {
	b.navigationList.SetActive(false)
	b.frontendTemplateList.SetActive(false)
	b.backendTemplateList.SetActive(false)
	b.dockerTemplateList.SetActive(false)
	for i := range b.frontendTemplateInputs {
		b.frontendTemplateInputs[i].SetActive(false)
	}

	for i := range b.backendTemplateInputs {
		b.backendTemplateInputs[i].SetActive(false)
	}

	for i := range b.dockerTemplateInputs {
		b.dockerTemplateInputs[i].SetActive(false)
	}
}

func (b *Bubble) resizeAllBubbles(vw int, vh int) {
	b.frontendTemplateList.SetSize(vw, vh)
	b.backendTemplateList.SetSize(vw, vh)
	b.dockerTemplateList.SetSize(vw, vh)
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
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy().Bold(false)

	navigationList = list.New(items, listDelegate, defaultWidth, listHeight)
	navigationList.SetShowTitle(false)

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
		frontendTemplateInputs = append(frontendTemplateInputs, _inputs.NewViteInputModel(v.Title))
	}

	listDelegate.Styles.SelectedTitle = frontendDelegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy().Bold(false)
	frontendTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	frontendTemplateList.SetShowTitle(false)

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
		backendTemplateInputs = append(backendTemplateInputs, _inputs.NewViteInputModel(v.Title))
	}
	listDelegate.Styles.SelectedTitle = backendDelegateStyle
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy().Bold(false)
	backendTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	backendTemplateList.SetShowTitle(false)

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
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy().Bold(false)
	dockerTemplateList = list.New(items, listDelegate, defaultWidth, listHeight)
	dockerTemplateList.SetShowTitle(false)

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
		keys:                   DefaultKeyMap(),
	}
}
