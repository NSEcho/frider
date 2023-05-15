package script

import (
	"github.com/spf13/cobra"
)

var ScriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Manage database scripts",
}
