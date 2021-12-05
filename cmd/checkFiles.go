package cmd

import (
	"TodoCmd/db"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// CheckFiles will check your db is existed or not; if not exist , will create Database and Table
func CheckFiles() {
	_, err := os.Stat("db/todoDB.db")
	if os.IsNotExist(err) {
		db.CreateDB()

		sqLiteDB, _ := sql.Open("sqlite3", "db/todoDB.db")
		defer func(sqLiteDB *sql.DB) {
			err := sqLiteDB.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}(sqLiteDB)

		db.CreateTable(sqLiteDB)
	}
}
