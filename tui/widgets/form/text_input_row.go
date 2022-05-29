package widgets

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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

func (r TextInputRow) Init() tea.Cmd {
	return textinput.Blink
}

func (r TextInputRow) Update(msg tea.Msg) (any, tea.Cmd) {
	if r.editing {
		var cmd tea.Cmd
		tf, cmd := (*(r.input)).Update(msg)
		r.input = &tf
		if r.checker != nil && !r.checker(r.input.Value()) {
			r.input.Prompt = ngMark
		} else {
			r.input.Prompt = okMark
		}
		return r, cmd
	}
	return r, nil
}

func (r TextInputRow) View() string {
	var doc strings.Builder
	doc.WriteString(r.label + "   ")
	if !r.editing {
		doc.WriteString(r.value)
		return doc.String()
	}

	doc.WriteString(r.input.View())
	return doc.String()
}
