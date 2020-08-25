package trade

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/convert"

	"github.com/Galdoba/utils"
)

//Cogri     0101   CA6A643-9 N  Ri Wa      A

const (
	planetStatSize = "Size"
	planetStatAtmo = "Atmo"
	planetStatHydr = "Hydr"
	planetStatPops = "Pops"
	planetStatGovr = "Govr"
	planetStatLaws = "Laws"
	planetStatTech = "Tech"
	//eXResources              = "Resources"
	//eXLabor                  = "Labor"
	//eXInfrastructure         = "Infrastructure"
	//eXCulture                = "Culture"
	tradeCodeAgricultural    = "Ag"
	tradeCodeAsteroid        = "As"
	tradeCodeBarren          = "Ba"
	tradeCodeDesert          = "De"
	tradeCodeFluidOceans     = "Fl"
	tradeCodeGarden          = "Ga"
	tradeCodeHighPopulation  = "Hi"
	tradeCodeHighTech        = "Ht"
	tradeCodeIceCapped       = "Ie"
	tradeCodeIndustrial      = "In"
	tradeCodeLowPopulation   = "Lo"
	tradeCodeLowTech         = "Lt"
	tradeCodeNonAgricultural = "Na"
	tradeCodeNonIndustrial   = "Nl"
	tradeCodePoor            = "Po"
	tradeCodeRich            = "Ri"
	tradeCodeVacuum          = "Va"
	tradeCodeWaterWorld      = "Wa"
	travelCodeAmber          = "AZ"
	travelCodeRed            = "RZ"
	portBaseNaval            = "N"
	portBaseScout            = "S"
	portBaseResearch         = "R"
	portBaseTAS              = "T"
	supplierTypeCommon       = "Common"
	supplierTypeTrade        = "Trade"
	supplierTypeNeutral      = "Neutral"
	supplierTypeIlligal      = "Illigal"

	iXImportance  = "Importance"
	tempFrosen    = "Frosen"
	tempCold      = "Cold"
	tempTemperate = "Termperate"
	tempHot       = "Hot"
	tempBoiling   = "Boiling"
)

//var TradeGoodsByCode map[string]*tradeGood

/*
Trade Checklist
1. Find a supplier .............skip
2. Determine goods available ...ок
3. Determine purchase price ....ok
4. Purchase goods
5. Travel to another market
6. Find a buyer
7. Determine sale price
*/

/*
start
+-+Trade
|  |
|  +Buy
|  +Sell
+-+Freight
   |
   +Mail
   +Passengers
   +Goods



*/

func startMENU() {
	utils.ClearScreen()
	fmt.Println("Imperial Trade Terminal welcomes you!")
	utils.TakeOptions("Select Action Menu:", "Trade", "Freight")
}

// //var TradeGoodsDataMap map[string][]string
// var CharacterBrokerSkill int
// var CharacterInvestigateSkill int
// var CharacterStreetwiseSkill int
// var CharacterSocialDM int
// var RefereeDM int
// var sourceWorldUPP string

// type tradeLot struct {
// 	lotTradeGoodR *TradeGoodR
// 	cargoVolume   int
// 	cost          int
// 	purchseDice   int
// }

func trimDefinition(def int) int {
	switch def {
	case 3, 4, 5:
		return 4
	case 6, 7, 8:
		return 7
	case 9, 10, 11:
		return 10
	}
	return def
}

