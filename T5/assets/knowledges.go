package assets

import "fmt"

//Knowledge - a body of information based on a field of science or
//experience.
type Knowledge struct {
	code            int
	alias           string
	rating          int
	maxRating       int
	affileatedSkill int
	err             error
}

func NewKnowledge(code int, optionalData ...string) *Knowledge {
	kn := Knowledge{}
	kn.code = code
	kn.maxRating = 6
	switch kn.code {
	default:
		kn.err = fmt.Errorf("creation failed: Code '%v' unregnized", code)
	case KNOWLEDGE_Rider:
		kn.alias = "Rider"
		kn.affileatedSkill = SKILL_Animals
	case KNOWLEDGE_Teamster:
		kn.alias = "Teamster"
		kn.affileatedSkill = SKILL_Animals
	case KNOWLEDGE_Trainer:
		kn.alias = "Trainer"
		kn.affileatedSkill = SKILL_Animals
	case KNOWLEDGE_ACV:
		kn.alias = "ACV"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_Automotive:
		kn.alias = "Automotive"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_GravDriver:
		kn.alias = "Grav Driver"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_Legged:
		kn.alias = "Legged"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_Mole:
		kn.alias = "Mole"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_Tracked:
		kn.alias = "Tracked"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_Wheeled:
		kn.alias = "Wheeled"
		kn.affileatedSkill = SKILL_Driver
	case KNOWLEDGE_JumpDrives:
		kn.alias = "Jump Drives"
		kn.affileatedSkill = SHIP_Engineer
	case KNOWLEDGE_LifeSupport:
		kn.alias = "Life Support"
		kn.affileatedSkill = SHIP_Engineer
	case KNOWLEDGE_ManeuverDrive:
		kn.alias = "Maneuver Drive"
		kn.affileatedSkill = SHIP_Engineer
	case KNOWLEDGE_PowerSystems:
		kn.alias = "Power Systems"
		kn.affileatedSkill = SHIP_Engineer
	case KNOWLEDGE_BattleDress:
		kn.alias = "Battle Dress"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_Beams:
		kn.alias = "Beams"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_Blades:
		kn.alias = "Blades"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_Exotics:
		kn.alias = "Exotics"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_SlugThrowers:
		kn.alias = "Slug Throwers"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_Sprays:
		kn.alias = "Sprays"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_Unarmed:
		kn.alias = "Unarmed"
		kn.affileatedSkill = SOLDER_Fighter
	case KNOWLEDGE_Aeronautics:
		kn.alias = "Aeronautics"
		kn.affileatedSkill = SKILL_Flyer
	case KNOWLEDGE_Flapper:
		kn.alias = "Flapper"
		kn.affileatedSkill = SKILL_Flyer
	case KNOWLEDGE_GravFlyer:
		kn.alias = "Grav Flyer"
		kn.affileatedSkill = SKILL_Flyer
	case KNOWLEDGE_LTA:
		kn.alias = "LTA"
		kn.affileatedSkill = SKILL_Flyer
	case KNOWLEDGE_Rotor:
		kn.alias = "Rotor"
		kn.affileatedSkill = SKILL_Flyer
	case KNOWLEDGE_Winged:
		kn.alias = "Winged"
		kn.affileatedSkill = SKILL_Flyer
	case KNOWLEDGE_BayWeapons:
		kn.alias = "Bay Weapons"
		kn.affileatedSkill = SHIP_Gunner
	case KNOWLEDGE_Ortillery:
		kn.alias = "Ortillery"
		kn.affileatedSkill = SHIP_Gunner
	case KNOWLEDGE_Screens:
		kn.alias = "Screens"
		kn.affileatedSkill = SHIP_Gunner
	case KNOWLEDGE_Spines:
		kn.alias = "Spines"
		kn.affileatedSkill = SHIP_Gunner
	case KNOWLEDGE_Turrets:
		kn.alias = "Turrets"
		kn.affileatedSkill = SHIP_Gunner
	case KNOWLEDGE_Artillery:
		kn.alias = "Artillery"
		kn.affileatedSkill = SOLDER_HeavyWeapons
	case KNOWLEDGE_Launcher:
		kn.alias = "Launcher"
		kn.affileatedSkill = SOLDER_HeavyWeapons
	case KNOWLEDGE_Ordnance:
		kn.alias = "Ordnance"
		kn.affileatedSkill = SOLDER_HeavyWeapons
	case KNOWLEDGE_WMD:
		kn.alias = "WMD"
		kn.affileatedSkill = SOLDER_HeavyWeapons
	case KNOWLEDGE_SmallCraft:
		kn.alias = "Small Craft"
		kn.affileatedSkill = SHIP_Pilot
	case KNOWLEDGE_SpacecraftACS:
		kn.alias = "Spacecraft ACS"
		kn.affileatedSkill = SHIP_Pilot
	case KNOWLEDGE_SpacecraftBCS:
		kn.alias = "Spacecraft BCS"
		kn.affileatedSkill = SHIP_Pilot
	case KNOWLEDGE_Aquanautics:
		kn.alias = "Aquanautics"
		kn.affileatedSkill = SKILL_Seafarer
	case KNOWLEDGE_GravSea:
		kn.alias = "Sea Grav"
		kn.affileatedSkill = SKILL_Seafarer
	case KNOWLEDGE_Boat:
		kn.alias = "Boat"
		kn.affileatedSkill = SKILL_Seafarer
	case KNOWLEDGE_ShipSea:
		kn.alias = "Sea Ship"
		kn.affileatedSkill = SKILL_Seafarer
	case KNOWLEDGE_Sub:
		kn.alias = "Sub"
		kn.affileatedSkill = SKILL_Seafarer
	case KNOWLEDGE_Archeology:
		kn.alias = "Archeology"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Biology:
		kn.alias = "Biology"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Chemistry:
		kn.alias = "Chemistry"
		kn.affileatedSkill = -1
	case KNOWLEDGE_History:
		kn.alias = "History"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Linguistics:
		kn.alias = "Linguistics"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Philosophy:
		kn.alias = "Philosophy"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Physics:
		kn.alias = "Physics"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Planetology:
		kn.alias = "Planetology"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Psionicology:
		kn.alias = "Psionicology"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Psychohistory:
		kn.alias = "Psychohistory"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Psychology:
		kn.alias = "Psychology"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Robotics:
		kn.alias = "Robotics"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Sophontology:
		kn.alias = "Sophontology"
		kn.affileatedSkill = -1
	case KNOWLEDGE_Specialized:
		kn.alias = "Specialized"
		kn.affileatedSkill = -1
	}
	return &kn
}

func (kn *Knowledge) Train() error {
	if kn.rating >= kn.maxRating {
		kn.rating = kn.maxRating
		return fmt.Errorf("knowledge '%v' cannot be trained: maximum level reached", kn.alias)
	}
	kn.rating++
	return nil
}
