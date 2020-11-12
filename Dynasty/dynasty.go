package dynasty

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/DateManager"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

const (
	Clv                 = "Cleverness"
	Grd                 = "Greed"
	Lty                 = "Loyalty"
	Mil                 = "Militarism"
	Pop                 = "Popularity"
	Sch                 = "Scheming"
	Tcy                 = "Tenacity"
	Tra                 = "Tradition"
	Culture             = "Culture"
	FiscalDfnce         = "Fiscal Defence"
	Fleet               = "Fleet"
	Technology          = "Technology"
	TerritorialDfnce    = "Territorial Defence"
	Acquisition         = "Acquisition"
	Bureaucracy         = "Bureaucracy"
	Conquest            = "Conquest"
	Economics           = "Economics"
	Entertain           = "Entertain"
	Expression          = "Expression"
	Hostility           = "Hostility"
	Illicit             = "Illicit"
	Intel               = "Intel"
	Maintenance         = "Maintenance"
	Politics            = "Politics"
	Posturing           = "Posturing"
	Propaganda          = "Propaganda"
	PublicRelations     = "Public Relations"
	Recruit             = "Recruit"
	Research            = "Research"
	Sabotage            = "Sabotage"
	Security            = "Security"
	Tactical            = "Tactical"
	Tutelage            = "Tutelage"
	Morale              = "Morale"
	Populace            = "Populace"
	Wealth              = "Wealth"
	ColonySettlement    = "Colony/Settlement"
	ConflictZone        = "Conflict Zone"
	Megalopolis         = "Megalopolis"
	MilitaryCompound    = "Military Compound"
	NobleEstate         = "Noble Estate"
	StarshipFlotilla    = "Starship/Flotilla"
	TempleHolyLand      = "Temple/Holy Land"
	UnchartedWilderness = "Uncharted Wilderness"
	UnderworldSlum      = "Underworld Slum"
	UrbanOffices        = "Urban Offices"
	Conglomerate        = "Conglomerate"
	MediaEmpire         = "Media Empire"
	MerchantMarket      = "Merchant Market"
	MilitaryCharter     = "Military Charter"
	NobleLine           = "Noble Line"
	ReligiousFaith      = "Religious Faith"
	Syndicate           = "Syndicate"
	BoardofDirectors    = "Board of Directors"
	CommandStaff        = "Command Staff"
	HeroicLeaders       = "Heroic Leaders"
	MatriarchPatriarch  = "Matriarch/Patriarch"
	Overlord            = "Overlord"
	Theocrat            = "Theocrat"
)

var currentDay int

func Test() {
	//code.ConstructStandardMethods()
	//validDyn := false
	dynMap := make(map[int]*Dynasty)
	d1 := NewDynasty("")
	//d1.name = "Dynasty 1"
	dynMap[0] = &d1
	d2 := NewDynasty("")
	//d2.name = "Dynasty 2"
	dynMap[1] = &d2
	// for len(dyns) < 2 {
	// 	d := Dynasty{}
	// 	for !validDyn {
	// 		d = NewDynasty("")
	// 		validDyn = Survived(d)
	// 	}
	// 	fmt.Print(d.Info())
	// 	dyns = append(dyns, d)
	// }

	currentDay = 402000
	for d := 0; d < 1000; d++ {
		currentDay++

		for i := range dynMap {
			if dynMap[i].nextActionDay <= currentDay {
				val := dynMap[i]
				val.DecrlareAction(currentDay)
			}

		}

	}

	// for i := curentDay; i < 402500; i++ {
	// 	for d := range dyns {
	// 		if dyns[d].TrackEvent(i) {
	// 			d2 := 0
	// 			if d == 1 {
	// 				d2 = 0
	// 			}
	// 			InitiateAction(&dyns[d], &dyns[d2], "Claiming neutral territory or resources", i)
	// 		}
	// 	}
	// }
}

func (d *Dynasty) DecrlareAction(curentDay int) {
	dp := dice.New(int64(curentDay))
	r := dp.RollFromList([]string{"Action 1", "Action 2", "Action 3", "Action 4"})
	fmt.Println("on day ", DateManager.FormatToDate(curentDay))
	fmt.Println(r + " declared by " + d.name)
	timeFactor := 1
	switch r {
	case "Action 1":
		timeFactor = dice.Roll("1d6").Sum() * (28 + dice.Flux())
	case "Action 2":
		timeFactor = dice.Roll("2d6").Sum() * (28 + dice.Flux())
	case "Action 3":
		timeFactor = dice.Roll("3d6").Sum() * (28 + dice.Flux())
	case "Action 4":
		timeFactor = dice.Roll("4d6").Sum() * (28 + dice.Flux())
	}
	fmt.Println("It took", timeFactor, "days")
	d.nextActionDay = curentDay + timeFactor
}

type Dynasty struct {
	name     string
	dicepool dice.Dicepool
	//characteristics map[string]int
	//traits          map[string]int
	//aptitudes       map[string]int
	//values          map[string]int
	stat           map[string]int
	boonsHinders   []string
	archetype      string
	powerBase      string
	fgBonus        string
	managment      string
	birthDate      string
	nextEventDay   int
	nextActionDay  int
	historicEvents int
	eventMap       map[int]string
	goals          []goal
	story          string
}

