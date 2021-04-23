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

func (t *Traveller) Sheet() string {
	sh := ""
	return sh
}

/*
+---INFO----------------------------+---ARMOR---------------------------------------------------------------------------+
| Name: [Traveller Name]            | TYPE              | RAD | PROTECTION | KG |             INSTALLED MODS            |
| UPP : 123456-7                    | Armor name 1      | XXX |     XX     | XX | [Loooooooooooooooooooong Description] |
| Rads: xxxx             Age: XXX   | Armor name 2      | XXX |     XX     | XX | Options:             [No Description] |
| Species: Human                    | Armor name 2      | XXX |     XX     | XX | Options:             [No Description] |
| Species Traits: _Mandatory_______ |                   |     |            |    |                                       |
=                 _Additionals_____ =  ____Additional Armor data__              |                                       =
| Homeworld: [Homeworld Name      ] +---FINANCES------------+---CAREER SUMMARY--+
| s123456-7 __ __ __ __ __ __ __ __ | Pension:              | Law Enforcement 2 |
+---CHARACTERISTICS-----------------+   XXXXXX Cr/Year      | Scavenger       1 |
| [XX] (+0) | [XX} (+0) | [XX] (-1) | Debt:                 | Marine Support  1 |
| Strenghh  | Dexterity | Endurance |   XXXXXXX xCr         | Drifter         8 |
+-----------+-----------+-----------+ Cash on Hand:         |                   |
| [XX] (+0) | [XX} (+0) | [XX] (-1) |   XXXXXXX xCr         |                   |
| Intellect | Education |   Social  | Ship Payments:        |                   |
+-----------+-----------+-----------+   XXXXXXX xCr/Month   |                   |
| Psionic Powers:         [XX] (-3) | Living Cost:          |                   |
|        [Untested or talents list] |   XXXXXXX  Cr/Month   |                   |
+---SKILLS--------------------------+-----------------------+-------------------+
| Jack of all Trades
| Pilot (capital ships) x
| 123456789012345678901234 | 123456789012345678901234 | 123456789012345678901234 | 123456789012345678901234 |
+----------------------------------------------------------------------------------------------------------------------+
123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890

*/
