package operations

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	EVENT_OPPORTUNITY = iota
	EVENT_INCIDENT
	EVENT_MISHAP
	EVENT_CRISIS
	NO_EVENT
)

type abstractTask struct {
	descr         string
	expectedDays  int
	taskModifiers map[string]int
	successIndex  int
	resolved      bool
}

func NewTask(descr string, expectedDays int) *abstractTask {
	tsk := abstractTask{}
	tsk.descr = descr
	tsk.expectedDays = expectedDays
	tsk.taskModifiers = make(map[string]int)
	return &tsk
}

func (tsk abstractTask) AddModifier(circumstance string, dm int) error {
	if tsk.resolved {
		return fmt.Errorf("can't add modifier to resolved task")
	}
	tsk.taskModifiers[circumstance] = dm
	return nil
}

func (tsk *abstractTask) addUpMods() int {
	m := 0
	for _, v := range tsk.taskModifiers {
		m += v
	}
	return m
}

func (t *abstractTask) ResolveBy(r Resolver) (string, int) {
	//логируем задание
	r.AddEntry("Resolving: " + t.descr)
	//расчитываем время - возвращаем потраченное
	t.successIndex = r.Resolve(t.addUpMods())
	t.resolved = true
	rep, evnt := resolutionReport(t.descr, t.successIndex)
	r.AddEntry(rep)
	return evnt, t.expectedDays
}

func resolutionReport(task string, tsi int) (string, string) {
	s := ""
	s += fmt.Sprintf("Resolving task was a %v (TSI:%v)", resolutionOutcome(tsi), tsi)
	eventCode := "No Event"
	switch resolutionEvent(tsi) {
	case EVENT_INCIDENT:
		eventCode = "Incident " + dice.New().RollNext("2d6").SumStr()
	case EVENT_OPPORTUNITY:
		eventCode = "Opportunity " + dice.New().RollNext("2d6").SumStr()
	case EVENT_MISHAP:
		eventCode = "Mishap " + dice.New().RollNext("2d6").SumStr()
	case EVENT_CRISIS:
		eventCode = "Crisis"
	}
	return s, eventCode
}

func resolutionEvent(i int) int {
	if i < 0 {
		i = 0
	}
	if i > 15 {
		i = 15
	}
	switch i {
	case 0, 1, 2, 5:
		if dice.Roll2D() == 2 {
			return EVENT_CRISIS
		}
		return EVENT_MISHAP
	case 3, 4:
		return EVENT_INCIDENT
	case 8, 13, 14, 15:
		return EVENT_OPPORTUNITY
	case 7, 9, 10, 11, 12:
	}
	return NO_EVENT
}

//
func resolutionOutcome(i int) string {
	if i <= 0 {
		return "Complete Failure"
	}
	switch i {
	case 1, 2:
		return "Failure"
	case 3, 4:
		return "Minimal Success"
	case 5:
		return "Partial Success"
	case 6, 7:
		return "Success"
	case 8, 9, 10:
		return "Solid Success"
	case 11, 12, 13, 14:
		return "Impressive Success"
	}
	return "Perfect Success"
}
