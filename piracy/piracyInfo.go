package piracy

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/worldBuilder"
)

const (
	directionUp   = "Heading to planet"
	directionDown = "Heading away from planet"
	directionFlat = "Stationary/Heading other direction"
)

func determineSystem(world *worldBuilder.World) [6]bool {
	types := 0
	var sysType [6]bool
	//TODO: Структурировать/синхронизировать порядок модификаторов системы для пиратства.
	if isBackwaterSystem(world) {
		types++
		sysType[0] = true
	}
	if isHighTrafficSystem(world) {
		types++
		sysType[1] = true
	}
	if isCapitalSystem(world) {
		types++
		sysType[2] = true
	}
	if isDangerousSystem(world) {
		types++
		sysType[3] = true
	}
	if isSecureSystem(world) {
		types++
		sysType[4] = true
	}

	if isNavalBaseSystem(world) {
		sysType[5] = true
	}
	return sysType
}

type PiracyEncounter struct {
	world         *worldBuilder.World
	sysType       [6]bool
	rollResult    string
	heading       string
	distance      string
	encounterType string
	pray          *pray
}

func EncounterSystemType(encounter *PiracyEncounter) string {
	var result string
	world := encounter.world
	result = world.Name() + " is "
	types := 0
	if encounter.sysType[0] {
		result = result + "Backwater "
		types++
	}
	if encounter.sysType[1] {
		result = result + "High traffic "
		types++
	}
	if encounter.sysType[2] {
		result = result + "Capital "
		types++
	}
	if encounter.sysType[3] {
		result = result + "Dangerous "
		types++
	}
	if encounter.sysType[4] {
		result = result + "Secure "
		types++
	}
	if types == 0 {
		result = result + "Average "
	}
	result = result + "System "
	if encounter.sysType[5] {
		result = result + "and have a Naval base"
		types++
	}
	return result
}

func NewEncounterPiracy(world *worldBuilder.World) *PiracyEncounter {
	pEnc := &PiracyEncounter{}
	pEnc.world = world
	pEnc.sysType = determineSystem(world)
	fmt.Println(EncounterSystemType(pEnc))
	pEnc.distance = pEnc.Distance()
	pEnc.rollResult = pEnc.RollResult()
	pEnc.pray = newPray(pEnc.rollResult, pEnc.world)
	return pEnc
}

func applyMods(systType [6]bool, mods [6]int) int {
	dm := 0
	for i := range systType {
		if systType[i] {
			dm = dm + mods[i]
		}
	}
	return dm
}

func (pEnc *PiracyEncounter) Response() {
	mods := [6]int{-1, 1, 2, -1, 1, 2}
	dm := applyMods(pEnc.sysType, mods)
	roll := utils.RollDice("2d6", dm)
	fmt.Println("RESPONCE DM:", dm, "result", roll)
	done := false
	for !done {
		if roll <= 5 {
			fmt.Println("No response; roll again in one hour")
			roll = utils.RollDice("2d6", dm)
			continue
		}
		if utils.InRange(roll, 6, 7) {
			fmt.Println("A vessel launches from the starport to investigate")
		}
		if utils.InRange(roll, 8, 9) {
			fmt.Println("A vessel launches from the starport or from the hundred-diameter jump limit, whichever is closer")
		}
		if roll == 10 {
			fmt.Println("A vessel in orbit responds; response time is", utils.RollDice("d6", 1), "hours")
		}
		if roll == 11 {
			fmt.Println("A vessel in orbit responds; response time is", utils.RollDice("d6"), "hours")
		}
		if roll == 12 {
			fmt.Println("A vessel in orbit responds; response time is", utils.RollDice("d6")*10, "minutes")
		}
		if utils.InRange(roll, 13, 99) {
			fmt.Println("A vessel in orbit responds; response time is", utils.RollDice("d6")*5, "minutes")
		}
		done = true
	}
}

func (pEnc *PiracyEncounter) Info() {
	//fmt.Println("pEnc.world", pEnc.world)
	//fmt.Println("pEnc.sysType", pEnc.sysType)
	//fmt.Println("pEnc.heading", pEnc.heading)
	//fmt.Println("pEnc.distance", pEnc.distance)
	//fmt.Println("pEnc.rollResult", pEnc.rollResult, encounterType(pEnc.rollResult))
	//fmt.Println("pEnc.pray", pEnc.pray)
	// fmt.Println("///////")
	// fmt.Println(randomVessel(encounterType(pEnc.rollResult)))
	// fmt.Println("///////")
}

func (enc *PiracyEncounter) Distance() string {
	if enc.distance != "" {
		return enc.distance
	}
	world := enc.world
	distance := utils.RollDice("3d6") * utils.BoundInt(world.Size(), 1, 99)
	enc.distance = convert.ItoS(distance*1000) + "km"
	return enc.distance
}

