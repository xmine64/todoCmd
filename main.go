package main

import (
	"TodoCmd/cmd"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(start())

	//CheckFiles will checking database and log
	cmd.CheckFiles()

	// declare variables
	var value, status string
	status = os.Args[1]
	value = os.Args[2]

	// open database
	db, err := sql.Open("sqlite3", "db/todoDB.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if status == "ADD" || status == "add" || status == "Add" {
		cmd.AddObject(db, value)
	}
}

//start is some ASCII ART code to show when app is start
func start() string {
	start :=
		`
  _______          _                           
 |__   __|        | |                          
    | | ___     __| | ___     __ _ _ __  _ __  
    | |/ _ \   / _'' |/ _ \   / _'' | '_ \| '_ \ 
    | | (_) | | (_| | (_) | | (_| | |_) | |_) |
    |_|\___/   \__,_|\___/   \__,_| .__/| .__/ 
                                  | |   | |    
                                  |_|   |_|

`

	return start
}
