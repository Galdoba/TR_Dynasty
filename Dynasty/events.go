package dynasty

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/DateManager"
	"github.com/Galdoba/TR_Dynasty/dice"
)

/*
событие:		// умеет менять данные
время старта	// данные
длительность	// данные
субъект			// умеет Выбирать объект и Решать
объект			// данные
имя				// данные
описание		// данные
метод решения	// данные - скорее функция
исход			// данные



*/

type chalenge struct {
	name       string
	descr      string
	roll       []string
	reward     string
	punishment string
}

type ChangeReceiver interface {
	ChangeStats(...string)
}

// func (d *Dynasty) ChangeStats(changes ...string) {
// 	rpa := parceChanges(changes...)
// 	for i, change := range rpa {

// 	}
// }

// func (d *Dynasty) applyChange(rp rewardPunishment) {
// 	for i, val := range listALL() {

// 	}
// }

type rewardPunishment struct {
	fieldName string
	changeby  int
	err       error
}

func Test3() {
	d := NewDynasty("Test Dynasty")
	rpa := parceChanges("-2 Popularity", "+1 AnyTrait", "+6 AnyAptitude", "-1 Hostility", "-1 Illicit")
	fmt.Println(d.Info())
	for _, rp := range rpa {
		d.changeStatBy(rp.fieldName, rp.changeby)
	}
	fmt.Println(d.Info())
}

func parceChanges(changes ...string) []rewardPunishment {
	rpa := []rewardPunishment{}
	entry := 1
	start := time.Now().UnixNano()
	for _, descr := range changes {
		fmt.Println(entry, "Go ", descr, time.Now().UnixNano()-start)
		entry++
		rp := rewardPunishment{}
		rp.changeby, rp.err = intFromDescr(descr)
		if rp.err != nil {
			continue
		}
		for _, field := range listALL() {
			if strings.Contains(descr, "AnyCharacteristic") {
				trait := dice.New(time.Now().UnixNano()).RollFromList(listCharacteristics())
				rp.fieldName = trait
				rpa = append(rpa, rp)
				break
			}
			if strings.Contains(descr, "AnyTrait") {
				trait := dice.New(time.Now().UnixNano()).RollFromList(listTraits())
				rp.fieldName = trait
				rpa = append(rpa, rp)
				break
			}
			if strings.Contains(descr, "AnyAptitude") {
				trait := dice.New(time.Now().UnixNano()).RollFromList(listAptitudes())
				rp.fieldName = trait
				rpa = append(rpa, rp)
				break
			}
			fmt.Println(entry, "check: ", descr, "|", field, time.Now().UnixNano()-start)
			entry++
			if strings.Contains(descr, field) {
				fmt.Println(entry, "FOUND!!! ", field, time.Now().UnixNano()-start)
				entry++
				rp.fieldName = field
				rpa = append(rpa, rp)
				break
			}

		}
	}
	fmt.Println("rpa:")
	for i := range rpa {
		fmt.Println(rpa[i])
	}

	return rpa
}

func intFromDescr(descr string) (int, error) {
	data := strings.Split(descr, " ")
	for _, str := range data {
		res, err := strconv.Atoi(str)
		if err == nil {
			return res, err
		}
	}
	return 0, errors.New("Parce failed")
}

type event struct {
	date        string
	lenght      string
	name        string
	description string
	outcome     string
	rollDescr   string
	rollEffect  int
}

func (d *Dynasty) BackGroundEvent(code string) {
	switch d.archetype {
	case Conglomerate:
		d.bgEventConglomerate(code)
	}
}

