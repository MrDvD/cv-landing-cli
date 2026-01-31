package action

import (
	"cv-landing-cli/pkg/model"
	"cv-landing-cli/pkg/model/form"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuEntry struct {
	Name  string
	Model tea.Model
}

type ActionModel struct {
	entries []MenuEntry
	cursor  int
}

func NewModel() ActionModel {
	activity := model.Activity{}
	model := ActionModel{
		cursor: 0,
	}
	model.entries = []MenuEntry{
		{
			Name: "Add items",
			Model: form.NewModel([]form.Field{
				{
					Label:  "Name",
					Value:  &activity.Name,
					Height: 1,
				},
				{
					Label:  "Description",
					Value:  &activity.Description,
					Height: 5,
				},
			}, activity, &model),
		},
		{
			Name: "Remove items",
		},
	}
	return model
}

func (m ActionModel) Init() tea.Cmd {
	return nil
}

func (m ActionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor+1 < len(m.entries) {
				m.cursor++
			}
		case "enter":
			entry := m.entries[m.cursor]
			if entry.Model != nil {
				return entry.Model, nil
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m ActionModel) View() string {
	var sb strings.Builder
	sb.WriteString("Please choose an action you'd like to perform:\n\n")
	for i, entry := range m.entries {
		if i == m.cursor {
			sb.WriteString("> ")
		} else {
			sb.WriteString("  ")
		}
		sb.WriteString(entry.Name)
		sb.WriteString("\n")
	}
	sb.WriteString("\nPress 'q' to quit")
	return sb.String()
}
