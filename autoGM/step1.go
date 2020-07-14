package autoGM

import (
	"fmt"

	"github.com/Galdoba/utils"
)

func step1JobHunting() {
	pos, ans := utils.TakeOptions("Describe City:", "Safe City", "Dangerous City")
	fmt.Println("picked", pos, ans)
	encounterDice := ""
	encounterTN := 0
	encounterPeriod := "d6"
	switch ans {
	case "Safe City":
		encounterDice = "d6"
		encounterTN = 6

	case "Dangerous City":
		encounterDice = "d6"
		encounterTN = 6
	}
	maxDays := utils.InputInt("Set expected Days: ")
	endScene := false
	daysSpent := 1
	nextRollIn := 0
	if daysSpent > maxDays {
		endScene = true
	}
	for !endScene {

		if nextRollIn == 0 {
			if encounterHappens(encounterDice, encounterTN) {
				table := urbanEventType()
				fmt.Println("Day", daysSpent, "event:  ", rollEncounterTable(table))
			} else {
				fmt.Println("Day", daysSpent, " - no encounter")
			}
			nextRollIn = utils.RollDice(encounterPeriod)
		} else {
			fmt.Println("Day", daysSpent, " - no events")
		}

		//4/;.;.;.;.;.;.;.
		if daysSpent >= maxDays {
			endScene = askYesNo(" End Scene?")
		}
		nextRollIn--
		daysSpent++

	}
	if askYesNo(" Job Offer?") {
		fmt.Println(rollRandomMission())
	}
}

func urbanEventType() string {
	if roll2d6() == 12 {
		return "Global Events"
	}
	return "Local Events"
}

func cityEventLocal() string {
	r := utils.RollDice("2d6")
	result := ""
	switch r {
	case 2:
		result = "Civil Unrest (page 9)"
	case 3:
		result = "Hostage Situation (page 10)"
	case 4:
		result = "Crime Spree (page 9)"
	case 5:
		result = "Scientific Mishap (page 11)"
	case 6:
		result = "Industrial Mishap (page 10)"
	case 7:
		result = "Surprising Discovery (page 12)"
	case 8:
		result = "Alien Visitors (page 9)"
	case 9:
		result = "Tournament (page 13)"
	case 10:
		result = "Festival (page 10)"
	case 11:
		result = "Accident (page 9)"
	case 12:
		result = "Chance Encounter (page 9)"
	}
	return result
}

func eventGlobal() string {
	res := ""
	switch rolld6() {
	case 1:
		res = "Civil War"
	case 2:
		res = "Industrial Disaster"
	case 3:
		res = "Invasion"
	case 4:
		res = "Natural Disaster"
	case 5:
		res = "Failed Experiment"
	case 6:
		res = "Zombie Apocalypse"
	}
	return res
}

func rollRandomMission() string {
	mission := ""
	switch rollD66() {
	case 11:
		mission = "Capture a ship (page 44) and bring it to a secret asteroid base."
	case 12:
		mission = "Locate and return a ship captured by pirates (page 44)."
	case 13:
		mission = "Locate and return a person whose ship crashed on an uncharted planet (page 28)."
	case 14:
		mission = "Kidnap a person (page 10) from the city."
	case 15:
		mission = "Help small community protect itself from danger (page 20)."
	case 16:
		mission = "Rescue kidnapped person (page 10)."
	case 21:
		mission = "Transport goods to a distant planet."
	case 22:
		mission = "Sell goods on a distant planet."
	case 23:
		mission = "Capture a rare animal (page 31) that lives on a newly discovered planet (page 105)."
	case 24:
		mission = "Assassinate a villainous figure (page 153)."
	case 25:
		mission = "Assassinate an innocent man."
	case 26:
		mission = "Destroy an enemy ship (page 44)."
	case 31:
		mission = "Destroy an enemy building."
	case 32:
		mission = "Infiltrate into a laboratory and steal secret information."
	case 33:
		mission = "Locate a higher entity (page 36) and ask it an existential question."
	case 34:
		mission = "Research a newly discovered wormhole (page 43) and come back sane enough to make a coherent report."
	case 35:
		mission = "Protect a person targeted for assassination."
	case 36:
		mission = "Provoke a war between two races or two nations on the same planet."
	case 41:
		mission = "Rob a bank."
	case 42:
		mission = "Catch an escaped criminal (page 9) who fled into space on a stolen spacecraft."
	case 43:
		mission = "Catch an escaped criminal (page 9) who fled into the wilderness (page 105)."
	case 44:
		mission = "Stop a mysterious figure moving about the city and killing members of a single power group (page 86)."
	case 45:
		mission = "Find a planet in neutral space fit for human colonisation."
	case 46:
		mission = "Cure or destroy a colossal alien that keeps eating merchant ships in space."
	case 51:
		mission = "Transport a very expensive item from one noble to another."
	case 52:
		mission = "Find cure for a mysterious illness afflicting an isolated and primitive community."
	case 53:
		mission = "Discover what happened to a child who disappeared without a trace (page 29)."
	case 54:
		mission = "Discover why kids keep getting sick in school and why the government does not do anything about it."
	case 55:
		mission = "Transport exotic and potentially deadly beast to a distant planet."
	case 56:
		mission = "Deliver a ransom to pirates and return with the kidnapped person."
	case 61:
		mission = "Escort a tourist on a mad tour across the galaxy."
	case 62:
		mission = "Establish peaceful contact with space natives (page 41)."
	case 63:
		mission = "Escort a civilian vessel on a long journey through enemy space."
	case 64:
		mission = "Gather reconnaissance on an approaching fleet of unknown alien ships."
	case 65:
		mission = "Break through enemy forces to bring supplies to a starving besieged colony."
	case 66:
		mission = "Bravely go where no man has gone before!"
	}
	destinations := ""
	switch rolld6() {
	case 1, 2:
		destinations = "Space Master Table (page 32)"
	case 3, 4:
		destinations = "Urban Global table (page 13)"
	case 5:
		destinations = "Village table (page 103)"
	case 6:
		destinations = "Wilderness Man made (page 25)"
	}
	distance := ""
	switch rolld6() {
	case 1:
		distance = "City (Less than 100 km)"
	case 2:
		distance = "State (101–5,000 km)"
	case 3:
		distance = "World (In System)"
	case 4:
		distance = "Neighbour System (1 Jump)"
	case 5:
		distance = "Subsector (2–12 Parsecs)"
	case 6:
		distance = "Sector (10–60 parsecs)"
	}
	return "Mission: " + mission + "\n" + "Distance: " + distance + "\n" + "Destination: " + destinations
}
