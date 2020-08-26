package trademgt2

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
)

var umcAvailable int
var umcPrice int

func MailProcedure() {
	/*
		2. Determine if Mail is Present.
		3. Determine time table.
		4. Negotiate hauling fees.
		5. Travel to destination world.
		6. Collect fees.
	*/

	mailDM := mailDMfromFTV()
	mailDM += 7
	switch sourceWorld.PlanetaryData("Tech") {
	case "0", "1", "2", "3", "4", "5":
		mailDM -= 5
	}
	mailDR := dice.Roll("2d6").DM(mailDM).Sum()
	//fmt.Println("Searching Mail Contracts...")
	//mailDR += userInputInt("Enter Diplomat(8) ir Investigate(8) check effect: ")
	r := dice.Roll("2d6").DM(-8).Sum()
	//fmt.Println("Auto Roll:", r)
	mailDR += r
	// if mailDR < 12 {
	// 	fmt.Println("No Mail Avalable")
	// 	return
	// }
	umcAvailable = mailDR - 11
	tn := brokerPersuadeDiff(ftv) + 2
	negRoll := dice.Roll("2d6").DM(-tn).Sum()
	//fmt.Println("Auto Roll with TN=", tn, "|", negRoll)
	umcPrice = mailNegotiationCheckEffect(negRoll)
	informAboutMail()

}

func mailDMfromFTV() int {
	if ftv <= -10 {
		return -2
	}
	if ftv <= -5 {
		return -1
	}
	if ftv <= 4 {
		return 0
	}
	if ftv <= 9 {
		return 1
	}
	return 10
}

func mailNegotiationCheckEffect(eff int) int {
	if eff <= -5 {
		return 5000
	}
	if eff <= -3 {
		return 10000
	}
	if eff <= -1 {
		return 15000
	}
	if eff == 0 {
		return 20000
	}
	if eff <= 2 {
		return 25000
	}
	if eff <= 4 {
		return 30000
	}
	return 40000
}

func informAboutMail() {
	defer fmt.Println("------------------------------------------------------")
	fmt.Println("There are", umcAvailable, "UMCs available for transport.")
	//fmt.Println("Transport fee is", umcPrice, "for each container.")
}
