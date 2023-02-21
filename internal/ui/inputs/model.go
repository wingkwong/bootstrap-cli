package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Bubble struct {
	FocusIndex int
	Inputs     []textinput.Model
	Finished   bool
	Active     bool
}

func (b Bubble) Init() tea.Cmd {
	return textinput.Blink
}

func (b Bubble) IsFinished() bool { return b.Finished }

func (b Bubble) GetInputs() []textinput.Model { return b.Inputs }

func (b *Bubble) SetActive(v bool) {
	b.Active = v
}
