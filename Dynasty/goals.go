package dynasty

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/dice"
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
	trigerDay                  int
	reward                     func(Dynasty) *Dynasty
	log                        string
}

func testNewGoal(urgency int) goal {
	g := goal{}
	g.trigerDay = urgency
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

func Test4() {

	dyn := NewDynasty("1")
	for i, val := range listAptitudes() {
		dyn.stat[val] = i
	}
	fmt.Println(dyn.Info())
	dyn.goals = append(dyn.goals, testNewGoal(2))

	for cd := 0; cd < 4; cd++ {
		fmt.Println("CD", cd)
		if cd < dyn.goals[0].trigerDay {
			continue
		}
		r := dice.Flux()
		dyn.goals[0].cumulativeEffectNeeded[Clv] = dyn.goals[0].cumulativeEffectNeeded[Clv] + r
		dyn.goals[0].conclude()
		fmt.Println(dyn)
		dyn.goals[0].reward(dyn)
	}
	fmt.Println(dyn.Info())
}

func (g *goal) conclude() {
	eventMap := make(map[string]func(Dynasty) *Dynasty)
	eventMap["test3"] = func(d Dynasty) *Dynasty {
		d.stat[Clv] = 99
		return &d
	}
	eventMap["test4"] = func(d Dynasty) *Dynasty {
		d.stat[Clv] = -99
		return &d
	}
	key := ""
	switch g.success() {
	default:
		fmt.Println("-----------")
	case true:
		key = "3"
		fmt.Println("++++++++++++++++")
	case false:
		key = "4"
		fmt.Println("=============")
	}
	name := "test"
	if val, ok := eventMap[name+key]; ok {
		g.reward = val
	} else {
		//fmt.Println("Null event")
		g.reward = func(d Dynasty) *Dynasty {
			return &d
		}
	}
}

func (g *goal) success() bool {
	for k, v := range g.cumulativeEffectNeeded {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " effect missing for " + k
			return false
		}
	}
	for k, v := range g.roll4Needed {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " 4+ effect rolls missing for " + k
			return false
		}
	}
	for k, v := range g.roll5Needed {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " 5+ effect rolls missing for " + k
			return false
		}
	}
	for k, v := range g.roll6Needed {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " 6+ effect rolls missing for " + k
			return false
		}
	}
	for k, v := range g.characteristicPointsNeeded {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " characteristic points missing for " + k
			return false
		}
	}
	for k, v := range g.traitPointsNeeded {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " trait points missing for " + k
			return false
		}
	}
	for k, v := range g.valuePointsNeeded {
		if v > 0 {
			g.log += "\n" + strconv.Itoa(v) + " value points missing for " + k
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
		if goal.trigerDay > 0 {
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
Fulfil a Successful Coup Dé'tat
Grow by Leaps and Bounds
Hold an Interstellar Peace Conference
Organise Order from Chaos
Start an Interstellar War
Teach a New Skill to the Masses
Utter Genocide
[NOT ESTABLESHED]



*/
