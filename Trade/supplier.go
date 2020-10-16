package trade

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/utils"
)

var tgDB map[string][]string

func init() {
	tgDB = TGoodsData()
}

//Merchant -
type Merchant struct {
	mType            string
	availableTGcodes []string
	localTC          []string
	localUWP         string
	tradeDice        int
	prices           map[string]int
	volume           map[string]int
}

//NewMerchant -
func NewMerchant() Merchant {
	m := Merchant{}
	m.tradeDice = 0
	m.prices = make(map[string]int)
	m.volume = make(map[string]int)
	return m
}

//SetLocalUWP -
func (m Merchant) SetLocalUWP(uwp string) Merchant {
	m.localUWP = uwp
	return m
}

//SetLocalTC -
func (m Merchant) SetLocalTC(tc []string) Merchant {
	m.localTC = tc
	return m
}

//CostPurchase -
func (m Merchant) CostPurchase(code string) int {
	pDM, sDM := PurchSaleDMs(categoryOf(code), m.localTC)
	base := GetBasePrice(code)
	pPrice := modifyPricePurchase(base, pDM-sDM+m.tradeDice+m.prices[categoryOf(code)])
	return pPrice
}

//CostSale -
func (m Merchant) CostSale(code string) int {
	pDM, sDM := PurchSaleDMs(categoryOf(code), m.localTC)
	base := GetBasePrice(code)
	sPrice := modifyPriceSale(base, sDM-pDM+m.tradeDice+m.prices[categoryOf(code)])
	return sPrice
}

//SetMType -
func (m Merchant) SetMType(mType string) Merchant {
	m.mType = mType
	return m
}

//SetTraderDice -
func (m Merchant) SetTraderDice(tDice int) Merchant {
	m.tradeDice = tDice
	return m
}

func (m Merchant) TraderDice() int {
	return m.tradeDice
}

//AvailableCategories -
func (m Merchant) AvailableCategories() []string {
	return m.availableTGcodes
}

func matchWorldsTC(code string, tc []string) bool {
	pMap := GetPurchaseDMmap(code)
	var keys []string
	for k := range pMap {
		keys = append(keys, k)
	}
	m, _ := matchTradeCodes(keys, tc)
	if m {
		return true
	}
	return false
}

//DetermineGoodsAvailable -
func (m Merchant) DetermineGoodsAvailable() Merchant {
	availableCategories := []string{"11", "12", "13", "14", "15", "16"}
	allCodes := allCategories()
	for i := range allCodes {
		if matchWorldsTC(allCodes[i]+"7", m.localTC) {
			availableCategories = append(availableCategories, allCodes[i])
		}
	}
	for i := 0; i < dice.Roll1D(); i++ {
		roll := dice.RollD66()
		for roll == "66" {
			roll = dice.RollD66() //исключает экзотические товары пока не имплементируется
		}
		availableCategories = append(availableCategories, roll)
	}
	switch m.mType { //case constant.MerchantTypeNeutral: имеет ВСЕ товары
	case constant.MerchantTypeCommon:
		var newAvail []string
		for _, val := range availableCategories {
			if val == "11" || val == "12" || val == "13" || val == "14" || val == "15" || val == "16" {
				newAvail = append(newAvail, val)
			}
		}
		availableCategories = newAvail
	case constant.MerchantTypeTrade:
		var newAvail []string
		for _, val := range availableCategories {
			if val != "61" && val != "62" && val != "63" && val != "64" && val != "65" {
				newAvail = append(newAvail, val)
			}
		}
		availableCategories = newAvail
	case constant.MerchantTypeIlligal:
		var newAvail []string
		for _, val := range availableCategories {
			if val == "61" || val == "62" || val == "63" || val == "64" || val == "65" || val == "66" {
				newAvail = append(newAvail, val)
			}
		}
		availableCategories = newAvail
	}
	sort.Strings(availableCategories)
	m.availableTGcodes = availableCategories
	for c := range availableCategories {
		key := availableCategories[c]
		m.prices[categoryOf(key)] = dice.Roll3D()
		m.volume[categoryOf(key)] = m.volume[categoryOf(key)] + RollMaximumForCategory(key)
	}
	return m
}

