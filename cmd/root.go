package cmd

import (
	"fmt"
	"os"

	"github.com/PolarNightCLI/dstm/config"
	"github.com/PolarNightCLI/dstm/dst"
	l10n "github.com/PolarNightCLI/dstm/localization"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PolarNightCLI/dstm/tui"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

const (
	appName = "DSTM"
	version = "v0.0.1"
)

var (
	appConf config.Config = config.LoadConfig()
	local                 = l10n.Singleton()
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runAPP(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Usage()
		os.Exit(1)
	}
	dst.Main()
	os.Exit(0)
	p := tea.NewProgram(tui.NewTuiApp(appName, version, &appConf), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func matchLangTag(str string) language.Tag {
	var matcher = language.NewMatcher([]language.Tag{
		language.English,
		language.Chinese,
		language.Japanese,
	})
	tag, _ := language.MatchStrings(matcher, str)
	return tag
}

var rootCmd = &cobra.Command{
	Use:     appName,
	Version: version,
	Short:   "",
	Long:    "",
	Args:    cobra.MinimumNArgs(0),
	Run:     runAPP,
}

func init() {
	l10n.Locale = matchLangTag(appConf.Common.Lang)
	rootCmd.Short = local.String("_short_des", l10n.MsgOnly, 0, nil)
	rootCmd.Long = local.String("_long_des", l10n.MsgOnly, 0, nil)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// cobra.Command 実行前の初期化処理を定義する
	// rootCmd.Execute > コマンドライン引数の処理 > cobra.OnInitialize > rootCmd.Run という順に実行される
	// cobra.OnInitialize(func() {})
}
