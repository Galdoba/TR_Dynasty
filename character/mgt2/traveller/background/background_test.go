package background

import (
	"testing"
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
				expectedInList = append(expectedInList, "Animals 0")
			case "As":
				expectedInList = append(expectedInList, "Athlithics 0")
			case "De":
				expectedInList = append(expectedInList, "Survival 0")
			case "Oc":
				expectedInList = append(expectedInList, "Seafarer 0")
			case "Ht":
				expectedInList = append(expectedInList, "Electronic 0")
			case "Hi":
				expectedInList = append(expectedInList, "Streetwise 0")
			case "Ic":
				expectedInList = append(expectedInList, "Vacc Suit 0")
			case "In":
				expectedInList = append(expectedInList, "Broker 0")
			case "Lt":
				expectedInList = append(expectedInList, "Survival 0")
			case "Po":
				expectedInList = append(expectedInList, "Animals 0")
			case "Ri":
				expectedInList = append(expectedInList, "Carouse 0")
			case "Wa":
				expectedInList = append(expectedInList, "Seafarer 0")
			case "Va":
				expectedInList = append(expectedInList, "Vacc Suit 0")
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
