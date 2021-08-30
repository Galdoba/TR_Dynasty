package assets

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	C1 = iota
	C2
	C3
	C4
	C5
	C6
	CS
	CP
	Strength
	Dexterity
	Agility
	Grace
	Endurance
	Stamina
	Vigor
	Intelligence
	Education
	Training
	Instinct
	Social
	Charisma
	Caste
	Sanity
	Psionics
)

type Characteristic struct {
	value        int
	genetics     int
	positionCode int
	code         int
	alias        string // not needed. TODO: get alias via code
	charDice     int    //сколько кубов ролится для создания х-ки
	err          error
}

func NewCharacteristic(code int, diceQty int) *Characteristic {
	c := Characteristic{}
	c.code = code
	c.charDice = diceQty
	c.positionCode, c.err = positionCodeOf(code)
	if c.charDice < 1 {
		c.err = fmt.Errorf("characteristic dice can not be less than 1")
	}
	if c.err != nil {
		return &c
	}
	for i := 0; i < diceQty; i++ {
		roll := dice.Roll1D()
		if i == 0 {
			c.genetics = roll
		}
		c.value = c.value + roll
	}
	return &c
}

func positionCodeOf(code int) (int, error) {
	pc := 0
	err := fmt.Errorf("code '%v' is unknown", code)
	switch code {
	default:
		return 0, err
	case Strength:
		pc = C1
	case Dexterity, Agility, Grace:
		pc = C2
	case Endurance, Stamina, Vigor:
		pc = C3
	case Intelligence:
		pc = C4
	case Education, Training, Instinct:
		pc = C5
	case Social, Charisma, Caste:
		pc = C6
	case Psionics:
		pc = CP
	case Sanity:
		pc = CS
	}
	return pc, nil
}

func (c *Characteristic) Value() int {
	return c.value
}

func (c *Characteristic) PositionCode() int {
	return c.positionCode
}

func (c *Characteristic) Genetics() int {
	return c.genetics
}

func (c *Characteristic) Name() string {
	switch c.code {
	default:
		return "ERROR"
	case Strength:
		return "Strength"
	case Dexterity:
		return "Dexterity"
	case Agility:
		return "Agility"
	case Grace:
		return "Grace"
	case Endurance:
		return "Endurance"
	case Stamina:
		return "Stamina"
	case Vigor:
		return "Vigor"
	case Intelligence:
		return "Intelligence"
	case Education:
		return "Education"
	case Training:
		return "Training"
	case Instinct:
		return "Instinct"
	case Social:
		return "Social"
	case Charisma:
		return "Charisma"
	case Caste:
		return "Caste"
	case Sanity:
		return "Sanity"
	case Psionics:
		return "Psi"
	}
}

func (c *Characteristic) Abb() string {
	switch c.code {
	default:
		return "ERROR"
	case Strength:
		return "Str"
	case Dexterity:
		return "Dex"
	case Agility:
		return "Agi"
	case Grace:
		return "Gra"
	case Endurance:
		return "End"
	case Stamina:
		return "Sta"
	case Vigor:
		return "Vig"
	case Intelligence:
		return "Int"
	case Education:
		return "Edu"
	case Training:
		return "Tra"
	case Instinct:
		return "Ins"
	case Social:
		return "Soc"
	case Charisma:
		return "Cha"
	case Caste:
		return "Cas"
	case Sanity:
		return "San"
	case Psionics:
		return "Psi"
	}
}

func (c *Characteristic) UseAs(charCode int) (int, error) {
	if charCode == c.code {
		return c.value, nil
	}
	c2 := NewCharacteristic(charCode, 1)
	if c2.positionCode != c.positionCode {
		return 0, fmt.Errorf("%v cannot be substituded with %v", c.Name(), c2.Name())
	}
	switch {
	case c.positionCode == C2 && c2.positionCode == C2:
		return halfValueRoundUp(c.value), nil
	case c.positionCode == C3 && c2.positionCode == C3:
		return halfValueRoundUp(c.value), nil
	case c.code == Education && c2.code == Training:
		return halfValueRoundUp(c.value), nil
	case c.code == Education && c2.code == Instinct:
		return 4, nil
	case c.code == Training && c2.code == Education:
		return halfValueRoundUp(c.value), nil

	}
}

func halfValueRoundUp(val int) int {
	return val/2 + val%2
}
