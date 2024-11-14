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

	. "twc/platform/kick"
	. "twc/platform/twitch"
	. "twc/platform/youtube"
)

/* ITEM */
type ItemWrapper[T any] struct {
	Data    T
	Sublist list.Model
}

func (i ItemWrapper[T]) Title() string {
	switch i := any(i.Data).(type) {
	case Channel:
		return i.Name
	case Video:
		return i.Name
	}
	return ""
}

func (i ItemWrapper[T]) Description() string {
	return ""
}

func (i ItemWrapper[T]) FilterValue() string {
	switch i := any(i.Data).(type) {
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

		switch listItem.(ItemWrapper[Channel]).Data.Platform.(type) {
		case Youtube:
			style = youtubeStyle
		case Kick:
			style = kickStyle
		case Twitch:
			style = twitchStyle
		}
	}

	var text string
	if i, ok := listItem.(list.DefaultItem); ok {
		text = i.Title()
	}

	switch i := listItem.(type) {
	case ItemWrapper[Channel]:
		if i.Data.Islive {
			style = style.Foreground(lipgloss.Color("11"))
		}
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
	state        state
	current_list *list.Model

	channels_list list.Model
	videos_list   list.Model

	// choice        itemWrapper
	// choice_channel itemWrapper[Channel]
	// choice_video   itemWrapper[Video]

	choice any
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		keypress := msg.String()
		switch keypress {
		case "l":
			switch m.state {
			case channels_list:
				m.channels_list = *m.current_list
				index := m.channels_list.Index()
				item := m.channels_list.SelectedItem().(ItemWrapper[Channel])

				if !item.Data.Islive {
					if reflect.ValueOf(item.Sublist).IsZero() {
						videos := item.Data.Platform.GetVods(item.Data)
						// fmt.Println(len(videos))

						item.Sublist = tweakList(list.New(convertToItems(videos), itemDelegate{}, 10, 20))
						m.channels_list.SetItem(index, item)
					}

					m.current_list = &item.Sublist
					m.state = videos_list
				} else {
					m.choice = item
					return m, tea.Quit
				}
			case videos_list:
				m.videos_list = *m.current_list
				item := m.videos_list.SelectedItem().(ItemWrapper[Video])
				m.choice = item
				return m, tea.Quit
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
		list_items[i] = ItemWrapper[T]{
			Data: item,
		}
	}

	return list_items
}

func initialModel[T any](items []T) model {
	m := model{}
	m.channels_list = tweakList(list.New(convertToItems(items), itemDelegate{}, 10, 20))
	m.current_list = &m.channels_list
	m.state = channels_list

	return m
}

/* THE MENU */
func Menu[T any](items []T) any {
	p := tea.NewProgram(initialModel(items), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}

	return m.(model).choice
}
