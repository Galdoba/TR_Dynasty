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
	attackType         string
	damageDice         int
	damageMod          int
	exoticWeapon       []string
	exoticWeaponEffect []string
	descr              string
	notes              string
	size               int
	generationRollDm   int
	evolutionRollDM    int
	expectedRolls      []string
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
	an.evolutionBase()
	//quirks := an.rollQuirks()
	benefits := an.rollBenefits()
	an.applyBenefits(benefits...)
	return an
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
	an.addGenerationRolls(rollBehavior)
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
	moveMap[0] = []string{"W", "W", "W", "W", "W +2", "F –6"}
	moveMap[1] = []string{"W", "W", "W", "W +2", "W +4", "F –6"}
	moveMap[2] = []string{"W", "W", "W", "W", "F –4", "F –6"}
	moveMap[3] = []string{"W", "W", "W", "W +2", "F –4", "F –6"}
	moveMap[4] = []string{"W", "W", "W", "F –2", "F –4", "F –6"}
	moveMap[5] = []string{"W", "W", "W", "W", "F –4", "F –6"}
	moveMap[6] = []string{"W", "W", "W", "W", "W", "F –6"}
	moveMap[7] = []string{"W", "W", "W", "W", "W +2", "F –6"}
	moveMap[8] = []string{"W", "W", "W", "W +2", "W +4", "F –6"}
	moveMap[9] = []string{"W", "W", "W", "W +2", "F –4", "F –6"}
	moveMap[10] = []string{"S –6", "S", "W", "W", "F –4", "F –6"}
	moveMap[11] = []string{"S +1", "S +1", "W", "W", "F –4", "F –6"}
	moveMap[12] = []string{"S –4", "S +2", "W", "W", "W", "F –6"}
	moveMap[13] = []string{"S +4", "S +2", "S", "S", "F –4", "F –6"}
	moveMap[14] = []string{"S +6", "S +4", "S +2", "S", "F –4", "F –6"}
	moveMap[15] = []string{"S +8", "S +6", "S +4", "S +2", "S", "S –2"}
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

func (an *animal) generationRolls(rollType ...string) {

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

/*
2 Whenever packs of these animals make any noise at all, they all make the exact same sound simultaneously several times in a row.
3 Apparently blind, these amphibians have no visible eyes or means of sight. Recon –1
4 These animals make no sound at all, even when they move in natural surroundings. They gain Stealth 0 as a result.
5 The colours of this amphibian’s hide are vivid and clashing, a sort of natural reverse camouflage. Natural predators dislike this display and leave it alone.
6 Seemingly everywhere, forms of this animal can be found in virtually every habitat type on their world. Survival +2
7 These amphibians emit a natural pheromone that other animals find highly attractive. They gain the Siren Behaviour in addition to their own.
8 When threatened, these amphibians emit a piercing scream that sounds like a sentient creature in terrible pain.
9 On rare occasions, these amphibians swarm viciously. Roll 2d6 at the start of any encounter. On a 2, replace their Behaviour with Killer.
10 The skin of these animals is naturally coated in a thick, foul-smelling emulsion. They possess the Stench Exotic Weapon.
11 Unusually for its kind, these amphibians have developed a rigid shell over their forelimbs and torsos. Armour +2
*/
