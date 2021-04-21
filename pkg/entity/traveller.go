package entity

import (
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/entity/asset"
)

type Traveller struct {
	Dice   *dice.Dicepool
	Info   map[string]string
	Chrctr map[string]asset.Characteristic
}

func NewTraveller(seed ...string) Traveller {
	t := Traveller{}
	t.Dice = dice.New()
	for i := range seed {
		switch i {
		case 0:
			t.Dice.SetSeed(seed[0])
		}
	}
	t.rollCharcteristics()
	return t
}

func (t *Traveller) rollCharcteristics() {
	t.Chrctr = make(map[string]asset.Characteristic)
	for _, val := range listCharacteristics() {
		score := t.Dice.RollNext("2d6").Sum()
		t.Chrctr[val] = asset.NewCharacteristic(val)
		t.Chrctr[val].SetCharacteristicValue(score)
	}
}
