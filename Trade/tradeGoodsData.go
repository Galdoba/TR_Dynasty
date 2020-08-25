package trade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

//TradeGoodR -
// type TradeGoodR struct {
// 	code                    string
// 	tgCategory              string
// 	availabilityTags        []string
// 	stockIncrementFormula   string
// 	basePrice               int
// 	purchaseDM              map[string]int
// 	saleDM                  map[string]int
// 	maximumRiskAssessmentDM int
// 	dangerousGoodsDM        int
// 	description             string
// }

// func TradeGood() *TradeGoodR {
// 	return &TradeGoodR{}
// }

// func (tgR *TradeGoodR) Info() string {
// 	fmt.Println("DEBUG: func (tgR *TradeGoodR) Info() string ")
// 	fmt.Println("tgR Code", tgR.code)
// 	fmt.Println("tgR Descr", tgR.description)
// 	fmt.Println("tgR basePrice", tgR.basePrice, "Cr/Ton")
// 	fmt.Println("END DEBUG")
// 	return ""
// }

//NewTradeGoodR -
// func NewTradeGoodR(code string) *TradeGoodR {
// 	//fmt.Println(code, "code")
// 	//TODO: попробывать оптимизировать создание товара из других модулей
// 	//(сейчас он создает 8 раз Мар на 600 записей при каждом создании TradeGoodR)
// 	tgr := &TradeGoodR{}
// 	tgr.code = code
// 	tgr.purchaseDM = make(map[string]int)
// 	tgr.saleDM = make(map[string]int)
// 	tgr.tgCategory = GetCategory(code)
// 	tgr.description = GetDescription(code)
// 	tgr.availabilityTags = GetAvailabilityTags(code)
// 	tgr.stockIncrementFormula = GetStockIncrementFormula(code)
// 	tgr.basePrice = GetBasePrice(code)
// 	tgr.purchaseDM = GetPurchaseDMmap(code)
// 	tgr.saleDM = GetSaleDMmap(code)
// 	tgr.maximumRiskAssessmentDM = GetMaximumRiskAssessment(code)
// 	tgr.dangerousGoodsDM = GetDangerousGoodsDM(code)

// 	return tgr
// }

// func (tgr *TradeGoodR) Info() {
// 	fmt.Println(tgr.code)
// 	fmt.Println(tgr.tgCategory)
// 	fmt.Println(tgr.availabilityTags)
// 	fmt.Println(tgr.stockIncrementFormula)
// 	fmt.Println(tgr.basePrice)
// 	fmt.Println(tgr.purchaseDM)
// 	fmt.Println(tgr.saleDM)
// 	fmt.Println(tgr.maximumRiskAssessmentDM)
// 	fmt.Println(tgr.dangerousGoodsDM)
// 	fmt.Println(tgr.description)
// }

// func actualPriceDM(actualTags []string, tgr *TradeGoodR, operationType string) int {
// 	maxVal := -999
// 	transactionDMmap := make(map[string]int)
// 	switch operationType {
// 	case "P":
// 		transactionDMmap = tgr.purchaseDM
// 	case "S":
// 		transactionDMmap = tgr.saleDM
// 	}
// 	for i := range actualTags {
// 		testTag := tradeCodeFullName(actualTags[i])
// 		if maxVal < transactionDMmap[testTag] {
// 			maxVal = transactionDMmap[testTag]
// 		}
// 	}
// 	return maxVal
// }

func GetCategory(code string) string {
	return tgDB[code][0]
}

func GetDescription(code string) string {
	//descrMerged := tgDB[code][1]
	//descriptionAll := strings.Split(descrMerged, "/")
	return tgDB[code][1]
}

func GetAvailabilityTags(code string) []string {
	tagsMerged := tgDB[code][2]
	tagsAll := strings.Split(tagsMerged, ", ")
	if tagsAll[0] == "All" {
		return tradeTagLIST()
	}
	return tagsAll
}

func GetStockIncrementFormula(code string) string {
	return tgDB[code][3]
}

func GetBasePrice(code string) int {
	res, err := strconv.Atoi(tgDB[code][4])
	if err != nil {
		panic("TODO EXPALANATION: GetBasePrice()")
	}
	return res
}

func GetPurchaseDMmap(code string) map[string]int {
	pMap := make(map[string]int)
	rawData := tgDB[code][5]
	tagsMerged := strings.Split(rawData, ", ")
	for i := range tagsMerged {
		key := tagsMerged[i][0 : len(tagsMerged[i])-3]
		val := tagsMerged[i][len(tagsMerged[i])-2:]
		pMap[key] = convert.StoI(val)
	}
	return pMap
}

func GetSaleDMmap(code string) map[string]int {
	//TradeGoodsDataMap := TradeGoodRData()
	pMap := make(map[string]int)
	rawData := tgDB[code][6]
	tagsMerged := strings.Split(rawData, ", ")
	for i := range tagsMerged {
		key := tagsMerged[i][0 : len(tagsMerged[i])-3]
		val := tagsMerged[i][len(tagsMerged[i])-2:]
		pMap[key] = convert.StoI(val)
	}
	return pMap
}

func GetMaximumRiskAssessment(code string) int {
	res, err := strconv.Atoi(tgDB[code][7])
	if err != nil {
		panic(err)
	}
	return res
}

func GetDangerousGoodsDM(code string) int {
	res, err := strconv.Atoi(tgDB[code][8])
	if err != nil {
		panic(err)
	}
	return res
}

func GetMaximumForCategoryFormula(code string) string {
	return tgDB[code][9]
}

//RollMaximumForCategory -
func RollMaximumForCategory(code string) int {
	qty := ""
	formula := GetMaximumForCategoryFormula(code + "7")
	rawAdd := strings.Split(formula, " + ")
	rawQty := strings.Split(rawAdd[0], " x ")
	if len(rawQty) < 2 {
		fmt.Println(rawQty[1])
		panic(code + " formula Error")
	}
	qty = rawQty[1]
	up := utils.RollDiceRandom(qty+"d6") * convert.StoI(qty)
	return up
}

//IncreseTG -
func IncreseTG(code string) int {
	qty := ""
	add := "1"
	formula := GetStockIncrementFormula(code)
	rawAdd := strings.Split(formula, " + ")
	if len(rawAdd) > 1 {
		add = rawAdd[1]
	}
	rawQty := strings.Split(rawAdd[0], " x ")
	if len(rawQty) < 2 {
		fmt.Println(rawQty[1])
		panic(code + " formula Error")
	}
	qty = rawQty[1]
	up := utils.RollDiceRandom(qty+"d6") * convert.StoI(add)
	return up
}

// func (tgr *TradeGoodR) IncreaseRandom() int {
// 	qty := ""
// 	add := "0"
// 	formula := tgr.stockIncrementFormula
// 	rawAdd := strings.Split(formula, " + ")
// 	if len(rawAdd) > 1 {
// 		add = rawAdd[1]
// 	}
// 	rawQty := strings.Split(rawAdd[0], " x ")
// 	if len(rawQty) < 2 {
// 		fmt.Println(rawQty[1])
// 		panic(tgr.code + " formula Error")
// 	}
// 	qty = rawQty[1]
// 	up := utils.RollDiceRandom(qty+"d6", convert.StoI(add))
// 	return up
// }
