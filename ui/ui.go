package ui

import (
	"fmt"
	"io"
	"os"
	"reflect"
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

/* ITEM DELEGATE */
type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	style := itemStyle
	if index == m.Index() {
		style = selectedStyle
	}

	var text string
	if i, ok := listItem.(list.DefaultItem); ok {
		text = i.Title()
	}

	fmt.Fprint(w, style.Render(text))
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

/* TWEAK LIST */
func tweakList(l list.Model) list.Model {
	l.SetDelegate(itemDelegate{})

	l.SetShowHelp(false)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetShowPagination(false)
	l.SetShowFilter(true)
	l.FilterInput.Prompt = "> "
	l.SetShowFilter(true)

	return l
}

/* STATE ENUM */
type state int

const (
	channels_list state = iota
	videos_list
)

/* MODEL */
type model struct {
	state         state
	current_list  *list.Model
	channels_list list.Model
	views_list    list.Model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		keypress := msg.String()
		switch keypress {
		case "l":
			switch m.state {
			case channels_list:
				index := m.current_list.Index()
				item := m.current_list.SelectedItem().(itemWrapper[Channel])

				if item.data.Islive {
					if reflect.ValueOf(item.sublist).IsZero() {
						videos := item.data.Platform.GetVods(item.data)
						item.sublist = tweakList(list.New(convertToItems(videos), itemDelegate{}, 10, 20))
						m.channels_list.SetItem(index, item)
					}

					m.current_list = &item.sublist
					m.state = videos_list
				}
			}
		case "h":
			switch m.state {
			case videos_list:
				m.current_list = &m.channels_list
				m.state = channels_list
			}
		}
	}

	// var cmd tea.Cmd
	l, cmd := m.current_list.Update(msg)
	m.current_list = &l

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
	m := model{}
	m.channels_list = tweakList(list.New(convertToItems(items), itemDelegate{}, 10, 20))
	m.current_list = &m.channels_list
	m.state = channels_list

	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}
}
