package assets

import (
	"fmt"
)

//Skills are the primary means by which characters do things in Traveller.
//	Each character has a variety of skills, and the higher a skill rating,
//	the more expert the character is with that skill. With training, any
//	character can eventually become proficient at any skill

const (
	SKILL_Admin = iota
	SKILL_Advocate
	SKILL_Animals
	SKILL_Athlete
	SKILL_Broker
	SKILL_Bureaucrat
	SKILL_Comms
	SKILL_Computer
	SKILL_Counsellor
	SKILL_Designer
	SKILL_Diplomat
	SKILL_Driver
	SKILL_Explosives
	SKILL_FleetTactics
	SKILL_Flyer
	SKILL_Forensics
	SKILL_Gambler
	SKILL_HighG
	SKILL_HostileEnviron
	SKILL_JOT
	SKILL_Language
	SKILL_Leader
	SKILL_Liaison
	SKILL_NavalArchitect
	SKILL_Seafarer
	SKILL_Stealth
	SKILL_Strategy
	SKILL_Streetwise
	SKILL_Survey
	SKILL_Survival
	SKILL_Tactics
	SKILL_Teacher
	SKILL_Trader
	SKILL_VaccSuit
	SKILL_ZeroG
	SKILL_Astrogator
	SHIP_Engineer
	SHIP_Gunner
	SHIP_Medic
	SHIP_Pilot
	SHIP_Sensors
	SHIP_Steward
	TRADE_Biologics
	TRADE_Craftsman
	TRADE_Electronics
	TRADE_Fluidics
	TRADE_Gravitics
	TRADE_Magnetics
	TRADE_Mechanic
	TRADE_Photonics
	TRADE_Polymers
	TRADE_Programmer
	ART_Actor
	ART_Artist
	ART_Author
	ART_Chef
	ART_Dancer
	ART_Musician
	SOLDER_Fighter
	SOLDER_ForwardObserver
	SOLDER_HeavyWeapons
	SOLDER_Navigator
	SOLDER_Recon
	SOLDER_Sapper
	TALENT_Compute
	TALENT_Empath
	TALENT_Hibernate
	TALENT_Hypno
	TALENT_Intuition
	TALENT_Math
	TALENT_MemAware
	TALENT_Memorize
	TALENT_MemPercept
	TALENT_MemScent
	TALENT_MemSight
	TALENT_MemSound
	TALENT_Morph
	TALENT_Rage
	TALENT_SoundMimic
	PRSNL_Carouse
	PRSNL_Query
	PRSNL_Persuade
	PRSNL_Command
	INTUIT_Curiosity
	INTUIT_Insight
	INTUIT_Luck
	KNOWLEDGE_Animals
	KNOWLEDGE_Rider
	KNOWLEDGE_Teamster
	KNOWLEDGE_Trainer
	KNOWLEDGE_ACV
	KNOWLEDGE_Automotive
	KNOWLEDGE_GravDriver
	KNOWLEDGE_Legged
	KNOWLEDGE_Mole
	KNOWLEDGE_Tracked
	KNOWLEDGE_Wheeled
	KNOWLEDGE_JumpDrives
	KNOWLEDGE_LifeSupport
	KNOWLEDGE_ManeuverDrive
	KNOWLEDGE_PowerSystems
	KNOWLEDGE_BattleDress
	KNOWLEDGE_Beams
	KNOWLEDGE_Blades
	KNOWLEDGE_Exotics
	KNOWLEDGE_SlugThrowers
	KNOWLEDGE_Sprays
	KNOWLEDGE_Unarmed
	KNOWLEDGE_Aeronautics
	KNOWLEDGE_Flapper
	KNOWLEDGE_GravFlyer
	KNOWLEDGE_LTA
	KNOWLEDGE_Rotor
	KNOWLEDGE_Winged
	KNOWLEDGE_BayWeapons
	KNOWLEDGE_Ortillery
	KNOWLEDGE_Screens
	KNOWLEDGE_Spines
	KNOWLEDGE_Turrets
	KNOWLEDGE_Artillery
	KNOWLEDGE_Launcher
	KNOWLEDGE_Ordnance
	KNOWLEDGE_WMD
	KNOWLEDGE_SmallCraft
	KNOWLEDGE_SpacecraftACS
	KNOWLEDGE_SpacecraftBCS
	KNOWLEDGE_Aquanautics
	KNOWLEDGE_GravSea
	KNOWLEDGE_Boat
	KNOWLEDGE_ShipSea
	KNOWLEDGE_Sub
	KNOWLEDGE_Archeology
	KNOWLEDGE_Biology
	KNOWLEDGE_Chemistry
	KNOWLEDGE_History
	KNOWLEDGE_Linguistics
	KNOWLEDGE_Philosophy
	KNOWLEDGE_Physics
	KNOWLEDGE_Planetology
	KNOWLEDGE_Psionicology
	KNOWLEDGE_Psychohistory
	KNOWLEDGE_Psychology
	KNOWLEDGE_Robotics
	KNOWLEDGE_Sophontology
	KNOWLEDGE_Specialized
)

