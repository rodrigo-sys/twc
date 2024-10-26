package ui

import (
	. "twc/types"

	"github.com/charmbracelet/bubbles/list"
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
