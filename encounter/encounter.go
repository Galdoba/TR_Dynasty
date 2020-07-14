package encounter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

/*
1. Terrain Dm Chart (type DM, Size Dm) - 			ok
2. Roll Creature Movement (move Type,  Size DM) - 	ok
3. Rural Encounter Table (Creature Type)			ok
4. Animal Type with typeDM							ok
5. Size Table with Size DM (Weight, S,D,E)			ok
6. Weapon Table (with animal type DM)
7. Armor Table (with animal type DM)
8. roll Pack
9. roll Instinct
10. distribute skills

*/

func testmatr() {
	mat := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	fmt.Println("test", mat[0][1])

}

func sizeTable() [][]string {
	sizeTable := [][]string{
		{"2d6", "Weight (kg)", "Strenght", "Dexterity", "Endurance"},
		{"1 or less", "1", "1d1", "1d6", "1d1"},
		{"2", "3", "2d1", "1d6", "2d1"},
		{"3", "6", "1d6", "2d6", "1d6"},
		{"4", "12", "1d6", "2d6", "1d6"},
		{"5", "25", "2d6", "3d6", "2d6"},
		{"6", "50", "2d6", "4d6", "2d6"},
		{"7", "100", "3d6", "3d6", "3d6"},
		{"8", "200", "3d6", "3d6", "3d6"},
		{"9", "400", "4d6", "2d6", "4d6"},
		{"10", "800", "4d6", "2d6", "4d6"},
		{"11", "1600", "5d6", "2d6", "5d6"},
		{"12", "3200", "6d6", "1d6", "6d6"},
		{"13", "5000", "7d6", "1d6", "7d6"},
	}
	return sizeTable
}

func weaponArmorTable() [][]string {
	weaponArmorTable := [][]string{
		{"2d6", "Weapon", "Armor"},
		{"1", "None", "0"},
		{"2", "Teeth", "0"},
		{"3", "Horns", "0"},
		{"4", "Hooves", "1"},
		{"5", "Hooves and Teeth", "1"},
		{"6", "Teeth", "2"},
		{"7", "Claws +1", "2"},
		{"8", "Stinger +1", "3"},
		{"9", "Thrasher +1", "3"},
		{"10", "Claws and Teeth +2", "4"},
		{"11", "Claws +2", "4"},
		{"12", "Teeth +2", "5"},
		{"13", "Thrasher +2", "5"},
	}
	return weaponArmorTable
}

func sizeTableRoll(sizeDM int) (weight, str, dex, end int) {

	tableRoll := utils.RollDice("2d6", sizeDM)
	row := boundInt(tableRoll, 1, 13)
	weight = convert.StoI(sizeTable()[row][1])
	str = utils.RollDice(sizeTable()[row][2])
	dex = utils.RollDice(sizeTable()[row][3])
	end = utils.RollDice(sizeTable()[row][4])
	return weight, str, dex, end
}

func boundInt(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}

type creature struct {
	creatureType string
	behavior     string
	moveType     string
	char         map[string]int
	skill        map[string]int
	weight       int
	weapon       string
	damage       string
	size         int
	armor        int
	hp           int
}

func NewEncounter() {
	// rangeBand := rollEncounterDistance()
	// //exactRange := rollExactDistance()
	// fmt.Println("Range:", rangeBand)
	creature := &creature{}
	creature.char = make(map[string]int)
	creature.skill = make(map[string]int)

	typeDM, sizeDM, moveType := terrainDM()
	creature.moveType = moveTypeStr(moveType)
	creature.creatureType = creatureType()
	wei, str, dex, end := sizeTableRoll(sizeDM)
	creature.char["str"] = str
	creature.char["dex"] = dex
	creature.char["end"] = end
	creature.char["inst"] = utils.RollDice("2d6")
	creature.char["pack"] = utils.RollDice("2d6")
	creature.behavior = animalType(typeDM, creature.creatureType)
	creature.setBehaviourBonus()
	creature.randomSkillsForCreature()
	creature.weight = utils.RollDice("1d"+convert.ItoS(wei), wei/2)
	weaponStr, armorStr := rollWeaponArmor(creature.creatureType)
	creature.weapon = weaponStr
	creature.armor = convert.StoI(armorStr)
	creature.hp = creature.char["str"] + creature.char["end"]
	creature.report()
}

