package entity

import (
	"errors"
	"strings"
)

const (
	skillCodeAdmin                  = "1-11-0"
	skillCodeAdvocate               = "1-12-0"
	skillCodeAnimalsHandling        = "1-13-1"
	skillCodeAnimalsTraining        = "1-13-2"
	skillCodeAnimalsVeterinary      = "1-13-3"
	skillCodeArtHolography          = "1-14-1"
	skillCodeArtInstrument          = "1-14-2"
	skillCodeArtPerformer           = "1-14-3"
	skillCodeArtVisualMedia         = "1-14-4"
	skillCodeArtWrite               = "1-14-5"
	skillCodeAstrogation            = "1-15-0"
	skillCodeAthleticsDEX           = "1-16-1"
	skillCodeAthleticsEND           = "1-16-2"
	skillCodeAthleticsSTR           = "1-16-3"
	skillCodeBroker                 = "1-17-0"
	skillCodeCarouse                = "1-18-0"
	skillCodeDeception              = "1-19-0"
	skillCodeDiplomat               = "1-1A-0"
	skillCodeDriveHovercraft        = "1-1B-1"
	skillCodeDriveMole              = "1-1B-2"
	skillCodeDriveTrack             = "1-1B-3"
	skillCodeDriveWalker            = "1-1B-4"
	skillCodeDriveWheel             = "1-1B-5"
	skillCodeElectronicsComms       = "1-1C-1"
	skillCodeElectronicsComputers   = "1-1C-2"
	skillCodeElectronicsRemoteops   = "1-1C-3"
	skillCodeElectronicsSensors     = "1-1C-4"
	skillCodeEngineerJdrive         = "1-1D-1"
	skillCodeEngineerLifesupport    = "1-1D-2"
	skillCodeEngineerMdrive         = "1-1D-3"
	skillCodeEngineerPower          = "1-1D-4"
	skillCodeExplosives             = "1-1E-0"
	skillCodeFlyerAirship           = "1-1F-1"
	skillCodeFlyerGrav              = "1-1F-2"
	skillCodeFlyerOrnithopter       = "1-1F-3"
	skillCodeFlyerRotor             = "1-1F-4"
	skillCodeFlyerWing              = "1-1F-5"
	skillCodeGambler                = "1-1G-0"
	skillCodeGuncombatArchaic       = "1-1J-1"
	skillCodeGuncombatEnergy        = "1-1J-2"
	skillCodeGuncombatSlug          = "1-1J-3"
	skillCodeGunnerCapital          = "1-1H-1"
	skillCodeGunnerOrtilery         = "1-1H-2"
	skillCodeGunnerScreen           = "1-1H-3"
	skillCodeGunnerTurret           = "1-1H-4"
	skillCodeHeavyweaponArtilery    = "1-1K-1"
	skillCodeHeavyweaponManportable = "1-1K-2"
	skillCodeHeavyweaponVehicle     = "1-1K-3"
	skillCodeInvestigate            = "1-1L-0"
	skillCodeJackofalltrades        = "1-1M-0"
	skillCodeLanguageAnglic         = "1-1N-1" //byRace
	skillCodeLeadership             = "1-1P-0"
	skillCodeMechanic               = "1-1Q-0"
	skillCodeMedic                  = "1-1R-0"
	skillCodeMeleeBlade             = "1-1S-1"
	skillCodeMeleeBludgeon          = "1-1S-2"
	skillCodeMeleeNatural           = "1-1S-3"
	skillCodeMeleeUnarmed           = "1-1S-4"
	skillCodeNavigation             = "1-1T-0"
	skillCodePersuade               = "1-1U-0"
	skillCodePilotcapitalships      = "1-1V-1"
	skillCodePilotSmallcraft        = "1-1V-2"
	skillCodePilotSpacecraft        = "1-1V-3"
	skillCodeProfession             = "1-1W-1" //byType
	skillCodeRecon                  = "1-11-0"
	skillCodeScience                = "1-12-1" //byField
	skillCodeSeafarerOceanships     = "1-13-1"
	skillCodeSeafarerPersonal       = "1-13-2"
	skillCodeSeafarerSail           = "1-13-3"
	skillCodeSeafarerSubmarine      = "1-13-4"
	skillCodeSteward                = "1-14-0"
	skillCodeStreetwise             = "1-15-0"
	skillCodeSurvival               = "1-16-0"
	skillCodeTacticsMilitary        = "1-17-1"
	skillCodeTacticsNavy            = "1-17-2"
	skillCodeVaccsuit               = "1-18-0"
)

//skill -
type skill struct {
	entity      string
	group       string
	specialitie string
	description string //low priority
	value       int
}

type Skill interface {
	SetSkill(string, int)
	TrainSkill(string)
	RemoveSkill(string)
	//DM(string) int - часть интерфейса TaskAsset
}

//SkillMap - объект на экспорт именно с ним должны работать внешние библиотеки
//носитель для интерфейса Skill
type SkillMap struct {
	skm map[string]skill
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

/*

 */
