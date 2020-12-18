package encounter

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore/ehex"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

const (
	terrainClear        = 0
	terrainPlain        = 1
	terrainDesert       = 2
	terrainHills        = 3
	terrainMountain     = 4
	terrainForest       = 5
	terrainWoods        = 6
	terrainJungle       = 7
	terrainRainforest   = 8
	terrainRoughBroken  = 9
	terrainSwampMarsh   = 10
	terrainBeachShore   = 11
	terrainRiverbank    = 12
	terrainShallowOcean = 13
	terrainOpenOcean    = 14
	terrainDeepOcean    = 15
	chrStrength         = "Strength"
	chrDexterity        = "Dexterity"
	chrEndurance        = "Endurance"
	chrIntelligence     = "Intelligence"
	chrInstinct         = "Instinct"
	chrPack             = "Pack"
	sklSurvival         = "Survival"
	sklAthlethics       = "Athlethics"
	sklDeception        = "Deception"
	sklPersuade         = "Persuade"
	sklRecon            = "Recon"
	sklMeleeNW          = "Melee (natural weapons)"
	sklStealth          = "Stealth"
	dietCarnivore       = "Carnivore"
	dietHerbivore       = "Herbivore"
	dietOmnivore        = "Omnivore"
	rollSocial          = "soc"
	rollPhysical        = "phy"
	rollBehavior        = "beh"
	rollEvolution       = "evo"
	//terrainGasWorld    = 16
	//terrainAsteroid    = 17

)

/*
Determine Terrain
	Determine Planet Conditions //2-я очередь
Determine Class
	Determine if preffered //2-я очередь
Determine Movement Type
Determine Behaviour Model
Numerical Details
	Size
	Weapons
	Armour
	Characteristics
	Skills
	Number Encountered
	Reaction Check
*/

func Test() {
	fmt.Println("Test Begin")
	// an := NewAnimal("B757754-9", -1, -1)
	// fmt.Println(AnimalBreedSheet(an))
	// fmt.Println("SHORT DESCR:\n", ShortDescr(an))
	EncounterTable()
	fmt.Println("Test End")
}

func EncounterTable() {
	fmt.Print("UWP: ")
	uwp, err := user.InputStr()
	fmt.Print("Biome: ")
	ter, err := user.InputInt()
	fmt.Print("Class: ")
	cl, err := user.InputInt()
	if err != nil {
		panic(2)
	}
	animalMap := make(map[int]animal)
	for i := 2; i <= 12; i++ {
		if i == 7 {
			fmt.Print("#7	Event\n")
			continue
		}
		animalMap[i] = NewAnimal(uwp+strconv.Itoa(i), ter, cl)
		fmt.Print("#", i, "	", ShortDescr(animalMap[i]), "\n")

	}
	fmt.Print("Select Animal #")
	num, err := user.InputInt()
	if an, ok := animalMap[num]; ok {
		fmt.Println(AnimalBreedSheet(an))
	} else {
		fmt.Println("No animal with #" + strconv.Itoa(num))
	}
	fmt.Print("Press 'ENTER' to exit")
	uwp, err = user.InputStr()
}

func ShortDescr(an animal) string {
	damage := strconv.Itoa(an.damageDice) + "d6"
	if an.damageMod > 0 {
		damage = damage + "+" + strconv.Itoa(an.damageMod)
	}
	sd := an.diet + " " + an.behaviour + " " + classString(an.classification) + "   " +
		ehex.New(an.characteristics[chrStrength]).String() +
		ehex.New(an.characteristics[chrDexterity]).String() +
		ehex.New(an.characteristics[chrEndurance]).String() +
		ehex.New(an.characteristics[chrIntelligence]).String() +
		ehex.New(an.characteristics[chrInstinct]).String() +
		ehex.New(an.characteristics[chrPack]).String() +
		"   Armor " + ehex.New(an.armorScore).String() + "   " + an.weapon + "/" + damage +
		"   F" + an.fleeIF + "/A" + an.attackIF
	return sd
}

func AnimalBreedSheet(an animal) string {
	abs := "ANIMAL BREED SHEET:\n"
	abs += "Species Name: <UNKNOWN>\n"
	abs += "Classification: " + classString(an.classification) + "\n"
	abs += "Preffered Terrain: " + terrainString(an.prefferedTerrain) + "\n"
	abs += "Diet: " + an.diet + "\n"
	abs += "Behaviour: " + an.behaviour + "\n"
	abs += "Movement Type: " + an.movementType + "\n"
	abs += "CHARACTERISTICS:\n"
	abs += "Strenght     : " + charPretty(an.characteristics[chrStrength]) + charDM(an.characteristics[chrStrength]) + "\n"
	abs += "Dexterity    : " + charPretty(an.characteristics[chrDexterity]) + charDM(an.characteristics[chrDexterity]) + "\n"
	abs += "Endurance    : " + charPretty(an.characteristics[chrEndurance]) + charDM(an.characteristics[chrEndurance]) + "\n"
	abs += "Intelligence : " + charPretty(an.characteristics[chrIntelligence]) + charDM(an.characteristics[chrIntelligence]) + "\n"
	abs += "Instinct     : " + charPretty(an.characteristics[chrInstinct]) + charDM(an.characteristics[chrInstinct]) + "\n"
	abs += "Pack         : " + charPretty(an.characteristics[chrPack]) + charDM(an.characteristics[chrPack]) + "\n"
	abs += "HITS: " + charPretty(an.characteristics[chrStrength]+an.characteristics[chrDexterity]+an.characteristics[chrEndurance]) + "\n"
	s := strconv.FormatFloat(an.weight, 'f', -1, 64)
	abs += "Size: " + charPretty(an.size) + " (" + s + " kg)\n"
	abs += "SKILLS: \n"
	for key := range an.skills {
		skName := key
		for len(skName) < 24 {
			skName += " "
		}
		abs += skName + charPretty(an.skills[key]) + "\n"
	}
	abs += "ENCOUNTER DATA: \n"
	rollNumber := strings.TrimSuffix(an.numbersEncountered, " x 3")
	mult := 1
	if strings.Contains(an.numbersEncountered, " x 3") {
		mult = 3
	}
	numEnc := dice.Roll(rollNumber).Sum() * mult
	abs += "Pack Size: " + an.numbersEncountered + " (" + strconv.Itoa(numEnc) + ")\n"
	damage := strconv.Itoa(an.damageDice) + "d6"
	if an.damageMod > 0 {
		damage = damage + "+" + strconv.Itoa(an.damageMod)
	}
	abs += "Weapon: " + an.weapon + " (damage: " + damage + ")\n"
	for i := range an.exoticWeapon {
		abs += "Exotic Weapon: " + an.exoticWeapon[i] + "\n"
		abs += "Exotic Weapon Effect: " + an.exoticWeaponEffect[i] + "\n"
	}
	abs += "Armor " + strconv.Itoa(an.armorScore) + "\n"
	abs += "Reaction roll: " + dice.Roll("2d6").SumStr() + "\n"
	abs += "Attack: " + an.attackIF + "\n"
	abs += "Flee  : " + an.fleeIF + "\n"
	abs += "NOTES:\n"
	abs += an.notes
	//moveTest := strconv.Itoa(an.size) + charDM(an.characteristics[chrDexterity])
	abs += "\nMovement: " + movement(&an) + "\n"
	abs += "Animal is " + edibility(an) + " and tastes " + taste(an) + " with most nutrient part being " + potencial(an) + "\n"
	abs += "Training is " + training(an) + "\n"
	return abs
}

func movement(an *animal) string {
	base := ((an.size / 2) + 2) * 100
	base += (charDMint(an.characteristics[chrDexterity]) * 100)
	if strings.Contains(an.movementType, "25%") {
		base = base / 4
	}
	if strings.Contains(an.movementType, "50%") {
		base = base / 2
	}
	if strings.Contains(an.movementType, "S") {
		base = base * 2
	}
	if strings.Contains(an.movementType, "F") {
		base = base * 3
	}
	base = base / 100
	if base < 0 {
		base = 0
	}
	return strconv.Itoa(base) + " m"
}

func roundFloat(fl float64) float64 {
	return float64(int(fl*10)) / 10.0
}

func charPretty(chr int) string {
	chrStr := ""

	if chr < 10 && chr > -1 {
		chrStr += " "
	}
	return chrStr + strconv.Itoa(chr)
}

func charDMint(chr int) int {
	dmi := 0
	if chr < 0 {
		chr = 0
	}
	switch chr {
	case 0:
		dmi = -3
	case 1, 2:
		dmi = -2
	case 3, 4, 5:
		dmi = -1
	default:
		dmi = (chr / 3) - 2
	}
	return dmi
}
func charDM(chr int) string {
	dm := " ("
	dmi := 0
	if chr < 0 {
		chr = 0
	}
	switch chr {
	case 0:
		dmi = -3
	case 1, 2:
		dmi = -2
	case 3, 4, 5:
		dmi = -1
	default:
		dmi = (chr / 3) - 2

		dm += "+"
	}
	dm += strconv.Itoa(dmi) + ")"
	return dm
}

type animal struct {
	dicepool           *dice.Dicepool
	name               string
	classification     int
	prefferedTerrain   int
	worldAtmo          int
	worldGravity       int
	worldHumidity      int
	characteristics    map[string]int //str/dex/end/Intel/instinct/pack
	skills             map[string]int
	diet               string
	behaviour          string
	movementType       string
	packSize           int
	armorType          string
	armorScore         int
	reactionDM         int
	attackIF           string
	fleeIF             string
	attackType         string
	damageDice         int
	damageMod          int
	exoticWeapon       []string
	weapon             string
	exoticWeaponEffect []string
	descr              string
	notes              string
	size               int
	generationRollDm   int
	evolutionRollDM    int
	expectedRolls      []string
	expectedSize       int
	weight             float64
	numbersEncountered string
}