func NewDynasty(name string) Dynasty {
	d := Dynasty{}
	d.eventMap = make(map[int]string)
	seed := utils.SeedFromString(name)

	if name == "" {
		seed = time.Now().UnixNano()
		name = "[RANDOM NAME] seed=" + strconv.Itoa(int(seed))
	}
	d.name = name

	d.dicepool = *dice.New(seed)
	//d.characteristics = make(map[string]int)
	//d.traits = make(map[string]int)
	//d.aptitudes = make(map[string]int)
	//d.values = make(map[string]int)
	d.stat = make(map[string]int)
	//1
	for _, val := range listCharacteristics() {
		d.stat[val] = d.dicepool.RollNext("2d6").Sum()
	}
	//2
	d.choosePowerBase("")
	d.gainPowerBaseBonuses()
	//3
	vArch := validArchetypes(d)
	for len(vArch) == 0 {
		d = NewDynasty("")
		vArch = validArchetypes(d)
	}
	arch := vArch[d.dicepool.RollNext("1d"+lenStr(vArch)).DM(-1).Sum()]
	d.chooseArchetype(arch)
	d.determineBaseTraitsAndAptitudes()
	d.gainFirstGenerationBonuses()
	d.chooseBoons()
	d.chooseManagementAsset()
	d.finalizeFirstGeneration()

	return d
}

func (d *Dynasty) finalizeFirstGeneration() {
	d.fgStep1() //train characteristics
	d.fgStep2() //train management
	d.fgStep3() //train aptitudes
	d.fgStep4() //Final Values Adjustments
}

func (d *Dynasty) fgStep1() {
	characteristicsPractice := []string{}
	for len(characteristicsPractice) < 3 {
		characteristicsPractice = append(characteristicsPractice, d.dicepool.RollFromList(listCharacteristics()))
	}
	for _, val := range characteristicsPractice {
		r := d.dicepool.RollNext("2d6").Sum()
		switch r {
		case 2:
			d.stat[val]--
		case 12:
			d.stat[val]++
		default:
			if DM(d.stat[val])+r > d.stat[val] {
				d.stat[val]++
			}
		}
	}
}

func (d *Dynasty) fgStep2() {
	max := DM(d.stat[Clv]) + 1
	for i := 0; i < max; i++ {
		moveFrom := d.dicepool.RollFromList(listTraits())
		moveTo := d.dicepool.RollFromList(listTraits())
		if d.stat[moveFrom] < 2 {
			continue
		}
		d.stat[moveFrom]--
		d.stat[moveTo]++
	}
}

func (d *Dynasty) fgStep3() {
	aptPractice := []string{}
	for len(aptPractice) < 5 {
		aptPractice = utils.AppendUniqueStr(aptPractice, d.dicepool.RollFromList(listAptitudes()))
	}
	for _, val := range aptPractice {
		r := d.dicepool.RollNext("2d6").DM(-8).Sum()
		e := -2
		if apt, ok := d.stat[val]; ok {
			e = apt
		}
		r = r + e
		if r > d.stat[val] {
			d.raiseApttitude(val)
		}
	}
}

func (d *Dynasty) fgStep4() {
	d.stat[Morale] = d.stat[Morale] + DM(d.stat[Pop]) + d.stat[Culture]
	d.stat[Populace] = d.stat[Populace] + DM(d.stat[Tra]) + d.stat[Technology]
	d.stat[Wealth] = d.stat[Wealth] + DM(d.stat[Clv]) + d.stat[FiscalDfnce]
	for _, val := range listValues() {
		if d.stat[val] < 1 {
			continue
		}
		d.stat[val]--
		d.stat[d.dicepool.RollFromList(listValues())]++
	}
}

func (d *Dynasty) chooseManagementAsset() {
	selected := -1
	for selected == -1 {
		r := d.dicepool.RollNext("2d6").Sum()
		switch r {
		case 3, 4:
			selected = 0
		case 5, 6:
			selected = 1
		case 7, 8, 9:
			selected = 2
		case 10, 11:
			selected = 3
		}
	}
	managementMap := make(map[string][]string)
	managementMap[Conglomerate] = []string{HeroicLeaders, Overlord, BoardofDirectors, CommandStaff}
	managementMap[MediaEmpire] = []string{Overlord, MatriarchPatriarch, BoardofDirectors, Theocrat}
	managementMap[MerchantMarket] = []string{HeroicLeaders, MatriarchPatriarch, BoardofDirectors, CommandStaff}
	managementMap[MilitaryCharter] = []string{Theocrat, HeroicLeaders, CommandStaff, Overlord}
	managementMap[NobleLine] = []string{HeroicLeaders, Overlord, MatriarchPatriarch, Theocrat}
	managementMap[ReligiousFaith] = []string{CommandStaff, HeroicLeaders, Theocrat, Overlord}
	managementMap[Syndicate] = []string{CommandStaff, BoardofDirectors, Overlord, MatriarchPatriarch}
	d.managment = managementMap[d.archetype][selected]
}

func DM(i int) int {
	switch i {
	default:
		if i < 1 {
			return -3
		}
		return 5
	case 1, 2:
		return -2
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	case 9, 10, 11:
		return 1
	case 12, 13, 14:
		return 2
	case 15, 16, 17:
		return 3
	case 18, 19, 20:
		return 4
	}
}

