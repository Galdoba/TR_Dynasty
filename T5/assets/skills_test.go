package assets

import "testing"

func TestSkills(t *testing.T) {
	for i := -2; i < 250; i++ {
		sk := NewSkill(i)
		if sk.err != nil {
			t.Errorf("creation error: %v", sk.err.Error())
		}
	}
}
