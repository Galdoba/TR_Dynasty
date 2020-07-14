package main

import (
	"strings"

	"github.com/Galdoba/utils"
)

const (
	skillAdmin                   = "Admin"
	skillAdvocate                = "Advocate"
	skillAnimals                 = "Animals"
	skillAnimalsHandling         = "Animals (Handling)"
	skillAnimalsVeterinary       = "Animals (Veterinary)"
	skillAnimalsTraining         = "Animals (Training)"
	skillArt                     = "Art"
	skillArtPerformer            = "Art (Performer)"
	skillArtHolography           = "Art (Holography)"
	skillArtInstrument           = "Art (Instrument)"
	skillArtVisualMedia          = "Art (Visual Media)"
	skillArtWrite                = "Art (Write)"
	skillAstrogation             = "Astrogation"
	skillAthletics               = "Athletics"
	skillAthleticsSTR            = "Athletics (STR)"
	skillAthleticsDEX            = "Athletics (DEX)"
	skillAthleticsEND            = "Athletics (END)"
	skillBroker                  = "Broker"
	skillCarouse                 = "Carouse"
	skillDeception               = "Deception"
	skillDiplomat                = "Diplomat"
	skillDrive                   = "Drive"
	skillDriveHovercraft         = "Drive (Hovercraft)"
	skillDriveMole               = "Drive (Mole)"
	skillDriveTrack              = "Drive (Track)"
	skillDriveWalker             = "Drive (Walker)"
	skillDriveWheel              = "Drive (Wheel)"
	skillElectronics             = "Electronics"
	skillElectronicsComms        = "Electronics (Comms)"
	skillElectronicsComputers    = "Electronics (Computers)"
	skillElectronicsRemoteOps    = "Electronics (Remote Ops)"
	skillElectronicsSensors      = "Electronics (Sensors)"
	skillEngineer                = "Engineer"
	skillEngineerMDrive          = "Engineer (M-Drive)"
	skillEngineerJDrive          = "Engineer (J-Drive)"
	skillEngineerLifeSupport     = "Engineer (Life Support)"
	skillEngineerPower           = "Engineer (Power)"
	skillExplosives              = "Explosives"
	skillFlyer                   = "Flyer"
	skillFlyerAirship            = "Flyer (Airship)"
	skillFlyerGrav               = "Flyer (Grav)"
	skillFlyerOrnithopter        = "Flyer (Ornithopter)"
	skillFlyerRotor              = "Flyer (Rotor)"
	skillFlyerWing               = "Flyer (Wing)"
	skillGambler                 = "Gambler"
	skillGunner                  = "Gunner"
	skillGunnerTurret            = "Gunner (Turret)"
	skillGunnerOrtilery          = "Gunner (Ortilery)"
	skillGunnerScreen            = "Gunner (Screen)"
	skillGunnerCapital           = "Gunner (Capital)"
	skillGunCombat               = "Gun Combat"
	skillGunCombatArchaic        = "Gun Combat (Archaic)"
	skillGunCombatEnergy         = "Gun Combat (Energy)"
	skillGunCombatSlug           = "Gun Combat (Slug)"
	skillHeavyWeapons            = "Heavy Weapons"
	skillHeavyWeaponsArtillery   = "Heavy Weapons (Аrtilery)"
	skillHeavyWeaponsManPortable = "Heavy Weapons (Man Portable)"
	skillHeavyWeaponsVechicle    = "Heavy Weapons (Vechicle)"
	skillInvestigate             = "Investigate"
	skillJackOfAllTrades         = "Jack-Of-All-Trades"
	skillLanguage                = "Language"
	skillLeadership              = "Leadership"
	skillMechanic                = "Mechanic"
	skillMedic                   = "Medic"
	skillMelee                   = "Melee"
	skillMeleeUnarmed            = "Melee (Unarmed)"
	skillMeleeBlade              = "Melee (Blade)"
	skillMeleeBludgeon           = "Melee (Bludgeon)"
	skillNavigation              = "Navigation"
	skillPersuade                = "Persuade"
	skillPilot                   = "Pilot"
	skillPilotSmallCraft         = "Pilot (Small Craft)"
	skillPilotSpacecraft         = "Pilot (Spacecraft)"
	skillPilotCapitalShips       = "Pilot (Capital Ships)"
	skillProfession              = "Profession"
	skillRecon                   = "Recon"
	skillScience                 = "Science"
	skillSeafarer                = "Seafarer"
	skillStealth                 = "Stealth"
	skillSteward                 = "Steward"
	skillStreetwise              = "Streetwise"
	skillSurvival                = "Survival"
	skillTactics                 = "Tactics"
	skillTacticsMilitary         = "Tactics (Military)"
	skillTacticsNaval            = "Tactics (Naval)"
	skillVaccSuit                = "Vacc Suit"
	difficultyVeryEasy           = -4
	difficultyEasy               = -2
	difficultyAverage            = 0
	difficultyHard               = 2
	difficultyVeryHard           = 4
	difficultyFormidable         = 6
	tFrameSeconds                = "Second"
	tFrameRoundCombat            = "Combat Round"
	tFrameSeconds10              = "10 Seconds"
	tFrameMinutes                = "Minute"
	tFrameRoundSpace             = "Space Combat Round"
	tFrameMinutes10              = "10 Minutes"
	tFrameHours                  = "Hour"
	tFrameHours4                 = "4 Hours"
	tFrameHours10                = "10 Hours"
	tFrameDays                   = "Day"
	tFrameDays4                  = "4 Days"
	tFrameWeeks                  = "Week"
	tFrameWeeks2                 = "2 Week"
	tFrameMonths                 = "Month"
	tFrameMonths6                = "6 Months"
	tFrameYears                  = "Year"
)

