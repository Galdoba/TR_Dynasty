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

type skill struct {
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

type Skill interface {
	Description() string
	Name() string
}

func ByCode(code int) Skill {
	s, _ := new(code)
	return &s
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
	return -1
}

func new(skillcode int) (skill, error) {
	s := skill{}
	switch skillcode {
	default:
		return s, errors.New("unknown skillcode=" + strconv.Itoa(skillcode))
	case Admin:
		s = skill{
			Code:   Admin,
			entity: entityTraveller,
			Group:  "Admin",
			Descr:  "This skill covers bureaucracies and administration of all sorts, including the navigation of bureaucratic obstacles or disasters. It also covers tracking inventories, ship manifests and other records.",
		}
	case Advocate:
		s = skill{
			Code:   Advocate,
			entity: entityTraveller,
			Group:  "Advocate",
			Descr:  "Advocate gives a knowledge of common legal codes and practises, especially interstellar law. It also gives the Traveller experience in oratory, debate and public speaking, making it an excellent skill for lawyers and politicians.",
		}
	case Animals:
		s = skill{
			Code:       Animals,
			entity:     entityTraveller,
			Group:      "Animals",
			Speciality: "",
			Descr:      "This skill, rare on industrialised or technologically advanced worlds, is for the care of animals.",
			hasSpec:    true,
		}
	case Animals_Handling:
		s = skill{
			Code:       Animals_Handling,
			entity:     entityTraveller,
			Group:      "Animals",
			Speciality: "Handling",
			Descr:      "The Traveller knows how to handle an animal and ride those trained to bear a rider. Unusual animals raise the difficulty of the check.",
		}
	case Animals_Veterinary:
		s = skill{
			Code:       Animals_Veterinary,
			entity:     entityTraveller,
			Group:      "Animals",
			Speciality: "Veterinary",
			Descr:      "The Traveller is trained in veterinary medicine and animal care.",
		}
	case Animals_Training:
		s = skill{
			Code:       Animals_Training,
			entity:     entityTraveller,
			Descr:      "The Traveller knows how to tame and train animals.",
			Group:      "Animals",
			Speciality: "Training",
		}
	case Art:
		s = skill{
			Code:       Art,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "",
			Descr:      "The Traveller is trained in a type of creative art.",
			hasSpec:    true,
		}
	case Art_Holography:
		s = skill{
			Code:       Art_Holography,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Holography",
			Descr:      "Recording and producing aesthetically pleasing and clear holographic images.",
		}
	case Art_Instrument:
		s = skill{
			Code:       Art_Instrument,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Instrument",
			Descr:      "Playing a particular musical instrument, such a flute, piano or organ.",
		}
	case Art_Performer:
		s = skill{
			Code:       Art_Performer,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Performer",
			Descr:      "The Traveller is a trained actor, dancer or singer at home on the stage, screen or holo.",
		}
	case Art_VisualMedia:
		s = skill{
			Code:       Art_VisualMedia,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Visual Media",
			Descr:      "Making artistic or abstract paintings or sculptures in a variety of media.",
		}
	case Art_Write:
		s = skill{
			Code:       Art_Write,
			entity:     entityTraveller,
			Group:      "Art",
			Speciality: "Write",
			Descr:      "Composing inspiring or interesting pieces of text.",
		}
	case Astrogation:
		s = skill{
			Code:   Astrogation,
			entity: entityTraveller,
			Group:  "Astrogation",
			Descr:  "This skill is for plotting the courses of starships and calculating accurate jumps.",
		}
	case Athletics:
		s = skill{
			Code:    Athletics,
			entity:  entityTraveller,
			Group:   "Athletics",
			Descr:   "The Traveller is a trained athlete and is physically fit.",
			hasSpec: true,
		}
	case Athletics_DEX:
		s = skill{
			Code:       Athletics_DEX,
			entity:     entityTraveller,
			Group:      "Athletics",
			Speciality: "DEX",
			Descr:      "Climbing, Juggling, Throwing. For alien races with wings, this also includes flying.",
		}
	case Athletics_END:
		s = skill{
			Code:       Athletics_END,
			entity:     entityTraveller,
			Group:      "Athletics",
			Speciality: "END",
			Descr:      "Long-distance running, hiking.",
		}
	case Athletics_STR:
		s = skill{
			Code:       Athletics_STR,
			entity:     entityTraveller,
			Group:      "Athletics",
			Speciality: "STR",
			Descr:      "Feats of strength, weight-lifting.",
		}
	case Broker:
		s = skill{
			Code:       Broker,
			entity:     entityTraveller,
			Group:      "Broker",
			Speciality: "",
			Descr:      "The Broker skill allows a Traveller to negotiate trades and arrange fair deals. It is heavily used when trading.",
		}
	case Carouse:
		s = skill{
			Code:       Carouse,
			entity:     entityTraveller,
			Group:      "Carouse",
			Speciality: "",
			Descr:      "Carousing is the art of socialising; having fun, but also ensuring other people have fun, and infectious good humour. It also covers social awareness and subterfuge in such situations.",
		}
	case Deception:
		s = skill{
			Code:       skillcode,
			entity:     entityTraveller,
			Group:      "Deception",
			Speciality: "",
			Descr:      "Deception allows a Traveller to lie fluently, disguise himself, perform sleight of hand and fool onlookers. Most underhanded ways of cheating and lying fall under deception.",
		}
	case Diplomat:
		s = skill{
			Code:       Diplomat,
			entity:     entityTraveller,
			Group:      "Diplomat",
			Speciality: "",
			Descr:      "The Diplomat skill is for negotiating deals, establishing peaceful contact and smoothing over social faux pas. It includes how to behave in high society and proper ways to address nobles. It is a much more formal skill than Persuade.",
		}
	case Drive:
		s = skill{
			Code:       Drive,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "",
			Descr:      "This skill is for controlling ground vehicles of various types. There are several specialities.",
			hasSpec:    true,
		}
	case Drive_Hovercraft:
		s = skill{
			Code:       Drive_Hovercraft,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Hovercraft",
			Descr:      "Vehicles that rely on a cushion of air and thrusters for motion.",
		}
	case Drive_Mole:
		s = skill{
			Code:       Drive_Mole,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Mole",
			Descr:      "For controlling vehicles that move through solid matter using drills or other earth-moving technologies, such as plasma torches or cavitation.",
		}
	case Drive_Track:
		s = skill{
			Code:       Drive_Track,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Track",
			Descr:      "For tanks and other vehicles that move on tracks.",
		}
	case Drive_Walker:
		s = skill{
			Code:       Drive_Walker,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Walker",
			Descr:      "Vehicles that use two or more legs to manoeuvre.",
		}
	case Drive_Wheel:
		s = skill{
			Code:       Drive_Wheel,
			entity:     entityTraveller,
			Group:      "Drive",
			Speciality: "Wheel",
			Descr:      "Vehicles that use two or more legs to manoeuvre.",
		}
	case Electronics:
		s = skill{
			Code:       Electronics,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "",
			Descr:      "This skill is used to operate electronic devices such as computers and ship-board systems. Higher levels represent the ability to repair and create electronic devices and systems. There are several specialities.",
			hasSpec:    true,
		}
	case Electronics_Comms:
		s = skill{
			Code:       Electronics_Comms,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Comms",
			Descr:      "The use of modern telecommunications; opening communications channels, querying computer networks, jamming signals and so on, as well as the proper protocols for communicating with starports and other spacecraft.",
		}
	case Electronics_Computers:
		s = skill{
			Code:       Electronics_Computers,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Computers",
			Descr:      "Using and controlling computer systems, and similar electronics and electrics.",
		}
	case Electronics_Remoteops:
		s = skill{
			Code:       Electronics_Remoteops,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Remote Ops",
			Descr:      "Using telepresence to remotely control drones, missiles, robots and other devices.",
		}
	case Electronics_Sensors:
		s = skill{
			Code:       Electronics_Sensors,
			entity:     entityTraveller,
			Group:      "Electronics",
			Speciality: "Sensors",
			Descr:      "The use and interpretation of data from electronic sensor devices, from observation satellites and remote probes to thermal imaging and densitometers.",
		}
	case Engineer:
		s = skill{
			Code:       Engineer,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "",
			Descr:      "The Engineer skill is used to operate and maintain spacecraft and advanced vehicles. Engineer can be used to make repairs on damaged systems on spacecraft and advanced vehicles. For repairs on simpler machines and systems, use the Mechanic skill.",
			hasSpec:    true,
		}
	case Engineer_Jdrive:
		s = skill{
			Code:       Engineer_Jdrive,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "J-drive",
			Descr:      "Maintaining and operating a spacecraft's Jump drive.",
		}
	case Engineer_Lifesupport:
		s = skill{
			Code:       Engineer_Lifesupport,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "Life Support",
			Descr:      "Covers oxygen generators, heating and lighting and other necessary life support systems.",
		}
	case Engineer_Mdrive:
		s = skill{
			Code:       Engineer_Mdrive,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "M-drive",
			Descr:      "Maintaining and operating a spacecraft's manoeuvre drive, as well as its artificial gravity.",
		}
	case Engineer_Power:
		s = skill{
			Code:       Engineer_Power,
			entity:     entityTraveller,
			Group:      "Engineer",
			Speciality: "Power",
			Descr:      "Maintaining and operating a spacecraft's power plant.",
		}
	case Explosives:
		s = skill{
			Code:       Explosives,
			entity:     entityTraveller,
			Group:      "Explosives",
			Speciality: "",
			Descr:      "The Explosives skill covers the use of demolition charges and other explosive devices, including assembling or disarming bombs. A failed Explosives check with an Effect of -4 or less can result in a bomb detonating prematurely.",
		}
	case Flyer:
		s = skill{
			Code:       Flyer,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "",
			Descr:      "The various specialities of this skill cover different types of flying vehicles. Flyers only work in an atmosphere; vehicles that can leave the atmosphere and enter orbit generally use the Pilot skill.",
			hasSpec:    true,
		}
	case Flyer_Airship:
		s = skill{
			Code:       Flyer_Airship,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Airship",
			Descr:      "Used for airships, dirigibles and other powered lighter than air craft.",
		}
	case Flyer_Grav:
		s = skill{
			Code:       Flyer_Grav,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Grav",
			Descr:      "This covers air/rafts, grav belts and other vehicles that use gravitic technology.",
		}
	case Flyer_Ornithopter:
		s = skill{
			Code:       Flyer_Ornithopter,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Ornithopter",
			Descr:      "For vehicles that fly through the use of flapping wings.",
		}
	case Flyer_Rotor:
		s = skill{
			Code:       Flyer_Rotor,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Rotor",
			Descr:      "For helicopters, tilt-rotors and aerodynes.",
		}
	case Flyer_Wing:
		s = skill{
			Code:       Flyer_Wing,
			entity:     entityTraveller,
			Group:      "Flyer",
			Speciality: "Wing",
			Descr:      "For jets, vectored thrust aircraft and aeroplanes using a lifting body.",
		}
	case Gambler:
		s = skill{
			Code:       Gambler,
			entity:     entityTraveller,
			Group:      "Gambler",
			Speciality: "",
			Descr:      "The Traveller is familiar with a wide variety of gambling games, such as poker, roulette, blackjack, horse-racing, sports betting and so on, and has an excellent grasp of statistics and probability. Gambler increases the rewards from Benefit Rolls, giving the Traveller DM+1 to his cash rolls if he has Gambler 1 or better.",
		}
	case Guncombat:
		s = skill{
			Code:       Guncombat,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "",
			Descr:      "The Gun Combat skill covers a variety of ranged weapons.",
			hasSpec:    true,
		}
	case Guncombat_Archaic:
		s = skill{
			Code:       Guncombat_Archaic,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "Archaic",
			Descr:      "For primitive weapons that are not thrown, such as bows and blowpipes.",
		}
	case Guncombat_Energy:
		s = skill{
			Code:       Guncombat_Energy,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "Energy",
			Descr:      "Using advanced energy weapons like laser pistols or plasma rifles.",
		}
	case Guncombat_Slug:
		s = skill{
			Code:       Guncombat_Slug,
			entity:     entityTraveller,
			Group:      "Gun Combat",
			Speciality: "Slug",
			Descr:      "Weapons that fire a solid projectile such as the autorifle or gauss rifle.",
		}
	case Gunner:
		s = skill{
			Code:       Gunner,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "",
			Descr:      "The various specialities of this skill deal with the operation of ship-mounted weapons in space combat. See Spacecraft Operations chapter for more details. Most Travellers have smaller ships equipped solely with turret weapons.",
			hasSpec:    true,
		}
	case Gunner_Capital:
		s = skill{
			Code:       Gunner_Capital,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Capital",
			Descr:      "Operating bay or spinal mount weapons on board a ship.",
		}
	case Gunner_Ortilery:
		s = skill{
			Code:       Gunner_Ortilery,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Ortilery",
			Descr:      "A contraction of Orbital Artillery – using a ship's weapons for planetary bombardment or attacks on stationary targets.",
		}
	case Gunner_Screen:
		s = skill{
			Code:       Gunner_Screen,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Screen",
			Descr:      "Activating and using a ship's energy screens like Black Globe generators or meson screens.",
		}
	case Gunner_Turret:
		s = skill{
			Code:       Gunner_Turret,
			entity:     entityTraveller,
			Group:      "Gunner",
			Speciality: "Turret",
			Descr:      "Operating turret-mounted weapons on board a ship.",
		}
	case Heavyweapon:
		s = skill{
			Code:       Heavyweapon,
			entity:     entityTraveller,
			Group:      "Heavy weapon",
			Speciality: "",
			Descr:      "The Heavy Weapons skill covers man-portable and larger weapons that cause extreme property damage, such as rocket launchers, artillery and large plasma weapons.",
			hasSpec:    true,
		}
	case Heavyweapon_Artilery:
		s = skill{
			Code:       Heavyweapon_Artilery,
			entity:     entityTraveller,
			Group:      "Heavy weapon",
			Speciality: "Artilery",
			Descr:      "Fixed guns, mortars and other indirect-fire weapons.",
		}
	case Heavyweapon_Manportable:
		s = skill{
			Code:       Heavyweapon_Manportable,
			entity:     entityTraveller,
			Group:      "Heavy Weapon",
			Speciality: "Man Portable",
			Descr:      "Missile launchers, flamethrowers and man portable fusion and plasma.",
		}
	case Heavyweapon_Vehicle:
		s = skill{
			Code:       Heavyweapon_Vehicle,
			entity:     entityTraveller,
			Group:      "Heavy Weapon",
			Speciality: "Vehicle",
			Descr:      "Large weapons typically mounted on vehicles or strongpoints such as tank guns and autocannon.",
		}
	case Investigate:
		s = skill{
			Code:       Investigate,
			entity:     entityTraveller,
			Group:      "Investigate",
			Speciality: "",
			Descr:      "The Investigate skill incorporates keen observation, forensics, and detailed analysis.",
		}
	case Jack_of_all_trades:
		s = skill{
			Code:       Jack_of_all_trades,
			entity:     entityTraveller,
			Group:      "Jack-of-All-Trades",
			Speciality: "",
			Descr:      "The Jack-of-All-Trades skill works differently to other skills. It reduces the unskilled penalty a Traveller receives for not having the appropriate skill by one for every level of Jack-of-All-Trades. For example, if a Traveller does not have the Pilot skill, he suffers DM-3 to all Pilot checks. If that Traveller has Jack-of-All-Trades 2, then the penalty is reduced by 2 to DM-1. With Jackof-All-Trades 3, a Traveller can totally negate the penalty for being unskilled.\nThere is no benefit for having Jack-of-All-Trades 0 or Jack-of-All-Trades 4 or more.",
		}
	case Language:
		s = skill{
			Code:       Language,
			entity:     entityTraveller,
			Group:      "Language",
			Speciality: "",
			Descr:      "There are numerous different Language specialities, each one covering reading and writing a different language. All Travellers can speak and read their native language without needing the Language skill, and automated computer translator programs mean Language skills are not always needed on other worlds. Having Language 0 implies the Traveller has a smattering of simple phrases in several languages.",
		}
	case Language_Anglic:
		s = skill{
			Code:       Language_Anglic,
			entity:     entityTraveller,
			Group:      "Language",
			Speciality: "Anglic",
			Descr:      "The common trade language of the Third Imperium, derived originally from the English spoken in the Rule of Man",
		}
	case Leadership:
		s = skill{
			Code:       Leadership,
			entity:     entityTraveller,
			Group:      "Leadership",
			Speciality: "",
			Descr:      "The Leadership skill is for directing, inspiring and rallying allies and comrades. A Traveller may make a Leadership action in combat, as detailed on page 72 (CRB).",
		}
	case Mechanic:
		s = skill{
			Code:       Mechanic,
			entity:     entityTraveller,
			Group:      "Mechanic",
			Speciality: "",
			Descr:      "The Mechanic skill allows a Traveller to maintain and repair most equipment – some advanced equipment and spacecraft components require the Engineer skill. Unlike the narrower and more focussed Engineer or Science skills, Mechanic does not allow a Traveller to build new devices or alter existing ones – it is purely for repairs and maintenance but covers all types of equipment.",
		}
	case Medic:
		s = skill{
			Code:       Medic,
			entity:     entityTraveller,
			Group:      "Medic",
			Speciality: "",
			Descr:      "The Medic skill covers emergency first aid and battlefield triage as well as diagnosis, treatment, surgery and longterm care. See Injury and Recovery on page 47.",
		}
	case Melee:
		s = skill{
			Code:       Melee,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "",
			Descr:      "The Melee skill covers attacking in hand-to-hand combat and the use of suitable weapons.",
			hasSpec:    true,
		}
	case Melee_Blade:
		s = skill{
			Code:       Melee_Blade,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Blade",
			Descr:      "Attacking with swords, rapiers, blades and other edged weapons.",
		}
	case Melee_Bludgeon:
		s = skill{
			Code:       Melee_Bludgeon,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Bludgeon",
			Descr:      "Attacking with maces, clubs, staves and so on.",
		}
	case Melee_Natural:
		s = skill{
			Code:       Melee_Natural,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Natural",
			Descr:      " Weapons that are part of an alien or creature, such as claws or teeth.",
		}
	case Melee_Unarmed:
		s = skill{
			Code:       Melee_Unarmed,
			entity:     entityTraveller,
			Group:      "Melee",
			Speciality: "Unarmed",
			Descr:      "Punching, kicking and wrestling; using improvised weapons in a bar brawl.",
		}
	case Navigation:
		s = skill{
			Code:       Navigation,
			entity:     entityTraveller,
			Group:      "Navigation",
			Speciality: "",
			Descr:      "Navigation is the planetside counterpart of astrogation, covering plotting courses and finding directions on the ground.",
		}
	case Persuade:
		s = skill{
			Code:       Persuade,
			entity:     entityTraveller,
			Group:      "Persuade",
			Speciality: "",
			Descr:      "Persuade is a more casual, informal version of Diplomat. It covers fast talking, bargaining, wheedling and bluffing. It also covers bribery or intimidation.",
		}
	case Pilot:
		s = skill{
			Code:       Pilot,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "",
			Descr:      "The Pilot skill specialities cover different forms of  spacecraft. See Spacecraft Operations chapter for more  details.",
			hasSpec:    true,
		}
	case Pilot_CapitalShips:
		s = skill{
			Code:       Pilot_CapitalShips,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "Capital Ships",
			Descr:      "Battleships and other ships over 5,000 tons.",
		}
	case Pilot_SmallCraft:
		s = skill{
			Code:       Pilot_SmallCraft,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "Small Craft",
			Descr:      "Shuttles and other craft under 100 tons.",
		}
	case Pilot_Spacecraft:
		s = skill{
			Code:       Pilot_Spacecraft,
			entity:     entityTraveller,
			Group:      "Pilot",
			Speciality: "Spacecraft",
			Descr:      ": Trade ships and other vessels between 100 and 5,000 tons.",
		}
	case Profession:
		s = skill{
			Code:       Profession,
			entity:     entityTraveller,
			Group:      "Profession",
			Speciality: "",
			Descr:      "A Traveller with a Profession skill is trained in producing useful goods or services. There are many different Profession specialities.",
			hasSpec:    true,
		}
	case Profession_Any:
		s = skill{}
	case Recon:
		s = skill{
			Code:       Recon,
			entity:     entityTraveller,
			Group:      "Recon",
			Speciality: "",
			Descr:      "A Traveller trained in Recon is able to scout out dangers and spot threats, unusual objects or out of place people. ",
		}
	case Science_Physical:
		s = skill{
			Code:       Science_Physical,
			entity:     entityTraveller,
			Group:      "Science Physical",
			Speciality: "",
		}
	case Science_Life:
		s = skill{
			Code:       Science_Life,
			entity:     entityTraveller,
			Group:      "Science Life",
			Speciality: "",
		}
	case Science_Social:
		s = skill{
			Code:       Science_Social,
			entity:     entityTraveller,
			Group:      "Science_Social",
			Speciality: "",
		}
	case Science_Space:
		s = skill{
			Code:       Science_Space,
			entity:     entityTraveller,
			Group:      "Science Space",
			Speciality: "",
		}
	case Science:
		s = skill{
			Code:       Science,
			entity:     entityTraveller,
			Group:      "Science",
			Speciality: "",
			Descr:      "The Science skill covers not just knowledge but also practical application of that knowledge where such practical application is possible. There are a large range of specialities.",
		}
	case Science_Any:
		s = skill{}
	case Seafarer:
		s = skill{
			Code:       Seafarer,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "",
			Descr:      "The Seafarer skill covers all manner of watercraft and ocean travel.",
			hasSpec:    true,
		}
	case Seafarer_Oceanships:
		s = skill{
			Code:       Seafarer_Oceanships,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Ocean Ships",
			Descr:      "For motorised sea-going vessels.",
		}
	case Seafarer_Personal:
		s = skill{
			Code:       Seafarer_Personal,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Personal",
			Descr:      "Used for very small waterborne craft such as canoes and rowboats.",
		}
	case Seafarer_Sail:
		s = skill{
			Code:       Seafarer_Sail,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Sail",
			Descr:      "This skill is for wind-driven watercraft.",
		}
	case Seafarer_Submarine:
		s = skill{
			Code:       Seafarer_Submarine,
			entity:     entityTraveller,
			Group:      "Seafarer",
			Speciality: "Submarine",
			Descr:      "For vehicles that travel underwater.",
		}
	case Stealth:
		s = skill{
			Code:       Stealth,
			entity:     entityTraveller,
			Group:      "Stealth",
			Speciality: "",
			Descr:      "A Traveller trained in the Stealth skill is adept at staying unseen, unheard, and unnoticed.",
		}
	case Steward:
		s = skill{
			Code:       skillcode,
			entity:     entityTraveller,
			Group:      "",
			Speciality: "",
			Descr:      "",
		}
	case Streetwise:
		s = skill{
			Code:       Streetwise,
			entity:     entityTraveller,
			Group:      "Streetwise",
			Speciality: "",
			Descr:      "The Steward skill allows the Traveller to serve and care for nobles and high-class passengers. It covers everything from proper address and behaviour to cooking and tailoring, as well as basic management skills. A Traveller with the Steward skill is necessary on any ship offering High Passage. See Spacecraft Operations chapter for more details.",
		}
	case Survival:
		s = skill{
			Code:       Survival,
			entity:     entityTraveller,
			Group:      "Survival",
			Speciality: "",
			Descr:      "",
		}
	case Tactics:
		s = skill{
			Code:       Tactics,
			entity:     entityTraveller,
			Group:      "Tactics",
			Speciality: "",
			Descr:      "This skill covers tactical planning and decision making, from board games to squad level combat to fleet engagements. For use in combat, see Combat chapter",
			hasSpec:    true,
		}
	case Tactics_Military:
		s = skill{
			Code:       Tactics_Military,
			entity:     entityTraveller,
			Group:      "Tactics",
			Speciality: "Military",
			Descr:      "Co-ordinating the attacks of foot troops or vehicles on the ground.",
		}
	case Tactics_Navy:
		s = skill{
			Code:       Tactics_Navy,
			entity:     entityTraveller,
			Group:      "Tactics",
			Speciality: "Navy",
			Descr:      "Co-ordinating the attacks of a spacecraft or fleet.",
		}
	case VaccSuit:
		s = skill{
			Code:       VaccSuit,
			entity:     entityTraveller,
			Group:      "Vacc Suit",
			Speciality: "",
			Descr:      "The Vacc Suit skill allows a Traveller to wear and operate spacesuits and environmental suits. A Traveller will rarely need to make Vacc Suit checks under ordinary circumstances – merely possessing the skill is enough. If the Traveller does not have the requisite Vacc Suit skill for the suit he is wearing, he suffers DM-2 to all skill checks made while wearing a suit for each missing level. This skill also permits the character to operate advanced battle armour. ",
		}
	}
	return s, nil
}

/*
fmt.Print(char.Skills[skill.Advocate].Description())
entity:
	traveller
		characteristic
		skill
		trait
	dynasty
		characteristic

	firm
*/

func (s *skill) Name() string {
	name := s.Group
	if s.Speciality != "" {
		name += " (" + s.Speciality + ")"
	}
	return name
}

func (s *skill) Description() string {
	return s.Descr
}

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
