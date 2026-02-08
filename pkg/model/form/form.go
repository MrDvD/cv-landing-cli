package form

import (
	"fmt"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Focusable interface {
	SetFocus(bool)
	IsFocused() bool
}

type AdaptiveWidthField interface {
	SetWidth(value int)
	GetWidth() int
}

type FormField interface {
	tea.Model
	Focusable
	GetHeight() int
}

type FormModel struct {
	fields []FormField
	labels []string
	cursor int
	width  int
	height int
	title  string
	tips   string
}

func NewModel(fields []FormField, labels []string, title string, tips string) FormModel {
	if len(fields) > 0 {
		fields[0].SetFocus(true)
	}
	return FormModel{
		fields: fields,
		labels: labels,
		cursor: 0,
		title:  title,
		tips:   tips,
	}
}

func (m FormModel) Init() tea.Cmd {
	return nil
}

func (m *FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.fields[m.cursor].SetFocus(false)
		switch msg.Type {
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.fields)-1 {
				m.cursor++
			}
		}
		m.fields[m.cursor].SetFocus(true)
	}

	updatedModel, cmd := m.fields[m.cursor].Update(msg)
	if assertedModel, ok := updatedModel.(FormField); ok {
		m.fields[m.cursor] = assertedModel
	}

	return m, cmd
}

func (m FormModel) View() string {
	availableHeight := m.height - 2
	pages := m.calculatePages(availableHeight)

	activePageIndex := 0
	for i, page := range pages {
		if slices.Contains(page, m.cursor) {
			activePageIndex = i
		}
	}

	var sb strings.Builder
	activePageFields := pages[activePageIndex]
	maxLabelLen := m.getMaxLabelLen()

	for _, idx := range activePageFields {
		field := m.fields[idx]
		label := m.labels[idx]
		if adaptiveField, ok := field.(AdaptiveWidthField); ok {
			adaptiveField.SetWidth(m.width - maxLabelLen - 1)
		}

		rendered := field.View()
		lines := strings.Split(rendered, "\n")

		for j, line := range lines {
			var formatted string
			if j == 0 {
				formatted = fmt.Sprintf("%*s %s", maxLabelLen, label, line)
			} else {
				indent := strings.Repeat(" ", maxLabelLen)
				formatted = fmt.Sprintf("%s %s", indent, line)
			}
			sb.WriteString(formatted + "\n")
		}
		sb.WriteString("\n")
	}

	footer := fmt.Sprintf("\n[ Page %d/%d | Use Up/Down to navigate ]", activePageIndex+1, len(pages))
	sb.WriteString(footer)

	return sb.String()
}

func (m FormModel) calculatePages(availableHeight int) [][]int {
	var pages [][]int
	var currentPage []int
	currentHeight := 0

	for i, field := range m.fields {
		fieldTotalHeight := field.GetHeight() + 1

		if currentHeight+fieldTotalHeight > availableHeight && len(currentPage) > 0 {
			pages = append(pages, currentPage)
			currentPage = []int{i}
			currentHeight = fieldTotalHeight
		} else {
			currentPage = append(currentPage, i)
			currentHeight += fieldTotalHeight
		}
	}

	if len(currentPage) > 0 {
		pages = append(pages, currentPage)
	}

	return pages
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

func (m *FormModel) SetWidth(value int) {
	m.width = value
}

func (m *FormModel) SetHeight(value int) {
	m.height = value
}

func (m FormModel) GetTitle() string {
	return m.title
}

func (m FormModel) GetTips() string {
	return m.tips
}
