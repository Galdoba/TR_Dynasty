package starsystem

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dwd"
	"github.com/Galdoba/TR_Dynasty/pkg/core/astronomical"
	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/names"
	"github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
	"github.com/Galdoba/TR_Dynasty/tab"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/constant"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"

	"github.com/Galdoba/TR_Dynasty/wrld"
)

func StarsystemTest() {
	world := wrld.PickWorld()
	ssData := From(world)
	ssData.PrintTable()
	key, _ := user.InputStr()
	for k, v := range ssData.bodyDetail {
		if v.stringKey == key {
			fmt.Println("User Selected:", k)
			fmt.Println(v)
		}
	}

}

//From - генерирует детали всех планетарных тел в системе на основе данных из SecondSurveyT5
func From(world wrld.World) SystemDetails {
	fmt.Print(world.SecondSurvey(), "\n")
	d := SystemDetails{}
	d.bodyDetail = make(map[numCode]BodyDetails)
	d.dicepool = dice.New().SetSeed(world.Name() + world.Name())
	starData := parseStellarData(world)
	stBodySlots := setupOrbitalBodySlots(starData)
	stBodySlots = d.placeMW(world, stBodySlots)
	stBodySlots = d.placeGG(world, stBodySlots)
	stBodySlots = d.placeBelts(world, stBodySlots)
	stBodySlots = d.placeOtherWorlds(world, stBodySlots)
	stBodySlots = d.assignWorldTypes(world, stBodySlots)
	stBodySlots = d.placeSatelites(world, stBodySlots)
	for k, v := range stBodySlots {
		if v == "--VALID--" {
			delete(stBodySlots, k)
		}
	}
	//fmt.Println(stBodySlots)
	for _, position := range allKeys2() {
		if v, ok := stBodySlots[position]; ok {
			bDetails := newBodyR(v, position, world)
			if position.sateliteCode() != -1 {
				bDetails.syncPlanetDistance(d)
			}
			bDetails.parentStar = starData[position.starCode()]
			bDetails.stringKey = code2string(position)
			//fmt.Println(position, v)
			//bDetails.DEBUGINFO()
			d.bodyDetail[position] = bDetails
		}
	}

	fmt.Println("DEBUG: OUTPUT INFO")

	fmt.Println(len(stBodySlots), "valid slots in total")

	//fmt.Println(starData, stBodySlots)
	return d
}

//Test -
func Test() {
	world := wrld.PickWorld()
	fmt.Println(world)
	d := From(world)
	d.PrintTable()
	fmt.Println("Enter planetary code for details:")
	key, _ := user.InputStr()
	k, err := string2code(key)
	if err != nil {
		fmt.Println(err)
	}
	body := d.bodyDetail[k]
	detailedWorldData := dwd.GenerateDetailedWorldData(body.parentStar, key, body.uwp, body.bodyType, body.nomena+body.uwp)
	detailedWorldData.PrintData()
	// if bd, ok := d.bodyDetail[k]; ok {
	// 	fmt.Println(bd.FullInfo())
	// 	fmt.Println(d.bodyDetail[k])
	// }
	fmt.Println("END PROGRAMM")
}

func (d *SystemDetails) StellarBodyDetails(code string) {

}

func (bd *BodyDetails) syncPlanetDistance(sd SystemDetails) {
	bd.orbitDistance = sd.bodyDetail[numCode{[3]int{bd.position.starCode(), bd.position.planetCode(), -1}}].orbitDistance
	if bd.position.sateliteCode() == -1 {
		return
	}
	plDiam := sd.bodyDetail[numCode{[3]int{bd.position.starCode(), bd.position.planetCode(), -1}}].diameter
	bd.orbitDistanceSat = calculateSateliteOrbitDistance(plDiam, bd.position.sateliteCode())
}

