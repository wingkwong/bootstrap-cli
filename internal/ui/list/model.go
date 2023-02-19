package list

import (
	_list "github.com/charmbracelet/bubbles/list"
)

type Bubble struct {
	List   _list.Model
	Active bool
}

func New(
	list _list.Model,
	active bool,
) Bubble {
	return Bubble{
		List:   list,
		Active: active,
	}
}

func (b *Bubble) SetActive(v bool) {
	b.Active = v
}

func (b *Bubble) SetSize(width, height int) {
	b.List.SetSize(
		width,
		height,
	)
}
