package law

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/worldBuilder"
)

//Lawer - тот кто выполняет законы
type Lawer interface {
	Interact(int) bool
}

type lawAgents struct {
	world    worldBuilder.World
	response int
}

func NewLawer(world *worldBuilder.World) *lawAgents {
	lc := &lawAgents{}
	lc.world = *world
	return lc
}

func (lc *lawAgents) Interact(dm int) bool {
	//defer lc.Report()
	r := TrvCore.Roll2D(dm)
	if r > lc.world.Stats()["Laws"] {
		return false
	}
	lc.response++
	return true
}

func (lc *lawAgents) Report() {
	fmt.Println("Report:", lc.response)
}
