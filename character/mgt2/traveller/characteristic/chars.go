package characteristic

import "github.com/Galdoba/TR_Dynasty/T5/ehex"

type characteristic struct {
	name         string
	abbreviation string
	chrDice      int
	value        ehex.Ehex
	genetics     int
}

//trv.Characteristic(C1)
//trv.Chr(C1)
//trv.Get(C1)
