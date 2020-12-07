package encounter

import (
	"errors"
	"fmt"
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
	exoticAttack       string
	exoticAttackEffect string
	descr              string
	notes              string
}

func NewAnimal(seed ...int64) animal {
	seed64 := int64(0)
	if len(seed) == 0 {
		seed64 = time.Now().UnixNano()
	} else {
		seed64 = seed[0]
	}
	an := animal{}
	an.dicepool = dice.New(seed64)
	an.selectTerrain()
	an.selectClass()
	an.setMovement()
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

func (an *animal) setMovement() {

}
