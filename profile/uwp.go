package profile

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/TrvCore/ehex"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

func RandomUWP(planetType ...string) string {
	var result string
	var pType string
	pType = constant.WTpHospitable
	if len(planetType) > 0 {
		pType = planetType[0]
		if !constant.WorldTypeValid(pType) {
			pType = constant.WTpHospitable
		}
	}
	if len(planetType) > 1 {
		utils.SetSeed(utils.SeedFromString(planetType[1]))
	}
	//fmt.Println("Set pType as:", pType)
	//////////SIZE
	var size int
	switch pType {
	default:
		size = rollStat(2, -2, 0)
		if size == 10 {
			size = rollStat(1, 9, 0)
		}
	case constant.WTpRadWorld, constant.WTpStormWorld:
		size = rollStat(2, 0, 0)
	case constant.WTpInferno:
		size = rollStat(1, 6, 0)
	case constant.WTpBigWorld:
		size = rollStat(2, 7, 0)
	case constant.WTpWorldlet:
		size = rollStat(1, -3, 0)
	case constant.WTpPlanetoid:
		size = 0
	case constant.WTpGG:
		size = rollStat(0, 26+flux(), 0)
	}
	size = utils.BoundInt(size, 0, 32)
	//uwp.data[constant.PrSize] = ehex(size)
	result = TrvCore.DigitToEhex(size)

	//////////ATMO
	var atmo int
	switch pType {
	default:
		atmo = rollStat(0, size+flux(), 0)
	case constant.WTpPlanetoid:
		atmo = 0
	case constant.WTpStormWorld:
		atmo = rollStat(0, size+flux(), 4)
	case constant.WTpInferno:
		atmo = TrvCore.EhexToDigit("B")
	}
	if size == 0 {
		atmo = 0
	}
	atmo = utils.BoundInt(atmo, 0, TrvCore.EhexToDigit("F"))
	//uwp.data[constant.PrAtmo] = ehex(atmo)
	result = result + TrvCore.DigitToEhex(atmo)

	//////////HYDRO
	var hydr int
	dm := 0
	if atmo < 2 || atmo > 9 {
		dm = -4
	}
	switch pType {
	default:
		hydr = rollStat(0, atmo+flux(), dm)
	case constant.WTpPlanetoid, constant.WTpInferno:
		hydr = 0
	case constant.WTpStormWorld, constant.WTpInnerWorld:
		hydr = rollStat(0, atmo+flux(), dm-4)
	}
	if size < 2 {
		hydr = 0
	}
	hydr = utils.BoundInt(hydr, 0, TrvCore.EhexToDigit("A"))
	//uwp.data[constant.PrHydr] = ehex(hydr)
	result = result + TrvCore.DigitToEhex(hydr)

	//////////POPS
	var pops int
	dm = 0
	switch pType {
	default:
		pops = rollStat(2, -2, dm)
		if pops == 10 {
			pops = rollStat(2, 3, dm)
		}
	case constant.WTpRadWorld, constant.WTpInferno, constant.WTpGG:
		pops = 0
	case constant.WTpIceWorld, constant.WTpStormWorld:
		pops = rollStat(2, -2, -6)
	case constant.WTpInnerWorld:
		pops = rollStat(2, -2, -4)
	}

	pops = utils.BoundInt(pops, 0, TrvCore.EhexToDigit("Y"))
	//uwp.data[constant.PrPops] = ehex(pops)
	result = result + TrvCore.DigitToEhex(pops)

	//////////GOVR
	var govr int
	switch pType {
	default:
		govr = rollStat(0, pops+flux(), 0)
	case constant.WTpRadWorld, constant.WTpInferno:
		govr = 0
	}
	if pops == 0 {
		govr = 0
	}
	govr = utils.BoundInt(govr, 0, TrvCore.EhexToDigit("F"))
	//uwp.data[constant.PrGovr] = ehex(govr)
	result = result + TrvCore.DigitToEhex(govr)

	//////////LAWS
	var laws int
	switch pType {
	default:
		laws = rollStat(0, govr+flux(), 0)
	}
	if pops == 0 {
		laws = 0
	}
	laws = utils.BoundInt(laws, 0, TrvCore.EhexToDigit("J"))
	//uwp.data[constant.PrLaws] = ehex(laws)
	result = result + TrvCore.DigitToEhex(laws)

	//////////Starport
	var stprt string
	switch pType {
	default:
		st := pops - rollStat(1, 0, 0)
		switch st {
		default:
			if st > 3 {
				stprt = "F"
			}
			if st < 1 {
				stprt = "Y"
			}
		case 3:
			stprt = "G"
		case 1, 2:
			stprt = "H"
		}
	case constant.WTpHospitable:
		dm = 0
		if pops > 7 {
			dm = dm + 1
		}
		if pops > 9 {
			dm = dm + 2
		}
		if pops < 5 {
			dm = dm - 1
		}
		if pops > 3 {
			dm = dm - 2
		}
		r := utils.RollDice("2d6", dm)
		switch r {
		default:
			if r < 3 {
				stprt = "X"
			}
			if r > 10 {
				stprt = "A"
			}
		case 3, 4:
			stprt = "E"
		case 5, 6:
			stprt = "D"
		case 7, 8:
			stprt = "C"
		case 9, 10:
			stprt = "B"
		}
	case constant.WTpInferno, constant.WTpGG:
		stprt = "Y"
	}
	//uwp.data[constant.PrStarport] = stprt
	result = stprt + result

	////////////////////TL
	dm = 0
	var tl int
	switch pType {
	default:
		switch stprt {
		case "A":
			dm += 6
		case "B":
			dm += 4
		case "C":
			dm += 2
		case "E": //Домысливание
			dm -= 2 //Домысливание
		case "X":
			dm -= 4
		case "F":
			dm += 1
		}
		switch size {
		case 0, 1:
			dm += 2
		case 2, 3, 4:
			dm += 1
		}
		switch atmo {
		case 0, 1, 2, 3:
			dm += 1
		case 10, 11, 12, 13, 14, 15:
			dm += 1
		}
		switch hydr {
		case 9:
			dm += 1
		case 10:
			dm += 2
		}
		switch pops {
		case 1, 2, 3, 4, 5:
			dm += 1
		case 9:
			dm += 2
		default:
			if pops > 9 {
				dm += 4
			}
		}
		switch govr {
		case 0, 5:
			dm += 1
		case 13:
			dm -= 2
		}
		tl = rollStat(1, 0, dm)
	case constant.WTpRadWorld, constant.WTpInferno, constant.WTpGG:
		tl = 0
	}

	tl = utils.BoundInt(tl, 0, TrvCore.EhexToDigit("Y"))
	//uwp.data[constant.PrTL] = ehex(tl)
	result = result + "-" + TrvCore.DigitToEhex(tl)
	return result
}

