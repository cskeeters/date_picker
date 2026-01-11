package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cskeeters/bubble-datepicker"
	"github.com/jehiah/go-strftime"
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
	var format string
	flag.StringVar(&format, "f", "%Y-%m-%d", "Output date format (POSIX strftime format)")
	flag.Parse()

	date := time.Now()

	lipgloss.SetDefaultRenderer(lipgloss.NewRenderer(os.Stderr))

	// Check for date argument after flag parsing
	args := flag.Args()
	if len(args) > 0 {
		input := args[0]

		// Try parsing as YYYY-MM-DD format
		if d, err := time.Parse("2006-01-02", input); err == nil {
			date = d
		} else if d, err := time.Parse("20060102T150405Z", input); err == nil {
			// Try parsing as ISO datetime format
			date = d
		} else {
			fmt.Fprintf(os.Stderr, "Error: Invalid date format: %s\n", input)
			os.Exit(1)
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
			fmt.Printf("%s\n", strftime.Format(format, finalModel.datepicker.Time))
		} else {
			os.Exit(1)
		}
	}

}