func (d *Dynasty) HistoricEvent() event {
	//story string
	//changes string
	//rollDescr string
	ev := event{}
	r := d.dicepool.RollNext("2d6").DM(d.historicEvents).Sum()
	d.historicEvents++
	switch r {
	default:
		ev.name = "Antient Visitor"
		ev.description = "A truly ancient and powerful being with abilities bordering on ‘magic’ comes forward to lend aid to the Dynasty for its own mysterious reasons."
		switch d.dicepool.RollNext("1d6").Sum() {
		case 1:
			ev.outcome = "+1 to all Traits"
			d.increaseAllTraits()
		case 2:
			ev.outcome = "+1 to all Values"
			d.increaseAllValues()
		case 3:
			ev.outcome = "+1 to "
			apts := []string{}
			for len(apts) < 3 {
				apts = append(apts, d.dicepool.RollFromList(listAptitudes()))
			}
			for _, val := range apts {
				d.raiseApttitude(val)
				ev.outcome += val + ", "
			}
			ev.outcome = strings.TrimSuffix(ev.outcome, ", ")
		case 4:
			ev.outcome = "+1 to "
			shr := []string{}
			for len(shr) < 1 {
				shr = append(shr, d.dicepool.RollFromList(listCharacteristics()))
			}
			for _, val := range shr {
				d.stat[val]++
				ev.outcome += val + ", "
			}
			ev.outcome = strings.TrimSuffix(ev.outcome, ", ")
		case 5:
			ev.outcome = "+2 to "
			shr := []string{}
			for len(shr) < 2 {
				shr = append(shr, d.dicepool.RollFromList(listCharacteristics()))
			}
			for _, val := range shr {
				d.stat[val]++
				ev.outcome += val + ", "
			}
			ev.outcome = strings.TrimSuffix(ev.outcome, ", ")
		case 6:
			ev.outcome = "+3 to "
			shr := []string{}
			for len(shr) < 3 {
				shr = append(shr, d.dicepool.RollFromList(listCharacteristics()))
			}
			for _, val := range shr {
				d.stat[val]++
				ev.outcome += val + ", "
			}
			ev.outcome = strings.TrimSuffix(ev.outcome, ", ")
		}
	case 2:
		ev.name = "War of the Worlds!"
		ev.description = "There is an interstellar war between planetary forces, sweeping them into the dangerous realm of battles and destruction."
		apts := []string{Conquest, Hostility, Security}
		for _, val := range apts {
			if d.rollAptitude(val) >= 8 {
				d.increaseAllTraits()
				ev.outcome += "Increase All Traits by 1\n"
			} else {
				d.decreaseAllValues()
				ev.outcome += "Decrease All Values by 1\n"
			}
		}
	case 3, 4:
		ev.name = "Foes on all Sides"
		ev.description = "A consolidation of enemies have targeted the Dynasty and are coming at them from all directions."
		apts := []string{Intel, Posturing, Security}
		for _, val := range apts {
			if d.rollAptitude(val) >= 8 {
				d.increaseAllValues()
				ev.outcome += "Increase All Values by 1\n"
			} else {
				d.stat[FiscalDfnce]--
				d.stat[TerritorialDfnce]--
				d.stat[Fleet]--
				ev.outcome += "Decrease All Fiscal Defence, Territorial Defence and Fleet by 1\n"
			}
		}
	case 5, 6:
		ev.name = "An Unlikely Hero Rises"
		ev.description = "One of the Dynasty’s inner members is given an opportunity to do something truly amazing – and does."
		chr := []string{}
		for len(chr) < 2 {
			chr = append(chr, d.dicepool.RollFromList(listCharacteristics()))
		}
		for _, val := range chr {
			d.stat[val]++
			ev.outcome += "Raise " + val + " by 1"
		}
		d.stat[Morale]++
		ev.outcome += "Raise Morale by 1"
	}
	return ev
}

