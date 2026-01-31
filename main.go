package main

import (
	"cv-landing-cli/pkg/model"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.NewActionModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("There's an error:", err)
		os.Exit(1)
	}
}
