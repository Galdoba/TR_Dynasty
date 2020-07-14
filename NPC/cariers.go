package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

const (
	careerAgent                = "Agent"
	careerArmy                 = "Army"
	careerCitizen              = "Citizen"
	careerDrifter              = "Drifter"
	careerEntertainer          = "Entertainer"
	careerMarine               = "Marine"
	careerMerchant             = "Merchant"
	careerNavy                 = "Navy"
	careerNoble                = "Noble"
	careerRogue                = "Rogue"
	careerScholar              = "Scholar"
	careerScout                = "Scout"
	careerPsionic              = "Psionic"
	careerPrisoner             = "Prisoner"
	assignmentLawEnforcement   = "Law Enforcement"
	assignmentIntelligence     = "Intelligence"
	assignmentCorporateAgent   = "Corporate Agent"
	assignmentSupportArmy      = "Support Military"
	assignmentInfantry         = "Infantry"
	assignmentCavalry          = "Cavalry"
	assignmentCorporateCitizen = "Corporate Citizen"
	assignmentWorker           = "Worker"
	assignmentColonist         = "Colonist"
	assignmentBarbarian        = "Barbarian"
	assignmentWanderer         = "Wanderer"
	assignmentScavenger        = "Scavenger"
	assignmentArtist           = "Artist"
	assignmentJournalist       = "Journalist"
	assignmentPerformer        = "Performer"
	assignmentSupportMarine    = "Support Marine"
	assignmentStarMarine       = "Star Marine"
	assignmentGroundAssault    = "Ground Assault"
	assignmentMerchantMarine   = "Merchant Marine"
	assignmentFreeTrader       = "Free Trader"
	assignmentBroker           = "Broker"
	assignmentLineCrew         = "Line/Crew"
	assignmentEngineerGunner   = "Engineer/Gunner"
	assignmentFlight           = "Flight"
	assignmentAdministrator    = "Administrator"
	assignmentDiplomat         = "Diplomat"
	assignmentDilettante       = "Dilettante"
	assignmentThief            = "Thief"
	assignmentEnforcer         = "Enforcer"
	assignmentPirate           = "Pirate"
	assignmentFieldResearcher  = "Field Researcher"
	assignmentScientist        = "Scientist"
	assignmentPhysician        = "Physician"
	assignmentCourier          = "Courier"
	assignmentSurveyor         = "Surveyor"
	assignmentExplorer         = "Explorer"
	testTypeQualification      = "Test: Qualification"
	testTypeSurvival           = "Test: Survival"
	testTypeAdvancment         = "Test: Advancment"
	testTypeCommision          = "Test: Commision"
	trainPersonalDevelopment   = "Personal Development"
	trainServiceSkills         = "Service Skills"
	trainAdvancedEducation     = "Advanced Education"
	trainAssignmentSkills      = "Assignment Skills"
	trainCommisioned           = "Commisioned"
)

type career struct {
	careerName       string
	nextTerm         string
	careerAssignment string
	qualification    string
	qualified        bool
	survival         string
	advancment       string
	advancmentMinEdu int
	commisioned      bool
	rank             int
	ended            bool
	training         map[string][]string
	benefitRolls     int
}

func setCareer(career string) *career {
	switch career {
	case careerAgent:
		return setCareerAgent()
	}
	return nil
}

func setCareerAgent() *career {
	cr := &career{}
	cr.training = make(map[string][]string)
	cr.careerName = careerAgent
	cr.advancmentMinEdu = 8
	r := utils.RollDice("d3")
	switch r {
	case 1:
		cr.careerAssignment = assignmentLawEnforcement
		cr.survival = "END 6+"
		cr.advancment = "INT 6+"
		cr.training[trainAssignmentSkills] = []string{skillInvestigate, skillRecon, skillStreetwise, skillStealth, skillMelee, skillAdvocate}
	case 2:
		cr.careerAssignment = assignmentIntelligence
		cr.survival = "INT 7+"
		cr.advancment = "INT 5+"
		cr.training[trainAssignmentSkills] = []string{skillInvestigate, skillRecon, skillElectronicsComms, skillStealth, skillPersuade, skillDeception}
	case 3:
		cr.careerAssignment = assignmentCorporateAgent
		cr.survival = "INT 5+"
		cr.advancment = "INT 7+"
		cr.training[trainAssignmentSkills] = []string{skillInvestigate, skillElectronicsComputers, skillStealth, skillCarouse, skillDeception, skillStreetwise}
	}
	cr.training[trainPersonalDevelopment] = []string{skillGunCombat, chrDEX, chrEND, skillMelee, chrINT, skillAthletics}
	cr.training[trainServiceSkills] = []string{skillStreetwise, skillDrive, skillInvestigate, skillFlyer, skillRecon, skillGunCombat}
	cr.training[trainAdvancedEducation] = []string{skillAdvocate, skillLanguage, skillExplosives, skillMedic, skillVaccSuit, skillElectronics}
	cr.qualification = "INT 6+"

	return cr
}

