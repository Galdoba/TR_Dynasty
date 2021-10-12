package calculations

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func EconomicExtention(uwp, pbg string, ix int) string {
	hex := strings.Split(uwp, "")
	tl := ehex.New(hex[8])
	pop := ehex.New(hex[4])
	hex2 := strings.Split(pbg, "")
	b := ehex.New(hex2[1])
	g := ehex.New(hex2[2])
	res, lab, inf, eff := 0, 0, 0, 0
	d := dice.New().SetSeed(uwp + pbg)
	//////////////////////////////////
	res = d.RollNext("2d6").Sum()
	if tl.Value() >= 8 {
		res += b.Value() + g.Value()
	}
	resH := ehex.New(res)
	//////////////////////////////////
	lab = pop.Value() - 1
	if lab == 0 {
		lab = 1
	}
	labH := ehex.New(lab)
	//////////////////////////////////
	switch pop.Value() {
	case 0:
		inf = 0
	case 1, 2, 3:
		inf = ix
	case 4, 5, 6:
		inf = d.RollNext("1d6").DM(ix).Sum()
	default:
		inf = d.RollNext("2d6").DM(ix).Sum()
	}
	infH := ehex.New(inf)
	//////////////////////////////////
	eff = d.FluxNext()
	if eff == 0 {
		eff = 1
	}
	sep := ""
	if eff > -1 {
		sep = "+"
	}
	//////////////////////////////////
	return fmt.Sprintf("(%v%v%v%v%v)", resH.String(), labH.String(), infH.String(), sep, eff)
}

func resValid(r ehex.DataRetriver) bool {
	switch {
	case r.Value() < 2:
		return false
	case r.Value() > 19:
		return false
	}
	return true
}

func reRollResourses(d *dice.Dicepool, tl ehex.DataRetriver, pbg string) ehex.DataRetriver {
	hex2 := strings.Split(pbg, "")
	b := ehex.New(hex2[1])
	g := ehex.New(hex2[2])
	res := d.RollNext("2d6").Sum()
	if tl.Value() >= 8 {
		res += b.Value() + g.Value()
	}
	return ehex.New(res)
}
