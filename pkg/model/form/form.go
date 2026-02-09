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
	fields           []FormField
	labels           []string
	activePageFields []int
	cursor           int
	width            int
	height           int
	title            string
	tips             string
}

func NewModel(fields []FormField, labels []string, title string) FormModel {
	return FormModel{
		fields: fields,
		labels: labels,
		cursor: 0,
		title:  title,
	}
}

func (m *FormModel) Init() tea.Cmd {
	if len(m.fields) > 0 {
		m.fields[0].SetFocus(true)
	}
	m.recalculateActivePages()
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
				m.recalculateActivePages()
			}
		case tea.KeyDown:
			if m.cursor < len(m.fields)-1 {
				m.cursor++
				m.recalculateActivePages()
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
	var sb strings.Builder
	maxLabelLen := m.getMaxLabelLen()
	for _, idx := range m.activePageFields {
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
	return sb.String()
}

func (m *FormModel) recalculateActivePages() {
	pages := m.calculatePages(m.height - 1)

	activePageIndex := 0
	for i, page := range pages {
		if slices.Contains(page, m.cursor) {
			activePageIndex = i
		}
	}
	m.activePageFields = pages[activePageIndex]
	m.tips = fmt.Sprintf("[ Page %d/%d | Use Up/Down to navigate ]", activePageIndex+1, len(pages))
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
