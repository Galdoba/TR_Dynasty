package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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
)

func Test() {
	for i := 0; i < 100; i++ {
		d := NewDynasty("")
		fmt.Println(i, d.powerBase)
		fmt.Println(d)
		fmt.Println(d.Info())
	}
	d := NewDynasty("")
	fmt.Print(d)
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
	managment       string
}

/*
1. Roll Characteristics and determine Characteristic modifiers.
2.  a. Choose a Power Base.
    b. Gain Trait and Aptitude Modifiers.
3.  a. Choose a Dynasty Archetype.
	b. Determine Base Traits and Aptitudes.
	c. Gain First Generation bonuses.
	d. Determine Dynasty Boons and Hinders.
4.  a. Determine Management Assets.
	b. Gain Management Asset Benefits.
5.  Calculate First Generation Values.
6. Move on to the Background and Historic Events process.

*/

func NewDynasty(name string) dynasty {
	d := dynasty{}
	seed := utils.SeedFromString(name)

	if name == "" {
		seed = time.Now().UnixNano()
		name = "[RANDOM NAME]" + strconv.Itoa(int(seed))
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
	d.applyPowerBase("")
	//3
	vArch := validArchetypes(d)
	for len(vArch) == 0 {
		d = NewDynasty("")
		vArch = validArchetypes(d)
	}
	a := dice.Roll("1d" + strconv.Itoa(len(vArch))).DM(-1).Sum()
	d.applyArchetypeBase(vArch[a])
	//
	d.chooseBoons()
	return d
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
	st := d.name + "\n"
	st += "POWER BASE: " + d.powerBase + "\n"
	st += "ARCHETYPE: " + d.archetype + "\n"
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

func (d *dynasty) applyPowerBase(pb string) error {
	if d.powerBase != "" {
		return errors.New("Power Base already chosen")
	}
	if pb == "" {
		i := dice.Roll("1d10").DM(-1).Sum()
		pb = listPowerBase()[i]
	}
	switch pb {
	default:
		return errors.New("Unknown Power Base '" + pb + "'")
	case ColonySettlement:
		d.traits[Culture]++
		d.traits[TerritorialDfnce]--
		d.aptitudes[Expression]++
		d.aptitudes[Recruit]++
		d.aptitudes[Maintenance]++
		d.aptitudes[Propaganda]++
		d.aptitudes[Tutelage]++
	case ConflictZone:
		d.traits[TerritorialDfnce]++
		d.traits[TerritorialDfnce]++
		d.traits[FiscalDfnce]--
		d.traits[Fleet]--
		d.aptitudes[Hostility]++
		d.aptitudes[Hostility]++
		d.aptitudes[Posturing]++
		d.aptitudes[Security]++
		d.aptitudes[Tactical]++
	case Megalopolis:
		d.traits[FiscalDfnce]++
		d.traits[Technology]++
		d.traits[Culture]--
		d.traits[Culture]--
		d.aptitudes[Bureaucracy]++
		d.aptitudes[Bureaucracy]++
		d.aptitudes[Economics]++
		d.aptitudes[PublicRelations]++
		d.aptitudes[Research]++
	case MilitaryCompound:
		d.traits[TerritorialDfnce]++
		d.traits[TerritorialDfnce]++
		d.traits[Fleet]++
		d.traits[FiscalDfnce]--
		d.traits[FiscalDfnce]--
		d.aptitudes[Conquest]++
		d.aptitudes[Conquest]++
		d.aptitudes[Tactical]++
		d.aptitudes[Tactical]++
		d.aptitudes[Politics]++
		d.aptitudes[Posturing]++
		d.aptitudes[Security]++
	case NobleEstate:
		d.traits[Culture]++
		d.traits[FiscalDfnce]++
		d.traits[TerritorialDfnce]--
		d.traits[TerritorialDfnce]--
		d.traits[Fleet]--
		d.aptitudes[Bureaucracy]++
		d.aptitudes[Bureaucracy]++
		d.aptitudes[Politics]++
		d.aptitudes[Politics]++
		d.aptitudes[Expression]++
		d.aptitudes[Posturing]++
		d.aptitudes[Security]++
	case StarshipFlotilla:
		d.traits[Fleet]++
		d.traits[Fleet]++
		d.traits[Technology]++
		d.traits[TerritorialDfnce]--
		d.traits[TerritorialDfnce]--
		d.aptitudes[Intel]++
		d.aptitudes[Intel]++
		d.aptitudes[Conquest]++
		d.aptitudes[Economics]++
		d.aptitudes[Maintenance]++
		d.aptitudes[Posturing]++
		d.aptitudes[Research]++
		d.aptitudes[Tactical]++
	case TempleHolyLand:
		d.traits[Culture]++
		d.traits[Culture]++
		d.traits[Technology]--
		d.traits[Technology]--
		d.aptitudes[Expression]++
		d.aptitudes[Expression]++
		d.aptitudes[Recruit]++
		d.aptitudes[Recruit]++
		d.aptitudes[Maintenance]++
		d.aptitudes[Propaganda]++
		d.aptitudes[PublicRelations]++
		d.aptitudes[Tutelage]++
	case UnchartedWilderness:
		d.traits[TerritorialDfnce]++
		d.traits[Technology]--
		d.aptitudes[Security]++
		d.aptitudes[Security]++
		d.aptitudes[Entertain]++
		d.aptitudes[Illicit]++
		d.aptitudes[Security]++ //или Conquest/Hostility?
	case UnderworldSlum:
		d.traits[FiscalDfnce]++
		d.traits[TerritorialDfnce]++
		d.traits[Culture]--
		d.traits[Culture]--
		d.aptitudes[Illicit]++
		d.aptitudes[Illicit]++
		d.aptitudes[Sabotage]++
		d.aptitudes[Sabotage]++
		d.aptitudes[Entertain]++
		d.aptitudes[Intel]++
		d.aptitudes[Posturing]++
		d.aptitudes[Security]++
	case UrbanOffices:
		d.traits[Culture]++
		d.traits[FiscalDfnce]++
		d.traits[Fleet]--
		d.aptitudes[Acquisition]++
		d.aptitudes[Acquisition]++
		d.aptitudes[Economics]++
		d.aptitudes[Economics]++
		d.aptitudes[Bureaucracy]++
		d.aptitudes[Intel]++
		d.aptitudes[PublicRelations]++
		d.aptitudes[Tutelage]++
	}
	d.powerBase = pb
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

func (d *dynasty) applyArchetypeBase(arch string) error {
	if d.archetype != "" {
		return errors.New("Archetype already chosen")
	}
	if arch == "" {
		i := dice.Roll("1d7").DM(-1).Sum()
		arch = listArchetypes()[i]
	}
	switch arch {
	default:
		return nil
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
	d.archetype = arch
	return nil
}

func (d *dynasty) ensureAptitude(apt string, val int) {
	if v, ok := d.aptitudes[apt]; ok {
		d.aptitudes[apt] = v + val
		return
	}
	d.aptitudes[apt] = val
}

func listBoons(d dynasty) []string {
	switch d.archetype {
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
		selected = utils.AppendUniqueStr(selected, utils.RandomFromList(listBoons(*d)))
	}
	d.boonsHinders = selected
}
