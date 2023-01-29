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

type Bubble struct {
	list   list.Model
	choice string
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) GetChoice() string {
	return b.choice
}

func (b Bubble) GetSelectedItem() Item {
	item, ok := b.list.SelectedItem().(Item)
	if ok {
		return item
	}
	return Item{}
}

func (b *Bubble) SetSize(width, height int) {
	b.list.SetSize(
		width,
		height,
	)
}

func New(listType string) Bubble {
	const defaultWidth = 20
	var l list.Model
	var items []list.Item
	switch listType {
	case _common.NAVIGATION_TEMPLATE_LIST:
		items = []list.Item{
			Item{title: _common.FRONTEND_FRAMEWORKS, desc: _common.FRONTEND_FRAMEWORKS_DESC},
			Item{title: _common.BACKEND_FRAMEWORKS, desc: _common.BACKEND_FRAMEWORKS_DESC},
			Item{title: _common.KUBERNETES_FRAMEWORKS, desc: _common.KUBERNETES_FRAMEWORKS_DESC},
			Item{title: _common.DOCKER_FRAMEWORKS, desc: _common.DOCKER_FRAMEWORKS_DESC},
		}
		l = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
		l.Title = "Select frameworks"
	case _common.FRONTEND_TEMPLATE_LIST:
		items = []list.Item{
			Item{title: "vue", desc: "Generate Vue.js App Template"},
			Item{title: "vue-ts", desc: "Generate Vue.js App Template in TypeScript"},
			Item{title: "react", desc: "Generate React App Template"},
			Item{title: "react-ts", desc: "Generate React App Template in TypeScript"},
			Item{title: "next", desc: "Generate Next.js App Template"},
			Item{title: "next-ts", desc: "Generate Next.js App Template in TypeScript"},
			Item{title: "vanilla", desc: "Generate Vanilla.js App Template"},
			Item{title: "vanilla-ts", desc: "Generate Vanilla.js App Template in TypeScript"},
			Item{title: "gatsby", desc: "Generate Gatsby App Template in TypeScript"},
			Item{title: "gatsby-ts", desc: "Generate Gatsby App Template in TypeScript"},
		}
		l = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
		l.Title = "Here's the available templates for Frontend Frameworks."
	case _common.BACKEND_TEMPLATE_LIST:
		items = []list.Item{
			Item{title: "express", desc: "Generate Express.js App Template"}}
		l = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
		l.Title = "Here's the available Templates for Backend Frameworks."
	}

	l.Styles.Title = titleStyle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// for newly created lists only
	// since WindowSizeMsg only gets sent on program start
	if vw != 0 && vh != 0 {
		l.SetWidth(vw)
		l.SetHeight(vh)
	}

	return Bubble{list: l}
}