func (char *character) CareerLoop() *career {
	done := false
	definedCareer := false
	targetCareer := utils.RandomFromList(listCareers)
	cr := setCareer(targetCareer)
	for !definedCareer {
		//fmt.Println("defineCareer:", targetCareer)
		if cr == nil {
			targetCareer = utils.RandomFromList(listCareers)
			cr = setCareer(targetCareer)
		} else {
			definedCareer = true
		}
	}
	qualRollData := decodeCheck(cr.qualification)
	survRollData := decodeCheck(cr.survival)
	advaRollData := decodeCheck(cr.advancment)

	for !done {
		if cr.nextTerm != "" {
			log("NEWXT CAREER NOT EMPTY")
			break
		}
		char.terms[targetCareer] = char.terms[targetCareer] + 1
		if char.age > 1000 || cr.ended {
			break
		}
		if !cr.qualified {
			cr = char.quaficationCheck(qualRollData, cr)
		}
		if cr.qualified {
			train := rollTraining(cr, char)
			char.train(train)
		}
		cr = char.survivalCheck(survRollData, cr)
		char.age = char.age + 4
		char.rollAgeingEffect()
		if cr.ended {
			break
		}
		log("Life Event:")
		char.rollLifeEventTable(cr)
		//fmt.Println("Roll Commision (maybe)")
		if cr.nextTerm == "" {
			log("Advancement:")

			cr = char.advancementCheck(advaRollData, cr)
		}

	}
	musterOut(char, cr)
	fmt.Println(char.age, targetCareer, char.terms[careerAgent], cr.benefitRolls)
	char.terms[cr.careerName] = cr.rank
	return cr
}

func (char *character) quaficationCheck(rollData chrTest, cr *career) *career {
	atrMod := charDM(char.characteristics[rollData.chr])
	tn := rollData.val
	r := utils.RollDice("2d6")
	//log("Qualification Roll: " + convert.ItoS(r) + "(" + convert.ItoS(atrMod) + ")")
	r = r + atrMod + char.qualifyDM
	char.qualifyDM = 0
	//fmt.Println("Qualification Roll: " + encodeCheck(rollData) + " result: " + convert.ItoS(r))

	if r < tn {
		fmt.Println("Not Qualificated!")
		return cr
	}
	cr.qualified = true
	//fmt.Println("Qualificated!")
	return cr
}

func (char *character) survivalCheck(rollData chrTest, cr *career) *career {
	if !cr.qualified {
		//	fmt.Println("Skip Survival Roll")
		return cr
	}
	atrMod := charDM(char.characteristics[rollData.chr])
	tn := rollData.val
	cr.benefitRolls++
	r := utils.RollDice("2d6")
	log("Survival Roll: " + convert.ItoS(r+atrMod) + "(" + convert.ItoS(rollData.val) + ")")
	r = r + atrMod
	//fmt.Println(r, "is R", tn, "TN", atrMod, "atrMod")
	if r < tn {
		//fmt.Println(r, "is less than", tn, "TN")
		//fmt.Println("Roll Mishap!!")
		rollMishapEvent(cr, char)
		return cr
	}
	return cr
}

func (char *character) advancementCheck(rollData chrTest, cr *career) *career {
	if !cr.qualified || cr.ended {
		fmt.Println("Skip Advancement Roll")
		return cr
	}
	//rollData = decodeCheck(cr.advancment)
	atrMod := charDM(char.characteristics[rollData.chr])
	tn := rollData.val

	r := utils.RollDice("2d6")
	log("Advancment Roll: " + convert.ItoS(r) + "(" + convert.ItoS(atrMod) + ")")
	log("Advancment TN: " + convert.ItoS(tn))
	r = r + atrMod
	log("Advancment Reuslt: " + convert.ItoS(r))
	if r >= tn {
		train := rollTraining(cr, char)
		char.train(train)
		Advance(char, cr)
		log("Advancement successful! New Rank = " + convert.ItoS(cr.rank))
	}
	if r <= char.terms[cr.careerName] {
		log("End career because of time")
		cr.ended = true
	}
	return cr
}