func (m *Merchant) Volume(code string) int {
	return m.volume[categoryOf(code)]
}

func (m *Merchant) MerchantType() string {
	return m.mType
}

func categoryOf(code string) string {
	return string([]byte(code)[0]) + string([]byte(code)[1])
}

func RandomTGCategory(w wrld.World) string {
	merch := NewMerchant().SetLocalUWP(w.UWP()).SetLocalTC(w.TradeCodes()).SetMType(constant.MerchantTypeTrade).DetermineGoodsAvailable()
	l := len(merch.availableTGcodes)
	//fmt.Print(".")
	return merch.availableTGcodes[dice.Roll("1d"+strconv.Itoa(l)).DM(-1).Sum()]
}

//AvailableTradeGoods -
func (m Merchant) AvailableTradeGoods() []string {
	return m.availableTGcodes
}

// type Seller interface {
// 	ProposeSell(string) Contract
// }

// //ProposeSell - Информация о передаче товара от игрока к купцу
// func (m Merchant) ProposeSell(code string) Contract {
// 	//dealDice := 10 //dice.Roll3D()
// 	dealDice := m.tradeDice
// 	saleDMmap := GetSaleDMmap(code)
// 	for i := range m.localTC {
// 		if val, ok := saleDMmap[m.localTC[i]]; ok {
// 			dealDice = dealDice + val
// 		}
// 	}
// 	return NewContract(1, code, dealDice+m.prices[categoryOf(code)])
// }

// type Buyer interface {
// 	ProposeBuy(string) Contract
// }

// //ProposeBuy - Информация о передаче товара от купца к игроку
// func (m Merchant) ProposeBuy(code string) Contract {
// 	dealDice := m.tradeDice
// 	purchaseDMmap := GetPurchaseDMmap(code) // + dice.Roll("2d6").SumStr())
// 	for i := range m.localTC {
// 		if val, ok := purchaseDMmap[m.localTC[i]]; ok {
// 			dealDice = dealDice + val
// 		}
// 	}
// 	return NewContract(2, code, dealDice+m.prices[categoryOf(code)])
// }

// func (m Merchant) EncodeContract(code string, cType int) string {
// 	//#[tgCode]-[die][cType][ta][te][TC]#
// 	cCode := "#"
// 	ta := string([]byte(m.localUWP)[5])
// 	te := string([]byte(m.localUWP)[6])
// 	tl := string([]byte(m.localUWP)[8])
// 	cCode += code + "-" + TrvCore.DigitToEhex(m.tradeDice) + TrvCore.DigitToEhex(cType) + ta + te + tl
// 	for i := range m.localTC {
// 		cCode += m.localTC[i]
// 	}
// 	return cCode
// }

// func (m Merchant) MakeOffer(code string, operation int) []Contract {

// 	var allCont []Contract
// 	maxTons := RollMaximumForCategory(code)
// 	fmt.Println("maxTons", maxTons)
// 	exactVolume := make(map[string]int)
// 	for i := 0; i < maxTons; i++ {
// 		exactVolume[code+dice.Roll("2d6").SumStr()]++
// 	}
// 	//categoryDice :=
// 	table := prettytable.New()
// 	table.AddRow([]string{"Category", "Maximum Tons", "Base Price", "Lot", "Price", "Trade Dice"})
// 	for k, v := range exactVolume {
// 		c := Contract{}
// 		c.lotCode = k
// 		c.volume = v
// 		c.cType = operation
// 		c.contractDice = m.tradeDice + dice.Roll3D()
// 		c.taxingAgent = string([]byte(m.localUWP)[5])
// 		c.taxingEnviroment = string([]byte(m.localUWP)[6])
// 		c.lotDescription = GetDescription(k)
// 		c.category = GetCategory(k)
// 		table.AddRow(c.SellShort())
// 		allCont = append(allCont, c)
// 	}
// 	table = prettytable.InsertSeparatorRow(table, 1)
// 	table.PTPrint()
// 	return allCont
// }