func (enc *PiracyEncounter) RollResult() string {
	if enc.rollResult != "" {
		return enc.rollResult
	}
	utils.RandomSeed()
	mods1 := [6]int{-1, 0, 1, 1, 2, 0}
	mods2 := [6]int{0, -1, 0, 0, 0, 2}
	dm1 := applyMods(enc.sysType, mods1)
	dm2 := applyMods(enc.sysType, mods2)
	r1 := utils.RollDice("d6", dm1)
	r2 := utils.RollDice("d6", dm2)
	r1 = utils.BoundInt(r1, 0, 7)
	r2 = utils.BoundInt(r2, 0, 8)
	result := convert.ItoS(r1) + convert.ItoS(r2)
	enc.rollResult = result
	return enc.rollResult
}

func (pray *pray) Heading() string {
	if pray.heading != "" {
		return pray.heading
	}
	switch utils.RollDice("d6") {
	case 1, 2, 3:
		pray.heading = directionUp
	case 4, 5:
		pray.heading = directionDown
	case 6:
		pray.heading = directionFlat
	}
	return pray.heading
}

func isBackwaterSystem(world *worldBuilder.World) bool {
	if !utils.ListContains([]string{"X", "E"}, world.StarPort()) {
		return false
	}
	statusBackwater := false
	wTradeCodes := world.TradeCodes()
	for i := range wTradeCodes {
		if utils.ListContains([]string{"Ba", "Lp", "Lt"}, wTradeCodes[i]) {
			statusBackwater = true
		}
	}
	return statusBackwater
}

func isHighTrafficSystem(world *worldBuilder.World) bool {
	if !utils.ListContains([]string{"A", "B"}, world.StarPort()) {
		return false
	}
	statusHighTraffic := false
	wTradeCodes := world.TradeCodes()
	for i := range wTradeCodes {
		if utils.ListContains([]string{"Ht", "Hp", "In", "Ag", "Ri"}, wTradeCodes[i]) {
			statusHighTraffic = true
		}
	}
	return statusHighTraffic
}

func isCapitalSystem(world *worldBuilder.World) bool {
	statusCapital := false
	wTradeCodes := world.TradeCodes()
	for i := range wTradeCodes {
		if utils.ListContains([]string{"Cp", "Cs", "Cx"}, wTradeCodes[i]) {
			statusCapital = true
		}
	}
	return statusCapital
}

func isDangerousSystem(world *worldBuilder.World) bool {
	statusDangerous := true
	if world.TravelZone() == "Amber Zone" {
		return statusDangerous
	}
	if world.TravelZone() == "Red Zone" {
		return statusDangerous
	}
	if world.Law() < 4 {
		return statusDangerous
	}
	return false
}

func isSecureSystem(world *worldBuilder.World) bool {
	statusSecure := true
	if world.Law() > 6 && world.TechLevel() > 8 {
		return statusSecure
	}
	return false
}

func isNavalBaseSystem(world *worldBuilder.World) bool {
	if utils.ListContains(world.Bases(), "N") {
		return true
	}
	return false
}

func encounterType(encCode string) (encType string) {
	switch encCode {

	case "00":
		encType = "Traveller"
	case "01":
		encType = "Traveller"
	case "02":
		encType = "No encounter"
	case "03":
		encType = "No encounter"
	case "04":
		encType = "Small Freighter"
	case "05":
		encType = "No encounter"
	case "06":
		encType = "No encounter"
	case "07":
		encType = "No encounter"
	case "08":
		encType = "Naval Patrol"
	case "10":
		encType = "Traveller"
	case "11":
		encType = "No encounter"
	case "12":
		encType = "No encounter"
	case "13":
		encType = "No encounter"
	case "14":
		encType = "Small Freighter"
	case "15":
		encType = "No encounter"
	case "16":
		encType = "No encounter"
	case "17":
		encType = "Medium Freighter"
	case "18":
		r := utils.RollDice("d6")
		if r < 4 {
			encType = "No encounter"
		} else {
			encType = "Naval Patrol"
		}
	case "20":
		encType = "Traveller"
	case "21":
		encType = "No encounter"
	case "22":
		encType = "No encounter"
	case "23":
		encType = "Small Freighter"
	case "24":
		encType = "Medium Freighter"
	case "25":
		encType = "No encounter"
	case "26":
		encType = "Unusual Vessel"
	case "27":
		encType = "System Defence Boat"
	case "28":
		r := utils.RollDice("d6")
		if r < 4 {
			encType = "No encounter"
		} else {
			encType = "Naval Patrol"
		}
	case "30":
		encType = "Traveller"
	case "31":
		encType = "Small Freighter"
	case "32":
		encType = "Convoy"
	case "33":
		encType = "Unusual Vessel"
	case "34":
		encType = "Medium Freighter"
	case "35":
		encType = "No encounter"
	case "36":
		encType = "No encounter"
	case "37":
		encType = "Rich Freighter"
	case "38":
		r := utils.RollDice("d6")
		if r < 4 {
			encType = "No encounter"
		} else {
			encType = "Naval Patrol"
		}
	case "40":
		encType = "Small Freighter"
	case "41":
		encType = "Traveller"
	case "42":
		encType = "Convoy"
	case "43":
		encType = "Heavy Freighter"
	case "44":
		encType = "No encounter"
	case "45":
		encType = "No encounter"
	case "46":
		encType = "Liner"
	case "47":
		encType = "System Defence Boat"
	case "48":
		r := utils.RollDice("d6")
		if r < 4 {
			encType = "No encounter"
		} else {
			encType = "Naval Patrol"
		}
	case "50":
		encType = "Traveller"
	case "51":
		encType = "Convoy"
	case "52":
		encType = "Small Freighter"
	case "53":
		encType = "Medium Freighter"
	case "54":
		encType = "No encounter"
	case "55":
		encType = "Heavy Freighter"
	case "56":
		encType = "Liner"
	case "57":
		encType = "System Defence Boat"
	case "58":
		r := utils.RollDice("d6")
		if r < 3 {
			encType = "No encounter"
		} else {
			encType = "Naval Patrol"
		}
	case "60":
		encType = "Traveller"
	case "61":
		encType = "Convoy"
	case "62":
		encType = "Small Freighter"
	case "63":
		encType = "Medium Freighter"
	case "64":
		encType = "Liner"
	case "65":
		encType = "Convoy"
	case "66":
		encType = "Rich Freighter"
	case "67":
		encType = "System Defence Boat"
	case "68":
		encType = "Naval Patrol"
	case "70":
		encType = "Traveller"
	case "71":
		encType = "Small Freighter"
	case "72":
		encType = "Medium Freighter"
	case "73":
		encType = "Convoy"
	case "74":
		encType = "Unusual Vessel"
	case "75":
		encType = "Liner"
	case "76":
		encType = "Rich Freighter"
	case "77":
		encType = "System Defence Boat"
	case "78":
		encType = "Naval Patrol"
	default:
		panic("Impossible code? " + encCode)
	}
	return encType
}

