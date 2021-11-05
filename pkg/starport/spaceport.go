package starport

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/T5/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type Starport struct {
	name             string
	dicepool         *dice.Dicepool
	dockingFee       []int //Land/Daily
	berthing         []int //Smallc Craft/Spaceships/Capital Ships
	fuel             []int //Unrefined/Refined
	warehousing      int   //
	hazmatCost       int   //per week per ton (twise of storage)
	storageCost      int   //per week per ton
	repairFacilities []int //Smallc Craft/Spaceships/Capital Ships
	upgrades         []string
	waitingTimes     []int //berthing/fuel/warehouse/hazmat/repair
	highport         bool
	class            string
}

type WorldDataDrawer interface {
	Name() string
	UWP() string
	ImportanceVal() int
}

func Assemble(wd WorldDataDrawer) (*Starport, error) {
	sp := Starport{}
	sp.dicepool = dice.New().SetSeed(wd.UWP() + wd.UWP())
	sp.name = " " + wd.Name()
	sp.class = spClass(wd)
	sp.calculateDockingFee(wd.ImportanceVal())
	sp.calculateBerthing(wd.ImportanceVal())
	sp.calculateFuelCost(wd)
	sp.warehousingCost()
	sp.calculateShipyard()
	// sp.berthing = append(sp.berthing, 1000)
	// sp.berthing = append(sp.berthing, 2000)
	// sp.berthing = append(sp.berthing, 3999)

	err := fmt.Errorf("Assembly not inplemented")
	return &sp, err
}

func (sp *Starport) calculateDockingFee(imp int) {
	dockingDice := sp.dicepool.RollNext("1d6").Sum() + imp
	dockFee := 0
	dockFeeDaily := 0
	switch sp.class {
	default:
		return
	case "A":
		dockFee = 1000 * dockingDice
		dockFeeDaily = 500 + imp*50
	case "B":
		dockFee = 500 * dockingDice
		dockFeeDaily = 200 + imp*20
	case "C":
		dockFee = 100 * dockingDice
		dockFeeDaily = 100 + imp*10
	case "D":
		dockFee = 10 * dockingDice
		dockFeeDaily = 10 + imp*1
	}
	if dockFee < 0 {
		dockFee = 0
	}
	if dockFeeDaily < 0 {
		dockFeeDaily = 0
	}
	sp.dockingFee = append(sp.dockingFee, dockFee)
	sp.dockingFee = append(sp.dockingFee, dockFeeDaily)
}

func (sp *Starport) calculateFuelCost(wd WorldDataDrawer) {
	uwp := wd.UWP()
	h := ehex.New().Set(string([]byte(uwp)[3])).Value()
	//g := ehex.New().Set(string([]byte(uwp)[5])).Value()
	//l := ehex.New().Set(string([]byte(uwp)[6])).Value()
	im := wd.ImportanceVal()
	f := 100 + ((6 - h) * 5) + (im * 5) // + (g * 2) + (9 - l*2)
	sp.fuel = append(sp.fuel, f)
	switch sp.class {
	case "A", "B", "C":
		sp.fuel = append(sp.fuel, f*5)
	default:
		sp.fuel = append(sp.fuel, -1)
	}

}

func (sp *Starport) calculateBerthing(imp int) {
	size := 0
	balance := []int{0, 0, 0}
	switch sp.class {
	default:
		return
	case "A":
		size = 300000
		balance = []int{35, 35, 30}
	case "B":
		size = 150000
		balance = []int{30, 35, 35}
	case "C":
		size = 30000
		balance = []int{20, 40, 40}
	case "D":
		size = 3000
		balance = []int{30, 70, 0}
	case "E":
		size = 1000
		balance = []int{40, 60, 0}
	case "X":
		size = 0
		balance = []int{35, 35, 30}
	}
	size = size + (sp.dicepool.FluxNext() * 10 * size / 100) + (size * imp * 10 / 100)
	for _, val := range balance {
		sp.berthing = append(sp.berthing, size*val/100)
	}
}

