package traveller

import (
	"fmt"
	"testing"
)

func TestEducation(t *testing.T) {
	fmt.Println("Test Education")
	trv := NewTravellerT5()
	trv.Educate()
	if trv.education == nil {
		t.Errorf("education Error: %v", trv.err.Error())
		return
	}
	if trv.err != nil {
		t.Errorf("education Error: %v", trv.err.Error())
		return
	}
	if trv.education.educationAtempts[EInst_ED5] > 1 {
		t.Errorf("ED5 Program can't be applied more than once: tries = %v", trv.education.educationAtempts[EInst_ED5])
	}

	cc := NewCard(trv)
	cc.PrintCard()
}

/*
Education Process:
Meet Prequsites (T5Charater) bool
Apply Check
Pass/Fail Check - yearly
Benefit - yearly
Graduate

Institution Struct:
majors []string
pre requisite check
pass check
//benefits

years int





*/
