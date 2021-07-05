package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/TR_Dynasty/cmd/autoGM/actions"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "v 0.0.0"
	app.Name = "agmtrv"
	app.Usage = "Collection of Referee tools for Mongoose Traveller 2E"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{},
	}
	app.Commands = []cli.Command{
		//////////////////////////////////////
		{
			Name:        "spaceencounter",
			Usage:       "Roll random Ship encounter MgT2 CRB",
			UsageText:   "Set Space Encounter at world's near space. Demands UWP for data generation",
			Description: "LONG DESCR",
			ArgsUsage:   "UWP: string represents world data",
			Category:    "Encounter",
			Action: func(c *cli.Context) error {
				fmt.Println("Start Space Encounter")
				return actions.SpaceEncounter(c)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "uwp",
					Usage:       "",
					EnvVar:      "",
					FilePath:    "",
					Required:    true,
					Value:       "",
					Destination: new(string),
				},
			},
		},
		//////////////////////////////////////

	}
	args := os.Args

	if err := app.Run(args); err != nil {
		fmt.Println(err.Error())
	}
}
