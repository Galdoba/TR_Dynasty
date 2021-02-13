package skimming

import (
	"errors"
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func Test() {
	SkimmFuel(40, 82, 1)
	fmt.Println("-----------------")
	SkimmFuel(40, 82, 2)
	fmt.Println("-----------------")
	SkimmFuel(40, 82, 3)
	fmt.Println("-----------------")
	SkimmFuel(40, 82, 4)
	fmt.Println("-----------------")
	SkimmFuel(40, 82, 5)
	fmt.Println("-----------------")
	SkimmFuel(40, 82, 6)
	fmt.Println("-----------------")
	SkimmFuel(40, 82, 7)
	fmt.Println("-----------------")
}

type layer struct {
	depth            int    //номер слоя
	name             string //название слоя
	skimmingRate     int    //эффективность забора топлива в тысячных
	skimmingTimeDice string
	pilotingDM       int    //модификатор сложности проверок пилота
	damageDice       int    //урон в дайсах за каждый раунд
	orbitDecay       string //время до падения в следующий слой при нерабочих двигателях
}

func SkimmFuel(shipTonnage int, targetVolume int, preferedlayer int) {
	l, err := newLayer(preferedlayer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fuelIncrement := shipTonnage / 100
	if fuelIncrement < 1 {
		fuelIncrement = 1
	}
	passesTotal := 0
	fuelSkimmed := 0
	totalMinutes := 0
	for fuelSkimmed < targetVolume {
		fuelSkimmed += fuelIncrement
		passesTotal++
		totalMinutes += dice.Roll(l.skimmingTimeDice).Sum()
	}
	fmt.Print("Using Ship with tonnage of ", shipTonnage, " tons in ", l.name, " targeting ", targetVolume, " tons of fuel: \n")
	fmt.Print("Will take about ", totalMinutes/60, " hours and ", totalMinutes%60, " minutes\n")
	fmt.Print("Piloting (", (-1 * (-8 + l.pilotingDM)), ") check MUST be taken\n")
}

func newLayer(i int) (layer, error) {
	l := layer{}
	switch i {
	default:
		return l, errors.New("Layer depth value unacceptable")
	case 0:
		l.name = "Space"
		l.pilotingDM = 0
		l.damageDice = 0
		l.orbitDecay = "Never"
	case 1:
		l.name = "Wisp"
		l.pilotingDM = 0
		l.damageDice = 0
		l.orbitDecay = dice.Roll("4d6").SumStr() + " days"
		l.skimmingTimeDice = "60d6"
	case 2:
		l.name = "Extreme Shallow"
		l.pilotingDM = 0
		l.damageDice = 0
		l.orbitDecay = dice.Roll("8d6").SumStr() + " hours"
		l.skimmingTimeDice = "20d6"
	case 3:
		l.name = "Shallow"
		l.pilotingDM = -1
		l.damageDice = 0
		l.orbitDecay = dice.Roll("2d6").SumStr() + " hours"
		l.skimmingTimeDice = "4d6"
	case 4:
		l.name = "Deep"
		l.pilotingDM = -2
		l.damageDice = 0
		l.orbitDecay = dice.Roll("2d6").SumStr() + " minutes"
		l.skimmingTimeDice = "2d6"
	case 5:
		l.name = "Extereme Deep"
		l.pilotingDM = -3
		l.damageDice = 2
		l.orbitDecay = dice.Roll("2d6").SumStr() + " minutes"
		l.skimmingTimeDice = "1d6"
	case 6:
		l.name = "Depths"
		l.pilotingDM = -4
		l.damageDice = 3
		l.orbitDecay = dice.Roll("2d6").SumStr() + " minutes"
		l.skimmingTimeDice = "1d2"
	case 7:
		l.name = "Abyssal Depths"
		l.pilotingDM = -5
		l.damageDice = 4
		l.orbitDecay = "None"
		l.skimmingTimeDice = "1d1"

	}
	l.depth = i
	return l, nil
}
