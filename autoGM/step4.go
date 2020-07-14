package autoGM

import (
	"fmt"

	"github.com/Galdoba/convert"

	"github.com/Galdoba/utils"
)

var Highport bool
var HighTrafficSpace bool
var SettledSpace bool
var WildSpace bool
var AmberZone bool
var RedZone bool
var Unexplored bool

func step4SpaceTravel() {
	//opt, systemType := utils.TakeOptions("Pick system Type: ", "Civilised", "Newly Colonised", "Uninhibited")
	eventDice := "d6"
	eventTN := 6
	Highport = askYesNo("Highport?")
	HighTrafficSpace = askYesNo("High-Traffic Space?")
	SettledSpace = askYesNo("Settled Space?")
	WildSpace = askYesNo("Wild Space?")
	AmberZone = askYesNo("Amber Zone?")
	RedZone = askYesNo("Red Zone?")
	Unexplored = askYesNo("Unexplored Space?")
	end := false
	for !end {
		if !encounterHappens(eventDice, eventTN) {
			fmt.Println("No encounter")
		} else {
			fmt.Println("")
			fmt.Println("While travelling in system some thing happens:")
			fmt.Println(rollSpaceEncounter())
			fmt.Println("///")
			fmt.Println(rollSpaceMasterTable())
			fmt.Println("")
			fmt.Println("Making a Sensor Scan: Routine (6+) Electronics (sensors) check (1D x 10 minutes, INT or EDU).")
			fmt.Println("Effect (-6): Adjacent  - 1 km or less")
			fmt.Println("Effect (-5): Close  - 1-10 km")
			fmt.Println("Effect (-4): Short  - 11-1,250 km")
			fmt.Println("Effect (-3): Medium  - 1,251-10,000 km")
			fmt.Println("Effect (-2): Long  - 10,001-25,000 km")
			fmt.Println("Effect (-1): Very Long  - 25,001-50,000 km")
			fmt.Println("Effect (0+): Distant  - More than 50,000 km")
			fmt.Println("Target's Effect is", roll2d6()-6)
		}
		if !askYesNo("Roll again? ") {
			end = true
		}
	}
}

func rollSpaceMasterTable() string {
	switch rollD66() {
	case 11:
		return "Life Event (page 67)"
	case 12:
		return "Alien Probe (page 33)"
	case 13:
		return "Alien Space (page 33)"
	case 14:
		return "Anti-Matter Bomb (page 33)"
	case 15:
		return "Automatic Guard System (page 34)"
	case 16:
		return "Ancient Jump Gate (page 34)"
	case 21:
		return "Asteroid, Empty (page 34)"
	case 22:
		return "Asteroid, Inhabited (page 34)"
	case 23:
		return "Onboard Event (page 60)"
	case 24:
		return "Black Box (page 35)"
	case 25:
		return "Colossal Alien (page 35)"
	case 26:
		return "Distress call (page 35)"
	case 31:
		return "Escape Pod (page 36)"
	case 32:
		return "Gas Giant (page 36)"
	case 33:
		return "Giant Space Battle"
	case 34:
		return "Ship Encounters (page 44)"
	case 35:
		return "Higher Entity (page 36)"
	case 36:
		return "Living Planet (page 37)"
	case 41:
		return "Lost Astronauts (page 37)"
	case 42:
		return "Meteor Swarm (page 38)"
	case 43:
		return "Onboard Event (page 60)"
	case 44:
		return "Ship Encounters (page 44)"
	case 45:
		return "Mine Field (page 38)"
	case 46:
		return "Mysterious Radiation (page 38)"
	case 51:
		return "Psychic Field (page 39)"
	case 52:
		return "Secret Star System (page 40)"
	case 53:
		return "Space garbage (page 40)"
	case 54:
		return "Onboard Event (page 60)"
	case 55:
		return "Ship Encounters (page 44)"
	case 56:
		return "Space Natives (page 41)"
	case 61:
		return "Space Parasites (page 41)"
	case 62:
		return "Strange Communiqué (page 41)"
	case 63:
		return "Super AI (page 43)"
	case 64:
		return "Temporal Anomaly (page 43)"
	case 65:
		return "Wormhole (page 43)"
	case 66:
		return "Ship Encounters (page 44)"
	}
	return "Error"
}

