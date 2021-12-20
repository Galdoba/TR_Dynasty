package cei

import (
	"fmt"
	"sort"

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

//Team - определяет способность абстрактной команды справиться с задачами
type Team struct {
	TeamType           string
	CrewEfficencyIndex int //CEI
	CEIModifier        map[string]int
	Division           map[string]*Team
	Morale             int
	Fatigue            fatigue
	Log                []string
	MissionDay         int
}

func (c *Team) AddEntry(entry string) {
	c.Log = append(c.Log, fmt.Sprintf("Day %v: %v", c.MissionDay, entry))
}

func NewTeam(teamtype string, baseIndex int) *Team {
	c := Team{}
	c.TeamType = teamtype
	c.CrewEfficencyIndex = baseIndex
	c.CEIModifier = make(map[string]int)
	c.Division = make(map[string]*Team)
	c.Morale = baseIndex
	c.Fatigue = fatigue{
		index:        0,
		fatigueState: Fresh,
		status:       EVENT_FatigueStatus_Fresh,
	}

	c.CallEvent(c.Fatigue.status)
	c.AddEntry(fmt.Sprintf("%v created", c.TeamType))
	return &c
}

func (t *Team) Update() error {
	fatigueEvent := t.Fatigue.update()
	if fatigueEvent {
		t.CallEvent(t.Fatigue.status)
	}
	t.MissionDay++
	return nil
}

func (t *Team) SetCEI(cei int) {
	t.CrewEfficencyIndex = cei
}

func (t *Team) SetMorale(mor int) {
	t.Morale = mor
}

func (c *Team) AddDivision(division string, dei int) {
	c.Division[division] = NewTeam(division, dei)
	c.AddEntry(fmt.Sprintf("Detachment '%v' formed", division))
}

func (c *Team) RemoveDivision(division string) {
	delete(c.Division, division)
	c.AddEntry(fmt.Sprintf("Detachment '%v' removed", division))
}

//CallDivision - вызывает дивизию вместо всей команды
func (c *Team) CallDivision(division string) *Team {
	if team, ok := c.Division[division]; ok {
		return team
	}
	return nil
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
	fmt.Printf("%v SitRep\n", c.TeamType)
	fmt.Printf("Mission Day        : %v\n", c.MissionDay)
	fmt.Printf("Crew Fatigue State : %v\n", c.Fatigue.State())
	fmt.Printf("Crew Fatigue Index : %v\n", c.Fatigue.index)
	fmt.Printf("Days until next Fatigue Check %v\n", c.Fatigue.daysToCheck)
	longestName := "CEIM"
	for _, div := range c.Division {
		if len(longestName) < len("DEI ("+div.TeamType+")") {
			longestName = "DEI (" + div.TeamType + ")"
		}
	}

	fmt.Printf("%v | %v\n", addSpaces("CEI", longestName), c.CrewEfficencyIndex)
	fmt.Printf("%v | %v\n", addSpaces("CEIM", longestName), c.CEIM())
	for _, v := range divesionsList() {
		if c.CallDivision(v) == nil {
			continue
		}
		div := c.CallDivision(v)
		fmt.Printf("%v | %v\n", addSpaces("DEI ("+v+")", longestName), div.ECEI())
	}
	fmt.Printf("%v | %v\n", addSpaces("MOR", longestName), c.Morale)
	modif := []string{}
	for k, _ := range c.CEIModifier {
		modif = append(modif, k)
	}
	sort.Strings(modif)
	if len(c.CEIModifier) > 0 {
		fmt.Printf("Active Modifiers:\n")
	}
	for _, m := range modif {
		fmt.Printf(" %v = %v\n", m, c.CEIModifier[m])
	}
	fmt.Printf("---END REPORT-------------------------------------\n")
}

func addSpaces(str, mask string) string {
	for len(str) < len(mask) {
		str += " "
	}
	return str
}

func (c *Team) sumMods() int {
	m := 0
	for _, mod := range c.CEIModifier {
		m += mod
	}
	return m
}

func (c *Team) ECEI() int {
	r := c.CrewEfficencyIndex
	r += c.CEIM()
	if r < 0 {
		return 0
	}
	if r > 15 {
		return 15
	}
	return r
}

func (c *Team) CEIM() int {
	mod := 0
	for _, m := range c.CEIModifier {
		mod += m
	}
	return mod
}

func divesionsList() []string {
	return []string{
		DIVISION_FLIGHT,
		DIVISION_GUNNERY,
		DIVISION_ENGINEERING,
		DIVISION_OTHER,
		DIVISION_COMMAND,
		DIVISION_OPERATIONS,
		DIVISION_MISSION,
	}
}

func (c *Team) ResolutionDM() int {
	return c.TaskDM()
}

// func (c *Team) CEIMchanges(eventDescr string, leadershipEffect int) {
// 	c.AddEntry(fmt.Sprintf("CEIM Changes: %v with leadership check effect %v", eventDescr, leadershipEffect))
// 	r := dice.Roll2D() + leadershipEffect
// 	c.AddEntry(fmt.Sprintf("Roll 2D: %v", r-leadershipEffect))
// 	mChange := 0
// 	switch r {
// 	case 1, 2:
// 		mChange = dice.Roll1D()
// 		c.Morale = c.Morale - mChange
// 		c.AddModifier(eventDescr, -2)
// 		c.AddEntry(fmt.Sprintf("MOR - %v, CEIM - 2", mChange))
// 	case 3, 4:
// 		mChange = dice.RollD3()
// 		c.Morale = c.Morale - mChange
// 		c.AddModifier(eventDescr, -1)
// 		c.AddEntry(fmt.Sprintf("MOR - %v, CEIM - 1", mChange))
// 	case 5, 6, 7, 8:
// 		c.AddEntry(fmt.Sprintf("No change"))
// 	case 9, 10, 11:
// 		c.AddEntry(fmt.Sprintf("The %v gains confidence. MOR + 1", c.TeamType))
// 	default:
// 		if r <= 0 {
// 			mChange = dice.Roll1D(3)
// 			c.Morale = c.Morale - mChange
// 			c.AddModifier(eventDescr, -3)
// 			c.AddEntry(fmt.Sprintf("Morale collapses (MOR - %v) and the crew is near mutiny. CEIM - 3", mChange))
// 			break
// 		}
// 		mChange = dice.RollD3()
// 		c.AddModifier(eventDescr, 1)
// 		c.AddEntry(fmt.Sprintf("Efficiency and morale increse. CEIM + 1, MOR + %v", mChange))
// 	}
// 	c.AddEntry(fmt.Sprintf("New morale is now MOR = %v", c.Morale))
// 	c.moraleStatus()
// }

func (c *Team) moraleStatus() {
	if c.Morale < 0 {
		c.Morale = 0
		c.AddEntry(fmt.Sprintf("%v morale fixed to 0", c.TeamType))
	}
	if c.Morale > 15 {
		c.Morale = 15
		c.AddEntry(fmt.Sprintf("%v morale fixed to 15", c.TeamType))
	}
}

//TaskDM - возвращает модификатор при тесте выполнения заданий
func (c *Team) TaskDM() int {
	switch c.ECEI() {
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

func (t *Team) Resolve(externalMods ...int) int {
	r := dice.Roll2D()
	dm := t.TaskDM()
	mods := sumOf(externalMods)
	sum := r + dm + mods
	t.AddEntry(fmt.Sprintf("%v [Roll=%v; DM=%v; Extertal Mods=%v]", sum, r, dm, mods))
	return sum
}

func (t *Team) String() string {
	return fmt.Sprintf("%v (%v)", t.TeamType, t.ECEI())
}

func (t *Team) ChangeMoraleBy(i int) {
	t.Morale = t.Morale + i
	gl := "+"
	if i < 0 {
		gl = ""
	}
	t.AddEntry(fmt.Sprintf("MOR %v%v", gl, i))
	t.moraleStatus()
}

func (t *Team) MoraleCheckMinor() {
	r := t.ECEI() + dice.Roll2D()
	switch {
	case r < 8:
		t.Morale--
		t.AddEntry("Minor morale check failed: MOR -1")
		t.moraleStatus()
	case r >= 8:
		t.AddEntry("Minor morale check passed")
	}
}

func (t *Team) MoraleCheckMajor() {
	r := t.TaskDM() + dice.Roll2D()
	switch {
	case r < 8:
		m := dice.Roll1D()
		t.Morale = t.Morale - m
		t.moraleStatus()
		t.AddEntry(fmt.Sprintf("Major morale check failed: MOR -%v", m))
		if m >= 3 {
			t.AddEntry(fmt.Sprintf("Leadership crisis occurs"))
			t.AddModifier("Leadership crisis not resolved", -1)
		}
	case r >= 8:
		t.AddEntry("Major morale check passed")
	}
}

func (t *Team) PrintLog() {
	fmt.Println(t, "Log:")
	for _, line := range t.Log {
		fmt.Println(line)
	}

}

func (t *Team) Check(tn int) int {
	r := dice.Roll2D() + t.TaskDM()
	return r - tn
}