//NewAnimal - Creates New Animal, uwp defines seed and planet features
//territory, class - defines location and class. -2 for pure random, -1 for semi-random
func NewAnimal(uwp string, territory, class int) animal {
	seed64 := utils.SeedFromString(uwp)
	an := animal{}

	an.skills = make(map[string]int)
	an.characteristics = make(map[string]int)
	an.dicepool = dice.New(seed64)
	an.selectTerrain(territory)
	an.selectClass(class)
	an.setMovement()
	an.setDiet()
	an.setAnimalBehaviour()
	an.evolutionBase()
	benefits := an.rollBenefits()
	an.generationRolls()
	an.size = an.size + an.dicepool.RollNext("2d6").Sum()
	an.applyBenefits(benefits...)
	an.defineChrs()
	an.getBehaviourBenefits()
	an.numbersInPack()
	an.attackFleeBeh()
	an.modifyFromPlanet(uwp)
	for strings.Contains(an.notes, "Imposible Animal") {
		an = NewAnimal(uwp+"r", territory, -1)
	}
	return an
}

func edibility(an animal) string {
	switch an.dicepool.FluxNext() {
	default:
		return "Unknown"
	case -5, -4, -3, -2, -1, 0:
		return "Not Edibile"
	case 1, 2:
		return "Marginal Edibility"
	case 3, 4, 5:
		return "Edible"
	}
}

func taste(an animal) string {
	switch an.dicepool.FluxNext() {
	default:
		return "Unknown"
	case -5, -4:
		return "Disgusting"
	case -3:
		return "Offensive"
	case -2:
		return "Bad"
	case -1:
		return "Slightly Off"
	case 0:
		return "Ordinary"
	case 1:
		return "Unusual"
	case 2:
		return "Good"
	case 3:
		return "Quite Tasty"
	case 4:
		return "Delocious"
	case 5:
		return "Exquisite"
	}
}

func potencial(an animal) string {
	switch an.dicepool.FluxNext() {
	default:
		return "Unknown"
	case -5:
		return "Sensory Organs"
	case -4:
		return "Brain"
	case -3:
		return "Skeleton"
	case -2:
		return "Digestive Organs"
	case -1:
		return "Circulary Organs"
	case 0:
		return "Eggs and Reproductive"
	case 1:
		return "Secretions"
	case 2:
		return "Interiour Fluids"
	case 3:
		return "Respiratory Organs"
	case 4:
		return "Outer Coverings"
	case 5:
		return "Waste Process Organs"
	}
}

func training(an animal) string {
	switch an.dicepool.FluxNext() {
	default:
		return "Unknown"
	case -5, -4, -3, -2, -1:
		return "Not Possible"
	case 0:
		return "Almost Impossible (16)"
	case 1:
		return "Formidable (14)"
	case 2:
		return "Very Difficult (12)"
	case 3:
		return "Difficult (10)"
	case 4:
		return "Average (8)"
	case 5:
		return "Routine (6)"
	}
}

func (an *animal) modifyFromPlanet(uwp string) {
	bt := []byte(uwp)
	switch string(bt[1]) {
	case "0", "1":
		an.armorScore = an.armorScore - 2
		an.addSkills(sklAthlethics)
		if an.movementType == "F" {
			an.movementType = "25% F"
		}
		an.addCharacteristic(chrStrength, -4)
		an.addCharacteristic(chrDexterity, 4)
		an.addCharacteristic(chrEndurance, -2)
	case "2", "3", "4", "5":
		an.armorScore = an.armorScore - 1
		an.addSkills(sklAthlethics)
		if an.movementType == "F" {
			an.movementType = "50% F"
		}
		an.addCharacteristic(chrStrength, -2)
		an.addCharacteristic(chrDexterity, 2)
		an.addCharacteristic(chrEndurance, -1)
	case "9", "A":
		an.armorScore = an.armorScore + 1
		an.addSkills(sklAthlethics)
		if an.movementType == "F" {
			an.movementType = "W"
		}
		an.addCharacteristic(chrStrength, 4)
		an.addCharacteristic(chrDexterity, -1)
		an.addCharacteristic(chrEndurance, 2)
	}
	switch string(bt[2]) {
	case "0", "1":
		an.armorScore = an.armorScore - 1
		an.addCharacteristic(chrDexterity, 1)
		an.addCharacteristic(chrEndurance, -1)
		an.addCharacteristic(chrInstinct, 1)
	case "2", "3":
		an.armorScore = an.armorScore - 1
		an.addCharacteristic(chrDexterity, 1)
		an.addCharacteristic(chrPack, 1)
	case "4", "5":
		an.armorScore = an.armorScore - 2
		an.addCharacteristic(chrDexterity, 2)
		an.addCharacteristic(chrEndurance, -2)
		an.addCharacteristic(chrPack, 1)
	case "8":
		an.armorScore = an.armorScore + 2
		if an.movementType == "F" {
			an.movementType = "W"
		}
		an.addCharacteristic(chrStrength, 2)
		an.addCharacteristic(chrDexterity, -1)
		an.addCharacteristic(chrEndurance, 3)
		an.addCharacteristic(chrInstinct, -1)
		an.addCharacteristic(chrPack, -1)
	case "9":
		an.armorScore = an.armorScore + 1
		an.addSkills(sklSurvival)
		if an.movementType == "F" {
			an.movementType = "W"
		}
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrInstinct, -2)
		an.addCharacteristic(chrPack, -2)
	case "B":
		an.addSkills(sklSurvival)
		an.notes += "\nImmune to acids"
		an.addCharacteristic(chrStrength, 1)
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrInstinct, -1)
		an.addCharacteristic(chrPack, -1)
	case "C":
		an.addSkills(sklSurvival)
		an.notes += "\nImmune to acids and radiation"
		an.addCharacteristic(chrEndurance, 3)
		an.addCharacteristic(chrInstinct, -1)
		an.addCharacteristic(chrPack, -1)
	}
	switch string(bt[3]) {
	case "0", "1", "2":
		an.addSkills(sklSurvival)
		an.addSkills(sklStealth)
		an.armorScore++
		an.addCharacteristic(chrStrength, 1)
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrInstinct, 2)
		an.addCharacteristic(chrPack, -2)
	case "3", "4":
		an.armorScore++
		if an.movementType == "S" {
			an.movementType = "50% S"
		}
		an.addCharacteristic(chrStrength, 1)
		an.addCharacteristic(chrDexterity, 1)
		an.addCharacteristic(chrPack, 1)
	case "8", "9", "A":
		an.addSkills(sklAthlethics)
		if an.movementType != "S" {
			an.movementType += "/50% S"
		}
		an.addCharacteristic(chrDexterity, 2)
		an.addCharacteristic(chrEndurance, 1)
		an.addCharacteristic(chrIntelligence, 1)
		an.addCharacteristic(chrPack, 1)
	}
	if an.armorScore < 0 {
		an.armorScore = 0
	}
	for key := range an.characteristics {
		if an.characteristics[key] < 0 {
			an.characteristics[key] = 0
		}
	}
}

func (an *animal) numbersInPack() string {
	if an.numbersEncountered != "" {
		return an.numbersEncountered
	}

	switch an.characteristics[chrPack] {
	case 0:
		an.numbersEncountered = "1"
	case 1, 2:
		an.numbersEncountered = "1d3"
	case 3, 4, 5:
		an.numbersEncountered = "1d6"
	case 6, 7, 8:
		an.numbersEncountered = "2d6"
	case 9, 10, 11:
		an.numbersEncountered = "3d6"
	case 12, 13, 14, 15:
		an.numbersEncountered = "4d6"
	default:
		an.numbersEncountered = "5d6"
	}
	if strings.Contains(an.notes, "Their numbers tripled") {
		an.numbersEncountered += " x 3"
	}
	return an.numbersEncountered
}

func (an *animal) defineChrs() {
	if an.size < 1 {
		an.size = 1
	}
	if an.size > 15 {
		an.size = 15
	}
	switch an.size {
	case 1:
		an.weightFlux(1)
		an.addCharacteristic(chrStrength, 1)
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrEndurance, 1)
	case 2:
		an.weightFlux(3)
		an.addCharacteristic(chrStrength, 2)
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrEndurance, 2)
	case 3:
		an.weightFlux(6)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 4:
		an.weightFlux(12)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.weightFlux(25)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("3d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("2d6").Sum())
	case 6:
		an.weightFlux(50)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("4d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("2d6").Sum())
	case 7:
		an.weightFlux(100)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("3d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("3d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("3d6").Sum())
	case 8:
		an.weightFlux(200)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("3d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("3d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("3d6").Sum())
	case 9:
		an.weightFlux(400)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("4d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("4d6").Sum())
	case 10:
		an.weightFlux(800)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("4d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("4d6").Sum())
	case 11:
		an.weightFlux(1600)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("5d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("2d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("5d6").Sum())
	case 12:
		an.weightFlux(3200)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("6d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("6d6").Sum())
	case 13:
		an.weightFlux(5000)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("7d6").Sum())
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("7d6").Sum())
	case 14:
		an.weightFlux(8000)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("8d6").Sum())
		an.addCharacteristic(chrDexterity, 2)
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("8d6").Sum())
	case 15:
		an.weightFlux(10000)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("9d6").Sum())
		an.addCharacteristic(chrDexterity, 1)
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("9d6").Sum())
	}
	an.weight = roundFloat(an.weight)
	an.addCharacteristic(chrPack, an.dicepool.RollNext("2d6").Sum())
	an.addCharacteristic(chrInstinct, an.dicepool.RollNext("2d6").Sum())
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1, 2, 3, 4:
		an.addCharacteristic(chrIntelligence, 0)
	case 5, 6:
		an.addCharacteristic(chrIntelligence, 1)
	}
	an.damageDice += (an.characteristics[chrStrength] / 10) + 1
	an.defineArmor()
	an.defineWeapon()
	//--------------------------

}

