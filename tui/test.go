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
		form.NewTextInputRow("服务器名", "请输入服务器名", "鸽子们的摸鱼日常", notEmpty),
		form.NewTextInputRow("服务器介绍", "请输入服务器介绍", "咕咕咕!", nil),
		form.NewTextInputRow("服务器密码", "请输入服务器密码", "", nil),
		form.NewSelectorRow("服务器语言", "中文", []string{"中文", "English", "日本語"}),
		form.NewSelectorRow("游戏风格", "合作", []string{"合作", "休闲", "竞赛", "疯狂"}),
		form.NewSelectorRow("游戏模式", "无尽", []string{"生存", "无尽", "荒野"}),
	}
	form := form.NewForm(rows)

	p := tea.NewProgram(form)

	if err := p.Start(); err != nil {
		panic(err)
	}
}
