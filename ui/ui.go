package ui

import (
	"fmt"
	"os"
	. "twc/types"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* ITEM */
type itemWrapper[T any] struct {
	data    T
	sublist list.Model
}

func (i itemWrapper[T]) Title() string {
	switch i := any(i.data).(type) {
	case Channel:
		return i.Name
	case Video:
		return i.Name
	}
	return ""
}

func (i itemWrapper[T]) Description() string {
	return ""
}

func (i itemWrapper[T]) FilterValue() string {
	switch i := any(i.data).(type) {
	case Channel:
		return i.Name
	case Video:
		return i.Name
	}
	return ""
}

/* STYLES */
var (
	itemStyle     = lipgloss.NewStyle().PaddingLeft(2)
	kickStyle     = selectedStyle.BorderForeground(lipgloss.Color("10"))
	youtubeStyle  = selectedStyle.BorderForeground(lipgloss.Color("9"))
	twitchStyle   = selectedStyle.BorderForeground(lipgloss.Color("13"))
	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.ThickBorder()).
			BorderLeft(true).
			BorderForeground(lipgloss.Color("12")).
			PaddingLeft(1)
)

/* STATE ENUM */
type state int

const (
	channels_list state = iota
	views_list
)

/* MODEL */
type model struct {
	// state state
	current_list  list.Model
	channels_list list.Model
	views_list    list.Model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	// return m, cmd
	var cmd tea.Cmd
	m.current_list, cmd = m.current_list.Update(msg)
	return m, cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.current_list.View()
}

/* HELPERS */
func convertToItems[T any](items []T) []list.Item {
	/* convert items[] to list.Item[] */
	list_items := make([]list.Item, len(items))
	for i, item := range items {
		list_items[i] = itemWrapper[T]{
			data: item,
		}
	}

	return list_items
}

/* THE MENU */
func Menu[T any](items []T) {
	// m := model{channels_list: list.New(convertToItems(items), list.NewDefaultDelegate(), 10, 20)}
	m := model{channels_list: list.New(convertToItems(items), itemDelegate{}, 10, 20)}
	m.current_list = m.channels_list

	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}
}
