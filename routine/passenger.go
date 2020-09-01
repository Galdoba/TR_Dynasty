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
	playerEffect1 := userInputInt("Enter Effect of Carouse(8) or Streetwise(8) check: ")

	low, basic, middle, high := availablePassengers(ptValue + playerEffect1)
	printSlow("Active passenger requests: " + strconv.Itoa(low+basic+middle+high) + "\n")
	if low > 0 {
		printSlow("   Low Passengers: " + strconv.Itoa(low) + " (" + strconv.Itoa(lowPassCost(jumpRoute)) + ")" + "\n")
	}
	if basic > 0 {
		printSlow(" Basic Passengers: " + strconv.Itoa(basic) + " (" + strconv.Itoa(basicPassCost(jumpRoute)) + ")" + "\n")
	}
	if middle > 0 {
		printSlow("Middle Passengers: " + strconv.Itoa(middle) + " (" + strconv.Itoa(middlePassCost(jumpRoute)) + ")" + "\n")
	}
	if high > 0 {
		printSlow("  High Passengers: " + strconv.Itoa(high) + " (" + strconv.Itoa(highPassCost(jumpRoute)) + ")" + "\n")
	}
	fmt.Println(https://docs.google.com/spreadsheets/d/1bThnY4Bgr0sU1NXcOJtDzKjO3THgx-VL-Lq8fhMDZk4/edit#gid=860849043&range=A11)
}

//ABCD
//EFGH
//IJKL
//MNOP
func passengerTrafficValue(sourceWorld, targetWorld world.World) int {
	dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrPops))
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
	switch sourceWorld.TravelZone() {
	case constant.TravelCodeAmber:
		dm += 2
	case constant.TravelCodeRed:
		dm += 4
	}
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
		}
	}
	switch targetWorld.TravelZone() {
	case constant.TravelCodeAmber:
		dm += -2
	case constant.TravelCodeRed:
		dm += -4
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

func availablePassengers(r int) (int, int, int, int) {
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

const (
	highPassenger = 0
	midPassenger  = 1
	basPassenger  = 2
	lowPassenger  = 3
)

func lowPassCost(jumpSeq []int) int {
	totalPrice := 0
	for _, val := range jumpSeq {
		if val < 1 {
			return 0
		}
		totalPrice = totalPrice + (800 + val*200)
	}
	return totalPrice
}

func basicPassCost(jumpSeq []int) int {
	totalPrice := 0
	for _, val := range jumpSeq {
		if val < 1 {
			return 0
		}
		if val == 1 {
			totalPrice = totalPrice + 1500
			continue
		}
		if val == 1 {
			totalPrice = totalPrice + 3000
			continue
		}
		totalPrice = totalPrice + ((val * 2500) - 2500)
	}
	return totalPrice
}

func middlePassCost(jumpSeq []int) int {
	// totalPrice := 0
	// for _, val := range jumpSeq {
	// 	if val < 1 {
	// 		return 0
	// 	}
	// 	if val == 1 {
	// 		totalPrice = totalPrice + 3000
	// 		continue
	// 	}
	// 	if val == 1 {
	// 		totalPrice = totalPrice + 6000
	// 		continue
	// 	}
	// 	totalPrice = totalPrice + ((val * 5000) - 5000)
	// }
	// return totalPrice
	return basicPassCost(jumpSeq) * 2
}

func highPassCost(jumpSeq []int) int {
	return middlePassCost(jumpSeq) * 2
}

// func basicPassCost() int {
// 	if dist < 1 {
// 		return 0
// 	}
// 	if dist == 1 {
// 		return 1500
// 	}
// 	if dist == 2 {
// 		return 3000
// 	}
// 	return (dist - 1) * 2500
// }

// func midddlePassCost() int {
// 	if dist < 1 {
// 		return 0
// 	}
// 	if dist == 1 {
// 		return 3000
// 	}
// 	if dist == 2 {
// 		return 6000
// 	}
// 	return (dist - 1) * 5000
// }

// func highPassCost() int {
// 	if dist < 1 {
// 		return 0
// 	}
// 	if dist == 1 {
// 		return 6000
// 	}
// 	if dist == 2 {
// 		return 12000
// 	}
// 	return (dist - 1) * 10000
// }

/*
Passage and Freight Costs
Parsecs Travelled	High Passage	Middle Passage	Basic Passage	Low Passage	Freight
1					Cr8500			Cr6200			Cr2200			Cr700		Cr1000
2					Cr12000			Cr9000			Cr2900			Cr1300		Cr1600
3					Cr20000			Cr15000			Cr4400			Cr2200		Cr3000
4					Cr41000			Cr31000			Cr8600			Cr4300		Cr7000
5					Cr45000			Cr34000			Cr9400			Cr13000		Cr7700
6					Cr470000		Cr350000		Cr93000			Cr96000		Cr86000


*/
