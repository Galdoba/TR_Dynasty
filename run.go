package main

import (
	"fmt"

	starport "github.com/Galdoba/TR_Dynasty/Starport"
	"github.com/Galdoba/TR_Dynasty/world"
)

func main() {

	w := world.NewWorld("Destiny").SetUWP("A540A98-E")
	sp := starport.From(w)
	fmt.Println(w)
	fmt.Println(sp)

	//Trade.Init()
	//	Trade.Run()
	// for i := 0; i < 300000; i++ {

	// 	fmt.Print(strconv.Itoa(i) + "/3000\r")
	// }
	// uwp := "A540A98-E"
	// tc := profile.TradeCodes(uwp)
	// fmt.Println(tc)
	// merch := Trade.NewMerchant().SetLocalUWP(uwp).SetLocalTC(tc).SetMType(constant.MerchantTypeTrade).SetTraderDice(3).DetermineGoodsAvailable()
	// merch.PurchaseList()
	// var tgCode string
	// for i := 0; i < 20; i++ {
	// 	tgCode = dice.RollD66() + dice.Roll("2d6").SumStr()
	// 	amount := dice.Roll("5d15").Sum()
	// 	merch.SaleProposalLegal(tgCode, amount)
	// }

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
