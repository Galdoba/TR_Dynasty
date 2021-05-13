package entity

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/core/skill"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/entity/asset"
	"github.com/Galdoba/TR_Dynasty/pkg/entity/career"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

type Traveller struct {
	Dice       *dice.Dicepool
	ManualMode bool
	Info       Info
	Chrctr     map[string]asset.Characteristic
	Skill      map[string]asset.Skill
	CareerMap  map[string]career.Career
}

type Info struct {
	Name        string
	Rads        int
	PsiStatus   string
	Species     string
	Education   string
	LastTerm    string
	TermsServed int
	TermsLog    []string
	TermsLogI   []int
	careerRank  map[string]int
}

func NewTraveller(seed ...string) Traveller {
	t := Traveller{}
	t.Info = Info{}
	t.Dice = dice.New()
	for i := range seed {
		switch i {
		case 0:
			t.Dice.SetSeed(seed[0])
		}
	}
	t.Info.Name = "Mark Miller"
	t.Info.PsiStatus = "Untested"
	t.Info.Species = "Human" //TODO: должно устанавливаться на прямую флагом с рассой
	t.rollCharcteristics()
	t.pickBackgroundSkills()
	allTerms := career.Allterms()
	maxTerms := 5
	d := strconv.Itoa(len(allTerms))
	r := t.Dice.RollNext("1d" + d).DM(-1).Sum()
	fmt.Println("pick ", allTerms[r])
	t.Info.careerRank = make(map[string]int)
	careerEnded := false
termLoop:
	for !careerEnded {
		fmt.Println("Term left", maxTerms)

		fmt.Print("Qualification:")
		switch ListContainsInt(t.Info.TermsLogI, allTerms[r]) {
		case true:
			fmt.Println("No Qualification Needed")
		case false:
			switch t.Dice.RollNext("2d6").Sum() > 6 {
			case true:
				fmt.Println("PASSED")
				t.Info.TermsLogI = append(t.Info.TermsLogI, allTerms[r])
			case false:
				fmt.Println("FAILED")
				continue
			}
		}
		fmt.Print("Survival:")
		maxTerms--
		switch t.Dice.RollNext("2d6").Sum() > 6 {
		case true:
			t.Info.TermsServed++
			t.Info.TermsLogI = append(t.Info.TermsLogI, allTerms[r])
			fmt.Println("PASSED")
			fmt.Println("EVENT ROLL")
		case false:
			fmt.Println("FAILED")
			fmt.Println("MISHAP ROLL")
			break termLoop
		}
		fmt.Print("Advansment:")
		switch t.Dice.RollNext("2d6").Sum() > 6 {
		case true:
			t.Info.TermsServed++
			t.Info.TermsLogI = append(t.Info.TermsLogI, allTerms[r])
			fmt.Println("PASSED")
			fmt.Println("EVENT ROLL")
		case false:
			fmt.Println("FAILED")
			fmt.Println("MISHAP ROLL")

			continue
		}
	}

	return t
}

func ListContainsInt(list []int, num int) bool {
	for _, val := range list {
		if val == num {
			return true
		}
	}
	return false
}

func (t *Traveller) rollCharcteristics() {
	t.Chrctr = make(map[string]asset.Characteristic)
	for _, val := range listCharacteristics() {
		score := t.Dice.RollNext("2d6").Sum()
		t.Chrctr[val] = asset.NewCharacteristic(val)
		t.Chrctr[val].SetCharacteristicValue(score)
	}
	if t.Dice.RollNext("2d6").Sum() > 10 {
		score := t.Dice.RollNext("2d6").Sum()
		t.Chrctr[PSI] = asset.NewCharacteristic(PSI)
		t.Chrctr[PSI].SetCharacteristicValue(score)
	}

	switch t.Info.Species {
	default:
	}
}

func (t *Traveller) pickBackgroundSkills() {
	t.Skill = make(map[string]asset.Skill)
	mod := t.Chrctr[EDU].Modifier()
	mod += 3
	picked := []string{}
	list := asset.BackgroundSkills()
	switch t.ManualMode {
	case false:
		for len(picked) < mod {
			picked = utils.AppendUniqueStr(picked, t.Dice.RollFromList(list))
		}
	case true:
		fmt.Println("func (t *Traveller) pickBackgroundSkills() - manual mode not implemented")
	}
	for i := range picked {
		t.Skill[picked[i]] = asset.BasicTraining(picked[i])
	}
}

