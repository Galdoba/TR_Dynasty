package cei

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type CEI struct {
	BaseIndex      int
	EffectiveIndex int
	Morale         int
}

func New(base int) *CEI {
	cei := CEI{}
	cei.BaseIndex = base
	cei.EffectiveIndex = base
	cei.Morale = base
	return &cei
}

func (cei *CEI) SkillsLevels() (primarySkills []int, secondarySkills []int) {
	switch cei.BaseIndex {
	default:
	case 0:
		primarySkills = append(primarySkills, 0)
		secondarySkills = append(secondarySkills, 0)
	case 1, 2:
		primarySkills = append(primarySkills, 1)
		secondarySkills = append(secondarySkills, 0)
	case 3, 4, 5:
		primarySkills = append(primarySkills, 1)
		primarySkills = append(primarySkills, 0)
		secondarySkills = append(secondarySkills, 0)
	case 6, 7, 8:
		primarySkills = append(primarySkills, 2)
		primarySkills = append(primarySkills, 1)
		secondarySkills = append(secondarySkills, 1)
		secondarySkills = append(secondarySkills, 1)
	case 9, 10, 11:
		primarySkills = append(primarySkills, 3)
		primarySkills = append(primarySkills, 1)
		secondarySkills = append(secondarySkills, 1)
		secondarySkills = append(secondarySkills, 1)
	case 12, 13, 14:
		primarySkills = append(primarySkills, 3)
		primarySkills = append(primarySkills, 1)
		secondarySkills = append(secondarySkills, 2)
		secondarySkills = append(secondarySkills, 1)
	case 15:
		primarySkills = append(primarySkills, 4)
		primarySkills = append(primarySkills, 2)
		secondarySkills = append(secondarySkills, 2)
		secondarySkills = append(secondarySkills, 2)
	}
	return primarySkills, secondarySkills
}

func (cei *CEI) TaskDM() (taskDM int) {
	switch cei.EffectiveIndex {
	default:
		taskDM = -999
	case 0:
		taskDM = -6
	case 1:
		taskDM = -5
	case 2:
		taskDM = -4
	case 3:
		taskDM = -3
	case 4:
		taskDM = -2
	case 5, 6:
		taskDM = -1
	case 7, 8:
		taskDM = 0
	case 9, 10:
		taskDM = 1
	case 11:
		taskDM = 2
	case 12:
		taskDM = 3
	case 13:
		taskDM = 4
	case 14:
		taskDM = 5
	case 15:
		taskDM = 6
	}
	return taskDM
}

func (cei *CEI) ResolveMission(descr string, addMods ...int) string {
	r := dice.Roll2D()
	r += cei.TaskDM()
	descr = "  " + descr + "\n    Resolution: "
	for _, m := range addMods {
		r += m
	}
	switch r {
	case 1, 2:
		descr += "The task dissolved into chaos and was only partialy completed. A Mishap occurs.\n"
		descr += Mishap()
	case 3, 4:
		descr += "An obviously poor performance which embarrasses the ship and her officers.An Incedent occurs. \n"
		descr += Incedent()
	case 5:
		descr += "A sloppy performance, but nobody saw the mistakes."
		if dice.Roll2D() >= 10 {
			descr += "\n" + Mishap()
		}
	case 6:
		descr += "A job got done, but there was an Incedent.\n" + Incedent()
	case 7:
		descr += "A job got done. Some crews will be satisfied with this level of performance, some would be relived that it went okay."
	case 8:
		descr += "A decent job, with room for some lessons learned. However, an Opportunity occurs during the operation.\n"
		descr += Opportunity()
	case 9, 10:
		descr += "A solid performance good enough to satisfy even a critical observer."
	case 11, 12:
		descr += "A text book performance."
		if dice.Roll2D() >= 12 {
			descr += " Morale increased."
			cei.Morale++
		}
	case 13, 14:
		descr += "Near perfect resolution. as good as anyone would expect in a exercise."
		if dice.Roll2D() >= 10 {
			descr += " Morale increased."
			cei.Morale++
		}
		descr += "An Opportunity occurs during the operation.\n"
		descr += Opportunity()
	}
	return descr
}

