package law

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/utils"
)

const (
	lawOverall     = "Overall"
	lawWeapon      = "Weapons"
	lawDrugs       = "Drugs"
	lawInformation = "Information"
	lawTechnology  = "Technology"
	lawTravellers  = "Travellers"
	lawPsionics    = "Psionics"
	goverment      = "Goverment"
)

type LawReport struct {
	dp         *dice.Dicepool
	ulp        string //L5-444444 [Wp Dr In Te Tr Ps]
	levelOf    map[string]int
	contraband [6]int
}

func New(uwpStr string) (LawReport, error) {
	lr := LawReport{}
	uwp, err := profile.NewUWP(uwpStr)
	if err != nil {
		return lr, err
	}
	lr.dp = dice.New(utils.SeedFromString(uwpStr))
	lr.levelOf = make(map[string]int)
	//lr.contraband = [6]int{2, 2, 2, 2, 2, 2}
	lr.levelOf[lawOverall] = TrvCore.EhexToDigit(uwp.Laws())
	lr.levelOf[goverment] = TrvCore.EhexToDigit(uwp.Govr())
	lr.levelOf[lawWeapon] = utils.BoundInt(lr.levelOf[lawOverall]+TrvCore.Flux(), 0, 30)
	lr.levelOf[lawDrugs] = utils.BoundInt(lr.levelOf[lawOverall]+TrvCore.Flux(), 0, 30)
	lr.levelOf[lawInformation] = utils.BoundInt(lr.levelOf[lawOverall]+TrvCore.Flux(), 0, 30)
	lr.levelOf[lawTechnology] = utils.BoundInt(lr.levelOf[lawOverall]+TrvCore.Flux(), 0, 30)
	lr.levelOf[lawTravellers] = utils.BoundInt(lr.levelOf[lawOverall]+TrvCore.Flux(), 0, 30)
	lr.levelOf[lawPsionics] = utils.BoundInt(lr.levelOf[lawOverall]+TrvCore.Flux(), 0, 30)
	lr.determineContraband()

	return lr, nil
}

func (lr *LawReport) determineContraband() {
	contrSl := [6]int{}
	switch lr.levelOf[goverment] {
	case 0:
		contrSl = [6]int{0, 0, 0, 0, 0, 0}
	case 1:
		contrSl = [6]int{1, 1, 0, 0, 1, 0}
	case 2:
		contrSl = [6]int{0, 1, 0, 0, 0, 0}
	case 3:
		contrSl = [6]int{1, 0, 0, 1, 1, 0}
	case 4:
		contrSl = [6]int{1, 1, 0, 0, 0, 1}
	case 5:
		contrSl = [6]int{1, 0, 1, 1, 0, 0}
	case 6:
		contrSl = [6]int{1, 0, 0, 1, 1, 0}
	case 7:
		contrSl = [6]int{2, 2, 2, 2, 2, 2}
	case 8:
		contrSl = [6]int{1, 1, 0, 0, 0, 0}
	case 9:
		contrSl = [6]int{1, 1, 0, 1, 1, 1}
	case 10:
		contrSl = [6]int{0, 0, 0, 0, 0, 0}
	case 11:
		contrSl = [6]int{1, 0, 1, 1, 0, 0}
	case 12:
		contrSl = [6]int{1, 0, 0, 0, 0, 0}
	default:
		contrSl = [6]int{2, 2, 2, 2, 2, 2}
	}
	for i := range contrSl {
		if contrSl[i] != 0 && contrSl[i] != 1 {
			r := lr.dp.RollNext("1d2").DM(-1).Sum()
			contrSl[i] = r
		}
	}
	lr.contraband = contrSl
}

func (lr LawReport) ULP() string {
	ulp := "L"
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawOverall]) + "-"
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawWeapon])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawDrugs])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawInformation])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawTechnology])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawTravellers])
	ulp += TrvCore.DigitToEhex(lr.levelOf[lawPsionics]) + " "
	//L5-444444 [Wp Dr In Te Tr Ps]
	if lr.contraband[0] == 1 {
		ulp += "Wp "
	}
	if lr.contraband[1] == 1 {
		ulp += "Dr "
	}
	if lr.contraband[2] == 1 {
		ulp += "In "
	}
	if lr.contraband[3] == 1 {
		ulp += "Te "
	}
	if lr.contraband[4] == 1 {
		ulp += "Tr "
	}
	if lr.contraband[5] == 1 {
		ulp += "Ps "
	}
	ulp = strings.TrimSuffix(ulp, " ")
	return ulp
}

func (lr LawReport) Report() string {
	str := "" + lr.ULP()
	return str
}
