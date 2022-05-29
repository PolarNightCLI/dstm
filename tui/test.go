package tui

import (

	//"github.com/charmbracelet/bubbles/textinput"
	widgets "github.com/PolarNightCLI/dstm/tui/widgets/form"
	tea "github.com/charmbracelet/bubbletea"
)

func notEmpty(str string) bool {
	return len(str) > 0
}

func Main() {
	rows := []widgets.TextInputRow{
		widgets.NewTextInputRow("房间名字", "请输入房间名字", "鸽子们的摸鱼日常", notEmpty),
		widgets.NewTextInputRow("房间介绍", "请输入房间介绍", "咕咕咕!", nil),
		widgets.NewTextInputRow("房间密码", "请输入房间密码", "", nil),
	}
	form := widgets.NewForm(rows)

	p := tea.NewProgram(form)

	if err := p.Start(); err != nil {
		panic(err)
	}
}