func (sp *Starport) calculateShipyard() {
	switch sp.class {
	default:
		sp.repairFacilities = []int{0, 0, 0}
	case "E":
		sp.repairFacilities = []int{1, 0, 0}
	case "D":
		sp.repairFacilities = []int{2, 1, 0}
	case "C":
		sp.repairFacilities = []int{3, 2, 1}
	case "B":
		sp.repairFacilities = []int{3, 3, 2}
	case "A":
		sp.repairFacilities = []int{3, 3, 3}
	}
}

func (sp *Starport) String() string {
	text := sp.name
	text += "\n Starport Class: " + sp.class
	text += "\n DOCKING FEE   : " + fmt.Sprintf("%v/%v Cr", sp.dockingFee[0], sp.dockingFee[1])
	text += "\n+---------------------------+---------------------------------+"
	text += "\n| BERTHING                  | WAITING TIME                    |"
	text += "\n| Small Craft  : " + formatFee(sp.berthing[0]) + " | " + waitingTimeBerthing(sp.berthing[0]) + "  |"
	text += "\n| Starships    : " + formatFee(sp.berthing[1]) + " | " + waitingTimeBerthing(sp.berthing[1]) + "  |"
	text += "\n| Capital Ships: " + formatFee(sp.berthing[2]) + " | " + waitingTimeBerthing(sp.berthing[2]) + "  |"
	text += "\n+---------------------------+---------------------------------+"
	text += "\n| FUEL COST    : " + sp.fuelSTR() + " | WAITING TIME                    |"
	text += "\n| Small Craft               | " + waitingTimeFuel(sp.berthing[0]) + "  |"
	text += "\n| Starships                 | " + waitingTimeFuel(sp.berthing[1]) + "  |"
	text += "\n| Capital Ships             | " + waitingTimeFuel(sp.berthing[2]) + "  |"
	text += "\n+---------------------------+---------------------------------+"
	text += "\n| WAREHOUSING  : " + sp.storageSTR() + " | WAITING TIME                    |"
	text += "\n| Small Craft               | " + waitingStorage(sp.berthing[0]) + "  |"
	text += "\n| Starships                 | " + waitingStorage(sp.berthing[1]) + "  |"
	text += "\n| Capital Ships             | " + waitingStorage(sp.berthing[2]) + "  |"
	text += "\n+---------------------------+---------------------------------+"
	text += "\n| Shipyard                                                    |"
	text += "\n| Small Craft  : " + shipyardServices(sp.repairFacilities[0]) + "  |"
	text += "\n| Starships    : " + shipyardServices(sp.repairFacilities[1]) + "  |"
	text += "\n| Capital Ships: " + shipyardServices(sp.repairFacilities[2]) + "  |"
	text += "\n+---------------------------+---------------------------------+"
	return text
	//Hull, Systems, Refit
	//Hull, Systems
	//Hull
	//N/A
}

func shipyardServices(i int) string {
	serv := ""
	switch i {
	default:
		serv = "N/A"
	case 3:
		serv = "Hull, Systems, Refit"
	case 2:
		serv = "Hull, Systems"
	case 1:
		serv = "Hull"
	}
	for len(serv) < 43 {
		serv += " "
	}
	return serv
}

func formatFee(fee int) string {
	f := strconv.Itoa(fee) + " "
	for len(f) < 10 {
		f += " "
	}
	return f
}

func (sp *Starport) warehousingCost() {
	switch sp.class {
	case "A":
		sp.storageCost = 500
	case "B":
		sp.storageCost = 400
	case "C":
		sp.storageCost = 300
	case "D":
		sp.storageCost = 200
	case "E":
		sp.storageCost = 100
	}
	sp.hazmatCost = sp.storageCost * 2
}

func (sp *Starport) fuelSTR() string {
	s := ""
	switch sp.fuel[1] {
	default:
		s = fmt.Sprintf("%v/%v", sp.fuel[0], sp.fuel[1])
	case -1:
		s = fmt.Sprintf("%v/NA", sp.fuel[0])
	}

	for len(s) < 10 {
		s += " "
	}
	return s
}

func (sp *Starport) storageSTR() string {
	s := fmt.Sprintf("%v/D-ton", sp.storageCost)

	for len(s) < 10 {
		s += " "
	}
	return s
}