func Mishap() string {
	ms := ""
	i := dice.Roll2D()
	switch i {
	default:
		ms = "TODO: Mishap description for index = " + strconv.Itoa(i)
	case 2:
		ms = "Structural damage is taken or weakness is detected. The ship loses '2D%' of its Hull points untill properly repaired."
	case 3:
		ms = "The ship is involved in a minor collision with small craft or object, or causes a similar incident to happen to another craft."
	case 4:
		ms = "A major system such as the spinal mount or a drive develops a fault which makes it erratic. Impose DM -1 on all checks involving that system until repaired."
	case 5:
		ms = "One of the ship’s minor systems, such as a single small craft or a point-defence battery, suffers a malfunction and is out of action until fully repaired."
	case 6:
		ms = "A crewmember is seriously injured, requiring investigation."
	case 7:
		ms = "A crewmember suffers a minor injury, which may well be his own fault. "
	case 8:
		ms = "A crewmember causes injury to someone outside the ship’s company, creating a possible incident."
	case 9:
		ms = "A crewmember manages to insult or offend someone important."
	case 10:
		ms = "One of the vehicles or small craft involved in a task suffers a serious malfuction, or a working party has an incident and require assistiance."
	case 11:
		ms = "The Travellers are given cause to suspect their plan for the current task is based on faulty data. This could be serious, such as failure to identify an atmospheric taint or a mis-estimate of the surface gravity of a nearby world. The bad information may impose DM-2 on all checks connected with it or a pose more direct hazard."
	case 12:
		ms = "A Crisis occurs."
	}

	return "      Mishap: " + ms
}

func Incedent() string {
	ms := ""
	i := dice.Roll2D()
	switch i {
	default:
		ms = "TODO: Incedent description for index = " + strconv.Itoa(i)
	case 2:
		ms = "Something goes sufficiently badly wrong that the incident becomes a Mishap."
	case 3:
		ms = "An equipment breakdown causes difficulties for the detachment or crew. For example, the main sensor processing centre suffers a software glitch and has to be rebooted whilst the ship is trying to scan local traffic. The Travellers will have to figure a way around the problem, which cannot be fixed in time to complete the task at hand."
	case 4:
		ms = "A discipline problem occurs during the operation, disrupting the work of everyone involved. A heavy-handed solution might cause further resentment, but letting crewmembers get away with indiscipline will lead to more problems in the future."
	case 5:
		ms = "The operation is made more complex by changing circumstances. This will usualy be something more inconvinient than deadly, such as equipment malfunction, unexpected enviromental changes, oe some other source of unnecessary delay or difficulty."
	case 6:
		ms = "Some of the necessary equipment is unexpectedly missing, offline or out of commission, requiring a creative workaround or a hurried fix."
	case 7:
		ms = "A minor hiccup occurs, such as an infraction of regulations that requires a hearing and disciplinary action. This might be embarrassing in front of outsiders and could have repercussions if the crew are discontent."
	case 8:
		ms = "Something unusual is spotted during the operation; a sensor blip, a strange atmosperic composition, an intriguing planetary formation or something equally puzzling but not immediately threatening."
	case 9:
		ms = "A crewmember or detachment demonstrates unexpected resourcefulness and gets the job done better and quicker than expected."
	case 10:
		ms = "A typicaly uncooperative crewmember pitches in with a will and is helpfull in completing the task. The reason for this change is not immideatly obvious."
	case 11:
		ms = "The operation requires dusting off an unusual piece of equipment or trying out a technique not normally used. Make a Difficult (10+) check using CEI and if successful, gain 1 point of Morale."
	case 12:
		ms = "Things seem to go pretty well, and in the middle of the operation an Opportunity occurs."
	}
	return "      Incedent: " + ms
}

