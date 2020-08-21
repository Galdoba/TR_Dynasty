package trademgt2

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

var sectorMapByHex map[string]string
var sectorMapByName map[string]string
var sectorMapByUWP map[string]string
var sourceWorld world.World
var targetWorld world.World

func init() {
	sectorData := otu.TrojanReachData()
	sectorMapByHex = otu.MapDataByHex(sectorData)
	sectorMapByName = otu.MapDataByName(sectorData)
	sectorMapByUWP = otu.MapDataByUWP(sectorData)

}

func RunMerchantPrince() {
	fmt.Println(" ")
	fmt.Println("///////////////////////")
	//sourceWorld := LoadWorld("Current World (Name, Hex or UWP): ")
	sourceWorld, _ = world.FromOTUdata(sectorMapByHex["2923"]) //2724
	//fmt.Println(sourceWorld)
	//targetWorld := LoadWorld("Target World (Name, Hex or UWP): ")
	targetWorld, _ = world.FromOTUdata(sectorMapByHex["2722"]) //2923
	dist := Astrogation.JumpDistance(sourceWorld.Hex(), targetWorld.Hex())
	fmt.Println("Jump Distance:", dist)
	fmt.Println("Freight Traffic Value: " + sourceWorld.Name() + " - " + targetWorld.Name())
	ftDMc := freightTrafficValue(sourceWorld, targetWorld)
	fl := 3
	inLot, mnLot, mjLot := AvailableFreightLots(ftDMc + fl)
	lots := LotList(inLot, mnLot, mjLot)
	for i := range lots {
		fmt.Println(lots[i])
		fmt.Println(lots[i].Price())
	}
	fmt.Println(ftDMc, fl)

	fmt.Println("///////////////////////")
	fmt.Println(sourceWorld)
	fmt.Println(targetWorld)
}

func LoadWorld(msg string) world.World {
	done := false
	key := ""
	otuData := ""
	for !done {
		key = userInputStr(msg)
		if val, ok := sectorMapByHex[key]; ok {
			otuData = val
			done = true
			continue
		}
		if val, ok := sectorMapByName[key]; ok {
			otuData = val
			done = true
			continue
		}
		key2 := strings.ToUpper(key)
		if val, ok := sectorMapByUWP[key2]; ok {
			otuData = val
			done = true
			continue
		}
		fmt.Println("No data by key '" + key + "'")
	}
	fmt.Println("Loading world data:")
	fmt.Println("Trojan Reach "+otu.Info{otuData}.Hex(), "-", otu.Info{otuData}.Name(), "("+otu.Info{otuData}.UWP()+")")
	w, err := world.FromOTUdata(otuData)
	if err != nil {
		panic(err)
	}
	return w
}

func userInputStr(msg string) string {
	done := false
	fmt.Print(msg)
	for !done {
		uwp, err := user.InputStr()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return uwp
	}
	return "Must not happen !!"
}

func userInputInt(msg string) int {
	done := false
	fmt.Print(msg)
	for !done {
		i, err := user.InputInt()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return i
	}
	return -999
}

func freightTrafficValue(sourceWorld, targetWorld world.World) int {
	dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData("Pops"))
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
	dm += TrvCore.EhexToDigit(targetWorld.PlanetaryData("Pops")) // - да хер пойми надо оно или нет (конфликт правил MgT1:MP p.66)
	for _, val := range targetWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 1
		case constant.TradeCodeAsteroid:
			dm += 1
		case constant.TradeCodeBarren:
			dm += -5
		case constant.TradeCodeDesert:
			dm += 0
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm += 1
		case constant.TradeCodeHighPopulation:
			dm += 0
		case constant.TradeCodeIceCapped:
			dm += 0
		case constant.TradeCodeIndustrial:
			dm += 2
		case constant.TradeCodeLowPopulation:
			dm += 0
		case constant.TradeCodeNonAgricultural:
			dm += 1
		case constant.TradeCodeNonIndustrial:
			dm += 1
		case constant.TradeCodePoor:
			dm += -3
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += 0
		case constant.TravelCodeAmber:
			dm += -5
		case constant.TravelCodeRed:
			return 0
		}
	}

	tl1 := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrTL))
	tl2 := TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrTL))
	tlDiff := utils.Max(tl1, tl2) - utils.Min(tl1, tl2)
	if tlDiff > 5 {
		tlDiff = 5
	}
	dm -= tlDiff
	return dm
}

func AvailableFreightLots(tfv int) (int, int, int) {
	inLots := utils.Max(0, dice.Roll("1d6").DM(tfv-13).Sum())
	mnLots := utils.Max(0, dice.Roll("1d6").DM(tfv-8).Sum())
	mjLots := utils.Max(0, dice.Roll("1d6").DM(tfv-6).Sum())
	fmt.Println(inLots, mnLots, mjLots)
	return inLots, mnLots, mjLots
}

type freightLot struct {
	tonns     int
	basePrice int
	source    string
	target    string
	dist      int
}

type lot interface {
	Price() int
	//Negotiate()
}

func NewFreightLot(sourceHex, targetHex string, tons int) freightLot {
	frL := freightLot{}
	frL.tonns = tons
	frL.source = sourceHex
	frL.target = targetHex
	frL.dist = Astrogation.JumpDistance(frL.source, frL.target)
	frL.basePrice = 500
	return frL
}

func (l freightLot) Price() int {
	price := 500 * l.tonns * l.dist
	bonus := (price / 5) * l.dist
	return price + bonus
}

func LotList(inLot, mnLot, mjLot int) []lot {
	var lots []lot
	var tons []int
	for i := 0; i < mjLot; i++ {
		tons = append(tons, dice.Roll("1d6").Sum()*10)
	}
	for i := 0; i < mnLot; i++ {
		tons = append(tons, dice.Roll("1d6").Sum()*5)
	}
	for i := 0; i < inLot; i++ {
		tons = append(tons, dice.Roll("1d6").Sum()*1)
	}
	for i := range tons {
		lots = append(lots, NewFreightLot(sourceWorld.Hex(), targetWorld.Hex(), tons[i]))
	}
	return lots
}
