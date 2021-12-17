package cei

import "github.com/Galdoba/TR_Dynasty/pkg/dice"

const (
	EVENT_NONE                              = "No Event"
	EVENT_FatigueStatus_Fresh               = "Fresh"
	EVENT_FatigueStatus_Fatigued            = "Fatigued"
	EVENT_FatigueStatus_HighlyFatigued      = "Highly Fatigued"
	EVENT_FatigueStatus_DangerouslyFatigued = "Dangerously Fatigued"
	EVENT_FatigueStatus_Exhausted           = "Exhausted"
	EVENT_FatigueStatus_Incapable           = "Incapable"
)

func (t *Team) CallEvent(event string) {
	switch event {
	default:
		return
	case EVENT_FatigueStatus_Fresh:
		t.AddModifier("Fatigue State", 1)
	case EVENT_FatigueStatus_Fatigued:
		t.AddModifier("Fatigue State", 0)
	case EVENT_FatigueStatus_HighlyFatigued:
		t.AddModifier("Fatigue State", -1)
	case EVENT_FatigueStatus_DangerouslyFatigued:
		t.AddModifier("Fatigue State", -2)
	case EVENT_FatigueStatus_Exhausted:
		t.AddModifier("Fatigue State", -3)
	case EVENT_FatigueStatus_Incapable:
		t.AddModifier("Fatigue State", -4)
	}
}

func (t *Team) Incident() {
	r := dice.Roll2D()
	text := "TODO Incident"
	switch r {
	case 2:
		text = "Something goes sufficiently badly wrong: " + Mishap() // that the Incident becomes a Mishap."
	case 3:
		text = "An equipment breakdown causes difficulties for the detachment or crew. For example, the main sensor processing centre suffers a software glitch and has to be rebooted whilst the ship is trying to scan local traffic. The Travellers will have to figure a way around the problem, which cannot be fixed in time to complete the task at hand."
	case 4:
		text = "A discipline problem occurs during the operation, disrupting the work of everyone involved. A heavy handed solution might cause further resentment, but letting crewmembers get away with indiscipline will lead to more problems in the future."
	case 5:
		text = "The operation is made more complex by changing circumstances. This will usually be something more inconvenient than deadly, such as equipment malfunction, unexpected environmental changes, or some other source of unnecessary delay and difficulty."
	case 6:
		text = "Some of the necessary equipment is unexpectedly missing, offline, or out of commission, requiring a creative workaround or hurried fix."
	case 7:
		text = "A minor hiccup occurs, such as an infraction of regulations that requires a hearing and disciplinary action. This might be embarrassing in front of outsiders and could have repercussions if the crew are discontent."
	case 8:
		text = "Something unusual is spotted; a sensor blip, an intriguing arrangement of planets and other bodies in the system, strange atmospheric composition, or something equally puzzling but not immediately threatening."
	case 9:
		text = "A crewmember or detachment demonstrates unexpected resourcefulness and gets the job done better and quicker than expected."
	case 10:
		text = "A typically uncooperative crewmember pitches in with a will and is helpful in completing the task. The reason for this change of heart is not immediately obvious."
	case 11:
		text = "The task requires dusting off an unusual piece of equipment or trying out a technique not normally used. Make a Difficult (10+) check using CEI and if successful, gain MOR +1."
		if t.Check(10) >= 0 {
			t.ChangeMoraleBy(1)
		}
	case 12:
		text = "Things seem to go pretty well, and in the middle of the task: " + Opportunity() //an Opportunity occurs."
	}
	t.AddEntry(text)
}

func Opportunity() string {
	r := dice.Roll2D()
	text := "TODO Opportunity"
	switch r {
	case 2:
	}
	return text
}

func Mishap() string {
	r := dice.Roll2D()
	text := "TODO Mishap"
	switch r {
	case 2:
	}
	return text
}

