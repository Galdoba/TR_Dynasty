package autoGM

import (
	"fmt"

	"github.com/Galdoba/convert"

	"github.com/Galdoba/utils"

	DateManager "github.com/Galdoba/TR_Dynasty/dateManager"
)

//	dateManager "github.com/Galdoba/TR_Dynasty/datemanager"

var curentDate *DateManager.ImperialDate

func printOption(optName string, optStatus bool, optNum int) {
	status := " "
	if optStatus == true {
		status = "X"
	}
	num := convert.ItoS(optNum)
	if utils.InRange(optNum, 0, 9) {
		num = " " + num
	}
	fmt.Println(num + " [" + status + "] -- " + optName)
}

func printAllOptions(optSlice []string, optStatuses []bool) {
	for i := range optSlice {
		if i == 0 {
			fmt.Println(optSlice[0])
		} else {

			printOption(optSlice[i], optStatuses[i], i)
		}
	}
}

func SelectionOptionsMult(descr string, opt ...string) ([]string, []bool) {
	//сбор данных
	var optSlice []string
	var optStatuses []bool
	optSlice = append(optSlice, descr)
	optStatuses = append(optStatuses, false)
	for i := range opt {
		optSlice = append(optSlice, opt[i])
		optStatuses = append(optStatuses, false)
	}
	optSlice = append(optSlice, "[DONE]")
	optStatuses = append(optStatuses, false)
	printAllOptions(optSlice, optStatuses)
	done := false
	for !done {
		//pick := utils.InputInt()
		var pick int
		fmt.Scan(&pick)
		if utils.InRange(pick, 1, len(optSlice)-1) {
			fmt.Println("\033[0A"+convert.ItoS(pick), "toggled   ") //, optSlice[pick]) //убрать текст опции
			optStatuses[pick] = !optStatuses[pick]
		}
		fmt.Print("\033[" + convert.ItoS(len(optStatuses)+1) + "A")
		printAllOptions(optSlice, optStatuses)
		if pick == len(optSlice)-1 {
			done = true
		}
	}
	//анализ и возврат
	fmt.Println("\n")
	var returnSlc []string
	var resultSlc []bool
	for i := range optSlice {
		if i == 0 || i == len(optSlice)-1 {
			continue
		}
		returnSlc = append(returnSlc, optSlice[i])
		resultSlc = append(resultSlc, optStatuses[i])

	}

	return returnSlc, resultSlc

}

func AutoGM2() {
	test()
}

func AutoGM() {
	// //startDate := DateManager.InputDate()
	// SelectionOptionsMult("Circumstance Modifiers",
	// 	"Item is considered to be highly specialised       –1",
	// 	"Item is typically reserved for military use       –2",
	// 	"Item’s TL is 3–4 steps away from World’s TL       –1",
	// 	"Item’s TL is 5 or more steps away from World’s TL –2",
	// 	"Purchaser willing to pay double listed cost       +1",
	// 	"Purchaser willing to pay triple listed cost       +2",
	// 	"Starport Class A or B                             +1",
	// 	"Starport Class X                                  –2",
	// 	"World has Hi, Ht, In and/or Ri Trade Codes        +1",
	// 	"World has Lt, Na, NI and/or Po Trade Codes        –1",
	// 	"population magnitude 4                            –3",
	// 	"population magnitude 5                            –2",
	// 	"population magnitude 6–7                          –1",
	// 	"population magnitude 8                            +1",
	// 	"population magnitude 9+                           +2",
	// )
	// masterList := MasterMap()
	// testMasterList(masterList)
	// return
	seed := utils.RandomSeed()
	fmt.Println(seed)
	curentDate = DateManager.NewImperialDate(1105001000000)
	fmt.Println(curentDate.TimeStamp())
	end := false
	for !end {
		switch selectionMenu() {
		case 1:
			step1JobHunting()
		case 2:
			step2Complication()
		case 3:
			step3JumpTravel()
		case 4:
			step4SpaceTravel()
		case -1:
			end = true
		}
	}

}

func selectionMenu() int {
	fmt.Println("")
	pick, opt := utils.TakeOptions("Select Scene Type", "Job Hunting", "Complication", "Jump Travel", "Space Travel", "Ground Travel", "Destination", "Resting", "END PROGRAM")
	switch pick {
	default:
		fmt.Println("TODO: ", opt)
	case 1:
		return pick
	case 2:
		fmt.Println(complication())
	case 3:
		return pick
	case 4:
		return pick
	case 8:
		return -1
	}
	return 0
}

func rolld6() int {
	return utils.RollDice("d6")
}

func roll2d6() int {
	return utils.RollDice("2d6")
}

func rollD66() int {
	d1 := utils.RollDice("d6")
	d2 := utils.RollDice("d6")
	dStr := convert.ItoS(d1) + convert.ItoS(d2)
	return convert.StoI(dStr)
}

func askYesNo(str string) bool {
	gotAnswer := false
	for !gotAnswer {
		fmt.Print(str + "(y/n) ")
		answer := utils.InputString()
		switch answer {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Error: Answer is incorrect. (Type 'y' or 'n')")
		}
	}
	return false
}

func encounterHappens(encounterDice string, tn int) bool {
	if utils.RollDice(encounterDice) >= tn {
		return true
	}
	return false
}

func rollEncounterTable(table string) string {
	switch table {
	default:
		return "TODO: " + table
	case "Local Events":
		return cityEventLocal()
	case "Global Events":
		return eventGlobal()

	}
}

// func cause() string {
// 	cause := ""
// 	fmt.Println("Cause on p.132")
// 	switch rollD66() {
// 	case 11:
// 		cause = "Stop all abuse of animals including experimentation, use in food industry, hunting and even pet ownership."
// 	case 12:
// 		cause = "Equal distribution of wealth between all sophonts on the planet, including illegal aliens and uplifted animals."
// 	case 13:
// 		cause = "Extermination of a minority considered racially inferior."
// 	case 14:
// 		cause = "Liberation of an occupied territory, possibly hundreds of parsecs away."
// 	case 15:
// 		cause = "Roll twice and keep both results."
// 	case 16:
// 		cause = "Roll again. The group is radically opposed to this cause."
// 	case 21:
// 		cause = "Banning a product considered illegal or immoral by the movement (drinking, sugar, the colour blue)."
// 	case 22:
// 		cause = "Banning an activity considered illegal or immoral by the movement (public displays of affection, spaceflight)."
// 	case 23:
// 		cause = "Absolute cosmic anarchy."
// 	case 24:
// 		cause = "Investigating corporate corruption."
// 	case 25:
// 		cause = "Roll twice and keep both results."
// 	case 26:
// 		cause = "Roll again. The group is radically opposed to this cause."
// 	case 31:
// 		cause = "Promotion of an immigrant religion incompatible with local laws."
// 	case 32:
// 		cause = "Making an allegorical statement through a series of high profile bombings."
// 	case 33:
// 		cause = "Toppling the local feudal lord."
// 	case 34:
// 		cause = "Toppling the local corporate manager."
// 	case 35:
// 		cause = "Roll twice and keep both results."
// 	case 36:
// 		cause = "Roll again. The group is radically opposed to this cause."
// 	case 41:
// 		cause = "Toppling the Imperium and replacing it with an alliance were all races, sexes, religions and classes are equal."
// 	case 42:
// 		cause = "Exterminating the working class and replacing it with robots."
// 	case 43:
// 		cause = "Uplifting all animals."
// 	case 44:
// 		cause = "Stopping an ongoing genocide deep inside Aslan space."
// 	case 45:
// 		cause = "Roll twice and keep both results."
// 	case 46:
// 		cause = "Roll again. The group is radically opposed to this cause."
// 	case 51:
// 		cause = "Assisting an alien species to take over the Imperium because this species is more advanced and moral than humans."
// 	case 52:
// 		cause = "Voting rights for illegal aliens who infiltrated the nation more than a century ago and were neither expelled nor given legal rights."
// 	case 53:
// 		cause = "Expulsion of the previous group."
// 	case 54:
// 		cause = "Stripping women of their voting and employment rights and returning them to the bedroom and kitchen."
// 	case 55:
// 		cause = "Roll twice and keep both results."
// 	case 56:
// 		cause = "Roll again. The group is radically opposed to this cause."
// 	case 61:
// 		cause = "Banning sexual reproduction because it fills the world with imperfect people who commit crimes."
// 	case 62:
// 		cause = "Legalisation of all banned substances, including the mutation-inducing Meshuginan extract."
// 	case 63:
// 		cause = "Re-conquering a recently liberated nearby nation and enslaving its racially inferior citizens."
// 	case 64:
// 		cause = "Banning all FTL travel and AI and killing everyone not native to this planet or with cybernetic implants."
// 	case 65:
// 		cause = "Roll twice and keep both results."
// 	case 66:
// 		cause = "Roll thrice and keep all results."
// 	}
// 	return cause
// }

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
// DELTA HERCULIS (SUBSECTOR A)
// Hex	Name	UWP	Bases	Trade	L/C/B	Temp	GG	Stars	Area
// 0103	NSSC 0103	X100000-0		Ba Va	0/0/X	Temp	G	K2V	Un
// 0104	Palmer's Landing (NSSC 0104)	X310000-0	Uh	Ba	0/0/+3	Cold	G	K2V M4V Un
// 0105	NSSC 0105	X84A000-0	Ba Wa B/0/X Temp G K8V	Un

