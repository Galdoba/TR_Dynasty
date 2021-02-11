package astronomical

import (
	"fmt"
	"strings"
)

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
