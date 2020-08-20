package world

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

//Cogri     0101   CA6A643-9 N  Ri Wa      A

const (
	//constant.PrSize     = "Size"
	typeHEX = "Hex"
	typeUWP = "UWP"

	planetStatPopDigit = "PopDigit"
	planetStatBelt     = "Belts"
	planetStatGasG     = "GG"
	eXImportance       = "Imp"
	statDensity        = "Planet Density"
	eXResources        = "Resources"
	eXLabor            = "Labor"
	eXInfrastructure   = "Infrastructure"
	eXCulture          = "Culture"

	tradeClassificationAsteroidBelt     = "As"
	tradeClassificationDesert           = "De"
	tradeClassificationFluid            = "Fl"
	tradeClassificationGardenWorld      = "Ga"
	tradeClassificationHellworld        = "He"
	tradeClassificationIceCapped        = "Ic"
	tradeClassificationOceanWorld       = "Oc"
	tradeClassificationVacuum           = "Va"
	tradeClassificationWaterWorld       = "Wa"
	tradeClassificationDieback          = "Di"
	tradeClassificationBarren           = "Ba"
	tradeClassificationLowPopulation    = "Lo"
	tradeClassificationNonIndustrial    = "Ni"
	tradeClassificationPreHigh          = "Ph"
	tradeClassificationHighPopulation   = "Hi"
	tradeClassificationLowTech          = "Lt"
	tradeClassificationHighTech         = "Ht"
	tradeClassificationPreAgricultural  = "Pa"
	tradeClassificationAgricultural     = "Ag"
	tradeClassificationNonAgricultural  = "Na"
	tradeClassificationPrisonExileCamp  = "Px"
	tradeClassificationPreIndustrial    = "Pi"
	tradeClassificationIndustrial       = "In"
	tradeClassificationPoor             = "Po"
	tradeClassificationPreRich          = "Pr"
	tradeClassificationRich             = "Ri"
	tradeClassificationFrozen           = "Fr"
	tradeClassificationHot              = "Ho"
	tradeClassificationCold             = "Co"
	tradeClassificationLocked           = "Lk"
	tradeClassificationTropic           = "Tr"
	tradeClassificationTundra           = "Tu"
	tradeClassificationTwilightZone     = "Tz"
	tradeClassificationFarming          = "Fa"
	tradeClassificationMining           = "Mi"
	tradeClassificationMilitaryRule     = "Mr"
	tradeClassificationPenalColony      = "Pe"
	tradeClassificationReserve          = "Re"
	tradeClassificationSubsectorCapital = "Cp"
	tradeClassificationSectorCapital    = "Cs"
	tradeClassificationCapital          = "Cx"
	tradeClassificationColony           = "Cy"
	tradeClassificationSatellite        = "Sa"
	tradeClassificationForbidden        = "Fo"
	tradeClassificationPuzzle           = "Pz"
	tradeClassificationDangerous        = "Da"
	tradeClassificationDataRepository   = "Ab"
	tradeClassificationAncientSite      = "An"

	portBaseNaval    = "N"
	portBaseScout    = "S"
	portBaseResearch = "R"
	portBaseTAS      = "T"

	tempFrozen    = "Frozen"
	tempCold      = "Cold"
	tempTemperate = "Termperate"
	tempHot       = "Hot"
	tempBoiling   = "Boiling"

	DataTypeStarport    = 0
	DataTypeSize        = 1
	DataTypeAtmosphere  = 2
	DataTypeHydrosphere = 3
	DataTypePopulation  = 4
	DataTypeGoverment   = 5
	DataTypeLaws        = 6
	DataTypeTechLevel   = 8
)

//World - объект отвечающий за описание планеты и ее непосредственного окружения
type World struct {
	stat map[string]int    //Избавиться
	data map[string]string //
	//temperature  string            // увести в карту
	name string //
	//uwp          string            // увести в карту или вообще избавиться
	tradeCodes   []string          //
	importanceEx string            // увести в карту или сделать интом
	economyEx    string            // увести в карту
	cultureEx    string            // увести в карту
	nobility     string            // увести в карту
	bases        []string          //
	travelCode   string            // увести в карту
	pbg          string            // увести в карту
	worlds       string            // это вообще что?
	stellar      map[string]string // технически должно быть в другом структе
	dice         *dice.Dicepool
	//esscStSyst   *esscStarSystem
}

//NewWorld -
func NewWorld(name string) World {
	world := World{}
	world.name = name
	world.stat = make(map[string]int)
	world.data = make(map[string]string)
	seed := utils.SeedFromString(name)
	world.dice = dice.New(seed)

	// if !uwpValid(uwp) {
	// 	uwp = "RANDOM"
	// }
	// world.SetUWP(uwp)

	//world.SecondSurvey()
	return world
}

//SetHex -
func (w *World) SetHex(hex string) {
	w.data[typeHEX] = hex
}

//Hex -
func (w *World) Hex() string {
	return w.data[typeHEX]
}

//UWP -
func (w *World) UWP() string {
	return w.data[constant.PrStarport] + w.data[constant.PrSize] + w.data[constant.PrAtmo] + w.data[constant.PrHydr] +
		w.data[constant.PrPops] + w.data[constant.PrGovr] + w.data[constant.PrLaws] + "-" + w.data[constant.PrTL]
}

//SetUWP -
func (w *World) SetUWP(uwp string) *World {
	uwp = strings.ToUpper(uwp)
	if uwp == "RANDOM" {
		rand.Seed(utils.SeedFromString(w.name))
		//rWorld := BuildPlanet() //вот тут должен быть рандомный сборщик
		w.rollSize()
		w.rollAtmo()
		w.rollTemperature()
		w.rollHydr()
		w.rollPops()
		w.rollGovr()
		w.rollLaws()
		w.rollSpaceport()
		w.bases = rollBases(w.data[constant.PrStarport])
		w.rollTech()
		w.buildUWP()
	}
	if !uwpValid(uwp) {
		return w
	}
	code := strings.Split(uwp, "")
	w.data[constant.PrStarport] = code[0]
	w.data[constant.PrSize] = code[1]
	w.data[constant.PrAtmo] = code[2]
	w.data[constant.PrHydr] = code[3]
	w.data[constant.PrPops] = code[4]
	w.data[constant.PrGovr] = code[5]
	w.data[constant.PrLaws] = code[6]
	w.data[constant.PrTL] = code[8]
	w.data["UWP"] = uwp
	return w
}

//TotalWorlds -
func (w *World) TotalWorlds() string {
	if w.worlds != "" {
		return w.worlds
	}
	w.PBG()
	wor := 1 + w.Stat(planetStatGasG) + w.Stat(planetStatBelt) + utils.RollDice("2d6")
	w.worlds = convert.ItoS(wor)
	//Great Rift:
	//w.worlds = esscPlanetaryBodies()
	return w.worlds
}

//JumpShadowRadius -
func (w *World) JumpShadowRadius() int {
	base := w.Size() * 160
	if base == 0 {
		base = 80
	}
	return base
}

// func (w *World) Size() int {
// 	if val, ok := w.Stat("Size"]; ok {
// 		return val
// 	}
// 	w.UWP()
// 	return w.Stat("Size"]
// }

//Name -
func (w *World) Name() string {
	return w.name
}

//Law -
func (w *World) Law() int {
	return w.Stat(constant.PrLaws)
}

//TechLevel -
func (w *World) TechLevel() int {
	return w.Stat(constant.PrTL)
}

func uwpValid(uwp string) bool {
	if uwp == "RANDOM" {
		return false
	}
	if len(uwp) != 9 {
		return false
	}
	data := strings.Split(uwp, "")
	for i := range data {
		if data[i] == "-" || data[i] == "_" {
			continue
		}
		if TrvCore.EhexToDigit(data[i]) == -999 {
			return false
		}
	}
	return true
}

func UWPisValid(uwp string) bool {
	return uwpValid(uwp)
}

//DebugInfo -
func (w *World) DebugInfo() {
	fmt.Println("stat        =", w.stat)
	fmt.Println("data        =", w.data)
	fmt.Println("temperature =", w.data["Temperature"])
	//fmt.Println("port        =", w.port)
	fmt.Println("hex         =", w.data[typeHEX])
	fmt.Println("name        =", w.name)
	fmt.Println("uwp         =", w.UWP())
	fmt.Println("tradeCodes  =", w.tradeCodes)
	fmt.Println("importanceEx=", w.importanceEx)
	fmt.Println("economyEx   =", w.economyEx)
	fmt.Println("cultureEx   =", w.cultureEx)
	fmt.Println("nobility    =", w.nobility)
	fmt.Println("bases       =", w.bases)
	fmt.Println("travelCode  =", w.travelCode)
	fmt.Println("pbg         =", w.pbg)
	fmt.Println("worlds      =", w.worlds)
	fmt.Println("stellar     =", w.stellar)
}

func (w *World) Factions() string {
	if val, ok := w.data["Factions"]; ok {
		return val
	}
	fact := "f"
	fDM := 0
	switch w.Stat(constant.PrGovr) {
	case 0, 7:
		fDM++
	case 10, 11, 12, 13, 14, 15, 16:
		fDM--
	}
	fNum := utils.RollDice("d3", fDM)
	for i := 0; i < fNum; i++ {
		switch utils.RollDice("2d6") {
		case 2, 3:
			fact += "1"
		case 4, 5:
			fact += "2"
		case 6, 7:
			fact += "3"
		case 8, 9:
			fact += "4"
		case 10, 11:
			fact += "5"
		case 12:
			fact += "6"
		}
	}
	w.data["Factions"] = fact
	return w.data["Factions"]
}

func (w *World) FirstSurvey() string {
	survey := ""
	rand.Seed(utils.SeedFromString(w.name))
	w.UWP()
	w.PBG()
	w.HomeStar()
	w.WorldOrbitHZ()
	w.PlanetType()
	w.UpdateTradeClassifications()
	return survey
}

