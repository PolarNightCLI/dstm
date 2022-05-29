package tui

import (

	//"github.com/charmbracelet/bubbles/textinput"
	"github.com/PolarNightCLI/dstm/tui/widgets"
	tea "github.com/charmbracelet/bubbletea"
)

func Main() {
	rows := []widgets.TextInputRow{
		widgets.NewTextInputRow("房间名字", "鸽子们的摸鱼日常", nil),
		widgets.NewTextInputRow("房间介绍", "咕咕咕!", nil),
		widgets.NewTextInputRow("房间密码", "", nil),
	}
	form := widgets.NewForm(rows)

	p := tea.NewProgram(form)

	if err := p.Start(); err != nil {
		panic(err)
	}
}
