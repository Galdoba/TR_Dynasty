package survey

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/universe/starsystem"
	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey/calculations"
	"github.com/Galdoba/utils"
)

const (
	cleanedDataPath = "c:\\Users\\Public\\TrvData\\cleanedData.txt"
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
	SubSector        string
	SubSectorInt     int
	Quadrant         int
	BasesOld         string
	SectorAbb        string
	SubSectorName    string
	AllegianceExt    string
	errors           []error
}

func Parse(input string) *SecondSurveyData {
	ssd := SecondSurveyData{}
	//ssd.input = input
	data := strings.Split(input, "|")
	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}
	ssd.MW_Name = data[1]
	ssd.Hex = data[2]
	ssd.MW_UWP = data[3]
	ssd.PBG = data[4]
	ssd.TravelZone = data[5]
	ssd.Bases = data[6]
	ssd.Allegiance = data[7]
	ssd.Stellar = data[8]
	ssd.SubSector = data[9]
	ssd.MW_Importance = data[10]
	impInt, errImp := strconv.Atoi(data[11])
	if errImp != nil {
		ssd.errors = append(ssd.errors, errImp)
	}
	ssd.MW_ImportanceInt = impInt
	ssd.MW_Economic = data[12]
	ssd.MW_Cultural = data[13]
	ssd.MW_Nobility = data[14]
	worlds, errWorlds := strconv.Atoi(data[15])
	if errWorlds != nil {
		ssd.errors = append(ssd.errors, errWorlds)
	}
	ssd.Worlds = worlds
	ru, errRu := strconv.Atoi(data[16])
	if errRu != nil {
		ssd.errors = append(ssd.errors, errRu)
	}
	ssd.RU = ru
	ssInt, errssInt := strconv.Atoi(data[17])
	if errssInt != nil {
		ssd.errors = append(ssd.errors, errssInt)
	}
	ssd.SubSectorInt = ssInt
	ssQuad, errQuad := strconv.Atoi(data[18])
	if errQuad != nil {
		ssd.errors = append(ssd.errors, errQuad)
	}
	ssd.Quadrant = ssQuad
	xCoord, errXcoord := strconv.Atoi(data[19])
	if errXcoord != nil {
		ssd.errors = append(ssd.errors, errXcoord)
	}
	ssd.CoordX = xCoord
	yCoord, errYcoord := strconv.Atoi(data[20])
	if errYcoord != nil {
		ssd.errors = append(ssd.errors, errYcoord)
	}
	ssd.CoordY = yCoord
	ssd.MW_Remarks = data[21]
	ssd.BasesOld = data[22]
	ssd.Sector = data[23]
	ssd.SubSectorName = data[24]
	ssd.SectorAbb = data[25]
	ssd.AllegianceExt = data[26]
	ssd.verify()
	return &ssd
}

