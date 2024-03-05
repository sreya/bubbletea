package main

import (
	"fmt"
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/term/ansi"
)

type model struct {
	prevKey    tea.KeyDownMsg
	kittyFlags int
}

func (m model) Init() tea.Cmd {
	// return nil
	return tea.EnableEnhancedKeyboard
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyDownMsg:
	out:
		switch m.prevKey.String() {
		case "q":
			if msg.String() == "q" {
				cmd = tea.Quit
			}
		case "r":
			switch msg.String() {
			case "b":
				execute(ansi.RequestBackgroundColor)
			case "d":
				execute(ansi.RequestPrimaryDeviceAttributes)
			case "k":
				execute(ansi.RequestKittyKeyboard)
			case "o":
				execute(ansi.RequestModifyOtherKeys)
			}
		case "k":
			switch msg.String() {
			case "0":
				m.kittyFlags = 0
			case "1":
				m.kittyFlags |= ansi.KittyDisambiguateEscapeCodes
			case "2":
				m.kittyFlags |= ansi.KittyReportEventTypes
			case "3":
				m.kittyFlags |= ansi.KittyReportAlternateKeys
			case "4":
				m.kittyFlags |= ansi.KittyReportAllKeys
			case "5":
				m.kittyFlags |= ansi.KittyReportAssociatedKeys
			default:
				break out
			}
			execute(ansi.PushKittyKeyboard(m.kittyFlags))
		}
		m.prevKey = msg
	}
	switch msg := msg.(type) {
	case string:
		return m, tea.Batch(tea.Println(msg), cmd)
	case fmt.Stringer:
		return m, tea.Batch(tea.Println(msg.String()), cmd)
	}
	return m, cmd
}

func (m model) View() string {
	return "Type any key and it will be printed to the terminal. Press qq to quit."
}

func main() {
	defer execute(ansi.PushKittyKeyboard(0))
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func execute(seq string) {
	io.WriteString(os.Stdout, seq)
}
