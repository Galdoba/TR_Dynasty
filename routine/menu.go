package routine

import "fmt"

var lastAction string

func enterMenu(menu string) {
	//fmt.Println("Start: enterMenu", menu)
	clrScrn()
	fmt.Println("Menu position-------------")
	if lastAction != "" {
		fmt.Println("lastAction =", lastAction)
	}
	switch menu {
	default:
		fmt.Println("TODO: menu", menu)
		quit = true
	case "":
		//clrScrn()
		fmt.Println("Start: startMenu", menu)
		startMenu()
	}

}

func startMenu() {
	opt, action := menu("Seclect Option:", "Disconnect", "Input Data", "Broker")
	switch opt {
	default:
	case 0:
		quit = true
	case 1:
		fmt.Println("\033[FAct 1")
	}
	lastAction = action
	fmt.Println("\033[F'" + action + "' action was chosen...")
}
