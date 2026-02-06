package form

import (
	"fmt"
	"os"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
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
	Fields        []FormField
	labels        []string
	cursor        int
	previousModel tea.Model
}

func NewModel(fields []FormField, labels []string, previousModel tea.Model) FormModel {
	if len(fields) > 0 {
		fields[0].SetFocus(true)
	}
	return FormModel{
		Fields:        fields,
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.Fields[m.cursor].SetFocus(false)
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
		}
		m.Fields[m.cursor].SetFocus(true)
	}

	updatedModel, cmd := m.Fields[m.cursor].Update(msg)
	if assertedModel, ok := updatedModel.(FormField); ok {
		m.Fields[m.cursor] = assertedModel
	}

	return m, cmd
}

func (m FormModel) View() string {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, height = 80, 24
	}

	availableHeight := height - 2
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
		field := m.Fields[idx]
		label := m.labels[idx]
		if adaptiveField, ok := field.(AdaptiveWidthField); ok {
			adaptiveField.SetWidth(min(width-maxLabelLen-1, 50))
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

	for i, field := range m.Fields {
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
