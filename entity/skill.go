package entity

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
)

type skill struct {
	group       string
	specialitie string
	value       int
}

func (s skill) String() string {
	output := ""
	output += s.group + " "
	if s.specialitie != "" {
		output += "(" + s.specialitie + ") "
	}
	if s.value >= 0 {
		output += strconv.Itoa(s.value)
	} else {
		output += "-"
	}
	return output
}

func newSkill(group, specialitie string) skill {
	s := skill{}
	s.group = group
	s.specialitie = specialitie
	s.value = 0
	return s
}

func Test() {
	fmt.Println("Start test")
	fmt.Println("End test")
}

//////////////////////////////////////////////////////////////////////
var nullSkillMap map[string]skill

func init() {
	nullSkillMap = make(map[string]skill)
	for _, groupe := range skillGroupsMasterList() {
		specs := validSpecialities(groupe)
		if len(specs) == 0 {
			specs = append(specs, "")
			nullSkillMap[groupe] = newSkill(groupe, "")
			continue
		}
		for _, spc := range specs {
			skillName := groupe + " (" + spc + ")"
			nullSkillMap[skillName] = newSkill(groupe, spc)
		}
	}
	for k, v := range nullSkillMap {
		fmt.Println(k, ": ", v)
	}
}

func skillGroupsMasterList() []string {
	mList := []string{
		constant.SKILLgroupeAdmin,
		constant.SKILLgroupeAdvovate,
		constant.SKILLgroupeAnimals,
		constant.SKILLgroupeArt,
		constant.SKILLgroupeAstrogation,
		constant.SKILLgroupeAthletics,
		constant.SKILLgroupeBroker,
		constant.SKILLgroupeCarouse,
		constant.SKILLgroupeDeception,
		constant.SKILLgroupeDiplomat,
		constant.SKILLgroupeDrive,
		constant.SKILLgroupeElectronics,
		constant.SKILLgroupeEngineer,
		constant.SKILLgroupeExplosives,
		constant.SKILLgroupeFlyer,
		constant.SKILLgroupeGambler,
		constant.SKILLgroupeGunner,
		constant.SKILLgroupeGuncombat,
		constant.SKILLgroupeHeavyweapons,
		constant.SKILLgroupeInvestigate,
		constant.SKILLgroupeJackofalltrades,
		constant.SKILLgroupeLanguage,
		constant.SKILLgroupeLeadership,
		constant.SKILLgroupeMechanic,
		constant.SKILLgroupeMedic,
		constant.SKILLgroupeMelee,
		constant.SKILLgroupeNavigation,
		constant.SKILLgroupePersuade,
		constant.SKILLgroupePilot,
		constant.SKILLgroupeProfession,
		constant.SKILLgroupeRecon,
		constant.SKILLgroupeScience,
		constant.SKILLgroupeSeafarer,
		constant.SKILLgroupeStealth,
		constant.SKILLgroupeSteward,
		constant.SKILLgroupeStreetwise,
		constant.SKILLgroupeSurvival,
		constant.SKILLgroupeTactics,
		constant.SKILLgroupeVaccsuit,
	}
	return mList
}

