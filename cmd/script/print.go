package script

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"os"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all the scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := database.NewDatabase()
		if err != nil {
			return err
		}

		scripts, err := db.Scripts()
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Name", "Category"})

		for i, script := range scripts {
			t.AppendRow(table.Row{i, script.Name, script.Category})
		}

		t.SetTitle("Scripts")
		t.Render()

		return nil
	},
}

func init() {
	ScriptCmd.AddCommand(printCmd)
}
