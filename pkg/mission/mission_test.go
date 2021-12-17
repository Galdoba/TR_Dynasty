package mission

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/ship/cei"
)

func Test_Mission(t *testing.T) {
	crew := cei.NewTeam("Crew", 8)
	crew.AddDivision(cei.DIVISION_COMMAND, 8)
	crew.AddDivision(cei.DIVISION_FLIGHT, 8)
	crew.AddDivision(cei.DIVISION_ENGINEERING, 8)
	crew.AddDivision(cei.DIVISION_MISSION, 8)
	crew.Morale = crew.CrewEfficencyIndex + dice.New().RollNext("2d3").Sum()

	reach := NewReach("Hushumaki Transit",
		Mission("Transit from A to C",
			MissionSegment("Jump from A to B",
				Operation("Skimm Fuel").AssignResolver(crew.CallDivision(cei.DIVISION_FLIGHT)),
				Operation("Jump to B").AssignResolver(crew.CallDivision(cei.DIVISION_ENGINEERING)),
			),
			MissionSegment("Jump from B to C",
				Operation("Skimm Fuel").AssignResolver(crew.CallDivision(cei.DIVISION_FLIGHT)),
				Operation("Jump to C").AssignResolver(crew.CallDivision(cei.DIVISION_ENGINEERING)),
			),
			MissionSegment("Jump from C to D",
				Operation("Skimm Fuel").AssignResolver(crew.CallDivision(cei.DIVISION_FLIGHT)),
				Operation("Jump to D").AssignResolver(crew.CallDivision(cei.DIVISION_ENGINEERING)),
			),
		),
		Mission("Obtain Artefact",
			MissionSegment("Conduct orbital survey"),
			MissionSegment("Locate source of signals"),
			MissionSegment("Search planetside location",
				Operation("Land shuttle").AssignResolver(crew.CallDivision(cei.DIVISION_FLIGHT)),
				Operation("Find source and investigate at ground level").AssignResolver(crew.CallDivision(cei.DIVISION_MISSION)),
				Operation("Retrieve Artefact").AssignResolver(crew.CallDivision(cei.DIVISION_MISSION)),
				Operation("Obtain any additional information").AssignResolver(crew.CallDivision(cei.DIVISION_MISSION)),
			),
			MissionSegment("Return to shuttle and climb back to parentship").AssignResolver(crew.CallDivision(cei.DIVISION_FLIGHT)),
		),
	)

	reach.AssignResolver(crew)

	if _, err := reach.AbstractResolve(); err != nil {
		t.Errorf("task not resolved: %v (not expected)\n", err.Error())
	}
	fmt.Println("------------------")
	fmt.Println(reach)
	fmt.Println(reach.sumEvents())
	crew.PrintLog()
}

func TestStaticFunctions(t *testing.T) {
	for i := -3; i < 20; i++ {
		//fmt.Printf("resolution outcome '%v' = '%v'\n", i, resolutionOutcome(i))
		//fmt.Printf("resolution event '%v' = '%v'\n", i, resolutionEvent(i))
	}
	fmt.Printf("                                                           \r")
}