func (d Dynasty) Info() string {
	st := "DYNASTY: " + d.name + "\n"
	st += "POWER BASE: " + d.powerBase + "\n"
	st += "ARCHETYPE: " + d.archetype + "\n"
	st += "FIRST GENERATION BONUS: " + d.fgBonus + "\n"
	st += "CHARACTERISTICS:\n"
	for _, val := range listCharacteristics() {
		st += val + ": " + strconv.Itoa(d.stat[val]) + " (" + strconv.Itoa(DM(d.stat[val])) + ")\n"
	}
	st += "TRAITS:\n"
	for _, val := range listTraits() {
		st += val + ": " + strconv.Itoa(d.stat[val]) + "\n"
	}
	st += "APTITUDES:\n"
	for _, val := range listAptitudes() {
		if data, ok := d.stat[val]; ok {
			st += val + ": " + strconv.Itoa(data) + "\n"
		} else {
			st += val + ": ---\n"
		}
	}
	st += "VALUES:\n"
	for _, val := range listValues() {
		st += val + ": " + strconv.Itoa(d.stat[val]) + "\n"
	}
	st += "BOONS&HINDERS: "
	for i := range d.boonsHinders {
		st += d.boonsHinders[i] + ", "
	}
	st = strings.TrimSuffix(st, ", ")
	st += "\n"
	return st
}

//LISTS:

func listALL() []string {
	masterList := []string{}
	masterList = append(masterList, listCharacteristics()...)
	masterList = append(masterList, listTraits()...)
	masterList = append(masterList, listAptitudes()...)
	masterList = append(masterList, listValues()...)
	return masterList
}

func listCharacteristics() []string {
	return []string{
		Clv,
		Grd,
		Lty,
		Mil,
		Pop,
		Sch,
		Tcy,
		Tra,
	}
}

func listTraits() []string {
	return []string{
		Culture,
		FiscalDfnce,
		Fleet,
		Technology,
		TerritorialDfnce,
	}
}

func listAptitudes() []string {
	return []string{
		Acquisition,
		Bureaucracy,
		Conquest,
		Economics,
		Entertain,
		Expression,
		Hostility,
		Illicit,
		Intel,
		Maintenance,
		Politics,
		Posturing,
		Propaganda,
		PublicRelations,
		Recruit,
		Research,
		Sabotage,
		Security,
		Tactical,
		Tutelage,
	}
}

func listValues() []string {
	return []string{
		Morale,
		Populace,
		Wealth,
	}
}

func listPowerBase() []string {
	return []string{
		ColonySettlement,
		ConflictZone,
		Megalopolis,
		MilitaryCompound,
		NobleEstate,
		StarshipFlotilla,
		TempleHolyLand,
		UnchartedWilderness,
		UnderworldSlum,
		UrbanOffices,
	}
}

func listArchetypes() []string {
	return []string{
		Conglomerate,
		MediaEmpire,
		MerchantMarket,
		MilitaryCharter,
		NobleLine,
		ReligiousFaith,
		Syndicate,
	}
}

//POWER BASE

func lenStr(sl []string) string {
	return strconv.Itoa(len(sl))
}

func (d *Dynasty) choosePowerBase(pb string) error {
	if d.powerBase != "" {
		return errors.New("Power Base already chosen")
	}
	if pb == "" {
		i := dice.Roll("1d" + lenStr(listPowerBase())).DM(-1).Sum()
		pb = listPowerBase()[i]
	}
	d.powerBase = pb
	return nil
}

func (d *Dynasty) gainPowerBaseBonuses() error {
	switch d.powerBase {
	default:
		return errors.New("Unknown bonuses for Power Base '" + d.powerBase + "'")
	case ColonySettlement:
		d.stat[Culture]++
		d.stat[TerritorialDfnce]--
		d.raiseApttitude(Expression)
		d.raiseApttitude(Recruit)
		d.raiseApttitude(Maintenance)
		d.raiseApttitude(Propaganda)
		d.raiseApttitude(Tutelage)
	case ConflictZone:
		d.stat[TerritorialDfnce]++
		d.stat[TerritorialDfnce]++
		d.stat[FiscalDfnce]--
		d.stat[Fleet]--
		d.raiseApttitude(Hostility)
		d.raiseApttitude(Hostility)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
		d.raiseApttitude(Tactical)
	case Megalopolis:
		d.stat[FiscalDfnce]++
		d.stat[Technology]++
		d.stat[Culture]--
		d.stat[Culture]--
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Economics)
		d.raiseApttitude(PublicRelations)
		d.raiseApttitude(Research)
	case MilitaryCompound:
		d.stat[TerritorialDfnce]++
		d.stat[TerritorialDfnce]++
		d.stat[Fleet]++
		d.stat[FiscalDfnce]--
		d.stat[FiscalDfnce]--
		d.raiseApttitude(Conquest)
		d.raiseApttitude(Conquest)
		d.raiseApttitude(Tactical)
		d.raiseApttitude(Tactical)
		d.raiseApttitude(Politics)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
	case NobleEstate:
		d.stat[Culture]++
		d.stat[FiscalDfnce]++
		d.stat[TerritorialDfnce]--
		d.stat[TerritorialDfnce]--
		d.stat[Fleet]--
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Politics)
		d.raiseApttitude(Politics)
		d.raiseApttitude(Expression)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
	case StarshipFlotilla:
		d.stat[Fleet]++
		d.stat[Fleet]++
		d.stat[Technology]++
		d.stat[TerritorialDfnce]--
		d.stat[TerritorialDfnce]--
		d.raiseApttitude(Intel)
		d.raiseApttitude(Intel)
		d.raiseApttitude(Conquest)
		d.raiseApttitude(Economics)
		d.raiseApttitude(Maintenance)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Research)
		d.raiseApttitude(Tactical)
	case TempleHolyLand:
		d.stat[Culture]++
		d.stat[Culture]++
		d.stat[Technology]--
		d.stat[Technology]--
		d.raiseApttitude(Expression)
		d.raiseApttitude(Expression)
		d.raiseApttitude(Recruit)
		d.raiseApttitude(Recruit)
		d.raiseApttitude(Maintenance)
		d.raiseApttitude(Propaganda)
		d.raiseApttitude(PublicRelations)
		d.raiseApttitude(Tutelage)
	case UnchartedWilderness:
		d.stat[TerritorialDfnce]++
		d.stat[Technology]--
		d.raiseApttitude(Security)
		d.raiseApttitude(Security)
		d.raiseApttitude(Entertain)
		d.raiseApttitude(Illicit)
		d.raiseApttitude(Security) //или Conquest/Hostility?
	case UnderworldSlum:
		d.stat[FiscalDfnce]++
		d.stat[TerritorialDfnce]++
		d.stat[Culture]--
		d.stat[Culture]--
		d.raiseApttitude(Illicit)
		d.raiseApttitude(Illicit)
		d.raiseApttitude(Sabotage)
		d.raiseApttitude(Sabotage)
		d.raiseApttitude(Entertain)
		d.raiseApttitude(Intel)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
	case UrbanOffices:
		d.stat[Culture]++
		d.stat[FiscalDfnce]++
		d.stat[Fleet]--
		d.raiseApttitude(Acquisition)
		d.raiseApttitude(Acquisition)
		d.raiseApttitude(Economics)
		d.raiseApttitude(Economics)
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Intel)
		d.raiseApttitude(PublicRelations)
		d.raiseApttitude(Tutelage)
	}
	return nil
}

