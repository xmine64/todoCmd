package main

import (
	"TodoCmd/cmd"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(start())

	//CheckFiles will checking database and log
	cmd.CheckFiles()

	// declare variables
	status := os.Args[1]
	value := os.Args[2]

	var IdFlag = flag.Int("id", 0, "id for delete row ID parameter")
	flag.Parse()

	// open database
	db, err := sql.Open("sqlite3", "db/todoDB.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if status == "ADD" || status == "add" || status == "Add" {
		cmd.AddObject(db, value)
	} else if status == "SHOW" || status == "show" || status == "Show" {
		cmd.Show(db)
	} else if status == "DELETE" || status == "Delete" || status == "delete" {
		cmd.DeleteByID(db, value)
	} else if status == "REPLACE" || status == "Replace" || status == "replace" {
		cmd.ReplaceByID(db, value, strconv.Itoa(*IdFlag))
	}
}

//start is some ASCII ART code to show when app is start
func start() string {
	start :=
		`
   __            .___      
_/  |_  ____   __| _/____  
\   __\/  _ \ / __ |/  _ \ 
 |  | (  <_> ) /_/ (  <_> )
 |__|  \____/\____ |\____/ 
                  \/

`
	return start
}