//SecondSurvey -
func (w *World) SecondSurvey() string {
	rand.Seed(utils.SeedFromString(w.name))
	w.UWP()
	w.PBG()
	w.HomeStar()
	w.WorldOrbitHZ()
	w.PlanetType()
	w.ImportanceEx()
	w.EconomicEx()
	w.CulturalEx()
	w.UpdateTradeClassifications()
	w.Nobility()
	w.Bases()
	w.UpdateTravelZone()
	w.NIL()
	w.SystemStars()
	w.Factions()

	// w.CalculateStarOrbits()
	// w.TotalWorlds()
	// w.PlaceWorlds()

	var survey string
	survey = w.Hex()
	survey = survey + "	" + w.Name()
	survey = survey + "	" + w.UWP()
	survey = survey + "	"
	tcodes := w.TradeCodes()
	for i := range tcodes {
		if i != 0 {
			survey = survey + " "
		}
		survey = survey + tcodes[i]
	}
	survey = survey + "	" + w.importanceEx
	survey = survey + "	" + w.economyEx
	survey = survey + "	" + w.cultureEx
	survey = survey + "	" + w.nobility
	survey = survey + "	"
	for i := range w.bases {
		if i == 0 {
			continue
		}
		if i != 1 {
			survey = survey + " "
		}
		survey = survey + w.bases[i]
	}
	survey = survey + "	" + w.travelCode
	survey = survey + "	" + w.pbg
	survey = survey + "	" + w.worlds
	survey = survey + "	"
	for _, st := range []string{"P", "Pc", "C", "Cc", "N", "Nc", "F", "Fc"} {
		if val, ok := w.stellar[st]; ok {
			survey = survey + " " + val
		}
		//survey = survey + w.stellar[i]
	}

	return survey
}

func (w *World) Star() string {
	if val, ok := w.stellar["P"]; ok {
		return val
	}
	w.SystemStars()
	return w.stellar["P"]
}

//PopulationDigit -
func (w *World) PopulationDigit() int {
	pbg := w.PBG()
	data := strings.Split(pbg, "")
	p, err := strconv.Atoi(data[0])
	if err != nil {
		panic("Population Digit unreadable! " + w.Name() + " " + w.UWP())
	}
	return p
}

//BeltsDigit -
func (w *World) BeltsDigit() int {
	pbg := w.PBG()
	data := strings.Split(pbg, "")
	p, err := strconv.Atoi(data[1])
	if err != nil {
		panic("Belts Digit unreadable! " + w.Name() + " " + w.UWP())
	}
	return p
}

//GasGigantsDigit -
func (w *World) GasGigantsDigit() int {
	pbg := w.PBG()
	data := strings.Split(pbg, "")
	p, err := strconv.Atoi(data[2])
	if err != nil {
		panic("Gas Gigants Digit unreadable! " + w.Name() + " " + w.UWP())
	}
	return p
}

//ULP - broken
func (w *World) ULP() string {
	if val, ok := w.data["ULP"]; ok {
		return val
	}
	codex, err := NewCodex(w)
	if err != nil {
		return err.Error()
	}
	return codex.lawProfile
}

func createGGCode() string {
	fl := TrvCore.Flux()
	dm := 0
	if fl > 2 {
		dm = 1
	}
	if fl < 2 {
		dm = -1
	}
	r := utils.RollDice("2d6", dm-1)
	size := []string{"SGGl", "SGGm", "SGGn", "LGGp", "LGGq", "LGGr", "LGGs", "LGGt", "LGGu", "LGGv", "LGGw", "LGGx", "LGGy"}
	return size[r]
}

//PlaceWorlds - broken
func (w *World) PlaceWorlds() {
	mWorldOrbit := calculateHZ(w.stellar["P"]) + convert.StoI(w.data["HZVariance"])
	ggSl := strings.Split(w.PBG(), "")
	freeGG := convert.StoI(ggSl[2])
	w.stellar[convert.ItoS(mWorldOrbit)] = "MainWorld"
	if w.data["PlanetType"] != "Planet" {
		if freeGG > 0 {
			w.stellar[convert.ItoS(mWorldOrbit)] = createGGCode()
		}
	}

}

//NIL -
func (w *World) NIL() string {
	if status, ok := w.data["Nil"]; ok {
		return status
	}
	status := "Absent"
	if w.Stat(constant.PrPops) == 0 && matchValue(w.Stat(constant.PrAtmo), 2, 3, 4, 5, 6, 7, 8, 9, 13, 15, 16) && w.Stat(constant.PrTL) == 0 {
		status = "Extinct Natives: Intelligent Life evolved here, but now extinct."
	}
	if w.Stat(constant.PrPops) == 0 && isInRange(w.Stat(constant.PrAtmo), 10, 12) && w.Stat(constant.PrTL) == 0 {
		status = "Extinct Exotic Natives: Intelligent Life evolved here, but now extinct."
	}
	if w.Stat(constant.PrPops) == 0 && matchValue(w.Stat(constant.PrAtmo), 2, 3, 4, 5, 6, 7, 8, 9, 13, 15, 16) && w.Stat(constant.PrTL) != 0 {
		status = "Catastrophic XN: Evidence of Extinct Natives remains."
	}
	if w.Stat(constant.PrPops) == 0 && isInRange(w.Stat(constant.PrAtmo), 10, 12) && w.Stat(constant.PrTL) != 0 {
		status = "Catastrophic EXN: Evidence of Exotic Extinct Natives remains."
	}
	if matchValue(w.Stat(constant.PrPops), 1, 2, 3) && w.Stat(constant.PrTL) != 0 {
		status = "Transients: Temporary commercial or scientific activity."
	}
	if matchValue(w.Stat(constant.PrPops), 4, 5, 6) && w.Stat(constant.PrTL) != 0 {
		status = "Settlers: The initial steps of creating a colony."
	}
	if w.Stat(constant.PrPops) > 6 && matchValue(w.Stat(constant.PrAtmo), 0, 1) && w.Stat(constant.PrTL) != 0 {
		status = "Transplants: Current locals evolved elsewhere."
	}
	if w.Stat(constant.PrPops) == 0 && matchValue(w.Stat(constant.PrAtmo), 0, 1) && w.Stat(constant.PrTL) != 0 {
		status = "Vanished Transplants: Evidence of Transplants, no longer present."
	}
	if w.Stat(constant.PrPops) > 6 && matchValue(w.Stat(constant.PrAtmo), 10, 11, 12) && w.Stat(constant.PrTL) != 0 {
		status = "Exotic Natives: Environment incompatible with humans"
	}
	if w.Stat(constant.PrPops) > 6 && matchValue(w.Stat(constant.PrAtmo), 2, 3, 4, 5, 6, 7, 8, 9, 13, 15, 16) && w.Stat(constant.PrTL) != 0 {
		status = "Natives: Intelligent Life evolved on this world."
	}
	if w.Stat(constant.PrGovr) == 1 {
		status = "Corporate: Locals are employees from elsewhere."
	}
	if w.Stat(constant.PrGovr) == 6 {
		status = "Colonists: Locals are colonists from another world."
	}

	w.data["NIL"] = status
	return status
}

//SystemStars -
func (w *World) SystemStars() string {
	if ststrs, ok := w.data["System Stars"]; ok {
		return ststrs
	}
	ststrs := "P"
	w.stellar = make(map[string]string)
	//w.stellar["P"] = rollSpectTypeAndSize(0)
	w.stellar["P"] = w.HomeStar()
	if TrvCore.Flux() > 2 {
		ststrs = ststrs + "c"
		w.stellar["Pc"] = rollSpectTypeAndSize(utils.RollDice("d6", -1))
	}
	if TrvCore.Flux() > 2 {
		ststrs = ststrs + " C"
		w.stellar["C"] = rollSpectTypeAndSize(utils.RollDice("d6", 2))
		if TrvCore.Flux() > 2 {
			ststrs = ststrs + "c"
			w.stellar["Cc"] = rollSpectTypeAndSize(utils.RollDice("d6", -1))
		}
	}
	if TrvCore.Flux() > 2 {
		ststrs = ststrs + " N"
		w.stellar["N"] = rollSpectTypeAndSize(utils.RollDice("d6", 2))

		if TrvCore.Flux() > 2 {
			ststrs = ststrs + "c"
			w.stellar["Nc"] = rollSpectTypeAndSize(utils.RollDice("d6", -1))
		}
	}
	if TrvCore.Flux() > 2 {
		ststrs = ststrs + " F"
		w.stellar["F"] = rollSpectTypeAndSize(utils.RollDice("d6", 2))
		if TrvCore.Flux() > 2 {
			ststrs = ststrs + "c"
			w.stellar["Fc"] = rollSpectTypeAndSize(utils.RollDice("d6", -1))
		}
	}

	w.data["System Stars"] = ststrs
	return w.data["System Stars"]

}

func (w *World) CalculateStarOrbits() {
	//systemModel := w.data["System Stars"]
	if val, ok := w.stellar["P"]; ok {
		//fmt.Println("System Center:", val)
		w.stellar["P"] = val
	}
	orb := 0
	if val, ok := w.stellar["Pc"]; ok {
		w.stellar[strconv.Itoa(orb)] = val
	}
	orb = utils.RollDice("d6", -1)
	if val, ok := w.stellar["C"]; ok {
		w.stellar[strconv.Itoa(orb)] = val
	}
	if val, ok := w.stellar["Cc"]; ok {
		w.stellar[strconv.Itoa(orb)+"_"] = val
	}
	orb = utils.RollDice("d6", 5)
	if val, ok := w.stellar["N"]; ok {
		w.stellar[strconv.Itoa(orb)] = val
	}
	if val, ok := w.stellar["Nc"]; ok {
		w.stellar[strconv.Itoa(orb)+"_"] = val
	}
	orb = utils.RollDice("d6", 11)
	if val, ok := w.stellar["F"]; ok {
		w.stellar[strconv.Itoa(orb)] = val
	}
	if val, ok := w.stellar["Fc"]; ok {
		w.stellar[strconv.Itoa(orb)+"_"] = val
	}
	//fmt.Println(w.stellar)
}

