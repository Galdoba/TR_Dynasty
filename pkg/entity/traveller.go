package entity

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/entity/asset"
	"github.com/Galdoba/utils"
)

type Traveller struct {
	Dice       *dice.Dicepool
	ManualMode bool
	Info       Info
	Chrctr     map[string]asset.Characteristic
	Skill      map[string]asset.Skill
}

type Info struct {
	Name        string
	Rads        int
	PsiStatus   string
	Species     string
	TermsServed int
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

	return t
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
	mod, _ := t.Chrctr[EDU].Modifier()
	mod += 3
	picked := []string{}
	//list := asset.BackgroundSkills()
	switch t.ManualMode {
	case false:
		for len(t.Skill) < mod {
			// d := strconv.Itoa(len(list))
			// sk := list[t.Dice.RollNext("1d"+d).DM(-1).Sum()]
			// skillName := skill.ByCode(sk).Name()
			// if _, ok := t.Skill[skillName]; !ok {
			// 	t.Skill[skillName] = asset.NewSkill(skillName)
			// }
		}
	case true:
		fmt.Println("func (t *Traveller) pickBackgroundSkills() - manual mode not implemented")
	}

	for i := range picked {
		t.Skill[picked[i]] = asset.NewSkill(picked[i])
		//t.Skill[picked[i]] = asset.NewSkill(skill.ByCode(i).Name())
	}
}

func (t *Traveller) Sheet() string {
	utils.ClearScreen()
	tName := FormatString(t.Info.Name, 27, false)       //set len to 27 left al
	tUPP := FormatString(t.UPPsheetString(), 27, false) //set len to 27 left al
	tRads := FormatInt(t.Info.Rads, 5, false)           //set len to 4 left al
	tAge := FormatInt(t.Age(), 3, false)                //set len to 3 left al
	tSpecies := FormatString(t.Info.Species, 24, false) //set len to 24 left al
	tSpeciesTraits := SpeciesTraitsSheet(t.Info.Species)

	//trvAtrArr := t.AtrArray()
	// str := FormatInt(trvAtrArr[0], 2, false) //set len to 2 left al
	// strM := SheetMod(str)
	// dex := FormatInt(trvAtrArr[1], 2, false) //set len to 2 left al
	// end := FormatInt(trvAtrArr[2], 2, false) //set len to 2 left al
	// int := FormatInt(trvAtrArr[3], 2, false) //set len to 2 left al
	// edu := FormatInt(trvAtrArr[4], 2, false) //set len to 2 left al
	// soc := FormatInt(trvAtrArr[5], 2, false) //set len to 2 left al

	sh := "+---INFO----------------------------+---ARMOR---------------------------------------------------------------------------+\n"
	sh += "| Name: " + tName + " | TYPE              | RAD | PROTECTION | KG |             INSTALLED MODS            |\n"
	sh += "| UPP : " + tUPP + " | Armor name 1      | XXX |     XX     | XX | [Loooooooooooooooooooong Description] |\n"
	sh += "| Rads: " + tRads + "            Age: " + tAge + "   | Armor name 2      | XXX |     XX     | XX | Options:             [No Description] |\n"
	sh += "| Species: " + tSpecies + " | Armor name 3      | XXX |     XX     | XX | Options:             [No Description] |\n"
	sh += "| Species Traits: " + tSpeciesTraits[0] + " | Armor name 4      | XXX |     XX     | XX | Options:             [No Description] |\n"
	//sh += "=                 _Additionals_____ =  ____Additional Armor data__              =                                       =\n"
	sh += "| Homeworld: [Homeworld Name      ] +---FINANCES------------+---CAREER SUMMARY--+---CAREER BENEFITS---------------------+\n"
	sh += "| s123456-7 __ __ __ __ __ __ __ __ | Pension:              | Law Enforcement 2 | 1234567890123456789012345678901234567 |\n"
	sh += "+---CHARACTERISTICS-----------------+   XXXXXX Cr/Year      | Scavenger       1 |                                       |\n"
	sh += "|" + AtrBox(t.Chrctr[STR]) + "|" + AtrBox(t.Chrctr[DEX]) + "|" + AtrBox(t.Chrctr[END]) + "| Debt:                 | Marine Support  1 |                                       |\n"
	sh += "| Strength  | Dexterity | Endurance |   XXXXXXX xCr         | Drifter         8 |                                       |\n"
	sh += "+-----------+-----------+-----------+ Cash on Hand:         |                   |                                       |\n"
	sh += "|" + AtrBox(t.Chrctr[INT]) + "|" + AtrBox(t.Chrctr[EDU]) + "|" + AtrBox(t.Chrctr[SOC]) + "|   XXXXXXX xCr         |                   |                                       |\n"
	sh += "| Intellect | Education |   Social  | Living Cost:          |                   |                                       |\n"
	sh += "+-----------+-----------+-----------+   XXXXXXX  Cr/Month   |                   |                                       |\n"
	sh += "| Psionic Powers: UNTESTED[XX] (-3) |                       |                   |                                       |\n"
	sh += "=        [Untested or talents list] =                       =                   = __Additional career benefits__        |\n"
	sh += "+---SKILLS--------------------------+-----------------------+-------------------+---------------------------------------+\n"
	sh += "+-----------------------------------+-----------------------+-------------------+---------------------------------------+\n"
	return sh
}

/*
+---INFO----------------------------+---ARMOR---------------------------------------------------------------------------+
| Name: [Traveller Name]            | TYPE              | RAD | PROTECTION | KG |             INSTALLED MODS            |
| UPP : 123456-7                    | Armor name 1      | XXX |     XX     | XX | [Loooooooooooooooooooong Description] |
| Rads: xxxx             Age: XXX   | Armor name 2      | XXX |     XX     | XX | Options:             [No Description] |
| Species: Human                    | Armor name 3      | XXX |     XX     | XX | Options:             [No Description] |
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
			v, _ := val.Value()
			upp += ehex.New(v).String()
		}
	}
	if val, ok := t.Chrctr[PSI]; ok { //PSI is a rare case
		v, _ := val.Value()
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
		v, _ := t.Chrctr[a].Value()
		aa = append(aa, v)
	}
	if psi, ok := t.Chrctr[PSI]; ok {
		psiVal, _ := psi.Value()
		aa = append(aa, psiVal)
	}
	return aa
}

func AtrBox(a asset.Characteristic) string {
	v, _ := a.Value()
	str := "  "
	if v < 10 && v > -1 {
		str += " "
	}
	str += strconv.Itoa(v) + "  ("
	aM, _ := a.Modifier()
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