func (d *SystemDetails) placeMW(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	tc := world.TradeClassificationsSl()
	starData := parseStellarData(world)
	hz := astronomical.HabitableOrbit(starData[0])
	mwOrbit := hz
	if sliceContains(tc, "Fr") {
		mwOrbit = hz + d.dicepool.RollNext("1d6").DM(1).Sum()
	}
	if sliceContains(tc, "Co") {
		mwOrbit = hz + 1
	}
	if sliceContains(tc, "Tu") {
		mwOrbit = hz + 1
	}
	if sliceContains(tc, "Tr") {
		mwOrbit = hz - 1
	}
	if sliceContains(tc, "Ho") {
		mwOrbit = hz - 1
	}
	if sliceContains(tc, "Sa") {
		//fmt.Print("TODO: Must be Satelite\n")
		satOrb := d.rollFarSatelite()
		if sliceContains(tc, "Lk") {
			//	fmt.Print("TODO: Roll Close Satelite\n")
			satOrb = d.rollCloseSatelite()
		}
		//fmt.Print("Suggest satelite orbit ", TrvCore.NumToAnglic(satOrb), "\n")
		stBodySlots[numCode{[3]int{0, mwOrbit, satOrb}}] = "MainWorld"
		if world.GasGigants() == 0 {
			stBodySlots[numCode{[3]int{0, mwOrbit, -1}}] = "Big World"
		}
		return stBodySlots
	}

	//fmt.Print("Suggest: ", mwOrbit, "\n")
	stBodySlots[numCode{[3]int{0, mwOrbit, -1}}] = "MainWorld"
	return stBodySlots
}

func mwOrbit(stBodySlots map[numCode]string) (int, int) {
	for k, v := range stBodySlots {
		if v == "MainWorld" {
			return k.planetCode(), k.sateliteCode()
		}
	}
	return -999, -999
}

func (d *SystemDetails) placeGG(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	sggCount := 0
	totalGG := world.GasGigants()
	ggTypes := []string{}
	starData := parseStellarData(world)
	starKeys := []int{}
	starHZ := []int{}
	for i := 0; i < totalGG; i++ {
		for l := 0; l < len(starData); l++ {
			starKeys = append(starKeys, l)
			starHZ = append(starHZ, astronomical.HabitableOrbit(starData[l]))
		}
		ggTypes = append(ggTypes, "Undefined GG")
		r := d.dicepool.RollNext("2d6").Sum()
		ggTypes[i] = "LGG"
		if r == 2 || r == 3 {
			ggTypes[i] = "SGG"
			if sggCount%2 == 1 {
				ggTypes[i] = "IG"
			}
			sggCount++
		}
	}
	starKeys = starKeys[0:len(ggTypes)]
	starHZ = starHZ[0:len(ggTypes)]
	for ggIndex, ggType := range ggTypes {
		suggest := starHZ[ggIndex] + d.rollOrbitPlacement(ggType)
		if ggIndex == 0 {
			mwOr, mwsatOrb := mwOrbit(stBodySlots)
			if mwsatOrb != -1 {
				suggest = mwOr
			}
		}

		//valid := false
		for {
			if val, ok := stBodySlots[numCode{[3]int{starKeys[ggIndex], suggest, -1}}]; ok {
				if val != "--VALID--" {
					//}
					//if stBodySlots[numCode{[3]int{starKeys[ggIndex], suggest, -1}}] != "--VALID--" {
					suggest++
					continue
				}
			}
			stBodySlots[numCode{[3]int{starKeys[ggIndex], suggest, -1}}] = ggType
			break
		}
	}

	return stBodySlots
}

func (d *SystemDetails) placeBelts(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	totalBelts := world.Belts()
	starData := parseStellarData(world)
	starKeys := []int{}
	starHZ := []int{}
	for i := 0; i < totalBelts; i++ {
		for l := 0; l < len(starData); l++ {
			starKeys = append(starKeys, l)
			starHZ = append(starHZ, astronomical.HabitableOrbit(starData[l]))
		}
	}
	starKeys = starKeys[0:totalBelts]
	starHZ = starHZ[0:totalBelts]
	for _, star := range starKeys {
		suggest := starHZ[star] + d.rollOrbitPlacement("Belt")
		//valid := false
		for {
			if val, ok := stBodySlots[numCode{[3]int{starKeys[star], suggest, -1}}]; ok {
				if val != "--VALID--" {
					//}
					//if stBodySlots[numCode{[3]int{starKeys[ggIndex], suggest, -1}}] != "--VALID--" {
					suggest++
					continue
				}
			}
			stBodySlots[numCode{[3]int{starKeys[star], suggest, -1}}] = "Belt"
			break
		}
	}

	return stBodySlots
}

