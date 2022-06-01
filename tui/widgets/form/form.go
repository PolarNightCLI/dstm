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

type Form struct {
	rows     []any
	cursor   int
	focusing bool
}

func NewForm(rows []any) Form {
	return Form{
		rows: rows,
	}
}

func (f Form) Init() tea.Cmd {
	//var cmds []tea.Cmd
	//for _, r := range f.rows {
	//	cmd := r.Init()
	//	cmds = append(cmds, cmd)
	//}
	//return tea.Batch(cmds...)
	return nil
}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case finishEditMsg:
		f.focusing = false
		row := f.rows[f.cursor]
		switch row.(type) {
		case TextInputRow:
			f.rows[f.cursor] = row.(TextInputRow).UnFocus()
		case SelectorRow:
			f.rows[f.cursor] = row.(SelectorRow).UnFocus()
		}
		return f, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return f, tea.Quit
		case "enter":
			if f.focusing {
				var cmd tea.Cmd
				row := f.rows[f.cursor]
				switch row.(type) {
				case TextInputRow:
					f.rows[f.cursor], cmd = row.(TextInputRow).Update(msg)
				case SelectorRow:
					f.rows[f.cursor], cmd = row.(SelectorRow).Update(msg)
				}
				return f, cmd
			}
			f.focusing = true
			row := f.rows[f.cursor]
			switch row.(type) {
			case TextInputRow:
				f.rows[f.cursor] = row.(TextInputRow).Focus()
			case SelectorRow:
				f.rows[f.cursor] = row.(SelectorRow).Focus()
			}
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
				var cmd tea.Cmd
				row := f.rows[f.cursor]
				switch row.(type) {
				case TextInputRow:
					f.rows[f.cursor], cmd = row.(TextInputRow).Update(msg)
				case SelectorRow:
					f.rows[f.cursor], cmd = row.(SelectorRow).Update(msg)
				}
				return f, cmd
			}
		}
	}
	return f, nil
}

func (f Form) View() string {
	var doc strings.Builder

	for i, row := range f.rows {
		var line string
		switch row.(type) {
		case TextInputRow:
			r := row.(TextInputRow)
			if f.cursor == i {
				if r.isEditing() {
					line = r.View()
				} else {
					line = selectedStyle.Render(r.View())
				}
			} else {
				line = normalStyle.Render(r.View())
			}
		case SelectorRow:
			r := row.(SelectorRow)
			if f.cursor == i {
				if r.isEditing() {
					line = r.View()
				} else {
					line = selectedStyle.Render(r.View())
				}
			} else {
				line = normalStyle.Render(r.View())
			}
		}

		doc.WriteString(line)
		doc.WriteString("\n")
	}

	return doc.String()
}
