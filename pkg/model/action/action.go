package action

import (
	"cv-landing-cli/pkg/model/form"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuEntry struct {
	Name     string
	NewModel func() tea.Model
}

type ActionModel struct {
	entries []MenuEntry
	cursor  int
}

func NewModel() ActionModel {
	model := ActionModel{
		cursor: 0,
	}
	model.entries = []MenuEntry{
		{
			Name: "Add items",
			NewModel: func() tea.Model {
				addItemsModel := form.NewModel([]form.FormField{
					form.NewTextField(1),
					form.NewTextField(1),
					form.NewTextField(3),
					form.NewSelectField([]string{"project", "education", "event"}, 0),
					form.NewTextField(1),
					form.NewDateField(),
					form.NewDateField(),
				}, []string{
					"Name",
					"Subtitle",
					"Description",
					"Type",
					"Meta label",
					"Date start",
					"Date end",
				}, &model)
				addItemsModel.Fields[0].SetFocus(true)
				return addItemsModel
			},
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
			if entry.NewModel != nil {
				return entry.NewModel(), nil
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