func (d *SystemDetails) placeOtherWorlds(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	numWo := world.NumOfWorlds()
	if numWo == "--NO DATA--" || numWo == "" {
		numWo = world.DicePool().RollNext("2d6").DM(world.Belts() + world.GasGigants() + 1).SumStr()
	}
	worldsTotal, err := strconv.Atoi(numWo)
	if err != nil {
		panic(err)
	}
	otherWorlds := worldsTotal - world.GasGigants() - world.Belts() - 1
	starData := parseStellarData(world)
	starKeys := []int{}
	for i := 0; i < otherWorlds; i++ {
		for l := 0; l < len(starData); l++ {
			starKeys = append(starKeys, l)
		}
	}
	starKeys = starKeys[0:otherWorlds]
	for index, star := range starKeys {
		suggest := d.rollOrbitPlacement("World1") // starHZ[star] +
		if index == otherWorlds-1 {
			suggest = d.rollOrbitPlacement("World2") //starHZ[star] +
		}

		for {
			if val, ok := stBodySlots[numCode{[3]int{starKeys[star], suggest, -1}}]; ok {
				if suggest < 0 {
					suggest++
					continue
				}
				if suggest < astronomical.ClosestPossibleOrbit(starData[starKeys[star]]) {
					suggest++
					continue
				}

				if val != "--VALID--" {
					suggest++
					continue
				}
			}
			stBodySlots[numCode{[3]int{starKeys[star], suggest, -1}}] = "Planet"
			break
		}
	}

	return stBodySlots
}

func keysSorted(stBodySlots map[numCode]string) []numCode {
	res := []numCode{}
	for str := 0; str < 5; str++ {
		for orb := -1; orb < 21; orb++ {
			for sat := -1; sat < 26; sat++ {
				if _, ok := stBodySlots[numCode{[3]int{str, orb, sat}}]; ok {
					res = append(res, numCode{[3]int{str, orb, sat}})
				}
			}
		}
	}
	return res
}

func (d *SystemDetails) assignWorldTypes(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	starData := parseStellarData(world)
	keys := keysSorted(stBodySlots)
	for _, k := range keys {
		if k.sateliteCode() != -1 {
			continue
		}
		if stBodySlots[k] != "Planet" {
			continue
		}
		hz := astronomical.HabitableOrbit(starData[k.starCode()])
		stBodySlots[k] = d.outerType()
		if k.planetCode() <= hz {
			stBodySlots[k] = d.innerType()
		}
	}
	return stBodySlots
}

func (d *SystemDetails) addRingAndSat(mod int) (int, int) {
	rings, sats := 0, 0
	r := d.dicepool.RollNext("1d6").DM(mod).Sum()
	for r == 0 {
		rings++
		r = d.dicepool.RollNext("1d6").DM(mod).Sum()
	}
	sats = r
	return sats, rings
}

func (d *SystemDetails) placeSatelites(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	starData := parseStellarData(world)
	keys := keysSorted(stBodySlots)
	for _, k := range keys {
		//fmt.Println("Go Key", k)
		if k.sateliteCode() != -1 {
			//	fmt.Println("Stop planet:", stBodySlots[k])
			continue
		}
		if stBodySlots[k] == "--VALID--" || stBodySlots[k] == "Star" {
			//fmt.Println("Stop Empty/Star:", stBodySlots[k])
			continue
		}
		hz := astronomical.HabitableOrbit(starData[k.starCode()])
		rings := 0
		numSat := -1
		numSat, rings = d.addRingAndSat(-4)
		if hz-k.planetCode() > 1 {
			numSat, rings = d.addRingAndSat(-3)
		}
		if hz-k.planetCode() < -1 {
			numSat, rings = d.addRingAndSat(-5)
		}
		if stBodySlots[k] == "LGG" || stBodySlots[k] == "SGG" || stBodySlots[k] == "IG" {
			numSat, rings = d.addRingAndSat(-1)
		}
		///ALT
		// if stBodySlots[k] == "LGG" {
		// 	numSat, rings = d.addRingAndSat(d.dicepool.RollNext("1d6").DM(0).Sum())
		// }
		// if stBodySlots[k] == "SGG" {
		// 	numSat, rings = d.addRingAndSat(d.dicepool.RollNext("1d6").DM(-4).Sum())
		// }
		// if stBodySlots[k] == "IG" {
		// 	numSat, rings = d.addRingAndSat(d.dicepool.RollNext("1d6").DM(-6).Sum())
		// }
		//fmt.Println("NumSat", numSat, rings, stBodySlots[k])
		for i := 0; i < rings; i++ {
			suggestOrbit := numCode{[3]int{k.starCode(), k.planetCode(), d.rollSatellitePosition()}}
			for stBodySlots[suggestOrbit] != "--VALID--" {
				suggestOrbit = numCode{[3]int{k.starCode(), k.planetCode(), d.rollSatellitePosition()}}
			}
			stBodySlots[suggestOrbit] = "Rings"
		}
		for i := 0; i < numSat; i++ {
			suggestOrbit := numCode{[3]int{k.starCode(), k.planetCode(), d.rollSatellitePosition()}}
			for stBodySlots[suggestOrbit] != "--VALID--" {
				suggestOrbit = numCode{[3]int{k.starCode(), k.planetCode(), d.rollSatellitePosition()}}
			}
			stBodySlots[suggestOrbit] = d.outerSateliteType()
			if k.planetCode()-hz < 2 {
				stBodySlots[suggestOrbit] = d.innerSateliteType()
			}
		}
	}
	return stBodySlots
}

