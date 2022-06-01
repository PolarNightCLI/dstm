package form

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
)

type TextInputRow struct {
	label   string
	value   string
	editing bool
	checker func(string) bool
	input   *textinput.Model
}

func newTextInputModel(placeholder string, charLimit, width int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = " " + placeholder
	ti.CharLimit = charLimit
	ti.Width = width
	return ti
}

func NewTextInputRow(label, placeholder, value string, checker func(string) bool) TextInputRow {
	tf := newTextInputModel(placeholder, 156, 20)
	return TextInputRow{
		label:   label,
		value:   value,
		checker: checker,
		input:   &tf,
	}
}

func (r TextInputRow) Focus() any {
	r.editing = true
	r.input.SetValue(r.value)
	r.input.Prompt = okMark
	r.input.Focus()
	return r
}

func (r TextInputRow) UnFocus() any {
	r.editing = false
	newValue := r.input.Value()

	if r.checker == nil || r.checker(newValue) {
		r.value = newValue
	}
	return r
}

func (r TextInputRow) isEditing() bool {
	return r.editing
}

func (r TextInputRow) Type() string {
	return "TextInputRow"
}

func (r TextInputRow) Init() tea.Cmd {
	return textinput.Blink
}

func (r TextInputRow) Update(msg tea.Msg) (any, tea.Cmd) {
	if !r.editing {
		return r, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			return r, finishEditCmd
		}
	}

	tf, cmd := (*(r.input)).Update(msg)
	r.input = &tf
	if r.checker != nil && !r.checker(r.input.Value()) {
		r.input.Prompt = ngMark
	} else {
		r.input.Prompt = okMark
	}
	return r, cmd
}

func (r TextInputRow) View() string {
	var doc strings.Builder
	l := runewidth.StringWidth(r.label)
	s := strings.Repeat(" ", 12-l)
	doc.WriteString(r.label + s)
	if !r.editing {
		doc.WriteString(r.value)
		return doc.String()
	}

	doc.WriteString(r.input.View())
	return doc.String()
}
