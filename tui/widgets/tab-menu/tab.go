package tabmenu

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabItem string

func (i TabItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(TabItem)
	if !ok {
		return
	}

	var (
		itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
		selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("214"))
	)
	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type Tab struct {
	list   list.Model
	choice string
}

func (t Tab) Init() tea.Cmd {
	return nil
}

func (t Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	changeRow := func(offset int) (Tab, tea.Cmd) {
		index := t.list.Index() + offset
		maxIndex := len(t.list.Items()) - 1
		if index < 0 {
			t.list.Select(maxIndex)
		} else if index > maxIndex {
			t.list.Select(0)
		} else {
			t.list.Select(index)
		}
		return t, nil
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.list.SetWidth(msg.Width)
		return t, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter", " ":
			i, ok := t.list.SelectedItem().(TabItem)
			if ok {
				t.choice = string(i)
			}
			return t, tea.Quit
		case "up":
			return changeRow(-1)
		case "down":
			return changeRow(1)
		}
	}

	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)
	return t, cmd
}

func (t Tab) View() string {
	return t.list.View()
}