// func MasterMap() map[string][]string {
// 	masterList := make(map[string][]string)
// 	//          Hex                Name         UWP       Bases Trade     L/C/B    Temp    GG  Stars   Area
// 	//DELTA HERCULIS (SUBSECTOR A)
// 	masterList["0103"] = []string{"NSSC 0103", "X100000-0", "", "Ba Va", "0/0/X", "Temp", "G", "K2V", "Un"}
// 	masterList["0104"] = []string{"Palmer's Landing (NSSC 0104)", "X310000-0", "Uh", "Ba", "0/0/+3", "Cold", "G", "K2V M4V", "Un"}
// 	masterList["0105"] = []string{"NSSC 0105", "X84A000-0", "", "Ba Wa", "B/0/X", "Temp", "G", "K8V", "Un"}
// 	masterList["0108"] = []string{"Vodyanoy (NSSC 0108)", "X636000-0", "Uh", "Ba", "0/0/-3", "Frozen", "G", "BD", "Un"}
// 	masterList["0109"] = []string{"NSSC 0109", "X413000-0", "", "Ba Ic", "1/0/+1", "Temp", "G", "M2V M4V", "Un"}
// 	masterList["0110"] = []string{"Alpha Serpentis", "X410000-0", "Um", "Ba", "0/0/+1", "Temp", "G", "K2III", "Un"}
// 	masterList["0203"] = []string{"NSSC 0203", "X200000-0", "", "Ba Va", "0/0/X", "Temp", "G", "M1V", "Un"}
// 	masterList["0204"] = []string{"NSSC 0204", "XAA8000-0", "", "Ba Fl", "6/0/+2", "Temp", "G", "M6V M9V", "Un"}
// 	masterList["0205"] = []string{"NSSC 0205", "X300000-0", "", "Ba Va", "0/0/X", "Frozen", "G", "M6V M9V", "Un"}
// 	masterList["0209"] = []string{"NSSC 0209", "X554000-0", "", "Ba Ga", "4/1/-1", "Cold", "G", "K3V M1V", "Un"}
// 	masterList["0301"] = []string{"NSSC 0301", "X310000-0", "", "Ba", "0/0/X", "Temp", "G", "M0V", "Un"}
// 	masterList["0302"] = []string{"NSSC 0302", "X100000-0", "", "Ba Va", "0/0/+2", "Temp", "G", "K4V M2V", "Un"}
// 	masterList["0304"] = []string{"NSSC 0304", "X100000-0", "", "Ba Va", "0/0/X", "Hot", "", "M0V M4V", "Un"}
// 	masterList["0305"] = []string{"NSSC 0305", "X668000-0", "", "Ba Ga", "8/7/-1", "Temp", "G", "K9V", "Un"}
// 	masterList["0310"] = []string{"NSSC 0310", "X770000-0", "", "Ba De", "0/0/+0", "Hot", "G", "K1V", "Un"}
// 	masterList["0401"] = []string{"NSSC 0401", "X200000-0", "", "Ba Va", "0/0/X", "Temp", "G", "M0V", "Un"}
// 	masterList["0402"] = []string{"NSSC 0402", "X738000-0", "", "Ba", "1/0/+1", "Frozen", "G", "M0V", "Un"}
// 	masterList["0403"] = []string{"GL632", "X883000-0", "Ua", "Ba", "7/4/X", "Hot", "G", "G0V", "Un"}
// 	masterList["0404"] = []string{"NSSC 0404", "X682000-0", "", "Ba", "8/0/+1", "Temp", "G", "K8V M7V", "Un"}
// 	masterList["0405"] = []string{"NSSC 0405", "X6A0000-0", "", "Ba De", "2/1/X", "Roast", "", "M3V", "Un"}
// 	masterList["0407"] = []string{"NSSC 0407", "X400000-0", "", "Ba Va", "0/0/X", "Temp", "G", "M1V", "Un"}
// 	masterList["0409"] = []string{"Derceto (NSSC 0409)", "C3A3353-8", "", "Lo", "7/5/X", "Cold", "G", "K7V M6V", "Un"}
// 	masterList["0504"] = []string{"Bruadair (NSSC 0504)", "D959322-6", "", "Lo", "A/8/X", "Temp", "G", "M2V", "Un"}
// 	masterList["0507"] = []string{"NSSC 0507", "X6B0000-0", "", "Ba De", "2/1/X", "Roast", "G", "K6V M0V", "Un"}
// 	masterList["0508"] = []string{"NSSC 0508", "XADA000-0", "", "Ba Fl", "1/8/X", "Temp", "", "K9V", "Un"}
// 	masterList["0603"] = []string{"Epsilon Scorpii", "X100000-0", "", "Ba Va", "0/0/X", "Frozen", "G", "K2III", "Un"}
// 	masterList["0604"] = []string{"NSSC 0604", "XAFA000-0", "", "Ba Wa", "4/0/X", "Temp", "", "M1V", "Un"}
// 	masterList["0606"] = []string{"Xi Serpentis", "X411000-0", "", "Ba Ic", "0/0/X", "Frozen", "", "F0III", "Un"}
// 	masterList["0607"] = []string{"NSSC 0607", "X200000-0", "", "Ba Va", "0/0/-1", "Temp", "G", "K6V M2V M4V", "Un"}
// 	masterList["0608"] = []string{"NSSC 0608", "X898000-0", "", "Ba Ga", "3/6/+1", "Temp", "G", "K2V", "Un"}
// 	masterList["0701"] = []string{"Delta Herculis", "X000000-0", "", "Ba Va", "0/0/X", "N/A", "", "A1V", "Un"}
// 	masterList["0703"] = []string{"NSSC 0703", "X100000-0", "", "Ba Va", "0/0/+1", "Temp", "G", "M1V", "Un"}
// 	masterList["0707"] = []string{"NSSC 0707", "X898000-0", "", "Ba Ga", "7/4/X", "Temp", "G", "M9V M9V M6V", "Un"}
// 	masterList["0709"] = []string{"NSSC 0709", "X3A1000-0", "", "Ba Fl", "4/2/X", "Temp", "G", "M2V BD", "Ex"}
// 	masterList["0710"] = []string{"Ivanhoe (NSSC 0710)", "X689000-0", "Uh", "Ba", "2/2/X", "Temp", "G", "M0V", "Ex"}
// 	masterList["0803"] = []string{"NSSC 0803", "X410000-0", "", "Ba", "0/0/X", "Hot", "G", "M5V", "Un"}
// 	masterList["0808"] = []string{"NSSC 0808", "X879000-0", "", "Ba", "9/3/+2", "Temp", "G", "K2V", "Ex"}
// 	masterList["0809"] = []string{"NSSC 0809", "X410000-0", "", "Ba", "0/0/X", "Temp", "", "M3V", "Ex"}
// 	masterList["0810"] = []string{"NSSC 0810", "X200000-0", "", "Ba Va", "0/0/+1", "Cold", "G", "M4V", "Ex"}
// 	masterList["0904"] = []string{"Daeva (NSSC 0904)", "X100000-0", "Uh", "Ba Va", "0/0/X", "Temp", "G", "M5V M7V", "Un"} //EPSILON SAGITTARIUS (SUBSECTOR B)
// 	masterList["0907"] = []string{"NSSC 0907", "X3A0000-0", "", "Ba De", "4/0/+1", "Hot", "", "M3V", "Un"}
// 	masterList["0908"] = []string{"Baba Yaga (NSSC 0908)", "E66A252-5", "", "Lo Lt Wa", "8/3/X", "Temp", "G", "K6V", "Ex"}
// 	masterList["0909"] = []string{"NSSC 0909", "X310000-0", "", "Ba", "0/0/+1", "Temp", "G", "M7V", "Ex"}
// 	masterList["1001"] = []string{"NSSC 1001", "X633000-0", "", "Ba Po", "4/3/+0", "Frozen", "G", "M9V", "Un"}
// 	masterList["1002"] = []string{"NSSC 1002", "X300000-0", "", "Ba Va", "0/0/-4", "Frozen", "", "BD", "Un"}
// 	masterList["1003"] = []string{"NSSC 1003", "X794000-0", "", "Ba Ga", "3/0/+1", "Temp", "G", "M9V", "Un"}
// 	masterList["1006"] = []string{"NSSC 1006", "X401000-0", "", "Ba Ic Va", "0/0/X", "Frozen", "", "BD", "Un"}
// 	masterList["1009"] = []string{"NSSC 1009", "X589000-0", "", "Ba", "4/7/+0", "Hot", "G", "M4V M5V", "Ex"}
// 	masterList["1010"] = []string{"NSSC 1010", "X503000-0", "", "Ba Ic Va", "0/0/X", "Temp", "G", "M9V M9V", "Ex"}
// 	masterList["1108"] = []string{"NSSC 1108", "X8D5000-0", "", "Ba Fl", "4/0/X", "Hot", "G", "K7V", "Ex"}
// 	masterList["1110"] = []string{"NSSC 1110", "X410000-0", "", "Ba", "0/0/X", "Temp", "G", "K6V M6V", "Ex"}
// 	masterList["1205"] = []string{"Mbwiri (GL700.2)", "X878000-0", "Uh", "Ba Ga", "8/0/-2", "Temp", "G", "K0V", "Un"}
// 	masterList["1206"] = []string{"NSSC 1206 X8B4000-0 Ba Fl 1/2/X Roast G M1V Ex"}
// 	masterList["1209"] = []string{"NSSC 1209 X544000-0 Ba Ga 3/0/X Temp G K1V Ex"}
// 	masterList["1301"] = []string{"NSSC 1301 X8A6000-0 Ba Fl 6/4/+1 Temp G K1V BD Un"}
// 	masterList["1303"] = []string{"NSSC 1303 X689000-0 Ba 6/1/-1 Hot G M2V Un"}
// 	masterList["1304"] = []string{"Epsilon Sagittarii X000000-0 Ba Va 0/0/X N/A A0II Un"}
// 	masterList["1306"] = []string{"NSSC 1306 X300000-0 Ba Va 0/0/X Temp G K3V M3V Un"}
// 	masterList["1307"] = []string{"NSSC 1307 X558000-0 Ba Ga 6/1/+0 Cold G K1V Ex"}
// 	masterList["1308"] = []string{"Naamah (NSSC 1308) X996000-0 Uh Ba Ga 1/1/-1 Temp G K3V Ex"}
// 	masterList["1309"] = []string{"NSSC 1309 X523000-0 Ba Po 3/0/-1 Temp G M7V Ex"}
// 	masterList["1310"] = []string{"NSSC 1310 X686000-0 Ba Ga 2/2/X Temp G K8V Ex"}
// 	masterList["1405"] = []string{"NSSC 1405 X79A000-0 Ba Wa 7/1/+0 Hot G M5V Un"}
// 	masterList["1407"] = []string{"NSSC 1407 X564000-0 Ba Ga 3/7/+0 Temp G M4V M5V Ex"}
// 	masterList["1408"] = []string{"NSSC 1408 X4A0000-0 Ba De 6/0/X Temp G K5V Ex"}
// 	masterList["1410"] = []string{"New Prospects (NSSC 1410) C200331-B R Lo Va 0/0/X Temp G M8V Ex/A"}
// 	masterList["1501"] = []string{"Narcissa (NSSC 1501) D3A0342-8 Lo De 2/0/X Temp G M3V Un"}
// 	masterList["1502"] = []string{"Trashim (NSSC 1502) X100000-0 Uhm Ba Va 0/0/+1 Temp M5V Un"}
// 	masterList["1503"] = []string{"Bezaleel (NSSC 1503) D200233-8 Lo Va 0/0/X Cold G BD Un"}
// 	masterList["1509"] = []string{"New Alberta (NSSC 1509) X567000-0 Uh Ba Ga 6/6/+0 Cold G K6V Ex"}
// 	masterList["1607"] = []string{"NSSC 1607 X000000-0 As Ba Va 0/0/+0 Temp G M4V Un"}
// 	masterList["1609"] = []string{"NSSC 1609 X400000-0 Ba Va 0/0/-1 Hot G M5V Ex"}
// 	masterList["1610"] = []string{"NSSC 1610 XA7A000-0 Ba Wa 7/4/X Temp M1V M2V Ex"}
// 	masterList["1701"] = []string{"NSSC 1701 X000000-0 As Ba Va 0/0/+1 Temp G M9V Un"}
// 	masterList["1703"] = []string{"Abyzou (NSSC 1703) X78A000-0 Uh Ba Wa 7/1/X Frozen G M4V M2V BD Un"}
// 	masterList["1704"] = []string{"NSSC 1704 X8B1000-0 Ba Fl 2/0/X Roast G M3V M6V Un"}
// 	masterList["1705"] = []string{"NSSC 1705 XA7A000-0 Ba Wa 8/5/+3 Temp G K5V Un"}
// 	masterList["1707"] = []string{"Zalambur (NSSC 1707) X8B0000-0 Uh Ba De 0/0/+0 Hot M1V Un"}
// 	masterList["1709"] = []string{"NSSC 1709 X877000-0 Ba Ga 7/5/X Temp K8V M9V Ex"}
// 	masterList["1710"] = []string{"NSSC 1710 X578000-0 Ba Ga 9/3/X Temp G M2V M2V Ex"}
// 	masterList["1801"] = []string{"Tau Sagittarii X400000-0 Uh Ba Va 0/0/+0 Hot G K2III Un"}
// 	masterList["1802"] = []string{"Baldoon (GL749) C200333-8 Lo Va 0/0/+3 Temp G F8V Un"}
// 	masterList["1803"] = []string{"NSSC 1803 X100000-0 Ba Va 0/0/X Temp G M1V Un"}
// 	masterList["1805"] = []string{"NSSC 1805 X657000-0 Ba Ga 7/0/-1 Temp G M2V M9V Un"}
// 	masterList["1806"] = []string{"NSSC 1806 E86A132-5 Lo Lt Wa 6/7/X Temp K8V M4V Un"}
// 	masterList["1807"] = []string{"Beta Pavonis X310000-0 Ba 0/0/X Roast A8V Un"}
// 	masterList["1808"] = []string{"NSSC 1808 X6A0000-0 Ba De 0/0/+1 Temp G K1V Un"}
// 	masterList["1809"] = []string{"NSSC 1809 X8B0000-0 Uh Ba De 0/0/X Hot K5V BD Ex"}
// 	masterList["1901"] = []string{"NSSC 1901 X698000-0 Ba Ga B/4/X Temp G M2V M4V Un"}
// 	masterList["1903"] = []string{"NSSC 1903 X520000-0 Ba De 1/1/X Cold M3V Un"}
// 	masterList["1905"] = []string{"NSSC 1905 X300000-0 Ba Va 0/0/+0 Cold G M7V Un"}
// 	masterList["1907"] = []string{"NSSC 1907 X200000-0 Ba Va 0/0/X Temp G K4V Un"}
// 	masterList["1909"] = []string{"NSSC 1909 X77A000-0 Ua Ba Wa 7/2/+3 Temp G K5V K9V Un"}
// 	masterList["1910"] = []string{"NSSC 1910 X85A000-0 Ba Wa 4/8/+0 Temp G K0V M0V Ex"}
// 	masterList["2006"] = []string{"Bolthole C3A0333-8 Ua Lo De 0/0/X Temp G M5V Un"}
// 	masterList["2010"] = []string{"NSSC 2010 X593000-0 Ua Ba 6/5/+0 Temp K0V Ex"}
// 	masterList["2104"] = []string{"GL756.1 X300000-0 Ba Va 0/0/X Temp G K5V Un"}
// 	masterList["2107"] = []string{"NSSC 2107 X8B0000-0 Ba De 0/0/-3 Roast G M6V M6V Un"}
// 	masterList["2201"] = []string{"NSSC 2201 X547000-0 Ba Ga 8/2/+0 Frozen G K2V Un"}
// 	masterList["2203"] = []string{"Styx (NSSC 2203) X733000-0 Uh Ba Po 1/5/+0 Cold G M0V M0V Un"}
// 	masterList["2204"] = []string{"Endless Blossom (NSSC 2204) X554200-4 Lo Ga 9/5/X Cold G K4V M3V Un"}
// 	masterList["2301"] = []string{"NSSC 2301 X545000-0 Ba Ga 5/4/X Temp M5V M6V BD Un"}
// 	masterList["2304"] = []string{"NSSC 2304 X8C0000-0 Ba De 0/0/X Roast M1V M5V Un"}
// 	masterList["2307"] = []string{"NSSC 2307 X535000-0 Ba 0/0/+1 Frozen G M1V M4V Un"}
// 	masterList["2308"] = []string{"NSSC 2308 X400000-0 Ba Va 0/0/X Temp G M5V M6V Un"}
// 	masterList["2309"] = []string{"NSSC 2309 X100000-0 Ba Va 0/0/X Temp M6V Un"}
// 	masterList["2310"] = []string{"Tootega (16 Cygni) C4A0232-8 Lo De 2/5/+1 Temp G G1V G2V M1V Un"}
// 	masterList["2405"] = []string{"NSSC 2405 X7C0000-0 Ba De 0/0/-3 Frozen BD Un"}
// 	masterList["2406"] = []string{"Senka (NSSC 2406) C9A6200-8 Lo Fl 0/0/X Temp G M1V Un"}
// 	masterList["2407"] = []string{"NSSC 2407 X310000-0 Ba 0/0/X Frozen G BD Un"}
// 	masterList["2502"] = []string{"NSSC 2502 X782000-0 Ba 4/7/+1 Hot G K9V Un"}
// 	masterList["2508"] = []string{"NSSC 2508 XA8A000-0 Ba Wa 8/4/X Temp M3V Un"}
// 	masterList["2509"] = []string{"NSSC 2509 X536000-0 Ba 4/2/+0 Frozen G M4V M6V Un"}
// 	masterList["2510"] = []string{"NSSC 2510 X200000-0 Ba Va 0/0/+3 Hot G K7V Un"}
// 	masterList["2602"] = []string{"NSSC 2602 X9B0000-0 Ba De 0/0/+0 Roast G M6V Un"}
// 	masterList["2603"] = []string{"NSSC 2603 X400000-0 Ba Va 0/0/+0 Cold G M4V M5V Un"}
// 	masterList["2604"] = []string{"NSSC 2604 X525000-0 Ua Ba 4/1/X Frozen G M3V Un"}
// 	masterList["2605"] = []string{"NSSC 2605 X9B0000-0 Ba De 0/0/+2 Hot G K2V K5V Un"}
// 	masterList["2606"] = []string{"NSSC 2606 X411000-0 Ba Ic 1/0/X Temp G M4V M4V M6V Un"}
// 	masterList["2607"] = []string{"NSSC 2607 X410000-0 Ba 0/0/-3 Frzoen G BD Un"}
// 	masterList["2609"] = []string{"NSSC 2609 XA63000-0 Ba 2/5/X Roast G M9V BD Un"}
// 	masterList["2610"] = []string{"NSSC 2610 X410000-0 Ba 0/0/X Temp M4V Un"}
// 	masterList["2701"] = []string{"NSSC 2701 X896000-0 Ba Ga 7/3/+1 Temp G K8V Un"}
// 	masterList["2702"] = []string{"NSSC 2702 X5A1000-0 Ba Fl 0/0/X Roast G M7V Un"}
// 	masterList["2704"] = []string{"NSSC 2704 X758000-0 Ba Ga 6/0/X Temp K3V Un"}
// 	masterList["2706"] = []string{"NSSC 2706 X567000-0 Ba Ga 8/5/+2 Temp K0V Un"}
// 	masterList["2707"] = []string{"NSSC 2707 X987000-0 Ba Ga 7/9/X Roast G K0V M6V Un"}
// 	masterList["2709"] = []string{"HD195019 X620000-0 Ba De 0/0/+0 Temp G G3IV Un"}
// 	masterList["2710"] = []string{"NSSC 2710 X300000-0 Ba Va 0/0/X Temp M2V Un"}
// 	masterList["2803"] = []string{"Jormungand (NSSC 2803) D569321-8 Ga Lo 6/5/+3 Temp G K5V K7V M5V Un"}
// 	masterList["2806"] = []string{"Rahab (NSSC 2806) X310000-0 Uh Ba 0/0/+0 Hot G K5V M5V Un"}
// 	masterList["2807"] = []string{"NSSC 2807 X655000-0 Ua Ba Ga 6/8/X Hot G M1V M1V Un"}
// 	masterList["2810"] = []string{"NSSC 2810 X300000-0 Ba Va 0/0/-1 Cold G M6V Un"}
// 	masterList["2901"] = []string{"NSSC 2901 X684000-0 Ba Ga 6/6/+0 Temp G K1V M2V Un"}
// 	masterList["2903"] = []string{"NSSC 2903 X200000-0 Uh Ba Va 0/0/+0 Temp G K2V Un"}
// 	masterList["2904"] = []string{"NSSC 2904 X666000-0 Ba Ga B/3/X Temp G K0V K6V Un"}
// 	masterList["2905"] = []string{"NSSC 2905 XA9A000-0 Ba Wa 5/5/+0 Temp G M1V Un"}
// 	masterList["2907"] = []string{"NSSC 2907 X87A000-0 Ba Wa A/1/X Temp G K3V M7V Un"}
// 	masterList["2909"] = []string{"NSSC 2909 X878000-0 Ba Ga 5/1/X Temp G K3V Un"}
// 	masterList["3003"] = []string{"NSSC 3003 X4A0000-0 Ba De 0/0/X Temp G M4V M6V Un"}
// 	masterList["3007"] = []string{"NSSC 3007 X3A0000-0 Ba De 0/0/+1 Roast G K3V Un"}
// 	masterList["3008"] = []string{"NSSC 3008 X200000-0 Ba Va 0/0/X Temp G K6V Un"}
// 	masterList["3009"] = []string{"NSSC 3009 X558000-0 Ba Ga 5/6/+0 Temp G M3V M5V Un"}
// 	masterList["3103"] = []string{"NSSC 3103 X100000-0 Ba Va 0/0/X Temp G M5V Un"}
// 	masterList["3105"] = []string{"NSSC 3105 X400000-0 Ba Va 0/0/X Temp M6V Un"}
// 	masterList["3108"] = []string{"NSSC 3108 X97000-0 Ba Ga 4/5/+2 Hot G K6V M8V Un"}
// 	masterList["3201"] = []string{"NSSC 3201 X733000-0 Ba Po 1/0/X Temp G K7V Un"}
// 	masterList["3202"] = []string{"NSSC 3202 X200000-0 Ba Va 0/0/X Temp M1V Un"}
// 	masterList["3203"] = []string{"NSSC 3203 X200000-0 Ua Ba Va 0/0/+0 Temp G M7V BD Un"}
// 	masterList["3205"] = []string{"NSSC 3205 X9A4000-0 Ba Fl 6/1/+0 Temp G K1V M3V Un"}
// 	masterList["3208"] = []string{"NSSC 3208 X660000-0 Ba De 0/0/X Roast G K9V Un"}
// 	masterList["3209"] = []string{"NSSC 3209 X523000-0 Ba Po 0/0/X Frozen G M1V Un"}
// 	masterList["0115"] = []string{"Gamma Bootis X300000-0 Ba Va 0/0/+1 Cold A7IV Ex"}
// 	masterList["0116"] = []string{"Baal (NSSC 0116) D758211-7 Ga Lo 4/7/X Temp G K6V M4V Ov"}
// 	masterList["0117"] = []string{"Sigma Librae I D400311-8 Lo Va 0/0/X Temp G M2III Ov"}
// 	masterList["0118"] = []string{"Chemosh (Theta Centauri) D560453-8 S De Ni 1/1/-1 Temp G K0III Ov"}
// 	masterList["0212"] = []string{"Alpha Circinis X700000-0 Um Ba Va 0/0/X Roast A7V Ex"}
// 	masterList["0216"] = []string{"NSSC 0216 X8A0000-0 Ba De 1/2/X Hot G M6V Ex"}
// 	masterList["0218"] = []string{"NSSC 0218 X300000-0 Ba Va 0/0/X Temp G M3V Ex"}
// 	masterList["0219"] = []string{"Ashima (NSSC 0219) C200384-9 S Lo Va 0/0/X Cold G M2V M5V Ov"}
// 	masterList["0220"] = []string{"Tau Bootis D731552-5 Lt Ni Po 6/1/X Frozen G F7V Ov"}
// 	masterList["0311"] = []string{"NSSC 0311 X875000-0 Ba Ga A/3/X Temp G K8V Ex"}
// 	masterList["0313"] = []string{"NSSC 0313 X976000-0 Ba Ga 8/3/+0 Hot K1V M3V Ex"}
// 	masterList["0316"] = []string{"NSSC 0316 X200000-0 Ba Va 0/0/X Temp G M9V Ex"}
// 	masterList["0318"] = []string{"Bootis Station (44 Bootis) E583121-6 Lo 6/2/X Temp G G0V Ov"}
// 	masterList["0319"] = []string{"Nu Lupi C799343-6 S Ga Lo A/4/X Temp G G2V Ov"}
// 	masterList["0411"] = []string{"14 Herculis X687000-0 Ba Ga A/8/+3 Hot G K0V Ex"}