func (t *Traveller) Sheet() string {
	//utils.ClearScreen()
	tName := FormatString(t.Info.Name, 27, false)       //set len to 27 left al
	tUPP := FormatString(t.UPPsheetString(), 27, false) //set len to 27 left al
	tRads := FormatInt(t.Info.Rads, 5, false)           //set len to 4 left al
	tAge := FormatInt(t.Age(), 3, false)                //set len to 3 left al
	tSpecies := FormatString(t.Info.Species, 24, false) //set len to 24 left al
	tSpeciesTraits := SpeciesTraitsSheet(t.Info.Species)
	skillList := listAllSkills(t)
	careerList := listAllCareers(t)

	sh := "+---INFO----------------------------+---ARMOR---------------------------------------------------------------------------+\n"
	sh += "| Name: " + tName + " | TYPE              | RAD | PROTECTION | KG |             INSTALLED MODS            |\n"
	sh += "| UPP : " + tUPP + " | Armor name 1      | XXX |     XX     | XX | [Loooooooooooooooooooong Description] |\n"
	sh += "| Rads: " + tRads + "            Age: " + tAge + "   | Armor name 2      | XXX |     XX     | XX | Options:             [No Description] |\n"
	sh += "| Species: " + tSpecies + " | Armor name 3      | XXX |     XX     | XX | Options:             [No Description] |\n"
	sh += "| Species Traits: " + tSpeciesTraits[0] + " | Armor name 4      | XXX |     XX     | XX | Options:             [No Description] |\n"
	//sh += "=                 _Additionals_____ =  ____Additional Armor data__              =                                       =\n"
	sh += "| Homeworld: [Homeworld Name      ] +---FINANCES------------+---CAREER SUMMARY--+---CAREER BENEFITS---------------------+\n"
	sh += "| s123456-7 __ __ __ __ __ __ __ __ | Pension:              | " + careerList[0] + " | 1234567890123456789012345678901234567 |\n"
	sh += "+---CHARACTERISTICS-----------------+   XXXXXX Cr/Year      | " + careerList[1] + " |                                       |\n"
	sh += "|" + AtrBox(t.Chrctr[STR]) + "|" + AtrBox(t.Chrctr[DEX]) + "|" + AtrBox(t.Chrctr[END]) + "| Debt:                 | " + careerList[2] + " |                                       |\n"
	sh += "| Strength  | Dexterity | Endurance |   XXXXXXX xCr         | " + careerList[3] + " |                                       |\n"
	sh += "+-----------+-----------+-----------+ Cash on Hand:         | " + careerList[4] + " |                                       |\n"
	sh += "|" + AtrBox(t.Chrctr[INT]) + "|" + AtrBox(t.Chrctr[EDU]) + "|" + AtrBox(t.Chrctr[SOC]) + "|   XXXXXXX xCr         | " + careerList[5] + " |                                       |\n"
	sh += "| Intellect | Education |   Social  | Living Cost:          | " + careerList[6] + " |                                       |\n"
	sh += "+-----------+-----------+-----------+   XXXXXXX  Cr/Month   | " + careerList[7] + " |                                       |\n"
	sh += "| Psionic Powers: UNTESTED[XX] (-3) |                       | " + careerList[8] + " |                                       |\n"
	sh += "=        [Untested or talents list] =                       =                   = __Additional career benefits__        |\n"
	sh += "+---SKILLS--------------------------+-----------------------+-------------------+---------------------------------------+\n"
	third := (len(skillList) / 4) + 1
	//fmt.Println(len(skillList), "|")
	for len(skillList) < third*4 {
		skillList = append(skillList, "                           ")
	}
	for i := range skillList {
		if i >= third {
			continue
		}
		//fmt.Println(third, len(skillList), "|", i, third+i, (third+third)+i, third+third+third+i)
		sh += "| " + skillList[i] + " | " + skillList[third+i] + " | " + skillList[(third+third)+i] + " | " + skillList[(third+third+third)+i] + " |\n"
	}
	sh += "+-----------------------------------+-----------------------+-------------------+---------------------------------------+\n"
	return sh
}

/*
+---INFO----------------------------+---ARMOR---------------------------------------------------------------------------+
| Name: [Traveller Name]            | TYPE              | RAD | PROTECTION | KG |             INSTALLED MODS            |
| UPP : 123456-7                    | Armor name 1      | XXX |     XX     | XX | [Loooooooooooooooooooong Description] |
| Rads: xxxx             Age: XXX   | Armor name 2      | XXX |     XX     | XX | Options:             [No Description] |
| Species: Human                    | Armor name 3      | XXX |     XX     | XX | 1234567890123456789012345678901234567 |
| Species Traits: _Mandatory_______ | Armor name 4      | XXX |     XX     | XX | Options:             [No Description] |
=                 _Additionals_____ =  ____Additional Armor data__              =                                       =
| Homeworld: [Homeworld Name      ] +---FINANCES------------+---CAREER SUMMARY--+---CAREER BENEFITS---------------------+
| s123456-7 __ __ __ __ __ __ __ __ | Pension:              | Law Enforcement 2 | 1234567890123456789012345678901234567 |
+---CHARACTERISTICS-----------------+   XXXXXX Cr/Year      | Scavenger       1 |                                       |
| [XX] (+0) | [XX} (+0) | [XX] (-1) | Debt:                 | Marine Support  1 |                                       |
|  Strengh  | Dexterity | Endurance |   XXXXXXX xCr         | Drifter         8 |                                       |
+-----------+-----------+-----------+ Cash on Hand:         |                   |                                       |
| [XX] (+0) | [XX} (+0) | [XX] (-1) |   XXXXXXX xCr         |                   |                                       |
| Intellect | Education |   Social  | Living Cost:          |                   |                                       |
+-----------+-----------+-----------+   XXXXXXX  Cr/Month   |                   |                                       |
| Psionic Powers: UNTESTED[XX] (-3) |                       |                   |                                       |
=        [Untested or talents list] =   XXXXXXX  Cr/Month   =                   = __Additional career benefits__        |
+---SKILLS--------------------------+-----------------------+-------------------+---------------------------------------+
= Jack of all Trades          =                             =                             =                             =  Alphabetical Order topdown
= Pilot (capital ships) x     =                             =                             =                             =
| 123456789012345678901234567 | 123456789012345678901234567 | 123456789012345678901234567 | 123456789012345678901234567 |
+-----------------------------------------------------------------------------------------------------------------------+
123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890

*/