type sklGroup struct {
	name    string
	trained bool
	specs   []string
}

var listBackgroundSkills []string
var listSkills []string
var listUniversitySkills []string
var listGraduateBenefits []string
var listCareers []string
var listCHARS []string

func preapareLists() {
	listCHARS = []string{
		chrSTR,
		chrDEX,
		chrEND,
		chrINT,
		chrEDU,
		chrSOC,
		chrPSI,
	}
	listBackgroundSkills = []string{
		skillAdmin,
		skillAnimals,
		skillArt,
		skillAthletics,
		skillCarouse,
		skillDrive,
		skillElectronics,
		skillFlyer,
		skillLanguage,
		skillMechanic,
		skillMedic,
		skillProfession,
		skillScience,
		skillSeafarer,
		skillStreetwise,
		skillSurvival,
		skillVaccSuit,
	}
	listSkills = []string{
		skillAdmin,
		skillAdvocate,
		skillAnimals,
		skillAnimalsHandling,
		skillAnimalsVeterinary,
		skillAnimalsTraining,
		skillArt,
		skillArtPerformer,
		skillArtHolography,
		skillArtInstrument,
		skillArtVisualMedia,
		skillArtWrite,
		skillAstrogation,
		skillAthletics,
		skillAthleticsSTR,
		skillAthleticsDEX,
		skillAthleticsEND,
		skillBroker,
		skillCarouse,
		skillDeception,
		skillDiplomat,
		skillDrive,
		skillDriveHovercraft,
		skillDriveMole,
		skillDriveTrack,
		skillDriveWalker,
		skillDriveWheel,
		skillElectronics,
		skillElectronicsComms,
		skillElectronicsComputers,
		skillElectronicsRemoteOps,
		skillElectronicsSensors,
		skillEngineer,
		skillEngineerMDrive,
		skillEngineerJDrive,
		skillEngineerLifeSupport,
		skillEngineerPower,
		skillExplosives,
		skillFlyer,
		skillFlyerAirship,
		skillFlyerGrav,
		skillFlyerOrnithopter,
		skillFlyerRotor,
		skillFlyerWing,
		skillGambler,
		skillGunner,
		skillGunnerTurret,
		skillGunnerOrtilery,
		skillGunnerScreen,
		skillGunnerCapital,
		skillGunCombat,
		skillGunCombatArchaic,
		skillGunCombatEnergy,
		skillGunCombatSlug,
		skillHeavyWeapons,
		skillHeavyWeaponsArtillery,
		skillHeavyWeaponsManPortable,
		skillHeavyWeaponsVechicle,
		skillInvestigate,
		skillJackOfAllTrades,
		skillLanguage,
		skillLeadership,
		skillMechanic,
		skillMedic,
		skillMelee,
		skillMeleeUnarmed,
		skillMeleeBlade,
		skillMeleeBludgeon,
		skillNavigation,
		skillPersuade,
		skillPilot,
		skillPilotSmallCraft,
		skillPilotSpacecraft,
		skillPilotCapitalShips,
		skillProfession,
		skillRecon,
		skillScience,
		skillSeafarer,
		skillStealth,
		skillSteward,
		skillStreetwise,
		skillSurvival,
		skillTactics,
		skillTacticsMilitary,
		skillTacticsNaval,
		skillVaccSuit,
	}
	listUniversitySkills = []string{
		skillAdmin,
		skillAdvocate,
		skillAnimalsTraining,
		skillAnimalsVeterinary,
		skillArtHolography,
		skillArtInstrument,
		skillArtPerformer,
		skillArtVisualMedia,
		skillArtWrite,
		skillAstrogation,
		skillElectronicsComms,
		skillElectronicsComputers,
		skillElectronicsRemoteOps,
		skillElectronicsSensors,
		skillEngineerJDrive,
		skillEngineerLifeSupport,
		skillEngineerMDrive,
		skillEngineerPower,
		skillLanguage,
		skillMedic,
		skillNavigation,
		skillProfession,
		skillScience,
	}
	listGraduateBenefits = []string{
		assignmentLawEnforcement,
		assignmentIntelligence,
		assignmentCorporateAgent,
		assignmentSupportArmy,
		assignmentInfantry,
		assignmentCavalry,
		assignmentCorporateCitizen,
		assignmentJournalist,
		assignmentSupportMarine,
		assignmentStarMarine,
		assignmentGroundAssault,
		assignmentLineCrew,
		assignmentEngineerGunner,
		assignmentFlight,
		assignmentFieldResearcher,
		assignmentScientist,
		assignmentPhysician,
		assignmentCourier,
		assignmentSurveyor,
		assignmentExplorer,
	}
	listCareers = []string{
		careerAgent,
		careerArmy,
		careerCitizen,
		careerDrifter,
		careerEntertainer,
		careerMarine,
		careerMerchant,
		careerNavy,
		careerNoble,
		careerRogue,
		careerScholar,
		careerScout,
		careerPsionic,
		careerPrisoner,
	}
}