func (d *SystemDetails) rollSatellitePosition() int {
	r := d.dicepool.FluxNext()
	switch d.dicepool.RollNext("2d6").Sum() {
	case 8, 9, 10, 11, 12:
		r += 7
	default:
		r += 20
	}
	return r
}

func setupOrbitalBodySlots(starData []string) map[numCode]string {
	numCodes := allKeys2()
	stBodySlots := make(map[numCode]string)
	for _, val := range numCodes {
		stBodySlots[val] = "--INVALID--"
	}
	for i := range starData {
		for k := range stBodySlots {
			if k.starCode() == i {
				stBodySlots[k] = "--VALID--"
			}
		}
	}
	for k, v := range stBodySlots {
		if v != "--VALID--" {
			delete(stBodySlots, k)
			continue
		}
		pc := k.planetCode()
		if pc == -1 && k.sateliteCode() == -1 && k.starCode() < len(starData) {
			stBodySlots[k] = "Star"
		}
	}
	return stBodySlots
}

func (d *SystemDetails) PrintTable() {
	tbl := tab.NewTST([]string{})
	for _, val := range allKeys2() {
		if bd, ok := d.bodyDetail[val]; ok {
			tbl.AddLine(bd.ShortInfo())
		}
	}
	tbl.PrintTable()
}

func parseStellarData(w wrld.World) []string {
	spectralData := strings.Split(w.Stellar(), " ")
	if w.Stellar() == "" {
		panic("\n\n---------------------------------------------------------\nTODO: решить что делать если не хватет официальных данных\n---------------------------------------------------------\n")
	}
	stars := []string{}
	for i, val := range spectralData {
		switch val {
		default:
			if i >= len(spectralData) {
				fmt.Println("Error: i >= len(spectralData)", i, len(spectralData))
				continue
			}
			stars = append(stars, spectralData[i]+" "+spectralData[i+1])
		case "BD":
			stars = append(stars, val)
		case "Ia", "Ib", "II", "III", "IV", "V", "VI", "D":
			continue
		}
	}
	return stars
}

/*
1. распределить доступные орбиты по звездам
2. поместить MW
3. поместить GG
4. поместить Belts
5. поместить Other Worlds
6. поместить сателиты
7. пройти по всем телам и раскидать характеристики.


*/

//SystemDetails - карта деталей планетарных тел
type SystemDetails struct {
	bodyDetail map[numCode]BodyDetails
	dicepool   *dice.Dicepool
}

func (bd *BodyDetails) DEBUGINFO() {
	fmt.Print("-------------------\n")
	fmt.Print("bd.nomena = '", bd.nomena, "' string\n")
	fmt.Print("bd.name = '", bd.name, "' string\n")
	fmt.Print("bd.uwp = '", bd.uwp, "' string\n")
	fmt.Print("bd.tags = '", bd.tags, "' string\n")
	fmt.Print("bd.bodyType = '", bd.bodyType, "' string\n")
	fmt.Print("bd.diameter = '", bd.diameter, "' float64/MegaMeters\n")
	fmt.Print("bd.orbitDistance = ", bd.orbitDistance, "' float64/au\n")
	fmt.Print("bd.jumpPointToBody = ", bd.jumpPointToBody, "' float64/MegaMeters\n")
	//fmt.Print("bd.orbitSpeed = ", bd.orbitSpeed, "' int\n")
	fmt.Print("bd.position = ", bd.position, "' numCode\n")
}

func code2string(nc numCode) string {
	s := ehex.New(nc.starCode()).String()
	p := ""
	if nc.planetCode() != -1 {
		p = ehex.New(nc.planetCode()).String()
	}
	st := ""
	if nc.sateliteCode() != -1 {
		st = ehex.New(nc.sateliteCode()).String()
	}
	return s + p + st
}

func string2code(str string) (numCode, error) {
	nc := numCode{}
	data := strings.Split(str, "")
	switch len(data) {
	default:
		nc.code = [3]int{-1, -1, -1}
		return nc, errors.New("strings lenght mismatch")
	case 3:
		nc.code = [3]int{ehex.New(data[0]).Value(), ehex.New(data[1]).Value(), ehex.New(data[2]).Value()}
	case 2:
		nc.code = [3]int{ehex.New(data[0]).Value(), ehex.New(data[1]).Value(), -1}
	case 1:
		nc.code = [3]int{ehex.New(data[0]).Value(), -1, -1}
	}

	return nc, nil
}

