package form

import (
	"strings"

	//"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type finishEditMsg struct{}

func finishEditCmd() tea.Msg {
	return finishEditMsg{}
}

type FormRow interface {
	TextInputRow | SelectorRow
	Focus() any
	UnFocus() any
	isEditing() bool
	Type() string

	Init() tea.Cmd
	Update(msg tea.Msg) (any, tea.Cmd)
	View() string
}

var (
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("100")).PaddingRight(3)
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("200")).PaddingRight(3)
	okMark        = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Render("✔ ")
	ngMark        = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("✘ ")
)

type Form[T FormRow] struct {
	rows     []T
	cursor   int
	focusing bool
}

func NewForm[T FormRow](rows []T) Form[T] {
	return Form[T]{
		rows: rows,
	}
}

func (f Form[T]) Init() tea.Cmd {
	//var cmds []tea.Cmd
	//for _, r := range f.rows {
	//	cmd := r.Init()
	//	cmds = append(cmds, cmd)
	//}
	//return tea.Batch(cmds...)
	return nil
}

func (f Form[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case finishEditMsg:
		f.focusing = false
		r := f.rows[f.cursor].UnFocus()
		f.rows[f.cursor] = r.(T)
		return f, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return f, tea.Quit
		case "enter":
			if f.focusing {
				r, cmd := f.rows[f.cursor].Update(msg)
				f.rows[f.cursor] = r.(T)
				return f, cmd
			}
			f.focusing = true
			r := f.rows[f.cursor].Focus()
			f.rows[f.cursor] = r.(T)
		case "up":
			if !f.focusing {
				f.cursor--
				if f.cursor < 0 {
					f.cursor = len(f.rows) - 1
				}
			}
		case "down":
			if !f.focusing {
				f.cursor++
				if f.cursor >= len(f.rows) {
					f.cursor = 0
				}
			}
		default:
			if f.focusing {
				r, cmd := f.rows[f.cursor].Update(msg)
				f.rows[f.cursor] = r.(T)
				return f, cmd
			}
		}
	}
	return f, nil
}

func (f Form[T]) View() string {
	var doc strings.Builder

	for i, row := range f.rows {
		var line string
		if f.cursor == i {
			if row.isEditing() {
				line = row.View()
			} else {
				line = selectedStyle.Render(row.View())
			}
		} else {
			line = normalStyle.Render(row.View())
		}
		doc.WriteString(line)
		doc.WriteString("\n")
	}

	return doc.String()
}