func (crtr *creature) report() {
	//fmt.Println("Type =", crtr.creatureType)
	fmt.Println("Behaviour =", crtr.creatureType+", "+crtr.behavior)
	fmt.Println("Movement =", crtr.moveType)
	fmt.Println("Speed =", crtr.char["dex"], "m")
	fmt.Println("Hits =", crtr.hp)
	fmt.Println("Weight =", crtr.weight, "kg")
	var skills []string
	for key, val := range crtr.skill {
		skills = append(skills, key+" "+convert.ItoS(val-1))
	}
	sort.Strings(skills)
	fmt.Println("Skills:")
	for i := range skills {
		fmt.Println(" " + skills[i])
	}
	fmt.Println("Traits:", " Armor", crtr.armor)

	fmt.Println("Attack:", crtr.weapon+"  ("+damageDice(crtr)+")")
}

func damageDice(crtr *creature) string {
	dd := crtr.char["str"]/10 + 1
	if strings.Contains(crtr.weapon, "+1") {
		dd++
	}
	if strings.Contains(crtr.weapon, "+2") {
		dd++
		dd++
	}
	return convert.ItoS(dd) + "d6"
}

func rollWeaponArmor(crtrType string) (string, string) {
	dm := 0
	switch crtrType {
	case "Herbivore":
		dm = -6
	case "Omnivore":
		dm = 4
	case "Carnivore":
		dm = 8

	}
	r1 := boundInt(utils.RollDice("2d6", dm), 1, 13)
	r2 := boundInt(utils.RollDice("2d6", dm), 1, 13)

	return weaponArmorTable()[r1][1], weaponArmorTable()[r2][2]
}

func (crtr *creature) randomSkillsForCreature() {
	r := utils.RollDice("d6")
	crtr.addSkill("Survival")
	crtr.addSkill("Athletics")
	crtr.addSkill("Recon")
	crtr.addSkill("Melee")
	for i := 0; i < r; i++ {
		for key, _ := range crtr.skill {
			crtr.addSkill(key)
			break
		}
	}

}

func (crtr *creature) setBehaviourBonus() {
	switch crtr.behavior {
	case "Carrion-Eater":
		crtr.addSkill("Recon")
		crtr.addChar("inst", 2)
	case "Chaser":
		crtr.addSkill("Athletics")
		crtr.addChar("dex", 4)
		crtr.addChar("inst", 2)
		crtr.addChar("pack", 2)
	case "Eater":
		crtr.addChar("end", 4)
		crtr.addChar("pack", 2)
	case "Filter":
		crtr.addChar("end", 4)
	case "Gatherer":
		crtr.addSkill("Stealth")
		crtr.addChar("pack", 2)
	case "Grazer":
		crtr.addChar("inst", 2)
		crtr.addChar("pack", 4)
	case "Hunter":
		crtr.addSkill("Survival")
		crtr.addChar("inst", 2)
	case "Hijacker":
		crtr.addChar("str", 2)
		crtr.addChar("pack", 2)
	case "Intimidator":
		crtr.addSkill("Persuade")
	case "Killer":
		crtr.addSkill("Melee")
		if utils.RandomBool() {
			crtr.addChar("str", 4)
		} else {
			crtr.addChar("dex", 4)
		}
		crtr.addChar("inst", 4)
		crtr.addChar("pack", -2)
	case "Intermittent":
		crtr.addChar("pack", 4)
	case "Pouncer":
		crtr.addSkill("Recon")
		crtr.addSkill("Stealth")
		crtr.addSkill("Athletics")
		crtr.addChar("dex", 4)
		crtr.addChar("inst", 4)
	case "Reducer":
		crtr.addChar("pack", 4)
	case "Siren":
		crtr.addSkill("Deception")
		crtr.addChar("pack", -4)
	case "Trapper":
		crtr.addChar("pack", -2)
	default:
		fmt.Println(crtr.behavior, " --- bonus")
	}
}

