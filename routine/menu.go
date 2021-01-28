package routine

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	starport "github.com/Galdoba/TR_Dynasty/Starport"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/profile"

	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/dice"
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
	case "INFORMATION":
		clrScrn()
		infoMenu()
	case "MARKET":
		clrScrn()
		marketMenu()
	case "DEPART":
		clrScrn()
		depart()
		if userConfirm("End program") {
		}
	}

}

func startMenu() {
	opt, action := menu("Select Action:", "Disconnect", "Input Data", "Search", "Hangar", "Information", "Market", "Depart")
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
	case 4:
		menuPosition = "INFORMATION"
	case 5:
		menuPosition = "MARKET"
	case 6:
		menuPosition = "DEPART"
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
		i, _ := menu("Select Action:", "Reserve Cargo Space", "Edit Cargo Space", "Add Cargo by Code")
		switch i {
		case 0:
			reserveCargoSpace()
		case 1:
			editCargoEntryVolume()
		case 2:
			addCargoByCode()
		}
		menuPosition = "HANGAR"
	}

	lastAction = action
}

func arrival() {
	unload := false
	if passengersDesignatedTo(sourceWorld) > 0 {
		unload = true
		unloadPassengers()
	}
	if cargoDesignatedTo(sourceWorld) > 0 {
		unload = true
		unloadCargo()
	}
	if unload {
		userConfirm("Continue")
	}

}

func infoMenu() {
	i, val := menu("Select category:", "Return", "Ship", "Port", "Planet")
	menuPosition = "INFORMATION"
	if i == 0 {
		menuPosition = ""
		return
	}
	menuPosition += " > " + val
	clrScrn()
	switch i {
	case 1:
		fmt.Println(shipInfo())
	case 2:
		fmt.Println(portInfo())
	case 3:
		fmt.Println("TODO: planetInfo()")
	}
	userConfirm("Continue ")
	menuPosition = "INFORMATION"
}

var localSupplier trade.Merchant
var localMarket []cargoLot

func marketMenu() {
	i, val := menu("Select category:", "Return", "Search Supplier", "Purchase Trade Goods", "Sell Trade Goods")

	//localSupplier = localSupplier.SetTraderDice(d.RollNext("3d6").Sum())
	menuPosition = "MARKET > " + val
	if i == 0 {
		menuPosition = ""
		return
	}
	eff := 0
	timeLimit := 999

	time := dice.Roll("1d6").Sum()
	switch i {
	case 1:
		clrScrn()
		options := []string{"Common-goods Supplier", "Trade Supplier", "Morally Neutral Supplier", "Black Market Supplier"}
		prf := profile.NewUWP(sourceWorld.UWP())
		tl := TrvCore.EhexToDigit(prf.TL().String())
		if tl >= 8 {
			options = append(options, "Finding an online supplier")
		}
		sup, _ := menu("Select Supplier Type:", options...)
		switch sup {
		case 0:
			//fmt.Println("Enter Effect of a Broker (6), EDU or SOC, 1–6 days:")
			fmt.Println("Search might take 1-6 days...")
			input := userInputIntSlice("Enter Effect of Broker (6), EDU or SOC check (limit in days after '/' if nesessary): ")
			eff, timeLimit = getEffTime(input)
		case 1:
			//fmt.Println("Enter Effect of a Broker (6), EDU or SOC, 1–6 days:")
			fmt.Println("Search might take 1-6 days...")
			input := userInputIntSlice("Enter Effect of Broker (8), EDU or SOC check (limit in days after '/' if nesessary): ")
			eff, timeLimit = getEffTime(input)
		case 2:
			//fmt.Println("Enter Effect of a Broker (6), EDU or SOC, 1–6 days:")
			fmt.Println("Search might take 2-12 days...")
			input := userInputIntSlice("Enter Effect of Streetwise or Investigate (10), EDU or SOC check (limit in days after '/' if nesessary): ")
			eff, timeLimit = getEffTime(input)
			time = dice.Roll("1d6").Sum() * 2
		case 3:
			//fmt.Println("Enter Effect of a Broker (6), EDU or SOC, 1–6 days:")
			fmt.Println("Search might take 1-6 days and WILL CAUSE negaive events if failed.")
			input := userInputIntSlice("Enter Effect of Streetwise (8), EDU or SOC check (limit in days after '/' if nesessary): ")
			eff, timeLimit = getEffTime(input)
		case 4:
			//fmt.Println("Enter Effect of a Broker (6), EDU or SOC, 1–6 days:")
			fmt.Println("Search will take few hours...")
			input := userInputInt("Enter Effect of Computers (8), EDU check: ")
			eff = input

		}
		switch prf.Starport().String() {
		case "A":
			eff += 6
		case "B":
			eff += 4
		case "C":
			eff += 2
		}
		eff, time, abort := mutateTestResultsByTime(eff, time, timeLimit)

		if abort {
			fmt.Println("Search aborted after", time, "days...")
		} else {
			fmt.Println("Search took", time, "days...")
		}

		menu("--------------------------------------------------------------------------------", "Continue")

		if eff >= 0 {
			newLocalSupplier(sup)
			menuPosition = "MARKET"
			return
		}

		r := dice.Roll("2d6").Sum()
		if r-eff > prf.Laws().Value() {
			fmt.Println(r, eff, prf.Laws().Value())
			menu("TODO: NEGATIVE EVENT Report To Referee", "Next")
		}

	case 2:
		fmt.Println("Purchase:")
		if len(localMarket) == 0 {
			clrScrn()
			menu("Local Supplier not Found", "Return")
			menuPosition = "MARKET"
			return
		}
		purchase()
	case 3:
		fmt.Println("Sale:")
		if len(localMarket) == 0 {
			clrScrn()
			menu("Local Supplier not Found", "Return")
			menuPosition = "MARKET"
			return
		}
		sellTradeGoods()
		menuPosition = "MARKET"
	}
}

