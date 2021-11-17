package mission

// import (
// 	"github.com/Galdoba/TR_Dynasty/pkg/cei"
// 	"github.com/Galdoba/TR_Dynasty/pkg/dice"
// )

// const (
// 	EVENT_OPPORTUNITY = "Opportunity"
// 	EVENT_INCIDENT    = "Incident"
// 	EVENT_MISHAP      = "Mishap"
// 	EVENT_CRISIS      = "Crisis"
// )

// type mission struct {
// 	title      string
// 	objectives []string
// 	segment    []*segment
// }

// func NewMission(title string)

// type segment struct {
// 	asignedCrew     *cei.CEI
// 	description     string
// 	degreeOfSuccess string
// 	operations      []*operation
// 	msi             int
// }

// func NewSegment(crew *cei.CEI, descr string, opers ...*operation) *segment {
// 	s := segment{
// 		asignedCrew: crew,
// 		description: descr,
// 	}
// 	for _, val := range opers {
// 		s.operations = append(s.operations, val)
// 	}
// 	return &s
// }

// func (s *segment) Resolve() {
// 	for _, op := range s.operations {
// 		op.Resolve()
// 	}
// }

// func (s *segment) Report() string {
// 	rep := "  Mission Segment: " + s.description
// 	for _, op := range s.operations {
// 		rep += "\n" + op.Report()
// 	}
// 	rep += "\n  Segment Summary: TODO"
// 	return rep
// }

// type operation struct {
// 	asignedCrew     *cei.CEI
// 	descr           string
// 	resolutionIndex int
// 	outcome         string
// 	event           string
// 	activeModifiers []resolutionModifier
// 	summary         string
// }

// type resolutionModifier struct {
// 	descr string
// 	val   int
// }

// func Modifier(descr string, val int) resolutionModifier {
// 	return resolutionModifier{descr, val}
// }

// //NewOperation - TODO: переделать *cei.CEI в интерфейс Resolver
// func NewOperation(crew *cei.CEI, descr string, mods ...resolutionModifier) *operation {
// 	o := operation{
// 		asignedCrew:     crew,
// 		descr:           descr,
// 		resolutionIndex: -999,
// 		outcome:         "Not concluded",
// 		event:           "Unknown",
// 		activeModifiers: mods,
// 	}
// 	return &o
// }

// func (o *operation) Resolve() {
// 	r := dice.Roll2D() + o.asignedCrew.TaskDM()
// 	for _, modifer := range o.activeModifiers {
// 		r += modifer.val
// 	}
// 	o.resolutionIndex = r
// 	o.event = resolutionEvent(r)
// 	o.outcome = resolutionOutcome(r)
// 	o.summary = "TODO: Summary"
// }

// func (o *operation) Report() string {
// 	report := "    Operation: " + o.descr + " (" + o.outcome + ")"
// 	if o.event != "" {
// 		report += "\n      Event: " + o.event
// 	}
// 	report += "\n      Summary: " + o.summary
// 	return report
// }

// func resolutionEvent(i int) string {
// 	switch i {
// 	case 0, 1, 2, 5:
// 		if dice.Roll2D() == 2 {
// 			return EVENT_CRISIS
// 		}
// 		return EVENT_MISHAP
// 	case 3, 4:
// 		return EVENT_INCIDENT
// 	case 8, 13, 14, 15:
// 		return EVENT_OPPORTUNITY
// 	case 7, 9, 10, 11, 12:
// 	}
// 	if i < 0 {
// 		if dice.Roll2D() == 2 {
// 			return EVENT_CRISIS
// 		}
// 		return EVENT_MISHAP
// 	}
// 	if i > 15 {
// 		return EVENT_OPPORTUNITY
// 	}
// 	return ""
// }

// func resolutionOutcome(i int) string {
// 	if i <= 0 {
// 		return "Complete Failure"
// 	}
// 	switch i {
// 	case 1, 2:
// 		return "Failure"
// 	case 3, 4:
// 		return "Minimal Success"
// 	case 5:
// 		return "Partial Success"
// 	case 6, 7:
// 		return "Success"
// 	case 8, 9, 10:
// 		return "Solid Success"
// 	case 11, 12, 13, 14:
// 		return "Impressive Success"
// 	}
// 	return "Perfect Success"
// }
