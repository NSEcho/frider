package script

import (
	"errors"
	"fmt"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [script name]",
	Short: "Show specific script",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing script path")
		}

		scriptName := args[0]

		db, err := database.NewDatabase()
		if err != nil {
			return err
		}
		defer db.Close()

		script, found, err := db.ScriptByName(scriptName)
		if err != nil {
			return err
		}

		if !found {
			return errors.New("no script with such name")
		}

		fmt.Printf("Name: %s\n", script.Name)
		fmt.Printf("Category: %s\n", script.Category)
		fmt.Println("========================================")
		fmt.Println(string(script.Content))

		return nil
	},
}

func init() {
	ScriptCmd.AddCommand(showCmd)
}
