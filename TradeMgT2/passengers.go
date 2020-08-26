package trademgt2

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/devtools/cli/prettytable"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

var ptv int
var lp int
var mp int
var hp int

func PassengerProcedure() {

	ptv = passengerTrafficValue(sourceWorld, targetWorld)
	//fmt.Println("PTV =", ptv)
	//fmt.Println("Searching potential passengers...")
	//mailDR += userInputInt("Enter Streetwise(8) or Carouse(8) check effect: ")
	r := dice.Roll("2d6").DM(-8).Sum()
	//fmt.Println("Auto Roll:", r)
	lp, mp, hp = AvailablePassengers(r)
	informAboutPassengers()
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
	return dm
}

func AvailablePassengers(r int) (int, int, int) {
	return lowPass(ptv + r), midPass(ptv + r), hghPass(ptv + r)
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

type passenger struct {
	pType string
	cost  int
}

func newPassenger(ptype string) passenger {
	p := passenger{}
	p.pType = ptype
	switch p.pType {
	case lowPassenger:
		p.cost = lowPassCost()
	case midPassenger:
		p.cost = midPassCost()
	case hghPassenger:
		p.cost = hghPassCost()
	default:
		p.cost = -1
	}
	return p
}

func (p passenger) passFee() int {
	return p.cost
}

func (p passenger) passType() string {
	return p.pType
}

const (
	lowPassenger = "Low Passage"
	midPassenger = "Middle Passage"
	hghPassenger = "High Passage"
)

func lowPassCost() int {
	if dist < 1 {
		return 0
	}
	return 800 + dist*200
}

func midPassCost() int {
	if dist < 1 {
		return 0
	}
	if dist == 1 {
		return 3000
	}
	if dist == 2 {
		return 6000
	}
	return (dist - 1) * 5000
}

func hghPassCost() int {
	if dist < 1 {
		return 0
	}
	if dist == 1 {
		return 6000
	}
	if dist == 2 {
		return 12000
	}
	return (dist - 1) * 10000
}

func informAboutPassengers() {
	defer fmt.Println("------------------------------------------------------")
	tb := prettytable.New()
	tb.AddRow([]string{"Passenger", "Destination", "Available", "Projected Fee"})
	tb.AddRow([]string{hghPassenger, targetWorld.Name(), strconv.Itoa(hp), strconv.Itoa(newPassenger(hghPassenger).passFee()) + " Cr"})
	tb.AddRow([]string{midPassenger, targetWorld.Name(), strconv.Itoa(mp), strconv.Itoa(newPassenger(midPassenger).passFee()) + " Cr"})
	tb.AddRow([]string{lowPassenger, targetWorld.Name(), strconv.Itoa(lp), strconv.Itoa(newPassenger(lowPassenger).passFee()) + " Cr"})
	tb.PTPrint()
}