func validArchetypes(d Dynasty) []string {
	vArch := []string{}
	if d.stat[Grd] >= 8 && d.stat[Pop] >= 6 && d.stat[Tcy] >= 5 {
		vArch = append(vArch, Conglomerate)
	}
	if d.stat[Clv] >= 6 && d.stat[Pop] >= 8 && d.stat[Sch] >= 5 {
		vArch = append(vArch, MediaEmpire)
	}
	if d.stat[Clv] >= 6 && d.stat[Grd] >= 8 && d.stat[Pop] >= 5 {
		vArch = append(vArch, MerchantMarket)
	}
	if d.stat[Lty] >= 5 && d.stat[Mil] >= 8 && d.stat[Tra] >= 6 {
		vArch = append(vArch, MilitaryCharter)
	}
	if d.stat[Lty] >= 6 && d.stat[Tcy] >= 5 && d.stat[Tra] >= 8 {
		vArch = append(vArch, NobleLine)
	}
	if d.stat[Lty] >= 8 && d.stat[Pop] >= 5 && d.stat[Tra] >= 6 {
		vArch = append(vArch, ReligiousFaith)
	}
	if d.stat[Grd] >= 6 && d.stat[Sch] >= 8 && d.stat[Tcy] >= 5 {
		vArch = append(vArch, Syndicate)
	}
	return vArch
}

func (d *Dynasty) chooseArchetype(arch string) error {
	if d.archetype != "" {
		return errors.New("Archetype already chosen")
	}
	if arch == "" {
		i := dice.Roll("1d" + lenStr(listArchetypes())).DM(-1).Sum()
		arch = listArchetypes()[i]
	}
	d.archetype = arch
	return nil
}

