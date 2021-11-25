package mission

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type AbstractTransit struct {
	RateOfAdvance           int
	EventAvoided            bool
	PointOfInterestDetected bool
	ParsecsCovered          int
	resolved                bool
}

const (
	EVENT_OPPORTUNITY = "Opportunity"
	EVENT_INCIDENT    = "Incident"
	EVENT_MISHAP      = "Mishap"
	EVENT_CRISIS      = "Crisis"
	SCALE_REACH       = 1
	SCALE_MISSION     = 2
	SCALE_SEGMENT     = 3
	SCALE_OPERATION   = 4
)

type operation struct {
	asignedCrew     Resolver
	descr           string
	objectives      []string
	scale           int
	resolutionIndex int
	outcome         string
	event           string
	summary         string
	activeModifiers []*resolutionModifier
	subOperations   []*operation
}

type Resolver interface {
	Resolve(...string) int
}

type resolutionModifier struct {
	descr string
	val   int
}

type Modifier interface {
	Modifier() (string, int)
}

func (rm *resolutionModifier) Modifier() (descr string, val int) {
	return rm.descr, rm.val
}

func NewModifier(descr string, val int) *resolutionModifier {
	rm := resolutionModifier{descr, val}
	return &rm
}

//NewOperation - TODO: переделать *cei.CEI в интерфейс Resolver
func NewReach(descr string, missions ...*operation) *operation {
	o := operation{
		descr: descr,
		scale: SCALE_REACH,
	}
	for _, mission := range missions {
		o.subOperations = append(o.subOperations, mission)
	}
	return &o
}

func (op *operation) SetScale(scale int) {
	switch scale {
	case SCALE_REACH, SCALE_MISSION, SCALE_SEGMENT, SCALE_OPERATION:
		op.scale = scale
	}
}

func (op *operation) AssignResolver(team Resolver) *operation {
	switch {
	case op.asignedCrew == nil:
		op.asignedCrew = team
	case op.asignedCrew != nil:
		return op
	}
	//op.asignedCrew = team
	for _, t1 := range op.subOperations {
		t1.AssignResolver(team)
	}
	return op
}

func (op *operation) SetModifiers(mods ...Modifier) *operation {
	for _, mod := range mods {
		descr, val := mod.Modifier()
		op.activeModifiers = append(op.activeModifiers, &resolutionModifier{descr: descr, val: val})
	}
	return op
}

func (op *operation) AbstractResolve() (int, error) {
	if op.asignedCrew == nil {
		return -999, fmt.Errorf("resolution team not asignned")
	}
	ri := 0
	switch len(op.subOperations) {
	case 0: //если нижний уровень
		ri = op.asignedCrew.Resolve(op.descr)
		for _, mod := range op.activeModifiers {
			_, v := mod.Modifier()
			ri += v

		}
		op.resolutionIndex = ri
		op.outcome = resolutionOutcome(op.resolutionIndex)
		op.event = resolutionEvent(op.resolutionIndex)
	default: //если НЕ нижний уровень
		segm := 0
		for i, task := range op.subOperations {
			tri, err := task.AbstractResolve()
			if err != nil {
				return -999, err
			}
			op.resolutionIndex += tri
			segm = i + 1
		}
		op.resolutionIndex = op.resolutionIndex / segm
		op.outcome = resolutionOutcome(op.resolutionIndex)
		//op.event = resolutionEvent(op.resolutionIndex)
	}
	return op.resolutionIndex, nil
}

func resolutionEvent(i int) string {
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
	if i < 0 {
		if dice.Roll2D() == 2 {
			return EVENT_CRISIS
		}
		return EVENT_MISHAP
	}
	if i > 15 {
		return EVENT_OPPORTUNITY
	}
	return "No Event"
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

func (op *operation) getScale() string {
	switch op.scale {
	case 1:
		return "Reach"
	case 2:
		return "Mission"
	case 3:
		return "Mission Segment"
	case 4:
		return "Operation"
	}
	return "Task"
}

func Operation(descr string) *operation {
	op := operation{
		descr: descr,
		scale: SCALE_OPERATION,
	}
	return &op
}

func MissionSegment(descr string, operations ...*operation) *operation {
	op := operation{
		descr: descr,
		scale: SCALE_SEGMENT,
	}
	for _, oper := range operations {
		op.subOperations = append(op.subOperations, oper)
	}
	return &op
}

func Mission(descr string, misSegm ...*operation) *operation {
	op := operation{
		descr: descr,
		scale: SCALE_MISSION,
	}
	for _, misSeg := range misSegm {
		op.subOperations = append(op.subOperations, misSeg)
	}
	return &op
}

func (op *operation) String() string {
	prefix := ""
	switch op.scale {
	case SCALE_MISSION:
		prefix = "    "
	case SCALE_SEGMENT:
		prefix = "        "
	case SCALE_OPERATION:
		prefix = "            "
	}
	r := fmt.Sprintf("%v%v: %v", prefix, op.getScale(), op.descr)
	//if len(op.subOperations) == 0 {
	r += fmt.Sprintf(" (%v) | %v %v", op.outcome, op.resolutionIndex, op.event)
	//}
	for _, s := range op.subOperations {
		r += "\n" + s.String()
	}
	return r
}

func (op *operation) sumEvents() []string {
	eventSl := []string{op.event}
	for _, val := range op.lowerScale() {
		eventSl = append(eventSl, val.sumEvents()...)
	}
	clear := []string{}
	for _, sl := range eventSl {
		if sl != "No Event" && sl != "" {
			clear = append(clear, sl)
		}
	}
	return clear
}

func (op *operation) lowerScale() []*operation {
	return op.subOperations
}

/*
REACH: Reach
Aim: Aim
	Mission: M1
		Mission Segment: S1
			Operation: O1
			Operation: O1
			Operation: O1
		Mission Segment: S2
			Operation: O1
			Operation: O1
			Operation: O1
	Mission: M1
		Mission Segment: S1
			Operation: O1
Summary:



*/
