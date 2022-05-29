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
	input   textinput.Model
	checker func() bool
}

func newTextInputModel(placeholder string, charLimit, width int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.CharLimit = charLimit
	ti.Width = width
	return ti
}

func NewTextInputRow(label, value string, checker func() bool) TextInputRow {
	return TextInputRow{
		input:   newTextInputModel(" "+value, 156, 20),
		label:   label,
		value:   value,
		checker: checker,
	}
}

func (r *TextInputRow) Focus() {
	r.editing = true
	r.input.Focus()
}

func (r *TextInputRow) UnFocus() {
	r.editing = false
	newValue := r.input.Value()
	if len(newValue) > 0 {
		r.value = newValue
	}
}

func (r TextInputRow) IsFocused() bool {
	return r.editing
}

func (r TextInputRow) Init() tea.Cmd {
	return textinput.Blink
}

func (r TextInputRow) Update(msg tea.Msg) (any, tea.Cmd) {
	if r.editing {
		var cmd tea.Cmd
		r.input, cmd = r.input.Update(msg)
		return r, cmd
	}
	return r, nil
}

func (r TextInputRow) View() string {
	var doc strings.Builder
	doc.WriteString(r.label + "   " + r.value)
	if r.editing {
		doc.WriteString("\n")
		doc.WriteString(r.input.View())
	}
	return doc.String()
}
