package routine

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/devtools/cli/features"
	"github.com/Galdoba/devtools/cli/user"
)

const (
	typingDelay = "15ms"
)

var delay time.Duration
var emmersiveMode bool
var sourceWorld world.World
var targetWorld world.World
var w int
var currentDate string
var dp *dice.Dicepool
var pFactor int

func init() {
	del, err := time.ParseDuration(typingDelay)
	if err != nil {
		fmt.Println(err.Error())
	}
	delay = del
	//emmersiveMode = true
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ = termbox.Size()
	defer termbox.Close()

}

func StartRoutine() {
	helloWorld()
	printSlow("TAS information terminal greets you, Traveller!\n")
	printSlow("Input current date: \n")
	currentDate = userInputStr()
	dp = dice.New(utils.SeedFromString(currentDate))
	printSlow("Select your current world: \n")
	sourceWorld = pickWorld()
	printSlow("Select your destination world: \n")
	targetWorld = pickWorld()
	pFactor = passengerTrafficValue(sourceWorld, targetWorld)
	clrScrn()
	PassengerRoutine()

}

func selectOperation() {

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
		output := "Data retrived: " + w.Name() + " (" + w.UWP() + ")\n"
		printSlow(output)
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
	printSlow("Clearence granted\n")
}

func printHead() {
	fmt.Println("         Date: ", currentDate)
	fmt.Println("Current World: ", sourceWorld.Hex()+" - "+sourceWorld.Name()+" ("+sourceWorld.UWP()+") "+sourceWorld.TradeCodesString()+" "+sourceWorld.TravelZone())
	fmt.Println("  Destination: ", targetWorld.Hex()+" - "+targetWorld.Name()+" ("+targetWorld.UWP()+") "+targetWorld.TradeCodesString()+" "+targetWorld.TravelZone())
	fmt.Println("Passenger Traffic Value:", pFactor)
	for i := 0; i < w; i++ {
		fmt.Print("-")
	}

}
