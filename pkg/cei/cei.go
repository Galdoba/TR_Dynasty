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

//Team - определяет способность абстрактной команды справиться с задачами
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

func NewTeam(teamtype string, baseIndex int) *Team {
	c := Team{}
	c.TeamType = teamtype
	c.CrewEfficencyIndex = baseIndex
	c.CEIModifier = make(map[string]int)
	c.Division = make(map[string]*Team)
	c.Morale = baseIndex
	c.AddEntry(fmt.Sprintf("%v created", c.TeamType))
	return &c
}

func (t *Team) SetCEI(cei int) {
	t.CrewEfficencyIndex = cei
}

func (t *Team) SetMorale(mor int) {
	t.Morale = mor
}

func (c *Team) AddDivision(division string) {
	c.Division[division] = NewTeam(division, c.CrewEfficencyIndex)
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

func (c *Team) sumMods() int {
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

func (t *Team) Resolve(descr ...string) int {
	r := dice.Roll2D()
	dm := t.TaskDM()
	if len(descr) > 0 {
		t.AddEntry(fmt.Sprintf("%v resolved: Roll=%v+(%v)", descr[0], r, dm))
	}
	return r + dm
}

func (t *Team) String() string {
	return fmt.Sprintf("%v (%v)", t.TeamType, t.ECEI())
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
	for _, d := range t.Division {
		d.PrintLog()
	}
}

/*
2 Something goes sufficiently badly wrong that the Incident becomes a Mishap.
3 An equipment breakdown causes difficulties for the detachment or crew. For example, the main sensor
processing centre suffers a software glitch and has to be rebooted whilst the ship is trying to scan local
traffic. The Travellers will have to figure a way around the problem, which cannot be fixed in time to
complete the task at hand.
4 A discipline problem occurs during the operation, disrupting the work of everyone involved. A heavy handed
solution might cause further resentment, but letting crewmembers get away with indiscipline will lead to
more problems in the future.
5 The operation is made more complex by changing circumstances. This will usually be something more
inconvenient than deadly, such as equipment malfunction, unexpected environmental changes, or some
other source of unnecessary delay and difficulty.
6 Some of the necessary equipment is unexpectedly missing, offline, or out of commission, requiring a
creative workaround or hurried fix.
7 A minor hiccup occurs, such as an infraction of regulations that requires a hearing and disciplinary action.
This might be embarrassing in front of outsiders and could have repercussions if the crew are discontent.
8 Something unusual is spotted; a sensor blip, an intriguing arrangement of planets and other bodies in the
system, strange atmospheric composition, or something equally puzzling but not immediately threatening.
9 A crewmember or detachment demonstrates unexpected resourcefulness and gets the job done better and
quicker than expected.
10 A typically uncooperative crewmember pitches in with a will and is helpful in completing the task. The
reason for this change of heart is not immediately obvious.
11 The task requires dusting off an unusual piece of equipment or trying out a technique not normally used.
Make a Difficult (10+) check using CEI and if successful, gain MOR +1.
12 Things seem to go pretty well, and in the middle of the task an Opportunity occurs


2 Structural damage is taken or a weakness is detected. The ship loses 2D% of its Hull points until
properly repaired.
3 The ship is involved in a minor collision with a small craft or object, or causes a similar incident to
happen to another craft.
4 A major system such as the spinal weapon or a drive develops a fault which makes it erratic. Impose
DM-1 on all checks involving that system until repaired.
5 One of the ship’s minor systems, such as a single small craft or a point-defence battery, suffers a
malfunction and is out of action until fully repaired.
6 A crewmember is seriously injured, requiring investigation.
7 A crewmember suffers a minor injury, which may well be his own fault.
8 A crewmember causes injury to someone, creating a possible Incident.
9 A crewmember manages to insult or offend someone.
10 One of the vehicles or small craft involved in the task suffers a serious malfunction, or a working party
has an accident and requires assistance.
11 The Travellers are given cause to suspect their plan for the current task is based on faulty data. This
could be serious, such as a failure to identify an atmospheric taint or a mis-estimate of the surface
gravity of a nearby world. The bad information may impose DM-2 on all checks connected with it, or
pose a more direct hazard.
12 A Crisis occurs. See page 27.


2 The Travellers gain knowledge of something very special in the local region. This might be a previously
unknown civilisation, wondrous phenomenon, or source of rare materials.
3 During the task the Travellers make a valuable find, such as a stash of components mislabelled and
forgotten in an obscure storeroom.
4 Some of the crew have found a way to fix a previously impossible problem. The Travellers gain a free repair
effort (see Repairs and Replacements on page 58) or similar advantage.
5 A piece of highly useful data is obtained. The Travellers may ‘cash in’ this find at a later date, in return for
information from the referee or a bonus to resolve tasks during a Mission.
6 Routine intelligence-collation produces a lead on something worthy of investigation, such as a planetoid belt
with rich resources or a star with unusual characteristics.
7 The current task turns out to be much easier than expected and is completed in record time or without
difficulties.
8 Inaccuracies in the supply logs turn out to be in favour of the Travellers for once. There are more of what is
needed than expected, allowing the current task to be completed without expending resources or supplies
expended this week to be effectively replaced.
9 Someone finds additional uses for standard equipment or a technique which improves efficiency. DM+1
applies to all checks for the remainder of this Mission.
10 An exceptionally good idea is presented to the Travellers, greatly simplifying current tasks. DM+1 applies to
all checks to resolve tasks throughout the current Reach.
11 A piece of equipment aboard the Travellers’ ship turns out to be a non-standard variant with additional
functions. These were not a great success and not integrated into future designs, which means few people
know about them. On one occasion when the Travellers face a technical problem (for example, inability to
resolve critical sensor data) this capability provides a solution.
12 A non-Traveller crewmember turns out to have had a very respectable academic career or interest in an
obscure subject that has now become useful. Their long-disused knowledge of obscure alien art, dead
languages, or rare jumpspace phenomena can be defined at any time the Travellers want, providing a
solution to an otherwise difficult situation.
*/
