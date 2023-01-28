package list

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_common "github.com/wingkwong/bootstrap-cli/internal/common"
)

const listHeight = 20

var (
	titleStyle        = lipgloss.NewStyle()
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("270"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Bubble struct {
	list list.Model
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) GetSelectedItem() Item {
	item, ok := b.list.SelectedItem().(Item)
	if ok {
		return item
	}
	return Item{}
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.SetWidth(msg.Width)
		return b, nil
	}

	var cmd tea.Cmd
	b.list, cmd = b.list.Update(msg)

	return b, cmd
}

func (b Bubble) View() string {
	return "\n" + b.list.View()
}

func New(listType string) Bubble {
	const defaultWidth = 20
	var l list.Model
	switch listType {
	case _common.NAVIGATION_TEMPLATE_LIST:
		items := []list.Item{
			Item{title: _common.FRONTEND_FRAMEWORKS, desc: _common.FRONTEND_FRAMEWORKS_DESC},
			Item{title: _common.BACKEND_FRAMEWORKS, desc: _common.BACKEND_FRAMEWORKS_DESC},
			Item{title: _common.KUBERNETES_FRAMEWORKS, desc: _common.KUBERNETES_FRAMEWORKS_DESC},
			Item{title: _common.DOCKER_FRAMEWORKS, desc: _common.DOCKER_FRAMEWORKS_DESC},
		}
		l = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
		l.Title = ""
	case _common.FRONTEND_TEMPLATE_LIST:
		items := []list.Item{
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
		l.Title = "Here's the available Templates."

	case _common.BACKEND_TEMPLATE_LIST:
		items := []list.Item{}
		l = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
		l.Title = "Here's the available Templates."
	}

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return Bubble{list: l}
}
