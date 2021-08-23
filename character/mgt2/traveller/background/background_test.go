package background

import (
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/core/skill"
)

func testInput() []*background {
	var input []*background
	input = append(input, NewBackground("[Erehwemos]", []string{"Ag", "Ni"}))
	input = append(input, NewBackground("[Regina]", []string{"Hi", "Ht", "Ri"}))
	return input
}

func TestBackgroundSkills(t *testing.T) {
	for i, bg := range testInput() {
		bgSkills := bg.BasicSkills()
		expectedInList := []string{}
		//fmt.Printf("Input %v: World %v with codes %v have these skills: %v\n", i, bg.worldName, bg.worldTC, bgSkills)
		for _, tc := range validBGCodes() {
			if !sliceContains(bg.worldTC, tc) {
				continue
			}
			switch tc {
			default:
			case "Ag":
				expectedInList = append(expectedInList, skill.ByCode(skill.Animals).Name())
			case "As":
				expectedInList = append(expectedInList, skill.ByCode(skill.Athletics).Name())
			case "De":
				expectedInList = append(expectedInList, skill.ByCode(skill.Survival).Name())
			case "Fl":
				expectedInList = append(expectedInList, skill.ByCode(skill.Seafarer).Name())
			case "Ht":
				expectedInList = append(expectedInList, skill.ByCode(skill.Electronics).Name())
			case "Hi":
				expectedInList = append(expectedInList, skill.ByCode(skill.Streetwise).Name())
			case "Ic":
				expectedInList = append(expectedInList, skill.ByCode(skill.VaccSuit).Name())
			case "In":
				expectedInList = append(expectedInList, skill.ByCode(skill.Profession).Name())
			case "Lt":
				expectedInList = append(expectedInList, skill.ByCode(skill.Survival).Name())
			case "Po":
				expectedInList = append(expectedInList, skill.ByCode(skill.Animals).Name())
			case "Ri":
				expectedInList = append(expectedInList, skill.ByCode(skill.Carouse).Name())
			case "Wa":
				expectedInList = append(expectedInList, skill.ByCode(skill.Seafarer).Name())
			case "Va":
				expectedInList = append(expectedInList, skill.ByCode(skill.VaccSuit).Name())
			}
		}
		if !match(bgSkills, expectedInList) {
			t.Errorf("\nInput %v: World %v with codes %v have skills: \n%v,\nbut expect \n%v\n", i, bg.worldName, bg.worldTC, bgSkills, expectedInList)
		}
	}
}

func validBGCodes() []string {
	return []string{
		"Ag",
		"As",
		"De",
		"Oc",
		"Ht",
		"Hi",
		"Ic",
		"In",
		"Lt",
		"Po",
		"Ri",
		"Wa",
		"Va",
	}
}

func sliceContains(sl []string, s string) bool {
	for _, val := range sl {
		if val == s {
			return true
		}
	}
	return false
}

func match(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, val := range sl1 {
		if val != sl2[i] {
			return false
		}
	}
	return true
}
