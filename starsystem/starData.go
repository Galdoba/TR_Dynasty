package starsystem

import (
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

//StarDiameter - возвращает диаметр звезды полученного класса. 1,00 = диаметр солнца
func StarDiameter(stellar string) float64 {
	massesMap := make(map[string][]float64)
	massesMap["b0"] = []float64{60, 50, 30, 25, 20, 18, -1, 0.26}
	massesMap["b5"] = []float64{30, 25, 20, 15, 10, 6.5, -1, 0.26}
	massesMap["a0"] = []float64{18, 16, 14, 12, 6, 3.2, -1, 0.36}
	massesMap["a5"] = []float64{15, 13, 11, 9, 4, 2.1, -1, 0.36}
	massesMap["f0"] = []float64{13, 12, 10, 8, 2.5, 1.7, -1, 0.42}
	massesMap["f5"] = []float64{12, 10, 8.1, 5, 2, 13, 0.8, 0.42}
	massesMap["g0"] = []float64{12, 10, 8.1, 2.5, 1.75, 1.04, 0.6, 0.63}
	massesMap["g5"] = []float64{13, 12, 10, 3.2, 2, 0.94, 0.528, 0.63}
	massesMap["k0"] = []float64{14, 13, 11, 4, 2.3, 0.825, 0.43, 0.83}
	massesMap["k5"] = []float64{18, 16, 14, 5, -1, 0.57, 0.33, 0.83}
	massesMap["m0"] = []float64{20, 16, 14, 6.3, -1, 0.489, 0.154, 1.11}
	massesMap["m5"] = []float64{25, 20, 16, 7.4, -1, 0.331, 0.104, 1.11}
	massesMap["m9"] = []float64{30, 25, 18, 9.2, -1, 0.215, 0.058, 1.11}
	class := strings.ToLower(string([]byte(stellar)[0]))
	dig := string([]byte(stellar)[1])
	size := ""
	data := strings.Split(stellar, " ")
	if len(data) < 2 {
		size = "d"
	} else {
		size = strings.ToLower(data[1])
	}
	num, _ := strconv.Atoi(dig)
	switch num {
	case 0, 1, 2, 3, 4:
		dig = "0"
	case 5, 6, 7, 8, 9:
		dig = "5"
	}
	if class == "m" && num == 9 {
		dig = "9"
	}
	low := class + dig
	high := ""
	grad := 5
	offset := num % 5
	switch low {
	case "b0":
		high = "b5"
	case "b5":
		high = "a0"
	case "a0":
		high = "a5"
	case "a5":
		high = "f0"
	case "f0":
		high = "f5"
	case "f5":
		high = "g0"
	case "g0":
		high = "g5"
	case "g5":
		high = "k0"
	case "k0":
		high = "k5"
	case "k5":
		high = "m0"
	case "m0":
		high = "m5"
	case "m5":
		high = "m9"
	case "m9":
		high = "m9"
		grad = 4
	}
	index := 0
	switch size {
	case "ia":
		index = 0
	case "ib":
		index = 1
	case "ii":
		index = 2
	case "iii":
		index = 3
	case "iv":
		index = 4
	case "v":
		index = 5
	case "vi":
		index = 6
	case "d":
		index = 7
	}
	lowVal := massesMap[low][index]
	if lowVal < 0 {
		lowVal = massesMap[low][5]
	}
	highVal := massesMap[high][index]
	if highVal < 0 {
		highVal = massesMap[high][5]
	}
	mass := aproximate(offset, grad, lowVal, highVal)

	return mass
}

func aproximate(offset, grad int, lowVal, highVal float64) float64 {
	// if lowVal == -1 || highVal == -1 {
	// 	return -999.999
	// }
	if offset < 0 || offset > grad {
		return -999.999
	}
	step := (highVal - lowVal) / float64(grad)
	r := lowVal + (step * float64(offset))
	r = utils.RoundFloat64(r, 3)
	return r
}

/*



 */
