package entity

import (
	"errors"
	"fmt"
	"strings"
)

const (
	skillCodeTravellerAdmin                  = "1-11-0"
	skillCodeTravellerAdvocate               = "1-12-0"
	skillCodeTravellerAnimalsHandling        = "1-13-1"
	skillCodeTravellerAnimalsTraining        = "1-13-2"
	skillCodeTravellerAnimalsVeterinary      = "1-13-3"
	skillCodeTravellerArtHolography          = "1-14-1"
	skillCodeTravellerArtInstrument          = "1-14-2"
	skillCodeTravellerArtPerformer           = "1-14-3"
	skillCodeTravellerArtVisualMedia         = "1-14-4"
	skillCodeTravellerArtWrite               = "1-14-5"
	skillCodeTravellerAstrogation            = "1-15-0"
	skillCodeTravellerAthleticsDEX           = "1-16-1"
	skillCodeTravellerAthleticsEND           = "1-16-2"
	skillCodeTravellerAthleticsSTR           = "1-16-3"
	skillCodeTravellerBroker                 = "1-17-0"
	skillCodeTravellerCarouse                = "1-18-0"
	skillCodeTravellerDeception              = "1-19-0"
	skillCodeTravellerDiplomat               = "1-1A-0"
	skillCodeTravellerDriveHovercraft        = "1-1B-1"
	skillCodeTravellerDriveMole              = "1-1B-2"
	skillCodeTravellerDriveTrack             = "1-1B-3"
	skillCodeTravellerDriveWalker            = "1-1B-4"
	skillCodeTravellerDriveWheel             = "1-1B-5"
	skillCodeTravellerElectronicsComms       = "1-1C-1"
	skillCodeTravellerElectronicsComputers   = "1-1C-2"
	skillCodeTravellerElectronicsRemoteops   = "1-1C-3"
	skillCodeTravellerElectronicsSensors     = "1-1C-4"
	skillCodeTravellerEngineerJdrive         = "1-1D-1"
	skillCodeTravellerEngineerLifesupport    = "1-1D-2"
	skillCodeTravellerEngineerMdrive         = "1-1D-3"
	skillCodeTravellerEngineerPower          = "1-1D-4"
	skillCodeTravellerExplosives             = "1-1E-0"
	skillCodeTravellerFlyerAirship           = "1-1F-1"
	skillCodeTravellerFlyerGrav              = "1-1F-2"
	skillCodeTravellerFlyerOrnithopter       = "1-1F-3"
	skillCodeTravellerFlyerRotor             = "1-1F-4"
	skillCodeTravellerFlyerWing              = "1-1F-5"
	skillCodeTravellerGambler                = "1-1G-0"
	skillCodeTravellerGuncombatArchaic       = "1-1J-1"
	skillCodeTravellerGuncombatEnergy        = "1-1J-2"
	skillCodeTravellerGuncombatSlug          = "1-1J-3"
	skillCodeTravellerGunnerCapital          = "1-1H-1"
	skillCodeTravellerGunnerOrtilery         = "1-1H-2"
	skillCodeTravellerGunnerScreen           = "1-1H-3"
	skillCodeTravellerGunnerTurret           = "1-1H-4"
	skillCodeTravellerHeavyweaponArtilery    = "1-1K-1"
	skillCodeTravellerHeavyweaponManportable = "1-1K-2"
	skillCodeTravellerHeavyweaponVehicle     = "1-1K-3"
	skillCodeTravellerInvestigate            = "1-1L-0"
	skillCodeTravellerJackofalltrades        = "1-1M-0"
	skillCodeTravellerLanguageAnglic         = "1-1N-1" //byRace
	skillCodeTravellerLeadership             = "1-1P-0"
	skillCodeTravellerMechanic               = "1-1Q-0"
	skillCodeTravellerMedic                  = "1-1R-0"
	skillCodeTravellerMeleeBlade             = "1-1S-1"
	skillCodeTravellerMeleeBludgeon          = "1-1S-2"
	skillCodeTravellerMeleeNatural           = "1-1S-3"
	skillCodeTravellerMeleeUnarmed           = "1-1S-4"
	skillCodeTravellerNavigation             = "1-1T-0"
	skillCodeTravellerPersuade               = "1-1U-0"
	skillCodeTravellerPilotcapitalships      = "1-1V-1"
	skillCodeTravellerPilotSmallcraft        = "1-1V-2"
	skillCodeTravellerPilotSpacecraft        = "1-1V-3"
	skillCodeTravellerProfession             = "1-1W-1" //byType
	skillCodeTravellerRecon                  = "1-11-0"
	skillCodeTravellerScience                = "1-12-1" //byField
	skillCodeTravellerSeafarerOceanships     = "1-13-1"
	skillCodeTravellerSeafarerPersonal       = "1-13-2"
	skillCodeTravellerSeafarerSail           = "1-13-3"
	skillCodeTravellerSeafarerSubmarine      = "1-13-4"
	skillCodeTravellerSteward                = "1-14-0"
	skillCodeTravellerStreetwise             = "1-15-0"
	skillCodeTravellerSurvival               = "1-16-0"
	skillCodeTravellerTacticsMilitary        = "1-17-1"
	skillCodeTravellerTacticsNavy            = "1-17-2"
	skillCodeTravellerVaccsuit               = "1-18-0"
)

