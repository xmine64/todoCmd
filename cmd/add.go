package cmd

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddObject(db *sql.DB, object string) {
	t := string(rune(time.Now().Hour())) + ":" + string(rune(time.Now().Minute())) + ":" + string(rune(time.Now().Second()))
	date := string(rune(time.Now().Day())) + string(rune(time.Now().Month())) + string(rune(time.Now().Year()))

	insertObject := `INSERT INTO todoTable(id, object, time, date) VALUES(?, ?, ?, ?)`
	statement, err := db.Prepare(insertObject)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(0, object, t, date)
	if err != nil {
		log.Fatal(err.Error())
	}
}