func newBodyR(planetType string, position numCode, w wrld.World) BodyDetails {
	bd := BodyDetails{}
	bd.position = position
	dp := w.DicePool()
	starData := parseStellarData(w)
	strCode := position.starCode()
	if planetType == "Belt" {
		planetType = constant.WTpPlanetoid
	}
	sDiam := utils.RoundFloat64(StarDiameter(starData[strCode]), 2) //диаметр звезды
	sShadow := Astrogation.StarJumpShadowAU(sDiam)                  //тень звезды в AU
	bd.calculateNomena(planetType, position, w)
	bd.bodyType = planetType
	switch planetType {
	default:
		//bd.uwp = uwp.RandomUWP(dp, planetType, w.UWP()) //TODO: Разбить функцию для создания профайла планеты и спутника (чтобы спутник не был больше чем планета)
		//if position.sateliteCode() >= 0 {
		alternative := uwp.GenerateOtherWorldUWP(dice.New().SetSeed(bd.nomena), w.UWP(), planetType, starData[bd.position.starCode()], bd.position.planetCode())

		bd.uwp = alternative

		if uwp.New(bd.uwp).Pops().Value() > 2 {
			bd.name = names.RandomPlace(w.Sector() + w.Hex() + bd.nomena)
		}
		if planetType == "MainWorld" {
			bd.uwp = w.UWP()
			bd.name = w.Name()
		}
		if planetType == "LGG" || planetType == "SGG" || planetType == "IG" {
			bd.uwp = uwp.NewGasGigant(dp, planetType)

		}
		bd.calculatePlanetDiameter(dp)
		bd.jumpPointToBody = Astrogation.JumpPointFromObject(bd.diameter)
		bd.calculateOrbitDistanceAU(position, dp)

		if sShadow-bd.orbitDistance > 0 {
			closestJumpPoint := utils.RoundFloat64(sShadow-bd.orbitDistance, 2)
			bd.jumpPointToBody = utils.RoundFloat64(Astrogation.AU2Megameters*(closestJumpPoint), 3)
		}
		hz := astronomical.HabitableOrbit(starData[position.starCode()])
		remarks := uwp.CalculateTradeCodesT5(bd.uwp, w.TradeClassificationsSl(), false, position.planetCode()-hz)
		if position.sateliteCode() > -1 {
			addTC := "Sa"
			if position.sateliteCode() < 14 {
				addTC = "Lk"
			}
			remarks = append(remarks, addTC)
		}
		if planetType != constant.WTpPlanetoid {
			remarks = removeFromSlice(remarks, "As")
		}
		bd.tags = strings.Join(remarks, " ")
	case "Planetoid":
		fmt.Println("Skip belt")
		bd.bodyType = "Belt"
		bd.calculateOrbitDistanceAU(position, dp)
		bd.parentStar = parseStellarData(w)[position.starCode()]
		bd.name = bd.AsteroidDetails()

	case "Star":
		bd.nomena = w.Sector() + " " + w.Hex() + " " + TrvCore.NumToGreek(strCode) + " " + starData[strCode]
		bd.diameter = utils.RoundFloat64(StarDiameter(starData[strCode])*Astrogation.SolDiametrMegameters, 2)
		js := StarDiameter(starData[strCode]) * 0.93
		jsstr := strconv.FormatFloat(js, 'f', 2, 64)
		bd.name = jsstr + " au"

	}

	bd.cleanDataSpecialType()

	return bd
}

func trimSatelliteUWP(uwp, tp string) string {
	data := strings.Split(uwp, "")
	if tp != constant.WTpPlanetoid && data[1] == "0" {
		data[1] = "S"
	}
	//0123456-8
	if data[4] == "0" && data[5] == "0" && data[6] == "0" {
		data[0] = "Y"
	}
	uwp = ""
	for _, d := range data {
		uwp += d
	}
	return uwp
}

func (bd *BodyDetails) cleanDataSpecialType() {
	switch bd.bodyType {
	case "Rings":
		bd.bodyType = "Ring System"
		bd.uwp = ""
		bd.tags = ""
	case "LGG":
		bd.bodyType = "Large Gas Gigant"
		bd.uwp = ""
		bd.tags = ""
	case "SGG":
		bd.bodyType = "Small Gas Gigant"
		bd.uwp = ""
		bd.tags = ""
	case "IG":
		bd.bodyType = "Ice Gigant"
		bd.uwp = ""
		bd.tags = ""
	}
	if bd.uwp == "" && bd.bodyType != "Star" && bd.bodyType != "Belt" {
		bd.name = ""
	}
}