func (an *animal) defineWeapon() {
	dm := 0
	switch an.diet {
	case dietCarnivore:
		dm = 8
	case dietHerbivore:
		dm = -6
	case dietOmnivore:
		dm = 4
	}
	r := an.dicepool.RollNext("2d6").DM(dm).Sum()
	if r < 1 {
		r = 1
	}
	if r > 20 {
		r = 20
	}
	switch r {
	case 1:
		an.weapon += "None"
	case 2:
		an.weapon += "Teeth"
	case 3:
		an.weapon += "Horns"
	case 4:
		an.weapon += "Hooves/Thrasher"
	case 5:
		an.weapon += "Hooves/Thrasher and Teeth"
	case 6:
		an.weapon += "Teeth"
	case 7:
		an.weapon += "Claws"
		an.damageMod = an.damageMod + 1
	case 8:
		an.weapon += "Stinger"
		an.damageMod = an.damageMod + 1
	case 9:
		an.weapon += "Thrasher"
		an.damageMod = an.damageMod + 1
	case 10:
		an.weapon += "Claws and Teeth"
		an.damageMod = an.damageMod + 2
	case 11:
		an.weapon += "Claws"
		an.damageMod = an.damageMod + 2
	case 12:
		an.weapon += "Teeth"
		an.damageMod = an.damageMod + 2
	case 13:
		an.weapon += "Thrasher"
		an.damageMod = an.damageMod + 2
	case 14:
		an.weapon += "Claws and Teeth"
		an.damageMod = an.damageMod + 2
	case 15:
		an.weapon += "Claws"
		an.damageMod = an.damageMod + 2
	case 16:
		an.weapon += "Stinger"
		an.damageMod = an.damageMod + 2
	case 17:
		an.weapon += "Thrasher"
		an.damageMod = an.damageMod + 2
	case 18:
		an.weapon += "Teeth"
		an.damageMod = an.damageMod + 3
	case 19:
		an.weapon += "Claws and Teeth"
		an.damageMod = an.damageMod + 3
	case 20:
		an.weapon += "Thrasher"
		an.damageMod = an.damageMod + 3
	}
}

func (an *animal) addExoticWeapon() {
	r := an.dicepool.RollNext("1d6").Sum()
	switch r {
	case 1:
		an.exoticWeapon = append(an.exoticWeapon, "Diseased Attack")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, "This animal has a bite or other physical attack that carries a dangerous disease. Any successful attack that damages the target’s Endurance may cause infection.")
	case 2:
		an.exoticWeapon = append(an.exoticWeapon, "Poisoned Attack")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, "This animal has the ability to inject its target with venom, poisoning it in the same manner as with the Diseased Attack.")
	case 3:
		an.exoticWeapon = append(an.exoticWeapon, "Bleeding Wound")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, "If a target suffered Endurance damage from one of this animal’s attacks, it continues to lose 1 Endurance point every round until it is given medical attention. This effect does not stack with multiple wounds.")
	case 4:
		an.exoticWeapon = append(an.exoticWeapon, "Bioelectricity")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, "This animal carries an electrical charge and can use it to shock and stun its foes. This can be done once an encounter, is usually the first attack used by an animal and doubles the dice it rolls for damage with that attack. This Exotic Weapon should be used in conjunction with the Knockout Blow rule.")
	case 5:
		an.exoticWeapon = append(an.exoticWeapon, "Concealing Mist")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, "The animal can emit a cloud of ink or vapour that conceals its actions. This is usable once in an encounter and grants the animal a +4 DM to one Stealth or Deception check taken at the same time.")
	case 6:
		an.exoticWeapon = append(an.exoticWeapon, "Ranged Strike")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, " The animal has developed the ability to use a projectile attack of some kind. This uses the animal’s Melee (natural weapons) skill, has a range of Ranged (thrown) and does its normal damage minus 1d6 to a minimum of 1 point.")
	}
}

func (an *animal) defineArmor() {
	dm := 0
	switch an.diet {
	case dietCarnivore:
		dm = 8
	case dietHerbivore:
		dm = -6
	case dietOmnivore:
		dm = 4
	}
	r := an.dicepool.RollNext("2d6").DM(dm).Sum()
	if r < 2 {
		r = 2
	}
	armrRolled := (r - 2) / 2
	an.armorScore = an.armorScore + armrRolled
}

func (an *animal) weightFlux(base float64) {
	flD1 := an.dicepool.RollNext("1d6").Sum()
	flD2 := an.dicepool.RollNext("1d6").Sum()
	fl := float64(flD1-flD2) * 0.1
	an.weight = base + (base * fl)
}

