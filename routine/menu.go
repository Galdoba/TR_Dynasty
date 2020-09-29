package routine

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
)

var lastAction string

func enterMenu(menu string) {
	//fmt.Println("Start: enterMenu", menu)
	clrScrn()
	if lastAction != "" {
		fmt.Println("lastAction =", lastAction)
	}
	fmt.Println("Menu position: MAIN >", menuPosition)

	switch menu {
	default:
		fmt.Println("TODO: menu", menu)
		quit = true
	case "":
		//clrScrn()
		startMenu()
	case "INPUT":
		//clrScrn()
		inputMenu()
	case "HIRE BROKER":
		//clrScrn()
		chooseBroker()
		menuPosition = ""
	}

}

func startMenu() {
	opt, action := menu("Select Action:", "Disconnect", "Input Data", "Broker")
	switch opt {
	default:
	case 0:
		quit = true
	case 1:
		menuPosition = "INPUT"
	case 2:
		menuPosition = "BROKER"
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
		//fmt.Println("MenuPosition = 21")
		targetWorld = pickWorld()
		distance = Astrogation.JumpDistance(sourceWorld.Hex(), targetWorld.Hex())
		ptValue = passengerTrafficValue(sourceWorld, targetWorld)
		ftValue = freightTrafficValue(sourceWorld, targetWorld)
		jumpRoutelocal, err := inputJumpRoute()
		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
		jumpRoute = jumpRoutelocal
	case 2:
		menuPosition = "HIRE BROKER"
		//fmt.Println("MenuPosition = 22")
		chooseBroker()
	}
	lastAction = action
	menuPosition = ""
	fmt.Println("\033[F'" + action + "' action was chosen...")
}