func (bd *BodyDetails) calculatePlanetDiameter(dp *dice.Dicepool) {
	sz := uwp.New(bd.uwp).Size().Value() // * 1000
	switch sz {
	case 21:
		sz = 30
	case 22:
		sz = 40
	case 23:
		sz = 50
	case 24:
		sz = 60
	case 25:
		sz = 70
	case 26:
		sz = 80
	case 27:
		sz = 90
	case 28:
		sz = 125
	case 29:
		sz = 180
	case 30:
		sz = 220
	case 31:
		sz = 250
	case 32:
		sz = 290
	}
	dMl := float64((sz * 1000) + ((dp.FluxNext() * 100) + (dp.FluxNext() * 10) + (dp.FluxNext()))) //диаметр планеты в милях
	if dMl < 0 {
		dMl = dMl * -1
	}
	dKm := dMl * 1.609 //диаметр планеты в километрах
	dMm := dKm / 1000  //диаметр планеты в мегаметрах
	bd.diameter = utils.RoundFloat64(dMm, 3)
}

func (bd *BodyDetails) calculateNomena(planetType string, position numCode, w wrld.World) {
	starData := parseStellarData(w)
	strCode := position.starCode()
	orbCode := position.planetCode()
	satCode := position.sateliteCode()
	switch planetType {
	default:
		bd.nomena = TrvCore.NumToGreek(strCode) + " " + strconv.Itoa(orbCode)
		if satCode != -1 {
			bd.nomena += " " + TrvCore.NumToAnglic(satCode)
		}
		if planetType == "Belt" {
			bd.nomena += " Asteroid Belt (TODO:)" //TODO: привентить механику определения залежей астеройдов)
			planetType = constant.WTpPlanetoid
		}
	case "Star":
		bd.nomena = w.Sector() + " " + w.Hex() + " " + TrvCore.NumToGreek(strCode) + " " + starData[strCode]
	}
}

func (bd *BodyDetails) calculateOrbitDistanceAU(position numCode, dp *dice.Dicepool) {
	orbCode := position.planetCode()
	orbDis := locateOrbitInt(dp, orbCode)
	bd.orbitDistance = orbDis
}

func calculateSateliteOrbitDistance(planetDiam float64, sateliteOrbit int) float64 {
	satDistance := planetDiam * float64(sateliteOrbit)
	satDistance = utils.RoundFloat64(satDistance, 3)
	return satDistance
}

//numCode - составляется из трех чисел (номер звезды, номер орбиты вокруг звезды и номер орбиты вокруг спутника)
//планета будет иметь вид [Х Х -1], звезда = [Х -1 -1]
type numCode struct {
	code [3]int
}

func (nc *numCode) starCode() int {
	return nc.code[0]
}

func (nc *numCode) planetCode() int {
	return nc.code[1]
}

func (nc *numCode) sateliteCode() int {
	return nc.code[2]
}

func allKeys2() (numKeys []numCode) {
	i := 0
	for starNum := 0; starNum < 5; starNum++ {
		numKeys = append(numKeys, numCode{[3]int{starNum, -1, -1}})
		i++
		for orbit := 0; orbit <= 20; orbit++ {
			numKeys = append(numKeys, numCode{[3]int{starNum, orbit, -1}})
			i++
			for satOrbit := 0; satOrbit < 26; satOrbit++ {
				numKeys = append(numKeys, numCode{[3]int{starNum, orbit, satOrbit}})
				i++
			}
		}
	}
	return numKeys
}

// func getHZ(star string) int {
// 	spectral := getStarSpectral(star)
// 	size := getStarSize(star)
// 	if star == "BD" {
// 		spectral = "B"
// 		size = "D"
// 	}
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
// 	if val, ok := hzMap[spectral+size]; !ok {
// 		fmt.Println(val, star, spectral+size)
// 		panic("star class unrecognized")
// 	}
// 	return hzMap[spectral+size]

// }

