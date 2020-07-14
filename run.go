package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/profile"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

func main() {
	utils.RandomSeed()

	world1 := world.NewWorld("Arnora")
	world1.SetValue(constant.PrAtmo, "5")
	trueUwp := profile.Compose("UWP", &world1)
	fmt.Println(trueUwp)
	uwp := profile.RandomUWP()
	fmt.Println(uwp)
	world1.MergeUWP(uwp)
	fmt.Println(world1, world1.UWP())
	fmt.Println("------------------")
	trueUwp = profile.Compose("UWP", &world1)
	fmt.Println(trueUwp)
	fmt.Println("-----------------")
	world1.SecondSurvey()
	fmt.Println(world1)
	// //system.AllNames("Sol")

	//SER
}

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