func (w *World) ImportanceEx() string {
	if w.importanceEx != "" {
		return w.importanceEx
	}
	imp := 0
	if w.data[constant.PrStarport] == "A" || w.data[constant.PrStarport] == "B" {
		imp++
	}
	if w.data[constant.PrStarport] == "D" || w.data[constant.PrStarport] == "E" || w.data[constant.PrStarport] == "X" {
		imp--
	}
	if w.Stat(constant.PrTL) >= eHexD("G") {
		imp++
	}
	if w.Stat(constant.PrTL) >= eHexD("A") {
		imp++
	}
	if w.Stat(constant.PrTL) <= eHexD("8") {
		imp--
	}
	if w.matchTradeClassification(tradeClassificationAgricultural) {
		imp++
	}
	if w.matchTradeClassification(tradeClassificationHighPopulation) {
		imp++
	}
	if w.matchTradeClassification(tradeClassificationIndustrial) {
		imp++
	}
	if w.matchTradeClassification(tradeClassificationRich) {
		imp++
	}
	if w.Stat(constant.PrPops) <= 6 {
		imp--
	}
	if BasePresent("N", w) && BasePresent("S", w) {
		imp++
	}
	if w.data["Way Station"] == "TRUE" {
		imp++
	}
	tag := ""
	if imp == 0 {
		tag = " "
	}
	if imp > 0 {
		tag = "+"
	}
	//w.data[eXImportance] = dEHex(imp)
	w.data[eXImportance] = "{" + tag + convert.ItoS(imp) + "}"
	w.importanceEx = "{" + tag + convert.ItoS(imp) + "}"
	return w.importanceEx
}

func (w *World) EconomicEx() string {
	if w.economyEx != "" {
		return w.economyEx
	}
	w.PBG() // удоставеряемяся что данные в планете есть
	//resourses := utils.RollDice("2d6")
	resourses := w.dice.RollNext("2d6").Sum()
	if w.Stat(constant.PrTL) >= 8 {
		resourses = resourses + w.Stat(planetStatBelt) + w.Stat(planetStatGasG)
	}
	resourses = utils.BoundInt(resourses, 0, 999)
	labor := w.Stat(constant.PrPops) - 1
	labor = utils.BoundInt(labor, 0, 999)
	infrastructure := 0
	if _, ok := w.data[eXImportance]; !ok { // удоставеряемяся что данные в планете есть Importance
		w.ImportanceEx()
	}
	pops := w.Stat(constant.PrPops)
	if pops == 0 {
		infrastructure = 0
	}
	if utils.InRange(pops, 1, 3) {
		infrastructure = w.Stat(eXImportance)
	}
	if utils.InRange(pops, 4, 6) {
		infrastructure = w.Stat(eXImportance) + utils.RollDice("d6")
	}
	if pops > 6 {
		infrastructure = w.Stat(eXImportance) + utils.RollDice("2d6")
	}
	infrastructure = utils.BoundInt(infrastructure, 0, 999)
	efficiency := TrvCore.Flux()
	sigil := "-"
	if efficiency >= 0 {
		sigil = "+"
	}
	effCode := ""
	if efficiency < 0 {
		effCode = convert.ItoS(efficiency * -1)
	} else {
		effCode = convert.ItoS(efficiency)
	}
	w.economyEx = "(" + dEHex(resourses) + dEHex(labor) + dEHex(infrastructure) + sigil + effCode + ")"
	eff := efficiency
	if eff == 0 {
		eff = 1
	}
	ru := utils.BoundInt(resourses, 1, 999) * utils.BoundInt(labor, 1, 999) * utils.BoundInt(infrastructure, 1, 999) * eff
	w.data["RU"] = strconv.Itoa(ru)
	return w.economyEx
}

func (w *World) CulturalEx() string {
	if w.cultureEx != "" {
		return w.cultureEx
	}
	pops := w.Stat(constant.PrPops)
	if _, ok := w.data[eXImportance]; !ok {
		w.ImportanceEx()
	}
	heterogenity := pops + TrvCore.Flux()
	acceptence := pops + w.Stat(eXImportance)
	strangeness := TrvCore.Flux() + 5
	symbols := TrvCore.Flux() + w.Stat(constant.PrTL)
	data := append([]int{}, heterogenity, acceptence, strangeness, symbols)
	for i := range data {
		if data[i] < 1 {
			data[i] = 1
		}
		if pops == 0 {
			data[i] = 0
		}
	}
	w.cultureEx = "[" + dEHex(data[0]) + dEHex(data[1]) + dEHex(data[2]) + dEHex(data[3]) + "]"
	return w.cultureEx
}

func (w *World) Nobility() string {
	if w.nobility != "" {
		return w.nobility
	}
	if _, ok := w.data[eXImportance]; !ok {
		w.ImportanceEx()
	}
	nob := "B"
	if w.matchTradeClassification(tradeClassificationPreAgricultural) || w.matchTradeClassification(tradeClassificationPreRich) {
		nob += "c"
	}
	if w.matchTradeClassification(tradeClassificationAgricultural) || w.matchTradeClassification(tradeClassificationRich) {
		nob += "C"
	}
	if w.matchTradeClassification(tradeClassificationPreIndustrial) {
		nob += "D"
	}
	if w.matchTradeClassification(tradeClassificationPreHigh) {
		nob += "e"
	}
	if w.matchTradeClassification(tradeClassificationIndustrial) || w.matchTradeClassification(tradeClassificationHighPopulation) {
		nob += "E"
	}
	if w.Stat(eXImportance) > 3 && (!w.matchTradeClassification(tradeClassificationSubsectorCapital) && !w.matchTradeClassification(tradeClassificationSectorCapital)) {
		nob += "f"
	}
	if w.matchTradeClassification(tradeClassificationSubsectorCapital) || w.matchTradeClassification(tradeClassificationSectorCapital) {
		nob += "F"
	}
	w.nobility = nob
	return w.nobility
}

func (w *World) matchTradeClassification(tc string) bool {
	tcSlc := w.TradeCodes()
	if len(tcSlc) < 1 {
		w.UpdateTradeClassifications()
		tcSlc = w.TradeCodes()
	}
	for i := range tcSlc {
		if tc == tcSlc[i] {
			return true
		}
	}
	return false
}

func (w *World) SetBases(bases ...string) {
	w.bases = nil
	w.bases = append(w.bases, "")
	for i := range bases {
		w.bases = append(w.bases, bases[i])
	}
}

func (w *World) Bases() []string {
	if len(w.bases) != 0 {
		return w.bases
	}
	if _, ok := w.data[constant.PrStarport]; !ok {
		w.StarPort()
	}
	w.bases = append(w.bases, "")
	switch w.data[constant.PrStarport] {
	case "A":
		r1 := utils.RollDice("2d6")
		r2 := utils.RollDice("2d6")
		if r1 <= 6 {
			w.bases = append(w.bases, "N")
		}
		if r2 <= 4 {
			w.bases = append(w.bases, "S")
		}
	case "B":
		r1 := utils.RollDice("2d6")
		r2 := utils.RollDice("2d6")
		if r1 <= 5 {
			w.bases = append(w.bases, "N")
		}
		if r2 <= 5 {
			w.bases = append(w.bases, "S")
		}
	case "C":
		r2 := utils.RollDice("2d6")
		if r2 <= 6 {
			w.bases = append(w.bases, "S")
		}
	case "D":
		r2 := utils.RollDice("2d6")
		if r2 <= 7 {
			w.bases = append(w.bases, "S")
		}
	}
	return w.bases
}

func BasePresent(base string, w *World) bool {
	for i := range w.bases {
		if base == w.bases[i] {
			return true
		}
	}
	return false
}

func (w *World) Stats() map[string]int {
	return w.stat
}

func (w *World) StarPort() string {
	if _, ok := w.data[constant.PrStarport]; !ok {
		w.UWP()
	}
	return w.data[constant.PrStarport]
}

// func (w *World) UWPshort() string {
// 	return w.data["uwpShort"]
// }

func (w *World) TradeCodes() []string {
	return w.tradeCodes
}

func (w World) SetTradeCodes(tc []string) World {
	w.tradeCodes = tc
	return w
}

func (w *World) SetPlanetaryData(dataType, dataVal string) {
	w.data[dataType] = dataVal
}

func (w *World) PlanetaryData(dataType string) string {
	return w.data[dataType]
}

// func BuildPlanetUWP(uwp string) *World {
// 	planet := &World{}
// 	code := strings.Split(uwp, "")
// 	planet.stat = make(map[string]int)
// 	planet.data = make(map[string]string)
// 	planet.data[constant.PrStarport] = code[0]
// 	planet.Stat(constant.PrSize] = eHexD(code[1])
// 	planet.Stat(constant.PrAtmo] = eHexD(code[2])
// 	planet.Stat(constant.PrHydr] = eHexD(code[3])
// 	planet.Stat(constant.PrPops] = eHexD(code[4])
// 	planet.Stat(constant.PrGovr] = eHexD(code[5])
// 	planet.Stat(constant.PrLaws] = eHexD(code[6])
// 	planet.Stat(constant.PrTL] = eHexD(code[8])
// 	planet.bases = rollBases(planet.data[constant.PrStarport])
// 	planet.data["uwpShort"] = uwp

// 	//DEBUG OUTPUT:

// 	return planet
// }

