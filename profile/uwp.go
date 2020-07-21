package profile

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

type Profiled interface {
	ValueOf(string) string
	SetValue(string, string)
}

type Profile struct {
	pType string
	data  string
}

func NewProfile(pType, data string) Profile {
	pr := Profile{}
	pr.pType = pType
	pr.data = data
	return pr
}

func (pr *Profile) order() []string {
	return orderByType(pr.pType)
}

type Puller interface {
	PullData() map[string]string
}

type Pusher interface {
	SetValue(string, string)
}

type PushPuller interface {
	Pusher
	Puller
}

func Compose(prfType string, p Puller) string {
	var profile string
	order := orderByType(prfType)
	if len(order) < 1 {
		return "Error: Unknown profile Type"
	}
	data := p.PullData()
	for _, val := range order {
		if ok, caseResolution := specialCase(val); ok {
			profile = profile + caseResolution
			continue
		}
		if _, ok := data[val]; !ok {
			profile = profile + constant.UNKNOWN
			//return "Error: DATA UNAVAILABLE(" + val + ")"
		}
		profile = profile + data[val]
	}
	return profile
}

func AssignTo(pp PushPuller, prf Profile) {

}

func specialCase(val string) (bool, string) {
	var output string
	switch val {
	default:
		return false, output
	case constant.DIVIDER:
		return true, "-"
	}
}

// type UWP struct {
// 	data map[string]string
// }

// func (uwp UWP) String() string {
// 	return uwp.data[constant.PrStarport] + uwp.data[constant.PrSize] + uwp.data[constant.PrAtmo] + uwp.data[constant.PrHydr] +
// 		uwp.data[constant.PrPops] + uwp.data[constant.PrGovr] + uwp.data[constant.PrLaws] + "-" + uwp.data[constant.PrTL]
// }

// func NewUWP(uwpStr string) (UWP, error) {
// 	var uwp UWP
// 	uwp.data = make(map[string]string)
// 	if len(uwpStr) != 9 {
// 		return uwp, errors.New("Invalid input size '" + uwpStr + "'")
// 	}
// 	code := strings.Split(uwpStr, "")

// 	uwp.data[constant.PrStarport] = code[0]
// 	uwp.data[constant.PrSize] = code[1]
// 	uwp.data[constant.PrAtmo] = code[2]
// 	uwp.data[constant.PrHydr] = code[3]
// 	uwp.data[constant.PrPops] = code[4]
// 	uwp.data[constant.PrGovr] = code[5]
// 	uwp.data[constant.PrLaws] = code[6]
// 	uwp.data[constant.PrTL] = code[8]
// 	return uwp, nil
// }

// func UWPFrom(container Profiled) UWP {
// 	order := orderByType("UWP")
// 	uwp := UWP{}
// 	uwp.data = make(map[string]string)
// 	for _, key := range order {
// 		uwp.data[key] = container.ValueOf(key)
// 	}
// 	return uwp
// }

// func PushTo(container Profiled, uwp *UWP) {
// 	for k, v := range uwp.data {
// 		if k == constant.DIVIDER {
// 			continue
// 		}
// 		container.SetValue(k, v)
// 	}
// }

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
	result = ehex(size)

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
	result = result + ehex(atmo)

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
	result = result + ehex(hydr)

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
	result = result + ehex(pops)

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
	result = result + ehex(govr)

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
	result = result + ehex(laws)

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
	result = result + "-" + ehex(tl)
	return result
}

func rollStat(dice, mod, dm int) int {
	d := strconv.Itoa(dice)
	r := utils.RollDice(d+"d6", mod+dm)
	return r
}

func flux() int {
	return TrvCore.Flux()
}

func ehex(i int) string {
	return TrvCore.DigitToEhex(i)
}

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

func From(pu Puller) Profile {
	p := Profile{}
	switch pu.(type) {
	case world.World:
		data := pu.PullData()
		p.pType = "UWP"
		p.data = data[constant.PrStarport] +
			data[constant.PrSize] + data[constant.PrAtmo] + data[constant.PrHydr] +
			data[constant.PrPops] + data[constant.PrGovr] + data[constant.PrLaws] +
			"-" + data[constant.PrTL]
		return p
	}
	return p
}
