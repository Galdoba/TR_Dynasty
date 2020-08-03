package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/essc"
)

func main() {
	//Trade.Initiate()

	// sourceWorld := world.FromUWP(profile.RandomUWP(constant.WTpHospitable)).UpdateTC()
	// fmt.Println(sourceWorld)
	// merch := Trade.NewMerchant().SetTraderDice(dice.Roll3D()).SetLocalUWP(sourceWorld.UWP())
	// merch = merch.SetLocalTC(sourceWorld.TradeCodes()).SetMType(constant.MerchantTypeTrade).DetermineGoodsAvailable()
	// fmt.Println(merch)

	//merch.ListPrices()
	for i := 0; i < 50; i++ {
		fmt.Println(i)
		ss := essc.Create()
		ss.Test()
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
