package hyperjump

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
)

type hyperJump struct {
	effA          int
	effE          int
	badJumpA      bool
	badJumpE      bool
	precipitation bool
	outcome       string
}

//HyperJump - модуль гиперпрыжка в MGT2 Traveller Companion (p. 141-143)
type HyperJump interface {
	//Stringer
}

func Test() {
	hj := New(TrvCore.Flux(), TrvCore.Flux())
	fmt.Println(hj.outcome)
}

func New(effA, effE int) *hyperJump {
	hj := hyperJump{}
	hj.effA = effA
	hj.effE = effE
	fmt.Println(hj.effA, hj.effE)
	dVar := hj.distVariance()
	fmt.Println(dVar)
	fmt.Println(hj.badJumpA)
	tVar := hj.timeVariance()
	fmt.Println(tVar)
	fmt.Println(hj.badJumpE)

	if hj.badJumpA == true || hj.badJumpE == true {
		fmt.Print("WARNING!! It was Bad Jump.")
	}
	if hj.badJumpA == true && hj.badJumpE == true {
		fmt.Print(".. A VERY BAD Jump.")
	}
	fmt.Print("\n")

	return &hj
}

func (hj *hyperJump) distVariance() string {
	variance := 100
	r := dice.Roll2D(hj.effA)
	switch r {
	case 2:
		variance += 10 - dice.Roll3D()
		hj.badJumpA = true
	case 3:
		variance += 10 - dice.Roll2D()
		hj.badJumpA = true
	case 4:
		variance += 5 - dice.Roll1D()
		hj.badJumpA = true
	case 5:
		variance += (dice.Roll2D() * 10)
		hj.badJumpA = true
	case 6:
		variance += (dice.Roll2D() * 5)
	case 7:
		variance += dice.Roll4D()
	case 8:
		variance += dice.Roll3D()
	case 9:
		variance += dice.Roll2D()
	case 10:
		variance += dice.Roll1D()
	case 11:
		variance += (dice.Roll1D(1) / 2)
	default:
		variance = 100
		if r < 2 {
			variance += 10 - dice.Roll3D()
			hj.badJumpA = true
		}
	}
	if variance < 100 {

	}
	text := "Ship emerged at " + strconv.Itoa(variance) + " diameters from intended planet"
	return text

}

func (hj *hyperJump) timeVariance() string {
	time := 160
	variance := 0
	r := dice.Roll2D(hj.effE)
	switch r {
	case 2:
		variance = dice.Roll("16d6").Sum()
		hj.badJumpE = true
	case 3:
		variance = dice.Roll("10d6").Sum()
		hj.badJumpE = true
	case 4:
		variance = dice.Roll("8d6").Sum()
		hj.badJumpE = true
	case 5:
		variance = dice.Roll("6d6").Sum()
		hj.badJumpE = true
	case 6:
		variance = dice.Roll("5d6").Sum()
	case 7:
		variance = dice.Roll("4d6").Sum()
	case 8:
		variance = dice.Roll("3d6").Sum()
	case 9:
		variance = dice.Roll("2d6").Sum()
	case 10:
		variance = dice.Roll("1d6").Sum()
	case 11:
		variance = dice.Roll("1d6").DM(1).Sum() / 2
	default:
		if r < 2 {
			variance = dice.Roll("16d6").Sum()
			hj.badJumpE = true
		}
	}
	r2 := dice.Roll("1d2").Sum()
	if r2 == 2 {
		variance = variance * -1
	}
	time = time + variance
	text := "Time spent in a Jump is " + strconv.Itoa(time) + " hours."
	return text
}