// func getStarClosestOrbit(star string) int {
// 	spectral := getStarSpectral(star)
// 	size := getStarSize(star)
// 	if star == "BD" {
// 		spectral = "B"
// 		size = "D"
// 	}
// 	hzMap := make(map[string]int)
// 	hzMap["OIa"] = 8
// 	hzMap["OIb"] = 7
// 	hzMap["OII"] = 6
// 	hzMap["OIII"] = 5
// 	hzMap["OIV"] = 4
// 	hzMap["OV"] = 5
// 	hzMap["OVI"] = -1
// 	hzMap["OD"] = 0
// 	hzMap["BIa"] = 7
// 	hzMap["BIb"] = 6
// 	hzMap["BII"] = 5
// 	hzMap["BIII"] = 4
// 	hzMap["BIV"] = 3
// 	hzMap["BV"] = 4
// 	hzMap["BVI"] = -1
// 	hzMap["BD"] = 0
// 	hzMap["AIa"] = 7
// 	hzMap["AIb"] = 5
// 	hzMap["AII"] = 3
// 	hzMap["AIII"] = 1
// 	hzMap["AIV"] = 1
// 	hzMap["AV"] = 0
// 	hzMap["AVI"] = -1
// 	hzMap["AD"] = 0
// 	hzMap["FIa"] = 6
// 	hzMap["FIb"] = 4
// 	hzMap["FII"] = 2
// 	hzMap["FIII"] = 0
// 	hzMap["FIV"] = 0
// 	hzMap["FV"] = 0
// 	hzMap["FVI"] = 0
// 	hzMap["FD"] = 0
// 	hzMap["GIa"] = 7
// 	hzMap["GIb"] = 5
// 	hzMap["GII"] = 2
// 	hzMap["GIII"] = 0
// 	hzMap["GIV"] = 0
// 	hzMap["GV"] = 0
// 	hzMap["GVI"] = 0
// 	hzMap["GD"] = 0
// 	hzMap["KIa"] = 7
// 	hzMap["KIb"] = 6
// 	hzMap["KII"] = 3
// 	hzMap["KIII"] = 0
// 	hzMap["KIV"] = 0
// 	hzMap["KV"] = 0
// 	hzMap["KVI"] = 0
// 	hzMap["KD"] = 0
// 	hzMap["MIa"] = 8
// 	hzMap["MIb"] = 7
// 	hzMap["MII"] = 6
// 	hzMap["MIII"] = 4
// 	hzMap["MIV"] = -1
// 	hzMap["MV"] = 0
// 	hzMap["MVI"] = 0
// 	hzMap["MD"] = 0
// 	if val, ok := hzMap[spectral+size]; !ok {
// 		fmt.Println(val, star, spectral+size)
// 		panic("star class unrecognized")
// 	}
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
// 		stSp = "K"
// 	}
// 	if strings.Contains(starCode, "M") {
// 		stSp = "M"
// 	}
// 	if strings.Contains(starCode, "BD") {
// 		stSp = "BD"
// 	}
// 	return stSp
// }

func (d *SystemDetails) innerType() string {
	return []string{
		constant.WTpInferno,
		constant.WTpInnerWorld,
		constant.WTpBigWorld,
		constant.WTpStormWorld,
		constant.WTpRadWorld,
		constant.WTpHospitable,
	}[d.dicepool.RollNext("1d6").DM(-1).Sum()]
}

func (d *SystemDetails) innerSateliteType() string {
	return []string{
		constant.WTpInferno,
		constant.WTpInnerWorld,
		constant.WTpBigWorld,
		constant.WTpStormWorld,
		constant.WTpRadWorld,
		constant.WTpHospitable,
	}[d.dicepool.RollNext("1d6").DM(-1).Sum()]
}

func (d *SystemDetails) outerType() string {
	return []string{
		constant.WTpWorldlet,
		constant.WTpIceWorld,
		constant.WTpBigWorld,
		constant.WTpIceWorld,
		constant.WTpRadWorld,
		constant.WTpIceWorld,
	}[d.dicepool.RollNext("1d6").DM(-1).Sum()]
}

func (d *SystemDetails) outerSateliteType() string {
	return []string{
		constant.WTpWorldlet,
		constant.WTpIceWorld,
		constant.WTpBigWorld,
		constant.WTpStormWorld,
		constant.WTpRadWorld,
		constant.WTpIceWorld,
	}[d.dicepool.RollNext("1d6").DM(-1).Sum()]
}

