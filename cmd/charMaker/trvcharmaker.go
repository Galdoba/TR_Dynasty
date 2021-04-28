package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/TR_Dynasty/pkg/entity"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = "v 0.0.0"
	app.Name = "trvcharmaker"
	app.Usage = "TODO"
	app.Flags = []cli.Flag{}

	app.Commands = []*cli.Command{
		//////////////////////////////////////
		{
			Name:  "new",
			Usage: "Creates New Traveller",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				trv := entity.NewTraveller()
				fmt.Println(trv)
				sheet := trv.Sheet()
				fmt.Print(sheet)
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
	if len(args) < 2 {
		args = append(args, "help") //Принудительно зовем помощь если нет других аргументов
	}
	if err := app.Run(args); err != nil {
		fmt.Println(err.Error())
	}
}
