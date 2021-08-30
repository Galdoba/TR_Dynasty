package assets

import (
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
			if c.err != nil {
				t.Errorf("creation Error with code %v and dice %v: %v", code, d, c.err.Error())
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
