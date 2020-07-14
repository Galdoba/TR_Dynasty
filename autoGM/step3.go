package autoGM

import (
	"fmt"

	"github.com/Galdoba/utils"
)

func step3JumpTravel() {
	jumpTimeTotal := utils.RollDice("6d6", 148)
	pasengers := utils.InputInt("Set number of Passengers: ")
	encounterDice := ""
	encounterTN := 9
	if utils.InRange(pasengers, 0, 10) {
		encounterDice = "2d6"
		encounterTN = 12
	}
	if utils.InRange(pasengers, 11, 50) {
		encounterDice = "d6"
		encounterTN = 6
	}
	if utils.InRange(pasengers, 51, 250) {
		encounterDice = "2d6"
		encounterTN = 12
	}
	if utils.InRange(pasengers, 251, 1000000) {
		encounterDice = "d6"
		encounterTN = 6
	}
	//encounter := false
	if !encounterHappens(encounterDice, encounterTN) {
		fmt.Println("Nothing notable happens in this jump. That was hell of a long week...")
		fmt.Println("After", jumpTimeTotal/24, "days and", jumpTimeTotal%24, "hours in hyperjump, you emerge to normal space.")
		return
	}
	eventType := ""
	switch rolld6() {
	case 1, 2, 3, 4:
		eventType = "Human"
	case 5, 6:
		eventType = "Technical"

	}
	fmt.Println("EventType:", eventType)
}
