package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli/v2"
)

func main(){
	app := getCli()

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

func getCli() *cli.App{
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
								Usage: "Id of the task to delete",
								Destination: &taskId,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("Removing task")
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