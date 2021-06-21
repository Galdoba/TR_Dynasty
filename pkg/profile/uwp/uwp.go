package uwp

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/pkg/core/astronomical"
	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

//NewGasGigant -
func NewGasGigant(dp *dice.Dicepool, ggType string) string {
	size := ""
	switch ggType {
	default:
		size = "0"
	case "LGG":
		for size == "" {
			r := dp.RollNext("2d6").DM(1).Sum()
			switch r {
			case 4:
				size = "P"
			case 5:
				size = "Q"
			case 6, 7:
				size = "Z"
			// case 7:
			// 	size = "S"
			case 8:
				size = "T"
			case 9:
				size = "U"
			case 10:
				size = "V"
			case 11:
				size = "W"
			case 12:
				size = "X"
			case 13:
				size = "Y"
			}
		}
	case "SGG", "IG":
		for size == "" {
			r := dp.RollNext("2d6").DM(-1).Sum()
			switch r {
			case 1:
				size = "L"
			case 2:
				size = "M"
			case 3:
				size = "N"
			}
		}
	}
	return "Y" + size + "XXXXX-X"
}

//GenerateOtherWorldUWP -
func GenerateOtherWorldUWP(dices *dice.Dicepool, mwUWP string, planetType string, star string, orbit int) string { //planetType string - используем ли?
	if planetType == constant.WTpPlanetoid {
		return "BELT     "
	}
	//SIZE
	sizeDM := sizeDM(star, orbit)
	size := rollSize(dices, planetType, sizeDM)
	//ATMO
	atmoDM := atmoDM(size, orbit, star)
	atmo := rollAtmo(dices, planetType, atmoDM, size)
	if astronomical.HabitableZoneScore(orbit, star) > 2 {
		if dices.RollNext("2d6").Sum() == 12 {
			atmo = 10
		}
	}
	//HYDR
	hydrDM := hydrDM(orbit, star, atmo)
	hydr := rollHydr(dices, size, atmo, orbit, star, hydrDM)
	//POPS
	popsDM := popsDM(star, orbit, atmo, size)
	pops := rollPops(dices, popsDM)
	mwPop := New(mwUWP).Pops().Value()
	pops = utils.BoundInt(pops, 0, mwPop-1)
	//GOVR
	govr := rollGovr(dices, pops, mwUWP)
	//LAWS
	laws := rollLaws(dices, mwUWP)
	if pops == 0 {
		laws = 0
	}
	//STARPORT
	stpt := rollStpt(dices, mwPop)
	//tl := rollTL(mwUWP)
	tl := rollTLt5(mwUWP, dices, stpt, size, atmo, hydr, pops, govr)
	if tl > ehex.New(mwUWP[8]).Value()-1 {
		tl = ehex.New(mwUWP[8]).Value() - 1
	}
	if pops == 0 {
		tl = 0
		govr = 0
		laws = 0
	}
	sizeGlyph := ehex.New(size).String()
	if planetType != constant.WTpGG && planetType != constant.WTpPlanetoid && sizeGlyph == "0" {
		sizeGlyph = "S"
	}
	uwp := stpt + sizeGlyph + ehex.New(atmo).String() + ehex.New(hydr).String() + ehex.New(pops).String() + ehex.New(govr).String() + ehex.New(laws).String() + "-" + ehex.New(tl).String()
	return uwp
}

func sizeDM(star string, orbit int) int {
	sizeDM := 0
	switch orbit {
	case 0:
		sizeDM += -5
	case 1:
		sizeDM += -4
	case 2:
		sizeDM += -2
	}
	if strings.Contains(star, "M") {
		sizeDM += -2
	}
	return sizeDM
}

