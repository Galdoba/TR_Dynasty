package main

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

func main() {
	Test()
}

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
	Populance           = "Populance"
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

func Test() {
	validDyn := false
	d := dynasty{}
	for !validDyn {
		d = NewDynasty("")
		validDyn = Survived(d)
	}
	fmt.Print(d.Info())
	curentDay := 403700
	date := DateManager.FormatToDate(curentDay)
	reverseDay, err := DateManager.FormatToDay(date)
	fmt.Println(reverseDay, err, date)
}

type dynasty struct {
	name            string
	dicepool        dice.Dicepool
	characteristics map[string]int
	traits          map[string]int
	aptitudes       map[string]int
	values          map[string]int
	boonsHinders    []string
	archetype       string
	powerBase       string
	fgBonus         string
	managment       string
	birthDate       string
	nextEventDate   string
	nextActionDate  string
	historicEvents  int
}

func NewDynasty(name string) dynasty {
	d := dynasty{}
	seed := utils.SeedFromString(name)

	if name == "" {
		seed = time.Now().UnixNano()
		name = "[RANDOM NAME] seed=" + strconv.Itoa(int(seed))
	}
	d.name = name

	d.dicepool = *dice.New(seed)
	d.characteristics = make(map[string]int)
	d.traits = make(map[string]int)
	d.aptitudes = make(map[string]int)
	d.values = make(map[string]int)
	//1
	for _, val := range listCharacteristics() {
		d.characteristics[val] = d.dicepool.RollNext("2d6").Sum()
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

func (d *dynasty) finalizeFirstGeneration() {
	d.fgStep1() //train characteristics
	d.fgStep2() //train management
	d.fgStep3() //train aptitudes
	d.fgStep4() //Final Values Adjustments
}

func (d *dynasty) fgStep1() {
	characteristicsPractice := []string{}
	for len(characteristicsPractice) < 3 {
		characteristicsPractice = append(characteristicsPractice, d.dicepool.RollFromList(listCharacteristics()))
	}
	for _, val := range characteristicsPractice {
		r := d.dicepool.RollNext("2d6").Sum()
		switch r {
		case 2:
			d.characteristics[val]--
		case 12:
			d.characteristics[val]++
		default:
			if DM(d.characteristics[val])+r > d.characteristics[val] {
				d.characteristics[val]++
			}
		}
	}
}

func (d *dynasty) fgStep2() {
	max := DM(d.characteristics[Clv]) + 1
	for i := 0; i < max; i++ {
		moveFrom := d.dicepool.RollFromList(listTraits())
		moveTo := d.dicepool.RollFromList(listTraits())
		if d.traits[moveFrom] < 2 {
			continue
		}
		d.traits[moveFrom]--
		d.traits[moveTo]++
	}
}

func (d *dynasty) fgStep3() {
	aptPractice := []string{}
	for len(aptPractice) < 5 {
		aptPractice = utils.AppendUniqueStr(aptPractice, d.dicepool.RollFromList(listAptitudes()))
	}
	for _, val := range aptPractice {
		r := d.dicepool.RollNext("2d6").DM(-8).Sum()
		e := -2
		if apt, ok := d.aptitudes[val]; ok {
			e = apt
		}
		r = r + e
		if r > d.aptitudes[val] {
			d.raiseApttitude(val)
		}
	}
}

func (d *dynasty) fgStep4() {
	d.values[Morale] = d.values[Morale] + DM(d.characteristics[Pop]) + d.traits[Culture]
	d.values[Populance] = d.values[Populance] + DM(d.characteristics[Tra]) + d.traits[Technology]
	d.values[Wealth] = d.values[Wealth] + DM(d.characteristics[Clv]) + d.traits[FiscalDfnce]
	for _, val := range listValues() {
		if d.values[val] < 1 {
			continue
		}
		d.values[val]--
		d.values[d.dicepool.RollFromList(listValues())]++
	}
}

func (d *dynasty) chooseManagementAsset() {
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

func (d dynasty) Info() string {
	st := "DYNASTY: " + d.name + "\n"
	st += "POWER BASE: " + d.powerBase + "\n"
	st += "ARCHETYPE: " + d.archetype + "\n"
	st += "FIRST GENERATION BONUS: " + d.fgBonus + "\n"
	st += "CHARACTERISTICS:\n"
	for _, val := range listCharacteristics() {
		st += val + ": " + strconv.Itoa(d.characteristics[val]) + " (" + strconv.Itoa(DM(d.characteristics[val])) + ")\n"
	}
	st += "TRAITS:\n"
	for _, val := range listTraits() {
		st += val + ": " + strconv.Itoa(d.traits[val]) + "\n"
	}
	st += "APTITUDES:\n"
	for _, val := range listAptitudes() {
		if data, ok := d.aptitudes[val]; ok {
			st += val + ": " + strconv.Itoa(data) + "\n"
		} else {
			st += val + ": ---\n"
		}
	}
	st += "VALUES:\n"
	for _, val := range listValues() {
		st += val + ": " + strconv.Itoa(d.values[val]) + "\n"
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
		Populance,
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

func (d *dynasty) choosePowerBase(pb string) error {
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

func (d *dynasty) gainPowerBaseBonuses() error {
	switch d.powerBase {
	default:
		return errors.New("Unknown bonuses for Power Base '" + d.powerBase + "'")
	case ColonySettlement:
		d.traits[Culture]++
		d.traits[TerritorialDfnce]--
		d.raiseApttitude(Expression)
		d.raiseApttitude(Recruit)
		d.raiseApttitude(Maintenance)
		d.raiseApttitude(Propaganda)
		d.raiseApttitude(Tutelage)
	case ConflictZone:
		d.traits[TerritorialDfnce]++
		d.traits[TerritorialDfnce]++
		d.traits[FiscalDfnce]--
		d.traits[Fleet]--
		d.raiseApttitude(Hostility)
		d.raiseApttitude(Hostility)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
		d.raiseApttitude(Tactical)
	case Megalopolis:
		d.traits[FiscalDfnce]++
		d.traits[Technology]++
		d.traits[Culture]--
		d.traits[Culture]--
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Economics)
		d.raiseApttitude(PublicRelations)
		d.raiseApttitude(Research)
	case MilitaryCompound:
		d.traits[TerritorialDfnce]++
		d.traits[TerritorialDfnce]++
		d.traits[Fleet]++
		d.traits[FiscalDfnce]--
		d.traits[FiscalDfnce]--
		d.raiseApttitude(Conquest)
		d.raiseApttitude(Conquest)
		d.raiseApttitude(Tactical)
		d.raiseApttitude(Tactical)
		d.raiseApttitude(Politics)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
	case NobleEstate:
		d.traits[Culture]++
		d.traits[FiscalDfnce]++
		d.traits[TerritorialDfnce]--
		d.traits[TerritorialDfnce]--
		d.traits[Fleet]--
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Bureaucracy)
		d.raiseApttitude(Politics)
		d.raiseApttitude(Politics)
		d.raiseApttitude(Expression)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
	case StarshipFlotilla:
		d.traits[Fleet]++
		d.traits[Fleet]++
		d.traits[Technology]++
		d.traits[TerritorialDfnce]--
		d.traits[TerritorialDfnce]--
		d.raiseApttitude(Intel)
		d.raiseApttitude(Intel)
		d.raiseApttitude(Conquest)
		d.raiseApttitude(Economics)
		d.raiseApttitude(Maintenance)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Research)
		d.raiseApttitude(Tactical)
	case TempleHolyLand:
		d.traits[Culture]++
		d.traits[Culture]++
		d.traits[Technology]--
		d.traits[Technology]--
		d.raiseApttitude(Expression)
		d.raiseApttitude(Expression)
		d.raiseApttitude(Recruit)
		d.raiseApttitude(Recruit)
		d.raiseApttitude(Maintenance)
		d.raiseApttitude(Propaganda)
		d.raiseApttitude(PublicRelations)
		d.raiseApttitude(Tutelage)
	case UnchartedWilderness:
		d.traits[TerritorialDfnce]++
		d.traits[Technology]--
		d.raiseApttitude(Security)
		d.raiseApttitude(Security)
		d.raiseApttitude(Entertain)
		d.raiseApttitude(Illicit)
		d.raiseApttitude(Security) //или Conquest/Hostility?
	case UnderworldSlum:
		d.traits[FiscalDfnce]++
		d.traits[TerritorialDfnce]++
		d.traits[Culture]--
		d.traits[Culture]--
		d.raiseApttitude(Illicit)
		d.raiseApttitude(Illicit)
		d.raiseApttitude(Sabotage)
		d.raiseApttitude(Sabotage)
		d.raiseApttitude(Entertain)
		d.raiseApttitude(Intel)
		d.raiseApttitude(Posturing)
		d.raiseApttitude(Security)
	case UrbanOffices:
		d.traits[Culture]++
		d.traits[FiscalDfnce]++
		d.traits[Fleet]--
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

func validArchetypes(d dynasty) []string {
	vArch := []string{}
	if d.characteristics[Grd] >= 8 && d.characteristics[Pop] >= 6 && d.characteristics[Tcy] >= 5 {
		vArch = append(vArch, Conglomerate)
	}
	if d.characteristics[Clv] >= 6 && d.characteristics[Pop] >= 8 && d.characteristics[Sch] >= 5 {
		vArch = append(vArch, MediaEmpire)
	}
	if d.characteristics[Clv] >= 6 && d.characteristics[Grd] >= 8 && d.characteristics[Pop] >= 5 {
		vArch = append(vArch, MerchantMarket)
	}
	if d.characteristics[Lty] >= 5 && d.characteristics[Mil] >= 8 && d.characteristics[Tra] >= 6 {
		vArch = append(vArch, MilitaryCharter)
	}
	if d.characteristics[Lty] >= 6 && d.characteristics[Tcy] >= 5 && d.characteristics[Tra] >= 8 {
		vArch = append(vArch, NobleLine)
	}
	if d.characteristics[Lty] >= 8 && d.characteristics[Pop] >= 5 && d.characteristics[Tra] >= 6 {
		vArch = append(vArch, ReligiousFaith)
	}
	if d.characteristics[Grd] >= 6 && d.characteristics[Sch] >= 8 && d.characteristics[Tcy] >= 5 {
		vArch = append(vArch, Syndicate)
	}
	return vArch
}

func (d *dynasty) chooseArchetype(arch string) error {
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

func (d *dynasty) determineBaseTraitsAndAptitudes() error {
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
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Grd]) + DM(d.characteristics[Tra])
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Grd]) + DM(d.characteristics[Tcy]) + 1
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Mil]) + 1
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Grd]) + DM(d.characteristics[Lty])
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Lty]) + DM(d.characteristics[Pop])
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
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Pop]) + DM(d.characteristics[Tra])
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Lty]) + 2
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Mil]) + 1
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Grd]) + DM(d.characteristics[Pop]) + 1
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Clv]) + DM(d.characteristics[Lty])
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
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Grd]) + DM(d.characteristics[Pop])
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Grd]) + DM(d.characteristics[Lty]) + 1
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Lty]) + DM(d.characteristics[Mil])
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Clv]) + DM(d.characteristics[Tra])
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Lty]) + 2
	case MilitaryCharter:
		d.ensureAptitude(Conquest, 1)
		d.ensureAptitude(Intel, 1)
		d.ensureAptitude(Maintenance, 0)
		d.ensureAptitude(Politics, 0)
		d.ensureAptitude(Recruit, 1)
		d.ensureAptitude(Security, 0)
		d.ensureAptitude(Tactical, 2)
		//TRA
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Tra]) + 1
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Grd]) + DM(d.characteristics[Mil])
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Mil]) + DM(d.characteristics[Tcy]) + 1
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Mil]) + 1
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Mil]) + DM(d.characteristics[Pop]) + 1
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
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Lty]) + DM(d.characteristics[Tra]) + 2
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Grd]) + DM(d.characteristics[Tcy])
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Mil]) + 1
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Tcy]) + 1
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Lty]) + DM(d.characteristics[Mil])
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
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Lty]) + DM(d.characteristics[Tra]) + 2
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Grd]) + 1
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Lty]) + 1
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Clv]) + DM(d.characteristics[Tcy])
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Lty]) + DM(d.characteristics[Mil]) + 1
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
		d.traits[Culture] = d.traits[Culture] + DM(d.characteristics[Grd]) + DM(d.characteristics[Sch])
		d.traits[FiscalDfnce] = d.traits[FiscalDfnce] + DM(d.characteristics[Lty]) + 1
		d.traits[Fleet] = d.traits[Fleet] + DM(d.characteristics[Mil]) + DM(d.characteristics[Sch])
		d.traits[Technology] = d.traits[Technology] + DM(d.characteristics[Mil]) + 2
		d.traits[TerritorialDfnce] = d.traits[TerritorialDfnce] + DM(d.characteristics[Lty]) + DM(d.characteristics[Mil]) + 1
	}
	return nil
}