func (d *Dynasty) bgEventConglomerate(code string) string {
	ev := ""
	switch code {
	case "14", "24", "34", "44", "54", "64":
		ev = "Historic Event – Roll on the Dynasty Historic Event Table. "
		stage := d.HistoricEvent()
		fmt.Println(stage)
		//
	case "11":
		ev = "Stocks are falling all over the galaxy for years; roll Greed 8+ or lose 1 point of Wealth. "
		if d.rollCharacteristic(Grd) < 8 {
			d.stat[Wealth]--
		}
	case "12":
		ev = "Scandal rocks the shareholders’ memo meetings and prices hit an all time low; roll Loyalty 8+ or lose 1 point of Morale. "
	case "13":
		ev = "Hostile takeovers try to devour the weak initial generation; roll Bureaucracy or Tenacity 8+ or lose 1 point of Wealth. "
	case "15":
		ev = "New ideas on the market test the Conglomerate’s ingenuity and adaptability; roll Economics or Research 7+ or lose 1 point of Fiscal Defence. "
	case "16":
		ev = "A rival has been moving in on your workers, roll Security 8+ or lose 1 point of Populace. "
	case "21":
		ev = "Heavy competition in the interplanetary market has really toughened things up around the power base; roll 1d6: 1–4, Gain +1 Territorial Defence; 5–6: Gain +1 Tenacity. "
	case "22":
		ev = "A massive media event provides management with a chance to make a name for itself; Gain +1 Popularity or +2 Morale. "
	case "23":
		ev = "Big business is good business these days; roll Bureaucracy 7+ to gain one Level in Wealth. "
	case "25":
		ev = "The territories are wearing your logo and there are not many locals who do not know your name. Gain +1 Popularity. "
	case "26":
		ev = "A major coup in the local government risks sweeping in the Conglomerate. Join in and roll Conquest 8+ to help the new regime. Avoid the conflict and roll Security 8+ to keep out of the line of fire. Succeed in either roll and gain +1 to any Value; fail and lose 1 point of Loyalty and Militarism. "
	case "31":
		ev = "Labour unions are not happy about the solidification of the management entities through the Conglomerate. Roll Propaganda and Security 7+; succeed in both and gain +1 to the Characteristic of the Player’s choice. "
	case "32":
		ev = "The management of the Conglomerate are contacted by tremendously powerful alien benefactors; Gain +1 Bureaucracy, Expression or Recruit. "
	case "33":
		ev = "Everything goes as planned for decades; add +1 to any Trait or Value. "
	case "35":
		ev = "The power base suffers a major natural disaster and the Conglomerate can lend charitable aid; you may spend 1 point of Wealth to increase Popularity by +1. "
	case "36":
		ev = "A sickness plagues the population and the workforce, putting the Conglomerate at risk but giving them a good idea to back medical resources. Roll Acquisition 8+ to gain 1 point of any Value. "
	case "41":
		ev = "A powerful client puts the Conglomerate through a vicious courtroom drama that lasts months, if not years; roll Politics or Security 8+ to avoid losing 1 point of Wealth. "
	case "42":
		ev = "A university grant is created in the Conglomerate’s honour; Gain +1 Popularity. "
	case "43":
		ev = "Everything goes as planned for decades; add +1 to any Aptitude or Trait. "
	case "45":
		ev = "High-credit gambling establishments become not only legal but encouraged among big businesses; Roll Illicit 7+ to gain +1 Wealth. "
	case "46":
		ev = "Primitives are in great supply to be exploited. If the Conglomerate treats them with respect, it gains +1 Popularity. If it uses them harshly, gain +1 Wealth and +1 Populace. "
	case "51":
		ev = "Industrial sabotage is rumoured to be targeting the Conglomerate; roll Security 8+ to protect itself, gaining +1 Territorial Defence. "
	case "52":
		ev = "Advanced aliens have chosen the Conglomerate to fabricate their devices, adding their tooling to their own; Roll Maintenance 8+ to gain +1 Technology. "
	case "53":
		ev = "Everything goes as planned for decades; add +1 to any Aptitude or Value. "
	case "55":
		ev = "War profiteers are looking to launder their ill-gotten gains through the Conglomerate; you may spend 1 point of Loyalty to gain 1 point of Scheming before rolling Illicit 9+; succeed in the Aptitude check to gain 1d6-4 Wealth (minimum of 1). "
	case "56":
		ev = "A celebrity enjoys associating on a business level with the Conglomerate; gain +1 Morale or blackmail the Celebrity with Illicit 8+ to gain +1 Wealth and +1 Scheming. "
	case "61":
		ev = "The government names a holiday after the Conglomerate’s founder(s); Gain +1 Loyalty, Popularity or Tradition. "
	case "62":
		ev = "An interstellar sports team needs a sponsor right before a major multi-planet tournament; buy the team by spending 1 point of Wealth, gaining +1 Culture and +1 Morale. "
	case "63":
		ev = "Things could not go any better for the Dynasty; add +1 to any Characteristic, Aptitude, Trait or Value. "
	case "65":
		ev = "An unexpected territory shift puts a new planet in the Conglomerate’s control, adding +1 to all Values. "
	case "66":
		ev = "A formerly powerful Conglomerate folds, leaving its resources and assets for the new one to claim unchallenged; Gain +1 to any one Characteristic and +1 to any two Aptitudes."

	}
	return ev
}

func Survived(d Dynasty) bool {
	for _, val := range listValues() {
		if d.stat[val] < 1 {
			//fmt.Println("Dynasty have crumbled and is no more...")
			return false
		}
	}
	zeroTraits := []string{}
	for _, val := range listTraits() {
		if d.stat[val] < 1 {
			zeroTraits = append(zeroTraits, val)
		}
	}
	if len(zeroTraits) > 1 {
		//fmt.Println("Dynasty is too weak to defend itself from the normal dangers it would face and is swiftly torn asunder by rivals...")
		return false
	}
	vitalChars := []string{Lty, Pop, Tra}
	for _, val := range vitalChars {
		if d.stat[val] < 1 {
			//	fmt.Println("Dynasty members riot and rise up from within, destroying the Dynasty’s power base until it cannot stand on its own...")
			return false
		}
	}
	return true
}

func (d *Dynasty) rollCharacteristic(chr string) int {
	dm := DM(d.stat[chr])
	return d.dicepool.RollNext("2d6").DM(dm).Sum()
}

