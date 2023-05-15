package script

import (
	"errors"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [script name]",
	Short: "Delete script from the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing script name")
		}
		scriptName := args[0]

		db, err := database.NewDatabase()
		if err != nil {
			return err
		}
		defer db.Close()

		return db.Delete(scriptName)
	},
}

func init() {
	ScriptCmd.AddCommand(deleteCmd)
}
