package Trade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

type TradeGoodR struct {
	code                    string
	tgCategory              string
	availabilityTags        []string
	stockIncrementFormula   string
	basePrice               int
	purchaseDM              map[string]int
	saleDM                  map[string]int
	maximumRiskAssessmentDM int
	dangerousGoodsDM        int
	description             string
}

func Initiate() {
	tgDB = TradeGoodRData()
}

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

func NewTradeGoodR(code string) *TradeGoodR {
	//fmt.Println(code, "code")
	//TODO: попробывать оптимизировать создание товара из других модулей
	//(сейчас он создает 8 раз Мар на 600 записей при каждом создании TradeGoodR)
	tgr := &TradeGoodR{}
	tgr.code = code
	tgr.purchaseDM = make(map[string]int)
	tgr.saleDM = make(map[string]int)
	tgr.tgCategory = getCategory(code)
	tgr.description = getDescription(code)
	tgr.availabilityTags = getAvailabilityTags(code)
	tgr.stockIncrementFormula = getStockIncrementFormula(code)
	tgr.basePrice = getBasePrice(code)
	tgr.purchaseDM = getPurchaseDMmap(code)
	tgr.saleDM = getSaleDMmap(code)
	tgr.maximumRiskAssessmentDM = getMaximumRiskAssessment(code)
	tgr.dangerousGoodsDM = getDangerousGoodsDM(code)

	return tgr
}

func (tgr *TradeGoodR) Info() {
	fmt.Println(tgr.code)
	fmt.Println(tgr.tgCategory)
	fmt.Println(tgr.availabilityTags)
	fmt.Println(tgr.stockIncrementFormula)
	fmt.Println(tgr.basePrice)
	fmt.Println(tgr.purchaseDM)
	fmt.Println(tgr.saleDM)
	fmt.Println(tgr.maximumRiskAssessmentDM)
	fmt.Println(tgr.dangerousGoodsDM)
	fmt.Println(tgr.description)
}

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

func getCategory(code string) string {
	return tgDB[code][0]
}

func getDescription(code string) string {
	//descrMerged := tgDB[code][1]
	//descriptionAll := strings.Split(descrMerged, "/")
	return tgDB[code][1]
}

func getAvailabilityTags(code string) []string {
	tagsMerged := tgDB[code][2]
	tagsAll := strings.Split(tagsMerged, ", ")
	if tagsAll[0] == "All" {
		return tradeTagLIST()
	}
	return tagsAll
}

func getStockIncrementFormula(code string) string {
	return tgDB[code][3]
}

func getBasePrice(code string) int {
	return convert.StoI(tgDB[code][4])
}

func getPurchaseDMmap(code string) map[string]int {
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

func getSaleDMmap(code string) map[string]int {
	TradeGoodsDataMap := TradeGoodRData()
	pMap := make(map[string]int)
	rawData := TradeGoodsDataMap[code][6]
	tagsMerged := strings.Split(rawData, ", ")
	for i := range tagsMerged {
		key := tagsMerged[i][0 : len(tagsMerged[i])-3]
		val := tagsMerged[i][len(tagsMerged[i])-2:]
		pMap[key] = convert.StoI(val)
	}
	return pMap
}

func getMaximumRiskAssessment(code string) int {
	res, err := strconv.Atoi(tgDB[code][7])
	if err != nil {
		panic(err)
	}
	return res
}

func getDangerousGoodsDM(code string) int {
	res, err := strconv.Atoi(tgDB[code][8])
	if err != nil {
		panic(err)
	}
	return res
}

func IncreseTG(code string) int {
	qty := ""
	add := "0"
	formula := getStockIncrementFormula(code)
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
	up := utils.RollDiceRandom(qty+"d6", convert.StoI(add))
	return up
}

func (tgr *TradeGoodR) IncreaseRandom() int {
	qty := ""
	add := "0"
	formula := tgr.stockIncrementFormula
	rawAdd := strings.Split(formula, " + ")
	if len(rawAdd) > 1 {
		add = rawAdd[1]
	}
	rawQty := strings.Split(rawAdd[0], " x ")
	if len(rawQty) < 2 {
		fmt.Println(rawQty[1])
		panic(tgr.code + " formula Error")
	}
	qty = rawQty[1]
	up := utils.RollDiceRandom(qty+"d6", convert.StoI(add))
	return up
}