func (ssd *SecondSurveyData) Compress() string {
	compressed := "|"
	compressed += fmt.Sprintf("%v|", ssd.MW_Name)
	compressed += fmt.Sprintf("%v|", ssd.Hex)
	compressed += fmt.Sprintf("%v|", ssd.MW_UWP)
	compressed += fmt.Sprintf("%v|", ssd.PBG)
	compressed += fmt.Sprintf("%v|", ssd.TravelZone)
	compressed += fmt.Sprintf("%v|", ssd.Bases)
	compressed += fmt.Sprintf("%v|", ssd.Allegiance)
	compressed += fmt.Sprintf("%v|", ssd.Stellar)
	compressed += fmt.Sprintf("%v|", ssd.SubSector)
	compressed += fmt.Sprintf("%v|", ssd.MW_Importance) //10
	compressed += fmt.Sprintf("%v|", ssd.MW_ImportanceInt)
	compressed += fmt.Sprintf("%v|", ssd.MW_Economic)
	compressed += fmt.Sprintf("%v|", ssd.MW_Cultural)
	compressed += fmt.Sprintf("%v|", ssd.MW_Nobility)
	compressed += fmt.Sprintf("%v|", ssd.Worlds)
	compressed += fmt.Sprintf("%v|", ssd.RU)
	compressed += fmt.Sprintf("%v|", ssd.SubSectorInt) //17
	compressed += fmt.Sprintf("%v|", ssd.Quadrant)
	compressed += fmt.Sprintf("%v|", ssd.CoordX)
	compressed += fmt.Sprintf("%v|", ssd.CoordY)
	compressed += fmt.Sprintf("%v|", ssd.MW_Remarks) //21
	compressed += fmt.Sprintf("%v|", ssd.BasesOld)   //22
	compressed += fmt.Sprintf("%v|", ssd.Sector)     //23
	compressed += fmt.Sprintf("%v|", ssd.SubSectorName)
	compressed += fmt.Sprintf("%v|", ssd.SectorAbb)
	compressed += fmt.Sprintf("%v", ssd.AllegianceExt)

	return compressed
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
	if calculations.AllegianceFull(ssd.Allegiance) == "UNKNOWN SHORTFORM" {
		ssd.Allegiance = "XXXX"
		ssd.AllegianceExt = calculations.AllegianceFull(ssd.Allegiance)
	}
	switch {
	default:
		return
	case ssd.MW_Name == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("mainworld name missing (fixed)"))
	case ssd.Stellar == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("stellar data missing (f)"))
	case ssd.Hex == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("hex data missing"))
	case !calculations.PBGvalid(ssd.PBG, ssd.MW_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf(" PBG data not valid"))
	case ssd.MW_Importance == "{+?}":
		ssd.errors = append(ssd.errors, fmt.Errorf("importance data does not present correctly (fixable)"))
	case ssd.MW_Economic == "(???+?)":
		ssd.errors = append(ssd.errors, fmt.Errorf("economic Not calculated"))
	case !calculations.ExValid(ssd.MW_Economic):
		ssd.errors = append(ssd.errors, fmt.Errorf("economic Not Valid"))
	case ssd.MW_Economic == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("economic data missing"))
	case ssd.MW_Cultural == "[????]":
		ssd.errors = append(ssd.errors, fmt.Errorf("cultural data Not calculated"))
	case importanceToInt(ssd.MW_Importance) != ssd.MW_ImportanceInt:
		ssd.errors = append(ssd.errors, fmt.Errorf("importance data does not match"))
	case calculations.RU(ssd.MW_Economic) != ssd.RU:
		ssd.errors = append(ssd.errors, fmt.Errorf("projected Ru does not match actual"))
	case !calculations.CxValid(ssd.MW_Cultural, ssd.MW_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf("culture data invalid"))
	case !calculations.WorldsValid(ssd.Worlds, ssd.PBG):
		ssd.errors = append(ssd.errors, fmt.Errorf("world number incorrect (have %v)", ssd.Worlds))
	case len(calculations.NobilityErrors(ssd.MW_Nobility, strings.Fields(ssd.MW_Remarks), ssd.MW_ImportanceInt)) != 0:
		ssd.errors = append(ssd.errors, calculations.NobilityErrors(ssd.MW_Nobility, strings.Fields(ssd.MW_Remarks), ssd.MW_ImportanceInt)...)
	case calculations.AllegianceFull(ssd.Allegiance) == "UNKNOWN SHORTFORM":
		ssd.errors = append(ssd.errors, fmt.Errorf("allegiance unknown"))
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
	for f := range fields {
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
		line := "|"
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

func Search(key string) ([]*SecondSurveyData, error) {
	err := fmt.Errorf("Search not implemented")
	var ssdArr []*SecondSurveyData
	lines := utils.LinesFromTXT(cleanedDataPath)
	for _, val := range lines {
		if strings.Contains(val, key) {
			fmt.Println(":::", val)
			ssdArr = append(ssdArr, Parse(val))
		}
	}
	return ssdArr, err
}
