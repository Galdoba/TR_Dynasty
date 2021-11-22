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
	SCALE_MISSION     = "Mission"
	SCALE_SEGMENT     = "Mission Segment"
	SCALE_OPERATION   = "Operation"
)

type operation struct {
	asignedCrew     Resolver
	descr           string
	opType          string
	resolutionIndex int
	outcome         string
	event           string
	activeModifiers []resolutionModifier
	summary         string
	subOperations   []*operation
}

type Resolver interface {
	Resolve() int
}

type resolutionModifier struct {
	descr string
	val   int
}

type Modifier interface {
	Modifier() (string, int)
}

func (rm *resolutionModifier) Modifier() (descr string, val int) {
	return rm.descr, val
}

func NewModifier(descr string, val int) resolutionModifier {
	rm := resolutionModifier{descr, val}
	return rm
}

//NewOperation - TODO: переделать *cei.CEI в интерфейс Resolver
func NewOperation(descr string, mods ...resolutionModifier) *operation {
	o := operation{
		asignedCrew:     nil,
		descr:           descr,
		resolutionIndex: -999,
		outcome:         "Not concluded",
		event:           "Unknown",
		activeModifiers: mods,
		opType:          "Undefined Task",
	}
	return &o
}

func (op *operation) AssignResolver(team Resolver) {
	op.asignedCrew = team
}

func (op *operation) SetModifiers(mods ...Modifier) {
	for _, mod := range mods {
		descr, val := mod.Modifier()
		op.activeModifiers = append(op.activeModifiers, resolutionModifier{descr: descr, val: val})
	}
}

func (op *operation) AbstractResolve() error {
	if op.asignedCrew == nil {
		return fmt.Errorf("resolution team not asignned")
	}
	ri := op.asignedCrew.Resolve()
	for _, mod := range op.activeModifiers {
		_, v := mod.Modifier()
		ri += v
	}
	op.resolutionIndex = ri
	op.outcome = resolutionOutcome(op.resolutionIndex)
	op.event = resolutionEvent(op.resolutionIndex)
	return nil
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

func (op *operation) String() string {
	r := fmt.Sprintf("%v: %v\n", op.opType, op.descr)
	r += fmt.Sprintf("Resolver: %v\n", op.asignedCrew)
	r += fmt.Sprintf("Resolution Outcome: %v\n", op.outcome)
	r += fmt.Sprintf("Resolution Index: %v\n", op.resolutionIndex)
	r += fmt.Sprintf("Resolution Event: %v", op.event)
	return r
}
