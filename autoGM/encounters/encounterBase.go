package encounters

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
)

type encounterEvent struct {
	dicepool   *dice.Dicepool
	name       string
	eventClass string
	location   string
	descr      string
	hook       string
}

const (
	EncounterShip = "Ship Encounter"
)

//New - Создает Энкаунтер
func New() encounterEvent {
	e := encounterEvent{}
	e.dicepool = dice.New()
	return e
}

func (e *encounterEvent) SetSeed(seed interface{}) {
	e.dicepool.SetSeed(seed)
}

func (e *encounterEvent) Express() string {
	str := "Class: " + e.eventClass
	str += "\nEvent: " + e.name
	if e.location != "" {
		str += "\nLocation: " + e.location
	}
	str += "\nDescription: " + e.descr
	if e.hook != "" {
		str += "\nHook: " + e.hook
	}
	fmt.Println(str)
	return str
}
