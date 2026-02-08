package shell

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SupplyMeta interface {
	GetTitle() string
	GetTips() string
}

type consumeMeta interface {
	setTitle(text string)
	setTips(text string)
}

type consumeSize interface {
	SetWidth(value int)
	SetHeight(value int)
}

type ShellContent interface {
	tea.Model
	SupplyMeta
	consumeSize
}

type Shell struct {
	title   string
	tips    string
	content ShellContent
	width   int
	height  int
}