func (d *Dynasty) determineBaseTraitsAndAptitudes() error {
	switch d.archetype {
	default:
		return errors.New("Unknown base for Archetype '" + d.archetype + "'")
	case Conglomerate:
		//APT
		d.ensureAptitude(Acquisition, 1)
		d.ensureAptitude(Bureaucracy, 1)
		d.ensureAptitude(Economics, 2)
		d.ensureAptitude(Intel, 0)
		d.ensureAptitude(Maintenance, 0)
		d.ensureAptitude(Propaganda, 0)
		d.ensureAptitude(PublicRelations, 0)
		d.ensureAptitude(Recruit, 0)
		d.ensureAptitude(Tutelage, 1)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Grd]) + DM(d.stat[Tra])
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Grd]) + DM(d.stat[Tcy]) + 1
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Mil]) + 1
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Grd]) + DM(d.stat[Lty])
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Lty]) + DM(d.stat[Pop])
	case MediaEmpire:
		d.ensureAptitude(Economics, 0)
		d.ensureAptitude(Entertain, 1)
		d.ensureAptitude(Expression, 0)
		d.ensureAptitude(Illicit, 0)
		d.ensureAptitude(Intel, 1)
		d.ensureAptitude(Politics, 0)
		d.ensureAptitude(Propaganda, 2)
		d.ensureAptitude(PublicRelations, 1)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Pop]) + DM(d.stat[Tra])
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Lty]) + 2
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Mil]) + 1
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Grd]) + DM(d.stat[Pop]) + 1
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Clv]) + DM(d.stat[Lty])
	case MerchantMarket:
		d.ensureAptitude(Acquisition, 1)
		d.ensureAptitude(Bureaucracy, 0)
		d.ensureAptitude(Economics, 2)
		d.ensureAptitude(Expression, 0)
		d.ensureAptitude(Intel, 0)
		d.ensureAptitude(Propaganda, 1)
		d.ensureAptitude(PublicRelations, 1)
		d.ensureAptitude(Recruit, 0)
		d.ensureAptitude(Research, 0)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Grd]) + DM(d.stat[Pop])
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Grd]) + DM(d.stat[Lty]) + 1
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Lty]) + DM(d.stat[Mil])
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Clv]) + DM(d.stat[Tra])
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Lty]) + 2
	case MilitaryCharter:
		d.ensureAptitude(Conquest, 1)
		d.ensureAptitude(Intel, 1)
		d.ensureAptitude(Maintenance, 0)
		d.ensureAptitude(Politics, 0)
		d.ensureAptitude(Recruit, 1)
		d.ensureAptitude(Security, 0)
		d.ensureAptitude(Tactical, 2)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Tra]) + 1
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Grd]) + DM(d.stat[Mil])
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Mil]) + DM(d.stat[Tcy]) + 1
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Mil]) + 1
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Mil]) + DM(d.stat[Pop]) + 1
	case NobleLine:
		d.ensureAptitude(Bureaucracy, 1)
		d.ensureAptitude(Economics, 0)
		d.ensureAptitude(Expression, 1)
		d.ensureAptitude(Illicit, 0)
		d.ensureAptitude(Politics, 2)
		d.ensureAptitude(Recruit, 0)
		d.ensureAptitude(Security, 0)
		d.ensureAptitude(Tutelage, 1)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Lty]) + DM(d.stat[Tra]) + 2
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Grd]) + DM(d.stat[Tcy])
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Mil]) + 1
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Tcy]) + 1
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Lty]) + DM(d.stat[Mil])
	case ReligiousFaith:
		d.ensureAptitude(Conquest, 0)
		d.ensureAptitude(Entertain, 0)
		d.ensureAptitude(Expression, 1)
		d.ensureAptitude(Politics, 0)
		d.ensureAptitude(Propaganda, 1)
		d.ensureAptitude(Recruit, 2)
		d.ensureAptitude(Security, 0)
		d.ensureAptitude(Tutelage, 1)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Lty]) + DM(d.stat[Tra]) + 2
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Grd]) + 1
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Lty]) + 1
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Clv]) + DM(d.stat[Tcy])
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Lty]) + DM(d.stat[Mil]) + 1
	case Syndicate:
		d.ensureAptitude(Conquest, 1)
		d.ensureAptitude(Entertain, 0)
		d.ensureAptitude(Expression, 0)
		d.ensureAptitude(Hostility, 0)
		d.ensureAptitude(Illicit, 2)
		d.ensureAptitude(Intel, 0)
		d.ensureAptitude(Posturing, 1)
		d.ensureAptitude(Sabotage, 1)
		d.ensureAptitude(Security, 0)
		//TRA
		d.stat[Culture] = d.stat[Culture] + DM(d.stat[Grd]) + DM(d.stat[Sch])
		d.stat[FiscalDfnce] = d.stat[FiscalDfnce] + DM(d.stat[Lty]) + 1
		d.stat[Fleet] = d.stat[Fleet] + DM(d.stat[Mil]) + DM(d.stat[Sch])
		d.stat[Technology] = d.stat[Technology] + DM(d.stat[Mil]) + 2
		d.stat[TerritorialDfnce] = d.stat[TerritorialDfnce] + DM(d.stat[Lty]) + DM(d.stat[Mil]) + 1
	}
	return nil
}

func (d *Dynasty) ensureAptitude(apt string, val int) {
	if v, ok := d.stat[apt]; ok {
		d.stat[apt] = v + val
		return
	}
	d.stat[apt] = val
}

func listBoons(arch string) []string {
	switch arch {
	default:
		return []string{}
	case Conglomerate:
		return []string{
			"Commercial Psions",
			"Endless Funds",
			"Governmental Backing",
			"Military Contracts",
			"Total Control",
			"Alien Extortions",
			"Market Mercenaries",
			"Spies in the Network",
			"Underworld Loans",
		}
	case MediaEmpire:
		return []string{
			"Bureaucratic Roots",
			"Gossip Rags",
			"Politics Engine",
			"Sports Contracts",
			"Voice of a Generation",
			"Hostile Paparazzi",
			"Pirate Comms Station",
			"Rumours of Corruption",
			"Translation Troubles",
		}
	case MerchantMarket:
		return []string{
			"Commercial Psions",
			"Interstellar Funding",
			"Naval Escorts",
			"Secure Production",
			"Vaulted Technologies",
			"Charitable Causes",
			"Depression Debts",
			"Pirate Problems",
			"Resource Mercenaries",
		}
	case MilitaryCharter:
		return []string{
			"Aggressive Politics",
			"Homeland Foundation",
			"Laurels of Victory",
			"Martial Law",
			"War Hero",
			"Enemies on All Fronts",
			"Gun Runner Gambles",
			"Tech Problems",
			"War Eternal",
		}
	case NobleLine:
		return []string{
			"Breeding Eugenics",
			"Inherited Fortunes",
			"Pocket Government",
			"Royal Family",
			"Secret Society",
			"Disease in the Genes",
			"Inbred Rumours",
			"Primitive Subjects",
			"Revolution in the Future",
		}
	case ReligiousFaith:
		return []string{
			"Alien Congregation",
			"Defenders of the Faith",
			"Holy Missionaries",
			"Tithes and Donations",
			"Words of Gods",
			"Atheist Coalition",
			"Controversial Clergy",
			"Superstitions Abound",
			"War Between Heavens",
		}
	case Syndicate:
		return []string{
			"Deadly Reputation",
			"Family of Crime",
			"Law Enforcement Spies",
			"Pirate Shipyard",
			"Rule Through Fear",
			"Bounty Hunters",
			"Grudges and Vendettas",
			"Most Wanted",
			"Question of Authority",
		}
	}
}

func (d *Dynasty) chooseBoons() {
	pick := d.dicepool.RollNext("1d6").DM(-2).Sum()
	if pick < 0 {
		pick = 0
	}
	selected := []string{}
	for i := 0; i < pick; i++ {
		selected = utils.AppendUniqueStr(selected, d.dicepool.RollFromList(listBoons(d.archetype)))
	}

	d.boonsHinders = selected
	d.initialBoonsEffect()
}

