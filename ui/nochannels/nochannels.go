package nochannels

import (
	"os"
	"twc/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var Style = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Bold(true)

type Model struct {
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "e":
			if channels_file := os.Getenv("CHANNELS_PATH"); channels_file != "" {
				utils.OpenWithDefaultApp(channels_file)
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	text :=
		"You currently have no channels added in\n" +
			lipgloss.NewStyle().Italic(true).Render(os.Getenv("CHANNELS_PATH")) + "\n" +
			"\nPress " + Style.Render("e") + " to edit the file.\n" +
			"Press " + Style.Render("q") + " to quit.\n" +
			"You can also use " + Style.Render("twc -e") + " to open the channels file."

	return text
}