func (crtr *creature) addSkill(sk string) {
	crtr.skill[sk] = crtr.skill[sk] + 1
}

func (crtr *creature) addChar(sk string, val int) {
	crtr.char[sk] = crtr.char[sk] + val
}

func rollEncounterDistance() string {
	rangeBand := ""
	dm := 0
	opt, _ := utils.TakeOptions("Pick Enviroment:",
		"Clear Terrain   (DM+3)",
		"Forest or Woods (DM-2)",
		"Crowded Area    (DM-2)",
		"In Space        (DM+4)",
	)
	switch opt {
	case 1:
		dm = 3
	case 2, 3:
		dm = -2
	case 4:
		dm = 4
	}
	r := utils.RollDice("2d6", dm)
	if utils.InRange(r, -10, 2) {
		rangeBand = "Close"
	}
	if utils.InRange(r, 3, 3) {
		rangeBand = "Short"
	}
	if utils.InRange(r, 4, 5) {
		rangeBand = "Medium"
	}
	if utils.InRange(r, 6, 9) {
		rangeBand = "Long"
	}
	if utils.InRange(r, 10, 11) {
		rangeBand = "Very Long"
	}
	if utils.InRange(r, 12, 30) {
		rangeBand = "Distant"
	}
	return rangeBand
}

func terrainDM() (typeDM, sizeDM, moveType int) {
	mvStr := ""
	opt, _ := utils.TakeOptions("Pick Enviroment:",
		"Clear Terrain          (typeDM = 3, sizeDM = 0)",
		"Plain or Prerie        (typeDM = 4, sizeDM = 0)",
		"Desert (hot or cold)   (typeDM = 3, sizeDM =-3)",
		"Hills, Foothils        (typeDM = 0, sizeDM = 0)",
		"Mountains              (typeDM = 0, sizeDM = 0)",
		"Forest                 (typeDM =-4, sizeDM =-4)",
		"Woods                  (typeDM =-2, sizeDM =-1)",
		"Jungle                 (typeDM =-4, sizeDM =-3)",
		"Rainforest             (typeDM =-2, sizeDM =-2)",
		"Rough, Broken         (typeDM =-3, sizeDM =-3)",
		"Swamp, Marsh          (typeDM =-2, sizeDM = 4)",
		"Beach, Shore          (typeDM = 3, sizeDM = 2)",
		"Riverbank             (typeDM = 1, sizeDM = 1)",
		"Ocean shallows        (typeDM = 4, sizeDM = 1)",
		"Open ocean            (typeDM = 4, sizeDM = -4)",
		"Deep ocean            (typeDM = 4, sizeDM = 2)",
		"RANDOM                (typeDM = X, sizeDM = X)",
	)
	if opt == 17 {
		opt = utils.RollDice("d16")
		fmt.Println("Random Pick =", opt)
	} else {
		fmt.Println("Pick =", opt)
	}
	switch opt {
	case 1:
		typeDM = 3
		sizeDM = 0
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 0", "W 2", "F -6"})
	case 2:
		typeDM = 4
		sizeDM = 0
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 2", "W 4", "F -6"})
	case 3:
		typeDM = 3
		sizeDM = -3
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 0", "F -4", "F -6"})
	case 4:
		typeDM = 0
		sizeDM = 0
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 2", "F -4", "F -6"})
	case 5:
		typeDM = 0
		sizeDM = 0
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "F -2", "F -4", "F -6"})
	case 6:
		typeDM = -4
		sizeDM = -4
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 0", "F -4", "F -6"})
	case 7:
		typeDM = -2
		sizeDM = -1
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 0", "W 0", "F -6"})
	case 8:
		typeDM = -4
		sizeDM = -3
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 0", "W 2", "F -6"})
	case 9:
		typeDM = -2
		sizeDM = -2
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 2", "W 4", "F -6"})
	case 10:
		typeDM = -3
		sizeDM = -3
		mvStr = utils.RandomFromList([]string{"W 0", "W 0", "W 0", "W 2", "F -4", "F -6"})
	case 11:
		typeDM = -2
		sizeDM = 4
		mvStr = utils.RandomFromList([]string{"S -6", "A 2", "W 0", "W 0", "F -4", "F -6"})
	case 12:
		typeDM = 3
		sizeDM = 2
		mvStr = utils.RandomFromList([]string{"S 1", "A 2", "W 0", "W 0", "F -4", "F -6"})
	case 13:
		typeDM = 1
		sizeDM = 1
		mvStr = utils.RandomFromList([]string{"S -4", "A 0", "W 0", "W 0", "W 0", "F -6"})
	case 14:
		typeDM = 4
		sizeDM = 1
		mvStr = utils.RandomFromList([]string{"S 4", "S 2", "S 0", "S 0", "F -4", "F -6"})
	case 15:
		typeDM = 4
		sizeDM = -4
		mvStr = utils.RandomFromList([]string{"S 6", "S 4", "S 2", "S 0", "F -4", "F -6"})
	case 16:
		typeDM = 4
		sizeDM = 2
		mvStr = utils.RandomFromList([]string{"S 8", "S 6", "S 4", "S 2", "S 0", "S -2"})
	}
	moveParts := strings.Split(mvStr, " ")
	switch moveParts[0] {
	case "A":
		moveType = 0
	case "S":
		moveType = 1
	case "F":
		moveType = 2
	case "W":
		moveType = 3
	}
	sizeDM = sizeDM + convert.StoI(moveParts[1])
	return typeDM, sizeDM, moveType
}

