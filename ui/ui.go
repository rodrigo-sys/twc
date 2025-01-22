package ui

import (
	"fmt"
	"io"
	"os"
	"reflect"

	. "twc/channel"
	. "twc/video"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	. "twc/platforms/kick"
	. "twc/platforms/twitch"
	. "twc/platforms/youtube"

	"twc/ui/nochannels"
)

/* ITEM */
type ItemWrapper[T any] struct {
	Data    T
	Sublist list.Model
}

func (i ItemWrapper[T]) Title() string {
	switch i := any(i.Data).(type) {
	case Channel:
		return i.Name()
	case Video:
		return i.Name()
	}
	return ""
}

func (i ItemWrapper[T]) Description() string {
	return ""
}

func (i ItemWrapper[T]) FilterValue() string {
	switch i := any(i.Data).(type) {
	case Channel:
		return i.Name()
	case Video:
		return i.Name()
	}
	return ""
}

/* ITEM DELEGATE */
type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	/* base style */
	style := itemStyle

	/* style selected item */
	if index == m.Index() {
		style = selectedStyle

		/* use different style per flatform */
		switch i := listItem.(type) {
		case ItemWrapper[Channel]:
			switch i.Data.(type) {
			case *Youtube:
				style = youtubeStyle
			case *Kick:
				style = kickStyle
			case *Twitch:
				style = twitchStyle
			}
		}
	}

	/* set text as Title return value */
	var text string
	if i, ok := listItem.(list.DefaultItem); ok {
		text = i.Title()
	}

	/* set style for live channels */
	switch i := listItem.(type) {
	case ItemWrapper[Channel]:
		if i.Data.Islive() {
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

	/* set keymaps */
	l.KeyMap.NextPage.SetKeys("n")
	l.KeyMap.PrevPage.SetKeys("p")

	return l
}

/* STATE ENUM */
type state int

const (
	channels_list state = iota
	videos_list
	loading
)

/* MODEL */
type model struct {
	spinner      spinner.Model
	state        state
	current_list list.Model

	channels_list list.Model
	videos_list   list.Model

	choice any
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case vodsListMsg:
		index := m.channels_list.Index()
		item := m.channels_list.SelectedItem().(ItemWrapper[Channel])

		item.Sublist = list.Model(msg)
		m.channels_list.SetItem(index, item)

		m.videos_list = item.Sublist
		m.current_list = m.videos_list
		m.state = videos_list
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if m.current_list.FilterState() == 1 {
			break
		}

		keypress := msg.String()
		switch keypress {
		case "l":
			switch m.state {
			case channels_list:
				m.channels_list = m.current_list
				item := m.channels_list.SelectedItem().(ItemWrapper[Channel])

				if !item.Data.Islive() {
					if reflect.ValueOf(item.Sublist).IsZero() {
						m.state = loading
						return m, tea.Batch(m.spinner.Tick, getVodsList(item))
					}
					m.videos_list = item.Sublist
					m.current_list = m.videos_list
					m.state = videos_list
				} else {
					m.choice = item
					return m, tea.Quit
				}
			case videos_list:
				m.videos_list = m.current_list
				item := m.videos_list.SelectedItem().(ItemWrapper[Video])
				m.choice = item
				return m, tea.Quit
			}
		case "h":
			switch m.state {
			case videos_list:
				if reflect.ValueOf(m.channels_list).IsZero() {
					break
				}
				m.current_list = m.channels_list
				m.state = channels_list
			}
		}
	}

	l, cmd := m.current_list.Update(msg)
	m.current_list = l

	return m, cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	// return m.current_list.View()
	if m.state == loading {
		return m.spinner.View()
	}
	return m.current_list.View()
}

/* CMD */
type vodsListMsg list.Model

func getVodsList(item ItemWrapper[Channel]) tea.Cmd {
	return func() tea.Msg {
		return vodsListMsg(tweakList(list.New(convertToItems(item.Data.GetVods()), itemDelegate{}, 10, 20)))
	}
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

	switch any(items).(type) {
	// case Channels:
	case []Channel:
		m.channels_list = tweakList(list.New(convertToItems(items), itemDelegate{}, 10, 20))
		m.current_list = m.channels_list
		m.state = channels_list
	// case Videos:
	case []Video:
		m.videos_list = tweakList(list.New(convertToItems(items), itemDelegate{}, 10, 20))
		m.current_list = m.videos_list
		m.state = videos_list
	}

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	m.spinner = s

	return m
}

func loguear(file_name string, text string) {
	f, _ := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(text)
}

/* THE MENU */
func Menu[T any](items []T) any {
	var m tea.Model
	var p *tea.Program

	if len(items) == 0 {
		if _, ok := any(items).([]Channel); ok {
			m = nochannels.Model{}
			p = tea.NewProgram(m)
		}
	} else {
		m = initialModel(items)
		p = tea.NewProgram(m, tea.WithAltScreen())
	}

	m, err := p.Run()
	if err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok {
		return m.choice
	} else {
		return nil
	}
}
