package actions

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

/*
TODO:
Setup sides

Outcome

Casualties




*/

type combat struct {
	sides       []combatSide
	globalMod   int
	totalPhases int
	dice        *dice.Dicepool
	outcomeStr  string
	sitRep      string
}

type combatSide struct {
	side         int
	participants int
}

func newCombat(seed ...int) *combat {
	comb := combat{}
	ppar := 30 //userinput
	playerSide := newSide(0, ppar)
	epar := 30 //userinput
	enemySide := newSide(1, epar)
	comb.globalMod = 1
	comb.sides = append(comb.sides, playerSide)
	comb.sides = append(comb.sides, enemySide)
	dice := dice.New()
	if len(seed) > 0 {
		dice.SetSeed(seed[0])
	}
	comb.dice = dice
	return &comb
}

func newSide(side, participants int) combatSide {
	cs := combatSide{}
	cs.side = side
	cs.participants = participants
	return cs
}

func (comb *combat) ResolvePhase() {
	fMod := 1
	eMod := 3
	outcome := comb.dice.RollNext("2d6").DM(comb.globalMod + fMod - eMod).Sum()
	fmt.Println(spellOutcome(outcome))
	comb.totalPhases++
}

func spellOutcome(i int) string {
	if i <= 0 {
		return "ROUT: The Traveller's side is totaly defeated, suffering heavy casualties and collapse of firepower. This fight is lost and another atempt, cannot be made without reinforcements, reorganisation and a period of recuperation."
	}
	if i >= 15 {
		return "TOTAL VICTORY: The Traveller's side is"
	}

	return "Error"
}
