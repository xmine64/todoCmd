package cmd

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"
)

// DeleteByID will delete Task by ID
func DeleteByID(db *sql.DB, ID int) error {
	insertQuery := `DELETE FROM todoTable WHERE id = ?`
	statement, err := db.Prepare(insertQuery)
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: %v", err.Error()))
		return err
	}

	_, err = statement.Exec(ID)
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: %v", err.Error()))
		return err
	}
	logger.AddLog(fmt.Sprintf("delete function used, value: %d", ID))

	return nil
}
