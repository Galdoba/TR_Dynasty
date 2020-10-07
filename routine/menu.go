package routine

import (
	"fmt"
)

var lastAction string

func enterMenu(menu string) {
	//fmt.Println("Start: enterMenu", menu)
	//clrScrn()
	if lastAction != "" {
		fmt.Println("lastAction =", lastAction)
	}

	switch menu {
	default:
		menuReturm()
	case "":
		clrScrn()
		startMenu()
	case "INPUT":
		clrScrn()
		inputMenu()
	case "HIRE BROKER":
		clrScrn()
		chooseBroker()
		//menuPosition = ""
	case "SEARCH":
		clrScrn()
		trafficMenu()
	case "HANGAR":
		clrScrn()
		hangarMenu()

	}

}

func startMenu() {
	opt, action := menu("Select Action:", "Disconnect", "Input Data", "Search", "Hangar")
	switch opt {
	default:
	case 0:
		quit = true
	case 1:
		menuPosition = "INPUT"
	case 2:
		menuPosition = "SEARCH"
	case 3:
		menuPosition = "HANGAR"
	}
	lastAction = action
	fmt.Println("\033[F'" + action + "' action was chosen...")

}

func inputMenu() {
	opt, action := menu("Select Action:", "Return", "Set Destination", "Hire Broker")
	switch opt {
	default:
	case 0:
		menuPosition = ""
	case 1:
		menuPosition = "SET DESTINATION"
		pickDestinationWorld()
		//fmt.Println("MenuPosition = 21")

	case 2:
		menuPosition = "HIRE BROKER"
		//fmt.Println("MenuPosition = 22")
		chooseBroker()

	}
	lastAction = action
	menuPosition = ""
	//fmt.Println("\033[F'" + action + "' action was chosen...")
}

func menuReturm() {
	opt, _ := menu("Select Action:", "Return to MAIN")
	switch opt {
	default:
		menuPosition = ""
	}
}

func trafficMenu() {
	if targetWorld.Name() == "--NO DATA--" {
		//menuPosition = "INPUT"
		fmt.Println("Destination World Undefined")
		pickDestinationWorld()
		clrScrn()
	}
	opt, action := menu("Select Action:", "Return", "Search Passengers", "Search Freight", "Search Mail", "Search All", "Let local Broker handle it")
	clrScrn()
	switch opt {
	default:
	case 0:
		menuPosition = ""
	case 1:
		menuPosition = "SEARCH PASSENGERS"
		PassengerRoutine()
	case 2:
		menuPosition = "SEARCH FREIGHT"
		FreightRoutine()

	case 3:
		menuPosition = "SEARCH MAIL"
		MailRoutine()

	case 4:
		menuPosition = "SEARCH ALL"
		PassengerRoutine()
		FreightRoutine()
		MailRoutine()

	case 5:
		autoMod = true
		menuPosition = "SEARCH ALL"
		if localBroker.cut == 0.0 {
			fmt.Println("Local Broker was not hired")
			chooseBroker()
		}
		PassengerRoutine()
		FreightRoutine()
		MailRoutine()
		autoMod = false

	}
	i, _ := menu("Load Cargo?", "Yes", "No")
	if i == 0 {
		loadCargo()
	}
	lastAction = action

	//fmt.Println("\033[F'" + action + "' action was chosen...")
}

func hangarMenu() {
	opt, action := menu("Select Action:", "Return", "Unload Freight", "Load Freight")
	clrScrn()
	switch opt {
	default:
	case 0:
		menuPosition = ""
	case 1:
		menuPosition = "HANGAR: UNLOAD CARGO"
		unloadCargo()

	case 2:
		menuPosition = "HANGAR: LOCAL FREIGHT"

		if len(portCargo) < 1 {
			fmt.Println("No Data on local Freight")
			break
		}
		loadCargo()
	}

	lastAction = action
}

func arrival() {
	i, _ := menu("Unload Freight for this planet?", "Yes", "No")
	if i == 0 {
		unloadCargo()
		menu("------------------", "Continue")
	}
}
