package tabmenu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewTabMenu(lists [][]list.Item, w, h int) TabMenu {
	var tabs []Tab
	var paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)

	for _, l := range lists {
		listModel := list.New(l, itemDelegate{}, w, h)
		listModel.SetFilteringEnabled(false)
		listModel.SetShowHelp(false)
		listModel.SetShowStatusBar(false)
		listModel.SetShowTitle(false)
		listModel.Styles.PaginationStyle = paginationStyle

		newTab := Tab{list: listModel}
		tabs = append(tabs, newTab)
	}
	return TabMenu{
		tabs: tabs,
	}
}

type TabMenu struct {
	tabs       []Tab
	currentTab int
}

func (t TabMenu) Choised() string {
	return t.tabs[t.currentTab].choice
}

func (t TabMenu) Init() tea.Cmd {
	for i := range t.tabs {
		if i == t.currentTab {
			continue
		}
		t.tabs[i].list.Select(-1)
	}
	return nil
}

func (t TabMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	changeTab := func(offset int) {
		index := t.tabs[t.currentTab].list.Index()
		t.tabs[t.currentTab].list.Select(-1)

		t.currentTab += offset
		if t.currentTab < 0 {
			t.currentTab += len(t.tabs)
		}
		if t.currentTab >= len(t.tabs) {
			t.currentTab -= len(t.tabs)
		}

		if index >= len(t.tabs[t.currentTab].list.Items()) {
			index = len(t.tabs[t.currentTab].list.Items()) - 1
		}
		t.tabs[t.currentTab].list.Select(index)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return t, tea.Quit
		case "enter", " ":
			tab, cmd := t.tabs[t.currentTab].Update(msg)
			t.tabs[t.currentTab] = tab
			return t, cmd
		case "up", "down":
			tab, cmd := t.tabs[t.currentTab].Update(msg)
			t.tabs[t.currentTab] = tab
			return t, cmd
		case "left":
			changeTab(-1)
		case "right":
			changeTab(1)
		}
	}

	var cmds []tea.Cmd
	for i, tab := range t.tabs {
		newTab, cmd := tab.Update(nil)
		t.tabs[i] = newTab
		cmds = append(cmds, cmd)
	}
	return t, tea.Batch(cmds...)
}

func (t TabMenu) View() string {
	var tabContents []string
	for _, tab := range t.tabs {
		tabContents = append(tabContents, tab.View())
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabContents...)
}
