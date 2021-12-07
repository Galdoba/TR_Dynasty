package operations

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/ship/cei"
)

func TestAdvance(t *testing.T) {
	for i, rate := range emulateRate() {

		at, err := NewAbstractTransit(rate, fmt.Sprintf("Test transit %v", i))
		if err != nil {
			t.Errorf("error: %v\n", err.Error())
		} else {
			if at.parsecsCovered == 0 {
				t.Errorf("Expected to advance but did not")
			}
			if at.event == "" {
				t.Errorf("Expected event to be familiar but have not")
			}
			if at.poi == "" {
				t.Errorf("Expected point of interest to be familiar but have not")
			}
			fmt.Println(at)
		}
	}
	for r1 := 2; r1 <= 12; r1++ {
		for r2 := 1; r2 <= 6; r2++ {
			poi := rollForPoi(r1, r2)
			if poi == "" {
				t.Errorf("blank poi on roll %v %v\n", r1, r2)
			}
		}
	}

	fmt.Print("\n")

}

func emulateRate() []int {
	return []int{
		FLANK_SPEED,
		RAPID_TRANSIT,
		CURSORY_EXPLORATION,
		DETAILED_EXPLORATION,
	}
}

func TestExecution(t *testing.T) {
	totalParsecs := 0
	ship := cei.NewTeam("Crew", 7)
	counter := 1
	for totalParsecs < 120 {

		transit, _ := NewAbstractTransit(FLANK_SPEED, "Test Transit "+strconv.Itoa(counter))
		transit.ExecuteBy(ship)
		if !transit.ready {
			t.Errorf("Transit supposed to be ready")
		}
		totalParsecs += transit.parsecsCovered
		fmt.Println(transit)
		fmt.Println(totalParsecs, "parsecs covered in total")
		fmt.Println("=============================")
		counter++
	}
	fmt.Println("Done")
}
