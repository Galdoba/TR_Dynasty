package cei

import "github.com/Galdoba/TR_Dynasty/pkg/dice"

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

func (cei *CEI) ResolveTask() {

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
