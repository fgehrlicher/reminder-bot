package main

import (
	"fmt"
	"github.com/fgehrlicher/reminder-bot/config"
	"log"
	"os"

	"github.com/urfave/cli"
)

const configFilePath = "./config.yml"

func main() {
	app := cli.NewApp()
	app.Name = "Reminder Bot"
	app.Action = func(c *cli.Context) error {
		conf, err := config.LoadConfig(configFilePath)
		if err != nil {
			return err
		}

		fmt.Println(conf.Database.Path)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}