//нужные функции
/*
1. Проверка есть ли у скила специализация
2. Метод выбора новой специализации
3. вернуть группу func skillGroupOf(skill string) (group string){}
*/

func (char *character) train(skillName string, val ...int) {
	//fmt.Println("Start:", skillName)
	//skillName - со специализацией или нет?
	//specWasSet := false
	//valWasSet := false
	//val - указано или нет?
	newVal := -1
	if len(val) < 1 {
		newVal = char.skills[skillName] + 1
	} else {
		if char.skills[skillName] < val[0] {
			//fmt.Println("val[0]>char.skills[skillName] - incresing")
			newVal = val[0]
		} else {
			//fmt.Println("val[0]<=char.skills[skillName] - not incresing")
			newVal = char.skills[skillName]
		}
		//newVal = val[0]
		//	valWasSet = true
	}
	if strings.Contains(skillName, "(") {
		//	specWasSet = true
	}

	for key := range char.characteristics {
		if key == skillName {
			char.characteristics[key] = char.characteristics[key] + 1
		}
	}
	// if specWasSet == true && valWasSet == true {
	// 	fmt.Println("Adding", skillName, newVal)
	// }
	// if specWasSet == false && valWasSet == true {
	// 	fmt.Println("Adding random SPEC", skillName, newVal)
	// 	skillName = utils.RandomFromList(skillSpecs(skillName))
	// }
	// if specWasSet == true && valWasSet == false {
	// 	fmt.Println("Adding ", skillName, "by +1")
	// }
	// if specWasSet == false && valWasSet == false {
	// 	skillName = utils.RandomFromList(skillSpecs(skillName))
	// 	fmt.Println("Adding random SPEC", skillName, "by +1")
	// }

	char.skills[skillName] = newVal
}

func isGroup(skl string) bool {
	if !strings.Contains(skl, "(") {
		return true
	}
	return false
}

func skillSpecs(skl string) []string {
	var result []string
	for i := range listSkills {
		if strings.Contains(listSkills[i], skl) && strings.Contains(listSkills[i], "(") {
			result = append(result, listSkills[i])
		}
	}
	return result
}

type CheckManager interface {
	ProbeSimpleCheck(atr string, tn int) float64
}

