package cei

import (
	"fmt"
	"testing"
)

func TestCEI(t *testing.T) {
	crew := NewTeam("Ship Crew", 8)
	crew.SetMorale(11)
	crew.AddDivision(DIVISION_FLIGHT, 8)
	crew.AddDivision(DIVISION_ENGINEERING, 8)
	crew.AddDivision(DIVISION_OPERATIONS, 8)
	crew.AddDivision(DIVISION_MISSION, 8)
	crew.Update()

	fmt.Println(crew.Resolve("Operation: Land shuttle"))
	fmt.Println(crew.Resolve("Operation: Find source and investigate at ground level"))
	fmt.Println("===LOG==============")
	crew.PrintLog()
	crew.Update()
	crew.Report()
	//fmt.Println(cei.ResolveMission("Operation: Retrive Artefect"))
	//fmt.Println(cei.ResolveMission("Operation: Obtain any additional information"))
}