func rollSize(dp *dice.Dicepool, planetType string, sizeDM int) int {
	sz := 0
	switch planetType {
	default:
		sz = dp.RollNext("2d6").DM(sizeDM - 2).Sum()
	case constant.WTpBigWorld:
		sz = dp.RollNext("2d6").DM(sizeDM + 7).Sum()
	case constant.WTpPlanetoid:
		sz = 0
	case constant.WTpRadWorld, constant.WTpStormWorld:
		sz = dp.RollNext("2d6").DM(sizeDM).Sum()
	case constant.WTpInferno:
		sz = dp.RollNext("1d6").DM(sizeDM + 6).Sum()
	case constant.WTpWorldlet:
		sz = dp.RollNext("1d6").DM(sizeDM - 3).Sum()
	}
	if sz < 0 {
		sz = 0
	}
	return sz
}

func atmoDM(size int, orbit int, star string) int {
	atmoDM := 0
	switch astronomical.Zone(orbit, star) {
	case astronomical.ZoneInner:
		atmoDM += -2
	case astronomical.ZoneOuter:
		atmoDM += -4
	}
	return atmoDM
}

func rollAtmo(dp *dice.Dicepool, planetType string, atmoDM int, size int) int {
	atmo := 0
	switch planetType {
	default:
		atmo = dp.RollNext("2d6").DM(size - 7 + atmoDM).Sum()
	case constant.WTpPlanetoid:
		atmo = 0
	}
	if atmo < 0 {
		atmo = 0
	}
	if size == 0 {
		atmo = 0
	}
	return atmo
}

func hydrDM(orbit int, star string, atmo int) int {
	hydrDM := 0
	switch astronomical.Zone(orbit, star) {
	case astronomical.ZoneOuter:
		hydrDM += -2
	case astronomical.ZoneInner:
		hydrDM += -4
	}
	if atmo < 2 {
		hydrDM += -4
	}
	if atmo > 9 {
		hydrDM += -4
	}
	return hydrDM
}

func rollHydr(dp *dice.Dicepool, size int, atmo int, orbit int, star string, hydrDM int) int {
	hydr := dp.RollNext("2d6").DM(size - 7 + hydrDM).Sum()
	if size < 2 {
		hydr = 0
	}
	if astronomical.Zone(orbit, star) == astronomical.ZoneInner {
		hydr = 0
	}
	hydr = utils.BoundInt(hydr, 0, 10)
	return hydr
}

func popsDM(star string, orbit int, atmo int, size int) int {
	dm := 0
	switch astronomical.Zone(orbit, star) {
	case astronomical.ZoneInner:
		dm += -5
	case astronomical.ZoneOuter:
		dm += -3
	}
	switch atmo {
	default:
		dm += -2
	case 0, 5, 6, 8:
	}
	if size < 5 || size > 9 {
		dm += -2
	}
	return dm
}

func govrDM(mwUWP string) int {
	dm := 0
	mwuwp := New(mwUWP)
	mwGovr := mwuwp.Govr().Value()
	switch mwGovr {
	case 6:
		dm += mwuwp.Pops().Value()
	default:
		if mwGovr > 6 {
			dm++
		}
	}
	return dm
}

func rollLaws(dp *dice.Dicepool, mwUWP string) int {
	mwLaw := New(mwUWP).Laws().Value()
	mwPop := New(mwUWP).Pops().Value()
	law := dp.RollNext("1d6").DM(-3 + mwLaw).Sum()
	if law < 0 {
		law = 0
	}
	if mwPop == 0 {
		law = 0
	}
	return law
}

func rollGovr(dp *dice.Dicepool, pops int, mwUWP string) int {
	mwGOV := ehex.New(mwUWP[5]).Value()
	dm := 0
	if mwGOV == 6 {
		dm = pops
	}
	if mwGOV >= 7 {
		dm = -1
	}
	// govr := dp.RollNext("1d6").DM(dm).Sum()
	// if pops == 0 {
	// 	govr = 0
	// }
	govr := 0
	r := dp.RollNext("1d6").DM(dm).Sum()
	switch r {
	case 1, 0:
		govr = 0
	case 2:
		govr = 1
	case 3:
		govr = 2
	case 4:
		govr = 3
	default:
		govr = 6
	}
	if pops == 0 {
		govr = 0
	}
	return govr
}

