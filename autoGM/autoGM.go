package autoGM

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/autoGM/encounters"
	"github.com/Galdoba/TR_Dynasty/dice"
)

func AutoGM() {
	//mission.Test()
	//RunACFlowchart(9)

	e := encounters.New()
	e.RollShipEncounterMGT1CG()
	e.Express()
}

func RunACFlowchart(tn int) {
	var event []string
	event = append(event, "1. Job Hunting (Planetside Events, page 7)")
	event = append(event, "2. Preparations (repeat previous step)")
	event = append(event, "3. Jump Travel (Onboard Events, page 60)")
	//Space Travel
	event = append(event, "4a. Space Events (page 32)")
	event = append(event, "4b. Life Events (page 67)")
	event = append(event, "5. Ground Travel (Planetside Events, page 7)")
	event = append(event, "6. Destination (Any)")
	event = append(event, "7. Return (repeat steps 3,4 and 5 in reverse order)")
	//Resting
	event = append(event, "8a. Planetside, page 7")
	event = append(event, "8b. Life events, page 67")
	event = append(event, "8c. Adventure Hooks, page 71"+" - Event("+dice.RollD66()+")")
	for i := range event {
		if dice.Roll("2d6").ResultTN(tn) {
			fmt.Println("event Happened:", event[i])
		}
	}
}

type shipEncounter struct {
	code          string
	encounter     string
	suggestedShip string
}

/*
PlanetSideEvent{
	type (Urban/Rural/Wilderness)
}


*/

/*
arrival system check list:
space travel event roll
ship encounter roll
pirate event roll





*/
