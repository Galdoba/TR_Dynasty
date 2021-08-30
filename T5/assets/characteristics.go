package assets

import "fmt"

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
	alias        string // not needed. TODO: get alias via code
	charDice     int    //сколько кубов ролится для создания х-ки
	err          error
}

func NewCharacteristic(code int, dice int) *Characteristic {
	c := Characteristic{}
	switch code {
	default:
		c.err = fmt.Errorf("code '%v' is unknown", code)
	case Strength:
		c.positionCode = C1
	case Dexterity, Agility, Grace:
		c.positionCode = C2
	case Endurance, Stamina, Vigor:
		c.positionCode = C3
	case Intelligence:
		c.positionCode = C4
	case Education, Training, Instinct:
		c.positionCode = C5
	case Social, Charisma, Caste:
		c.positionCode = C6
	case Psionics:
		c.positionCode = CP
	case Sanity:
		c.positionCode = CS
	}
	for i := 0; i < dice; i++ {
		c.value = dice.Roll1D()
	}
	// 649 8 12
	// 165 +0 +2
	//  84 +2 +8
	return &c
}
