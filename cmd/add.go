package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// AddObject will add task and insert into database
func AddObject(db *sql.DB, object string) {
	t := fmt.Sprintf("%v:%v:%v", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	date := fmt.Sprintf("%v-%v-%v", time.Now().Day(), time.Now().Month(), time.Now().Year())

	insertObject := `INSERT INTO todoTable( object, time, date) VALUES(?, ?, ?)`
	statement, err := db.Prepare(insertObject)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(object, t, date)
	if err != nil {
		log.Fatal(err.Error())
	}
}
