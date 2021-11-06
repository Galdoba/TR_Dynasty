package traffic

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile"
)

func PassengerTrafficCard(pt *passengerTrafficMgT2) {
	fmt.Printf("Passenger Traffic from %v to %v:\n", pt.source, pt.target)
	fmt.Printf("Low passengers   : %v\n", pt.lp)
	fmt.Printf("Basic passengers : %v\n", pt.bp)
	fmt.Printf("Middle passengers: %v\n", pt.mp)
	fmt.Printf("High passengers  : %v\n", pt.hp)
}

type passengerTrafficMgT2 struct {
	source          string
	target          string
	effect          int //эффект проверки Broker, Carouse или Streetwise
	stewardMod      int
	sPopMod         int
	sSPortMod       int
	sTravellZoneMod int
	tPopMod         int
	tSPortMod       int
	tTravellZoneMod int
	distanceMod     int
	lp              int
	bp              int
	mp              int
	hp              int
}

func NewPassengerTrafficMgT2(sw, tw *world) (*passengerTrafficMgT2, error) {
	err := fmt.Errorf("Effect&Steward mods not Implemented")
	pt := passengerTrafficMgT2{}
	pt.source = sw.name
	pt.target = tw.name
	swUwp := profile.NewUWP(sw.uwp)
	twUwp := profile.NewUWP(tw.uwp)
	dist := Astrogation.Distance(coordinates(sw), coordinates(tw)) * -1
	pt.distanceMod = dist
	if pt.distanceMod > 0 {
		pt.distanceMod = 0
	}
	switch {
	case swUwp.Population().Value() == 0, swUwp.Population().Value() == 1:
		pt.sPopMod = -4
	case swUwp.Population().Value() == 6, swUwp.Population().Value() == 7:
		pt.sPopMod = 1
	case swUwp.Population().Value() > 7:
		pt.sPopMod = 3
	case swUwp.Starport().Code() == "A":
		pt.sSPortMod = 2
	case swUwp.Starport().Code() == "B":
		pt.sSPortMod = 1
	case swUwp.Starport().Code() == "E":
		pt.sSPortMod = -1
	case swUwp.Starport().Code() == "X":
		pt.sSPortMod = -3
	case sw.travelZone == "A":
		pt.sSPortMod = 1
	case sw.travelZone == "R":
		pt.sSPortMod = -4
		//////////////////////////////////////////////////////
	case twUwp.Population().Value() == 0, twUwp.Population().Value() == 1:
		pt.tPopMod = -4
	case twUwp.Population().Value() == 6, twUwp.Population().Value() == 7:
		pt.tPopMod = 1
	case twUwp.Population().Value() > 7:
		pt.tPopMod = 3
	case twUwp.Starport().Code() == "A":
		pt.tSPortMod = 2
	case twUwp.Starport().Code() == "B":
		pt.tSPortMod = 1
	case twUwp.Starport().Code() == "E":
		pt.tSPortMod = -1
	case twUwp.Starport().Code() == "X":
		pt.tSPortMod = -3
	case tw.travelZone == "A":
		pt.tSPortMod = 1
	case tw.travelZone == "R":
		pt.tSPortMod = -4
	}
	pt.rollPassengers()
	return &pt, err
}

func (pt *passengerTrafficMgT2) constMods() int {
	return pt.sPopMod + pt.sSPortMod + pt.sTravellZoneMod + pt.tPopMod + pt.tSPortMod + pt.tTravellZoneMod + pt.effect + pt.stewardMod + pt.distanceMod
}

func (pt *passengerTrafficMgT2) setCheckEffect(e int) {
	pt.effect = e
}

func (pt *passengerTrafficMgT2) setStewardEffect(e int) {
	pt.stewardMod = e
}

func (pt *passengerTrafficMgT2) rollPassengers() {
	dp := dice.New()
	mods := pt.constMods()
	pt.lp = dice.Roll2D(mods + 1)
	pt.bp = dice.Roll2D(mods)
	pt.mp = dice.Roll2D(mods)
	pt.hp = dice.Roll2D(mods - 4)
	pt.lp = dp.RollNext(fmt.Sprint(PassengerTrafficTableMgT2(pt.lp)) + "d6").Sum()
	pt.bp = dp.RollNext(fmt.Sprint(PassengerTrafficTableMgT2(pt.bp)) + "d6").Sum()
	pt.mp = dp.RollNext(fmt.Sprint(PassengerTrafficTableMgT2(pt.mp)) + "d6").Sum()
	pt.hp = dp.RollNext(fmt.Sprint(PassengerTrafficTableMgT2(pt.hp)) + "d6").Sum()
}

func PassengerTrafficTableMgT2(i int) int {
	switch i {
	case 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9, 10:
		return 3
	case 11, 12, 13:
		return 4
	case 14, 15:
		return 5
	case 16:
		return 6
	case 17:
		return 7
	case 18:
		return 8
	case 19:
		return 9
	}
	if i <= 1 {
		return 0
	}
	if i >= 20 {
		return 10
	}
	return -1
}
