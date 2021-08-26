package assets

import (
	"testing"
)

func TestSkills(t *testing.T) {
	for i := 0; i < 64; i++ {
		sk := NewSkill(i)
		if sk.err != nil {
			t.Errorf("creation error: %v", sk.err.Error())
		} else {
			for j := 0; j < 15; j++ {
				if err := sk.Train(); err != nil {
					t.Errorf("	Train %v failed:	%v\n", j, err.Error())
				}
			}
		}
	}
}
