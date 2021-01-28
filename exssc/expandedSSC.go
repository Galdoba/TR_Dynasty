package exssc

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/Galdoba/TR_Dynasty/TrvCore"

// 	"github.com/Galdoba/TR_Dynasty/Astrogation"
// 	system "github.com/Galdoba/TR_Dynasty/starsystem"
// 	"github.com/Galdoba/utils"
// )

// type systemData struct {
// 	star      []starSyst
// 	hexName   string
// 	hexCoords string
// }

// type starSyst struct {
// 	starName  string
// 	starClass string
// 	tpb       int
// 	hz        int
// 	js        int
// 	pl        []planetaryBody
// }

// func newSystemData(hexName string, hexCoords string) systemData {
// 	sD := systemData{}
// 	sD.hexName = hexName
// 	sD.hexCoords = hexCoords
// 	return sD
// }

// func AddStarTo(sys systemData, star starSyst) systemData {
// 	i := len(sys.star)
// 	star.starName = sys.hexName + " " + TrvCore.NumToGreek(i)
// 	sys.star = append(sys.star, star)
// 	return sys
// }

// func newStarSyst(starCode string) starSyst {
// 	stSys := starSyst{}
// 	stSys.starClass = starCode
// 	stSys.tpb = totalPlanetaryBodies(0)
// 	stSys.hz = getHZ(starCode)
// 	stSys.js = Astrogation.JumpDriveLimitStar(starCode)
// 	return stSys
// }

// func NameStar(starSystem starSyst, newName string) starSyst {
// 	starSystem.starName = newName
// 	return starSystem
// }

// func getHZ(star string) int {
// 	spectral := getStarSpectral(star)
// 	size := getStarSize(star)
// 	hzMap := make(map[string]int)
// 	hzMap["OIa"] = 15
// 	hzMap["OIb"] = 15
// 	hzMap["OII"] = 14
// 	hzMap["OIII"] = 13
// 	hzMap["OIV"] = 12
// 	hzMap["OV"] = 11
// 	hzMap["OVI"] = -1
// 	hzMap["OD"] = 1
// 	hzMap["BIa"] = 13
// 	hzMap["BIb"] = 13
// 	hzMap["BII"] = 12
// 	hzMap["BIII"] = 11
// 	hzMap["BIV"] = 10
// 	hzMap["BV"] = 9
// 	hzMap["BVI"] = -1
// 	hzMap["BD"] = 0
// 	hzMap["AIa"] = 12
// 	hzMap["AIb"] = 11
// 	hzMap["AII"] = 9
// 	hzMap["AIII"] = 7
// 	hzMap["AIV"] = 7
// 	hzMap["AV"] = 7
// 	hzMap["AVI"] = -1
// 	hzMap["AD"] = 0
// 	hzMap["FIa"] = 11
// 	hzMap["FIb"] = 10
// 	hzMap["FII"] = 9
// 	hzMap["FIII"] = 6
// 	hzMap["FIV"] = 6
// 	hzMap["FV"] = 4
// 	hzMap["FVI"] = 3
// 	hzMap["FD"] = 0
// 	hzMap["GIa"] = 12
// 	hzMap["GIb"] = 10
// 	hzMap["GII"] = 9
// 	hzMap["GIII"] = 7
// 	hzMap["GIV"] = 5
// 	hzMap["GV"] = 3
// 	hzMap["GVI"] = 2
// 	hzMap["GD"] = 0
// 	hzMap["KIa"] = 12
// 	hzMap["KIb"] = 10
// 	hzMap["KII"] = 9
// 	hzMap["KIII"] = 8
// 	hzMap["KIV"] = 5
// 	hzMap["KV"] = 2
// 	hzMap["KVI"] = 1
// 	hzMap["KD"] = 0
// 	hzMap["MIa"] = 12
// 	hzMap["MIb"] = 11
// 	hzMap["MII"] = 10
// 	hzMap["MIII"] = 9
// 	hzMap["MIV"] = -1
// 	hzMap["MV"] = 0
// 	hzMap["MVI"] = 0
// 	hzMap["MD"] = 0
// 	return hzMap[spectral+size]

// }

