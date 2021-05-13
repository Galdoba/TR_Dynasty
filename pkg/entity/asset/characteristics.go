package asset

type Characteristic interface {
	Value() int
	Modifier() int
	SetCharacteristicValue(int)
}

func NewCharacteristic(name string) Characteristic {
	c := asset{}
	c.name = name
	return &c
}

func (a *asset) Value() int {
	if len(a.numericalValues) < 1 {
		return 0
	}
	return a.numericalValues[0]
}

func (a *asset) Modifier() int {
	if len(a.numericalValues) < 2 {
		return 0
	}
	return a.numericalValues[1]
}

func (a *asset) SetCharacteristicValue(newVal int) {
	for len(a.numericalValues) < 2 {
		a.numericalValues = append(a.numericalValues, 0)
	}
	a.numericalValues[0] = newVal
	a.numericalValues[1] = characteristicModifier(newVal)
}

//////////////////////

func characteristicModifier(val int) int {
	switch val {
	case 1, 2:
		return -2
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	case 9, 10, 11:
		return 1
	case 12, 13, 14:
		return 2
	}
	if val < 1 {
		return -3
	}
	return 3
}