func (an *animal) selectTerrain(t int) {
	if t < -1 || t > 15 {
		fmt.Println("Select terrain:")
		fmt.Print("Roll Random   = [", -1, "]\n")
		fmt.Print("Clear         = [", terrainClear, "]\n")
		fmt.Print("Plain         = [", terrainPlain, "]\n")
		fmt.Print("Desert        = [", terrainDesert, "]\n")
		fmt.Print("Hills         = [", terrainHills, "]\n")
		fmt.Print("Mountain      = [", terrainMountain, "]\n")
		fmt.Print("Forest        = [", terrainForest, "]\n")
		fmt.Print("Woods         = [", terrainWoods, "]\n")
		fmt.Print("Jungle        = [", terrainJungle, "]\n")
		fmt.Print("Rainforest    = [", terrainRainforest, "]\n")
		fmt.Print("Rough/Broken  = [", terrainRoughBroken, "]\n")
		fmt.Print("Swamp/Marsh   = [", terrainSwampMarsh, "]\n")
		fmt.Print("Beach/Shore   = [", terrainBeachShore, "]\n")
		fmt.Print("Riverbank     = [", terrainRiverbank, "]\n")
		fmt.Print("Shallow Ocean = [", terrainShallowOcean, "]\n")
		fmt.Print("Open Ocean    = [", terrainOpenOcean, "]\n")
		fmt.Print("Deep Ocean    = [", terrainDeepOcean, "]\n")
	} else {
		if t == -1 {
			t = an.dicepool.RollNext("1d16").DM(-1).Sum()
		}
		an.prefferedTerrain = t
		return
	}
	err := errors.New("No Input")
	for err != nil {
		t, err = user.InputInt()
		if t == -1 {
			t = dice.Roll("1d16").DM(-1).Sum()
		}
		if t > 15 || t < 0 {
			err = errors.New("Invalid Value")
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Print("\r")
	an.prefferedTerrain = t
}

func (an *animal) selectExpectedSize() {
	fmt.Println("Select expected size [1-15] (0 - RANDOM):")
	t := 0
	err := errors.New("No Input")
	for err != nil {
		t, err = user.InputInt()
		if t == 0 {
			t = dice.Roll("2d6").Sum()
			an.expectedSize = -1
			fmt.Print("\r")
			return
		}
		if t > 15 || t < 0 {
			err = errors.New("Invalid Value")
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Print("\r")
	an.expectedSize = t
}

func terrainString(t int) string {
	switch t {
	case terrainClear:
		return "Clear"
	case terrainPlain:
		return "Plain"
	case terrainDesert:
		return "Desert"
	case terrainHills:
		return "Hills"
	case terrainMountain:
		return "Mountain"
	case terrainForest:
		return "Forest"
	case terrainWoods:
		return "Woods"
	case terrainJungle:
		return "Jungle"
	case terrainRainforest:
		return "Rainforest"
	case terrainRoughBroken:
		return "Rough/Broken"
	case terrainSwampMarsh:
		return "Swamp/Marsh"
	case terrainBeachShore:
		return "Beach/Shore"
	case terrainRiverbank:
		return "Riverbank"
	case terrainShallowOcean:
		return "Shallow Ocean"
	case terrainOpenOcean:
		return "Open Ocean"
	case terrainDeepOcean:
		return "Deep Ocean"
	default:
		return "Unknown Terrain"
	}
}

func classString(class int) string {
	switch class {
	default:
		return "<UNKNOWN>"
	case 0:
		return "Amphibian"
	case 1:
		return "Aquatic"
	case 2:
		return "Avian"
	case 3:
		return "Fungal"
	case 4:
		return "Insect"
	case 5:
		return "Mammal"
	case 6:
		return "Reptile"
	}
}

func (an *animal) selectClass(t int) {
	if t < -1 || t > 6 {
		fmt.Println("Select Class:")
		fmt.Print("Roll Random  = [", -1, "]\n")
		fmt.Print("Amphibians   = [", 0, "]\n")
		fmt.Print("Aquatic      = [", 1, "]\n")
		fmt.Print("Avians       = [", 2, "]\n")
		fmt.Print("Fungals      = [", 3, "]\n")
		fmt.Print("Insect       = [", 4, "]\n")
		fmt.Print("Mammals      = [", 5, "]\n")
		fmt.Print("Reptiles     = [", 6, "]\n")
	} else {
		if t == -1 {
			t = an.dicepool.RollNext("1d7").DM(-1).Sum()
		}
		an.classification = t
		return
	}
	err := errors.New("No Input")
	for err != nil {
		t, err = user.InputInt()

		if t == -1 {
			t = dice.Roll("1d7").DM(-1).Sum()
		}
		if t > 6 || t < 0 {
			err = errors.New("Invalid Value")
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Print("\r")
	an.classification = t

}

func (an *animal) setDiet() {
	dietRoll := an.dicepool.RollNext("1d6").DM(-1).Sum()
	//an.addGenerationRolls(rollBehavior)
	switch an.classification {
	case 0: //Amphibian
		dietSl := []string{dietCarnivore, dietCarnivore, dietHerbivore, dietOmnivore, dietOmnivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklAthlethics, sklRecon, sklSurvival)
		switch an.prefferedTerrain {
		case 7, 10, 11, 12, 13:
		default:
			an.evolutionRollDM = an.evolutionRollDM - 2
		}
		switch an.diet {
		case dietCarnivore:
			an.addCharacteristic(chrStrength, 1)
			an.evolutionRollDM++
			an.addGenerationRolls(rollPhysical)
			an.addSkills(sklMeleeNW)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, 1)
			an.addCharacteristic(chrInstinct, 1)
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.addCharacteristic(chrPack, 4)
			an.addCharacteristic(chrInstinct, 1)
			an.addGenerationRolls(rollPhysical, rollSocial)
		}
	case 1: //Aquatic
		dietSl := []string{dietCarnivore, dietCarnivore, dietCarnivore, dietHerbivore, dietHerbivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklAthlethics, sklRecon, sklSurvival)
		switch an.prefferedTerrain {
		case 10, 11, 12, 13:
		default:
			an.notes = "Imposible Animal"
		}
		an.movementType = "S"
		switch an.diet {
		case dietCarnivore:
			an.addCharacteristic(chrDexterity, 1)
			an.addCharacteristic(chrPack, 4)
			an.addGenerationRolls(rollPhysical)
			an.addSkills(sklMeleeNW)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.addCharacteristic(chrPack, 2)
			an.addCharacteristic(chrInstinct, 1)
			an.addGenerationRolls(rollPhysical, rollSocial)
			an.addSkills(sklMeleeNW)
		}
	case 2: //Avian
		dietSl := []string{dietCarnivore, dietCarnivore, dietHerbivore, dietHerbivore, dietOmnivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklAthlethics, sklRecon, sklRecon, sklSurvival)
		switch an.prefferedTerrain {
		case 13, 14, 15:
			if an.movementType == "W" {
				an.movementType = "S"
			}
		default:
		}
		switch an.diet {
		case dietCarnivore:
			an.addCharacteristic(chrDexterity, 2)
			an.evolutionRollDM++
			an.addGenerationRolls(rollPhysical)
			an.addSkills(sklMeleeNW)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, 2)
			an.addCharacteristic(chrPack, 2)
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.addCharacteristic(chrPack, 2)
			an.evolutionRollDM++
			an.addGenerationRolls(rollPhysical, rollSocial)
			an.addSkills(sklMeleeNW)
		}
	case 3: //Fungals
		dietSl := []string{dietCarnivore, dietHerbivore, dietOmnivore, dietOmnivore, dietOmnivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklAthlethics, sklRecon, sklStealth, sklSurvival)
		switch an.prefferedTerrain {
		case 5, 6, 7, 8, 10:
		default:
			an.generationRollDm--
		}
		switch an.diet {
		case dietCarnivore:
			an.addCharacteristic(chrStrength, 2)
			an.evolutionRollDM++
			an.addGenerationRolls(rollPhysical)
			an.addSkills(sklMeleeNW)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, 2)
			an.addCharacteristic(chrPack, 2)
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.evolutionRollDM++
			an.addGenerationRolls(rollPhysical)
			an.addSkills(sklMeleeNW)
		}
	case 4: //Insect
		dietSl := []string{dietCarnivore, dietCarnivore, dietCarnivore, dietHerbivore, dietOmnivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklAthlethics, sklMeleeNW, sklRecon, sklSurvival)
		switch an.prefferedTerrain {
		case 13, 14, 15:
			an.notes = "Do not have their numbers tripled"
		default:
			an.notes = "Their numbers tripled"
		}
		switch an.diet {
		case dietCarnivore:
			an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
			an.addGenerationRolls(rollPhysical)
			an.addSkills(sklMeleeNW)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, 2)
			an.addCharacteristic(chrPack, 2)
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.evolutionRollDM++
			an.addCharacteristic(chrPack, 1)
			an.addGenerationRolls(rollPhysical, rollSocial)
			an.addSkills(sklMeleeNW)
		}
	case 5: //Mammal
		dietSl := []string{dietCarnivore, dietCarnivore, dietHerbivore, dietHerbivore, dietOmnivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklAthlethics, sklMeleeNW, sklRecon, sklSurvival)
		switch an.prefferedTerrain {
		case 0, 1, 3, 5, 6, 12:
		default:
			an.addCharacteristic(chrPack, 1)
		}
		switch an.diet {
		case dietCarnivore:
			an.addCharacteristic(chrStrength, 1)
			an.addCharacteristic(chrDexterity, 1)
			an.addGenerationRolls(rollPhysical)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, 2)
			an.addCharacteristic(chrPack, 2)
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.evolutionRollDM++
			an.addCharacteristic(chrIntelligence, 1)
			an.addGenerationRolls(rollPhysical, rollSocial)
		}
	case 6: //Reptiles
		dietSl := []string{dietCarnivore, dietCarnivore, dietHerbivore, dietHerbivore, dietOmnivore, dietOmnivore}
		an.diet = dietSl[dietRoll]
		an.addSkills(sklMeleeNW, sklRecon, sklSurvival)
		switch an.prefferedTerrain {
		case 1, 2, 3, 7, 8, 9, 10, 12:
		default:
			an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
			an.characteristics[chrDexterity] = an.characteristics[chrDexterity] - 2
		}
		switch an.diet {
		case dietCarnivore:
			an.addSkills(sklAthlethics)
			an.addCharacteristic(chrStrength, 1)
			an.addCharacteristic(chrDexterity, 1)
			an.addGenerationRolls(rollPhysical)
		case dietHerbivore:
			an.addCharacteristic(chrEndurance, 2)
			an.addCharacteristic(chrPack, 1)
			an.addGenerationRolls(rollSocial)
		case dietOmnivore:
			an.addSkills(sklAthlethics)
			an.evolutionRollDM++
			an.addCharacteristic(chrEndurance, 1)
			an.addGenerationRolls(rollPhysical, rollSocial)
		}
	}
}

func (an *animal) addGenerationRolls(rolls ...string) {
	for i := range rolls {
		an.expectedRolls = append(an.expectedRolls, rolls[i])
	}
}

func (an *animal) setMovement() {
	r := an.dicepool.RollNext("1d6").DM(-1).Sum()
	moveMap := make(map[int][]string)
	moveMap[0] = []string{"W", "W", "W", "W", "W +2", "F -6"}
	moveMap[1] = []string{"W", "W", "W", "W +2", "W +4", "F -6"}
	moveMap[2] = []string{"W", "W", "W", "W", "F -4", "F -6"}
	moveMap[3] = []string{"W", "W", "W", "W +2", "F -4", "F -6"}
	moveMap[4] = []string{"W", "W", "W", "F -2", "F -4", "F -6"}
	moveMap[5] = []string{"W", "W", "W", "W", "F -4", "F -6"}
	moveMap[6] = []string{"W", "W", "W", "W", "W", "F -6"}
	moveMap[7] = []string{"W", "W", "W", "W", "W +2", "F -6"}
	moveMap[8] = []string{"W", "W", "W", "W +2", "W +4", "F -6"}
	moveMap[9] = []string{"W", "W", "W", "W +2", "F -4", "F -6"}
	moveMap[10] = []string{"S -6", "S", "W", "W", "F -4", "F -6"}
	moveMap[11] = []string{"S +1", "S +1", "W", "W", "F -4", "F -6"}
	moveMap[12] = []string{"S -4", "S +2", "W", "W", "W", "F -6"}
	moveMap[13] = []string{"S +4", "S +2", "S", "S", "F -4", "F -6"}
	moveMap[14] = []string{"S +6", "S +4", "S +2", "S", "F -4", "F -6"}
	moveMap[15] = []string{"S +8", "S +6", "S +4", "S +2", "S", "S -2"}
	mDataRaw := moveMap[an.prefferedTerrain][r]
	mData := strings.Split(mDataRaw, " ")
	an.movementType = mData[0]
	if an.classification == 1 {
		an.movementType = "S"
	}
	if len(mData) > 1 {
		sizeChange, _ := strconv.Atoi(mData[1])
		an.size = an.size + sizeChange
	}
	switch an.prefferedTerrain {
	case 13, 14, 15:
	default:
		if an.dicepool.RollNext("2d6").Sum() > 9 && an.movementType == "F" {
			an.movementType = "B"
			an.addSkills(sklStealth)
			an.addCharacteristic(chrInstinct, 2)
		}

	}
}

func (an *animal) addSkills(skills ...string) {
	for _, skill := range skills {
		if _, ok := an.skills[skill]; ok {
			an.skills[skill]++
			return
		}
		an.skills[skill] = 0
	}
}

func (an *animal) addCharacteristic(chr string, val int) {
	if _, ok := an.characteristics[chr]; ok {
		an.characteristics[chr] = an.characteristics[chr] + val
		return
	}
	an.characteristics[chr] = val
}

func (an *animal) evolutionBase() {
	r := an.dicepool.RollNext("1d6").DM(an.evolutionRollDM).Sum()
	switch an.classification {
	case 0:
		switch r {
		case 1, 2:
			an.addGenerationRolls(rollSocial)
		case 3, 4:
			an.addGenerationRolls(rollSocial, rollEvolution)
		case 5:
			an.addGenerationRolls(rollPhysical, rollEvolution)
		case 6:
			an.addGenerationRolls(rollSocial, rollPhysical, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollSocial, rollPhysical, rollEvolution)
		}
	case 1:
		switch r {
		case 1:
			an.addGenerationRolls(rollSocial)
		case 2:
			an.addGenerationRolls(rollPhysical)
		case 3:
			an.addGenerationRolls(rollSocial)
		case 4:
			an.addGenerationRolls(rollSocial, rollEvolution)
		case 5:
			an.addGenerationRolls(rollPhysical, rollEvolution)
		case 6:
			an.addGenerationRolls(rollSocial, rollPhysical, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollSocial, rollPhysical, rollEvolution)
		}
	case 2:
		switch r {
		case 1, 2:
			an.addGenerationRolls(rollSocial)
		case 3:
			an.addGenerationRolls(rollSocial, rollPhysical)
		case 4:
			an.addGenerationRolls(rollSocial, rollEvolution)
		case 5:
			an.addGenerationRolls(rollPhysical, rollEvolution)
		case 6:
			an.addGenerationRolls(rollSocial, rollPhysical, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollSocial, rollPhysical, rollEvolution)
		}
	case 3:
		switch r {
		case 1, 2:
			an.addGenerationRolls(rollSocial)
		case 3:
			an.addGenerationRolls(rollSocial, rollPhysical)
		case 4:
			an.addGenerationRolls(rollSocial, rollEvolution)
		case 5:
			an.addGenerationRolls(rollPhysical, rollEvolution)
		case 6:
			an.addGenerationRolls(rollPhysical, rollPhysical, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollPhysical, rollEvolution)
		}
	case 4:
		switch r {
		case 1, 2:
			an.addGenerationRolls(rollPhysical)
		case 3:
			an.addGenerationRolls(rollSocial, rollPhysical)
		case 4:
			an.addGenerationRolls(rollPhysical, rollSocial)
		case 5:
			an.addGenerationRolls(rollPhysical, rollEvolution)
		case 6:
			an.addGenerationRolls(rollPhysical, rollPhysical, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollPhysical, rollEvolution)
		}
	case 5:
		switch r {
		case 1:
			an.addGenerationRolls(rollPhysical)
		case 2:
			an.addGenerationRolls(rollSocial)
		case 3:
			an.addGenerationRolls(rollSocial, rollPhysical)
		case 4:
			an.addGenerationRolls(rollPhysical, rollSocial)
		case 5:
			an.addGenerationRolls(rollSocial, rollEvolution)
		case 6:
			an.addGenerationRolls(rollPhysical, rollSocial, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollSocial, rollPhysical, rollEvolution)
		}
	case 6:
		switch r {
		case 1:
			an.addGenerationRolls(rollPhysical)
		case 2:
			an.addGenerationRolls(rollPhysical)
		case 3:
			an.addGenerationRolls(rollSocial, rollPhysical)
		case 4:
			an.addGenerationRolls(rollPhysical, rollPhysical)
		case 5:
			an.addGenerationRolls(rollSocial, rollEvolution)
		case 6:
			an.addGenerationRolls(rollPhysical, rollPhysical, rollEvolution)
		case 7:
			an.addGenerationRolls(rollSocial, rollPhysical, rollEvolution)
		}
	}
}

func (an *animal) rollBenefits() []int {
	roll := 1
	rArr := []int{}
	for roll > 0 {
		r := an.dicepool.RollNext("1d6").DM(an.evolutionRollDM).Sum()
		if r < 1 {
			r = 1
		}
		switch r {
		case 7:
			roll++
			continue
		default:
			roll--
			rArr = append(rArr, r)
		}
	}
	return rArr
}

func (an *animal) applyBenefits(benefits ...int) {
	for _, b := range benefits {
		switch an.classification {
		case 0:
			an.amphibianBenefit(b)
		case 1:
			an.aquaticBenefit(b)
		case 2:
			an.avianBenefit(b)
		case 3:
			an.fungalBenefit(b)
		case 4:
			an.insectBenefit(b)
		case 5:
			an.mammalBenefit(b)
		case 6:
			an.reptileBenefit(b)
		}
	}
}

func (an *animal) amphibianBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrInstinct, 2)
	case 2:
		an.addCharacteristic(chrPack, 2)
	case 3:
		an.addCharacteristic(chrIntelligence, 1)
	case 4:
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) aquaticBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrInstinct, 2)
	case 2:
		an.addCharacteristic(chrPack, 2)
	case 3:
		an.addCharacteristic(chrIntelligence, 1)
	case 4:
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) avianBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrInstinct, 2)
	case 2:
		an.addCharacteristic(chrPack, 2)
	case 3:
		an.addCharacteristic(chrInstinct, 1)
		an.addCharacteristic(chrPack, 2)
	case 4:
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrPack, 2)
	case 5:
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrPack, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) fungalBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrInstinct, 2)
	case 2:
		an.addCharacteristic(chrEndurance, 2)
	case 3:
		an.addExoticWeapon()
	case 4:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrPack, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) insectBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrInstinct, 2)
	case 2:
		an.addGenerationRolls(rollEvolution)
	case 3:
		an.addExoticWeapon()
	case 4:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) mammalBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrInstinct, 2)
	case 2:
		an.addGenerationRolls(rollEvolution)
	case 3:
		an.addCharacteristic(chrStrength, 1)
		an.addCharacteristic(chrEndurance, 1)
		an.addCharacteristic(chrPack, 1)
	case 4:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.size++
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) reptileBenefit(b int) {
	switch b {
	case 1:
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrDexterity, 1)
	case 2:
		an.addCharacteristic(chrStrength, 2)
	case 3:
		an.addCharacteristic(chrDexterity, 1)
		an.addCharacteristic(chrPack, 1)
	case 4:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 5:
		an.size++
	case 6:
		an.rollQuirks()
	}
}

