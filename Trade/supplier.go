package Trade

import (
	"fmt"
	"sort"

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
	localUWP         string
	tradeDice        int
	prices           map[string]int
	volume           map[string]int
}

func NewMerchant() Merchant {
	m := Merchant{}
	m.tradeDice = 0
	m.prices = make(map[string]int)
	m.volume = make(map[string]int)
	return m
}

func (m Merchant) SetLocalUWP(uwp string) Merchant {
	m.localUWP = uwp
	return m
}

func (m Merchant) SetLocalTC(tc []string) Merchant {
	m.localTC = tc
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

func (m Merchant) DetermineGoodsAvailable() Merchant {
	var avGoodsCodes []string
	availableCategories := m.availableTGcodes
	switch m.mType {
	default:
		for i := 0; i < dice.Roll1D(); i++ {
			roll := dice.RollD66()
			availableCategories = append(availableCategories, roll)
		}
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
	sort.Strings(availableCategories)
	fmt.Println(availableCategories)
	for c := range availableCategories {

		key := availableCategories[c]
		m.prices[key] = dice.Roll3D()
		key = key + dice.Roll("2d6").SumStr()
		m.volume[categoryOf(key)] = m.volume[categoryOf(key)] + RollMaximumForCategory(key)
		avGoodsCodes = append(avGoodsCodes, key)

	}
	sort.Strings(avGoodsCodes)

	m.availableTGcodes = avGoodsCodes
	fmt.Println(m.availableTGcodes)
	return m
}

func categoryOf(code string) string {
	return string([]byte(code)[0]) + string([]byte(code)[1])
}

func (m Merchant) AvailableTradeGoods() []string {
	return m.availableTGcodes
}

type Seller interface {
	ProposeSell(string) Contract
}

//ProposeSell - Информация о передаче товара от игрока к купцу
func (m Merchant) ProposeSell(code string) Contract {
	//dealDice := 10 //dice.Roll3D()
	dealDice := m.tradeDice
	saleDMmap := getSaleDMmap(code)
	for i := range m.localTC {
		if val, ok := saleDMmap[m.localTC[i]]; ok {
			dealDice = dealDice + val
		}
	}
	return NewContract(1, code, dealDice+m.prices[categoryOf(code)])
}

type Buyer interface {
	ProposeBuy(string) Contract
}

//ProposeBuy - Информация о передаче товара от купца к игроку
func (m Merchant) ProposeBuy(code string) Contract {
	dealDice := m.tradeDice
	purchaseDMmap := getPurchaseDMmap(code) // + dice.Roll("2d6").SumStr())
	for i := range m.localTC {
		if val, ok := purchaseDMmap[m.localTC[i]]; ok {
			dealDice = dealDice + val
		}
	}
	return NewContract(2, code, dealDice+m.prices[categoryOf(code)])
}

func (m Merchant) EncodeContract(code string, cType int) string {
	//#[tgCode]-[die][cType][ta][te][TC]#
	cCode := "#"
	ta := string([]byte(m.localUWP)[5])
	te := string([]byte(m.localUWP)[6])
	tl := string([]byte(m.localUWP)[8])
	cCode += code + "-" + TrvCore.DigitToEhex(m.tradeDice) + TrvCore.DigitToEhex(cType) + ta + te + tl
	for i := range m.localTC {
		cCode += m.localTC[i]
	}
	return cCode
}

func (m Merchant) ListAvailable() {
	fmt.Println(m.availableTGcodes, "----")
	for i := range m.availableTGcodes {
		c := m.ProposeBuy(m.availableTGcodes[i])
		fmt.Println(c.ShowShort())
	}
}

func (m Merchant) ListPrices() {
	allCodes := allTradeGoodsRCodes()
	for i := range allCodes {
		c := m.ProposeSell(allCodes[i])
		fmt.Println(c.ShowShort())
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
