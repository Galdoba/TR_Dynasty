package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/devtools/cli/user"

	"github.com/Galdoba/TR_Dynasty/npc/npcmakerv2"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func main() {
	plnt := wrld.PickWorld()

	//uwp := profile.NewUWP(plnt.UWP())
	//encounter.EncounterTable(uwp.String())

	//mission.Test()

	//ehex.TestEhex()
	//dynasty.Test4()

	//skimming.Test()
	//routine.StartRoutine()
	fmt.Print("Enter Career: ")
	carArgs, err := user.InputStr()
	if err != nil {
		panic(err)
	}
	trv := npcmakerv2.NewTraveller(plnt, dice.New(0).RollFromList(npcmakerv2.SearchCareers(carArgs)))
	fmt.Println(trv.String())
	//w := pickWorld()
	//entity.Test()
	//hyperjump.StartJumpEvent(w)
	// starport.FullInfo(w)

	//autoGM.AutoGM()
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