func (an *animal) rollQuirks() {
	roll := 2
	rArr := []int{}
	for roll > 0 {
		r := an.dicepool.RollNext("2d6").Sum()
		switch r {
		case 12:
			roll++
			continue
		default:
			roll--
			//rArr = append(rArr, r)
			rArr = utils.AppendUniqueInt(rArr, r)
		}
	}
	for _, q := range rArr {
		switch an.classification {
		case 0:
			an.amphibianQuirk(q)
		case 1:
			an.aquaticQuirk(q)
		case 2:
			an.avianQuirk(q)
		case 3:
			an.fungalQuirk(q)
		case 4:
			an.insectQuirk(q)
		case 5:
			an.mammalQuirk(q)
		case 6:
			an.reptileQuirk(q)
		}
	}
}

func (an *animal) amphibianQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nWhenever packs of these animals make any noise at all, they all make the exact same sound simultaneously several times in a row."
	case 3:
		an.notes += "\nApparently blind, these amphibians have no visible eyes or means of sight."
		an.skills[sklRecon] = an.skills[sklRecon] - 1
	case 4:
		an.notes += "\nThese animals make no sound at all, even when they move in natural surroundings."
		an.addSkills(sklStealth)
	case 5:
		an.notes += "\nThe colours of this amphibian’s hide are vivid and clashing, a sort of natural reverse camouflage. Natural predators dislike this display and leave it alone."
	case 6:
		an.notes += "\nSeemingly everywhere, forms of this animal can be found in virtually every habitat type on their world."
		an.addSkills(sklSurvival)
		an.addSkills(sklSurvival)
	case 7:
		an.notes += "\nThese amphibians emit a natural pheromone that other animals find highly attractive."
		an.behaviour += "-Siren"
	case 8:
		an.notes += "\nWhen threatened, these amphibians emit a piercing scream that sounds like a sentient creature in terrible pain."
	case 9:
		an.notes += "\nOn rare occasions, these amphibians swarm viciously."
		if an.dicepool.RollNext("2d6").Sum() == 2 {
			an.behaviour = "Killer"
		}
	case 10:
		an.notes += "\nThe skin of these animals is naturally coated in a thick, foul-smelling emulsion."
		an.exoticWeapon = append(an.exoticWeapon, "Stench Exotic Weapon")
		an.exoticWeaponEffect = append(an.exoticWeaponEffect, "The skin of these animals is naturally coated in a thick, foul-smelling emulsion.")
	case 11:
		an.notes += "\nUnusually for its kind, these amphibians have developed a rigid shell over their forelimbs and torsos."
		an.armorScore = an.armorScore + 2
	}
}

func (an *animal) aquaticQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nThis aquatic is found in the darkest parts of its habitat and sees through bioluminescent eyes."
	case 3:
		an.notes += "\nThis creature is never found alone and will die within " + an.dicepool.RollNext("1d6").SumStr() + " days of natural causes if it cannot find a pack to join."
		if an.characteristics[chrPack] == 0 {
			an.characteristics[chrPack] = 1
		}
	case 4:
		an.notes += "\nPosseses a frail physique and has the ability to engage in extremely swift movement."
		an.behaviour += "-Pouncer"
		an.armorScore = 0
	case 5:
		an.notes += "\nPossessed of a unique biology, this aquatic can survive for " + an.dicepool.RollNext("1d6").SumStr() + " hours on dry land and has a W movement type in addition to its ability to swim."
		an.movementType += " and W"
	case 6:
		an.notes += "\nUnnaturally large for the local ecology."
		an.size++
	case 7:
		an.notes += "\nCapable of surviving for long periods of time without any nourishment, this aquatic goes dormant for long periods of time, awaking for " + an.dicepool.RollNext("3d6").SumStr() + " days at a time to feed and breed."
	case 8:
		an.notes += "\nThis aquatic breed has volatile genetics and is prone to mutation."
		an.evolutionRollDM++
	case 9:
		an.notes += "\nUnlike most aquatics, this species reproduces asexually and is never encountered with others of its kind."
		an.characteristics[chrPack] = 0
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 10:
		an.notes += "\nExtremely vicious, this animal gains a +1 DM to all Melee (natural weapons) and damage rolls after it or its opponent suffers damage in combat."
	case 11:
		an.notes += "\nUnusually bright and clever."
		an.notes += "\nHas a minimum Instinct of 9."
		an.characteristics[chrIntelligence] = 2
	}
}

