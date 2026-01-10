package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethanefung/bubble-datepicker"
)

type model struct {
	selected   bool
	datepicker datepicker.Model
}

func initialModel(t time.Time) tea.Model {
	dp := datepicker.New(t)
	dp.SelectDate()

	return model{
		datepicker: dp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "enter":
			m.selected = true
			return m, tea.Quit

		// These keys should exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		default:
			datepicker, cmd := m.datepicker.Update(msg)
			m.datepicker = datepicker
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {

	// Send the UI for rendering
	return m.datepicker.View()
}

func main() {
	date := time.Now()

	lipgloss.SetDefaultRenderer(lipgloss.NewRenderer(os.Stderr))

	if len(os.Args) > 1 {
		// input := "20260103T150856Z"
		d, err := time.Parse("2006-01-02", "2025-12-24") // os.Args[1])
		if err == nil {
			fmt.Printf("yes date\n")
			date = d
		}

		d, err = time.Parse("20060102T150405Z", os.Args[1])
		if err == nil {
			fmt.Printf("yes datetime\n")
			date = d
		}

	}
	tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		// Fallback to Stderr if TTY isn't available
		tty = os.Stderr
	}
	defer tty.Close()

	p := tea.NewProgram(initialModel(date), tea.WithOutput(tty), tea.WithAltScreen())
	teaModel, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if finalModel, ok := teaModel.(model); ok {
		if finalModel.selected {
			fmt.Printf("%s\n", finalModel.datepicker.Time.Format("2006-01-02"))
		} else {
			os.Exit(1)
		}
	}

}
