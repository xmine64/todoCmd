package cmd

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// ReplaceByID will replace your Task By ID
func ReplaceByID(db *sql.DB, object string, id int) error {
	insertQuery := `UPDATE todoTable SET object = ? WHERE id = ?`
	statement, err := db.Prepare(insertQuery)
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: %v", err))
		return err
	}

	if _, err = statement.Exec(object, id); err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: %v", err))
		return err
	}

	logger.AddLog(fmt.Sprintf("Replace Func Used ID: %v, value: %v", id, object))
	return nil
}
