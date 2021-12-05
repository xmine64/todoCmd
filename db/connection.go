package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	file, err := os.Create("db/todoDB.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	file.Close()
}

func CreateTable(db *sql.DB) {
	createTodoTable := `CREATE TABLE todoTable(
		"id" integer NOT NULL PRIMARY KEY ,
		"object" TEXT NOT NULL,
		"time" TEXT NOT NULL,
		"date" TEXT NOT NULL
);`

	statement, err := db.Prepare(createTodoTable)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
}
