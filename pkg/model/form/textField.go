package form

import (
	"cv-landing-cli/pkg/view"

	tea "github.com/charmbracelet/bubbletea"
)

type TextField struct {
	value   *string
	height  int
	width   int
	column  int
	focused bool
}

func NewTextField(height int) *TextField {
	return &TextField{
		value:   new(string),
		height:  height,
		focused: false,
	}
}

func (m *TextField) SetFocus(focused bool) {
	m.focused = focused
}

func (m *TextField) IsFocused() bool {
	return m.focused
}

func (m TextField) Init() tea.Cmd {
	return nil
}

func (m *TextField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyLeft:
			if m.column > 0 {
				m.column--
			}
		case tea.KeyRight:
			n := m.currentLen()
			if n != nil && m.column < *n {
				m.column++
			}
		case tea.KeyCtrlLeft:
			m.column = 0
		case tea.KeyCtrlRight:
			n := m.currentLen()
			if n != nil {
				m.column = *n
			}
		case tea.KeyBackspace:
			m.handleBackspace()
		case tea.KeyDelete:
			m.handleDelete()
		case tea.KeyRunes, tea.KeySpace:
			m.handleInput(msg.String())
		}
	}
	return m, nil
}

func (m *TextField) handleInput(input string) {
	if m.value == nil {
		return
	}
	runes := []rune(*m.value)
	newRunes := append(runes[:m.column], []rune(input)...)
	newRunes = append(newRunes, runes[m.column:]...)
	*m.value = string(newRunes)
	m.column += len([]rune(input))
}

func (m *TextField) handleBackspace() {
	if m.column == 0 || m.value == nil {
		return
	}
	runes := []rune(*m.value)
	*m.value = string(append(runes[:m.column-1], runes[m.column:]...))
	m.column--
}

func (m *TextField) handleDelete() {
	v := m.value
	if v == nil {
		return
	}
	runes := []rune(*v)
	if m.column >= len(runes) {
		return
	}
	*v = string(append(runes[:m.column], runes[m.column+1:]...))
}

func (m TextField) View() string {
	return view.Textfield(*m.value, m.column, m.focused, m.width, m.height)
}

func (m *TextField) currentLen() *int {
	if m.value == nil {
		return nil
	}
	n := len([]rune(*m.value))
	return &n
}

func (m *TextField) SetWidth(value int) {
	m.width = value
}

func (m TextField) GetWidth() int {
	return m.width
}

func (m TextField) GetHeight() int {
	return m.height
}
