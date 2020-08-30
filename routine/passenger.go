package routine

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/dice"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
)

//PassengerRoutine -
func PassengerRoutine() {

	printSlow("Searching for Passengers...\n")
	playerEffect := userInputInt("Enter Effect of Broker, Carouse or Streetwise: ")

	//pFactor += pFactorPops(sourceWorld.PlanetaryData(constant.PrPops))
	//pFactor += pfFactorSp(sourceWorld.PlanetaryData(constant.PrStarport))

	//fmt.Println("pFactor", pFactor)
	clrScrn()
	// lowPass := rollPassengerTrafficTable(pFactor + 1 + playerEffect)
	// midPass := rollPassengerTrafficTable(pFactor + playerEffect)
	// basicPass := rollPassengerTrafficTable(pFactor + playerEffect)
	// highPass := rollPassengerTrafficTable(pFactor - 4 + playerEffect)
	// fmt.Println(lowPass, basicPass, midPass, highPass)
	fmt.Println(AvailablePassengers(pFactor + playerEffect))
}

// func pFactorPops(pops string) int {
// 	popInt := TrvCore.EhexToDigit(pops)
// 	switch popInt {
// 	case 0, 1:
// 		return -4
// 	case 6, 7:
// 		return 1
// 	default:
// 		if popInt >= 8 {
// 			return 3
// 		}
// 	}
// 	return 0
// }

// func pfFactorSp(sp string) int {
// 	switch sp {
// 	case "A":
// 		return 2
// 	case "B":
// 		return 1
// 	case "E":
// 		return -1
// 	case "X":
// 		return -3
// 	}
// 	return 0
// }

func rollPassengerTrafficTable(pF int) int {
	pass := pF // dp.RollNext("2d6").DM(pF).Sum()
	d := 0
	switch pass {
	default:
		if pass <= 1 {
			return 0
		}
		if pass >= 20 {
			d = 10
		}
	case 2, 3:
		d = 1
	case 4, 5, 6:
		d = 2
	case 7, 8, 9, 10:
		d = 3
	case 11, 12, 13:
		d = 4
	case 14, 15:
		d = 5
	case 16:
		d = 6
	case 17:
		d = 7
	case 18:
		d = 8
	case 19:
		d = 9
	}
	ps := dp.RollNext(strconv.Itoa(d) + "d6").Sum()
	if ps < 0 {
		ps = 0
	}
	return ps
}

func passengerTrafficValue(sourceWorld, targetWorld world.World) int {
	dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData("Pops"))
	for _, val := range sourceWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 0
		case constant.TradeCodeAsteroid:
			dm += 1
		case constant.TradeCodeBarren:
			dm += 5
		case constant.TradeCodeDesert:
			dm += -1
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm += 2
		case constant.TradeCodeHighPopulation:
			dm += 0
		case constant.TradeCodeIceCapped:
			dm += 1
		case constant.TradeCodeIndustrial:
			dm += 2
		case constant.TradeCodeLowPopulation:
			dm += 0
		case constant.TradeCodeNonAgricultural:
			dm += 0
		case constant.TradeCodeNonIndustrial:
			dm += 0
		case constant.TradeCodePoor:
			dm += -2
		case constant.TradeCodeRich:
			dm += -1
		case constant.TradeCodeWaterWorld:
			dm += 0
		case constant.TravelCodeAmber:
			dm += 2
		case constant.TravelCodeRed:
			dm += 4
		}
	}
	//Starport Value
	// switch sourceWorld.StarPort() {
	// case "A":
	// 	dm += 2
	// case "B":
	// 	dm += 1
	// case "E":
	// 	dm += -1
	// case "X":
	// 	dm += -3
	// }
	//dm += TrvCore.EhexToDigit(targetWorld.PlanetaryData("Pops")) // - да хер пойми надо оно или нет (конфликт правил MgT1:MP p.66)
	for _, val := range targetWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 0
		case constant.TradeCodeAsteroid:
			dm += -1
		case constant.TradeCodeBarren:
			dm += -5
		case constant.TradeCodeDesert:
			dm += -1
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm += 0
		case constant.TradeCodeHighPopulation:
			dm += 4
		case constant.TradeCodeIceCapped:
			dm += -1
		case constant.TradeCodeIndustrial:
			dm += 1
		case constant.TradeCodeLowPopulation:
			dm += -4
		case constant.TradeCodeNonAgricultural:
			dm += 0
		case constant.TradeCodeNonIndustrial:
			dm += -1
		case constant.TradeCodePoor:
			dm += -1
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += 0
		case constant.TravelCodeAmber:
			dm += -2
		case constant.TravelCodeRed:
			dm += -4
		}
	}

	tl1 := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrTL))
	tl2 := TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrTL))
	tlDiff := utils.Max(tl1, tl2) - utils.Min(tl1, tl2)
	if tlDiff > 5 {
		tlDiff = 5
	}
	dm -= tlDiff
	fmt.Println(dm)
	return dm
}

