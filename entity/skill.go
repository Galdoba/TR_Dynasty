package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
)

const (
	SCEntity                    = 0
	SCGroup                     = 1
	SCSpeciality                = 2
	SCDescription               = 3
	SCTrvAdmin                  = "1=11=0=Admin"
	SCTrvAdvocate               = "1=12=0=Advocate"
	SCTrvAnimals                = "1=13=0=Animals"
	SCTrvAnimalsHandling        = "1=13=1=Animals (Handling)"
	SCTrvAnimalsTraining        = "1=13=2=Animals (Training)"
	SCTrvAnimalsVeterinary      = "1=13=3=Animals (Veterinary)"
	SCTrvArt                    = "1=14=0=Art"
	SCTrvArtHolography          = "1=14=1=Art (Holography)"
	SCTrvArtInstrument          = "1=14=2=Art (Instrument)"
	SCTrvArtPerformer           = "1=14=3=Art (Performer)"
	SCTrvArtVisualMedia         = "1=14=4=Art (Visual Media)"
	SCTrvArtWrite               = "1=14=5=Art (Write)"
	SCTrvAstrogation            = "1=15=0=Astrogation"
	SCTrvAthletics              = "1=16=0=Athletics"
	SCTrvAthleticsDEX           = "1=16=1=Athletics (DEX)"
	SCTrvAthleticsEND           = "1=16=2=Athletics (END)"
	SCTrvAthleticsSTR           = "1=16=3=Athletics (STR)"
	SCTrvBroker                 = "1=17=0=Broker"
	SCTrvCarouse                = "1=18=0=Carouse"
	SCTrvDeception              = "1=19=0=Deception"
	SCTrvDiplomat               = "1=1A=0=Diplomat"
	SCTrvDrive                  = "1=1B=0=Drive"
	SCTrvDriveHovercraft        = "1=1B=1=Drive (Hovercraft)"
	SCTrvDriveMole              = "1=1B=2=Drive (Mole)"
	SCTrvDriveTrack             = "1=1B=3=Drive (Track)"
	SCTrvDriveWalker            = "1=1B=4=Drive (Walker)"
	SCTrvDriveWheel             = "1=1B=5=Drive (Wheel)"
	SCTrvElectronics            = "1=1C=0=Electronics"
	SCTrvElectronicsComms       = "1=1C=1=Electronics (Comms)"
	SCTrvElectronicsComputers   = "1=1C=2=Electronics (Computers)"
	SCTrvElectronicsRemoteops   = "1=1C=3=Electronics (Remote Ops)"
	SCTrvElectronicsSensors     = "1=1C=4=Electronics (Sensors)"
	SCTrvEngineer               = "1=1D=0=Engineer"
	SCTrvEngineerJdrive         = "1=1D=1=Engineer (J-drive)"
	SCTrvEngineerLifesupport    = "1=1D=2=Engineer (Life Support)"
	SCTrvEngineerMdrive         = "1=1D=3=Engineer (M-drive)"
	SCTrvEngineerPower          = "1=1D=4=Engineer (Power)"
	SCTrvExplosives             = "1=1E=0=Explosives"
	SCTrvFlyer                  = "1=1F=0=Flyer"
	SCTrvFlyerAirship           = "1=1F=1=Flyer (Airship)"
	SCTrvFlyerGrav              = "1=1F=2=Flyer (Grav)"
	SCTrvFlyerOrnithopter       = "1=1F=3=Flyer (Ornithopter)"
	SCTrvFlyerRotor             = "1=1F=4=Flyer (Rotor)"
	SCTrvFlyerWing              = "1=1F=5=Flyer (Wing)"
	SCTrvGambler                = "1=1G=0=Gambler"
	SCTrvGuncombat              = "1=1J=0=Guncombat"
	SCTrvGuncombatArchaic       = "1=1J=1=Guncombat (Archaic)"
	SCTrvGuncombatEnergy        = "1=1J=2=Guncombat (Energy)"
	SCTrvGuncombatSlug          = "1=1J=3=Guncombat (Slug)"
	SCTrvGunner                 = "1=1H=0=Gunner"
	SCTrvGunnerCapital          = "1=1H=1=Gunner (Capital)"
	SCTrvGunnerOrtilery         = "1=1H=2=Gunner (Ortilery)"
	SCTrvGunnerScreen           = "1=1H=3=Gunner (Screen)"
	SCTrvGunnerTurret           = "1=1H=4=Gunner (Turret)"
	SCTrvHeavyweapon            = "1=1K=0=Heavyweapon"
	SCTrvHeavyweaponArtilery    = "1=1K=1=Heavyweapon (Artilery)"
	SCTrvHeavyweaponManportable = "1=1K=2=Heavyweapon (Man Portable)"
	SCTrvHeavyweaponVehicle     = "1=1K=3=Heavyweapon (Vehicle)"
	SCTrvInvestigate            = "1=1L=0=Investigate"
	SCTrvJackofalltrades        = "1=1M=0=Jack-of-all-Trades"
	SCTrvLanguage               = "1=1N=0=Language"          //byRac-LanguageAnglic         e
	SCTrvLanguageAnglic         = "1=1N=1=Language (Anglic)" //byRac-LanguageAnglic         e
	SCTrvLeadership             = "1=1P=0=Leadership"
	SCTrvMechanic               = "1=1Q=0=Mechanic"
	SCTrvMedic                  = "1=1R=0=Medic"
	SCTrvMelee                  = "1=1S=0=Melee"
	SCTrvMeleeBlade             = "1=1S=1=Melee (Blade)"
	SCTrvMeleeBludgeon          = "1=1S=2=Melee (Bludgeon)"
	SCTrvMeleeNatural           = "1=1S=3=Melee (Natural)"
	SCTrvMeleeUnarmed           = "1=1S=4=Melee (Unarmed)"
	SCTrvNavigation             = "1=1T=0=Navigation"
	SCTrvPersuade               = "1=1U=0=Persuade"
	SCTrvPilot                  = "1=1V=0=Pilot"
	SCTrvPilotCapitalships      = "1=1V=1=Pilot (Capital Ships)"
	SCTrvPilotSmallcraft        = "1=1V=2=Pilot (Small craft)"
	SCTrvPilotSpacecraft        = "1=1V=3=Pilot (Spacecraft)"
	SCTrvProfession             = "1=1W=0=Profession" //byTyp-Profession             e
	SCTrvProfessionAny          = "1=1W=1=Profession (Any)"
	SCTrvRecon                  = "1=21=0=Recon"
	SCTrvSciencePhysical        = "1=22=0=Science" //byField Science                d
	SCTrvScienceLife            = "1=23=0=Science" //byField Science                d
	SCTrvScienceSocial          = "1=24=0=Science" //byField Science                d
	SCTrvScienceSpace           = "1=25=0=Science" //byField Science                d
	SCTrvScienceAny             = "1=26=1=Science (TODO)"
	SCTrvSeafarer               = "1=27=0=Seafarer"
	SCTrvSeafarerOceanships     = "1=27=1=Seafarer (Ocean Ships)"
	SCTrvSeafarerPersonal       = "1=27=2=Seafarer (Personal)"
	SCTrvSeafarerSail           = "1=27=3=Seafarer (Sail)"
	SCTrvSeafarerSubmarine      = "1=27=4=Seafarer (Submarine)"
	SCTrvStealth                = "1=28=0=Stealth"
	SCTrvSteward                = "1=29=0=Steward"
	SCTrvStreetwise             = "1=2A=0=Streetwise"
	SCTrvSurvival               = "1=2B=0=Survival"
	SCTrvTactics                = "1=2C=0=Tactics"
	SCTrvTacticsMilitary        = "1=2C=1=Tactics (Military)"
	SCTrvTacticsNavy            = "1=2C=2=Tactics (Navy)"
	SCTrvVaccsuit               = "1=2D=0=Vaccsuit"
)

