package system

import (
	"errors"

	core "github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/utils"
)

type StarSystemEssc struct {
	name       string
	systemCode string
}

type star struct {
	position string
	sequense string
	size     string
	digit    string
	planets  int
}

func newStar(position string) *star {
	st := &star{}

	return st
}

func rollSequense(position string) string {
	dm := 0
	switch position {
	default:
		dm = 0
	case "Pc":
		dm = -1
	case "D":
		dm = -2
	case "Dc":
		dm = -3
	}
	switch core.Roll2D(dm) {
	default:
		return "BD"
	case 3, 4, 5, 6:
		return "M"
	case 7, 8:
		return "K"
	case 9, 10:
		return "G"
	case 11:
		return "F"
	case 12:
		return utils.RandomFromList([]string{"A", "O", "B"})
	}
}

func NewStarSystemEssc() *StarSystemEssc {
	ss := &StarSystemEssc{}
	// ss.systemCode += core.DigitToEhex(core.Roll2D())
	// ss.systemCode += core.DigitToEhex(core.Roll2D())
	// ss.systemCode += core.DigitToEhex(core.Roll2D())
	// syst, err := readGlyph(ss.systemCode, 0)
	// if err != nil {
	// 	return ss
	// }
	// if syst == "2" {
	// 	ss.systemCode += core.DigitToEhex(core.RerollIf(core.Roll2D, 2))
	// }
	//testSS()
	return ss
}

func readGlyph(code string, n int) (string, error) {
	codeBt := []byte(code)
	if n >= len(codeBt) {
		return "", errors.New("Can't read glyph from '" + code + "'. 'n' is too high")
	}
	if n < 0 {
		return "", errors.New("Can't read glyph from '" + code + "'. 'n' is too low")
	}
	return string(codeBt[n]), nil
}

func tableSystemType(index int) string {
	switch index {
	case 2:
		return "Special (roll on the Special System table)"
	case 3, 4:
		return "Trinary (close and distant companion)"
	case 5, 6:
		return "Binary (close companion)"
	case 7, 8:
		return "Solo star"
	case 9, 10:
		return "Binary (distant companion)"
	case 11:
		return "Trinary (distant companion with close companion of its own)"
	case 12:
		return "Multiple star system (four or more stellar bodies)"
	default:
		return ""
	}
}

func systemCode() string {
	code := ""
	code += core.DigitToEhex(core.Roll2D())
	if code != "2" {
		return code
	}
	alternative := core.RerollIf(core.Roll2D, 2)

	code = core.DigitToEhex(alternative) + core.DigitToEhex(core.Roll2D())
	return code
}

/*
3-4  PcD
5-6  Pc
7-8  P
9-10 PD
11   PDc

Star System
	Stellar Objects
		Planetary Bodies


*/

func AllNames(systName string) {

	testSS()
}

func testSS() {

}

func getPB(star string, posCode string) int {

	r := 0
	switch core.Roll2D() {
	default:
		r = 0
	case 2:
		r = 1
	case 3:
		r = utils.RollDice("d3")
	case 4, 5:
		r = utils.RollDice("d6", 1)
	case 6, 7, 8:
		r = utils.RollDice("2d6")
	case 9, 10:
		r = utils.RollDice("2d6", 3)
	case 11:
		r = utils.RollDice("3d6")
	case 12:
		r = utils.RollDice("4d6")
	}
	return r
}

func natureOfPlanet(hz int) string {
	if hz > -2 {
		return innerHZPlanet()
	}
	return outerHZPlanet()
}

func innerHZPlanet() string {
	return utils.RandomFromList([]string{
		constant.WTpInferno,
		constant.WTpInnerWorld,
		constant.WTpBigWorld,
		constant.WTpStormWorld,
		constant.WTpRadWorld,
		constant.WTpHospitable,
		constant.WTpPlanetoid,
	})
}

func outerHZPlanet() string {
	return utils.RandomFromList([]string{
		constant.WTpWorldlet,
		constant.WTpIceWorld,
		constant.WTpBigWorld,
		constant.WTpIceWorld,
		constant.WTpRadWorld,
		constant.WTpIceWorld,
		constant.WTpPlanetoid,
		constant.WTpPlanetoid,
		constant.WTpGG,
		constant.WTpGG,
		constant.WTpGG,
	})
}

func natureOfSatellite(hz int) string {
	if hz > -2 {
		return innerHZSatellite()
	}
	return outerHZSatellite()
}

func innerHZSatellite() string {
	return utils.RandomFromList([]string{
		constant.WTpInferno,
		constant.WTpInnerWorld,
		constant.WTpBigWorld,
		constant.WTpStormWorld,
		constant.WTpRadWorld,
		constant.WTpHospitable,
	})
}

func outerHZSatellite() string {
	return utils.RandomFromList([]string{
		constant.WTpWorldlet,
		constant.WTpIceWorld,
		constant.WTpBigWorld,
		constant.WTpStormWorld,
		constant.WTpRadWorld,
		constant.WTpIceWorld,
	})
}

func numberOfSatellites(wt string) int {
	switch wt {
	case "Inner":
		return utils.RollDice("d6", -5)
	case "Hospitable":
		return utils.RollDice("d6", -4)
	case "Outer":
		return utils.RollDice("d6", -3)
	}
	return -1
}

func natureOfBody() string {
	switch core.Roll2D() {
	case 2:
		return "Unusual " + natureOfBody()
	case 3, 4:
		return "Planetoid belt"
	case 5, 6, 7, 8:
		return "Terrestrial (rocky) planet"
	case 9, 10:
		return "Small gas giant"
	case 11:
		return "Large gas giant"
	case 12:
		return "Anomalous " + natureOfBody()
	default:
		return "Error"
	}
}

func checkArrayVal(arr []string, val string, pos int) bool {
	if len(arr) <= pos {
		return false
	}
	if arr[pos] != val {
		return false
	}
	return true
}

func arrayContainsAnyFrom(arr []string, subArr []string) bool {
	for i := range arr {
		for j := range subArr {
			if arr[i] == subArr[j] {
				return true
			}
		}
	}
	return false
}

func inList(str string, arr []string) bool {
	for i := range arr {
		if arr[i] == str {
			return true
		}
	}
	return false
}
