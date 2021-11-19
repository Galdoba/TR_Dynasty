package cei

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	CREW                         = "Crew"
	DIVISION_FLIGHT              = "Flight"
	DIVISION_GUNNERY             = "Gunnery"
	DIVISION_ENGINEERING         = "Engineering"
	DIVISION_OTHER               = "Other"
	DIVISION_COMMAND             = "Command"
	DIVISION_OPERATIONS          = "Operations"
	DIVISION_MISSION             = "Mission"
	FATIGUE_Fresh                = 0
	FATIGUE_Fatigued             = 0
	FATIGUE_Highly_Fatigued      = 0
	FATIGUE_Dangerously_Fatigued = 0
	FATIGUE_Exhausted            = 0
	FATIGUE_Incapable            = 0
	//DETACHMENT_GUNNERY     = "Gunnery"
)

//Crew - определяет способность команды справиться с задачами
type Team struct {
	TeamType           string
	CrewEfficencyIndex int //CEI
	CEIModifier        map[string]int
	Division           map[string]*Team
	Morale             int
	Fatigue            int
	Log                []string
}

func (c *Team) AddEntry(entry string) {
	c.Log = append(c.Log, entry)
}

func NewTeam(teamtype string, divisions ...string) *Team {
	c := Team{}
	c.TeamType = teamtype
	c.AddEntry(fmt.Sprintf("%v created", c.TeamType))
	c.CrewEfficencyIndex = -1
	c.CEIModifier = make(map[string]int)
	c.Division = make(map[string]*Team)
	for _, dei := range divisions {
		div := Team{TeamType: dei}
		div.CrewEfficencyIndex = -1
		div.CEIModifier = make(map[string]int)
		div.Division = make(map[string]*Team)
		c.Division[dei] = &div
		c.AddEntry(fmt.Sprintf("%v division added", dei))
	}

	return &c
}

func (c *Team) Assemble() {
	if c.CrewEfficencyIndex < 0 {
		c.CrewEfficencyIndex = 7
		c.Morale = 7
		c.Fatigue = 0
		c.AddEntry(fmt.Sprintf("%v assembled", c.TeamType))
	}
	for _, detachment := range c.Division {
		detachment.Assemble()
	}
}

func (c *Team) AddModifier(name string, effect int) {
	c.CEIModifier[name] = effect
	c.AddEntry(fmt.Sprintf("Modifier Added: '%v' (%v)", name, effect))
}

func (c *Team) RemoveModifier(name string) {
	for n, e := range c.CEIModifier {
		if n == name {
			delete(c.CEIModifier, name)
			c.AddEntry(fmt.Sprintf("Modifier Removed: '%v' (%v)", n, e))
		}
	}
}

func (c *Team) Report() {
	longestName := "CEIM"
	for _, div := range c.Division {
		if len(longestName) < len("DEI ("+div.TeamType+")") {
			longestName = "DEI (" + div.TeamType + ")"
		}
	}
	fmt.Printf("CEI | %v\n", c.CrewEfficencyIndex)
	fmt.Printf("CEIM | %v\n", c.ECEI())
	for k, v := range c.Division {
		fmt.Printf("DEI %v | %v\n", k, v.ECEI())
	}
	fmt.Printf("MOR | %v\n", c.Morale)
}

func (c *Team) SumMods() int {
	m := 0
	for _, mod := range c.CEIModifier {
		m += mod
	}
	return m
}

func (c *Team) ECEI() int {
	r := c.CrewEfficencyIndex
	for _, val := range c.CEIModifier {
		r += val
	}
	return r
}

func (c *Team) CEIMchanges(eventDescr string, leadershipEffect int) {
	c.AddEntry(fmt.Sprintf("CEIM Changes: %v with leadership check effect %v", eventDescr, leadershipEffect))
	r := dice.Roll2D() + leadershipEffect
	c.AddEntry(fmt.Sprintf("Roll 2D: %v", r-leadershipEffect))
	mChange := 0

	switch r {
	case 1, 2:
		mChange = dice.Roll1D()
		c.Morale = c.Morale - mChange
		c.AddModifier(eventDescr, -2)
		c.AddEntry(fmt.Sprintf("MOR - %v, CEIM - 2", mChange))
	case 3, 4:
		mChange = dice.RollD3()
		c.Morale = c.Morale - mChange
		c.AddModifier(eventDescr, -1)
		c.AddEntry(fmt.Sprintf("MOR - %v, CEIM - 1", mChange))
	case 5, 6, 7, 8:
		c.AddEntry(fmt.Sprintf("No change"))
	case 9, 10, 11:
		c.AddEntry(fmt.Sprintf("The %v gains confidence. MOR + 1", c.TeamType))
	default:
		if r <= 0 {
			mChange = dice.Roll1D(3)
			c.Morale = c.Morale - mChange
			c.AddModifier(eventDescr, -3)
			c.AddEntry(fmt.Sprintf("Morale collapses (MOR - %v) and the crew is near mutiny. CEIM - 3", mChange))
			break
		}
		mChange = dice.RollD3()
		c.AddModifier(eventDescr, 1)
		c.AddEntry(fmt.Sprintf("Efficiency and morale increse. CEIM + 1, MOR + %v", mChange))
	}
	c.AddEntry(fmt.Sprintf("New morale is now MOR = %v", c.Morale))
	c.moraleStatus()
}

func (c *Team) moraleStatus() {
	if c.Morale < 0 {
		c.Morale = 0
		c.AddEntry(fmt.Sprintf("%v morale fixed to 0", c.TeamType))
	}
}

//TaskDM - возвращает модификатор при тесте выполнения заданий
func (c *Team) TaskDM() int {
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