func BuildPlanet() *World { //заменить функцию на рандомный сборщих по типу
	planet := &World{}
	planet.stat = make(map[string]int)
	planet.data = make(map[string]string)

	planet.rollSize()
	planet.rollAtmo()
	planet.rollTemperature()
	planet.rollHydr()
	planet.rollPops()
	planet.rollGovr()
	planet.rollLaws()
	planet.rollSpaceport()
	planet.bases = rollBases(planet.data[constant.PrStarport])
	planet.rollTech()
	//travellClassification(planet)
	//tradeClassification(planet)
	planet.buildUWP()

	//DEBUG OUTPUT:
	// fmt.Println("Testing table:")
	// table := utils.NewTxtTable("densityTable.txtTable")
	// fmt.Println(table.Cell(2, 2))
	// fmt.Println("End Testing table:")

	return planet
}

func (w *World) buildUWP() {
	uwp := w.data[constant.PrStarport]
	uwp += dEHex(w.Stat(constant.PrSize))
	uwp += dEHex(w.Stat(constant.PrAtmo))
	uwp += dEHex(w.Stat(constant.PrHydr))
	uwp += dEHex(w.Stat(constant.PrPops))
	uwp += dEHex(w.Stat(constant.PrGovr))
	uwp += dEHex(w.Stat(constant.PrLaws))
	uwp += "-" + dEHex(w.Stat(constant.PrTL))
	//w.data["uwpShort"] = uwp
}

func isDoomed(w *World) bool {
	atmo := w.Stat(constant.PrAtmo)
	tech := w.Stat(constant.PrTL)
	switch atmo {
	case 0, 1:
		if tech < 8 {
			return true
		}
	case 2, 3:
		if tech < 5 {
			return true
		}
	case 4, 7, 9:
		if tech < 3 {
			return true
		}
	case 10:
		if tech < 8 {
			return true
		}
	case 11:
		if tech < 9 {
			return true
		}
	case 12:
		if tech < 10 {
			return true
		}
	case 13, 14:
		if tech < 5 {
			return true
		}
	case 15:
		if tech < 8 {
			return true
		}
	}
	return false
}

func rollPGB() (int, int, int) {
	p := utils.RollDice("1d10", -1)
	r := utils.RollDice("2d6")
	gg := 0
	switch r {
	case 2, 3, 4:
		gg = 0
	case 5:
		gg = 1
	case 6:
		gg = 2
	case 7, 8:
		gg = 3
	case 9, 10:
		gg = 4
	case 11, 12:
		gg = 5
	}
	b := utils.RollDice("2d6", gg)
	belts := 0
	switch b {
	case 2, 3, 4, 5, 6, 7:
		belts = 0
	case 8, 9:
		belts = 1
	case 10, 11, 12:
		belts = 2
	default: //13+
		belts = 3
	}
	return p, gg, belts
}

func calculateEconomic(w *World) string {
	eX := ""

	resourses := utils.RollDice("2d6", -2)
	if utils.ListContains(w.tradeCodes, tradeClassificationBarren) || utils.ListContains(w.tradeCodes, tradeClassificationAsteroidBelt) || utils.ListContains(w.tradeCodes, tradeClassificationPoor) {
		resourses = utils.RollDice("d6", -1)
	}
	if w.Stat(constant.PrTL) >= 8 {
		resourses = resourses + w.Stat("GG") + w.Stat("belts")
	}

	labor := w.Stat(constant.PrPops) - 1

	infrastructure := utils.RollDice("2d6", -2)
	if utils.ListContains(w.tradeCodes, tradeClassificationBarren) {
		switch utils.RollDice("d6") {
		case 1, 2:
			infrastructure = 0
		case 3, 4, 5:
			infrastructure = 1
		case 6:
			infrastructure = 2
		}
	}

	culture := utils.RollDice("2d6", -2)
	if utils.ListContains(w.tradeCodes, tradeClassificationBarren) {
		switch utils.RollDice("d6") {
		default:
			culture = 0
		case 5, 6:
			culture = 1
		}
	}
	if w.Stat(constant.PrPops) == 0 {
		culture = 0
	}

	if utils.ListContains(w.tradeCodes, tradeClassificationAgricultural) {
		resourses++
		culture--
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationAsteroidBelt) {
		infrastructure--
		culture++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationDesert) {
		resourses--
		culture++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationHighPopulation) {
		resourses++
		infrastructure++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationIceCapped) {
		resourses--
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationIndustrial) {
		resourses++
		resourses++
		infrastructure++
		infrastructure++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationLowPopulation) {
		infrastructure--
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationNonAgricultural) {
		resourses--
		culture--
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationNonIndustrial) {
		infrastructure--
		culture--
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationPoor) {
		infrastructure--
		infrastructure--
		culture++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationRich) {
		resourses++
		infrastructure++
		infrastructure++
		culture++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationVacuum) && !utils.ListContains(w.tradeCodes, tradeClassificationAsteroidBelt) {
		resourses--
		culture++
	}
	if utils.ListContains(w.tradeCodes, tradeClassificationWaterWorld) {
		infrastructure--
	}
	switch w.data[constant.PrStarport] {
	case "A":
		resourses += 2
		infrastructure += 4
		culture++
	case "B":
		resourses++
		infrastructure += 3
	case "C":
		infrastructure += 2
	case "D":
		infrastructure++
	}

	if infrastructure > resourses {
		infrastructure = resourses
	}
	if infrastructure > w.Stat(constant.PrTL) {
		infrastructure = w.Stat(constant.PrTL)
	}

	eX = dEHex(resourses) + dEHex(labor) + dEHex(infrastructure) + dEHex(culture)
	w.data[eXResources] = dEHex(resourses)
	w.data[eXLabor] = dEHex(labor)
	w.data[eXInfrastructure] = dEHex(infrastructure)
	w.data[eXCulture] = dEHex(culture)
	return eX
}

func (w *World) Size() int {
	if val, ok := w.data[constant.PrSize]; ok {
		return eHexD(val)
	}
	//fmt.Println("Roll Size")
	w.rollSize()
	return w.Stat(constant.PrSize)
}

func (w *World) rollSize() {
	size := utils.RollDice("2d6", -2)
	if size == 10 {
		size = utils.RollDice("d6", 9)
	}
	w.data[constant.PrSize] = dEHex(size)
}

func (w *World) rollAtmo(dm ...int) {
	size := w.Stat(constant.PrSize)
	atmo := size + TrvCore.Flux()
	if len(dm) > 0 {
		atmo += dm[0]
	}
	if atmo < 0 || size == 0 {
		atmo = 0
	}
	if atmo > eHexD("F") {
		atmo = eHexD("F")
	}
	w.data[constant.PrAtmo] = dEHex(atmo)
}

func (w *World) rollHydr(dm ...int) {
	flux := TrvCore.Flux()
	if len(dm) > 0 {
		flux += dm[0]
	}
	atm := w.Stat(constant.PrAtmo)
	size := w.Stat(constant.PrSize)
	mod := 0
	if size < 2 {
		w.data[constant.PrHydr] = dEHex(0)
		return
	}
	if atm < 2 || atm > 9 {
		mod = mod - 4
	}
	if w.data["Temperature"] == tempHot && w.Stat(constant.PrAtmo) != 13 {
		mod = mod - 2
	}
	if w.data["Temperature"] == tempBoiling && w.Stat(constant.PrAtmo) != 13 {
		mod = mod - 6
	}
	hyd := flux + atm + mod
	hyd = utils.BoundInt(hyd, 0, eHexD("A"))
	w.data[constant.PrHydr] = dEHex(hyd)
}

func (w *World) rollPops(dm ...int) {
	pop := utils.RollDice("2d6", -2)
	if len(dm) > 0 {
		pop += dm[0]
	}
	if pop >= 10 {
		utils.RollDice("2d6", 3)
	}
	if pop < 0 {
		pop = 0
	}
	w.data[constant.PrPops] = dEHex(pop)
}

func (w *World) rollGovr() {
	pop := w.Stat(constant.PrPops)
	gov := TrvCore.Flux() + pop
	gov = utils.BoundInt(gov, 0, eHexD("F"))
	if w.Stat(constant.PrPops) < 1 {
		gov = 0
	}
	w.data[constant.PrGovr] = dEHex(gov)
}

func (w *World) rollLaws() {
	law := TrvCore.Flux() + w.Stat(constant.PrGovr)
	if law < 0 {
		law = 0
	}
	law = utils.BoundInt(law, 0, eHexD("J"))
	if w.Stat(constant.PrPops) < 1 {
		law = 0
	}
	w.data[constant.PrLaws] = dEHex(law)
}

func (w *World) rollSpaceport() {
	dm := 0
	if w.Stat(constant.PrPops) > 7 {
		dm = dm + 1
	}
	if w.Stat(constant.PrPops) > 9 {
		dm = dm + 2
	}
	if w.Stat(constant.PrPops) < 5 {
		dm = dm - 1
	}
	if w.Stat(constant.PrPops) < 3 {
		dm = dm - 2
	}
	port := utils.RollDice("2d6", dm)
	switch port {
	default:
		if port < 3 {
			w.data[constant.PrStarport] = "X"
		}
		if port > 10 {
			w.data[constant.PrStarport] = "A"
		}
	case 3, 4:
		w.data[constant.PrStarport] = "E"
	case 5, 6:
		w.data[constant.PrStarport] = "D"
	case 7, 8:
		w.data[constant.PrStarport] = "C"
	case 9, 10:
		w.data[constant.PrStarport] = "B"
	}
}

func rollBases(port string) []string {
	var bases []string
	switch port {
	case "D":
		if utils.RollDice("2d6") >= 7 {
			bases = append(bases, portBaseScout)
		}
	case "C":
		if utils.RollDice("2d6") >= 8 {
			bases = append(bases, portBaseScout)
		}
		if utils.RollDice("2d6") >= 10 {
			bases = append(bases, portBaseResearch)
		}
		if utils.RollDice("2d6") >= 10 {
			bases = append(bases, portBaseTAS)
		}
	case "B":
		if utils.RollDice("2d6") >= 8 {
			bases = append(bases, portBaseNaval)
		}
		if utils.RollDice("2d6") >= 8 {
			bases = append(bases, portBaseScout)
		}
		if utils.RollDice("2d6") >= 10 {
			bases = append(bases, portBaseResearch)
		}
		if utils.RollDice("2d6") >= 2 {
			bases = append(bases, portBaseTAS)
		}
	case "A":
		if utils.RollDice("2d6") >= 8 {
			bases = append(bases, portBaseNaval)
		}
		if utils.RollDice("2d6") >= 10 {
			bases = append(bases, portBaseScout)
		}
		if utils.RollDice("2d6") >= 8 {
			bases = append(bases, portBaseResearch)
		}
		if utils.RollDice("2d6") >= 2 {
			bases = append(bases, portBaseTAS)
		}
	}
	return bases
}

