package mission

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
)

/*
Checklist
Roll on Random Patron Table (pg.81)
Roll for a Random Trait (pg.76)
Roll for a Random Encounter (either Starport, Rural or Urban) (pg.82)
Select or randomly decide whether focus is a Person or Other.
	If a Person, roll once on the Contact, Allies and Enemies table (pg. 76)
	If Other, roll for a Random Mission Target (pg. 81)
	If an object, vessel or trade good was rolled for Other, then roll on the Nature of the Item table (see below)
Roll for a Random Patron Mission (pg.81)
Roll a Random Journey (see below)
Roll a Random Place (see below)
Roll for a Random Opposition (pg. 82)
Roll for a second Random Encounter (either Starport, Rural or Urban) (pg.82)
*/

func Test() {
	patron := randomPatron()
	fmt.Println("Patron:", patron)
	trait := randomTrait()
	fmt.Println("Trait:", trait)
	encounter := randomEncounter()
	fmt.Println("Encounter:", encounter)
	focus := randomFocus()
	fmt.Println("Focus:", focus)
	patronMission := randomPatronMission()
	fmt.Println("Patron Mission:", patronMission)
	journey := randomJourney()
	fmt.Println("Journey:", journey)
	place := randomPlace()
	fmt.Println("Place:", place)
	opposition := randomOpposition()
	fmt.Println("Opposition:", opposition)
}

func randomPatron() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Assassin "
	case "41":
		return "Merchant "
	case "12":
		return "Smuggler "
	case "42":
		return "Free Trader "
	case "13":
		return "Terrorist "
	case "43":
		return "Broker "
	case "14":
		return "Embezzler "
	case "44":
		return "Corporate Executive "
	case "15":
		return "Thief "
	case "45":
		return "Corporate Agent "
	case "16":
		return "Revolutionary "
	case "46":
		return "Financier"
	case "21":
		return "Clerk "
	case "51":
		return "Belter "
	case "22":
		return "Administrator "
	case "52":
		return "Researcher "
	case "23":
		return "Mayor "
	case "53":
		return "Naval Officer "
	case "24":
		return "Minor Noble "
	case "54":
		return "Pilot "
	case "25":
		return "Physician "
	case "55":
		return "Starport Administrator "
	case "26":
		return "Tribal Leader "
	case "56":
		return "Scout"
	case "31":
		return "Diplomat "
	case "61":
		return "Alien "
	case "32":
		return "Courier "
	case "62":
		return "Playboy "
	case "33":
		return "Spy "
	case "63":
		return "Stowaway "
	case "34":
		return "Ambassador "
	case "64":
		return "Family Relative "
	case "35":
		return "Noble "
	case "65":
		return "Agent of a Foreign Power "
	case "36":
		return "Police Officer "
	case "66":
		return "Imperial Agent"
	}
}

func randomTrait() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Loyal "
	case "41":
		return "Rumour-monger "
	case "12":
		return "Distracted by other worries "
	case "42":
		return "Unusually provincial "
	case "13":
		return "In debt to criminals "
	case "43":
		return "Drunkard or drug addict "
	case "14":
		return "Makes very bad jokes "
	case "44":
		return "Government informant "
	case "15":
		return "Will betray characters "
	case "45":
		return "Mistakes a player character for someone else "
	case "16":
		return "Aggressive "
	case "46":
		return "Possesses unusually advanced technology "
	case "21":
		return "Has secret allies "
	case "51":
		return "Unusually handsome or beautiful "
	case "22":
		return "Secret anagathic user "
	case "52":
		return "Spying on the characters "
	case "23":
		return "Looking for something "
	case "53":
		return "Possesses TAS membership "
	case "24":
		return "Helpful "
	case "54":
		return "Is secretly hostile towards the characters "
	case "25":
		return "Forgetful "
	case "55":
		return "Wants to borrow money "
	case "26":
		return "Wants to hire the chracters "
	case "56":
		return "Is convinced the characters are dangerous "
	case "31":
		return "Has useful contacts "
	case "61":
		return "Involved in political intrigue "
	case "32":
		return "Artistic "
	case "62":
		return "Has a dangerous secret "
	case "33":
		return "Easily confused "
	case "63":
		return "Wants to get offplanet as soon as possible "
	case "34":
		return "Unusually ugly "
	case "64":
		return "Attracted to a player character "
	case "35":
		return "Worried about current situation "
	case "65":
		return "From offworld "
	case "36":
		return "Shows pictures of his children "
	case "66":
		return "Possesses telepathy or other unusual quality"
	}
}

