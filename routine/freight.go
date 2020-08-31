package routine

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

func freightTrafficValue(sourceWorld, targetWorld world.World) int {
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

	dm += TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrPops)) // - для грузов важно население targetWorld
	for _, val := range targetWorld.TradeCodes() {
		fmt.Println("check", val)
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

	tl1 := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrTL))
	tl2 := TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrTL))
	tlDiff := utils.Max(tl1, tl2) - utils.Min(tl1, tl2)
	if tlDiff > 5 {
		tlDiff = 5
	}
	dm -= tlDiff
	return dm
}

func availableFreight(tfv int) (int, int, int) {

	inLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-13).Sum())
	mnLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-8).Sum())
	// mjLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-6).Sum())

	mjLots := utils.Max(0, dp.RollNext("1d6").DM(tfv-6).Sum())

	return inLots, mnLots, mjLots
}
