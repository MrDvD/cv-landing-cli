package form

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type FormField interface {
	tea.Model
	SetFocus(bool)
	IsFocused() bool
}

type Cursor struct {
	column int
	row    int
}

type FormModel struct {
	Fields        []FormField
	labels        []string
	cursor        int
	previousModel tea.Model
}

func NewModel(Fields []FormField, labels []string, previousModel tea.Model) FormModel {
	return FormModel{
		Fields:        Fields,
		labels:        labels,
		cursor:        0,
		previousModel: previousModel,
	}
}

func (m FormModel) Init() tea.Cmd {
	return nil
}

func (m FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.Fields[m.cursor].SetFocus(false)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m.previousModel, nil
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.Fields)-1 {
				m.cursor++
			}
		case tea.KeyCtrlUp:
			m.cursor = 0
		case tea.KeyCtrlDown:
			m.cursor = len(m.Fields) - 1
		}
	}
	updatedModel, cmd := m.Fields[m.cursor].Update(msg)
	if assertedModel, ok := updatedModel.(FormField); ok {
		m.Fields[m.cursor] = assertedModel
	}
	m.Fields[m.cursor].SetFocus(true)
	return m, cmd
}

func (m FormModel) View() string {
	var sb strings.Builder

	maxLen := m.getMaxLabelLen()
	for i, field := range m.Fields {
		rendered := field.View()
		lines := strings.Split(rendered, "\n")

		for j, line := range lines {
			if j == 0 {
				fmt.Fprintf(&sb, "%*s %s\n", maxLen, m.labels[i], line)
			} else {
				indent := strings.Repeat(" ", maxLen+1)
				fmt.Fprintf(&sb, "%s%s\n", indent, line)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (m FormModel) getMaxLabelLen() int {
	maxLen := 0
	for _, label := range m.labels {
		labelLen := len([]rune(label))
		if labelLen > maxLen {
			maxLen = labelLen
		}
	}
	return maxLen
}