func (w *World) rollTech() {
	tl := utils.RollDice("d6")
	dm := 0
	switch w.data[constant.PrStarport] {
	case "A":
		dm = dm + 6
	case "B":
		dm = dm + 4
	case "C":
		dm = dm + 2
	case "X":
		dm = dm - 4
	case "F":
		dm = dm + 1
	}
	switch w.Stat(constant.PrSize) {
	case 0, 1:
		dm = dm + 2
	case 2, 3, 4:
		dm = dm + 1
	}
	switch w.Stat(constant.PrAtmo) {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		dm = dm + 1
	}
	switch w.Stat(constant.PrHydr) {
	case 0, 9:
		dm = dm + 1
	case 10:
		dm = dm + 2
	}
	switch w.Stat(constant.PrPops) {
	case 1, 2, 3, 4, 5:
		dm = dm + 1
	case 9:
		dm = dm + 2
	case 10, 11, 12, 13, 14, 15:
		dm = dm + 4
	}
	switch w.Stat(constant.PrGovr) {
	case 0, 5:
		dm = dm + 1
	// case 7:
	// 	dm = dm + 2
	case 13, 14:
		dm = dm - 2
	}
	w.data[constant.PrTL] = dEHex(tl + dm)
	if w.Stat(constant.PrTL) < 0 {
		w.data[constant.PrTL] = dEHex(0)
	}
}

func (w *World) rollTemperature() {
	dm := 0
	switch w.Stat(constant.PrAtmo) {
	case 2, 3:
		dm = -2
	case 4, 5, 14:
		dm = -1
	case 6, 7:
		dm = 0
	case 8, 9:
		dm = 1
	case 10, 13, 15:
		dm = 2
	case 11, 12:
		dm = 6
	}
	temp := utils.RollDice("2d6", dm)
	switch temp {
	case 3, 4:
		w.data["Temperature"] = tempCold
	case 5, 6, 7, 8, 9:
		w.data["Temperature"] = tempTemperate
	case 10, 11:
		w.data["Temperature"] = tempHot
	default:
		if temp < 3 {
			w.data["Temperature"] = tempFrozen
		}
		if temp > 11 {
			w.data["Temperature"] = tempBoiling
		}
	}

}

func boundInt(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}

func isInRange(i, min, max int) bool {
	if i < min {
		return false
	}
	if i > max {
		return false
	}
	return true
}

func matchValue(val int, check ...int) bool {
	for i := range check {
		if check[i] == val {
			return true
		}
	}
	return false
}

func TradeClassificationsFULLLIST() []string {
	tcFL := []string{
		tradeClassificationAsteroidBelt,
		tradeClassificationDesert,
		tradeClassificationFluid,
		tradeClassificationGardenWorld,
		tradeClassificationHellworld,
		tradeClassificationIceCapped,
		tradeClassificationOceanWorld,
		tradeClassificationVacuum,
		tradeClassificationWaterWorld,
		tradeClassificationDieback,
		tradeClassificationBarren,
		tradeClassificationLowPopulation,
		tradeClassificationNonIndustrial,
		tradeClassificationPreHigh,
		tradeClassificationHighPopulation,
		tradeClassificationLowTech,
		tradeClassificationHighTech,
		tradeClassificationPreAgricultural,
		tradeClassificationAgricultural,
		tradeClassificationNonAgricultural,
		tradeClassificationPrisonExileCamp,
		tradeClassificationPreIndustrial,
		tradeClassificationIndustrial,
		tradeClassificationPoor,
		tradeClassificationPreRich,
		tradeClassificationRich,
		tradeClassificationFrozen,
		tradeClassificationHot,
		tradeClassificationCold,
		tradeClassificationLocked,
		tradeClassificationTropic,
		tradeClassificationTundra,
		tradeClassificationTwilightZone,
		tradeClassificationFarming,
		tradeClassificationMining,
		tradeClassificationMilitaryRule,
		tradeClassificationPenalColony,
		tradeClassificationReserve,
		tradeClassificationSubsectorCapital,
		tradeClassificationSectorCapital,
		tradeClassificationCapital,
		tradeClassificationColony,
		tradeClassificationSatellite,
		tradeClassificationForbidden,
		tradeClassificationPuzzle,
		tradeClassificationDangerous,
		tradeClassificationDataRepository,
		tradeClassificationAncientSite,
	}
	return tcFL
}

func (w *World) UpdateTradeClassifications() {
	tradeCodesFullList := TradeClassificationsFULLLIST()
	for i := range tradeCodesFullList {
		if TradeCodeViable(w, tradeCodesFullList[i]) {
			w.EnsureTradeCode(tradeCodesFullList[i])
		}
	}
}

func (w World) UpdateTC() World {
	tradeCodesFullList := TradeClassificationsFULLLIST()
	for i := range tradeCodesFullList {
		if TradeCodeViable(&w, tradeCodesFullList[i]) {
			w.EnsureTradeCode(tradeCodesFullList[i])
		}
	}
	return w
}

func eHex(s string) int {
	return eHexD(s)
}

func checkStat(stat int, valArray string) bool {
	valRange := strings.Split(valArray, "")
	for i := range valRange {
		if stat == eHexD(valRange[i]) {
			return true
		}
	}
	return false
}

func matchTradeClassificationRequirements(w *World, reqLine string) bool {
	stats := strings.Split(reqLine, " ") //-- 23456789 0 -- -- --
	ehexList := TrvCore.ValidEhexs()
	fullArray := ""
	for i := range ehexList {
		fullArray = fullArray + ehexList[i]
	}
	statArray := []string{constant.PrSize, constant.PrAtmo, constant.PrHydr, constant.PrPops, constant.PrGovr, constant.PrLaws}
	for i := range stats { //собираем аррэй
		array := stats[i]
		if array == "--" {
			array = fullArray
		}
		if !checkStat(w.Stat(statArray[i]), array) {
			return false
		}
	}
	return true
}

//TradeCodeViable - эта версия считает TC по корнику Mongoose Traveller 2E (p. 228)
//TODO: переписать функцию так чтобы исключить зависмость структуры World (поместить ее в Profile)
func TradeCodeViable(w *World, tc string) bool {
	switch tc {
	default:
		return false
	case tradeClassificationAgricultural:
		if matchTradeClassificationRequirements(w, "-- 456789 45678 567 -- --") {
			return true
		}
	case tradeClassificationAsteroidBelt:
		if matchTradeClassificationRequirements(w, "0 0 0 -- -- --") {
			return true
		}
	case tradeClassificationBarren:
		if matchTradeClassificationRequirements(w, "-- -- -- 0 0 0") {
			return true
		}
	case tradeClassificationDesert:
		if matchTradeClassificationRequirements(w, "-- 23456789ABCDEFS 0 -- -- --") {
			return true
		}
	case tradeClassificationFluid:
		if matchTradeClassificationRequirements(w, "-- ABCDEF 123456789A -- -- --") {
			return true
		}
	case tradeClassificationGardenWorld:
		if matchTradeClassificationRequirements(w, "678 568 567 -- -- --") {
			return true
		}
	case tradeClassificationHighPopulation:
		if matchTradeClassificationRequirements(w, "-- -- -- 9ABCDEF -- --") {
			return true
		}
	case tradeClassificationHighTech:
		if isInRange(w.Stat(constant.PrTL), 12, 30) {
			return true
		}
	case tradeClassificationIceCapped:
		if matchTradeClassificationRequirements(w, "-- 01 123456789A -- -- --") {
			return true
		}
	case tradeClassificationIndustrial:
		if matchTradeClassificationRequirements(w, "-- 012479 -- 9ABCDEF -- --") {
			return true
		}
	case tradeClassificationLowPopulation:
		if matchTradeClassificationRequirements(w, "-- -- -- 123 -- --") {
			return true
		}
	case tradeClassificationLowTech:
		if isInRange(w.Stat(constant.PrTL), 0, 5) {
			return true
		}
	case tradeClassificationNonAgricultural:
		if matchTradeClassificationRequirements(w, "-- 0123 0123 6789ABCDEF -- --") {
			return true
		}
	case tradeClassificationNonIndustrial:
		if matchTradeClassificationRequirements(w, "-- -- -- 123456 -- --") {
			return true
		}
	case tradeClassificationPoor:
		if matchTradeClassificationRequirements(w, "-- 2345 0123 -- -- --") {
			return true
		}
	case tradeClassificationRich:
		if matchTradeClassificationRequirements(w, "-- 68 -- 678 456789 --") {
			return true
		}
	case tradeClassificationVacuum:
		if matchTradeClassificationRequirements(w, "-- 0 -- -- -- --") {
			return true
		}
	case tradeClassificationWaterWorld:
		if matchTradeClassificationRequirements(w, "-- -- A -- -- --") {
			return true
		}

	}
	return false
}