func rollPops(dp *dice.Dicepool, dm int) int {
	pop := dp.RollNext("2d6").DM(-2 + dm).Sum()
	if pop < 0 {
		pop = 0
	}
	return pop
}

func rollStpt(dp *dice.Dicepool, mwPop int) string {
	//dm := 0
	r := dp.RollNext("1d6").Sum()
	spIndex := mwPop - r
	stpt := ""
	switch spIndex {
	default:
		if spIndex >= 4 {
			stpt = "F"
		}
		if spIndex <= 0 {
			stpt = "Y"
		}
	case 1, 2:
		stpt = "H"
	case 3:
		stpt = "G"
	}
	// switch mwPop {
	// case 0:
	// 	dm += -3
	// case 1:
	// 	dm += -2
	// default:
	// 	if mwPop > 5 {
	// 		dm += 2
	// 	}
	// }
	// r := dp.RollNext("1d6").DM(dm).Sum()
	// r = utils.BoundInt(r, 1, 6)
	// stpt := ""
	// switch r {
	// case 1, 2:
	// 	stpt = "Y"
	// case 3:
	// 	stpt = "H"
	// case 4, 5:
	// 	stpt = "G"
	// case 6:
	// 	stpt = "F"
	// }
	return stpt
}

func rollTL(mwUWP string) int {
	mwuwp := New(mwUWP)
	mwTL := mwuwp.TL().Value()
	//mwAtmo := mwuwp.Atmo().Value()
	tl := mwTL - 1
	// if tl < 7 {
	// 	// switch mwAtmo {
	// 	// default:
	// 	// 	tl = 0
	// 	// case 5, 6, 8:
	// 	// }
	// 	tl = 0
	// }
	return tl
}

func rollTLt5(mwUWP string, dices *dice.Dicepool, stpt string, siz, atm, hyd, pop, gov int) int {
	r := dices.RollNext("1d6").Sum()
	tl := r
	switch siz {
	case 0, 1:
		tl = tl + 2
	case 2, 3, 4:
		tl++
	}
	switch atm {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		tl++
	}
	switch hyd {
	case 9:
		tl++
	case 10:
		tl = tl + 2
	}
	switch pop {
	case 1, 2, 3, 4, 5:
		tl++
	case 9:
		tl = tl + 2
	case 10:
		tl = tl + 4
	}
	switch gov {
	case 0, 5:
		tl++
	case 13:
		tl = tl - 2
	}
	switch stpt {
	case "A":
		tl = tl + 6
	case "B":
		tl = tl + 4
	case "C":
		tl = tl + 2
	case "X":
		tl = tl - 4
	case "F":
		tl++
	}
	mwuwp := New(mwUWP)
	mwTL := mwuwp.TL().Value()
	if tl >= mwTL {
		tl = mwTL - 1
	}
	if tl < 0 {
		tl = 0
	}
	return tl
}

