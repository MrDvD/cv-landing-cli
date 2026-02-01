package form

import (
	"cv-landing-cli/pkg/view"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

type DateField struct {
	digits  string
	focused bool
}

func NewDateField() *DateField {
	return &DateField{}
}

func (m *DateField) SetFocus(focused bool) { m.focused = focused }
func (m *DateField) IsFocused() bool       { return m.focused }
func (m DateField) Init() tea.Cmd          { return nil }

func (m *DateField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyBackspace:
			if len(m.digits) > 0 {
				m.digits = m.digits[:len(m.digits)-1]
			}
		case tea.KeyRunes:
			r := msg.Runes[0]
			if unicode.IsDigit(r) && len(m.digits) < 8 {
				m.digits += string(r)
			}
		}
	}
	return m, nil
}

func (m DateField) View() string {
	return view.Date(m.digits, m.focused)
}