// func getStarSize(starCode string) string {
// 	stSp := "Ia"
// 	if strings.Contains(starCode, "Ib") {
// 		stSp = "Ib"
// 	}
// 	if strings.Contains(starCode, "II") {
// 		stSp = "II"
// 	}
// 	if strings.Contains(starCode, "III") {
// 		stSp = "III"
// 	}
// 	if strings.Contains(starCode, "V") {
// 		stSp = "V"
// 	}
// 	if strings.Contains(starCode, "IV") {
// 		stSp = "IV"
// 	}
// 	if strings.Contains(starCode, "VI") {
// 		stSp = "VI"
// 	}
// 	if strings.Contains(starCode, "D") {
// 		stSp = "D"
// 	}
// 	return stSp
// }
// func getStarDecimal(starCode string) string {
// 	stDec := ""
// 	for i := 0; i < 10; i++ {
// 		if strings.Contains(starCode, strconv.Itoa(i)) {
// 			return strconv.Itoa(i)
// 		}
// 	}
// 	return stDec
// }
// func getStarSpectral(starCode string) string {
// 	stSp := ""
// 	if strings.Contains(starCode, "O") {
// 		stSp = "O"
// 	}
// 	if strings.Contains(starCode, "B") {
// 		stSp = "B"
// 	}
// 	if strings.Contains(starCode, "A") {
// 		stSp = "A"
// 	}
// 	if strings.Contains(starCode, "F") {
// 		stSp = "F"
// 	}
// 	if strings.Contains(starCode, "G") {
// 		stSp = "G"
// 	}
// 	if strings.Contains(starCode, "K") {
// 		stSp = "M"
// 	}
// 	if strings.Contains(starCode, "M") {
// 		stSp = "M"
// 	}
// 	if strings.Contains(starCode, "BD") {
// 		stSp = "BD"
// 	}
// 	return stSp
// }

// func PrintStarData(star starSyst) {
// 	fmt.Println("------")
// 	fmt.Println("Star Name", star.starName)
// 	fmt.Println("Star Class", star.starClass)

// 	fmt.Println("Star tpb", star.tpb)
// 	fmt.Println("Star hz", star.hz)
// 	fmt.Println("Star js", star.js)
// }

// type planetaryBody struct {
// 	name         string
// 	plType       string
// 	orbit        int
// 	habitability int
// 	sOrbits      []string
// 	statelites   []satelite
// 	shadowed     bool
// }

// func newPlanet(orbit int) planetaryBody {
// 	pl := planetaryBody{}
// 	pl.orbit = orbit
// 	return pl
// }

// func syncDataFrom(star starSyst, pl planetaryBody) planetaryBody {
// 	pl.name = star.starName + " " + strconv.Itoa(pl.orbit)
// 	pl.habitability = pl.orbit - star.hz
// 	if pl.orbit <= star.js {
// 		pl.shadowed = true
// 	}

// 	return pl
// }

// func rollPlanetaryType(pl planetaryBody) planetaryBody {
// 	flType := TrvCore.Flux()
// 	switch flType {
// 	default:
// 		pl.plType = innerType()
// 		if pl.habitability > 1 {
// 			pl.plType = outerType()
// 		}
// 	case -4, 4:
// 		pl.plType = "Gas Gigant"
// 	case -5, 5:
// 		pl.plType = "Asteroid Belt"
// 		return pl
// 	}
// 	return pl
// }

// func rollSatOrbits(pl planetaryBody) int {
// 	s := 0
// 	switch pl.orbit {
// 	default:
// 		if pl.orbit < -1 {
// 			s = utils.RollDice("d6", -5)
// 		}
// 		if pl.orbit > 1 {
// 			s = utils.RollDice("d6", -3)
// 		}
// 	case -1, 0, 1:
// 		s = utils.RollDice("d6", -4)
// 	}
// 	if pl.plType == "Gas Gigant" {
// 		s = utils.RollDice("d6", -1)
// 	}
// 	if pl.plType == "Asteroid Belt" {
// 		s = -1
// 	}
// 	return s
// }

// func addPlanet(star starSyst, pl planetaryBody) starSyst {
// 	star.pl = append(star.pl, pl)
// 	return star
// }

// type satelite struct {
// 	name  string
// 	sType string
// }

// func newSatelite(pl planetaryBody) satelite {
// 	sat := satelite{}
// 	sat.sType = innerSateliteType()
// 	if pl.habitability > 1 {
// 		sat.sType = outerSateliteType()
// 	}
// 	return sat
// }

// func nameSatelite(sat satelite, newName string) satelite {
// 	sat.name = newName
// 	return sat
// }

// func generateSatelites(pl planetaryBody) planetaryBody {
// 	so := rollSatOrbits(pl)
// 	satPositions := system.GenerateSateliteOrbits(so)
// 	for current := range satPositions {
// 		if satPositions[current] == "None" || satPositions[current] == "Ring" {
// 			break
// 		}
// 		sat := newSatelite(pl)
// 		sat = nameSatelite(sat, pl.name+" "+satPositions[current])
// 		pl.statelites = append(pl.statelites, sat)
// 	}
// 	return pl
// }