func (an *animal) avianQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nThe plumage of this animal is highly exotic and valuable, exhibiting colours rarely found within its habitat."
	case 3:
		an.notes += "\nExtremely social, these animals live in immense flocks.\nIf the animal’s Pack score is 6 or less increase it to 12. For animals with a Pack score of 7 or more double the number."
	case 4:
		an.notes += "\nQuite at home on the ground, this species has evolved away from flight and no longer has an F movement type."
		an.movementType = "W"
	case 5:
		an.notes += "\nFar smaller than their evolutionary niche would suggest."
		an.size = an.size - 4
	case 6:
		an.notes += "\nThese avians have adapted a very unusual way of dealing with enemies."
		an.addExoticWeapon()
	case 7:
		an.notes += "\nThese avians have developed a way to emit calls that sound exactly like the cries of wounded prey. Using these to lure meals closer."
		an.behaviour = "Siren"
	case 8:
		an.notes += "\nEnvironmental pressures have forced this animal to adapt to a hostile environment, granting +1 to both Endurance and Armour."
		an.armorScore++
		an.addCharacteristic(chrEndurance, 1)
	case 9:
		an.notes += "\nNot just ground bound, this flightless species has no F movement rate and thrives because of it."
		an.movementType = "W"
		an.behaviour = "Chaser"
		an.addCharacteristic(chrEndurance, 1)
	case 10:
		an.notes += "\nPossessed of a deadly main attack, these avians are truly vicious and always press their attack once they wound an enemy."
	case 11:
		an.notes += "\nThese avians mate for life, are never encountered in packs larger than a pair of adults. If one is killed the other will automatically flee if possible."
	}
}

func (an *animal) fungalQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nThis Fungal is an absolutely bizarre colour and smells rancid. It loses all ranks in Stealth, cannot succeed at Stealth rolls and gains an Exotic Weapon."
		delete(an.skills, sklStealth)
	case 3:
		an.notes += "\nUnlike other fungus-based life, this species has developed a rudimentary vocal structure. The sounds it can make may be extremely strange, similar to nothing else found in nature."
	case 4:
		an.notes += "\nThe Fungal can inflate itself with a light gas, allowing for a slow form of flight."
	case 5:
		an.notes += "\nThough capable of physical movement to attack or defend itself, this Fungal species is stationary and cannot change location."
		//"Its behaviour changes to Siren and it gains +1d6 End. If the base species was herbivorous, it is now specialises in luring other fungals to their doom."
		an.behaviour = "Siren"
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.notes += "\nThis species propagates very quickly and easily, dwelling in large family structures with its progeny."
		an.addCharacteristic(chrPack, an.dicepool.RollNext("1d6").Sum())
	case 7:
		an.notes += "\nVery soft in bodily structure, this fungal loses all Armour but gains +1d6 End."
		an.armorScore = 0
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
	case 8:
		an.notes += "\nThe scent and outlandish appearance of this fungal terrifies other animals."
		an.behaviour += "-Hijacker"
	case 9:
		an.notes += "\nUnfortunately for this fungal, its biological structure is extremely nutritious, capable of feeding even carnivores in its environment.\nWhen encountered, there is a 50% chance that a predator of another species is also in the area."
	case 10:
		an.notes += "\nCapable of rapid regrowth from even very small samples, this species must be completely destroyed or it will regenerate completely in " + an.dicepool.RollNext("1d6").SumStr() + " days."
	case 11:
		an.notes += "\nAlmost liquid in structure, this extremely slimy fungal moves at normal speed and is capable of extremely rapid motion when it hunts."
		an.behaviour += "-Pouncer"
		an.addCharacteristic(chrDexterity, 2)
	}
}

func (an *animal) insectQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nExtremely unusual in appearance, these insects have apparently useless and garish physical structures and barely fit in their own ecosystems."
	case 3:
		an.notes += "\nSlow moving because of heavy exoskeleton plating, these insects travel at half speed but gain 1 Armour in return."
		an.armorScore++
	case 4:
		an.notes += "\nIf this insect has a Flight mode of travel, it loses it and gains +1d6 Str instead. If it does not, it gains flight and loses 1 Str and 1 Armour."
		if an.movementType == "F" {
			an.movementType = "W"
			an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
		} else {
			an.movementType = "F"
			an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum()*-1)
		}
	case 5:
		an.notes += "\nThese insects form veritable swarms. They never have a Pack score less than 2 and they appear in four times the number as opposed to three."
	case 6:
		an.notes += "\nSolitary by nature, these insects have a Pack of 0 and exchange their behaviour type for Trapper. If they have the ability to fly, they become Pouncers instead. If the insects are herbivores, they just leave their prey to rot and eat the resulting fungus."
		an.characteristics[chrPack] = 0
		an.behaviour = "Trapper"
		if an.movementType == "F" {
			an.behaviour = "Pouncer"
		}
	case 7:
		an.notes += "\nAcutely self-aware, these no longer triple their numbers when encountered."
		an.addCharacteristic(chrIntelligence, 2)
	case 8:
		an.notes += "\nThese insects have a hive mind and a minimum Pack score of 6. One of their number has an Intelligence of 2, all the rest are 0 and serve its will without question."
		if an.characteristics[chrPack] < 6 {
			an.characteristics[chrPack] = 6
		}
	case 9:
		an.notes += "\nEvolved in a particularly dangerous habitat, these insects developed an unusual defence. They gain an Exotic Weapon."
		an.addExoticWeapon()
	case 10:
		an.notes += "\nThese insects have a decentralised nervous system and can be hacked apart into smaller creatures. In combat, any attack that inflicts Endurance damage has a 50% chance of splitting the insect in half. The resulting insects have their attack damage dice halved and divide their remaining Endurance between them. If this would result in an insect with a starting End of 3 or less, the insect dies instead of splitting."
	case 11:
		an.notes += "\nThe insect can generate a hypnotic drone. It gains the Siren behaviour in addition to its own."
		an.behaviour += "-Siren"
	}
}

func (an *animal) mammalQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nThis mammal has an unusual mode of travel, be it gliding or swinging between trees in its home environment. If the animal is an omnivore, it gains the Pouncer behaviour instead of its normal one."
		if an.diet == dietOmnivore {
			an.behaviour = "Pouncer"
		}
	case 3:
		an.notes += "\nExtremely swift."
		an.size = an.size / 2
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 4:
		an.notes += "\nThese animals have remarkably fast metabolisms, enabling them to recover quickly from injuries. They regain one lost Endurance point every other round of combat starting at the beginning of the second round."
	case 5:
		an.notes += "\nBright even for its class, these mammals show a devious cunning that borders on compulsive mischief. They gain Stealth and Deception."
		an.addSkills(sklStealth)
		an.addSkills(sklDeception)
	case 6:
		an.notes += "\nProfuse body hair marks this species as a sign of its innate adaptability."
		an.addSkills(sklSurvival)
	case 7:
		an.notes += "\nHerd-oriented and nomadic, these mostly peaceful mammals."
		an.addCharacteristic(chrPack, an.dicepool.RollNext("1d6").Sum())
	case 8:
		an.notes += "\nThese animals have prodigious horns and know how to use them in combat. They gain horns as a weapon type if they did not have them before and a rank of Melee (natural weapons)."
		an.addSkills(sklMeleeNW)
		an.weapon = "Horns and "
	case 9:
		an.notes += "\nUnusually vicious, these mammals are hostile to any species but their own. They gain the Killer behaviour type in addition to their own. If they are already Killers, their Reaction Modifier increases to +2 and they gain 2 Str."
		if an.behaviour == "Killer" {
			an.reactionDM = an.reactionDM + 2
			an.addCharacteristic(chrStrength, 2)
		} else {
			an.behaviour += "-Killer"
		}
	case 10:
		an.notes += "\nAdapted to an aquatic environment even if they do not normally live near one."
		if an.movementType != "S" {
			an.movementType = "S"
		} else {
			an.notes += "Move at twice the normal speed."
		}
	case 11:
		an.notes += "\nThis animal species is on the verge of evolving into sentience.\nInstinct score is 12 at a minimum."
		an.addCharacteristic(chrIntelligence, 2)
	}
}

func (an *animal) reptileQuirk(q int) {
	switch q {
	default:
		panic("Quirk unknown: " + strconv.Itoa(q))
	case 2:
		an.notes += "\nOutlandish colours and adaptations make this reptile a bizarre sight and remarkably intimidating to other non-sentient species."
	case 3:
		an.notes += "\nMottled in appearance and adapted to its surroundings."
		an.addSkills(sklStealth)
	case 4:
		an.notes += "\nSeveral of the scales on this reptile are jagged and sharp, letting it inflict 4 + the Effect in damage when it grapples. This becomes its main way to hunt if the animal eats live prey."
	case 5:
		an.notes += "\nAble to go dormant for long periods of time, may go for weeks or even months between meals."
		an.addSkills(sklSurvival)
	case 6:
		an.notes += "\nThis reptile buries itself in its terrain, blending in and waiting for prey to ensnare."
		an.addSkills(sklStealth)
		an.behaviour = "Trapper"
		if an.diet == dietHerbivore {
			switch an.dicepool.RollNext("1d2").Sum() {
			case 1:
				an.diet = dietCarnivore
			case 2:
				an.diet = dietOmnivore
			}
		}
	case 7:
		an.notes += "\nThese reptiles see heat, allowing them to have normal vision even in total darkness."
	case 8:
		an.notes += "\nCapable of flying, these reptiles have adapted body structures that generate heat through wind friction, allowing them to stay warm during flight. They do not sleep, they never land intentionally and will die within " + an.dicepool.RollNext("1d6").SumStr() + " hours if grounded."
		an.movementType = "F"
	case 9:
		an.notes += "\nUnlike other reptiles, these animals have no scales and rely on a dense hide for defence."
		an.armorScore = an.armorScore / 2
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 10:
		an.notes += "\nAn oddity even within an evolutionarily diverse class, this reptile has a very complex genetic history and gains two Exotic Weapons as a result."
		an.addExoticWeapon()
		an.addExoticWeapon()
	case 11:
		an.notes += "\nRelative safety in its environment has allowed this species to evolve mentally"
		an.addCharacteristic(chrIntelligence, 1)
	}
}

