package traveller

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/core/skill"
)

func (t *traveller) BasicTraining(skCode int) error {
	sk := skill.ByCode(skCode)
	skGrp := sk.Group
	list := skill.ByGroup(skGrp)
	for _, skl := range list {
		if _, ok := t.skills[skl.Code]; !ok {
			t.skills[skl.Code] = skill.ByCode(skl.Code)
		}
	}
	return nil
}

func (t *traveller) Train(skCode int) error {
	skl := skill.ByCode(skCode)
	if _, ok := t.skills[skCode]; !ok {
		t.skills[skCode] = skl
	}
	s := t.skills[skCode]
	s.Value++
	t.skills[skCode] = s
	t.fixGroupValues()
	return nil
}

func (t *traveller) fixGroupValues() {
	for _, val := range t.skills {
		if val.Speciality == "" && val.Value > 0 {
			list := skill.ByGroup(val.Group)
			if len(list) < 2 {
				continue
			}
			toincr := []int{}
			for _, s := range list {
				if s.Speciality == "" {
					continue
				}
				toincr = append(toincr, s.Code)
			}
			val.Value--
			t.skills[val.Code] = val
			r := t.dice.RollNext("1d" + strconv.Itoa(len(toincr))).DM(-1).Sum()
			t.Train(toincr[r])
		}
	}
}

func (t *traveller) EnsureSkill(code, val int) {
	if _, ok := t.skills[code]; !ok {
		t.skills[code] = skill.ByCode(code)
	}
	sk := t.skills[code]
	if sk.Value < val {
		sk.Value = val
	}
	t.skills[code] = sk
}