func skillSpecialitiesMasterList() []string {
	mList := []string{
		constant.SKILLSpecialityANY,
		constant.SKILLSpecialityNULL,
		constant.SKILLSpecialityHandling,
		constant.SKILLSpecialityVeterinary,
		constant.SKILLSpecialityTraining,
		constant.SKILLSpecialityPerformer,
		constant.SKILLSpecialityHolography,
		constant.SKILLSpecialityInstrument,
		constant.SKILLSpecialityVisualmedia,
		constant.SKILLSpecialityWrite,
		constant.SKILLSpecialityDexterity,
		constant.SKILLSpecialityEndurance,
		constant.SKILLSpecialityStrength,
		constant.SKILLSpecialityHovercraft,
		constant.SKILLSpecialityMole,
		constant.SKILLSpecialityTrack,
		constant.SKILLSpecialityWalker,
		constant.SKILLSpecialityWheel,
		constant.SKILLSpecialityComms,
		constant.SKILLSpecialityComputers,
		constant.SKILLSpecialityRemoteops,
		constant.SKILLSpecialitySensors,
		constant.SKILLSpecialityMdrive,
		constant.SKILLSpecialityJdrive,
		constant.SKILLSpecialityLifesupport,
		constant.SKILLSpecialityPower,
		constant.SKILLSpecialityAirship,
		constant.SKILLSpecialityGrav,
		constant.SKILLSpecialityOrnithopter,
		constant.SKILLSpecialityRotor,
		constant.SKILLSpecialityWing,
		constant.SKILLSpecialityTurret,
		constant.SKILLSpecialityOrtilery,
		constant.SKILLSpecialityScreen,
		constant.SKILLSpecialityCapital,
		constant.SKILLSpecialityArchaic,
		constant.SKILLSpecialityEnergy,
		constant.SKILLSpecialitySlug,
		constant.SKILLSpecialityArtilery,
		constant.SKILLSpecialityManportable,
		constant.SKILLSpecialityVehicle,
		constant.SKILLSpecialityAnglic,
		constant.SKILLSpecialityVilani,
		constant.SKILLSpecialityZdetl,
		constant.SKILLSpecialityOynprith,
		constant.SKILLSpecialityUnarmed,
		constant.SKILLSpecialityBlade,
		constant.SKILLSpecialityBludgeon,
		constant.SKILLSpecialityNatural,
		constant.SKILLSpecialitySmallcraft,
		constant.SKILLSpecialitySpacecraft,
		constant.SKILLSpecialityCapitalships,
		constant.SKILLSpecialityBelter,
		constant.SKILLSpecialityBiologicals,
		constant.SKILLSpecialityArchaeology,
		constant.SKILLSpecialityAstronomy,
		constant.SKILLSpecialityBiology,
		constant.SKILLSpecialityChemistry,
		constant.SKILLSpecialityCosmology,
		constant.SKILLSpecialityCybernetics,
		constant.SKILLSpecialityEconimics,
		constant.SKILLSpecialityGenetics,
		constant.SKILLSpecialityHistory,
		constant.SKILLSpecialityLinguistics,
		constant.SKILLSpecialityPhilosophy,
		constant.SKILLSpecialityPhysics,
		constant.SKILLSpecialityPlanetology,
		constant.SKILLSpecialityPsionicology,
		constant.SKILLSpecialityPsyhology,
		constant.SKILLSpecialityRobotics,
		constant.SKILLSpecialitySophontology,
		constant.SKILLSpecialityXenology,
		constant.SKILLSpecialityOceanships,
		constant.SKILLSpecialityPersonal,
		constant.SKILLSpecialitySail,
		constant.SKILLSpecialitySubmarine,
		constant.SKILLSpecialityMilitary,
		constant.SKILLSpecialityNavy,
	}
	return mList
}

