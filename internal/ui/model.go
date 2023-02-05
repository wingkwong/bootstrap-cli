package ui

import (
	"github.com/wingkwong/bootstrap-cli/internal/ui/list"
)

type Bubble struct {
	l    list.Bubble
	keys KeyMap
}

func New() Bubble {
	return Bubble{
		l:    list.New(),
		keys: DefaultKeyMap(),
	}
}
