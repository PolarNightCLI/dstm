package form

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
)

type SelectorRow struct {
	label   string
	value   string
	editing bool
	options []string
	picker  *Selector
}

func NewSelectorRow(label, value string, opts []string) SelectorRow {
	p := NewSelector(opts, false, false, false)
	return SelectorRow{
		label:   label,
		value:   value,
		options: opts,
		picker:  &p,
	}
}

func (s SelectorRow) Focus() any {
	s.editing = true
	return s
}

func (s SelectorRow) UnFocus() any {
	s.editing = false
	s.value = s.picker.GetSelected()[0]
	return s
}

func (s SelectorRow) isEditing() bool {
	return s.editing
}

func (s SelectorRow) Type() string {
	return "SelectorRow"
}

func (s SelectorRow) Init() tea.Cmd {
	return nil
}

func (s SelectorRow) Update(msg tea.Msg) (any, tea.Cmd) {
	if !s.editing {
		return s, nil
	}

	selector, cmd := (*(s.picker)).Update(msg)
	if cmd != nil {
		switch cmd().(type) {
		case finishSelectMsg:
			return s, finishEditCmd
		}
	}

	s.picker = &selector
	return s, cmd
}

func (s SelectorRow) View() string {
	var doc strings.Builder
	l := runewidth.StringWidth(s.label)
	ss := strings.Repeat(" ", 12-l)
	doc.WriteString(s.label + ss)
	if !s.editing {
		doc.WriteString(s.value)
		return doc.String()
	}

	doc.WriteString(s.picker.View())
	return doc.String()
}
