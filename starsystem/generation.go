package starsystem

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/profile"

	"github.com/Galdoba/TR_Dynasty/constant"

	"github.com/Galdoba/TR_Dynasty/dice"

	"github.com/Galdoba/TR_Dynasty/wrld"
)

//From - генерирует дители всех планитарных тел в системе на основе данных из SecondSurveyT5
func From(world wrld.World) SystemDetails {
	fmt.Print(world.SecondSurvey(), "\n")
	d := SystemDetails{}
	d.bodyDetail = make(map[string]bodyDetails)
	d.dicepool = dice.New().SetSeed(world.Name() + world.Name())
	starData := parseStellarData(world)
	stBodySlots := setupOrbitalBodySlots(starData)
	stBodySlots = d.placeMW(world, stBodySlots)
	stBodySlots = d.placeGG(world, stBodySlots)
	stBodySlots = d.placeBelts(world, stBodySlots)
	stBodySlots = d.placeOtherWorlds(world, stBodySlots)

	stBodySlots = d.assignWorldTypes(world, stBodySlots)
	stBodySlots = d.placeSatelites(world, stBodySlots)
	fmt.Println("DEBUG: OUTPUT INFO")
	fmt.Println(len(stBodySlots), "valid slots in total")
	for _, val := range allKeys2() {
		if val.starCode() > len(starData)-1 {
			continue
		}
		if stBodySlots[val] == "Star" {
			fmt.Print("\n")
		}
		if stBodySlots[val] != "--VALID--" {
			if val.planetCode() != -1 {
				fmt.Print("	")
			}
			if val.sateliteCode() != -1 {
				fmt.Print("	")
			}
			fmt.Print(val, "	- ", stBodySlots[val], "\n")

		} else {
			delete(stBodySlots, val)
		}

	}
	fmt.Println(len(stBodySlots), "valid slots in total")
	//fmt.Println(starData, stBodySlots)
	return d
}

