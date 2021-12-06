package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Show(db *sql.DB) {
	row, err := db.Query("SELECT * FROM todoTable")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()
	for row.Next() {
		var id int
		var object string
		var time string
		var date string

		row.Scan(&id, &object, &time, &date)
		fmt.Sprintf("Number :%v, object: %v, time added: %v, date added: %v", id, object, time, date)
	}
}