func randomEncounter() string {
	d := dice.Roll("1d3").Sum()
	switch d {
	default:
		return "Error"
	case 1:
		return starportEncounter()
	case 2:
		return urbanEncounter()
	case 3:
		return ruralEncounter()
	}
}

func starportEncounter() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Maintenance robot at work "
	case "41":
		return "Traders offer spare parts and supplies at cut-price rates "
	case "12":
		return "Trade ship arrives or departs "
	case "42":
		return "Repair yard catches fire "
	case "13":
		return "Captain argues about fuel prices "
	case "43":
		return "Passenger liner arrives or departs "
	case "14":
		return "News report about pirate activity on a starport screen draws a crowd "
	case "44":
		return "Servant robot offers to guide characters around the spaceport "
	case "15":
		return "Bored clerk makes life difficult for the characters "
	case "45":
		return "Trader from a distant system selling strange curios "
	case "16":
		return "Local merchant with cargo to transport seeks a ship "
	case "46":
		return "Old crippled belter asks for spare change and complains about drones taking his job "
	case "21":
		return "Dissident tries to claim sanctuary from planetary authorities "
	case "51":
		return "Patron offers the characters a job "
	case "22":
		return "Traders from offworld argue with local brokers "
	case "52":
		return "Passenger looking for a ship "
	case "23":
		return "Technician repairing starport computer system "
	case "53":
		return "Religious pilgrims try to convert the characters "
	case "24":
		return "Reporter asks for news from offworld "
	case "54":
		return "Cargo hauler arrives or departs "
	case "25":
		return "Bizarre cultural performance "
	case "55":
		return "Scout ship arrives or departs "
	case "26":
		return "Patron argues with another group of travellers "
	case "56":
		return "Illegal or dangerous goods are impounded "
	case "31":
		return "Military vessel arrives or departs "
	case "61":
		return "Pickpocket tries to steal from the characters "
	case "32":
		return "Demonstration outside starport "
	case "62":
		return "Drunken crew pick a fight "
	case "33":
		return "Escaped prisoners begs for passage offworld "
	case "63":
		return "Government officials investigate the characters "
	case "34":
		return "Impromptu bazaar of bizarre items "
	case "64":
		return "Random security sweep scans characters & baggage "
	case "35":
		return "Security patrol "
	case "65":
		return "Starport is temporarily shut down for security reasons "
	case "36":
		return "Unusual alien "
	case "66":
		return "Damaged ship makes emergency docking"
	}
}

func urbanEncounter() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Street riot in progress "
	case "41":
		return "Security Patrol "
	case "12":
		return "Characters pass a charming restaurant "
	case "42":
		return "Ancient building or archive "
	case "13":
		return "Trader in illegal goods "
	case "43":
		return "Festival "
	case "14":
		return "Public argument "
	case "44":
		return "Someone is following the characters "
	case "15":
		return "Sudden change of weather "
	case "45":
		return "Unusual cultural group or event "
	case "16":
		return "NPC asks for the character’s help "
	case "46":
		return "Planetary official "
	case "21":
		return "Characters pass a bar or pub "
	case "51":
		return "Characters spot someone they recognise "
	case "22":
		return "Characters pass a theatre or other entertainment venue "
	case "52":
		return "Public demonstration "
	case "23":
		return "Curiosity Shop "
	case "53":
		return "Robot or other servant passes characters "
	case "24":
		return "Street market stall tries to sell the characters something "
	case "54":
		return "Prospective patron "
	case "25":
		return "Fire, dome breach or other emergency in progress "
	case "55":
		return "Crime such as robbery or attack in progress "
	case "26":
		return "Attempted robbery of characters "
	case "56":
		return "Street preacher rants at the characters "
	case "31":
		return "Vehicle accident involving characters "
	case "61":
		return "News broadcast on public screens "
	case "32":
		return "Low-ﬂying spacecraft ﬂies overhead "
	case "62":
		return "Sudden curfew or other restriction on movement "
	case "33":
		return "Alien or other offworlder "
	case "63":
		return "Unusually empty or quiet street "
	case "34":
		return "Random NPC bumps into character "
	case "64":
		return "Public announcement "
	case "35":
		return "Pickpocket "
	case "65":
		return "Sports event "
	case "36":
		return "Media team or journalist "
	case "66":
		return "Imperial Dignitary"

	}
}

