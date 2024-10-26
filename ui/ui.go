package ui

import (
	. "twc/types"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
