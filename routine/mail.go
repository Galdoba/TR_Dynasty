package routine

import (
	"fmt"
	"strconv"
)

var mailOffer string

func MailRoutine() {
	fmt.Println("Mail hauling offer:")
	if mailOffer != "" {
		fmt.Println(mailOffer)
		fmt.Println("-----------------------------------------------------")
		return
	}
	mailLot := newCargoLot()
	qty := 0
	fee := 0
	mailDice := dp.RollNext("2d6").DM(mailDM()).Sum()
	if mailDice >= 12 {
		qty = dp.RollNext("1d6").Sum()
		fee = qty * 25000
		if autoMod {
			fee -= localBroker.CutFrom(fee)
		}
		mailOffer = strconv.Itoa(qty*5) + " tons of Universal mail containers are ready to pick up.\n" + "Hauling fee: " + strconv.Itoa(fee) + " Cr"

	} else {
		mailOffer = "No mail available"
	}
	fmt.Println(mailOffer)
	fmt.Println("-----------------------------------------------------")
	if qty != 0 && fee != 0 {
		mailLot.FillMailData(qty, fee)
		portCargo = append(portCargo, mailLot)
	}
}

func (mailLot *cargoLot) FillMailData(qty, fee int) {
	mailLot.SetETA("NOETA")
	mailLot.SetTGCode("MAIL")
	mailLot.SetDescr("Universal mail containers")
	mailLot.SetVolume(qty * 5)
	mailLot.SetDestination(targetWorld.Hex())
	mailLot.SetComment("Transport fee " + strconv.Itoa(fee))
	mailLot.SetCost(fee)
	mailLot.SetInsurance(100)
	mailLot.SetDangerDM(-1 * qty)
	mailLot.SetOrigin(sourceWorld.Hex())
	mailLot.SetLegality(true)
	mailLot.SetRiskDM(0)
	mailLot.SetSupplierType("Goverment")
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
