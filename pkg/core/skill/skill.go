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
	Pilot_SpaceCraft
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
}

func New(skillcode int) (skill, error) {
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
			Code:   skillcode,
			entity: entityTraveller,
			Group:  "Astrogation",
			Descr:  "This skill is for plotting the courses of starships and calculating accurate jumps.",
		}
	case Athletics:
		s = skill{}
	case Athletics_DEX:
		s = skill{}
	case Athletics_END:
		s = skill{}
	case Athletics_STR:
		s = skill{}
	case Broker:
		s = skill{}
	case Carouse:
		s = skill{}
	case Deception:
		s = skill{}
	case Diplomat:
		s = skill{}
	case Drive:
		s = skill{}
	case Drive_Hovercraft:
		s = skill{}
	case Drive_Mole:
		s = skill{}
	case Drive_Track:
		s = skill{}
	case Drive_Walker:
		s = skill{}
	case Drive_Wheel:
		s = skill{}
	case Electronics:
		s = skill{}
	case Electronics_Comms:
		s = skill{}
	case Electronics_Computers:
		s = skill{}
	case Electronics_Remoteops:
		s = skill{}
	case Electronics_Sensors:
		s = skill{}
	case Engineer:
		s = skill{}
	case Engineer_Jdrive:
		s = skill{}
	case Engineer_Lifesupport:
		s = skill{}
	case Engineer_Mdrive:
		s = skill{}
	case Engineer_Power:
		s = skill{}
	case Explosives:
		s = skill{}
	case Flyer:
		s = skill{}
	case Flyer_Airship:
		s = skill{}
	case Flyer_Grav:
		s = skill{}
	case Flyer_Ornithopter:
		s = skill{}
	case Flyer_Rotor:
		s = skill{}
	case Flyer_Wing:
		s = skill{}
	case Gambler:
		s = skill{}
	case Guncombat:
		s = skill{}
	case Guncombat_Archaic:
		s = skill{}
	case Guncombat_Energy:
		s = skill{}
	case Guncombat_Slug:
		s = skill{}
	case Gunner:
		s = skill{}
	case Gunner_Capital:
		s = skill{}
	case Gunner_Ortilery:
		s = skill{}
	case Gunner_Screen:
		s = skill{}
	case Gunner_Turret:
		s = skill{}
	case Heavyweapon:
		s = skill{}
	case Heavyweapon_Artilery:
		s = skill{}
	case Heavyweapon_Manportable:
		s = skill{}
	case Heavyweapon_Vehicle:
		s = skill{}
	case Investigate:
		s = skill{}
	case Jack_of_all_trades:
		s = skill{}
	case Language:
		s = skill{}
	case Language_Anglic:
		s = skill{}
	case Leadership:
		s = skill{}
	case Mechanic:
		s = skill{}
	case Medic:
		s = skill{}
	case Melee:
		s = skill{}
	case Melee_Blade:
		s = skill{}
	case Melee_Bludgeon:
		s = skill{}
	case Melee_Natural:
		s = skill{}
	case Melee_Unarmed:
		s = skill{}
	case Navigation:
		s = skill{}
	case Persuade:
		s = skill{}
	case Pilot:
		s = skill{}
	case Pilot_CapitalShips:
		s = skill{}
	case Pilot_SmallCraft:
		s = skill{}
	case Pilot_SpaceCraft:
		s = skill{}
	case Profession:
		s = skill{}
	case Profession_Any:
		s = skill{}
	case Recon:
		s = skill{}
	case Science_Physical:
		s = skill{}
	case Science_Life:
		s = skill{}
	case Science_Social:
		s = skill{}
	case Science_Space:
		s = skill{}
	case Science:
		s = skill{}
	case Science_Any:
		s = skill{}
	case Seafarer:
		s = skill{}
	case Seafarer_Oceanships:
		s = skill{}
	case Seafarer_Personal:
		s = skill{}
	case Seafarer_Sail:
		s = skill{}
	case Seafarer_Submarine:
		s = skill{}
	case Stealth:
		s = skill{}
	case Steward:
		s = skill{}
	case Streetwise:
		s = skill{}
	case Survival:
		s = skill{}
	case Tactics:
		s = skill{}
	case Tactics_Military:
		s = skill{}
	case Tactics_Navy:
		s = skill{}
	case VaccSuit:
		s = skill{}

	}
	return s, nil
}

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


