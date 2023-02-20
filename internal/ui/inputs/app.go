package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func NewViteInputModel() Bubble {
	b := Bubble{
		Inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range b.Inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Enter App Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Enter the directory"
		}
		b.Inputs[i] = t
	}

	return b
}

func NewMSSQLInputModel() Bubble {
	b := Bubble{
		Inputs: make([]textinput.Model, 5),
	}

	var t textinput.Model
	for i := range b.Inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "MSSQL_USER"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "MSSQL_SA_PASSWORD"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		case 2:
			t.Placeholder = "MSSQL_PID"
		case 3:
			t.Placeholder = "name"
		case 4:
			t.Placeholder = "port"
		}

		b.Inputs[i] = t
	}

	return b
}
