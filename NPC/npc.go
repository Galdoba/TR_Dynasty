package main

import (
	"fmt"

	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

const (
	chrSTR = "STR"
	chrDEX = "DEX"
	chrEND = "END"
	chrINT = "INT"
	chrEDU = "EDU"
	chrSOC = "SOC"
	chrPSI = "PSI"
)

type skill struct {
	skill string
	spec  string
	val   int
}

type character struct {
	name            string
	characteristics map[string]int
	skills          map[string]int
	//qualifyDM        map[string]int
	qualifyDM        int
	age              int
	education        string
	commisionAllowed bool
	currentTerm      int
	terms            map[string]int
}

//Log -
var Log []string

type testfail struct {
	thing string
}

func main() {
	preapareLists()
	wantAss := careerAgent
	wantRank := 7
	done := false
	for !done {
		char := newCharacter()

		cr := char.CareerLoop()

		// for i := range Log {
		// 	fmt.Println(Log[i])
		// }
		fmt.Println("------")
		fmt.Println(char.toString())
		fmt.Println("Archived Rank: ", cr.careerName, cr.rank)
		if cr.careerName == wantAss && cr.rank >= wantRank {
			done = true
			for i := range Log {
				fmt.Println(Log[i])
			}
		} else {
			Log = nil
		}
	}

	//fmt.Println(char.toString())

	// test := &testfail{"sdf"}
	// fmt.Println(MakeProbe(test, chrDEX, 7))

	// task := NewTask()
	// task.SetParameters(skillAdmin, chrINT)
	// fmt.Println(char.skillCheck(task))

}

func (tst *testfail) ProbeSimpleCheck(chr string, tn int) float64 {
	return 0.018
}

func MakeProbe(cman CheckManager, chr string, tn int) float64 {
	return cman.ProbeSimpleCheck(chr, tn)
}

func log(message string) {
	Log = append(Log, message)
}

func newCharacter() *character {
	utils.RandomSeed()
	char := &character{}
	char.skills = make(map[string]int)
	//char.qualifyDM = make(map[string]int)
	char.terms = make(map[string]int)
	for i := range listSkills {
		char.skills[listSkills[i]] = -1
	}
	char.rollCharacteristics()
	char.chooseBackgroundSkills()
	char.currentTerm = 1
	char.age = 18

	return char
}

func (char *character) chooseBackgroundSkills() {
	times := charDM(char.characteristics[chrEDU]) + 3
	bSkills := utils.PickFewUniqueFromList(listBackgroundSkills, times)
	for i := range bSkills {
		char.train(bSkills[i], 0)
	}

}

func (char *character) toString() string {
	str := ""
	str += "Name: " + char.name + "\n"
	str += "UPP: " + char.charUPP() + "\n"
	str += "Age: " + convert.ItoS(char.age) + "\n"
	str += "Known Skills\n"
	for i := range listSkills {
		if char.skills[listSkills[i]] > -1 {
			str += listSkills[i] + ": " + convert.ItoS(char.skills[listSkills[i]]) + "\n"
		}
	}
	str += "\n"

	return str
}

func (char *character) rollCharacteristics() {
	char.characteristics = make(map[string]int)
	char.characteristics[chrSTR] = utils.RollDice("2d6")
	char.characteristics[chrDEX] = utils.RollDice("2d6")
	char.characteristics[chrEND] = utils.RollDice("2d6")
	char.characteristics[chrINT] = utils.RollDice("2d6")
	char.characteristics[chrEDU] = utils.RollDice("2d6")
	char.characteristics[chrSOC] = utils.RollDice("2d6")
}

func numberCode(i int) string {
	switch i {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	case 0:
		return "0"

	}
	return "X"
}

func charDM(char int) int {
	if char < 1 {
		return -3
	}
	if char < 3 {
		return -2
	}
	if char < 6 {
		return -1
	}
	if char < 9 {
		return 0
	}
	if char < 12 {
		return 1
	}
	if char < 15 {
		return 2
	}
	if char < 18 {
		return 3
	}
	if char < 21 {
		return 4
	}
	return 5
}

func (char *character) charUPP() string {
	upp := ""
	upp += numberCode(char.characteristics[chrSTR])
	upp += numberCode(char.characteristics[chrDEX])
	upp += numberCode(char.characteristics[chrEND])
	upp += numberCode(char.characteristics[chrINT])
	upp += numberCode(char.characteristics[chrEDU])
	upp += numberCode(char.characteristics[chrSOC])
	if char.characteristics[chrPSI] > 0 {
		upp += "-" + numberCode(char.characteristics[chrSOC])
	}

	return upp
}

func (char *character) setChr(chr string, newVal int) {
	char.characteristics[chr] = newVal
}

func (char *character) changeChrBy(chr string, delta int) {
	current := char.characteristics[chr]
	char.setChr(chr, current+delta)
}
