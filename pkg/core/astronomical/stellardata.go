package astronomical

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

//StellarData - характеристики описывающие звезду
type StellarData struct {
	SpectralType         string
	Decimal              string
	Size                 string
	tableVal             string
	BolometricMagnitude  float64
	Luminosity           float64
	EffectiveTemperature int
	Radii                float64
	Mass                 float64
}

//NewStellarData - возвращает данные о звезде (тест)
func NewStellarData(stellarCode string) StellarData {
	sd := StellarData{}
	sd.SpectralType, sd.Decimal, sd.Size = DecodeStellar(stellarCode)
	sd.tableVal = tableCode(stellarCode)
	return sd
}

func tableCode(stellarCode string) string {
	spectral, decimal, size := DecodeStellar(stellarCode)
	decI := -1
	if decimal == "" {
		return "D" + spectral
	}
	decimal = strconv.Itoa(((decI / 5) * 5))
	tk := spectral + decimal
	if spectral == "M" && decI == 9 {
		tk = "M9"
	}
	return tk + size
}

const (
	//ZoneUnavailable - зона орбит невозможных для существования планет
	ZoneUnavailable = "Unavailable"
	//ZoneInner - зона орбит получающих слишком много света от звезды
	ZoneInner = "Inner zone"
	//ZoneHabitable - зона орбит оптимальных для развития жизни
	ZoneHabitable = "Habitable zone"
	//ZoneOuter - зона орбит получающих слишком мало света от звезды
	ZoneOuter = "Outer zone"
)

//Zone - Возвращает название зоны в которой находится орбита для текущей звезды
func Zone(orbit int, star string) string {
	closest := ClosestPossibleOrbit(star)
	if orbit < closest {
		return ZoneUnavailable
	}
	hz := HabitableOrbit(star)
	pos := orbit - hz
	if pos < 0 {
		return ZoneInner
	}
	if pos > 0 {
		return ZoneOuter
	}
	return ZoneHabitable
}

//HabitableZoneScore - возвращает отступ орбиты от благоприятной для заданной звезды.
//TODO: подобрать более оптимальное имя
func HabitableZoneScore(orbit int, star string) int {
	hz := HabitableOrbit(star)
	return orbit - hz
}

//HabitableOrbit - возвращает оптимальныю для развития жизни орбиту
func HabitableOrbit(star string) int {
	spectral := starSpectral(star)
	size := starSize(star)
	if star == "BD" {
		spectral = "B"
		size = "D"
	}
	hzMap := make(map[string]int)
	hzMap["OIa"] = 15
	hzMap["OIb"] = 15
	hzMap["OII"] = 14
	hzMap["OIII"] = 13
	hzMap["OIV"] = 12
	hzMap["OV"] = 11
	hzMap["OVI"] = -1
	hzMap["OD"] = 1
	hzMap["BIa"] = 13
	hzMap["BIb"] = 13
	hzMap["BII"] = 12
	hzMap["BIII"] = 11
	hzMap["BIV"] = 10
	hzMap["BV"] = 9
	hzMap["BVI"] = -1
	hzMap["BD"] = 0
	hzMap["AIa"] = 12
	hzMap["AIb"] = 11
	hzMap["AII"] = 9
	hzMap["AIII"] = 7
	hzMap["AIV"] = 7
	hzMap["AV"] = 7
	hzMap["AVI"] = -1
	hzMap["AD"] = 0
	hzMap["FIa"] = 11
	hzMap["FIb"] = 10
	hzMap["FII"] = 9
	hzMap["FIII"] = 6
	hzMap["FIV"] = 6
	hzMap["FV"] = 4
	hzMap["FVI"] = 3
	hzMap["FD"] = 0
	hzMap["GIa"] = 12
	hzMap["GIb"] = 10
	hzMap["GII"] = 9
	hzMap["GIII"] = 7
	hzMap["GIV"] = 5
	hzMap["GV"] = 3
	hzMap["GVI"] = 2
	hzMap["GD"] = 0
	hzMap["KIa"] = 12
	hzMap["KIb"] = 10
	hzMap["KII"] = 9
	hzMap["KIII"] = 8
	hzMap["KIV"] = 5
	hzMap["KV"] = 2
	hzMap["KVI"] = 1
	hzMap["KD"] = 0
	hzMap["MIa"] = 12
	hzMap["MIb"] = 11
	hzMap["MII"] = 10
	hzMap["MIII"] = 9
	hzMap["MIV"] = -1
	hzMap["MV"] = 0
	hzMap["MVI"] = 0
	hzMap["MD"] = 0
	if val, ok := hzMap[spectral+size]; !ok {
		fmt.Println(val, star, spectral+size)
		panic("star class unrecognized")
	}
	return hzMap[spectral+size]
}