func validSpecialities(skGroup string) []string {
	valid := []string{}
	switch skGroup {
	default:
		return valid
	case constant.SKILLgroupeAnimals:
		valid = []string{
			constant.SKILLSpecialityHandling,
			constant.SKILLSpecialityVeterinary,
			constant.SKILLSpecialityTraining,
		}
	case constant.SKILLgroupeArt:
		valid = []string{
			constant.SKILLSpecialityPerformer,
			constant.SKILLSpecialityHolography,
			constant.SKILLSpecialityInstrument,
			constant.SKILLSpecialityVisualmedia,
			constant.SKILLSpecialityWrite,
		}
	case constant.SKILLgroupeAthletics:
		valid = []string{
			constant.SKILLSpecialityDexterity,
			constant.SKILLSpecialityEndurance,
			constant.SKILLSpecialityStrength,
		}
	case constant.SKILLgroupeDrive:
		valid = []string{
			constant.SKILLSpecialityHovercraft,
			constant.SKILLSpecialityMole,
			constant.SKILLSpecialityTrack,
			constant.SKILLSpecialityWalker,
			constant.SKILLSpecialityWheel,
		}
	case constant.SKILLgroupeElectronics:
		valid = []string{
			constant.SKILLSpecialityComms,
			constant.SKILLSpecialityComputers,
			constant.SKILLSpecialityRemoteops,
			constant.SKILLSpecialitySensors,
		}
	case constant.SKILLgroupeEngineer:
		valid = []string{
			constant.SKILLSpecialityMdrive,
			constant.SKILLSpecialityJdrive,
			constant.SKILLSpecialityLifesupport,
			constant.SKILLSpecialityPower,
		}
	case constant.SKILLgroupeFlyer:
		valid = []string{
			constant.SKILLSpecialityAirship,
			constant.SKILLSpecialityGrav,
			constant.SKILLSpecialityOrnithopter,
			constant.SKILLSpecialityRotor,
			constant.SKILLSpecialityWing,
		}
	case constant.SKILLgroupeGunner:
		valid = []string{
			constant.SKILLSpecialityTurret,
			constant.SKILLSpecialityOrtilery,
			constant.SKILLSpecialityScreen,
			constant.SKILLSpecialityCapital,
		}

	case constant.SKILLgroupeGuncombat:
		valid = []string{
			constant.SKILLSpecialityArchaic,
			constant.SKILLSpecialityEnergy,
			constant.SKILLSpecialitySlug,
		}
	case constant.SKILLgroupeHeavyweapons:
		valid = []string{
			constant.SKILLSpecialityManportable,
			constant.SKILLSpecialityArtilery,
			constant.SKILLSpecialityVehicle,
		}
	case constant.SKILLgroupeLanguage:
		valid = []string{
			constant.SKILLSpecialityAnglic,
			constant.SKILLSpecialityVilani,
			constant.SKILLSpecialityZdetl,
			constant.SKILLSpecialityOynprith,
		}
	case constant.SKILLgroupeMelee:
		valid = []string{
			constant.SKILLSpecialityUnarmed,
			constant.SKILLSpecialityBlade,
			constant.SKILLSpecialityBludgeon,
			constant.SKILLSpecialityNatural,
		}
	case constant.SKILLgroupePilot:
		valid = []string{
			constant.SKILLSpecialitySmallcraft,
			constant.SKILLSpecialitySpacecraft,
			constant.SKILLSpecialityCapitalships,
		}
	case constant.SKILLgroupeProfession:
		valid = []string{
			constant.SKILLSpecialityBelter,
		}
	case constant.SKILLgroupeRecon:

	case constant.SKILLgroupeScience:
		valid = []string{
			constant.SKILLSpecialityArchaeology,
			constant.SKILLSpecialityAstronomy,
			constant.SKILLSpecialityBiology,
			constant.SKILLSpecialityChemistry,
			constant.SKILLSpecialityCosmology,
			constant.SKILLSpecialityCybernetics,
			constant.SKILLSpecialityEconimics,
			constant.SKILLSpecialityGenetics,
			constant.SKILLSpecialityHistory,
			constant.SKILLSpecialityLinguistics,
			constant.SKILLSpecialityPhilosophy,
			constant.SKILLSpecialityPhysics,
			constant.SKILLSpecialityPlanetology,
			constant.SKILLSpecialityPsionicology,
			constant.SKILLSpecialityPsyhology,
			constant.SKILLSpecialityRobotics,
			constant.SKILLSpecialitySophontology,
			constant.SKILLSpecialityXenology,
		}
	case constant.SKILLgroupeSeafarer:
		valid = []string{
			constant.SKILLSpecialityOceanships,
			constant.SKILLSpecialityPersonal,
			constant.SKILLSpecialitySail,
			constant.SKILLSpecialitySubmarine,
		}
	case constant.SKILLgroupeTactics:
		valid = []string{
			constant.SKILLSpecialityMilitary,
			constant.SKILLSpecialityNavy,
		}
	}
	return valid
}

/*
Skill

*/
