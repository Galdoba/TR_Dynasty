package cei

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TR_Dynasty/pkg/ship/operations"
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
	eventDay := 0
	eventCode := ""
	for i := 0; i < 1000; i++ {
		crew.Update()
		fmt.Println("test", i)
		crew.Report()
		time.Sleep(time.Millisecond * 10)
		if i%10 == 0 {
			task := operations.NewTask("Abstract task", 5)

			eventCode0, resDays := task.ResolveBy(crew)
			eventDay = resDays + crew.MissionDay
			eventCode = eventCode0

		}
		if crew.MissionDay == eventDay {
			crew.CallEvent(eventCode)
		}
	}
	crew.PrintLog()
	//fmt.Println(cei.ResolveMission("Operation: Retrive Artefect"))
	//fmt.Println(cei.ResolveMission("Operation: Obtain any additional information"))
}
