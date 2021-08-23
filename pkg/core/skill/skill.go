package skill

import (
	"errors"
	"strconv"
)

const (
	entityTraveller = 0
	//
	Admin = iota
	Advocate
	Animals
	Animals_Handling
	Animals_Veterinary
	Animals_Training
	Art
	Art_Holography
	Art_Instrument
	Art_Performer
	Art_VisualMedia
	Art_Write
	Astrogation
	Athletics
	Athletics_DEX
	Athletics_END
	Athletics_STR
	Broker
	Carouse
	Deception
	Diplomat
	Drive
	Drive_Hovercraft
	Drive_Mole
	Drive_Track
	Drive_Walker
	Drive_Wheel
	Electronics
	Electronics_Comms
	Electronics_Computers
	Electronics_Remoteops
	Electronics_Sensors
	Engineer
	Engineer_Jdrive
	Engineer_Lifesupport
	Engineer_Mdrive
	Engineer_Power
	Explosives
	Flyer
	Flyer_Airship
	Flyer_Grav
	Flyer_Ornithopter
	Flyer_Rotor
	Flyer_Wing
	Gambler
	Guncombat
	Guncombat_Archaic
	Guncombat_Energy
	Guncombat_Slug
	Gunner
	Gunner_Capital
	Gunner_Ortilery
	Gunner_Screen
	Gunner_Turret
	Heavyweapon
	Heavyweapon_Artilery
	Heavyweapon_Manportable
	Heavyweapon_Vehicle
	Investigate
	Jack_of_all_trades
	Language
	Language_Anglic
	Leadership
	Mechanic
	Medic
	Melee
	Melee_Blade
	Melee_Bludgeon
	Melee_Natural
	Melee_Unarmed
	Navigation
	Persuade
	Pilot
	Pilot_CapitalShips
	Pilot_SmallCraft
	Pilot_Spacecraft
	Profession
	Profession_Any
	Recon
	Science_Physical
	Science_Life
	Science_Social
	Science_Space
	Science
	Science_Any
	Seafarer
	Seafarer_Oceanships
	Seafarer_Personal
	Seafarer_Sail
	Seafarer_Submarine
	Stealth
	Steward
	Streetwise
	Survival
	Tactics
	Tactics_Military
	Tactics_Navy
	VaccSuit
)

type Skill struct {
	Code       int //ID скила
	entity     int //сущность использующая этот скил
	Group      string
	Speciality string
	//name       string //Метод интерфейса?
	//nameShort  string
	Descr   string
	Value   int
	trained bool
	hasSpec bool
}

func ByCode(code int) Skill {
	s, _ := new(code)
	return s
}

func ByGroup(grp string) (list []Skill) {
	for i := 0; i < 150; i++ {
		skl, err := new(i)
		if err != nil {
			continue
		}
		if skl.Group == grp {
			list = append(list, skl)
		}
	}
	return list
}

func CodeFromName(name string) int {
	for i := 0; i < 1000; i++ {
		sk, err := new(i)
		if err != nil {
			return -1
		}
		if sk.Name() == name {
			return i
		}
	}
	return -2
}

