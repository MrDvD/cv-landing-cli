package form

import (
	"cv-landing-cli/pkg/view"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Field struct {
	Label  string
	Key    string
	Value  *string
	Height int
}

type Cursor struct {
	column int
	row    int
}

type FormModel struct {
	fields        []Field
	cursor        Cursor
	previousModel tea.Model
}

func NewModel(fields []Field, result any, previousModel tea.Model) FormModel {
	return FormModel{
		fields: fields,
		cursor: Cursor{
			row:    0,
			column: 0,
		},
		previousModel: previousModel,
	}
}

func (m FormModel) Init() tea.Cmd {
	return nil
}

func (m FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m.previousModel, nil
		case tea.KeyLeft:
			if m.cursor.column > 0 {
				m.cursor.column--
			}
		case tea.KeyRight:
			n := m.currentLen()
			if n != nil && m.cursor.column < *n {
				m.cursor.column++
			}
		case tea.KeyCtrlUp:
			m.cursor.row = 0
		case tea.KeyCtrlDown:
			m.cursor.row = len(m.fields) - 1
		case tea.KeyCtrlLeft:
			m.cursor.column = 0
		case tea.KeyCtrlRight:
			n := m.currentLen()
			if n != nil {
				m.cursor.column = *n
			}
		case tea.KeyUp:
			if m.cursor.row > 0 {
				m.cursor.row--
				n := m.currentLen()
				if n != nil {
					m.cursor.column = *n
				}
			}
		case tea.KeyDown, tea.KeyEnter:
			if m.cursor.row < len(m.fields)-1 {
				m.cursor.row++
				n := m.currentLen()
				if n != nil {
					m.cursor.column = *n
				}
			}
		case tea.KeyBackspace:
			m.handleBackspace()
		case tea.KeyDelete:
			m.handleDelete()
		case tea.KeyRunes:
			m.handleInput(msg.String())
		}
	}
	return m, nil
}

func (m *FormModel) currentVal() *string {
	return m.fields[m.cursor.row].Value
}

func (m *FormModel) currentLen() *int {
	v := m.currentVal()
	if v == nil {
		return nil
	}
	n := len([]rune(*v))
	return &n
}

func (m *FormModel) handleInput(input string) {
	v := m.currentVal()
	if v == nil {
		return
	}
	runes := []rune(*v)
	newRunes := append(runes[:m.cursor.column], []rune(input)...)
	newRunes = append(newRunes, runes[m.cursor.column:]...)
	*v = string(newRunes)
	m.cursor.column += len([]rune(input))
}

func (m *FormModel) handleBackspace() {
	v := m.currentVal()
	if m.cursor.column == 0 || v == nil {
		return
	}
	runes := []rune(*v)
	*v = string(append(runes[:m.cursor.column-1], runes[m.cursor.column:]...))
	m.cursor.column--
}

func (m *FormModel) handleDelete() {
	v := m.currentVal()
	if v == nil {
		return
	}
	runes := []rune(*v)
	if m.cursor.column >= len(runes) {
		return
	}
	*v = string(append(runes[:m.cursor.column], runes[m.cursor.column+1:]...))
}

func (m FormModel) View() string {
	var sb strings.Builder

	maxLen := 0
	for _, field := range m.fields {
		labelLen := len([]rune(field.Label))
		if labelLen > maxLen {
			maxLen = labelLen
		}
	}

	for i, field := range m.fields {
		isFocused := (i == m.cursor.row)
		val := m.fields[i].Value
		rendered := view.Textfield(*val, m.cursor.column, isFocused)
		fmt.Fprintf(&sb, "%-*s : %s\n", maxLen, field.Label, rendered)
	}

	return sb.String()
}
