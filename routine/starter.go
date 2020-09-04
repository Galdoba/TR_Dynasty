package routine

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/devtools/cli/features"
	"github.com/Galdoba/devtools/cli/user"
)

const (
	typingDelay = "4ms"
)

var delay time.Duration
var emmersiveMode bool
var sourceWorld world.World
var targetWorld world.World
var distance int
var currentDate string
var dp *dice.Dicepool
var ptValue int
var ftValue int
var jumpRoute []int
var day int
var year int
var autoMod bool
var gmMode bool

func init() {
	printSlow("Initialisation...\n")
	del, err := time.ParseDuration(typingDelay)
	if err != nil {
		fmt.Println(err.Error())
	}
	delay = del
	emmersiveMode = true
	gmMode = true
	freightBase = 500
	localBroker = broker{0, 0.0}
}

func StartRoutine() {

	clrScrn()
	printSlow("Start...\n")
	helloWorld()
	printSlow("TAS information terminal greets you, Traveller!\n")
	printSlow("Gathering data...\n")
	//printSlow("Input current date: \n")
	//currentDate = userInputStr()
	currentDate = userInputDate()
	clrScrn()
	dp = dice.New(utils.SeedFromString(currentDate))
	printSlow("Select your current world: \n")
	sourceWorld = pickWorld()
	clrScrn()
	printSlow("Select your destination world: \n")
	targetWorld = pickWorld()

	distance = Astrogation.JumpDistance(sourceWorld.Hex(), targetWorld.Hex())
	ptValue = passengerTrafficValue(sourceWorld, targetWorld)
	ftValue = freightTrafficValue(sourceWorld, targetWorld)
	clrScrn()
	jumpRoute = []int{distance}
	if distance > 2 {
		routeValid := false
		for !routeValid {
			jumpRouteTest, err := userInputJumpRoute()
			if err != nil {
				printSlow(err.Error())
				continue
			}
			jumpRoute = jumpRouteTest
			routeValid = true
		}

	}
	clrScrn()
	printOptions()
	selectOperation()
}

func userInputDate() string {
	input := 0
	output := ""
	input = userInputInt("Enter current Imperial Date (day only): ")

	year = getYear()
	day = input
	for day > 365 {
		year++
		day = day - 365
	}
	for day < 1 {
		year--
		day = day + 365
	}

	if day < 100 {
		output = "0" + output
	}
	if day < 10 {
		output = "0" + output
	}
	output = output + strconv.Itoa(day) + "-" + strconv.Itoa(year)
	return output
}

func printOptions() {
	printSlow("Select operation: \n")
	printSlow(" [0] - Disconnect \n")
	printSlow(" [1] - Hire Local Broker\n")
	printSlow(" [2] - Let Local Broker do all the work\n")
	printSlow(" [4] - Search Passengers\n")
	printSlow(" [5] - Search Freight \n")
	printSlow(" [6] - Search Mail \n")
	printSlow(" [7] - Search ALL \n")
}

func selectOperation() {
	for {
		input := userInputStr("Initiate ")
		autoMod = false
		switch input {
		default:
			printSlow("Sorry, command '" + input + "' unrecognised\n")
		case "0":
			printSlow("Have a nice day!")
			os.Exit(0)
		case "1":
			clrScrn()
			chooseBroker()
			clrScrn()
		case "4":
			clrScrn()
			PassengerRoutine()
		case "5":
			clrScrn()
			FreightRoutine()
		case "6":
			clrScrn()
			MailRoutine()
		case "7":
			clrScrn()
			PassengerRoutine()
			FreightRoutine()
			MailRoutine()
		case "2":
			clrScrn()
			autoMod = true
			if localBroker.cut == 0 {
				clrScrn()
				printSlow("Local Broker is not hired...\n")
			} else {
				PassengerRoutine()
				FreightRoutine()
				MailRoutine()
			}

		}
		printOptions()
	}
}

/*
Automatic Campaign Flowc hart
1. Job Hunting (Planetside Events, page 7)
2. Preparations (repeat previous step)
3. Jump Travel (Onboard Events, page 60)
4. Space Travel
a. Space Events (page 32)
b. Life Events (page 67)
5. Ground Travel (Planetside Events, page 7)
6. Destination (Any)
7. Return (repeat steps 3,4 and 5 in reverse order)
8. Resting
a. Planetside, page 7
b. Life events, page 67
c. Adventure Hooks, page 71

*/