func ShadowOrbit(star string) int {
	spectral := starSpectral(star)
	size := starSize(star)
	if star == "BD" {
		spectral = "B"
		size = "D"
	}
	hzMap := make(map[string]int)
	hzMap["OIa"] = 15
	hzMap["OIb"] = 15
	hzMap["OII"] = 14
	hzMap["OIII"] = 13
	hzMap["OIV"] = 12
	hzMap["OV"] = 11
	hzMap["OVI"] = -1
	hzMap["OD"] = 1
	hzMap["BIa"] = 13
	hzMap["BIb"] = 13
	hzMap["BII"] = 12
	hzMap["BIII"] = 11
	hzMap["BIV"] = 10
	hzMap["BV"] = 9
	hzMap["BVI"] = -1
	hzMap["BD"] = 0
	hzMap["AIa"] = 10
	hzMap["AIb"] = 9
	hzMap["AII"] = 7
	hzMap["AIII"] = 5
	hzMap["AIV"] = 4
	hzMap["AV"] = 4
	hzMap["AVI"] = -1
	hzMap["AD"] = 0
	hzMap["FIa"] = 11
	hzMap["FIb"] = 9
	hzMap["FII"] = 7
	hzMap["FIII"] = 5
	hzMap["FIV"] = 4
	hzMap["FV"] = 3
	hzMap["FVI"] = 3
	hzMap["FD"] = 0
	hzMap["GIa"] = 12
	hzMap["GIb"] = 10
	hzMap["GII"] = 8
	hzMap["GIII"] = 7
	hzMap["GIV"] = 4
	hzMap["GV"] = 2
	hzMap["GVI"] = 1
	hzMap["GD"] = 0
	hzMap["KIa"] = 13
	hzMap["KIb"] = 12
	hzMap["KII"] = 10
	hzMap["KIII"] = 9
	hzMap["KIV"] = -1
	hzMap["KV"] = 1
	hzMap["KVI"] = 0
	hzMap["KD"] = 0
	hzMap["MIa"] = 15
	hzMap["MIb"] = 14
	hzMap["MII"] = 13
	hzMap["MIII"] = 11
	hzMap["MIV"] = -1
	hzMap["MV"] = 0
	hzMap["MVI"] = 0
	hzMap["MD"] = 0
	if val, ok := hzMap[spectral+size]; !ok {
		fmt.Println(val, star, spectral+size)
		panic("star class unrecognized")
	}
	return hzMap[spectral+size]
}

func starSpectral(starCode string) string {
	stSp := ""
	if strings.Contains(starCode, "BD") {
		stSp = "BD"
	}
	if strings.Contains(starCode, "O") {
		stSp = "O"
	}
	if strings.Contains(starCode, "B") {
		stSp = "B"
	}
	if strings.Contains(starCode, "A") {
		stSp = "A"
	}
	if strings.Contains(starCode, "F") {
		stSp = "F"
	}
	if strings.Contains(starCode, "G") {
		stSp = "G"
	}
	if strings.Contains(starCode, "K") {
		stSp = "K"
	}
	if strings.Contains(starCode, "M") {
		stSp = "M"
	}

	return stSp
}

func starSize(starCode string) string {
	stSp := "Ia"
	if strings.Contains(starCode, "Ib") {
		stSp = "Ib"
	}
	if strings.Contains(starCode, "II") {
		stSp = "II"
	}
	if strings.Contains(starCode, "III") {
		stSp = "III"
	}
	if strings.Contains(starCode, "V") {
		stSp = "V"
	}
	if strings.Contains(starCode, "IV") {
		stSp = "IV"
	}
	if strings.Contains(starCode, "VI") {
		stSp = "VI"
	}
	if strings.Contains(starCode, "D") {
		stSp = "D"
	}
	return stSp
}