func TradeCodeViableT5(w *World, tc string) bool {
	switch tc {
	default:
		return false
	case tradeClassificationAsteroidBelt:
		if matchTradeClassificationRequirements(w, "0 0 0 -- -- --") {
			return true
		}
	case tradeClassificationDesert:
		if matchTradeClassificationRequirements(w, "-- 23456789 0 -- -- --") {
			return true
		}
	case tradeClassificationFluid:
		if matchTradeClassificationRequirements(w, "-- ABC 123456789A -- -- --") {
			return true
		}
	case tradeClassificationGardenWorld:
		if matchTradeClassificationRequirements(w, "678 568 567 -- -- --") {
			return true
		}
	case tradeClassificationHellworld:
		if matchTradeClassificationRequirements(w, "3456789ABC 2479ABC 012 -- -- --") {
			return true
		}
	case tradeClassificationIceCapped:
		if matchTradeClassificationRequirements(w, "-- 01 123456789A -- -- --") {
			return true
		}
	case tradeClassificationOceanWorld:
		if matchTradeClassificationRequirements(w, "ABCDEF 3456789DEF A -- -- --") {
			return true
		}
	case tradeClassificationVacuum:
		if matchTradeClassificationRequirements(w, "-- 0 -- -- -- --") {
			return true
		}
	case tradeClassificationWaterWorld:
		if matchTradeClassificationRequirements(w, "3456789 3456789DEF A -- -- --") {
			return true
		}
	case tradeClassificationDieback:
		if matchTradeClassificationRequirements(w, "-- -- -- 0 0 0") {
			return true
		}
	case tradeClassificationBarren:
		if matchTradeClassificationRequirements(w, "-- -- -- 0 0 0") {
			return true
		}
	case tradeClassificationLowPopulation:
		if matchTradeClassificationRequirements(w, "-- -- -- 123 -- --") {
			return true
		}
	case tradeClassificationNonIndustrial:
		if matchTradeClassificationRequirements(w, "-- -- -- 456 -- --") {
			return true
		}
	case tradeClassificationPreHigh:
		if matchTradeClassificationRequirements(w, "-- -- -- 8 -- --") {
			return true
		}
	case tradeClassificationHighPopulation:
		if matchTradeClassificationRequirements(w, "-- -- -- 9ABCDEF -- --") {
			return true
		}
	case tradeClassificationPreAgricultural:
		if matchTradeClassificationRequirements(w, "-- 456789 45678 48 -- --") {
			return true
		}
	case tradeClassificationAgricultural:
		if matchTradeClassificationRequirements(w, "-- 456789 45678 567 -- --") {
			return true
		}
	case tradeClassificationNonAgricultural:
		if matchTradeClassificationRequirements(w, "-- 0123 0123 6789ABCDEF -- --") {
			return true
		}
	case tradeClassificationPrisonExileCamp:
		if matchTradeClassificationRequirements(w, "-- 23AB 12345 3456 6789") {
			return true
		}
	case tradeClassificationPreIndustrial:
		if matchTradeClassificationRequirements(w, "-- 012479 -- 78 -- --") {
			return true
		}
	case tradeClassificationIndustrial:
		if matchTradeClassificationRequirements(w, "-- 012479ABC -- 9ABCDEF -- --") {
			return true
		}
	case tradeClassificationPoor:
		if matchTradeClassificationRequirements(w, "-- 2345 0123 -- -- --") {
			return true
		}
	case tradeClassificationPreRich:
		if matchTradeClassificationRequirements(w, "-- 68 -- 59 -- --") {
			return true
		}
	case tradeClassificationRich:
		if matchTradeClassificationRequirements(w, "-- 68 -- 678 -- --") {
			return true
		}
	case tradeClassificationFrozen:
		if matchTradeClassificationRequirements(w, "23456789 -- 123456789A -- -- --") {
			//if w.WorldOrbitHZ() > 1 {
			if w.data["Temperature"] == tempFrozen {
				return true
			}
			//}
			//return false
		}
	case tradeClassificationHot:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			//if w.WorldOrbitHZ() == -1 {
			if w.data["Temperature"] == tempHot {
				return true
			}
			//}
			//return false
		}
	case tradeClassificationCold:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			//if w.WorldOrbitHZ() == 1 {
			if w.data["Temperature"] == tempCold {
				return true
			}
			//}
			//return false
		}
	case tradeClassificationLocked:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.data["PlanetType"] == "Close Satellite" {
				return true
			}
			if w.data["Orbit"] == "0" {
				return true
			}
			// if w.PlanetType() == "Close Satellite" {
			// 	return true
			// }
			// if w.Orbit() == 0 || w.Orbit() == 1 {
			// 	return true
			// }

		}
	case tradeClassificationTropic:
		if matchTradeClassificationRequirements(w, "6789 456789 34567 -- -- --") {
			//if w.WorldOrbitHZ() == -1 {
			if w.data["Temperature"] == "Tropical" {
				return true
			}
			//}
			//return true
		}
	case tradeClassificationTundra:
		if matchTradeClassificationRequirements(w, "6789 456789 34567 -- -- --") {
			//if w.WorldOrbitHZ() == 1 {
			if w.data["Temperature"] == "Tundra" {
				return true
			}
			//}
		}
	case tradeClassificationTwilightZone:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			//return true TODO: Орбита 0-1
		}
	case tradeClassificationFarming:
		if matchTradeClassificationRequirements(w, "-- 456789 45678 23456 -- --") {
			//return true TODO mainworld = false
			if w.data["MainWorld"] == "FALSE" {
				return true
			}
		}
	case tradeClassificationMining:
		if matchTradeClassificationRequirements(w, "-- -- -- 23456 -- --") {
			if w.data["MainWorld"] == "FALSE" {
				return true
			}
		}
	case tradeClassificationMilitaryRule:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.data["MainWorld"] == "FALSE" {
				return true
			}
		}
	case tradeClassificationPenalColony:
		if matchTradeClassificationRequirements(w, "-- 23AB 12345 3456 6 6789") {
			if w.data["MainWorld"] == "FALSE" {
				return true
			}
		}
	case tradeClassificationReserve:
		if matchTradeClassificationRequirements(w, "-- -- -- 1234 6 45") {
			if w.data["MainWorld"] == "FALSE" {
				return true
			}
		}
	case tradeClassificationSubsectorCapital:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.data["SubsectorCapital"] == "TRUE" {
				return true
			}
		}
	case tradeClassificationSectorCapital:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.data["SectorCapital"] == "TRUE" {
				return true
			}
		}
	case tradeClassificationCapital:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.data["Capital"] == "TRUE" {
				return true
			}
		}
	case tradeClassificationColony:
		if matchTradeClassificationRequirements(w, "-- -- -- 56789A 6 0123") {
			if w.data["Colony"] == "TRUE" {
				return true
			}
		}
	case tradeClassificationSatellite:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.PlanetType() == "Far Satellite" {
				return true
			}

		}
	case tradeClassificationForbidden:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if w.travelCode == "Red Zone" {
				return true
			}

		}
	case tradeClassificationPuzzle:
		if matchTradeClassificationRequirements(w, "-- -- -- 789ABCDEF -- --") {
			if w.travelCode == "Amber Zone" {
				return true
			}
		}
	case tradeClassificationDangerous:
		if matchTradeClassificationRequirements(w, "-- -- -- 0123456 -- --") {
			if w.travelCode == "Amber Zone" {
				return true
			}
		}
	case tradeClassificationDataRepository:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if val, ok := w.data["DataRepository"]; ok {
				if val == "TRUE" {
					return true
				}
			}
		}
	case tradeClassificationAncientSite:
		if matchTradeClassificationRequirements(w, "-- -- -- -- -- --") {
			if val, ok := w.data["AncientSite"]; ok {
				if val == "TRUE" {
					return true
				}
			}
		}
	case tradeClassificationLowTech:
		if isInRange(w.Stat(constant.PrTL), 0, 5) {
			return true
		}
	case tradeClassificationHighTech:
		if isInRange(w.Stat(constant.PrTL), 12, 30) {
			return true
		}

	}
	return false
}

// func tradeClassification(w *World) []string {
// 	var tradecodes []string

// 	if isInRange(w.Stat(constant.PrAtmo], 4, 9) && isInRange(w.Stat(constant.PrHydr], 4, 8) && matchValue(w.Stat(constant.PrPops], 5, 6, 7) {
// 		tradecodes = append(tradecodes, tradeClassificationAgricultural)
// 	}
// 	if w.Stat(constant.PrSize] == 0 && w.Stat(constant.PrAtmo] == 0 && w.Stat(constant.PrHydr] == 0 {
// 		tradecodes = append(tradecodes, tradeClassificationAsteroidBelt)
// 	}
// 	if isInRange(w.Stat(constant.PrPops], 0, 0) && isInRange(w.Stat(constant.PrGovr], 0, 0) && isInRange(w.Stat(constant.PrLaws], 0, 0) {
// 		tradecodes = append(tradecodes, tradeClassificationBarren)
// 	}
// 	if isInRange(w.Stat(constant.PrAtmo], 2, 15) && w.Stat(constant.PrHydr] == 0 {
// 		tradecodes = append(tradecodes, tradeClassificationDesert)
// 	}
// 	if isInRange(w.Stat(constant.PrAtmo], 10, 25) && isInRange(w.Stat(constant.PrHydr], 1, 15) {
// 		tradecodes = append(tradecodes, tradeClassificationFluid)
// 	}
// 	if isInRange(w.Stat(constant.PrSize], 6, 8) && matchValue(w.Stat(constant.PrAtmo], 5, 6, 8) && isInRange(w.Stat(constant.PrHydr], 5, 7) {
// 		tradecodes = append(tradecodes, tradeClassificationGardenWorld)
// 	}
// 	if isInRange(w.Stat(constant.PrPops], 9, 30) {
// 		tradecodes = append(tradecodes, tradeClassificationHighPopulation)
// 	}
// 	if isInRange(w.Stat(constant.PrTL], 12, 30) {
// 		tradecodes = append(tradecodes, tradeClassificationHighTech)
// 	}
// 	if matchValue(w.Stat(constant.PrAtmo], 0, 1) && isInRange(w.Stat(constant.PrHydr], 1, 10) {
// 		tradecodes = append(tradecodes, tradeClassificationIceCapped)
// 	}
// 	if matchValue(w.Stat(constant.PrAtmo], 0, 1, 2, 4, 7, 9) && isInRange(w.Stat(constant.PrPops], 9, 25) {
// 		tradecodes = append(tradecodes, tradeClassificationIndustrial)
// 	}
// 	if isInRange(w.Stat(constant.PrPops], 0, 3) {
// 		tradecodes = append(tradecodes, tradeClassificationLowPopulation)
// 	}
// 	if isInRange(w.Stat(constant.PrTL], 0, 5) {
// 		tradecodes = append(tradecodes, tradeClassificationLowTech)
// 	}
// 	if isInRange(w.Stat(constant.PrAtmo], 0, 3) && isInRange(w.Stat(constant.PrHydr], 0, 3) && isInRange(w.Stat(constant.PrPops], 6, 20) {
// 		tradecodes = append(tradecodes, tradeClassificationNonAgricultural)
// 	}
// 	if isInRange(w.Stat(constant.PrPops], 0, 6) {
// 		tradecodes = append(tradecodes, tradeCodeNonIndustrial)
// 	}
// 	if isInRange(w.Stat(constant.PrAtmo], 2, 5) && isInRange(w.Stat(constant.PrHydr], 0, 3) {
// 		tradecodes = append(tradecodes, tradeCodePoor)
// 	}
// 	if matchValue(w.Stat(constant.PrAtmo], 6, 8) && matchValue(w.Stat(constant.PrPops], 6, 7, 8) && isInRange(w.Stat(constant.PrGovr], 4, 9) {
// 		tradecodes = append(tradecodes, tradeCodeRich)
// 	}
// 	if matchValue(w.Stat(constant.PrAtmo], 0) {
// 		tradecodes = append(tradecodes, tradeCodeVacuum)
// 	}
// 	if isInRange(w.Stat(constant.PrHydr], 10, 10) {
// 		tradecodes = append(tradecodes, tradeCodeWaterWorld)
// 	}
// 	w.tradeCodes = tradecodes
// 	return tradecodes
// }

