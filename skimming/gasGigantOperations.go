package skimming

import (
	"errors"

	"github.com/Galdoba/TR_Dynasty/dice"
)

type layer struct {
	depth        int    //номер слоя
	name         string //название слоя
	skimmingRate int    //эффективность забора топлива в тысячных
	pilotingDM   int    //модификатор сложности проверок пилота
	damageDice   int    //урон в дайсах за каждый раунд
	orbitDecay   string //время до падения в следующий слой при нерабочих двигателях
}

func newLayer(i int) (layer, error) {
	l := layer{}

	switch i {
	default:
		return l, errors.New("Layer depth value unacceptable")
	case 0:
		l.depth = 0
		l.name = "Space"
		l.skimmingRate = 0
		l.pilotingDM = 0
		l.damageDice = 0
		l.orbitDecay = "Never"

	case 1:
		l.depth = 1
		l.name = "Wisp"
		l.skimmingRate = 0
		l.pilotingDM = 0
		l.damageDice = 0
		l.orbitDecay = dice.Roll("4d6").SumStr() + " days"
	case 2:
		l.depth = 1
		l.name = "Extreme Shallow"
		l.skimmingRate = 100
		l.pilotingDM = 0
		l.damageDice = 0
		l.orbitDecay = dice.Roll("4d6").SumStr() + " days"
	case 3:
		l.depth = 1
		l.name = "Shallow"
		l.skimmingRate = 500
		l.pilotingDM = -1
		l.damageDice = 0
		l.orbitDecay = dice.Roll("4d6").SumStr() + " days"

	}
}