func rollStat(die, mod, dm int) int {
	d := strconv.Itoa(die)
	//r := utils.RollDice(d+"d6", mod+dm)
	r := dice.Roll(d + "d6").DM(mod + dm).Sum()
	return r
}

func flux() int {
	return TrvCore.Flux()
}

// func ehex(i int) string {
// 	return TrvCore.DigitToEhex(i)
// }

func orderByType(profileType string) (order []string) {
	switch profileType {
	case "UWP":
		order = []string{
			constant.PrStarport,
			constant.PrSize,
			constant.PrAtmo,
			constant.PrHydr,
			constant.PrPops,
			constant.PrGovr,
			constant.PrLaws,
			constant.DIVIDER,
			constant.PrTL,
		}
	}
	return order
}

// func From(pu Puller) Profile {
// 	p := Profile{}
// 	switch pu.(type) {
// 	case world.World:
// 		data := pu.PullData()
// 		p.pType = "UWP"
// 		p.data = data[constant.PrStarport] +
// 			data[constant.PrSize] + data[constant.PrAtmo] + data[constant.PrHydr] +
// 			data[constant.PrPops] + data[constant.PrGovr] + data[constant.PrLaws] +
// 			"-" + data[constant.PrTL]
// 		return p
// 	}
// 	return p
// }

