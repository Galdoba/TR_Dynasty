package mission

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/cei"
)

func Test_Mission(t *testing.T) {
	crew := cei.NewTeam("Crew", 7)
	reach := NewReach("Hushumaki Transit",
		Mission("Transit from A to C",
			MissionSegment("Jump from A to B",
				Operation("Skimm Fuel"),
				Operation("Jump to B"),
			),
			MissionSegment("Jump from B to C",
				Operation("Skimm Fuel"),
				Operation("Jump to C"),
			),
			MissionSegment("Jump from C to D",
				Operation("Skimm Fuel"),
				Operation("Jump to C"),
			),
		),
	)

	reach.AssignResolver(crew)

	if _, err := reach.AbstractResolve(); err != nil {
		t.Errorf("task not resolved: %v (not expected)\n", err.Error())
	}
	fmt.Println("------------------")
	fmt.Println(reach)

}

func TestStaticFunctions(t *testing.T) {
	for i := -3; i < 20; i++ {
		//fmt.Printf("resolution outcome '%v' = '%v'\n", i, resolutionOutcome(i))
		//fmt.Printf("resolution event '%v' = '%v'\n", i, resolutionEvent(i))
	}
	fmt.Printf("                                                           \r")
}
