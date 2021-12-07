package cmd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ReplaceByID(db *sql.DB, object string, id string) {
	insertQuery := `UPDATE todoTable SET object = ? WHERE id = ?`
	statement, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec(object, id)
}
