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
	typingDelay = "10ms"
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

func init() {
	printSlow("Initialisation...\n")
	del, err := time.ParseDuration(typingDelay)
	if err != nil {
		fmt.Println(err.Error())
	}
	delay = del
	emmersiveMode = true
	freightBase = 500
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
	valid := false
	input := "000-0000"
	for !valid {
		input = userInputStr("Enter current Imperial Date (format: ddd-yyyy): ")
		data := strings.Split(input, "-")
		if len(data) != 2 {
			printSlow("WARNING: Unknown format '" + input + "'\n")
			continue
		}
		for i := range data {
			test, err := strconv.Atoi(data[i])
			if err != nil {
				printSlow("WARNING: " + err.Error() + "\n")
			}
			switch i {
			case 0:
				if test < 100 {
					input = "0" + input
				}
				if test < 10 {
					input = "0" + input
				}
				day = test
			case 1:
				year = test
			}
		}
		valid = true
	}
	return input
}

func printOptions() {
	printSlow("Selelect operation: \n")
	printSlow("[0] - Disconnect \n")
	printSlow("[1] - Search Passengers \n")
	printSlow("[2] - Search Freight \n")
	printSlow("[3] - Search Mail \n")
}

func selectOperation() {
	for {
		input := userInputStr("Initiate ")
		switch input {
		default:
			printSlow("Sorry, command '" + input + "' unrecognised\n")
		case "0":
			printSlow("Have a nice day!")
			os.Exit(0)
		case "1":
			PassengerRoutine()
		case "2":
			FreightRoutine()
		case "3":
			MailRoutine()

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
	printSlow("Clearance granted!\n")
}

func printHead() {
	fmt.Println("         Date: ", currentDate)
	fmt.Println("Current World: ", sourceWorld.Hex()+" - "+sourceWorld.Name()+" ("+sourceWorld.UWP()+") "+sourceWorld.TradeCodesString()+" "+sourceWorld.TravelZone())
	fmt.Println("  Destination: ", targetWorld.Hex()+" - "+targetWorld.Name()+" ("+targetWorld.UWP()+") "+targetWorld.TradeCodesString()+" "+targetWorld.TravelZone())
	fmt.Println("          ETA: ", formatDate(day+(len(jumpRoute)*7), year))
	fmt.Println("Passenger Traffic Value:", ptValue)
	fmt.Println("  Freight Traffic Value:", ftValue)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Expected Jump Sequance: ", jumpRoute)
	fmt.Println("        Total Distance: ", distance)
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