//RandomUWP -
func RandomUWP(dicepool *dice.Dicepool, planetType ...string) string {
	var result string
	var pType string
	pType = constant.WTpHospitable
	if len(planetType) > 0 {
		pType = planetType[0]
		if !constant.WorldTypeValid(pType) {
			pType = constant.WTpHospitable
		}
	}
	mainworldUWP := UWP{}
	mainworldPops := 15
	//mainworldTL := 30
	if len(planetType) > 1 {
		mainworldUWP = *New(planetType[1])
		mainworldPops = mainworldUWP.Pops().Value()
		//	mainworldTL = mainworldUWP.TL().Value()
	}

	// if len(planetType) > 1 {
	// 	utils.SetSeed(utils.SeedFromString(planetType[1]))
	// }
	//fmt.Println("Set pType as:", pType)
	//////////SIZE
	var size int
	switch pType {
	default:
		size = rollStat(dicepool, 2, -2, 0)
		if size == 10 {
			size = rollStat(dicepool, 1, 9, 0)
		}
	case constant.WTpRadWorld, constant.WTpStormWorld:
		size = rollStat(dicepool, 2, 0, 0)
	case constant.WTpInferno:
		size = rollStat(dicepool, 1, 6, 0)
	case constant.WTpBigWorld:
		size = rollStat(dicepool, 2, 7, 0)
	case constant.WTpWorldlet:
		size = rollStat(dicepool, 1, -3, 0)
	case constant.WTpPlanetoid:
		size = 0
	case constant.WTpGG:
		size = rollStat(dicepool, 0, 26+dicepool.FluxNext(), 0)
	}
	size = utils.BoundInt(size, 0, 32)
	//uwp.data[constant.PrSize] = ehex(size)
	result = TrvCore.DigitToEhex(size)

	//////////ATMO
	var atmo int
	switch pType {
	default:
		atmo = rollStat(dicepool, 0, size+dicepool.FluxNext(), 0)
	case constant.WTpPlanetoid:
		atmo = 0
	case constant.WTpStormWorld:
		atmo = rollStat(dicepool, 0, size+dicepool.FluxNext(), 4)
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
		hydr = rollStat(dicepool, 0, atmo+dicepool.FluxNext(), dm)
	case constant.WTpPlanetoid, constant.WTpInferno:
		hydr = 0
	case constant.WTpStormWorld, constant.WTpInnerWorld:
		hydr = rollStat(dicepool, 0, atmo+dicepool.FluxNext(), dm-4)
	}
	if size < 2 {
		hydr = 0
	}
	hydr = utils.BoundInt(hydr, 0, TrvCore.EhexToDigit("A"))
	//uwp.data[constant.PrHydr] = ehex(hydr)
	result = result + TrvCore.DigitToEhex(hydr)

	//////////POPS
	var pops int
	msp := 15
	dm = -3
	if &mainworldUWP != nil && mainworldUWP.TL().Value() < 9 {
		//dm = dm + (-1 * (9 - utils.BoundInt(mainworldUWP.TL().Value(), 0, 9)))
		msp = 10
		if mainworldUWP.TL().Value() < 9 {
			msp = 0
		}
		switch atmo {
		case 0, 1, 2, 3, 10, 11, 12:
			msp = 0
		case 5, 7, 9:
			msp = msp - 1
		case 4:
			msp = msp - 2
		case 13, 14, 15:
			msp = msp - 3
		}
		switch size {
		case 5, 6, 7:
			msp = msp - 1
		case 1, 2, 3, 4:
			msp = msp - 2
		}
		switch hydr {
		case 1, 2, 10:
			msp = msp - 1
		case 0:
			msp = msp - 2
		}

	}
	switch pType {
	default:
		pops = rollStat(dicepool, 2, -2, dm)
		if pops == 10 {
			pops = rollStat(dicepool, 2, 3, dm)
		}
	case constant.WTpRadWorld, constant.WTpInferno, constant.WTpGG:
		pops = 0
	case constant.WTpIceWorld, constant.WTpStormWorld:
		pops = rollStat(dicepool, 2, -2, -6+dm)
	case constant.WTpInnerWorld:
		pops = rollStat(dicepool, 2, -2, -4+dm)
	}
	msp = utils.BoundInt(msp, 0, 15)
	pops = utils.BoundInt(pops, 0, mainworldPops-1)
	pops = utils.BoundInt(pops, 0, msp)

	//uwp.data[constant.PrPops] = ehex(pops)
	result = result + TrvCore.DigitToEhex(pops)

	//////////GOVR
	var govr int
	switch pType {
	default:
		govr = rollStat(dicepool, 0, pops+dicepool.FluxNext(), 0)
		if &mainworldUWP != nil {
			govrRoll := dicepool.RollNext("1d6").Sum()
			if mainworldUWP.Govr().Value() == 6 {
				govrRoll = govrRoll + pops
			}
			if mainworldUWP.Govr().Value() > 6 {
				govrRoll--
			}
			govrRoll = utils.BoundInt(govrRoll, 1, 6)
			switch govrRoll {
			case 1:
				govr = 0
			case 2:
				govr = 1
			case 3:
				govr = 2
			case 4:
				govr = 3
			case 5, 6:
				govr = 6
			}
		}
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
		laws = rollStat(dicepool, 0, govr+dicepool.FluxNext(), 0)
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
		st := pops - rollStat(dicepool, 1, 0, 0)
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
		if pops < 3 {
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
			dm++
		}
		switch size {
		case 0, 1:
			dm += 2
		case 2, 3, 4:
			dm++
		}
		switch atmo {
		case 0, 1, 2, 3:
			dm++
		case 10, 11, 12, 13, 14, 15:
			dm++
		}
		switch hydr {
		case 9:
			dm++
		case 10:
			dm += 2
		}
		switch pops {
		case 1, 2, 3, 4, 5:
			dm++
		case 9:
			dm += 2
		default:
			if pops > 9 {
				dm += 4
			}
		}
		switch govr {
		case 0, 5:
			dm++
		case 13:
			dm -= 2
		case 14, 15:
			dm--
		}
		tl = rollStat(dicepool, 1, 0, dm)
	case constant.WTpRadWorld, constant.WTpInferno, constant.WTpGG:
		tl = 0
	}
	// if pops > 0 {
	// decline := 0
	// switch tl {
	// case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	// 	decline = dicepool.RollNext("1d6").DM(-3).Sum()
	// case 9, 10:
	// 	decline = dicepool.RollNext("1d6").Sum()
	// case 11, 12, 13:
	// 	decline = dicepool.RollNext("2d6").Sum()
	// default:
	// 	decline = dicepool.RollNext("3d6").Sum()
	// }
	// decline = utils.BoundInt(decline, 0, 18)
	// tl = tl - decline

	// }
	if pops == 0 && tl < 9 {
		tl = 0
	}
	//if stprt == "Y" ||
	if &mainworldUWP != nil {
		if mainworldUWP.TL().Value() <= tl {
			tl = 0
		}
	}

	tl = utils.BoundInt(tl, 0, TrvCore.EhexToDigit("Y"))

	//uwp.data[constant.PrTL] = ehex(tl)
	result = result + "-" + TrvCore.DigitToEhex(tl)
	return result
}