//skill -
type skill struct {
	entity      string //low priority
	group       string
	speciality  string
	description string //low priority
	value       int
}

func newSkill(entity, group, speciality string) skill {
	sk := skill{}
	sk.entity = entity
	sk.group = group
	sk.speciality = speciality
	sk.value = 0
	return sk
}

func setValue(s skill, newVal int) skill {
	s.value = newVal
	return s
}

type Skill interface {
	Set(string, int)
	Train(string)
	Remove(string)
	//DM(string) int - часть интерфейса TaskAsset
}

//SkillMap - объект на экспорт именно с ним должны работать внешние библиотеки
//носитель для интерфейса Skill
type SkillMap struct {
	skm map[string]skill
}

func NewSkillMap() *SkillMap {
	sm := SkillMap{}
	sm.skm = make(map[string]skill)
	return &sm
}

//Set - Устанавливает значение скила равное val
// и удостоверяется что вся группа имеет хотябы 0
func (sm *SkillMap) Set(skillCode string, val int) {
	ent, grp, spc, err := disassebleCode(skillCode)
	if err != nil {
		return
	}
	if _, ok := sm.skm[skillCode]; !ok {
		newskill := newSkill(ent, grp, spc)
		sm.skm[skillCode] = newskill
	}
	sm.skm[skillCode] = setValue(sm.skm[skillCode], val)
	sm.ensureByCodeDangerous("-"+grp+"-", 0)
}

//Train - увеличивает значение скила на 1
// и удостоверяется что вся группа имеет хотябы 0
func (sm *SkillMap) Train(skillCode string) {
	ent, grp, spc, err := disassebleCode(skillCode)
	if err != nil {
		return
	}
	if _, ok := sm.skm[skillCode]; !ok {
		newskill := newSkill(ent, grp, spc)
		sm.skm[skillCode] = newskill
	}
	sm.skm[skillCode] = setValue(sm.skm[skillCode], sm.skm[skillCode].value+1)
	sm.ensureByCodeDangerous("-"+grp+"-", 0)
}

//Remove - Удаляет запись о навыке - в теории эта функция вообще не должна использоваться
func (sm *SkillMap) Remove(skillCode string) {
	delete(sm.skm, skillCode)
}

func (sm *SkillMap) ensureByCodeDangerous(codeSample string, val int) {
	for _, code := range skillCodesList() {
		if !strings.Contains(code, codeSample) {
			continue
		}
		if _, ok := sm.skm[code]; !ok {
			ent, grp, spc, _ := disassebleCode(code)
			newskill := newSkill(ent, grp, spc)
			sm.skm[code] = newskill
		}
		if sm.skm[code].value < val {
			sm.skm[code] = setValue(sm.skm[code], val)
		}
	}
}

/*
Каждый скил должен транслироваться в код
в коде должно быть зафиксировано:
-специализация
-группа
-валидный тип сущности (не уверен)
*/

func disassebleCode(skillCode string) (string, string, string, error) {
	data := strings.Split(skillCode, "-")
	if len(data) != 3 {
		return "", "", "", errors.New("skillCode '" + skillCode + "' unreadable")
	}
	return data[0], data[1], data[2], nil
}