func ruralEncounter() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Wild Animal "
	case "41":
		return "Wild Animal "
	case "12":
		return "Agricultural robots "
	case "42":
		return "Small community – quiet place to live "
	case "13":
		return "Crop sprayer drone ﬂies overhead "
	case "43":
		return "Small community – on a trade route "
	case "14":
		return "Damaged agricultural robot being repaired "
	case "44":
		return "Small community – festival in progress "
	case "15":
		return "Small, isolationist community "
	case "45":
		return "Small community – in danger "
	case "16":
		return "Noble hunting party "
	case "46":
		return "Small community – not what it seems "
	case "21":
		return "Wild Animal "
	case "51":
		return "Wild Animal "
	case "22":
		return "Local landing field "
	case "52":
		return "Unusual weather "
	case "23":
		return "Lost child "
	case "53":
		return "Difficult terrain "
	case "24":
		return "Travelling merchant caravan "
	case "54":
		return "Unusual creature "
	case "25":
		return "Cargo convoy "
	case "55":
		return "Isolated homestead – welcoming "
	case "26":
		return "Police chase "
	case "56":
		return "Isolated homestead – unfriendly "
	case "31":
		return "Wild Animal "
	case "61":
		return "Wild Animal "
	case "32":
		return "Telecommunications black spot "
	case "62":
		return "Private villa "
	case "33":
		return "Security patrol "
	case "63":
		return "Monastery or retreat "
	case "34":
		return "Military facility "
	case "64":
		return "Experimental farm "
	case "35":
		return "Bar or waystation "
	case "65":
		return "Ruined structure "
	case "36":
		return "Grounded spacecraft "
	case "66":
		return "Research facility"

	}
}

func randomFocus() string {
	d := dice.Roll("1d2").Sum()
	switch d {
	default:
		return "Error"
	case 1:
		return personFocus()
	case 2:
		return objectFocus()

	}
}

func personFocus() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Naval Officer "
	case "41":
		return "Bored Noble "
	case "12":
		return "Imperial Diplomat "
	case "42":
		return "Planetary Governor "
	case "13":
		return "Crooked Trader "
	case "43":
		return "Inveterate Gambler "
	case "14":
		return "Medical Doctor "
	case "44":
		return "Crusading Journalist "
	case "15":
		return "Eccentric Scientist "
	case "45":
		return "Doomsday Cultist "
	case "16":
		return "Mercenary "
	case "46":
		return "Corporate Agent "
	case "21":
		return "Famous Performer "
	case "51":
		return "Criminal Syndicate "
	case "22":
		return "Alien Thief "
	case "52":
		return "Military Governor "
	case "23":
		return "Free Trader "
	case "53":
		return "Army Quartermaster "
	case "24":
		return "Explorer "
	case "54":
		return "Private Investigator "
	case "25":
		return "Marine Captain "
	case "55":
		return "Starport Administrator "
	case "26":
		return "Corporate Executive "
	case "56":
		return "Retired Admiral "
	case "31":
		return "Researcher "
	case "61":
		return "Alien Ambassador "
	case "32":
		return "Cultural Attaché "
	case "62":
		return "Smuggler "
	case "33":
		return "Religious Leader "
	case "63":
		return "Weapons Inspector "
	case "34":
		return "Conspirator "
	case "64":
		return "Elder Statesman "
	case "35":
		return "Rich Noble "
	case "65":
		return "Planetary Warlord "
	case "36":
		return "Artificial Intelligence "
	case "66":
		return "Imperial Agent"
	}
}

