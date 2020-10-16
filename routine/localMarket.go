package routine

import (
	"fmt"
	"strconv"
	"time"

	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/utils"
)

var sTryes int

func newLocalSupplier(supType int) {
	merchType := ""
	localMarket = nil
	switch supType {
	default:
		merchType = constant.MerchantTypeTrade
	case 0:
		merchType = constant.MerchantTypeCommon
	case 1:
		merchType = constant.MerchantTypeTrade
	case 2:
		merchType = constant.MerchantTypeNeutral
	case 3:
		merchType = constant.MerchantTypeIlligal
	}
	d := dice.New(utils.SeedFromString(sourceWorld.Name() + formatDate(day, year)))
	if localMarket == nil {
		tdie := dice.Roll("3d6").DM(-1 * sTryes).Sum()
		sTryes++
		localSupplier = trade.NewMerchant().SetLocalUWP(sourceWorld.UWP()).SetTraderDice(tdie).SetLocalTC(sourceWorld.TradeCodes()).SetMType(merchType).DetermineGoodsAvailable()
		for _, val0 := range localSupplier.AvailableTradeGoods() {
			val := val0 + d.RollNext("2d6").SumStr()
			marketLot := newCargoLot()
			marketLot.SetTGCode(val)
			marketLot.SetOrigin(sourceWorld.Hex())
			marketLot.SetCost(localSupplier.CostPurchase(val))
			marketLot.SetVolume(d.RollNext("1d" + strconv.Itoa(localSupplier.Volume(val))).Sum())
			marketLot.SetDangerDM(trade.GetDangerousGoodsDM(val))
			marketLot.SetRiskDM(trade.GetMaximumRiskAssessment(val))
			marketLot.SetDescr(trade.GetDescription(val))
			marketLot.SetLegality(true)
			marketLot.SetComment("Bought at " + sourceWorld.Name() + " on day " + formatDate(day, year))
			switch val0 {
			case "61", "62", "63", "64", "65":
				marketLot.SetLegality(false)
			}
			localMarket = append(localMarket, marketLot)
			//		fmt.Println(i, marketLot)
		}
	}
	//fmt.Println(localMarket)
	//	testTrader()
}

func testTrader() {
	initialTD := localSupplier.TraderDice()
	for i, val := range localMarket {

		localSupplier = localSupplier.SetTraderDice(initialTD)
		fmt.Println(localSupplier.TraderDice())
		fmt.Println(i, "val.GetCost()", val.GetCost(), val.GetDescr())
		fmt.Println(i, localSupplier.CostSale(val.GetTGCode()))
		for d := 0; d < 20; d++ {
			localSupplier = localSupplier.SetTraderDice(d)
			fmt.Println(d, localSupplier.CostSale(val.GetTGCode()), localSupplier.CostPurchase(val.GetTGCode()))
		}

	}

}

func purchase() {
	done := false
	message := ""
	cm := loadCargoManifest()
	for !done {
		freeVolume := freeCargoVolume()
		clrScrn()
		if message != "" {
			fmt.Println(message)
		}
		fmt.Println("Free Volume: ", freeVolume)
		fmt.Println("Trader Dice: ", localSupplier.TraderDice())
		allLots := []string{}
		allLots = append(allLots, "Return")
		for i := range localMarket {
			//if freeVolume >= localMarket[i].GetVolume() {
			allLots = append(allLots, marketlotInfo(localMarket[i]))
			//}
		}
		if len(allLots) == 1 {
			menu("Nothing to load...", "Return")
			break
		}
		selected, _ := menu("Load Cargo:", allLots...)
		if selected == 0 {
			done = true
			continue
		}
		purchased := localMarket[selected-1]
		purchased.SetID(int(time.Now().UnixNano()))
		canBuy := utils.Min(freeVolume, localMarket[selected-1].GetVolume())
		//fmt.Println("canBuy", canBuy, localMarket[selected-1].GetVolume(), localMarket[selected-1].GetTGCode())
		fmt.Println("Selected:", purchased.GetDescr(), "with price of", purchased.GetCost(), "Cr per ton")
		//fmt.Print("Set tons to buy [0-" + strconv.Itoa(canBuy) + "]: ")
		qty := userInputIntMinMax("Set tons to buy", 0, canBuy)
		if qty == 0 {
			continue
		}
		purchased.SetVolume(qty)

		fmt.Println("Total Cost:", strconv.Itoa(purchased.GetCost()*purchased.GetVolume())+" Cr")
		if userConfirm("Confirm Transaction?") {
			cm.entry = append(cm.entry, purchased)
			message = strconv.Itoa(purchased.GetVolume()) + " tons of " + purchased.GetDescr() + " was loaded to ship"
			saveCargoManifest(cm)
			localMarket[selected-1].SetVolume(localMarket[selected-1].GetVolume() - qty)
		}

	}
	menuPosition = "MARKET"
	clrScrn()
}

