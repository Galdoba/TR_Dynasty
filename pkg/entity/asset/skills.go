package asset

import "errors"

const (
	GroupAdmin           = "Admin"
	GroupAdvocate        = "Advocate"
	GroupAnimals         = "Animals"
	GroupArt             = "Art"
	GroupAstrogation     = "Astrogation"
	GroupAthletics       = "Athletics"
	GroupBroker          = "Broker"
	GroupCarouse         = "Carouse"
	GroupDeception       = "Deception"
	GroupDiplomat        = "Diplomat"
	GroupDrive           = "Drive"
	GroupElectronics     = "Electronics"
	GroupEngineer        = "Engineer"
	GroupExplosives      = "Explosives"
	GroupFlyer           = "Flyer"
	GroupGambler         = "Gambler"
	GroupGuncombat       = "Gun Combat"
	GroupGunner          = "Gunner"
	GroupHeavyweapon     = "Heavy Weapon"
	GroupInvestigate     = "Investigate"
	GroupJackofalltrades = "Jack-of-all-Trades"
	GroupLanguage        = "Language" //byRac-LanguageAnglic         e
	GroupLeadership      = "Leadership"
	GroupMechanic        = "Mechanic"
	GroupMedic           = "Medic"
	GroupMelee           = "Melee"
	GroupNavigation      = "Navigation"
	GroupPersuade        = "Persuade"
	GroupPilot           = "Pilot"
	GroupProfession      = "Profession" //byTyp-Profession             e
	GroupRecon           = "Recon"
	GroupSciencePhysical = "Science Physical" //byField Science                d
	GroupScienceLife     = "Science Life"     //byField Science                d
	GroupScienceSocial   = "Science Social"   //byField Science                d
	GroupScienceSpace    = "Science Space"    //byField Science                d
	GroupScience         = "Science"
	GroupSeafarer        = "Seafarer"
	GroupStealth         = "Stealth"
	GroupSteward         = "Steward"
	GroupStreetwise      = "Streetwise"
	GroupSurvival        = "Survival"
	GroupTactics         = "Tactics"
	GroupVaccsuit        = "Vacc Suit"
)

func BackgroundSkills() []string {
	return []string{
		GroupAdmin,
		GroupAnimals,
		GroupArt,
		GroupAthletics,
		GroupCarouse,
		GroupDrive,
		GroupElectronics,
		GroupFlyer,
		GroupLanguage,
		GroupMechanic,
		GroupMedic,
		GroupProfession,
		GroupScienceLife,
		GroupSciencePhysical,
		GroupScienceSocial,
		GroupScienceSpace,
		GroupSeafarer,
		GroupStreetwise,
		GroupSurvival,
		GroupVaccsuit,
	}
}

// type asset struct {
// 	name              string
// 	description       string
// 	code              string
// 	usage             string
// 	numericalValues   []int
// 	numericalValuesFl []float64
// 	list1             []string
// }

type Skill interface {
	Proficiency() int
	Specialities() ([]string, []int)
	Train(string) error
}

//Proficiency() int
//Train()
//Ensure()
//Set()

func BasicTraining(name string) Skill {
	c := asset{} //в данном случае ассетом является вся группа скилов. если я хочу увеличить специализацию то мне нужна проверка
	//является ли эта специализация частью данной группы. При создании группы я добавляю все ее специализации в лист ассета 1
	specs := groupSpecs(name)
	for i := range specs {
		c.list1 = append(c.list1, specs[i])
		c.numericalValues = append(c.numericalValues, 0)
	}
	if len(specs) == 0 {
		c.list1 = append(c.list1, name)
		c.numericalValues = append(c.numericalValues, 0)
	}
	c.numericalValues = append(c.numericalValues, 0)
	return &c
}

func (a *asset) IsPresent(skill string) bool {
	for _, val := range a.list1 {
		if val == skill {
			return true
		}
	}
	return false
}

