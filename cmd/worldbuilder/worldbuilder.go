package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = "v 0.0.0"
	app.Name = "worldbuilder"
	app.Usage = "Generates maximum possible data about starsystem, stars, planets and satellites based on T5 Book3 rules for system generation, and MT World Builder's handbook rules for planetary details generation"
	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		//////////////////////////////////////
		{
			Name:  "build",
			Usage: "Generates Data structure containing all data about World",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				ar := c.Args()
				fmt.Println("args num:", len(ar))
				for i := range ar {
					fmt.Println("arg", i+1, ar[i])
				}
				// for i := range trv.Skill {
				// 	//fmt.Println(trv.Skill[i].Name(), "- picked as Background")
				// }
				//придумать как красиво образаться к ассетам
				//trv.Call["STR"].Value()?
				return nil
			},
		},
		//////////////////////////////////////
	}
	args := os.Args

	if err := app.Run(args); err != nil {
		fmt.Println(err.Error())
	}
}
