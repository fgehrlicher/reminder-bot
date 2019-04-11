package main

import (
	"fmt"
	"github.com/fgehrlicher/reminder-bot/conf"
	"github.com/fgehrlicher/reminder-bot/db"
	"log"
	"os"

	"github.com/urfave/cli"
)

const configFilePath = "./conf.yml"

func main() {
	app := cli.NewApp()
	app.Name = "Reminder Bot"
	app.Action = func(c *cli.Context) error {
		config, err := conf.LoadConfig(configFilePath)
		if err != nil {
			return err
		}

		database, err := db.GetConnection(config)
		if err != nil {
			return err
		}

		fmt.Println(config.Database.Path)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
