package dynasty

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/task"
)

const (
	goalAcquireAncientTechnology          = "Acquire Ancient Technology"
	goalBanishAnEnemy                     = "Banish an Enemy"
	goalFulfilASuccessfulCoupDetat        = "Fulfil a Successful Coup De'tat"
	goalGrowByLeapsAndBounds              = "Grow by Leaps and Bounds"
	goalInventANewTechnologicalMarvel     = "Invent a New Technological Marvel"
	goalHoldAnInterstellarPeaceConference = "Hold an Interstellar Peace Conference"
	goalOrganiseOrderFromChaos            = "Organise Order from Chaos"
	goalStartAnInterstellarWar            = "Start an Interstellar War"
	goalTeachANewSkillToTheMasses         = "Teach a New Skill to the Masses"
	goalUtterGenocide                     = "Utter Genocide"
)

type goal struct {
	name                       string
	descr                      string
	scale                      int            //instant, 5year, 10year,30year
	cumulativeEffectNeeded     map[string]int //карта необходимых кумулятивных успехов
	roll4Needed                map[string]int
	roll5Needed                map[string]int
	roll6Needed                map[string]int
	characteristicPointsNeeded map[string]int
	traitPointsNeeded          map[string]int
	valuePointsNeeded          map[string]int
	urgency                    int
	reward                     func(Dynasty) *Dynasty
	log                        string
}

func listGoals() []string {
	return []string{
		goalAcquireAncientTechnology,
		goalBanishAnEnemy,
		goalFulfilASuccessfulCoupDetat,
		goalGrowByLeapsAndBounds,
		goalInventANewTechnologicalMarvel,
		goalHoldAnInterstellarPeaceConference,
		goalOrganiseOrderFromChaos,
		goalStartAnInterstellarWar,
		goalTeachANewSkillToTheMasses,
		goalUtterGenocide,
	}
}

func testNewGoal(urgency int) goal {
	g := goal{}
	g.urgency = urgency
	g.cumulativeEffectNeeded = make(map[string]int)
	g.roll4Needed = make(map[string]int)
	g.roll5Needed = make(map[string]int)
	g.roll6Needed = make(map[string]int)
	g.characteristicPointsNeeded = make(map[string]int)
	g.traitPointsNeeded = make(map[string]int)
	g.valuePointsNeeded = make(map[string]int)
	g.name = "test Goal 1"
	g.cumulativeEffectNeeded[Clv] = 1
	return g
}

func NewGoal(name string, scale int) goal {
	g := goal{}
	switch scale {
	default:
		g.urgency = 0
	case 1:
		g.urgency = 60 * months()
	case 2:
		g.urgency = 120 * months()
	case 3:
		g.urgency = 360 * months()
	}
	g.cumulativeEffectNeeded = make(map[string]int)
	g.roll4Needed = make(map[string]int)
	g.roll5Needed = make(map[string]int)
	g.roll6Needed = make(map[string]int)
	g.characteristicPointsNeeded = make(map[string]int)
	g.traitPointsNeeded = make(map[string]int)
	g.valuePointsNeeded = make(map[string]int)
	g.name = name
	g.setupChecklist(name)
	return g
}

func (g goal) triggered() bool {
	return true
}

func Test4() {

	dyn := NewDynasty("")

	fmt.Println(dyn.Info())
	dyn.goals = append(dyn.goals, NewGoal("Utter Genocide", 3))

	// for cd := 0; cd < 12000; cd++ {
	// 	// fmt.Print("CD: ", cd, "  \r")
	// 	// if cd < dyn.goals[0].urgency {
	// 	// 	continue
	// 	// }
	// 	// dyn.goals[0].conclude()

	// 	// fmt.Println(dyn)
	// 	// dyn.goals[0].reward(dyn)
	// 	// dyn.ensureAptValidRange()
	// 	// break
	// 	dyn.updateGoals()
	// }
	fmt.Print("  \n")
	fmt.Println(dyn.Info())
	dyn.rollCheckAptitude(Mil, Conquest, task.DifficultyAverage)
}

func (d *Dynasty) updateGoals() {
	for i := range d.goals {
		if d.goals[i].urgency > 0 {
			d.goals[i].urgency--
			continue
		}
		d.goals[i].conclude()
		d.goals[i].reward(*d)
		d.ensureAptValidRange()
	}
}

func (g *goal) conclude() {
	// eventMap := make(map[string]func(Dynasty) *Dynasty)
	// eventMap["test3"] = func(d Dynasty) *Dynasty {
	// 	d.stat[Clv] = 99
	// 	return &d
	// }
	// eventMap["test4"] = func(d Dynasty) *Dynasty {
	// 	d.stat[Clv] = -99
	// 	return &d
	// }
	key := ""
	switch g.success() {
	default:
		fmt.Println("-----------")
	case true:
		key = g.name + "|SUCCESS"
		fmt.Println("++++++++++++++++")

	case false:
		key = g.name + "|FAILURE"

		fmt.Println("=============")
	}
	//name := "test"
	if val, ok := EventMap()[key]; ok {
		fmt.Println(key)
		g.reward = val
	} else {
		fmt.Println("Null event")
		g.reward = func(d Dynasty) *Dynasty {
			return &d
		}
	}
	g = nil

}

