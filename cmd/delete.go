package cmd

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"
	"log"
)

// DeleteByID will delete Task by ID
func DeleteByID(db *sql.DB, ID string) {
	insertQuery := `DELETE FROM todoTable WHERE id = ?`
	statement, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err.Error())
		logger.AddLog(fmt.Sprintf("ERROR: %v", err.Error()))
	}

	_, err = statement.Exec(ID)
	if err != nil {
		log.Fatal(err.Error())
		logger.AddLog(fmt.Sprintf("ERROR: %v", err.Error()))
	}
	logger.AddLog(fmt.Sprintf("delete function used, value: %s", ID))

	fmt.Sprintf("CONSOLE: Delete item successfully !\n")
}
