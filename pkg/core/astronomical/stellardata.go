package astronomical

import (
	"fmt"
	"strconv"
	"strings"
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

// func init() {
// 	fmt.Println("Initiating `pkg/core/astronomical`")
// 	StellarDataMap = make(map[string][]string)
// 	for _, val := range allStarCodes() {
// 		switch val {
// 		case "BD":
// 			StellarDataMap["BD"] = append(StellarDataMap["BD"], "D")
// 			StellarDataMap["BD"] = append(StellarDataMap["BD"], "")
// 			StellarDataMap["BD"] = append(StellarDataMap["BD"], "D")

// 		case "D":
// 			StellarDataMap["D"] = append(StellarDataMap["D"], "D")
// 			StellarDataMap["D"] = append(StellarDataMap["D"], "")
// 			StellarDataMap["D"] = append(StellarDataMap["D"], "D")

// 		case "DB":
// 			StellarDataMap["DB"] = append(StellarDataMap["DB"], "B")
// 			StellarDataMap["DB"] = append(StellarDataMap["DB"], "")
// 			StellarDataMap["DB"] = append(StellarDataMap["DB"], "D")

// 		case "DA":
// 			StellarDataMap["DA"] = append(StellarDataMap["DA"], "A")
// 			StellarDataMap["DA"] = append(StellarDataMap["DA"], "")
// 			StellarDataMap["DA"] = append(StellarDataMap["DA"], "D")

// 		case "DF":
// 			StellarDataMap["DF"] = append(StellarDataMap["DF"], "F")
// 			StellarDataMap["DF"] = append(StellarDataMap["DF"], "")
// 			StellarDataMap["DF"] = append(StellarDataMap["DF"], "D")

// 		case "DG":
// 			StellarDataMap["DG"] = append(StellarDataMap["DG"], "G")
// 			StellarDataMap["DG"] = append(StellarDataMap["DG"], "")
// 			StellarDataMap["DG"] = append(StellarDataMap["DG"], "D")

// 		case "DK":
// 			StellarDataMap["DK"] = append(StellarDataMap["DK"], "K")
// 			StellarDataMap["DK"] = append(StellarDataMap["DK"], "")
// 			StellarDataMap["DK"] = append(StellarDataMap["DK"], "D")

// 		case "DM":
// 			StellarDataMap["DM"] = append(StellarDataMap["DM"], "M")
// 			StellarDataMap["DM"] = append(StellarDataMap["DM"], "")
// 			StellarDataMap["DM"] = append(StellarDataMap["DM"], "D")

// 		case "F0Ia":
// 			StellarDataMap["F0Ia"] = append(StellarDataMap["F0Ia"], "F")
// 			StellarDataMap["F0Ia"] = append(StellarDataMap["F0Ia"], "0")
// 			StellarDataMap["F0Ia"] = append(StellarDataMap["F0Ia"], "Ia")

// 		case "F1Ia":
// 			StellarDataMap["F1Ia"] = append(StellarDataMap["F1Ia"], "F")
// 			StellarDataMap["F1Ia"] = append(StellarDataMap["F1Ia"], "1")
// 			StellarDataMap["F1Ia"] = append(StellarDataMap["F1Ia"], "Ia")

// 		case "F2Ia":
// 			StellarDataMap["F2Ia"] = append(StellarDataMap["F2Ia"], "F")
// 			StellarDataMap["F2Ia"] = append(StellarDataMap["F2Ia"], "2")
// 			StellarDataMap["F2Ia"] = append(StellarDataMap["F2Ia"], "Ia")

// 		case "F3Ia":
// 			StellarDataMap["F3Ia"] = append(StellarDataMap["F3Ia"], "F")
// 			StellarDataMap["F3Ia"] = append(StellarDataMap["F3Ia"], "3")
// 			StellarDataMap["F3Ia"] = append(StellarDataMap["F3Ia"], "Ia")

// 		case "F4Ia":
// 			StellarDataMap["F4Ia"] = append(StellarDataMap["F4Ia"], "F")
// 			StellarDataMap["F4Ia"] = append(StellarDataMap["F4Ia"], "4")
// 			StellarDataMap["F4Ia"] = append(StellarDataMap["F4Ia"], "Ia")

// 		case "F5Ia":
// 			StellarDataMap["F5Ia"] = append(StellarDataMap["F5Ia"], "F")
// 			StellarDataMap["F5Ia"] = append(StellarDataMap["F5Ia"], "5")
// 			StellarDataMap["F5Ia"] = append(StellarDataMap["F5Ia"], "Ia")

// 		case "F6Ia":
// 			StellarDataMap["F6Ia"] = append(StellarDataMap["F6Ia"], "F")
// 			StellarDataMap["F6Ia"] = append(StellarDataMap["F6Ia"], "6")
// 			StellarDataMap["F6Ia"] = append(StellarDataMap["F6Ia"], "Ia")

// 		case "F7Ia":
// 			StellarDataMap["F7Ia"] = append(StellarDataMap["F7Ia"], "F")
// 			StellarDataMap["F7Ia"] = append(StellarDataMap["F7Ia"], "7")
// 			StellarDataMap["F7Ia"] = append(StellarDataMap["F7Ia"], "Ia")

// 		case "F8Ia":
// 			StellarDataMap["F8Ia"] = append(StellarDataMap["F8Ia"], "F")
// 			StellarDataMap["F8Ia"] = append(StellarDataMap["F8Ia"], "8")
// 			StellarDataMap["F8Ia"] = append(StellarDataMap["F8Ia"], "Ia")

// 		case "F9Ia":
// 			StellarDataMap["F9Ia"] = append(StellarDataMap["F9Ia"], "F")
// 			StellarDataMap["F9Ia"] = append(StellarDataMap["F9Ia"], "9")
// 			StellarDataMap["F9Ia"] = append(StellarDataMap["F9Ia"], "Ia")

// 		case "F0Ib":
// 			StellarDataMap["F0Ib"] = append(StellarDataMap["F0Ib"], "F")
// 			StellarDataMap["F0Ib"] = append(StellarDataMap["F0Ib"], "0")
// 			StellarDataMap["F0Ib"] = append(StellarDataMap["F0Ib"], "Ib")

// 		case "F1Ib":
// 			StellarDataMap["F1Ib"] = append(StellarDataMap["F1Ib"], "F")
// 			StellarDataMap["F1Ib"] = append(StellarDataMap["F1Ib"], "1")
// 			StellarDataMap["F1Ib"] = append(StellarDataMap["F1Ib"], "Ib")

// 		case "F2Ib":
// 			StellarDataMap["F2Ib"] = append(StellarDataMap["F2Ib"], "F")
// 			StellarDataMap["F2Ib"] = append(StellarDataMap["F2Ib"], "2")
// 			StellarDataMap["F2Ib"] = append(StellarDataMap["F2Ib"], "Ib")

// 		case "F3Ib":
// 			StellarDataMap["F3Ib"] = append(StellarDataMap["F3Ib"], "F")
// 			StellarDataMap["F3Ib"] = append(StellarDataMap["F3Ib"], "3")
// 			StellarDataMap["F3Ib"] = append(StellarDataMap["F3Ib"], "Ib")

// 		case "F4Ib":
// 			StellarDataMap["F4Ib"] = append(StellarDataMap["F4Ib"], "F")
// 			StellarDataMap["F4Ib"] = append(StellarDataMap["F4Ib"], "4")
// 			StellarDataMap["F4Ib"] = append(StellarDataMap["F4Ib"], "Ib")

// 		case "F5Ib":
// 			StellarDataMap["F5Ib"] = append(StellarDataMap["F5Ib"], "F")
// 			StellarDataMap["F5Ib"] = append(StellarDataMap["F5Ib"], "5")
// 			StellarDataMap["F5Ib"] = append(StellarDataMap["F5Ib"], "Ib")

// 		case "F6Ib":
// 			StellarDataMap["F6Ib"] = append(StellarDataMap["F6Ib"], "F")
// 			StellarDataMap["F6Ib"] = append(StellarDataMap["F6Ib"], "6")
// 			StellarDataMap["F6Ib"] = append(StellarDataMap["F6Ib"], "Ib")

// 		case "F7Ib":
// 			StellarDataMap["F7Ib"] = append(StellarDataMap["F7Ib"], "F")
// 			StellarDataMap["F7Ib"] = append(StellarDataMap["F7Ib"], "7")
// 			StellarDataMap["F7Ib"] = append(StellarDataMap["F7Ib"], "Ib")

// 		case "F8Ib":
// 			StellarDataMap["F8Ib"] = append(StellarDataMap["F8Ib"], "F")
// 			StellarDataMap["F8Ib"] = append(StellarDataMap["F8Ib"], "8")
// 			StellarDataMap["F8Ib"] = append(StellarDataMap["F8Ib"], "Ib")

// 		case "F9Ib":
// 			StellarDataMap["F9Ib"] = append(StellarDataMap["F9Ib"], "F")
// 			StellarDataMap["F9Ib"] = append(StellarDataMap["F9Ib"], "9")
// 			StellarDataMap["F9Ib"] = append(StellarDataMap["F9Ib"], "Ib")

// 		case "F0II":
// 			StellarDataMap["F0II"] = append(StellarDataMap["F0II"], "F")
// 			StellarDataMap["F0II"] = append(StellarDataMap["F0II"], "0")
// 			StellarDataMap["F0II"] = append(StellarDataMap["F0II"], "II")

// 		case "F1II":
// 			StellarDataMap["F1II"] = append(StellarDataMap["F1II"], "F")
// 			StellarDataMap["F1II"] = append(StellarDataMap["F1II"], "1")
// 			StellarDataMap["F1II"] = append(StellarDataMap["F1II"], "II")

// 		case "F2II":
// 			StellarDataMap["F2II"] = append(StellarDataMap["F2II"], "F")
// 			StellarDataMap["F2II"] = append(StellarDataMap["F2II"], "2")
// 			StellarDataMap["F2II"] = append(StellarDataMap["F2II"], "II")

// 		case "F3II":
// 			StellarDataMap["F3II"] = append(StellarDataMap["F3II"], "F")
// 			StellarDataMap["F3II"] = append(StellarDataMap["F3II"], "3")
// 			StellarDataMap["F3II"] = append(StellarDataMap["F3II"], "II")

// 		case "F4II":
// 			StellarDataMap["F4II"] = append(StellarDataMap["F4II"], "F")
// 			StellarDataMap["F4II"] = append(StellarDataMap["F4II"], "4")
// 			StellarDataMap["F4II"] = append(StellarDataMap["F4II"], "II")

// 		case "F5II":
// 			StellarDataMap["F5II"] = append(StellarDataMap["F5II"], "F")
// 			StellarDataMap["F5II"] = append(StellarDataMap["F5II"], "5")
// 			StellarDataMap["F5II"] = append(StellarDataMap["F5II"], "II")

// 		case "F6II":
// 			StellarDataMap["F6II"] = append(StellarDataMap["F6II"], "F")
// 			StellarDataMap["F6II"] = append(StellarDataMap["F6II"], "6")
// 			StellarDataMap["F6II"] = append(StellarDataMap["F6II"], "II")

// 		case "F7II":
// 			StellarDataMap["F7II"] = append(StellarDataMap["F7II"], "F")
// 			StellarDataMap["F7II"] = append(StellarDataMap["F7II"], "7")
// 			StellarDataMap["F7II"] = append(StellarDataMap["F7II"], "II")

// 		case "F8II":
// 			StellarDataMap["F8II"] = append(StellarDataMap["F8II"], "F")
// 			StellarDataMap["F8II"] = append(StellarDataMap["F8II"], "8")
// 			StellarDataMap["F8II"] = append(StellarDataMap["F8II"], "II")

// 		case "F9II":
// 			StellarDataMap["F9II"] = append(StellarDataMap["F9II"], "F")
// 			StellarDataMap["F9II"] = append(StellarDataMap["F9II"], "9")
// 			StellarDataMap["F9II"] = append(StellarDataMap["F9II"], "II")

// 		case "F0III":
// 			StellarDataMap["F0III"] = append(StellarDataMap["F0III"], "F")
// 			StellarDataMap["F0III"] = append(StellarDataMap["F0III"], "0")
// 			StellarDataMap["F0III"] = append(StellarDataMap["F0III"], "III")

// 		case "F1III":
// 			StellarDataMap["F1III"] = append(StellarDataMap["F1III"], "F")
// 			StellarDataMap["F1III"] = append(StellarDataMap["F1III"], "1")
// 			StellarDataMap["F1III"] = append(StellarDataMap["F1III"], "III")

// 		case "F2III":
// 			StellarDataMap["F2III"] = append(StellarDataMap["F2III"], "F")
// 			StellarDataMap["F2III"] = append(StellarDataMap["F2III"], "2")
// 			StellarDataMap["F2III"] = append(StellarDataMap["F2III"], "III")

// 		case "F3III":
// 			StellarDataMap["F3III"] = append(StellarDataMap["F3III"], "F")
// 			StellarDataMap["F3III"] = append(StellarDataMap["F3III"], "3")
// 			StellarDataMap["F3III"] = append(StellarDataMap["F3III"], "III")

// 		case "F4III":
// 			StellarDataMap["F4III"] = append(StellarDataMap["F4III"], "F")
// 			StellarDataMap["F4III"] = append(StellarDataMap["F4III"], "4")
// 			StellarDataMap["F4III"] = append(StellarDataMap["F4III"], "III")

// 		case "F5III":
// 			StellarDataMap["F5III"] = append(StellarDataMap["F5III"], "F")
// 			StellarDataMap["F5III"] = append(StellarDataMap["F5III"], "5")
// 			StellarDataMap["F5III"] = append(StellarDataMap["F5III"], "III")

// 		case "F6III":
// 			StellarDataMap["F6III"] = append(StellarDataMap["F6III"], "F")
// 			StellarDataMap["F6III"] = append(StellarDataMap["F6III"], "6")
// 			StellarDataMap["F6III"] = append(StellarDataMap["F6III"], "III")

// 		case "F7III":
// 			StellarDataMap["F7III"] = append(StellarDataMap["F7III"], "F")
// 			StellarDataMap["F7III"] = append(StellarDataMap["F7III"], "7")
// 			StellarDataMap["F7III"] = append(StellarDataMap["F7III"], "III")

// 		case "F8III":
// 			StellarDataMap["F8III"] = append(StellarDataMap["F8III"], "F")
// 			StellarDataMap["F8III"] = append(StellarDataMap["F8III"], "8")
// 			StellarDataMap["F8III"] = append(StellarDataMap["F8III"], "III")

// 		case "F9III":
// 			StellarDataMap["F9III"] = append(StellarDataMap["F9III"], "F")
// 			StellarDataMap["F9III"] = append(StellarDataMap["F9III"], "9")
// 			StellarDataMap["F9III"] = append(StellarDataMap["F9III"], "III")

// 		case "F0IV":
// 			StellarDataMap["F0IV"] = append(StellarDataMap["F0IV"], "F")
// 			StellarDataMap["F0IV"] = append(StellarDataMap["F0IV"], "0")
// 			StellarDataMap["F0IV"] = append(StellarDataMap["F0IV"], "IV")

// 		case "F1IV":
// 			StellarDataMap["F1IV"] = append(StellarDataMap["F1IV"], "F")
// 			StellarDataMap["F1IV"] = append(StellarDataMap["F1IV"], "1")
// 			StellarDataMap["F1IV"] = append(StellarDataMap["F1IV"], "IV")

// 		case "F2IV":
// 			StellarDataMap["F2IV"] = append(StellarDataMap["F2IV"], "F")
// 			StellarDataMap["F2IV"] = append(StellarDataMap["F2IV"], "2")
// 			StellarDataMap["F2IV"] = append(StellarDataMap["F2IV"], "IV")

// 		case "F3IV":
// 			StellarDataMap["F3IV"] = append(StellarDataMap["F3IV"], "F")
// 			StellarDataMap["F3IV"] = append(StellarDataMap["F3IV"], "3")
// 			StellarDataMap["F3IV"] = append(StellarDataMap["F3IV"], "IV")

// 		case "F4IV":
// 			StellarDataMap["F4IV"] = append(StellarDataMap["F4IV"], "F")
// 			StellarDataMap["F4IV"] = append(StellarDataMap["F4IV"], "4")
// 			StellarDataMap["F4IV"] = append(StellarDataMap["F4IV"], "IV")

// 		case "F5IV":
// 			StellarDataMap["F5IV"] = append(StellarDataMap["F5IV"], "F")
// 			StellarDataMap["F5IV"] = append(StellarDataMap["F5IV"], "5")
// 			StellarDataMap["F5IV"] = append(StellarDataMap["F5IV"], "IV")

// 		case "F6IV":
// 			StellarDataMap["F6IV"] = append(StellarDataMap["F6IV"], "F")
// 			StellarDataMap["F6IV"] = append(StellarDataMap["F6IV"], "6")
// 			StellarDataMap["F6IV"] = append(StellarDataMap["F6IV"], "IV")

// 		case "F7IV":
// 			StellarDataMap["F7IV"] = append(StellarDataMap["F7IV"], "F")
// 			StellarDataMap["F7IV"] = append(StellarDataMap["F7IV"], "7")
// 			StellarDataMap["F7IV"] = append(StellarDataMap["F7IV"], "IV")

// 		case "F8IV":
// 			StellarDataMap["F8IV"] = append(StellarDataMap["F8IV"], "F")
// 			StellarDataMap["F8IV"] = append(StellarDataMap["F8IV"], "8")
// 			StellarDataMap["F8IV"] = append(StellarDataMap["F8IV"], "IV")

// 		case "F9IV":
// 			StellarDataMap["F9IV"] = append(StellarDataMap["F9IV"], "F")
// 			StellarDataMap["F9IV"] = append(StellarDataMap["F9IV"], "9")
// 			StellarDataMap["F9IV"] = append(StellarDataMap["F9IV"], "IV")

// 		case "F0V":
// 			StellarDataMap["F0V"] = append(StellarDataMap["F0V"], "F")
// 			StellarDataMap["F0V"] = append(StellarDataMap["F0V"], "0")
// 			StellarDataMap["F0V"] = append(StellarDataMap["F0V"], "V")

// 		case "F1V":
// 			StellarDataMap["F1V"] = append(StellarDataMap["F1V"], "F")
// 			StellarDataMap["F1V"] = append(StellarDataMap["F1V"], "1")
// 			StellarDataMap["F1V"] = append(StellarDataMap["F1V"], "V")

// 		case "F2V":
// 			StellarDataMap["F2V"] = append(StellarDataMap["F2V"], "F")
// 			StellarDataMap["F2V"] = append(StellarDataMap["F2V"], "2")
// 			StellarDataMap["F2V"] = append(StellarDataMap["F2V"], "V")

// 		case "F3V":
// 			StellarDataMap["F3V"] = append(StellarDataMap["F3V"], "F")
// 			StellarDataMap["F3V"] = append(StellarDataMap["F3V"], "3")
// 			StellarDataMap["F3V"] = append(StellarDataMap["F3V"], "V")

// 		case "F4V":
// 			StellarDataMap["F4V"] = append(StellarDataMap["F4V"], "F")
// 			StellarDataMap["F4V"] = append(StellarDataMap["F4V"], "4")
// 			StellarDataMap["F4V"] = append(StellarDataMap["F4V"], "V")

// 		case "F5V":
// 			StellarDataMap["F5V"] = append(StellarDataMap["F5V"], "F")
// 			StellarDataMap["F5V"] = append(StellarDataMap["F5V"], "5")
// 			StellarDataMap["F5V"] = append(StellarDataMap["F5V"], "V")

// 		case "F6V":
// 			StellarDataMap["F6V"] = append(StellarDataMap["F6V"], "F")
// 			StellarDataMap["F6V"] = append(StellarDataMap["F6V"], "6")
// 			StellarDataMap["F6V"] = append(StellarDataMap["F6V"], "V")

// 		case "F7V":
// 			StellarDataMap["F7V"] = append(StellarDataMap["F7V"], "F")
// 			StellarDataMap["F7V"] = append(StellarDataMap["F7V"], "7")
// 			StellarDataMap["F7V"] = append(StellarDataMap["F7V"], "V")

// 		case "F8V":
// 			StellarDataMap["F8V"] = append(StellarDataMap["F8V"], "F")
// 			StellarDataMap["F8V"] = append(StellarDataMap["F8V"], "8")
// 			StellarDataMap["F8V"] = append(StellarDataMap["F8V"], "V")

// 		case "F9V":
// 			StellarDataMap["F9V"] = append(StellarDataMap["F9V"], "F")
// 			StellarDataMap["F9V"] = append(StellarDataMap["F9V"], "9")
// 			StellarDataMap["F9V"] = append(StellarDataMap["F9V"], "V")

// 		case "F0VI":
// 			StellarDataMap["F0VI"] = append(StellarDataMap["F0VI"], "F")
// 			StellarDataMap["F0VI"] = append(StellarDataMap["F0VI"], "0")
// 			StellarDataMap["F0VI"] = append(StellarDataMap["F0VI"], "VI")

// 		case "F1VI":
// 			StellarDataMap["F1VI"] = append(StellarDataMap["F1VI"], "F")
// 			StellarDataMap["F1VI"] = append(StellarDataMap["F1VI"], "1")
// 			StellarDataMap["F1VI"] = append(StellarDataMap["F1VI"], "VI")

// 		case "F2VI":
// 			StellarDataMap["F2VI"] = append(StellarDataMap["F2VI"], "F")
// 			StellarDataMap["F2VI"] = append(StellarDataMap["F2VI"], "2")
// 			StellarDataMap["F2VI"] = append(StellarDataMap["F2VI"], "VI")

// 		case "F3VI":
// 			StellarDataMap["F3VI"] = append(StellarDataMap["F3VI"], "F")
// 			StellarDataMap["F3VI"] = append(StellarDataMap["F3VI"], "3")
// 			StellarDataMap["F3VI"] = append(StellarDataMap["F3VI"], "VI")

// 		case "F4VI":
// 			StellarDataMap["F4VI"] = append(StellarDataMap["F4VI"], "F")
// 			StellarDataMap["F4VI"] = append(StellarDataMap["F4VI"], "4")
// 			StellarDataMap["F4VI"] = append(StellarDataMap["F4VI"], "VI")

// 		case "F5VI":
// 			StellarDataMap["F5VI"] = append(StellarDataMap["F5VI"], "F")
// 			StellarDataMap["F5VI"] = append(StellarDataMap["F5VI"], "5")
// 			StellarDataMap["F5VI"] = append(StellarDataMap["F5VI"], "VI")

// 		case "G0Ia":
// 			StellarDataMap["G0Ia"] = append(StellarDataMap["G0Ia"], "G")
// 			StellarDataMap["G0Ia"] = append(StellarDataMap["G0Ia"], "0")
// 			StellarDataMap["G0Ia"] = append(StellarDataMap["G0Ia"], "Ia")

// 		case "G1Ia":
// 			StellarDataMap["G1Ia"] = append(StellarDataMap["G1Ia"], "G")
// 			StellarDataMap["G1Ia"] = append(StellarDataMap["G1Ia"], "1")
// 			StellarDataMap["G1Ia"] = append(StellarDataMap["G1Ia"], "Ia")

// 		case "G2Ia":
// 			StellarDataMap["G2Ia"] = append(StellarDataMap["G2Ia"], "G")
// 			StellarDataMap["G2Ia"] = append(StellarDataMap["G2Ia"], "2")
// 			StellarDataMap["G2Ia"] = append(StellarDataMap["G2Ia"], "Ia")

// 		case "G3Ia":
// 			StellarDataMap["G3Ia"] = append(StellarDataMap["G3Ia"], "G")
// 			StellarDataMap["G3Ia"] = append(StellarDataMap["G3Ia"], "3")
// 			StellarDataMap["G3Ia"] = append(StellarDataMap["G3Ia"], "Ia")

// 		case "G4Ia":
// 			StellarDataMap["G4Ia"] = append(StellarDataMap["G4Ia"], "G")
// 			StellarDataMap["G4Ia"] = append(StellarDataMap["G4Ia"], "4")
// 			StellarDataMap["G4Ia"] = append(StellarDataMap["G4Ia"], "Ia")

// 		case "G5Ia":
// 			StellarDataMap["G5Ia"] = append(StellarDataMap["G5Ia"], "G")
// 			StellarDataMap["G5Ia"] = append(StellarDataMap["G5Ia"], "5")
// 			StellarDataMap["G5Ia"] = append(StellarDataMap["G5Ia"], "Ia")

// 		case "G6Ia":
// 			StellarDataMap["G6Ia"] = append(StellarDataMap["G6Ia"], "G")
// 			StellarDataMap["G6Ia"] = append(StellarDataMap["G6Ia"], "6")
// 			StellarDataMap["G6Ia"] = append(StellarDataMap["G6Ia"], "Ia")

// 		case "G7Ia":
// 			StellarDataMap["G7Ia"] = append(StellarDataMap["G7Ia"], "G")
// 			StellarDataMap["G7Ia"] = append(StellarDataMap["G7Ia"], "7")
// 			StellarDataMap["G7Ia"] = append(StellarDataMap["G7Ia"], "Ia")

// 		case "G8Ia":
// 			StellarDataMap["G8Ia"] = append(StellarDataMap["G8Ia"], "G")
// 			StellarDataMap["G8Ia"] = append(StellarDataMap["G8Ia"], "8")
// 			StellarDataMap["G8Ia"] = append(StellarDataMap["G8Ia"], "Ia")

// 		case "G9Ia":
// 			StellarDataMap["G9Ia"] = append(StellarDataMap["G9Ia"], "G")
// 			StellarDataMap["G9Ia"] = append(StellarDataMap["G9Ia"], "9")
// 			StellarDataMap["G9Ia"] = append(StellarDataMap["G9Ia"], "Ia")

// 		case "G0Ib":
// 			StellarDataMap["G0Ib"] = append(StellarDataMap["G0Ib"], "G")
// 			StellarDataMap["G0Ib"] = append(StellarDataMap["G0Ib"], "0")
// 			StellarDataMap["G0Ib"] = append(StellarDataMap["G0Ib"], "Ib")

// 		case "G1Ib":
// 			StellarDataMap["G1Ib"] = append(StellarDataMap["G1Ib"], "G")
// 			StellarDataMap["G1Ib"] = append(StellarDataMap["G1Ib"], "1")
// 			StellarDataMap["G1Ib"] = append(StellarDataMap["G1Ib"], "Ib")

// 		case "G2Ib":
// 			StellarDataMap["G2Ib"] = append(StellarDataMap["G2Ib"], "G")
// 			StellarDataMap["G2Ib"] = append(StellarDataMap["G2Ib"], "2")
// 			StellarDataMap["G2Ib"] = append(StellarDataMap["G2Ib"], "Ib")

// 		case "G3Ib":
// 			StellarDataMap["G3Ib"] = append(StellarDataMap["G3Ib"], "G")
// 			StellarDataMap["G3Ib"] = append(StellarDataMap["G3Ib"], "3")
// 			StellarDataMap["G3Ib"] = append(StellarDataMap["G3Ib"], "Ib")

// 		case "G4Ib":
// 			StellarDataMap["G4Ib"] = append(StellarDataMap["G4Ib"], "G")
// 			StellarDataMap["G4Ib"] = append(StellarDataMap["G4Ib"], "4")
// 			StellarDataMap["G4Ib"] = append(StellarDataMap["G4Ib"], "Ib")

// 		case "G5Ib":
// 			StellarDataMap["G5Ib"] = append(StellarDataMap["G5Ib"], "G")
// 			StellarDataMap["G5Ib"] = append(StellarDataMap["G5Ib"], "5")
// 			StellarDataMap["G5Ib"] = append(StellarDataMap["G5Ib"], "Ib")

// 		case "G6Ib":
// 			StellarDataMap["G6Ib"] = append(StellarDataMap["G6Ib"], "G")
// 			StellarDataMap["G6Ib"] = append(StellarDataMap["G6Ib"], "6")
// 			StellarDataMap["G6Ib"] = append(StellarDataMap["G6Ib"], "Ib")

// 		case "G7Ib":
// 			StellarDataMap["G7Ib"] = append(StellarDataMap["G7Ib"], "G")
// 			StellarDataMap["G7Ib"] = append(StellarDataMap["G7Ib"], "7")
// 			StellarDataMap["G7Ib"] = append(StellarDataMap["G7Ib"], "Ib")

// 		case "G8Ib":
// 			StellarDataMap["G8Ib"] = append(StellarDataMap["G8Ib"], "G")
// 			StellarDataMap["G8Ib"] = append(StellarDataMap["G8Ib"], "8")
// 			StellarDataMap["G8Ib"] = append(StellarDataMap["G8Ib"], "Ib")

// 		case "G9Ib":
// 			StellarDataMap["G9Ib"] = append(StellarDataMap["G9Ib"], "G")
// 			StellarDataMap["G9Ib"] = append(StellarDataMap["G9Ib"], "9")
// 			StellarDataMap["G9Ib"] = append(StellarDataMap["G9Ib"], "Ib")

// 		case "G0II":
// 			StellarDataMap["G0II"] = append(StellarDataMap["G0II"], "G")
// 			StellarDataMap["G0II"] = append(StellarDataMap["G0II"], "0")
// 			StellarDataMap["G0II"] = append(StellarDataMap["G0II"], "II")

// 		case "G1II":
// 			StellarDataMap["G1II"] = append(StellarDataMap["G1II"], "G")
// 			StellarDataMap["G1II"] = append(StellarDataMap["G1II"], "1")
// 			StellarDataMap["G1II"] = append(StellarDataMap["G1II"], "II")

// 		case "G2II":
// 			StellarDataMap["G2II"] = append(StellarDataMap["G2II"], "G")
// 			StellarDataMap["G2II"] = append(StellarDataMap["G2II"], "2")
// 			StellarDataMap["G2II"] = append(StellarDataMap["G2II"], "II")

// 		case "G3II":
// 			StellarDataMap["G3II"] = append(StellarDataMap["G3II"], "G")
// 			StellarDataMap["G3II"] = append(StellarDataMap["G3II"], "3")
// 			StellarDataMap["G3II"] = append(StellarDataMap["G3II"], "II")

// 		case "G4II":
// 			StellarDataMap["G4II"] = append(StellarDataMap["G4II"], "G")
// 			StellarDataMap["G4II"] = append(StellarDataMap["G4II"], "4")
// 			StellarDataMap["G4II"] = append(StellarDataMap["G4II"], "II")

// 		case "G5II":
// 			StellarDataMap["G5II"] = append(StellarDataMap["G5II"], "G")
// 			StellarDataMap["G5II"] = append(StellarDataMap["G5II"], "5")
// 			StellarDataMap["G5II"] = append(StellarDataMap["G5II"], "II")

// 		case "G6II":
// 			StellarDataMap["G6II"] = append(StellarDataMap["G6II"], "G")
// 			StellarDataMap["G6II"] = append(StellarDataMap["G6II"], "6")
// 			StellarDataMap["G6II"] = append(StellarDataMap["G6II"], "II")

// 		case "G7II":
// 			StellarDataMap["G7II"] = append(StellarDataMap["G7II"], "G")
// 			StellarDataMap["G7II"] = append(StellarDataMap["G7II"], "7")
// 			StellarDataMap["G7II"] = append(StellarDataMap["G7II"], "II")

// 		case "G8II":
// 			StellarDataMap["G8II"] = append(StellarDataMap["G8II"], "G")
// 			StellarDataMap["G8II"] = append(StellarDataMap["G8II"], "8")
// 			StellarDataMap["G8II"] = append(StellarDataMap["G8II"], "II")

// 		case "G9II":
// 			StellarDataMap["G9II"] = append(StellarDataMap["G9II"], "G")
// 			StellarDataMap["G9II"] = append(StellarDataMap["G9II"], "9")
// 			StellarDataMap["G9II"] = append(StellarDataMap["G9II"], "II")

// 		case "G0III":
// 			StellarDataMap["G0III"] = append(StellarDataMap["G0III"], "G")
// 			StellarDataMap["G0III"] = append(StellarDataMap["G0III"], "0")
// 			StellarDataMap["G0III"] = append(StellarDataMap["G0III"], "III")

// 		case "G1III":
// 			StellarDataMap["G1III"] = append(StellarDataMap["G1III"], "G")
// 			StellarDataMap["G1III"] = append(StellarDataMap["G1III"], "1")
// 			StellarDataMap["G1III"] = append(StellarDataMap["G1III"], "III")

// 		case "G2III":
// 			StellarDataMap["G2III"] = append(StellarDataMap["G2III"], "G")
// 			StellarDataMap["G2III"] = append(StellarDataMap["G2III"], "2")
// 			StellarDataMap["G2III"] = append(StellarDataMap["G2III"], "III")

// 		case "G3III":
// 			StellarDataMap["G3III"] = append(StellarDataMap["G3III"], "G")
// 			StellarDataMap["G3III"] = append(StellarDataMap["G3III"], "3")
// 			StellarDataMap["G3III"] = append(StellarDataMap["G3III"], "III")

// 		case "G4III":
// 			StellarDataMap["G4III"] = append(StellarDataMap["G4III"], "G")
// 			StellarDataMap["G4III"] = append(StellarDataMap["G4III"], "4")
// 			StellarDataMap["G4III"] = append(StellarDataMap["G4III"], "III")

// 		case "G5III":
// 			StellarDataMap["G5III"] = append(StellarDataMap["G5III"], "G")
// 			StellarDataMap["G5III"] = append(StellarDataMap["G5III"], "5")
// 			StellarDataMap["G5III"] = append(StellarDataMap["G5III"], "III")

// 		case "G6III":
// 			StellarDataMap["G6III"] = append(StellarDataMap["G6III"], "G")
// 			StellarDataMap["G6III"] = append(StellarDataMap["G6III"], "6")
// 			StellarDataMap["G6III"] = append(StellarDataMap["G6III"], "III")

// 		case "G7III":
// 			StellarDataMap["G7III"] = append(StellarDataMap["G7III"], "G")
// 			StellarDataMap["G7III"] = append(StellarDataMap["G7III"], "7")
// 			StellarDataMap["G7III"] = append(StellarDataMap["G7III"], "III")

// 		case "G8III":
// 			StellarDataMap["G8III"] = append(StellarDataMap["G8III"], "G")
// 			StellarDataMap["G8III"] = append(StellarDataMap["G8III"], "8")
// 			StellarDataMap["G8III"] = append(StellarDataMap["G8III"], "III")

// 		case "G9III":
// 			StellarDataMap["G9III"] = append(StellarDataMap["G9III"], "G")
// 			StellarDataMap["G9III"] = append(StellarDataMap["G9III"], "9")
// 			StellarDataMap["G9III"] = append(StellarDataMap["G9III"], "III")

// 		case "G0IV":
// 			StellarDataMap["G0IV"] = append(StellarDataMap["G0IV"], "G")
// 			StellarDataMap["G0IV"] = append(StellarDataMap["G0IV"], "0")
// 			StellarDataMap["G0IV"] = append(StellarDataMap["G0IV"], "IV")

// 		case "G1IV":
// 			StellarDataMap["G1IV"] = append(StellarDataMap["G1IV"], "G")
// 			StellarDataMap["G1IV"] = append(StellarDataMap["G1IV"], "1")
// 			StellarDataMap["G1IV"] = append(StellarDataMap["G1IV"], "IV")

// 		case "G2IV":
// 			StellarDataMap["G2IV"] = append(StellarDataMap["G2IV"], "G")
// 			StellarDataMap["G2IV"] = append(StellarDataMap["G2IV"], "2")
// 			StellarDataMap["G2IV"] = append(StellarDataMap["G2IV"], "IV")

// 		case "G3IV":
// 			StellarDataMap["G3IV"] = append(StellarDataMap["G3IV"], "G")
// 			StellarDataMap["G3IV"] = append(StellarDataMap["G3IV"], "3")
// 			StellarDataMap["G3IV"] = append(StellarDataMap["G3IV"], "IV")

// 		case "G4IV":
// 			StellarDataMap["G4IV"] = append(StellarDataMap["G4IV"], "G")
// 			StellarDataMap["G4IV"] = append(StellarDataMap["G4IV"], "4")
// 			StellarDataMap["G4IV"] = append(StellarDataMap["G4IV"], "IV")

// 		case "G5IV":
// 			StellarDataMap["G5IV"] = append(StellarDataMap["G5IV"], "G")
// 			StellarDataMap["G5IV"] = append(StellarDataMap["G5IV"], "5")
// 			StellarDataMap["G5IV"] = append(StellarDataMap["G5IV"], "IV")

// 		case "G6IV":
// 			StellarDataMap["G6IV"] = append(StellarDataMap["G6IV"], "G")
// 			StellarDataMap["G6IV"] = append(StellarDataMap["G6IV"], "6")
// 			StellarDataMap["G6IV"] = append(StellarDataMap["G6IV"], "IV")

// 		case "G7IV":
// 			StellarDataMap["G7IV"] = append(StellarDataMap["G7IV"], "G")
// 			StellarDataMap["G7IV"] = append(StellarDataMap["G7IV"], "7")
// 			StellarDataMap["G7IV"] = append(StellarDataMap["G7IV"], "IV")

// 		case "G8IV":
// 			StellarDataMap["G8IV"] = append(StellarDataMap["G8IV"], "G")
// 			StellarDataMap["G8IV"] = append(StellarDataMap["G8IV"], "8")
// 			StellarDataMap["G8IV"] = append(StellarDataMap["G8IV"], "IV")

// 		case "G9IV":
// 			StellarDataMap["G9IV"] = append(StellarDataMap["G9IV"], "G")
// 			StellarDataMap["G9IV"] = append(StellarDataMap["G9IV"], "9")
// 			StellarDataMap["G9IV"] = append(StellarDataMap["G9IV"], "IV")

// 		case "G0V":
// 			StellarDataMap["G0V"] = append(StellarDataMap["G0V"], "G")
// 			StellarDataMap["G0V"] = append(StellarDataMap["G0V"], "0")
// 			StellarDataMap["G0V"] = append(StellarDataMap["G0V"], "V")

// 		case "G1V":
// 			StellarDataMap["G1V"] = append(StellarDataMap["G1V"], "G")
// 			StellarDataMap["G1V"] = append(StellarDataMap["G1V"], "1")
// 			StellarDataMap["G1V"] = append(StellarDataMap["G1V"], "V")

// 		case "G2V":
// 			StellarDataMap["G2V"] = append(StellarDataMap["G2V"], "G")
// 			StellarDataMap["G2V"] = append(StellarDataMap["G2V"], "2")
// 			StellarDataMap["G2V"] = append(StellarDataMap["G2V"], "V")

// 		case "G3V":
// 			StellarDataMap["G3V"] = append(StellarDataMap["G3V"], "G")
// 			StellarDataMap["G3V"] = append(StellarDataMap["G3V"], "3")
// 			StellarDataMap["G3V"] = append(StellarDataMap["G3V"], "V")

// 		case "G4V":
// 			StellarDataMap["G4V"] = append(StellarDataMap["G4V"], "G")
// 			StellarDataMap["G4V"] = append(StellarDataMap["G4V"], "4")
// 			StellarDataMap["G4V"] = append(StellarDataMap["G4V"], "V")

// 		case "G5V":
// 			StellarDataMap["G5V"] = append(StellarDataMap["G5V"], "G")
// 			StellarDataMap["G5V"] = append(StellarDataMap["G5V"], "5")
// 			StellarDataMap["G5V"] = append(StellarDataMap["G5V"], "V")

// 		case "G6V":
// 			StellarDataMap["G6V"] = append(StellarDataMap["G6V"], "G")
// 			StellarDataMap["G6V"] = append(StellarDataMap["G6V"], "6")
// 			StellarDataMap["G6V"] = append(StellarDataMap["G6V"], "V")

// 		case "G7V":
// 			StellarDataMap["G7V"] = append(StellarDataMap["G7V"], "G")
// 			StellarDataMap["G7V"] = append(StellarDataMap["G7V"], "7")
// 			StellarDataMap["G7V"] = append(StellarDataMap["G7V"], "V")

// 		case "G8V":
// 			StellarDataMap["G8V"] = append(StellarDataMap["G8V"], "G")
// 			StellarDataMap["G8V"] = append(StellarDataMap["G8V"], "8")
// 			StellarDataMap["G8V"] = append(StellarDataMap["G8V"], "V")

// 		case "G9V":
// 			StellarDataMap["G9V"] = append(StellarDataMap["G9V"], "G")
// 			StellarDataMap["G9V"] = append(StellarDataMap["G9V"], "9")
// 			StellarDataMap["G9V"] = append(StellarDataMap["G9V"], "V")

// 		case "G0VI":
// 			StellarDataMap["G0VI"] = append(StellarDataMap["G0VI"], "G")
// 			StellarDataMap["G0VI"] = append(StellarDataMap["G0VI"], "0")
// 			StellarDataMap["G0VI"] = append(StellarDataMap["G0VI"], "VI")

// 		case "G1VI":
// 			StellarDataMap["G1VI"] = append(StellarDataMap["G1VI"], "G")
// 			StellarDataMap["G1VI"] = append(StellarDataMap["G1VI"], "1")
// 			StellarDataMap["G1VI"] = append(StellarDataMap["G1VI"], "VI")

// 		case "G2VI":
// 			StellarDataMap["G2VI"] = append(StellarDataMap["G2VI"], "G")
// 			StellarDataMap["G2VI"] = append(StellarDataMap["G2VI"], "2")
// 			StellarDataMap["G2VI"] = append(StellarDataMap["G2VI"], "VI")

// 		case "G3VI":
// 			StellarDataMap["G3VI"] = append(StellarDataMap["G3VI"], "G")
// 			StellarDataMap["G3VI"] = append(StellarDataMap["G3VI"], "3")
// 			StellarDataMap["G3VI"] = append(StellarDataMap["G3VI"], "VI")

// 		case "G4VI":
// 			StellarDataMap["G4VI"] = append(StellarDataMap["G4VI"], "G")
// 			StellarDataMap["G4VI"] = append(StellarDataMap["G4VI"], "4")
// 			StellarDataMap["G4VI"] = append(StellarDataMap["G4VI"], "VI")

// 		case "G5VI":
// 			StellarDataMap["G5VI"] = append(StellarDataMap["G5VI"], "G")
// 			StellarDataMap["G5VI"] = append(StellarDataMap["G5VI"], "5")
// 			StellarDataMap["G5VI"] = append(StellarDataMap["G5VI"], "VI")

// 		case "G6VI":
// 			StellarDataMap["G6VI"] = append(StellarDataMap["G6VI"], "G")
// 			StellarDataMap["G6VI"] = append(StellarDataMap["G6VI"], "6")
// 			StellarDataMap["G6VI"] = append(StellarDataMap["G6VI"], "VI")

// 		case "G7VI":
// 			StellarDataMap["G7VI"] = append(StellarDataMap["G7VI"], "G")
// 			StellarDataMap["G7VI"] = append(StellarDataMap["G7VI"], "7")
// 			StellarDataMap["G7VI"] = append(StellarDataMap["G7VI"], "VI")

// 		case "G8VI":
// 			StellarDataMap["G8VI"] = append(StellarDataMap["G8VI"], "G")
// 			StellarDataMap["G8VI"] = append(StellarDataMap["G8VI"], "8")
// 			StellarDataMap["G8VI"] = append(StellarDataMap["G8VI"], "VI")

// 		case "G9VI":
// 			StellarDataMap["G9VI"] = append(StellarDataMap["G9VI"], "G")
// 			StellarDataMap["G9VI"] = append(StellarDataMap["G9VI"], "9")
// 			StellarDataMap["G9VI"] = append(StellarDataMap["G9VI"], "VI")

// 		case "K0Ia":
// 			StellarDataMap["K0Ia"] = append(StellarDataMap["K0Ia"], "K")
// 			StellarDataMap["K0Ia"] = append(StellarDataMap["K0Ia"], "0")
// 			StellarDataMap["K0Ia"] = append(StellarDataMap["K0Ia"], "Ia")

// 		case "K1Ia":
// 			StellarDataMap["K1Ia"] = append(StellarDataMap["K1Ia"], "K")
// 			StellarDataMap["K1Ia"] = append(StellarDataMap["K1Ia"], "1")
// 			StellarDataMap["K1Ia"] = append(StellarDataMap["K1Ia"], "Ia")

// 		case "K2Ia":
// 			StellarDataMap["K2Ia"] = append(StellarDataMap["K2Ia"], "K")
// 			StellarDataMap["K2Ia"] = append(StellarDataMap["K2Ia"], "2")
// 			StellarDataMap["K2Ia"] = append(StellarDataMap["K2Ia"], "Ia")

// 		case "K3Ia":
// 			StellarDataMap["K3Ia"] = append(StellarDataMap["K3Ia"], "K")
// 			StellarDataMap["K3Ia"] = append(StellarDataMap["K3Ia"], "3")
// 			StellarDataMap["K3Ia"] = append(StellarDataMap["K3Ia"], "Ia")

// 		case "K4Ia":
// 			StellarDataMap["K4Ia"] = append(StellarDataMap["K4Ia"], "K")
// 			StellarDataMap["K4Ia"] = append(StellarDataMap["K4Ia"], "4")
// 			StellarDataMap["K4Ia"] = append(StellarDataMap["K4Ia"], "Ia")

// 		case "K5Ia":
// 			StellarDataMap["K5Ia"] = append(StellarDataMap["K5Ia"], "K")
// 			StellarDataMap["K5Ia"] = append(StellarDataMap["K5Ia"], "5")
// 			StellarDataMap["K5Ia"] = append(StellarDataMap["K5Ia"], "Ia")

// 		case "K6Ia":
// 			StellarDataMap["K6Ia"] = append(StellarDataMap["K6Ia"], "K")
// 			StellarDataMap["K6Ia"] = append(StellarDataMap["K6Ia"], "6")
// 			StellarDataMap["K6Ia"] = append(StellarDataMap["K6Ia"], "Ia")

// 		case "K7Ia":
// 			StellarDataMap["K7Ia"] = append(StellarDataMap["K7Ia"], "K")
// 			StellarDataMap["K7Ia"] = append(StellarDataMap["K7Ia"], "7")
// 			StellarDataMap["K7Ia"] = append(StellarDataMap["K7Ia"], "Ia")

// 		case "K8Ia":
// 			StellarDataMap["K8Ia"] = append(StellarDataMap["K8Ia"], "K")
// 			StellarDataMap["K8Ia"] = append(StellarDataMap["K8Ia"], "8")
// 			StellarDataMap["K8Ia"] = append(StellarDataMap["K8Ia"], "Ia")

// 		case "K9Ia":
// 			StellarDataMap["K9Ia"] = append(StellarDataMap["K9Ia"], "K")
// 			StellarDataMap["K9Ia"] = append(StellarDataMap["K9Ia"], "9")
// 			StellarDataMap["K9Ia"] = append(StellarDataMap["K9Ia"], "Ia")

// 		case "K0Ib":
// 			StellarDataMap["K0Ib"] = append(StellarDataMap["K0Ib"], "K")
// 			StellarDataMap["K0Ib"] = append(StellarDataMap["K0Ib"], "0")
// 			StellarDataMap["K0Ib"] = append(StellarDataMap["K0Ib"], "Ib")

// 		case "K1Ib":
// 			StellarDataMap["K1Ib"] = append(StellarDataMap["K1Ib"], "K")
// 			StellarDataMap["K1Ib"] = append(StellarDataMap["K1Ib"], "1")
// 			StellarDataMap["K1Ib"] = append(StellarDataMap["K1Ib"], "Ib")

// 		case "K2Ib":
// 			StellarDataMap["K2Ib"] = append(StellarDataMap["K2Ib"], "K")
// 			StellarDataMap["K2Ib"] = append(StellarDataMap["K2Ib"], "2")
// 			StellarDataMap["K2Ib"] = append(StellarDataMap["K2Ib"], "Ib")

// 		case "K3Ib":
// 			StellarDataMap["K3Ib"] = append(StellarDataMap["K3Ib"], "K")
// 			StellarDataMap["K3Ib"] = append(StellarDataMap["K3Ib"], "3")
// 			StellarDataMap["K3Ib"] = append(StellarDataMap["K3Ib"], "Ib")

// 		case "K4Ib":
// 			StellarDataMap["K4Ib"] = append(StellarDataMap["K4Ib"], "K")
// 			StellarDataMap["K4Ib"] = append(StellarDataMap["K4Ib"], "4")
// 			StellarDataMap["K4Ib"] = append(StellarDataMap["K4Ib"], "Ib")

// 		case "K5Ib":
// 			StellarDataMap["K5Ib"] = append(StellarDataMap["K5Ib"], "K")
// 			StellarDataMap["K5Ib"] = append(StellarDataMap["K5Ib"], "5")
// 			StellarDataMap["K5Ib"] = append(StellarDataMap["K5Ib"], "Ib")

// 		case "K6Ib":
// 			StellarDataMap["K6Ib"] = append(StellarDataMap["K6Ib"], "K")
// 			StellarDataMap["K6Ib"] = append(StellarDataMap["K6Ib"], "6")
// 			StellarDataMap["K6Ib"] = append(StellarDataMap["K6Ib"], "Ib")

// 		case "K7Ib":
// 			StellarDataMap["K7Ib"] = append(StellarDataMap["K7Ib"], "K")
// 			StellarDataMap["K7Ib"] = append(StellarDataMap["K7Ib"], "7")
// 			StellarDataMap["K7Ib"] = append(StellarDataMap["K7Ib"], "Ib")

// 		case "K8Ib":
// 			StellarDataMap["K8Ib"] = append(StellarDataMap["K8Ib"], "K")
// 			StellarDataMap["K8Ib"] = append(StellarDataMap["K8Ib"], "8")
// 			StellarDataMap["K8Ib"] = append(StellarDataMap["K8Ib"], "Ib")

// 		case "K9Ib":
// 			StellarDataMap["K9Ib"] = append(StellarDataMap["K9Ib"], "K")
// 			StellarDataMap["K9Ib"] = append(StellarDataMap["K9Ib"], "9")
// 			StellarDataMap["K9Ib"] = append(StellarDataMap["K9Ib"], "Ib")

// 		case "K0II":
// 			StellarDataMap["K0II"] = append(StellarDataMap["K0II"], "K")
// 			StellarDataMap["K0II"] = append(StellarDataMap["K0II"], "0")
// 			StellarDataMap["K0II"] = append(StellarDataMap["K0II"], "II")

// 		case "K1II":
// 			StellarDataMap["K1II"] = append(StellarDataMap["K1II"], "K")
// 			StellarDataMap["K1II"] = append(StellarDataMap["K1II"], "1")
// 			StellarDataMap["K1II"] = append(StellarDataMap["K1II"], "II")

// 		case "K2II":
// 			StellarDataMap["K2II"] = append(StellarDataMap["K2II"], "K")
// 			StellarDataMap["K2II"] = append(StellarDataMap["K2II"], "2")
// 			StellarDataMap["K2II"] = append(StellarDataMap["K2II"], "II")

// 		case "K3II":
// 			StellarDataMap["K3II"] = append(StellarDataMap["K3II"], "K")
// 			StellarDataMap["K3II"] = append(StellarDataMap["K3II"], "3")
// 			StellarDataMap["K3II"] = append(StellarDataMap["K3II"], "II")

// 		case "K4II":
// 			StellarDataMap["K4II"] = append(StellarDataMap["K4II"], "K")
// 			StellarDataMap["K4II"] = append(StellarDataMap["K4II"], "4")
// 			StellarDataMap["K4II"] = append(StellarDataMap["K4II"], "II")

// 		case "K5II":
// 			StellarDataMap["K5II"] = append(StellarDataMap["K5II"], "K")
// 			StellarDataMap["K5II"] = append(StellarDataMap["K5II"], "5")
// 			StellarDataMap["K5II"] = append(StellarDataMap["K5II"], "II")

// 		case "K6II":
// 			StellarDataMap["K6II"] = append(StellarDataMap["K6II"], "K")
// 			StellarDataMap["K6II"] = append(StellarDataMap["K6II"], "6")
// 			StellarDataMap["K6II"] = append(StellarDataMap["K6II"], "II")

// 		case "K7II":
// 			StellarDataMap["K7II"] = append(StellarDataMap["K7II"], "K")
// 			StellarDataMap["K7II"] = append(StellarDataMap["K7II"], "7")
// 			StellarDataMap["K7II"] = append(StellarDataMap["K7II"], "II")

// 		case "K8II":
// 			StellarDataMap["K8II"] = append(StellarDataMap["K8II"], "K")
// 			StellarDataMap["K8II"] = append(StellarDataMap["K8II"], "8")
// 			StellarDataMap["K8II"] = append(StellarDataMap["K8II"], "II")

// 		case "K9II":
// 			StellarDataMap["K9II"] = append(StellarDataMap["K9II"], "K")
// 			StellarDataMap["K9II"] = append(StellarDataMap["K9II"], "9")
// 			StellarDataMap["K9II"] = append(StellarDataMap["K9II"], "II")

// 		case "K0III":
// 			StellarDataMap["K0III"] = append(StellarDataMap["K0III"], "K")
// 			StellarDataMap["K0III"] = append(StellarDataMap["K0III"], "0")
// 			StellarDataMap["K0III"] = append(StellarDataMap["K0III"], "III")

// 		case "K1III":
// 			StellarDataMap["K1III"] = append(StellarDataMap["K1III"], "K")
// 			StellarDataMap["K1III"] = append(StellarDataMap["K1III"], "1")
// 			StellarDataMap["K1III"] = append(StellarDataMap["K1III"], "III")

// 		case "K2III":
// 			StellarDataMap["K2III"] = append(StellarDataMap["K2III"], "K")
// 			StellarDataMap["K2III"] = append(StellarDataMap["K2III"], "2")
// 			StellarDataMap["K2III"] = append(StellarDataMap["K2III"], "III")

// 		case "K3III":
// 			StellarDataMap["K3III"] = append(StellarDataMap["K3III"], "K")
// 			StellarDataMap["K3III"] = append(StellarDataMap["K3III"], "3")
// 			StellarDataMap["K3III"] = append(StellarDataMap["K3III"], "III")

// 		case "K4III":
// 			StellarDataMap["K4III"] = append(StellarDataMap["K4III"], "K")
// 			StellarDataMap["K4III"] = append(StellarDataMap["K4III"], "4")
// 			StellarDataMap["K4III"] = append(StellarDataMap["K4III"], "III")

// 		case "K5III":
// 			StellarDataMap["K5III"] = append(StellarDataMap["K5III"], "K")
// 			StellarDataMap["K5III"] = append(StellarDataMap["K5III"], "5")
// 			StellarDataMap["K5III"] = append(StellarDataMap["K5III"], "III")

// 		case "K6III":
// 			StellarDataMap["K6III"] = append(StellarDataMap["K6III"], "K")
// 			StellarDataMap["K6III"] = append(StellarDataMap["K6III"], "6")
// 			StellarDataMap["K6III"] = append(StellarDataMap["K6III"], "III")

// 		case "K7III":
// 			StellarDataMap["K7III"] = append(StellarDataMap["K7III"], "K")
// 			StellarDataMap["K7III"] = append(StellarDataMap["K7III"], "7")
// 			StellarDataMap["K7III"] = append(StellarDataMap["K7III"], "III")

// 		case "K8III":
// 			StellarDataMap["K8III"] = append(StellarDataMap["K8III"], "K")
// 			StellarDataMap["K8III"] = append(StellarDataMap["K8III"], "8")
// 			StellarDataMap["K8III"] = append(StellarDataMap["K8III"], "III")

// 		case "K9III":
// 			StellarDataMap["K9III"] = append(StellarDataMap["K9III"], "K")
// 			StellarDataMap["K9III"] = append(StellarDataMap["K9III"], "9")
// 			StellarDataMap["K9III"] = append(StellarDataMap["K9III"], "III")

// 		case "K0IV":
// 			StellarDataMap["K0IV"] = append(StellarDataMap["K0IV"], "K")
// 			StellarDataMap["K0IV"] = append(StellarDataMap["K0IV"], "0")
// 			StellarDataMap["K0IV"] = append(StellarDataMap["K0IV"], "IV")

// 		case "K1IV":
// 			StellarDataMap["K1IV"] = append(StellarDataMap["K1IV"], "K")
// 			StellarDataMap["K1IV"] = append(StellarDataMap["K1IV"], "1")
// 			StellarDataMap["K1IV"] = append(StellarDataMap["K1IV"], "IV")

// 		case "K2IV":
// 			StellarDataMap["K2IV"] = append(StellarDataMap["K2IV"], "K")
// 			StellarDataMap["K2IV"] = append(StellarDataMap["K2IV"], "2")
// 			StellarDataMap["K2IV"] = append(StellarDataMap["K2IV"], "IV")

// 		case "K3IV":
// 			StellarDataMap["K3IV"] = append(StellarDataMap["K3IV"], "K")
// 			StellarDataMap["K3IV"] = append(StellarDataMap["K3IV"], "3")
// 			StellarDataMap["K3IV"] = append(StellarDataMap["K3IV"], "IV")

// 		case "K4IV":
// 			StellarDataMap["K4IV"] = append(StellarDataMap["K4IV"], "K")
// 			StellarDataMap["K4IV"] = append(StellarDataMap["K4IV"], "4")
// 			StellarDataMap["K4IV"] = append(StellarDataMap["K4IV"], "IV")

// 		case "K5IV":
// 			StellarDataMap["K5IV"] = append(StellarDataMap["K5IV"], "K")
// 			StellarDataMap["K5IV"] = append(StellarDataMap["K5IV"], "5")
// 			StellarDataMap["K5IV"] = append(StellarDataMap["K5IV"], "IV")

// 		case "K0V":
// 			StellarDataMap["K0V"] = append(StellarDataMap["K0V"], "K")
// 			StellarDataMap["K0V"] = append(StellarDataMap["K0V"], "0")
// 			StellarDataMap["K0V"] = append(StellarDataMap["K0V"], "V")

// 		case "K1V":
// 			StellarDataMap["K1V"] = append(StellarDataMap["K1V"], "K")
// 			StellarDataMap["K1V"] = append(StellarDataMap["K1V"], "1")
// 			StellarDataMap["K1V"] = append(StellarDataMap["K1V"], "V")

// 		case "K2V":
// 			StellarDataMap["K2V"] = append(StellarDataMap["K2V"], "K")
// 			StellarDataMap["K2V"] = append(StellarDataMap["K2V"], "2")
// 			StellarDataMap["K2V"] = append(StellarDataMap["K2V"], "V")

// 		case "K3V":
// 			StellarDataMap["K3V"] = append(StellarDataMap["K3V"], "K")
// 			StellarDataMap["K3V"] = append(StellarDataMap["K3V"], "3")
// 			StellarDataMap["K3V"] = append(StellarDataMap["K3V"], "V")

// 		case "K4V":
// 			StellarDataMap["K4V"] = append(StellarDataMap["K4V"], "K")
// 			StellarDataMap["K4V"] = append(StellarDataMap["K4V"], "4")
// 			StellarDataMap["K4V"] = append(StellarDataMap["K4V"], "V")

// 		case "K5V":
// 			StellarDataMap["K5V"] = append(StellarDataMap["K5V"], "K")
// 			StellarDataMap["K5V"] = append(StellarDataMap["K5V"], "5")
// 			StellarDataMap["K5V"] = append(StellarDataMap["K5V"], "V")

// 		case "K6V":
// 			StellarDataMap["K6V"] = append(StellarDataMap["K6V"], "K")
// 			StellarDataMap["K6V"] = append(StellarDataMap["K6V"], "6")
// 			StellarDataMap["K6V"] = append(StellarDataMap["K6V"], "V")

// 		case "K7V":
// 			StellarDataMap["K7V"] = append(StellarDataMap["K7V"], "K")
// 			StellarDataMap["K7V"] = append(StellarDataMap["K7V"], "7")
// 			StellarDataMap["K7V"] = append(StellarDataMap["K7V"], "V")

// 		case "K8V":
// 			StellarDataMap["K8V"] = append(StellarDataMap["K8V"], "K")
// 			StellarDataMap["K8V"] = append(StellarDataMap["K8V"], "8")
// 			StellarDataMap["K8V"] = append(StellarDataMap["K8V"], "V")

// 		case "K9V":
// 			StellarDataMap["K9V"] = append(StellarDataMap["K9V"], "K")
// 			StellarDataMap["K9V"] = append(StellarDataMap["K9V"], "9")
// 			StellarDataMap["K9V"] = append(StellarDataMap["K9V"], "V")

// 		case "K0VI":
// 			StellarDataMap["K0VI"] = append(StellarDataMap["K0VI"], "K")
// 			StellarDataMap["K0VI"] = append(StellarDataMap["K0VI"], "0")
// 			StellarDataMap["K0VI"] = append(StellarDataMap["K0VI"], "VI")

// 		case "K1VI":
// 			StellarDataMap["K1VI"] = append(StellarDataMap["K1VI"], "K")
// 			StellarDataMap["K1VI"] = append(StellarDataMap["K1VI"], "1")
// 			StellarDataMap["K1VI"] = append(StellarDataMap["K1VI"], "VI")

// 		case "K2VI":
// 			StellarDataMap["K2VI"] = append(StellarDataMap["K2VI"], "K")
// 			StellarDataMap["K2VI"] = append(StellarDataMap["K2VI"], "2")
// 			StellarDataMap["K2VI"] = append(StellarDataMap["K2VI"], "VI")

// 		case "K3VI":
// 			StellarDataMap["K3VI"] = append(StellarDataMap["K3VI"], "K")
// 			StellarDataMap["K3VI"] = append(StellarDataMap["K3VI"], "3")
// 			StellarDataMap["K3VI"] = append(StellarDataMap["K3VI"], "VI")

// 		case "K4VI":
// 			StellarDataMap["K4VI"] = append(StellarDataMap["K4VI"], "K")
// 			StellarDataMap["K4VI"] = append(StellarDataMap["K4VI"], "4")
// 			StellarDataMap["K4VI"] = append(StellarDataMap["K4VI"], "VI")

// 		case "K5VI":
// 			StellarDataMap["K5VI"] = append(StellarDataMap["K5VI"], "K")
// 			StellarDataMap["K5VI"] = append(StellarDataMap["K5VI"], "5")
// 			StellarDataMap["K5VI"] = append(StellarDataMap["K5VI"], "VI")

// 		case "K6VI":
// 			StellarDataMap["K6VI"] = append(StellarDataMap["K6VI"], "K")
// 			StellarDataMap["K6VI"] = append(StellarDataMap["K6VI"], "6")
// 			StellarDataMap["K6VI"] = append(StellarDataMap["K6VI"], "VI")

// 		case "K7VI":
// 			StellarDataMap["K7VI"] = append(StellarDataMap["K7VI"], "K")
// 			StellarDataMap["K7VI"] = append(StellarDataMap["K7VI"], "7")
// 			StellarDataMap["K7VI"] = append(StellarDataMap["K7VI"], "VI")

// 		case "K8VI":
// 			StellarDataMap["K8VI"] = append(StellarDataMap["K8VI"], "K")
// 			StellarDataMap["K8VI"] = append(StellarDataMap["K8VI"], "8")
// 			StellarDataMap["K8VI"] = append(StellarDataMap["K8VI"], "VI")

// 		case "K9VI":
// 			StellarDataMap["K9VI"] = append(StellarDataMap["K9VI"], "K")
// 			StellarDataMap["K9VI"] = append(StellarDataMap["K9VI"], "9")
// 			StellarDataMap["K9VI"] = append(StellarDataMap["K9VI"], "VI")

// 		case "M0Ia":
// 			StellarDataMap["M0Ia"] = append(StellarDataMap["M0Ia"], "M")
// 			StellarDataMap["M0Ia"] = append(StellarDataMap["M0Ia"], "0")
// 			StellarDataMap["M0Ia"] = append(StellarDataMap["M0Ia"], "Ia")

// 		case "M1Ia":
// 			StellarDataMap["M1Ia"] = append(StellarDataMap["M1Ia"], "M")
// 			StellarDataMap["M1Ia"] = append(StellarDataMap["M1Ia"], "1")
// 			StellarDataMap["M1Ia"] = append(StellarDataMap["M1Ia"], "Ia")

// 		case "M2Ia":
// 			StellarDataMap["M2Ia"] = append(StellarDataMap["M2Ia"], "M")
// 			StellarDataMap["M2Ia"] = append(StellarDataMap["M2Ia"], "2")
// 			StellarDataMap["M2Ia"] = append(StellarDataMap["M2Ia"], "Ia")

// 		case "M3Ia":
// 			StellarDataMap["M3Ia"] = append(StellarDataMap["M3Ia"], "M")
// 			StellarDataMap["M3Ia"] = append(StellarDataMap["M3Ia"], "3")
// 			StellarDataMap["M3Ia"] = append(StellarDataMap["M3Ia"], "Ia")

// 		case "M4Ia":
// 			StellarDataMap["M4Ia"] = append(StellarDataMap["M4Ia"], "M")
// 			StellarDataMap["M4Ia"] = append(StellarDataMap["M4Ia"], "4")
// 			StellarDataMap["M4Ia"] = append(StellarDataMap["M4Ia"], "Ia")

// 		case "M5Ia":
// 			StellarDataMap["M5Ia"] = append(StellarDataMap["M5Ia"], "M")
// 			StellarDataMap["M5Ia"] = append(StellarDataMap["M5Ia"], "5")
// 			StellarDataMap["M5Ia"] = append(StellarDataMap["M5Ia"], "Ia")

// 		case "M6Ia":
// 			StellarDataMap["M6Ia"] = append(StellarDataMap["M6Ia"], "M")
// 			StellarDataMap["M6Ia"] = append(StellarDataMap["M6Ia"], "6")
// 			StellarDataMap["M6Ia"] = append(StellarDataMap["M6Ia"], "Ia")

// 		case "M7Ia":
// 			StellarDataMap["M7Ia"] = append(StellarDataMap["M7Ia"], "M")
// 			StellarDataMap["M7Ia"] = append(StellarDataMap["M7Ia"], "7")
// 			StellarDataMap["M7Ia"] = append(StellarDataMap["M7Ia"], "Ia")

// 		case "M8Ia":
// 			StellarDataMap["M8Ia"] = append(StellarDataMap["M8Ia"], "M")
// 			StellarDataMap["M8Ia"] = append(StellarDataMap["M8Ia"], "8")
// 			StellarDataMap["M8Ia"] = append(StellarDataMap["M8Ia"], "Ia")

// 		case "M9Ia":
// 			StellarDataMap["M9Ia"] = append(StellarDataMap["M9Ia"], "M")
// 			StellarDataMap["M9Ia"] = append(StellarDataMap["M9Ia"], "9")
// 			StellarDataMap["M9Ia"] = append(StellarDataMap["M9Ia"], "Ia")

// 		case "M0Ib":
// 			StellarDataMap["M0Ib"] = append(StellarDataMap["M0Ib"], "M")
// 			StellarDataMap["M0Ib"] = append(StellarDataMap["M0Ib"], "0")
// 			StellarDataMap["M0Ib"] = append(StellarDataMap["M0Ib"], "Ib")

// 		case "M1Ib":
// 			StellarDataMap["M1Ib"] = append(StellarDataMap["M1Ib"], "M")
// 			StellarDataMap["M1Ib"] = append(StellarDataMap["M1Ib"], "1")
// 			StellarDataMap["M1Ib"] = append(StellarDataMap["M1Ib"], "Ib")

// 		case "M2Ib":
// 			StellarDataMap["M2Ib"] = append(StellarDataMap["M2Ib"], "M")
// 			StellarDataMap["M2Ib"] = append(StellarDataMap["M2Ib"], "2")
// 			StellarDataMap["M2Ib"] = append(StellarDataMap["M2Ib"], "Ib")

// 		case "M3Ib":
// 			StellarDataMap["M3Ib"] = append(StellarDataMap["M3Ib"], "M")
// 			StellarDataMap["M3Ib"] = append(StellarDataMap["M3Ib"], "3")
// 			StellarDataMap["M3Ib"] = append(StellarDataMap["M3Ib"], "Ib")

// 		case "M4Ib":
// 			StellarDataMap["M4Ib"] = append(StellarDataMap["M4Ib"], "M")
// 			StellarDataMap["M4Ib"] = append(StellarDataMap["M4Ib"], "4")
// 			StellarDataMap["M4Ib"] = append(StellarDataMap["M4Ib"], "Ib")

// 		case "M5Ib":
// 			StellarDataMap["M5Ib"] = append(StellarDataMap["M5Ib"], "M")
// 			StellarDataMap["M5Ib"] = append(StellarDataMap["M5Ib"], "5")
// 			StellarDataMap["M5Ib"] = append(StellarDataMap["M5Ib"], "Ib")

// 		case "M6Ib":
// 			StellarDataMap["M6Ib"] = append(StellarDataMap["M6Ib"], "M")
// 			StellarDataMap["M6Ib"] = append(StellarDataMap["M6Ib"], "6")
// 			StellarDataMap["M6Ib"] = append(StellarDataMap["M6Ib"], "Ib")

// 		case "M7Ib":
// 			StellarDataMap["M7Ib"] = append(StellarDataMap["M7Ib"], "M")
// 			StellarDataMap["M7Ib"] = append(StellarDataMap["M7Ib"], "7")
// 			StellarDataMap["M7Ib"] = append(StellarDataMap["M7Ib"], "Ib")

// 		case "M8Ib":
// 			StellarDataMap["M8Ib"] = append(StellarDataMap["M8Ib"], "M")
// 			StellarDataMap["M8Ib"] = append(StellarDataMap["M8Ib"], "8")
// 			StellarDataMap["M8Ib"] = append(StellarDataMap["M8Ib"], "Ib")

// 		case "M9Ib":
// 			StellarDataMap["M9Ib"] = append(StellarDataMap["M9Ib"], "M")
// 			StellarDataMap["M9Ib"] = append(StellarDataMap["M9Ib"], "9")
// 			StellarDataMap["M9Ib"] = append(StellarDataMap["M9Ib"], "Ib")

// 		case "M0II":
// 			StellarDataMap["M0II"] = append(StellarDataMap["M0II"], "M")
// 			StellarDataMap["M0II"] = append(StellarDataMap["M0II"], "0")
// 			StellarDataMap["M0II"] = append(StellarDataMap["M0II"], "II")

// 		case "M1II":
// 			StellarDataMap["M1II"] = append(StellarDataMap["M1II"], "M")
// 			StellarDataMap["M1II"] = append(StellarDataMap["M1II"], "1")
// 			StellarDataMap["M1II"] = append(StellarDataMap["M1II"], "II")

// 		case "M2II":
// 			StellarDataMap["M2II"] = append(StellarDataMap["M2II"], "M")
// 			StellarDataMap["M2II"] = append(StellarDataMap["M2II"], "2")
// 			StellarDataMap["M2II"] = append(StellarDataMap["M2II"], "II")

// 		case "M3II":
// 			StellarDataMap["M3II"] = append(StellarDataMap["M3II"], "M")
// 			StellarDataMap["M3II"] = append(StellarDataMap["M3II"], "3")
// 			StellarDataMap["M3II"] = append(StellarDataMap["M3II"], "II")

// 		case "M4II":
// 			StellarDataMap["M4II"] = append(StellarDataMap["M4II"], "M")
// 			StellarDataMap["M4II"] = append(StellarDataMap["M4II"], "4")
// 			StellarDataMap["M4II"] = append(StellarDataMap["M4II"], "II")

// 		case "M5II":
// 			StellarDataMap["M5II"] = append(StellarDataMap["M5II"], "M")
// 			StellarDataMap["M5II"] = append(StellarDataMap["M5II"], "5")
// 			StellarDataMap["M5II"] = append(StellarDataMap["M5II"], "II")

// 		case "M6II":
// 			StellarDataMap["M6II"] = append(StellarDataMap["M6II"], "M")
// 			StellarDataMap["M6II"] = append(StellarDataMap["M6II"], "6")
// 			StellarDataMap["M6II"] = append(StellarDataMap["M6II"], "II")

// 		case "M7II":
// 			StellarDataMap["M7II"] = append(StellarDataMap["M7II"], "M")
// 			StellarDataMap["M7II"] = append(StellarDataMap["M7II"], "7")
// 			StellarDataMap["M7II"] = append(StellarDataMap["M7II"], "II")

// 		case "M8II":
// 			StellarDataMap["M8II"] = append(StellarDataMap["M8II"], "M")
// 			StellarDataMap["M8II"] = append(StellarDataMap["M8II"], "8")
// 			StellarDataMap["M8II"] = append(StellarDataMap["M8II"], "II")

// 		case "M9II":
// 			StellarDataMap["M9II"] = append(StellarDataMap["M9II"], "M")
// 			StellarDataMap["M9II"] = append(StellarDataMap["M9II"], "9")
// 			StellarDataMap["M9II"] = append(StellarDataMap["M9II"], "II")

// 		case "M0III":
// 			StellarDataMap["M0III"] = append(StellarDataMap["M0III"], "M")
// 			StellarDataMap["M0III"] = append(StellarDataMap["M0III"], "0")
// 			StellarDataMap["M0III"] = append(StellarDataMap["M0III"], "III")

// 		case "M1III":
// 			StellarDataMap["M1III"] = append(StellarDataMap["M1III"], "M")
// 			StellarDataMap["M1III"] = append(StellarDataMap["M1III"], "1")
// 			StellarDataMap["M1III"] = append(StellarDataMap["M1III"], "III")

// 		case "M2III":
// 			StellarDataMap["M2III"] = append(StellarDataMap["M2III"], "M")
// 			StellarDataMap["M2III"] = append(StellarDataMap["M2III"], "2")
// 			StellarDataMap["M2III"] = append(StellarDataMap["M2III"], "III")

// 		case "M3III":
// 			StellarDataMap["M3III"] = append(StellarDataMap["M3III"], "M")
// 			StellarDataMap["M3III"] = append(StellarDataMap["M3III"], "3")
// 			StellarDataMap["M3III"] = append(StellarDataMap["M3III"], "III")

// 		case "M4III":
// 			StellarDataMap["M4III"] = append(StellarDataMap["M4III"], "M")
// 			StellarDataMap["M4III"] = append(StellarDataMap["M4III"], "4")
// 			StellarDataMap["M4III"] = append(StellarDataMap["M4III"], "III")

// 		case "M5III":
// 			StellarDataMap["M5III"] = append(StellarDataMap["M5III"], "M")
// 			StellarDataMap["M5III"] = append(StellarDataMap["M5III"], "5")
// 			StellarDataMap["M5III"] = append(StellarDataMap["M5III"], "III")

// 		case "M6III":
// 			StellarDataMap["M6III"] = append(StellarDataMap["M6III"], "M")
// 			StellarDataMap["M6III"] = append(StellarDataMap["M6III"], "6")
// 			StellarDataMap["M6III"] = append(StellarDataMap["M6III"], "III")

// 		case "M7III":
// 			StellarDataMap["M7III"] = append(StellarDataMap["M7III"], "M")
// 			StellarDataMap["M7III"] = append(StellarDataMap["M7III"], "7")
// 			StellarDataMap["M7III"] = append(StellarDataMap["M7III"], "III")

// 		case "M8III":
// 			StellarDataMap["M8III"] = append(StellarDataMap["M8III"], "M")
// 			StellarDataMap["M8III"] = append(StellarDataMap["M8III"], "8")
// 			StellarDataMap["M8III"] = append(StellarDataMap["M8III"], "III")

// 		case "M9III":
// 			StellarDataMap["M9III"] = append(StellarDataMap["M9III"], "M")
// 			StellarDataMap["M9III"] = append(StellarDataMap["M9III"], "9")
// 			StellarDataMap["M9III"] = append(StellarDataMap["M9III"], "III")

// 		case "M0V":
// 			StellarDataMap["M0V"] = append(StellarDataMap["M0V"], "M")
// 			StellarDataMap["M0V"] = append(StellarDataMap["M0V"], "0")
// 			StellarDataMap["M0V"] = append(StellarDataMap["M0V"], "V")

// 		case "M1V":
// 			StellarDataMap["M1V"] = append(StellarDataMap["M1V"], "M")
// 			StellarDataMap["M1V"] = append(StellarDataMap["M1V"], "1")
// 			StellarDataMap["M1V"] = append(StellarDataMap["M1V"], "V")

// 		case "M2V":
// 			StellarDataMap["M2V"] = append(StellarDataMap["M2V"], "M")
// 			StellarDataMap["M2V"] = append(StellarDataMap["M2V"], "2")
// 			StellarDataMap["M2V"] = append(StellarDataMap["M2V"], "V")

// 		case "M3V":
// 			StellarDataMap["M3V"] = append(StellarDataMap["M3V"], "M")
// 			StellarDataMap["M3V"] = append(StellarDataMap["M3V"], "3")
// 			StellarDataMap["M3V"] = append(StellarDataMap["M3V"], "V")

// 		case "M4V":
// 			StellarDataMap["M4V"] = append(StellarDataMap["M4V"], "M")
// 			StellarDataMap["M4V"] = append(StellarDataMap["M4V"], "4")
// 			StellarDataMap["M4V"] = append(StellarDataMap["M4V"], "V")

// 		case "M5V":
// 			StellarDataMap["M5V"] = append(StellarDataMap["M5V"], "M")
// 			StellarDataMap["M5V"] = append(StellarDataMap["M5V"], "5")
// 			StellarDataMap["M5V"] = append(StellarDataMap["M5V"], "V")

// 		case "M6V":
// 			StellarDataMap["M6V"] = append(StellarDataMap["M6V"], "M")
// 			StellarDataMap["M6V"] = append(StellarDataMap["M6V"], "6")
// 			StellarDataMap["M6V"] = append(StellarDataMap["M6V"], "V")

// 		case "M7V":
// 			StellarDataMap["M7V"] = append(StellarDataMap["M7V"], "M")
// 			StellarDataMap["M7V"] = append(StellarDataMap["M7V"], "7")
// 			StellarDataMap["M7V"] = append(StellarDataMap["M7V"], "V")

// 		case "M8V":
// 			StellarDataMap["M8V"] = append(StellarDataMap["M8V"], "M")
// 			StellarDataMap["M8V"] = append(StellarDataMap["M8V"], "8")
// 			StellarDataMap["M8V"] = append(StellarDataMap["M8V"], "V")

// 		case "M9V":
// 			StellarDataMap["M9V"] = append(StellarDataMap["M9V"], "M")
// 			StellarDataMap["M9V"] = append(StellarDataMap["M9V"], "9")
// 			StellarDataMap["M9V"] = append(StellarDataMap["M9V"], "V")

// 		case "M0VI":
// 			StellarDataMap["M0VI"] = append(StellarDataMap["M0VI"], "M")
// 			StellarDataMap["M0VI"] = append(StellarDataMap["M0VI"], "0")
// 			StellarDataMap["M0VI"] = append(StellarDataMap["M0VI"], "VI")

// 		case "M1VI":
// 			StellarDataMap["M1VI"] = append(StellarDataMap["M1VI"], "M")
// 			StellarDataMap["M1VI"] = append(StellarDataMap["M1VI"], "1")
// 			StellarDataMap["M1VI"] = append(StellarDataMap["M1VI"], "VI")

// 		case "M2VI":
// 			StellarDataMap["M2VI"] = append(StellarDataMap["M2VI"], "M")
// 			StellarDataMap["M2VI"] = append(StellarDataMap["M2VI"], "2")
// 			StellarDataMap["M2VI"] = append(StellarDataMap["M2VI"], "VI")

// 		case "M3VI":
// 			StellarDataMap["M3VI"] = append(StellarDataMap["M3VI"], "M")
// 			StellarDataMap["M3VI"] = append(StellarDataMap["M3VI"], "3")
// 			StellarDataMap["M3VI"] = append(StellarDataMap["M3VI"], "VI")

// 		case "M4VI":
// 			StellarDataMap["M4VI"] = append(StellarDataMap["M4VI"], "M")
// 			StellarDataMap["M4VI"] = append(StellarDataMap["M4VI"], "4")
// 			StellarDataMap["M4VI"] = append(StellarDataMap["M4VI"], "VI")

// 		case "M5VI":
// 			StellarDataMap["M5VI"] = append(StellarDataMap["M5VI"], "M")
// 			StellarDataMap["M5VI"] = append(StellarDataMap["M5VI"], "5")
// 			StellarDataMap["M5VI"] = append(StellarDataMap["M5VI"], "VI")

// 		case "M6VI":
// 			StellarDataMap["M6VI"] = append(StellarDataMap["M6VI"], "M")
// 			StellarDataMap["M6VI"] = append(StellarDataMap["M6VI"], "6")
// 			StellarDataMap["M6VI"] = append(StellarDataMap["M6VI"], "VI")

// 		case "M7VI":
// 			StellarDataMap["M7VI"] = append(StellarDataMap["M7VI"], "M")
// 			StellarDataMap["M7VI"] = append(StellarDataMap["M7VI"], "7")
// 			StellarDataMap["M7VI"] = append(StellarDataMap["M7VI"], "VI")

// 		case "M8VI":
// 			StellarDataMap["M8VI"] = append(StellarDataMap["M8VI"], "M")
// 			StellarDataMap["M8VI"] = append(StellarDataMap["M8VI"], "8")
// 			StellarDataMap["M8VI"] = append(StellarDataMap["M8VI"], "VI")

// 		case "M9VI":
// 			StellarDataMap["M9VI"] = append(StellarDataMap["M9VI"], "M")
// 			StellarDataMap["M9VI"] = append(StellarDataMap["M9VI"], "9")
// 			StellarDataMap["M9VI"] = append(StellarDataMap["M9VI"], "VI")

// 		case "O0Ia":
// 			StellarDataMap["O0Ia"] = append(StellarDataMap["O0Ia"], "O")
// 			StellarDataMap["O0Ia"] = append(StellarDataMap["O0Ia"], "0")
// 			StellarDataMap["O0Ia"] = append(StellarDataMap["O0Ia"], "Ia")

// 		case "O1Ia":
// 			StellarDataMap["O1Ia"] = append(StellarDataMap["O1Ia"], "O")
// 			StellarDataMap["O1Ia"] = append(StellarDataMap["O1Ia"], "1")
// 			StellarDataMap["O1Ia"] = append(StellarDataMap["O1Ia"], "Ia")

// 		case "O2Ia":
// 			StellarDataMap["O2Ia"] = append(StellarDataMap["O2Ia"], "O")
// 			StellarDataMap["O2Ia"] = append(StellarDataMap["O2Ia"], "2")
// 			StellarDataMap["O2Ia"] = append(StellarDataMap["O2Ia"], "Ia")

// 		case "O3Ia":
// 			StellarDataMap["O3Ia"] = append(StellarDataMap["O3Ia"], "O")
// 			StellarDataMap["O3Ia"] = append(StellarDataMap["O3Ia"], "3")
// 			StellarDataMap["O3Ia"] = append(StellarDataMap["O3Ia"], "Ia")

// 		case "O4Ia":
// 			StellarDataMap["O4Ia"] = append(StellarDataMap["O4Ia"], "O")
// 			StellarDataMap["O4Ia"] = append(StellarDataMap["O4Ia"], "4")
// 			StellarDataMap["O4Ia"] = append(StellarDataMap["O4Ia"], "Ia")

// 		case "O5Ia":
// 			StellarDataMap["O5Ia"] = append(StellarDataMap["O5Ia"], "O")
// 			StellarDataMap["O5Ia"] = append(StellarDataMap["O5Ia"], "5")
// 			StellarDataMap["O5Ia"] = append(StellarDataMap["O5Ia"], "Ia")

// 		case "O6Ia":
// 			StellarDataMap["O6Ia"] = append(StellarDataMap["O6Ia"], "O")
// 			StellarDataMap["O6Ia"] = append(StellarDataMap["O6Ia"], "6")
// 			StellarDataMap["O6Ia"] = append(StellarDataMap["O6Ia"], "Ia")

// 		case "O7Ia":
// 			StellarDataMap["O7Ia"] = append(StellarDataMap["O7Ia"], "O")
// 			StellarDataMap["O7Ia"] = append(StellarDataMap["O7Ia"], "7")
// 			StellarDataMap["O7Ia"] = append(StellarDataMap["O7Ia"], "Ia")

// 		case "O8Ia":
// 			StellarDataMap["O8Ia"] = append(StellarDataMap["O8Ia"], "O")
// 			StellarDataMap["O8Ia"] = append(StellarDataMap["O8Ia"], "8")
// 			StellarDataMap["O8Ia"] = append(StellarDataMap["O8Ia"], "Ia")

// 		case "O9Ia":
// 			StellarDataMap["O9Ia"] = append(StellarDataMap["O9Ia"], "O")
// 			StellarDataMap["O9Ia"] = append(StellarDataMap["O9Ia"], "9")
// 			StellarDataMap["O9Ia"] = append(StellarDataMap["O9Ia"], "Ia")

// 		case "O0Ib":
// 			StellarDataMap["O0Ib"] = append(StellarDataMap["O0Ib"], "O")
// 			StellarDataMap["O0Ib"] = append(StellarDataMap["O0Ib"], "0")
// 			StellarDataMap["O0Ib"] = append(StellarDataMap["O0Ib"], "Ib")

// 		case "O1Ib":
// 			StellarDataMap["O1Ib"] = append(StellarDataMap["O1Ib"], "O")
// 			StellarDataMap["O1Ib"] = append(StellarDataMap["O1Ib"], "1")
// 			StellarDataMap["O1Ib"] = append(StellarDataMap["O1Ib"], "Ib")

// 		case "O2Ib":
// 			StellarDataMap["O2Ib"] = append(StellarDataMap["O2Ib"], "O")
// 			StellarDataMap["O2Ib"] = append(StellarDataMap["O2Ib"], "2")
// 			StellarDataMap["O2Ib"] = append(StellarDataMap["O2Ib"], "Ib")

// 		case "O3Ib":
// 			StellarDataMap["O3Ib"] = append(StellarDataMap["O3Ib"], "O")
// 			StellarDataMap["O3Ib"] = append(StellarDataMap["O3Ib"], "3")
// 			StellarDataMap["O3Ib"] = append(StellarDataMap["O3Ib"], "Ib")

// 		case "O4Ib":
// 			StellarDataMap["O4Ib"] = append(StellarDataMap["O4Ib"], "O")
// 			StellarDataMap["O4Ib"] = append(StellarDataMap["O4Ib"], "4")
// 			StellarDataMap["O4Ib"] = append(StellarDataMap["O4Ib"], "Ib")

// 		case "O5Ib":
// 			StellarDataMap["O5Ib"] = append(StellarDataMap["O5Ib"], "O")
// 			StellarDataMap["O5Ib"] = append(StellarDataMap["O5Ib"], "5")
// 			StellarDataMap["O5Ib"] = append(StellarDataMap["O5Ib"], "Ib")

// 		case "O6Ib":
// 			StellarDataMap["O6Ib"] = append(StellarDataMap["O6Ib"], "O")
// 			StellarDataMap["O6Ib"] = append(StellarDataMap["O6Ib"], "6")
// 			StellarDataMap["O6Ib"] = append(StellarDataMap["O6Ib"], "Ib")

// 		case "O7Ib":
// 			StellarDataMap["O7Ib"] = append(StellarDataMap["O7Ib"], "O")
// 			StellarDataMap["O7Ib"] = append(StellarDataMap["O7Ib"], "7")
// 			StellarDataMap["O7Ib"] = append(StellarDataMap["O7Ib"], "Ib")

// 		case "O8Ib":
// 			StellarDataMap["O8Ib"] = append(StellarDataMap["O8Ib"], "O")
// 			StellarDataMap["O8Ib"] = append(StellarDataMap["O8Ib"], "8")
// 			StellarDataMap["O8Ib"] = append(StellarDataMap["O8Ib"], "Ib")

// 		case "O9Ib":
// 			StellarDataMap["O9Ib"] = append(StellarDataMap["O9Ib"], "O")
// 			StellarDataMap["O9Ib"] = append(StellarDataMap["O9Ib"], "9")
// 			StellarDataMap["O9Ib"] = append(StellarDataMap["O9Ib"], "Ib")

// 		case "O0II":
// 			StellarDataMap["O0II"] = append(StellarDataMap["O0II"], "O")
// 			StellarDataMap["O0II"] = append(StellarDataMap["O0II"], "0")
// 			StellarDataMap["O0II"] = append(StellarDataMap["O0II"], "II")

// 		case "O1II":
// 			StellarDataMap["O1II"] = append(StellarDataMap["O1II"], "O")
// 			StellarDataMap["O1II"] = append(StellarDataMap["O1II"], "1")
// 			StellarDataMap["O1II"] = append(StellarDataMap["O1II"], "II")

// 		case "O2II":
// 			StellarDataMap["O2II"] = append(StellarDataMap["O2II"], "O")
// 			StellarDataMap["O2II"] = append(StellarDataMap["O2II"], "2")
// 			StellarDataMap["O2II"] = append(StellarDataMap["O2II"], "II")

// 		case "O3II":
// 			StellarDataMap["O3II"] = append(StellarDataMap["O3II"], "O")
// 			StellarDataMap["O3II"] = append(StellarDataMap["O3II"], "3")
// 			StellarDataMap["O3II"] = append(StellarDataMap["O3II"], "II")

// 		case "O4II":
// 			StellarDataMap["O4II"] = append(StellarDataMap["O4II"], "O")
// 			StellarDataMap["O4II"] = append(StellarDataMap["O4II"], "4")
// 			StellarDataMap["O4II"] = append(StellarDataMap["O4II"], "II")

// 		case "O5II":
// 			StellarDataMap["O5II"] = append(StellarDataMap["O5II"], "O")
// 			StellarDataMap["O5II"] = append(StellarDataMap["O5II"], "5")
// 			StellarDataMap["O5II"] = append(StellarDataMap["O5II"], "II")

// 		case "O6II":
// 			StellarDataMap["O6II"] = append(StellarDataMap["O6II"], "O")
// 			StellarDataMap["O6II"] = append(StellarDataMap["O6II"], "6")
// 			StellarDataMap["O6II"] = append(StellarDataMap["O6II"], "II")

// 		case "O7II":
// 			StellarDataMap["O7II"] = append(StellarDataMap["O7II"], "O")
// 			StellarDataMap["O7II"] = append(StellarDataMap["O7II"], "7")
// 			StellarDataMap["O7II"] = append(StellarDataMap["O7II"], "II")

// 		case "O8II":
// 			StellarDataMap["O8II"] = append(StellarDataMap["O8II"], "O")
// 			StellarDataMap["O8II"] = append(StellarDataMap["O8II"], "8")
// 			StellarDataMap["O8II"] = append(StellarDataMap["O8II"], "II")

// 		case "O9II":
// 			StellarDataMap["O9II"] = append(StellarDataMap["O9II"], "O")
// 			StellarDataMap["O9II"] = append(StellarDataMap["O9II"], "9")
// 			StellarDataMap["O9II"] = append(StellarDataMap["O9II"], "II")

// 		case "O0III":
// 			StellarDataMap["O0III"] = append(StellarDataMap["O0III"], "O")
// 			StellarDataMap["O0III"] = append(StellarDataMap["O0III"], "0")
// 			StellarDataMap["O0III"] = append(StellarDataMap["O0III"], "III")

// 		case "O1III":
// 			StellarDataMap["O1III"] = append(StellarDataMap["O1III"], "O")
// 			StellarDataMap["O1III"] = append(StellarDataMap["O1III"], "1")
// 			StellarDataMap["O1III"] = append(StellarDataMap["O1III"], "III")

// 		case "O2III":
// 			StellarDataMap["O2III"] = append(StellarDataMap["O2III"], "O")
// 			StellarDataMap["O2III"] = append(StellarDataMap["O2III"], "2")
// 			StellarDataMap["O2III"] = append(StellarDataMap["O2III"], "III")

// 		case "O3III":
// 			StellarDataMap["O3III"] = append(StellarDataMap["O3III"], "O")
// 			StellarDataMap["O3III"] = append(StellarDataMap["O3III"], "3")
// 			StellarDataMap["O3III"] = append(StellarDataMap["O3III"], "III")

// 		case "O4III":
// 			StellarDataMap["O4III"] = append(StellarDataMap["O4III"], "O")
// 			StellarDataMap["O4III"] = append(StellarDataMap["O4III"], "4")
// 			StellarDataMap["O4III"] = append(StellarDataMap["O4III"], "III")

// 		case "O5III":
// 			StellarDataMap["O5III"] = append(StellarDataMap["O5III"], "O")
// 			StellarDataMap["O5III"] = append(StellarDataMap["O5III"], "5")
// 			StellarDataMap["O5III"] = append(StellarDataMap["O5III"], "III")

// 		case "O6III":
// 			StellarDataMap["O6III"] = append(StellarDataMap["O6III"], "O")
// 			StellarDataMap["O6III"] = append(StellarDataMap["O6III"], "6")
// 			StellarDataMap["O6III"] = append(StellarDataMap["O6III"], "III")

// 		case "O7III":
// 			StellarDataMap["O7III"] = append(StellarDataMap["O7III"], "O")
// 			StellarDataMap["O7III"] = append(StellarDataMap["O7III"], "7")
// 			StellarDataMap["O7III"] = append(StellarDataMap["O7III"], "III")

// 		case "O8III":
// 			StellarDataMap["O8III"] = append(StellarDataMap["O8III"], "O")
// 			StellarDataMap["O8III"] = append(StellarDataMap["O8III"], "8")
// 			StellarDataMap["O8III"] = append(StellarDataMap["O8III"], "III")

// 		case "O9III":
// 			StellarDataMap["O9III"] = append(StellarDataMap["O9III"], "O")
// 			StellarDataMap["O9III"] = append(StellarDataMap["O9III"], "9")
// 			StellarDataMap["O9III"] = append(StellarDataMap["O9III"], "III")

// 		case "O0IV":
// 			StellarDataMap["O0IV"] = append(StellarDataMap["O0IV"], "O")
// 			StellarDataMap["O0IV"] = append(StellarDataMap["O0IV"], "0")
// 			StellarDataMap["O0IV"] = append(StellarDataMap["O0IV"], "IV")

// 		case "O1IV":
// 			StellarDataMap["O1IV"] = append(StellarDataMap["O1IV"], "O")
// 			StellarDataMap["O1IV"] = append(StellarDataMap["O1IV"], "1")
// 			StellarDataMap["O1IV"] = append(StellarDataMap["O1IV"], "IV")

// 		case "O2IV":
// 			StellarDataMap["O2IV"] = append(StellarDataMap["O2IV"], "O")
// 			StellarDataMap["O2IV"] = append(StellarDataMap["O2IV"], "2")
// 			StellarDataMap["O2IV"] = append(StellarDataMap["O2IV"], "IV")

// 		case "O3IV":
// 			StellarDataMap["O3IV"] = append(StellarDataMap["O3IV"], "O")
// 			StellarDataMap["O3IV"] = append(StellarDataMap["O3IV"], "3")
// 			StellarDataMap["O3IV"] = append(StellarDataMap["O3IV"], "IV")

// 		case "O4IV":
// 			StellarDataMap["O4IV"] = append(StellarDataMap["O4IV"], "O")
// 			StellarDataMap["O4IV"] = append(StellarDataMap["O4IV"], "4")
// 			StellarDataMap["O4IV"] = append(StellarDataMap["O4IV"], "IV")

// 		case "O5IV":
// 			StellarDataMap["O5IV"] = append(StellarDataMap["O5IV"], "O")
// 			StellarDataMap["O5IV"] = append(StellarDataMap["O5IV"], "5")
// 			StellarDataMap["O5IV"] = append(StellarDataMap["O5IV"], "IV")

// 		case "O6IV":
// 			StellarDataMap["O6IV"] = append(StellarDataMap["O6IV"], "O")
// 			StellarDataMap["O6IV"] = append(StellarDataMap["O6IV"], "6")
// 			StellarDataMap["O6IV"] = append(StellarDataMap["O6IV"], "IV")

// 		case "O7IV":
// 			StellarDataMap["O7IV"] = append(StellarDataMap["O7IV"], "O")
// 			StellarDataMap["O7IV"] = append(StellarDataMap["O7IV"], "7")
// 			StellarDataMap["O7IV"] = append(StellarDataMap["O7IV"], "IV")

// 		case "O8IV":
// 			StellarDataMap["O8IV"] = append(StellarDataMap["O8IV"], "O")
// 			StellarDataMap["O8IV"] = append(StellarDataMap["O8IV"], "8")
// 			StellarDataMap["O8IV"] = append(StellarDataMap["O8IV"], "IV")

// 		case "O9IV":
// 			StellarDataMap["O9IV"] = append(StellarDataMap["O9IV"], "O")
// 			StellarDataMap["O9IV"] = append(StellarDataMap["O9IV"], "9")
// 			StellarDataMap["O9IV"] = append(StellarDataMap["O9IV"], "IV")

// 		case "O0V":
// 			StellarDataMap["O0V"] = append(StellarDataMap["O0V"], "O")
// 			StellarDataMap["O0V"] = append(StellarDataMap["O0V"], "0")
// 			StellarDataMap["O0V"] = append(StellarDataMap["O0V"], "V")

// 		case "O1V":
// 			StellarDataMap["O1V"] = append(StellarDataMap["O1V"], "O")
// 			StellarDataMap["O1V"] = append(StellarDataMap["O1V"], "1")
// 			StellarDataMap["O1V"] = append(StellarDataMap["O1V"], "V")

// 		case "O2V":
// 			StellarDataMap["O2V"] = append(StellarDataMap["O2V"], "O")
// 			StellarDataMap["O2V"] = append(StellarDataMap["O2V"], "2")
// 			StellarDataMap["O2V"] = append(StellarDataMap["O2V"], "V")

// 		case "O3V":
// 			StellarDataMap["O3V"] = append(StellarDataMap["O3V"], "O")
// 			StellarDataMap["O3V"] = append(StellarDataMap["O3V"], "3")
// 			StellarDataMap["O3V"] = append(StellarDataMap["O3V"], "V")

// 		case "O4V":
// 			StellarDataMap["O4V"] = append(StellarDataMap["O4V"], "O")
// 			StellarDataMap["O4V"] = append(StellarDataMap["O4V"], "4")
// 			StellarDataMap["O4V"] = append(StellarDataMap["O4V"], "V")

// 		case "O5V":
// 			StellarDataMap["O5V"] = append(StellarDataMap["O5V"], "O")
// 			StellarDataMap["O5V"] = append(StellarDataMap["O5V"], "5")
// 			StellarDataMap["O5V"] = append(StellarDataMap["O5V"], "V")

// 		case "O6V":
// 			StellarDataMap["O6V"] = append(StellarDataMap["O6V"], "O")
// 			StellarDataMap["O6V"] = append(StellarDataMap["O6V"], "6")
// 			StellarDataMap["O6V"] = append(StellarDataMap["O6V"], "V")

// 		case "O7V":
// 			StellarDataMap["O7V"] = append(StellarDataMap["O7V"], "O")
// 			StellarDataMap["O7V"] = append(StellarDataMap["O7V"], "7")
// 			StellarDataMap["O7V"] = append(StellarDataMap["O7V"], "V")

// 		case "O8V":
// 			StellarDataMap["O8V"] = append(StellarDataMap["O8V"], "O")
// 			StellarDataMap["O8V"] = append(StellarDataMap["O8V"], "8")
// 			StellarDataMap["O8V"] = append(StellarDataMap["O8V"], "V")

// 		case "O9V":
// 			StellarDataMap["O9V"] = append(StellarDataMap["O9V"], "O")
// 			StellarDataMap["O9V"] = append(StellarDataMap["O9V"], "9")
// 			StellarDataMap["O9V"] = append(StellarDataMap["O9V"], "V")

// 		case "B0Ia":
// 			StellarDataMap["B0Ia"] = append(StellarDataMap["B0Ia"], "B")
// 			StellarDataMap["B0Ia"] = append(StellarDataMap["B0Ia"], "0")
// 			StellarDataMap["B0Ia"] = append(StellarDataMap["B0Ia"], "Ia")

// 		case "B1Ia":
// 			StellarDataMap["B1Ia"] = append(StellarDataMap["B1Ia"], "B")
// 			StellarDataMap["B1Ia"] = append(StellarDataMap["B1Ia"], "1")
// 			StellarDataMap["B1Ia"] = append(StellarDataMap["B1Ia"], "Ia")

// 		case "B2Ia":
// 			StellarDataMap["B2Ia"] = append(StellarDataMap["B2Ia"], "B")
// 			StellarDataMap["B2Ia"] = append(StellarDataMap["B2Ia"], "2")
// 			StellarDataMap["B2Ia"] = append(StellarDataMap["B2Ia"], "Ia")

// 		case "B3Ia":
// 			StellarDataMap["B3Ia"] = append(StellarDataMap["B3Ia"], "B")
// 			StellarDataMap["B3Ia"] = append(StellarDataMap["B3Ia"], "3")
// 			StellarDataMap["B3Ia"] = append(StellarDataMap["B3Ia"], "Ia")

// 		case "B4Ia":
// 			StellarDataMap["B4Ia"] = append(StellarDataMap["B4Ia"], "B")
// 			StellarDataMap["B4Ia"] = append(StellarDataMap["B4Ia"], "4")
// 			StellarDataMap["B4Ia"] = append(StellarDataMap["B4Ia"], "Ia")

// 		case "B5Ia":
// 			StellarDataMap["B5Ia"] = append(StellarDataMap["B5Ia"], "B")
// 			StellarDataMap["B5Ia"] = append(StellarDataMap["B5Ia"], "5")
// 			StellarDataMap["B5Ia"] = append(StellarDataMap["B5Ia"], "Ia")

// 		case "B6Ia":
// 			StellarDataMap["B6Ia"] = append(StellarDataMap["B6Ia"], "B")
// 			StellarDataMap["B6Ia"] = append(StellarDataMap["B6Ia"], "6")
// 			StellarDataMap["B6Ia"] = append(StellarDataMap["B6Ia"], "Ia")

// 		case "B7Ia":
// 			StellarDataMap["B7Ia"] = append(StellarDataMap["B7Ia"], "B")
// 			StellarDataMap["B7Ia"] = append(StellarDataMap["B7Ia"], "7")
// 			StellarDataMap["B7Ia"] = append(StellarDataMap["B7Ia"], "Ia")

// 		case "B8Ia":
// 			StellarDataMap["B8Ia"] = append(StellarDataMap["B8Ia"], "B")
// 			StellarDataMap["B8Ia"] = append(StellarDataMap["B8Ia"], "8")
// 			StellarDataMap["B8Ia"] = append(StellarDataMap["B8Ia"], "Ia")

// 		case "B9Ia":
// 			StellarDataMap["B9Ia"] = append(StellarDataMap["B9Ia"], "B")
// 			StellarDataMap["B9Ia"] = append(StellarDataMap["B9Ia"], "9")
// 			StellarDataMap["B9Ia"] = append(StellarDataMap["B9Ia"], "Ia")

// 		case "B0Ib":
// 			StellarDataMap["B0Ib"] = append(StellarDataMap["B0Ib"], "B")
// 			StellarDataMap["B0Ib"] = append(StellarDataMap["B0Ib"], "0")
// 			StellarDataMap["B0Ib"] = append(StellarDataMap["B0Ib"], "Ib")

// 		case "B1Ib":
// 			StellarDataMap["B1Ib"] = append(StellarDataMap["B1Ib"], "B")
// 			StellarDataMap["B1Ib"] = append(StellarDataMap["B1Ib"], "1")
// 			StellarDataMap["B1Ib"] = append(StellarDataMap["B1Ib"], "Ib")

// 		case "B2Ib":
// 			StellarDataMap["B2Ib"] = append(StellarDataMap["B2Ib"], "B")
// 			StellarDataMap["B2Ib"] = append(StellarDataMap["B2Ib"], "2")
// 			StellarDataMap["B2Ib"] = append(StellarDataMap["B2Ib"], "Ib")

// 		case "B3Ib":
// 			StellarDataMap["B3Ib"] = append(StellarDataMap["B3Ib"], "B")
// 			StellarDataMap["B3Ib"] = append(StellarDataMap["B3Ib"], "3")
// 			StellarDataMap["B3Ib"] = append(StellarDataMap["B3Ib"], "Ib")

// 		case "B4Ib":
// 			StellarDataMap["B4Ib"] = append(StellarDataMap["B4Ib"], "B")
// 			StellarDataMap["B4Ib"] = append(StellarDataMap["B4Ib"], "4")
// 			StellarDataMap["B4Ib"] = append(StellarDataMap["B4Ib"], "Ib")

// 		case "B5Ib":
// 			StellarDataMap["B5Ib"] = append(StellarDataMap["B5Ib"], "B")
// 			StellarDataMap["B5Ib"] = append(StellarDataMap["B5Ib"], "5")
// 			StellarDataMap["B5Ib"] = append(StellarDataMap["B5Ib"], "Ib")

// 		case "B6Ib":
// 			StellarDataMap["B6Ib"] = append(StellarDataMap["B6Ib"], "B")
// 			StellarDataMap["B6Ib"] = append(StellarDataMap["B6Ib"], "6")
// 			StellarDataMap["B6Ib"] = append(StellarDataMap["B6Ib"], "Ib")

// 		case "B7Ib":
// 			StellarDataMap["B7Ib"] = append(StellarDataMap["B7Ib"], "B")
// 			StellarDataMap["B7Ib"] = append(StellarDataMap["B7Ib"], "7")
// 			StellarDataMap["B7Ib"] = append(StellarDataMap["B7Ib"], "Ib")

// 		case "B8Ib":
// 			StellarDataMap["B8Ib"] = append(StellarDataMap["B8Ib"], "B")
// 			StellarDataMap["B8Ib"] = append(StellarDataMap["B8Ib"], "8")
// 			StellarDataMap["B8Ib"] = append(StellarDataMap["B8Ib"], "Ib")

// 		case "B9Ib":
// 			StellarDataMap["B9Ib"] = append(StellarDataMap["B9Ib"], "B")
// 			StellarDataMap["B9Ib"] = append(StellarDataMap["B9Ib"], "9")
// 			StellarDataMap["B9Ib"] = append(StellarDataMap["B9Ib"], "Ib")

// 		case "B0II":
// 			StellarDataMap["B0II"] = append(StellarDataMap["B0II"], "B")
// 			StellarDataMap["B0II"] = append(StellarDataMap["B0II"], "0")
// 			StellarDataMap["B0II"] = append(StellarDataMap["B0II"], "II")

// 		case "B1II":
// 			StellarDataMap["B1II"] = append(StellarDataMap["B1II"], "B")
// 			StellarDataMap["B1II"] = append(StellarDataMap["B1II"], "1")
// 			StellarDataMap["B1II"] = append(StellarDataMap["B1II"], "II")

// 		case "B2II":
// 			StellarDataMap["B2II"] = append(StellarDataMap["B2II"], "B")
// 			StellarDataMap["B2II"] = append(StellarDataMap["B2II"], "2")
// 			StellarDataMap["B2II"] = append(StellarDataMap["B2II"], "II")

// 		case "B3II":
// 			StellarDataMap["B3II"] = append(StellarDataMap["B3II"], "B")
// 			StellarDataMap["B3II"] = append(StellarDataMap["B3II"], "3")
// 			StellarDataMap["B3II"] = append(StellarDataMap["B3II"], "II")

// 		case "B4II":
// 			StellarDataMap["B4II"] = append(StellarDataMap["B4II"], "B")
// 			StellarDataMap["B4II"] = append(StellarDataMap["B4II"], "4")
// 			StellarDataMap["B4II"] = append(StellarDataMap["B4II"], "II")

// 		case "B5II":
// 			StellarDataMap["B5II"] = append(StellarDataMap["B5II"], "B")
// 			StellarDataMap["B5II"] = append(StellarDataMap["B5II"], "5")
// 			StellarDataMap["B5II"] = append(StellarDataMap["B5II"], "II")

// 		case "B6II":
// 			StellarDataMap["B6II"] = append(StellarDataMap["B6II"], "B")
// 			StellarDataMap["B6II"] = append(StellarDataMap["B6II"], "6")
// 			StellarDataMap["B6II"] = append(StellarDataMap["B6II"], "II")

// 		case "B7II":
// 			StellarDataMap["B7II"] = append(StellarDataMap["B7II"], "B")
// 			StellarDataMap["B7II"] = append(StellarDataMap["B7II"], "7")
// 			StellarDataMap["B7II"] = append(StellarDataMap["B7II"], "II")

// 		case "B8II":
// 			StellarDataMap["B8II"] = append(StellarDataMap["B8II"], "B")
// 			StellarDataMap["B8II"] = append(StellarDataMap["B8II"], "8")
// 			StellarDataMap["B8II"] = append(StellarDataMap["B8II"], "II")

// 		case "B9II":
// 			StellarDataMap["B9II"] = append(StellarDataMap["B9II"], "B")
// 			StellarDataMap["B9II"] = append(StellarDataMap["B9II"], "9")
// 			StellarDataMap["B9II"] = append(StellarDataMap["B9II"], "II")

// 		case "B0III":
// 			StellarDataMap["B0III"] = append(StellarDataMap["B0III"], "B")
// 			StellarDataMap["B0III"] = append(StellarDataMap["B0III"], "0")
// 			StellarDataMap["B0III"] = append(StellarDataMap["B0III"], "III")

// 		case "B1III":
// 			StellarDataMap["B1III"] = append(StellarDataMap["B1III"], "B")
// 			StellarDataMap["B1III"] = append(StellarDataMap["B1III"], "1")
// 			StellarDataMap["B1III"] = append(StellarDataMap["B1III"], "III")

// 		case "B2III":
// 			StellarDataMap["B2III"] = append(StellarDataMap["B2III"], "B")
// 			StellarDataMap["B2III"] = append(StellarDataMap["B2III"], "2")
// 			StellarDataMap["B2III"] = append(StellarDataMap["B2III"], "III")

// 		case "B3III":
// 			StellarDataMap["B3III"] = append(StellarDataMap["B3III"], "B")
// 			StellarDataMap["B3III"] = append(StellarDataMap["B3III"], "3")
// 			StellarDataMap["B3III"] = append(StellarDataMap["B3III"], "III")

// 		case "B4III":
// 			StellarDataMap["B4III"] = append(StellarDataMap["B4III"], "B")
// 			StellarDataMap["B4III"] = append(StellarDataMap["B4III"], "4")
// 			StellarDataMap["B4III"] = append(StellarDataMap["B4III"], "III")

// 		case "B5III":
// 			StellarDataMap["B5III"] = append(StellarDataMap["B5III"], "B")
// 			StellarDataMap["B5III"] = append(StellarDataMap["B5III"], "5")
// 			StellarDataMap["B5III"] = append(StellarDataMap["B5III"], "III")

// 		case "B6III":
// 			StellarDataMap["B6III"] = append(StellarDataMap["B6III"], "B")
// 			StellarDataMap["B6III"] = append(StellarDataMap["B6III"], "6")
// 			StellarDataMap["B6III"] = append(StellarDataMap["B6III"], "III")

// 		case "B7III":
// 			StellarDataMap["B7III"] = append(StellarDataMap["B7III"], "B")
// 			StellarDataMap["B7III"] = append(StellarDataMap["B7III"], "7")
// 			StellarDataMap["B7III"] = append(StellarDataMap["B7III"], "III")

// 		case "B8III":
// 			StellarDataMap["B8III"] = append(StellarDataMap["B8III"], "B")
// 			StellarDataMap["B8III"] = append(StellarDataMap["B8III"], "8")
// 			StellarDataMap["B8III"] = append(StellarDataMap["B8III"], "III")

// 		case "B9III":
// 			StellarDataMap["B9III"] = append(StellarDataMap["B9III"], "B")
// 			StellarDataMap["B9III"] = append(StellarDataMap["B9III"], "9")
// 			StellarDataMap["B9III"] = append(StellarDataMap["B9III"], "III")

// 		case "B0IV":
// 			StellarDataMap["B0IV"] = append(StellarDataMap["B0IV"], "B")
// 			StellarDataMap["B0IV"] = append(StellarDataMap["B0IV"], "0")
// 			StellarDataMap["B0IV"] = append(StellarDataMap["B0IV"], "IV")

// 		case "B1IV":
// 			StellarDataMap["B1IV"] = append(StellarDataMap["B1IV"], "B")
// 			StellarDataMap["B1IV"] = append(StellarDataMap["B1IV"], "1")
// 			StellarDataMap["B1IV"] = append(StellarDataMap["B1IV"], "IV")

// 		case "B2IV":
// 			StellarDataMap["B2IV"] = append(StellarDataMap["B2IV"], "B")
// 			StellarDataMap["B2IV"] = append(StellarDataMap["B2IV"], "2")
// 			StellarDataMap["B2IV"] = append(StellarDataMap["B2IV"], "IV")

// 		case "B3IV":
// 			StellarDataMap["B3IV"] = append(StellarDataMap["B3IV"], "B")
// 			StellarDataMap["B3IV"] = append(StellarDataMap["B3IV"], "3")
// 			StellarDataMap["B3IV"] = append(StellarDataMap["B3IV"], "IV")

// 		case "B4IV":
// 			StellarDataMap["B4IV"] = append(StellarDataMap["B4IV"], "B")
// 			StellarDataMap["B4IV"] = append(StellarDataMap["B4IV"], "4")
// 			StellarDataMap["B4IV"] = append(StellarDataMap["B4IV"], "IV")

// 		case "B5IV":
// 			StellarDataMap["B5IV"] = append(StellarDataMap["B5IV"], "B")
// 			StellarDataMap["B5IV"] = append(StellarDataMap["B5IV"], "5")
// 			StellarDataMap["B5IV"] = append(StellarDataMap["B5IV"], "IV")

// 		case "B6IV":
// 			StellarDataMap["B6IV"] = append(StellarDataMap["B6IV"], "B")
// 			StellarDataMap["B6IV"] = append(StellarDataMap["B6IV"], "6")
// 			StellarDataMap["B6IV"] = append(StellarDataMap["B6IV"], "IV")

// 		case "B7IV":
// 			StellarDataMap["B7IV"] = append(StellarDataMap["B7IV"], "B")
// 			StellarDataMap["B7IV"] = append(StellarDataMap["B7IV"], "7")
// 			StellarDataMap["B7IV"] = append(StellarDataMap["B7IV"], "IV")

// 		case "B8IV":
// 			StellarDataMap["B8IV"] = append(StellarDataMap["B8IV"], "B")
// 			StellarDataMap["B8IV"] = append(StellarDataMap["B8IV"], "8")
// 			StellarDataMap["B8IV"] = append(StellarDataMap["B8IV"], "IV")

// 		case "B9IV":
// 			StellarDataMap["B9IV"] = append(StellarDataMap["B9IV"], "B")
// 			StellarDataMap["B9IV"] = append(StellarDataMap["B9IV"], "9")
// 			StellarDataMap["B9IV"] = append(StellarDataMap["B9IV"], "IV")

// 		case "B0V":
// 			StellarDataMap["B0V"] = append(StellarDataMap["B0V"], "B")
// 			StellarDataMap["B0V"] = append(StellarDataMap["B0V"], "0")
// 			StellarDataMap["B0V"] = append(StellarDataMap["B0V"], "V")

// 		case "B1V":
// 			StellarDataMap["B1V"] = append(StellarDataMap["B1V"], "B")
// 			StellarDataMap["B1V"] = append(StellarDataMap["B1V"], "1")
// 			StellarDataMap["B1V"] = append(StellarDataMap["B1V"], "V")

// 		case "B2V":
// 			StellarDataMap["B2V"] = append(StellarDataMap["B2V"], "B")
// 			StellarDataMap["B2V"] = append(StellarDataMap["B2V"], "2")
// 			StellarDataMap["B2V"] = append(StellarDataMap["B2V"], "V")

// 		case "B3V":
// 			StellarDataMap["B3V"] = append(StellarDataMap["B3V"], "B")
// 			StellarDataMap["B3V"] = append(StellarDataMap["B3V"], "3")
// 			StellarDataMap["B3V"] = append(StellarDataMap["B3V"], "V")

// 		case "B4V":
// 			StellarDataMap["B4V"] = append(StellarDataMap["B4V"], "B")
// 			StellarDataMap["B4V"] = append(StellarDataMap["B4V"], "4")
// 			StellarDataMap["B4V"] = append(StellarDataMap["B4V"], "V")

// 		case "B5V":
// 			StellarDataMap["B5V"] = append(StellarDataMap["B5V"], "B")
// 			StellarDataMap["B5V"] = append(StellarDataMap["B5V"], "5")
// 			StellarDataMap["B5V"] = append(StellarDataMap["B5V"], "V")

// 		case "B6V":
// 			StellarDataMap["B6V"] = append(StellarDataMap["B6V"], "B")
// 			StellarDataMap["B6V"] = append(StellarDataMap["B6V"], "6")
// 			StellarDataMap["B6V"] = append(StellarDataMap["B6V"], "V")

// 		case "B7V":
// 			StellarDataMap["B7V"] = append(StellarDataMap["B7V"], "B")
// 			StellarDataMap["B7V"] = append(StellarDataMap["B7V"], "7")
// 			StellarDataMap["B7V"] = append(StellarDataMap["B7V"], "V")

// 		case "B8V":
// 			StellarDataMap["B8V"] = append(StellarDataMap["B8V"], "B")
// 			StellarDataMap["B8V"] = append(StellarDataMap["B8V"], "8")
// 			StellarDataMap["B8V"] = append(StellarDataMap["B8V"], "V")

// 		case "B9V":
// 			StellarDataMap["B9V"] = append(StellarDataMap["B9V"], "B")
// 			StellarDataMap["B9V"] = append(StellarDataMap["B9V"], "9")
// 			StellarDataMap["B9V"] = append(StellarDataMap["B9V"], "V")

// 		case "B0VI":
// 			StellarDataMap["B0VI"] = append(StellarDataMap["B0VI"], "B")
// 			StellarDataMap["B0VI"] = append(StellarDataMap["B0VI"], "0")
// 			StellarDataMap["B0VI"] = append(StellarDataMap["B0VI"], "VI")

// 		case "A0Ia":
// 			StellarDataMap["A0Ia"] = append(StellarDataMap["A0Ia"], "A")
// 			StellarDataMap["A0Ia"] = append(StellarDataMap["A0Ia"], "0")
// 			StellarDataMap["A0Ia"] = append(StellarDataMap["A0Ia"], "Ia")

// 		case "A1Ia":
// 			StellarDataMap["A1Ia"] = append(StellarDataMap["A1Ia"], "A")
// 			StellarDataMap["A1Ia"] = append(StellarDataMap["A1Ia"], "1")
// 			StellarDataMap["A1Ia"] = append(StellarDataMap["A1Ia"], "Ia")

// 		case "A2Ia":
// 			StellarDataMap["A2Ia"] = append(StellarDataMap["A2Ia"], "A")
// 			StellarDataMap["A2Ia"] = append(StellarDataMap["A2Ia"], "2")
// 			StellarDataMap["A2Ia"] = append(StellarDataMap["A2Ia"], "Ia")

// 		case "A3Ia":
// 			StellarDataMap["A3Ia"] = append(StellarDataMap["A3Ia"], "A")
// 			StellarDataMap["A3Ia"] = append(StellarDataMap["A3Ia"], "3")
// 			StellarDataMap["A3Ia"] = append(StellarDataMap["A3Ia"], "Ia")

// 		case "A4Ia":
// 			StellarDataMap["A4Ia"] = append(StellarDataMap["A4Ia"], "A")
// 			StellarDataMap["A4Ia"] = append(StellarDataMap["A4Ia"], "4")
// 			StellarDataMap["A4Ia"] = append(StellarDataMap["A4Ia"], "Ia")

// 		case "A5Ia":
// 			StellarDataMap["A5Ia"] = append(StellarDataMap["A5Ia"], "A")
// 			StellarDataMap["A5Ia"] = append(StellarDataMap["A5Ia"], "5")
// 			StellarDataMap["A5Ia"] = append(StellarDataMap["A5Ia"], "Ia")

// 		case "A6Ia":
// 			StellarDataMap["A6Ia"] = append(StellarDataMap["A6Ia"], "A")
// 			StellarDataMap["A6Ia"] = append(StellarDataMap["A6Ia"], "6")
// 			StellarDataMap["A6Ia"] = append(StellarDataMap["A6Ia"], "Ia")

// 		case "A7Ia":
// 			StellarDataMap["A7Ia"] = append(StellarDataMap["A7Ia"], "A")
// 			StellarDataMap["A7Ia"] = append(StellarDataMap["A7Ia"], "7")
// 			StellarDataMap["A7Ia"] = append(StellarDataMap["A7Ia"], "Ia")

// 		case "A8Ia":
// 			StellarDataMap["A8Ia"] = append(StellarDataMap["A8Ia"], "A")
// 			StellarDataMap["A8Ia"] = append(StellarDataMap["A8Ia"], "8")
// 			StellarDataMap["A8Ia"] = append(StellarDataMap["A8Ia"], "Ia")

// 		case "A9Ia":
// 			StellarDataMap["A9Ia"] = append(StellarDataMap["A9Ia"], "A")
// 			StellarDataMap["A9Ia"] = append(StellarDataMap["A9Ia"], "9")
// 			StellarDataMap["A9Ia"] = append(StellarDataMap["A9Ia"], "Ia")

// 		case "A0Ib":
// 			StellarDataMap["A0Ib"] = append(StellarDataMap["A0Ib"], "A")
// 			StellarDataMap["A0Ib"] = append(StellarDataMap["A0Ib"], "0")
// 			StellarDataMap["A0Ib"] = append(StellarDataMap["A0Ib"], "Ib")

// 		case "A1Ib":
// 			StellarDataMap["A1Ib"] = append(StellarDataMap["A1Ib"], "A")
// 			StellarDataMap["A1Ib"] = append(StellarDataMap["A1Ib"], "1")
// 			StellarDataMap["A1Ib"] = append(StellarDataMap["A1Ib"], "Ib")

// 		case "A2Ib":
// 			StellarDataMap["A2Ib"] = append(StellarDataMap["A2Ib"], "A")
// 			StellarDataMap["A2Ib"] = append(StellarDataMap["A2Ib"], "2")
// 			StellarDataMap["A2Ib"] = append(StellarDataMap["A2Ib"], "Ib")

// 		case "A3Ib":
// 			StellarDataMap["A3Ib"] = append(StellarDataMap["A3Ib"], "A")
// 			StellarDataMap["A3Ib"] = append(StellarDataMap["A3Ib"], "3")
// 			StellarDataMap["A3Ib"] = append(StellarDataMap["A3Ib"], "Ib")

// 		case "A4Ib":
// 			StellarDataMap["A4Ib"] = append(StellarDataMap["A4Ib"], "A")
// 			StellarDataMap["A4Ib"] = append(StellarDataMap["A4Ib"], "4")
// 			StellarDataMap["A4Ib"] = append(StellarDataMap["A4Ib"], "Ib")

// 		case "A5Ib":
// 			StellarDataMap["A5Ib"] = append(StellarDataMap["A5Ib"], "A")
// 			StellarDataMap["A5Ib"] = append(StellarDataMap["A5Ib"], "5")
// 			StellarDataMap["A5Ib"] = append(StellarDataMap["A5Ib"], "Ib")

// 		case "A6Ib":
// 			StellarDataMap["A6Ib"] = append(StellarDataMap["A6Ib"], "A")
// 			StellarDataMap["A6Ib"] = append(StellarDataMap["A6Ib"], "6")
// 			StellarDataMap["A6Ib"] = append(StellarDataMap["A6Ib"], "Ib")

// 		case "A7Ib":
// 			StellarDataMap["A7Ib"] = append(StellarDataMap["A7Ib"], "A")
// 			StellarDataMap["A7Ib"] = append(StellarDataMap["A7Ib"], "7")
// 			StellarDataMap["A7Ib"] = append(StellarDataMap["A7Ib"], "Ib")

// 		case "A8Ib":
// 			StellarDataMap["A8Ib"] = append(StellarDataMap["A8Ib"], "A")
// 			StellarDataMap["A8Ib"] = append(StellarDataMap["A8Ib"], "8")
// 			StellarDataMap["A8Ib"] = append(StellarDataMap["A8Ib"], "Ib")

// 		case "A9Ib":
// 			StellarDataMap["A9Ib"] = append(StellarDataMap["A9Ib"], "A")
// 			StellarDataMap["A9Ib"] = append(StellarDataMap["A9Ib"], "9")
// 			StellarDataMap["A9Ib"] = append(StellarDataMap["A9Ib"], "Ib")

// 		case "A0II":
// 			StellarDataMap["A0II"] = append(StellarDataMap["A0II"], "A")
// 			StellarDataMap["A0II"] = append(StellarDataMap["A0II"], "0")
// 			StellarDataMap["A0II"] = append(StellarDataMap["A0II"], "II")

// 		case "A1II":
// 			StellarDataMap["A1II"] = append(StellarDataMap["A1II"], "A")
// 			StellarDataMap["A1II"] = append(StellarDataMap["A1II"], "1")
// 			StellarDataMap["A1II"] = append(StellarDataMap["A1II"], "II")

// 		case "A2II":
// 			StellarDataMap["A2II"] = append(StellarDataMap["A2II"], "A")
// 			StellarDataMap["A2II"] = append(StellarDataMap["A2II"], "2")
// 			StellarDataMap["A2II"] = append(StellarDataMap["A2II"], "II")

// 		case "A3II":
// 			StellarDataMap["A3II"] = append(StellarDataMap["A3II"], "A")
// 			StellarDataMap["A3II"] = append(StellarDataMap["A3II"], "3")
// 			StellarDataMap["A3II"] = append(StellarDataMap["A3II"], "II")

// 		case "A4II":
// 			StellarDataMap["A4II"] = append(StellarDataMap["A4II"], "A")
// 			StellarDataMap["A4II"] = append(StellarDataMap["A4II"], "4")
// 			StellarDataMap["A4II"] = append(StellarDataMap["A4II"], "II")

// 		case "A5II":
// 			StellarDataMap["A5II"] = append(StellarDataMap["A5II"], "A")
// 			StellarDataMap["A5II"] = append(StellarDataMap["A5II"], "5")
// 			StellarDataMap["A5II"] = append(StellarDataMap["A5II"], "II")

// 		case "A6II":
// 			StellarDataMap["A6II"] = append(StellarDataMap["A6II"], "A")
// 			StellarDataMap["A6II"] = append(StellarDataMap["A6II"], "6")
// 			StellarDataMap["A6II"] = append(StellarDataMap["A6II"], "II")

// 		case "A7II":
// 			StellarDataMap["A7II"] = append(StellarDataMap["A7II"], "A")
// 			StellarDataMap["A7II"] = append(StellarDataMap["A7II"], "7")
// 			StellarDataMap["A7II"] = append(StellarDataMap["A7II"], "II")

// 		case "A8II":
// 			StellarDataMap["A8II"] = append(StellarDataMap["A8II"], "A")
// 			StellarDataMap["A8II"] = append(StellarDataMap["A8II"], "8")
// 			StellarDataMap["A8II"] = append(StellarDataMap["A8II"], "II")

// 		case "A9II":
// 			StellarDataMap["A9II"] = append(StellarDataMap["A9II"], "A")
// 			StellarDataMap["A9II"] = append(StellarDataMap["A9II"], "9")
// 			StellarDataMap["A9II"] = append(StellarDataMap["A9II"], "II")

// 		case "A0III":
// 			StellarDataMap["A0III"] = append(StellarDataMap["A0III"], "A")
// 			StellarDataMap["A0III"] = append(StellarDataMap["A0III"], "0")
// 			StellarDataMap["A0III"] = append(StellarDataMap["A0III"], "III")

// 		case "A1III":
// 			StellarDataMap["A1III"] = append(StellarDataMap["A1III"], "A")
// 			StellarDataMap["A1III"] = append(StellarDataMap["A1III"], "1")
// 			StellarDataMap["A1III"] = append(StellarDataMap["A1III"], "III")

// 		case "A2III":
// 			StellarDataMap["A2III"] = append(StellarDataMap["A2III"], "A")
// 			StellarDataMap["A2III"] = append(StellarDataMap["A2III"], "2")
// 			StellarDataMap["A2III"] = append(StellarDataMap["A2III"], "III")

// 		case "A3III":
// 			StellarDataMap["A3III"] = append(StellarDataMap["A3III"], "A")
// 			StellarDataMap["A3III"] = append(StellarDataMap["A3III"], "3")
// 			StellarDataMap["A3III"] = append(StellarDataMap["A3III"], "III")

// 		case "A4III":
// 			StellarDataMap["A4III"] = append(StellarDataMap["A4III"], "A")
// 			StellarDataMap["A4III"] = append(StellarDataMap["A4III"], "4")
// 			StellarDataMap["A4III"] = append(StellarDataMap["A4III"], "III")

// 		case "A5III":
// 			StellarDataMap["A5III"] = append(StellarDataMap["A5III"], "A")
// 			StellarDataMap["A5III"] = append(StellarDataMap["A5III"], "5")
// 			StellarDataMap["A5III"] = append(StellarDataMap["A5III"], "III")

// 		case "A6III":
// 			StellarDataMap["A6III"] = append(StellarDataMap["A6III"], "A")
// 			StellarDataMap["A6III"] = append(StellarDataMap["A6III"], "6")
// 			StellarDataMap["A6III"] = append(StellarDataMap["A6III"], "III")

// 		case "A7III":
// 			StellarDataMap["A7III"] = append(StellarDataMap["A7III"], "A")
// 			StellarDataMap["A7III"] = append(StellarDataMap["A7III"], "7")
// 			StellarDataMap["A7III"] = append(StellarDataMap["A7III"], "III")

// 		case "A8III":
// 			StellarDataMap["A8III"] = append(StellarDataMap["A8III"], "A")
// 			StellarDataMap["A8III"] = append(StellarDataMap["A8III"], "8")
// 			StellarDataMap["A8III"] = append(StellarDataMap["A8III"], "III")

// 		case "A9III":
// 			StellarDataMap["A9III"] = append(StellarDataMap["A9III"], "A")
// 			StellarDataMap["A9III"] = append(StellarDataMap["A9III"], "9")
// 			StellarDataMap["A9III"] = append(StellarDataMap["A9III"], "III")

// 		case "A0IV":
// 			StellarDataMap["A0IV"] = append(StellarDataMap["A0IV"], "A")
// 			StellarDataMap["A0IV"] = append(StellarDataMap["A0IV"], "0")
// 			StellarDataMap["A0IV"] = append(StellarDataMap["A0IV"], "IV")

// 		case "A1IV":
// 			StellarDataMap["A1IV"] = append(StellarDataMap["A1IV"], "A")
// 			StellarDataMap["A1IV"] = append(StellarDataMap["A1IV"], "1")
// 			StellarDataMap["A1IV"] = append(StellarDataMap["A1IV"], "IV")

// 		case "A2IV":
// 			StellarDataMap["A2IV"] = append(StellarDataMap["A2IV"], "A")
// 			StellarDataMap["A2IV"] = append(StellarDataMap["A2IV"], "2")
// 			StellarDataMap["A2IV"] = append(StellarDataMap["A2IV"], "IV")

// 		case "A3IV":
// 			StellarDataMap["A3IV"] = append(StellarDataMap["A3IV"], "A")
// 			StellarDataMap["A3IV"] = append(StellarDataMap["A3IV"], "3")
// 			StellarDataMap["A3IV"] = append(StellarDataMap["A3IV"], "IV")

// 		case "A4IV":
// 			StellarDataMap["A4IV"] = append(StellarDataMap["A4IV"], "A")
// 			StellarDataMap["A4IV"] = append(StellarDataMap["A4IV"], "4")
// 			StellarDataMap["A4IV"] = append(StellarDataMap["A4IV"], "IV")

// 		case "A5IV":
// 			StellarDataMap["A5IV"] = append(StellarDataMap["A5IV"], "A")
// 			StellarDataMap["A5IV"] = append(StellarDataMap["A5IV"], "5")
// 			StellarDataMap["A5IV"] = append(StellarDataMap["A5IV"], "IV")

// 		case "A6IV":
// 			StellarDataMap["A6IV"] = append(StellarDataMap["A6IV"], "A")
// 			StellarDataMap["A6IV"] = append(StellarDataMap["A6IV"], "6")
// 			StellarDataMap["A6IV"] = append(StellarDataMap["A6IV"], "IV")

// 		case "A7IV":
// 			StellarDataMap["A7IV"] = append(StellarDataMap["A7IV"], "A")
// 			StellarDataMap["A7IV"] = append(StellarDataMap["A7IV"], "7")
// 			StellarDataMap["A7IV"] = append(StellarDataMap["A7IV"], "IV")

// 		case "A8IV":
// 			StellarDataMap["A8IV"] = append(StellarDataMap["A8IV"], "A")
// 			StellarDataMap["A8IV"] = append(StellarDataMap["A8IV"], "8")
// 			StellarDataMap["A8IV"] = append(StellarDataMap["A8IV"], "IV")

// 		case "A9IV":
// 			StellarDataMap["A9IV"] = append(StellarDataMap["A9IV"], "A")
// 			StellarDataMap["A9IV"] = append(StellarDataMap["A9IV"], "9")
// 			StellarDataMap["A9IV"] = append(StellarDataMap["A9IV"], "IV")

// 		case "A0V":
// 			StellarDataMap["A0V"] = append(StellarDataMap["A0V"], "A")
// 			StellarDataMap["A0V"] = append(StellarDataMap["A0V"], "0")
// 			StellarDataMap["A0V"] = append(StellarDataMap["A0V"], "V")

// 		case "A1V":
// 			StellarDataMap["A1V"] = append(StellarDataMap["A1V"], "A")
// 			StellarDataMap["A1V"] = append(StellarDataMap["A1V"], "1")
// 			StellarDataMap["A1V"] = append(StellarDataMap["A1V"], "V")

// 		case "A2V":
// 			StellarDataMap["A2V"] = append(StellarDataMap["A2V"], "A")
// 			StellarDataMap["A2V"] = append(StellarDataMap["A2V"], "2")
// 			StellarDataMap["A2V"] = append(StellarDataMap["A2V"], "V")

// 		case "A3V":
// 			StellarDataMap["A3V"] = append(StellarDataMap["A3V"], "A")
// 			StellarDataMap["A3V"] = append(StellarDataMap["A3V"], "3")
// 			StellarDataMap["A3V"] = append(StellarDataMap["A3V"], "V")

// 		case "A4V":
// 			StellarDataMap["A4V"] = append(StellarDataMap["A4V"], "A")
// 			StellarDataMap["A4V"] = append(StellarDataMap["A4V"], "4")
// 			StellarDataMap["A4V"] = append(StellarDataMap["A4V"], "V")

// 		case "A5V":
// 			StellarDataMap["A5V"] = append(StellarDataMap["A5V"], "A")
// 			StellarDataMap["A5V"] = append(StellarDataMap["A5V"], "5")
// 			StellarDataMap["A5V"] = append(StellarDataMap["A5V"], "V")

// 		case "A6V":
// 			StellarDataMap["A6V"] = append(StellarDataMap["A6V"], "A")
// 			StellarDataMap["A6V"] = append(StellarDataMap["A6V"], "6")
// 			StellarDataMap["A6V"] = append(StellarDataMap["A6V"], "V")

// 		case "A7V":
// 			StellarDataMap["A7V"] = append(StellarDataMap["A7V"], "A")
// 			StellarDataMap["A7V"] = append(StellarDataMap["A7V"], "7")
// 			StellarDataMap["A7V"] = append(StellarDataMap["A7V"], "V")

// 		case "A8V":
// 			StellarDataMap["A8V"] = append(StellarDataMap["A8V"], "A")
// 			StellarDataMap["A8V"] = append(StellarDataMap["A8V"], "8")
// 			StellarDataMap["A8V"] = append(StellarDataMap["A8V"], "V")

// 		case "A9V":
// 			StellarDataMap["A9V"] = append(StellarDataMap["A9V"], "A")
// 			StellarDataMap["A9V"] = append(StellarDataMap["A9V"], "9")
// 			StellarDataMap["A9V"] = append(StellarDataMap["A9V"], "V")

// 		}
// 	}
// }
