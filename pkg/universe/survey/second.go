package survey

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/universe/starsystem"
	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey/calculations"
)

type SecondSurveyData struct {
	CoordX           int
	CoordY           int
	Sector           string
	Hex              string
	MW_Name          string
	MW_UWP           string
	MW_Remarks       string
	MW_Importance    string
	MW_ImportanceInt int
	MW_Economic      string
	MW_Cultural      string
	MW_Nobility      string
	Bases            string
	TravelZone       string
	PBG              string
	Worlds           int
	Allegiance       string
	Stellar          string
	RU               int
	input            string //temp
	errors           []error
}

func Parse(input string) *SecondSurveyData {
	ssd := SecondSurveyData{}
	//ssd.input = input
	data := strings.Split(input, "|")
	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}
	xCoord, errXcoord := strconv.Atoi(data[19])
	ssd.errors = append(ssd.errors, errXcoord)
	yCoord, errYcoord := strconv.Atoi(data[20])
	ssd.errors = append(ssd.errors, errYcoord)
	ssd.CoordX = xCoord
	ssd.CoordY = yCoord
	ssd.Sector = data[23]
	ssd.Hex = data[2]
	ssd.MW_Name = data[1]
	ssd.MW_UWP = data[3]
	ssd.MW_Remarks = data[21]
	ssd.MW_Importance = data[10]
	impInt, errImp := strconv.Atoi(data[11])
	ssd.errors = append(ssd.errors, errImp)
	ssd.MW_ImportanceInt = impInt
	ssd.MW_Economic = data[12]
	ssd.MW_Cultural = data[13]
	ssd.MW_Nobility = data[14]
	ssd.Bases = data[6]
	ssd.TravelZone = data[5]
	ssd.PBG = data[4]
	worlds, errWorlds := strconv.Atoi(data[15])
	ssd.errors = append(ssd.errors, errWorlds)
	ssd.Worlds = worlds
	ssd.Allegiance = data[5]
	ssd.Stellar = data[5]
	ru, errRu := strconv.Atoi(data[16])
	ssd.errors = append(ssd.errors, errRu)
	ssd.RU = ru
	ssd.verify()
	return &ssd
}

func (ssd *SecondSurveyData) containsErrors() bool {
	for _, val := range ssd.errors {
		if val != nil {
			return true
		}
	}
	return false
}

func (ssd *SecondSurveyData) verify() {
	if ssd.MW_Name == "" {
		ssd.MW_Name = ssd.NameByConvention()
	}
	if ssd.Stellar == "" {
		ssd.Stellar = starsystem.RollStellar(ssd.NameByConvention())
	}
	if !calculations.UWPvalid(ssd.MW_UWP) {
		ssd.MW_UWP = calculations.FixUWP(ssd.MW_UWP, ssd.NameByConvention())
	}
	if !calculations.PBGvalid(ssd.PBG, ssd.MW_UWP) {
		old := ssd.PBG
		ssd.PBG = calculations.FixPBG(ssd.PBG, ssd.MW_UWP, ssd.NameByConvention())
		new := ssd.PBG
		fmt.Printf("%v corrected to %v - %v\n", old, new, ssd.NameByConvention())
	}
	if ssd.MW_Importance == "{+?}" {
		ssd.MW_Importance = importanceToString(ssd.MW_ImportanceInt)
	}
	switch {
	default:
		return
	case ssd.MW_Name == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Mainworld name missing"))
	case ssd.Stellar == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Stellar data missing"))
	case ssd.Hex == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Hex data missing"))
	case !calculations.PBGvalid(ssd.PBG, ssd.MW_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf("PBG data not valid"))
	case ssd.MW_Importance == "{+?}":

		ssd.errors = append(ssd.errors, fmt.Errorf("Importance data does not present correctly (fixable)"))
	//case !strings.Contains(ssd.MW_UWP, "?") && ssd.MW_ImportanceInt != calculations.Importance(uwp.Starport().String(), uwp.TL().String(), uwp.Pops().String(), ssd.Bases, ssd.MW_Remarks):
	//	ssd.errors = append(ssd.errors, fmt.Errorf("Importance data does not match predicted"))
	case ssd.MW_Economic == "(???+?)":
		ssd.errors = append(ssd.errors, fmt.Errorf("Economic Not calculated"))
	case ssd.MW_Economic == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Economic data missing"))
	case ssd.MW_Cultural == "[????]":
		ssd.errors = append(ssd.errors, fmt.Errorf("Cultural data Not calculated"))
	case importanceToInt(ssd.MW_Importance) != ssd.MW_ImportanceInt:
		ssd.errors = append(ssd.errors, fmt.Errorf("Importance data does not match"))

	}
}

func (ssd *SecondSurveyData) NameByConvention() string {
	x := ssd.CoordX
	pX := "S"
	if x < 0 {
		x = x * -1
		pX = "T"
	}
	y := ssd.CoordY
	pY := "R"
	if y < 0 {
		y = y * -1
		pY = "C"
	}
	return fmt.Sprintf("%v %v/%v%v-%v%v", ssd.Sector, ssd.Hex, pX, x, pY, y)
}

func importanceToInt(str string) int {
	switch str {
	default:
		return -999
	case "{ -5 }":
		return -5
	case "{ -4 }":
		return -4
	case "{ -3 }":
		return -3
	case "{ -2 }":
		return -2
	case "{ -1 }":
		return -1
	case "{ 0 }":
		return 0
	case "{ 1 }":
		return 1
	case "{ 2 }":
		return 2
	case "{ 3 }":
		return 3
	case "{ 4 }":
		return 4
	case "{ 5 }":
		return 5
	case "{ 6 }":
		return 6
	}
}

func importanceToString(i int) string {
	switch i {
	default:
		return "{+?}"
	case -5:
		return "{ -5 }"
	case -4:
		return "{ -4 }"
	case -3:
		return "{ -3 }"
	case -2:
		return "{ -2 }"
	case -1:
		return "{ -1 }"
	case 0:
		return "{ 0 }"
	case 1:
		return "{ 1 }"
	case 2:
		return "{ 2 }"
	case 3:
		return "{ 3 }"
	case 4:
		return "{ 4 }"
	case 5:
		return "{ 5 }"
	case 6:
		return "{ 6 }"
	}
}
