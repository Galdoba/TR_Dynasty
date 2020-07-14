package system

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/utils"
)

/*
Sector										Astergoth
	-system									Astergoth B0323
		--name
		--position
		--stars								Astergoth B0323 Alpha
			---name
			---Classification
			---position
			---hab zone						Astergoth B0323 Alpha 2
				----name
				----planet					Astergoth B0323 Alpha 2
					-----worldData
					-----name
					-----position
					-----satellite			Astergoth B0323 Alpha 2 Ay
						------worldData
						------name
						------position

Astergoth B0323 Alpha 2 Ay
*/

func TestName(name string) {
	valid := worldNameValid(name)
	fmt.Println(name, " = ", valid)

}

func worldNameValid(worldname string) bool {
	nameParts := strings.Split(worldname, "  ")
	if len(nameParts) > 5 || len(nameParts) < 1 {
		return false
	}
	for i, val := range nameParts {
		switch i {
		case 1:
			if len([]byte(val)) != 5 {
				return false
			}
		case 2:
			if TrvCore.GreekToNum(val) == -1 {
				return false
			}
		case 3:
			_, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
		case 4:
			if TrvCore.AnglicToNum(val) == -1 {
				return false
			}
		}
	}
	return true
}

//Sector - описывает сектор, расположение звездных систем и межзвездную навигацию (Уровень хексов)
type Sector struct{}

//StarSystem - описывает звезды, планеты и спутники в рамках одного хекса
type StarSystem struct {
	//ssMap map[orbitName]*Planet
	//ssOrbits map[float32]uwp (string)
}

//Planet - описывает поверхность звездного тела
type Planet struct{}

func NewStarSystem(systemName string) StarSystem {
	utils.SetSeed(utils.SeedFromString(systemName))
	//preData
	gg := rollGG()
	belts := rollBelts()
	other := utils.RollDice("2d6")
	totalWorlds := gg + belts + other
	fmt.Println(systemName+" has", totalWorlds, "worlds")
	fmt.Println(systemName+" has", gg, "GG and", belts, "belts")

	starword := generateStarWord()
	starCodes := starTypesFromCode(starword)
	fmt.Println("There are total of", len(starword), "stars", "|", starword)
	for i, pos := range starCodes {
		starclass := GenerateStarClass(pos)
		starName := systemName + " " + TrvCore.NumToGreek(i)
		fmt.Println("Star:", starName, " has class", starclass)
	}

	ss := StarSystem{}
	return ss
}

func rollStarOrbit(starType string) int {
	switch starType {
	default:
		return -1
	case "C":
		return utils.RollDice("d6", -1)
	case "N":
		return utils.RollDice("d6", 5)
	case "F":
		return utils.RollDice("d6", 11)
	}
}

func rollGG() int {
	gg := utils.RollDice("2d6")
	return utils.BoundInt((gg/2)-2, 0, 5)
}

func rollBelts() int {
	return utils.BoundInt(utils.RollDice("d6", -3), 0, 6)
}

func orbitFor(pos string) int {
	switch pos {
	case "P":
		return -1
	case "C":
		return utils.RollDice("d6", -1)
	case "N":
		return utils.RollDice("d6", 5)
	default:
		return utils.RollDice("d6", 11)
	}
}

func starPositionsS() []string {
	return []string{
		"P",
		"C",
		"N",
		"F",
	}
}

func rollPlanetsPresent() int {
	return utils.RollDice("2d6")
}

func starTypesFromCode(stCode string) (starSlice []string) {
	if strings.Contains(stCode, "P") {
		starSlice = append(starSlice, "P")
	}
	if strings.Contains(stCode, "Pc") {
		starSlice = append(starSlice, "Pc")
	}
	if strings.Contains(stCode, "C") {
		starSlice = append(starSlice, "C")
	}
	if strings.Contains(stCode, "Cc") {
		starSlice = append(starSlice, "Cc")
	}
	if strings.Contains(stCode, "N") {
		starSlice = append(starSlice, "N")
	}
	if strings.Contains(stCode, "Nc") {
		starSlice = append(starSlice, "Nc")
	}
	if strings.Contains(stCode, "F") {
		starSlice = append(starSlice, "F")
	}
	if strings.Contains(stCode, "Fc") {
		starSlice = append(starSlice, "Fc")
	}
	return starSlice
}

func generateStarWord() string {
	word := "P"
	if TrvCore.Flux() > 2 {
		word += "c"
	}
	if TrvCore.Flux() > 2 {
		word += "C"
		if TrvCore.Flux() > 2 {
			word += "c"
		}
	}
	if TrvCore.Flux() > 2 {
		word += "N"
		if TrvCore.Flux() > 2 {
			word += "c"
		}
	}
	if TrvCore.Flux() > 2 {
		word += "F"
		if TrvCore.Flux() > 2 {
			word += "c"
		}
	}
	return word
}