//skill -
type skill struct {
	entity      string //low priority
	group       string
	speciality  string
	description string
	value       int
}

func newSkill(skillCode string) skill {
	sk := skill{}
	sk.entity = GetFromCode(SCEntity, skillCode)
	sk.group = GetFromCode(SCGroup, skillCode)
	sk.speciality = GetFromCode(SCSpeciality, skillCode)
	sk.description = GetFromCode(SCDescription, skillCode)
	sk.value = 0
	return sk
}

// func setValue(s skill, newVal int) skill {
// 	s.value = newVal
// 	return s
// }

func (s *skill) setValue(newVal int) {
	s.value = newVal
}

type Parameter interface {
	Set(string, int)
	GetValue(string) (int, error)
	Train(string)
	Remove(string)
}

type Skill interface {
	Parameter
	Get(string) skill
	FPrintSkillGroupS(string) string
	//DM(string) int - часть интерфейса TaskAsset
}

func (s skill) String() string {
	if s.speciality != "0" && s.value == 0 {
		return ""
	}
	return s.description + " " + strconv.Itoa(s.value)
}

func (sm *SkillMap) Get(skillName string) skill {
	if val, ok := sm.skm[skillName]; ok {
		return val
	}
	return skill{}
}

func (sm *SkillMap) FPrintSkillGroupS(skillName string) string {
	if GetFromCode(SCSpeciality, skillName) != "0" {
		descr := GetFromCode(SCDescription, skillName)
		str := ""
		if val, ok := sm.skm[skillName]; ok {
			if val.value > 0 {
				str += descr + " " + strconv.Itoa(val.value)
			}
		}
		return str
	}
	str := ""
	grp := GetFromCode(SCGroup, skillName)
	codes := codesByGroup(grp)
	rem := ""
	for i, cCode := range codes {
		spc := GetFromCode(SCSpeciality, cCode)
		if val, ok := sm.skm[codes[i]]; ok {
			if spc == "0" {
				rem = val.description
				str += GetFromCode(SCDescription, codes[i]) + " " + strconv.Itoa(val.value) + ", "
			}
			if val.value > 0 && spc != "0" {
				str += GetFromCode(SCDescription, codes[i]) + " " + strconv.Itoa(val.value) + ", "
			}

		}
	}
	str = strings.TrimSuffix(str, ", ")
	str = strings.TrimPrefix(str, rem+" 0, ")
	return str
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
	groupeCodes := codesByGroup(GetFromCode(SCGroup, skillCode))
	if _, ok := sm.skm[skillCode]; !ok { //Если такого скила нет - создаем всю группу со значением 0
		for _, code := range groupeCodes {
			sm.skm[code] = newSkill(code)
		}
	}
	skl := newSkill(skillCode)
	skl.setValue(val)
	sm.skm[skillCode] = skl
}

