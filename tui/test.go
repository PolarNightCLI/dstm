package tui

import (

	//"github.com/charmbracelet/bubbles/textinput"
	"github.com/PolarNightCLI/dstm/tui/widgets/form"
	tea "github.com/charmbracelet/bubbletea"
)

func notEmpty(str string) bool {
	return len(str) > 0
}

func Main() {
	//rows := []form.TextInputRow{
	//	form.NewTextInputRow("房间名字", "请输入房间名字", "鸽子们的摸鱼日常", notEmpty),
	//	form.NewTextInputRow("房间介绍", "请输入房间介绍", "咕咕咕!", nil),
	//	form.NewTextInputRow("房间密码", "请输入房间密码", "", nil),
	//}
	//rows := []form.SelectorRow{
	//	form.NewSelectorRow("选项A", "a", []string{"a", "aa", "aaa"}),
	//	form.NewSelectorRow("选项B", "b", []string{"b", "bb", "bbb"}),
	//	form.NewSelectorRow("选项C", "c", []string{"c", "cc", "ccc"}),
	//}
	rows := []any{
		form.NewTextInputRow("房间名字", "请输入房间名字", "鸽子们的摸鱼日常", notEmpty),
		form.NewTextInputRow("房间介绍", "请输入房间介绍", "咕咕咕!", nil),
		form.NewTextInputRow("房间密码", "请输入房间密码", "", nil),
		form.NewSelectorRow("选项A", "a", []string{"a", "aa", "aaa"}),
		form.NewSelectorRow("选项B", "b", []string{"b", "bb", "bbb"}),
		form.NewSelectorRow("选项C", "c", []string{"c", "cc", "ccc"}),
	}
	form := form.NewForm(rows)

	p := tea.NewProgram(form)

	if err := p.Start(); err != nil {
		panic(err)
	}
}
