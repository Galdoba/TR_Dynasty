package main

import (
	"fmt"
	"os"

	. "github.com/Galdoba/TR_Dynasty/dateManager"

	"github.com/Galdoba/utils"
)

var programError error

func errorDetected() bool {
	if programError != nil {
		return true
	}
	return false
}

type upcomingDate struct {
	actionDate    string
	genUpkeepDate string
	fiveYearEvent string
	decadeEvent   string
}

var listTRAITS []string
var listCHARS []string
var listAPTITUDES []string
var listVALUES []string

func main() {
	//trvRoll("2D")

	seed := utils.RandomSeed()
	listTRAITS = traitsLIST()
	listCHARS = characteristicsLIST()
	listAPTITUDES = aptitudeLIST()
	listVALUES = valuesLIST()
	if len(listAPTITUDES) == len(listCHARS) && len(listTRAITS) == len(listVALUES) {
		os.Exit(1)
	}

	seeAge()

	gameDate := NewImperialDate(1105, 52, 23, 59, 57)
	for i := 0; i < 20; i++ {
		gameDate.PassTime("Second", 5)
		fmt.Println(gameDate.ToString())
		fmt.Println(trvRoll("4D+3"))
	}
	os.Exit(4)
	var dyn *Dynasty
	for dyn == nil {
		dyn = NewDynasty()
	}

	fmt.Println("Seed =", seed)
	fmt.Println(dyn)
	fmt.Println(dyn.toString())

	fmt.Println("Gather resources:", dyn.aptitideCheck(apttAcquisition, charCleverness, difficultyEasy))
	fmt.Println("Gather resources:", dyn.aptitideCheck(apttAcquisition, charCleverness, difficultyEasy))
	fmt.Println(dyn.probeAptitideCheck(apttAcquisition, charCleverness, difficultyEasy))

}
