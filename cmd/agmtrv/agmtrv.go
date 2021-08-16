package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/TR_Dynasty/cmd/agmtrv/actions"
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
				return actions.SpaceEncounter(c)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "uwp",
					Usage:    "Set UWP to determine generator options (mandatory)",
					Required: true,
				},
				&cli.IntFlag{
					Name:  "days, d",
					Usage: "Set number of days players spend in space. Each day 1D is rolled, if 6 - encounter occurs",
					Value: 0,
				},
			},
		},
		{
			Name:        "ticket",
			Usage:       "Construct random Mercenary Ticket MgT1 S1 p42-62",
			UsageText:   "Set Space Encounter at world's near space. Demands UWP for data generation",
			Description: "LONG DESCR",
			ArgsUsage:   "UWP: string represents world data",
			Category:    "Encounter",
			Action: func(c *cli.Context) error {
				return actions.NewTicket()
			},
		},
		//////////////////////////////////////

	}
	args := os.Args

	if err := app.Run(args); err != nil {
		fmt.Println(err.Error())
	}
}