package main

import (
	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

//INJURY TABLE

func decodeRollIndex(rollType string, index ...int) int {
	r := 0
	if len(index) != 0 {
		r = index[0]
	} else {
		r = utils.RollDice(rollType)
	}
	return r
}

func (char *character) rollInjuryTable(index ...int) {
	r := decodeRollIndex("d6", index...)
	switch r {
	case 1:
		char.injury1()
	case 2:
		char.injury2()
	case 3:
		char.injury3()
	case 4:
		char.injury4()
	case 5:
		char.injury5()
	case 6:
		char.injury6()
	}
}

func (char *character) injury1() {
	log("Nearly Killed!!")
	var atrs []string
	for len(atrs) < 3 {
		atrs = utils.AppendUniqueStr(atrs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	n := utils.RollDice("d6")
	char.changeChrBy(atrs[0], -n)
	log("Reduce " + atrs[0] + " by " + convert.ItoS(n))
	char.changeChrBy(atrs[1], -2)
	log("Reduce " + atrs[1] + " by 2")
	char.changeChrBy(atrs[2], -2)
	log("Reduce " + atrs[2] + " by 2")
}

func (char *character) injury2() {
	log("Severly injured!!")
	var atrs []string
	for len(atrs) < 1 {
		atrs = utils.AppendUniqueStr(atrs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	n := utils.RollDice("d6")
	char.characteristics[atrs[0]] = char.characteristics[atrs[0]] - n
	log("Reduce " + atrs[0] + " by " + convert.ItoS(n))
}

func (char *character) injury3() {
	log("Lost eye or limb!!")
	var atrs []string
	for len(atrs) < 1 {
		atrs = utils.AppendUniqueStr(atrs, utils.RandomFromList([]string{chrSTR, chrDEX}))
	}
	char.characteristics[atrs[0]] = char.characteristics[atrs[0]] - 2
	log("Reduce " + atrs[0] + " by 2")
}

func (char *character) injury4() {
	log("Scarred and injured!!")
	var atrs []string
	for len(atrs) < 1 {
		atrs = utils.AppendUniqueStr(atrs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.characteristics[atrs[0]] = char.characteristics[atrs[0]] - 2
	log("Reduce " + atrs[0] + " by 2")

}

func (char *character) injury5() {
	log("Injured!!")
	var atrs []string
	for len(atrs) < 1 {
		atrs = utils.AppendUniqueStr(atrs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.characteristics[atrs[0]] = char.characteristics[atrs[0]] - 1
	log("Reduce " + atrs[0] + " by 1")

}

func (char *character) injury6() {
	log("Lightly injured! No permanent effect.")
}

//AGEING TABLE

func (char *character) rollAgeingEffect() {
	if char.age < 34 {
		return
	}
	dm := termsCompleted(char)
	ageRoll := utils.RollDice("2d6")
	effect := ageRoll - dm
	log("ROLL AGEING   " + "dm=" + convert.ItoS(dm) + " ageRoll=" + convert.ItoS(ageRoll) + " effect=" + convert.ItoS(effect))
	log("Ageing Effects:")
	if effect > 0 {
		log("No Effect")
		return
	}
	switch effect {
	default:
		char.ageing6()
	case -5:
		char.ageing5()
	case -4:
		char.ageing4()
	case -3:
		char.ageing3()
	case -2:
		char.ageing2()
	case -1:
		char.ageing1()
	case 0:
		char.ageing0()
	}
}

func (char *character) ageing6() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	var mCHRs []string
	for len(mCHRs) < 3 {
		mCHRs = utils.AppendUniqueStr(mCHRs, utils.RandomFromList([]string{chrINT, chrEDU, chrSOC}))
	}
	char.changeChrBy(pCHRs[0], -2)
	log("Reduce " + pCHRs[0] + " by 2")
	char.changeChrBy(pCHRs[1], -2)
	log("Reduce " + pCHRs[1] + " by 2")
	char.changeChrBy(pCHRs[2], -2)
	log("Reduce " + pCHRs[2] + " by 2")
	char.changeChrBy(mCHRs[0], -1)
	log("Reduce " + mCHRs[0] + " by 1")
}

func (char *character) ageing5() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.changeChrBy(pCHRs[0], -2)
	log("Reduce " + pCHRs[0] + " by 2")
	char.changeChrBy(pCHRs[1], -2)
	log("Reduce " + pCHRs[1] + " by 2")
	char.changeChrBy(pCHRs[2], -2)
	log("Reduce " + pCHRs[2] + " by 2")
}

func (char *character) ageing4() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.changeChrBy(pCHRs[0], -2)
	log("Reduce " + pCHRs[0] + " by 2")
	char.changeChrBy(pCHRs[1], -2)
	log("Reduce " + pCHRs[1] + " by 2")
	char.changeChrBy(pCHRs[2], -1)
	log("Reduce " + pCHRs[2] + " by 1")
}

func (char *character) ageing3() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.changeChrBy(pCHRs[0], -2)
	log("Reduce " + pCHRs[0] + " by 2")
	char.changeChrBy(pCHRs[1], -1)
	log("Reduce " + pCHRs[1] + " by 1")
	char.changeChrBy(pCHRs[2], -1)
	log("Reduce " + pCHRs[2] + " by 1")
}

func (char *character) ageing2() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.changeChrBy(pCHRs[0], -1)
	log("Reduce " + pCHRs[0] + " by 1")
	char.changeChrBy(pCHRs[1], -1)
	log("Reduce " + pCHRs[1] + " by 1")
	char.changeChrBy(pCHRs[2], -1)
	log("Reduce " + pCHRs[2] + " by 1")
}

func (char *character) ageing1() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.changeChrBy(pCHRs[0], -1)
	log("Reduce " + pCHRs[0] + " by 1")
	char.changeChrBy(pCHRs[1], -1)
	log("Reduce " + pCHRs[1] + " by 1")
}

func (char *character) ageing0() {
	var pCHRs []string
	for len(pCHRs) < 3 {
		pCHRs = utils.AppendUniqueStr(pCHRs, utils.RandomFromList([]string{chrSTR, chrDEX, chrEND}))
	}
	char.changeChrBy(pCHRs[0], -1)
	log("Reduce " + pCHRs[0] + " by 1")
}

//LIFE EVENTS TABLE

func (char *character) rollLifeEventTable(cr *career) {
	event := utils.RollDice("2d6")
	switch event {
	case 2:
		char.lifeEvent2()
	case 3:
		char.lifeEvent3()
	case 4:
		char.lifeEvent4()
	case 5:
		char.lifeEvent5()
	case 6:
		char.lifeEvent6()
	case 7:
		char.lifeEvent7()
	case 8:
		char.lifeEvent8()
	case 9:
		char.lifeEvent9()
	case 10:
		char.lifeEvent10()
	case 11:
		char.lifeEvent11()
		if utils.RandomBool() {
			cr.benefitRolls--
		} else {
			cr.nextTerm = careerPrisoner
		}
	case 12:
		char.lifeEvent12()

	}
}

func (char *character) lifeEvent2() {
	log("You are injured or become sick")
	char.rollInjuryTable()
}

func (char *character) lifeEvent3() {
	log("Someone close to you dies, like a friend or family member. Alternatively, someone close to you gives birth (or is born!).  You are involved in some fashion (father or mother, relative, godparent, etc).")
}

func (char *character) lifeEvent4() {
	log("A romantic relationship involving you ends. Badly. Gain a Rival or Enemy.")
}

func (char *character) lifeEvent5() {
	log("A romantic relationship involving you deepens, possibly leading to marriage or some other emotional commitment. Gain an Ally.")
}

func (char *character) lifeEvent6() {
	log("You become involved in a romantic relationship. Gain an Ally.")
}

func (char *character) lifeEvent7() {
	log("You gain a new Contact.")
}

func (char *character) lifeEvent8() {
	log("You are betrayed in some fashion by a friend.  If you have any Contacts or Allies, convert one into a Rival or Enemy. Otherwise, gain a Rival or an Enemy.")
}

func (char *character) lifeEvent9() {
	log("You move to another world. You gain DM+2 to your next Qualification roll.")
	char.qualifyDM = char.qualifyDM + 2
}

func (char *character) lifeEvent10() {
	log("Something good happens to you; you come into money unexpectedly, have a lifelong dream come true, get a book published or have some other stroke of good fortune. Gain DM+2 to any one Benefit roll.")
}

func (char *character) lifeEvent11() {
	log("You commit or are the victim (or are accused) of a crime. Lose one Benefit roll or take the Prisoner career in your next term.")
}

func (char *character) lifeEvent12() {
	log("TODO: unusualEvent() - needMechanical side")
	char.unusualEvent()
}

func (char *character) unusualEvent() {
	r := utils.RollDice("d6")
	switch r {
	case 1:
		log("You encounter a Psionic institute. You may immediately test your Psionic Strength and, if you qualify, take the Psion career in your next term.")
	case 2:
		log("You spend time among an alien race.")
		char.train(skillScience, 1)
	case 3:
		log("You have a strange and unusual device from an alien culture that is not normally available to humans.")
	case 4:
		log("Something happened to you, but you do not know what it was.")
	case 5:
		log("You briefly came into contact with the highest echelons of the Imperium – an Archduke or the Emperor, perhaps, or Imperial intelligence.")
	case 6:
		log("You have something older than the Imperium, or even something older than humanity.")
	}
}

func rollMishapTableAgent(char *character, cr *career, index ...int) {
	r := decodeRollIndex("d6", index...)
	log("Mishap Roll:")
	switch r {
	case 1:
		agentMishap1(char, cr)
	case 2:
		agentMishap2(char, cr)
	case 3:
		agentMishap3(char, cr)
	case 4:
		agentMishap4(char, cr)
	case 5:
		agentMishap5(char, cr)
	case 6:
		agentMishap6(char, cr)
	}
	cr.benefitRolls--
}

func rollMishapEvent(cr *career, char *character) {
	switch cr.careerName {
	default:
		log("TODO: " + cr.careerName + " mishap event")
	case careerAgent:
		rollMishapTableAgent(char, cr)
	}
	cr.ended = true
}

func agentMishap1(char *character, cr *career) {
	log("Severely injured")
	char.injury1()
}

func agentMishap2(char *character, cr *career) {
	log("A criminal or other figure under investigation offers you a deal.")
	if utils.RandomBool() {
		log("Offer was accepted:")
		log("Career ended")
	} else {
		log("Offer was refused:")
		r := utils.Min(utils.RollDice("d6"), utils.RollDice("d6"))
		char.rollInjuryTable(r)
		skill := utils.RandomFromList(listSkills)
		log("Train: " + skill)
		char.train(skill)
	}

}

func agentMishap3(char *character, cr *career) {
	log("An investigation goes critically wrong or leads to the top, ruining your career.")
	task := NewTask()
	task.SetParameters(skillAdvocate)
	d, _, _, s := char.skillCheck(task)
	if d == 2 {
		cr.nextTerm = careerPrisoner
	}
	if s {
		cr.benefitRolls++
	}

}

func agentMishap4(char *character, cr *career) {
	log("You learn something you should not know, and people want to kill you for it.")
	log("Gain an Enemy")
	char.train(skillDeception, 1)
}

func agentMishap5(char *character, cr *career) {
	log("Your work ends up coming home with you, and someone gets hurt.")
	log("Lost Ally, Contact or gain an Enemy")

}

func agentMishap6(char *character, cr *career) {
	log("Injured")
	char.rollInjuryTable()
}
