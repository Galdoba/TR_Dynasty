package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/npc"

	starport "github.com/Galdoba/TR_Dynasty/Starport"
	"github.com/Galdoba/TR_Dynasty/world"
)

func main() {
	npc.RandomNPC()

	return

	w := world.NewWorld("Destiny").SetUWP("C540A98-E")
	sp := starport.From(w)
	fmt.Println(w)
	fmt.Println(sp)
	servises := []string{"Berthing", "Refuiling", "Warehousing", "Hazmat", "Repairs"}
	fmt.Println("Services available:")
	for i, service := range servises {
		fmt.Println(service+": ", sp.ServiseTime(i))
	}
	if dice.Roll("2d6").ResultTN(8) {
		fmt.Println("General event!")
	}
	if dice.Roll("2d6").ResultTN(11) {
		fmt.Println("Significant event!")
	}

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