type ship struct {
	tonnage    int
	hp         int
	mDrive     int
	sensors    int
	weapons    []string
	crew       int
	crewBonus  int
	passengers [3]int
	cargoBay   int
}

type staterooms struct {
	lowBerth int
	standard int
	high     int
	luxury   int
}

type vesselData struct {
	class      string
	function   string
	page       string
	tons       int
	thrust     int
	cargo      int
	crew       int
	staterooms staterooms
	prayType   string
}

func randomVessel(category string) vesselData {
	vData := VesselData()
	fmt.Printf("Category %v\n", category)
	if category == "No encounter" {
		fmt.Println("Abort: No Encounter")
		return vesselData{}
	}
	if category == "Convoy" {
		fmt.Println("Abort: Convoy", TrvCore.Roll2D(), "ships")
		return vesselData{}
	}
	var vesselPool []vesselData
	for i := range vData {
		if category == "Traveller" {
			if vData[i].prayType == "Traveller" || vData[i].prayType == "Unusual Vessel" || vData[i].prayType == "Liner" {
				vesselPool = append(vesselPool, vData[i])
			}
		}
		if category == "Rich Freighter" {
			if vData[i].prayType == "Small Freighter" || vData[i].prayType == "Medium Freighter" || vData[i].prayType == "Heavy Freighter" {
				vesselPool = append(vesselPool, vData[i])
			}
		}

		if vData[i].prayType == category {
			vesselPool = append(vesselPool, vData[i])
		}

	}
	r := utils.RollDice("d"+convert.ItoS(len(vesselPool)), -1)
	//fmt.Println("DEBUG:", category, r, "--------", len(vesselPool), vesselPool)
	fmt.Printf("Ship Type: %v-class %v (%v)\n", vesselPool[r].class, vesselPool[r].function, vesselPool[r].page)
	fmt.Printf("Tonnage: %v tonns\n", vesselPool[r].tons)

	return vesselPool[r]
}