func starDecimal(starCode string) (int, string) {
	if strings.Contains(starCode, "0") {
		return 0, "0"
	}
	if strings.Contains(starCode, "1") {
		return 1, "1"
	}
	if strings.Contains(starCode, "2") {
		return 2, "2"
	}
	if strings.Contains(starCode, "3") {
		return 3, "3"
	}
	if strings.Contains(starCode, "4") {
		return 4, "4"
	}
	if strings.Contains(starCode, "5") {
		return 5, "5"
	}
	if strings.Contains(starCode, "6") {
		return 6, "6"
	}
	if strings.Contains(starCode, "7") {
		return 7, "7"
	}
	if strings.Contains(starCode, "8") {
		return 8, "8"
	}
	if strings.Contains(starCode, "9") {
		return 9, "9"
	}
	return -1, ""
}

//ClosestPossibleOrbit - возвращает орбиту максимально близкую к звезде их тех
//где возможна планета
func ClosestPossibleOrbit(star string) int {
	spectral := starSpectral(star)
	size := starSize(star)
	if star == "BD" {
		spectral = "B"
		size = "D"
	}
	hzMap := make(map[string]int)
	hzMap["OIa"] = 8
	hzMap["OIb"] = 7
	hzMap["OII"] = 6
	hzMap["OIII"] = 5
	hzMap["OIV"] = 4
	hzMap["OV"] = 5
	hzMap["OVI"] = -1
	hzMap["OD"] = 0
	hzMap["BIa"] = 7
	hzMap["BIb"] = 6
	hzMap["BII"] = 5
	hzMap["BIII"] = 4
	hzMap["BIV"] = 3
	hzMap["BV"] = 4
	hzMap["BVI"] = -1
	hzMap["BD"] = 0
	hzMap["AIa"] = 7
	hzMap["AIb"] = 5
	hzMap["AII"] = 3
	hzMap["AIII"] = 1
	hzMap["AIV"] = 1
	hzMap["AV"] = 0
	hzMap["AVI"] = -1
	hzMap["AD"] = 0
	hzMap["FIa"] = 6
	hzMap["FIb"] = 4
	hzMap["FII"] = 2
	hzMap["FIII"] = 0
	hzMap["FIV"] = 0
	hzMap["FV"] = 0
	hzMap["FVI"] = 0
	hzMap["FD"] = 0
	hzMap["GIa"] = 7
	hzMap["GIb"] = 5
	hzMap["GII"] = 2
	hzMap["GIII"] = 0
	hzMap["GIV"] = 0
	hzMap["GV"] = 0
	hzMap["GVI"] = 0
	hzMap["GD"] = 0
	hzMap["KIa"] = 7
	hzMap["KIb"] = 6
	hzMap["KII"] = 3
	hzMap["KIII"] = 0
	hzMap["KIV"] = 0
	hzMap["KV"] = 0
	hzMap["KVI"] = 0
	hzMap["KD"] = 0
	hzMap["MIa"] = 8
	hzMap["MIb"] = 7
	hzMap["MII"] = 6
	hzMap["MIII"] = 4
	hzMap["MIV"] = -1
	hzMap["MV"] = 0
	hzMap["MVI"] = 0
	hzMap["MD"] = 0
	if val, ok := hzMap[spectral+size]; !ok {
		fmt.Println(val, star, spectral+size)
		panic("star class unrecognized")
	}
	return hzMap[spectral+size]
}

//K2 IV BD D
//
func DecodeStellar(stellarCode string) (string, string, string) {
	switch stellarCode {
	case "DB":
		return "B", "", "D"
	case "DA":
		return "A", "", "D"
	case "DF":
		return "F", "", "D"
	case "DG":
		return "G", "", "D"
	case "DK":
		return "K", "", "D"
	case "DM":
		return "M", "", "D"
	}
	spectral := starSpectral(stellarCode)
	_, decimal := starDecimal(stellarCode)
	size := starSize(stellarCode)
	return spectral, decimal, size
}