func Advance(char *character, cr *career) {
	cr.rank++
	switch cr.careerAssignment {
	default:
		log("TODO ADVANCEMENT BONUS FOR " + cr.careerAssignment)
	case assignmentLawEnforcement:
		if cr.rank == 1 {
			char.train(skillStreetwise, 1)
		}
		if cr.rank == 4 {
			char.train(skillInvestigate, 1)
		}
		if cr.rank == 5 {
			char.train(skillAdmin, 1)
		}
		if cr.rank == 6 {
			char.train(chrSOC)
		}
	case assignmentIntelligence, assignmentCorporateAgent:
		if cr.rank == 1 {
			char.train(skillDeception, 1)
			log("tarin: " + skillDeception)
		}
		if cr.rank == 2 {
			char.train(skillInvestigate, 1)
		}
		if cr.rank == 4 {
			char.train(skillGunCombat, 1)
		}
	}
}

func validTraningTable(char *character, cr *career) []string {
	var list []string
	for key := range cr.training {
		if key == trainAdvancedEducation {
			if char.characteristics[chrEDU] < cr.advancmentMinEdu {
				continue
			}
		}
		if key == trainCommisioned {
			if !cr.commisioned {
				continue
			}
		}
		list = append(list, key)
	}
	return list
}

func pickFromTrainingTable(cr *career, trainingTable string) string {
	return utils.RandomFromList(cr.training[trainingTable])
}

func rollTraining(cr *career, char *character) string {
	trainingTable := validTraningTable(char, cr)
	train := utils.RandomFromList(cr.training[utils.RandomFromList(trainingTable)])
	log("Training: " + train)
	return train
}

/*
	mainLoop:

	1 Выбор карьеры
	1.а Составить список доступных карьер и выбрать случайный
	0 Образование (50%)
	0.а Если да выбрать тип института
	0.б Выбрать (специализацию и навыки)
	2 Выбор назначения
	2.а выбрать случайный из карьеры
	3 Квалификация
	конец карьеры = false
	3.а Если это первый срок - бросить квалификацию, иначе пропустить
	3.б Выбрать таблицу и бросить на результат
	3.в Бросить Выживание
		если да бросить Life Event
		если нет конец карьеры = true
*/

func desidionPreCareerEducation(char *character) bool {
	if char.currentTerm > 3 {
		return false
	}
	if utils.RandomBool() {
		return false
	}
	return true
}

type chrTest struct {
	chr string
	val int
}

func (char *character) runTest(test chrTest, dm int) bool {
	r := utils.RollDice("2d6", charDM(char.characteristics[test.chr])+dm)
	if r < test.val {
		return false
	}
	return true
}

type term struct {
	career     string
	assignment string
	rank       int
}