//RandomUWPShort -
func RandomUWPShort(dicepool *dice.Dicepool, planetType ...string) string {
	var result string
	var pType string
	pType = constant.WTpHospitable
	if len(planetType) > 0 {
		pType = planetType[0]
		if !constant.WorldTypeValid(pType) {
			pType = constant.WTpHospitable
		}
	}
	// if len(planetType) > 1 {
	// 	utils.SetSeed(utils.SeedFromString(planetType[1]))
	// }
	//fmt.Println("Set pType as:", pType)
	//////////SIZE
	var size int
	switch pType {
	default:
		size = rollStat(dicepool, 2, -2, 0)
		if size == 10 {
			size = rollStat(dicepool, 1, 9, 0)
		}
	case constant.WTpRadWorld, constant.WTpStormWorld:
		size = rollStat(dicepool, 2, 0, 0)
	case constant.WTpInferno:
		size = rollStat(dicepool, 1, 6, 0)
	case constant.WTpBigWorld:
		size = rollStat(dicepool, 2, 7, 0)
	case constant.WTpWorldlet:
		size = rollStat(dicepool, 1, -3, 0)
	case constant.WTpPlanetoid:
		size = 0
	case constant.WTpGG:
		size = rollStat(dicepool, 0, 26+dicepool.FluxNext(), 0)
	}
	size = utils.BoundInt(size, 0, 32)
	//uwp.data[constant.PrSize] = ehex(size)
	result = TrvCore.DigitToEhex(size)

	//////////ATMO
	var atmo int
	switch pType {
	default:
		atmo = rollStat(dicepool, 0, size+dicepool.FluxNext(), 0)
	case constant.WTpPlanetoid:
		atmo = 0
	case constant.WTpStormWorld:
		atmo = rollStat(dicepool, 0, size+dicepool.FluxNext(), 4)
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
		hydr = rollStat(dicepool, 0, atmo+dicepool.FluxNext(), dm)
	case constant.WTpPlanetoid, constant.WTpInferno:
		hydr = 0
	case constant.WTpStormWorld, constant.WTpInnerWorld:
		hydr = rollStat(dicepool, 0, atmo+dicepool.FluxNext(), dm-4)
	}
	if size < 2 {
		hydr = 0
	}
	hydr = utils.BoundInt(hydr, 0, TrvCore.EhexToDigit("A"))
	//uwp.data[constant.PrHydr] = ehex(hydr)
	result = result + TrvCore.DigitToEhex(hydr)

	result = result + TrvCore.DigitToEhex(0)
	result = result + TrvCore.DigitToEhex(0)
	result = result + TrvCore.DigitToEhex(0)

	//////////Starport
	var stprt string
	switch pType {
	default:
		stprt = "X"
	case constant.WTpInferno, constant.WTpGG:
		stprt = "Y"
	}
	//uwp.data[constant.PrStarport] = stprt
	result = stprt + result

	////////////////////TL
	tl := 0
	result = result + "-" + TrvCore.DigitToEhex(tl)
	return result
}