// func travellClassification(w *World) string {
// 	var travellcode string
// 	if isInRange(w.Stat(constant.PrAtmo], 10, 99) || matchValue(w.Stat(constant.PrGovr], 0, 7, 10) || isInRange(w.Stat(constant.PrLaws], 9, 99) || matchValue(w.Stat(constant.PrLaws], 0) {
// 		travellcode = "Amber Zone"
// 	}
// 	if w.Stat(constant.PrLaws]+w.Stat(constant.PrGovr) >= 22 {
// 		travellcode = "Red Zone"
// 	}
// 	if w.data[constant.PrStarport] == "X" {
// 		switch utils.RollDice("d6") {
// 		default:
// 			travellcode = "Red Zone"
// 		case 5, 6:
// 		}

// 	}
// 	w.travelCode = travellcode
// 	return travellcode
// }

func (w *World) HomeStar(refereeDM ...int) string {
	if _, ok := w.data["HomeStar"]; ok {
		return w.data["HomeStar"]
	}
	w.data["HomeStar"] = table2aHomeStar(refereeDM...)
	//w.stellar = append(w.stellar, w.data["HomeStar"])
	return w.data["HomeStar"]
}

func table2aHomeStar(refereeDM ...int) string {
	dm := 0
	if len(refereeDM) > 1 {
		dm = refereeDM[0]
	}
	dm = utils.BoundInt(dm, -1, 1)
	flux1 := TrvCore.Flux() + dm
	flux2 := TrvCore.Flux()
	class := []string{"O", "B", "A", "A", "F", "F", "G", "K", "K", "M", "M", "M", "M"}
	dec := utils.RollDice("d10", -1)
	hsMap := make(map[int][]string)
	hsMap[-6] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "V", "V", "V", "IV", "D", "D"}
	hsMap[-5] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "III", "V", "V", "IV", "D", "D"}
	hsMap[-4] = []string{"Ia", "Ia", "Ib", "II", "III", "IV", "V", "V", "V", "V", "V", "D", "D"}
	hsMap[-3] = []string{"Ia", "Ia", "Ib", "II", "III", "IV", "V", "V", "V", "V", "V", "D", "D"}
	hsMap[-2] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[-1] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+0] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+1] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+2] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+3] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+4] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+5] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "D"}
	hsMap[+6] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "D"}

	//fmt.Println(class[flux1+6], convert.ItoS(dec), hsMap[flux1][flux2+6])
	homeStar := class[flux1+6] + convert.ItoS(dec) + "" + hsMap[flux1][flux2+6]
	return homeStar
}

func (w *World) WorldOrbitHZ() int {
	if w.data["HZVariance"] != "" {
		return convert.StoI(w.data["HZVariance"])
	}
	w.data["HZVariance"] = table2bWorldOrbit()

	if w.data["HZVariance"] == "0" {
		w.data["Temperature"] = tempTemperate
	}
	w.data["Orbit"] = dEHex(calculateHZ(w.HomeStar()) + convert.StoI(w.data["HZVariance"]))
	if w.Stat("Orbit") < 0 {
		w.data["Orbit"] = dEHex(0)
	}
	return convert.StoI(w.data["HZVariance"])
}

func (w *World) Orbit() int {
	if val, ok := w.data["Orbit"]; ok {
		return eHexD(val)
	}
	w.HomeStar()
	w.WorldOrbitHZ()
	return w.Stat("Orbit")
}

func table2bWorldOrbit() string {
	hz := []string{"-2", "-1", "-1", "0", "0", "0", "0", "0", "+1", "+1", "+2"}
	flux := TrvCore.Flux()
	return hz[flux+5]
}

func (w *World) EnsureTradeCode(tCode string) {
	w.tradeCodes = utils.AppendUniqueStr(w.tradeCodes, tCode)
}

func (w *World) PlanetType() string {
	if w.data["PlanetType"] != "" {
		return w.data["PlanetType"]
	}
	pt := []string{"Far Satellite", "Far Satellite", "Close Satellite", "Planet", "Planet", "Planet", "Planet", "Planet", "Planet", "Planet", "Planet"}
	flux := TrvCore.Flux()
	w.data["PlanetType"] = pt[flux+5]
	return w.data["PlanetType"]
}

func (w *World) PBG() string {
	if w.pbg != "" {
		return w.pbg
	}
	w.SetPBG("RANDOM")
	return w.pbg
}

func readPBG(pbg string) (p int, b int, g int) {
	data := strings.Split(pbg, "")
	if len(data) < 3 {
		panic("PBG data not available!")
	}
	p = eHexD(data[0])
	b = eHexD(data[1])
	g = eHexD(data[2])
	return p, b, g
}

func (w *World) SetPBG(pbg string) {
	if pbg == "RANDOM" {
		p := 0
		for p == 0 {
			p = utils.RollDice("d10", -1)
		}
		if w.Stat(constant.PrPops) == 0 {
			p = 0
		}
		b := utils.RollDice("d6", -3)
		if b < 0 {
			b = 0
		}
		g := (utils.RollDice("2d6") / 2) - 2
		if g < 0 {
			g = 0
		}
		pbg = convert.ItoS(p) + convert.ItoS(b) + convert.ItoS(g)
		w.data[planetStatPopDigit] = dEHex(p)
		w.data[planetStatBelt] = dEHex(b)
		w.data[planetStatGasG] = dEHex(g)
	}
	w.pbg = pbg
	if !pbgValid(pbg) {
		return
		//pbg = "RANDOM"
	}
}

func pbgValid(pbg string) bool {
	if len(pbg) != 3 {
		return false
	}
	data := strings.Split(pbg, "")
	for i := range data {
		_, err := strconv.Atoi(data[i])
		if err != nil {
			return false
		}
	}
	return true
}

func (w *World) IsMainWorld(mw bool) {
	if mw {
		w.data["MainWorld"] = "TRUE"
	} else {
		w.data["MainWorld"] = "FALSE"
	}
}

func (w *World) UpdateTravelZone() {
	w.travelCode = "Green Zone"
	if w.Stat(constant.PrGovr)+w.Stat(constant.PrLaws) >= 20 {
		w.travelCode = "Amber Zone"
	}
	if w.Stat(constant.PrGovr)+w.Stat(constant.PrLaws) >= 22 {
		w.travelCode = "Red Zone"
	}
}

func (w *World) TravelZone() string {
	if w.travelCode == "" {
		w.UpdateTravelZone()
	}
	return w.travelCode
}

/*
0 0 0 -- -- --
-- 23456789 0 -- -- --
-- ABC 123456789A -- -- --
678 568 567 -- -- --
3456789ABC 2479ABC 012 -- -- --
-- 01 123456789A -- -- --
ABCDEF 3456789DEF A -- -- --
-- 0 -- -- -- --
3456789 3456789DEF A -- -- --
-- -- -- 0 0 0
-- -- -- 0 0 0
-- -- -- 123 -- --
-- -- -- 456 -- --
-- -- 8 -- --
-- -- -- 9ABCDEF -- --
-- 456789 45678 48 -- --
-- 456789 45678 567 -- --
-- 0123 0123 6789ABCDEF -- --
-- 23AB 12345 3456 6789
-- 012479 -- 78 -- --
-- 012479ABC -- 9ABCDEF -- --
-- 2345 0123 -- -- --
-- 68 -- 59 -- --
-- 68 -- 678 -- --
23456789 -- 123456789A -- -- --
-- -- -- -- -- -- Hot
-- -- -- -- -- -- Cold
-- -- -- -- -- -- Locked
6789 456789 34567 -- -- --
6789 456789 34567 -- -- --
-- -- -- -- -- -- Twilight Zone
-- 456789 45678 23456 -- --
-- -- -- 23456 -- --
-- -- -- -- -- --
-- 23AB 12345 3456 6 6789
-- -- -- 1234 6 45
-- -- -- -- -- --
-- -- -- -- -- --
-- -- -- -- -- --
-- -- -- 56789A 6 0123 Colony --
-- -- -- -- -- --
-- -- -- -- -- --
-- -- -- 789ABCDEF -- --
-- -- -- 0123456 -- --
-- -- -- -- -- --
-- -- -- -- -- --
*/