The Traveller is a trained athlete and is physically fit.
Climbing, Juggling, Throwing. For alien races with wings, this also includes flying.
Long-distance running, hiking. 
Feats of strength, weight-lifting.
The Broker skill allows a Traveller to negotiate trades and arrange fair deals. It is heavily used when trading.
Carousing is the art of socialising; having fun, but also ensuring other people have fun, and infectious good humour. It also covers social awareness and subterfuge in such situations.
Deception allows a Traveller to lie fluently, disguise himself, perform sleight of hand and fool onlookers. Most underhanded ways of cheating and lying fall under deception.
Convincing a Guard to let you Past Without ID: Very
Difficult (12+) Deception check (1D minutes, INT).
Alternatively, oppose with a Recon check.
Palming a Credit Chit: Average (8+) Deception check
(1D seconds, DEX).
Disguising Yourself as a Wealthy Noble to Fool a Client:
Difficult (10+) Deception check (1D x 10 minutes, INT
or SOC). Alternatively, oppose with a Recon check.
Diplomat
The Diplomat skill is for negotiating deals, establishing
peaceful contact and smoothing over social faux pas. It
includes how to behave in high society and proper ways
to address nobles. It is a much more formal skill than
Persuade.
Greeting the Emperor Properly: Difficult (10+) Diplomat
check (1D minutes, SOC).
Negotiating a Peace Treaty: Average (8+) Diplomat
check (1D days, EDU).
Transmitting a Formal Surrender: Average (8+) Diplomat
check (1D x 10 seconds, INT).
63
Drive
This skill is for controlling ground vehicles of various
types. There are several specialities.
SPECIALITIES
● Hovercraft: Vehicles that rely on a cushion of air and
thrusters for motion.
Manoeuvring a Hovercraft Through a Tight Canal:
Difficult (10+) Drive (hovercraft) check (1D
minutes, DEX).
● Mole: For controlling vehicles that move through
solid matter using drills or other earth-moving
technologies, such as plasma torches or cavitation.
Surfacing in the Right Place: Average (8+) Drive
(mole) check (1D x 10 minutes, INT).
Precisely Controlling a Dig to Expose a Vein of
Minerals: Difficult (10+) Drive (mole) check (1D x
10 minutes, DEX).
● Track: For tanks and other vehicles that move on
tracks.
Manoeuvring (or Smashing, Depending on the
Vehicle) Through a Forest: Difficult (10+) Drive
(tracked) check (1D minutes, DEX).
Driving a Tank into a Cargo Bay: Average (8+) Drive
(tracked) check (1D x 10 seconds, DEX).
● Walker: Vehicles that use two or more legs to
manoeuvre.
Negotiating Rough Terrain: Difficult (10+) Drive
(walker) check (1D minutes, DEX).
● Wheel: For automobiles and similar groundcars.
Driving a Groundcar in a Short Race: Opposed Drive
(wheeled) check (1D minutes, DEX). Longer races
use END instead of DEX.
Avoiding an Unexpected Obstacle on the Road: Average
(8+) Drive (wheeled) check (1D seconds, DEX).
Electronics
This skill is used to operate electronic devices such
as computers and ship-board systems. Higher levels
represent the ability to repair and create electronic
devices and systems. There are several specialities.
SPECIALITIES
● Comms: The use of modern telecommunications;
opening communications channels, querying
computer networks, jamming signals and so on, as
well as the proper protocols for communicating with
starports and other spacecraft.
Requesting Landing Privileges at a Starport: Routine
(6+) Electronic (comms) check (1D minutes, EDU).
Accessing Publicly Available but Obscure Data Over
Comms: Average (8+) Electronic (comms) check (1D
x 10 minutes, EDU).
Bouncing a Signal off Orbiting Satellite to Hide Your
Transmitter: Difficult (10+) Electronics (comms)
check (1D x 10 minutes, INT).
Jamming a Comms System: Opposed Electronics
(comms) check (1D minutes, INT). Difficult
(10+) for radio, Very Difficult (12+) for laser, and
Formidable (14+) for masers. A Traveller using a
comms system with a higher Technology Level than
his opponent gains DM+1 for every TL of difference.
● Computers: Using and controlling computer systems,
and similar electronics and electrics.
Accessing Publicly Available Data: Easy (4+)
Electronics (computers) check (1D minutes, INT or
EDU).
Activating a Computer Program on a Ship’s
Computer: Routine (6+) Electronics (computers)
check (1D x 10 seconds, INT or EDU).
Searching a Corporate Database for Evidence
of Illegal Activity: Difficult (10+) Electronics
(computers) check (1D hours, INT).
Hacking into a Secure Computer Network:
Formidable (14+) Electronics (computers) check
(1D x 10 hours, INT). Hacking is aided by Intrusion
programs and made more difficult by Security
programs. The Effect determines the amount of data
retrieved; failure means the targeted system may be
able to trace the hacking attempt.
• Remote Ops: Using telepresence to remotely control
drones, missiles, robots and other devices.
Using a Mining Drone to Excavate an Asteroid:
Routine (6+) Electronics (remote ops) check (1D
hours, DEX).
• Sensors: The use and interpretation of data from
electronic sensor devices, from observation
satellites and remote probes to thermal imaging and
densitometers.
Making a Detailed Sensor Scan: Routine (6+)
Electronics (sensors) check (1D x 10 minutes, INT
or EDU).
Analysing Sensor Data: Average (8+) Electronics
(sensors) check (1D hours, INT).
Engineer
The Engineer skill is used to operate and maintain
spacecraft and advanced vehicles. Engineer can be used
to make repairs on damaged systems on spacecraft and
advanced vehicles. For repairs on simpler machines and
systems, use the Mechanic skill.
SPECIALITIES
● M-drive: Maintaining and operating a spacecraft’s
manoeuvre drive, as well as its artificial gravity.
Overcharging a Thruster Plate to Increase a Ship’s
Estimating a Ship’s Tonnage From its Observed
Performance: Average (8+) Engineer (m-drive) check
(1D x 10 seconds, INT).
● J-drive: Maintaining and operating a spacecraft’s
Jump drive.
64
Making a Jump: Easy (4+) Engineer (j-drive) check
(1D x 10 minutes, EDU).
● Life Support: Covers oxygen generators, heating and
lighting and other necessary life support systems.
Safely Reducing Power to Life Support to Prolong
a Ship’s Battery Life: Average (8+) Engineer (life
support) check (1D minutes, EDU).
● Power: Maintaining and operating a spacecraft’s
power plant.
Monitoring an Enemy ship's Power Output to
Determine its Capabilities: Difficult (10+) Engineer
(power) check (1D minutes, INT).
Explosives
The Explosives skill covers the
use of demolition charges and
other explosive devices, including
assembling or disarming bombs.
A failed Explosives check with an
Effect of -4 or less can result in a bomb
detonating prematurely.
Planting Charges to Collapse
a Wall in a Building: Average
(8+) Explosives check (1D x 10
minutes, EDU).
Planting a Breaching Charge:
Average (8+) Explosives check (1D
x 10 seconds, EDU). The damage
from the explosive is multiplied by
the Effect.
Disarming a Bomb Equipped with
Anti-Tamper Trembler Detonators:
Formidable (14+) Explosives check
(1D minutes, DEX).
Flyer
The various specialities of this skill cover
different types of flying vehicles. Flyers
only work in an atmosphere; vehicles that
can leave the atmosphere and enter orbit
generally use the Pilot skill.
SPECIALITIES
• Airship: Used for airships, dirigibles and other
powered lighter than air craft.
• Grav: This covers air/rafts, grav belts and other
vehicles that use gravitic technology.
• Ornithopter: For vehicles that fly through the use of
flapping wings.
• Rotor: For helicopters, tilt-rotors and aerodynes.
• Wing: For jets, vectored thrust aircraft and
aeroplanes using a lifting body.
Landing Safely: Routine (6+) Flyer check (1D minutes, DEX).
Racing Another Flyer: Opposed Flyer check (1D x 10
minutes, DEX).
Gambler
The Traveller is familiar with a wide variety of gambling
games, such as poker, roulette, blackjack, horse-racing,
sports betting and so on, and has an excellent grasp of
statistics and probability. Gambler increases the rewards
from Benefit Rolls, giving the Traveller DM+1 to his cash
rolls if he has Gambler 1 or better.
65
A Casual Game of Poker: Opposed Gambler check (1D
hours, INT).
Picking the Right Horse to Bet On: Average (8+)
Gambler check (1D minutes, INT).
Gunner
The various specialities of this skill deal with the operation
of ship-mounted weapons in space combat. See Spacecraft
Operations chapter for more details. Most Travellers have
smaller ships equipped solely with turret weapons.
SPECIALITIES
● Turret: Operating turret-mounted weapons on board
a ship.
Firing a Turret at an Enemy Ship: Average (8+)
Gunner (turret) check (1D seconds, DEX).
● Ortillery: A contraction of Orbital Artillery – using
a ship’s weapons for planetary bombardment or
attacks on stationary targets.
Firing Ortillery: Average (8+) Gunner (ortillery) check
(1D minutes, INT).
● Screen: Activating and using a ship’s energy screens
like Black Globe generators or meson screens.
Activating a Screen to Intercept Enemy Fire: Difficult
(10+) Gunner (screen) check (1D seconds, DEX).
● Capital: Operating bay or spinal mount weapons on
board a ship.
Firing a Spinal Mount Weapon: Average (8+) Gunner
(capital) check (1D minutes, INT).
Gun Combat
The Gun Combat skill covers a variety of ranged weapons.
See Combat chapter for details on using guns in combat.
SPECIALITIES
● Archaic: For primitive weapons that are not thrown,
such as bows and blowpipes.
● Energy: Using advanced energy weapons like laser
pistols or plasma rifles.
● Slug: Weapons that fire a solid projectile such as the
autorifle or gauss rifle.
Firing a Gun: Average (8+) Gun Combat check (1D
seconds, DEX).
Heavy Weapons
The Heavy Weapons skill covers man-portable and larger
weapons that cause extreme property damage, such as
rocket launchers, artillery and large plasma weapons.
SPECIALITIES
● Artillery: Fixed guns, mortars and other indirect-fire
weapons.
● Man Portable: Missile launchers, flamethrowers and
man portable fusion and plasma.
● Vehicle: Large weapons typically mounted on
vehicles or strongpoints such as tank guns and
autocannon.
Firing an Artillery Piece at a Visible Target: Average (8+)
Heavy Weapons (artillery) check (1D seconds, DEX).
Firing an Artillery Piece Using Indirect Fire: Difficult (10+)
Heavy Weapons (artillery) check (1D x 10 seconds, INT).
Investigate
The Investigate skill incorporates keen observation,
forensics, and detailed analysis.
Searching a Crime Scene For Clues: Average (8+)
Investigate check (1D x 10 minutes, INT).
Watching a Bank of Security Monitors in a Starport,
Watching for a Specific Criminal: Difficult (10+)
Investigate check (1D hours, INT).
Jack-of-All-Trades
The Jack-of-All-Trades skill works differently to other
skills. It reduces the unskilled penalty a Traveller
receives for not having the appropriate skill by one
for every level of Jack-of-All-Trades. For example, if a
Traveller does not have the Pilot skill, he suffers DM-3 to
all Pilot checks. If that Traveller has Jack-of-All-Trades
2, then the penalty is reduced by 2 to DM-1. With Jackof-
All-Trades 3, a Traveller can totally negate the penalty
for being unskilled.
There is no benefit for having Jack-of-All-Trades 0 or
Jack-of-All-Trades 4 or more.
Language
There are numerous different Language specialities,
each one covering reading and writing a different
language. All Travellers can speak and read their
native language without needing the Language skill,
and automated computer translator programs mean
Language skills are not always needed on other
worlds. Having Language 0 implies the Traveller has a
smattering of simple phrases in several languages.
SPECIALITIES
There are, of course, as many specialities of Language
as there are actual languages. Those presented here are
examples from the Third Imperium setting.
Anglic: The common trade language of the Third
Imperium, derived originally from the English spoken in
the Rule of Man.
66
Vilani: The language spoken by the Vilani of the First
Imperium; the ‘Latin’ of the Third Imperium.
Zdetl: The Zhodani spoken language.
Oynprith: The Droyne ritual language.
Ordering a Meal, Asking for Basic Directions: Routine
(6+) Language check (1D seconds, EDU).
Holding a Simple Conversation: Average (8+) Language
check (1D x 10 seconds, EDU).
Understanding a Complex Technical Document or
Report: Very Difficult (12+) Language check (1D
minutes, EDU).
Leadership
The Leadership skill is for directing, inspiring and
rallying allies and comrades. A Traveller may make a
Leadership action in combat, as detailed on page 72.
Shouting an Order: Average (8+) Leadership check (1D
seconds, SOC).
Rallying Shaken Troops: Difficult (10+) Leadership
check (1D seconds, SOC).
Mechanic
The Mechanic skill allows a Traveller to maintain and
repair most equipment – some advanced equipment
and spacecraft components require the Engineer skill.
Unlike the narrower and more focussed Engineer or
Science skills, Mechanic does not allow a Traveller to
build new devices or alter existing ones – it is purely
for repairs and maintenance but covers all types of
equipment.
Repairing a Damaged System in the Field: Average (8+)
Mechanic check (1D minutes, INT or EDU).
Medic
The Medic skill covers emergency first aid and battlefield
triage as well as diagnosis, treatment, surgery and longterm
care. See Injury and Recovery on page 47.
First Aid: Average (8+) Medic check (1D minutes, EDU).
The patient regains lost characteristic points equal to
the Effect.
Treat Poison or Disease: Average (8+) Medic check (1D
hours, EDU).
Long-term Care: Average (8+) Medic check (1D hours,
EDU).
Melee
The Melee skill covers attacking in hand-to-hand combat
and the use of suitable weapons.
SPECIALITIES
● Unarmed: Punching, kicking and wrestling; using
improvised weapons in a bar brawl.
● Blade: Attacking with swords, rapiers, blades and
other edged weapons.
● Bludgeon: Attacking with maces, clubs, staves and
so on.
● Natural: Weapons that are part of an alien or
creature, such as claws or teeth.
Swinging a Sword: Average (8+) Melee (blade) check
(1D seconds, STR or DEX).
67
given here. Also note that on some worlds other skills,
such as Animals or Computer, may be used to earn a
living in the same manner as Profession skills.
SPECIALITIES
● Belter: Mining asteroids for valuable ores and
minerals.
● Biologicals: Engineering and managing artificial
organisms.
● Civil Engineering: Designing structures and buildings.
● Construction: Building orbital habitats and
megastructures.
● Hydroponics: Growing crops in hostile environments.
● Polymers: Designing and using polymers.
Recon
A Traveller trained in Recon is able to scout out dangers
and spot threats, unusual objects or out of place people.
Working Out the Routine of a Trio of Guard Patrols:
Average (8+) Recon check (1D x 10 minutes, INT).
Spotting the Sniper Before he Shoots You: Recon check
(1D x 10 seconds, INT) opposed by Stealth (DEX) check.
Science
The Science skill covers not just knowledge but also
practical application of that knowledge where such
practical application is possible. There are a large range
of specialities.
SPECIALITIES
● Archaeology: The study of ancient civilisations,
including the previous Imperiums and Ancients. It also
covers techniques of investigation and excavations.
● Astronomy: The study of stars and celestial pheonomina.
● Biology: The study of living organisms.
● Chemistry: The study of matter at the atomic,
molecular, and macromolecular levels.
● Cosmology: The study of universe and its creation.
● Cybernetics: The study of blending living and
synthetic life.
● Economics: The study of trade and markets.
● Genetics: The study of genetic codes and
engineering.
● History: The study of the past, as seen through
documents and records as opposed to physical artefacts.
● Linguistics: The study of languages.
● Philosophy: The study of beliefs and religions.
● Physics: The study of the fundamental forces.
● Planetology: The study of planet formation and
evolution.
● Psionicology: The study of psionic powers and
phenomena.
Navigation
Navigation is the planetside counterpart of astrogation,
covering plotting courses and finding directions on the
ground.
Plotting a Course Using an Orbiting Satellite Beacon:
Average (8+) Navigation check (1D x 10 minutes, INT
or EDU).
Avoiding Getting Lost in Thick Jungle: Difficult (10+)
Navigation check (1D hours, INT).
Persuade
Persuade is a more casual, informal version of Diplomat.
It covers fast talking, bargaining, wheedling and
bluffing. It also covers bribery or intimidation.
Bluffing Your Way Past a Guard: Opposed Persuade
check (1D minutes, INT or SOC).
Haggling in a Bazaar: Opposed Persuade check (1D
minutes, INT or SOC).
Intimidating a Thug: Opposed Persuade check (1D
minutes, STR or SOC).
Asking the Alien Space Princess to Marry You: Formidable
(14+) Persuade check (1D x 10 minutes, SOC).
Pilot
The Pilot skill specialities cover different forms of
spacecraft. See Spacecraft Operations chapter for more
details.
SPECIALITIES
Small Craft: Shuttles and other craft under 100 tons.
Spacecraft: Trade ships and other vessels between 100
and 5,000 tons.
Capital Ships: Battleships and other ships over 5,000 tons.
Profession
A Traveller with a Profession skill is trained in producing
useful goods or services. There are many different
Profession specialities, but each one works the same
way – the Traveller can make a Profession check to earn
money on a planet that supports that trade. The amount
of money raised is Cr250 x the Effect of the check per
month. Unlike other skills with specialties, levels in
the Profession skill do not grant the ability to use other
specialties at level 0. Each specialty must be learned
individually. Someone with a Profession skill of 0 has a
general grasp of working for a living but little experience
beyond the most menial jobs.
There are a huge range of potential specialities for this
skill, one for every possible profession in the universe.
Some examples suitable to a science fiction setting are
68
● Psychology: The study of thought and society.
● Robotics: The study of robot construction and use.
● Sophontology: The study of intelligent living
creatures.
● Xenology: The study of alien life forms.
Remembering a Commonly Known Fact: Routine (6+)
Science check (1D minutes, EDU).
Researching a Problem Related to a Field of Science:
Average (8+) Science check (1D days, INT).
Seafarer
The Seafarer skill covers all manner of watercraft and
ocean travel.
SPECIALITIES
● Ocean Ships: For motorised sea-going
vessels. Personal: Used for very small waterborne
craft such as canoes and rowboats.
● Sail: This skill is for wind-driven watercraft.
● Submarine: For vehicles that travel underwater.
Stealth
A Traveller trained in the Stealth skill is adept at staying
unseen, unheard, and unnoticed.
Sneaking Past a Guard: Stealth check (1D x 10 seconds,
DEX) opposed by Recon (INT) check.
Avoiding Detection by a Security Patrol: Stealth check
(1D minutes, DEX) opposed by Recon (INT) check.
Steward
The Steward skill allows the Traveller to serve and
care for nobles and high-class passengers. It covers
everything from proper address and behaviour to
cooking and tailoring, as well as basic management
skills. A Traveller with the Steward skill is necessary
on any ship offering High Passage. See Spacecraft
Operations chapter for more details.
Cooking a Fine Meal: Average (8+) Steward check (1D
hours, EDU).
Calming Down an Angry Duke who has Just Been Told
you Will not be Jumping to his Destination on Time:
Difficult (10+) Steward check (1D minutes, SOC).
Streetwise
A Traveller with the Streetwise skill understands the urban
environment and the power structures in society. On his
homeworld and in related systems, he knows criminal
contacts and fixers. On other worlds, he can quickly intuit
power structures and can fit into local underworlds.
Finding a Dealer in Illegal Materials or Technologies:
Average (8+) Streetwise check (1D x 10 hours, INT).
Evading a Police Search: Streetwise check (1D x 10
minutes, INT) opposed by Recon (INT) check.
Survival
The Survival skill is the wilderness counterpart of the urban
Streetwise skill – the Traveller is trained to survive in the
wild, build shelters, hunt or trap animals, avoid exposure
and so forth. He can recognise plants and animals of
his homeworld and related planets, and can pick up on
common clues and traits even on unfamiliar worlds.
Gathering Supplies in the Wilderness to Survive for a
Week: Average (8+) Survival check (1D days, EDU).
Identifying a Poisonous Plant: Average (8+) Survival
check (1D x 10 seconds, INT or EDU).
Tactics
This skill covers tactical planning and decision making,
from board games to squad level combat to fleet
engagements. For use in combat, see Combat chapter.
SPECIALITIES
● Military: Co-ordinating the attacks of foot troops or
vehicles on the ground.
● Naval: Co-ordinating the attacks of a spacecraft or
fleet.
Developing a Strategy for Attacking an Enemy Base:
Average (8+) Tactics (military) check (1D x 10 hours,
INT).
Vacc Suit
The Vacc Suit skill allows a Traveller to wear and operate
spacesuits and environmental suits. A Traveller will
rarely need to make Vacc Suit checks under ordinary
circumstances – merely possessing the skill is enough. If
the Traveller does not have the requisite Vacc Suit skill
for the suit he is wearing, he suffers DM-2 to all skill
checks made while wearing a suit for each missing level.
This skill also permits the character to operate advanced
battle armour.
Performing a Systems Check on Battle Dress: Average
(8+) Vacc Suit check (1D minutes, EDU).


*/
