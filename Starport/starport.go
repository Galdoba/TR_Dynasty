package starport

import (
	"strconv"

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
)

//Starport -
type Starport struct {
	sType    string
	quality  int
	yards    int
	repairs  int
	downport int
	highport int
	tl       int
	bases    []string
}

//Planet - Штука которая может получить UWP
type Planet interface {
	Bases() []string
	UWP() string
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
	switch spCode {
	case "A":
		sp.quality = qualityExellent
		sp.yards = yardStarships
		sp.repairs = repairsOverhaul
		sp.downport = downportYes
		if TrvCore.EhexToDigit(popsCode) >= 7 {
			sp.highport = highportYes
		}
	case "B":
		sp.quality = qualityGood
		sp.yards = yardSpacecraft
		sp.repairs = repairsOverhaul
		sp.downport = downportYes
		if TrvCore.EhexToDigit(popsCode) >= 8 {
			sp.highport = highportYes
		}
	case "C":
		sp.quality = qualityRoutine
		sp.repairs = repairsMajor
		sp.downport = downportYes
		if TrvCore.EhexToDigit(popsCode) >= 9 {
			sp.highport = highportYes
		}
	case "D":
		sp.quality = qualityPoor
		sp.repairs = repairsMinor
		sp.downport = downportYes
	case "E":
		sp.quality = qualityFrontier
		sp.downport = downportBeacon

	}

	return sp
}

func (sp Starport) String() string {
	str := ""

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

//WaitingTime - test
func WaitingTime(i int) string {
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
