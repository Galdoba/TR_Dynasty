package Trade

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

var tgDB map[string][]string

func init() {
	tgDB = TradeGoodRData()
}

type supplier struct {
	sType string
	//cargo    []*tradeLot
	//cargo     map[string]*tradeLot
	cargoNew  *Cargo
	planet    *world.World
	tradeDice int
}

type Merchant struct {
	mType            string
	availableTGcodes []string
	localTC          []string
	tradeDice        int
}

func NewMerchantOn(localTradeCodes []string) Merchant {
	m := Merchant{}
	m.mType = supplierTypeTrade
	m.tradeDice = 0
	m.localTC = localTradeCodes
	m.availableTGcodes = availableFromTCodes(localTradeCodes)
	return m
}

func (m Merchant) SetMType(mType string) Merchant {
	mTypeErr := true
	for _, val := range []string{supplierTypeCommon, supplierTypeTrade, supplierTypeNeutral, supplierTypeCommon} {
		if mType == val {
			mTypeErr = false
			break
		}
	}
	if mTypeErr {
		return m
	}
	m.mType = mType
	return m
}

func (m Merchant) SetTraderDice(tDice int) Merchant {
	m.tradeDice = tDice
	return m
}

func (m Merchant) DetermineGoodsAvailable() []string {
	var avGoodsCodes []string
	availableCategories := m.availableTGcodes
	switch m.mType {
	default:
		return avGoodsCodes
	case supplierTypeCommon:
		availableCategories = []string{"11", "12", "13", "14", "15", "16"}
		add := utils.RollDiceRandom("d6")
		for i := 1; i < add; i++ {
			roll := TrvCore.RollD66()
			if !utils.ListContains([]string{"11", "12", "13", "14", "15", "16"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeTrade:
		availableCategories = append([]string{"11", "12", "13", "14", "15", "16"}, availableCategories...)
		add := utils.RollDiceRandom("d6")
		for i := 0; i < add; i++ {
			roll := dice.RollD66()
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
		//ExcludeFromSliceStr(availableCategories, "66")
	case supplierTypeNeutral:
		add := dice.Roll1D()
		for i := 0; i < add; i++ {
			roll := dice.RollD66()
			if utils.ListContains([]string{"66"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeIlligal:
		availableCategories = []string{"61", "62", "63", "64", "65"}
		add := dice.Roll1D()
		for i := 0; i < add; i++ {
			roll := dice.RollD66()
			if !utils.ListContains([]string{"61", "62", "63", "64", "65"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	}

	for c := range availableCategories {
		definition := dice.Roll("2d6").SumStr()
		key := availableCategories[c] + definition
		avGoodsCodes = append(avGoodsCodes, key)
		// if _, ok := sup.cargo[key]; ok {
		// 	//do something here
		// 	sup.cargo[key].cargoVolume = sup.cargo[key].cargoVolume + sup.cargo[key].lotTradeGoodR.IncreaseRandom()
		// } else {
		// 	newLot := NewTradeLot(availableCategories[c]+definition, sup.planet)
		// 	sup.cargo[availableCategories[c]+definition] = newLot
		// }
		//sup.cargoNew.Add(tgr, tgr.IncreaseRandom())
		//fmt.Println(tgDB[key])
	}
	return avGoodsCodes
}

type Seller interface {
	ProposeSell(string) Contract
}

//ProposeSell - Информация о передаче товара от игрока к купцу
func (m Merchant) ProposeSell(code string) Contract {
	dealDice := dice.Roll3D()
	dealDice += m.tradeDice
	saleDMmap := getSaleDMmap(code)
	for i := range m.localTC {
		if val, ok := saleDMmap[m.localTC[i]]; ok {
			dealDice = dealDice + val
		}
	}
	fmt.Println(encodeContractCode(code, dealDice, m.localTC), "--------------------------------")
	return NewContract(1, code, dealDice)
}

type Buyer interface {
	ProposeBuy(string) Contract
}

//ProposeBuy - Информация о передаче товара от купца к игроку
func (m Merchant) ProposeBuy(code string) Contract {
	dealDice := dice.Roll3D()
	dealDice += m.tradeDice
	purchaseDMmap := getPurchaseDMmap(code)
	for i := range m.localTC {
		if val, ok := purchaseDMmap[m.localTC[i]]; ok {
			dealDice = dealDice + val
		}
	}
	return NewContract(2, code, dealDice)
}

func (m Merchant) ListForSale() {
	//TODO: нужно переписать метод так чтобы по запросу [Merchant.ProposeSale(code)] он выдавал строку с информацией о типе и стоймости товара
	//
	var rows [][]string
	rows = append(rows, []string{"CODE", "description", "price per ton"})
	for _, category := range m.availableTGcodes {
		code := category + dice.Roll("2d6").SumStr()
		basePrice := getBasePrice(code)
		description := getDescription(code)
		priceDie := dice.Roll("3d6").Sum()
		saleDMmap := getSaleDMmap(code)
		for i := range m.localTC {
			if val, ok := saleDMmap[m.localTC[i]]; ok {
				priceDie = priceDie + val
			}
		}
		actualPrice := modifyPriceSale(basePrice, priceDie)
		//fmt.Println(description, basePrice, priceDie, actualPrice, stock)
		//fmt.Println(code, "	", description, "for", actualPrice, "per ton", basePrice)
		rows = append(rows, []string{code, description, strconv.Itoa(actualPrice)})
	}
}

func Test() {

}

func NewSupplier(stype string, planet *world.World) *supplier {
	sup := &supplier{}
	// seed := utils.CurrentSeed()
	sup.sType = stype
	sup.planet = planet
	//sup.cargo = make(map[string]*tradeLot)
	sup.cargoNew = NewCargo()
	//sup.determineGoodsAvailable()
	return sup
}

func (sup *supplier) CargoNewShow() *Cargo {
	return sup.cargoNew
}

// func (sup *supplier) determineGoodsAvailable() []string {
// 	var avGoodsCodes []string
// 	availableCategories := availableCategories(sup.planet)
// 	switch sup.sType {
// 	default:
// 		return avGoodsCodes
// 	case supplierTypeCommon:
// 		availableCategories = []string{"11", "12", "13", "14", "15", "16"}
// 		add := utils.RollDiceRandom("d6")
// 		for i := 0; i < add; i++ {
// 			roll := TrvCore.RollD66()
// 			if !utils.ListContains([]string{"11", "12", "13", "14", "15", "16"}, roll) {
// 				i--
// 				continue
// 			}
// 			availableCategories = append(availableCategories, roll)
// 		}
// 	case supplierTypeTrade:
// 		add := utils.RollDiceRandom("d6")
// 		for i := 0; i < add; i++ {
// 			roll := TrvCore.RollD66()
// 			if utils.ListContains([]string{"61", "62", "63", "64", "65", "66"}, roll) {
// 				i--
// 				continue
// 			}
// 			availableCategories = append(availableCategories, roll)
// 		}
// 		ExcludeFromSliceStr(availableCategories, "61")
// 		ExcludeFromSliceStr(availableCategories, "62")
// 		ExcludeFromSliceStr(availableCategories, "63")
// 		ExcludeFromSliceStr(availableCategories, "64")
// 		ExcludeFromSliceStr(availableCategories, "65")
// 	case supplierTypeNeutral:
// 		add := utils.RollDiceRandom("d6")
// 		for i := 0; i < add; i++ {
// 			roll := TrvCore.RollD66()
// 			if utils.ListContains([]string{"66"}, roll) {
// 				i--
// 				continue
// 			}
// 			availableCategories = append(availableCategories, roll)
// 		}
// 	case supplierTypeIlligal:
// 		availableCategories = []string{"61", "62", "63", "64", "65"}
// 		add := utils.RollDiceRandom("d6")
// 		for i := 0; i < add; i++ {
// 			roll := TrvCore.RollD66()
// 			if !utils.ListContains([]string{"61", "62", "63", "64", "65"}, roll) {
// 				i--
// 				continue
// 			}
// 			availableCategories = append(availableCategories, roll)
// 		}
// 	}

// 	for c := range availableCategories {
// 		definition := convert.ItoS(trimDefinition(utils.RollDiceRandom("2d6")))
// 		key := availableCategories[c] + definition
// 		// if _, ok := sup.cargo[key]; ok {
// 		// 	//do something here
// 		// 	sup.cargo[key].cargoVolume = sup.cargo[key].cargoVolume + sup.cargo[key].lotTradeGoodR.IncreaseRandom()
// 		// } else {
// 		// 	newLot := NewTradeLot(availableCategories[c]+definition, sup.planet)
// 		// 	sup.cargo[availableCategories[c]+definition] = newLot
// 		// }
// 		tgr := NewTradeGoodR(key)
// 		sup.cargoNew.Add(tgr, tgr.IncreaseRandom())
// 	}

// 	return avGoodsCodes
// }

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
	//sup.determineGoodsAvailable()
}
