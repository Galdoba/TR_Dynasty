package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/world"
)

func main() {
	Trade.Initiate()

	for i := 0; i < 5; i++ {
		sourceWorld := world.FromUWP(profile.RandomUWP(constant.WTpHospitable)).SetName("Mars").UpdateTC()
		//fmt.Println(sourceWorld)
		fmt.Println(sourceWorld.UWP(), sourceWorld.TradeCodes())
		merch := Trade.NewMerchantOn(sourceWorld.TradeCodes())
		fmt.Println("--------------------------", i)
		fmt.Println(merch)
		code := dice.RollD66() + dice.Roll("2d6").SumStr()
		fmt.Println(code)
		if len(merch.DetermineGoodsAvailable()) < 9 {
			break
		}
		contract := merch.ProposeSell(code).SetVolume(3)
		fmt.Println(contract.String())
		contract = contract.Negotiate(2)
		fmt.Println(fmt.Println(contract.String()))

	}
	//merch.DetermineGoodsAvailable()
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
