package shell

import (
	"cv-landing-cli/pkg/view"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
)

func NewShell(content ShellContent) *Shell {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	width, height = convertSizes(width, height)
	return &Shell{
		content: content,
		width:   width,
		height:  height,
	}
}

func (m Shell) Init() tea.Cmd {
	return m.content.Init()
}

func (m *Shell) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = convertSizes(msg.Width, msg.Height)
		m.content.SetWidth(m.width)
		m.content.SetHeight(m.height - 2)

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	newContent, cmd := m.content.Update(msg)
	m.content = newContent.(ShellContent)
	m.setTitle(m.content.GetTitle())
	m.setTips(m.content.GetTips())
	return m, cmd
}

func (m Shell) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	childView := m.content.View()
	return view.Window(m.width, m.height, m.title, childView, m.tips)
}

func convertSizes(width int, height int) (int, int) {
	return min(width, 80), height
}

func (m Shell) getHeight() int {
	return m.height
}

func (m Shell) getWidth() int {
	return m.width
}

func (m *Shell) setTitle(text string) {
	m.title = text
}

func (m *Shell) setTips(text string) {
	m.tips = text
}
