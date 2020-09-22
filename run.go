package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/entity"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/devtools/cli/user"
)

func main() {
	//npcmakerv2.Test()
	entity.Test()

	// w := pickWorld()

	// hyperjump.StartJumpEvent(w)
	// starport.FullInfo(w)
}

func pickWorld() wrld.World {
	dataFound := false
	for !dataFound {
		input := userInputStr("Enter world's Name, Hex or UWP: ")
		data, err := otu.GetDataOn(input)
		if err != nil {
			fmt.Print("WARNING: " + err.Error() + "\n")
			continue
		}
		w, err := wrld.FromOTUdata(data)
		if err != nil {
			fmt.Print(err.Error() + "\n")
			continue
		}
		return w
	}
	fmt.Println("This must not happen!")
	return wrld.World{}
}

func userInputStr(msg ...string) string {
	for i := range msg {
		fmt.Print(msg[i])
	}
	str, err := user.InputStr()
	if err != nil {
		fmt.Print(err.Error())
		return err.Error()
	}
	return str
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
