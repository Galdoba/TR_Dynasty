package encounter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/devtools/cli/user"
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
	an := NewAnimal(0)
	fmt.Println(an)
	fmt.Println("Test End")
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
	attackIF           int
	fleeIF             int
	attackType         string
	damageDice         int
	damageMod          int
	exoticWeapon       []string
	weapon             []string
	exoticWeaponEffect []string
	descr              string
	notes              string
	size               int
	generationRollDm   int
	evolutionRollDM    int
	expectedRolls      []string
	weight             float64
}

func NewAnimal(seed ...int64) animal {
	seed64 := int64(0)
	if len(seed) == 0 {
		seed64 = time.Now().UnixNano()
	} else {
		seed64 = seed[0]
	}
	an := animal{}

	an.skills = make(map[string]int)
	an.characteristics = make(map[string]int)
	an.dicepool = dice.New(seed64)

	an.selectTerrain()
	an.selectClass()
	an.setMovement()
	an.setDiet()
	an.setAnimalBehaviour()
	an.evolutionBase()
	//quirks := an.rollQuirks()
	benefits := an.rollBenefits()
	an.generationRolls()
	an.getBehaviourBenefits()
	an.applyBenefits(benefits...)
	an.size = an.size + an.dicepool.RollNext("2d6").Sum()
	an.defineChrs()
	return an
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
	//--------------------------
	fmt.Println("Animal Stats:")
	fmt.Println(chrStrength, an.characteristics[chrStrength])
	fmt.Println(chrDexterity, an.characteristics[chrDexterity])
	fmt.Println(chrEndurance, an.characteristics[chrEndurance])
	fmt.Println(chrIntelligence, an.characteristics[chrIntelligence])
	fmt.Println(chrInstinct, an.characteristics[chrInstinct])
	fmt.Println(chrPack, an.characteristics[chrPack])
	fmt.Println("Total HP:", an.characteristics[chrStrength]+an.characteristics[chrDexterity]+an.characteristics[chrEndurance])
	fmt.Println("Weight:", an.weight)
	fmt.Print("Attack: ", an.damageDice, "d6\n")
	fmt.Print("Armor: ", an.armorScore, "\n")
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

func (an *animal) selectTerrain() {
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
	t := -1
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

func (an *animal) selectClass() {
	fmt.Println("Select Class:")
	fmt.Print("Roll Random  = [", -1, "]\n")
	fmt.Print("Amphibians   = [", 0, "]\n")
	fmt.Print("Aquatic      = [", 1, "]\n")
	fmt.Print("Avians       = [", 2, "]\n")
	fmt.Print("Fungals      = [", 3, "]\n")
	fmt.Print("Insect       = [", 4, "]\n")
	fmt.Print("Mammals      = [", 5, "]\n")
	fmt.Print("Reptiles     = [", 6, "]\n")

	t := -1
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
		//Exotic Weapon
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
		//Exotic Weapon
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
			rArr = append(rArr, r)
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		an.weapon = append(an.weapon, "Horns")
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
		//Exotic Weapon
		//Exotic Weapon
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
		rawData = "Pouncer -1 Filter -1 Carrion-Eater -1 "
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
		rawData = "Eater -1 Filter -1 Carrion-Eater -1 "
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
		rawData = "Hunter -1 Intimidator -1 Carrion-Eater -1 "
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
		rawData = "Hunter -2 Intermittent -2 Carrion-Eater -1 "
	case 20:
		rawData = "Hunter -1 Intermittent -1 Carrion-Eater +0 "
	case 21:
		rawData = "Hunter -0 Intermittent -1 Eater +0 "
	case 22:
		rawData = "Siren +0 Intermittent +0 Reducer +0 "
	case 23:
		rawData = "Siren +1 Grazer -1 Reducer -1 "
	case 24:
		rawData = "Killer +0 Grazer -2 Reducer -2 "
	case 25:
		rawData = "Pouncer +0 Eater -2 Carrion-Eater -1 "
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
		rawData = "Pouncer +0 Eater -2 Carrion-Eater -1 "
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
		rawData = "Pouncer +0 Gatherer -1 Carrion-Eater -1 "
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
func (an *animal) getBehaviourBenefits() {
	if strings.Contains(an.behaviour, "Carrion-Eater") {
		an.addCharacteristic(chrInstinct, 2)
		an.size = an.size - 2
		// If a Carrion-Eater has an Exotic Weapon(s), its first one is always Diseased Attack.
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		//Exotic Weapon
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
		//Exotic Weapon
	}
}
