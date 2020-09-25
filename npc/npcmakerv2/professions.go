package npcmakerv2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/entity"
	"github.com/Galdoba/utils"
)

const (
	CareerB0AgentLawenforcement             = "B0=11=1==Law Enforcement Agent"
	CareerB0Agentintelligence               = "B0=11=2==Intelligence Agent"
	CareerB0AgentCorporate                  = "B0=11=3==Corporate Agent"
	CareerB0ArmySupport                     = "B0=12=1==Army Support"
	CareerB0ArmyInfatry                     = "B0=12=2==Army Infatry"
	CareerB0ArmyCavalry                     = "B0=12=3==Army Cavalry"
	CareerB0CitizenCorporate                = "B0=13=1==Corporate Citizen"
	CareerB0CitizenWorker                   = "B0=13=2==Worker (Citizen)"
	CareerB0CitizenColonist                 = "B0=13=3==Colonist (Citizen)"
	CareerB0DrifterBarbarian                = "B0=14=1==Barbarian (Drifter)"
	CareerB0DrifterWandere                  = "B0=14=2==Wanderer (Drifter)"
	CareerB0DrifterScavenger                = "B0=14=3==Scavenger (Drifter)"
	CareerB0EntertainerArtist               = "B0=15=1==Artist (Entertainer)"
	CareerB0EntertainerJournalist           = "B0=15=2==Journalist (Entertainer)"
	CareerB0EntertainerPerformer            = "B0=15=3==Performer (Entertainer)"
	CareerB0MarinesSupport                  = "B0=16=1==Support Marines"
	CareerB0MarinesStarMarines              = "B0=16=2==Star Marines"
	CareerB0MarinesGroundAssault            = "B0=16=3==Ground Assault Marines"
	CareerB0MerchantMarine                  = "B0=17=1==Merchant Marines"
	CareerB0MerchantFreeTrader              = "B0=17=2==Free Trader"
	CareerB0MerchantBroker                  = "B0=17=3==Broker"
	CareerB0NavyLineCrew                    = "B0=18=1==Navy Line/Crew"
	CareerB0NavyEngineering                 = "B0=18=2==Navy Engineering/Gunnery"
	CareerB0NavyFlight                      = "B0=18=3==Navy Flight"
	CareerB0NobilityAdministrator           = "B0=19=1==Noble Administrator"
	CareerB0NobilityDiplomat                = "B0=19=2==Noble Diplomat"
	CareerB0NobilityDilettante              = "B0=19=3==Noble Dilettante"
	CareerB0RogueThief                      = "B0=1A=1==Thief"
	CareerB0RogueEnforcer                   = "B0=1A=2==Enforcer"
	CareerB0RoguePirate                     = "B0=1A=3==Pirate"
	CareerB0ScholarFieldResearcher          = "B0=1B=1==Field Researcher"
	CareerB0ScholarScientist                = "B0=1B=2==Scientist"
	CareerB0ScholarPhysician                = "B0=1B=3==Physician"
	CareerB0ScoutCourier                    = "B0=1C=1==Scout Courier"
	CareerB0ScoutSurvey                     = "B0=1C=2==Scout Survey"
	CareerB0ScoutExploration                = "B0=1C=3==Scout Exploration"
	CareerB1LawEnforcementPatroller         = "B1=11=1==Law Enforcement (Patroller)"
	CareerB1LawEnforcementSpecialOperations = "B1=11=2==Law Enforcement (Special Operations)"
	CareerB1LawEnforcementCustoms           = "B1=11=3==Law Enforcement (Customs)"
	////
	STR                    = entity.CharCodeTrvC1STRENGTH
	DEX                    = entity.CharCodeTrvC2DEXTERITY
	END                    = entity.CharCodeTrvC3ENDURANCE
	INT                    = entity.CharCodeTrvC4INTELLIGENCE
	EDU                    = entity.CharCodeTrvC5EDUCATION
	SOC                    = entity.CharCodeTrvC6SOCIAL
	Admin                  = entity.SCTrvAdmin
	Advocate               = entity.SCTrvAdvocate
	Animals                = entity.SCTrvAnimals
	AnimalsHandling        = entity.SCTrvAnimalsHandling
	AnimalsTraining        = entity.SCTrvAnimalsTraining
	AnimalsVeterinary      = entity.SCTrvAnimalsVeterinary
	Art                    = entity.SCTrvArt
	ArtHolography          = entity.SCTrvArtHolography
	ArtInstrument          = entity.SCTrvArtInstrument
	ArtPerformer           = entity.SCTrvArtPerformer
	ArtVisualMedia         = entity.SCTrvArtVisualMedia
	ArtWrite               = entity.SCTrvArtWrite
	Astrogation            = entity.SCTrvAstrogation
	Athletics              = entity.SCTrvAthletics
	AthleticsDEX           = entity.SCTrvAthleticsDEX
	AthleticsEND           = entity.SCTrvAthleticsEND
	AthleticsSTR           = entity.SCTrvAthleticsSTR
	Broker                 = entity.SCTrvBroker
	Carouse                = entity.SCTrvCarouse
	Deception              = entity.SCTrvDeception
	Diplomat               = entity.SCTrvDiplomat
	Drive                  = entity.SCTrvDrive
	DriveHovercraft        = entity.SCTrvDriveHovercraft
	DriveMole              = entity.SCTrvDriveMole
	DriveTrack             = entity.SCTrvDriveTrack
	DriveWalker            = entity.SCTrvDriveWalker
	DriveWheel             = entity.SCTrvDriveWheel
	Electronics            = entity.SCTrvElectronics
	Comms                  = entity.SCTrvElectronicsComms
	Computers              = entity.SCTrvElectronicsComputers
	RemoteOps              = entity.SCTrvElectronicsRemoteops
	Sensors                = entity.SCTrvElectronicsSensors
	Engineer               = entity.SCTrvEngineer
	EngineerJdrive         = entity.SCTrvEngineerJdrive
	EngineerLifesupport    = entity.SCTrvEngineerLifesupport
	EngineerMdrive         = entity.SCTrvEngineerMdrive
	EngineerPower          = entity.SCTrvEngineerPower
	Explosives             = entity.SCTrvExplosives
	Flyer                  = entity.SCTrvFlyer
	FlyerAirship           = entity.SCTrvFlyerAirship
	FlyerGrav              = entity.SCTrvFlyerGrav
	FlyerOrnithopter       = entity.SCTrvFlyerOrnithopter
	FlyerRotor             = entity.SCTrvFlyerRotor
	FlyerWing              = entity.SCTrvFlyerWing
	Gambler                = entity.SCTrvGambler
	GunCombat              = entity.SCTrvGuncombat
	GuncombatArchaic       = entity.SCTrvGuncombatArchaic
	GuncombatEnergy        = entity.SCTrvGuncombatEnergy
	GuncombatSlug          = entity.SCTrvGuncombatSlug
	Gunner                 = entity.SCTrvGunner
	GunnerCapital          = entity.SCTrvGunnerCapital
	GunnerOrtilery         = entity.SCTrvGunnerOrtilery
	GunnerScreen           = entity.SCTrvGunnerScreen
	GunnerTurret           = entity.SCTrvGunnerTurret
	HeavyWeapon            = entity.SCTrvHeavyweapon
	HeavyweaponArtilery    = entity.SCTrvHeavyweaponArtilery
	HeavyweaponManportable = entity.SCTrvHeavyweaponManportable
	HeavyweaponVehicle     = entity.SCTrvHeavyweaponVehicle
	Investigate            = entity.SCTrvInvestigate
	Jackofalltrades        = entity.SCTrvJackofalltrades
	Language               = entity.SCTrvLanguage
	LanguageAnglic         = entity.SCTrvLanguageAnglic
	Leadership             = entity.SCTrvLeadership
	Mechanic               = entity.SCTrvMechanic
	Medic                  = entity.SCTrvMedic
	Melee                  = entity.SCTrvMelee
	MeleeBlade             = entity.SCTrvMeleeBlade
	MeleeBludgeon          = entity.SCTrvMeleeBludgeon
	MeleeNatural           = entity.SCTrvMeleeNatural
	MeleeUnarmed           = entity.SCTrvMeleeUnarmed
	Navigation             = entity.SCTrvNavigation
	Persuade               = entity.SCTrvPersuade
	Pilot                  = entity.SCTrvPilot
	PilotCapitalships      = entity.SCTrvPilotCapitalships
	PilotSmallcraft        = entity.SCTrvPilotSmallcraft
	PilotSpacecraft        = entity.SCTrvPilotSpacecraft
	Profession             = entity.SCTrvProfession
	ProfessionAny          = entity.SCTrvProfessionAny
	Recon                  = entity.SCTrvRecon
	SciencePhysical        = entity.SCTrvSciencePhysical
	ScienceLife            = entity.SCTrvScienceLife
	ScienceSocial          = entity.SCTrvScienceSocial
	ScienceSpace           = entity.SCTrvScienceSpace
	ScienceAny             = entity.SCTrvScienceAny
	Seafarer               = entity.SCTrvSeafarer
	SeafarerOceanships     = entity.SCTrvSeafarerOceanships
	SeafarerPersonal       = entity.SCTrvSeafarerPersonal
	SeafarerSail           = entity.SCTrvSeafarerSail
	SeafarerSubmarine      = entity.SCTrvSeafarerSubmarine
	Stealth                = entity.SCTrvStealth
	Steward                = entity.SCTrvSteward
	Streetwise             = entity.SCTrvStreetwise
	Survival               = entity.SCTrvSurvival
	Tactics                = entity.SCTrvTactics
	TacticsMilitary        = entity.SCTrvTacticsMilitary
	TacticsNavy            = entity.SCTrvTacticsNavy
	Vaccsuit               = entity.SCTrvVaccsuit
)

