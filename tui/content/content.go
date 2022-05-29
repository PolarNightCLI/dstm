package content

import (
	"fmt"

	l10n "github.com/PolarNightCLI/dstm/localization"
	tea "github.com/charmbracelet/bubbletea"
)

var local = l10n.Singleton()

func NewContent() {
	fmt.Println("new content")
}

type Content struct {
	options map[string][]string
}

func (c Content) Init() tea.Cmd {
	return nil
}

func (c Content) Update(msg tea.Msg) (Content, tea.Cmd) {
	return c, nil
}

func (c Content) View() string {
	return "content"
}
