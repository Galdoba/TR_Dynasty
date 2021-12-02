package maintenance

import (
	"fmt"
	"testing"
)

func TestReporter(t *testing.T) {
	dcm := NewDamageControlManager()
	fmt.Printf("%v\n", dcm)
	if dcm == nil {
		t.Errorf("dcm cannot ne NIL")
	}
	if err := dcm.Assign(DetermineIssues(40)...); err != nil {
		t.Errorf("Assign error: %v\n", err.Error())
	} else {
		fmt.Printf("defect assigned...\n")
	}

	fmt.Println(dcm)
	for _, v := range dcm.problemWith {
		fmt.Println(v)
	}
}