func getEffTime(input []int) (int, int) {
	eff := 0
	timeLim := 999
	if len(input) > 0 {
		eff = input[0]
	}
	if len(input) > 1 {
		timeLim = input[1]
	}
	return eff, timeLim
}

func depart() {
	fmt.Println("Preaparing for Depart:")
	timeSpent := userInputInt("Enter number of days spent on the planet: ")
	advanceTime(timeSpent)
	sp, _ := starport.From(sourceWorld)
	total := 0
	bf, err := sp.BerthingFee(day - startDay)
	reportErr(err)
	total += bf
	fmt.Println(bf, "Cr was charged for Berthing Servises")
	lsPayment := countHeads() * 250
	total += lsPayment
	fmt.Println(lsPayment, "Cr was charged for Life Support Supplies")
	prf := profile.NewUWP(sourceWorld.UWP())
	//reportErr(err2)
	popDM := prf.Pops().Value()
	lawsDM := prf.Laws().Value()

	calculateETA()

	r := dice.Roll("2d6").DM(popDM).Sum()
	if r < lawsDM {
		fmt.Println("------------------------------")
		fmt.Println("Departure delayed for", (r-lawsDM)*-1, "hours")
		fmt.Println("------------------------------")
	}
	if userConfirm("Exit") {

	}
	tm := time.Now().Format("2006102150405")
	reportErr(copyFile(cargofile, "CargoManifestArhive"+tm+".txt"))
	reportErr(copyFile(passengerfile, "PassengerManifestArhive"+tm+".txt"))

	os.Exit(1)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func calculateETA() {
	cm := loadCargoManifest()
	for i, val := range cm.entry {
		if val.GetETA() == "NOETA" {
			etaDay := day + len(jumpRoute)*8 + dice.Flux()
			newRawDay := etaDay + 365*year
			//newEta := formatDate(etaDay, year)
			cm.entry[i].SetETA(integerToEhexCode(newRawDay))
		}
	}
	saveCargoManifest(cm)
}
