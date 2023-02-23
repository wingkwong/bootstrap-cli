package inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type inputBubble struct {
	placeholder string
	echoMode    textinput.EchoMode
}

func NewInputModel(data []inputBubble) Bubble {
	n := len(data)
	b := Bubble{
		Inputs: make([]textinput.Model, n),
	}
	var t textinput.Model
	for i := range b.Inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32
		if i == 0 {
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}
		t.Placeholder = data[i].placeholder
		t.EchoMode = data[i].echoMode
		if data[i].echoMode == 2 {
			t.EchoCharacter = '•'
		}
		b.Inputs[i] = t
	}
	return b
}

func NewViteInputModel() Bubble {
	return NewInputModel([]inputBubble{
		{placeholder: "Enter App Name", echoMode: textinput.EchoNormal},
		{placeholder: "Enter the directory", echoMode: textinput.EchoNormal},
	})
}

func NewMSSQLInputModel() Bubble {
	return NewInputModel([]inputBubble{
		{placeholder: "MSSQL_USER", echoMode: textinput.EchoNormal},
		{placeholder: "MSSQL_SA_PASSWORD", echoMode: textinput.EchoNormal},
		{placeholder: "MSSQL_PID", echoMode: textinput.EchoPassword},
		{placeholder: "name", echoMode: textinput.EchoNormal},
		{placeholder: "port", echoMode: textinput.EchoNormal},
	})
}