// 	masterList["0414"] = []string{"Berith (Rho Coronae) E594132-5 Ga Lo Lt 7/1/X Hot G G5V Ov"}
// 	masterList["0416"] = []string{"Eshmun (26 Draconis) E651222-3 Ua Lo Lt 2/8/+1 Frozen G G0V K3V Ov"}
// 	masterList["0418"] = []string{"Kothar (NSSC 0418) C310411-9 Ni 0/0/X Temp M7V BD Ov"}
// 	masterList["0419"] = []string{"Gamma Serpentis C410434-9 M Ua Ni 0/0/X Hot G F6V Ov"}
// 	masterList["0420"] = []string{"Dagon (Lambda Serpentis) C66A431-6 Ni Wa 4/6/+2 Temp G G0V Ov"}
// 	masterList["0512"] = []string{"GL651 X592000-0 Ba 7/1/+1 Temp G G8V Ex"}
// 	masterList["0514"] = []string{"Moloch (NSSC 0514) E4A0255-8 P De Lo 0/0/+2 Temp G M7V Ov"}
// 	masterList["0516"] = []string{"Melqart (NSSC 0516) E100222-9 Lo Va 0/0/X Temp K5V BD Ov"}
// 	masterList["0517"] = []string{"Beta Trianguli EA36242-8 Lo 1/2/-1 Temp G F2III Ov"}
// 	masterList["0518"] = []string{"Psi Serpentis E543300-5 Lo Lt Po 8/4/-1 Hot G G5V Ov"}
// 	masterList["0613"] = []string{"Astarte (NSSC 0613) C787311-8 Ga Lo A/4/X Hot K2V M3V Ov"}
// 	masterList["0615"] = []string{"Lena's Legacy (18 Scorpii) E100231-8 Lo Va 0/0/+1 Temp G G2V Ov"}
// 	masterList["0617"] = []string{"NSSC 0617 X410000-0 Ba 0/0/-2 Temp M5V Ov"}
// 	masterList["0618"] = []string{"Zeta Herculis E000220-8 As Lo Va 0/0/+3 Temp G G0IV K0V Ov/A"}
// 	masterList["0620"] = []string{"Ashera (NSSC 0620) B744586-9 N S Ag Ga Ni 8/4/-2 Temp G K4V Fr"}
// 	masterList["0713"] = []string{"Hadad (NSSC 0713) E785242-6 Ga Lo 6/7/+1 Temp G K3V M5V Ov"}
// 	masterList["0716"] = []string{"Zeta Trianguli C583551-8 Ni 6/7/X Cold G F6V G1V Ov"}
// 	masterList["0719"] = []string{"Anat (12 Ophiuchi) B567686-A N M S R Ag Ga Ni Ri 8/5/X Temp G K2V Fr"}
// 	masterList["0811"] = []string{"Eta Ophiuchi X410000-0 Ba 0/0/X Roast G A2V Ex"}
// 	masterList["0812"] = []string{"Mot (NSSC 0812) E300353-9 Lo Va 0/0/X Temp G M4V Ov"}
// 	masterList["0814"] = []string{"19 Draconis X200000-0 Ba Va 0/0/+0 Cold G F6V Ex"}
// 	masterList["0815"] = []string{"GL620 X8A4000-0 Ba Fl 6/0/X Temp G G5V DA Ex"}
// 	masterList["0817"] = []string{"Gorynych (NSSC 0817) EAB2373-9 Ua Fl Lo 4/2/+0 Hot M5V M6V Ov/A"}
// 	masterList["0818"] = []string{"Mamlambo (GL638) C584642-8 S Ua Ag Ga Ni Ri B/6/-1 Temp G K5V Fr"}
// 	masterList["0818a"] = []string{"uKqili (GL638) D311341-8 Ic Ni 1/1/-1 Frozen Fr"}
// 	masterList["0819"] = []string{"Atargatis (NSSC 0819) C649583-9 Ga Ni 6/2/-2 Temp G M2V BD Fr"}