//A vessel other than a cargo ship, such as a scout vessel, small military ship, fast courier or even another pirate.
func VesselData() []vesselData {
	var vData []vesselData
	vData = append(vData, vesselData{"Type-A3", "FAST TRADER", "PoD3 2", 200, 4, 36, 5, staterooms{20, 10, 0, 0}, "Small Freighter"})
	vData = append(vData, vesselData{"STAR RAY", "INTERCEPTOR", "PoD3 4", 200, 3, 44, 13, staterooms{8, 6, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"STAR RAY", "INTERCEPTOR", "PoD3 4", 200, 3, 44, 13, staterooms{8, 6, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"HERALD", "FAST MESSENGER", "PoD3 6", 300, 2, 14, 7, staterooms{0, 4, 2, 2}, "Unusual Vessel"})
	vData = append(vData, vesselData{"INDIGO", "PIRATE CARRIER", "PoD3 8", 300, 1, 65, 16, staterooms{8, 10, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"BUCCANEER", "BLOCKADE RUNNER", "PoD3 10", 400, 3, 62, 11, staterooms{0, 8, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"FIERY", "GUNSHIP", "PoD3 12", 400, 6, 6, 41, staterooms{34, 9, 0, 0}, "Liner"})
	vData = append(vData, vesselData{"GHOST OF THE REACH", "HEAVY SCOUT", "PoD3 15", 400, 3, 36, 17, staterooms{0, 10, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"Type RQ", "SUBSIDISED MERCHANT", "PoD3 18", 400, 4, 76, 39, staterooms{16, 14, 0, 0}, "System Defence Boat"})
	vData = append(vData, vesselData{"VULTURE", "SALVAGE HAULER", "PoD3 21", 400, 1, 199, 5, staterooms{0, 4, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"VULTURE", "SALVAGE HAULER", "PoD3 21", 400, 1, 199, 5, staterooms{0, 4, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"WATCHDOG", "FLEET PICKET", "PoD3 23", 500, 3, 16, 12, staterooms{8, 8, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"CORSAIR", "POCKET WARSHIP", "PoD3 26", 600, 3, 279, 16, staterooms{20, 10, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"MAGENTA", "REPAIR SHIP", "PoD3 29", 700, 1, 19, 13, staterooms{10, 16, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"Unknown", "PATROL TENDER", "PoD3 32", 1000, 2, 238, 15, staterooms{10, 9, 0, 0}, "Heavy Freighter"})
	vData = append(vData, vesselData{"QUEEN ELIZABETH", "LINER", "PoD3 36", 1200, 1, 5, 14, staterooms{0, 78, 10, 4}, "Liner"})
	vData = append(vData, vesselData{"ULFHEDNAR", "ESCORT CARRIER", "PoD3 40", 2000, 4, 7, 47, staterooms{0, 25, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"RITCHEY", "ESCORT", "PoD3 44", 8000, 6, 83, 192, staterooms{0, 110, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"GALOOF", "MEGAFREIGHTER", "PoD3 48", 30000, 1, 15248, 165, staterooms{0, 87, 1, 0}, "Heavy Freighter"})
	vData = append(vData, vesselData{"PLANET", "HEAVY CRUISER", "PoD3 53", 75000, 6, 376, 1087, staterooms{225, 600, 3, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"HRAYE", "SCOUT", "PoD3 62", 100, 2, 15, 3, staterooms{0, 4, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"IHATEISHO", "SCOUT", "PoD3 64", 100, 2, 4, 3, staterooms{1, 1, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"KTIYHUI", "COURIER", "PoD3 66", 200, 4, 3, 6, staterooms{0, 6, 0, 1}, "Traveller"})
	vData = append(vData, vesselData{"KTEIROA", "SEEKER", "PoD3 68", 200, 2, 61, 4, staterooms{0, 4, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"KTEIROA", "SEEKER", "PoD3 68", 200, 2, 61, 4, staterooms{0, 4, 0, 0}, "Small Freighter"})
	vData = append(vData, vesselData{"IYELIY", "MESSENGER", "PoD3 70", 200, 1, 8, 4, staterooms{0, 2, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"AOA'IW", "LIGHT TRADER", "PoD3 72", 300, 1, 86, 6, staterooms{12, 8, 0, 0}, "Small Freighter"})
	vData = append(vData, vesselData{"EAKHAU", "TRADER", "PoD3 74", 400, 1, 173, 4, staterooms{16, 13, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"HKIYRERAO", "RESEARCHER", "PoD3 76", 400, 1, 28, 6, staterooms{10, 15, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"KHTUKHAO", "CLAN TRANSPORT", "PoD3 78", 600, 2, 149, 5, staterooms{30, 25, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"KHTUKHAO", "CLAN TRANSPORT", "PoD3 78", 600, 2, 149, 5, staterooms{30, 25, 0, 0}, "Liner"})
	vData = append(vData, vesselData{"OWATARL", "TENDER", "PoD3 80", 600, 1, 211, 10, staterooms{0, 10, 0, 0}, "Liner"})
	vData = append(vData, vesselData{"OWATARL", "TENDER", "PoD3 80", 600, 1, 211, 10, staterooms{0, 10, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"EKAWSIYKUA", "ESCORT", "PoD3 82", 800, 4, 19, 40, staterooms{0, 20, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"AOSITAOH", "CRUISER", "PoD3 84", 1000, 4, 13, 103, staterooms{0, 54, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"HKISYELEAA", "SLAVER", "PoD3 87", 1000, 2, 114, 13, staterooms{750, 12, 0, 0}, "Liner"})
	vData = append(vData, vesselData{"HALAHEIKE", "POCKET WARSHIP", "PoD3 90", 1200, 3, 65, 40, staterooms{0, 30, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"SAKHAI", "ASSAULT CARRIER", "PoD3 93", 2000, 3, 6, 211, staterooms{360, 17, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"LANCER", "CORVETTE", "ReachAd4 55", 200, 4, 14, 7, staterooms{0, 10, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"X-BOAT", "EXPRESS BOAT", "HG 108", 100, 0, 26, 0, staterooms{0, 1, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"Type-S", "SCOUT/COURIER", "HG 110", 100, 2, 12, 3, staterooms{0, 4, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"Type-S", "SCOUT/COURIER", "HG 110", 100, 2, 12, 3, staterooms{0, 4, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"Type-J", "SEEKER MINING SHIP", "HG 112", 100, 2, 26, 3, staterooms{0, 2, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"SERPENT", "SCOUT", "HG 114", 200, 2, 6, 3, staterooms{0, 4, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"EMPRESS MARAVA", "FAR TRADER", "HG 116", 200, 1, 57, 6, staterooms{0, 10, 0, 0}, "Small Freighter"})
	vData = append(vData, vesselData{"Type-A2", "FAR TRADER", "HG 118", 200, 1, 64, 5, staterooms{6, 10, 0, 0}, "Small Freighter"})
	vData = append(vData, vesselData{"Type-A", "FREE TRADER", "HG 120", 200, 1, 82, 5, staterooms{20, 10, 0, 0}, "Small Freighter"})
	vData = append(vData, vesselData{"Unknown", "SAFARI SHIP", "HG 122", 200, 1, 6, 5, staterooms{0, 11, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"Unknown", "SYSTEM DEFENCE BOAT", "HG 124", 200, 9, 23, 15, staterooms{0, 15, 0, 0}, "System Defence Boat"})
	vData = append(vData, vesselData{"Type-Y", "YACHT", "HG 126", 200, 1, 11, 5, staterooms{0, 12, 0, 1}, "Unusual Vessel"})
	vData = append(vData, vesselData{"ROCK", "ASTEROID SHIP", "HG 128", 300, 1, 26, 6, staterooms{0, 5, 0, 0}, "Traveller"})
	vData = append(vData, vesselData{"GAZELLE", "CLOSE ESCORT", "HG 130", 400, 6, 33, 21, staterooms{0, 11, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"Unknown", "FLEET COURIER", "HG 132", 400, 2, 9, 16, staterooms{0, 10, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"Type-L", "LABORATORY SHIP", "HG 134", 400, 2, 3, 4, staterooms{0, 20, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"Type-T", "PATROL CORVETTE", "HG 136", 400, 4, 46, 17, staterooms{4, 12, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"Type-R", "SUBSIDISED MERCHANT", "HG 138", 400, 1, 199, 5, staterooms{9, 19, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"DONOSEV", "SURVEY SCOUT", "HG 140", 400, 2, 15, 5, staterooms{0, 5, 0, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"Unknown", "SYSTEM DEFENCE BOAT", "HG 142", 400, 6, 100, 13, staterooms{0, 13, 0, 0}, "System Defence Boat"})
	vData = append(vData, vesselData{"Unknown", "ANNIC NOVA", "HG 144", 520, 0, 295, 13, staterooms{0, 11, 0, 0}, "Medium Freighter"})
	vData = append(vData, vesselData{"Type-M", "SUBSIDISED LINER", "HG 146", 600, 1, 119, 6, staterooms{20, 30, 0, 0}, "Liner"})
	vData = append(vData, vesselData{"Type-C", "MERCENARY CRUISER", "HG 148", 800, 3, 72, 6, staterooms{0, 25, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"CHRYSANTHEMUM", "DESTROYER ESCORT", "HG 150", 1000, 6, 31, 41, staterooms{0, 23, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"FER-DE-LANCE", "DESTROYER ESCORT", "HG 153", 1000, 6, 51, 39, staterooms{0, 23, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"Unknown", "EXPRESS BOAT TENDER", "156", 1000, 1, 25, 11, staterooms{0, 10, 1, 0}, "Unusual Vessel"})
	vData = append(vData, vesselData{"KINUNIR", "COLONIAL CRUISER", "HG 159", 1250, 4, 37, 87, staterooms{31, 30, 0, 0}, "Liner"})
	vData = append(vData, vesselData{"LEVIATHAN", "MERCHANT CRUISER", "HG 161", 1800, 4, 188, 34, staterooms{0, 31, 1, 1}, "Heavy Freighter"})
	vData = append(vData, vesselData{"MIDU AGASHAM", "DESTROYER", "HG 165", 3000, 6, 65, 106, staterooms{0, 106, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"P.F. SLOAN", "FLEET ESCORT", "HG 169", 5000, 6, 239, 114, staterooms{0, 114, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"SKIMKISH", "LIGHT CARRRIER", "HG 173", 29000, 2, 2002, 962, staterooms{0, 490, 1, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"GIONETTI", "LIGHT CRUISER", "HG 176", 30000, 5, 477, 305, staterooms{0, 167, 2, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"GHALALK", "ARMOURED CRUISER", "HG 180", 50000, 6, 463, 749, staterooms{180, 407, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"EMPRESS TROYHUNE", "PLANETOID MONITOR", "HG 184", 50000, 6, 5403, 1238, staterooms{0, 631, 0, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"ARAKOINE", "STRIKE CRUISER", "HG 188", 50000, 4, 173, 1107, staterooms{0, 598, 1, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"AZHANTI", "FRONTIER CRUISER", "HG 192", 60000, 2, 316, 775, staterooms{0, 401, 1, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"ATLANTIC", "HEAVY CRUISER", "HG 196", 75000, 5, 1174, 1351, staterooms{0, 688, 1, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"WIND", "STRIKE CARRIER", "HG 200", 75000, 6, 446, 1424, staterooms{0, 726, 1, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"ANTIAMA", "FLEET CARRIER", "HG 204", 100000, 2, 5075, 2798, staterooms{0, 1424, 1, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"KOKIRRAK", "DREADNOUGHT", "HG 208", 200000, 6, 6771, 4330, staterooms{0, 2214, 3, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"PLANKWELL", "DREADNOUGHT", "HG 212", 200000, 5, 1625, 3887, staterooms{0, 1984, 2, 0}, "Naval Patrol"})
	vData = append(vData, vesselData{"TIGRESS", "DREADNOUGHT", "HG 216", 500000, 6, 3787, 8715, staterooms{0, 4619, 3, 0}, "Naval Patrol"})
	//vData = append(vData, vesselData{"", "", "", 0, 0, 0, 0, staterooms{0, 0, 0, 0}, ""})

	return vData
}

type pray struct {
	vessel       vesselData
	heading      string
	distance     string
	quirk        string
	complication string
	morale       int
	passengers   [3]int
	routeTarget  string
	response     string
	cargo        *Trade.TradeGoodR
}

func newPray(encCode string, world *worldBuilder.World) *pray {
	pray := &pray{}
	encType := encounterType(encCode)
	pray.morale = utils.RollDice("d6", 6)
	if encType == "Naval Patrol" && encType == "System Defence Boat" {
		pray.morale = utils.RollDice("d6", 8)
	}
	if encType == "Small Freighter" || encType == "Medium Freighter" || encType == "Heavy Freighter" {
		pray.morale = utils.RollDice("d6", 3)
	}
	pray.vessel = randomVessel(encType)
	if (pray.vessel == vesselData{}) {
		return pray
	}
	pray.heading = pray.Heading()
	pray.distance = convert.ItoS(world.Size()*utils.RollDice("3d6")*1000) + " km"
	pray.routeTarget = "Imperium"
	if utils.RandomBool() {
		pray.routeTarget = "Hierate"
	}
	pray.complication = prayComplication()
	pray.quirk = prayQuirk()
	fmt.Print("Encounter time Limit: ")
	dist := float64(world.JumpShadowRadius())
	thrust := float64(pray.vessel.thrust)
	if thrust < 0.2 {
		thrust = 0.1
	}
	if pray.heading == directionDown {
		fmt.Println(TrvCore.TravelTime(dist, thrust), "hours")
	} else {
		fmt.Println(TrvCore.TravelTime(dist*2, thrust)/2, "hours")
	}
	fmt.Println("Thrust:", thrust)
	fmt.Println("Distance:", pray.distance)
	fmt.Println(pray.heading)

	fmt.Println("Encounter Quirk =", pray.quirk)
	fmt.Println("Encounter Complication =", pray.complication)
	ans, _ := utils.TakeOptions("Attack?", "Yes", "No")
	if ans == 1 {
		fmt.Println("Target MORALE =", pray.morale)
		fmt.Println("Target Cargo =", ((pray.vessel.cargo * 10 * utils.RollDice("d6", 4)) / 100), "of", pray.vessel.cargo, "filled")
		fmt.Println("Target Crew =", pray.vessel.crew)
		fmt.Println("Target Free StateRooms =", pray.vessel.staterooms)

	}
	// fmt.Println(TrvCore.TravelTime(dist, trust), "hours")
	// fmt.Println("Getaway:")
	// fmt.Println(TrvCore.TravelTime(dist*2, trust)/2, "hours")
	// fmt.Println("Thrust:", trust)

	lowPass, stPass, hiPass, luxPass := Trade.SeekPassengers(world, 0)

	fmt.Println(lowPass, stPass, hiPass, luxPass)

	return pray
}

// func (pray *pray) ComparePass(loP int, stP int, hiP int, luP int) {
// 	cr := pray.vessel.crew
// 	stRoomFree := pray.vessel.staterooms
// 	freeSt := stRoomFree.standard
// 	for freeSt := stRoomFree.standard; freeSt > 0; freeSt-- {
// 		if cr == 0 {
// 			break
// 		}
// 		cr--
// 	}

// }

func prayComplication() string {
	d66 := TrvCore.RollD66()
	var complication string
	switch d66 {
	case "11":
		complication = "Solar Flares: The system’s primary sun spits out huge flares and high levels of radiation. All ships take 3D x 100 rads per hour."
	case "12":
		complication = "Debris Field: The encounter takes place in a debris field. Pilot checks are needed to avoid floating obstacles; on the bright side, there may be some salvage here."
	case "13":
		complication = "Ice Field: The planet’s surrounded by a ring of ice particles, and the quarry takes refuge there. Direct-fire weapons are limited to Short range."
	case "14":
		complication = "Comms Jamming: Something in the system blocks communications. The victim can’t call for help."
	case "15":
		complication = "Behind The Moon: There’s a nearby moon. What’s lurking there? Another pirate? An interceptor? An Aslan spy?"
	case "16":
		complication = "Incoming Escort: The merchant has an escort, but they haven’t jumped in yet. They’ll be here any minute. "
	case "21":
		complication = "Rival Pirate: There’s another pirate after the same prize "
	case "22":
		complication = "Slow Leak: The Travellers’ fuel tank has a slow leak; they’re losing 1D tons of fuel per round. "
	case "23":
		complication = "Out of Control: The quarry loses control of its attitude thrusters and starts spinning wildly. It’s now easy to catch but very hard to dock with."
	case "24":
		complication = "Cargo Spilled: In a panic, the merchant jettisons most of its cargo, sending an expanding flock of canisters into space."
	case "25":
		complication = "Collision Warning! Both ships nearly collide with a small asteroid or other piece of space debris. "
	case "26":
		complication = "Misjump: The first ship to jump out misjumps when they flee "
	case "31":
		complication = "Observer: There’s another ship nearby. They steer clear of the dogfight, but they’re watching... "
	case "32":
		complication = "Bad Jump: This was a bad jump – the pirates have arrived well outside the travelled parts of the system."
	case "33":
		complication = "Unfortunate Timing: Another ship jumps right into the middle of the battle."
	case "34":
		complication = "Crew Dissent: One of the crew on board the Travellers’ ship is having problems that affect the battle. Perhaps they object to this particular target, are drunk, or are deliberately sabotaging the attack."
	case "35":
		complication = "System Failure: A key system fails on board the pirate ship. Roll for a random critical hit with a Severity of D3."
	case "36":
		complication = "Escape Pods: The merchant’s crew flee their ship in escape pods and small craft. They could be carrying treasure on board those pods – but the pirates have time to only chase down one of them..."
	case "41":
		complication = "Rapid Reaction: The security forces here respond very quickly – apply DM+4 to the response time roll. "
	case "42":
		complication = "Corrupt Cops: The security forces can be bribed to ignore the attack."
	case "43":
		complication = "Nearby Asteroid: There’s an asteroid close to the battle; the merchant can fly to the refuge and hide behind it. The asteroid might even be inhabited."
	case "44":
		complication = "Sensor Jamming: Conditions in the system block sensors."
	case "45":
		complication = "Imperial Patrol: There’s an Imperial or Aslan patrol in the system, hunting for pirates. They’re far enough away that the Travellers might be able to complete the attack before the first fighters arrive..."
	case "46":
		complication = "Distress Call: The Travellers detect a distress call from a stricken ship. Do they call off their attack?"
	case "51":
		complication = "High Guard: There’s an unexpected ship refuelling at the system’s gas giant (or at another source of hydrogen, like a lake). Why are they avoiding the starport?"
	case "52":
		complication = "Spy in the System: A spy in the system contacts the Travellers by radio, offering them useful information about traffic."
	case "53":
		complication = "Screamer: The merchant ship frantically warns everyone who’ll listen about the pirates – not just in this system, but in every other system the merchant visits"
	case "54":
		complication = "Incoming!: The starport below launches groundto-space missiles. The first missile hits in 1D+10 rounds..."
	case "55":
		complication = "Tricky Calculation: The complex arrangement of moons and planets in this system make jump calculations harder. Apply DM-4 to any Astrogation checks."
	case "56":
		complication = "Pull Up!: The merchant doesn’t slow down as it approaches the planet – instead, they plan to use aerobraking to slow their dissent."
	case "61":
		complication = "The Black Signal: The pirates pick up the fabled ‘black signal’ on the ship; a pattern of radiation burned into the hull, denoting that this ship is an enemy of the pirates of Theev."
	case "62":
		complication = "Familiar Ship: The Travellers have encountered this merchant ship before..."
	case "63":
		complication = "Aslan Raiders: Several Aslan raiders led by an ambitious ihatei warlord arrive in the system."
	case "64":
		complication = "Under The Shield of the Sunburst: An Imperial patrol jumps in; they’re not pirate hunting, they’re here to enforce the Third Imperium’s will on the planetary government."
	case "65":
		complication = "Didn’t Expect To Find You Here: A Contact (or Ally, or Enemy) of a Traveller is on board the merchant."
	case "66":
		complication = "Anomaly: The Travellers run into something unusual, like a wrecked ship or a spatial anomaly.	"
	}
	return complication

}

func prayQuirk() string {
	var quirk string
	code := TrvCore.RollD66()
	switch code {
	case "11":
		quirk = "Coward: Surrenders easily. Reduce starting MOR by 1D."
	case "12":
		quirk = "Deceitful: Pretends to surrender in order to lure the pirates into docking, then fights back at Short range."
	case "13":
		quirk = "Smuggler: The really valuable cargo is hidden in a secret compartment"
	case "14":
		quirk = "Eccentric: The captain is insane, drunk or otherwise eccentric."
	case "15":
		quirk = "No Surrender: The crew will not surrender under any circumstances. Ignore MOR."
	case "16":
		quirk = "Duel of Honour: The captain challenges one of the pirates to a rapier duel in vacc suits on the exterior hull of his ship."
	case "21":
		quirk = "Noble: There’s a noble on board. If ransomed, she’s worth considerably more than normal."
	case "22":
		quirk = "Alien: There’s an exotic alien like a Hiver on board. "
	case "23":
		quirk = "Family: The captain’s family travel on board the ship."
	case "24":
		quirk = "Diplomat: There is an Imperial or Aslan diplomat on board, carrying a secret message"
	case "25":
		quirk = "Stowaway: Someone’s hidden inside a cargo container that the pirates just stole."
	case "26":
		quirk = "Prisoner: There’s a criminal – perhaps a captured pirate – in the ship’s brig."
	case "31":
		quirk = "Plague Ship: The crew are infected with a potentially fatal disease."
	case "32":
		quirk = "Dying Ship: The ship misjumped and is running low on food, oxygen or fuel."
	case "33":
		quirk = "Damaged Ship: The ship has sustained 1D critical hits, each of Severity D3 already."
	case "34":
		quirk = "Treasure Map: While looting the ship, the Travellers find a map pointing to a hidden supply cache, mineral deposit or other valuable treasure."
	case "35":
		quirk = "Important Document: The ship’s safe contains the deeds to a property, a letter of marque, a corporate contract or some other valuable document."
	case "36":
		quirk = "Message Pod: The ship carries a 5-dton data drum containing mail. Decoding this data may reveal useful information."
	case "41":
		quirk = "Heavily Armed: The merchant ship is ready for a fight. Any hardpoints are equipped with turrets filled with lasers or missile racks."
	case "42":
		quirk = "Berserker: One of the merchant crew is a trained marine equipped with battle dress or boarding vacc suit, and a heavy weapon."
	case "43":
		quirk = "Self Destruct: The captain would rather die than lose his ship. Unless the pirates can stop him, he’ll scuttle his ship rather than lose his cargo."
	case "44":
		quirk = "Mission of Mercy: The ship is carrying vitally needed supplies, like medicine or food, to a troubled colony."
	case "45":
		quirk = "Die Hard: One of the merchant’s crew hides when the ship is boarded, and sneaks onto the Travellers’ ship to sabotage them."
	case "46":
		quirk = "Psionic Defender: One of the crew of the merchant ship is a psion."
	case "51":
		quirk = "Unlikely Cargo: The merchant ship is carrying an unexpected cargo – what are they doing out here?"
	case "52":
		quirk = "Perishable Cargo: The merchant’s cargo is valuable, but only if sold within the month."
	case "53":
		quirk = "Dangerous Cargo: The merchant’s cargo is dangerous to have on board."
	case "54":
		quirk = "Living Cargo: The cargo is alive – animals, insects, or even slaves."
	case "55":
		quirk = "Hot Cargo: The cargo was stolen – and the real owner wants it back."
	case "56":
		quirk = "Alien Cargo: The merchant is carrying something from a very distant part of space, or even an Ancient relic."
	case "61":
		quirk = "Traitor: One of the merchant’s crew is willing to betray his shipmates for a large payoff."
	case "62":
		quirk = "Infestation: There’s something alive on board ship. "
	case "63":
		quirk = "Ghost Ship: This ship has been drifting dead for centuries. The Travellers were attacked by automated weapons."
	case "64":
		quirk = "Strange Curio: There’s a relic or other strange item in the captain’s cabin."
	case "65":
		quirk = "It’s a Trap: This ‘merchant’ is actually a disguised q-ship or pirate hunter."
	case "66":
		quirk = "Drinaxian on Board: One of the characters from The Trojan Reach page 24 is on board... what are they doing here?"
	}
	return quirk
}

/*
Planet
Size 				Thrust 1 Thrust 2 Thrust 3 Thrust 4 Thrust 5 Thrust 6
					Up Down Up Down Up Down Up Down Up Down Up Down
0 (80,000 km) 		1hr 1.5hrs 45mins 1.2hrs 40mins 1hr 33mins 45mins 30mins 42mins 27mins 38mins
1 (160,000 km) 		1.5hrs 2 hrs 1.2hrs 1.5hrs 1hr 1.2hrs 45mins 1.1hrs 42mins 1hr 38mins 0.9hrs
2 (320,000 km) 		2hrs 3 hrs 1.5hrs 2.25hrs 1.2hrs 2hrs 1.1hrs 1.5hrs 1hr 1.5hrs 0.9hrs 1.2hrs
3 (480,000 km) 		3hrs 4 hrs 2hrs 2.75hrs 1.5hrs 2.2hrs 1.3hrs 2hrs 1.2hrs 1.75hrs 1.1hrs 1.5hrs
4 (640,000 km) 		3 hrs 4.3 hrs 2.25hrs 3hrs 1.9hrs 2.5hrs 1.5hrs 2.25hrs 1.5hrs 2hrs 1.2hrs 1.9hrs
5 (800,000 km) 		3.5hrs 5 hrs 2.5hrs 3.5hrs 2hrs 3hrs 1.75hrs 2.5hrs 1.6hrs 2.2hrs 1.5hrs 2hrs
6 (960,000 km) 		4hrs 5.5hrs 2.75hrs 4hrs 2.25hrs 3.2hrs 2hrs 2.75hrs 1.75hrs 2.5hrs 1.5hrs 2.25hrs
7 (1,120,000 km)	 4.2hrs 6hrs 3hrs 4.2hrs 2.4hrs 3.5hrs 2.1hrs 3hrs 1.9hrs 2.6hrs 1.75hrs 2.5hrs
8 (1,280,000 km) 4.5hrs 6.3hrs 3.2hrs 4.5hrs 2.6hrs 3.7hrs 2.25hrs 3.1hrs 2hrs 2.9hrs 1.9hrs 2.4hrs
9 (1,440,000 km) 4.75hrs 6.75hrs 3.3hrs 4.75hrs 2.75hrs 3.9hrs 2.3hrs 3.3hrs 2.1hrs 3hrs 2hrs 2.75hrs
A (1,600,000 km) 5hrs 7hrs 3.5hrs 5hrs 3hrs 4.1hrs 2.5hrs 3.5hrs 2.25hrs 3.1hrs 2hrs 3hrs


+ ++ +++ ++++ +++++ +++++ ++++ +++ ++ +

*/
