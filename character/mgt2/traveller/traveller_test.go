package traveller

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/core/skill"
)

func TestNewCharacter(t *testing.T) {
	chr, err := NewCharacter()
	if err != nil {
		t.Errorf("func error: '%v'\n", err.Error())
	}
	testCharacteristics := []string{"STR", "DEX", "END", "INT", "EDU", "SOC", "PSI", "SAN"}
	for i, val := range testCharacteristics {
		chr.SetChar(val, i)
		if chr.chars[val].Value() != i {
			t.Errorf("SetChar: have'%v' = '%v', expect %v = %v\n", val, chr.chars[val].Value(), val, i)
		}
	}
	chr.Train(skill.Admin)
	chr.Train(skill.Admin)
	chr.Train(skill.Admin)
	chr.Train(skill.Animals)
	chr.Train(skill.Animals)
	chr.Train(skill.Animals)
	chr.BasicTraining(skill.Electronics)
	chr.EnsureSkill(skill.Animals_Handling, 1)
	fmt.Println(chr.skills)

}

func TestCharacterCharacteristics(t *testing.T) {
	chr, _ := NewCharacter()
	chr.rollCharacteristics()
	for key, val := range chr.chars {
		if val.Value() == 0 {
			t.Errorf("This test MUST NOT fail, %v = %v\n", key, val)
		}

	}
}
