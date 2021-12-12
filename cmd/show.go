package cmd

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Show will show your task u added!
func Show(db *sql.DB) error {
	row, err := db.Query("SELECT * FROM todoTable")
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: %v", err.Error()))
		return err
	}

	defer row.Close()

	for row.Next() {
		var id int
		var object string
		var time string
		var date string

		if err := row.Scan(&id, &object, &time, &date); err != nil {
			return err
		}

		log.Println(":", id, object, time, date)
		fmt.Println("--------------------------------------")
	}

	logger.AddLog("show function Used!")
	return nil
}
