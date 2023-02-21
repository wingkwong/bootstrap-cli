package ui

type Item struct {
	id                                      int
	title, name, desc, command, commandArgs string
}

func (i Item) Id() int { return i.id }

func (i Item) Title() string { return i.title }

func (i Item) Description() string { return i.desc }

func (i Item) Name() string { return i.name }

func (i Item) Command() string { return i.command }

func (i Item) CommandArgs() string { return i.commandArgs }

func (i Item) FilterValue() string { return i.title }