func new(skillcode int) (Skill, error) {
	s := Skill{}
	switch skillcode {
	default:
		return s, errors.New("unknown skillcode=" + strconv.Itoa(skillcode))
	case Admin:
		s = Skill{
			Code:   Admin,
			entity: entityTraveller,
			Group:  "Admin",
			Descr:  "This Skill covers bureaucracies and administration of all sorts, including the navigation of bureaucratic obstacles or disasters. It also covers tracking inventories, ship manifests and other records.",
		}
	case Advocate:
		s = Skill{
			Code:   Advocate,
			entity: entityTraveller,
			Group:  "Advocate",
			Descr:  "Advocate gives a knowledge of common legal codes and practises, especially interstellar law. It also gives the Traveller experience in oratory, debate and public speaking, making it an excellent Skill for lawyers and politicians.",
		}
	case Animals:
		s = Skill{
			Code:       Animals,
			entity:     entityTraveller,
			Group:      "Animals",
			Speciality: "",
			Descr:      "This Skill, rare on industrialised or technologically advanced worlds, is for the care of animals.",
			hasSpec:    true,
		}
	case Animals_Handling:
		s = Skill{
			Code:       Animals_Handling,
			entity:     entityTraveller,
			Group:      "Animals",
			Speciality: "Handling",
			Descr:      "The Traveller knows how to handle an animal and ride those trained to bear a rider. Unusual animals raise the difficulty of the check.",
		}
	case Animals_Veterinary:
		s = Skill{
			Code:       Animals_Veterinary,
			entity:     entityTraveller,
			Group:      "Animals",
			Speciality: "Veterinary",
			Descr:      "The Traveller is trained in veterinary medicine and animal care.",
		}
	case Animals_Training:
		s = Skill{
			Code:       Animals_Training,
			entity:     entityTraveller,
			Descr:      "The Traveller knows how to tame and train animals.",
			Group:      "Animals",
			Speciality: "Training",
		}
	case Art:
		s = Skill{
			Code:       Art,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "",
			Descr:      "The Traveller is trained in a type of creative art.",
			hasSpec:    true,
		}
	case Art_Holography:
		s = Skill{
			Code:       Art_Holography,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Holography",
			Descr:      "Recording and producing aesthetically pleasing and clear holographic images.",
		}
	case Art_Instrument:
		s = Skill{
			Code:       Art_Instrument,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Instrument",
			Descr:      "Playing a particular musical instrument, such a flute, piano or organ.",
		}
	case Art_Performer:
		s = Skill{
			Code:       Art_Performer,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Performer",
			Descr:      "The Traveller is a trained actor, dancer or singer at home on the stage, screen or holo.",
		}
	case Art_VisualMedia:
		s = Skill{
			Code:       Art_VisualMedia,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Visual Media",
			Descr:      "Making artistic or abstract paintings or sculptures in a variety of media.",
		}
	case Art_Write:
		s = Skill{
			Code:       Art_Write,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Write",
			Descr:      "Composing inspiring or interesting pieces of text.",
		}
	case Astrogation:
		s = Skill{
			Code:   Astrogation,
			entity: entityTraveller,
			Group:  "Astrogation",
			Descr:  "This Skill is for plotting the courses of starships and calculating accurate jumps.",
		}
	case Athletics:
		s = Skill{
			Code:    Athletics,
			entity:  entityTraveller,
			Group:   "Athletics",
			Descr:   "The Traveller is a trained athlete and is physically fit.",
			hasSpec: true,
		}
	case Athletics_DEX:
		s = Skill{
			Code:       Athletics_DEX,
			entity:     entityTraveller,
			Group:      "Athletics",
			Speciality: "DEX",
			Descr:      "Climbing, Juggling, Throwing. For alien races with wings, this also includes flying.",
		}
	case Athletics_END:
		s = Skill{
			Code:       Athletics_END,
			entity:     entityTraveller,
			Group:      "Athletics",
			Speciality: "END",
			Descr:      "Long-distance running, hiking.",
		}
	case Athletics_STR:
		s = Skill{
			Code:       Athletics_STR,
			entity:     entityTraveller,
			Group:      "Athletics",
			Speciality: "STR",
			Descr:      "Feats of strength, weight-lifting.",
		}
	case Broker:
		s = Skill{
			Code:       Broker,
			entity:     entityTraveller,
			Group:      "Broker",
			Speciality: "",
			Descr:      "The Broker Skill allows a Traveller to negotiate trades and arrange fair deals. It is heavily used when trading.",
		}
	case Carouse:
		s = Skill{
			Code:       Carouse,
			entity:     entityTraveller,
			Group:      "Carouse",
			Speciality: "",
			Descr:      "Carousing is the art of socialising; having fun, but also ensuring other people have fun, and infectious good humour. It also covers social awareness and subterfuge in such situations.",
		}
	case Deception:
		s = Skill{
			Code:       skillcode,
			entity:     entityTraveller,
			Group:      "Deception",
			Speciality: "",
			Descr:      "Deception allows a Traveller to lie fluently, disguise himself, perform sleight of hand and fool onlookers. Most underhanded ways of cheating and lying fall under deception.",
		}
	case Diplomat:
		s = Skill{
			Code:       Diplomat,
			entity:     entityTraveller,
			Group:      "Diplomat",
			Speciality: "",
			Descr:      "The Diplomat Skill is for negotiating deals, establishing peaceful contact and smoothing over social faux pas. It includes how to behave in high society and proper ways to address nobles. It is a much more formal Skill than Persuade.",
		}
	case Drive:
		s = Skill{
			Code:       Drive,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "",
			Descr:      "This Skill is for controlling ground vehicles of various types. There are several specialities.",
			hasSpec:    true,
		}
	case Drive_Hovercraft:
		s = Skill{
			Code:       Drive_Hovercraft,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Hovercraft",
			Descr:      "Vehicles that rely on a cushion of air and thrusters for motion.",
		}
	case Drive_Mole:
		s = Skill{
			Code:       Drive_Mole,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Mole",
			Descr:      "For controlling vehicles that move through solid matter using drills or other earth-moving technologies, such as plasma torches or cavitation.",
		}
	case Drive_Track:
		s = Skill{
			Code:       Drive_Track,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Track",
			Descr:      "For tanks and other vehicles that move on tracks.",
		}
	case Drive_Walker:
		s = Skill{
			Code:       Drive_Walker,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Walker",
			Descr:      "Vehicles that use two or more legs to manoeuvre.",
		}
	case Drive_Wheel:
		s = Skill{
			Code:       Drive_Wheel,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Wheel",
			Descr:      "Vehicles that use two or more legs to manoeuvre.",
		}
	case Electronics:
		s = Skill{
			Code:       Electronics,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "",
			Descr:      "This Skill is used to operate electronic devices such as computers and ship-board systems. Higher levels represent the ability to repair and create electronic devices and systems. There are several specialities.",
			hasSpec:    true,
		}
	case Electronics_Comms:
		s = Skill{
			Code:       Electronics_Comms,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Comms",
			Descr:      "The use of modern telecommunications; opening communications channels, querying computer networks, jamming signals and so on, as well as the proper protocols for communicating with starports and other spacecraft.",
		}
	case Electronics_Computers:
		s = Skill{
			Code:       Electronics_Computers,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Computers",
			Descr:      "Using and controlling computer systems, and similar electronics and electrics.",
		}
	case Electronics_Remoteops:
		s = Skill{
			Code:       Electronics_Remoteops,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Remote Ops",
			Descr:      "Using telepresence to remotely control drones, missiles, robots and other devices.",
		}
	case Electronics_Sensors:
		s = Skill{
			Code:       Electronics_Sensors,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Sensors",
			Descr:      "The use and interpretation of data from electronic sensor devices, from observation satellites and remote probes to thermal imaging and densitometers.",
		}
	case Engineer:
		s = Skill{
			Code:       Engineer,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "",
			Descr:      "The Engineer Skill is used to operate and maintain spacecraft and advanced vehicles. Engineer can be used to make repairs on damaged systems on spacecraft and advanced vehicles. For repairs on simpler machines and systems, use the Mechanic Skill.",
			hasSpec:    true,
		}
	case Engineer_Jdrive:
		s = Skill{
			Code:       Engineer_Jdrive,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "J-drive",
			Descr:      "Maintaining and operating a spacecraft's Jump drive.",
		}
	case Engineer_Lifesupport:
		s = Skill{
			Code:       Engineer_Lifesupport,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "Life Support",
			Descr:      "Covers oxygen generators, heating and lighting and other necessary life support systems.",
		}
	case Engineer_Mdrive:
		s = Skill{
			Code:       Engineer_Mdrive,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "M-drive",
			Descr:      "Maintaining and operating a spacecraft's manoeuvre drive, as well as its artificial gravity.",
		}
	case Engineer_Power:
		s = Skill{
			Code:       Engineer_Power,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "Power",
			Descr:      "Maintaining and operating a spacecraft's power plant.",
		}
	case Explosives:
		s = Skill{
			Code:       Explosives,
			entity:     entityTraveller,
			Group:      "Explosives",
			Speciality: "",
			Descr:      "The Explosives Skill covers the use of demolition charges and other explosive devices, including assembling or disarming bombs. A failed Explosives check with an Effect of -4 or less can result in a bomb detonating prematurely.",
		}
	case Flyer:
		s = Skill{
			Code:       Flyer,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "",
			Descr:      "The various specialities of this Skill cover different types of flying vehicles. Flyers only work in an atmosphere; vehicles that can leave the atmosphere and enter orbit generally use the Pilot Skill.",
			hasSpec:    true,
		}
	case Flyer_Airship:
		s = Skill{
			Code:       Flyer_Airship,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Airship",
			Descr:      "Used for airships, dirigibles and other powered lighter than air craft.",
		}
	case Flyer_Grav:
		s = Skill{
			Code:       Flyer_Grav,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Grav",
			Descr:      "This covers air/rafts, grav belts and other vehicles that use gravitic technology.",
		}
	case Flyer_Ornithopter:
		s = Skill{
			Code:       Flyer_Ornithopter,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Ornithopter",
			Descr:      "For vehicles that fly through the use of flapping wings.",
		}
	case Flyer_Rotor:
		s = Skill{
			Code:       Flyer_Rotor,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Rotor",
			Descr:      "For helicopters, tilt-rotors and aerodynes.",
		}
	case Flyer_Wing:
		s = Skill{
			Code:       Flyer_Wing,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Wing",
			Descr:      "For jets, vectored thrust aircraft and aeroplanes using a lifting body.",
		}
	case Gambler:
		s = Skill{
			Code:       Gambler,
			entity:     entityTraveller,
			Group:      "Gambler",
			Speciality: "",
			Descr:      "The Traveller is familiar with a wide variety of gambling games, such as poker, roulette, blackjack, horse-racing, sports betting and so on, and has an excellent grasp of statistics and probability. Gambler increases the rewards from Benefit Rolls, giving the Traveller DM+1 to his cash rolls if he has Gambler 1 or better.",
		}
	case Guncombat:
		s = Skill{
			Code:       Guncombat,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "",
			Descr:      "The Gun Combat Skill covers a variety of ranged weapons.",
			hasSpec:    true,
		}
	case Guncombat_Archaic:
		s = Skill{
			Code:       Guncombat_Archaic,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "Archaic",
			Descr:      "For primitive weapons that are not thrown, such as bows and blowpipes.",
		}
	case Guncombat_Energy:
		s = Skill{
			Code:       Guncombat_Energy,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "Energy",
			Descr:      "Using advanced energy weapons like laser pistols or plasma rifles.",
		}
	case Guncombat_Slug:
		s = Skill{
			Code:       Guncombat_Slug,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "Slug",
			Descr:      "Weapons that fire a solid projectile such as the autorifle or gauss rifle.",
		}
	case Gunner:
		s = Skill{
			Code:       Gunner,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "",
			Descr:      "The various specialities of this Skill deal with the operation of ship-mounted weapons in space combat. See Spacecraft Operations chapter for more details. Most Travellers have smaller ships equipped solely with turret weapons.",
			hasSpec:    true,
		}
	case Gunner_Capital:
		s = Skill{
			Code:       Gunner_Capital,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Capital",
			Descr:      "Operating bay or spinal mount weapons on board a ship.",
		}
	case Gunner_Ortilery:
		s = Skill{
			Code:       Gunner_Ortilery,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Ortilery",
			Descr:      "A contraction of Orbital Artillery – using a ship's weapons for planetary bombardment or attacks on stationary targets.",
		}
	case Gunner_Screen:
		s = Skill{
			Code:       Gunner_Screen,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Screen",
			Descr:      "Activating and using a ship's energy screens like Black Globe generators or meson screens.",
		}
	case Gunner_Turret:
		s = Skill{
			Code:       Gunner_Turret,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Turret",
			Descr:      "Operating turret-mounted weapons on board a ship.",
		}
	case Heavyweapon:
		s = Skill{
			Code:       Heavyweapon,
			entity:     entityTraveller,
			Group:      "Heavy weapon",
			Speciality: "",
			Descr:      "The Heavy Weapons Skill covers man-portable and larger weapons that cause extreme property damage, such as rocket launchers, artillery and large plasma weapons.",
			hasSpec:    true,
		}
	case Heavyweapon_Artilery:
		s = Skill{
			Code:       Heavyweapon_Artilery,
			entity:     entityTraveller,
			Group:      "Heavy weapon",
			Speciality: "Artilery",
			Descr:      "Fixed guns, mortars and other indirect-fire weapons.",
		}
	case Heavyweapon_Manportable:
		s = Skill{
			Code:       Heavyweapon_Manportable,
			entity:     entityTraveller,
			Group:      "Heavy Weapon",
			Speciality: "Man Portable",
			Descr:      "Missile launchers, flamethrowers and man portable fusion and plasma.",
		}
	case Heavyweapon_Vehicle:
		s = Skill{
			Code:       Heavyweapon_Vehicle,
			entity:     entityTraveller,
			Group:      "Heavy Weapon",
			Speciality: "Vehicle",
			Descr:      "Large weapons typically mounted on vehicles or strongpoints such as tank guns and autocannon.",
		}
	case Investigate:
		s = Skill{
			Code:       Investigate,
			entity:     entityTraveller,
			Group:      "Investigate",
			Speciality: "",
			Descr:      "The Investigate Skill incorporates keen observation, forensics, and detailed analysis.",
		}
	case Jack_of_all_trades:
		s = Skill{
			Code:       Jack_of_all_trades,
			entity:     entityTraveller,
			Group:      "Jack-of-All-Trades",
			Speciality: "",
			Descr:      "The Jack-of-All-Trades Skill works differently to other skills. It reduces the unskilled penalty a Traveller receives for not having the appropriate Skill by one for every level of Jack-of-All-Trades. For example, if a Traveller does not have the Pilot Skill, he suffers DM-3 to all Pilot checks. If that Traveller has Jack-of-All-Trades 2, then the penalty is reduced by 2 to DM-1. With Jackof-All-Trades 3, a Traveller can totally negate the penalty for being unskilled.\nThere is no benefit for having Jack-of-All-Trades 0 or Jack-of-All-Trades 4 or more.",
		}
	case Language:
		s = Skill{
			Code:       Language,
			entity:     entityTraveller,
			Group:      "Language",
			Speciality: "",
			Descr:      "There are numerous different Language specialities, each one covering reading and writing a different language. All Travellers can speak and read their native language without needing the Language Skill, and automated computer translator programs mean Language skills are not always needed on other worlds. Having Language 0 implies the Traveller has a smattering of simple phrases in several languages.",
		}
	case Language_Anglic:
		s = Skill{
			Code:       Language_Anglic,
			entity:     entityTraveller,
			Group:      "Language",
			Speciality: "Anglic",
			Descr:      "The common trade language of the Third Imperium, derived originally from the English spoken in the Rule of Man",
		}
	case Leadership:
		s = Skill{
			Code:       Leadership,
			entity:     entityTraveller,
			Group:      "Leadership",
			Speciality: "",
			Descr:      "The Leadership Skill is for directing, inspiring and rallying allies and comrades. A Traveller may make a Leadership action in combat, as detailed on page 72 (CRB).",
		}
	case Mechanic:
		s = Skill{
			Code:       Mechanic,
			entity:     entityTraveller,
			Group:      "Mechanic",
			Speciality: "",
			Descr:      "The Mechanic Skill allows a Traveller to maintain and repair most equipment – some advanced equipment and spacecraft components require the Engineer Skill. Unlike the narrower and more focussed Engineer or Science skills, Mechanic does not allow a Traveller to build new devices or alter existing ones – it is purely for repairs and maintenance but covers all types of equipment.",
		}
	case Medic:
		s = Skill{
			Code:       Medic,
			entity:     entityTraveller,
			Group:      "Medic",
			Speciality: "",
			Descr:      "The Medic Skill covers emergency first aid and battlefield triage as well as diagnosis, treatment, surgery and longterm care. See Injury and Recovery on page 47.",
		}
	case Melee:
		s = Skill{
			Code:       Melee,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "",
			Descr:      "The Melee Skill covers attacking in hand-to-hand combat and the use of suitable weapons.",
			hasSpec:    true,
		}
	case Melee_Blade:
		s = Skill{
			Code:       Melee_Blade,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Blade",
			Descr:      "Attacking with swords, rapiers, blades and other edged weapons.",
		}
	case Melee_Bludgeon:
		s = Skill{
			Code:       Melee_Bludgeon,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Bludgeon",
			Descr:      "Attacking with maces, clubs, staves and so on.",
		}
	case Melee_Natural:
		s = Skill{
			Code:       Melee_Natural,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Natural",
			Descr:      " Weapons that are part of an alien or creature, such as claws or teeth.",
		}
	case Melee_Unarmed:
		s = Skill{
			Code:       Melee_Unarmed,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Unarmed",
			Descr:      "Punching, kicking and wrestling; using improvised weapons in a bar brawl.",
		}
	case Navigation:
		s = Skill{
			Code:       Navigation,
			entity:     entityTraveller,
			Group:      "Navigation",
			Speciality: "",
			Descr:      "Navigation is the planetside counterpart of astrogation, covering plotting courses and finding directions on the ground.",
		}
	case Persuade:
		s = Skill{
			Code:       Persuade,
			entity:     entityTraveller,
			Group:      "Persuade",
			Speciality: "",
			Descr:      "Persuade is a more casual, informal version of Diplomat. It covers fast talking, bargaining, wheedling and bluffing. It also covers bribery or intimidation.",
		}
	case Pilot:
		s = Skill{
			Code:       Pilot,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "",
			Descr:      "The Pilot Skill specialities cover different forms of  spacecraft. See Spacecraft Operations chapter for more  details.",
			hasSpec:    true,
		}
	case Pilot_CapitalShips:
		s = Skill{
			Code:       Pilot_CapitalShips,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "Capital Ships",
			Descr:      "Battleships and other ships over 5,000 tons.",
		}
	case Pilot_SmallCraft:
		s = Skill{
			Code:       Pilot_SmallCraft,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "Small Craft",
			Descr:      "Shuttles and other craft under 100 tons.",
		}
	case Pilot_Spacecraft:
		s = Skill{
			Code:       Pilot_Spacecraft,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "Spacecraft",
			Descr:      ": Trade ships and other vessels between 100 and 5,000 tons.",
		}
	case Profession:
		s = Skill{
			Code:       Profession,
			entity:     entityTraveller,
			Group:      "Profession",
			Speciality: "",
			Descr:      "A Traveller with a Profession Skill is trained in producing useful goods or services. There are many different Profession specialities.",
			hasSpec:    true,
		}
	case Profession_Any:
		s = Skill{}
	case Recon:
		s = Skill{
			Code:       Recon,
			entity:     entityTraveller,
			Group:      "Recon",
			Speciality: "",
			Descr:      "A Traveller trained in Recon is able to scout out dangers and spot threats, unusual objects or out of place people. ",
		}
	case Science_Physical:
		s = Skill{
			Code:       Science_Physical,
			entity:     entityTraveller,
			Group:      "Science Physical",
			Speciality: "",
		}
	case Science_Life:
		s = Skill{
			Code:       Science_Life,
			entity:     entityTraveller,
			Group:      "Science Life",
			Speciality: "",
		}
	case Science_Social:
		s = Skill{
			Code:       Science_Social,
			entity:     entityTraveller,
			Group:      "Science_Social",
			Speciality: "",
		}
	case Science_Space:
		s = Skill{
			Code:       Science_Space,
			entity:     entityTraveller,
			Group:      "Science Space",
			Speciality: "",
		}
	case Science:
		s = Skill{
			Code:       Science,
			entity:     entityTraveller,
			Group:      "Science",
			Speciality: "",
			Descr:      "The Science Skill covers not just knowledge but also practical application of that knowledge where such practical application is possible. There are a large range of specialities.",
		}
	case Science_Any:
		s = Skill{}
	case Seafarer:
		s = Skill{
			Code:       Seafarer,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "",
			Descr:      "The Seafarer Skill covers all manner of watercraft and ocean travel.",
			hasSpec:    true,
		}
	case Seafarer_Oceanships:
		s = Skill{
			Code:       Seafarer_Oceanships,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Ocean Ships",
			Descr:      "For motorised sea-going vessels.",
		}
	case Seafarer_Personal:
		s = Skill{
			Code:       Seafarer_Personal,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Personal",
			Descr:      "Used for very small waterborne craft such as canoes and rowboats.",
		}
	case Seafarer_Sail:
		s = Skill{
			Code:       Seafarer_Sail,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Sail",
			Descr:      "This Skill is for wind-driven watercraft.",
		}
	case Seafarer_Submarine:
		s = Skill{
			Code:       Seafarer_Submarine,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Submarine",
			Descr:      "For vehicles that travel underwater.",
		}
	case Stealth:
		s = Skill{
			Code:       Stealth,
			entity:     entityTraveller,
			Group:      "Stealth",
			Speciality: "",
			Descr:      "A Traveller trained in the Stealth Skill is adept at staying unseen, unheard, and unnoticed.",
		}
	case Steward:
		s = Skill{
			Code:       skillcode,
			entity:     entityTraveller,
			Group:      "",
			Speciality: "",
			Descr:      "",
		}
	case Streetwise:
		s = Skill{
			Code:       Streetwise,
			entity:     entityTraveller,
			Group:      "Streetwise",
			Speciality: "",
			Descr:      "The Steward Skill allows the Traveller to serve and care for nobles and high-class passengers. It covers everything from proper address and behaviour to cooking and tailoring, as well as basic management skills. A Traveller with the Steward Skill is necessary on any ship offering High Passage. See Spacecraft Operations chapter for more details.",
		}
	case Survival:
		s = Skill{
			Code:       Survival,
			entity:     entityTraveller,
			Group:      "Survival",
			Speciality: "",
			Descr:      "",
		}
	case Tactics:
		s = Skill{
			Code:       Tactics,
			entity:     entityTraveller,
			Group:      "Tactics",
			Speciality: "",
			Descr:      "This Skill covers tactical planning and decision making, from board games to squad level combat to fleet engagements. For use in combat, see Combat chapter",
			hasSpec:    true,
		}
	case Tactics_Military:
		s = Skill{
			Code:       Tactics_Military,
			entity:     entityTraveller,
			Group:      "Tactics",
			Speciality: "Military",
			Descr:      "Co-ordinating the attacks of foot troops or vehicles on the ground.",
		}
	case Tactics_Navy:
		s = Skill{
			Code:       Tactics_Navy,
			entity:     entityTraveller,
			Group:      "Tactics",
			Speciality: "Navy",
			Descr:      "Co-ordinating the attacks of a spacecraft or fleet.",
		}
	case VaccSuit:
		s = Skill{
			Code:       VaccSuit,
			entity:     entityTraveller,
			Group:      "Vacc Suit",
			Speciality: "",
			Descr:      "The Vacc Suit Skill allows a Traveller to wear and operate spacesuits and environmental suits. A Traveller will rarely need to make Vacc Suit checks under ordinary circumstances – merely possessing the Skill is enough. If the Traveller does not have the requisite Vacc Suit Skill for the suit he is wearing, he suffers DM-2 to all Skill checks made while wearing a suit for each missing level. This Skill also permits the character to operate advanced battle armour. ",
		}
	}
	return s, nil
}

/*
fmt.Print(char.Skills[Skill.Advocate].Description())
entity:
	traveller
		characteristic
		Skill
		trait
	dynasty
		characteristic

	firm
*/

func (s Skill) Name() string {
	name := s.Group
	if s.Speciality != "" {
		name += " (" + s.Speciality + ")"
	}
	return name
}

func (s Skill) String() string {
	return s.Name() + " " + strconv.Itoa(s.Value)
}

// func (s Skill) Increase() {
// 	s.Value++
// }

func BackgroundSkills() []int {
	return []int{
		Admin,
		Animals,
		Art,
		Athletics,
		Carouse,
		Drive,
		Electronics,
		Flyer,
		Language,
		Mechanic,
		Medic,
		Profession,
		Science,
		Seafarer,
		Streetwise,
		Survival,
		VaccSuit,
	}
}

func NameIsValid(name string) bool {
	if CodeFromName(name) == -2 {
		return false
	}
	return true
}
