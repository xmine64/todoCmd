package db

import (
	"TodoCmd/logger"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// CreateDB will create DataBase if not exist
func CreateDB(dbPath string) error {
	file, err := os.Create(dbPath)
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: CreateDB %v", err.Error()))
		return err
	}

	file.Close()
	fmt.Println("CONSOLE: Create database successfully")
	return nil
}

// CreateTable will create DataBase Table if there's no db.
func CreateTable(db *sql.DB) error {
	createTodoTable := `CREATE TABLE IF NOT EXISTS todoTable(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		object TEXT NOT NULL,
		time TEXT NOT NULL,
		date TEXT NOT NULL,
		UNIQUE (id)
);
`

	statement, err := db.Prepare(createTodoTable)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	if _, err := statement.Exec(); err != nil {
		return err
	}

	return nil
}
