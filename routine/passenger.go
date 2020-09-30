package routine

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/wrld"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
)

//PassengerRoutine -
func PassengerRoutine() {

	printSlow("Searching for Passengers...\n")
	timeLimit := 0
	playerEffect1 := 0
	switch autoMod {
	case false:
		//playerEffect1 = userInputInt("Enter Effect of Carouse(8) or Streetwise(8) check: ")
		input := userInputIntSlice("Enter Effect of Carouse(8) or Streetwise(8) check (and time limit in days after ' ' if nesessary): ")
		if len(input) > 0 {
			playerEffect1 = input[0]
		}
		if len(input) > 1 {
			timeLimit = input[1]
		}
	case true:
		playerEffect1 = autoFlux()
	}
	playerEffect1, time, abort := mutateTestResultsByTime(playerEffect1, dice.Roll("1d6").Sum(), timeLimit)
	if abort {
		fmt.Println("Search aborted after", time, "days...")
	}
	fmt.Println("Search took", time, "days...")
	// if gmMode {
	// 	fmt.Println("GM TIP: Passenger Roll:", ptValue, playerEffect1, localBroker.DM(), "|", ptValue+playerEffect1+localBroker.DM())
	// }

	low, basic, middle, high := availablePassengers(ptValue + playerEffect1 + localBroker.DM())
	printSlow("Active passenger requests: " + strconv.Itoa(low+basic+middle+high) + "\n")
	if low > 0 {
		fee := lowPassCost(jumpRoute) - localBroker.CutFrom(lowPassCost(jumpRoute))
		printSlow("   Low Passengers: " + strconv.Itoa(low) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	if basic > 0 {
		fee := basicPassCost(jumpRoute) - localBroker.CutFrom(basicPassCost(jumpRoute))
		printSlow(" Basic Passengers: " + strconv.Itoa(basic) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	if middle > 0 {
		fee := middlePassCost(jumpRoute) - localBroker.CutFrom(middlePassCost(jumpRoute))
		printSlow("Middle Passengers: " + strconv.Itoa(middle) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	if high > 0 {
		fee := highPassCost(jumpRoute) - localBroker.CutFrom(highPassCost(jumpRoute))
		printSlow("  High Passengers: " + strconv.Itoa(high) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	fmt.Println("-----------------------------------------------------")
}

//ABCD
//EFGH
//IJKL
//MNOP
func passengerTrafficValue(sourceWorld, targetWorld wrld.World) int {
	//dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrPops))
	uwp, _ := profile.NewUWP(sourceWorld.UWP())
	dm := TrvCore.EhexToDigit(uwp.Pops())
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

	dm += techDifferenceDM()
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
		dm = dp.RollNext("1d6").Sum()
	case 5:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
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
	lp := dp.RollNext(d).DM(dm).Sum()
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
		dm = dp.RollNext("1d6").Sum()
	case 5:
		d = "2d6"
		dm = dp.RollNext("1d6").Sum()
	case 6:
		d = "2d6"
	case 7:
		d = "3d6"
		dm = dp.RollNext("2d6").Sum()
	case 8:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 9:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
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
	mp := dp.RollNext(d).DM(dm).Sum()
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
		dm = dp.RollNext("1d6").Sum()
	case 4:
		d = "2d6"
		dm = dp.RollNext("2d6").Sum()
	case 5:
		d = "2d6"
		dm = dp.RollNext("1d6").Sum()
	case 6:
		d = "3d6"
		dm = dp.RollNext("2d6").Sum()
	case 7:
		d = "3d6"
		dm = dp.RollNext("2d6").Sum()
	case 8:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 9:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 10:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
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
	hp := dp.RollNext(d).DM(dm).Sum()
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
