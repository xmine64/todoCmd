package cmd

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// AddObject will add task and insert into database
func AddObject(db *sql.DB, object string) error {
	t := fmt.Sprintf("%v:%v:%v", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	date := fmt.Sprintf("%v-%v-%v", time.Now().Day(), time.Now().Month(), time.Now().Year())

	insertObject := `INSERT INTO todoTable( object, time, date) VALUES(?, ?, ?)`
	statement, err := db.Prepare(insertObject)
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: error on insert object: %v", err.Error()))
		return err
	}

	if _, err = statement.Exec(object, t, date); err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: error on statement.Exec: %v", err.Error()))
		return err
	}

	logger.AddLog(fmt.Sprintf("AddObject used, value: %v", object))
	return nil
}