func (an *animal) setAnimalBehaviour() {
	//r := an.dicepool.RollNext("1d6").Sum()
	dietMode := 0
	switch an.diet {
	case dietHerbivore:
		dietMode = 2
	case dietOmnivore:
		dietMode = 4
	}
	rawData := ""
	rll := an.dicepool.RollNext("1d6").Sum()
	switch rll + (6 * an.classification) {
	case 1:
		rawData = "Pouncer -1 Filter -1 Carrion_Еater -1 "
	case 2:
		rawData = "Trapper -2 Filter +0 Gatherer -1 "
	case 3:
		rawData = "Hunter -2 Intermittent -2 Eater -1 "
	case 4:
		rawData = "Hunter -1 Intermittent -1 Hunter +0 "
	case 5:
		rawData = "Hunter +0 Intermittent +0 Intermittent -1 "
	case 6:
		rawData = "Chaser -2 Grazer -2 Reducer -2 "
	case 7:
		rawData = "Eater -1 Filter -1 Carrion_Еater -1 "
	case 8:
		rawData = "Hunter -2 Filter +0 Eater -1 "
	case 9:
		rawData = "Killer -1 Filter +1 Eater -0 "
	case 10:
		rawData = "Killer +0 Intermittent -1 Eater +1 "
	case 11:
		rawData = "Killer +1 Intermittent +0 Reducer -1 "
	case 12:
		rawData = "Chaser -2 Grazer -2 Reducer -2 "
	case 13:
		rawData = "Hunter -1 Intimidator -1 Carrion_Еater -1 "
	case 14:
		rawData = "Hunter +0 Intermittent +0 Eater -1 "
	case 15:
		rawData = "Hunter +0 Intermittent +0 Eater -0 "
	case 16:
		rawData = "Chaser +0 Intermittent +1 Intimidator +1 "
	case 17:
		rawData = "Killer +1 Intermittent +2 Reducer -1 "
	case 18:
		rawData = "Pouncer -2 Grazer -2 Reducer -2 "
	case 19:
		rawData = "Hunter -2 Intermittent -2 Carrion_Еater -1 "
	case 20:
		rawData = "Hunter -1 Intermittent -1 Carrion_Еater +0 "
	case 21:
		rawData = "Hunter -0 Intermittent -1 Eater +0 "
	case 22:
		rawData = "Siren +0 Intermittent +0 Reducer +0 "
	case 23:
		rawData = "Siren +1 Grazer -1 Reducer -1 "
	case 24:
		rawData = "Killer +0 Grazer -2 Reducer -2 "
	case 25:
		rawData = "Pouncer +0 Eater -2 Carrion_Еater -1 "
	case 26:
		rawData = "Hunter +1 Intermittent -1 Eater +0 "
	case 27:
		rawData = "Hunter +2 Filter +0 Eater +2 "
	case 28:
		rawData = "Killer +0 Intermittent +0 Reducer +0 "
	case 29:
		rawData = "Trapper -1 Grazer -1 Reducer -1 "
	case 30:
		rawData = "Chaser +0 Gatherer +0 Reducer -2 "
	case 31:
		rawData = "Pouncer +0 Eater -2 Carrion_Еater -1 "
	case 32:
		rawData = "Killer +1 Intermittent -1 Gatherer +0 "
	case 33:
		rawData = "Trapper +0 Intermittent +0 Gatherer +1 "
	case 34:
		rawData = "Chaser +0 Intermittent +0 Hunter +0 "
	case 35:
		rawData = "Hunter -1 Grazer -1 Intimidator -1 "
	case 36:
		rawData = "Hijacker +0 Gatherer +0 Reducer +0 "
	case 37:
		rawData = "Pouncer +0 Gatherer -1 Carrion_Еater -1 "
	case 38:
		rawData = "Killer +1 Intermittent -1 Gatherer +0 "
	case 39:
		rawData = "Killer +2 Intermittent +0 Hijacker +1 "
	case 40:
		rawData = "Intimidator +0 Intermittent +1 Hunter +0 "
	case 41:
		rawData = "Hunter +1 Grazer -1 Hunter +1 "
	case 42:
		rawData = "Hijacker +0 Grazer +0 Reducer +0"
	}
	spData := strings.Split(rawData, " ")
	an.behaviour = spData[0+dietMode]
	rdm, err := strconv.Atoi(spData[1+dietMode])
	if err != nil {
		fmt.Println(rdm, err, spData[1+dietMode], 1+dietMode, rawData, rll, an.classification*6)
		panic(nil)
	}
	an.reactionDM = rdm

}

