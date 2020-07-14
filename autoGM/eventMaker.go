package autoGM

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/worldBuilder"
	"github.com/Galdoba/utils"
)

func test() {
	fmt.Println("Start test")
	utils.RandomSeed()
	mw := WorldMaker("MW")
	as := WorldMaker("AS")
	tst := WorldMaker("test")
	if world, ok := mw.(*worldBuilder.World); ok {
		world.SecondSurvey()
	}
	fmt.Println(mw)
	fmt.Println(as)
	fmt.Println(tst)
	test := worldBuilder.NewWorld("Test World", "RANDOM")
	fmt.Println(test)
	fmt.Println("End test")
}

type PlanetBody interface {
	Orbit() int
}

type asteroid struct {
	name  string
	orbit int
}

func (as *asteroid) Orbit() int {
	return as.orbit
}

func WorldMaker(worldType string) PlanetBody {
	switch worldType {
	default:
		return nil
	case "MW":
		return worldBuilder.NewWorld("Astar", "A555679-B")
	case "AS":
		return &asteroid{name: "Asteroid", orbit: 2}
	}

}
