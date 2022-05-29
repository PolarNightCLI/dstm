package header

import (
	"os"
	"strings"

	l10n "github.com/PolarNightCLI/dstm/localization"
	widgets "github.com/PolarNightCLI/dstm/tui/widgets"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var local = l10n.Singleton()

func NewHeader(t string, ts, shardsName []string) Header {
	var ss []widgets.ShardStatus
	for _, name := range shardsName {
		ss = append(ss, widgets.NewShardStatus(name))
	}
	return Header{
		title:      t,
		tips:       ts,
		shards:     ss,
		checkedNum: 0,
	}
}

type Header struct {
	title string
	tips  []string

	shards     []widgets.ShardStatus
	checkedNum int
}

func (h *Header) AddNewShard(name string) {
	newShard := widgets.NewShardStatus(name)
	h.shards = append(h.shards, newShard)
}

func (h *Header) RefreshAll() tea.Cmd {
	h.checkedNum = 0
	var cmds []tea.Cmd
	for i := range h.shards {
		h.shards[i].Reset()
		cmds = append(cmds, h.shards[i].Init())
	}
	return tea.Batch(cmds...)
}

func (h Header) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, s := range h.shards {
		cmds = append(cmds, s.Init())
	}
	return tea.Batch(cmds...)
}

func (h Header) Update(msg tea.Msg) (Header, tea.Cmd) {
	if len(h.shards) == 0 || len(h.shards) == h.checkedNum {
		return h, nil
	}

	var cmds []tea.Cmd
	for i, s := range h.shards {
		sm, sc := s.Update(msg)
		h.shards[i] = sm
		cmds = append(cmds, sc)
	}
	return h, tea.Batch(cmds...)
}

func (h Header) View() string {
	w, _, _ := term.GetSize(int(os.Stdout.Fd()))
	var border01 = lipgloss.Border{Top: "=", Bottom: "-", Left: "│", Right: "│", TopLeft: "╭", TopRight: "╮", BottomLeft: "┘", BottomRight: "└"}
	var border02 = lipgloss.Border{Top: "─", Bottom: "─", Left: "│", Right: "│", TopLeft: "├", TopRight: "┤", BottomLeft: "├", BottomRight: "┤"}
	var border03 = lipgloss.Border{Top: "─", Bottom: "=", Left: "│", Right: "│", TopLeft: "╭", TopRight: "╮", BottomLeft: "└", BottomRight: "┘"}
	basicBorder := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("63")).
		Width(w - 2)

	var headerBlock strings.Builder

	version := basicBorder.Copy().
		Border(border01, true, true, false).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("110")).
		Render(h.title)
	headerBlock.WriteString(version + "\n")

	tips := lipgloss.JoinVertical(lipgloss.Left, h.tips...)
	tipsBlock := basicBorder.Copy().
		Border(border02, true, true, true, true).
		PaddingLeft(2).PaddingRight(2).
		Foreground(lipgloss.Color("208")).
		Render(tips)
	headerBlock.WriteString(tipsBlock + "\n")

	length := 0
	var shards strings.Builder
	if len(h.shards) == 0 {
		shards.WriteString(local.String("_no_running_shards", l10n.MsgOnly, 0, nil))
	} else {
		v := h.shards[0].View()
		length += lipgloss.Width(v)
		shards.WriteString(v)
	}
	for i := 1; i < len(h.shards); i++ {
		v := h.shards[i].View()
		wl := lipgloss.Width(v)
		if length+wl+3 >= w {
			shards.WriteString("\n")
			length = 0
		} else {
			shards.WriteString("   ")
			length += 3
		}
		shards.WriteString(v)
		length += wl
	}

	shardsBlock := basicBorder.Copy().
		Border(border03, false, true, true).
		Foreground(lipgloss.Color("30")).
		Align(lipgloss.Center).
		Render(shards.String())
	headerBlock.WriteString(shardsBlock)

	return headerBlock.String()
}
