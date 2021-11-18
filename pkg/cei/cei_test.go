package cei

import (
	"fmt"
	"testing"
)

func TestCEI(t *testing.T) {
	crew, err := NewCrew(7,
		*NewDivision(DIVISION_COMMAND),
		*NewDivision(DIVISION_FLIGHT),
		*NewDivision(DIVISION_ENGINEERING),
		*NewDivision(DIVISION_MISSION),
	)
	fmt.Println(crew)
	if err != nil {
		t.Errorf(err.Error())
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
