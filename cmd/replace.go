package cmd

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// ReplaceByID will replace your Task By ID
func ReplaceByID(db *sql.DB, object string, id string) {
	insertQuery := `UPDATE todoTable SET object = ? WHERE id = ?`
	statement, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err.Error())
		logger.AddLog(fmt.Sprintf("ERROR: %v", err))
	}
	_, err = statement.Exec(object, id)

	logger.AddLog(fmt.Sprintf("Replace Func Used ID: %v, value: %v", id, object))
	fmt.Sprintf("CONSOLE: replace successfully!\n")
}