func testFor(assignment string, testType string) *chrTest {
	//chrTest := &chrTest{}
	switch testType {
	case testTypeQualification:
		//AGENT
		if assignment == assignmentLawEnforcement || assignment == assignmentIntelligence || assignment == assignmentCorporateAgent {
			return &chrTest{chrINT, 6}
		}
		//ARMY
		if assignment == assignmentSupportArmy || assignment == assignmentInfantry || assignment == assignmentCavalry {
			return &chrTest{chrEND, 5}
		}
		//CITIZEN
		if assignment == assignmentCorporateCitizen || assignment == assignmentWorker || assignment == assignmentColonist {
			return &chrTest{chrEDU, 5}
		}
		//DRIFTER
		if assignment == assignmentBarbarian || assignment == assignmentWanderer || assignment == assignmentScavenger {
			return &chrTest{chrEND, -100} //AUTOMATIC
		}
		//ENTERTAINMENT
		if assignment == assignmentArtist || assignment == assignmentJournalist || assignment == assignmentPerformer {
			return &chrTest{chrINT, 5} //DEX or INT
		}
		//MARINE
		if assignment == assignmentSupportMarine || assignment == assignmentStarMarine || assignment == assignmentGroundAssault {
			return &chrTest{chrEND, 6}
		}
		//MERCHANT
		if assignment == assignmentMerchantMarine || assignment == assignmentFreeTrader || assignment == assignmentBroker {
			return &chrTest{chrINT, 4}
		}
		//NAVY
		if assignment == assignmentLineCrew || assignment == assignmentEngineerGunner || assignment == assignmentFlight {
			return &chrTest{chrINT, 6}
		}
		//NOBLE
		if assignment == assignmentAdministrator || assignment == assignmentDiplomat || assignment == assignmentDilettante {
			return &chrTest{chrSOC, 10}
		}
		//ROGUE
		if assignment == assignmentThief || assignment == assignmentEnforcer || assignment == assignmentPirate {
			return &chrTest{chrDEX, 6}
		}
		//SCHOLAR
		if assignment == assignmentFieldResearcher || assignment == assignmentScientist || assignment == assignmentPhysician {
			return &chrTest{chrINT, 6}
		}
		//SCOUT
		if assignment == assignmentCourier || assignment == assignmentSurveyor || assignment == assignmentExplorer {
			return &chrTest{chrINT, 5}
		}
	case testTypeSurvival:
		switch assignment {
		case assignmentLawEnforcement:
			return &chrTest{chrEND, 6}
		case assignmentIntelligence:
			return &chrTest{chrINT, 7}
		case assignmentCorporateAgent:
			return &chrTest{chrINT, 5}
		case assignmentSupportArmy:
			return &chrTest{chrEND, 5}
		case assignmentInfantry:
			return &chrTest{chrSTR, 6}
		case assignmentCavalry:
			return &chrTest{chrDEX, 7}
		case assignmentCorporateCitizen:
			return &chrTest{chrSOC, 6}
		case assignmentWorker:
			return &chrTest{chrEND, 4}
		case assignmentColonist:
			return &chrTest{chrINT, 7}
		case assignmentBarbarian:
			return &chrTest{chrEND, 7}
		case assignmentWanderer:
			return &chrTest{chrEND, 7}
		case assignmentScavenger:
			return &chrTest{chrDEX, 7}
		case assignmentArtist:
			return &chrTest{chrSOC, 6}
		case assignmentJournalist:
			return &chrTest{chrEDU, 7}
		case assignmentPerformer:
			return &chrTest{chrINT, 5}
		case assignmentSupportMarine:
			return &chrTest{chrEND, 5}
		case assignmentStarMarine:
			return &chrTest{chrEND, 6}
		case assignmentGroundAssault:
			return &chrTest{chrEND, 7}
		case assignmentMerchantMarine:
			return &chrTest{chrEDU, 5}
		case assignmentFreeTrader:
			return &chrTest{chrDEX, 6}
		case assignmentBroker:
			return &chrTest{chrEDU, 5}
		case assignmentLineCrew:
			return &chrTest{chrINT, 5}
		case assignmentEngineerGunner:
			return &chrTest{chrINT, 6}
		case assignmentFlight:
			return &chrTest{chrDEX, 7}
		case assignmentAdministrator:
			return &chrTest{chrINT, 4}
		case assignmentDiplomat:
			return &chrTest{chrINT, 5}
		case assignmentDilettante:
			return &chrTest{chrSOC, 3}
		case assignmentThief:
			return &chrTest{chrINT, 6}
		case assignmentEnforcer:
			return &chrTest{chrEND, 6}
		case assignmentPirate:
			return &chrTest{chrDEX, 6}
		case assignmentFieldResearcher:
			return &chrTest{chrEND, 6}
		case assignmentScientist:
			return &chrTest{chrEDU, 4}
		case assignmentPhysician:
			return &chrTest{chrEDU, 4}
		case assignmentCourier:
			return &chrTest{chrEND, 5}
		case assignmentSurveyor:
			return &chrTest{chrEND, 6}
		case assignmentExplorer:
			return &chrTest{chrEND, 7}
		}
	case testTypeAdvancment:
		switch assignment {
		case assignmentLawEnforcement:
			return &chrTest{chrINT, 6}
		case assignmentIntelligence:
			return &chrTest{chrINT, 5}
		case assignmentCorporateAgent:
			return &chrTest{chrINT, 7}
		case assignmentSupportArmy:
			return &chrTest{chrEDU, 7}
		case assignmentInfantry:
			return &chrTest{chrEDU, 6}
		case assignmentCavalry:
			return &chrTest{chrINT, 5}
		case assignmentCorporateCitizen:
			return &chrTest{chrINT, 6}
		case assignmentWorker:
			return &chrTest{chrEDU, 8}
		case assignmentColonist:
			return &chrTest{chrEND, 5}
		case assignmentBarbarian:
			return &chrTest{chrSTR, 7}
		case assignmentWanderer:
			return &chrTest{chrINT, 7}
		case assignmentScavenger:
			return &chrTest{chrEND, 7}
		case assignmentArtist:
			return &chrTest{chrINT, 6}
		case assignmentJournalist:
			return &chrTest{chrINT, 5}
		case assignmentPerformer:
			return &chrTest{chrDEX, 7}
		case assignmentSupportMarine:
			return &chrTest{chrEDU, 7}
		case assignmentStarMarine:
			return &chrTest{chrEDU, 6}
		case assignmentGroundAssault:
			return &chrTest{chrEDU, 5}
		case assignmentMerchantMarine:
			return &chrTest{chrINT, 7}
		case assignmentFreeTrader:
			return &chrTest{chrINT, 6}
		case assignmentBroker:
			return &chrTest{chrINT, 7}
		case assignmentLineCrew:
			return &chrTest{chrEDU, 7}
		case assignmentEngineerGunner:
			return &chrTest{chrEDU, 6}
		case assignmentFlight:
			return &chrTest{chrEDU, 5}
		case assignmentAdministrator:
			return &chrTest{chrEDU, 6}
		case assignmentDiplomat:
			return &chrTest{chrSOC, 7}
		case assignmentDilettante:
			return &chrTest{chrINT, 8}
		case assignmentThief:
			return &chrTest{chrDEX, 6}
		case assignmentEnforcer:
			return &chrTest{chrSTR, 6}
		case assignmentPirate:
			return &chrTest{chrINT, 6}
		case assignmentFieldResearcher:
			return &chrTest{chrINT, 6}
		case assignmentScientist:
			return &chrTest{chrINT, 8}
		case assignmentPhysician:
			return &chrTest{chrEDU, 8}
		case assignmentCourier:
			return &chrTest{chrEDU, 9}
		case assignmentSurveyor:
			return &chrTest{chrINT, 8}
		case assignmentExplorer:
			return &chrTest{chrEDU, 7}
		}

	}
	return nil
}

