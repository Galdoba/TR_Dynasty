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
	ssd.Allegiance = data[7]
	ssd.Stellar = data[8]
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
		ssd.PBG = calculations.FixPBG(ssd.PBG, ssd.MW_UWP, ssd.NameByConvention())
	}
	if ssd.MW_Importance == "{+?}" {
		ssd.MW_Importance = importanceToString(ssd.MW_ImportanceInt)
		calc := calculations.Importance(ssd.MW_UWP, ssd.Bases, ssd.MW_Remarks)
		if calc != ssd.MW_ImportanceInt && ssd.MW_ImportanceInt == 0 {
			ssd.MW_Importance = importanceToString(calc)
			ssd.MW_ImportanceInt = calc
		}
	}
	if importanceToInt(ssd.MW_Importance) != ssd.MW_ImportanceInt {
		ssd.MW_Importance = importanceToString(ssd.MW_ImportanceInt)
	}
	if !calculations.ExValid(ssd.MW_Economic) {
		ssd.MW_Economic = calculations.FixEconomicExtention(ssd.MW_Economic, ssd.MW_UWP, ssd.PBG, ssd.NameByConvention(), ssd.MW_ImportanceInt)
	}
	if calculations.RU(ssd.MW_Economic) != ssd.RU {
		ssd.RU = calculations.RU(ssd.MW_Economic)
	}
	if !calculations.CxValid(ssd.MW_Cultural, ssd.MW_UWP) {
		ssd.MW_Cultural = calculations.Cultural(ssd.MW_UWP, ssd.NameByConvention(), ssd.MW_ImportanceInt)
	}
	if !calculations.WorldsValid(ssd.Worlds, ssd.PBG) {
		ssd.Worlds = calculations.FixWorlds(ssd.PBG, ssd.NameByConvention())
	}
	if len(calculations.NobilityErrors(ssd.MW_Nobility, strings.Fields(ssd.MW_Remarks), ssd.MW_ImportanceInt)) != 0 {
		ssd.MW_Nobility = calculations.FixNobility(strings.Fields(ssd.MW_Remarks), ssd.MW_ImportanceInt)
	}
	if !calculations.CxValid(ssd.MW_Cultural, ssd.MW_UWP) {
		fmt.Println("invalid culture data:", ssd.MW_Cultural)
	}
	switch {
	default:
		return
	case ssd.MW_Name == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Mainworld name missing (fixed)"))
	case ssd.Stellar == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Stellar data missing (f)"))
	case ssd.Hex == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Hex data missing"))
	case !calculations.PBGvalid(ssd.PBG, ssd.MW_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf("PBG data not valid"))
	case ssd.MW_Importance == "{+?}":
		ssd.errors = append(ssd.errors, fmt.Errorf("Importance data does not present correctly (fixable)"))
	case ssd.MW_Economic == "(???+?)":
		ssd.errors = append(ssd.errors, fmt.Errorf("Economic Not calculated"))
	case !calculations.ExValid(ssd.MW_Economic):
		ssd.errors = append(ssd.errors, fmt.Errorf("Economic Not Valid"))
	case ssd.MW_Economic == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("Economic data missing"))
	case ssd.MW_Cultural == "[????]":
		ssd.errors = append(ssd.errors, fmt.Errorf("Cultural data Not calculated"))
	case importanceToInt(ssd.MW_Importance) != ssd.MW_ImportanceInt:
		ssd.errors = append(ssd.errors, fmt.Errorf("Importance data does not match"))
	case calculations.RU(ssd.MW_Economic) != ssd.RU:
		ssd.errors = append(ssd.errors, fmt.Errorf("Projected Ru does not match actual"))
	case !calculations.CxValid(ssd.MW_Cultural, ssd.MW_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf("Culture data invalid"))
	case !calculations.WorldsValid(ssd.Worlds, ssd.PBG):
		ssd.errors = append(ssd.errors, fmt.Errorf("world number incorrect (have %v)", ssd.Worlds))
	case len(calculations.NobilityErrors(ssd.MW_Nobility, strings.Fields(ssd.MW_Remarks), ssd.MW_ImportanceInt)) != 0:
		ssd.errors = append(ssd.errors, calculations.NobilityErrors(ssd.MW_Nobility, strings.Fields(ssd.MW_Remarks), ssd.MW_ImportanceInt)...)

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
	case "{ +0 }":
		return 0
	case "{ +1 }":
		return 1
	case "{ +2 }":
		return 2
	case "{ +3 }":
		return 3
	case "{ +4 }":
		return 4
	case "{ +5 }":
		return 5
	case "{ +6 }":
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
		return "{ +0 }"
	case 1:
		return "{ +1 }"
	case 2:
		return "{ +2 }"
	case 3:
		return "{ +3 }"
	case 4:
		return "{ +4 }"
	case 5:
		return "{ +5 }"
	case 6:
		return "{ +6 }"
	}
}

func (ssd *SecondSurveyData) String() string {
	rep := ssd.Hex + "   "
	rep += ssd.MW_Name + "   "
	rep += ssd.MW_UWP + "   "
	rep += ssd.MW_Remarks + "   "
	rep += ssd.MW_Importance + "   "
	rep += ssd.MW_Economic + "   "
	rep += ssd.MW_Cultural + "   "
	rep += ssd.MW_Nobility + "   "
	rep += ssd.Bases + "   "
	rep += ssd.TravelZone + "   "
	rep += ssd.PBG + "   "
	rep += strconv.Itoa(ssd.Worlds) + "   "
	rep += ssd.Allegiance + "   "
	rep += ssd.Stellar
	return rep
}

func ListOf(ssds []*SecondSurveyData) []string {
	if len(ssds) < 1 {
		return nil
	}
	sample := ssds[0].String()
	fields := strings.Split(sample, "   ")
	colMap := make(map[int]int)
	for f, _ := range fields {
		for _, ssd := range ssds {
			testFields := strings.Split(ssd.String(), "   ")
			if colMap[f] < len(testFields[f]) {
				colMap[f] = len(testFields[f])
			}
		}
	}
	table := []string{}
	for _, ssd := range ssds {
		newFields := strings.Split(ssd.String(), "   ")
		line := ""
		for n, fld := range newFields {
			for len(fld) < colMap[n] {
				fld += " "
			}
			line += fld + "|"
		}
		table = append(table, line)
	}
	return table
}
