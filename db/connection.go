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
func CreateDB() {
	file, err := os.Create("db/todoDB.db")
	if err != nil {
		logger.AddLog(fmt.Sprintf("ERROR: %v", err.Error()))
		log.Fatal(err.Error())
	}

	file.Close()
	fmt.Println("CONSOLE: Create database successfully")
}

// CreateTable will create DataBase Table if there's no db.
func CreateTable(db *sql.DB) {
	createTodoTable := `CREATE TABLE todoTable(
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
	}

	statement.Exec()
	fmt.Println("CONSOLE: Create Table successfully")
}