func marketlotInfo(lot cargoLot) string {
	lotInfo := lot.GetTGCode() + "	" + strconv.Itoa(lot.GetVolume()) + " tons  	" + strconv.Itoa(lot.GetCost()) + " Cr  	" + lot.GetDescr()
	if !lot.GetLegality() {
		lotInfo += "(!!!)"
	}
	return lotInfo
}

func sellTradeGoods() {
	done := false

	for !done {
		cm := loadCargoManifest()
		//totalProfit := 0
		allLots := []string{}
		ids := []int{}
		allLots = append(allLots, "Return")
		idMap := make(map[int]int)
		for i, val := range cm.entry {
			if val.GetOrigin() == sourceWorld.Hex() {
				continue
			}
			if val.GetTGCode() == "XXX" {
				continue
			}
			price := localSupplier.CostSale(val.GetTGCode())
			allLots = append(allLots, strconv.Itoa(price)+" Cr		"+val.GetDescr()+" ("+strconv.Itoa(val.GetCost())+" Cr) "+strconv.Itoa(val.GetVolume())+" tons "+val.GetComment())
			ids = append(ids, val.GetID())
			idMap[val.GetID()] = i

		}

		i, _ := menu("Select Cargo to Sell:", allLots...)
		if i == 0 {
			menuPosition = "MARKET"
			return
		}
		selectedID := ids[i-1]
		salePosition := byID(selectedID)
		qty := userInputIntMinMax("Set tons to sell", 0, salePosition.GetVolume())
		if qty == 0 {
			menuPosition = "MARKET"
			return
		}
		salePosition.SetVolume(qty)

		if userConfirm("Sell " + strconv.Itoa(salePosition.GetVolume()) + " tons for " + strconv.Itoa(salePosition.GetVolume()*localSupplier.CostSale(salePosition.GetTGCode())) + " Cr") {
			for i, val := range cm.entry {
				if val.GetID() != salePosition.GetID() {
					continue
				}
				cm.entry[i].SetVolume(cm.entry[i].GetVolume() - salePosition.GetVolume())
				profit := salePosition.GetVolume() * (localSupplier.CostSale(salePosition.GetTGCode()) - cm.entry[i].GetCost())
				if profit < 0 {
					profit = 0
				}
				taxes := trade.TaxingAmount(profit, taxingAgent())
				fmt.Println(taxes, "Cr was charged as Taxes")
				fmt.Println("------------")
				fmt.Println("Pure profit:", salePosition.GetVolume()*localSupplier.CostSale(salePosition.GetTGCode())-taxes, "Cr")
				if cm.entry[i].GetVolume() <= 0 {
					cm = deleteFromCargoManifest(salePosition.GetID())
				}
				if userConfirm("Continue") {
				}
			}
			saveCargoManifest(cm)
		}

	}

}

func taxingAgent() string {
	if localSupplier.MerchantType() == constant.MerchantTypeIlligal {
		return "CM"
	}
	prf, err := profile.NewUWP(sourceWorld.UWP())
	reportErr(err)
	return prf.Govr()
}
