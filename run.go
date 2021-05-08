package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/pkg/t5trade"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func main() {
	fmt.Println("Select Source World:")
	sourceworld := wrld.PickWorld()
	fmt.Println("Select Market World:")
	marketworld := wrld.PickWorld()

	cargo := t5trade.NewCargo(sourceworld.GetСharacteristic(constant.PrTL).Value(), sourceworld.TradeCodes())
	sell := t5trade.SellPrice(cargo, marketworld.GetСharacteristic(constant.PrTL).Value(), sourceworld.TradeCodes())
	fmt.Println(cargo)
	fmt.Println(sell)
	//t5trade.Test()

	//fmt.Println(astronomical.NewStellarData("G2 V"))
	//os.Exit(3)
	//starsystem.Test()
	// world := wrld.PickWorld()
	// ssData := starsystem.From(world)
	// ssData.PrintTable()

	//w := wrld.PickWorld()
	//fmt.Println(w.SecondSurvey())
	//stats := spaceship.Statistics("Gazelle")
	//fmt.Println(stats)
	//autoGM.AutoGM()
	// plnt := wrld.PickWorld()
	// st, err := starport.From(plnt)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(st.Info())
	//uwp := profile.NewUWP(plnt.UWP())
	//encounter.EncounterTable(uwp.String())

	//mission.Test()

	//ehex.TestEhex()
	//dynasty.Test4()

	//skimming.Test()
	//routine.StartRoutine()
	// fmt.Print("Enter Career: ")
	// carArgs, err := user.InputStr()
	// if err != nil {
	// 	panic(err)
	// }
	// trv := npcmakerv2.NewTraveller(plnt, dice.New(0).RollFromList(npcmakerv2.SearchCareers(carArgs)))
	// fmt.Println(trv.String())
	//w := pickWorld()
	//entity.Test()
	//hyperjump.StartJumpEvent(w)
	// starport.FullInfo(w)

	//autoGM.AutoGM()
}

/*
B34+3
11 - Resorses
3 - Labor
4 - Infrastructure
3 - Culture

Base Demand: 11
Total Demand: 4   [2d6-1-2]

Export: 11-4 = 7
Export Benefit: 0.5

Resource Available := 3.5 + 4 = 7.5

Labor Base = 0.0001 * 7

GWP = ((14*0.1*7.5)*(0.0001*7)*4) / (3+1)
GWP = 10.5 * 0.0007 * 4 / 4
GWP = 0.007 RU

Trading Finished Goods = 1.106
Aggregated Demand = 1.0

Final GWP = 0.008 RU



*/

//OB Ia Ia Ia II II II II
/*
 ___     ___     ___     ___
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
/   \___/   \___/   \___/   \___
\___/   \___/   \___/   \___/
*/
