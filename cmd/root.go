package cmd

import (
	"fmt"
	"os"

	"github.com/qaqland/dstm/config"
	l10n "github.com/qaqland/dstm/localization"

	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

// tea "github.com/charmbracelet/bubbletea"
// "github.com/qaqland/dstm/tui"

const (
	appName = "DSTM"
	version = "v0.0.1"
)

var (
	appConf config.Config = config.LoadConfig()
	local                 = l10n.NewLocalizer()
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
	//p := tea.NewProgram(tui.NewTuiApp(appName, version))
	//if err := p.Start(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	fmt.Println(appConf.Common)
	fmt.Println(appConf.Path)
	fmt.Println(appConf.Url)
	fmt.Println(appConf.Color)
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
	l10n.Locale = language.Chinese
	rootCmd.Short = local.String("_short_des", l10n.MsgOnly, 0, nil)
	rootCmd.Long = local.String("_long_des", l10n.MsgOnly, 0, nil)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// cobra.Command 実行前の初期化処理を定義する
	// rootCmd.Execute > コマンドライン引数の処理 > cobra.OnInitialize > rootCmd.Run という順に実行される
	// cobra.OnInitialize(func() {})
}