func (sm *SkillMap) GetValue(code string) (int, error) {
	if val, ok := sm.skm[code]; ok {
		return val.value, nil
	}
	return 0, errors.New("No Value for '" + code + "'")
}

//Train - увеличивает значение скила на 1
// и удостоверяется что вся группа имеет хотябы 0
func (sm *SkillMap) Train(skillCode string) {
	groupeCodes := codesByGroup(GetFromCode(SCGroup, skillCode))
	if _, ok := sm.skm[skillCode]; !ok { //Если такого скила нет - создаем всю группу со значением 0
		for _, code := range groupeCodes {
			sm.skm[code] = newSkill(code)
		}
		return
	} //мы точно знаем что скилл есть
	if GetFromCode(SCSpeciality, skillCode) != "0" { //если это не общая группа,
		sm.Set(skillCode, sm.Get(skillCode).value+1) //то просто увеличиваем на 1
		return
	}
	spCode := distributeSpeciality(sm, groupeCodes)
	sm.Set(spCode, sm.Get(spCode).value+1)
}

func distributeSpeciality(sm *SkillMap, groupeCodes []string) string {
	if len(groupeCodes) == 1 {
		return groupeCodes[0]
	}
	array := []string{}
	for i := 0; i < 5; i++ {
		for g, code := range groupeCodes {
			if g == 0 {
				continue
			}
			skl := sm.Get(code).value
			if skl >= i {
				array = append(array, code)
			}
		}
	}
	l := len(array)
	d := dice.Roll("1d" + strconv.Itoa(l)).DM(-1).Sum()
	return array[d]
}

//Remove - Удаляет запись о навыке - в теории эта функция вообще не должна использоваться
func (sm *SkillMap) Remove(skillCode string) {
	delete(sm.skm, skillCode)
}

/*
Каждый скил должен транслироваться в код
в коде должно быть зафиксировано:
-специализация
-группа
-валидный тип сущности (не уверен)
*/

func disassebleCode(skillCode string) (string, string, string, string, error) {
	data := strings.Split(skillCode, "=")
	if len(data) != 4 {
		return "", "", "", "", errors.New("skillCode '" + skillCode + "' unreadable")
	}
	return data[0], data[1], data[2], data[3], nil
}

func GetFromCode(entry int, code string) string {
	data := strings.Split(code, "=")
	// if len(data) != 4 {
	// 	return ""
	// }
	return data[entry]
}