func (char *character) ProbeSimpleCheck(atr string, tn int) float64 {
	atrMod := 0
	if isCharacteristic(atr) {
		atrMod = charDM(char.characteristics[atr])
	} else {
		atrMod = char.skills[atr]
		if atrMod < 0 {
			atrMod = -3
		}
	}
	return successOf2d6(atrMod, tn) //Дает вероятность позитивного изхода
}

func isCharacteristic(val string) bool {
	for i := range listCHARS {
		if val == listCHARS[i] {
			return true
		}
	}
	return false
}

func successOf2d6(dm int, tn int) float64 {
	var validRolls int
	for d1 := 1; d1 < 7; d1++ {
		for d2 := 1; d2 < 7; d2++ {
			if d1+d2+dm >= tn {
				validRolls++
			}
		}
	}
	flChance := utils.RoundFloat64(float64(validRolls)/36.0, 3)
	return flChance
}

func (char *character) skillDM(skill string) int {
	for key, val := range char.skills {
		//	fmt.Println("check " + key + " search " + skill)
		if key == skill {
			//		fmt.Println("search result", val)
			if val == -1 {
				//			fmt.Println("result", -3)
				return -3
			}

			return val
		}
	}
	//fmt.Println("search result", -3)
	return -3
}

func (char *character) skillCheck(task *Task) (dice, time, effect int, success bool) {
	chDM := 0
	if task.characteristics != "" {
		chDM = charDM(char.characteristics[task.characteristics])
	}
	skillDM := char.skillDM(task.skill)
	tn := task.difficulty + 8

	r := 0
	switch task.boonBane {
	default:
		r = roll2D()
	case "boon":
		r = roll2DBoon()
	case "bane":
		r = roll2DBane()
	}
	time = utils.RollDice("d6")
	dice = r
	roll := r + chDM + skillDM
	effect = boundEffect(roll - tn)
	if effect >= 0 {
		success = true
	}
	return dice, time, effect, success
}

type Task struct {
	difficulty      int
	characteristics string
	skill           string
	timeFrame       string
	dm              int
	boonBane        string
}

func NewTask() *Task {
	skch := &Task{}
	return skch
}

func (skch *Task) SetParameters(dataParameter ...string) {
	for i := range dataParameter {
		data := dataParameter[i]
		switch data {
		default:
			if listed(data, listSkills) {
				skch.skill = data
			}
			if listed(data, listCHARS) {
				skch.characteristics = data
			}
		case "dif-4":
			skch.difficulty = difficultyVeryEasy
		case "dif-2":
			skch.difficulty = difficultyEasy
		case "dif-0":
			skch.difficulty = difficultyAverage
		case "dif2":
			skch.difficulty = difficultyHard
		case "dif4":
			skch.difficulty = difficultyVeryHard
		case "dif6":
			skch.difficulty = difficultyFormidable
		case "boon":
			skch.boonBane = "boon"
		case "bane":
			skch.boonBane = "bane"
		case tFrameSeconds:
			skch.timeFrame = tFrameSeconds
		}
	}
}

func min(array ...int) int {
	var min int
	if len(array) > 0 {
		min = array[0]
	}
	for i := range array {
		if min > array[i] {
			min = array[i]
		}
	}
	return min
}

func max(array ...int) int {
	var max int
	if len(array) > 0 {
		max = array[0]
	}
	for i := range array {
		if max < array[i] {
			max = array[i]
		}
	}
	return max
}

func roll2D() int {
	d1 := utils.RollDice("d6")
	d2 := utils.RollDice("d6")
	return d1 + d2
}

func roll2DBoon() int {
	d1 := utils.RollDice("d6")
	d2 := utils.RollDice("d6")
	d3 := utils.RollDice("d6")
	min := min(d1, d2, d3)
	return d1 + d2 + d3 - min
}

func roll2DBane() int {
	d1 := utils.RollDice("d6")
	d2 := utils.RollDice("d6")
	d3 := utils.RollDice("d6")
	max := max(d1, d2, d3)
	return d1 + d2 + d3 - max
}

func listed(arg string, list []string) bool {
	for i := range list {
		if list[i] == arg {
			return true
		}
	}
	return false
}

func setBounds(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}

func boundEffect(i int) int {
	return setBounds(i, -6, 6)
}
