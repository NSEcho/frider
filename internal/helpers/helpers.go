package helpers

import (
	"os"
	"path/filepath"
)

func CreateDatabasePath(dbName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	friderDir := filepath.Join(home, "frider")
	if err := os.Mkdir(friderDir, os.ModePerm); os.IsNotExist(err) {
		return "", err
	}
	return filepath.Join(friderDir, dbName), nil
}
