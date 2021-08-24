package background

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/core/skill"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type background struct {
	worldName string
	worldTC   []string
}

func NewBackground(name string, codes []string) *background {
	bg := background{}
	bg.worldName = name
	bg.worldTC = codes
	return &bg
}

func (bg *background) LearnOrder(dice *dice.Dicepool) []int {
	lo := bg.basicSkills()
	for len(lo) < 6 {
		pick := dice.RollFromListInt(skill.BackgroundSkills())
		lo = addIfUniqueInt(lo, pick)
	}
	return lo
}

func (bg *background) basicSkills() []int {
	list := []int{}
	for _, code := range bg.worldTC {
		switch code {
		case "Ag", "Po":
			list = append(list, skill.Animals)
		case "As", "Va", "Ic":
			list = append(list, skill.VaccSuit)
		case "De", "Lt":
			list = append(list, skill.Survival)
		case "Fl", "Wa":
			list = append(list, skill.Seafarer)
		case "Ht":
			list = append(list, skill.Electronics)
		case "Hi":
			list = append(list, skill.Streetwise)
		case "In":
			list = append(list, skill.Profession)
		case "Ri":
			list = append(list, skill.Carouse)
		}
		fmt.Print(code + "\r   ")
	}
	return []int{}
}

func addIfUniqueInt(sl []int, new int) []int {
	unique := true
	for _, val := range sl {
		if val == new {
			unique = false
		}
	}
	if unique {
		sl = append(sl, new)
	}
	return sl
}
