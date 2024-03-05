package main

// A simple program that opens the alternate screen buffer and displays mouse
// coordinates and events.

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model{}, tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	mouseEvent tea.MouseMsg
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyDownMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return m, tea.Quit
		}

	case tea.MouseDownMsg:
		return m, tea.Printf("(X: %d, Y: %d) %s press", msg.X, msg.Y, msg)
	case tea.MouseUpMsg:
		return m, tea.Printf("(X: %d, Y: %d) %s release", msg.X, msg.Y, msg)
	case tea.MouseMoveMsg:
		s := msg.String()
		if s != "" {
			s += " "
		}
		s += "motion"
		return m, tea.Printf("(X: %d, Y: %d) %s", msg.X, msg.Y, s)
	}

	return m, nil
}

func (m model) View() string {
	s := "Do mouse stuff. When you're done press q to quit.\n"

	return s
}
