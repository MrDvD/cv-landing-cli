package form

import (
	"cv-landing-cli/pkg/view"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectField struct {
	options  []string
	selected *int
	cursor   int
	focused  bool
}

func NewSelectField(options []string, defaultIndex int) *SelectField {
	idx := defaultIndex
	return &SelectField{
		options:  options,
		selected: &idx,
		cursor:   idx,
		focused:  false,
	}
}

func (m *SelectField) SetFocus(focused bool) {
	m.focused = focused
}

func (m *SelectField) IsFocused() bool {
	return m.focused
}

func (m SelectField) Init() tea.Cmd {
	return nil
}

func (m *SelectField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyLeft:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyRight:
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case tea.KeyEnter:
			*m.selected = m.cursor
		}
	}
	return m, nil
}

func (m SelectField) View() string {
	return view.Select(m.options, m.cursor, *m.selected, m.focused)
}
