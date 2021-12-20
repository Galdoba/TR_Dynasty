package cei

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	Interval_Initial          = 10
	Interval_Standard         = 6
	Interval_Stressful        = 4
	Interval_Highly_Stressful = 2
	Accomodation_Generous     = 1
	Accomodation_Cramped      = -1
	Supplies_Reduced          = -1
	Supplies_Luxuries         = 2
	Fresh                     = 0
	Fatigued                  = 1
	HighlyFatigued            = 2
	DangerouslyFatigued       = 3
	Exhaused                  = 4
	Incapable                 = 5
)

type fatigue struct {
	index        int //Crew Fatigue Index
	fatigueState int
	daysToCheck  int //
	status       string
}

func (ftg *fatigue) newInterval(interval int, factors ...int) {
	diceCode := fmt.Sprintf("%vd6", interval+sumOf(factors))
	ftg.daysToCheck = dice.New().RollNext(diceCode).Sum()
}

func (ftg *fatigue) update() bool {
	if ftg.daysToCheck > 0 {
		ftg.daysToCheck--
		return false
	}
	ftg.index++
	if dice.Roll2D() < ftg.index {
		ftg.fatigueState++
		ftg.index = 0
		switch ftg.fatigueState {
		case Fresh:
			ftg.status = EVENT_FatigueStatus_Fresh
		case Fatigued:
			ftg.status = EVENT_FatigueStatus_Fatigued
		case HighlyFatigued:
			ftg.status = EVENT_FatigueStatus_HighlyFatigued
		case DangerouslyFatigued:
			ftg.status = EVENT_FatigueStatus_DangerouslyFatigued
		case Exhaused:
			ftg.status = EVENT_FatigueStatus_Exhausted
		case Incapable:
			ftg.status = EVENT_FatigueStatus_Incapable
		}
		return true
	}
	ftg.newInterval(Interval_Standard)
	ftg.check()
	return false
}

func (ftg *fatigue) check() {
	if ftg.fatigueState < Fresh {
		ftg.fatigueState = Fresh
	}
	if ftg.fatigueState > Incapable {
		ftg.fatigueState = Incapable
	}
	if ftg.index < 0 {
		ftg.index = 0
	}
	if ftg.daysToCheck < 0 {
		ftg.daysToCheck = 0
	}
}

func sumOf(intSl []int) int {
	s := 0
	for _, v := range intSl {
		s += v
	}
	return s
}

func (ftg *fatigue) State() string {
	switch ftg.fatigueState {
	default:
		return "Unknown"
	case Fresh:
		return EVENT_FatigueStatus_Fresh
	case Fatigued:
		return EVENT_FatigueStatus_Fatigued
	case HighlyFatigued:
		return EVENT_FatigueStatus_HighlyFatigued
	case DangerouslyFatigued:
		return EVENT_FatigueStatus_DangerouslyFatigued
	case Exhaused:
		return EVENT_FatigueStatus_Exhausted
	case Incapable:
		return EVENT_FatigueStatus_Incapable
	}
}
