package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli/v2"
)

func main(){
	app := &cli.App{
        Name:  "greet",
        Usage: "fight the loneliness!",
        Action: func(*cli.Context) error {
            fmt.Println("Hello friend!")
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}