func ListCareers() []string {
	return []string{
		CareerB0AgentLawenforcement,
		CareerB0Agentintelligence,
		CareerB0AgentCorporate,
		CareerB0ArmySupport,
		CareerB0ArmyInfatry,
		CareerB0ArmyCavalry,
		CareerB0CitizenCorporate,
		CareerB0CitizenWorker,
		CareerB0CitizenColonist,
		CareerB0DrifterBarbarian,
		CareerB0DrifterWandere,
		CareerB0DrifterScavenger,
		CareerB0EntertainerArtist,
		CareerB0EntertainerJournalist,
		CareerB0EntertainerPerformer,
		CareerB0MarinesSupport,
		CareerB0MarinesStarMarines,
		CareerB0MarinesGroundAssault,
		CareerB0MerchantMarine,
		CareerB0MerchantFreeTrader,
		CareerB0MerchantBroker,
		CareerB0NavyLineCrew,
		CareerB0NavyEngineering,
		CareerB0NavyFlight,
		CareerB0NobilityAdministrator,
		CareerB0NobilityDiplomat,
		CareerB0NobilityDilettante,
		CareerB0RogueThief,
		CareerB0RogueEnforcer,
		CareerB0RoguePirate,
		CareerB0ScholarFieldResearcher,
		CareerB0ScholarScientist,
		CareerB0ScholarPhysician,
		CareerB0ScoutCourier,
		CareerB0ScoutSurvey,
		CareerB0ScoutExploration,
		CareerB1LawEnforcementPatroller,
		CareerB1LawEnforcementSpecialOperations,
		CareerB1LawEnforcementCustoms,
	}
}