func FormatString(str string, maxlen int, alighnRight bool) string {
	switch alighnRight {
	case true:
		for len(str) < maxlen {
			str = " " + str
		}
	case false:
		for len(str) < maxlen {
			str = str + " "
		}
	}
	if len(str) > maxlen {
		str = string([]byte(str)[0:maxlen])
	}
	return str
}

func FormatInt(data int, maxlen int, alighnRight bool) string {
	str := strconv.Itoa(data)
	return FormatString(str, maxlen, alighnRight)
}

func (t *Traveller) UPPsheetString() string {
	upp := ""
	for _, chr := range listCharacteristics() {
		if val, ok := t.Chrctr[chr]; ok {
			v := val.Value()
			upp += ehex.New(v).String()
		}
	}
	if val, ok := t.Chrctr[PSI]; ok { //PSI is a rare case
		v := val.Value()
		upp += "-" + ehex.New(v).String()
	}

	return upp
}

func (t *Traveller) Age() int {
	return 18 + t.Info.TermsServed*4
}

func (t *Traveller) AtrArray() []int {
	aa := []int{}
	for _, a := range listCharacteristics() {
		v := t.Chrctr[a].Value()
		aa = append(aa, v)
	}
	if psi, ok := t.Chrctr[PSI]; ok {
		psiVal := psi.Value()
		aa = append(aa, psiVal)
	}
	return aa
}

func AtrBox(a asset.Characteristic) string {
	v := a.Value()
	str := "  "
	if v < 10 && v > -1 {
		str += " "
	}
	str += strconv.Itoa(v) + "  ("
	aM := a.Modifier()
	if aM > -1 {
		str += "+"
	}
	str += strconv.Itoa(aM) + ") "
	return str
}

func SpeciesTraitsSheet(race string) []string {
	traits := listSpeciesTraits(race)
	for i := range traits {
		traits[i] = FormatString(traits[i], 17, false)
	}
	return traits
}

func listAllSkills(t *Traveller) []string {
	list := []string{}
	for _, v := range t.Skill {
		specs, vals := v.Specialities()
		for i := range specs {
			for len(specs[i]) < 25 {
				specs[i] += " "
			}
			list = append(list, specs[i]+" "+strconv.Itoa(vals[i]))
		}
	}
	sort.Strings(list)
	return list
}

func listAllCareers(t *Traveller) []string {
	allList := []string{}
	for _, val := range t.Info.TermsLog {
		allList = append(allList, FormatString(val, 17, false))
	}
	for len(allList) < 9 {
		allList = append(allList, FormatString("", 17, false))
	}
	return allList
}

func (t *Traveller) SkillPresent(sk string) bool {
	for _, skil := range t.Skill {
		spec, _ := skil.Specialities()
		for _, spk := range spec {
			if spk == sk {
				return true
			}
		}
	}
	return false
}

func (t *Traveller) Train(thing string) error {
	switch thing {
	case STR, DEX, END, INT, EDU, SOC:
		val := t.Chrctr[thing].Value()
		t.Chrctr[thing].SetCharacteristicValue(val + 1)
		return nil
	}

	if !skill.NameIsValid(thing) {
		panic(errors.New("TrvCore_SkillCode unknown for '" + thing + "'"))
	}
	if !t.SkillPresent(thing) {
		data := strings.Split(thing, " (")
		switch len(data) {
		case 2:
			t.Skill[data[0]] = asset.BasicTraining(data[0])
			return nil
		case 1:
			t.Skill[thing] = asset.BasicTraining(thing)
		}
	}
	for key, v := range t.Skill {
		spec, _ := v.Specialities()
		if thing == key {
			message := "Choose Speciality to train:"
			thing = t.chooseFromList(message, spec)
			return t.Skill[key].Train(thing)
		}
		if utils.ListContains(spec, thing) {
			return t.Skill[key].Train(thing)
		}

	}

	return nil
}

func (t *Traveller) chooseFromList(message string, list []string) string {
	chosen := ""
	switch t.ManualMode {
	case false:
		chosen = t.Dice.RollFromList(list)
	case true:
		chosen = chooseManualy(message, list)
	}
	return chosen
}

func chooseManualy(message string, list []string) string {
	pick, err := user.ChooseOne(message, list)
	if err != nil {
		panic(err.Error())
	}
	return list[pick]
}