func objectFocus() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Common Trade Goods " + natureOfItem()
	case "41":
		return "Roll on the Patron Table "
	case "12":
		return "Common Trade Goods " + natureOfItem()
	case "42":
		return "Roll on the Patron Table "
	case "13":
		return "Random Trade Goods " + natureOfItem()
	case "43":
		return "Roll on the Patron Table "
	case "14":
		return "Random Trade Goods " + natureOfItem()
	case "44":
		return "Roll on the Random Opposition table "
	case "15":
		return "Illegal Trade Goods " + natureOfItem()
	case "45":
		return "Roll on the Random Opposition table "
	case "16":
		return "Illegal Trade Goods " + natureOfItem()
	case "46":
		return "Roll on the Random Opposition table "
	case "21":
		return "Computer Data " + natureOfItem()
	case "51":
		return "Local Government "
	case "22":
		return "Alien Artefact " + natureOfItem()
	case "52":
		return "Planetary Government "
	case "23":
		return "Personal Effects " + natureOfItem()
	case "53":
		return "Corporation "
	case "24":
		return "Work of Art " + natureOfItem()
	case "54":
		return "Imperial Intelligence "
	case "25":
		return "Historical Artefact " + natureOfItem()
	case "55":
		return "Criminal Syndicate "
	case "26":
		return "Weapon " + natureOfItem()
	case "56":
		return "Criminal Gang "
	case "31":
		return "Starport "
	case "61":
		return "Free Trader " + natureOfItem()
	case "32":
		return "Asteroid Base "
	case "62":
		return "Yacht " + natureOfItem()
	case "33":
		return "City "
	case "63":
		return "Cargo Hauler " + natureOfItem()
	case "34":
		return "Research station "
	case "64":
		return "Police Cutter " + natureOfItem()
	case "35":
		return "Bar or Nightclub "
	case "65":
		return "Space Station " + natureOfItem()
	case "36":
		return "Medical Facility "
	case "66":
		return "Warship" + natureOfItem()

	}
}

func natureOfItem() string {
	switch dice.Roll("1d6").Sum() {
	default:
		return "Error"
	case 1:
		return "(Lost)"
	case 2:
		return "(Hidden)"
	case 3:
		return "(Destroyed)"
	case 4:
		return "(Captured)"
	case 5:
		return "(Dangerous)"
	case 6:
		return "(Stolen)"
	}
}

func randomPatronMission() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Assassinate a target "
	case "41":
		return "Investigate a crime "
	case "12":
		return "Frame a target "
	case "42":
		return "Investigate a theft "
	case "13":
		return "Destroy a target "
	case "43":
		return "Investigate a murder "
	case "14":
		return "Steal from a target "
	case "44":
		return "Investigate a mystery "
	case "15":
		return "Aid in a burglary "
	case "45":
		return "Investigate a target "
	case "16":
		return "Stop a burglary "
	case "46":
		return "Investigate an event "
	case "21":
		return "Retrieve data or an object from a secure facility "
	case "51":
		return "Join an expedition "
	case "22":
		return "Discredit a target "
	case "52":
		return "Survey a planet "
	case "23":
		return "Find a lost cargo "
	case "53":
		return "Explore a new system "
	case "24":
		return "Find a lost person "
	case "54":
		return "Explore a ruin "
	case "25":
		return "Deceive a target "
	case "55":
		return "Salvage a ship "
	case "26":
		return "Sabotage a target "
	case "56":
		return "Capture a creature "
	case "31":
		return "Transport goods "
	case "61":
		return "Hijack a ship "
	case "32":
		return "Transport a person "
	case "62":
		return "Entertain a noble "
	case "33":
		return "Transport data "
	case "63":
		return "Protect a target "
	case "34":
		return "Transport goods secretly "
	case "64":
		return "Save a target "
	case "35":
		return "Transport goods quickly "
	case "65":
		return "Aid a target "
	case "36":
		return "Transport dangerous goods"
	case "66":
		return "It’s a trap – the patron intends to betray the characters"
	}
}

