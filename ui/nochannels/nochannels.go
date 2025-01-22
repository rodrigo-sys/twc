package nochannels

import (
	"os"
	"twc/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var keyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Bold(true)

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
		"You dont have any channel added\n" +
			"Press " + keyStyle.Render("e") + " to edit the file\n" +
			"Press " + keyStyle.Render("q") + " to quit\n" +
			"you also can use twc -e to open the channels file\n"

	return text
}
