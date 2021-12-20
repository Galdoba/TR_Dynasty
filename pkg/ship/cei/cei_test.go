package cei

import (
	"fmt"
	"testing"
	"time"
)

func TestCEI(t *testing.T) {
	crew := NewTeam("Ship Crew", 8)
	crew.SetMorale(11)
	crew.AddDivision(DIVISION_FLIGHT, 8)
	crew.AddDivision(DIVISION_ENGINEERING, 8)
	crew.AddDivision(DIVISION_OPERATIONS, 8)
	crew.AddDivision(DIVISION_MISSION, 8)
	crew.Update()

	//operations.NewTask("Land shuttle", 0).ExecuteBy(crew)

	fmt.Println("===LOG==============")
	crew.PrintLog()
	for i := 0; i < 100; i++ {
		crew.Update()
		fmt.Println("test", i)
		crew.Report()
		time.Sleep(time.Millisecond * 10)
		if i%10 == 0 {
			//at, err := operations.NewAbstractTransit(operations.FLANK_SPEED, "test transit "+strconv.Itoa(i))
			//if err != nil {

			//}
			//at.ExecuteBy(crew)
		}
	}
	crew.PrintLog()
	//fmt.Println(cei.ResolveMission("Operation: Retrive Artefect"))
	//fmt.Println(cei.ResolveMission("Operation: Obtain any additional information"))
}
