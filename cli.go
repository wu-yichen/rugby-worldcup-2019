package main

import (
	"log"
	"os"

	"github.com/subosito/gotenv"

	"github.com/followme1987/wcwhen/helpers"
	"github.com/urfave/cli"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start()
}

func start() {
	app := cli.NewApp()
	app.Name = "wcwhen"
	app.Usage = "A CLI for rugby world cup 2019"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:        "team",
			Usage:       "team -name <teamName>",
			Description: "Retrieve upcoming fixtures of the team indicated",
			Action:      helpers.GetFixturesByTeam,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "The team name",
				},
			},
		},
		{
			Name:        "group",
			Usage:       "group -name <groupName>",
			Description: "Retrieve all members in the group indicated with full and short name",
			Action:      helpers.GetTeamsByGroupName,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "The group name",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
