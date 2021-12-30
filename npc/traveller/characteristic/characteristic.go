package characteristic

import (
	"fmt"
	"strings"
)

const (
	STRENGTH  = "STR"
	DEXTERITY = "DEX"
	ENDURANCE = "END"
	INTELLECT = "INT"
	EDUCATION = "EDU"
	SOCIAL    = "SOC"
	MORALE    = "MOR"
	LUCK      = "LCK"
	SANITY    = "SAN"
	CHARM     = "CRM"
	PSIONC    = "PSI"
	TERRITORY = "TER"
	CHARISMA  = "CHA"
)

func RollCharacteristics(race string) (map[string]Characteristic, error) {
	chrs := make(map[string]Characteristic)
	switch race {
	default:
		return chrs, fmt.Errorf("race '%v' not implemented")
	case strings.ToTitle("human"), strings.ToTitle("vilani"), strings.ToTitle("solomani"):

	}
}

type Characteristic struct {
	NameFull  string
	NameShort string
	Value     int
}
