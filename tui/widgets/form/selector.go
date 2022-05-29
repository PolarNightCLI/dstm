package form

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type finishSelectMsg struct{}

func finishSelectCmd() tea.Msg {
	return finishSelectMsg{}
}

func NewSelector(opts []string, allowCancel, allowMultiSelect, listVertical bool) Selector {
	return Selector{
		options:          opts,
		selected:         map[int]struct{}{},
		allowCancel:      allowCancel,
		allowMultiSelect: allowMultiSelect,
		listVertical:     listVertical,
	}
}

type Selector struct {
	options  []string
	selected map[int]struct{}
	cursor   int

	allowCancel      bool
	allowMultiSelect bool
	listVertical     bool
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

func (s Selector) scalMaxIndex() int {
	maxIndex := len(s.options) - 1

	if s.allowCancel {
		maxIndex += 2
	} else {
		maxIndex += 1
	}

	return maxIndex
}

func (s *Selector) checkOption(index int) {
	if _, ok := s.selected[index]; ok {
		delete(s.selected, index)
	} else {
		if !s.allowMultiSelect && len(s.selected) > 0 {
			s.selected = map[int]struct{}{}
		}
		s.selected[index] = struct{}{}
	}
}

func (s Selector) Update(msg tea.Msg) (Selector, tea.Cmd) {
	maxIndex := s.scalMaxIndex()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if s.cursor == maxIndex { // ok
				if len(s.selected) > 0 {
					return s, finishSelectCmd
				} else {
					return s, nil
				}
			}
			if s.allowCancel && s.cursor == maxIndex-1 { // cancel
				s.selected = map[int]struct{}{}
				return s, tea.Quit
			}
			s.checkOption(s.cursor)
		case "left", "up":
			s.cursor -= 1
			if s.cursor < 0 {
				s.cursor = maxIndex
			}
		case "right", "down":
			s.cursor += 1
			if s.cursor > maxIndex {
				s.cursor = 0
			}
		}
	}
	return s, nil
}

func (s Selector) View() string {
	normalStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("100")).PaddingRight(3)
	selectedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("200")).PaddingRight(3)

	var opts []string
	var optsBlock string
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
	if s.listVertical {
		optsBlock = lipgloss.JoinVertical(lipgloss.Center, opts...)
	} else {
		optsBlock = lipgloss.JoinHorizontal(lipgloss.Center, opts...)
	}

	opts = []string{}
	var actionsBlock string
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
	actionsBlock = lipgloss.JoinHorizontal(lipgloss.Center, opts...)

	if s.listVertical {
		return lipgloss.JoinVertical(lipgloss.Center, optsBlock, actionsBlock)
	} else {
		return lipgloss.JoinHorizontal(lipgloss.Center, optsBlock, actionsBlock)
	}
}