func randomJourney() string {
	switch dice.Roll("1d6").Sum() {
	default:
		return "Error"
	case 1:
		return "Wilderness Trek – Across the surface of a world. "
	case 2:
		return "Wilderness Trek – Across the surface of a world. "
	case 3:
		return "Space Travel – Either insystem, or to another world via jump drive. "
	case 4:
		return "City – Moving across town meeting contacts, checking locations. "
	case 5:
		return "Building – The infiltration of a large building or other complex. "
	case 6:
		return "Building – The infiltration of a large building or other complex. "
	}
}

func randomPlace() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Warehouse "
	case "31":
		return "Junkyard "
	case "51":
		return "Skyscraper "
	case "12":
		return "Hotel "
	case "32":
		return "Police HQ "
	case "52":
		return "Uptown Office "
	case "13":
		return "Temple "
	case "33":
		return "Laboratory "
	case "53":
		return "Industrial Unit "
	case "14":
		return "Starport Terminal "
	case "34":
		return "Bar "
	case "54":
		return "Factory "
	case "15":
		return "Powerstation "
	case "35":
		return "Nightclub "
	case "55":
		return "Fuel Dump "
	case "16":
		return "Space-station "
	case "36":
		return "Restaurant "
	case "56":
		return "Government Building "
	case "21":
		return "Starship "
	case "41":
		return "Back Alley "
	case "61":
		return "Penthouse Appt. "
	case "22":
		return "Ship’s Boat "
	case "42":
		return "Vehicle Park "
	case "62":
		return "Crime Base "
	case "23":
		return "Remote Outpost "
	case "43":
		return "Fast-food Bar "
	case "63":
		return "Tenement Block "
	case "24":
		return "Museum "
	case "44":
		return "Casino "
	case "64":
		return "Suburban House "
	case "25":
		return "Shopping Mall "
	case "45":
		return "Villa "
	case "65":
		return "Sewers "
	case "26":
		return "Farming Complex "
	case "46":
		return "Carnival/Parade "
	case "66":
		return "Theatre"

	}
}

func randomOpposition() string {
	switch dice.RollD66() {
	default:
		return "Error"
	case "11":
		return "Animals "
	case "41":
		return "Target is in deep space "
	case "12":
		return "Large animal "
	case "42":
		return "Target is in orbit "
	case "13":
		return "Bandits and thieves "
	case "43":
		return "Hostile weather conditions "
	case "14":
		return "Fearful peasants "
	case "44":
		return "Dangerous organisms or radiation "
	case "15":
		return "Local authorities "
	case "45":
		return "Target is in a dangerous region "
	case "16":
		return "Local lord "
	case "46":
		return "Target is in a restricted area"
	case "21":
		return "Criminals – thugs or corsairs "
	case "51":
		return "Target is under electronic observation "
	case "22":
		return "Criminals – thieves or saboteurs "
	case "52":
		return "Hostile guard robots or ships "
	case "23":
		return "Police – ordinary security forces "
	case "53":
		return "Biometric identification required "
	case "24":
		return "Police – inspectors and detectives "
	case "54":
		return "Mechanical failure or computer hacking "
	case "25":
		return "Corporate – agents "
	case "55":
		return "Characters are under surveillance "
	case "26":
		return "Corporate – legal "
	case "56":
		return "Out of fuel or ammunition"
	case "31":
		return "Starport security "
	case "61":
		return "Police investigation "
	case "32":
		return "Imperial marines "
	case "62":
		return "Legal barriers "
	case "33":
		return "Interstellar corporation "
	case "63":
		return "Nobility "
	case "34":
		return "Alien – private citizen or corporation "
	case "64":
		return "Government officials "
	case "35":
		return "Alien – government "
	case "65":
		return "Target is protected by a third party "
	case "36":
		return "Space travellers or rival ship "
	case "66":
		return "Hostages "
	}
}
