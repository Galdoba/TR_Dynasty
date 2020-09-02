package routine

import (
	"fmt"
	"strconv"
)

func MailRoutine() {
	clrScrn()
	mailDice := dp.RollNext("2d6").DM(mailDM()).Sum()
	if mailDice >= 12 {
		qty := dp.RollNext("1d6").Sum()
		printSlow("Mail hauling offer:\n")
		printSlow(strconv.Itoa(qty*5) + " tons of Universal mail containers are ready to pick up.\n")
		printSlow("Hauling fee: " + strconv.Itoa(qty*25000) + " Cr\n")
		fmt.Println("-----------------------------------------------------")
	} else {
		printSlow("No mail available\n")
		fmt.Println("-----------------------------------------------------")
	}

}

func mailDM() int {
	mailDM := 0
	if ftValue < -9 {
		mailDM = -2
	}
	if ftValue < -5 {
		mailDM = -1
	}
	if ftValue < 4 {
		mailDM = 0
	}
	if ftValue < 9 {
		mailDM = 1
	}
	if ftValue > 9 {
		mailDM = 2
	}
	if shipArmed() {
		mailDM += 2
	}
	mailDM += getCrewNavyScoutMerchantRank()
	mailDM += getCrewSOCdm()
	mailDM += techDifferenceDM()
	return mailDM
}