//Assets - Skills, Knowledges, Talents, Characteristics, and Modifiers used in a task.
type Assets struct {
	Skills map[int]Skill
}

/////////////////////////////////////

//Skill - statement of ability based on a job, vocation, or interest.
type Skill struct {
	code       int
	alias      string
	rating     int
	knowledges []int
	isDefault  bool
	category   string
	err        error
}

func NewSkill(code int) *Skill {
	sk := Skill{}
	sk.code = code
	switch code {
	default:
		sk.err = fmt.Errorf("can not create skill: Code '%v' is unknown", code)
	case SKILL_Admin:
		sk.alias = "Admin"
		sk.category = "Broad"
	case SKILL_Advocate:
		sk.alias = "Advocate"
		sk.category = "Broad"
	case SKILL_Animals:
		sk.alias = "Animals"
		sk.category = "Broad"
		sk.knowledges = []int{KNOWLEDGE_Rider, KNOWLEDGE_Teamster, KNOWLEDGE_Trainer}
		sk.isDefault = true
	case SKILL_Athlete:
		sk.alias = "Athlete"
		sk.category = "Broad"
		sk.isDefault = true
	case SKILL_Broker:
		sk.alias = "Broker"
		sk.category = "Broad"
	case SKILL_Bureaucrat:
		sk.alias = "Bureaucrat"
		sk.category = "Broad"
	case SKILL_Comms:
		sk.alias = "Comms"
		sk.category = "Broad"
		sk.isDefault = true
	case SKILL_Computer:
		sk.alias = "Computer"
		sk.category = "Broad"
		sk.isDefault = true
	case SKILL_Counsellor:
		sk.alias = "Counsellor"
		sk.category = "Broad"
	case SKILL_Designer:
		sk.alias = "Designer"
		sk.category = "Broad"
	case SKILL_Diplomat:
		sk.alias = "Diplomat"
		sk.category = "Broad"
	case SKILL_Driver:
		sk.alias = "Driver"
		sk.knowledges = []int{KNOWLEDGE_ACV, KNOWLEDGE_Automotive, KNOWLEDGE_GravDriver, KNOWLEDGE_Legged, KNOWLEDGE_Mole, KNOWLEDGE_Tracked, KNOWLEDGE_Wheeled}
		sk.category = "Broad"
		sk.isDefault = true
	case SKILL_Explosives:
		sk.alias = "Explosives"
		sk.category = "Broad"
	case SKILL_FleetTactics:
		sk.alias = "Fleet Tactics"
		sk.category = "Broad"
	case SKILL_Flyer:
		sk.alias = "Flyer"
		sk.knowledges = []int{KNOWLEDGE_Aeronautics, KNOWLEDGE_Flapper, KNOWLEDGE_GravFlyer, KNOWLEDGE_LTA, KNOWLEDGE_Rotor, KNOWLEDGE_Winged}
		sk.category = "Broad"
	case SKILL_Forensics:
		sk.alias = "Forensics"
		sk.category = "Broad"
	case SKILL_Gambler:
		sk.alias = "Gambler"
		sk.category = "Broad"
	case SKILL_HighG:
		sk.alias = "High-G"
		sk.category = "Broad"
	case SKILL_HostileEnviron:
		sk.alias = "HostileEnviron"
		sk.category = "Broad"
	case SKILL_JOT:
		sk.alias = "Jack-Of-All-Trades"
		sk.category = "Broad"
	case SKILL_Language:
		sk.alias = "Language"
		sk.category = "Broad"
		sk.knowledges = []int{999}
	case SKILL_Leader:
		sk.alias = "Leader"
		sk.category = "Broad"
	case SKILL_Liaison:
		sk.alias = "Liaison"
		sk.category = "Broad"
	case SKILL_NavalArchitect:
		sk.alias = "NavalArchitect"
		sk.category = "Broad"
	case SKILL_Seafarer:
		sk.alias = "Seafarer"
		sk.knowledges = []int{KNOWLEDGE_Aquanautics, KNOWLEDGE_GravSea, KNOWLEDGE_Boat, KNOWLEDGE_ShipSea, KNOWLEDGE_Sub}
		sk.category = "Broad"
	case SKILL_Stealth:
		sk.alias = "Stealth"
		sk.category = "Broad"
	case SKILL_Strategy:
		sk.alias = "Strategy"
		sk.category = "Broad"
	case SKILL_Streetwise:
		sk.alias = "Streetwise"
		sk.category = "Broad"
	case SKILL_Survey:
		sk.alias = "Survey"
		sk.category = "Broad"
	case SKILL_Survival:
		sk.alias = "Survival"
		sk.category = "Broad"
	case SKILL_Tactics:
		sk.alias = "Tactics"
		sk.category = "Broad"
	case SKILL_Teacher:
		sk.alias = "Teacher"
		sk.category = "Broad"
	case SKILL_Trader:
		sk.alias = "Trader"
		sk.category = "Broad"
	case SKILL_VaccSuit:
		sk.alias = "VaccSuit"
		sk.category = "Broad"
		sk.isDefault = true
	case SKILL_ZeroG:
		sk.alias = "ZeroG"
	case SKILL_Astrogator:
		sk.alias = "Astrogator"
		sk.category = "Starship"
	case SHIP_Engineer:
		sk.alias = "Engineer"
		sk.knowledges = []int{KNOWLEDGE_JumpDrives, KNOWLEDGE_LifeSupport, KNOWLEDGE_ManeuverDrive, KNOWLEDGE_PowerSystems}
		sk.category = "Starship"
	case SHIP_Gunner:
		sk.alias = "Gunner"
		sk.knowledges = []int{KNOWLEDGE_BayWeapons, KNOWLEDGE_Ortillery, KNOWLEDGE_Screens, KNOWLEDGE_Spines, KNOWLEDGE_Turrets}
		sk.category = "Starship"
	case SHIP_Medic:
		sk.alias = "Medic"
		sk.category = "Starship"
	case SHIP_Pilot:
		sk.alias = "Pilot"
		sk.knowledges = []int{KNOWLEDGE_SmallCraft, KNOWLEDGE_SpacecraftACS, KNOWLEDGE_SpacecraftBCS}
		sk.category = "Starship"
	case SHIP_Sensors:
		sk.alias = "Sensors"
		sk.category = "Starship"
	case SHIP_Steward:
		sk.alias = "Steward"
		sk.isDefault = true
		sk.category = "Starship"
	case TRADE_Biologics:
		sk.alias = "Biologics"
		sk.category = "Trade"
	case TRADE_Craftsman:
		sk.alias = "Craftsman"
		sk.category = "Trade"
	case TRADE_Electronics:
		sk.alias = "Electronics"
		sk.category = "Trade"
	case TRADE_Fluidics:
		sk.alias = "Fluidics"
		sk.category = "Trade"
	case TRADE_Gravitics:
		sk.alias = "Gravitics"
		sk.category = "Trade"
	case TRADE_Magnetics:
		sk.alias = "Magnetics"
		sk.category = "Trade"
	case TRADE_Mechanic:
		sk.alias = "Mechanic"
		sk.category = "Trade"
		sk.isDefault = true
	case TRADE_Photonics:
		sk.alias = "Photonics"
		sk.category = "Trade"
	case TRADE_Polymers:
		sk.alias = "Polymers"
		sk.category = "Trade"
	case TRADE_Programmer:
		sk.alias = "Programmer"
		sk.category = "Trade"
	case ART_Actor:
		sk.alias = "Actor"
		sk.isDefault = true
		sk.category = "Art"
	case ART_Artist:
		sk.alias = "Artist"
		sk.isDefault = true
		sk.category = "Art"
	case ART_Author:
		sk.alias = "Author"
		sk.category = "Art"
	case ART_Chef:
		sk.alias = "Chef"
		sk.category = "Art"
	case ART_Dancer:
		sk.alias = "Dancer"
		sk.category = "Art"
	case ART_Musician:
		sk.alias = "Musician"
		sk.category = "Art"
	case SOLDER_Fighter:
		sk.alias = "Fighter"
		sk.knowledges = []int{KNOWLEDGE_BattleDress, KNOWLEDGE_Beams, KNOWLEDGE_Blades, KNOWLEDGE_Exotics, KNOWLEDGE_SlugThrowers, KNOWLEDGE_Sprays, KNOWLEDGE_Unarmed}
		sk.isDefault = true
		sk.category = "Solder"
	case SOLDER_ForwardObserver:
		sk.alias = "Forward Observer"
		sk.category = "Solder"
	case SOLDER_HeavyWeapons:
		sk.alias = "Heavy Weapons"
		sk.knowledges = []int{KNOWLEDGE_Artillery, KNOWLEDGE_Launcher, KNOWLEDGE_Ordnance, KNOWLEDGE_WMD}
		sk.category = "Solder"
	case SOLDER_Navigator:
		sk.alias = "Navigator"
		sk.category = "Solder"
	case SOLDER_Recon:
		sk.alias = "Recon"
		sk.category = "Solder"
	case SOLDER_Sapper:
		sk.alias = "Sapper"
		sk.category = "Solder"
	}
	return &sk
}

func (s *Skill) Rating() int {
	return s.rating
}

func (s *Skill) Category() string {
	return s.category
}

func (s *Skill) String() string {
	return fmt.Sprint("%v-%v", s.alias, s.rating)
}

func (s *Skill) Train() error {
	s.rating++
	if s.rating > 15 {
		s.rating = 15
		return fmt.Errorf("skill '%v' cannot be trained: maximum level reached", s.alias)
	}
	return nil
}

////////////////////////////////////////////////////

//Knowledge - a body of information based on a field of science or
//experience.
type Knowledge struct {
	code            int
	alias           string
	rating          int
	affileatedSkill int
}