// 	masterList["0919"] = []string{"Diwata", "D545573-8", "", "", "", "", "", "G", "", "If"}
// 	masterList["1215"] = []string{"Fort Chang", "C400364-9", "M", "", "", "", "", "G", "", "Fr"}
// 	masterList["1220"] = []string{"Rusalka", "B678686-A", "N M S", "", "", "", "", "G", "K8VI", "If"}

// 	masterList["1711"] = []string{"Enkidu (NSSC 1711) D545374-6 Ga Lo A/3/+0 Temp G K0V K6V Ex"}
// 	masterList["1712"] = []string{"NSSC 1712 X000000-0 As Ba Va 0/0/-1 Temp G M2V Ex"}
// 	masterList["1713"] = []string{"Humbaba (NSSC 1713) D3A0300-8 P De Lo 2/0/X Temp G M4V Ov/A"}
// 	masterList["1716"] = []string{"NSSC 1716 X567100-0 Ga Lo Lt 5/7/X Temp G K3V Ex"}
// 	masterList["1717"] = []string{"Hanbi (NSSC 1717) D65743A-4 S Ga Lt Ni 9/2/+0 Temp G K2V Ov/A"}
// 	masterList["1720"] = []string{"Malakbel (NSSC 1720) E998121-7 Ga Lo 6/3/+2 Temp K2V BD BD Ov"}
// 	masterList["1812"] = []string{"NSSC 1812 X584000-0 Ba Ga 3/4/X Hot G K4V M9V BD Ex"}
// 	masterList["1815"] = []string{"Alshain E585220-2 Ga Lo Lt B/4/X Hot G G8V M3V Ov"}
// 	masterList["1817"] = []string{"NSSC 1817 XA93000-0 Um Ba 4/4/X Roast G M9V Ov/A"}
// 	masterList["1911"] = []string{"NSSC 1911 X795000-0 Ba Ga 7/2/X Temp G K6V Ex"}
// 	masterList["1912"] = []string{"NSSC 1912 X400000-0 Ba Va 0/0/-1 Temp G M7V Ex"}
// 	masterList["1913"] = []string{"Arsu (NSSC 1913) X8A1000-0 Uh Ba Fl 3/0/X Frozen G BD Ex"}
// 	masterList["1917"] = []string{"Poludnitsa (NSSC 1917) E661300-7 Lo 8/1/+1 Roast G K6V Ov/A"}
// 	masterList["2011"] = []string{"NSSC 2011 X748000-0 Ba Ga 5/4/+2 Temp G K6V Ex"}
// 	masterList["2013"] = []string{"NSSC 2013 X201000-0 Ba Ic Va 0/0/+0 Temp M5V Ex"}
// 	masterList["2016"] = []string{"Nanna (GL777) E100212-9 Lo Va 0/0/+0 Temp G G7IV M5V Ov"}
// 	masterList["2018"] = []string{"Avanim (NSSC 2018)", "B411585-A", "N M S", "Ic Ni Va", "0/0/X", "Frozen", "G", "M2V M4V", "Ov"}
// 	masterList["2020"] = []string{"Gamma Pavonis V C668321-7 Ga Lo 3/8/X Cold G F6V Ov"}
// 	masterList["2111"] = []string{"NSSC 2111 X568000-0 Ba Ga 3/2/X Temp M7V Ex"}
// 	masterList["2112"] = []string{"HD192263 X311000-0 Ba Ic 0/0/X Temp G K2V Ex"}
// 	masterList["2115"] = []string{"NSSC 2115 X300000-0 Ba Va 0/0/+3 Temp G K7V Ex"}
// 	masterList["2119"] = []string{"Al-Qaum (NSSC 2119) C100483-A M R Ni Va 0/0/X Hot G M2V Ov"}
// 	masterList["2216"] = []string{"Epsilon Cygnii E687244-5 Ua Ga Lo Lt A/3/+0 Temp G K0III Ov"}
// 	masterList["2217"] = []string{"Al-Tawhid (NSSC 2217) E552285-7 Lo Po 9/3/X Temp G K8V Ov"}
// 	masterList["2218"] = []string{"Psi Capricorni E000132-9 As Lo Va 0/0/+3 Temp F5V Ov"}
// 	masterList["2219"] = []string{"Iota Pegasi X000000-0 As Ba Va 0/0/+0 Temp F5V G9V Ov/A"}
// 	masterList["2220"] = []string{"Al-Kalimah (NSSC 2220) C544484-6 Ga Ni 2/1/X Temp G K4V K7V Ov"}
// 	masterList["2312"] = []string{"NSSC 2312 X500000-0 Ba Va 0/0/X Temp G M1V Ex"}
// 	masterList["2314"] = []string{"NSSC 2314 X412000-0 Ba Ic 0/0/X Temp G M2V Ex"}
// 	masterList["2316"] = []string{"NSSC 2316 X410000-0 Ba 0/0/X Cold G M5V Ex"}
// 	masterList["2317"] = []string{"Aglibol (NSSC 2317) D200342-9 Lo Va 0/0/X Temp G M5V M6V Ov"}
// 	masterList["2319"] = []string{"Basmala (NSSC 2319) DA85321-7 Ga Lo 5/5/-1 Hot G K0V Ov"}
// 	masterList["2320"] = []string{"Delta Capricorni C9A3384-9 Ba Fl 0/0/+0 Roast G A5V F2V Ov"}
// 	masterList["2412"] = []string{"NSSC 2412 X786000-0 Ba Ga A/7/+1 Hot G K8V Ex"}
// 	masterList["2414"] = []string{"NSSC 2414 X571000-0 Ba 2/2/+3 Temp G K8V M3V Ex"}
// 	masterList["2415"] = []string{"NSSC 2415 X100000-0 Ba Va 0/0/X Temp G M4V BD Ex"}
// 	masterList["2416"] = []string{"Chukwu (NSSC 2416) E886342-4 Ga Lo Lt 6/6/X Temp G K6V Ov"}
// 	masterList["2417"] = []string{"Eta Cephei D200455-9 S Ni Va 0/0/+0 Temp K0IV Ov"}
// 	masterList["2419"] = []string{"Nippur C100332-8 Um Lo Va 0/0/-1 Temp G M7V BD Ov"}
// 	masterList["2511"] = []string{"NSSC 2511 X94A000-0 Ba Wa 3/5/X Temp K2V Un"}
// 	masterList["2512"] = []string{"NSSC 2512 X100000-0 Ba Va 0/0/X Temp G M9V Un"}
// 	masterList["2514"] = []string{"NSSC 2514 X4A0000-0 Ba De 4/1/-3 Roast G M2V M6V Ex"}
// 	masterList["2515"] = []string{"NSSC 2515 X300000-0 Ba Va 0/0/-3 Temp G M6V Ex"}
// 	masterList["2517"] = []string{"NSSC 2517 X200000-0 Ba Va 0/0/X Cold G M9V Ex"}
// 	masterList["2518"] = []string{"Fenghuang (NSSC 2518) E98A385-7 Lo Wa 2/3/+1 Hot G M0V Ov"}
// 	masterList["2519"] = []string{"Nueva Vilcabamba (NSSC 2519) B898553-A M N S R Ag Ga Ni 9/2/X Temp G K8V M3V Ov"}
// 	masterList["2614"] = []string{"NSSC 2614 X533000-0 Ba Po 3/3/+0 Cold G M1V M3V Ex"}
// 	masterList["2615"] = []string{"NSSC 2615 X4A0000-0 Ba de 1/0/X Cold G M2V M6V Ex"}
// 	masterList["2618"] = []string{"Fusang (NSSC 2618) E9D9311-8 Fl Lo 2/4/+0 Temp G K9V Ov/A"}
// 	masterList["2712"] = []string{"NSSC 2712 X848000-0 Ua Ba Ga A/5/+1 Cold G K0V Un"}
// 	masterList["2714"] = []string{"Medraut (NSSC 2714) C562372-9 P Lo 5/6/+0 Cold K5V M6V Ex"}
// 	masterList["2715"] = []string{"NSSC 2715 X410000-0 Ba 0/0/-3 Frozen G BD Ex"}
// 	masterList["2716"] = []string{"NSSC 2716 X3A0000-0 Ba De 0/0/X Roast M6V M6V Ex"}
// 	masterList["2718"] = []string{"NSSC 2718 X571000-0 Ba 9/7/X Cold G K1V Ex"}
// 	masterList["2719"] = []string{"NSSC 2719 X623000-0 Ba Po 0/0/+0 Temp G M6V Ex"}
// 	masterList["2811"] = []string{"NSSC 2811 X310000-0 Ba 0/0/+2 Temp G K6V M3V Un"}
// 	masterList["2812"] = []string{"NSSC 2812 X4A0000-0 Ba De 6/5/+0 Temp G M2V Un"}
// 	masterList["2814"] = []string{"NSSC 2814 X410000-0 Ba 0/0/+1 Temp G M3V M7V Ex"}
// 	masterList["2815"] = []string{"NSSC 2815 X6B4000-0 Ba Fl 0/0/X Roast G M1V BD Ex"}
// 	masterList["2816"] = []string{"NSSC 2816 X657000-0 Ba Ga 3/6/X Cold G M1V Ex"}
// 	masterList["2817"] = []string{"NSSC 2817 X988000-0 Ba Ga 5/3X Temp G K0V M1V Ex"}
// 	masterList["2819"] = []string{"Tinophoth (NSSC 2819) E879374-7 Lo 4/1/+3 Temp G K6V Ov/A"}
// 	masterList["2914"] = []string{"NSSC 2914 X7A2000-0 Ba Fl 7/4/+1 Hot G K6V K7V Un"}
// 	masterList["2915"] = []string{"NSSC 2915 X644000-0 Ba Ga 5/6/X Cold K9V Ex"}
// 	masterList["2918"] = []string{"NSSC 2918 X573000-0 Ba 8/6/X Temp G K6V M2V Ex"}
// 	masterList["3011"] = []string{"NSSC 3011 X645000-0 Ba Ga 6/5/+2 Temp G K9V Un"}
// 	masterList["3012"] = []string{"NSSC 3012 X410000-0 Ba 0/0/X Temp G M1V M9V Un"}
// 	masterList["3015"] = []string{"NSSC 3015 X200000-0 Ba Va 0/0/X Cold G K5V Ex"}
// 	masterList["3016"] = []string{"NSSC 3016 X634000-0 Ba 1/3/+0 Temp G M6V M9V Ex"}
// 	masterList["3017"] = []string{"Ehcatl (NSSC 3017) E764200-5 Ga Lo A/3/X Temp G K2V Ex"}
// 	masterList["3018"] = []string{"NSSC 3018 X99A000-0 Ba Wa 9/0/X Hot G K8V Ex"}
// 	masterList["3020"] = []string{"NSSC 3020 X77A000-0 Ba Wa 5/0/+1 Hot G K5V Ex"}
// 	masterList["3113"] = []string{"NSSC 3113 X642000-0 Ba Po 7/0/+1 Temp K3V Un"}
// 	masterList["3114"] = []string{"NSSC 3114 X764000-0 Ba Ga 4/8/+2 Temp G K7V Un"}
// 	masterList["3115"] = []string{"NSSC 3115 X3A0000-0 Ba De 3/0/+2 Temp G M2V M9V Un"}
// 	masterList["3116"] = []string{"GL9769 X200000-0 Ba Va 0/0/X Temp G G0V Ex"}
// 	masterList["3117"] = []string{"Itzli (NSSC 3117) X651000-0 Uh Ba Po 3/1/X Temp G K6V Ex"}
// 	masterList["3118"] = []string{"Rho Indi X988000-0 Ba Ga 7/6/-1 Temp G K3V M5V Ex"}
// 	masterList["3120"] = []string{"HD217107 X300000-0 Ba Va 0/0/+0 Temp G G8IV Ex"}
// 	masterList["3213"] = []string{"NSSC 3213 X410000-0 Ba 0/0/X Temp G M9V Un"}
// 	masterList["3214"] = []string{"Amimitl (NSSC 3214) X687222-4 Ga Lo Lt 5/6/X Temp G K0V M2V Un"}
// 	masterList["3216"] = []string{"Theta Pegasi X3A0000-0 Ba De 0/0/+1 Hot G A2V Ex"}
// 	masterList["3219"] = []string{"Epsilon Grus X6A1000-0 Ba Fl 0/0/X Temp G A2V Ex"}
// 	masterList["3220"] = []string{"Delta Aquari X3A0000-0 Ba De 0/0/+1 Temp A3V Ex"}
// 	masterList["0122"] = []string{"Muphrid X6A0000-0 Ba De 4/1/X Hot G G0IV BD Ov/A"}
// 	masterList["0124"] = []string{"Alpha Corvi X620000-0 Ba De 0/0/+0 Cold F0V Ov/A"}
// 	masterList["0126"] = []string{"Porrima X6A0000-0 Ba De 0/0/+1 Hot G F0V F0V Ov/A"}
// 	masterList["0127"] = []string{"Alaraph V D672574-8 S Ni 4/5/+0 Cold G F8V Fr"}
// 	masterList["0128"] = []string{"Zhuravlyova (GL436) D000132-A As Lo Va 0/0/+3 Temp G M2V Fr"}
// 	masterList["0129"] = []string{"Denebola IV E6A2253-8 Fl Lo 0/0/X Temp A3V Fr"}
// 	masterList["0130"] = []string{"Catequil (NSSC 0130) CAD9586-A S Fl Ni 3/2/X Hot G K0V Fr"}
// 	masterList["0222"] = []string{"Korolyov (NSSC 0222) E200412-8 Lo Va 0/0/X Temp G M1V M5V Fr"}
// 	masterList["0226"] = []string{"Goddard (GL432) E100332-9 Lo Va 0/0/-2 Hot G K0V M5V Fr"}
// 	masterList["0229"] = []string{"Shenlong (NSSC 0229) C555675-B M S R Ag Ga Ni 8/7/X Temp G K1V Fr"}
// 	masterList["0230"] = []string{"Snegurochka (NSSC 0230) C402586-A Ic Ni Va 0/0/X Frozen G BD Fr"}
// 	masterList["0321"] = []string{"Theta Bootis IV DA87544-7 S P Ag Ga Ni 5/3/X Temp G F7V M2V Fr/A"}
// 	masterList["0322"] = []string{"Arcturus II E978442-6 Ga Ni 3/8/+0 Temp G K2III Fr"}
// 	masterList["0323"] = []string{"Alpha Comae Berenices X7C0000-0 Ba De 0/0/X Roast F5V F5V Ov"}
// 	masterList["0325"] = []string{"NSSC 0325 X3A0000-0 Ba De 2/5/+0 Temp G M6V Ov"}
// 	masterList["0326"] = []string{"Beta Comae Berenices D583642-6 S P Ni 6/2/+1 Temp G0V Fr/A"}
// 	masterList["0327"] = []string{"Zhongli Quan (HR4523) D755632-5 S Ag Ga Lt Ni 4/6/+2 Temp G4V M4V Fr"}
// 	masterList["0328"] = []string{"Yaghuth (Groombridge 1830) D558353-6 Ga Lo 6/3/+1 Frozen G8IV Fr"}
// 	masterList["0329"] = []string{"Nyankopon (NSSC 0329) C732431-9 S Ni Po 4/4/X Temp G M2V Fr"}
// 	masterList["0425"] = []string{"Pachamama (61 Virginis) C678496-8 S P Ga Ni B/4/X Temp G G5V Fr/A"}
// 	masterList["0426"] = []string{"Proserpina (Beta Canum Venaticorum) C5646BB-8 S P Ag Ga Ni Ri 7/6/X Temp G G0V Fr/A"}
// 	masterList["0427"] = []string{"Mama Cocha (61 Ursae Majoris) DA8A322-9 Lo Wa 9/2/+1 Hot G G8V Fr"}
// 	masterList["0429"] = []string{"Herut (NSSC 0429) C200321-A Um Lo Va 0/0/-1 Hot K0V Fr"}
// 	masterList["0523"] = []string{"Erecura (HN Librae) E546344-7 Ga Lo 1/4/X Cold G M3V Fr"}
// 	masterList["0524"] = []string{"Dis Pater (Xi Bootis) D635575-7 S Ni 4/1/X Temp G G8V K4V Fr/A"}
// 	masterList["0527"] = []string{"Nezha (GJ Virginis) C531385-7 M S Lo Po 1/3/X Cold G M5V Fr"}
// 	masterList["0529"] = []string{"Diyu (GL408) D8B0435-9 S De Ni 1/2/X Roast G M2V Fr/A"}
// 	masterList["0530"] = []string{"Sisterhood (Wolf 358) C7A9321-A Ua Fl Lo 4/1/X Temp G M4V BD Fr"}
// 	masterList["0622"] = []string{"Zlota Baba (Wolf 562) B410653-A M N R Ni 0/0/X Temp G M3V Fr"}
// 	masterList["0624"] = []string{"Skoll (Wolf 498) C410484-9 N Ni 0/0/X Temp M1V Fr"}
// 	masterList["0721"] = []string{"Hildegard (NSSC 0721) C4A3511-9 R Fl Ni 5/1/X Cold G M6V Fr"}
// 	masterList["0723"] = []string{"Ganswindt (GL570) E200221-A Lo Va 0/0/+0 Cold G K5V M1V M3V BD Fr"}
// 	masterList["0727"] = []string{"Fenrir (Wolf 424) C635687-8 R Ua Ni 1/1/-1 Temp M5V M7V If"}
// 	masterList["0727a"] = []string{"Gleipnir E200466-8 Ni Va 0/0/-1 Temp If"}
// 	masterList["0729"] = []string{"Changing Winds (WX Ursae Majoris) C8A039C-A De Lo 4/0/-1 Hot G M1V M5V If/A"}
// 	masterList["0821"] = []string{"Prosperity (G180-060) C610252-8 Lo 0/0/+1 Frozen WD Fr"}
// 	masterList["0825"] = []string{"Vayu (NSSC 0825) B4A0613-B R Ua De Ni 4/1/X Temp G M2V In"}
// 	masterList["0828"] = []string{"Aningan (Ross 128) B412754-B M N S Ic 0/0/X Frozen G M4V If"}
// 	masterList["0829"] = []string{"New Horizons (AD Leonis) C300647-A Ni Va 0/0/X Temp G M3V If"}
// 	masterList["0830"] = []string{"Canceri Belt (EI Canceri) C000485-B S As Ni Va 0/0/+0 Temp M5V M5V If"}
// 	masterList["0921"] = []string{"Glushko (Wolf 630) C200356-A Ua Lo Va 0/0/X Hot M2V M2V M4V M7V Fr"}
// 	masterList["0928"] = []string{"Nkrumah (Lalande 21185) B310744-9 S Na 0/0/X Temp G M2V In"}
// 	masterList["1026a"] = []string{"Medea (Alpha Centauri B II) A768988-B M N S R Ga Hi 6/7/-4 Temp G G2V K0V M5V Cr"}
// 	masterList["1026b"] = []string{"Hecate (Alpha Centauri A IV) B400889-B Na Va 0/0/X Temp Cr"}
// 	masterList["1121"] = []string{"Nuwa (36 Ophiuchi) B200699-B Na Ni Va 0/0/+0 Temp G K1V K1V K5V If"}
// 	masterList["1122"] = []string{"Suribachi (GL674) C400311-A Lo Va 0/0/X Temp G M3V If"}
// 	masterList["1123"] = []string{"Keynes (SCR 1845-6357) C100686-A Na Ni Va 0/0/X Temp G M8V BD In"}
// 	masterList["1124"] = []string{"Subarashii (Ross 154) B764746-9 M N Ua Ag Ga 8/4/X Temp M3V In"}
// 	masterList["1125a"] = []string{"Barnard's Belt A000998-B R As Hi In Na Va 0/0/-3 Hot G M4V Cr"}
// 	masterList["1125b"] = []string{"Barnard's Eye B100668-B Ni Va 0/0/-3 Frozen Cr"}
// 	masterList["1130"] = []string{"DX Facility (DX Canceri) E40049B-8 Ni Va 0/0/X Temp M6V If/R"}
// 	masterList["1222"] = []string{"Opportunity (Struve 2398) B8A7657-8 M S Fl Ni A/0/X Temp G M3V M4V If"}
// 	masterList["1226a"] = []string{"Earth (Sol) A867A79-B M N S R Ga Hi B/A/-4 Temp G G2V Cr"}
// 	masterList["1226c"] = []string{"Luna (Sol) A200988-B M N S R Hi In Na Va 0/0/-4 Temp Cr"}
// 	masterList["1226b"] = []string{"Mars (Sol) A422984-B M R Ua Hi In Na Po 1/A/-4 Frozen Cr"}
// 	masterList["1226d"] = []string{"Sol Belt (Sol) A000945-B N R As Hi In Na Va 0/0/-4 Frozen Cr"}
// 	masterList["1226e"] = []string{"Callisto (Sol) B302886-B N R Ic Na Va 1/1/-4 Frozen Cr"}
// 	masterList["1229"] = []string{"Sirius VII B400797-B M Na Va 0/0/X Frozen A1V WD In"}
// 	masterList["1230"] = []string{"Procyon III B753513-9 M Ni Po 1/3/X Cold G F5IV WD If"}
// 	masterList["1325"] = []string{"Gandhi (NSSC 1325) B513742-B M N S Ic 1/0/X Temp M9V Cr"}
// 	masterList["1328"] = []string{"Beira (NSSC 1328) A000759-B As Na Va 0/0/-4 Frozen G BD Cr"}
// 	masterList["1330"] = []string{"Luyten's World (Luyten's Star) D620386-8 Um De Lo 0/0/X Roast G M3V If"}
// 	masterList["1421"] = []string{"Altair Vf D410633-8 Na Ni 0/0/+0 Frozen G A7IV If"}
// 	masterList["1423"] = []string{"Roosevelt (61 Cygni) C400754-B M Na Va 0/0/-1 Temp G K5V K7V In"}
// 	masterList["1429"] = []string{"Kapteyn (Kapteyn's Star) C533735-A R Na Po 0/0/X Cold M1V In"}
// 	masterList["1521"] = []string{"Lincoln (Herschel 5173) C564686-9 Ua Ag Ga Ni Ri 5/7/-1 Temp G K3V M3V Fr"}
// 	masterList["1522"] = []string{"Sigma Draconis III C579697-8 M S Ni 8/5/+1 Temp G K0V If"}
// 	masterList["1524"] = []string{"Temazcal (EZ Aquari) BAF9736-A S Fl 3/0/X Hot G M5V M8V BD In"}
// 	masterList["1527"] = []string{"UV Ceti I C411645-9 Ic Na Ni 0/0/X Temp G M5V M6V In"}
// 	masterList["1529"] = []string{"Teegarden (Teegarden's Star) B630631-B De Ni Na Po 0/0/-4 Temp G M7V In"}
// 	masterList["1624"] = []string{"Epsilon Indi II B868788-9 M S Ag Ga Ri 6/8/-1 Cold G K5V BD BD If"}
// 	masterList["1624a"] = []string{"Epsilon Indi IIIk C4A1488-8 Ni Fl 0/0/-1 Frozen If"}
// 	masterList["1625"] = []string{"Novi Volgograd (Ross 248) B7A0655-B S De Ni 1/1/+0 Hot M6V If"}
// 	masterList["1626"] = []string{"New Canberra (Groombridge 34) C7A0412-A De Ni 0/0/+0 Roast G M2V M3V If"}
// 	masterList["1627"] = []string{"Tau Ceti V B554735-9 M S Ag Ga 5/5/X Temp G G8V If"}
// 	masterList["1627a"] = []string{"Rebecca (Tau Ceti IV) C622365-9 R Lo Po 2/0/X Hot If"}
// 	masterList["1628"] = []string{"Epsilon Eridani III A564713-B M S Ag Ga 8/2/+1 Temp G K2V If"}
// 	masterList["1630"] = []string{"Keid C4A1111-9 S R Fl Lo 9/1/X Temp G K1V WD M5V If"}
// 	masterList["1722"] = []string{"Wa kabout (Ross 775) C546554-9 S Ua Ag Ga Ni 8/1/X Temp G M3V Fr"}
// 	masterList["1725"] = []string{"Ross Station (Ross 780) D530331-9 S De Lo Po 0/0/-2 Hot G M3V If"}
// 	masterList["1727"] = []string{"Van Maanen (Van Maanen's Star) D510512-9 S Ni 0/0/X Frozen G WD If"}
// 	masterList["1729"] = []string{"Black Sands (GL832) D520388-7 S De Lo Po 0/0/-1 Frozen G M3V If"}
// 	masterList["1823"] = []string{"EV Lacertae C100453-B Ni Va 0/0/X Cold G M3V Fr"}
// 	masterList["1827"] = []string{"Eta Cassiopeia E000411-8 As Ni Va 0/0/+1 Hot G G3V K7V If"}
// 	masterList["1830"] = []string{"Novaya Pechenga (82 Eridani) C885649-7 M Ag Ga Ni 6/4/X Hot G G5V BD If"}
// 	masterList["1924"] = []string{"Nikkal (GL892) C675343-4 S Ga Lo Lt 5/0/X Temp G K3V Fr"}
// 	masterList["1925"] = []string{"Pegasus (EQ Pegasi) C410523-8 M S Ni 0/0/-3 Hot G M3V M4V Fr"}
// 	masterList["1929"] = []string{"Dracul (GL33) C875344-8 M S Ga Lo 7/2/X Temp G K2V Fr"}
// 	masterList["2021"] = []string{"New Chryse (GL849) C521411-9 Ni Po 0/0/X Temp G M3V Ov"}
// 	masterList["2022"] = []string{"Fomalhaut VI C8A5384-9 S R Fl Lo 1/0/X Hot G A3V Ov"}
// 	masterList["2026"] = []string{"Beta Hydri III B563686-A M N S Ua Ni 6/8/X Temp G G2V Fr"}
// 	masterList["2026a"] = []string{"Hydri Belt C000586-A As Ni Va 0/0/-1 Frozen Fr"}
// 	masterList["2027"] = []string{"Snowball (Mu Cassipeia) C737311-5 S Lo 5/3/X Frozen G5V M5V Fr"}
// 	masterList["2028"] = []string{"107 Piscium X510000-0 Ba 0/0/X Temp G K1V Fr"}
// 	masterList["2030"] = []string{"Thyestes (HR 753) D687531-5 Ag Ga Lt Ni 7/5/X Temp G K3V M3V M7V Fr"}
// 	masterList["2123"] = []string{"Hubbert (GL884) D789551-6 Ni 3/2/X Temp G K5V Ov"}
// 	masterList["2125"] = []string{"Chelomei (GL1289) E200222-8 Lo Va 0/0/X Temp G M4V Fr"}
// 	masterList["2127"] = []string{"Zeta Tucanae E000342-8 As Lo Va 0/0/-1 Temp G F9V Fr"}
// 	masterList["2129"] = []string{"Rho Eridani E4A0321-8 Lo 0/0/X Temp G K2V K3V Fr"}
// 	masterList["2221"] = []string{"Marcos (NSSC 2221) C310322-9 S Lo 0/0/X Temp G M2V Ov"}
// 	masterList["2224"] = []string{"Apep (NSSC 2224) E743431-6 Ni Po 1/0/X Hot G K6V M2V Ov"}
// 	masterList["2226"] = []string{"Xiuhcoatl (NSSC 2226) C5A0421-A De Ni 2/1/X Roast G M0V Fr"}
// 	masterList["2227"] = []string{"Svoboda (NSSC 2227) C100321-8 S Lo Va 0/0/X Temp G M0V Fr"}
// 	masterList["2229"] = []string{"Pigulim (HR511) E4A2333-8 De Fl Lo 1/0/X Temp G K0V Fr"}
// 	masterList["2230"] = []string{"Pantethys (NSSC 2230) C86A621-9 Ni Wa 5/3/X Cold G K5V Fr"}
// 	masterList["2230a"] = []string{"Panlythos (NSSC 2230) E200300-8 Lo Va 0/0/X Roast Fr"}
// 	masterList["2321"] = []string{"Clytemnestra (NSSC 2321) D545520-6 Ag Ga Ni 7/4/+1 Temp G K4V Ov/A"}
// 	masterList["2322"] = []string{"NSSC 2322 X4A0000-0 Ba De 0/0/+1 Temp G M6V Ov/A"}
// 	masterList["2323"] = []string{"Hideout (NSSC 2323) X653100-2 Lo Po A/3/X Temp G K9V M0V Ov/A"}
// 	masterList["2324"] = []string{"New Moonsmouth (NSSC 2324) X733000-0 Uh Ba Po 9/0/X Temp M2V Ov"}
// 	masterList["2325"] = []string{"Alrai C668586-8 M S Ag Ga Ni A/4/X Cold G K1V M4V Ov"}
// 	masterList["2326"] = []string{"NSSC 2326 X100000-0 Ba Va 0/0/X Temp G M1V Ov/A"}
// 	masterList["2328"] = []string{"54 Piscium I X8A4000-0 Uh Ba Fl A/1/-1 Hot G K0V BD Ov/A"}
// 	masterList["2329"] = []string{"Jiaolong (NSSC 2329) E85A411-7 Lt Ni Wa A/8/X Temp G K2V M3V Ov"}
// 	masterList["2423"] = []string{"Nantosuelta (51 Pegasi) D684574-8 S Ag Ga Ni A/7/+2 Temp G G2IV Ov"}
// 	masterList["2426"] = []string{"Cerberus (85 Pegasi) E789285-8 Lo 3/8/X Temp G G5V K7V Ov"}
// 	masterList["2427"] = []string{"Glenn's Memorial (NSSC 2427) C310441-8 M S Ni 0/0/+2 Temp G K5V K7V Ov"}
// 	masterList["2430"] = []string{"Chantico (NSSC 2430) E696252-7 Ga Lo 3/0/X Hot G K9V Ov"}
// 	masterList["2521"] = []string{"Shalhevet C200231-9 Ba Va 0/0/-1 Roast G K1V M2V Ov"}
// 	masterList["2524"] = []string{"Iota Piscium D310132-8 S Lo 0/0/X Temp F7V Ov"}
// 	masterList["2525"] = []string{"Pirithous (NSSC 2525) E957220-5 Ga Lo Lt 3/7/X Temp G K7V Ov/A"}
// 	masterList["2526"] = []string{"NSSC 2526 X100000-0 Ba Va 0/0/+1 Temp G K2V M4V M6V Ex"}
// 	masterList["2529"] = []string{"Heraclitus (NSSC 2529) C3A0421-9 S De Ni 2/0/-3 Frozen G BD Ov"}
// 	masterList["2621"] = []string{"Hypatia (NSSC 2621) C747553-9 R Ag Ga Ni 8/3/+1 Temp G K5V Ov"}
// 	masterList["2622"] = []string{"Spinoza (NSSC 2622) B667553-A M N S Ag Ga Ni A/7/X Temp G K7V M3V Ov"}
// 	masterList["2624"] = []string{"NSSC 2624 X401000-0 Ba Ic Va 0/0/X Temp M7V Ex"}
// 	masterList["2629"] = []string{"NSSC 2629 X4A0000-0 Ba De 5/0/X Hot G M1V BD Ov"}
// 	masterList["2721"] = []string{"Pythagoras (NSSC 2721) C9D8596-6 Fl Ni 7/1/+0 Temp G K8V M1V Ov"}
// 	masterList["2722"] = []string{"Cunitz (NSSC 2722) E000100-9 As Lo Va 0/0/+0 Temp G M5V Ov"}
// 	masterList["2726"] = []string{"Abnoba (NSSC 2726) E557283-8 Ga Lo 9/3/+0 Cold G K6V Ov"}
// 	masterList["2727"] = []string{"Mush ki (GL3021) D695231-7 S Ga Lo 8/6/X Hot G G6V Ov"}
// 	masterList["2728"] = []string{"Beta Ceti E569242-2 Lo Lt 2/4/X Temp G K0III Ov"}
// 	masterList["2730"] = []string{"Kropotkin (NSSC 2730) D310420-8 Ua Ni 0/0/X Temp G M1V M8V Ov/A"}
// 	masterList["2821"] = []string{"NSSC 2821 X310000-0 Ba 0/0/-1 Hot G M8V Ex"}
// 	masterList["2825"] = []string{"NSSC 2825 XA8A100-1 Lo Lt Wa 6/3/X Temp K3V M8V Ex"}
// 	masterList["2829"] = []string{"Aesara (NSSC 2829) C683474-9 Ni 6/9/+0 Roast G K7V Ov/A"}
// 	masterList["2830"] = []string{"Tito (NSSC 2830) C66A484-8 Ni Wa 5/5/X Temp G K2V Ov"}
// 	masterList["2921"] = []string{"NSSC 2921 X87A000-0 Ba Wa 4/4/+1 Temp G K4V Ex"}
// 	masterList["2922"] = []string{"NSSC 2922 XA7A000-0 Ba Wa A/0/X Cold G K7V M3V Ex"}
// 	masterList["2924"] = []string{"NSSC 2924 X4A0000-0 Ba De 5/0/X Temp G M5V BD Ex"}
// 	masterList["2925"] = []string{"NSSC 2925 X500000-0 Ba Va 0/0/X Temp G K2V Ex"}
// 	masterList["2926"] = []string{"NSSC 2926 X674000-0 Ba Ga 9/1/+2 Temp G K9V Ex"}
// 	masterList["2927"] = []string{"NSSC 2927 X310000-0 Ba 0/0/-4 Frozen BD Ex"}
// 	masterList["2928"] = []string{"Teach's Folly (NSSC 2928) C721220-7 P De Po 2/0/X Frozen M2V Ex"}
// 	masterList["2929"] = []string{"NSSC 2929 X310000-0 Ba 0/0/X Hot G M3V Ex"}
// 	masterList["3021"] = []string{"NSSC 3021 X410000-0 Ba 0/0/X Hot G M7V Ex"}
// 	masterList["3022"] = []string{"NSSC 3022 X695000-0 Ba Ga 3/3/+2 Cold G K6V M9V Ex"}
// 	masterList["3024"] = []string{"NSSC 3024 X3A0000-0 Ba De 0/0/+1 Hot G M6V M8V Ex"}
// 	masterList["3025"] = []string{"NSSC 3025 X200000-0 Ba Va 0/0/-1 Roast G M2V M7V Ex"}
// 	masterList["3026"] = []string{"Alpha Phoenicis X683000-0 Ba 6/0/-1 Hot G K0III Ex"}
// 	masterList["3027"] = []string{"NSSC 3027 X856000-0 Ba Ga 8/4/-1 Temp G K3V M4V Ex"}
// 	masterList["3029"] = []string{"Epiphany (NSSC 3029) X76A233-1 Lo Lt Wa 5/8/+0 Cold G K5V K9V Ex"}
// 	masterList["3121"] = []string{"Newer Lanark (NSSC 3121) X310000-0 Uh Ba 0/0/X Temp K8V M0V Ex"}
// 	masterList["3122"] = []string{"NSSC 3122 X568000-0 Ba Ga 4/4/X Temp G M1V BD Ex"}
// 	masterList["3125"] = []string{"NSSC 3125 X994000-0 Ba Ga 8/3/X Roast G M1V Ex"}
// 	masterList["3126"] = []string{"HD142 X992000-0 Ba 6/4/X Roast G G1IV Ex"}
// 	masterList["3127"] = []string{"NSSC 3127 X685000-0 Ba Ga 7/4/+0 Hot K7V M3V Ex"}
// 	masterList["3128"] = []string{"NSSC 3128 X310000-0 Ba 0/0/X Temp G K7V M0V Ex"}
// 	masterList["3222"] = []string{"Alpha Pegasi X000000-0 Ua As Ba Va 0/0/+2 Roast A0IV Ex"}
// 	masterList["3225"] = []string{"NSSC 3225 X100000-0 Ba Va 0/0/X Frozen G K8V M9V Ex"}
// 	masterList["3229"] = []string{"GL9028 X98A000-0 Ba Wa A/8/+0 Temp G G5V Ex"}
// 	masterList["0134"] = []string{"36 Ursae Majoris X400000-0 Ba Va 0/0/+0 Hot G F8V Ex"}
// 	masterList["0135"] = []string{"NSSC 0135 X300000-0 Ba Va 0/0/+0 Temp G K7V Ex"}
// 	masterList["0136"] = []string{"Nueva Catalonia (NSSC 0136) C410420-9 S Ni 0/0/X Temp G K7V BD Ov"}
// 	masterList["0137"] = []string{"NSSC 0137 X200000-0 Ba Va 0/0/X Temp G M2V Ex"}
// 	masterList["0231"] = []string{"Novi Magnitogorsk (20 Leonis Minoris) B663584-A N Ni 9/7/+1 Temp G G1V Fr"}
// 	masterList["0234"] = []string{"Maximon (NSSC 0234) C751574-9 S Ni Po 4/6/+1 Hot G K5V Ov"}
// 	masterList["0236"] = []string{"Liberta (Iota Ursae Majoris) D310300-9 Lo 0/0/+0 Temp G A7IV Ov"}
// 	masterList["0238"] = []string{"NSSC 0238 X200000-0 Ba Va 0/0/X Temp G M4V BD Ex"}
// 	masterList["0239"] = []string{"Beta Carina X000000-0 Ba Va 0/0/X N/A A1III Ex"}
// 	masterList["0332"] = []string{"United Harmony (Pi Ursae Majoris) C552454-6 Ni Po 9/4/X Cold G G1V Fr"}
// 	masterList["0334"] = []string{"Three Peaks (NSSC 0334) C8D3373-9 S Fl Lo 6/2/+0 Hot G K5V M4V Ov/A"}
// 	masterList["0336"] = []string{"Theta Ursae Majoris X534000-0 Ba 1/2/X Cold G F6V M6V Ex"}
// 	masterList["0339"] = []string{"NSSC 0339 X100000-0 Um Ba Va 0/0/-1 Temp G M0V M5V Ex"}
// 	masterList["0340"] = []string{"Delta Vela X000000-0 Ba Va 0/0/-1 N/A A1V Ex"}
// 	masterList["0431"] = []string{"Alula C000511-9 As Ni Va 0/0/+0 Cold F9V G3V M3V BD Fr"}
// 	masterList["0432"] = []string{"11 Leonis Minoris X200000-0 Ua Ba Va 0/0/+1 Hot G G8V M5V Ov/A"}
// 	masterList["0433"] = []string{"Liberty's Fruits (NSSC 0433) E546342-7 Ga Lo 6/3/X Temp G K4V Ov"}
// 	masterList["0434"] = []string{"NSSC 0434 X4A0000-0 Ua Ba De 0/0/-1 Temp G K8V M0V Ex"}
// 	masterList["0436"] = []string{"Schrier's Memory (47 Ursae Majoris) C666364-9 M Lo Ni 7/8/+0 Cold G G1V Ov"}
// 	masterList["0439"] = []string{"Xi Gemini X776000-0 Ba Ga 7/6/+1 Hot G F5IV Ex"}
// 	masterList["0532"] = []string{"Gliese's World (GL357) C300431-A S Ni Va 0/0/X Temp G M2V Fr"}
// 	masterList["0533"] = []string{"Lakhish (NSSC 0533) D300331-A Lo Va 0/0/X Temp M3V Fr"}
// 	masterList["0535"] = []string{"Stopover (NSSC 0535) C554231-9 Ga Lo A/5/-1 Frozen G G8V M4V Ov"}
// 	masterList["0536"] = []string{"Dunayevskaya (NSSC 0536) E665221-3 Ga Lo Lt 8/6/+0 Temp G K6V M7V Ov"}
// 	masterList["0538"] = []string{"NSSC 0538 X310000-0 Ba 0/0/X Cold M1V Ex"}
// 	masterList["0634"] = []string{"Ladon (NSSC 0634) C310283-8 Uh Lo 0/0/X Cold G M4V M7V Ov"}
// 	masterList["0635"] = []string{"NSSC 0635 X411000-0 Ba Ic 0/0/-1 Temp G M3V Ex"}
// 	masterList["0637"] = []string{"Castor X000000-0 Ba Va 0/0/X N/A A1V A2V A2V A5V Ex"}
// 	masterList["0638"] = []string{"Gamma Gemini X000000-0 Ba Va 0/0/X N/A A1IV Ex"}
// 	masterList["0731"] = []string{"Novi Kerch (SFT1321) C949485-8 M Ni 5/2/+1 Temp K7V M0V If"}
// 	masterList["0733"] = []string{"Black Winds (NSSC 0733) C310496-8 Ni 0/0/+0 Temp G M7V Fr"}
// 	masterList["0734"] = []string{"Mutual Balance (NSSC 0734) E99A222-6 Lo Wa 7/1/X Temp G K4V Fr"}
// 	masterList["0736"] = []string{"Echidna (GL302) E578212-6 Ga Lo 6/0/X Temp G G5V Ov"}
// 	masterList["0739"] = []string{"Delta Gemini VI E5322A8-7 Lo Po 1/2/+1 Temp F0IV Ov/A"}
// 	masterList["0831"] = []string{"Mat Zemlya (LHS2090) B677786-B M N S Ag Ga 9/5/+2 Temp M6V If"}
// 	masterList["0831a"] = []string{"Zilyonaya Voda (LHS2090) D5AA566-B Fl Ni 3/1/X Frozen If"}
// 	masterList["0835"] = []string{"Pollux E84A200-9 Lo Wa 3/0/X Temp G K0III Ov/A"}
// 	masterList["0838"] = []string{"Risar (NSSC 0838) E200272-9 Ux Lo Va 0/0/X Hot G M2V M9V Ov/A"}
// 	masterList["0839"] = []string{"NSSC 0839 X6A0000-0 Ba De 0/0/+0 Temp G M7V Ex"}
// 	masterList["0932"] = []string{"Tacitus (Ross 619) C200595-B R Ni Va 0/0/X Temp G M4V If"}
// 	masterList["0937"] = []string{"NSSC 0937 X633000-0 Ba Po 6/6/+0 Temp G K3V Ex"}
// 	masterList["0940"] = []string{"NSSC 0940 X8A5000-0 Ba Fl 6/4/X Temp M1V Ex"}
// 	masterList["1031"] = []string{"Orthrus (YZ Canis Minoris) B636685-B N S Na Ni 4/4/X Temp G M4V If"}
// 	masterList["1031a"] = []string{"Erytheia (YZ Canis Minoris) C301365-B Lo Ic 0/0/X Temp If"}
// 	masterList["1033"] = []string{"Tomorrow's Promise (GL293) B412658-B Ic Na Ni 0/0/-3 Frozen G DQ9 If"}
// 	masterList["1037"] = []string{"Xenophon (NSSC 1037) C556412-7 Ga Ni 7/3/+0 Temp G K6V Ov"}
// 	masterList["1038"] = []string{"Beta Aurigae E510286-8 Lo 0/0/X Temp G A7V Ov"}
// 	masterList["1039"] = []string{"Klipa (NSSC 1039) C300111-9 R Lo Va 0/0/X Temp G M3V Ov"}
// 	masterList["1040"] = []string{"Alpha Pictoris X400000-0 Ba Va 0/0/X Temp A6V Ex"}
// 	masterList["1132"] = []string{"New Detroit (Wolf 294) A558787-B M N S Ag Ga 7/4/X Temp G M3V If"}
// 	masterList["1133"] = []string{"Aurigae Facility (QY Aurigae) C200396-9 S R Um Lo Va 0/0/X Temp G M4V M9V If"}
// 	masterList["1135"] = []string{"Erebus (GL250) D6A5552-8 S Fl Ni 9/0/X Temp M2V M3V Fr"}
// 	masterList["1138"] = []string{"Purity (NSSC 1138) DADA433-9 S Fl Ni 1/1/+0 Temp G K7V M3V Ov/A"}
// 	masterList["1140"] = []string{"NSSC 1140 X6A0000-0 Ba De 2/0/+0 Roast G M9V BD Ex"}
// 	masterList["1236"] = []string{"Atra-Hasis (NSSC 1236) D400654-8 S Na Ni Va 0/0/X Temp G K3V Fr"}
// 	masterList["1237"] = []string{"Alpha Mensae E577441-7 P Ga Ni 6/5/X Temp G G5V Fr/A"}
// 	masterList["1333"] = []string{"New Ballarat (GL233.2) C510577-A S Ni 0/0/X Frozen G WD Fr"}
// 	masterList["1335"] = []string{"Chi Orionis E000384-8 As Lo Va 0/0/+2 Temp G0V M6V Fr"}
// 	masterList["1338"] = []string{"Eta Leporis X200000-0 Ba Va 0/0/X Cold G F1V Ov/A"}
// 	masterList["1339"] = []string{"Maayanot (NSSC 1339) D883253-3 Lo Lt 5/9/X Hot G K2V Ov"}
// 	masterList["1431"] = []string{"Maat (Luyten 1723) D200413-A S Ni Va 0/0/X Temp G M4V If"}
// 	masterList["1434"] = []string{"Gamma Leporis D669596-4 Lt Ni 6/3/X Temp G F7V K2V Fr"}
// 	masterList["1437"] = []string{"Novi Krondstadt (111 Tauri) E747200-4 Ga Lo Lt 6/7/+0 Temp G F8V Ov/A"}
// 	masterList["1438"] = []string{"Strange Echoes (NSSC 1438) X201000-0 Uh Ba Va 0/0/X Temp G M4V BD Ov/A"}
// 	masterList["1440"] = []string{"Pi Mensae X797000-0 Ba Ga A/0/X Temp G G1IV Ex"}
// 	masterList["1531"] = []string{"Bad Vibe (Stein 2051) C513598-9 Ic Ni 0/0/X Frozen G M4V WD If"}
// 	masterList["1534"] = []string{"Gernsback (GL183) B732551-9 Ni Po 7/6/+0 Cold G K4III Fr"}
// 	masterList["1537"] = []string{"Capella VII C896510-8 Ag Ga Ni 4/6/+2 Hot G0III G5III M2V M4V Ov"}
// 	masterList["1538"] = []string{"NSSC 1538 X4A0000-0 Ba De 0/0/+1 Hot G K4V Ex"}
// 	masterList["1632"] = []string{"Bolotnikov (Pi 3 Orionis) C661541-8 R Ua Ni 7/6/+2 Hot G F6V Fr"}
// 	masterList["1636"] = []string{"Lambda Aurigae III C846511-7 Ag Ga 4/3/+3 Temp G G1IV Ov"}
// 	masterList["1638"] = []string{"Speck (NSSC 1638) X100000-0 Uh Ba Va 0/0/-1 Temp M2V M3V M5V Ov/R"}
// 	masterList["1639"] = []string{"NSSC 1639 X790000-0 Ba De 8/0/X Roast G K8V Ex"}
// 	masterList["1640"] = []string{"Beta Pictoris X5A4000-0 Ba Fl 7/3/X Temp A5V Ex"}
// 	masterList["1731"] = []string{"Rana (Rana II) A777784-B M N S Ag Ga 5/5/X Cold G K0IV If"}
// 	masterList["1731a"] = []string{"Lyagushka (Rana I) C400584-B Ni Va 3/0/X Hot If"}
// 	masterList["1734"] = []string{"Nizoz (NSSC 1734) C4A0573-B De Ni 2/0/X Temp K4V M0V M1V Fr"}
// 	masterList["1735"] = []string{"Panta Rhei (NSSC 1735) E652273-8 Lo Po 1/6/X Temp K9V Ov/A"}
// 	masterList["1738"] = []string{"Wisp (NSSC 1738) D310220-9 Lo 0/0/X Frozen G M8V Ov/A"}
// 	masterList["1832"] = []string{"New Salvation (NSSC 1832) C4A1595-A S Fl Ni 0/0/X Hot G K5V Fr"}
// 	masterList["1833"] = []string{"Democritus (NSSC 1833) B300554-B M N Ni Va 0/0/-4 Hot G M6V BD Fr"}
// 	masterList["1836"] = []string{"Novi Voronezh (NSSC 1836) E624231-7 Lo 0/0/X Cold G M1V Ov"}
// 	masterList["1837"] = []string{"Freeman's Belt (58 Eridani) E000221-8 P As Lo Va 0/0/+3 Cold G3V Ov/A"}
// 	masterList["1838"] = []string{"NSSC 1838 X7B0000-0 Ba De 0/0/X Roast G K6V M4V Ov/A"}
// 	masterList["1839"] = []string{"Epsilon Reticuli V C534111-9 R Lo 3/4/X Frozen G K2IV WD Ov"}
// 	masterList["1933"] = []string{"Qareen (NSSC 1933) C6B0421-9 De Ni 0/0/X Roast G M6V Fr/A"}
// 	masterList["1934"] = []string{"Unseelee (NSSC 1934) E547484-6 Ni B/3/X Cold G K1V K9V Ov"}
// 	masterList["1936"] = []string{"NSSC 1936 XAFA000-0 Ba Fl 2/0/+0 Temp G K7V M7V Ex"}
// 	masterList["1938"] = []string{"NSSC 1938 X300000-0 Ba Va 0/0/+0 Temp G K9V M2V Ex"}
// 	masterList["1940"] = []string{"Beta Eridani X520000-0 Ba De 0/0/X Cold G A3IV Ex"}
// 	masterList["2032"] = []string{"Kappa Ceti II C64A451-A Ni Wa 8/5/+1 Cold G G5V Fr"}
// 	masterList["2035"] = []string{"Istiklal (NSSC 2035) E100252-9 Lo Va 0/0/X Hot G BD Ov"}
// 	masterList["2036"] = []string{"Zipacna (NSSC 2036) D411331-8 Ic Lo 0/0/X Cold M5V Ex"}
// 	masterList["2132"] = []string{"Meng Po (NSSC 2132) D20056A-8 M S Ua Ni Va 0/0/X Temp M7V M9V Fr/R"}
// 	masterList["2133"] = []string{"Iota Persei III E97A252-4 Lo Lt Wa 7/4/X Temp G G0V Ov"}
// 	masterList["2134"] = []string{"Balor's Belt (NSSC 2134) D000231-9 As Lo Va 0/0/+2 Temp G K2V Ov"}
// 	masterList["2136"] = []string{"Ptitsa (NSSC 2136) E8A6232-8 Fl Lo 6/5/-1 Hot G M9V Ov"}
// 	masterList["2137"] = []string{"NSSC 2137 X8C4000-0 Ua Ba Fl 1/0/X Temp K6V M3V Ov/R"}
// 	masterList["2139"] = []string{"NSSC 2139 X200000-0 Ba Va 0/0/+1 Cold G K0V Ex"}
// 	masterList["2140"] = []string{"Fenja (NSSC 2140) X656000-0 Uha Ba Ga 7/4/+2 Cold G K1V Ex"}
// 	masterList["2232"] = []string{"Theta Persei VI C666532-8 Ag Ga Ni 5/8/+1 Temp G F7V M1V Ov"}
// 	masterList["2233"] = []string{"Zeta Reticuli 2 II D554411-4 Ua Ga Lt Ni 6/3/X Temp G G1V G2V Ov"}
// 	masterList["2236"] = []string{"Nightshade (NSSC 2236) E400281-8 Lo Va 0/0/X Temp G K2V Ov"}
// 	masterList["2237"] = []string{"New Canaan (NSSC 2237) D5666AA-6 Ag Ga Ni 5/3/+0 Temp G K8V Ov"}
// 	masterList["2238"] = []string{"Imprimatur (NSSC 2238) D200242-8 Lo Va 0/0/X Temp G M4V Ov"}
// 	masterList["2240"] = []string{"Aldebaran II D564430-8 P Ga Ni 9/3/+1 Temp G K5III M2V Ex"}
// 	masterList["2331"] = []string{"Buarainech (GL86) X201000-0 Uha Ic Ba Va 0/0/-1 Frozen G K1V WD Ov/A"}
// 	masterList["2332"] = []string{"Delta Trianguli IV C782531-7 Ni 8/7/+0 Hot G G0V K4V Ov"}
// 	masterList["2334"] = []string{"Skratti (NSSC 2334) X400000-0 Uh Ba Va 0/0/+0 Hot G M8V Ov/A"}
// 	masterList["2336"] = []string{"Magna Mater (10 Tauri) D662483-6 S Ni 9/4/X Roast G F9IV Ov"}
// 	masterList["2337"] = []string{"Cybele (NSSC 2337) E768551-7 Ag Ga Ni 7/3/+2 Temp G K1V Ov"}
// 	masterList["2338"] = []string{"Go bniu (NSSC 2338) E310475-9 Ni 0/0/X Temp G K6V M2V Ov/A"}
// 	masterList["2433"] = []string{"Cethlenn (NSSC 2433) E766373-7 Uh Ga Lo 6/5/+1 Hot G K9V Ov/A"}
// 	masterList["2440"] = []string{"Algol X100000-0 Ua Ba Va 0/0/+2 Cold G B8V A5V K2IV Un"}
// 	masterList["2531"] = []string{"Upsilon Andromedae V C576431-9 M S Ga Ni 5/2/X Cold G F8V Ov"}
// 	masterList["2533"] = []string{"Alpha Trianguli VIIf X300000-0 Uh Ba Va 0/0/X Temp G F6IV Ex"}
// 	masterList["2534"] = []string{"NSSC 2534 X300000-0 Ba Va 0/0/X Temp G M6V M6V Ex"}
// 	masterList["2535"] = []string{"Tau 1 Eridani IV C684211-8 Ga Lo 2/1/X Cold G F5V Ov"}
// 	masterList["2536"] = []string{"Alpha Fornacis B IV E86A49E-6 Ni Wa 7/5/+2 Temp G F8V G7V Ov/A"}
// 	masterList["2540"] = []string{"NSSC 2540 X533000-0 Ba Po 3/3/X Frozen G M3V Un"}
// 	masterList["2631"] = []string{"Nu Pheonicis III E300311-9 Lo Va 0/0/X Temp F8V Ov"}
// 	masterList["2632"] = []string{"NSSC 2632 X400000-0 Ba Va 0/0/X Temp G M1V Ex"}
// 	masterList["2640"] = []string{"NSSC 2640 X9B4000-0 Ua Ba Fl 1/6/X Temp G M6V Un"}
// 	masterList["2731"] = []string{"NSSC 2731 X100000-0 Ba Va 0/0/-2 Temp G M6V Ex"}
// 	masterList["2732"] = []string{"NSSC 2732 X100000-0 Ba Va 0/0/X Temp G K2V M2V Ov"}
// 	masterList["2733"] = []string{"Hamdir (NSSC 2633) E587284-7 Ga Lo 5/3/+0 Temp G K9V M7V Ov"}
// 	masterList["2737"] = []string{"Gamma Ceti X000000-0 Ba Va 0/0/X N/A A2V K5V Ex"}
// 	masterList["2738"] = []string{"NSSC 2738 X632000-0 Ba Po 7/7/+0 Cold M3V Ex"}
// 	masterList["2739"] = []string{"NSSC 2739 X310000-0 Ba 0/0/X Temp G M0V BD Ex"}
// 	masterList["2833"] = []string{"Porphyrion (NSSC 2833) D999331-5 S Lo Lt 8/4/+1 Temp G K2V K5V Ov"}
// 	masterList["2837"] = []string{"NSSC 2837 X410000-0 Ba 0/0/+0 Temp G M0V Ex"}
// 	masterList["2840"] = []string{"Theta Eridani X000000-0 As Ba Va 0/0/+2 Hot G A5IV Un"}
// 	masterList["2934"] = []string{"Alpha Hydri X300000-0 Ba Va 0/0/X Temp G F0V Ex"}
// 	masterList["2936"] = []string{"NSSC 2936 X200000-0 Ba Va 0/0/X Temp G M5V M7V Ex"}
// 	masterList["2938"] = []string{"NSSC 2938 X410000-0 Ua Ba 0/0/X Temp G M1V M5V Ex"}
// 	masterList["3031"] = []string{"NSSC 3031 X866000-0 Ba Ga 7/6/+0 Temp G K1V M6V Ex"}
// 	masterList["3033"] = []string{"Alcyoneus (NSSC 3033) X510000-0 Uh Ba 0/0/X Temp G M3V Ex"}
// 	masterList["3038"] = []string{"NSSC 3038 X4A0000-0 Ba De 1/3/-1 Roast G M1V M5V M9V Un"}
// 	masterList["3039"] = []string{"94 Ceti X585000-0 Ba Ga 8/2/+1 Cold G F8V M3V Un"}
// 	masterList["3131"] = []string{"NSSC 3131 X626000-0 Ba 3/1/X Cold G M2V Ex"}
// 	masterList["3132"] = []string{"NSSC 3132 X200000-0 Ba Va 0/0/-2 Temp M4V BD Ex"}
// 	masterList["3133"] = []string{"NSSC 3133 X654000-0 Ba Ga 5/3/+0 Temp G K8V Ex"}
// 	masterList["3134"] = []string{"Delta Cassiopeia X8B5000-0 Ba Fl 0/0/X Temp G A5V Ex"}
// 	masterList["3140"] = []string{"Hiisi (NSSC 3140) X310000-0 Uh Ba 0/0/-1 Temp G M1V M6V Un"}
// 	masterList["3231"] = []string{"NSSC 3231 X6B1000-0 Ua Ba Fl 1/6/X Temp G K3V M0V Ex"}
// 	masterList["3234"] = []string{"NSSC 3234 X8B3000-0 Ua Ba Fl 1/7/X Temp G K3V Ex"}
// 	masterList["3235"] = []string{"Alpha Eridani X000000-0 Ba Va 0/0/X N/A B3V Ex"}
// 	masterList["3236"] = []string{"NSSC 3236 XA9A000-0 Ba Wa 7/4/-1 Hot G K1V Ex"}
// 	masterList["3238"] = []string{"Novi Saratov (NSSC 3238) CAA0350-9 R De Lo 2/4/X Roast G K5V M7V Un"}
// 	masterList["3239"] = []string{"NSSC 3239 X673000-0 Um Ba 7/4/+0 Cold G K2V Un"}
// 	masterList["3240"] = []string{"HD20367 X545000-0 Ba Ga 7/2/+0 Temp G G0V Un"}

// 	return masterList
// }