func printSlow(text string) {
	if emmersiveMode {
		features.TypingSlowly(text, delay)
	} else {
		fmt.Print(text)
	}
}

func userInputStr(msg ...string) string {
	for i := range msg {
		printSlow(msg[i])
	}
	str, err := user.InputStr()
	if err != nil {
		printSlow(err.Error())
		return err.Error()
	}
	return str
}

func userInputInt(msg ...string) int {
	str := userInputStr(msg...)
	i, err := strconv.Atoi(str)
	for err != nil {
		printSlow(err.Error())
		str = userInputStr(msg...)
		i, err = strconv.Atoi(str)
	}
	return i
}

func pickWorld() world.World {
	dataFound := false
	for !dataFound {
		input := userInputStr("Enter world's Name, Hex or UWP: ")
		otuData, err := otu.GetDataOn(input)
		if err != nil {
			printSlow("WARNING: " + err.Error() + "\n")
			continue
		}
		w, err := world.FromOTUdata(otuData.Info)
		if err != nil {
			printSlow(err.Error() + "\n")
			continue
		}
		//output := "Data retrived: " + w.Name() + " (" + w.UWP() + ")\n"
		//printSlow(output)
		return w

	}
	fmt.Println("This must not happen!")
	return world.World{}
}

func loadWorld(key string) (world.World, error) {
	otuData, err := otu.GetDataOn(key)
	if err != nil {
		return world.World{}, err
	}
	w, err := world.FromOTUdata(otuData.Info)
	if err != nil {
		return world.World{}, err
	}
	return w, nil
}

func userInputJumpRoute() ([]int, error) {
	route := userInputStr("Enter route sequence (format: 'XXYY XXYY ... XXYY'): ")
	var routeSl []int
	jumpPoints := strings.Split(route, " ")
	for i := 1; i < len(jumpPoints); i++ {
		locDist := Astrogation.JumpDistance(jumpPoints[i], jumpPoints[i-1])
		if locDist > getJumpDrive() {
			fmt.Println(routeSl)
			return routeSl, errors.New("Jump route invalid: Distance > JumpDrive")
		}
		routeSl = append(routeSl, locDist)
	}
	return routeSl, nil
}

func techDifferenceDM() int {
	tl1 := TrvCore.EhexToDigit(sourceWorld.PlanetaryData(constant.PrTL))
	tl2 := TrvCore.EhexToDigit(targetWorld.PlanetaryData(constant.PrTL))
	tlDiff := utils.Max(tl1, tl2) - utils.Min(tl1, tl2)
	if tlDiff > 5 {
		tlDiff = 5
	}
	return -tlDiff
}

func clrScrn() {
	var clear map[string]func()
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
	printHead()
}

func helloWorld() {
	printSlow("   LOGIN: ***********\n")
	printSlow("PASSWORD: *************\n")
	//printSlow("Clearance granted!\n")

}

func printHead() {
	fmt.Println("         Date: ", formatDate(day, year))
	fmt.Println("Current World: ", sourceWorld.Hex()+" - "+sourceWorld.Name()+" ("+sourceWorld.UWP()+") "+sourceWorld.TradeCodesString()+" "+sourceWorld.TravelZone())
	fmt.Println("  Destination: ", targetWorld.Hex()+" - "+targetWorld.Name()+" ("+targetWorld.UWP()+") "+targetWorld.TradeCodesString()+" "+targetWorld.TravelZone())
	fmt.Println("Passenger Traffic Value:", ptValue)
	fmt.Println("  Freight Traffic Value:", ftValue)
	fmt.Println("     Local Broker's Cut:", localBroker.cut, "%")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Expected Jump Sequence: ", jumpRoute)
	fmt.Println("        Total Distance: ", distance)
	fmt.Println("                   ETA: ", formatDate(day+(len(jumpRoute)*7), year))
	fmt.Println("-----------------------------------------------------")
}

func formatDate(day, year int) string {
	date := ""
	if day < 100 {
		date = "0" + date
	}
	if day < 10 {
		date = "0" + date
	}
	if day > 365 {
		day = day - 365
		year++
	}
	date += strconv.Itoa(day)
	date += "-"
	date += strconv.Itoa(year)
	return date
}

func spendTime() {
	if localBroker.cut == 0 {
		dTook := dice.Roll("1d6").Sum()
		printSlow("This operation took " + strconv.Itoa(dTook) + " days...\n")
	} else {
		printSlow("This operation took few hours...\n")
	}
}
