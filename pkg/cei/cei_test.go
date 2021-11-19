package cei

import (
	"fmt"
	"testing"
)

func TestCEI(t *testing.T) {
	crew := NewTeam("Ship Crew",
		DIVISION_COMMAND,
		DIVISION_FLIGHT,
		DIVISION_ENGINEERING,
		DIVISION_MISSION,
	)
	crew.Assemble()
	crew.Report()
	crew.CEIMchanges("Bad Event", -6)
	for _, entry := range crew.Log {
		fmt.Println(entry)
	}
	crew.Report()
	// fmt.Println("Mission: Search planet side location:")
	// fmt.Println(cei.CheckTask(8))
	// fmt.Println(cei.ResolveTask())
	//fmt.Println(cei.ResolveMission("Operation: Land shuttle"))
	//fmt.Println(cei.ResolveMission("Operation: Find source and investigate at ground level"))
	//fmt.Println(cei.ResolveMission("Operation: Retrive Artefect"))
	//fmt.Println(cei.ResolveMission("Operation: Obtain any additional information"))
}