func AvailablePassengers(r int) (int, int, int, int) {
	return lowPass(r), midPass(r), midPass(r), hghPass(r)
}

func lowPass(ptv int) int {
	if ptv <= 0 {
		return 0
	}
	dm := 0
	d := ""
	switch ptv {
	case 1:
		d = "2d6"
		dm = -6
	case 2:
		d = "2d6"
	case 3:
		d = "2d6"
	case 4:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 5:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 6:
		d = "3d6"
	case 7:
		d = "3d6"
	case 8:
		d = "4d6"
	case 9:
		d = "4d6"
	case 10:
		d = "5d6"
	case 11:
		d = "5d6"
	case 12:
		d = "6d6"
	case 13:
		d = "6d6"
	case 14:
		d = "7d6"
	case 15:
		d = "8d6"
	default:
		if ptv > 15 {
			d = "9d6"
		}
	}
	lp := dice.Roll(d).DM(dm).Sum()
	if lp < 0 {
		lp = 0
	}
	return lp
}

func midPass(ptv int) int {
	if ptv <= 0 {
		return 0
	}
	dm := 0
	d := ""
	switch ptv {
	case 1:
		return 0
	case 2:
		d = "1d6"
		dm = -2
	case 3:
		d = "1d6"
	case 4:
		d = "2d6"
		dm = dice.Roll("1d6").Sum()
	case 5:
		d = "2d6"
		dm = dice.Roll("1d6").Sum()
	case 6:
		d = "2d6"
	case 7:
		d = "3d6"
		dm = dice.Roll("2d6").Sum()
	case 8:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 9:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 10:
		d = "3d6"
	case 11:
		d = "4d6"
	case 12:
		d = "4d6"
	case 13:
		d = "4d6"
	case 14:
		d = "5d6"
	case 15:
		d = "5d6"
	default:
		if ptv > 15 {
			d = "6d6"
		}
	}
	mp := dice.Roll(d).DM(dm).Sum()
	if mp < 0 {
		mp = 0
	}
	return mp
}

func hghPass(ptv int) int {
	if ptv <= 0 {
		return 0
	}
	dm := 0
	d := ""
	switch ptv {
	case 1:
		return 0
	case 2:
		return 0
	case 3:
		d = "1d6"
		dm = dice.Roll("1d6").Sum()
	case 4:
		d = "2d6"
		dm = dice.Roll("2d6").Sum()
	case 5:
		d = "2d6"
		dm = dice.Roll("1d6").Sum()
	case 6:
		d = "3d6"
		dm = dice.Roll("2d6").Sum()
	case 7:
		d = "3d6"
		dm = dice.Roll("2d6").Sum()
	case 8:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 9:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 10:
		d = "3d6"
		dm = dice.Roll("1d6").Sum()
	case 11:
		d = "3d6"
	case 12:
		d = "3d6"
	case 13:
		d = "4d6"
	case 14:
		d = "4d6"
	case 15:
		d = "4d6"
	default:
		if ptv > 15 {
			d = "5d6"
		}
	}
	hp := dice.Roll(d).DM(dm).Sum()
	if hp < 0 {
		hp = 0
	}
	return hp
}