//ExcludeFromSliceStr -
func ExcludeFromSliceStr(slice []string, elem string) []string {
	for i := 0; i < len(slice); i++ {
		if elem == slice[i] {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}

// func NewTradeLot(tgCode string, planet *world.World) *tradeLot {
// 	tLot := &tradeLot{}
// 	tLot.lotTradeGoodR = NewTradeGoodR(tgCode)
// 	tLot.cargoVolume = tLot.cargoVolume + tLot.lotTradeGoodR.IncreaseRandom()
// 	purchaseDM := actualPriceDM(planet.TradeCodes(), tLot.lotTradeGoodR, "P")
// 	saleDM := actualPriceDM(planet.TradeCodes(), tLot.lotTradeGoodR, "S")
// 	tLot.purchseDice = utils.RollDice("3d6") + purchaseDM - saleDM + CharacterBrokerSkill + CharacterSocialDM + RefereeDM
// 	tLot.cost = modifyPricePurchase(tLot.lotTradeGoodR.basePrice, tLot.purchseDice)
// 	return tLot
// }

func rollTons(data string) int {
	unsplitted := strings.Split(data, " + ")
	diceUn := strings.Split(unsplitted[0], " x ")
	dice := diceUn[1] + "d6"
	add := 0
	if len(unsplitted) > 1 {
		add = convert.StoI(unsplitted[1])
	}
	increment := utils.RollDice(dice, add)
	return increment
}

// func availableCategories(planet *world.World) (availableCategories []string) {
// 	planetCodes := planet.TradeCodes()
// 	categories := categoryCodesLIST()
// 	for i := range planetCodes {
// 		for j := range categories {
// 			categ := categories[j]
// 			testTgr := NewTradeGoodR(categ + "7")
// 			for k := range testTgr.availabilityTags {
// 				if testTgr.availabilityTags[k] == tradeCodeFullName(planetCodes[i]) {
// 					availableCategories = utils.AppendUniqueStr(availableCategories, categ)
// 				}
// 			}
// 		}

// 	}
// 	return availableCategories
// }

func availableFromTCodes(planetCodes []string) (availableCategories []string) {
	categories := categoryCodesLIST()
	for j := range categories {
		categ := categories[j]
		if ok, _ := matchTradeCodes(planetCodes, GetAvailabilityTags(categ+"7")); ok {
			availableCategories = utils.AppendUniqueStr(availableCategories, categ)
		}
	}
	return availableCategories
}

func matchTradeCodes(sl1, sl2 []string) (bool, string) {
	for i := range sl1 {
		for j := range sl2 {
			if sl1[i] == sl2[j] {
				return true, sl1[i]
			}
		}
	}
	return false, ""
}

func categoryCodesLIST() []string {
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
	}
}

func tradeGoodsCodesLIST() []string {
	list := []string{}
	categorylist := categoryCodesLIST()
	for i := range categorylist {
		for definition := 2; definition <= 12; definition++ {
			list = append(list, categorylist[i]+convert.ItoS(definition))
		}
	}
	return list
}

func modifyPricePurchase(basePrice int, result int) int {
	x := 0
	switch result {
	default:
		if result < -1 {
			x = 400
		}
		if result > 24 {
			x = 25
		}
	case -1:
		x = 400
	case 0:
		x = 300
	case 1:
		x = 200
	case 2:
		x = 175
	case 3:
		x = 150
	case 4:
		x = 135
	case 5:
		x = 125
	case 6:
		x = 120
	case 7:
		x = 115
	case 8:
		x = 110
	case 9:
		x = 105
	case 10:
		x = 100
	case 11:
		x = 95
	case 12:
		x = 90
	case 13:
		x = 85
	case 14:
		x = 80
	case 15:
		x = 75
	case 16:
		x = 70
	case 17:
		x = 65
	case 18:
		x = 60
	case 19:
		x = 55
	case 20:
		x = 50
	case 21:
		x = 45
	case 22:
		x = 40
	case 23:
		x = 30
	case 24:
		x = 25
	}
	modifiedPrice := basePrice / 100 * x
	return modifiedPrice
}

func modifyPriceSale(basePrice int, result int) int {
	x := 0
	switch result {
	default:
		if result < -1 {
			x = 25
		}
		if result > 24 {
			x = 400
		}
	case -1:
		x = 25
	case 0:
		x = 30
	case 1:
		x = 40
	case 2:
		x = 45
	case 3:
		x = 50
	case 4:
		x = 55
	case 5:
		x = 60
	case 6:
		x = 65
	case 7:
		x = 70
	case 8:
		x = 75
	case 9:
		x = 80
	case 10:
		x = 85
	case 11:
		x = 90
	case 12:
		x = 95
	case 13:
		x = 100
	case 14:
		x = 105
	case 15:
		x = 110
	case 16:
		x = 115
	case 17:
		x = 120
	case 18:
		x = 125
	case 19:
		x = 135
	case 20:
		x = 150
	case 21:
		x = 175
	case 22:
		x = 200
	case 23:
		x = 300
	case 24:
		x = 400
	}
	modifiedPrice := basePrice / 100 * x
	return modifiedPrice
}

func commonElements(array1, array2 []string) []string {
	result := []string{}
	for i := range array1 {
		for j := range array2 {
			if array1[i] == array2[j] {
				result = append(result, array1[i])
			}
		}
	}
	return result
}

func parseAvailability(data string) (tags []string) {
	if data == "All" {
		return listTradeCodes()
	}
	tcLongNames := listTradeCodesLong()
	for i := range tcLongNames {
		if strings.Contains(data, tcLongNames[i]) {
			tags = append(tags, listTradeCodes()[i])
		}
	}

	return tags
}

func parseTradeDM(data string) map[string]int {
	tradeMDmap := make(map[string]int)
	tags := strings.Split(data, ", ")
	tcLongNames := listTradeCodesLong()
	tcLongNames = append(tcLongNames, "Amber Zone")
	tcLongNames = append(tcLongNames, "Red Zone")
	for next := range tags {
		//fmt.Println("Tag:", tags[next])
		for i := range tcLongNames {
			if !strings.Contains(tags[next], tcLongNames[i]) {
				//		fmt.Println("remove", tcLongNames[i])
				continue
			}
			//	fmt.Println("perse", tcLongNames[i])
			val := utils.ParseValueInt(tags[next], tcLongNames[i])
			tradeMDmap[listTradeCodes()[i]] = val

			break
		}
	}
	return tradeMDmap
}

func trvlRoll(diceCode string) int {
	dieType := "6"
	diceCode = strings.ToLower(diceCode)
	return utils.RollDice(diceCode + dieType)
}

func codeFromLine(line string) string {
	parts := strings.Split(line, "	")
	return parts[0]
}

func typeFromLine(line string) string {
	parts := strings.Split(line, "	")
	return parts[1]
}

func listTradeCodes() []string {
	str := []string{
		tradeCodeAgricultural,
		tradeCodeAsteroid,
		tradeCodeBarren,
		tradeCodeDesert,
		tradeCodeFluidOceans,
		tradeCodeGarden,
		tradeCodeHighPopulation,
		tradeCodeHighTech,
		tradeCodeIceCapped,
		tradeCodeIndustrial,
		tradeCodeLowPopulation,
		tradeCodeLowTech,
		tradeCodeNonAgricultural,
		tradeCodeNonIndustrial,
		tradeCodePoor,
		tradeCodeRich,
		tradeCodeVacuum,
		tradeCodeWaterWorld,
		travelCodeAmber,
		travelCodeRed,
	}
	return str
}

func listTradeCodesLong() []string {
	str := []string{
		"Agricultural",
		"Asteroid",
		"Barren",
		"Desert",
		"Fluid Oceans",
		"Garden",
		"High Population",
		"High Tech",
		"Ice-Capped",
		"Industrial",
		"Low Population",
		"Low Tech",
		"Non-Agricultural",
		"Non-Industrial",
		"Poor",
		"Rich",
		"Vacuum",
		"Water World",
	}
	return str
}

func listD66() []string {
	d66 := []string{
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
		"66",
	}
	return d66
}

func checkDangerousGoods() {
	factors := []string{
		"Having a Rival in local area (+1DM)", //0
		"Having an Enemy in local area (+2DM)",
		"Having a Contact in local area (-1DM)", //2
		"Having an Ally in local area (-2DM)",
		"Having a Buyer in local area (-3DM)", //4
		"Supplier in Amber Zone (+1DM)",
		"Supplier in Red Zone (+2DM)", //6
		"Supplier is Moraly Neutral (+1DM)",
		"Supplier is affiliated with Black market (+2DM)", //8
		"Supplier has major Govermental Backing... (+ (1d6 - 4DM))",
		"Supplier has major Govermental Backing... (- (1d6 - 4DM))", //10
	}
	_, pickedBool := utils.SelectionOptionsMult("Pick all factors: ", factors...)
	checkDM := 0
	for i := range pickedBool {
		switch i {
		case 0:
			if pickedBool[i] {
				checkDM = checkDM + 1
			}
		case 1:
			if pickedBool[i] {
				checkDM = checkDM + 2
			}
		case 2:
			if pickedBool[i] {
				checkDM = checkDM - 1
			}
		case 3:
			if pickedBool[i] {
				checkDM = checkDM - 2
			}
		case 4:
			if pickedBool[i] {
				checkDM = checkDM - 3
			}
		case 5:
			if pickedBool[i] {
				checkDM = checkDM + 1
			}
		case 6:
			if pickedBool[i] {
				checkDM = checkDM + 2
			}
		case 7:
			if pickedBool[i] {
				checkDM = checkDM + 1
			}
		case 8:
			if pickedBool[i] {
				checkDM = checkDM + 2
			}
		case 9:
			if pickedBool[i] {
				checkDM = checkDM + (TrvCore.Roll1D() - 4)
			}
		case 10:
			if pickedBool[i] {
				checkDM = checkDM - (TrvCore.Roll1D() - 4)
			}
		}
	}
	checkRoll := utils.RollDice("2d6", checkDM)
	fmt.Println("The Dangerous Cargo", checkRoll)
	switch checkRoll {
	default:

	}
}
