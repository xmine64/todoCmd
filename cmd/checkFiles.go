package cmd

import (
	"TodoCmd/db"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CheckFiles() {
	_, err := os.Stat("db/todoDB.db")
	if os.IsNotExist(err) {
		db.CreateDB()

		sqLiteDB, _ := sql.Open("sqlite3", "db/todoDB.db")
		defer sqLiteDB.Close()

		db.CreateTable(sqLiteDB)
	}
}