func randomAssignment(career string) string {
	result := ""
	switch career {
	case careerAgent:
		return utils.RandomFromList([]string{assignmentLawEnforcement, assignmentIntelligence, assignmentCorporateAgent})
	}
	return result
}

func termsCompleted(char *character) int {
	//temp TODO в последствии считаем по карьере
	term := ((char.age - 18) / 4)
	return term
}

func newTest(atr string, tn int) *chrTest {
	test := &chrTest{}
	test.chr = atr
	test.val = tn
	return test
}

func decodeCheck(code string) chrTest {
	test := chrTest{}
	parts := strings.Split(code, " ")
	atr := parts[0]
	tnStr := strings.Split(parts[1], "+")
	tn, _ := strconv.Atoi(tnStr[0])
	test.chr = atr
	test.val = tn
	return test
}

func encodeCheck(rollData chrTest) string {
	return rollData.chr + " " + convert.ItoS(rollData.val) + "+"
}

// chrEND, 6  chrINT, 6
// chrINT, 7  chrINT, 5
// chrINT, 5  chrINT, 7
// chrEND, 5  chrEDU, 7
// chrSTR, 6  chrEDU, 6
// chrDEX, 7  chrINT, 5
// chrSOC, 6  chrINT, 6
// chrEND, 4  chrEDU, 8
// chrINT, 7  chrEND, 5
// chrEND, 7  chrSTR, 7
// chrEND, 7  chrINT, 7
// chrDEX, 7  chrEND, 7
// chrSOC, 6  chrINT, 6
// chrEDU, 7  chrINT, 5
// chrINT, 5  chrDEX, 7
// chrEND, 5  chrEDU, 7
// chrEND, 6  chrEDU, 6
// chrEND, 7  chrEDU, 5
// chrEDU, 5  chrINT, 7
// chrDEX, 6  chrINT, 6
// chrEDU, 5  chrINT, 7
// chrINT, 5  chrEDU, 7
// chrINT, 6  chrEDU, 6
// chrDEX, 7  chrEDU, 5
// chrINT, 4  chrEDU, 6
// chrINT, 5  chrSOC, 7
// chrSOC, 3  chrINT, 8
// chrINT, 6  chrDEX, 6
// chrEND, 6  chrSTR, 6
// chrDEX, 6  chrINT, 6
// chrEND, 6  chrINT, 6
// chrEDU, 4  chrINT, 8
// chrEDU, 4  chrEDU, 8
// chrEND, 5  chrEDU, 9
// chrEND, 6  chrINT, 8
// chrEND, 7  chrEDU, 7

//LISTS

func musterOut(char *character, cr *career) {
	log("Mustering Out from " + cr.careerName + " " + convert.ItoS(cr.benefitRolls))
	switch cr.careerName {
	case careerAgent:
		musterOutAgent(char, cr)
	}
}

func musterOutAgent(char *character, cr *career) {
	for cr.benefitRolls > 0 {
		cr.benefitRolls--
		r := utils.RollDice("d6")
		switch r {
		case 2:
			char.train(chrINT)
			log("Train: " + chrINT)
		case 6:
			char.train(chrSOC)
			log("Train: " + chrSOC)
		}

	}
}
