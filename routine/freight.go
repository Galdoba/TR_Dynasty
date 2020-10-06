package routine

import (
	"fmt"
	"sort"
	"strconv"

	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/utils"
)

var freightBase int

func FreightRoutine() {
	printSlow("Searching for Freight...\n")

	//diff := freightDiff(ftValue)
	//playerEffect2 := userInputInt("Enter Effect of Diplomat(" + strconv.Itoa(diff) + "), Investigate(" + strconv.Itoa(diff) + ") or Streetwise(" + strconv.Itoa(diff) + ") check: ")
	//playerEffect2 := userInputInt("Enter Effect of Diplomat(8), Investigate(8) or Streetwise(8) check: ")
	timeLimit := 0
	playerEffect2 := 0
	switch autoMod {
	case false:
		//playerEffect2 = userInputStr("Enter Effect of Diplomat(8), Investigate(8) or Streetwise(8) check: ")
		input := userInputIntSlice("Enter Effect of Diplomat(8), Investigate(8) or Streetwise(8) check (and time limit in days after ' ' if nesessary): ")
		if len(input) > 0 {
			playerEffect2 = input[0]
		}
		if len(input) > 1 {
			timeLimit = input[1]
		}

	case true:
		playerEffect2 = autoFlux()

	}
	playerEffect2, time, abort := mutateTestResultsByTime(playerEffect2, dice.Roll("1d6").Sum(), timeLimit)
	if abort {
		fmt.Println("Search aborted after", time, "days...")
	}
	fmt.Println("Search took", time, "days...")
	//spendTime(playerEffect2, timeLimit)
	// if gmMode {
	// 	fmt.Println("GM TIP: Freight Roll:", ftValue, playerEffect2, localBroker.DM(), "|", ftValue+playerEffect2+localBroker.DM())
	// }
	inLot, mnLot, mjLot := availableFreight(ftValue + playerEffect2 + localBroker.DM())
	//fmt.Println(inLot, mnLot, mjLot)
	frList := freightListed(inLot, mnLot, mjLot)
	fmt.Println(frList)
	//TODO: Create lots here
	if len(frList) < 1 {
		printSlow("No freight lots available\n")
	}
	for i := range frList {
		base := 500
		fee := frList[i]*freightCostPerTon(base) - localBroker.CutFrom(frList[i]*freightCostPerTon(base))
		lot := newCargoLot()
		lot.detailsFreight(determineLot(), frList[i], fee)
		portCargo = append(portCargo, lot)
		//fmt.Println(lot)
		//printSlow("Freight lot " + strconv.Itoa(i+1) + " 		" + strconv.Itoa(frList[i]) + " tons		Hauling fee: " + strconv.Itoa(fee) + " Cr\n")
		fmt.Println("Lot", i, "	"+lot.GetTGCode()+"	", lot.GetVolume(), "tons  	", lot.GetComment()+"	", lot.GetDescr())
	}
	fmt.Println("-----------------------------------------------------")

}

func lotInfo(lot cargoLot) string {
	return lot.GetTGCode() + "	" + strconv.Itoa(lot.GetVolume()) + " tons  	" + lot.GetComment() + "	" + lot.GetDescr()
}

func determineLot() string {
	randomLot := trade.RandomTGCategory(sourceWorld) + dice.Roll("2d6").SumStr()
	return randomLot
}

func freightDiff(ftValue int) int {
	diff := 6
	switch ftValue {
	case 1, 2, 3, 4:
		diff = 6
	case 5, 6, 7:
		diff = 7
	case 8, 9, 10, 11:
		diff = 8
	case 12, 13, 14:
		diff = 9
	}
	if ftValue > 14 {
		diff = 10
	}
	return diff
}

func freightCostPerTon(base int) int {
	cpt := 0
	for _, val := range jumpRoute {
		cpt = cpt + (base * val)
	}

	return cpt
}

func freightListed(inLot, mnLot, mjLot int) []int {
	var tons []int
	//printSlow("Searching available Freight lots...\n")
	for i := 0; i < mjLot; i++ {
		tons = append(tons, dp.RollNext(dp.RollNext("d6").DM(6).SumStr()+"d6").Sum()) //*10)
	}
	for i := 0; i < mnLot; i++ {
		tons = append(tons, dp.RollNext(dp.RollNext("1d6").SumStr()+"d6").Sum()) //*5)
	}
	for i := 0; i < inLot; i++ {
		tons = append(tons, dp.RollNext("1d6").Sum()*1)
	}
	sort.Ints(tons)
	printSlow("Found " + strconv.Itoa(len(tons)) + " active requests...\n")
	return tons

}

func freightTrafficValue(sourceWorld, targetWorld wrld.World) int {
	//dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData("Pops"))
	dm := 0
	for _, val := range sourceWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 2
		case constant.TradeCodeAsteroid:
			dm += -3
		case constant.TradeCodeBarren:
			return 0
		case constant.TradeCodeDesert:
			dm += -3
		case constant.TradeCodeFluidOceans:
			dm += -3
		case constant.TradeCodeGarden:
			dm += 2
		case constant.TradeCodeHighPopulation:
			dm += 2
		case constant.TradeCodeIceCapped:
			dm += -3
		case constant.TradeCodeIndustrial:
			dm += 3
		case constant.TradeCodeLowPopulation:
			dm += -5
		case constant.TradeCodeNonAgricultural:
			dm += -3
		case constant.TradeCodeNonIndustrial:
			dm += -3
		case constant.TradeCodePoor:
			dm += -3
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += -3
		case constant.TravelCodeAmber:
			dm += 5
		case constant.TravelCodeRed:
			dm += -5
		}
	}
	switch sourceWorld.TravelZone() {
	case constant.TravelCodeAmber:
		dm += 5
	case constant.TravelCodeRed:
		dm += -5
	}

	//dm += TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrPops)) // - для грузов важно население targetWorld
	uwp, _ := profile.NewUWP(targetWorld.UWP())
	dm += TrvCore.EhexToDigit(uwp.Pops())
	for _, val := range targetWorld.TradeCodes() {
		//fmt.Println("check", val)
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm++
		case constant.TradeCodeAsteroid:
			dm++
		case constant.TradeCodeBarren:
			dm += -5
		case constant.TradeCodeDesert:
			dm += 0
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm++
		case constant.TradeCodeHighPopulation:
			dm += 0
		case constant.TradeCodeIceCapped:
			dm += 0
		case constant.TradeCodeIndustrial:
			dm += 2
		case constant.TradeCodeLowPopulation:
			dm += 0
		case constant.TradeCodeNonAgricultural:
			dm++
		case constant.TradeCodeNonIndustrial:
			dm++
		case constant.TradeCodePoor:
			dm += -3
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += 0
		}
	}
	switch targetWorld.TravelZone() {
	case constant.TravelCodeAmber:
		dm += -5
	case constant.TravelCodeRed:
		return -999
	}
	//}

	dm += techDifferenceDM()
	return dm
}

func availableFreight(tfv int) (int, int, int) {

	inLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-13).Sum())
	mnLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-8).Sum())
	// mjLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-6).Sum())

	mjLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-6).Sum())

	return inLots, mnLots, mjLots
}
