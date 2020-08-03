package essc

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
)

//StmKey -
type StmKey struct {
	starSystem         int
	specialSystem      int
	highlyUnususalBody int
	sequense           int
	planetaryBodies    int
	natureOfBodies     []int
	unusualBodies      []int
	anomalusBodies     []int
	orbiatlZones       []int
	mainworld          int
	planetUWPs         []string
	moons              []int
}

func (ss StmKey) Test() {

}

//Create -
func Create() StmKey {
	ss := StmKey{}
	ss.starSystem = 2 //dice.Roll2D()
	ss.specialSystem = dice.Roll2D()

	ss.highlyUnususalBody = dice.Roll2D()
	fmt.Println(starSystem(ss.starSystem))
	for ss.starSystem == 2 {
		ss.starSystem = dice.Roll2D()
		fmt.Println(specialSystem(ss.specialSystem))
		if ss.specialSystem == 2 {
			fmt.Println(highlyUnusualSystem(ss.highlyUnususalBody))

		}
		fmt.Println(starSystem(ss.starSystem))
	}

	ss.sequense = dice.Roll2D()
	ss.planetaryBodies = pbNumber(0)
	for i := 0; i < ss.planetaryBodies; i++ {
		ss.natureOfBodies = append(ss.natureOfBodies, dice.Roll2D())
		ss.unusualBodies = append(ss.unusualBodies, dice.Roll2D())
		ss.anomalusBodies = append(ss.anomalusBodies, dice.Roll2D())
	}
	//ss.natureOfBodies = dice.Roll2D()

	//ss.unusualBodies =
	//ss.anomalusBodies =
	//ss.orbiatlZones =
	ss.mainworld = dice.Roll2D()
	//ss.planetUWPs =
	//ss.moons =
	return ss
}

func pbNumber(dm int) int {
	r := dice.Roll2D() - dm
	switch r {
	default:
		return 0
	case 2:
		return 1
	case 3:
		return dice.Roll("d3").Sum()
	case 4, 5:
		return dice.Roll("d6").DM(1).Sum()
	case 6, 7, 8:
		return dice.Roll("2d6").Sum()
	case 9, 10:
		return dice.Roll("2d6").DM(3).Sum()

	case 11:
		return dice.Roll("3d6").Sum()

	case 12:
		return dice.Roll("4d6").Sum()
	}
}

func starSystem(i int) string {
	s := ""
	switch i {
	case 2:
		s = "Special (roll on the Special System table)"
	case 3, 4:
		s = "Trinary (close and distant companion)"
	case 5, 6:
		s = "Binary (close companion)"
	case 7, 8:
		s = "Solo star"
	case 9, 10:
		s = "Binary (distant companion)"
	case 11:
		s = "Trinary (distant companion with close companion of its own)"
	case 12:
		s = "Multiple star system (four or more stellar bodies)"
	}
	return s
}

func specialSystem(i int) string {
	s := ""
	switch i {
	case 2:
		s = "Highly unusual body (roll on the Highly Unusual Body table)"
	case 3:
		s = "Expanding pre-giant star"
	case 4, 5:
		s = "Brown dwarf system: Primary is a brown dwarf, as are all other stars"
	case 6, 7, 8:
		s = "Empty system: No planetary bodies present"
	case 9, 10:
		s = "White dwarf star: Inner system destroyed"
	case 11:
		s = "Giant star"
	case 12:
		s = "Unstable star: Prone to nova events"
	}
	return s
}

func highlyUnusualSystem(i int) string {
	s := ""
	switch i {
	case 2:
		s = "Black hole"
	case 3, 4:
		s = "Anomaly, e.g. white dwarf with life-bearing inner planets"
	case 5, 6, 7, 8, 9:
		s = "Nebula or protostar"
	case 10, 11:
		s = "Highly complex multiple star system"
	case 12:
		s = "Neutron star"
	}
	return s
}
