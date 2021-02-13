package encounters

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/utils"
)

type encounterEvent struct {
	dicepool   *dice.Dicepool
	world      *wrld.World
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

func (e *encounterEvent) SetWorld(w *wrld.World) {
	e.world = w
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

func (e *encounterEvent) RollKeySpaceEncounter() string {
	if e.world == nil {
		return "World not defined!"
	}
	dm := e.world.ImportanceVal()
	//fmt.Print("Min:", 1+dm, " Max:", 6+dm, "\n")
	d1 := e.dicepool.RollNext("1d6").DM(dm).Sum()
	d1 = utils.BoundInt(d1, 0, 9)
	d2 := e.dicepool.RollNext("1d6").Sum()
	key := strconv.Itoa(d1) + strconv.Itoa(d2)
	return key
}

/*
логика:
arrival
--
spaceFlight
HighPort Dock





раз в день в космосе делается бросок d6. если выпадает 6 - ролим энкаунтер.
если энкаунтер (
	берем модификатор от орбитальной зоны:
	highport (DM+3): The space near an orbital starport
	high-Traffic space (DM+2): planet have 'In' tag
	settled space (DM+1): default
	border systems (DM+0): LawLevel <= 6
	Wild space (DM-1): Amber or Red worlds. LawLevel <= 3
	empty space (DM-4): Untravelled space or unexplored systems.

	ролим код энкаунтера

)
*/
