package routine

import (
	"fmt"
	"strconv"
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
	advanceTime(longestSearchTime)
	if userConfirm("Load Cargo?") {
		loadCargo()
	}
	lastAction = action
	menuPosition = ""
	//fmt.Println("\033[F'" + action + "' action was chosen...")
}

func hangarMenu() {
	opt, action := menu("Select Action:", "Return", "Unload Freight", "Load Freight", "Edit Cargo Space")
	clrScrn()
	switch opt {
	default:
	case 0:
		menuPosition = ""
	case 1:
		menuPosition = "HANGAR: UNLOAD CARGO"
		clrScrn()
		d := cargoDesignatedTo(sourceWorld)
		if d < 1 {
			menu("No designnated Cargo for this planet", "Return to HANGAR")
			menuPosition = "HANGAR"
			return
		}
		if userConfirm("Unload Freight for this planet? (" + strconv.Itoa(d) + " Lots)") {
			unloadCargo()
		}

	case 2:
		menuPosition = "HANGAR: LOCAL FREIGHT"
		clrScrn()
		if len(portCargo) < 1 {
			menu("No Data on local Freight", "Return to HANGAR")
			menuPosition = "HANGAR"
			break
		}
		loadCargo()
		menuPosition = "HANGAR"

	case 3:
		menuPosition = "HANGAR: EDIT CARGO SPACE"

		i, _ := menu("Select Action:", "Reserve Cargo Space", "Edit Cargo Space")
		switch i {
		case 0:
			reserveCargoSpace()
		case 1:
			editCargoEntryVolume()
		}
		menuPosition = "HANGAR"
	}

	lastAction = action
}

func arrival() {
	d := cargoDesignatedTo(sourceWorld)
	if d < 1 {
		return
	}
	if userConfirm("Unload Freight for this planet? (" + strconv.Itoa(d) + " Lots)") {
		unloadCargo()
		menu("------------------", "Continue")
	}

}
