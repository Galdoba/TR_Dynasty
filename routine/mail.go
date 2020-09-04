package routine

import (
	"fmt"
	"strconv"
)

func MailRoutine() {
	mailDice := dp.RollNext("2d6").DM(mailDM()).Sum()
	if mailDice >= 12 {
		qty := dp.RollNext("1d6").Sum()
		printSlow("Mail hauling offer:\n")
		printSlow(strconv.Itoa(qty*5) + " tons of Universal mail containers are ready to pick up.\n")
		fee := 0
		if autoMod {
			fee = qty*25000 - localBroker.CutFrom(qty*25000)
		} else {
			fee = qty * 25000
		}
		printSlow("Hauling fee: " + strconv.Itoa(fee) + " Cr\n")
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
