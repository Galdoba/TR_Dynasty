package Trade

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Galdoba/devtools/cli/prettytable"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

var tgDB map[string][]string

func init() {
	tgDB = TradeGoodRData()
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

func (m Merchant) CostPurchase(code string) int {
	pDM, sDM := TradeDMs(categoryOf(code), m.localTC)
	base := getBasePrice(code)
	pPrice := modifyPricePurchase(base, pDM-sDM+m.tradeDice+m.prices[code])
	return pPrice
}

func (m Merchant) CostSale(code string) int {
	pDM, sDM := TradeDMs(categoryOf(code), m.localTC)
	base := getBasePrice(code)
	sPrice := modifyPriceSale(base, sDM-pDM+m.tradeDice+m.prices[code])
	return sPrice
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

func (m Merchant) AvailableCategories() []string {
	return m.availableTGcodes
}

func matchWorldsTC(code string, tc []string) bool {
	pMap := getPurchaseDMmap(code)
	var keys []string
	for k, _ := range pMap {
		keys = append(keys, k)
	}
	m, _ := matchTradeCodes(keys, tc)
	if m {
		return true
	}
	// for k, _ := range pMap {
	// 	for i := range tc {
	// 		if tc[i] == k {
	// 			return true
	// 		}
	// 	}
	// }
	return false
}

func (m Merchant) DetermineGoodsAvailable() Merchant {
	var avGoodsCodes []string
	//availableCategories := m.availableTGcodes
	availableCategories := []string{"11", "12", "13", "14", "15", "16"}
	allCodes := allCategories()
	for i := range allCodes {
		if matchWorldsTC(allCodes[i]+"7", m.localTC) {
			availableCategories = append(availableCategories, allCodes[i])
		}
	}
	for i := 0; i < dice.Roll1D(); i++ {
		roll := dice.RollD66()
		fmt.Println(roll)
		availableCategories = append(availableCategories, roll)
	}
	//TODO: дальше в зависимости от типа торговцев исключать не подходящие товары и считать их тоннаж

	return m
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
		//availableCategories = append([]string{"11", "12", "13", "14", "15", "16"}, availableCategories...)

		for i := range allCodes {
			fmt.Println(allCodes[i] + "7")
			tgTCList := getAvailabilityTags(allCodes[i] + "7")
			if len(commonElements(tgTCList, m.localTC)) > 0 {
				availableCategories = append(availableCategories, allCodes[i])
			}
		}

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
	//fmt.Println(availableCategories)
	for c := range availableCategories {

		key := availableCategories[c]
		m.prices[key] = dice.Roll3D()
		//key = key + dice.Roll("2d6").SumStr()
		m.volume[categoryOf(key)] = m.volume[categoryOf(key)] + RollMaximumForCategory(key)
		avGoodsCodes = append(avGoodsCodes, key)

	}
	sort.Strings(avGoodsCodes)

	m.availableTGcodes = avGoodsCodes
	//fmt.Println(m.availableTGcodes)

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

func (m Merchant) MakeOffer(code string, operation int) []Contract {

	var allCont []Contract
	maxTons := RollMaximumForCategory(code)
	fmt.Println("maxTons", maxTons)
	exactVolume := make(map[string]int)
	for i := 0; i < maxTons; i++ {
		exactVolume[code+dice.Roll("2d6").SumStr()]++
	}
	//categoryDice :=
	table := prettytable.New()
	table.AddRow([]string{"Category", "Maximum Tons", "Base Price", "Lot", "Price", "Trade Dice"})
	for k, v := range exactVolume {
		c := Contract{}
		c.lotCode = k
		c.volume = v
		c.cType = operation
		c.contractDice = m.tradeDice + dice.Roll3D()
		c.taxingAgent = string([]byte(m.localUWP)[5])
		c.taxingEnviroment = string([]byte(m.localUWP)[6])
		c.lotDescription = getDescription(k)
		c.category = getCategory(k)
		table.AddRow(c.SellShort())
		allCont = append(allCont, c)
	}
	table = prettytable.InsertSeparatorRow(table, 1)
	table.PTPrint()
	return allCont
}

func cleanSlice(sl []string) []string {
	var newSl []string
	for i := range sl {
		newSl = utils.AppendUniqueStr(newSl, sl[i])
	}
	return newSl
}

func (m Merchant) PurchaseList() {
	tb := prettytable.New()
	tb.AddRow([]string{"Item", "Maximum Tons", "Tons per Defined Trade Good (Base Price)", "Cost", "Purchase DM"})
	//cleanaTGcodes := cleanSlice(m.availableTGcodes)
	aC := allCategories()
	for i := range aC {

		catList := listCategory(m, aC[i])
		for l := range catList {
			tb.AddRow(catList[l])

		}
		fmt.Print(".")
	}
	tb = prettytable.InsertSeparatorRow(tb, 1)
	fmt.Print("OK!\n")
	tb.PTPrintSlow(0)
}

func listCategory(m Merchant, code string) [][]string {
	if code == "66" {
		return [][]string{}
	}
	maxTons := RollMaximumForCategory(code) * countElement(code, m.availableTGcodes)
	purchaseDM := -999
	//var price int
	var dataSheet [][]string
	exactVolume := make(map[string]int)
	purchDMmap := getPurchaseDMmap(code + "7")
	for k, val := range purchDMmap {
		for i := range m.localTC {
			if m.localTC[i] == k {
				if purchaseDM < val {
					purchaseDM = val
				}
			}
		}
	}
	if purchaseDM == -999 {
		purchaseDM = 0
	}

	tradePurch := m.tradeDice + dice.Roll3D() + purchaseDM
	for i := 0; i < maxTons; i++ {
		descr := dice.Roll("2d6").SumStr()
		switch descr {
		case "3", "4", "5":
			descr = "4"
		case "6", "7", "8":
			descr = "7"
		case "9", "10", "11":
			descr = "10"
		}
		exactVolume[code+descr]++
	}
	madeCat := false
	for _, descr := range []string{"2", "4", "7", "10", "12"} {
		if exactVolume[code+descr] == 0 {
			continue
		}
		dataline := make([]string, 5)
		if !madeCat {
			dataline[0] = getCategory(code + descr)
			dataline[1] = strconv.Itoa(maxTons)
			dataline[4] = strconv.Itoa(purchaseDM)
			if purchaseDM >= 0 {
				dataline[4] = "+" + dataline[4]
			}

			madeCat = true
		}
		basePrice := getBasePrice(code + descr)
		dataline[2] = strconv.Itoa(exactVolume[code+descr]) + " x " + getDescription(code+descr) + " (" + strconv.Itoa(basePrice) + ")"

		costP := modifyPricePurchase(basePrice, tradePurch)

		if exactVolume[code+descr] > 0 {
			dataline[3] = strconv.Itoa(costP) // + " (" + strconv.Itoa(basePrice) + ")"
		}

		//dataline[5] = strconv.Itoa(basePrice)
		dataSheet = append(dataSheet, dataline)
	}
	return dataSheet
}

func CategoryList() []string {
	return allCategories()
}

func allCategories() []string {
	return []string{
		"11",
		"12",
		"13",
		"14",
		"15",
		"16",
		"21",
		"22",
		"23",
		"24",
		"25",
		"26",
		"31",
		"32",
		"33",
		"34",
		"35",
		"36",
		"41",
		"42",
		"43",
		"44",
		"45",
		"46",
		"51",
		"52",
		"53",
		"54",
		"55",
		"56",
		"61",
		"62",
		"63",
		"64",
		"65",
		//"66",
	}

}

func countElement(elem string, sl []string) int {
	var n int
	for i := range sl {
		if sl[i] == elem {
			n++
		}
	}
	return n
}

func TradeDMs(code string, tc []string) (pDM int, sDM int) {
	pDM = -999
	purchDMmap := getPurchaseDMmap(code + "7")
	for k, val := range purchDMmap {
		for i := range tc {
			if tc[i] == k {
				if pDM < val {
					pDM = val
				}
			}
		}
	}
	if pDM == -999 {
		pDM = 0
	}
	sDM = -999
	saleDMmap := getSaleDMmap(code + "7")
	for k, val := range saleDMmap {
		for i := range tc {
			if tc[i] == k {
				if sDM < val {
					sDM = val
				}
			}
		}
	}
	if sDM == -999 {
		sDM = 0
	}
	return pDM, sDM
}
