package main

import (
	"fmt"
	"time"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/world"
)

func main() {

	for i := 0; i < 20; i++ {
		dp := dice.Roll("2d6")
		fmt.Println("before Boon:", dp.Result())
		dp.Boon()
		fmt.Println(" after Boon:", dp.Result())
		time.Sleep(40 * time.Millisecond)
		fmt.Println(" ------------------------ ")
	}

	worl := world.NewWorld("Drinax").SetUWP(profile.RandomUWP())
	fmt.Println(worl)
	r := dice.Roll("3d6")
	if ok := r.ResultTN(15); ok {
		fmt.Println("Result positive")

	}
	fmt.Println(r.Result())
	fmt.Println("End")

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
