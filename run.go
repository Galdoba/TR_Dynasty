package main

import (
	"fmt"

	law "github.com/Galdoba/TR_Dynasty/Law"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/world"
)

func main() {

	fmt.Println("Test:")

	fmt.Println("\033[1m BOLD \033[0m")
	fmt.Print("\033c")
	fmt.Print("\033c\n")
	fmt.Print("\x1bc")
	fmt.Print("\x1bc\n")
	fmt.Println("\033c")
	fmt.Println(" dark_yellow ")

	fmt.Println("End Test:")
	return
	uwp := profile.RandomUWP()
	w := world.FromUWP(uwp)

	fmt.Println(w.SecondSurvey())
	lr, _ := law.New(uwp)
	fmt.Println(lr)
	fmt.Println(lr.ULP())
	fmt.Println("")
	fmt.Println(lr.Report())

	// sp, err := starport.From(uwp)
	// //fmt.Println(sp, err)
	// fmt.Println(uwp, err)
	// fmt.Println(sp.Info())
	//starport.FullInfo()
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
