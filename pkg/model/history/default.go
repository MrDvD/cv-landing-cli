package history

import (
	"cv-landing-cli/pkg/model/shell"

	tea "github.com/charmbracelet/bubbletea"
)

type History struct {
	current  shell.ShellContent
	previous *History
}

func New() HistoryModel {
	return &History{}
}

func (h History) Push(nextChild shell.ShellContent) HistoryModel {
	return &History{
		current:  nextChild,
		previous: &h,
	}
}

func (h History) Init() tea.Cmd {
	return h.current.Init()
}

func (h *History) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" && h.previous != nil && h.previous.Current() != nil {
			return h.previous, nil
		}
	}

	var nextM tea.Model
	newChild, cmd := h.current.Update(msg)
	switch newChild := newChild.(type) {
	case HistoryModel:
		nextM = h.Push(newChild.Current())
	case shell.ShellContent:
		h.current = newChild
		nextM = h
	}
	return nextM, cmd
}

func (h History) View() string {
	return h.current.View()
}

func (h History) GetTitle() string {
	return h.current.GetTitle()
}

func (h History) GetTips() string {
	return h.current.GetTips()
}

func (h *History) SetHeight(value int) {
	h.current.SetHeight(value)
}

func (h *History) SetWidth(value int) {
	h.current.SetWidth(value)
}

func (h History) Current() shell.ShellContent {
	return h.current
}