func locateOrbitInt(dp *dice.Dicepool, dOrbit int) float64 {
	flux := dp.FluxNext()
	dO := []float64{}
	switch dOrbit {
	default:
		return -999.9
	case 0:
		dO = []float64{0.15, 0.16, 0.17, 0.18, 0.19, 0.20, 0.22, 0.24, 0.26, 0.28, 0.30}
	case 1:
		dO = []float64{0.30, 0.32, 0.34, 0.36, 0.38, 0.40, 0.43, 0.46, 0.49, 0.52, 0.55}
	case 2:
		dO = []float64{0.55, 0.58, 0.61, 0.64, 0.67, 0.70, 0.73, 0.76, 0.79, 0.82, 0.85}
	case 3:
		dO = []float64{0.85, 0.88, 0.91, 0.94, 0.97, 1.00, 1.06, 1.12, 1.18, 1.24, 1.30}
	case 4:
		dO = []float64{1.30, 1.36, 1.42, 1.48, 1.54, 1.60, 1.72, 1.84, 1.96, 2.08, 2.20}
	case 5:
		dO = []float64{2.20, 2.32, 2.44, 2.56, 2.68, 2.80, 3.04, 3.28, 3.52, 3.76, 4.00}
	case 6:
		dO = []float64{4.00, 4.20, 4.40, 4.70, 4.90, 5.20, 5.60, 6.10, 6.60, 7.10, 7.60}
	case 7:
		dO = []float64{7.60, 8.10, 8.50, 9.00, 9.50, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0}
	case 8:
		dO = []float64{15.0, 16.0, 17.0, 18.0, 19.0, 20.0, 22.0, 24.0, 26.0, 28.0, 30.0}
	case 9:
		dO = []float64{30.0, 32.0, 34.0, 36.0, 38.0, 40.0, 43.0, 47.0, 51.0, 54.0, 58.0}
	case 10:
		dO = []float64{58.0, 62.0, 65.0, 69.0, 73.0, 77.0, 84.0, 92.0, 100, 107, 115}
	case 11:
		dO = []float64{115, 123, 130, 138, 146, 154, 169, 184, 200, 215, 231}
	case 12:
		dO = []float64{231, 246, 261, 277, 292, 308, 338, 369, 400, 430, 461}
	case 13:
		dO = []float64{461, 492, 522, 553, 584, 615, 676, 738, 799, 861, 922}
	case 14:
		dO = []float64{922, 984, 1045, 1107, 1168, 1230, 1352, 1475, 1598, 1721, 1844}
	case 15:
		dO = []float64{1844, 1966, 2089, 2212, 2335, 2458, 2703, 2949, 3195, 3441, 3687}
	case 16:
		dO = []float64{3687, 3932, 4178, 4424, 4670, 4916, 5407, 5898, 6390, 6881, 7373}
	case 17:
		dO = []float64{7373, 7864, 8355, 8847, 9338, 9830, 10797, 11764, 12731, 13698, 14665}
	}

	return dO[flux+5]
}

// func cyclePlanetbodyNames() []string {
// 	var names []string
// 	for _, star := range []string{"Alpha", "Beta", "Gamma"} {
// 		names = append(names, star)
// 		for p := 0; p < 20; p++ {
// 			planet := strconv.Itoa(p)
// 			names = append(names, star+" "+planet)
// 			for _, sat := range []string{"a", "b", "c", "d", "e"} {
// 				names = append(names, star+" "+planet+" "+sat)
// 			}
// 		}
// 	}
// 	return names
// }

//Missing Details:
func (d *SystemDetails) rollCloseSatelite() int {
	return d.dicepool.RollNext("2d6").Sum()
}

func (d *SystemDetails) rollFarSatelite() int {
	return d.dicepool.RollNext("2d6").DM(13).Sum()
}

func (d *SystemDetails) rollOrbitPlacement(ggType string) int {
	switch ggType {
	default:
		panic("Unknown Body Type")
	case "LGG":
		return d.dicepool.RollNext("2d6").DM(-5).Sum()
	case "SGG":
		return d.dicepool.RollNext("2d6").DM(-4).Sum()
	case "IG":
		return d.dicepool.RollNext("2d6").DM(-2).Sum()
	case "Belt":
		return d.dicepool.RollNext("2d6").DM(-3).Sum()
	case "World1":
		arr := []int{11, 10, 8, 6, 4, 2, 0, 1, 3, 5, 7, 9}

		suggest := arr[d.dicepool.RollNext("2d6").DM(-2).Sum()]
		//fmt.Println("Suggest Roll:", suggest)
		return suggest
	case "World2":
		arr := []int{18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7}
		return arr[d.dicepool.RollNext("2d6").DM(-2).Sum()]
	}
}

/*


1000 - star
10 - planet
1 - satelite

1050 - пятая планета от первой звезды
*/
func cls() {
	var clear map[string]func()
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func removeFromSlice(sl []string, elem string) []string {
	res := []string{}
	for _, val := range sl {
		if val != elem {
			res = append(res, val)
		}
	}
	return res
}
