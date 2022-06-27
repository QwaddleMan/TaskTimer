package main

import (
	"fmt"
	"os"
	"log"
	"database/sql"

	"github.com/urfave/cli/v2"
	_ "github.com/mattn/go-sqlite3"
)

const db_file string = "TaskTimer.db"
const db_create string = `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL,
		name TEXT NOT NULL,
		desc TEXT,
		PRIMARY KEY(id)
	);

	CREATE TABLE IF NOT EXISTS logs(
		id INTEGER NOT NULL,
		task_id INTEGER NOT NULL,
		start DATETIME,
		end DATETIME,
		notes TEXT,
		PRIMARY KEY(id),
		FOREIGN KEY(task_id) REFERENCES tasks(id)
	);
`

func main(){
	db := dbInit()
	app := getCli(db)

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
	db.Close()
}

func dbInit() *sql.DB {
	db, err := sql.Open("sqlite3", db_file)

	if err != nil{
		panic(err)
	}

	if _, err := db.Exec(db_create); err != nil{
		panic(err)
	}

	return db
}

func getCli(db *sql.DB) *cli.App{
	var taskName string
	var taskDesc string
	var taskId int

	app := &cli.App{
        Name:  "TaskTimer",
        Usage: "Create, log, and take notes on time spent on tasks.",
		Commands: []*cli.Command{
            {
                Name:    "task",
                Usage:   "tasks interface",
				Subcommands: []*cli.Command{
                    {
                        Name:  "list",
                        Usage: "List all tasks",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("listing tasks")
                            return nil
                        },
					},
					{
                        Name:  "create",
                        Usage: "Create new task.",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Aliases: []string{"n"},
								Usage: "Name of the new task.",
								Destination: &taskName,
							},
							&cli.StringFlag{
								Name: "desc",
								Aliases: []string{"d"},
								Usage: "Describe the task",
								Destination: &taskDesc,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("creating task")
                            return nil
                        },
                    },
					{
                        Name:  "remove",
                        Usage: "Delete a task.",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "ID of the task to delete",
								Destination: &taskId,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("Removing task")
                            return nil
                        },
                    },
					{
                        Name:  "start",
                        Usage: "start a task or resume a paused task.",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "ID of the task to start",
								Destination: &taskId,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("starting task")
                            return nil
                        },
                    },
					{
                        Name:  "stop",
                        Usage: "stop the current task",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("starting task")
                            return nil
                        },
                    },
					{
                        Name:  "pause",
                        Usage: "pause the current task",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("starting task")
                            return nil
                        },
                    },
				},
            },
		},
        Action: func(*cli.Context) error {
			fmt.Println("Welcome to TaskTimer. Use -h for more information on usage.")
            return nil
        },
    }

	return app
}