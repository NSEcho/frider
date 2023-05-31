package script

import (
	"bytes"
	"errors"
	"github.com/apex/log"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var saveCmd = &cobra.Command{
	Use:   "save [script path]",
	Short: "Save script",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing script path")
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		if name == "" {
			return errors.New("script name cannot be empty")
		}

		category, err := cmd.Flags().GetString("category")
		if err != nil {
			return err
		}

		if category == "" {
			return errors.New("script category cannot be empty")
		}

		scriptPath := args[0]
		f, err := os.Open(scriptPath)
		if err != nil {
			return err
		}
		defer f.Close()

		db, err := database.NewDatabase()
		if err != nil {
			return err
		}
		defer db.Close()

		buff := new(bytes.Buffer)
		if _, err := io.Copy(buff, f); err != nil {
			return err
		}

		if err := db.Save(database.Script{
			Name:     name,
			Category: category,
			Content:  buff.Bytes(),
		}); err != nil {
			return err
		}

		log.Infof("Saved script from %s with name %s and category %s",
			scriptPath, name, category)

		return nil
	},
}

func init() {
	saveCmd.Flags().StringP("name", "n", "", "name of the script")
	saveCmd.Flags().StringP("category", "c", "", "script category")
	ScriptCmd.AddCommand(saveCmd)
}
