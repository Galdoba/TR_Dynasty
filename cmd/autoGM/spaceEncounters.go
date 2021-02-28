package autoGM

import (
	"fmt"

	law "github.com/Galdoba/TR_Dynasty/Law"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

//EncounterMgT2Core -
func EncounterMgT2Core() {
	event := "no encounter"
	i := 0
	for {
		i++
		eventRoll := dice.Roll("1d6").Sum()
		event = "no encounter"
		if eventRoll == 6 {
			event = "encounter " + dice.RollD66()
			fmt.Print("Day ", i, ": ", event+" happen"+"\n")
			break
		}

	}
	fmt.Print(event + "\n")
	w := wrld.PickWorld()
	m1, m2 := encounterMods(w)
	fmt.Print(m1, m2)

}

func encounterMods(w wrld.World) (int, int) {
	m1, m2 := 0, 0
	if isHighTraffic(w) {
		fmt.Print("High Traffic\n")
		m1 = m1 + 2
	}
	if isBackwater(w) {
		fmt.Print("Backwater\n")
		m1--
	}
	if isDangerous(w) {
		fmt.Print("Dangerous\n")
		m1--
	}
	if isSettled(w) {
		fmt.Print("Settled\n")
		m1++
	}
	//w := pickWorld()

	return m1, m2
}

///////////////////////////////////
//helpers

func starport(w wrld.World) string {
	prf := uwp.From(&w)
	return prf.Starport().String()
}

func lawLevel(w wrld.World) int {
	prf := uwp.From(&w)
	return prf.Laws().Value()
}

func isHighTraffic(w wrld.World) bool {
	sp := starport(w)
	if sp != "A" && sp != "B" {
		return false
	}
	fmt.Println(w.TradeCodes())
	for _, val := range w.TradeCodes() {
		switch val {
		case constant.TradeCodeHighTech, constant.TradeCodeHighPopulation, constant.TradeCodeIndustrial, constant.TradeCodeAgricultural, constant.TradeCodeRich:
			return true
		}
	}
	return false
}

func isBackwater(w wrld.World) bool {
	sp := starport(w)
	if sp != "X" && sp != "E" {
		return false
	}
	return true
}

func isDangerous(w wrld.World) bool {
	tz := w.TravelZone()
	if tz == "A" || tz == "R" {
		return true
	}
	if lawLevel(w) <= 3 {
		return true
	}
	sec := law.NewSecurity(&w)
	if sec.OrbitalPresence() == 0 {
		return true
	}
	return false
}

func isSettled(w wrld.World) bool {
	for _, val := range w.TradeCodes() {
		if val == constant.TradeCodeLowPopulation {
			return false
		}
	}
	return true
}