func CalculateTradeCodes(uwp string) []string {
	tradeCodes := constant.TravelCodesMgT2()
	var res []string
	for _, tc := range tradeCodes {
		switch tc {
		default:
		case constant.TradeCodeAgricultural:
			if matchTradeClassificationRequirements(uwp, "-- 456789 45678 567 -- -- --") {
				res = append(res, constant.TradeCodeAgricultural)
			}
		case constant.TradeCodeAsteroid:
			if matchTradeClassificationRequirements(uwp, "0 0 0 -- -- -- --") {
				res = append(res, constant.TradeCodeAsteroid)
			}
		case constant.TradeCodeBarren:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 0 0 0") {
				res = append(res, constant.TradeCodeBarren)
			}
		case constant.TradeCodeDesert:
			if matchTradeClassificationRequirements(uwp, "-- 23456789ABCDEFS 0 -- -- -- --") {
				res = append(res, constant.TradeCodeDesert)
			}
		case constant.TradeCodeFluidOceans:
			if matchTradeClassificationRequirements(uwp, "-- ABCDEF 123456789A -- -- -- --") {
				res = append(res, constant.TradeCodeFluidOceans)
			}
		case constant.TradeCodeGarden:
			if matchTradeClassificationRequirements(uwp, "678 568 567 -- -- -- --") {
				res = append(res, constant.TradeCodeGarden)
			}
		case constant.TradeCodeHighPopulation:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 9ABCDEF -- -- --") {
				res = append(res, constant.TradeCodeHighPopulation)
			}
		case constant.TradeCodeHighTech:
			if matchTradeClassificationRequirements(uwp, "-- -- -- -- -- -- CDEFGH") {
				res = append(res, constant.TradeCodeHighTech)
			}
		case constant.TradeCodeIceCapped:
			if matchTradeClassificationRequirements(uwp, "-- 01 123456789A -- -- -- --") {
				res = append(res, constant.TradeCodeIceCapped)
			}
		case constant.TradeCodeIndustrial:
			if matchTradeClassificationRequirements(uwp, "-- 012479 -- 9ABCDEF -- -- --") {
				res = append(res, constant.TradeCodeIndustrial)
			}
		case constant.TradeCodeLowPopulation:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 123 -- -- --") {
				res = append(res, constant.TradeCodeLowPopulation)
			}
		case constant.TradeCodeLowTech:
			if matchTradeClassificationRequirements(uwp, "-- -- -- -- -- -- 12345") {
				res = append(res, constant.TradeCodeLowTech)
			}
		case constant.TradeCodeNonAgricultural:
			if matchTradeClassificationRequirements(uwp, "-- 0123 0123 6789ABCDEF -- -- --") {
				res = append(res, constant.TradeCodeNonAgricultural)
			}
		case constant.TradeCodeNonIndustrial:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 123456 -- -- --") {
				res = append(res, constant.TradeCodeNonIndustrial)
			}
		case constant.TradeCodePoor:
			if matchTradeClassificationRequirements(uwp, "-- 2345 0123 -- -- -- --") {
				res = append(res, constant.TradeCodePoor)
			}
		case constant.TradeCodeRich:
			if matchTradeClassificationRequirements(uwp, "-- 68 -- 678 456789 -- --") {
				res = append(res, constant.TradeCodeRich)
			}
		case constant.TradeCodeVacuum:
			if matchTradeClassificationRequirements(uwp, "-- 0 -- -- -- -- --") {
				res = append(res, constant.TradeCodeVacuum)
			}
		case constant.TradeCodeWaterWorld:
			if matchTradeClassificationRequirements(uwp, "-- -- A -- -- -- --") {
				res = append(res, constant.TradeCodeWaterWorld)
			}

		}
	}
	return res
}

