package routine

import (
	"fmt"
	"strconv"
	"time"

	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

func newLocalSupplier(supType int) {
	merchType := ""
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
		localSupplier = trade.NewMerchant().SetLocalUWP(sourceWorld.UWP()).SetLocalTC(sourceWorld.TradeCodes()).SetMType(merchType).DetermineGoodsAvailable()
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
		allLots := []string{}
		allLots = append(allLots, "Return")
		for i := range localMarket {
			if freeVolume >= localMarket[i].GetVolume() {
				allLots = append(allLots, marketlotInfo(localMarket[i]))
			}
		}
		if len(allLots) == 1 {
			menu("Nothing to load...", "Return")
			break
		}
		selected, lot := menu("Load Cargo:", allLots...)
		if selected == 0 {
			done = true
			continue
		}
		purchased := localMarket[selected-1]
		purchased.SetID(int(time.Now().UnixNano()))
		canBuy := utils.Min(freeVolume, localMarket[selected-1].GetVolume())
		fmt.Println("canBuy", canBuy)
		//fmt.Print("Set tons to buy [0-" + strconv.Itoa(canBuy) + "]: ")
		qty := userInputIntMinMax("Set tons to buy", 0, canBuy)

		purchased.SetVolume(qty)
		localMarket[selected-1].SetVolume(localMarket[selected-1].GetVolume() - qty)
		cm.entry = append(cm.entry, purchased)
		message = lot + " was loaded to ship"
		saveCargoManifest(cm)

	}
	clrScrn()
}

func marketlotInfo(lot cargoLot) string {
	lotInfo := lot.GetTGCode() + "	" + strconv.Itoa(lot.GetVolume()) + " tons  	" + strconv.Itoa(lot.GetCost()) + " Cr  	" + lot.GetDescr()
	if !lot.GetLegality() {
		lotInfo += "(!!!)"
	}
	return lotInfo
}
