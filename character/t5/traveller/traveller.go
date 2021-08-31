package traveller

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
)

type TravellerT5 struct {
	notHuman       bool
	characteristic map[int]assets.Characteristic
	skills         map[int]assets.Skill
	knowledges     map[int]assets.Knowledge
	homeworld      string
	err            error
}

func NewTravellerT5() *TravellerT5 {
	trv := TravellerT5{}
	switch trv.notHuman {
	case true:
		trv.err = fmt.Errorf("generation of non-Human not implemented")
		return &trv
	case false:
		charList := []int{
			assets.Strength,
			assets.Dexterity,
			assets.Endurance,
			assets.Intelligence,
			assets.Education,
			assets.Social,
			assets.Psionics,
			assets.Sanity,
		}
		for _, val := range charList {
			trv.characteristic[val] = *assets.NewCharacteristic(val, 2)
			if trv.characteristic[val].Err != nil {
				trv.err = trv.characteristic[val].Err
				return &trv
			}
		}
	}
	return &trv
}
