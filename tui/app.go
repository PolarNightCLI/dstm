package tui

import (
	"fmt"

	"github.com/PolarNightCLI/dstm/config"
	l10n "github.com/PolarNightCLI/dstm/localization"
	"github.com/PolarNightCLI/dstm/tui/header"
	tea "github.com/charmbracelet/bubbletea"
)

var local = l10n.Singleton()

func newHeader(t string) header.Header {
	tips := []string{
		local.String("_tip01", l10n.MsgOnly, 0, nil),
		local.String("_tip02", l10n.MsgOnly, 0, nil),
	}
	shardsName := []string{
		"a-01", "a-02",
	}
	return header.NewHeader(t, tips, shardsName)
}

func NewTuiApp(name, ver string, conf *config.Config) tuiApp {
	return tuiApp{
		name:      name,
		version:   ver,
		appConf:   conf,
		headBlock: newHeader(name + " " + ver),
	}
}

type tuiApp struct {
	name    string
	version string
	appConf *config.Config

	headBlock header.Header
}

func (app tuiApp) Init() tea.Cmd {
	return app.headBlock.Init()
}

func (app tuiApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			fmt.Println()
			return app, tea.Quit
		}
	}
	h, cmd := app.headBlock.Update(msg)
	app.headBlock = h
	return app, cmd
}

func (app tuiApp) View() string {
	return app.headBlock.View()
}
