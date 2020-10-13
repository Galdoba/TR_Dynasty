package routine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/name"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
)

//PassengerRoutine -
func PassengerRoutine() {

	printSlow("Searching for Passengers...\n")
	timeLimit := 0
	playerEffect1 := 0
	switch autoMod {
	case false:
		//playerEffect1 = userInputInt("Enter Effect of Carouse(8) or Streetwise(8) check: ")
		input := userInputIntSlice("Enter Effect of Carouse(8) or Streetwise(8) check (and time limit in days after ' ' if nesessary): ")
		if len(input) > 0 {
			playerEffect1 = input[0]
		}
		if len(input) > 1 {
			timeLimit = input[1]
		}
	case true:
		playerEffect1 = autoFlux()
	}
	playerEffect1, time, abort := mutateTestResultsByTime(playerEffect1, dice.Roll("1d6").Sum(), timeLimit)
	if abort {
		fmt.Println("Search aborted after", time, "days...")
	}
	fmt.Println("Search took", time, "days...")
	if longestSearchTime < time {
		longestSearchTime = time
	}
	low, basic, middle, high := availablePassengers(ptValue + playerEffect1 + localBroker.DM())
	passPrices := make(map[int]int)
	passPrices[lowPassenger] = lowPassCost(jumpRoute) - localBroker.CutFrom(lowPassCost(jumpRoute))
	passPrices[basPassenger] = basicPassCost(jumpRoute) - localBroker.CutFrom(basicPassCost(jumpRoute))
	passPrices[midPassenger] = middlePassCost(jumpRoute) - localBroker.CutFrom(middlePassCost(jumpRoute))
	passPrices[highPassenger] = highPassCost(jumpRoute) - localBroker.CutFrom(highPassCost(jumpRoute))

	printSlow("Active passenger requests: " + strconv.Itoa(low+basic+middle+high) + "\n")
	if low > 0 {
		fee := passPrices[lowPassenger]
		printSlow("   Low Passengers: " + strconv.Itoa(low) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	if basic > 0 {
		fee := passPrices[basPassenger]
		printSlow(" Basic Passengers: " + strconv.Itoa(basic) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	if middle > 0 {
		fee := passPrices[midPassenger]
		printSlow("Middle Passengers: " + strconv.Itoa(middle) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	if high > 0 {
		fee := passPrices[highPassenger]
		printSlow("  High Passengers: " + strconv.Itoa(high) + "		Transport fee: " + strconv.Itoa(fee) + "\n")
	}
	fmt.Println("-----------------------------------------------------")
	if userConfirm("Take Passengers?") {
		takePassengers(passPrices)
	}
}

//ABCD
//EFGH
//IJKL
//MNOP
func passengerTrafficValue(sourceWorld, targetWorld wrld.World) int {
	//dm := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrPops))
	uwp, _ := profile.NewUWP(sourceWorld.UWP())
	dm := TrvCore.EhexToDigit(uwp.Pops())
	for _, val := range sourceWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 0
		case constant.TradeCodeAsteroid:
			dm += 1
		case constant.TradeCodeBarren:
			dm += 5
		case constant.TradeCodeDesert:
			dm += -1
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm += 2
		case constant.TradeCodeHighPopulation:
			dm += 0
		case constant.TradeCodeIceCapped:
			dm += 1
		case constant.TradeCodeIndustrial:
			dm += 2
		case constant.TradeCodeLowPopulation:
			dm += 0
		case constant.TradeCodeNonAgricultural:
			dm += 0
		case constant.TradeCodeNonIndustrial:
			dm += 0
		case constant.TradeCodePoor:
			dm += -2
		case constant.TradeCodeRich:
			dm += -1
		case constant.TradeCodeWaterWorld:
			dm += 0
		case constant.TravelCodeAmber:
			dm += 2
		case constant.TravelCodeRed:
			dm += 4
		}
	}
	switch sourceWorld.TravelZone() {
	case constant.TravelCodeAmber:
		dm += 2
	case constant.TravelCodeRed:
		dm += 4
	}
	//dm += TrvCore.EhexToDigit(targetWorld.PlanetaryData("Pops")) // - да хер пойми надо оно или нет (конфликт правил MgT1:MP p.66)
	for _, val := range targetWorld.TradeCodes() {
		switch val {
		default:
		case constant.TradeCodeAgricultural:
			dm += 0
		case constant.TradeCodeAsteroid:
			dm += -1
		case constant.TradeCodeBarren:
			dm += -5
		case constant.TradeCodeDesert:
			dm += -1
		case constant.TradeCodeFluidOceans:
			dm += 0
		case constant.TradeCodeGarden:
			dm += 0
		case constant.TradeCodeHighPopulation:
			dm += 4
		case constant.TradeCodeIceCapped:
			dm += -1
		case constant.TradeCodeIndustrial:
			dm += 1
		case constant.TradeCodeLowPopulation:
			dm += -4
		case constant.TradeCodeNonAgricultural:
			dm += 0
		case constant.TradeCodeNonIndustrial:
			dm += -1
		case constant.TradeCodePoor:
			dm += -1
		case constant.TradeCodeRich:
			dm += 2
		case constant.TradeCodeWaterWorld:
			dm += 0
		}
	}
	switch targetWorld.TravelZone() {
	case constant.TravelCodeAmber:
		dm += -2
	case constant.TravelCodeRed:
		dm += -4
	}

	dm += techDifferenceDM()
	return dm
}

func availablePassengers(r int) (int, int, int, int) {
	return lowPass(r), midPass(r), midPass(r), hghPass(r)
}

func lowPass(ptv int) int {
	if ptv <= 0 {
		return 0
	}
	dm := 0
	d := ""
	switch ptv {
	case 1:
		d = "2d6"
		dm = -6
	case 2:
		d = "2d6"
	case 3:
		d = "2d6"
	case 4:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 5:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 6:
		d = "3d6"
	case 7:
		d = "3d6"
	case 8:
		d = "4d6"
	case 9:
		d = "4d6"
	case 10:
		d = "5d6"
	case 11:
		d = "5d6"
	case 12:
		d = "6d6"
	case 13:
		d = "6d6"
	case 14:
		d = "7d6"
	case 15:
		d = "8d6"
	default:
		if ptv > 15 {
			d = "9d6"
		}
	}
	lp := dp.RollNext(d).DM(dm).Sum()
	if lp < 0 {
		lp = 0
	}
	return lp
}

func midPass(ptv int) int {
	if ptv <= 0 {
		return 0
	}
	dm := 0
	d := ""
	switch ptv {
	case 1:
		return 0
	case 2:
		d = "1d6"
		dm = -2
	case 3:
		d = "1d6"
	case 4:
		d = "2d6"
		dm = dp.RollNext("1d6").Sum()
	case 5:
		d = "2d6"
		dm = dp.RollNext("1d6").Sum()
	case 6:
		d = "2d6"
	case 7:
		d = "3d6"
		dm = dp.RollNext("2d6").Sum()
	case 8:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 9:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 10:
		d = "3d6"
	case 11:
		d = "4d6"
	case 12:
		d = "4d6"
	case 13:
		d = "4d6"
	case 14:
		d = "5d6"
	case 15:
		d = "5d6"
	default:
		if ptv > 15 {
			d = "6d6"
		}
	}
	mp := dp.RollNext(d).DM(dm).Sum()
	if mp < 0 {
		mp = 0
	}
	return mp
}

func hghPass(ptv int) int {
	if ptv <= 0 {
		return 0
	}
	dm := 0
	d := ""
	switch ptv {
	case 1:
		return 0
	case 2:
		return 0
	case 3:
		d = "1d6"
		dm = dp.RollNext("1d6").Sum()
	case 4:
		d = "2d6"
		dm = dp.RollNext("2d6").Sum()
	case 5:
		d = "2d6"
		dm = dp.RollNext("1d6").Sum()
	case 6:
		d = "3d6"
		dm = dp.RollNext("2d6").Sum()
	case 7:
		d = "3d6"
		dm = dp.RollNext("2d6").Sum()
	case 8:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 9:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 10:
		d = "3d6"
		dm = dp.RollNext("1d6").Sum()
	case 11:
		d = "3d6"
	case 12:
		d = "3d6"
	case 13:
		d = "4d6"
	case 14:
		d = "4d6"
	case 15:
		d = "4d6"
	default:
		if ptv > 15 {
			d = "5d6"
		}
	}
	hp := dp.RollNext(d).DM(dm).Sum()
	if hp < 0 {
		hp = 0
	}
	return hp
}

const (
	crewMember     = -2
	highPassenger  = 0
	midPassenger   = 1
	basPassenger   = 2
	lowPassenger   = 3
	guestPassenger = 4
)

func lowPassCost(jumpSeq []int) int {
	totalPrice := 0
	for _, val := range jumpSeq {
		if val < 1 {
			return 0
		}
		totalPrice = totalPrice + (800 + val*200)
	}
	return totalPrice
}

func basicPassCost(jumpSeq []int) int {
	totalPrice := 0
	for _, val := range jumpSeq {
		if val < 1 {
			return 0
		}
		if val == 1 {
			totalPrice = totalPrice + 1500
			continue
		}
		if val == 1 {
			totalPrice = totalPrice + 3000
			continue
		}
		totalPrice = totalPrice + ((val * 2500) - 2500)
	}
	return totalPrice
}

func middlePassCost(jumpSeq []int) int {
	// totalPrice := 0
	// for _, val := range jumpSeq {
	// 	if val < 1 {
	// 		return 0
	// 	}
	// 	if val == 1 {
	// 		totalPrice = totalPrice + 3000
	// 		continue
	// 	}
	// 	if val == 1 {
	// 		totalPrice = totalPrice + 6000
	// 		continue
	// 	}
	// 	totalPrice = totalPrice + ((val * 5000) - 5000)
	// }
	// return totalPrice
	return basicPassCost(jumpSeq) * 2
}

func highPassCost(jumpSeq []int) int {
	return middlePassCost(jumpSeq) * 2
}

// func basicPassCost() int {
// 	if dist < 1 {
// 		return 0
// 	}
// 	if dist == 1 {
// 		return 1500
// 	}
// 	if dist == 2 {
// 		return 3000
// 	}
// 	return (dist - 1) * 2500
// }

// func midddlePassCost() int {
// 	if dist < 1 {
// 		return 0
// 	}
// 	if dist == 1 {
// 		return 3000
// 	}
// 	if dist == 2 {
// 		return 6000
// 	}
// 	return (dist - 1) * 5000
// }

// func highPassCost() int {
// 	if dist < 1 {
// 		return 0
// 	}
// 	if dist == 1 {
// 		return 6000
// 	}
// 	if dist == 2 {
// 		return 12000
// 	}
// 	return (dist - 1) * 10000
// }

/*
Passage and Freight Costs
Parsecs Travelled	High Passage	Middle Passage	Basic Passage	Low Passage	Freight
1					Cr8500			Cr6200			Cr2200			Cr700		Cr1000
2					Cr12000			Cr9000			Cr2900			Cr1300		Cr1600
3					Cr20000			Cr15000			Cr4400			Cr2200		Cr3000
4					Cr41000			Cr31000			Cr8600			Cr4300		Cr7000
5					Cr45000			Cr34000			Cr9400			Cr13000		Cr7700
6					Cr470000		Cr350000		Cr93000			Cr96000		Cr86000


*/

// type passengers struct {
// 	pType             int
// 	pQty              int
// 	cost              int
// 	destiantion       string
// 	notablePassengers int
// }

/*
TranzitHazardDM - Следует использовать в прыжковой программе
*/

// func newPassengers(pType int, pQty int, cost int) passengers {
// 	p := passengers{}
// 	p.pType = pType
// 	p.pQty = pQty
// 	p.cost = cost
// 	p.destiantion = targetWorld.Hex()
// 	if p.pType != lowPassenger {
// 		p.notablePassengers = notablePassengers(p.pQty)
// 	}
// 	return p
// }

// func notablePassengers(i int) int {
// 	n := 0
// 	for i > 0 {
// 		if dice.Roll("1d6").Sum() > 3 {
// 			n++
// 		}
// 		i = i - 6
// 	}
// 	return n
// }

func takePassengers(passPrices map[int]int) {
	clrScrn()
	//err := errors.New("No Value for 'pQty'")
	done := false
	for !done {
		fstr := freeStateRooms()
		ptype, _ := menu("Select Passenger Type:", "High", "Middle", "Basic", "Low", "Guest", "[End Operation]")
		if ptype == 5 {
			return
		}
		err := errors.New("Not valid number")
		num := 0
		for err != nil {
			fmt.Print("Enter number of passengers [0-" + strconv.Itoa(fstr) + "]: ")
			num, err = user.InputInt()
			if num < 0 {
				err = errors.New("Can't have negative value")
			}
			if num > fstr {
				err = errors.New("Not enough Free Staterooms")
			}
			if num == 0 {
				err = errors.New("Zero value entered")
				reportErr(err)
				if userConfirm("End Operation?") {
					return
				}
				continue
			}
			reportErr(err)
		}
		pickUpPassengers(num, ptype, passPrices[ptype])
	}

}

/*
	highPassenger   = 0
	midPassenger    = 1
	basPassenger    = 2
	lowPassenger    = 3
	guestyPassenger = 4

D66 Passenger Transit Hazard DM
11 Refugee – political +4
12 Refugee – economic +3
13 Starting a new life offworld –5
14 Mercenary +0
15 Spy +1
16 Corporate executive +0
21 Out to see the universe –4
22 Tourist. (1–3 Irritating, 4–6 Charming) –4
23 Wide-eyed yokel –5
24 Adventurer +1
25 Explorer +0
26 Claustrophobic +0
31 Expectant Mother +0
32 Wants to stowaway or join the crew +1
33 Possesses something dangerous or illegal +3
34 Troublemaker (1–3 drunkard, 4–5 violent, 6 insane) +0
35 Unusually pretty or handsome +1
36 Engineer (Mechanic and Engineer of 1d6-1 each) +0
41 Ex-scout +1
42 Wanderer –2
43 Thief or other criminal +2
44 Scientist +1
45 Journalist or researcher +1
46 Entertainer (Steward and Perform of 1d6-1 each) +1
51 Gambler (Gambling skill of 1d6-1) +2
52 Rich noble – complains a lot +1
53 Rich noble – eccentric +2
54 Rich noble – raconteur +2
55 Diplomat on a mission +3
56 Agent on a mission +3
61 Patron +1
62 Alien (roll again; ignoring ‘62’ for alien’s further deﬁ nition) +0
63 Bounty Hunter +3
64 On the run +4
65 Wants to be on board the PC’s ship for some reason +1
66 Hijacker or pirate agent +5

*/

type passengerManifest struct {
	entry []passenger
}

type passenger struct {
	id          int
	name        string
	category    string
	quality     string
	thDM        int
	fee         int
	origin      string
	destination string
	stateroom   string
}

func newPassenger() passenger {
	p := passenger{}
	p.id = int(time.Now().UnixNano())
	p.name = name.RandomNew()
	p.origin = sourceWorld.Hex()
	p.destination = targetWorld.Hex()
	return p
}

func (p *passenger) ID() int {
	return p.id
}

func (p *passenger) setID(id int) {
	p.id = id
}

func (p *passenger) IDstr() string {
	return strconv.Itoa(p.id)
}

func (p *passenger) Name() string {
	return p.name
}

func (p *passenger) setName(name string) {
	p.name = name
}

func (p *passenger) Category() string {
	return p.category
}

func (p *passenger) setCategory(category string) {
	p.category = category
}

func (p *passenger) Quality() string {
	return p.quality
}

func (p *passenger) setQuality(quality string) {
	p.quality = quality
}

func (p *passenger) TransitHazardDM() int {
	return p.thDM
}

func (p *passenger) setTransitHazardDM(th int) {
	p.thDM = th
}

func (p *passenger) Fee() int {
	return p.fee
}

func (p *passenger) setFee(fee int) {
	p.fee = fee
}

func (p *passenger) Origin() string {
	return p.origin
}

func (p *passenger) setOrigin(origin string) {
	p.origin = origin
}

func (p *passenger) Destination() string {
	return p.destination
}

func (p *passenger) setDestination(destination string) {
	p.destination = destination
}

func (p *passenger) Stateroom() string {
	return p.stateroom
}

func (p *passenger) setStateroom(sttrm string) {
	p.stateroom = sttrm
}

func loadPassengerManifest() passengerManifest {
	pm := passengerManifest{}
	rawData := getPassengers()
	for i := range rawData {
		lot := newPassenger()
		if lot.LeachData(rawData[i]) != nil {
			continue
		}
		pm.entry = append(pm.entry, lot)
	}
	return pm
}

func getPassengers() []string {
	lines := utils.LinesFromTXT(passengerfile)

	lineNums := utils.InFileContainsN(passengerfile, "ENTRY")
	passengersInfo := []string{}

	for _, i := range lineNums {
		currentLine := lines[i]
		data := strings.Split(currentLine, ":")
		dataParts := strings.Split(data[1], "_")
		if len(dataParts) != 9 {
			for e := 0; e < len(dataParts); e++ {
				fmt.Println(e, dataParts[e])
			}
			panic(errors.New("Data Corrupted: " + data[1]))
		}
		passengersInfo = append(passengersInfo, data[1])
	}
	return passengersInfo
}

func (p *passenger) LeachData(rawData string) error {
	err := errors.New("Null error")
	data := strings.Split(rawData, "_")
	p.id, err = strconv.Atoi(data[0])
	p.name = data[1]
	p.category = data[2]
	p.quality = data[3]
	p.thDM, err = strconv.Atoi(data[4])
	p.fee, err = strconv.Atoi(data[5])
	p.origin = data[6]
	p.destination = data[7]
	p.stateroom = data[8]
	return err
}

func (p *passenger) SeedData() string {
	str := "CARGOENTRY:"
	//ENTRY:ID_NAME_CATEGORY_QUALITY_THDM_FEE_ORIGIN_DESTINATION
	str += p.IDstr() + "_" + p.Name() + "_" + p.Category() + "_" + p.Quality() + "_" + strconv.Itoa(p.TransitHazardDM()) + "_" +
		strconv.Itoa(p.Fee()) + "_" + p.Origin() + "_" + p.Destination() + "_" + p.Stateroom()
	return str

}

func pType2Category(pType int) string {
	switch pType {
	case lowPassenger:
		return "Low"
	case basPassenger:
		return "Basic"
	case midPassenger:
		return "Middle"
	case highPassenger:
		return "High"
	case guestPassenger:
		return "Guest"
	case crewMember:
		return "Crew"
	}
	return ""
}

func category2pType(cat string) int {
	switch cat {
	case "High":
		return highPassenger
	case "Middle":
		return midPassenger
	case "Basic":
		return basPassenger
	case "Low":
		return lowPassenger
	case "Guest":
		return guestPassenger
	case "Crew":
		return crewMember
	}
	return -999
}

func pickUpPassengers(pQty int, pType int, fee int) {
	pm := loadPassengerManifest()
	for i := 0; i < pQty; i++ {
		ps := newPassenger()
		ps.setCategory(pType2Category(pType))
		ps.setFee(fee)
		ps.setStateroom(bestStateroom(ps))
		pm.entry = append(pm.entry, ps)
	}
	savePassengerManifest(pm)
}

func savePassengerManifest(pm passengerManifest) {
	lines := utils.LinesFromTXT(exPath + passengerfile)
entry:
	for i := range pm.entry {
		for l, val := range lines {
			if strings.Contains(val, strconv.Itoa(pm.entry[i].ID())) {
				utils.EditLineInFile(exPath+passengerfile, l, pm.entry[i].SeedData())
				continue entry
			}
		}
		utils.AddLineToFile(exPath+passengerfile, pm.entry[i].SeedData())
	}
}

const (
	stroomLux  = "STATEROOMLUX"
	stroomHigh = "STATEROOMHIGH"
	stroomStd  = "STATEROOMSSTD"
	stroomLB   = "LOWBIRTH"
)

func bestStateroom(p passenger) string {
	switch p.Category() {
	default:
		return stroomStd
	}
}

func getOccupiedStaterooms() int { //Временная мера
	pm := loadPassengerManifest()
	count := len(pm.entry)
	return count
}

func freeStateRooms() int {
	total := getShipData("SHIP_STATEROOMS_STANDARD")
	occ := getOccupiedStaterooms()
	return total - occ
}

func unloadPassengers() {
	pm := loadPassengerManifest()
	totalProfit := 0
	for _, val := range pm.entry {
		if val.Destination() == sourceWorld.Hex() {
			fmt.Print("Unloaded ", val.Name(), " fee: "+strconv.Itoa(val.Fee()), " Cr\n")
			profit := val.Fee()
			totalProfit += profit
			deleteFromPassengerManifest(val.ID())
		}

	}
	fmt.Print("------------------------------\n")
	fmt.Print("Passenger freight fee: ", totalProfit, " Cr\n")
}

func deleteFromPassengerManifest(idValue int) passengerManifest {
	id := strconv.Itoa(idValue)
	n := utils.InFileContains(exPath+passengerfile, id)
	utils.DeleteLineFromFileN(exPath+passengerfile, n)
	return loadPassengerManifest()
}