func allStarCodes() []string {
	return []string{
		"BD",
		"D",
		"DB",
		"DA",
		"DF",
		"DG",
		"DK",
		"DM",
		"F0Ia",
		"F1Ia",
		"F2Ia",
		"F3Ia",
		"F4Ia",
		"F5Ia",
		"F6Ia",
		"F7Ia",
		"F8Ia",
		"F9Ia",
		"F0Ib",
		"F1Ib",
		"F2Ib",
		"F3Ib",
		"F4Ib",
		"F5Ib",
		"F6Ib",
		"F7Ib",
		"F8Ib",
		"F9Ib",
		"F0II",
		"F1II",
		"F2II",
		"F3II",
		"F4II",
		"F5II",
		"F6II",
		"F7II",
		"F8II",
		"F9II",
		"F0III",
		"F1III",
		"F2III",
		"F3III",
		"F4III",
		"F5III",
		"F6III",
		"F7III",
		"F8III",
		"F9III",
		"F0IV",
		"F1IV",
		"F2IV",
		"F3IV",
		"F4IV",
		"F5IV",
		"F6IV",
		"F7IV",
		"F8IV",
		"F9IV",
		"F0V",
		"F1V",
		"F2V",
		"F3V",
		"F4V",
		"F5V",
		"F6V",
		"F7V",
		"F8V",
		"F9V",
		"F0VI",
		"F1VI",
		"F2VI",
		"F3VI",
		"F4VI",
		"F5VI",
		"G0Ia",
		"G1Ia",
		"G2Ia",
		"G3Ia",
		"G4Ia",
		"G5Ia",
		"G6Ia",
		"G7Ia",
		"G8Ia",
		"G9Ia",
		"G0Ib",
		"G1Ib",
		"G2Ib",
		"G3Ib",
		"G4Ib",
		"G5Ib",
		"G6Ib",
		"G7Ib",
		"G8Ib",
		"G9Ib",
		"G0II",
		"G1II",
		"G2II",
		"G3II",
		"G4II",
		"G5II",
		"G6II",
		"G7II",
		"G8II",
		"G9II",
		"G0III",
		"G1III",
		"G2III",
		"G3III",
		"G4III",
		"G5III",
		"G6III",
		"G7III",
		"G8III",
		"G9III",
		"G0IV",
		"G1IV",
		"G2IV",
		"G3IV",
		"G4IV",
		"G5IV",
		"G6IV",
		"G7IV",
		"G8IV",
		"G9IV",
		"G0V",
		"G1V",
		"G2V",
		"G3V",
		"G4V",
		"G5V",
		"G6V",
		"G7V",
		"G8V",
		"G9V",
		"G0VI",
		"G1VI",
		"G2VI",
		"G3VI",
		"G4VI",
		"G5VI",
		"G6VI",
		"G7VI",
		"G8VI",
		"G9VI",
		"K0Ia",
		"K1Ia",
		"K2Ia",
		"K3Ia",
		"K4Ia",
		"K5Ia",
		"K6Ia",
		"K7Ia",
		"K8Ia",
		"K9Ia",
		"K0Ib",
		"K1Ib",
		"K2Ib",
		"K3Ib",
		"K4Ib",
		"K5Ib",
		"K6Ib",
		"K7Ib",
		"K8Ib",
		"K9Ib",
		"K0II",
		"K1II",
		"K2II",
		"K3II",
		"K4II",
		"K5II",
		"K6II",
		"K7II",
		"K8II",
		"K9II",
		"K0III",
		"K1III",
		"K2III",
		"K3III",
		"K4III",
		"K5III",
		"K6III",
		"K7III",
		"K8III",
		"K9III",
		"K0IV",
		"K1IV",
		"K2IV",
		"K3IV",
		"K4IV",
		"K5IV",
		"K0V",
		"K1V",
		"K2V",
		"K3V",
		"K4V",
		"K5V",
		"K6V",
		"K7V",
		"K8V",
		"K9V",
		"K0VI",
		"K1VI",
		"K2VI",
		"K3VI",
		"K4VI",
		"K5VI",
		"K6VI",
		"K7VI",
		"K8VI",
		"K9VI",
		"M0Ia",
		"M1Ia",
		"M2Ia",
		"M3Ia",
		"M4Ia",
		"M5Ia",
		"M6Ia",
		"M7Ia",
		"M8Ia",
		"M9Ia",
		"M0Ib",
		"M1Ib",
		"M2Ib",
		"M3Ib",
		"M4Ib",
		"M5Ib",
		"M6Ib",
		"M7Ib",
		"M8Ib",
		"M9Ib",
		"M0II",
		"M1II",
		"M2II",
		"M3II",
		"M4II",
		"M5II",
		"M6II",
		"M7II",
		"M8II",
		"M9II",
		"M0III",
		"M1III",
		"M2III",
		"M3III",
		"M4III",
		"M5III",
		"M6III",
		"M7III",
		"M8III",
		"M9III",
		"M0V",
		"M1V",
		"M2V",
		"M3V",
		"M4V",
		"M5V",
		"M6V",
		"M7V",
		"M8V",
		"M9V",
		"M0VI",
		"M1VI",
		"M2VI",
		"M3VI",
		"M4VI",
		"M5VI",
		"M6VI",
		"M7VI",
		"M8VI",
		"M9VI",
		"O0Ia",
		"O1Ia",
		"O2Ia",
		"O3Ia",
		"O4Ia",
		"O5Ia",
		"O6Ia",
		"O7Ia",
		"O8Ia",
		"O9Ia",
		"O0Ib",
		"O1Ib",
		"O2Ib",
		"O3Ib",
		"O4Ib",
		"O5Ib",
		"O6Ib",
		"O7Ib",
		"O8Ib",
		"O9Ib",
		"O0II",
		"O1II",
		"O2II",
		"O3II",
		"O4II",
		"O5II",
		"O6II",
		"O7II",
		"O8II",
		"O9II",
		"O0III",
		"O1III",
		"O2III",
		"O3III",
		"O4III",
		"O5III",
		"O6III",
		"O7III",
		"O8III",
		"O9III",
		"O0IV",
		"O1IV",
		"O2IV",
		"O3IV",
		"O4IV",
		"O5IV",
		"O6IV",
		"O7IV",
		"O8IV",
		"O9IV",
		"O0V",
		"O1V",
		"O2V",
		"O3V",
		"O4V",
		"O5V",
		"O6V",
		"O7V",
		"O8V",
		"O9V",
		"B0Ia",
		"B1Ia",
		"B2Ia",
		"B3Ia",
		"B4Ia",
		"B5Ia",
		"B6Ia",
		"B7Ia",
		"B8Ia",
		"B9Ia",
		"B0Ib",
		"B1Ib",
		"B2Ib",
		"B3Ib",
		"B4Ib",
		"B5Ib",
		"B6Ib",
		"B7Ib",
		"B8Ib",
		"B9Ib",
		"B0II",
		"B1II",
		"B2II",
		"B3II",
		"B4II",
		"B5II",
		"B6II",
		"B7II",
		"B8II",
		"B9II",
		"B0III",
		"B1III",
		"B2III",
		"B3III",
		"B4III",
		"B5III",
		"B6III",
		"B7III",
		"B8III",
		"B9III",
		"B0IV",
		"B1IV",
		"B2IV",
		"B3IV",
		"B4IV",
		"B5IV",
		"B6IV",
		"B7IV",
		"B8IV",
		"B9IV",
		"B0V",
		"B1V",
		"B2V",
		"B3V",
		"B4V",
		"B5V",
		"B6V",
		"B7V",
		"B8V",
		"B9V",
		"B0VI",
		"A0Ia",
		"A1Ia",
		"A2Ia",
		"A3Ia",
		"A4Ia",
		"A5Ia",
		"A6Ia",
		"A7Ia",
		"A8Ia",
		"A9Ia",
		"A0Ib",
		"A1Ib",
		"A2Ib",
		"A3Ib",
		"A4Ib",
		"A5Ib",
		"A6Ib",
		"A7Ib",
		"A8Ib",
		"A9Ib",
		"A0II",
		"A1II",
		"A2II",
		"A3II",
		"A4II",
		"A5II",
		"A6II",
		"A7II",
		"A8II",
		"A9II",
		"A0III",
		"A1III",
		"A2III",
		"A3III",
		"A4III",
		"A5III",
		"A6III",
		"A7III",
		"A8III",
		"A9III",
		"A0IV",
		"A1IV",
		"A2IV",
		"A3IV",
		"A4IV",
		"A5IV",
		"A6IV",
		"A7IV",
		"A8IV",
		"A9IV",
		"A0V",
		"A1V",
		"A2V",
		"A3V",
		"A4V",
		"A5V",
		"A6V",
		"A7V",
		"A8V",
		"A9V",
	}
}

