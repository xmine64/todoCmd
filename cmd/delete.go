package cmd

import (
	"database/sql"
	"log"
)

func DeleteByID(db *sql.DB, ID string) {
	insertQuery := `DELETE FROM todoTable WHERE id = ?`
	statement, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(ID)
	if err != nil {
		log.Fatal(err.Error())
	}
}