func moveTypeStr(moveType int) string {
	crType := ""
	switch moveType {
	case 0:
		crType = "Amphibious"
	case 1:
		crType = "Swimmer"
	case 2:
		crType = "Flyer"
	case 3:
		crType = "Walker"
	default:
		crType = "UNKNOWN"
	}
	return crType
}

func creatureType() string {
	crType := ""
	r := utils.RollDice("2d6")
	switch r {
	case 2, 4:
		crType = "Scavenger"
	case 3, 5:
		crType = "Omnivore"
	case 6, 7, 8:
		crType = "Herbivore"
	case 9, 11, 12:
		crType = "Carnivore"
	default:
		crType = "Strange"
	}
	return crType
}

func animalType(typeDM int, creatureType string) string {
	animalType := ""
	r := utils.RollDice("2d6", typeDM)
	switch creatureType {
	case "Herbivore":
		animalType = herbivoreType(r)
	case "Omnivore":
		animalType = omnivoreType(r)
	case "Carnivore":
		animalType = carnivoreType(r)
	case "Scavenger":
		animalType = scavengerType(r)
	}
	return animalType
}

func herbivoreType(i int) string {
	aType := ""
	switch i {
	case 3, 4, 5, 6:
		aType = "Intermittent"
	default:
		if i < 3 {
			aType = "Filter"
		}
		if i > 6 {
			aType = "Grazer"
		}
	}
	return aType
}

func omnivoreType(i int) string {
	aType := ""
	switch i {
	case 3, 5, 9, 12:
		aType = "Gatherer"
	case 2, 4, 10:
		aType = "Eater"
	case 6, 7, 8, 11:
		aType = "Hunter"
	default:
		aType = "Gatherer"
	}
	return aType
}

func carnivoreType(i int) string {
	aType := ""
	switch i {
	case 2, 12:
		aType = "Siren"
	case 3, 6:
		aType = "Pouncer"
	case 4, 10:
		aType = "Killer"
	case 5:
		aType = "Trapper"
	case 7, 8, 9, 11:
		aType = "Chaser"
	default:
		if i < 2 {
			aType = "Pouncer"
		}
		if i > 12 {
			aType = "Chaser"
		}
	}
	return aType
}

func scavengerType(i int) string {
	aType := ""
	switch i {
	case 2, 6, 8, 11:
		aType = "Reducer"
	case 3, 9, 12:
		aType = "Hijacker"
	case 4, 7:
		aType = "Carrion-Eater"
	case 5, 10:
		aType = "Intimidator"
	default:
		if i < 2 {
			aType = "Carrion-Eater"
		}
		if i > 12 {
			aType = "Intimidator"
		}
	}
	return aType
}
