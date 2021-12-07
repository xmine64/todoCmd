package main

import (
	"TodoCmd/cmd"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

func main() {
	// declare variables
	status := os.Args[1]
	if status == "" {
		cmd.StartAscII()
	}

	var valueFlag = flag.String("value", "show", "it's just value dude :/")
	var idFlag = flag.Int("id", 0, "id for delete row ID parameter")
	flag.Parse()

	//CheckFiles will checking database and log
	cmd.CheckFiles()

	// open database
	db, err := sql.Open("sqlite3", "db/todoDB.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if status == "ADD" || status == "add" || status == "Add" {
		cmd.AddAscII()
		cmd.AddObject(db, *valueFlag)

	} else if status == "SHOW" || status == "show" || status == "Show" {
		cmd.ShowAscII()
		cmd.Show(db)

	} else if status == "DELETE" || status == "Delete" || status == "delete" {
		cmd.DeleteAscII()
		cmd.DeleteByID(db, *valueFlag)

	} else if status == "REPLACE" || status == "Replace" || status == "replace" {
		cmd.ReplaceAscII()
		cmd.ReplaceByID(db, *valueFlag, strconv.Itoa(*idFlag))
	}
}
