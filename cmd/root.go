package cmd

import (
	"github.com/nsecho/frider/cmd/app"
	"github.com/nsecho/frider/cmd/backup"
	"github.com/nsecho/frider/cmd/script"
	"github.com/spf13/cobra"

	_ "github.com/nsecho/frider/cmd/script"
)

var rootCmd = &cobra.Command{
	Use:   "frider",
	Short: "Frida helper tool",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(app.AppCmd)
	rootCmd.AddCommand(backup.BackupCmd)
	rootCmd.AddCommand(script.ScriptCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