func GenerateStarClass(pos string) string {
	spDM := 0
	// switch utils.RollDice("2d6") { пока без них
	// case 2, 3:
	// 	spDM--
	// case 11, 12:
	// 	spDM++
	// }
	fl := TrvCore.Flux()
	if pos != "P" {
		spDM = utils.RollDice("d6", -1)
	}
	i := utils.BoundInt(spDM+fl+6, 0, 14)
	spArray := []string{"OB", "A", "A", "F", "F", "G", "G", "K", "K", "M", "M", "M", "BD", "BD", "BD"}
	stSp := spArray[i]
	if stSp == "OB" {
		stSp = utils.RandomFromList([]string{"O", "B"})
	}
	if stSp == "BD" {
		return stSp
	}
	size := ""
	szFlux := TrvCore.Flux()
	if pos != "P" {
		szFlux = szFlux + utils.RollDice("d6", 2)
	}
	j := utils.BoundInt(szFlux+5, 0, 10)
	k := 0
	switch stSp {
	case "B":
		k = 1
	case "A":
		k = 2
	case "F":
		k = 3
	case "G":
		k = 4
	case "K":
		k = 5
	case "M":
		k = 6
	}
	sizeArr := []string{}
	switch j {
	case 0:
		sizeArr = []string{"Ia", "Ia", "Ia", "II", "II", "II", "II"}
	case 1:
		sizeArr = []string{"Ib", "Ib", "Ib", "III", "III", "III", "II"}
	case 2:
		sizeArr = []string{"II", "II", "II", "IV", "IV", "IV", "II"}
	case 3:
		sizeArr = []string{"III", "III", "III", "V", "V", "V", "III"}
	case 4:
		sizeArr = []string{"III", "III", "IV", "V", "V", "V", "V"}
	case 5:
		sizeArr = []string{"III", "III", "V", "V", "V", "V", "V"}
	case 6:
		sizeArr = []string{"V", "III", "V", "V", "V", "V", "V"}
	case 7:
		sizeArr = []string{"V", "V", "V", "V", "V", "V", "V"}
	case 8:
		sizeArr = []string{"V", "V", "V", "V", "V", "V", "V"}
	case 9:
		sizeArr = []string{"IV", "IV", "V", "VI", "VI", "VI", "VI"}
	case 10:
		sizeArr = []string{"D", "D", "D", "D", "D", "D", "D"}
	}
	size = sizeArr[k]
	if size == "D" {
		return stSp + "D"
	}
	decInt := utils.RollDice("d10", -1)
	dec := strconv.Itoa(decInt)
	if size == "IV" && stSp == "K" && decInt > 4 {
		size = "V"
	}
	if size == "IV" && stSp == "M" {
		size = "V"
	}
	if size == "VI" && stSp == "A" {
		size = "V"
	}
	if size == "VI" && stSp == "F" && decInt < 5 {
		size = "V"
	}
	return stSp + dec + size
}

func orbitOrderList() []string {
	return []string{"P", "Pc", "C", "Cc", "N", "Nc", "F", "Fc"}
}

func getStarSpectral(starCode string) string {
	stSp := ""
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
		stSp = "M"
	}
	if strings.Contains(starCode, "M") {
		stSp = "M"
	}
	if strings.Contains(starCode, "BD") {
		stSp = "BD"
	}
	return stSp
}

func getStarSize(starCode string) string {
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
func getStarDecimal(starCode string) string {
	stDec := ""
	for i := 0; i < 10; i++ {
		if strings.Contains(starCode, strconv.Itoa(i)) {
			return strconv.Itoa(i)
		}
	}
	return stDec
}

func getHZ(star string) int {
	spectral := getStarSpectral(star)
	size := getStarSize(star)
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
	return hzMap[spectral+size]

}

func GenerateSateliteOrbits(totalOrbits int) []string {
	var orbitOrder []int
	var orb int
	if totalOrbits < 0 {
		return []string{"None"}
	}
	if totalOrbits == 0 {
		return []string{"Ring"}
	}
	if totalOrbits > 22 {
		totalOrbits = 22
	}
	for len(orbitOrder) < totalOrbits {
		if TrvCore.Roll2D() < 8 {
			orb = 6 + TrvCore.Flux()
		} else {
			orb = 18 + TrvCore.Flux()
		}
		orbitOrder = utils.AppendUniqueInt(orbitOrder, orb)

	}
	sort.Ints(orbitOrder)
	var namedOrder []string
	for i := range orbitOrder {
		namedOrder = append(namedOrder, TrvCore.NumToAnglic(orbitOrder[i]))
	}
	return namedOrder
}

// 2D System Type

/*
P Pc 0 1 2 ... 19   19
if
C Cc 0 1 2 ... 5    2
if
N Nc 0 1 2 ... 11   9
if
F Fc 0 1 2 ... 17   14

12+1+3+5





*/
