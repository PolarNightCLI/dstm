package cmd

import (
	"fmt"
	"os"

	l10n "github.com/qaqland/dstm/localization"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var local = l10n.NewLocalizer()

func runAPP(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Usage()
		os.Exit(1)
	}
	fmt.Println("Let's play don't strave together!")
}

var rootCmd = &cobra.Command{
	Use:     "DSTM",
	Version: "v0.0.1",
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
}
