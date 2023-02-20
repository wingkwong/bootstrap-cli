package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Bubble struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode textinput.CursorMode
	Finished   bool
	Active     bool
}

func (b Bubble) IsFinished() bool { return b.Finished }

func (b Bubble) GetInputs() []textinput.Model { return b.Inputs }

func (b *Bubble) SetActive(v bool) {
	b.Active = v
}
