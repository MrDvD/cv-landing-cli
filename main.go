package main

import (
	"cv-landing-cli/pkg/model/action"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(action.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("There's an error:", err)
		os.Exit(1)
	}
}
