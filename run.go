package main

import (
	"fmt"

	trademgt2 "github.com/Galdoba/TR_Dynasty/TradeMgT2"
	"github.com/Galdoba/TR_Dynasty/otu"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
)

func main() {
	//trade.RunTraffic()
	neib := Astrogation.JumpCoordinatesFrom("2923", 2)
	hexMap := otu.MapDataByHex(otu.TrojanReachData())
	for _, hex := range neib {
		if val, ok := hexMap[hex]; ok {
			fmt.Println(val)
		}
	}

	//sectorData := otu.TrojanReachData()
	// h1 := Astrogation.Hex("2624")
	// h2 := Astrogation.Hex("2520")
	// fmt.Println(h1)
	// fmt.Println(h2)
	// dist := Astrogation.JumpDistance(h1, h2)
	// fmt.Println(dist)

	trademgt2.RunMerchantPrince()

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

func checkCoords(coords string) bool {
	if len(coords) != 2 {
		return false
	}
	for _, v := range coords {
		switch string(v) {
		default:
			return false
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			continue
		}
		//fmt.Println(i, string(v))
	}
	return true
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