func SearchCareers(input string) []string {
	valid := []string{}
	for _, val := range ListCareers() {
		if strings.Contains(val, input) {
			valid = append(valid, val)
		}
	}
	if len(valid) == 0 {
		return ListCareers()
	}
	return valid
}

var SaT map[string][]string

func init() {
	SaT = make(map[string][]string)
	SaT["B0=11pd"] = []string{GunCombat, DEX, END, Melee, INT, Athletics}
	SaT["B0=11ss"] = []string{Streetwise, Drive, Investigate, Computers, Recon, GunCombat}
	SaT["B0=11ae"] = []string{Advocate, Drive, Computers, Medic, Stealth, RemoteOps}
	SaT["B0=11=1"] = []string{Investigate, Recon, Streetwise, Stealth, Melee, Advocate}
	SaT["B0=11=1ma"] = []string{"INT 6", "END 6", "INT 6"}
	SaT["B0=11=2"] = []string{Investigate, Recon, Comms, Stealth, Persuade, Deception}
	SaT["B0=11=2ma"] = []string{"INT 6", "INT 7", "INT 5"}
	SaT["B0=11=3"] = []string{Investigate, Computers, Stealth, GunCombat, Deception, Streetwise}
	SaT["B0=11=3ma"] = []string{"INT 6", "INT 5", "INT 7"}

	SaT["B0=12pd"] = []string{STR, DEX, END, Gambler, Medic, MeleeUnarmed}
	SaT["B0=12ss"] = []string{Drive, Athletics, GunCombat, Recon, Melee, HeavyWeapon}
	SaT["B0=12ae"] = []string{Comms, Sensors, Navigation, Explosives, Engineer, Survival}
	SaT["B0=12os"] = []string{TacticsMilitary, Leadership, Advocate, Diplomat, TacticsMilitary, Admin}
	SaT["B0=12=1"] = []string{Mechanic, Drive, Flyer, Explosives, Comms, Medic}
	SaT["B0=12=1ma"] = []string{"END 5", "END 5", "EDU 7"}
	SaT["B0=12=2"] = []string{GunCombat, Melee, HeavyWeapon, Stealth, Athletics, Recon}
	SaT["B0=12=2ma"] = []string{"END 5", "STR 6", "EDU 6"}
	SaT["B0=12=3"] = []string{Mechanic, Drive, Flyer, Recon, Gunner, Sensors}
	SaT["B0=12=3ma"] = []string{"END 5", "DEX 7", "INT 5"}

	SaT["B0=13pd"] = []string{EDU, INT, Carouse, Gambler, Drive, Jackofalltrades}
	SaT["B0=13ss"] = []string{Drive, Flyer, Streetwise, Melee, Steward, Profession}
	SaT["B0=13ae"] = []string{Art, Advocate, Diplomat, Language, Computers, Medic}
	SaT["B0=13os"] = []string{}
	SaT["B0=13=1"] = []string{Advocate, Admin, Broker, Computers, Diplomat, Leadership}
	SaT["B0=13=1ma"] = []string{"EDU 5", "SOC 6", "INT 6"}
	SaT["B0=13=2"] = []string{Drive, Mechanic, Profession, Engineer, Profession, ScienceAny}
	SaT["B0=13=2ma"] = []string{"EDU 5", "END 4", "EDU 8"}
	SaT["B0=13=3"] = []string{Animals, Athletics, Jackofalltrades, Drive, Survival, Recon}
	SaT["B0=13=3ma"] = []string{"EDU 5", "INT 7", "END 5"}

	SaT["B0=14pd"] = []string{STR, END, DEX, Jackofalltrades, END, INT}
	SaT["B0=14ss"] = []string{Athletics, MeleeUnarmed, Recon, Streetwise, Stealth, Survival}
	SaT["B0=14ae"] = []string{}
	SaT["B0=14os"] = []string{}
	SaT["B0=14=1"] = []string{Animals, Carouse, MeleeBlade, Stealth, Seafarer, Survival}
	SaT["B0=14=1ma"] = []string{"EDU 2", "END 7", "STR 7"}
	SaT["B0=14=2"] = []string{Athletics, Deception, Recon, Stealth, Streetwise, Survival}
	SaT["B0=14=2ma"] = []string{"EDU 2", "END 7", "INT 7"}
	SaT["B0=14=3"] = []string{PilotSmallcraft, Mechanic, Astrogation, Vaccsuit, AthleticsDEX, GunCombat}
	SaT["B0=14=3ma"] = []string{"EDU 2", "DEX 7", "END 7"}

	SaT["B0=15pd"] = []string{DEX, INT, SOC, EDU, Carouse, Stealth}
	SaT["B0=15ss"] = []string{Art, Art, Carouse, Deception, Persuade, Steward}
	SaT["B0=15ae"] = []string{Advocate, Art, Deception, ScienceAny, Streetwise, Diplomat}
	SaT["B0=15os"] = []string{}
	SaT["B0=15=1"] = []string{Art, Carouse, Computers, Gambler, Persuade, Profession}
	SaT["B0=15=1ma"] = []string{"INT 5", "SOC 6", "INT 6"}
	SaT["B0=15=2"] = []string{pickAnyOf(ArtWrite, ArtHolography), Comms, Computers, Investigate, Recon, Streetwise}
	SaT["B0=15=2ma"] = []string{"INT 5", "EDU 7", "INT 5"}
	SaT["B0=15=3"] = []string{pickAnyOf(ArtPerformer, ArtVisualMedia, ArtInstrument), pickAnyOf(AthleticsDEX, AthleticsEND), Carouse, Deception, Stealth, Streetwise}
	SaT["B0=15=3ma"] = []string{"INT 5", "INT 5", "DEX 7"}

	SaT["B0=16pd"] = []string{STR, DEX, END, Gambler, MeleeUnarmed, MeleeBlade}
	SaT["B0=16ss"] = []string{Athletics, Vaccsuit, Tactics, HeavyWeapon, GunCombat, Stealth}
	SaT["B0=16ae"] = []string{Medic, Survival, Explosives, Engineer, Pilot, Medic}
	SaT["B0=16os"] = []string{Leadership, Tactics, Admin, Advocate, Vaccsuit, Leadership}
	SaT["B0=16=1"] = []string{Comms, Mechanic, pickAnyOf(Drive, Flyer), Medic, HeavyWeapon, GunCombat}
	SaT["B0=16=1ma"] = []string{"END 6", "END 5", "EDU 7"}
	SaT["B0=16=2"] = []string{Vaccsuit, AthleticsDEX, Gunner, MeleeBlade, Sensors, GunCombat}
	SaT["B0=16=2ma"] = []string{"END 6", "END 6", "EDU 6"}
	SaT["B0=16=3"] = []string{Vaccsuit, HeavyWeapon, Recon, MeleeBlade, TacticsMilitary, GunCombat}
	SaT["B0=16=3ma"] = []string{"END 6", "END 7", "EDU 5"}

	SaT["B0=17pd"] = []string{STR, DEX, END, INT, MeleeBlade, Streetwise}
	SaT["B0=17ss"] = []string{Drive, Vaccsuit, Broker, Steward, Comms, Persuade}
	SaT["B0=17ae"] = []string{ScienceSocial, Astrogation, Computers, Pilot, Admin, Advocate}
	SaT["B0=17os"] = []string{}
	SaT["B0=17=1"] = []string{pickAnyOf(PilotSpacecraft, PilotCapitalships), Vaccsuit, AthleticsDEX, Mechanic, Engineer, Gunner}
	SaT["B0=17=1ma"] = []string{"INT 4", "EDU 5", "INT 7"}
	SaT["B0=17=2"] = []string{PilotSpacecraft, Vaccsuit, AthleticsDEX, Mechanic, Engineer, Sensors}
	SaT["B0=17=2ma"] = []string{"INT 4", "DEX 6", "INT 6"}
	SaT["B0=17=3"] = []string{Admin, Advocate, Broker, Streetwise, Deception, Persuade}
	SaT["B0=17=3ma"] = []string{"INT 4", "EDU 5", "INT 7"}

	SaT["B0=18pd"] = []string{STR, DEX, END, INT, EDU, SOC}
	SaT["B0=18ss"] = []string{Pilot, Vaccsuit, AthleticsDEX, Gunner, Mechanic, GunCombat}
	SaT["B0=18ae"] = []string{RemoteOps, Astrogation, Engineer, Computers, Navigation, Admin}
	SaT["B0=18os"] = []string{Leadership, TacticsNavy, Pilot, MeleeBlade, Admin, TacticsNavy}
	SaT["B0=18=1"] = []string{Comms, Mechanic, GunCombat, Sensors, Melee, Vaccsuit}
	SaT["B0=18=1ma"] = []string{"INT 6", "INT 5", "EDU 7"}
	SaT["B0=18=2"] = []string{Engineer, Mechanic, Sensors, Engineer, Gunner, Computers}
	SaT["B0=18=2ma"] = []string{"INT 6", "INT 6", "EDU 6"}
	SaT["B0=18=3"] = []string{Pilot, Flyer, Gunner, PilotSmallcraft, Astrogation, AthleticsDEX}
	SaT["B0=18=3ma"] = []string{"INT 6", "DEX 7", "EDU 5"}

	SaT["B0=19pd"] = []string{Carouse, EDU, Deception, DEX, MeleeBlade, SOC}
	SaT["B0=19ss"] = []string{Admin, Advocate, Comms, Diplomat, Investigate, Persuade}
	SaT["B0=19ae"] = []string{Admin, Advocate, Language, Leadership, Diplomat, Computers}
	SaT["B0=19os"] = []string{}
	SaT["B0=19=1"] = []string{Admin, Advocate, Broker, Diplomat, Leadership, Persuade}
	SaT["B0=19=1ma"] = []string{"SOC 10", "INT 4", "EDU 6"}
	SaT["B0=19=2"] = []string{Advocate, Carouse, Comms, Steward, Diplomat, Deception}
	SaT["B0=19=2ma"] = []string{"SOC 10", "INT 5", "SOC 7"}
	SaT["B0=19=3"] = []string{Carouse, Deception, Flyer, Streetwise, Gambler, Jackofalltrades}
	SaT["B0=19=3ma"] = []string{"SOC 10", "SOC 3", "INT 8"}

	SaT["B0=1Apd"] = []string{Carouse, DEX, END, Gambler, Melee, GunCombat}
	SaT["B0=1Ass"] = []string{Deception, Recon, Athletics, GunCombat, Stealth, Streetwise}
	SaT["B0=1Aae"] = []string{Computers, Comms, Medic, Investigate, Persuade, Advocate}
	SaT["B0=1Aos"] = []string{}
	SaT["B0=1A=1"] = []string{Stealth, Computers, RemoteOps, Streetwise, Deception, AthleticsDEX}
	SaT["B0=1A=1ma"] = []string{"DEX 6", "INT 6", "DEX 6"}
	SaT["B0=1A=2"] = []string{GunCombat, Melee, Streetwise, Persuade, Athletics, Drive}
	SaT["B0=1A=2ma"] = []string{"DEX 6", "END 6", "STR 6"}
	SaT["B0=1A=3"] = []string{Pilot, Astrogation, Gunner, Engineer, Vaccsuit, Melee}
	SaT["B0=1A=3ma"] = []string{"DEX 6", "DEX 6", "INT 6"}

	SaT["B0=1Bpd"] = []string{INT, EDU, SOC, DEX, END, Computers}
	SaT["B0=1Bss"] = []string{Comms, Computers, Diplomat, Medic, Investigate, ScienceAny}
	SaT["B0=1Bae"] = []string{Art, Advocate, Computers, Language, Engineer, ScienceAny}
	SaT["B0=1Bos"] = []string{}
	SaT["B0=1B=1"] = []string{Sensors, Diplomat, Language, Survival, Investigate, ScienceAny}
	SaT["B0=1B=1ma"] = []string{"INT 6", "END 6", "INT 6"}
	SaT["B0=1B=2"] = []string{Admin, Engineer, ScienceAny, Sensors, Computers, ScienceAny}
	SaT["B0=1B=2ma"] = []string{"INT 6", "EDU 4", "INT 8"}
	SaT["B0=1B=3"] = []string{Medic, Comms, Investigate, Medic, Persuade, ScienceAny}
	SaT["B0=1B=3ma"] = []string{"INT 6", "EDU 4", "EDU 8"}

	SaT["B0=1Cpd"] = []string{STR, DEX, END, INT, EDU, Jackofalltrades}
	SaT["B0=1Css"] = []string{pickAnyOf(PilotSpacecraft, PilotSmallcraft), Survival, Mechanic, Astrogation, Comms, GunCombat}
	SaT["B0=1Cae"] = []string{Medic, Navigation, Engineer, Computers, ScienceSpace, Jackofalltrades}
	SaT["B0=1Cos"] = []string{}
	SaT["B0=1C=1"] = []string{Comms, Sensors, PilotSpacecraft, Vaccsuit, AthleticsDEX, Astrogation}
	SaT["B0=1C=1ma"] = []string{"INT 5", "END 5", "EDU 9"}
	SaT["B0=1C=2"] = []string{Sensors, Persuade, PilotSmallcraft, Navigation, Diplomat, Streetwise}
	SaT["B0=1C=2ma"] = []string{"INT 5", "END 6", "INT 8"}
	SaT["B0=1C=3"] = []string{Sensors, PilotSpacecraft, PilotSmallcraft, ScienceLife, Stealth, Recon}
	SaT["B0=1C=3ma"] = []string{"INT 5", "END 7", "EDU 7"}

	SaT["B1=11pd"] = []string{SOC, DEX, END, EDU, STR, Carouse}
	SaT["B1=11ss"] = []string{Investigate, Streetwise, Admin, GunCombat, Athletics, Drive}
	SaT["B1=11ae"] = []string{Advocate, Comms, Computers, Drive, Investigate, Diplomat}
	SaT["B1=11os"] = []string{}
	SaT["B1=11=1"] = []string{Drive, Melee, Athletics, Recon, Streetwise, Investigate}
	SaT["B1=11=1ma"] = []string{"INT 5", "INT 7", "EDU 7"}
	SaT["B1=11=2"] = []string{GunCombat, Stealth, TacticsMilitary, Recon, HeavyWeapon, Vaccsuit}
	SaT["B1=11=2ma"] = []string{"INT 5", "END 8", "INT 6"}
	SaT["B1=11=3"] = []string{pickAnyOf(Pilot, Seafarer), Sensors, Comms, Recon, GunnerTurret, Vaccsuit}
	SaT["B1=11=3ma"] = []string{"INT 5", "DEX 6", "EDU 8"}
}