func spClass(wd WorldDataDrawer) string {
	cls := ""
	uwp := wd.UWP()
	cls = string([]byte(uwp)[0])
	return cls
}

func highportPresence(cls string) bool {
	switch cls {
	default:
		return false
	case "A":
		return true
	}
}

func waitingTimeBerthing(size int) string {
	r := dice.Roll1D()
	switch {
	case size > 100000:
		r = r - 5
		return waitingTimeString(r)
	case size > 50000:
		r = r - 4
		return waitingTimeString(r)
	case size > 10000:
		r = r - 3
		return waitingTimeString(r)
	case size > 3000:
		r = r - 2
		return waitingTimeString(r)
	case size > 1000:
		r = r - 2
		return waitingTimeString(r)
	case size > 0:
		r = r - 1
		return waitingTimeString(r)
	}
	return "                              "
}

func waitingTimeFuel(size int) string {
	r := dice.Roll1D()
	switch {
	case size > 100000:
		r = r - 3
		return waitingTimeString(r)
	case size > 50000:
		r = r - 3
		return waitingTimeString(r)
	case size > 10000:
		r = r - 2
		return waitingTimeString(r)
	case size > 3000:
		r = r - 2
		return waitingTimeString(r)
	case size > 1000:
		r = r - 1
		return waitingTimeString(r)
	case size > 0:
		r = r - 0
		return waitingTimeString(r)
	}
	return "                              "
}

func waitingStorage(size int) string {
	r := dice.Roll1D()
	size = size / 2
	switch {
	case size > 100000:
		r = r - 3
		return waitingTimeString(r)
	case size > 50000:
		r = r - 2
		return waitingTimeString(r)
	case size > 10000:
		r = r - 2
		return waitingTimeString(r)
	case size > 3000:
		r = r - 1
		return waitingTimeString(r)
	case size > 1000:
		r = r - 0
		return waitingTimeString(r)
	case size > 0:
		r = r + 1
		return waitingTimeString(r)
	}
	return "                              "
}

func waitingTimeString(i int) string {
	if i < 0 {
		i = 0
	}
	if i > 7 {
		i = 7
	}
	r := dice.Roll1D()
	switch i {
	case 0:
		return "Service available immediately "
	case 1:
		st := fmt.Sprintf("%v minute", r)
		if r > 1 {
			st += "s"
		}
		for len(st) < 30 {
			st += " "
		}
		return st
	case 2:
		st := fmt.Sprintf("%v minutes", r*10)
		for len(st) < 30 {
			st += " "
		}
		return st
	case 3:
		st := fmt.Sprintf("1 hour")
		for len(st) < 30 {
			st += " "
		}
		return st
	case 4:
		st := fmt.Sprintf("%v hour", r)
		if r > 1 {
			st += "s"
		}
		for len(st) < 30 {
			st += " "
		}
		return st
	case 5:
		r = dice.Roll2D()
		st := fmt.Sprintf("%v hours", r)
		for len(st) < 30 {
			st += " "
		}
		return st
	case 6:
		st := fmt.Sprintf("1 day")
		for len(st) < 30 {
			st += " "
		}
		return st
	case 7:
		st := fmt.Sprintf("%v day", r)
		if r > 1 {
			st += "s"
		}
		for len(st) < 30 {
			st += " "
		}
		return st
	}
	return "unavailable"
}

/*
Space Station Characteristics Profile
========================
 [Station Name]
 DOCKING FEE: [xxxxx Cr]
+---------------------------+-------------------------------+
| BERTHING                  | WAITING TIME                  |
| Small Craft  : [xxxxx Cr] | Service available immediately |
| Starships    : [xxxxx Cr] |                               |
| Capital Ships: [xxxxx Cr] |                               |
+---------------------------+-------------------------------+


*/

// A = 300k
//65 + 30
//40+100
//0 + 100
// B = 150k
//25+20
//30+20
//50
// C = 30k
//4+1.5 5
//7+3
//10
// D = 3k
// E = 1k
//size + importance*10%*size + flux*size*1%