func Opportunity() string {
	ms := ""
	i := dice.Roll2D()
	switch i {
	default:
		ms = "TODO: Opportunity description for index = " + strconv.Itoa(i)
	case 2:
		ms = "The Travellers obtain a lead on an intelligence windfall, such as a lost vessel from a foreign power or a defector. If the lead can be successfully followed up, significant advantages will be obtained."
	case 3:
		ms = "The Travellers are made aware of an equipment and supply cache hidden in the area, which may contain spares and supplies they need."
	case 4:
		ms = "A crewmember has insider information about the current situation. This might be a recruit from the world the Travellers are visiting or someone who served in an intelligence team collecting data about this subject."
	case 5:
		ms = "A piece of highly useful data is obtained. The Travellers may ‘cash in’ this find at a later date, in return for information from the referee or a bonus to resolve tasks during a mission segment."
	case 6:
		ms = "Routine intelligence-gathering produces a lead on some minor threat such as the haven of a small pirate band. With some additional investigation the Travellers may be able to locate it."
	case 7:
		ms = "The Travellers encounter a friend in an unexpected place, such as a retired former colleague in a new job. The old friend is probably willing to help out but may need something in return."
	case 8:
		ms = "A friendly vessel arrives unexpectedly and is in a position to help out or support the Travellers for a few weeks."
	case 9:
		ms = "A faction in the region asks for help from the navy, in return for which they will provide help and support. "
	case 10:
		ms = "The Travellers’ ship is unexpectedly met by a support vessel carrying equipment that will provide significant advantages, such as a couple of dozen experimental high-yield missiles sent out for field testing."
	case 11:
		ms = "The Travellers know or become aware of a circumstance in the region that could help them greatly, such as a recently laid minefield in a certain gas giant or the presence of a forward-deployed cruiser squadron which could offer support."
	case 12:
		ms = "A highly favourable circumstance occurs, such as a revolt against the government of a nearby world. Which, if successful, will cause the world to align itself more firmly with Imperial interests."
	}
	return "      Opportunity: " + ms
}

func (cei *CEI) MoraleDM() (moraleDM int) {
	switch cei.Morale {
	default:
		moraleDM = -999
	case 0:
		moraleDM = -6
	case 1:
		moraleDM = -5
	case 2:
		moraleDM = -4
	case 3:
		moraleDM = -3
	case 4:
		moraleDM = -2
	case 5, 6:
		moraleDM = -1
	case 7, 8:
		moraleDM = 0
	case 9, 10:
		moraleDM = 1
	case 11:
		moraleDM = 2
	case 12:
		moraleDM = 3
	case 13:
		moraleDM = 4
	case 14:
		moraleDM = 5
	case 15:
		moraleDM = 6
	}
	return moraleDM
}

//MinorMoraleCheck - Roll minor Morale check.
//If passed return true
//if failed return false and reduce morale by 1
//optionals: 0 - replace default difficulty (8) with this value
//           1 - cumulative DM for this check
func (cei *CEI) MinorMoraleCheck(optionals ...int) bool {
	diff := 8
	dm := 0
	for i, val := range optionals {
		if i == 0 {
			diff = val
			continue
		}
		dm += val
	}
	if dice.New().RollNext("2d6").DM(dm).Sum() >= diff {
		return true
	}
	cei.Morale--
	if cei.Morale < 0 {
		cei.Morale = 0
	}
	return false
}

//MinorMoraleCheck - Roll mojor Morale check.
//If passed return true
//if failed return false and reduce morale by 1d6
//optionals: 0 - replace default difficulty (8) with this value
//           1 - cumulative DM for this check
func (cei *CEI) MajorMoraleCheck(optionals ...int) bool {
	diff := 8
	dm := 0
	for i, val := range optionals {
		if i == 0 {
			diff = val
			continue
		}
		dm += val
	}
	if dice.New().RollNext("2d6").DM(dm).Sum() >= diff {
		return true
	}
	cei.Morale = cei.Morale - dice.Roll1D()
	if cei.Morale < 0 {
		cei.Morale = 0
	}
	return false
}