func rollSpaceEncounter() string {
	dm := 0
	if Highport {
		dm = dm + 3
	}
	if HighTrafficSpace {
		dm = dm + 2
	}
	if SettledSpace {
		dm = dm + 1
	}
	if WildSpace {
		dm = dm - 1
	}
	if AmberZone {
		dm = dm - 2
	}
	if RedZone {
		dm = dm - 3
	}
	if Unexplored {
		dm = dm - 4
	}
	r1 := rolld6() + dm
	r1 = utils.BoundInt(r1, 0, 9)
	r2 := rolld6()
	rrStr := convert.ItoS(r1) + convert.ItoS(r2)
	return spaceEncounter(rrStr)
}

func spaceEncounter(rrStr string) string {
	switch rrStr {
	case "01":
		return "Alien derelict (possible salvage)"
	case "02":
		return "Solar flare (1D x 100 rads)"
	case "03":
		return "Asteroid (empty rock)"
	case "04":
		return "Ore-bearing asteroid (possible mining)"
	case "05":
		return "Alien vessel (on a mission)"
	case "06":
		return "Rock hermit (inhabited rock)"
	case "11":
		return "Pirate"
	case "12":
		return "Derelict vessel (possible salvage)"
	case "13":
		return "Space station (1-4: derelict)"
	case "14":
		return "Comet (may be ancient derelict at its core)"
	case "15":
		return "Ore-bearing asteroid (possible mining)"
	case "16":
		return "Ship in distress (roll again for type)"
	case "21":
		return "Pirate"
	case "22":
		return "Free trader"
	case "23":
		return "Micrometeorite Storm (collision!)\nCollision!: Almost any collision at high speed is lethal even for the most powerful spacecraft. In this case, the ship has collided with a tiny object that has nevertheless smashed into the hull. The ship suffers 1D damage."
	case "24":
		return "Hostile vessel (roll again for type)"
	case "25":
		return "Mining ship"
	case "26":
		return "Scout Ship"
	case "31":
		return "Alien vessel (1-3: trader, 4-6: explorer, 6: spy)"
	case "32":
		return "Space junk (possible salvage)"
	case "33":
		return "Far Trader"
	case "34":
		return "Derelict (possible salvage)"
	case "35":
		return "Safari or science vessel"
	case "36":
		return "Escape pod"
	case "41":
		return "Passenger liner"
	case "42":
		return "Ship in distress (roll again for type)\nDistress Signals: Ships transmit the standard timestamped SOS message (also known as Mayday in Solomani or Signal GK in Vilani within the Third Imperium setting) when in distress. Any vessel who detects an SOS is legally required to respond and offer assistance or contact the authorities. Failure to render assistance is a criminal offence, but the harsh requirements of life support and orbital mechanics mean that many deaths in space are slow ones, where a crew know they are doomed but have days or weeks in which to contemplate it.  Most ships carry emergency low berths where the crew can freeze themselves and wait for rescue. Some distress calls are fakes, intended to draw ships in so they can be attacked."
	case "43":
		return "Colony ship or passenger liner"
	case "44":
		return "Scout ship"
	case "45":
		return "Space station"
	case "46":
		return "X-Boat Courier"
	case "51":
		return "Hostile vessel (roll again for type)"
	case "52":
		return "Garbage ejected from a ship"
	case "53":
		return "Medical ship or hospital"
	case "54":
		return "Lab ship or scout"
	case "55":
		return "Patron"
	case "56":
		return "Police ship"
	case "61":
		return "Unusually daring pirate"
	case "62":
		return "Noble yacht"
	case "63":
		return "Warship"
	case "64":
		return "Cargo vessel"
	case "65":
		return "Navigational Buoy or Beacon"
	case "66":
		return "Unusual ship"
	case "71":
		return "Collision with space junk (collision!)\nCollision!: Almost any collision at high speed is lethal even for the most powerful spacecraft. In this case, the ship has collided with a tiny object that has nevertheless smashed into the hull. The ship suffers 1D damage."
	case "72":
		return "Automated vessel"
	case "73":
		return "Free Trader"
	case "74":
		return "Dumped cargo pod (roll on random trade goods CRB p.211)"
	case "75":
		return "Police vessel"
	case "76":
		return "Cargo hauler"
	case "81":
		return "Passenger liner"
	case "82":
		return "Orbital factory (roll on random trade goods CRB p.211)"
	case "83":
		return "Orbital habitat"
	case "84":
		return "Orbital habitat"
	case "85":
		return "Communications Satellite"
	case "86":
		return "Defence Satellite"
	case "91":
		return "Pleasure craft"
	case "92":
		return "Space station"
	case "93":
		return "Police vessel"
	case "94":
		return "Cargo hauler"
	case "95":
		return "System Defence Boat"
	case "96":
		return "Grand Fleet warship"
	}
	return "Error"
}