/*
INCIDENT
2 Something goes sufficiently badly wrong that the Incident becomes a Mishap.
3 An equipment breakdown causes difficulties for the detachment or crew. For example, the main sensor processing centre suffers a software glitch and has to be rebooted whilst the ship is trying to scan local traffic. The Travellers will have to figure a way around the problem, which cannot be fixed in time to complete the task at hand.
4 A discipline problem occurs during the operation, disrupting the work of everyone involved. A heavy handed solution might cause further resentment, but letting crewmembers get away with indiscipline will lead to more problems in the future.
5 The operation is made more complex by changing circumstances. This will usually be something more inconvenient than deadly, such as equipment malfunction, unexpected environmental changes, or some other source of unnecessary delay and difficulty.
6 Some of the necessary equipment is unexpectedly missing, offline, or out of commission, requiring a creative workaround or hurried fix.
7 A minor hiccup occurs, such as an infraction of regulations that requires a hearing and disciplinary action. This might be embarrassing in front of outsiders and could have repercussions if the crew are discontent.
8 Something unusual is spotted; a sensor blip, an intriguing arrangement of planets and other bodies in the system, strange atmospheric composition, or something equally puzzling but not immediately threatening.
9 A crewmember or detachment demonstrates unexpected resourcefulness and gets the job done better and quicker than expected.
10 A typically uncooperative crewmember pitches in with a will and is helpful in completing the task. The reason for this change of heart is not immediately obvious.
11 The task requires dusting off an unusual piece of equipment or trying out a technique not normally used. Make a Difficult (10+) check using CEI and if successful, gain MOR +1.
12 Things seem to go pretty well, and in the middle of the task an Opportunity occurs

MISHAP
2 Structural damage is taken or a weakness is detected. The ship loses 2D% of its Hull points until properly repaired.
3 The ship is involved in a minor collision with a small craft or object, or causes a similar incident to happen to another craft.
4 A major system such as the spinal weapon or a drive develops a fault which makes it erratic. Impose DM-1 on all checks involving that system until repaired.
5 One of the ship’s minor systems, such as a single small craft or a point-defence battery, suffers a malfunction and is out of action until fully repaired.
6 A crewmember is seriously injured, requiring investigation.
7 A crewmember suffers a minor injury, which may well be his own fault.
8 A crewmember causes injury to someone, creating a possible Incident.
9 A crewmember manages to insult or offend someone.
10 One of the vehicles or small craft involved in the task suffers a serious malfunction, or a working party has an accident and requires assistance.
11 The Travellers are given cause to suspect their plan for the current task is based on faulty data. This could be serious, such as a failure to identify an atmospheric taint or a mis-estimate of the surface gravity of a nearby world. The bad information may impose DM-2 on all checks connected with it, or pose a more direct hazard.
12 A Crisis occurs. See page 27.

OPPORTUNITY
2 The Travellers gain knowledge of something very special in the local region. This might be a previously unknown civilisation, wondrous phenomenon, or source of rare materials.
3 During the task the Travellers make a valuable find, such as a stash of components mislabelled and forgotten in an obscure storeroom.
4 Some of the crew have found a way to fix a previously impossible problem. The Travellers gain a free repair effort (see Repairs and Replacements on page 58) or similar advantage.
5 A piece of highly useful data is obtained. The Travellers may ‘cash in’ this find at a later date, in return for information from the referee or a bonus to resolve tasks during a Mission.
6 Routine intelligence-collation produces a lead on something worthy of investigation, such as a planetoid belt with rich resources or a star with unusual characteristics.
7 The current task turns out to be much easier than expected and is completed in record time or without difficulties.
8 Inaccuracies in the supply logs turn out to be in favour of the Travellers for once. There are more of what is needed than expected, allowing the current task to be completed without expending resources or supplies expended this week to be effectively replaced.
9 Someone finds additional uses for standard equipment or a technique which improves efficiency. DM+1 applies to all checks for the remainder of this Mission.
10 An exceptionally good idea is presented to the Travellers, greatly simplifying current tasks. DM+1 applies to all checks to resolve tasks throughout the current Reach.
11 A piece of equipment aboard the Travellers’ ship turns out to be a non-standard variant with additional functions. These were not a great success and not integrated into future designs, which means few people know about them. On one occasion when the Travellers face a technical problem (for example, inability to resolve critical sensor data) this capability provides a solution.
12 A non-Traveller crewmember turns out to have had a very respectable academic career or interest in an obscure subject that has now become useful. Their long-disused knowledge of obscure alien art, dead languages, or rare jumpspace phenomena can be defined at any time the Travellers want, providing a solution to an otherwise difficult situation.
*/
