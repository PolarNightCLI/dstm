package widgets

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewSelector(opts []string, allowCancel bool) Selector {
	return Selector{options: opts, selected: map[int]struct{}{}, allowCancel: allowCancel}
}

type Selector struct {
	options  []string
	selected map[int]struct{}
	cursor   int

	allowCancel bool
}

func (s Selector) Init() tea.Cmd {
	return nil
}

func (s Selector) GetSelected() []string {
	result := []string{}
	for i := 0; i < len(s.options); i++ {
		if _, ok := s.selected[i]; ok {
			result = append(result, s.options[i])
		}
	}
	return result
}

func (s Selector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	maxIndex := len(s.options)
	if s.allowCancel {
		maxIndex++
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			fmt.Println()
			return s, tea.Quit
		case "enter", " ":
			if s.cursor == maxIndex { // ok
				return s, tea.Quit
			}
			if s.allowCancel && s.cursor == maxIndex-1 { // cancel
				s.selected = map[int]struct{}{}
				return s, tea.Quit
			}

			if _, ok := s.selected[s.cursor]; ok {
				delete(s.selected, s.cursor)
			} else {
				s.selected[s.cursor] = struct{}{}
			}
		case "left":
			s.cursor -= 1
			if s.cursor < 0 {
				s.cursor = maxIndex
			}
		case "right":
			s.cursor += 1
			if s.cursor > maxIndex {
				s.cursor = 0
			}
		case "up":
			if s.cursor >= len(s.options) {
				s.cursor -= len(s.options)
			}
		case "down":
			if s.cursor > 0 && s.allowCancel {
				s.cursor = len(s.options) + 1
			} else {
				s.cursor = len(s.options)
			}
		}
	}
	return s, nil
}

func (s Selector) View() string {
	normalStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("100")).PaddingRight(3)
	selectedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("200")).PaddingRight(3)
	var opts []string

	for i, opt := range s.options {
		prefix := "  "
		if i == s.cursor {
			prefix = "> "
		}

		item := fmt.Sprintf("%s%d.%s", prefix, i+1, opt)
		var line string
		if _, ok := s.selected[i]; ok {
			line = selectedStyle.Render(item)
		} else {
			line = normalStyle.Render(item)
		}
		opts = append(opts, line)
	}
	optsBlock := lipgloss.JoinHorizontal(lipgloss.Center, opts...)

	opts = []string{}
	actions := []string{"[Cancel]", "[OK]"}
	for i, action := range actions {
		if !s.allowCancel && i == 0 {
			continue
		}

		var item string
		dontAllowCancelAndSelectOK := !s.allowCancel && s.cursor == len(s.options)
		allowCancelAndSelectAction := s.cursor == len(s.options)+i
		if dontAllowCancelAndSelectOK || allowCancelAndSelectAction {
			item = selectedStyle.Render("> " + action)
		} else {
			item = normalStyle.Render("  " + action)
		}

		opts = append(opts, item)
	}
	actionsBlock := lipgloss.JoinHorizontal(lipgloss.Center, opts...)

	return lipgloss.JoinVertical(lipgloss.Center, optsBlock, actionsBlock)
}