func rollSpectTypeAndSize(dm int) string {
	flux1 := TrvCore.Flux() + dm
	flux1 = utils.BoundInt(flux1, -6, 8)
	SpecClass := []string{utils.RandomFromList([]string{"O", "B"}), "A", "A", "F", "F", "G", "G", "K", "K", "M", "M", "M", "BD", "BD", "BD"}
	spectral := SpecClass[flux1+6]
	if spectral == "BD" {
		return spectral + " "
	}
	flux2 := TrvCore.Flux() + dm
	flux2 = utils.BoundInt(flux2, -6, 8)
	size := make(map[string][]string)
	size["O"] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "V", "V", "V", "IV", "D", "IV", "IV", "IV"}
	size["B"] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "III", "V", "V", "IV", "D", "IV", "IV", "IV"}
	size["A"] = []string{"Ia", "Ia", "Ib", "II", "III", "IV", "V", "V", "V", "V", "V", "D", "V", "V", "V"}
	size["F"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	size["G"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	size["K"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	size["M"] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	//TODO: Size IV not for K5-K9 and M0-M9.
	//TODO: Size VI not for A0-A9 and F0-F4.
	//fmt.Println(spectral, flux2+6)
	sizeStr := size[spectral][flux2+6]
	dec := utils.RollDice("d10", -1)
	decStr := strconv.Itoa(dec)
	return spectral + decStr + sizeStr
}

func calculateHZ(starType string) (hz int) {
	// starType M6III
	data := strings.Split(starType, "")
	spec := data[0]
	size := strings.Join(data[2:], "")
	hzMap := make(map[string][]int)
	hzMap["O"] = []int{15, 15, 14, 13, 12, 11, -99, 1}
	hzMap["B"] = []int{13, 13, 12, 11, 10, 9, -99, 0}
	hzMap["A"] = []int{12, 11, 9, 7, 7, 7, -99, 0}
	hzMap["F"] = []int{11, 10, 9, 6, 6, 4, 3, 0}
	hzMap["G"] = []int{12, 10, 9, 7, 5, 3, 2, 0}
	hzMap["K"] = []int{12, 10, 9, 8, 5, 2, 1, 0}
	hzMap["M"] = []int{12, 11, 10, 9, -99, 0, 0, 0}
	key := 0
	switch size {
	case "Ia":
		key = 0
	case "Ib":
		key = 1
	case "II":
		key = 2
	case "III":
		key = 3
	case "IV":
		key = 4
	case "V":
		key = 5
	case "VI":
		key = 6
	case "D":
		key = 7
	}
	return hzMap[spec][key]
}

/*
System Presence
System Details
System Nature
Primary Star Type and Size
Primary Star Decimal Classification
Companion Star Type and Size
Companion Orbit
Far Companion
Far Orbit
Additional Stars
Maximum Orbits
Available Orbits
Orbit Zones for Star:
	Orbit Zones for Star Size Ia
	Orbit Zones for Star Size Ib
	Orbit Zones for Star Size II
	Orbit Zones for Star Size III
	Orbit Zones for Star Size IV
	Orbit Zones for Star Size V
	Orbit Zones for Star Size VI
	Orbit Zones for Star Size D
Gas Gigants
Planetoid Belts
Empty Orbits
Captured Planets
Place Gas Gigants
Place Planetoid Belts
Place Mainworld
Other Worlds
World Size
Atmosphere
Hydrographics
Population
Sattellite
Sattellite Size
Sattellite Orbit
Sattellite Atmosphere
Sattellite Hydrosphere
Sattellite Population
Social Data
Subordinate Goverment
Subordinate Law Level
Facilities
Subordinate Tech Level
Subordinate Spaceport
Orbital Distances



*/

func rollNMWStarport(pop int) string {
	r := utils.RollDice("d6")
	d := pop - r
	switch d {
	case 0:
		return "Y"
	case 1, 2:
		return "H"
	case 3:
		return "G"
	default:
		if d < 1 {
			return "Y"
		}
		return "F"
	}
}

func NewGasGigant() *World {
	gg := World{}
	gg.stat = make(map[string]int)
	gg.data = make(map[string]string)
	size := utils.RollDice("2d6", 19)
	gg.data[constant.PrSize] = dEHex(size)
	gg.SetUWP("Y" + dEHex(size) + "00000-0")
	fmt.Println(distributeSatteliteOrbits(utils.RollDice("d6", -1)))
	return &gg
}

func distributeSatteliteOrbits(orbitQ int) []string {
	orbs := []string{}
	for len(orbs) < orbitQ {
		if utils.RandomBool() {
			orbs = utils.AppendUniqueStr(orbs, rollCloseOrbit())
		} else {
			orbs = utils.AppendUniqueStr(orbs, rollFarOrbit())
		}
	}
	return orbs
}

func rollCloseOrbit() string {
	switch TrvCore.Flux() {
	case -6:
		return "Ay"
	case -5:
		return "Bee"
	case -4:
		return "See"
	case -3:
		return "Dee"
	case -2:
		return "Ee"
	case -1:
		return "Eff"
	case 0:
		return "Gee"
	case 1:
		return "Aitch"
	case 2:
		return "Eye"
	case 3:
		return "Jay"
	case 4:
		return "Kay"
	case 5:
		return "Ell"
	case 6:
		return "Em"
	}
	return "ERROR"
}

func rollFarOrbit() string {
	switch TrvCore.Flux() {
	case -6:
		return "En"
	case -5:
		return "Oh"
	case -4:
		return "Pee"
	case -3:
		return "Cue"
	case -2:
		return "Arr"
	case -1:
		return "Ess"
	case 0:
		return "Tee"
	case 1:
		return "You"
	case 2:
		return "Vee"
	case 3:
		return "Double-You"
	case 4:
		return "Eks"
	case 5:
		return "Wye"
	case 6:
		return "Zee"
	}
	return "ERROR"
}

// func (uwp *UWP) SetData(dataType int, data string) {
// 	bt := []byte(uwp.data)
// 	return string(bt[dataType])
// }

// func NewUWP(planetType ...string) UWP {
// 	plType := "Main World"
// 	if len(planetType) != 0 {
// 		plType = planetType[0]
// 	}
// 	uwp := UWP{data: "_______-_"}
// 	switch plType {
// 	default:
// 		if uwp.DataType(1) == "_" {

// 		}
// 	}
// 	return uwp
// }

func (w World) Stat(key string) int {
	return eHexD(w.data[key])
}

func (w World) PullData() map[string]string {
	return w.data
}

func (w World) ValueOf(key string) string {
	if key == constant.DIVIDER {
		return "-"
	}
	if _, ok := w.data[key]; ok {
		return w.data[key]
	}
	return "_"
}

func (w World) SetValue(dataKey, val string) {
	if val == "_" {
		return
	}
	//	if _, ok := w.data[dataKey]; ok {
	w.data[dataKey] = val
	//	}
}

func dEHex(i int) string {
	return TrvCore.DigitToEhex(i)
}

func eHexD(s string) int {
	return TrvCore.EhexToDigit(s)
}

func (w *World) MergeUWP(uwp string) {
	if !uwpValid(uwp) {
		return
	}
	data := strings.Split(uwp, "")
	w.SetValue(constant.PrStarport, data[0])
	w.SetValue(constant.PrSize, data[1])
	w.SetValue(constant.PrAtmo, data[2])
	w.SetValue(constant.PrHydr, data[3])
	w.SetValue(constant.PrPops, data[4])
	w.SetValue(constant.PrGovr, data[5])
	w.SetValue(constant.PrLaws, data[6])
	w.SetValue(constant.PrTL, data[8])
}

func FromUWP(uwp string) World {
	if !uwpValid(uwp) {
		fmt.Println("UWP invalid")
		return World{}
	}
	w := World{}
	w.data = make(map[string]string)
	w.MergeUWP(uwp)
	return w
}

func (w World) SetName(newName string) World {
	w.name = newName
	return w
}

func FromOTUdata(otuData string) (World, error) {
	w := World{}
	data := strings.Split(otuData, "	")
	if len(data) != 17 {
		return w, errors.New("OTU data unparseble: (Len != 17)")
	}
	w = NewWorld(otu.Info{otuData}.Name())
	w.data["SS"] = otu.Info{otuData}.SubSector()
	w.data["Hex"] = otu.Info{otuData}.Hex()
	w.SetUWP(otu.Info{otuData}.UWP())
	w.bases = otu.Info{otuData}.Bases()
	w.tradeCodes = otu.Info{otuData}.Remarks()
	w.checkLtHtTradeCodes()
	w.travelCode = otu.Info{otuData}.Zone()
	w.pbg = otu.Info{otuData}.PBG()
	w.data["Allegiance"] = otu.Info{otuData}.Allegiance()
	w.data["Stars"] = otu.Info{otuData}.Stars()
	w.data["Ix"] = otu.Info{otuData}.Iextention()
	w.data["Ex"] = otu.Info{otuData}.Eextention()
	w.data["Cx"] = otu.Info{otuData}.Cextention()
	w.data["Nobility"] = otu.Info{otuData}.Nobility()
	w.data["Worlds"] = otu.Info{otuData}.Worlds()
	w.data["RU"] = otu.Info{otuData}.RU()

	return w, nil
}

func (w *World) checkLtHtTradeCodes() {
	if TrvCore.EhexToDigit(w.data[constant.PrTL]) <= TrvCore.EhexToDigit("5") {
		w.tradeCodes = append(w.tradeCodes, constant.TradeCodeLowTech)
	}
	if TrvCore.EhexToDigit(w.data[constant.PrTL]) >= TrvCore.EhexToDigit("C") {
		w.tradeCodes = append(w.tradeCodes, constant.TradeCodeHighTech)
	}
}
