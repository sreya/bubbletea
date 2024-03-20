package main

// A simple program demonstrating the textarea component from the Bubbles
// component library.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram()

	m, err := tea.Run(p, model{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %T\n", m)
}

type errMsg error

type model struct {
	textarea textarea.Model
	err      error
}

func (m model) Init(ctx *tea.Context) (model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Init(ctx)
	m.textarea.Placeholder = "Once upon a time..."
	m.textarea.Focus()

	return m, cmd
}

func (m model) Update(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd := m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}
