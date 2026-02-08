package action

import (
	"cv-landing-cli/pkg/model/form"
	"cv-landing-cli/pkg/model/history"
	"cv-landing-cli/pkg/model/shell"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuEntry struct {
	Name     string
	NewModel func() shell.ShellContent
}

type ActionModel struct {
	entries []MenuEntry
	cursor  int
	title   string
	tips    string
	width   int
	height  int
}

func NewModel(title string, tips string) *ActionModel {
	model := ActionModel{
		cursor: 0,
		title:  title,
		tips:   tips,
	}
	model.entries = []MenuEntry{
		{
			Name: "Add items",
			NewModel: func() shell.ShellContent {
				addItemsModel :=
					form.NewModel([]form.FormField{
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
						"Start date",
						"End date",
					}, "Add activity", "asd")
				addItemsModel.SetWidth(model.width)
				addItemsModel.SetHeight(model.height)
				history := history.New()
				return history.Push(&addItemsModel)
			},
		},
		{
			Name: "Remove items",
		},
	}
	return &model
}

func (m ActionModel) Init() tea.Cmd {
	return nil
}

func (m *ActionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	for i, entry := range m.entries {
		if i == m.cursor {
			sb.WriteString("> ")
		} else {
			sb.WriteString("  ")
		}
		sb.WriteString(entry.Name)
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m *ActionModel) SetHeight(value int) {
	m.height = value
}

func (m *ActionModel) SetWidth(value int) {
	m.width = value
}

func (m ActionModel) GetTitle() string {
	return m.title
}

func (m ActionModel) GetTips() string {
	return m.tips
}