func (d *SystemDetails) placeMW(world wrld.World, stBodySlots map[numCode]string) map[numCode]string {
	tc := world.TradeClassificationsSl()
	starData := parseStellarData(world)
	hz := getHZ(starData[0])
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
			starHZ = append(starHZ, getHZ(starData[l]))
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
			starHZ = append(starHZ, getHZ(starData[l]))
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
	worldsTotal, err := strconv.Atoi(world.NumOfWorlds())
	if err != nil {
		panic(err)
	}
	otherWorlds := worldsTotal - world.GasGigants() - world.Belts() - 1
	starData := parseStellarData(world)
	starKeys := []int{}
	starHZ := []int{}
	for i := 0; i < otherWorlds; i++ {
		for l := 0; l < len(starData); l++ {
			starKeys = append(starKeys, l)
			starHZ = append(starHZ, getHZ(starData[l]))
		}
	}
	starKeys = starKeys[0:otherWorlds]
	starHZ = starHZ[0:otherWorlds]
	for index, star := range starKeys {
		suggest := starHZ[star] + d.rollOrbitPlacement("World1")
		if index == otherWorlds-1 {
			suggest = starHZ[star] + d.rollOrbitPlacement("World2")
		}

		//valid := false
		for {
			if val, ok := stBodySlots[numCode{[3]int{starKeys[star], suggest, -1}}]; ok {
				if suggest < 0 {
					suggest++
					continue
				}
				if val != "--VALID--" {
					//}
					//if stBodySlots[numCode{[3]int{starKeys[ggIndex], suggest, -1}}] != "--VALID--" {
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
		hz := getHZ(starData[k.starCode()])
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
		hz := getHZ(starData[k.starCode()])
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
		sc := k.starCode()
		pc := k.planetCode()
		if pc < getStarClosestOrbit(starData[sc]) && pc != -1 {
			delete(stBodySlots, k)
		}
		if pc == -1 && k.sateliteCode() == -1 && k.starCode() < len(starData) {
			stBodySlots[k] = "Star"
		}
	}
	return stBodySlots
}

// func (d *SystemDetails) newGG() bodyDetails {

// }

//Test -
func Test() {
	world := wrld.PickWorld()
	fmt.Println(world)
	d := From(world)
	for _, val := range cyclePlanetbodyNames() {
		if bd, ok := d.bodyDetail[val]; ok {
			fmt.Println(bd.ShortInfo())
		}
	}
	//fmt.Println("Enter planetary code for details:")
	key, _ := user.InputStr()
	if bd, ok := d.bodyDetail[key]; ok {
		fmt.Println(bd.FullInfo())
		fmt.Println(d.bodyDetail[key])
	}
	fmt.Println("END PROGRAMM")
}

func parseStellarData(w wrld.World) []string {
	spectralData := strings.Split(w.Stellar(), " ")
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
	bodyDetail map[string]bodyDetails
	dicepool   *dice.Dicepool
	//PositionFromStar__PositionFromPlanet__Name__UWP__Actual Orbit
	/*
		Primary		Fijari	Ka V
		0		Keetle	SGG
		5		Ra-La-Lantra	LGG
		5	0	Ring System	YR00000-0	3.28
		5	8	B'kolior	X621000-0	3.28
	*/
}

type bodyDetails struct {
	nomena           string
	name             string
	uwp              string
	tags             string
	bodyType         string
	diameter         float64
	orbitDistance    float64
	jumpPointToOrbit float64
	orbitSpeed       int
}

func newBody(str string, dp *dice.Dicepool) bodyDetails {
	bd := bodyDetails{}
	data := strings.Split(str, "	")
	bd.nomena = data[0] + " " + data[2] + " " + data[3] + " " + data[4]
	bd.name = data[1]
	if len(data[6]) >= 3 {
		bd.uwp = data[6]
		sz := profile.NewUWP(bd.uwp).Size().Value() * 1000
		if sz == 0 {
			sz = 400
		}
		dKm := sz + (((dp.FluxNext()*100)+dp.FluxNext()*10+dp.FluxNext())*1600)/1000 //диаметр планеты в километрах
		dMm := utils.RoundFloat64(float64(dKm)/1000, 3)                              //диаметр планеты в мегаметрах
		jp := dMm * 100                                                              //точка JP без учета тени звезды
		sDiam, _ := strconv.ParseFloat(data[10], 64)                                 //диаметр звезды
		sShadow := Astrogation.StarJumpShadowAU(sDiam)                               //тень звезды в AU
		dFl := utils.RoundFloat64(float64(dKm), 3)                                   //
		fl, _ := strconv.ParseFloat(data[7], 64)                                     //
		bd.orbitDistance = fl                                                        // орбита в AU
		bd.diameter = dFl                                                            //
		trueJP := jp                                                                 //
		if sShadow > bd.orbitDistance {                                              //если орбита покрывает планету то точка выхода считается от солнца и вычитает орбиту планеты от тени звезды
			addTravell := sShadow - bd.orbitDistance
			trueJP = addTravell * Astrogation.AU2Megameters
		}
		trueJP = utils.RoundFloat64(trueJP, 3)
		bd.jumpPointToOrbit = trueJP
		bd.tags = data[8]
		bd.bodyType = data[5]
		//fmt.Println(data)
	} else {
		bd.bodyType = data[5]
		sDiam, _ := strconv.ParseFloat(data[10], 64)                                          //диаметр звезды
		bd.jumpPointToOrbit = Astrogation.StarJumpShadowAU(sDiam * Astrogation.AU2Megameters) //тень звезды в AU
	}

	return bd
}

func from(world wrld.World) SystemDetails {
	fmt.Print("", world.SecondSurvey(), "\n")
	d := SystemDetails{}
	tabl := []string{}
	d.bodyDetail = make(map[string]bodyDetails)
	d.dicepool = dice.New().SetSeed(world.Name() + world.Name())
	starData := parseStellarData(world)
	starMap := make(map[int][]string)
	for i := 0; i < 17; i++ {
		starMap[1] = append(starMap[1], "EMPTY")
	}
	if len(starData) > 1 {
		for i := 0; i < d.dicepool.RollNext("2d6").DM(-1).Sum(); i++ {
			starMap[2] = append(starMap[2], "EMPTY")
		}
	}
	if len(starData) > 2 {
		for i := 0; i < d.dicepool.RollNext("2d6").DM(-1).Sum(); i++ {
			starMap[3] = append(starMap[3], "EMPTY")
		}
	}
	stObj := []string{"Mainworld"}
	for i := 0; i < getGG(world); i++ {
		stObj = append(stObj, constant.WTpGG)
	}
	for i := 0; i < getBelts(world); i++ {
		stObj = append(stObj, constant.WTpPlanetoid)
	}
	pl, _ := strconv.Atoi(world.NumOfWorlds())
	for i := 0; i < pl-getBelts(world)-getGG(world)-1; i++ {
		stObj = append(stObj, "Planet")
	}
	for _, val := range stObj {
		if val == "Mainworld" {
			starMap[1][getHZ(starData[0])] = val
			continue
		}
		added := false
		runs := -1
		for !added {
			runs++
			r := d.dicepool.RollNext("1d" + strconv.Itoa(len(starMap))).Sum()
			//ln := d.dicepool.RollNext("1d" + strconv.Itoa(len(starMap[r])-1)).Sum()
			orb := getHZ(starData[r-1]) + d.dicepool.FluxNext() + runs
			if orb > len(starMap[r])-1 || orb < 0 {
				runs--
				continue
			}
			if starMap[r][orb] == "EMPTY" {
				starMap[r][orb] = val
				added = true
			}
		}
	}
	for s := 1; s <= len(starMap); s++ {
		//hz := strconv.Itoa(getHZ(starData[s-1]))
		detailLine := "	 	" + starData[s-1] + "	-1	 	 	 	**	" + TrvCore.NumToGreek(s-1)
		tabl = append(tabl, detailLine)
		for j := range starMap[s] {
			nSat := 0
			if j <= getHZ(starData[s-1]) && starMap[s][j] == "Planet" {
				starMap[s][j] = d.innerType()
				nSat = d.dicepool.RollNext("1d6").DM(-5).Sum()
			}
			if j > getHZ(starData[s-1]) && starMap[s][j] == "Planet" {
				starMap[s][j] = d.outerType()
				nSat = d.dicepool.RollNext("1d6").DM(-3).Sum()
			}
			if starMap[s][j] == constant.WTpGG {
				//starMap[s][i] = d.rollGG()
				nSat = d.dicepool.RollNext("1d6").DM(-1).Sum()
			}
			if starMap[s][j] == constant.WTpHospitable || starMap[s][j] == "Mainworld" {
				nSat = d.dicepool.RollNext("1d6").DM(-4).Sum()
			}

			if nSat < 0 {
				nSat = 0
			}

			detailLine := d.makeDetailLine(s, j, starMap[s][j], "", world, getHZ(starData[s-1]))
			if strings.Contains(detailLine, "Mainworld") {
				detailLine = strings.TrimSuffix(detailLine, "Mainworld	")
				detailLine += world.Name() + "	" + world.UWP()

			}

			if detailLine != "" {
				detailLine = world.Hex() + "	" + world.Name() + "	" + detailLine + "	"
				if nSat > 0 {
					detailLine += strconv.Itoa(nSat)
				}
				detailLine += "	"
				tabl = append(tabl, detailLine)
			}

			for sat := 0; sat < nSat; sat++ {
				satType := d.outerSateliteType()
				if j <= getHZ(starData[s-1]) {
					satType = d.innerSateliteType()
				}
				detailLine := d.makeDetailLine(s, j, satType, strconv.Itoa(sat), world, getHZ(starData[s-1]))
				if strings.Contains(detailLine, "Mainworld") {
					detailLine = strings.TrimSuffix(detailLine, "Mainworld	")
					detailLine += world.Name() + "	" + world.UWP() + "	"

				}

				if detailLine != "" {
					detailLine = world.Hex() + "	" + world.Name() + "	" + detailLine + "	" + "	"
					tabl = append(tabl, detailLine)
				}
			}
		}
	}
	plOrbit := ""
	starNum := 0
	for k, v := range tabl {
		//fmt.Println(k, "||", v)
		lnData := strings.Split(v, "	")
		lnData = append(lnData, "")
		lnData = append(lnData, "")

		hz := 0

		switch lnData[2] {
		case "Alpha":
			starNum = 0
		case "Beta":
			starNum = 1
		case "Gamma":
			starNum = 2
		}
		starHZ := getHZ(starData[starNum])
		planetHZ, _ := strconv.Atoi(lnData[3])
		hz = planetHZ - starHZ
		solDiamMm := 13927.7
		starJZ := strconv.FormatFloat(StarDiameter(starData[starNum])*solDiamMm/Astrogation.AU2Megameters, 'f', 2, 64)
		lnData[9] = strconv.Itoa(hz)
		lnData[1] = ""
		// if lnData[0] == "Star" {
		// 	lnData[1] = "Star"
		// 	fmt.Println("Star***********")

		// }
		if lnData[5] == world.Name() {
			lnData[1] = world.Name() + " "
			for _, val := range profile.CalculateTradeCodesT5(lnData[6], nil, true, 0) {
				lnData[8] += val + " "
			}
			lnData[5] = "Mainworld"
		} else {
			if lnData[6] != " " {

				uwp := profile.NewUWP(lnData[6])

				for _, val := range profile.CalculateTradeCodesT5(lnData[6], nil, false, hz) {
					lnData[8] += val + " "
				}
				if uwp.Pops().Value() > 0 {
					lnData[1] += "(!)"
				}
				if uwp.Starport().String() == "A" {
					if uwp.Pops().Value() >= 7 {
						lnData[1] += "H"
					}
					lnData[1] += "D"
				}
				if uwp.Starport().String() == "B" {
					if uwp.Pops().Value() >= 8 {
						lnData[1] += "H"
					}
					lnData[1] += "D"
				}
				if uwp.Starport().String() == "C" {
					if uwp.Pops().Value() >= 9 {
						lnData[1] += "H"
					}
					lnData[1] += "D"
				}
				if uwp.Starport().String() == "D" {
					lnData[1] += "D"
				}
				if uwp.Starport().String() == "H" || uwp.Starport().String() == "E" {
					lnData[1] += "B"
				}

				if uwp.TL().Value() > 9 {
					lnData[1] += "*"
				}
				if uwp.TL().Value() > 0 {
					lnData[1] += "*"
				}
				if uwp.TL().Value() > 12 {
					lnData[1] += "*"
				}

			}
		}

		if lnData[5] == constant.WTpPlanetoid {
			lnData[5] = "Asteroid Belt"
		}

		if lnData[4] == "" {
			plOrbit = locateOrbit(d, lnData[3])
			lnData[7] = plOrbit
		}
		lnData[7] = plOrbit
		if lnData[6] == " " {
			lnData[9] = lnData[2]
			lnData[2] = lnData[8]
			lnData[0] = world.Hex()
			//lnData[1] = ""
			lnData[3] = ""
			lnData[4] = ""
			lnData[5] = lnData[9]
			lnData[7] = ""
			lnData[8] = ""
			lnData[9] = ""
			lnData[10] = ""

		}
		lnData[10] = starJZ
		tabl[k] = concatSlice(lnData)
	}

	for _, val := range tabl {
		key := drawKey(val)
		//fmt.Println(key, "||", val)
		d.bodyDetail[key] = newBody(val, d.dicepool)
	}

	//fmt.Println(d.bodyDetail)

	//fmt.Println(d.bodyDetail[k])
	return d
}

func drawKey(val string) string {
	data := strings.Split(val, "	")
	key := data[2]
	if data[3] != "" {
		key += " " + data[3]
	}
	if data[4] != "" {
		key += " " + data[4]
	}
	return key
}

func concatSlice(sl []string) string {
	str := ""
	for i := range sl {
		str += sl[i] + "	"
	}
	str = strings.TrimSuffix(str, "	")
	return str
}

func (d *SystemDetails) makeDetailLine(s, i int, pType string, st string, mainworld wrld.World, hz int) string {
	if pType == "EMPTY" {
		return ""
	}
	line := ""
	switch s {
	case 1:
		line += "Alpha"
	case 2:
		line += "Beta"
	case 3:
		line += "Gamma"
	}
	if st != "" {
		n, _ := strconv.Atoi(st)
		st = string([]byte(strings.ToLower(TrvCore.NumToAnglic(n)))[0])
	}
	line += "	" + strconv.Itoa(i) + "	"
	line += st + "	" + pType + "	"

	if pType != "Mainworld" {
		//mwUWP := profile.NewUWP(mainworld.UWP())

		//mwTags := profile.CalculateTradeCodesT5(mainworld.UWP(), []string{}, true, hz)
		line += profile.RandomUWP(d.dicepool, pType, mainworld.UWP())

	}
	return line
}

func allKeys() []string {
	keys := []string{}
	for i := 0; i <= 20; i++ {
		for j := -1; j < 10; j++ {
			k := strconv.Itoa(i) + "	"
			if j > -1 {
				k += strconv.Itoa(j)
			}
			keys = append(keys, k)
		}
	}
	return keys
}

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
	//keys := []string{}
	//numKeys := []string{}
	//ehexKeys := []string{}
	i := 0
	for starNum := 0; starNum < 5; starNum++ {
		numKeys = append(numKeys, numCode{[3]int{starNum, -1, -1}})
		//fmt.Print(i, " N := '", numKeys[i], "'\n")
		i++
		for orbit := 0; orbit <= 20; orbit++ {
			numKeys = append(numKeys, numCode{[3]int{starNum, orbit, -1}})
			//	fmt.Print(i, " N := '", numKeys[i], "'\n")
			i++
			for satOrbit := 0; satOrbit < 26; satOrbit++ {
				numKeys = append(numKeys, numCode{[3]int{starNum, orbit, satOrbit}})
				//		fmt.Print(i, " N := '", numKeys[i], "'\n")
				i++
			}
		}
	}
	return numKeys
}

func (d *SystemDetails) rollGG() string {
	switch d.dicepool.FluxNext() {
	case -5, -4:
		ggType := "Small Gas Gigant"
		if d.dicepool.RollNext("1d2").Sum() == 2 {
			ggType = "Ice Gigant"
		}
		return ggType
	default:
		return "Large Gas Gigant"
	}
}

func (d *SystemDetails) rollSatelliteName() string {
	r := d.dicepool.FluxNext()
	switch d.dicepool.RollNext("2d6").Sum() {
	case 8, 9, 10, 11, 12:
		r += 7
	default:
		r += 20
	}
	return TrvCore.NumToAnglic(r)
}

func (d *SystemDetails) rollBeltposition() int {
	r := d.dicepool.RollNext("2d6").DM(-2).Sum()
	beltPos := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return beltPos[r]
}

func (d *SystemDetails) rollPlanetposition() int {
	r := d.dicepool.RollNext("2d6").DM(-2).Sum()
	beltPos := []int{10, 8, 6, 4, 2, 0, 1, 3, 5, 7, 9}
	return beltPos[r]
}

func getGG(world wrld.World) int {
	pbg := world.PBG()
	gg, err := strconv.Atoi(string([]byte(pbg)[2]))
	if err != nil {
		panic(err)
	}
	return gg
}

func getBelts(world wrld.World) int {
	pbg := world.PBG()
	gg, err := strconv.Atoi(string([]byte(pbg)[1]))
	if err != nil {
		panic(err)
	}
	return gg
}

func starPositions() []string {
	return []string{
		"Primary Star",
		"Close Star",
		"Near Star",
		"Far Star",
	}
}

func getHZ(star string) int {
	spectral := getStarSpectral(star)
	size := getStarSize(star)
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

func getStarClosestOrbit(star string) int {
	spectral := getStarSpectral(star)
	size := getStarSize(star)
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
		stSp = "K"
	}
	if strings.Contains(starCode, "M") {
		stSp = "M"
	}
	if strings.Contains(starCode, "BD") {
		stSp = "BD"
	}
	return stSp
}

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

func locateOrbit(d SystemDetails, bOrbit string) string {
	dOrbit, err := strconv.Atoi(bOrbit)
	if err != nil {
		panic(err)
	}
	flux := d.dicepool.FluxNext()
	dO := []float64{}
	switch dOrbit {
	default:
		return "Star"
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

	return strconv.FormatFloat(dO[flux+5], 'f', 2, 64)
}

func cyclePlanetbodyNames() []string {
	var names []string
	for _, star := range []string{"Alpha", "Beta", "Gamma"} {
		names = append(names, star)
		for p := 0; p < 20; p++ {
			planet := strconv.Itoa(p)
			names = append(names, star+" "+planet)
			for _, sat := range []string{"a", "b", "c", "d", "e"} {
				names = append(names, star+" "+planet+" "+sat)
			}
		}
	}
	return names
}

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
		return arr[d.dicepool.RollNext("2d6").DM(-2).Sum()]
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