func codesByGroup(grp string) []string {
	groupeCodes := []string{}
	for _, code := range SkillCodesList() {
		if GetFromCode(SCGroup, code) == grp {
			groupeCodes = append(groupeCodes, code)
		}
	}
	return groupeCodes
}

func SkillCodesList() []string {
	return []string{
		SCTrvAdmin,
		SCTrvAdvocate,
		SCTrvAnimals,
		SCTrvAnimalsHandling,
		SCTrvAnimalsTraining,
		SCTrvAnimalsVeterinary,
		SCTrvArt,
		SCTrvArtHolography,
		SCTrvArtInstrument,
		SCTrvArtPerformer,
		SCTrvArtVisualMedia,
		SCTrvArtWrite,
		SCTrvAstrogation,
		SCTrvAthletics,
		SCTrvAthleticsDEX,
		SCTrvAthleticsEND,
		SCTrvAthleticsSTR,
		SCTrvBroker,
		SCTrvCarouse,
		SCTrvDeception,
		SCTrvDiplomat,
		SCTrvDrive,
		SCTrvDriveHovercraft,
		SCTrvDriveMole,
		SCTrvDriveTrack,
		SCTrvDriveWalker,
		SCTrvDriveWheel,
		SCTrvElectronics,
		SCTrvElectronicsComms,
		SCTrvElectronicsComputers,
		SCTrvElectronicsRemoteops,
		SCTrvElectronicsSensors,
		SCTrvEngineer,
		SCTrvEngineerJdrive,
		SCTrvEngineerLifesupport,
		SCTrvEngineerMdrive,
		SCTrvEngineerPower,
		SCTrvExplosives,
		SCTrvFlyer,
		SCTrvFlyerAirship,
		SCTrvFlyerGrav,
		SCTrvFlyerOrnithopter,
		SCTrvFlyerRotor,
		SCTrvFlyerWing,
		SCTrvGambler,
		SCTrvGuncombat,
		SCTrvGuncombatArchaic,
		SCTrvGuncombatEnergy,
		SCTrvGuncombatSlug,
		SCTrvGunner,
		SCTrvGunnerCapital,
		SCTrvGunnerOrtilery,
		SCTrvGunnerScreen,
		SCTrvGunnerTurret,
		SCTrvHeavyweapon,
		SCTrvHeavyweaponArtilery,
		SCTrvHeavyweaponManportable,
		SCTrvHeavyweaponVehicle,
		SCTrvInvestigate,
		SCTrvJackofalltrades,
		SCTrvLanguage,
		SCTrvLanguageAnglic,
		SCTrvLeadership,
		SCTrvMechanic,
		SCTrvMedic,
		SCTrvMelee,
		SCTrvMeleeBlade,
		SCTrvMeleeBludgeon,
		SCTrvMeleeNatural,
		SCTrvMeleeUnarmed,
		SCTrvNavigation,
		SCTrvPersuade,
		SCTrvPilot,
		SCTrvPilotCapitalships,
		SCTrvPilotSmallcraft,
		SCTrvPilotSpacecraft,
		SCTrvProfession,
		SCTrvProfessionAny,
		SCTrvRecon,
		SCTrvSciencePhysical,
		SCTrvScienceLife,
		SCTrvScienceSocial,
		SCTrvScienceSpace,
		SCTrvScienceAny,
		SCTrvSeafarer,
		SCTrvSeafarerOceanships,
		SCTrvSeafarerPersonal,
		SCTrvSeafarerSail,
		SCTrvSeafarerSubmarine,
		SCTrvStealth,
		SCTrvSteward,
		SCTrvStreetwise,
		SCTrvSurvival,
		SCTrvTactics,
		SCTrvTacticsMilitary,
		SCTrvTacticsNavy,
		SCTrvVaccsuit,
	}
}

func Test() {
	fmt.Println("Start test")
	sklMap := NewSkillMap()
	sklMap.Train(SCTrvGuncombatEnergy)
	sklMap.Train(SCTrvGuncombatEnergy)
	sklMap.Train(SCTrvGuncombatEnergy)
	sklMap.Train(SCTrvGuncombatEnergy)
	sklMap.Train(SCTrvGuncombatEnergy)
	sklMap.Set(SCTrvGuncombatSlug, 1)
	sklMap.Remove(SCTrvGuncombatSlug)
	fmt.Println(sklMap)
	fmt.Println("End test")
}
