package asset

type Skill interface {
	Value() (int, error)
	Modifier() (int, error)
	SetCharacteristicValue(int)
}

func NewSkill(name string) Characteristic {
	c := asset{}
	c.name = name
	return &c
}
