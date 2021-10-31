package actions

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

const (
	ROUT          = "ROUT"
	DEFEAT        = "DEFEAT"
	SETBACK       = "SETBACK"
	DEADLOCK      = "DEADLOCK"
	ADVANCE       = "ADVANCE"
	VICTORY       = "VICTORY"
	TOTALVICTORY  = "TOTAL VICTORY"
	DM_Protection = "Protection"
	DM_Firepower  = "Firepower"
	DM_Position   = "Position"
	DM_RSize      = "Relative Size"
	DM_Combat     = "Combat"
)

/*
TODO:
Setup sides

Outcome

Casualties




*/

type combat struct {
	sides       []combatSide
	dms         map[string]int
	totalPhases int
	dice        *dice.Dicepool
	outcomeStr  string
	sitRep      string
}

type combatSide struct {
	side         int
	participants int
}

func NewCombat(seed ...int) error {
	comb := combat{}
	comb.dms = make(map[string]int)
	comb.dms[DM_Firepower] = userInputInt("Define Firepower Modifier")
	comb.dms[DM_Protection] = userInputInt("Define Protection Modifier")
	comb.dms[DM_Position] = userInputInt("Define Position Modifier")
	ppar := userInputInt("Define Traveller's side numbers")
	playerSide := newSide(0, ppar)
	epar := userInputInt("Define Enemy's side numbers")
	enemySide := newSide(1, epar)
	comb.dms[DM_RSize] = numbersDM(playerSide.participants, enemySide.participants)
	//comb.globalMod = defineGlobalModifier()
	comb.sides = append(comb.sides, playerSide)
	comb.sides = append(comb.sides, enemySide)
	dice := dice.New()
	if len(seed) > 0 {
		dice.SetSeed(seed[0])
	}
	comb.dice = dice
	comb.Resolve()
	comb.Casualties()
	return nil
}

func numbersDM(t, e int) int {
	m := 1
	max, min := t, e
	if e > t {
		m = -1
		max, min = e, t
	}
	dm := 0
	rs := (max * 100) / min
	if rs >= 150 {
		dm = 1
	}
	if rs >= 200 {
		dm = 2
	}
	if rs >= 300 {
		dm = 3
	}
	return dm * m
}

func userInputInt(msg string) int {
	status := fmt.Errorf("null input err")
	val := 0
	for status != nil {
		fmt.Printf(msg + ": ")
		val, status = user.InputInt()
		//if status != nil && status.Error() != "EOF" {
		//	fmt.Printf("Error: %v\n", status.Error())
		//}
	}
	return val
}

func (c *combat) globalMod() int {
	return c.dms[DM_Firepower] + c.dms[DM_Protection] + c.dms[DM_Position] + c.dms[DM_RSize]
}

func newSide(side, participants int) combatSide {
	cs := combatSide{}
	cs.side = side
	cs.participants = participants
	return cs
}

func (comb *combat) Resolve() {
	fmt.Println("===Combat Resolution=========")
	for {
		comb.totalPhases++
		fmt.Printf("Begin Phase %v:\n", comb.totalPhases)
		fmsg := fmt.Sprintf("Define Traveller's side tactical mod for phase %v", comb.totalPhases)
		fMod := userInputInt(fmsg)
		emsg := fmt.Sprintf("Define Enemy's side tactical mod for phase %v", comb.totalPhases)
		eMod := userInputInt(emsg)
		rr := comb.dice.RollNext("2d6").Sum()
		fmt.Println("Roll:", rr)
		rResult := rr + comb.globalMod() + fMod - eMod + comb.dms[DM_Combat]
		fmt.Println("DM Relative Size:", comb.dms[DM_RSize])
		fmt.Println("DM Firepower    :", comb.dms[DM_Firepower])
		fmt.Println("DM Position     :", comb.dms[DM_Position])
		fmt.Println("DM Protection   :", comb.dms[DM_Protection])
		fmt.Println("DM Combat       :", comb.dms[DM_Combat])
		fmt.Println("Tactics (trvlrs):", fMod)
		fmt.Println("Tactics (enemy) :", eMod*-1)
		outcome := spellOutcome(rResult)
		fmt.Printf("Phase %v\nOutcome: %v (%v)\n", comb.totalPhases, rResult, outcome)
		fmt.Printf("=============================\n")
		switch outcome {
		case SETBACK:
			comb.dms[DM_Combat]--
		case ADVANCE:
			comb.dms[DM_Combat]++
		case ROUT, DEFEAT, VICTORY, TOTALVICTORY:
			return
		}
	}
}

func (c *combat) Casualties() {
	sideName := []string{"Traveller", "Enemy"}
	for side, name := range sideName {

		medicEffect := userInputInt(name + "'s side medic effect")
		phaseDm := c.totalPhases
		setupDM := c.globalMod()
		if side != 0 {
			setupDM = setupDM * -1
		}
		r := c.dice.RollNext("2d6").Sum()
		r += medicEffect - phaseDm + setupDM
		casDice := utils.BoundInt(15-r, 0, 15)
		fmt.Println(casDice)
		cas := c.dice.RollNext(strconv.Itoa(casDice) + "d6").Sum()
		fmt.Println(cas)
		fmt.Printf("[reporting %v's side casualties]\n", name)
		fmt.Printf("[%v percent from %v]\n", cas, c.sides[side].participants)
		wounded := int(float64(c.sides[side].participants)*(float64(cas)/100) + 1)
		casualtiesMap := make(map[string]int)
		for i := 0; i < wounded; i++ {
			switch c.dice.RollNext("1d9").Sum() {
			case 1, 2, 3:
				casualtiesMap["slight"]++
			case 4, 5, 6:
				casualtiesMap["minor"]++
			case 7:
				casualtiesMap["serious"]++
			case 8:
				casualtiesMap["critical"]++
			case 9:
				casualtiesMap["dead"]++
			}
		}
		fmt.Println("slight", casualtiesMap["slight"])
		fmt.Println("minor", casualtiesMap["minor"])
		fmt.Println("serious", casualtiesMap["serious"])
		fmt.Println("critical", casualtiesMap["critical"])
		fmt.Println("dead", casualtiesMap["dead"])

	}
}

func spellOutcome(i int) string {
	if i <= 0 {
		return ROUT
	}
	if i >= 15 {
		return TOTALVICTORY
	}
	switch i {
	default:
		return "OUTCOME Error"
	case 1, 2:
		return DEFEAT
	case 3, 4, 5:
		return SETBACK
	case 6, 7, 8:
		return DEADLOCK
	case 9, 10, 11:
		return ADVANCE
	case 12, 13, 14:
		return VICTORY
	}

}
