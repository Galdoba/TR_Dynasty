package cei

import (
	"fmt"
)

const (
	DIVISION_FLIGHT      = "Flight"
	DIVISION_GUNNERY     = "Gunnery"
	DIVISION_ENGINEERING = "Engineering"
	DIVISION_CREW        = "Crew"
	DIVISION_COMMAND     = "Command & Administration"
	DIVISION_OPERATIONS  = "Operations"
	DIVISION_MISSION     = "Mission"
	//DETACHMENT_GUNNERY     = "Gunnery"
)

//Crew - определяет способность команды справиться с задачами
type Crew struct {
	CrewEfficencyIndex int //CEI
	CEIModifier        []activeModifier
	Division           map[string]detachment
	Morale             int
	Fatigue            string
}

type activeModifier struct {
	descr string
	value int
}

type detachment struct {
	dType           string
	efficiencyIndex int
}

func NewCrew(baseIndex int, divisions ...detachment) (*Crew, error) {
	err := fmt.Errorf("Not implemented")
	c := Crew{}
	c.Division = make(map[string]detachment)
	for _, dei := range divisions {
		dei.efficiencyIndex = c.CrewEfficencyIndex
		c.Division[dei.dType] = dei
	}
	return &c, err
}

func (c *Crew) Report() {
	longestName := "CEIM"
	for _, div := range c.Division {
		if len(longestName) < len("DEI ("+div.dType+")") {
			longestName = "DEI (" + div.dType + ")"
		}
	}
	fmt.Printf("CEI			| %v\n", c.CrewEfficencyIndex)
	fmt.Printf("CEIM		| %v\n", c.CrewEfficencyIndex)

}

func NewDivision(name string) *detachment {
	return &detachment{name, -1}
}

func (c *Crew) ECEI() int {
	r := c.CrewEfficencyIndex
	for _, mod := range c.CEIModifier {
		r += mod.value
	}
	return r
}

//TaskDM - возвращает модификатор при тесте выполнения заданий
func (c *Crew) TaskDM() int {
	switch c.CrewEfficencyIndex {
	default:
		return -999
	case 0:
		return -6
	case 1:
		return -5
	case 2:
		return -4
	case 3:
		return -3
	case 4:
		return -2
	case 5:
		return -1
	case 6:
		return -1
	case 7:
		return 0
	case 8:
		return 0
	case 9:
		return 1
	case 10:
		return 1
	case 11:
		return 2
	case 12:
		return 3
	case 13:
		return 4
	case 14:
		return 5
	case 15:
		return 6
	}
}
