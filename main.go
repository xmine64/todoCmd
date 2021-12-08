package main

import (
	"TodoCmd/cmd"
	"TodoCmd/logger"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

func main() {

	// declare variables
	var valueFlag string
	var idFlag int
	var statusFlag string

	// Note: Probably you ask yourself why I didn't use os.Args and use flags
	// so , I tell u this, go is shit in os.Args and flag same time , and it didn't work
	// I will update and commit when go fix this shit
	// P.S: this shit didn't fix from go v1.13

	// Parsing Flags
	flag.StringVar(&valueFlag, "object", "if you seen this, it means flag is not working", "it's just value dude :/")
	flag.IntVar(&idFlag, "id", 0, "id for delete row ID parameter")
	flag.StringVar(&statusFlag, "status", "hi", "status will decide what do ")

	flag.Parse()
	flag.Args()

	// open database
	db, err := sql.Open("sqlite3", "db/todoDB.db")
	if err != nil {
		log.Fatal(err.Error())
		logger.AddLog(fmt.Sprintf("ERORR: %v", err.Error()))
	}
	// Use defer for closing when program is done.
	defer db.Close()

	if statusFlag == "ADD" || statusFlag == "add" || statusFlag == "Add" {
		cmd.AddAscII()
		cmd.AddObject(db, valueFlag)

		fmt.Println(fmt.Sprintf("CONSOLE: %v added into database successfully", valueFlag))

	} else if statusFlag == "SHOW" || statusFlag == "show" || statusFlag == "Show" {
		cmd.ShowAscII()
		cmd.Show(db)

	} else if statusFlag == "DELETE" || statusFlag == "Delete" || statusFlag == "delete" {
		cmd.DeleteAscII()
		cmd.DeleteByID(db, strconv.Itoa(idFlag))

		fmt.Println(fmt.Sprintf("%v ID sucessfully deleted", idFlag))

	} else if statusFlag == "REPLACE" || statusFlag == "Replace" || statusFlag == "replace" {
		cmd.ReplaceAscII()
		cmd.ReplaceByID(db, valueFlag, strconv.Itoa(idFlag))

		fmt.Println(fmt.Sprintf("%v ID successfully replace , new Value: %v", idFlag, valueFlag))

	} else {
		cmd.FirstUse()
		cmd.CheckFiles()
	}
}