func matchTradeClassificationRequirements(uwp, reqLine string) bool {
	stats := strings.Split(reqLine, " ") //-- 23456789 0 -- -- --
	ehexList := TrvCore.ValidEhexs()
	fullArray := ""
	for i := range ehexList {
		fullArray = fullArray + ehexList[i]
	}
	statArray := []string{string([]byte(uwp)[1]), string([]byte(uwp)[2]), string([]byte(uwp)[3]), string([]byte(uwp)[4]), string([]byte(uwp)[5]), string([]byte(uwp)[6]), string([]byte(uwp)[8])}
	for i := range stats { //собираем аррэй
		array := stats[i]
		if array == "--" {
			array = fullArray
		}
		if !strings.Contains(array, statArray[i]) {
			return false
		}
	}
	return true
}

// func Goverment(uwp string) string {
// 	return string([]byte(uwp)[5])
// }

// type UWP struct {
// 	data string
// }

// func NewUWP(s string) (UWP, error) {
// 	uwp := UWP{}
// 	if !uwpValid(s) {
// 		return uwp, errors.New("NewUWP: can't parse UWP from string")
// 	}
// 	uwp.data = s
// 	return uwp, nil
// }

// func uwpValid(uwp string) bool {
// 	if len(uwp) != 9 {
// 		return false
// 	}
// 	data := strings.Split(uwp, "")
// 	if data[7] != constant.DIVIDER {
// 		return false
// 	}
// 	for i := range data {
// 		if i == 7 {
// 			continue
// 		}
// 		if data[i] == "_" {
// 			continue
// 		}
// 		if TrvCore.EhexToDigit(data[i]) == -999 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (uwp UWP) Starport() string {
// 	return string([]byte(uwp.data)[0])
// }

// func (uwp UWP) Size() string {
// 	return string([]byte(uwp.data)[1])
// }
// func (uwp UWP) Atmo() string {
// 	return string([]byte(uwp.data)[2])
// }
// func (uwp UWP) Hydr() string {
// 	return string([]byte(uwp.data)[3])
// }
// func (uwp UWP) Pops() string {
// 	return string([]byte(uwp.data)[4])
// }
// func (uwp UWP) Govr() string {
// 	return string([]byte(uwp.data)[5])
// }
// func (uwp UWP) Laws() string {
// 	return string([]byte(uwp.data)[6])
// }

// func (uwp UWP) TL() string {
// 	return string([]byte(uwp.data)[8])
// }

type newUWPr struct {
	data map[string]ehex.DataRetriver
}

func (u *newUWPr) String() string {
	return u.Starport().String() + u.Size().String() + u.Atmo().String() + u.Hydr().String() + u.Pops().String() + u.Govr().String() + u.Laws().String() + "-" + u.TL().String()
}

func NewUWP(str string) *newUWPr {
	u := newUWPr{}
	u.data = make(map[string]ehex.DataRetriver)
	u.data[constant.PrStarport] = ehex.New(str[0])
	u.data[constant.PrSize] = ehex.New(str[1])
	u.data[constant.PrAtmo] = ehex.New(str[2])
	u.data[constant.PrHydr] = ehex.New(str[3])
	u.data[constant.PrPops] = ehex.New(str[4])
	u.data[constant.PrGovr] = ehex.New(str[5])
	u.data[constant.PrLaws] = ehex.New(str[6])
	u.data[constant.PrTL] = ehex.New(str[8])
	return &u
}

func (u *newUWPr) Starport() ehex.DataRetriver {
	return u.data[constant.PrStarport]
}

func (u *newUWPr) Size() ehex.DataRetriver {
	return u.data[constant.PrSize]
}
func (u *newUWPr) Atmo() ehex.DataRetriver {
	return u.data[constant.PrAtmo]
}
func (u *newUWPr) Hydr() ehex.DataRetriver {
	return u.data[constant.PrHydr]
}
func (u *newUWPr) Pops() ehex.DataRetriver {
	return u.data[constant.PrPops]
}
func (u *newUWPr) Govr() ehex.DataRetriver {
	return u.data[constant.PrGovr]
}
func (u *newUWPr) Laws() ehex.DataRetriver {
	return u.data[constant.PrLaws]
}

func (u *newUWPr) TL() ehex.DataRetriver {
	return u.data[constant.PrTL]
}
