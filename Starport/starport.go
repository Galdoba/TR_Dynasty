package starport

import (
	"strconv"

	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/world"
)

const (
	qualityNone     = 0
	qualityBasic    = 1
	qualityFrontier = 2
	qualityPoor     = 3
	qualityRoutine  = 4
	qualityGood     = 5
	qualityExellent = 6
	yardNo          = 0
	yardSpacecraft  = 1
	yardStarships   = 2
	repairsNo       = 0
	repairsMinor    = 1
	repairsMajor    = 2
	repairsOverhaul = 3
	downportNo      = 0
	downportBeacon  = 1
	downportYes     = 2
	highportNo      = 0
	highportYes     = 1
	dtStarport      = world.DataTypeStarport
	dtSize          = world.DataTypeSize
	dtAtmosphere    = world.DataTypeAtmosphere
	dtHydrosphere   = world.DataTypeHydrosphere
	dtPopulation    = world.DataTypePopulation
	dtGoverment     = world.DataTypeGoverment
	dtLaws          = world.DataTypeLaws
	dtTechLevel     = world.DataTypeTechLevel
	//ServiceBerthing -
	ServiceBerthing = 0
	//ServiceRefuiling -
	ServiceRefuiling = 1
	//ServiceWarehousing -
	ServiceWarehousing = 2
	//ServiceHazmat -
	ServiceHazmat = 3
	//ServiceRepairs -
	ServiceRepairs = 4
)

//Starport -
type Starport struct {
	sType     string
	quality   int
	yards     int
	repairs   int
	downport  int
	highport  int
	tl        int
	bases     []string
	dice      *dice.Dicepool //а нужен ли?
	modules   map[string]bool
	serviceDM [5]int
	governor  string
}

/*
РУТИНА:
1. Проверка очереди на стыковку.				 (Rand)
2. Проверка Law Interraction (Check) 			 (Rand)
3. Проверка фонового события (General) 8+		 (Rand)
4. Проверка фонового события (Significant) 11+	 (Rand)
5. Создание местного Непися						 (Rand)
6. Создание местного Губернатора				 (Perma).
6. Создание отличительной черты для Космопорта 	 (Perma).
*/

//Planet - Штука которая может получить UWP
type Planet interface {
	Bases() []string
	UWP() string
	Name() string
}

//From - создает старпорт и детали от планеты
func From(planet Planet) Starport {
	sp := Starport{}
	uwp, err := profile.NewUWP(planet.UWP())
	if err != nil {
		panic(err.Error())
	}
	spCode := uwp.Starport()
	sp.sType = spCode
	sp.tl = TrvCore.EhexToDigit(uwp.TL())
	//sp.tl = TrvCore.EhexToDigit(uwp.DataType(dtTechLevel))
	popsCode := uwp.Pops()
	sp.dice = dice.New(utils.SeedFromString(planet.Name() + planet.UWP()))
	switch spCode {
	default:
		sp.serviceDM = [5]int{-1, 0, 1, 1, 0}
	case "A":
		sp.quality = qualityExellent
		sp.yards = yardStarships
		sp.repairs = repairsOverhaul
		sp.downport = downportYes
		if TrvCore.EhexToDigit(popsCode) >= 7 {
			sp.highport = highportYes
		}
		sp.serviceDM = [5]int{-5, -3, -3, -2, -3}
	case "B":
		sp.quality = qualityGood
		sp.yards = yardSpacecraft
		sp.repairs = repairsOverhaul
		sp.downport = downportYes
		if TrvCore.EhexToDigit(popsCode) >= 8 {
			sp.highport = highportYes
		}
		sp.serviceDM = [5]int{-4, -3, -2, -2, -3}
	case "C":
		sp.quality = qualityRoutine
		sp.repairs = repairsMajor
		sp.downport = downportYes
		if TrvCore.EhexToDigit(popsCode) >= 9 {
			sp.highport = highportYes
		}
		sp.serviceDM = [5]int{-3, -2, -2, -1, -2}
	case "D":
		sp.quality = qualityPoor
		sp.repairs = repairsMinor
		sp.downport = downportYes
		sp.serviceDM = [5]int{-2, -2, -1, -1, -2}
	case "E":
		sp.quality = qualityFrontier
		sp.downport = downportBeacon
		sp.serviceDM = [5]int{-2, -1, 0, 0, -1}
	}
	sp.rollGovernor()
	return sp
}