func (d *Dynasty) initialBoonsEffect() {
	for _, val := range d.boonsHinders {
		switch val {
		case "Commercial Psions":
			d.stat[Pop]--
		case "Endless Funds":
			d.stat[FiscalDfnce]--
			d.stat[FiscalDfnce]--
		case "Governmental Backing":
			d.stat[Tra]--
		case "Military Contracts":
			d.stat[Pop]--
			d.stat[Grd]--
		case "Total Control":
			d.stat[TerritorialDfnce]--
			d.stat[TerritorialDfnce]--
		case "Alien Extortions":
			d.stat[Grd]++
		case "Market Mercenaries":
			d.stat[Clv]++
			d.stat[Mil]++
		case "Spies in the Network":
			d.stat[Sch]++
		case "Underworld Loans":
			d.stat[FiscalDfnce]++
			d.stat[FiscalDfnce]++
		case "Bureaucratic Roots":
			d.stat[Grd]--
		case "Gossip Rags":
			d.stat[Culture]--
		case "Politics Engine":
			switch d.randomAction("-1 Lty", "-1 Sch") {
			case 1:
				d.stat[Lty]--
			case 2:
				d.stat[Sch]--
			}
		case "Sports Contracts":
			d.stat[Pop]--
			d.stat[Culture]--
		case "Voice of a Generation":
			d.stat[Pop]--
		case "Hostile Paparazzi":
			d.stat[Culture]++
			d.stat[Culture]++
		case "Pirate Comms Station":
			d.stat[Pop]++
			d.stat[Tcy]++
		case "Rumours of Corruption":
			d.stat[Clv]++
		case "Translation Troubles":
			d.stat[TerritorialDfnce]++
		case "Interstellar Funding":
			switch d.randomAction("-1 Tra", "-2 Culture") {
			case 1:
				d.stat[Tra]--
			case 2:
				d.stat[Culture]--
				d.stat[Culture]--
			}
		case "Naval Escorts":
			d.stat[Mil]--
		case "Secure Production":
			d.stat[FiscalDfnce]--
			d.stat[TerritorialDfnce]--
		case "Vaulted Technologies":
			d.stat[Technology]--
		case "Charitable Causes":
			d.stat[Culture]++
		case "Depression Debts":
			d.stat[Grd]++
		case "Pirate Problems":
			d.stat[TerritorialDfnce]++
			d.stat[TerritorialDfnce]++
		case "Resource Mercenaries":
			d.stat[Clv]++
			d.stat[Mil]++
		case "Aggressive Politics":
			d.stat[Pop]--
		case "Homeland Foundation":
			d.stat[Fleet]--
		case "Laurels of Victory":
			d.stat[Tcy]--
		case "Martial Law":
			d.stat[Lty]--
			d.stat[Culture]--
		case "War Hero":
			d.stat[Sch]--
		case "Enemies on All Fronts":
			d.stat[Clv]++
			d.stat[Mil]++
		case "Gun Runner Gambles":
			d.stat[Technology]++
		case "Tech Problems":
			d.stat[Tcy]++
		case "War Eternal":
			d.stat[TerritorialDfnce]++
			d.stat[TerritorialDfnce]++
		case "Breeding Eugenics":
			d.stat[Technology]--
		case "Inherited Fortunes":
			d.stat[FiscalDfnce]--
		case "Pocket Government":
			d.stat[Fleet]--
		case "Royal Family":
			d.stat[Lty]--
			d.stat[Technology]--
		case "Secret Society":
			d.stat[Sch]--
		case "Disease in the Genes":
			d.stat[Tra]++
		case "Inbred Rumours":
			d.stat[Culture]++
		case "Primitive Subjects":
			d.stat[FiscalDfnce]++
			d.stat[FiscalDfnce]++
		case "Revolution in the Future":
			d.stat[Sch]++
			d.stat[Mil]++
		case "Alien Congregation":
			d.stat[Pop]--
			d.stat[Culture]--
		case "Defenders of the Faith":
			d.stat[Sch]--
		case "Holy Missionaries":
			d.stat[Mil]--
		case "Tithes and Donations":
			d.stat[Culture]--
		case "Words of Gods":
			d.stat[Tra]--
		case "Atheist Coalition":
			d.stat[Tcy]++
			d.stat[Culture]++
		case "Controversial Clergy":
			d.stat[Lty]++
		case "Superstitions Abound":
			d.stat[Culture]++
		case "War Between Heavens":
			switch d.randomAction("+2 Td", "+1 Mil") {
			case 1:
				d.stat[TerritorialDfnce]++
				d.stat[TerritorialDfnce]++
			case 2:
				d.stat[Mil]++
			}
		case "Deadly Reputation":
			d.stat[Pop]--
		case "Family of Crime":
			d.stat[Lty]--
		case "Law Enforcement Spies":
			d.stat[Mil]--
		case "Pirate Shipyard":
			d.stat[Grd]--
			d.stat[FiscalDfnce]++
		case "Rule Through Fear":
			d.stat[Lty]++
		case "Bounty Hunters":
			d.stat[Clv]++
			d.stat[Mil]++
		case "Grudges and Vendettas":
			d.stat[Lty]++
		case "Most Wanted":
			d.stat[Sch]++
			d.stat[Culture]++
		case "Question of Authority":
			d.stat[Grd]++
		}
	}
}