func (d *Dynasty) rollAptitude(apt string) int {
	dm := -2
	if val, ok := d.stat[apt]; ok {
		dm = val
	}
	return d.dicepool.RollNext("2d6").DM(dm).Sum()
}

func (d *Dynasty) increaseAllTraits() {
	for _, val := range listTraits() {
		d.stat[val]++
	}
}

func (d *Dynasty) decreaseAllTraits() {
	for _, val := range listTraits() {
		d.stat[val]--
	}
}

func (d *Dynasty) increaseAllValues() {
	for _, val := range listValues() {
		d.stat[val]++
	}
}

func (d *Dynasty) decreaseAllValues() {
	for _, val := range listValues() {
		d.stat[val]--
	}
}

func NewEvent(name string) event {
	ev := event{}
	ev.name = name
	return ev
}

func (ev *event) SetDescription(descr string) {
	ev.description = descr
}

type EventTracker interface {
	TrackEvent(day int) bool
}

func (d *Dynasty) TrackEvent(currentDay int) bool {
	if currentDay < d.nextActionDay {
		return false
	}
	//d.nextActionDay = currentDay + dice.Roll("2d6").DM(28).Sum()
	fmt.Println("Date: ", DateManager.FormatToDate(currentDay))
	fmt.Println("Launch Action until:", DateManager.FormatToDate(d.nextActionDay))
	fmt.Println("----------------------")
	return true
}

const (
	DifficultyVeryEasy      = 6
	DifficultyEasy          = 4
	DifficultyRoutine       = 2
	DifficultyAverage       = 0
	DifficultyDifficult     = -2
	DifficultyVeryDifficult = -4
	DifficultyFormidable    = -6
)

type apttAction struct {
	name          string
	sourceApt     string
	sourceChr     []string
	opposedBy     []string
	difficulty    int
	source        Dynasty
	target        Dynasty
	startDay      int
	conclusionDay int
}

func InitiateAction(source, target *Dynasty, name string, currentDay int) {
	aAct := apttAction{}
	aAct.name = name
	switch name {
	default:
		fmt.Println("TODO: ", name)
		//return aAct
	case "Claiming neutral territory or resources":
		aAct.difficulty = DifficultyVeryEasy
		aAct.startDay = currentDay
		aAct.conclusionDay = currentDay + dice.Roll1D()*months()
		aAct.sourceApt = Conquest
		aAct.sourceChr = []string{Grd, Mil}
		aAct.source = *source
		aAct.target = *source
		//return aAct

	}
	source.nextActionDay = aAct.conclusionDay
	fmt.Println(aAct)
	fmt.Println(source.name, " DO ", name, " AGAINST ", target.name)
	fmt.Println("EFFECT NULL")
}

func months() int {
	return 31 + dice.Flux()
}

func EventMap(name string) map[string]func(Dynasty) *Dynasty {
	evmap := make(map[string]func(Dynasty) *Dynasty)
	////GENERATION GOALS
	evmap["Acquire Ancient Technology|SUCCESS"] = func(d Dynasty) *Dynasty {
		for i := 0; i < d.dicepool.RollNext("1d6").Sum(); i++ {
			val := d.dicepool.RollFromList([]string{Technology, Wealth})
			if i%2 == 1 {
				val = Technology
			}
			d.changeStatBy(val, 1)
		}
		return &d
	}
	evmap["Acquire Ancient Technology|FAILURE"] = func(d Dynasty) *Dynasty {
		for i := 0; i < d.dicepool.RollNext("1d6").Sum()+1; i++ {
			val := d.dicepool.RollFromList(listValues())
			if i%2 == 1 {
				val = Wealth
			}
			d.changeStatBy(val, -1)
		}
		return &d
	}

	evmap["Banish an Enemy|SUCCESS"] = func(d Dynasty) *Dynasty {
		d.changeStatBy(d.anyCharacteristic(), 1)
		d.changeStatBy(d.anyCharacteristic(), 1)
		for i := 0; i < d.dicepool.RollNext("1d6").Sum(); i++ {
			val := d.dicepool.RollFromList(listTraits())
			d.changeStatBy(val, 1)
		}
		return &d
	}
	evmap["Banish an Enemy|FAILURE"] = func(d Dynasty) *Dynasty {
		for i := 0; i < d.dicepool.RollNext("1d6").Sum(); i++ {
			val := d.anyCharacteristic()
			if i == 1 {
				val = Lty
			}
			d.changeStatBy(val, -1)
		}
		return &d
	}

	return evmap
}
