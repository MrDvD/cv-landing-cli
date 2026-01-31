package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ActionModel struct {
	choices []string
	cursor  int
}

func NewActionModel() ActionModel {
	return ActionModel{
		choices: []string{"Add items", "Remove items"},
		cursor:  0,
	}
}

func (a ActionModel) Init() tea.Cmd {
	return nil
}

func (a ActionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if a.cursor > 0 {
				a.cursor--
			}
		case "down":
			if a.cursor+1 < len(a.choices) {
				a.cursor++
			}
		case "q":
			return a, tea.Quit
		}
	}
	return a, nil
}

func (a ActionModel) View() string {
	var sb strings.Builder
	sb.WriteString("Please choose an action you'd like to perform:\n\n")
	for i, actionName := range a.choices {
		if i == a.cursor {
			sb.WriteString("> ")
		} else {
			sb.WriteString("  ")
		}
		sb.WriteString(actionName)
		sb.WriteString("\n")
	}
	sb.WriteString("\nPress 'q' to quit")
	return sb.String()
}
