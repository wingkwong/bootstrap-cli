package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Bubble struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
	finished   bool
}

func (b Bubble) IsFinished() bool { return b.finished }

func (b Bubble) GetInputs() []textinput.Model { return b.inputs }
