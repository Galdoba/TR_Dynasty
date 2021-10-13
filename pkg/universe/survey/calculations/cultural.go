package calculations

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func Cultural(uwp, seed string, ix int) string {
	hex := strings.Split(uwp, "")
	pop := ehex.New(hex[4])
	tl := ehex.New(hex[8])
	if pop.Value() == 0 {
		return "[0000]"
	}
	h, a, st, sy := 1, 1, 1, 1
	d := dice.New().SetSeed(seed)
	h = pop.Value() + d.FluxNext()
	if h < 1 {
		h = 1
	}
	a = pop.Value() + ix
	if a < 1 {
		a = 1
	}
	st = d.FluxNext() + 5
	if st < 1 {
		st = 1
	}
	sy = d.FluxNext() + tl.Value()
	if sy < 1 {
		sy = 1
	}
	return "[" + ehex.New(h).String() + ehex.New(a).String() + ehex.New(st).String() + ehex.New(sy).String() + "]"
}

func CxValid(cx, uwp string) bool {
	culturalInvalid := []string{"[????]", "", "----", "[]"}
	for _, val := range culturalInvalid {
		if cx == val {
			return false
		}
	}
	if len(cx) != 6 {
		return false
	}
	hexu := strings.Split(uwp, "")
	pop := ehex.New(hexu[4])
	if pop.Value() == 0 && cx != "[0000]" {
		return false
	}
	hex := strings.Split(cx, "")
	if !hetValid(ehex.New(hex[1]).Value()) {
		return false
	}
	if !accValid(ehex.New(hex[2]).Value()) {
		return false
	}
	if !strValid(ehex.New(hex[3]).Value()) {
		return false
	}
	if !symValid(ehex.New(hex[4]).Value()) {
		return false
	}
	return true
}

func reRollHeterogenity(d *dice.Dicepool, pop ehex.DataRetriver) int {
	h := pop.Value() + d.FluxNext()
	if h < 1 {
		h = 1
	}
	return h
}

func hetValid(h int) bool {
	if h > 20 {
		return false
	}
	return true
}

func accValid(a int) bool {
	if a > 21 {
		return false
	}
	return true
}

func reRollAcceptance(d *dice.Dicepool, ix int, pop ehex.DataRetriver) int {
	return 0
}

func strValid(st int) bool {
	if st > 11 {
		return false
	}
	return true
}

func symValid(sy int) bool {
	if sy > 23 {
		return false
	}
	return true
}