func (d *Dynasty) gainFirstGenerationBonuses() error {
	r := d.dicepool.RollNext("2d6").Sum()
	//fmt.Print("First Generation Bonus: ")
	fgbM := fgbMap()
	switch r {
	case 2:
		d.fgBonus = fgbM[d.archetype][0]
	case 3, 4:
		d.fgBonus = fgbM[d.archetype][1]
	case 5, 6:
		d.fgBonus = fgbM[d.archetype][2]
	case 7:
		d.fgBonus = fgbM[d.archetype][3]
	case 8, 9:
		d.fgBonus = fgbM[d.archetype][4]
	case 10, 11:
		d.fgBonus = fgbM[d.archetype][5]
	case 12:
		d.fgBonus = fgbM[d.archetype][6]
	}
	d.applyFGB()
	return nil
}

func fgbMap() map[string][]string {
	fgbM := make(map[string][]string)
	fgbM[Conglomerate] = []string{
		"University Board Members",
		"Monopoly",
		"Shipyard Access",
		"Noble Investors",
		"Inherited Pride",
		"Multi-stellar Benefactor",
		"Royal Backing",
	}
	fgbM[MediaEmpire] = []string{
		"Psionic Investigators",
		"Pyramid Structure",
		"High-Tech Communications",
		"Noble Investors",
		"Military Reporters",
		"Interstellar Cover Story",
		"Total Media Monopoly",
	}
	fgbM[MerchantMarket] = []string{
		"Barter Over Sales",
		"Monopoly",
		"Intense Collegiate Training",
		"Patents Upon Patents",
		"Governmental Acquisitions",
		"A Republic in Good Fortune",
		"Perfect Economy",
	}
	fgbM[MilitaryCharter] = []string{
		"Intense Generational Training",
		"War Coffers",
		"Naval Partners",
		"An Armed Populace",
		"Victory over Invasion",
		"War Colleges",
		"Noble Armada",
	}
	fgbM[NobleLine] = []string{
		"Royal Guard",
		"Of Pawns and Kings",
		"The Love of the People",
		"Order of Protectors",
		"Military Honour",
		"Interstellar Marriages",
		"No Peers In Sight",
	}
	fgbM[ReligiousFaith] = []string{
		"Clergy Scholars",
		"Knights and Templars",
		"Holy Treasures",
		"Family Comes First",
		"Online Scripture",
		"Blessings from Beyond",
		"Living Legends",
	}
	fgbM[Syndicate] = []string{
		"Undeniable Success",
		"Art Thieves and Extortions",
		"Pirate Captains",
		"Gangs Upon Gangs",
		"Tougher than the Street",
		"Empire of Crime",
		"Intergalactic Mafia",
	}

	return fgbM
}

func (d *Dynasty) applyFGB() {
	switch d.fgBonus {
	case "University Board Members":
		d.raiseAny3AptitudesToLevel1()
	case "Monopoly":
		d.stat[FiscalDfnce]++
	case "Shipyard Access":
		d.stat[Fleet]++
	case "Noble Investors":
		d.stat[Wealth]++
	case "Inherited Pride":
		d.stat[Culture]++
	case "Multi-stellar Benefactor":
		d.raiseAny2TraitsBy1()
	case "Royal Backing":
		d.add1d6PointsToValues()
	case "Psionic Investigators":
		d.raiseAny2AptitudesBy1()
	case "Pyramid Structure":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.stat[Wealth] = d.stat[Wealth] + d.dicepool.RollNext("1d6").Sum()
		case 2:
			d.stat[Grd]++
		}
	case "High-Tech Communications":
		d.stat[Technology]++
	case "Military Reporters":
		d.ensureAptitude(Conquest, 1)
		d.ensureAptitude(Security, 1)
	case "Interstellar Cover Story":
		d.ensureAptitude(Expression, 1)
		d.ensureAptitude(Politics, 1)
		d.ensureAptitude(Posturing, 1)
	case "Total Media Monopoly":
		d.add1d6PointsToValues()
	case "Barter Over Sales":
		d.raiseAnyLevel0AptitudeTo1()
		d.raiseAnyLevel0AptitudeTo1()
	case "Intense Collegiate Training":
		d.raiseApttitude(d.dicepool.RollFromList(listAptitudes()))
	case "Patents Upon Patents":
		d.stat[Wealth]++
	case "Governmental Acquisitions":
		d.stat[Fleet]++
	case "A Republic in Good Fortune":
		d.stat[Wealth] = d.stat[Wealth] + d.dicepool.RollNext("1d6").Sum()
	case "Perfect Economy":
		d.add1d6PointsToValues()
	case "Intense Generational Training":
		d.raiseAny3AptitudesToLevel1()
	case "War Coffers":
		d.stat[Wealth]++
	case "Naval Partners":
		d.stat[Fleet]++
	case "An Armed Populace":
		d.stat[TerritorialDfnce]++
	case "Victory over Invasion":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.stat[Culture] = d.stat[Culture] + 2
		case 2:
			d.stat[Tra]++
		}
	case "War Colleges":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.raiseAnyLevel0AptitudeTo1()
			d.raiseAnyLevel0AptitudeTo1()
			d.raiseAnyLevel0AptitudeTo1()
		case 2:
			d.raiseAnyLevel1AptitudeBy1()
			d.raiseAnyLevel1AptitudeBy1()
		}
	case "Noble Armada":
		d.stat[Tra]++
		d.stat[Fleet]++
		d.stat[Fleet]++
	case "Royal Guard":
		d.stat[Mil]++
		d.stat[TerritorialDfnce]++
	case "Of Pawns and Kings":
		d.stat[Sch]++
	case "The Love of the People":
		d.stat[Morale]++
	case "Order of Protectors":
		d.stat[TerritorialDfnce]++
	case "Military Honour":
		d.stat[Mil]++
	case "Interstellar Marriages":
		d.stat[Culture]++
		d.stat[Fleet]++
	case "No Peers In Sight":
		d.add1d6PointsToValues()
	case "Clergy Scholars":
		d.raiseAny3AptitudesToLevel1()
	case "Knights and Templars":
		d.stat[TerritorialDfnce]++
	case "Holy Treasures":
		d.stat[Wealth]++
	case "Family Comes First":
		d.stat[Populace]++
	case "Online Scripture":
		d.stat[Technology]++
	case "Blessings from Beyond":
		d.stat[d.dicepool.RollFromList(listTraits())]++
		d.stat[d.dicepool.RollFromList(listTraits())]++
	case "Living Legends":
		d.add1d6PointsToValues()
	case "Undeniable Success":
		d.raiseAny3AptitudesToLevel1()
	case "Art Thieves and Extortions":
		d.stat[FiscalDfnce]++
	case "Pirate Captains":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.stat[Fleet]++
		case 2:
			d.stat[Mil]++
		}
	case "Gangs Upon Gangs":
		d.stat[Populace]++
	case "Tougher than the Street":
		d.stat[TerritorialDfnce]++
		d.stat[Technology]++
	case "Empire of Crime":
		d.stat[d.dicepool.RollFromList(listTraits())]++
		d.stat[d.dicepool.RollFromList(listTraits())]++
	case "Intergalactic Mafia":
		d.add1d6PointsToValues()
	}
}

