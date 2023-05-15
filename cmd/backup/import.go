package backup

import (
	"encoding/gob"
	"fmt"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"os"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import exported .frider scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}

		f, err := os.Open(input)
		if err != nil {
			return err
		}
		defer f.Close()

		var scripts []database.Script

		if err := gob.NewDecoder(f).Decode(&scripts); err != nil {
			return err
		}

		db, err := database.NewDatabase()
		if err != nil {
			return err
		}
		defer db.Close()

		for _, script := range scripts {
			if err := db.Save(script); err != nil {
				return err
			}
		}

		fmt.Printf("[*] Imported %d scripts from %s\n",
			len(scripts), input)

		return nil
	},
}

func init() {
	importCmd.Flags().StringP("input", "i", "backup.frider", "backup file to load")
	BackupCmd.AddCommand(importCmd)
}
