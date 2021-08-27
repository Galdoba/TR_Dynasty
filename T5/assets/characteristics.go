package assets

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
	char     int
	alias    string
	charDice int //сколько кубов ролится для создания х-ки
	err      error
}

func NewCharacteristic(code int, dice ...int) *Characteristic {
	c := Characteristic{}
	// 649 8 12
	// 165 +0 +2
	//  84 +2 +8
	return &c
}
