package assets

import (
	"fmt"
	"testing"
)

func testCharCodes() []int {
	return []int{
		Strength,
		Dexterity,
		Agility,
		Grace,
		Endurance,
		Stamina,
		Vigor,
		Intelligence,
		Education,
		Training,
		Instinct,
		Social,
		Charisma,
		Caste,
		Sanity,
		Psionics,
	}
}

func TestCharacteristic(t *testing.T) {
	charCodes := testCharCodes()
	for _, code := range charCodes {
		for d := 1; d < 9; d++ {
			c := NewCharacteristic(code, d)
			if c.Err != nil {
				t.Errorf("creation Error with code %v and dice %v: %v", code, d, c.Err.Error())
				continue
			}
			if c.Value() < 0 || c.Value() > c.charDice*6 {
				t.Errorf("value Error with code %v and dice %v: expect something between 0 and %v, have %v", code, d, c.charDice*6, c.value)
			}
			if c.Genetics() < 1 || c.Genetics() > 6 {
				t.Errorf("genetics Error with code %v and dice %v: expect something between 1 and 6, have %v", code, d, c.genetics)
			}
		}
	}
}

func TestSubstitution(t *testing.T) {
	list1 := testCharCodes()
	list2 := testCharCodes()
	for _, valO := range list1 {
		for _, valS := range list2 {
			chr := NewCharacteristic(valO, 2)
			chr2 := NewCharacteristic(valS, 2)
			valR, err := chr.UseAs(valS)
			if err != nil {
				expectedError := fmt.Errorf("%v cannot be substituded with %v", chr.Name(), chr2.Name())
				if err.Error() == expectedError.Error() {
					continue
				}
				t.Errorf("substitution error: '%v' | (origin=%v, target=%v, result=%v)", err.Error(), valO, valS, valR)
			} else {
				//t.Errorf("substitution sucseeds: %v use as %v origin val=%v, result=%v", valO, valS, chr.value, valR)
			}

		}
	}
}
