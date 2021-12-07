package operations

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	FLANK_SPEED = iota
	RAPID_TRANSIT
	CURSORY_EXPLORATION
	DETAILED_EXPLORATION
)

type abstractTransit struct {
	name           string
	rate           int
	event          string
	poi            string //Point of Interest
	parsecsCovered int
	ready          bool
	eventTN        int
	poiTN          int
	eventAvoided   bool
	poiDetected    bool
}

func NewAbstractTransit(rate int, name string) (*abstractTransit, error) {
	at := abstractTransit{}
	at.name = name
	switch rate {
	default:
		return nil, fmt.Errorf("invalid 'rate of advance' input (%v)", rate)
	case FLANK_SPEED:
		at.parsecsCovered = 6
		at.eventTN = 10
		at.poiTN = 12
	case RAPID_TRANSIT:
		at.parsecsCovered = 4
		at.eventTN = 8
		at.poiTN = 10
	case CURSORY_EXPLORATION:
		at.parsecsCovered = 2
		at.eventTN = 6
		at.poiTN = 8
	case DETAILED_EXPLORATION:
		at.parsecsCovered = 0
		at.eventTN = 4
		at.poiTN = 6
	}
	at.parsecsCovered += dice.Roll1D()
	at.event = rollForEvent(dice.Roll2D())
	at.poi = rollForPoi(dice.Roll2D(), dice.Roll1D())
	return &at, nil
}

func rollForEvent(r1 int) string {
	switch r1 {
	default:
		return ""
	case 2:
		return "Major Supply Problem"
	case 3:
		return "Major Crew Problem"
	case 4:
		return "Bad Data"
	case 5:
		return "Cargo Problem"
	case 6:
		return "Minor Crew Problem"
	case 7:
		return "Minor Supply Problem"
	case 8:
		return "Crewmember Taken Ill"
	case 9:
		return "Non-Critical System Malfunction"
	case 10:
		return "Critical System Malfunction"
	case 11:
		return "Non-Critical System Breakdown"
	case 12:
		return "Critical System Breakdown"
	}
}

func rollForPoi(r1, r2 int) string {
	switch r1 {
	default:
		return ""
	case 2:
		text := "Anomaly: "
		switch r2 {
		case 1:
			text += "Windfall"
		case 2:
			text += "Major Anomaly"
		case 3, 4, 5, 6:
			text += "Minor Anomaly"
		}
		return text
	case 3:
		text := "Stellar Body: "
		switch r2 {
		case 1:
			text += "Highly unusual stellar body type"
		case 2, 3:
			text += "Unusual stellar body type"
		case 4, 5, 6:
			text += "Unusual stellar body characteristics"
		}
		return text
	case 4:
		text := "System Composition: "
		switch r2 {
		case 1:
			text += "Multi-star system"
		case 2:
			text += "Large Companion System"
		case 3:
			text += "Unusual distributions of bodies"
		case 4:
			text += "Unusual compositions of bodies"
		case 5:
			text += "Unusual orbital characteristics"
		case 6:
			text += "Multiple habitable-zone bodies"
		}
		return text
	case 5:
		text := "Rogue Bodies: "
		switch r2 {
		case 1, 2:
			text += "System is in process of ejecting multiple bodies"
		case 3, 4:
			text += "System is in process of capturing of major body such as gas gigant with moons"
		case 5, 6:
			text += "A rogue body is the process of passing through the central system, causing disruption"

		}
		return text
	case 6, 7, 8:
		return "An interesting or impressive but mundane phenominon exist in the system"
	case 9:
		text := "Mainworld: "
		switch r2 {
		case 1:
			text += "Paradise world"
		case 2:
			text += "Habitable world"
		case 3:
			text += "Hell world"
		case 4:
			text += "Unusual ecosphere"
		case 5:
			text += "Unusual temporary conditions"
		case 6:
			text += "Unusual permanent conditions"
		}
		return text
	case 10, 11:
		text := "Outersystem World: "
		switch r2 {
		case 1:
			text += "Unusual orbital path"
		case 2:
			text += "Binary planet"
		case 3:
			text += "Complex moon system"
		case 4:
			text += "High radiation"
		case 5:
			text += "Life of an unusual sort present"
		case 6:
			text += "High-gravity super-earth"
		}
		return text
	case 12:
		text := "Encounter: "
		switch r2 {
		case 1, 2:
			text += "Ruins"
		case 3, 4:
			text += "Intelligent Beings"
		case 5:
			text += "Transmission"
		case 6:
			text += "Sighting"
		}
		return text
	}
}

type Resolver interface {
	TaskDM() int
	Resolve(...string) int
	AddEntry(string)
}

func (at *abstractTransit) ExecuteBy(r Resolver) string {
	report := "No formed"
	bonus := 0
	if r.TaskDM() > 1 {
		bonus = r.TaskDM()
	}
	bonuses := []int{}
	for b := 0; b < bonus; b++ {
		bonuses = append(bonuses, dice.Roll("1d3").Sum())
	}
	for _, b := range bonuses {
		switch b {
		case 1:
			at.eventTN--
		case 2:
			at.poiTN--
		case 3:
			at.parsecsCovered++
		}
	}
	if r.Resolve() >= at.eventTN {
		r.AddEntry("EVENT AVOIDED: " + at.event)
		at.eventAvoided = true
	}
	if r.Resolve() >= at.poiTN {
		r.AddEntry("POINT OF INTEREST DETECTED: " + at.poi)
		at.poiDetected = true
	}
	r.AddEntry(fmt.Sprintf("Transit concluded covering %v parsecs", at.parsecsCovered))
	at.ready = true
	return report
}

func (at *abstractTransit) String() string {
	rep := at.name
	if !at.ready {
		rep += "\nNot Ready"
		rep += fmt.Sprintf("\nEvent TN = %v", at.eventTN)
		rep += fmt.Sprintf("\nPOI TN = %v", at.poiTN)
		rep += fmt.Sprintf("\nLenght = %v", at.parsecsCovered)
		return rep
	}
	if !at.eventAvoided {
		rep += "\nEvent: " + at.event
	}
	if at.poiDetected {
		rep += "\nPoint of Interest: " + at.poi
	}
	rep += fmt.Sprintf("\n%v parsecs covered", at.parsecsCovered)
	return rep
}
