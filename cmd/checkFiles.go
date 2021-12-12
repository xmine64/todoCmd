package cmd

import (
	"TodoCmd/db"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// CheckFiles will check your db and log are existed or not; if not exist , will create Database and Table
func CheckFiles(dbPath string) error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if err := db.CreateDB(dbPath); err != nil {
			return err
		}
	}

	if _, err := os.Stat(".log"); os.IsNotExist(err) {
		file, err := os.Create(".log")
		if err != nil {
			return err
		}
		file.Close()
		fmt.Printf("CONSOLE: .log created!\n")
	}

	return nil
}