// func Test() {
// 	sysData := newSystemData("Argos", "A0203")
// 	ssType := starSystemType()
// 	fmt.Println("")
// 	fmt.Println("Generating System:", ssType)
// 	//hexName := "Amani"
// 	starCodes := strings.Split(ssType, " ")
// 	for i := range starCodes {
// 		starSystem := newStarSyst(starCodes[i])
// 		sysData = AddStarTo(sysData, starSystem)
// 		fmt.Println(sysData)
// 	}
// 	for i := range sysData.star {
// 		star := sysData.star[i]
// 		PrintStarData(star)
// 		for o := 0; o < sysData.star[i].tpb; o++ {
// 			pl := newPlanet(o)
// 			pl = syncDataFrom(sysData.star[i], pl)
// 			pl = rollPlanetaryType(pl)
// 			pl = generateSatelites(pl)
// 			///////////
// 			//TODO: Сателиты
// 			//////////
// 			sysData.star[i] = addPlanet(sysData.star[i], pl)

// 		}
// 		for j, planet := range sysData.star[i].pl {
// 			fmt.Println("	", sysData.star[i].pl[j])
// 			for k := range planet.statelites {
// 				fmt.Println("		", planet.statelites[k])
// 			}
// 		}

// 	}
// }

// func innerType() string {
// 	return utils.RandomFromList([]string{
// 		"Inferno world",
// 		"Inner world",
// 		"Big world",
// 		"Storm world",
// 		"Rad world",
// 		"Hospitable world",
// 	})
// }

// func innerSateliteType() string {
// 	return utils.RandomFromList([]string{
// 		"Inferno world",
// 		"Inner world",
// 		"Big world",
// 		"Storm world",
// 		"Rad world",
// 		"Hospitable world",
// 	})
// }

// func outerType() string {
// 	return utils.RandomFromList([]string{
// 		"Worldlet",
// 		"Ice world",
// 		"Big world",
// 		"Ice world",
// 		"Rad world",
// 		"Ice world",
// 	})
// }

// func outerSateliteType() string {
// 	return utils.RandomFromList([]string{
// 		"Worldlet",
// 		"Ice world",
// 		"Big world",
// 		"Storm world",
// 		"Rad world",
// 		"Ice world",
// 	})
// }

// func starSystemType() string {
// 	switch utils.RollDice("2d6") {
// 	case 2:
// 		return starSystemType() // + " " + starSystemSpecial()
// 	case 3, 4:
// 		//	return "Trinary (CF)" + "[" + system.GenerateStar("P") + " with " + system.GenerateStar("C") + " and " + system.GenerateStar("F")
// 		return system.GenerateStar("P") + " " + system.GenerateStar("C") + " " + system.GenerateStar("F")
// 	case 5, 6:
// 		//return "Binary (C)"
// 		return system.GenerateStar("P") + " " + system.GenerateStar("Pc")
// 	case 7, 8:
// 		//return "Solo"
// 		return system.GenerateStar("P")
// 	case 9, 10:
// 		//return "Binary (F)"
// 		return system.GenerateStar("P") + " " + system.GenerateStar("F")
// 	case 11:
// 		//return "Trinary (FC)"
// 		return system.GenerateStar("P") + " " + system.GenerateStar("F") + " " + system.GenerateStar("Fc")
// 	case 12:
// 		//return "Multiple " + starSystemType()
// 		return starSystemType() // + " " + starSystemSpecial()
// 	}
// 	return "THIS CAN'NOT BE"
// }

// func starSystemSpecial() string {
// 	switch utils.RollDice("2d6") {
// 	case 2:
// 		return starSystemHighlyUnusual() + " " + starSystemSpecial()
// 	case 3:
// 		return "Expanding Pre-Gigant"
// 	case 4, 5:
// 		return "Brown Dwarf"
// 	case 6, 7, 8:
// 		return "Empty"
// 	case 9, 10:
// 		return "White Dwarf"
// 	case 11:
// 		return "Gigant"
// 	case 12:
// 		return "Unstable"
// 	}
// 	return "THIS CAN'NOT BE"
// }

// func starSystemHighlyUnusual() string {
// 	switch utils.RollDice("2d6") {
// 	case 2:
// 		return "Black hole"
// 	case 3, 4:
// 		return "Anomalous"
// 	case 5, 6, 7, 8, 9:
// 		return "Nebula"
// 	case 10, 11:
// 		return "Highly Complex Multiple Starsystem"
// 	case 12:
// 		return "Neutron Star"
// 	}
// 	return "THIS CAN'NOT BE"
// }

// type ssystemData struct {
// 	starType string
// 	tpb      int
// }

// func totalPlanetaryBodies(dm int) int {
// 	r := utils.RollDice("2d6", dm)
// 	switch r {
// 	case 2:
// 		return 1
// 	case 3:
// 		return utils.RollDice("d3")
// 	case 4, 5:
// 		return utils.RollDice("d6", 1)
// 	case 6, 7, 8:
// 		return utils.RollDice("2d6")
// 	case 9, 10:
// 		return utils.RollDice("2d6", 3)
// 	case 11:
// 		return utils.RollDice("3d6")
// 	case 12:
// 		return utils.RollDice("3d6", 2)

// 	}
// 	return 0
// }
