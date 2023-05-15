package backup

import (
	"encoding/gob"
	"fmt"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"os"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		f, err := os.Create(output)
		if err != nil {
			return err
		}
		defer f.Close()

		db, err := database.NewDatabase()
		if err != nil {
			return err
		}
		defer db.Close()

		scripts, err := db.Scripts()
		if err != nil {
			return err
		}

		if err := gob.NewEncoder(f).Encode(scripts); err != nil {
			return err
		}

		fmt.Printf("[*] Exported %d scripts to %s\n",
			len(scripts), output)

		return nil
	},
}

func init() {
	exportCmd.Flags().StringP("output", "o", "backup.frider", "name of the file where to store backup")
	BackupCmd.AddCommand(exportCmd)
}