func (an *animal) attackFleeBeh() {
	behParts := strings.Split(an.behaviour, "-")
	an.behaviour = ""
	cleaned := []string{}
	for i := range behParts {
		cleaned = utils.AppendUniqueStr(cleaned, behParts[i])
	}
	sort.Strings(cleaned)
	for i := range cleaned {
		an.behaviour += cleaned[i] + "-"
	}
	an.behaviour = strings.TrimSuffix(an.behaviour, "-")
	trueBeh := an.behaviour
	an.behaviour = cleaned[0]
	switch an.behaviour {
	default:
		panic("Unknown Behavoiur: " + an.behaviour)
	case "Filter":
		an.attackIF = strconv.Itoa(10+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(5+an.reactionDM) + "-"
	case "Intermittent":
		an.attackIF = strconv.Itoa(10+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(4+an.reactionDM) + "-"
	case "Intermittent-Siren":
		an.attackIF = strconv.Itoa(10+an.reactionDM) + "+, If it has surprise, it attacks"
		an.fleeIF = strconv.Itoa(4+an.reactionDM) + "-, If it is Immobile, it cannot flee and attacks."
	case "Grazer":
		an.attackIF = strconv.Itoa(8+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(6+an.reactionDM) + "-"
	case "Gatherer":
		an.attackIF = strconv.Itoa(9+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(7+an.reactionDM) + "-"
	case "Eater", "Hijacker-Eater":
		an.attackIF = strconv.Itoa(5+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(4+an.reactionDM) + "-"
	case "Killer", "Carrion_Еater-Killer", "Hijacker-Killer", "Gatherer-Killer-Reducer":
		an.attackIF = strconv.Itoa(6+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(3+an.reactionDM) + "-"
	case "Hijacker":
		an.attackIF = strconv.Itoa(7+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(6+an.reactionDM) + "-"
	case "Intimidator", "Carrion_Еater-Intimidator", "Intimidator-Carrion_Еater":
		an.attackIF = strconv.Itoa(8+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(7+an.reactionDM) + "-"
	case "Carrion_Еater":
		an.attackIF = strconv.Itoa(11+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(7+an.reactionDM) + "-"
	case "Reducer", "Reducer-Carrion_Еater":
		an.attackIF = strconv.Itoa(10+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(7+an.reactionDM) + "-"
	case "Trapper":
		an.attackIF = "If the trapper has surprise, it attacks."
		an.fleeIF = strconv.Itoa(5+an.reactionDM) + "-"
	case "Pouncer":
		an.attackIF = " If the Pouncer has surprise, it attacks."
		an.fleeIF = " If the Pouncer is surprised, it flees. If it cannot flee, it attacks."
	case "Hunter", "Hunter-Siren":
		an.attackIF = " If the Hunter is heavier than a least one foe, it attacks on a " + strconv.Itoa(6+an.reactionDM) + "+. Otherwise it attacks on a " + strconv.Itoa(10+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(5+an.reactionDM) + "-"
	case "Hunter-Killer":
		an.attackIF = " If the Hunter-Killer is heavier than a least one foe, it attacks on a " + strconv.Itoa(6+an.reactionDM) + "+. Otherwise it attacks on a " + strconv.Itoa(10+an.reactionDM) + "+"
		an.fleeIF = strconv.Itoa(3+an.reactionDM) + "-"
	case "Chaser", "Chaser-Trapper":
		an.attackIF = " If the chasers outnumber the foes, they attack."
		an.fleeIF = strconv.Itoa(5+an.reactionDM) + "-"
	case "Siren":
		an.attackIF = " If the Siren has surprise, it attacks"
		an.fleeIF = strconv.Itoa(4+an.reactionDM) + "-, If it is Immobile, it cannot flee and attacks."
	}
	an.behaviour = trueBeh
}

func (an *animal) getBehaviourBenefits() {
	if strings.Contains(an.behaviour, "Carrion_Еater") {
		an.addCharacteristic(chrInstinct, 2)
		an.size = an.size - 2
		if len(an.exoticWeapon) > 0 {
			an.exoticWeapon[0] = "Diseased Attack"
			an.exoticWeaponEffect[0] = "This animal has a bite or other physical attack that carries a dangerous disease. Any successful attack that damages the target’s Endurance may cause infection."
		}
	} else {
		if strings.Contains(an.behaviour, "Eater") {
			an.addCharacteristic(chrEndurance, 4)
			an.addCharacteristic(chrPack, 4)
		}
	}
	if strings.Contains(an.behaviour, "Chaser") {
		an.addCharacteristic(chrDexterity, 4)
		an.addCharacteristic(chrInstinct, 2)
		an.addCharacteristic(chrPack, 2)
	}
	if strings.Contains(an.behaviour, "Filter") {
		an.addCharacteristic(chrEndurance, 4)
		an.addCharacteristic(chrPack, -2)
	}
	if strings.Contains(an.behaviour, "Gatherer") {
		an.addSkills(sklStealth)
		an.addCharacteristic(chrInstinct, 1)
		an.addCharacteristic(chrPack, 2)
	}
	if strings.Contains(an.behaviour, "Grazer") {
		an.addCharacteristic(chrInstinct, 2)
		an.addCharacteristic(chrPack, 4)
	}
	if strings.Contains(an.behaviour, "Hunter") {
		an.addSkills(sklSurvival)
		an.addSkills(sklRecon)
		an.addCharacteristic(chrInstinct, 2)
	}
	if strings.Contains(an.behaviour, "Hijacker") {
		an.addCharacteristic(chrStrength, 2)
		an.addCharacteristic(chrPack, 2)
	}
	if strings.Contains(an.behaviour, "Intimidator") {
		an.addSkills(sklPersuade)
		an.addCharacteristic(chrInstinct, 2)
	}
	if strings.Contains(an.behaviour, "Killer") {
		an.addSkills(sklMeleeNW)
		an.addCharacteristic(chrInstinct, 4)
		an.addCharacteristic(chrPack, -2)
		switch an.dicepool.RollNext("1d6").Sum() {
		case 1, 2, 3:
			an.addCharacteristic(chrStrength, 4)
		case 4, 5, 6:
			an.addCharacteristic(chrDexterity, 4)
		}
	}
	if strings.Contains(an.behaviour, "Intermittent") {
		an.addSkills(sklSurvival)
		an.addCharacteristic(chrPack, 4)
		an.size = an.size + 2
	}
	if strings.Contains(an.behaviour, "Pouncer") {
		an.addSkills(sklStealth)
		an.addSkills(sklRecon)
		an.addSkills(sklAthlethics)
		an.addCharacteristic(chrDexterity, 2)
		an.addCharacteristic(chrInstinct, 2)
	}
	if strings.Contains(an.behaviour, "Reducer") {
		an.addCharacteristic(chrEndurance, 2)
		an.addCharacteristic(chrPack, 4)
	}
	if strings.Contains(an.behaviour, "Siren") {
		an.addSkills(sklDeception)
		an.addCharacteristic(chrPack, -4)
		if an.dicepool.RollNext("1d4").Sum() == 4 {
			an.movementType = "Immobile"
			an.addSkills(sklSurvival)
			an.addCharacteristic(chrEndurance, 2)
		}
	}
	if strings.Contains(an.behaviour, "Trapper") {
		an.addSkills(sklStealth)
		an.addCharacteristic(chrEndurance, 1)
		an.addCharacteristic(chrPack, -2)
	}
}

func (an *animal) generationRolls() {
	for _, roll := range an.expectedRolls {
		switch an.classification {
		case 0:
			if roll == rollPhysical {
				an.amphibianPhysical()
			}
			if roll == rollSocial {
				an.amphibianSocial()
			}
			if roll == rollEvolution {
				an.amphibianEvolution()
			}
		case 1:
			if roll == rollPhysical {
				an.aquaticPhysical()
			}
			if roll == rollSocial {
				an.aquaticSocial()
			}
			if roll == rollEvolution {
				an.aquaticEvolution()
			}
		case 2:
			if roll == rollPhysical {
				an.avianPhysical()
			}
			if roll == rollSocial {
				an.avianSocial()
			}
			if roll == rollEvolution {
				an.avianEvolution()
			}
		case 3:
			if roll == rollPhysical {
				an.fungalPhysical()
			}
			if roll == rollSocial {
				an.fungalSocial()
			}
			if roll == rollEvolution {
				an.fungalEvolution()
			}
		case 4:
			if roll == rollPhysical {
				an.insectPhysical()
			}
			if roll == rollSocial {
				an.insectSocial()
			}
			if roll == rollEvolution {
				an.insectEvolution()
			}
		case 5:
			if roll == rollPhysical {
				an.mammalPhysical()
			}
			if roll == rollSocial {
				an.mammalSocial()
			}
			if roll == rollEvolution {
				an.mammalEvolution()
			}
		case 6:
			if roll == rollPhysical {
				an.reptilePhysical()
			}
			if roll == rollSocial {
				an.reptileSocial()
			}
			if roll == rollEvolution {
				an.reptileEvolution()
			}
		}
	}
}

func (an *animal) amphibianPhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrDexterity, 1)
	case 2:
		an.addCharacteristic(chrStrength, 2)
	case 3:
		an.addCharacteristic(chrEndurance, 1)
	case 4:
		an.addCharacteristic(chrEndurance, 2)
	case 5:
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) amphibianSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 4)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklDeception)
	case 4:
		an.addCharacteristic(chrInstinct, 1)
	case 5:
		an.addSkills(sklDeception)
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) amphibianEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.damageDice++
	case 2:
		an.armorScore++
	case 3:
		an.addCharacteristic(chrIntelligence, 1)
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}

func (an *animal) aquaticPhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrDexterity, 1)
	case 2:
		an.addCharacteristic(chrDexterity, 2)
	case 3:
		an.addCharacteristic(chrEndurance, 2)
	case 4:
		an.addCharacteristic(chrEndurance, 4)
	case 5:
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) aquaticSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 2)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklDeception)
	case 4:
		an.addCharacteristic(chrInstinct, 1)
	case 5:
		an.addCharacteristic(chrPack, 5)
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) aquaticEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.damageDice++
	case 2:
		an.addSkills(sklMeleeNW)
	case 3:
		an.addCharacteristic(chrInstinct, 1)
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}

func (an *animal) avianPhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrDexterity, 1)
	case 2:
		an.addCharacteristic(chrDexterity, 2)
	case 3:
		an.addCharacteristic(chrEndurance, 2)
	case 4:
		an.addCharacteristic(chrEndurance, 4)
	case 5:
		an.addCharacteristic(chrDexterity, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) avianSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 2)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklDeception)
	case 4:
		an.addCharacteristic(chrInstinct, 2)
	case 5:
		an.addCharacteristic(chrPack, 5)
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) avianEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.damageDice++
	case 2:
		an.addSkills(sklMeleeNW)
	case 3:
		an.addCharacteristic(chrInstinct, 1)
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}

func (an *animal) fungalPhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrEndurance, 1)
	case 2:
		an.addCharacteristic(chrEndurance, 2)
	case 3:
		an.addCharacteristic(chrEndurance, 4)
	case 4:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("2d6").Sum())
	case 5:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) fungalSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 2)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklStealth)
	case 4:
		an.addCharacteristic(chrInstinct, 2)
	case 5:
		an.addCharacteristic(chrPack, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) fungalEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.damageDice++
	case 2:
		an.addSkills(sklMeleeNW)
	case 3:
		an.addCharacteristic(chrInstinct, 1)
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}

func (an *animal) insectPhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrEndurance, 1)
	case 2:
		an.addCharacteristic(chrEndurance, 2)
	case 3:
		an.armorScore++
	case 4:
		an.addCharacteristic(chrEndurance, an.dicepool.RollNext("1d6").Sum())
		an.armorScore++
	case 5:
		an.addCharacteristic(chrStrength, an.dicepool.RollNext("1d6").Sum())
		an.armorScore++
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) insectSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 2)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklStealth)
	case 4:
		an.addCharacteristic(chrInstinct, 2)
	case 5:
		an.addCharacteristic(chrPack, an.dicepool.RollNext("1d6").Sum())
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) insectEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.damageDice++
	case 2:
		an.addSkills(sklMeleeNW)
	case 3:
		an.addCharacteristic(chrPack, 1)
		an.addCharacteristic(chrInstinct, an.dicepool.RollNext("1d6").Sum())
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}

func (an *animal) mammalPhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrStrength, 1)
	case 2:
		an.addCharacteristic(chrEndurance, 2)
	case 3:
		an.addCharacteristic(chrDexterity, 1)
	case 4:
		an.size++
	case 5:
		an.addCharacteristic(chrStrength, 2)
		an.addCharacteristic(chrEndurance, 1)
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) mammalSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 2)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklStealth)
	case 4:
		an.addCharacteristic(chrIntelligence, 1)
	case 5:
		an.addCharacteristic(chrPack, 1)
		an.addSkills(sklSurvival)
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) mammalEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.setAnimalBehaviour()
		beh1 := an.behaviour
		an.setAnimalBehaviour()
		beh2 := an.behaviour
		an.reactionDM = 0
		an.behaviour = beh1 + "-" + beh2
	case 2:
		an.addSkills(sklMeleeNW)
	case 3:
		an.addCharacteristic(chrPack, 1)
		an.addCharacteristic(chrInstinct, 1)
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}

func (an *animal) reptilePhysical() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrStrength, 1)
	case 2:
		an.addCharacteristic(chrEndurance, 2)
	case 3:
		an.addCharacteristic(chrDexterity, 2)
	case 4:
		an.size++
	case 5:
		an.addCharacteristic(chrStrength, 2)
		an.addCharacteristic(chrEndurance, 1)
	case 6:
		an.addSkills(sklMeleeNW)
	}
}
func (an *animal) reptileSocial() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addCharacteristic(chrPack, 2)
	case 2:
		an.addCharacteristic(chrInstinct, 1)
	case 3:
		an.addSkills(sklStealth)
	case 4:
		an.addSkills(sklSurvival)
	case 5:
		an.addCharacteristic(chrPack, 1)
		an.addCharacteristic(chrInstinct, 1)
	case 6:
		an.addSkills(sklRecon)
	}
}
func (an *animal) reptileEvolution() {
	switch an.dicepool.RollNext("1d6").Sum() {
	case 1:
		an.addExoticWeapon()
	case 2:
		an.addSkills(sklMeleeNW)
	case 3:
		an.addCharacteristic(chrPack, 1)
		an.addCharacteristic(chrDexterity, 1)
	case 4:
		an.addGenerationRolls(rollPhysical)
	case 5:
		an.addGenerationRolls(rollSocial)
	case 6:
		an.addExoticWeapon()
	}
}
