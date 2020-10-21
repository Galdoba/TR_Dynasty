package autoGM

import (
	"errors"
	"fmt"

	law "github.com/Galdoba/TR_Dynasty/Law"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/devtools/cli/user"
)

func EncounterMgT2Core() {
	err := errors.New("No Input")

	var travellDays int
	for err != nil {
		fmt.Print("Enter days: ")
		travellDays, err = user.InputInt()
	}
	event := "no encounter"
	for i := 0; i < travellDays; i++ {
		eventRoll := dice.Roll("1d6").Sum()
		event = "no encounter"
		if eventRoll == 6 {
			event = "encounter " + dice.RollD66()
			fmt.Print("Day ", i+1, ": ", event+" happen"+"\n")
			break
		}
		fmt.Print("Day ", i+1, ": ", event+"\n")
	}
	fmt.Print(event + "\n")
	w := pickWorld()
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

func pickWorld() wrld.World {
	err := errors.New("No Input")
	for err != nil {
		input := ""
		fmt.Print("Enter world's Name, Hex or UWP: ")
		input, err = user.InputStr()
		otuData, errI := otu.GetDataOn(input)
		if errI != nil {
			fmt.Print("WARNING: " + err.Error() + "\n")
			continue
		}
		w, errO := wrld.FromOTUdata(otuData)
		if errO != nil {
			fmt.Print(err.Error() + "\n")
			continue
		}
		//output := "Data retrived: " + w.Name() + " (" + w.UWP() + ")\n"
		//printSlow(output)
		return w

	}
	fmt.Println("This must not happen!")
	return wrld.World{}
}

func starport(w wrld.World) string {
	prf, _ := profile.NewUWP(w.UWP())
	return prf.Starport()
}

func lawLevel(w wrld.World) int {
	prf, _ := profile.NewUWP(w.UWP())
	return TrvCore.EhexToDigit(prf.Laws())
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