func groupSpecs(group string) []string {
	specs := []string{}
	//specs = append(specs, group) - если список специализаций пуст - то качаем группу как отдельный навык
	switch group {
	default:
	case GroupAnimals:
		specs = append(specs, group+" (Handling)")
		specs = append(specs, group+" (Veterinary)")
		specs = append(specs, group+" (Training)")
	case GroupArt:
		specs = append(specs, group+" (Performer)")
		specs = append(specs, group+" (Holography)")
		specs = append(specs, group+" (Instrument)")
		specs = append(specs, group+" (Visual Media)")
		specs = append(specs, group+" (Write)")
	case GroupAthletics:
		specs = append(specs, group+" (Dexterity)")
		specs = append(specs, group+" (Endurance)")
		specs = append(specs, group+" (Strength)")
	case GroupDrive:
		specs = append(specs, group+" (Hovercraft)")
		specs = append(specs, group+" (Mole)")
		specs = append(specs, group+" (Track)")
		specs = append(specs, group+" (Walker)")
		specs = append(specs, group+" (Wheel)")
	case GroupElectronics:
		specs = append(specs, group+" (Comms)")
		specs = append(specs, group+" (Computers)")
		specs = append(specs, group+" (Remote Ops)")
		specs = append(specs, group+" (Sensors)")
	case GroupEngineer:
		specs = append(specs, group+" (M-drive)")
		specs = append(specs, group+" (J-drive)")
		specs = append(specs, group+" (Life Support)")
		specs = append(specs, group+" (Power)")
	case GroupFlyer:
		specs = append(specs, group+" (Airship)")
		specs = append(specs, group+" (Grav)")
		specs = append(specs, group+" (Ornithopter)")
		specs = append(specs, group+" (Rotor)")
		specs = append(specs, group+" (Wing)")
	case GroupGunner:
		specs = append(specs, group+" (Turret)")
		specs = append(specs, group+" (Ortillery)")
		specs = append(specs, group+" (Screen)")
		specs = append(specs, group+" (Capital)")
	case GroupGuncombat:
		specs = append(specs, group+" (Archaic)")
		specs = append(specs, group+" (Energy)")
		specs = append(specs, group+" (Slug)")
	case GroupHeavyweapon:
		specs = append(specs, group+" (Artillery)")
		specs = append(specs, group+" (Man Portable)")
		specs = append(specs, group+" (Vehicle)")
	case GroupLanguage:
		specs = append(specs, group+" (Anglic)")
		specs = append(specs, group+" (Vilani)")
		specs = append(specs, group+" (Zdetl)")
		specs = append(specs, group+" (Oynprith)")
		specs = append(specs, group+" (Aslan)")
		specs = append(specs, group+" (Vargr)")
	case GroupMelee:
		specs = append(specs, group+" (Unarmed)")
		specs = append(specs, group+" (Blade)")
		specs = append(specs, group+" (Bludgeon)")
		specs = append(specs, group+" (Natural)")
	case GroupPilot:
		specs = append(specs, group+" (Small Craft)")
		specs = append(specs, group+" (Spacecraft)")
		specs = append(specs, group+" (Capital Ships)")
	case GroupProfession:
		specs = append(specs, group+" (Profession 1)")
		specs = append(specs, group+" (Profession 2)")
		specs = append(specs, group+" (Profession 3)")
	case GroupScienceLife:
		group = "Sciense"
		specs = append(specs, group+" (Biology)")
		specs = append(specs, group+" (Cybernetics)")
		specs = append(specs, group+" (Genetics)")
		specs = append(specs, group+" (Psionicology)")
	case GroupSciencePhysical:
		group = "Sciense"
		specs = append(specs, group+" (Physics)")
		specs = append(specs, group+" (Chemistry)")
		specs = append(specs, group+" (Electronics)")
	case GroupScienceSocial:
		group = "Sciense"
		specs = append(specs, group+" (Archeology)")
		specs = append(specs, group+" (Economics)")
		specs = append(specs, group+" (History)")
		specs = append(specs, group+" (Linguistics)")
		specs = append(specs, group+" (Philosophy)")
		specs = append(specs, group+" (Psychology)")
		specs = append(specs, group+" (Sophontology)")
	case GroupScienceSpace:
		group = "Sciense"
		specs = append(specs, group+" (Planetology)")
		specs = append(specs, group+" (Robotics)")
		specs = append(specs, group+" (Xenology)")
	case GroupSeafarer:
		specs = append(specs, group+" (Ocean Ships)")
		specs = append(specs, group+" (Personal)")
		specs = append(specs, group+" (Sail)")
		specs = append(specs, group+" (Submarine)")
	case GroupTactics:
		specs = append(specs, group+" (Military)")
		specs = append(specs, group+" (Naval)")

	}
	return specs
}

func (a *asset) Proficiency() int {
	return a.numericalValues[0]
}

func (a *asset) Specialities() ([]string, []int) {
	specs := []string{}
	vals := []int{}
	for i, val := range a.list1 {
		specs = append(specs, val)
		vals = append(vals, a.numericalValues[i])
	}
	return specs, vals
}

func (a *asset) Train(spec string) error {
	for i := range a.list1 {
		if spec == a.list1[i] {
			a.numericalValues[i]++
			return nil
		}
	}
	return errors.New("'" + spec + "' not found among specialisations")
}