func rollStat(dp *dice.Dicepool, die, mod, dm int) int {
	d := strconv.Itoa(die)
	//r := utils.RollDice(d+"d6", mod+dm)
	r := dp.RollNext(d + "d6").DM(mod + dm).Sum()
	return r
}

//CalculateTradeCodes -
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

//CalculateTradeCodesT5 -
func CalculateTradeCodesT5(uwp string, mwTags []string, mw bool, hz int) []string {
	tradeCodes := constant.TravelCodesT5()
	var res []string
	for _, tc := range tradeCodes {
		switch tc {
		default:
		case constant.TradeCodeAsteroid:
			if matchTradeClassificationRequirements(uwp, "0 0 0 -- -- -- --") {
				res = append(res, constant.TradeCodeAsteroid)
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
		case constant.TradeCodeHellworld:
			if matchTradeClassificationRequirements(uwp, "3456789ABC 2479ABC 012 -- -- -- --") || hz <= -2 {
				res = append(res, constant.TradeCodeHellworld)
			}
		case constant.TradeCodeIceCapped:
			if matchTradeClassificationRequirements(uwp, "-- 01 123456789A -- -- -- --") {
				res = append(res, constant.TradeCodeIceCapped)
			}
		case constant.TradeCodeOceanWorld:
			if matchTradeClassificationRequirements(uwp, "ABCDEF -- A -- -- -- --") {
				res = append(res, constant.TradeCodeOceanWorld)
			}
		case constant.TradeCodeVacuum:
			if matchTradeClassificationRequirements(uwp, "-- 0 -- -- -- -- --") {
				res = append(res, constant.TradeCodeVacuum)
			}
		case constant.TradeCodeWaterWorld:
			if matchTradeClassificationRequirements(uwp, "-- -- A -- -- -- --") {
				res = append(res, constant.TradeCodeWaterWorld)
			}
			//Population
		case constant.TradeCodeDieback:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 0 0 0 123456789ABCDEFG") {
				res = append(res, constant.TradeCodeDieback)
			}
		case constant.TradeCodeBarren:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 0 0 0 0") {
				res = append(res, constant.TradeCodeBarren)
			}
		case constant.TradeCodeLowPopulation:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 123 -- -- --") {
				res = append(res, constant.TradeCodeLowPopulation)
			}
		case constant.TradeCodeNonIndustrial:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 456 -- -- --") {
				res = append(res, constant.TradeCodeNonIndustrial)
			}
		case constant.TradeCodePreHigh:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 8 -- -- --") {
				res = append(res, constant.TradeCodePreHigh)
			}
		case constant.TradeCodeHighPopulation:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 9ABCDEF -- -- --") {
				res = append(res, constant.TradeCodeHighPopulation)
			}
			//Secondary
		case constant.TradeCodeFarming:
			if matchTradeClassificationRequirements(uwp, "-- 456789 45678 23456 -- -- --") && mw == false {
				res = append(res, constant.TradeCodeFarming)
			}
			//
		case constant.TradeCodeMining:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 23456 -- -- --") && mw == false {
				for _, val := range mwTags {
					if val != "In" {
						continue
					}
					res = append(res, constant.TradeCodeMining)
				}
				//res = append(res, constant.TradeCodeAgricultural)
			}
		case constant.TradeCodeMilitaryRule:
			if matchTradeClassificationRequirements(uwp, "-- -- -- -- -- -- --") && mw == false {
				for _, val := range mwTags {
					if val != "Ph" && val != "Hi" {
						continue
					}
					dp := dice.New().SetSeed(uwp + uwp)
					dm := 0
					if val == "Hi" {
						dm++
					}
					if dp.RollNext("2d6").DM(dm).Sum() >= 12 {
						res = append(res, constant.TradeCodeMining)
					}
				}
				//				res = append(res, constant.TradeCodeMilitaryRule)
			}
		case constant.TradeCodePenalColony:
			if matchTradeClassificationRequirements(uwp, "-- 23AB 12345 3456 6 6789 --") && mw == false {
				res = append(res, constant.TradeCodePenalColony)
			}
		case constant.TradeCodeReserve:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 01234 6 045 --") && mw == false {
				res = append(res, constant.TradeCodeReserve)
			}
			//Political
		case constant.TradeCodeColony:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 56789A 6 0123 --") {
				res = append(res, constant.TradeCodeColony)
			}
			//Climate
		case constant.TradeCodeFrozen:
			if matchTradeClassificationRequirements(uwp, "23456789 -- 123456789A -- -- -- --") && hz > 1 {
				res = append(res, constant.TradeCodeFrozen)
			}
		case constant.TradeCodeHot:
			if matchTradeClassificationRequirements(uwp, "-- -- -- -- -- -- --") && hz == -1 {
				res = append(res, constant.TradeCodeHot)
			}
		case constant.TradeCodeCold:
			if matchTradeClassificationRequirements(uwp, "-- -- -- -- -- -- --") && hz == 1 {
				res = append(res, constant.TradeCodeCold)
			}
		case constant.TradeCodeTropic:
			if matchTradeClassificationRequirements(uwp, "6789 456789 34567 -- -- -- --") && hz == -1 {
				res = append(res, constant.TradeCodeTropic)
			}
		case constant.TradeCodeTundra:
			if matchTradeClassificationRequirements(uwp, "6789 456789 34567 -- -- -- --") && hz == 1 {
				res = append(res, constant.TradeCodeTundra)
			}

		//Economic
		case constant.TradeCodeHighTech:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 456789ABCDEF -- -- CDEFGH") {
				res = append(res, constant.TradeCodeHighTech)
			}
		case constant.TradeCodeLowTech:
			if matchTradeClassificationRequirements(uwp, "-- -- -- 456789ABCDEF -- -- 12345") {
				res = append(res, constant.TradeCodeLowTech)
			}
		case constant.TradeCodePreAgricultural:
			if matchTradeClassificationRequirements(uwp, "-- 456789 45678 48 -- -- --") {
				res = append(res, constant.TradeCodePreAgricultural)
			}
		case constant.TradeCodeAgricultural:
			if matchTradeClassificationRequirements(uwp, "-- 456789 45678 567 -- -- --") {
				res = append(res, constant.TradeCodeAgricultural)
			}
		case constant.TradeCodeNonAgricultural:
			if matchTradeClassificationRequirements(uwp, "-- 0123 0123 6789ABCDEF -- -- --") {
				res = append(res, constant.TradeCodeNonAgricultural)
			}
		case constant.TradeCodePrisonExileCamp:
			if matchTradeClassificationRequirements(uwp, "-- 23AB 12345 3456 -- 6789 --") && mw == true {
				res = append(res, constant.TradeCodePrisonExileCamp)
			}
		case constant.TradeCodePreIndustrial:
			if matchTradeClassificationRequirements(uwp, "-- 012479 -- 78 -- -- --") {
				res = append(res, constant.TradeCodePreIndustrial)
			}
		case constant.TradeCodeIndustrial:
			if matchTradeClassificationRequirements(uwp, "-- 012479ABC -- 9ABCDEF -- -- --") {
				res = append(res, constant.TradeCodeIndustrial)
			}
		case constant.TradeCodePoor:
			if matchTradeClassificationRequirements(uwp, "-- 2345 0123 -- -- -- --") {
				res = append(res, constant.TradeCodePoor)
			}
		case constant.TradeCodePreRich:
			if matchTradeClassificationRequirements(uwp, "-- 68 -- 59 -- -- --") {
				res = append(res, constant.TradeCodePreRich)
			}
		case constant.TradeCodeRich:
			if matchTradeClassificationRequirements(uwp, "-- 68 -- 678 -- -- --") {
				res = append(res, constant.TradeCodeRich)
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

type UWP struct {
	data map[string]ehex.DataRetriver
}

func (u *UWP) String() string {
	return u.Starport().String() + u.Size().String() + u.Atmo().String() + u.Hydr().String() + u.Pops().String() + u.Govr().String() + u.Laws().String() + "-" + u.TL().String()
}

//New -
func New(str string) *UWP {
	u := UWP{}
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

type UWPer interface {
	UWP() string
}

func From(w UWPer) *UWP {
	u := UWP{}
	u.data = make(map[string]ehex.DataRetriver)
	u.data[constant.PrStarport] = ehex.New(w.UWP()[0])
	u.data[constant.PrSize] = ehex.New(w.UWP()[1])
	u.data[constant.PrAtmo] = ehex.New(w.UWP()[2])
	u.data[constant.PrHydr] = ehex.New(w.UWP()[3])
	u.data[constant.PrPops] = ehex.New(w.UWP()[4])
	u.data[constant.PrGovr] = ehex.New(w.UWP()[5])
	u.data[constant.PrLaws] = ehex.New(w.UWP()[6])
	u.data[constant.PrTL] = ehex.New(w.UWP()[8])
	return &u
}

func (u *UWP) Starport() ehex.DataRetriver {
	return u.data[constant.PrStarport]
}

func (u *UWP) Size() ehex.DataRetriver {
	return u.data[constant.PrSize]
}

func (u *UWP) SizeValue() float64 {
	ehex := u.data[constant.PrSize]
	if ehex.String() == "S" {
		return 0.6
	}
	return float64(ehex.Value())
}

func (u *UWP) Atmo() ehex.DataRetriver {
	return u.data[constant.PrAtmo]
}
func (u *UWP) Hydr() ehex.DataRetriver {
	return u.data[constant.PrHydr]
}
func (u *UWP) Pops() ehex.DataRetriver {
	return u.data[constant.PrPops]
}
func (u *UWP) Govr() ehex.DataRetriver {
	return u.data[constant.PrGovr]
}
func (u *UWP) Laws() ehex.DataRetriver {
	return u.data[constant.PrLaws]
}

func (u *UWP) TL() ehex.DataRetriver {
	return u.data[constant.PrTL]
}