func pickAnyOf(list ...string) string {
	sl := []string{}
	for _, opt := range list {
		sl = append(sl, opt)
	}
	return utils.RandomFromList(sl)
}

func (trv *Traveller) getCareerCodes() {
	var sat []string
	data := strings.Split(trv.career, "=")
	book, career, spec := data[0], data[1], data[2]
	ae := "---"
	co := "---"
	edu, err := trv.characteristics.GetValue(EDU)
	soc, err := trv.characteristics.GetValue(SOC)
	if err != nil {
		fmt.Println(err)
	}
	if edu >= 8 {
		ae = "ae"
	}
	if soc >= 8 {
		co = "co"
	}
	sat = append(sat, SaT[book+"="+career+"pd"]...)
	sat = append(sat, SaT[book+"="+career+"ss"]...)
	sat = append(sat, SaT[book+"="+career+ae]...)
	sat = append(sat, SaT[book+"="+career+co]...)
	sat = append(sat, SaT[book+"="+career+"="+spec]...)
	//for i := range sat {
	//fmt.Println(sat[i])
	//trv.Learn(sat[i])
	//}
	for _, data := range SaT[book+"="+career+"="+spec+"ma"] {
		dataParts := strings.Split(data, " ")
		val := chrFullCode(dataParts[0])

		num, err := strconv.Atoi(dataParts[1])
		if err != nil {
			fmt.Println(err)
		}
		chr, err2 := trv.characteristics.GetValue(val)
		if err2 != nil {
			fmt.Println(err2)
		}
		//fmt.Println("Ensure", entity.GetFromCode(3, val), dataParts[1]+"+")
		if num > chr {
			trv.characteristics.Set(val, num)
			//	fmt.Println("Initial Training:", entity.GetFromCode(3, val))

		}
		//fmt.Println("Initial Training:", entity.GetFromCode(3, val))
	}

	for _, val := range SaT[book+"="+career+"ss"] {
		trv.Learn(val)
		//fmt.Println("Basic Training:", entity.GetFromCode(3, val))
	}
	//fmt.Println("Stop Basic Training")

	for i := 0; i < trv.term; i++ {
		val := utils.RandomFromList(sat)
		trv.Learn(val)
		if dice.Roll("2d6").Sum() > 6 {
			val := utils.RandomFromList(sat)
			trv.Learn(val)
			trv.rank++
		}
		//	fmt.Println("Career Training:", entity.GetFromCode(3, val))
	}
}

func (trv *Traveller) Learn(code string) {
	//fmt.Println("Learn:", code)
	for _, val := range entity.SkillCodesList() {
		if strings.Contains(val, code) {
			trv.skills.Train(code)
			return
		}
	}
	for _, val := range entity.CharacteristicsCodesList() {
		if strings.Contains(val, code) {
			trv.characteristics.Train(code)
			return
		}
	}
}

func chrFullCode(code string) string {
	switch code {
	default:
		return ""
	case "STR":
		return STR
	case "DEX":
		return DEX
	case "END":
		return END
	case "INT":
		return INT
	case "EDU":
		return EDU
	case "SOC":
		return SOC
	}
}
