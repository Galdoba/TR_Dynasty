package autoGM

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/autoGM/encounters"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func AutoGM() {
	//mission.Test()
	//RunACFlowchart(9)
	w := wrld.PickWorld()
	fmt.Println(w.SecondSurvey())
	e := encounters.New()
	e.SetWorld(&w)
	key := e.RollKeySpaceEncounter()
	fmt.Println("key = ", key, encounters.EncounterMapCRB(key))
	fmt.Println("Danger Case:", dice.Flux())
	//e.RollShipEncounterMGT1CG()
	//e.Express()
}

type shipEncounter struct {
	code          string
	encounter     string
	suggestedShip string
}

/*
arrival system check list:
space travel event roll
ship encounter roll
pirate event roll

In System Activity check list (Merchant):
I	Preapare for jump.
		A. Check charts for new System.
		B. Plot Course.
II	Jump to System.
III	Arrive in Star System.
	A. Scan area for potential danger, problems or other data.
	B. Access local communicator directory.										новости, сообщения, контакты и враги в системе.
	C. Set course in system.
		1. Gas Gigant (go to IV).
		2. Mainworld (go to V).
		3. Other world (go to VI).
	D. Possible ship encounter.													бросок space encounter исходя из параметров текущей системы.
IV	Local Gas Gigant.
	A. Move to Orbit/Possible ship encounter.									Проверка пилота и броски space encounter за каждый тень перелета.
	B. Refuel and return to orbit.
	C. Set course in system.
		1. Mainworld (go to V).
		2. Other world (go to VI).
		3. Jump Point (go to VII).
V	Mainworld.
	A. Achive orbit.
	B. Orbital Starport (unstreamlined ships).									Проверка Legal Encounter (досмотр корабля)
	C. Surface Starport.
		1. Unload high passengers.
		2. Unload mail.
		3. Unload middle passengers.
		4. Unload cargo.
		5. Defrost and unload low passengers.
		6. Conclude low lottery.
	D. Refit and maintainance.
		1. Refuel from Starport.
		2. Renew ship life support.
	E. Commodity activity.														загрузка торгового модуля.
		1. Sell speculative cargo.
		2. Buy speculative cargo.
	F. Ship's business.
		1. Pay Berthing costs.
		2. Pay Banking payment.
		3. Pay maintainance fund.
		4. Pay crew salaries.
	G. Miscellanious activity.
		1. Patron encounter.													поиск и генерация мелких квестов.
		2. Planetary exploration.												генерирование точек интересов на планете.
	H. Preapare for Departure.
		1. Load Cargo.
		2. Load low passengers.
		3. Load middle passengers.
		4. Load high passengers.
		5. Load mail.
		6. Collect income for all aspects of current trip.
	J. Departure
		1. Lift-off.
		2. Achive orbit.
		3. Set cource.
			a. Other world (go to VI).
			b. Gas Gigant (go to IV).
			c. Jump Point (go to VII).
		4. Possible ship encounter.
VI	Other World.
	A. Achieve orbit.
		1. Scan and map surface.
		2. Investigate interesting details.
	B. Move to surface.
		1. Refuel.
		2. Explore.
	C. Departure.
		1. Lift-off.
		2. Achive orbit.
		3. Set cource.
			a. Mainworld (go to V).
			b. Other world (go to VI).
			c. Gas Gigant (go to IV).
			d. Jump Point (go to VII).
		4. Possible ship encounter.
VII	Jump Point.
	A. Approach jump point.
	B. Handle necessary calculations (go to I).
	C. Jump (go to II).
*/
