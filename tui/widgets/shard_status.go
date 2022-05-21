package widgets

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewShardStatus(name string) ShardStatus {
	return ShardStatus{
		shardName: name,
		checking:  true,
		running:   false,
		spinner:   newSpinner(),
	}
}

func newSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("228"))
	return s
}

// tmp task
func checkSomeShard(shard string) tea.Cmd {
	return func() tea.Msg {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(2)
		time.Sleep(time.Duration(n+1) * time.Second)
		if rand.Intn(2) == 0 {
			return ShardStatusMsg("t" + shard)
		}
		return ShardStatusMsg("f" + shard)
	}
}

type ShardStatusMsg string

type ShardStatus struct {
	shardName string
	checking  bool
	running   bool

	spinner spinner.Model
}

func (m *ShardStatus) Reset() {
	m.checking = true
	m.running = false
	m.spinner = newSpinner()
}

func (m ShardStatus) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, checkSomeShard(m.shardName))
}

func (m ShardStatus) Update(msg tea.Msg) (ShardStatus, tea.Cmd) {
	if !m.checking {
		return m, nil
	}
	switch msg := msg.(type) {
	case ShardStatusMsg:
		str := string(msg)
		if str[1:] == m.shardName {
			m.checking = false
			if str[0:1] == "t" {
				m.running = true
			}
		}
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m ShardStatus) View() string {
	prefix := ""
	if m.checking {
		prefix = m.spinner.View()
	} else {
		if m.running {
			prefix = lipgloss.NewStyle().Foreground(lipgloss.Color("112")).Render("⣿")
		} else {
			prefix = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Render("⣿")
		}
	}

	s := lipgloss.NewStyle().Foreground(lipgloss.Color("30")).Render(m.shardName)
	return fmt.Sprintf(" %s %s", prefix, s)
}