func (g *goal) success() bool {
	for _, v := range g.cumulativeEffectNeeded {
		if v > 0 {
			return false
		}
	}
	for _, v := range g.roll4Needed {
		if v > 0 {
			return false
		}
	}
	for _, v := range g.roll5Needed {
		if v > 0 {
			return false
		}
	}
	for _, v := range g.roll6Needed {
		if v > 0 {
			return false
		}
	}
	for _, v := range g.characteristicPointsNeeded {
		if v > 0 {
			return false
		}
	}
	for _, v := range g.traitPointsNeeded {
		if v > 0 {
			return false
		}
	}
	for _, v := range g.valuePointsNeeded {
		if v > 0 {
			return false
		}
	}
	return true
}

type testStruct struct {
	value    int
	testgoal goal
}

type RewardReciever interface {
	GetRewardFor() Dynasty
}

func (d *Dynasty) GetRewardFor() Dynasty {
	newState := *d
	for _, goal := range d.goals {
		if goal.urgency > 0 {
			continue
		}
		//goal.conclude() // проверяет выполнена ли цель и решает какую дать награду
		//newState = goal.reward(newState)
	}
	return newState
	//return ts.testgoal.reward(*ts)

}

// func Test2() {
// 	ts := testStruct{3, goal{}}
// 	ts.testgoal.reward = increasetestvalue
// 	fmt.Println(ts)
// 	ts = ts.Reward()

// 	fmt.Println(ts)
// }

// func (ts *testStruct) Reward() testStruct {
// 	return ts.testgoal.reward(*ts)

// }

// func increasetestvalue(ts testStruct) testStruct {
// 	ts.value++
// 	return ts
// }

/*
Acquire Ancient Technology
Banish an Enemy
Fulfil a Successful Coup De'tat
Grow by Leaps and Bounds
Hold an Interstellar Peace Conference
Organise Order from Chaos
Start an Interstellar War
Teach a New Skill to the Masses
Utter Genocide
[NOT ESTABLESHED]



*/

func (g *goal) setupChecklist(name string) {
	switch name {
	case goalAcquireAncientTechnology:
		g.cumulativeEffectNeeded[Intel] = 20
		g.cumulativeEffectNeeded[Research] = 30
		g.roll5Needed[Conquest] = 3
		g.roll6Needed[Clv] = 2
		g.traitPointsNeeded[Technology] = -3
	case goalBanishAnEnemy:
		g.cumulativeEffectNeeded[Sabotage] = 10
		g.cumulativeEffectNeeded[Hostility] = 20
		g.cumulativeEffectNeeded[Propaganda] = 30
		g.roll5Needed[Mil] = 2
		g.valuePointsNeeded[Populace] = -3 //TODO: Cannot lose more than 5 points between Populace and Wealth
		g.valuePointsNeeded[Wealth] = -3
	case goalFulfilASuccessfulCoupDetat:
		g.cumulativeEffectNeeded[Intel] = 12
		g.cumulativeEffectNeeded[Politics] = 30
		g.roll4Needed[Clv] = 3
		g.roll5Needed[Sch] = 2
		g.characteristicPointsNeeded[Pop] = 1
	case goalGrowByLeapsAndBounds:
		g.cumulativeEffectNeeded[Acquisition] = 10
		g.cumulativeEffectNeeded[PublicRelations] = 15
		g.cumulativeEffectNeeded[Recruit] = 20
		g.characteristicPointsNeeded[Lty] = 1
		g.valuePointsNeeded[Populace] = 2
	case goalHoldAnInterstellarPeaceConference:
		g.cumulativeEffectNeeded[Expression] = 10
		g.cumulativeEffectNeeded[Security] = 10
		g.cumulativeEffectNeeded[Posturing] = 15
		g.roll5Needed[Pop] = 3
		g.valuePointsNeeded[Morale] = -2
	case goalInventANewTechnologicalMarvel:
		g.cumulativeEffectNeeded[Intel] = 10
		g.cumulativeEffectNeeded[Security] = 15
		g.cumulativeEffectNeeded[Research] = 15
		g.roll6Needed[Clv] = 2
		g.traitPointsNeeded[Technology] = 2
	case goalOrganiseOrderFromChaos:
		g.cumulativeEffectNeeded[Expression] = 10
		g.cumulativeEffectNeeded[PublicRelations] = 10
		g.cumulativeEffectNeeded[Tutelage] = 15
		g.roll5Needed[Pop] = 3
		g.roll6Needed[Tra] = 1
	case goalStartAnInterstellarWar:
		g.cumulativeEffectNeeded[Expression] = 10
		g.cumulativeEffectNeeded[Politics] = 10
		g.cumulativeEffectNeeded[Hostility] = 15
		g.cumulativeEffectNeeded[Tactical] = 15
		g.roll4Needed[Pop] = 4
	case goalTeachANewSkillToTheMasses:
		g.cumulativeEffectNeeded[Recruit] = 10
		g.cumulativeEffectNeeded[Tutelage] = 10
		g.roll5Needed[Pop] = 2
		g.roll4Needed[Tra] = 3
		g.traitPointsNeeded[Culture] = 2
	case goalUtterGenocide:
		g.cumulativeEffectNeeded[Conquest] = 15
		g.cumulativeEffectNeeded[Hostility] = 15
		g.cumulativeEffectNeeded[Intel] = 15
		g.cumulativeEffectNeeded[Sabotage] = 15
		g.valuePointsNeeded[Morale] = -2
	}
}
