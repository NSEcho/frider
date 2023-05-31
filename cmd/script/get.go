package script

import (
	"errors"
	"github.com/apex/log"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get [script name]",
	Short: "Get content of the script as file",
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

		script, found, err := db.ScriptByName(scriptName)
		if err != nil {
			return err
		}

		if !found {
			return errors.New("no script with such name")
		}

		outputFilename, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		f, err := os.Create(outputFilename)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(script.Content); err != nil {
			return err
		}

		log.Infof("Saved %s to %s", scriptName, outputFilename)

		return nil
	},
}

func init() {
	getCmd.Flags().StringP("output", "o", "script.js", "Where to save the script")
	ScriptCmd.AddCommand(getCmd)
}
