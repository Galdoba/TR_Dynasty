package Trade

import (
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

type supplier struct {
	sType string
	//cargo    []*tradeLot
	//cargo     map[string]*tradeLot
	cargoNew  *Cargo
	planet    *world.World
	tradeDice int
}

func NewSupplier(stype string, planet *world.World) *supplier {
	sup := &supplier{}
	// seed := utils.CurrentSeed()
	sup.sType = stype
	sup.planet = planet
	//sup.cargo = make(map[string]*tradeLot)
	sup.cargoNew = NewCargo()
	sup.determineGoodsAvailable()
	return sup
}

func (sup *supplier) CargoNewShow() *Cargo {
	return sup.cargoNew
}

func (sup *supplier) determineGoodsAvailable() []string {
	var avGoodsCodes []string
	availableCategories := availableCategories(sup.planet)
	switch sup.sType {
	default:
		return avGoodsCodes
	case supplierTypeCommon:
		availableCategories = []string{"11", "12", "13", "14", "15", "16"}
		add := utils.RollDiceRandom("d6")
		for i := 0; i < add; i++ {
			roll := TrvCore.RollD66()
			if !utils.ListContains([]string{"11", "12", "13", "14", "15", "16"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeTrade:
		add := utils.RollDiceRandom("d6")
		for i := 0; i < add; i++ {
			roll := TrvCore.RollD66()
			if utils.ListContains([]string{"61", "62", "63", "64", "65", "66"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
		ExcludeFromSliceStr(availableCategories, "61")
		ExcludeFromSliceStr(availableCategories, "62")
		ExcludeFromSliceStr(availableCategories, "63")
		ExcludeFromSliceStr(availableCategories, "64")
		ExcludeFromSliceStr(availableCategories, "65")
	case supplierTypeNeutral:
		add := utils.RollDiceRandom("d6")
		for i := 0; i < add; i++ {
			roll := TrvCore.RollD66()
			if utils.ListContains([]string{"66"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeIlligal:
		availableCategories = []string{"61", "62", "63", "64", "65"}
		add := utils.RollDiceRandom("d6")
		for i := 0; i < add; i++ {
			roll := TrvCore.RollD66()
			if !utils.ListContains([]string{"61", "62", "63", "64", "65"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	}

	for c := range availableCategories {
		definition := convert.ItoS(trimDefinition(utils.RollDiceRandom("2d6")))
		key := availableCategories[c] + definition
		// if _, ok := sup.cargo[key]; ok {
		// 	//do something here
		// 	sup.cargo[key].cargoVolume = sup.cargo[key].cargoVolume + sup.cargo[key].lotTradeGoodR.IncreaseRandom()
		// } else {
		// 	newLot := NewTradeLot(availableCategories[c]+definition, sup.planet)
		// 	sup.cargo[availableCategories[c]+definition] = newLot
		// }
		tgr := NewTradeGoodR(key)
		sup.cargoNew.Add(tgr, tgr.IncreaseRandom())
	}

	return avGoodsCodes
}

// func (sup *supplier) CargoInfo() (output []string) {
// 	tradeGoodsCodesLIST := tradeGoodsCodesLIST()
// 	var validKeys []string
// 	for keyNum := range tradeGoodsCodesLIST {
// 		key := tradeGoodsCodesLIST[keyNum]
// 		if _, ok := sup.cargo[key]; ok {
// 			validKeys = append(validKeys, key)
// 		}
// 	}
// 	for i := range validKeys {
// 		lot := sup.cargo[validKeys[i]]
// 		outPutDesc := lot.lotTradeGoodR.pickRandomDescription()
// 		outputTons := convert.ItoS(lot.cargoVolume)
// 		basePrice := convert.ItoS(lot.lotTradeGoodR.basePrice) + " Cr"
// 		output = append(output, outPutDesc+" -- "+outputTons+" -- "+basePrice)
// 	}
// 	return output
// }

func (sup *supplier) RerollTradeGoods() {
	//sup.cargo = make(map[string]*tradeLot)
	sup.cargoNew = NewCargo()
	sup.determineGoodsAvailable()
}
