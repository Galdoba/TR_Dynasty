package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func main() {
	//gui.Test()
	//entity.Test()
	data, err := otu.GetDataOn("Borite")
	fmt.Println("DATA:", data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DATA2:", data.Info)
	w, err := wrld.FromOTUdata(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w.SecondSurvey())
	dist := 1000 * w.GetСharacteristic("Size").Value()
	fmt.Println(dist)
	// uwp := w.UWP()
	// fmt.Println("WORLD", w.PBG())
	// fmt.Println(w.SecondSurvey())

	// fmt.Println(law.Describe(uwp))

	// sp, err := starport.From(w)
	// fmt.Println(sp, err)
	// fmt.Println(uwp, err)
	// fmt.Println(sp.Info())
	// starport.FullInfo()

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