func (d *dynasty) ensureAptitude(apt string, val int) {
	if v, ok := d.aptitudes[apt]; ok {
		d.aptitudes[apt] = v + val
		return
	}
	d.aptitudes[apt] = val
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

func (d *dynasty) chooseBoons() {
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

func (d *dynasty) initialBoonsEffect() {
	for _, val := range d.boonsHinders {
		switch val {
		case "Commercial Psions":
			d.characteristics[Pop]--
		case "Endless Funds":
			d.traits[FiscalDfnce]--
			d.traits[FiscalDfnce]--
		case "Governmental Backing":
			d.characteristics[Tra]--
		case "Military Contracts":
			d.characteristics[Pop]--
			d.characteristics[Grd]--
		case "Total Control":
			d.traits[TerritorialDfnce]--
			d.traits[TerritorialDfnce]--
		case "Alien Extortions":
			d.characteristics[Grd]++
		case "Market Mercenaries":
			d.characteristics[Clv]++
			d.characteristics[Mil]++
		case "Spies in the Network":
			d.characteristics[Sch]++
		case "Underworld Loans":
			d.traits[FiscalDfnce]++
			d.traits[FiscalDfnce]++
		case "Bureaucratic Roots":
			d.characteristics[Grd]--
		case "Gossip Rags":
			d.traits[Culture]--
		case "Politics Engine":
			switch d.randomAction("-1 Lty", "-1 Sch") {
			case 1:
				d.characteristics[Lty]--
			case 2:
				d.characteristics[Sch]--
			}
		case "Sports Contracts":
			d.characteristics[Pop]--
			d.traits[Culture]--
		case "Voice of a Generation":
			d.characteristics[Pop]--
		case "Hostile Paparazzi":
			d.traits[Culture]++
			d.traits[Culture]++
		case "Pirate Comms Station":
			d.characteristics[Pop]++
			d.characteristics[Tcy]++
		case "Rumours of Corruption":
			d.characteristics[Clv]++
		case "Translation Troubles":
			d.traits[TerritorialDfnce]++
		case "Interstellar Funding":
			switch d.randomAction("-1 Tra", "-2 Culture") {
			case 1:
				d.characteristics[Tra]--
			case 2:
				d.traits[Culture]--
				d.traits[Culture]--
			}
		case "Naval Escorts":
			d.characteristics[Mil]--
		case "Secure Production":
			d.traits[FiscalDfnce]--
			d.traits[TerritorialDfnce]--
		case "Vaulted Technologies":
			d.traits[Technology]--
		case "Charitable Causes":
			d.traits[Culture]++
		case "Depression Debts":
			d.characteristics[Grd]++
		case "Pirate Problems":
			d.traits[TerritorialDfnce]++
			d.traits[TerritorialDfnce]++
		case "Resource Mercenaries":
			d.characteristics[Clv]++
			d.characteristics[Mil]++
		case "Aggressive Politics":
			d.characteristics[Pop]--
		case "Homeland Foundation":
			d.traits[Fleet]--
		case "Laurels of Victory":
			d.characteristics[Tcy]--
		case "Martial Law":
			d.characteristics[Lty]--
			d.traits[Culture]--
		case "War Hero":
			d.characteristics[Sch]--
		case "Enemies on All Fronts":
			d.characteristics[Clv]++
			d.characteristics[Mil]++
		case "Gun Runner Gambles":
			d.traits[Technology]++
		case "Tech Problems":
			d.characteristics[Tcy]++
		case "War Eternal":
			d.traits[TerritorialDfnce]++
			d.traits[TerritorialDfnce]++
		case "Breeding Eugenics":
			d.traits[Technology]--
		case "Inherited Fortunes":
			d.traits[FiscalDfnce]--
		case "Pocket Government":
			d.traits[Fleet]--
		case "Royal Family":
			d.characteristics[Lty]--
			d.traits[Technology]--
		case "Secret Society":
			d.characteristics[Sch]--
		case "Disease in the Genes":
			d.characteristics[Tra]++
		case "Inbred Rumours":
			d.traits[Culture]++
		case "Primitive Subjects":
			d.traits[FiscalDfnce]++
			d.traits[FiscalDfnce]++
		case "Revolution in the Future":
			d.characteristics[Sch]++
			d.characteristics[Mil]++
		case "Alien Congregation":
			d.characteristics[Pop]--
			d.traits[Culture]--
		case "Defenders of the Faith":
			d.characteristics[Sch]--
		case "Holy Missionaries":
			d.characteristics[Mil]--
		case "Tithes and Donations":
			d.traits[Culture]--
		case "Words of Gods":
			d.characteristics[Tra]--
		case "Atheist Coalition":
			d.characteristics[Tcy]++
			d.traits[Culture]++
		case "Controversial Clergy":
			d.characteristics[Lty]++
		case "Superstitions Abound":
			d.traits[Culture]++
		case "War Between Heavens":
			switch d.randomAction("+2 Td", "+1 Mil") {
			case 1:
				d.traits[TerritorialDfnce]++
				d.traits[TerritorialDfnce]++
			case 2:
				d.characteristics[Mil]++
			}
		case "Deadly Reputation":
			d.characteristics[Pop]--
		case "Family of Crime":
			d.characteristics[Lty]--
		case "Law Enforcement Spies":
			d.characteristics[Mil]--
		case "Pirate Shipyard":
			d.characteristics[Grd]--
			d.traits[FiscalDfnce]++
		case "Rule Through Fear":
			d.characteristics[Lty]++
		case "Bounty Hunters":
			d.characteristics[Clv]++
			d.characteristics[Mil]++
		case "Grudges and Vendettas":
			d.characteristics[Lty]++
		case "Most Wanted":
			d.characteristics[Sch]++
			d.traits[Culture]++
		case "Question of Authority":
			d.characteristics[Grd]++
		}
	}
}

func (d *dynasty) gainFirstGenerationBonuses() error {
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

func (d *dynasty) applyFGB() {
	switch d.fgBonus {
	case "University Board Members":
		d.raiseAny3AptitudesToLevel1()
	case "Monopoly":
		d.traits[FiscalDfnce]++
	case "Shipyard Access":
		d.traits[Fleet]++
	case "Noble Investors":
		d.values[Wealth]++
	case "Inherited Pride":
		d.traits[Culture]++
	case "Multi-stellar Benefactor":
		d.raiseAny2TraitsBy1()
	case "Royal Backing":
		d.add1d6PointsToValues()
	case "Psionic Investigators":
		d.raiseAny2AptitudesBy1()
	case "Pyramid Structure":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.values[Wealth] = d.values[Wealth] + d.dicepool.RollNext("1d6").Sum()
		case 2:
			d.characteristics[Grd]++
		}
	case "High-Tech Communications":
		d.traits[Technology]++
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
		d.values[Wealth]++
	case "Governmental Acquisitions":
		d.traits[Fleet]++
	case "A Republic in Good Fortune":
		d.values[Wealth] = d.values[Wealth] + d.dicepool.RollNext("1d6").Sum()
	case "Perfect Economy":
		d.add1d6PointsToValues()
	case "Intense Generational Training":
		d.raiseAny3AptitudesToLevel1()
	case "War Coffers":
		d.values[Wealth]++
	case "Naval Partners":
		d.traits[Fleet]++
	case "An Armed Populace":
		d.traits[TerritorialDfnce]++
	case "Victory over Invasion":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.traits[Culture] = d.traits[Culture] + 2
		case 2:
			d.characteristics[Tra]++
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
		d.characteristics[Tra]++
		d.traits[Fleet]++
		d.traits[Fleet]++
	case "Royal Guard":
		d.characteristics[Mil]++
		d.traits[TerritorialDfnce]++
	case "Of Pawns and Kings":
		d.characteristics[Sch]++
	case "The Love of the People":
		d.values[Morale]++
	case "Order of Protectors":
		d.traits[TerritorialDfnce]++
	case "Military Honour":
		d.characteristics[Mil]++
	case "Interstellar Marriages":
		d.traits[Culture]++
		d.traits[Fleet]++
	case "No Peers In Sight":
		d.add1d6PointsToValues()
	case "Clergy Scholars":
		d.raiseAny3AptitudesToLevel1()
	case "Knights and Templars":
		d.traits[TerritorialDfnce]++
	case "Holy Treasures":
		d.values[Wealth]++
	case "Family Comes First":
		d.values[Populance]++
	case "Online Scripture":
		d.traits[Technology]++
	case "Blessings from Beyond":
		d.traits[d.dicepool.RollFromList(listTraits())]++
		d.traits[d.dicepool.RollFromList(listTraits())]++
	case "Living Legends":
		d.add1d6PointsToValues()
	case "Undeniable Success":
		d.raiseAny3AptitudesToLevel1()
	case "Art Thieves and Extortions":
		d.traits[FiscalDfnce]++
	case "Pirate Captains":
		switch d.dicepool.RollNext("1d2").Sum() {
		case 1:
			d.traits[Fleet]++
		case 2:
			d.characteristics[Mil]++
		}
	case "Gangs Upon Gangs":
		d.values[Populance]++
	case "Tougher than the Street":
		d.traits[TerritorialDfnce]++
		d.traits[Technology]++
	case "Empire of Crime":
		d.traits[d.dicepool.RollFromList(listTraits())]++
		d.traits[d.dicepool.RollFromList(listTraits())]++
	case "Intergalactic Mafia":
		d.add1d6PointsToValues()
	}
}

func (d *dynasty) raiseAny3AptitudesToLevel1() {
	validToRaise := []string{}
	for _, apt := range listAptitudes() {
		if val, ok := d.aptitudes[apt]; !ok {
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
		d.aptitudes[val] = 1
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (d *dynasty) raiseAny2TraitsBy1() {
	for i := 0; i < 2; i++ {
		d.traits[d.dicepool.RollFromList(listTraits())]++
	}
}

func (d *dynasty) add1d6PointsToValues() {
	r := d.dicepool.RollNext("1d6").Sum()
	for i := 0; i < r; i++ {
		d.values[utils.RandomFromList(listValues())]++
	}
}

func (d *dynasty) raiseApttitude(apt string) {
	if _, ok := d.aptitudes[apt]; ok {
		d.aptitudes[apt]++
		return
	}
	d.aptitudes[apt] = 0
}

func (d *dynasty) raiseAny2AptitudesBy1() {
	for i := 0; i < 2; i++ {
		d.raiseApttitude(utils.RandomFromList(listAptitudes()))
	}
}

func (d *dynasty) raiseAnyLevel0AptitudeTo1() {
	validToRaise := []string{}
	for _, apt := range listAptitudes() {
		if val, ok := d.aptitudes[apt]; !ok {
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
		d.aptitudes[val] = 1
	}
}

func (d *dynasty) raiseAnyLevel1AptitudeBy1() {
	validToRaise := []string{}
	for _, apt := range listAptitudes() {
		if val, ok := d.aptitudes[apt]; !ok {
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
		d.aptitudes[val] = 2
	}
}

func (d *dynasty) randomAction(act ...string) int {
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
