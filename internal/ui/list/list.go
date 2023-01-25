package list

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type Bubble struct {
	list     list.Model
	choice   string
	quitting bool
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.SetWidth(msg.Width)
		return b, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			item, ok := b.list.SelectedItem().(item)
			if ok {
				b.choice = item.title
			}
			return b, nil
		}
	}

	var cmd tea.Cmd
	b.list, cmd = b.list.Update(msg)
	return b, cmd
}

func (b Bubble) View() string {
	if b.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", b.choice))
	}
	if b.quitting {
		return quitTextStyle.Render("Quitting? That’s cool.")
	}
	return "\n" + b.list.View()
}

func New() Bubble {
	items := []list.Item{
		item{title: "vue", desc: "Generate Vue.js App Template"},
		item{title: "vue-ts", desc: "Generate Vue.js App Template in TypeScript"},
		item{title: "react", desc: "Generate React App Template"},
		item{title: "react-ts", desc: "Generate React App Template in TypeScript"},
		item{title: "next", desc: "Generate Next.js App Template"},
		item{title: "next-ts", desc: "Generate Next.js App Template in TypeScript"},
		item{title: "vanilla", desc: "Generate Vanilla.js App Template"},
		item{title: "vanilla-ts", desc: "Generate Vanilla.js App Template in TypeScript"},
		item{title: "gatsby", desc: "Generate Gatsby App Template in TypeScript"},
		item{title: "gatsby-ts", desc: "Generate Gatsby App Template in TypeScript"},
	}
	const defaultWidth = 20
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	l.Title = "Here's the available Templates."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return Bubble{list: l}
}