func skillCodesList() []string {
	return []string{
		skillCodeTravellerAdmin,
		skillCodeTravellerAdvocate,
		skillCodeTravellerAnimalsHandling,
		skillCodeTravellerAnimalsTraining,
		skillCodeTravellerAnimalsVeterinary,
		skillCodeTravellerArtHolography,
		skillCodeTravellerArtInstrument,
		skillCodeTravellerArtPerformer,
		skillCodeTravellerArtVisualMedia,
		skillCodeTravellerArtWrite,
		skillCodeTravellerAstrogation,
		skillCodeTravellerAthleticsDEX,
		skillCodeTravellerAthleticsEND,
		skillCodeTravellerAthleticsSTR,
		skillCodeTravellerBroker,
		skillCodeTravellerCarouse,
		skillCodeTravellerDeception,
		skillCodeTravellerDiplomat,
		skillCodeTravellerDriveHovercraft,
		skillCodeTravellerDriveMole,
		skillCodeTravellerDriveTrack,
		skillCodeTravellerDriveWalker,
		skillCodeTravellerDriveWheel,
		skillCodeTravellerElectronicsComms,
		skillCodeTravellerElectronicsComputers,
		skillCodeTravellerElectronicsRemoteops,
		skillCodeTravellerElectronicsSensors,
		skillCodeTravellerEngineerJdrive,
		skillCodeTravellerEngineerLifesupport,
		skillCodeTravellerEngineerMdrive,
		skillCodeTravellerEngineerPower,
		skillCodeTravellerExplosives,
		skillCodeTravellerFlyerAirship,
		skillCodeTravellerFlyerGrav,
		skillCodeTravellerFlyerOrnithopter,
		skillCodeTravellerFlyerRotor,
		skillCodeTravellerFlyerWing,
		skillCodeTravellerGambler,
		skillCodeTravellerGuncombatArchaic,
		skillCodeTravellerGuncombatEnergy,
		skillCodeTravellerGuncombatSlug,
		skillCodeTravellerGunnerCapital,
		skillCodeTravellerGunnerOrtilery,
		skillCodeTravellerGunnerScreen,
		skillCodeTravellerGunnerTurret,
		skillCodeTravellerHeavyweaponArtilery,
		skillCodeTravellerHeavyweaponManportable,
		skillCodeTravellerHeavyweaponVehicle,
		skillCodeTravellerInvestigate,
		skillCodeTravellerJackofalltrades,
		skillCodeTravellerLanguageAnglic,
		skillCodeTravellerLeadership,
		skillCodeTravellerMechanic,
		skillCodeTravellerMedic,
		skillCodeTravellerMeleeBlade,
		skillCodeTravellerMeleeBludgeon,
		skillCodeTravellerMeleeNatural,
		skillCodeTravellerMeleeUnarmed,
		skillCodeTravellerNavigation,
		skillCodeTravellerPersuade,
		skillCodeTravellerPilotcapitalships,
		skillCodeTravellerPilotSmallcraft,
		skillCodeTravellerPilotSpacecraft,
		skillCodeTravellerProfession,
		skillCodeTravellerRecon,
		skillCodeTravellerScience,
		skillCodeTravellerSeafarerOceanships,
		skillCodeTravellerSeafarerPersonal,
		skillCodeTravellerSeafarerSail,
		skillCodeTravellerSeafarerSubmarine,
		skillCodeTravellerSteward,
		skillCodeTravellerStreetwise,
		skillCodeTravellerSurvival,
		skillCodeTravellerTacticsMilitary,
		skillCodeTravellerTacticsNavy,
		skillCodeTravellerVaccsuit,
	}
}

func Test() {
	fmt.Println("Start test")
	sklMap := NewSkillMap()
	sklMap.Train(skillCodeTravellerGuncombatEnergy)
	sklMap.Train(skillCodeTravellerGuncombatEnergy)
	sklMap.Train(skillCodeTravellerGuncombatEnergy)
	sklMap.Train(skillCodeTravellerGuncombatEnergy)
	sklMap.Train(skillCodeTravellerGuncombatEnergy)
	sklMap.Set(skillCodeTravellerGuncombatSlug, 1)
	sklMap.Remove(skillCodeTravellerGuncombatSlug)
	fmt.Println(sklMap)
	fmt.Println("End test")
}