//StellarDataMap - хранит в себе Spectral Types, Stellar Sizes и Stellar Digit
var StellarDataMap map[string][]string

func stellarToNum(spCls string) int {
	class := " "
	dec := " "
	spDat := strings.Split(spCls, "")
	validClass := []string{"O", "B", "A", "F", "G", "K", "M"}
	validDec := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if len(spDat) < 2 {
		fmt.Println("ERROR: Stellar Data not parsed")
		return -3
	}
	switch {
	case utils.ListContains(validClass, spDat[0]) && utils.ListContains(validDec, spDat[1]):
		class = spDat[0]
		dec = spDat[1]
	default:
		fmt.Println("ERROR: Stellar Data not parsed")
		return -2
	}
	index := 0
	for i, val := range validClass {
		if val == class {
			index = index + (i * 10)
		}
	}
	for i, val := range validDec {
		if val == dec {
			index = index + i
		}
	}
	return index
}

/*



SecondSurvey <= Name

Placement <= SecondSurvey
Flux <= SecondSurvey

0 ORBITAL
0.1- Orbit Zones <= placement

1 SIZE
1.1 Basic World Type				placement
1.2a Planet Diameter				UWP, Flux
1.2b Planet Density					UWP, Flux
1.3a GG UWP Size					placement
1.3b GG Diameter					1.3a, Flux
1.3c GG Density						Flux
1.4a Belt Diameter					Flux
1.4b Belt Zones						placement, SecondSurvey
1.4c Belt Orbit Width				placement
1.4d Belt Profile					1.4a, 1.4b, 1.4c
1.5 World Mass						UWP, 1.2b
1.6 World Gravity					UWP, 1.5
1.7 Planet Orbital Priod
1.7a Stellar Mass					placement, SecondSurvey
1.7b Orbital Distance				placement
1.7c Orbital Period					1.7a, 1.7b
1.8 Satellite Orbital Priod
1.8a Satellite Orbital Distance		placement, 1.2a
1.8b Orbital Priod					1.8a, 1.5
1.9 Rotational Period				Flux, (1.7a || 1.5), 1.7b
1.10 Axial Tilt						Flux
1.11 Orbital Eccentricity			Flux
1.12 Seismic Stress Factor			1.2b, placement, (1.7a || 1.5)

2 ATMO
2.1 Atmospheric Composition			UWP, Flux
2.2 Surface Atmospheric Pressure	UWP, Flux
2.3 Surface Temperature
2.3a Stellar Luminosity				placement, SecondSurvey
2.3b Orbit Factor					placement
2.3c Energy Absorbtion				UWP
2.3d Greenhouse Effect				UWP, Flux
2.3e Base Temperature				2.3a, 2.3b, 2.3c, 2.3d
2.4 Orbital Eccentricity Effects	1.11
2.5 Latitude Temperature Effects	UWP
2.6a Axial Tilt Base Increase		2.3e, 1.10
2.6b Axial Tilt Base Decrease		2.3e, 1.10
2.6c Axial Tilt Latitude Effects	1.10
2.7a Lenght of Day and Night		1.9
2.7b Rotation-Luminocity Factor		2.3a, 1.7b
2.7c Daytime Rotation Effects		UWP, 1.9, 2.3e,
2.7d Nighttime Rotation Effects		UWP, 1.9, 2.3e,


*/
