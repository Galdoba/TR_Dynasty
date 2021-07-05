package actions

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
	"github.com/urfave/cli"
)

//SpaceEncounter - основная функция для SpaceEncounter
func SpaceEncounter(c *cli.Context) error {
	uwp, err := uwp.FromString(c.String("uwp"))
	//fmt.Println(err)
	if err != nil {
		return fmt.Errorf("action: SpaceEncounter:\n  %v", err.Error())
	}
	dm := 0
	switch {
	case isUntravelled(uwp):
		dm = -4
	case isWildSpace(uwp):
		dm = -1
	case hasHighport(uwp):
		dm = 3
	case hasHighTraffic(uwp):
		dm = 2
	case isSettled(uwp):
		dm = 1
	}
	fmt.Println("Dm =", dm)
	r1 := dice.Roll("1d6").DM(dm).Sum()
	r2 := dice.Roll("1d6").Sum()
	if r1 < 0 {
		r1 = 0
	}
	r66 := strconv.Itoa(r1) + strconv.Itoa(r2)
	fmt.Print("Roll ", r66)
	return nil
}

func hasHighport(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Starport().String() == "A" && uwp.Pops().Value() >= 7:
		return true
	case uwp.Starport().String() == "B" && uwp.Pops().Value() >= 8:
		return true
	case uwp.Starport().String() == "C" && uwp.Pops().Value() >= 9:
		return true
	}
}

func hasHighTraffic(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Pops().Value() >= 9:
		return true
	}
}

func isSettled(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Pops().Value() > 6:
		return true
	}
}

func isWildSpace(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Govr().Value()+uwp.Laws().Value() >= 20:
		return true
	case uwp.Pops().Value() <= 3:
		return true
	}
}

func isUntravelled(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Starport().String() == "X":
		return true
	case uwp.Pops().Value() == 0 || uwp.Govr().Value() == 0 || uwp.Laws().Value() == 0:
		return true
	}
}
