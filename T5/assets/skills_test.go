package assets

import (
	"fmt"
	"testing"
)

func TestSkills(t *testing.T) {
	for i := 0; i < 64; i++ {
		sk := NewSkill(i)
		if sk.err != nil {
			t.Errorf("creation error: %v", sk.err.Error())
		} else {
			fmt.Printf("	creation pass:	Code '%v' is valid and goes to skill '%v'\n", i, sk.alias)
			fmt.Printf("	Test sk.String() = %v, %v\n", sk.String(), sk)

			for j := 0; j < 17; j++ {
				if err := sk.Train(); err != nil {
					fmt.Printf("	Train %v failed:	%v\n", j, err.Error())
				}
			}
			fmt.Printf("	Test sk.Train(): name after Train '%v'|%v\n", sk.String(), sk)
		}
		fmt.Println("")
	}
}
