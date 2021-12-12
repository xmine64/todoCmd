package main

import (
	"TodoCmd/cmd"
	"TodoCmd/db"
	"TodoCmd/logger"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

func main() {
	var database *sql.DB
	var dbPath string
	var objectFlag string
	var idFlag int

	app := &cli.App{
		Name:  "todoCmd",
		Usage: "Manage your tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "db",
				Usage:       "Database path",
				Value:       "db/todoDB.db",
				TakesFile:   true,
				EnvVars:     []string{"TODOCMD_DB"},
				Destination: &dbPath,
			},
		},
		Before: func(c *cli.Context) error {
			if c.NArg() < 1 {
				ascII()
				os.Exit(1)
			}

			// create required files if they don't exist
			if err := cmd.CheckFiles(dbPath); err != nil {
				return err
			}

			// open database
			var err error
			if database, err = sql.Open("sqlite3", dbPath); err != nil {
				return err
			}

			// create table if required
			db.CreateTable(database)

			return nil
		},
		After: func(c *cli.Context) error {
			// close database on exit
			if err := database.Close(); err != nil {
				return err
			}
			return nil
		},
		ExitErrHandler: func(context *cli.Context, err error) {
			if err != nil {
				logger.AddLog(fmt.Sprintf("ERORR: %v", err.Error()))
				log.Fatal(err)
			}
		},
		CommandNotFound: func(c *cli.Context, s string) {
			log.Fatalf(`command "%s" not found`, s)
		},
		Commands: []*cli.Command{
			{
				Name:      "add",
				ArgsUsage: "-object <object>",
				Aliases:   []string{"a"},
				Usage:     "Add a task",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "object",
						Aliases:     []string{"o"},
						Usage:       "task content",
						Required:    true,
						TakesFile:   false,
						Destination: &objectFlag,
					},
				},
				Action: func(c *cli.Context) error {
					if err := cmd.AddObject(database, objectFlag); err != nil {
						return err
					}
					fmt.Println(fmt.Sprintf("CONSOLE: %v added into database successfully", objectFlag))
					return nil
				},
			},
			{
				Name:      "show",
				ArgsUsage: "",
				Aliases:   []string{"print", "s"},
				Usage:     "Show tasks",
				Action: func(c *cli.Context) error {
					if err := cmd.Show(database); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:      "delete",
				ArgsUsage: "-id <id>",
				Aliases:   []string{"complete", "remove", "del", "rm", "d"},
				Usage:     "Delete a task",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "id",
						Usage:       "task id",
						DefaultText: "",
						Required:    true,
						Destination: &idFlag,
					},
				},
				Action: func(c *cli.Context) error {
					if err := cmd.DeleteByID(database, idFlag); err != nil {
						return err
					}
					fmt.Println(fmt.Sprintf("%v ID sucessfully deleted", idFlag))
					return nil
				},
			},
			{
				Name:      "replace",
				ArgsUsage: "-id <id> -object <object>",
				Usage:     "Replace a task with something else",
				Aliases:   []string{"edit", "r"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "id",
						Usage:       "task id",
						DefaultText: "",
						Required:    true,
						Destination: &idFlag,
					},
					&cli.StringFlag{
						Name:        "object",
						Aliases:     []string{"o"},
						Usage:       "task content",
						Required:    true,
						TakesFile:   false,
						Destination: &objectFlag,
					},
				},
				Action: func(c *cli.Context) error {
					if err := cmd.ReplaceByID(database, objectFlag, idFlag); err != nil {
						return err
					}
					fmt.Println(fmt.Sprintf("%v ID successfully replace , new Value: %v", idFlag, objectFlag))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func ascII() {
	fmt.Print(`
d888888P                dP  
   88                   88 
   88    .d8888b. .d888b88 .d8888b. 
   88    88'  '88 88'  '88 88'  '88
   88    88.  .88 88.  .88 88.  .88
   dP    '88888P' '88888P8 '88888P' 
   ooooooooooooooooooooooooooooooo // written by dozheiny

CONSOLE : for see help, type ./todocmd -help

`)
}