func (d *Dynasty) raiseAny3AptitudesToLevel1() {
	validToRaise := []string{}
	for _, apt := range listAptitudes() {
		if val, ok := d.stat[apt]; !ok {
			if val < 1 {
				validToRaise = append(validToRaise, apt)
			}
		} else {
			validToRaise = append(validToRaise, apt)
		}
	}
	for len(validToRaise) > 3 {
		r := d.dicepool.RollNext("1d" + lenStr(validToRaise)).DM(-1).Sum()
		validToRaise = remove(validToRaise, r)
	}
	for _, val := range validToRaise {
		d.stat[val] = 1
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (d *Dynasty) raiseAny2TraitsBy1() {
	for i := 0; i < 2; i++ {
		d.stat[d.dicepool.RollFromList(listTraits())]++
	}
}

func (d *Dynasty) add1d6PointsToValues() {
	r := d.dicepool.RollNext("1d6").Sum()
	for i := 0; i < r; i++ {
		d.stat[utils.RandomFromList(listValues())]++
	}
}

func (d *Dynasty) raiseApttitude(apt string) {
	if _, ok := d.stat[apt]; ok {
		d.stat[apt]++
		return
	}
	d.stat[apt] = 0
}

func (d *Dynasty) raiseAny2AptitudesBy1() {
	for i := 0; i < 2; i++ {
		d.raiseApttitude(utils.RandomFromList(listAptitudes()))
	}
}

func (d *Dynasty) raiseAnyLevel0AptitudeTo1() {
	validToRaise := []string{}
	for _, apt := range listAptitudes() {
		if val, ok := d.stat[apt]; !ok {
			if val == 0 {
				validToRaise = append(validToRaise, apt)
			}
		}
	}
	for len(validToRaise) > 1 {
		r := d.dicepool.RollNext("1d" + lenStr(validToRaise)).DM(-1).Sum()
		validToRaise = remove(validToRaise, r)
	}
	for _, val := range validToRaise {
		d.stat[val] = 1
	}
}

func (d *Dynasty) raiseAnyLevel1AptitudeBy1() {
	validToRaise := []string{}
	for _, apt := range listAptitudes() {
		if val, ok := d.stat[apt]; !ok {
			if val == 1 {
				validToRaise = append(validToRaise, apt)
			}
		}
	}
	for len(validToRaise) > 1 {
		r := d.dicepool.RollNext("1d" + lenStr(validToRaise)).DM(-1).Sum()
		validToRaise = remove(validToRaise, r)
	}
	for _, val := range validToRaise {
		d.stat[val] = 2
	}
}

func (d *Dynasty) randomAction(act ...string) int {
	l := strconv.Itoa(len(act))
	r := d.dicepool.RollNext("1d" + l).Sum()
	return r
}

func reportErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func (d *Dynasty) changeStatBy(stat string, increment int) {
	d.story += "\n " + stat + " " + strconv.Itoa(increment)
	d.stat[stat] = d.stat[stat] + increment
	// if utils.ListContains(listAptitudes(), stat) { - не в этом месте.
	// 	if d.stat[stat] < 0 {
	// 		d.stat[stat] = 0
	// 	}
	// }
}

func (d *Dynasty) ensureAptValidRange() {
	for _, val := range listAptitudes() {
		if d.stat[val] < 0 {
			d.stat[val] = 0
		}
		if d.stat[val] > 5 {
			d.stat[val] = 5
		}
	}
}

func (d *Dynasty) anyCharacteristic() string {
	return d.dicepool.RollFromList(listCharacteristics())
}

func (d *Dynasty) aptitudeValue(apt string) int {
	if val, ok := d.stat[apt]; ok {
		return val
	}
	return -2
}

func (d *Dynasty) characteristicDM(chr string) int {
	if val, ok := d.stat[chr]; ok {
		return DM(val)
	}
	panic("Dynasty Inconsistent!")
	return -2
}