func cleanSlice(sl []string) []string {
	var newSl []string
	for i := range sl {
		newSl = utils.AppendUniqueStr(newSl, sl[i])
	}
	return newSl
}

func sumMap(m map[string]int) int {
	s := 0
	for _, v := range m {
		s += v
	}
	return s
}

//CategoryList -
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

//PurchSaleDMs -
func PurchSaleDMs(code string, tc []string) (pDM int, sDM int) {
	pDM = -999
	purchDMmap := GetPurchaseDMmap(code + "7")
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
	saleDMmap := GetSaleDMmap(code + "7")
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

//////////////////////////////////////////////SALE

func confirm(tgCodes []string) []string {
	var confirmed []string
	validCodes := allTradeGoodsRCodes()
	for i := range tgCodes {
		for j := range validCodes {
			if validCodes[j] == tgCodes[i] {
				confirmed = append(confirmed, tgCodes[i])
				continue
			}
		}
	}
	return confirmed
}

// func (m Merchant) SaleProposalLegal(code string, amount int) string {
// 	basePrice := GetBasePrice(code)
// 	salePrice := m.CostSale(code)
// 	profit := (salePrice - basePrice) * amount
// 	if profit < 0 {
// 		//	profit = 0
// 	}
// 	tax := taxingAmount(profit, string([]byte(m.localUWP)[5]))
// 	proposal := ""
// 	//fmt.Println("Base", basePrice, "sale", salePrice, "#profit", profit, "||tax", tax, GetDescription(code))
// 	proposal += "Trade Lot: " + strconv.Itoa(amount) + " x " + GetDescription(code) + "\n"
// 	proposal += " Proposal: " + strconv.Itoa(salePrice) + " x " + strconv.Itoa(amount) + " (" + strconv.Itoa(salePrice*amount) + " Cr)" + "\n"
// 	proposal += "      Tax: " + strconv.Itoa(tax) + " Cr" + "\n"
// 	proposal += "---------------------" + "\n"
// 	proposal += "   Profit: " + strconv.Itoa((salePrice-tax)*amount) + " Cr" + "\n"
// 	return proposal
// }

/*
sale table:
|Tons per Defined Trade Good|Proposal|Tax|Total Profit|Sell DM|

*/

func (m *Merchant) ListExport() []string {
	allCodes := allCategories()
	availableCategories := []string{}
	for i := range allCodes {
		pdm, _ := PurchSaleDMs(allCodes[i], m.localTC)
		fmt.Println("pdm", pdm)
		if matchWorldsTC(allCodes[i]+"7", m.localTC) {
			fmt.Println("Add Export", allCodes[i], GetCategory(allCodes[i]+"7"))
			availableCategories = append(availableCategories, allCodes[i])
		}
	}
	exportList := []string{}
	for i := range availableCategories {
		exportList = append(exportList, availableCategories[i]+" "+GetCategory(availableCategories[i]+"7"))
	}

	return exportList
}

func (m *Merchant) ListImport() []string {
	allCodes := allCategories()
	importList := []string{}
	for i := range allCodes {
		_, sdm := PurchSaleDMs(allCodes[i], m.localTC)
		fmt.Println(sdm)
		if sdm > 0 {
			importList = append(importList, allCodes[i]+" "+GetCategory(allCodes[i]+"7"))
			fmt.Println("add import: " + allCodes[i] + " " + GetCategory(allCodes[i]+"7"))
		}
	}

	return importList
}

func (m *Merchant) ListImportExport() ([]string, []string) {
	allCodes := allCategories()
	exportList := []string{}
	importList := []string{}
	for i := range allCodes {
		pdm, sdm := PurchSaleDMs(allCodes[i], m.localTC)
		if sdm > pdm && sdm > 0 {
			importList = append(importList, GetCategory(allCodes[i]+"7")+" ["+allCodes[i]+"]")
		}
		if pdm > sdm && pdm > 0 {
			exportList = append(exportList, GetCategory(allCodes[i]+"7")+" ["+allCodes[i]+"]")
		}
	}
	return importList, exportList
}