func (sp Starport) Info() string {
	str := ""
	// sType    string
	str += "Starport Class : " + sp.sType + "\n"
	// quality  int
	str += "       Quality : " + sp.Quality() + "\n"
	// yards    int
	str += "         Yards : " + sp.Yards() + "\n"
	// repairs  int
	str += "       Repairs : " + sp.Repairs() + "\n"
	// downport int
	str += "      Downport : " + sp.Downport() + "\n"
	// highport int
	str += "      Highport : " + sp.Highport() + "\n"
	// tl       int
	str += "            TL : " + strconv.Itoa(sp.tl) + "\n"
	// bases    []string

	return str
}

//Quality -
func (sp *Starport) Quality() string {
	q := ""
	switch sp.quality {
	case 0:
		q = "None"
	case 1:
		q = "Basic"
	case 2:
		q = "Frontier"
	case 3:
		q = "Poor"
	case 4:
		q = "Routine"
	case 5:
		q = "Good"
	case 6:
		q = "Exellent"
	}
	return q
}

//Yards -
func (sp *Starport) Yards() string {
	y := ""
	switch sp.yards {
	case 0:
		y = "No"
	case 1:
		y = "Spacecraft"
	case 2:
		y = "Starships"
	}
	return y
}

//Repairs -
func (sp *Starport) Repairs() string {
	r := ""
	switch sp.repairs {
	case 0:
		r = "No"
	case 1:
		r = "Minor"
	case 2:
		r = "Major"
	case 3:
		r = "Overhaul"
	}
	return r
}

//Downport -
func (sp *Starport) Downport() string {
	dp := ""
	switch sp.downport {
	case 0:
		dp = "No"
	case 1:
		dp = "Beacon"
	case 2:
		dp = "Yes"
	}
	return dp
}

//Highport -
func (sp *Starport) Highport() string {
	hp := ""
	switch sp.highport {
	case 0:
		hp = "No"
	case 1:
		hp = "Yes"
	}
	return hp
}

func waitingTime(i int) string {
	switch i {
	default:
		if i < 1 {
			return "Immidiatly"
		}
		num := dice.Roll("d6").Sum()
		return strconv.Itoa(num) + " days"
	case 1:
		num := dice.Roll("d6").Sum()
		return strconv.Itoa(num) + " minutes"
	case 2:
		num := dice.Roll("d6").Sum()
		return strconv.Itoa(num) + "0 minutes"
	case 3:
		return "1 Hour"
	case 4:
		num := dice.Roll("d6").Sum()
		return strconv.Itoa(num) + " hours"
	case 5:
		num := dice.Roll("2d6").Sum()
		return strconv.Itoa(num) + " hours"
	case 6:
		return "1 day"
	}
}

func (sp Starport) ServiseTime(s int) string {
	time := ""
	switch s {
	default:
		return "Unknown Service"
	case ServiceBerthing, ServiceRefuiling, ServiceWarehousing, ServiceHazmat, ServiceRepairs:
		time = waitingTime(dice.Roll("1d6").Sum() + sp.serviceDM[s])
	}
	return time
}

func (sp *Starport) rollGovernor() {
	switch sp.dice.RollNext("1d6").Sum() {
	case 1:
		sp.governor = "Hands-on: The governor constantly gets involved in the running of every aspect of port life, from the cleaning of corridors to the organisation of berthing space. He can usually be seen bustling around the port, even taking on other employee’s jobs for an hour or two."
	case 2:
		sp.governor = "Omni-Present: Huge viewing screens display images of the governor night and day, as he gazes down on his domain. Regular announcements and reminders of Starport protocols are read out in his voice."
	case 3:
		sp.governor = "Out-of-his Depth: The Governor should never have been given this post and he knows it. He constantly defers decisions and delegates responsibility to his underlings. If they are good at their jobs, this problem could persist for some time; if they are not, the port will soon run into trouble."
	case 4:
		sp.governor = "Puritanical: The Governor sees himself as the moral guardian of those passing through the port. He will not tolerate anything that might be seen as depraved or deviant."
	case 5:
		sp.governor = "Alien-ophile: Cultural and racial diversity are encouraged by an open-minded governor who sees all sentient beings as equal."
	case 6:
		sp.governor = "About to Retire: After a lifetime’s service, the governor is months away from retirement. His subordinates struggle to establish themselves as his heir-apparent, whilst he sizes them up for the post and often asks visitors for their impressions of his staff."
	}
}

func lawInteractionCheck(law string) bool {
	if dice.Roll("2d6").Sum() >= TrvCore.EhexToDigit(law) {
		return true
	}
	return false
}


func StartRoutine(sp Starport) {
	
}