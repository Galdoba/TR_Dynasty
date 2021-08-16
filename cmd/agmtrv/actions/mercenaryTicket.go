package actions

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

const (
	mType_ActiveDuty         = "Active Duty"
	mType_Assasination       = "Assassination"
	mType_Commerce           = "Commerce"
	mType_Defence            = "Defence"
	mType_CounterStrike      = "Counter Strike"
	mType_Elimination        = "Elimination"
	mType_Escort             = "Escort"
	mType_FieldExercise      = "Field Exercise"
	mType_FirstStrike        = "First Strike"
	mType_Raid               = "Raid"
	mType_Recon              = "Recon"
	mType_Retrival           = "Retrieval"
	mType_Sabotage           = "Sabotage"
	mType_TechnologicalTest  = "Technological Test"
	mType_Terrorise          = "Terrorise"
	mType_Train              = "Train"
	mType_UnlawfulAquisition = "Unlawful Aquisition"
	sType_Criminal           = "Criminal"
	sType_Guerilla           = "Guerilla"
	sType_Cadre              = "Cadre"
	sType_Commando           = "Commando"
	sType_Striker            = "Striker"
	sType_Security           = "Security"
	sType_Warmonger          = "Warmonger"
	sType_Dream              = "Dream"
)

type Ticket struct {
	employer          string
	employee          string
	service           service
	preSupport        []string
	postSupport       []string
	compensation      string
	repatriation      string
	ticketAdjustments int
	patronDM          int
}

type service struct {
	genericServiceType string
	mission            string
	payGrade           int
	lenghtOfService    string
	exposure           string
	targetType         string
	targetDescriptor   string
	risk               string
}

func NewTicket() error {
	t := &Ticket{}
	t.patronDM = rollPatronDM()
	fmt.Printf("Patron DM is %v\n", t.patronDM)
	err := errors.New("Initial")
	status(t)
	t.administrationPhase()
	t.creationPhase()
	return err
}

func (t *Ticket) creationPhase() {
	t.creationStep1() //WORKING OUT THE EMPLOYER DETAILS
	t.creationStep2() //INCLUDING THE EMPLOYEE DETAILS
	t.creationStep3() //SERVICE REQUIRED
	t.creationStep4() //PRE-TICKET SUPPORT
	t.service.payGrade = t.service.payGrade + (t.ticketAdjustments / 2)
	status(t)
	os.Exit(1)
}

func (t *Ticket) creationStep1() error {
	status(t)
	t.employer = employerDetails()
	fmt.Println("Employer Details defined as: " + t.employer)

	confirmed, err := user.Confirm("Spend 1 Ticket Adjustment Point to reveal Employer Details?")
	if confirmed {
		fmt.Println("Player must pass Admin(ANY) 7+ check...")
		t.ticketAdjustments--
		fmt.Println("Roll result is: ")
		res, _ := user.InputInt()
		if res >= 7 {
			t.employer = t.employer + " [REVEALED]"
		}
	}
	return err
}
func (t *Ticket) creationStep2() error {
	status(t)
	behavour := []string{"Share info as needed", "Hide info", "Be completly honest"}
	beh, err := user.ChooseOne("Choose Mercenary unit behaviour:", behavour)
	switch beh {
	case 0:
	case 1:
		t.ticketAdjustments--
	case 2:
		t.ticketAdjustments++
	}
	t.employee = behavour[beh]
	return err
}
func (t *Ticket) creationStep3() error {
	status(t)
	t.service = service{}
	t.genericServiceType()
	switch t.service.genericServiceType {
	default:
		user.Confirm(t.service.genericServiceType + " missions not implemented!")
	case sType_Criminal:
		t.criminalTickets()
	case sType_Guerilla:
		t.guerillaTickets()
	case sType_Cadre:
		t.cadreTickets()
	case sType_Commando:
		t.commandoTickets()
	case sType_Striker:
		t.strikerTickets()
	case sType_Security:
		t.securityTickets()
	case sType_Warmonger:
		t.warmongerTickets()
	case sType_Dream:
		t.dreamTickets()
	}
	status(t)
	t.creationStep3b() //LENGTH OF SERVICE
	t.creationStep3c() //TICKET EXPOSURE
	t.creationStep3d() //DETERMINE TARGET
	t.creationStep3e() //DETERMINE TARGET DESCRIPTOR
	t.creationStep3f() //DETERMINE RISK
	fmt.Println(defineMission(t))
	return nil
}

func (t *Ticket) creationStep3b() {
	status(t)
	t.lenghtOfService()
	t.rollLoS()
}

func (t *Ticket) creationStep3c() {
	status(t)
	printExposureMissionsTable()
	dm := 0
	switch t.service.genericServiceType {
	case sType_Criminal:
		dm = -3
	case sType_Guerilla:
		dm = -2
	case sType_Warmonger:
		dm = -1
	case sType_Dream:
		dm = 2
	}
	r := dice.Roll2D()
	rollValue := r + dm
	fmt.Printf("Result is %v\nThe mercenary administrator can spend Ticket Adjustments to raise or lower the result by +/– 1 per Ticket Adjustment.\n", rollValue)
	if userConfirmed("Adjust Exposure?") {
		valid := false
		for !valid {
			fmt.Println("Enter New Value:")
			userValue, err := user.InputInt()
			if err != nil {
				fmt.Println("Error: " + err.Error())
				continue
			}
			if difference(rollValue, userValue) > t.ticketAdjustments {
				fmt.Printf("Error: not enough Ticket Adjustment Points to set value as %v\n", userValue)
				continue
			}
			t.ticketAdjustments = t.ticketAdjustments - difference(rollValue, userValue)
			rollValue = userValue
			valid = true
		}
	}
	t.service.exposure = exposureTable(rollValue)
}

func (t *Ticket) creationStep3d() {
	status(t)
	targetList := []string{}
	switch t.service.mission {
	default:
		panic("Unknown mission " + t.service.mission)
	case mType_ActiveDuty:
		targetList = append(targetList, "d")
		targetList = append(targetList, "o")
	case mType_Defence:
		targetList = append(targetList, "d")
		targetList = append(targetList, "n")
	case mType_Escort:
		targetList = append(targetList, "d")
	case mType_Retrival:
		targetList = append(targetList, "d")
		targetList = append(targetList, "n")
	case mType_Train:
		targetList = append(targetList, "d")
	case mType_Commerce:
		targetList = append(targetList, "n")
	case mType_FieldExercise:
		targetList = append(targetList, "n")
		targetList = append(targetList, "o")
	case mType_Recon:
		targetList = append(targetList, "n")
		targetList = append(targetList, "o")
	case mType_TechnologicalTest:
		targetList = append(targetList, "n")
	case mType_Assasination: //SPECIAL CASE
		t.service.targetType = "Individual"
		t.service.payGrade++
		return
	case mType_CounterStrike:
		targetList = append(targetList, "o")
	case mType_Elimination:
		targetList = append(targetList, "o")
	case mType_FirstStrike:
		targetList = append(targetList, "o")
	case mType_Raid:
		targetList = append(targetList, "o")
	case mType_Sabotage:
		targetList = append(targetList, "o")
	case mType_Terrorise:
		targetList = append(targetList, "o")
	case mType_UnlawfulAquisition:
		targetList = append(targetList, "o")
	}
	target := dice.New().RollFromList(targetList)
	rollValue := dice.Roll1D()
	valid := false
	targ := ""
	payInc := 0
	for !valid {
		switch target {
		case "d":
			printDefensiveTargetTypes()
			i := selectFromTable(t, rollValue, 1, 1, 6)
			targ, payInc = defensiveTargetTypes(i)
		case "n":
			printNeutralTargetTypes()
			i := selectFromTable(t, rollValue, 1, 1, 6)
			targ, payInc = neutralTargetTypes(i)
		case "o":
			printOffensiveTargetTypes()
			i := selectFromTable(t, rollValue, 1, 1, 6)
			targ, payInc = offensiveTargetTypes(i)
		}
		if strings.Contains(targ, "Error") {
			fmt.Println(targ)
			continue
		}
		t.service.targetType = targ
		t.service.payGrade = t.service.payGrade + payInc
		valid = true
	}
	status(t)
}

func (t *Ticket) creationStep3e() {
	done := false
	mobile := false
	alien := false
	desc := ""
	for !done {
		switch dice.Roll1D() {
		case 1:
			desc += "Political"
		case 2:
			desc += "Military"
		case 3:
			desc += "Civilian"
		case 4:
			desc += "Commercial"
		case 5:
			mobile = true
			continue
		case 6:
			alien = true
			continue
		}
		done = true
	}
	if alien {
		desc = "Alien " + desc
	}
	if mobile {
		desc = "Mobile " + desc
	}
	t.service.targetDescriptor = desc
	status(t)
}

func (t *Ticket) creationStep3f() {
	rollResult := dice.Roll1D()
	switch rollResult {
	case 1:
		t.service.risk = "Too Easy – This is well beneath the unit’s level of training; it is unlikely they will even break a sweat."
	case 2:
		t.service.risk = "Easy – This ticket will not cost the unit much in the way of resources or stress."
	case 3:
		t.service.risk = "Average – This is what the unit is trained for, and should serve as a good reminder what ticket work should be."
	case 4:
		t.service.risk = "Worthy Test – This is a fantastic place to test the unit’s skills, even some of the obscure ones. They might suffer some wounds or even casualties."
	case 5:
		t.service.risk = "Difficult – This ticket will be a tough one for the whole unit, and the members will need to be diligent in their training or they might not make it back home."
	case 6:
		t.service.risk = "Arduous – This mission is a nightmare. If anyone makes it back in one piece, they will have been pushed to the very limit."
	}
	t.service.payGrade += (rollResult - 3)
	///////////////////////////////////////
	if t.ticketAdjustments > 0 {
		if userConfirmed("Reveal Mission Risk? (cost 1 Adjustment point)") {
			t.service.risk += " [REVEALED]"
			t.ticketAdjustments--
		}
	}
	status(t)
}

func (t *Ticket) creationStep4() {
	status(t)
	fmt.Printf("//PRE-TICKET SUPPORT:\n")
	if t.service.mission == mType_TechnologicalTest {
		t.applyPreTicketSupport(2)
		t.service.payGrade = t.service.payGrade + 3
	}
	fmt.Printf("The mercenary Administrator can refuse from Pre-Ticket Support (will increase Pay Grade by 1)\n")
	if userConfirmed("Waive Pre-Ticket support?") {
		t.service.payGrade++
		status(t)
		return
	}
	tablesRolled := []bool{false, false, false} //funds = 0, Services = 1, Equipment = 2
	options := []string{"Roll Broker 9+ to select single support type", "spend x Adjustment point for each support type (maximum 3)"}
	selected, _ := user.ChooseOne("Select negotiations type:", options)

	switch selected {
	case 0:
		fmt.Println("Roll Broker (ANY) 9+ check. Result: ")
		rollResult, _ := user.InputInt()
		if rollResult < 9 {
			return
		}
		fmt.Println("Select Support Type:")
		r, _ := user.ChooseOne("Select negotiations type:", []string{"Advance Funds", "Services", "Equpment"})
		t.applyPreTicketSupport(r)
	case 1:
		uInp := userInputIntBounded(fmt.Sprintf("Select how many Ticket Adjustment Points to spend (0-%v):\n ", utils.Min(3, t.ticketAdjustments)), 0, utils.Min(3, t.ticketAdjustments))
		for i := 0; i < uInp; i++ {
			r := dice.Roll1D() % 3
			if tablesRolled[r] {
				i--
				continue
			}
			tablesRolled[r] = true
			t.applyPreTicketSupport(r)
		}
	}

}

func (t *Ticket) applyPreTicketSupport(r int) {
	switch r {
	case 0:
		t.service.payGrade--
		status(t)
		printSupportAdvanceFunds()
		t.preSupport = append(t.preSupport, supportAdvanceFundsTable(dice.Roll1D()))
	case 1:
		t.service.payGrade--
		status(t)
		printSupportServices()
		t.preSupport = append(t.preSupport, supportServicesTable(dice.Roll1D()))
	case 2:
		t.service.payGrade = t.service.payGrade - 3
		status(t)
		printSupportEquipment()
		t.preSupport = append(t.preSupport, supportEquipmentTable(dice.Roll1D()))
	}
}

func userInputIntBounded(descr string, min, max int) int {
	fmt.Println(descr)
	for {
		uInp, err := user.InputInt()
		if err != nil {
			fmt.Println("Error: " + err.Error())
			continue
		}
		if uInp < min || uInp > max {
			fmt.Printf("Error: cannot assign value %v (only values between %v and %v are valid)\n", uInp, min, max)
			continue
		}
		return uInp
	}
}

func userConfirmed(question string) bool {
	err := errors.New("Initial")
	answer := false
	for err != nil {
		answer, err = user.Confirm(question)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			continue
		}
	}
	return answer
}

func difference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (t *Ticket) lenghtOfService() {
	length := ""
	switch t.service.mission {
	case mType_Assasination, mType_CounterStrike, mType_Elimination, mType_FirstStrike, mType_Raid, mType_Sabotage, mType_UnlawfulAquisition:
		length = "SHORT"
	case mType_Commerce, mType_Defence, mType_Escort, mType_FieldExercise, mType_Recon, mType_Retrival, mType_Terrorise, mType_Train:
		length = "MEDIUM"
	case mType_ActiveDuty, mType_TechnologicalTest:
		length = "LONG"
	default:
		panic(t.service.mission + " mission lenght NOT DEFINED")
	}
	t.service.lenghtOfService = length
	if t.ticketAdjustments < 3 {
		return
	}
	mustChange, _ := user.Confirm("This mission is expected to be " + length + ". It will take 3 Ticket Adjustment Points to pick other lenght. Pick other?")
	if mustChange {
		newLen, _ := user.ChooseOneStr("Pick Mission Lenght Table:", []string{"SHORT", "MEDIUM", "LONG"})
		if newLen != length {
			t.ticketAdjustments = t.ticketAdjustments - 3
		}
		t.service.lenghtOfService = newLen
	}
}

func (t *Ticket) rollLoS() {
	patronOffer := dice.Roll1D()
	list := []string{}
	switch t.service.lenghtOfService {
	default:
		panic("Unknown LoS")
	case "SHORT":
		printSHORTlMissionsTable()
		list = []string{"1d6 Days", "1d6 Days", "2d6 Days", "2d6 Days", "1d6 Weeks", "1d6 Weeks", "1d6+2 Weeks"}
	case "MEDIUM":
		printMEDIUMMissionsTable()
		list = []string{"1d6 Weeks", "1d6+1 Weeks", "2d6 Weeks", "1d6 Months", "1d6+1 Months", "2d6 Months", "2d6+1 Months"}
	case "LONG":
		printLONGMissionsTable()
		list = []string{"1d6+1 Months", "2d6 Months", "2d6+1 Months", "3d6 Months", "3d6+2 Months", "4d6 Months", "1d6 Years"}
	}
	chosens := selectFromTable(t, patronOffer, 1, 1, 7)
	t.service.lenghtOfService = list[chosens-1]
}

func (t *Ticket) genericServiceType() {
	printGenericServiceTypeTable()
	patronOffer := dice.Roll2D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 13)
	gst := []string{sType_Criminal, sType_Guerilla, sType_Cadre, sType_Commando, sType_Striker, sType_Security, sType_Warmonger, sType_Dream}
	switch chosen {
	case 1:
		t.service.genericServiceType = gst[0]
	case 2:
		t.service.genericServiceType = gst[1]
	case 3, 4:
		t.service.genericServiceType = gst[2]
	case 5, 6:
		t.service.genericServiceType = gst[3]
	case 7, 8, 9:
		t.service.genericServiceType = gst[4]
	case 10, 11:
		t.service.genericServiceType = gst[5]
	case 12:
		t.service.genericServiceType = gst[6]
	case 13:
		t.service.genericServiceType = gst[7]
	default:
		panic("Invalid result from selectFromTable() - genericServiceType()")
	}

	status(t)
}

func selectFromTable(t *Ticket, patronOffer, payMult, minVal, maxVal int) int {
	agreedUpon := false
	chosen := 0
	err := errors.New("Init")
	fmt.Printf("Предложение нанимателя: %v\n", patronOffer)
	for !agreedUpon {

		fmt.Println("Выберите конечный результат (ВНИМАНИЕ! -1 к Ticket Adjustment Point за каждый отступ от предложения нанимателя):")
		chosen, err = user.InputInt()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		min := utils.Min(patronOffer, chosen)
		max := utils.Max(patronOffer, chosen)
		tapCost := (max - min) * payMult
		if t.ticketAdjustments-tapCost < 0 {
			err = fmt.Errorf("Not enough Ticket Adjustment points (have %v, need %v for this agreement)", t.ticketAdjustments, tapCost)
			fmt.Println(err.Error())
			continue
		}
		if chosen < minVal {
			err = fmt.Errorf("Invalid selection (have %v, need >= %v for agreement)", chosen, minVal)
			fmt.Println(err.Error())
			continue
		}
		if chosen > maxVal {
			err = fmt.Errorf("Invalid selection (have %v, need >= %v for agreement)", chosen, maxVal)
			fmt.Println(err.Error())
			continue
		}
		t.ticketAdjustments = t.ticketAdjustments - tapCost
		agreedUpon = true
	}
	return chosen
}

func (t *Ticket) criminalTickets() {
	printCriminalMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_Assasination, mType_Raid, mType_Raid, mType_Raid, mType_Sabotage, mType_UnlawfulAquisition}
	smPayment := []int{7, 3, 4, 5, 3, 6}
	switch chosen {
	default:
		panic("Criminal Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) guerillaTickets() {
	printGuerillaMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_Sabotage, mType_Sabotage, mType_Terrorise, mType_Assasination, mType_Recon, mType_FirstStrike}
	smPayment := []int{3, 4, 6, 6, 4, 5}
	switch chosen {
	default:
		panic("Guerilla Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) cadreTickets() {
	printCadreMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_Train, mType_Train, mType_FieldExercise, mType_FieldExercise, mType_ActiveDuty, mType_Recon}
	smPayment := []int{3, 4, 4, 5, 6, 5}
	switch chosen {
	default:
		panic("Guerilla Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) commandoTickets() {
	printCommandoMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_FirstStrike, mType_Raid, mType_ActiveDuty, mType_ActiveDuty, mType_Retrival, mType_Elimination}
	smPayment := []int{5, 4, 6, 7, 5, 6}
	switch chosen {
	default:
		panic("Guerilla Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) strikerTickets() {
	printStrikerMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_CounterStrike, mType_Recon, mType_FirstStrike, mType_FirstStrike, mType_Elimination, mType_Elimination}
	smPayment := []int{7, 3, 5, 6, 5, 6}
	switch chosen {
	default:
		panic("Guerilla Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) securityTickets() {
	printSecurityMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_Defence, mType_Defence, mType_Defence, mType_ActiveDuty, mType_Escort, mType_Escort}
	smPayment := []int{3, 4, 5, 6, 4, 5}
	switch chosen {
	default:
		panic("Security Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) warmongerTickets() {
	printWarmongerMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_Escort, mType_Escort, mType_FieldExercise, mType_Commerce, mType_Commerce, mType_Raid}
	smPayment := []int{4, 5, 5, 6, 7, 5}
	switch chosen {
	default:
		panic("Security Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func (t *Ticket) dreamTickets() {
	printDreamMissionsTable()
	patronOffer := dice.Roll1D()
	chosen := selectFromTable(t, patronOffer, 1, 1, 6)
	sm := []string{mType_Recon, mType_Escort, mType_FirstStrike, mType_FieldExercise, mType_Elimination, mType_TechnologicalTest}
	smPayment := []int{6, 7, 8, 6, 9, 6}
	switch chosen {
	default:
		panic("Dream Mission creation Failed")
	case 1, 2, 3, 4, 5, 6:
		t.service.mission = sm[chosen-1]
		t.service.payGrade = smPayment[chosen-1]
	}
}

func printGenericServiceTypeTable() {
	fmt.Println("Generic Service Type table:")
	fmt.Println("---------------------------")
	fmt.Println("1	Criminal")
	fmt.Println("2	Guerilla")
	fmt.Println("3	Cadre")
	fmt.Println("4	Cadre")
	fmt.Println("5	Commando")
	fmt.Println("6	Commando")
	fmt.Println("7	Striker")
	fmt.Println("8	Striker")
	fmt.Println("9	Striker")
	fmt.Println("10	Security")
	fmt.Println("11	Security")
	fmt.Println("12	Warmonger")
	fmt.Println("13	Dream")
	fmt.Println("---------------------------")
}

func printSHORTlMissionsTable() {
	fmt.Println("SHORT Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Ticket Service Lenght")
	fmt.Println("1   1d6 Days")
	fmt.Println("2   1d6 Days")
	fmt.Println("3   2d6 Days")
	fmt.Println("4   2d6 Days")
	fmt.Println("5   1d6 Weeks")
	fmt.Println("6   1d6 Weeks")
	fmt.Println("7   1d6+2 Weeks")
	fmt.Println("------------------------------------")
}

func printMEDIUMMissionsTable() {
	fmt.Println("MEDIUM Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Ticket Service Lenght")
	fmt.Println("1   1d6 Weeks")
	fmt.Println("2   1d6+1 Weeks")
	fmt.Println("3   2d6 Weeks")
	fmt.Println("4   1d6 Months")
	fmt.Println("5   1d6+1 Months")
	fmt.Println("6   2d6 Months")
	fmt.Println("7   2d6+2 Months")
	fmt.Println("------------------------------------")
}

func printLONGMissionsTable() {
	fmt.Println("LONG Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Ticket Service Lenght")
	fmt.Println("1   1d6+1 Months")
	fmt.Println("2   2d6 Months")
	fmt.Println("3   2d6+1 Months")
	fmt.Println("4   3d6 Months")
	fmt.Println("5   3d6+2 Months")
	fmt.Println("6   4d6 Months")
	fmt.Println("7   1d6 Years")
	fmt.Println("------------------------------------")
}

func printCriminalMissionsTable() {
	fmt.Println("Criminal Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission       Compensation Grade")
	fmt.Println("1   Assassination 7     (100,000 Cr)")
	fmt.Println("2   Raid          3      (20,000 Cr)")
	fmt.Println("3   Raid          4      (30,000 Cr)")
	fmt.Println("4   Raid          5      (50,000 Cr)")
	fmt.Println("5   Sabotage      3      (20,000 Cr)")
	fmt.Println("6   Unlawful      6      (75,000 Cr)")
	fmt.Println("    Acquisition ")
	fmt.Println("------------------------------------")
}

func printGuerillaMissionsTable() {
	fmt.Println("Guerilla Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission       Compensation Grade")
	fmt.Println("1   Sabotage      3      (20,000 Cr)")
	fmt.Println("2   Sabotage      4      (30,000 Cr)")
	fmt.Println("3   Terrorise     6      (75,000 Cr)")
	fmt.Println("4   Assassination 6      (75,000 Cr)")
	fmt.Println("5   Recon         4      (30,000 Cr)")
	fmt.Println("6   First Strike  5      (50,000 Cr)")
	fmt.Println("------------------------------------")
}

func printCadreMissionsTable() {
	fmt.Println("Cadre Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission        Compensation Grade")
	fmt.Println("1   Train          3      (20,000 Cr)")
	fmt.Println("2   Train          4      (30,000 Cr)")
	fmt.Println("3   Field Exercise 4      (30,000 Cr)")
	fmt.Println("4   Field Exercise 5      (50,000 Cr)")
	fmt.Println("5   Active Duty    6      (75,000 Cr)")
	fmt.Println("6   Recon          5      (50,000 Cr)")
	fmt.Println("------------------------------------")
}

func printCommandoMissionsTable() {
	fmt.Println("Commando Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission        Compensation Grade")
	fmt.Println("1   First Strike   5	(" + paygrade2Credits(5) + ")")
	fmt.Println("2   Raid           4	(" + paygrade2Credits(4) + ")")
	fmt.Println("3   Active Duty    6	(" + paygrade2Credits(6) + ")")
	fmt.Println("4   Active Duty    7	(" + paygrade2Credits(7) + ")")
	fmt.Println("5   Retrieval      5	(" + paygrade2Credits(5) + ")")
	fmt.Println("6   Elimination    6	(" + paygrade2Credits(6) + ")")
	fmt.Println("------------------------------------")
}

func printStrikerMissionsTable() {
	fmt.Println("Striker Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission        Compensation Grade")
	fmt.Println("1   Counter Strike 7	(" + paygrade2Credits(7) + ")")
	fmt.Println("2   Recon          3	(" + paygrade2Credits(3) + ")")
	fmt.Println("3   First Strike   5	(" + paygrade2Credits(5) + ")")
	fmt.Println("4   First Strike   6	(" + paygrade2Credits(6) + ")")
	fmt.Println("5   Elimination    5	(" + paygrade2Credits(5) + ")")
	fmt.Println("6   Elimination    6	(" + paygrade2Credits(6) + ")")
	fmt.Println("------------------------------------")
}

func printSecurityMissionsTable() {
	fmt.Println("Security Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission        Compensation Grade")
	fmt.Println("1   Defence        3	(" + paygrade2Credits(3) + ")")
	fmt.Println("2   Defence        4	(" + paygrade2Credits(4) + ")")
	fmt.Println("3   Defence        5	(" + paygrade2Credits(5) + ")")
	fmt.Println("4   Active Duty    6	(" + paygrade2Credits(6) + ")")
	fmt.Println("5   Escort         4	(" + paygrade2Credits(4) + ")")
	fmt.Println("6   Escort         5	(" + paygrade2Credits(5) + ")")
	fmt.Println("------------------------------------")
}

func printWarmongerMissionsTable() {
	fmt.Println("Warmonger Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission        Compensation Grade")
	fmt.Println("1   Escort         4	(" + paygrade2Credits(4) + ")")
	fmt.Println("2   Escort         5	(" + paygrade2Credits(5) + ")")
	fmt.Println("3   Field Exercise 5	(" + paygrade2Credits(5) + ")")
	fmt.Println("4   Commerce       6	(" + paygrade2Credits(6) + ")")
	fmt.Println("5   Commerce       7	(" + paygrade2Credits(7) + ")")
	fmt.Println("6   Raid           5	(" + paygrade2Credits(5) + ")")
	fmt.Println("------------------------------------")
}

func printDreamMissionsTable() {
	fmt.Println("Security Missions table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Mission        Compensation Grade")
	fmt.Println("1   Recon              6	(" + paygrade2Credits(6) + ")")
	fmt.Println("2   Escort             7	(" + paygrade2Credits(7) + ")")
	fmt.Println("3   First Strike       8	(" + paygrade2Credits(8) + ")")
	fmt.Println("4   Field Exercise     6	(" + paygrade2Credits(6) + ")")
	fmt.Println("5   Elimination        9	(" + paygrade2Credits(9) + ")")
	fmt.Println("6   Technological Test 6	(" + paygrade2Credits(6) + ")")
	fmt.Println("------------------------------------")
}

func printExposureMissionsTable() {
	fmt.Println("Public Exposure table:")
	fmt.Println("------------------------------------")
	fmt.Println("#   Result        Compensation Grade")
	fmt.Println("2-           Hidden – Without doing research, no one knows the ticket existed.")
	fmt.Println("3 - 4        Obscure – Only the local public is aware of the ticket’s actions.")
	fmt.Println("5 - 6        Low Profile – Much of the planet is aware of the ticket’s actions, and the local public know the name of the mercenary unit.")
	fmt.Println("7 - 8        Uncommon – The ticket has received some media attention in the local area and the mercenary unit’s involvement is locally public.")
	fmt.Println("9 - 10       Common – The media has spread the mercenary unit’s name throughout the planet and it has spread to neighbouring planets.")
	fmt.Println("11 - 12      Exposed – The mercenary unit’s name is publicly known on a planetary level; even a few specific members’ names are being said.")
	fmt.Println("13+          High Profile – The ticket and the mercenary unit are being talked about throughout the system. At least one member of the unit is being named specifically.")
	fmt.Println("------------------------------------")
}

func (t *Ticket) administrationPhase() {
	dp := dice.New()
	fmt.Println("The Offer:")
	fmt.Println("A mercenary administrator sits down to judge the opening offer of a ticket – to ensure it is good enough, but not too good – he must throw his Admin skill 8+. Depending on the Effect of the throw, the mercenary can have a good deal of sway over the honesty of the Ticket Offer segment of a ticket (see Creating the Ticket below). The degrees of Effect on this throw are below.")
	fmt.Println("----------------")
	fmt.Println("Effect		Ticket Offer Adjustment")
	fmt.Println("1 or lower	-4 DM")
	fmt.Println("2		-2 DM")
	fmt.Println("3		+0 DM")
	fmt.Println("4		+0 DM")
	fmt.Println("5		+1 DM")
	fmt.Println("6 or more	+2 DM")
	fmt.Println("----------------")
	fmt.Println("проверка Adnin проив сложности 8\nЭффект:")
	effect, _ := user.InputInt()
	tOfferAdjust := ticketOfferAdjustmentDM(effect)
	fmt.Println("The Negotiation:")
	fmt.Println("Negotiation process requires both primary administrators (the employer and the employee) to throw their Broker skill. If the employer’s result is higher than the mercenary’s, then the ticket is more or less arranged as the employer needs it to be, and all of the tables used to generate a mercenary ticket (see Creating the Ticket below) are rolled normally, with no modifiers. If the mercenary administrator manages to roll a higher overall result than the employer, his Effect is compared to the table below. This shows just how much sway the mercenaries have in creating the ticket and adjusting it according to their wishes. Ticket adjustments are discussed further later in this chapter.")
	fmt.Print("встречная проверка патрона против игрока\nСумма броска игрока: ")
	mercRoll, _ := user.InputInt()
	patronRoll := dp.RollNext("2d6").DM(t.patronDM).Shout().Sum()
	fmt.Printf("Результат патрона: %v\n", patronRoll)
	nta := NumberOfTicketAdjustment(mercRoll - patronRoll + tOfferAdjust)
	fmt.Printf("Суммарное кол-во очков изменений: %v\n", nta)
	t.ticketAdjustments = nta
}

func NumberOfTicketAdjustment(effect int) int {
	dp := dice.New()
	switch effect {
	case 2:
		return dp.RollNext("1d6").DM(1).Shout().Sum()
	case 3:
		return dp.RollNext("1d6").DM(2).Shout().Sum()
	case 4:
		return dp.RollNext("1d6").DM(3).Shout().Sum()
	case 5:
		return dp.RollNext("2d6").DM(0).Shout().Sum()
	}
	if effect < 2 {
		return dp.RollNext("1d6").DM(0).Shout().Sum()
	}
	return dp.RollNext("2d6").DM(2).Shout().Sum()
}

func ticketOfferAdjustmentDM(effect int) int {
	switch effect {
	case 2:
		return -2
	case 3, 4:
		return 0
	case 5:
		return 1
	}
	if effect < 2 {
		return -4
	}
	return 2
}

func rollPatronDM() int {
	switch dice.Roll2D() {
	case 2, 3:
		return 0
	case 4, 5:
		return 1
	case 6, 7, 8:
		return 2
	case 9, 10:
		return 3
	case 11, 12:
		return 4
	}
	return -999
}

func status(t *Ticket) {
	utils.ClearScreen()
	fmt.Println(t)
	fmt.Printf("Ticket Adjustment Points : %v\n", t.ticketAdjustments)
	fmt.Printf("Patron Negotiation Bonus : %v\n", t.patronDM)
	fmt.Printf("Employer Details         : %v\n", t.employer)
	fmt.Printf("Employee Details         : %v\n", t.employee)
	fmt.Println("-------------------------")
	fmt.Printf("Mission Service Type     : %v\n", t.service.genericServiceType)
	fmt.Printf("Mission Type             : %v\n", t.service.mission)
	fmt.Printf("Mission Paygrade         : %v (%v)\n", t.service.payGrade, paygrade2Credits(t.service.payGrade))
	fmt.Printf("Mission Service Lenght   : %v \n", t.service.lenghtOfService)
	fmt.Printf("Mission Exposure         : %v \n", t.service.exposure)
	fmt.Printf("Mission Target           : %v \n", t.service.targetType)
	fmt.Printf("Mission Target Descriptor: %v \n", t.service.targetDescriptor)
	fmt.Printf("Mission Risk             : %v \n", t.service.risk)
	fmt.Println("-------------------------")
	if len(t.preSupport) > 0 {
		for i, val := range t.preSupport {
			switch i {
			case 0:
				fmt.Printf("Pre-Ticket Support: %v\n", val)
			default:
				fmt.Printf("                    %v\n", val)
			}
		}
	}
	if len(t.postSupport) > 0 {
		for i, val := range t.postSupport {
			switch i {
			case 0:
				fmt.Printf("Post-Ticket Support: %v\n", val)
			default:
				fmt.Printf("                     %v\n", val)
			}
		}
	}

}

func employerDetails() string {
	dp := dice.New()
	sum := dp.RollNext("2d6").Sum()
	details := ""
	switch sum {
	case 2:
		details = "Employer is trying to remain anonymous and use false nomenclature to protect itself."
	case 3, 4, 5:
		details = "Employer is purposefully vague on important details."
	case 6, 7, 8:
		details = "Employer is perfectly honest in the ticket, but details are little more than title and mode of communication."
	case 9, 10:
		details = "Honest details; including the employing agent’s name and direct communication."
	case 11:
		details = "Honest and very detailed information about the employer."
	case 12:
		details = "Private Ticket; employer is honest – but is willing to pay extra to keep the information secret."
	}
	return details
}

func paygrade2Credits(pg int) string {
	pay := ""
	switch pg {
	default:
		pay = "Undefined"
	case 1:
		pay = "5,000 Cr"
	case 2:
		pay = "10,000 Cr"
	case 3:
		pay = "20,000 Cr"
	case 4:
		pay = "30,000 Cr"
	case 5:
		pay = "50,000 Cr"
	case 6:
		pay = "75,000 Cr"
	case 7:
		pay = "100,000 Cr"
	case 8:
		pay = "150,000 Cr"
	case 9:
		pay = "200,000 Cr"
	case 10:
		pay = "250,000 Cr"
	case 11:
		pay = "325,000 Cr"
	case 12:
		pay = "400,000 Cr"
	case 13:
		pay = "500,000 Cr"
	case 14:
		pay = "750,000 Cr"
	case 15:
		pay = "1 MCr"
	case 16:
		pay = "1.5 MCr"
	case 17:
		pay = "2 MCr"
	case 18:
		pay = "3 MCr"
	case 19:
		pay = "4 MCr"
	case 20:
		pay = "5 MCr"
	case 21:
		pay = "7 MCr"
	case 22:
		pay = "10 MCr"
	case 23:
		pay = "15 MCr"
	case 24:
		pay = "20 MCr"
	case 25:
		pay = "25 MCr"
	case 26:
		pay = "30 MCr"
	case 27:
		pay = "40 MCr"
	case 28:
		pay = "50 MCr"

	}
	return pay
}

func defineMission(t *Ticket) string {
	def := ""
	switch t.service.mission {
	default:
		def = "UNDEFINED:" + t.service.mission
	case mType_ActiveDuty:
		def = "The mission assigns the mercenary unit to serve on a battlefield that is currently home to some sort of military conflict."
		switch t.service.genericServiceType {
		case sType_Cadre:
			def += " Cadre missions of this type tend to be ‘hands on’ training runs that let the mercenary unit help one side of the conflict learn how to survive and hopefully be victorious."
		case sType_Commando:
			def += " Commando missions in active duty roles tend to be simple fillin positions for regular military units."
		case sType_Security:
			def += " Security units placed in active duty are used to protect important battlefield assets or personalities."
		}
	case mType_Assasination:
		def = "The mission involves the killing of a specific target or targets. "
		switch t.service.genericServiceType {
		case sType_Criminal:
			def += "Criminal missions of this kind are unsanctioned murders, plain and simple."
		case sType_Guerilla:
			def += "Guerrilla units sent to assassinate someone are almost always targeting the head of a company, governmental office or rival force in the way of the employer’s progress."
		}
	case mType_Commerce:
		def = "This Warmonger-only mission type is used to describe a ticket that revolves around the sale of specific goods or services in a hostile area, or to a hostile client. Including arms deals, smuggling and personal contracting, this mission is all about making profits."
	case mType_CounterStrike:
		def = "This Striker-only mission is the directed use of force in retaliation for an attack of some kind against the employer. Paying largely for expedience and rapid-response, these missions are used normally at the beginning of smaller conflicts before they become full-fledged wars."
	case mType_Defence:
		def = "This Security-only mission is based around the hiring and assignment of mercenary personnel to a single location or person, protecting them from outside harm. This could be a corporate location, a specific item, or even a travelling starship."
	case mType_Elimination:
		def = "The mission is based around the active seeking of a target and destroying it. It could be a location, an item, a group or a specific piece of information. "
		switch t.service.genericServiceType {
		case sType_Commando:
			def += "Commando missions of this type frequently are a single piece of a larger effort. The unit must destroy a specific target that would be helpful to the greater conflict effort."
		case sType_Striker:
			def += "Striker units undertaking these missions are frequently targeting hard-to-get locations or groups, using their fast insertion techniques to deal with the target before any further defences can be raised."
		case sType_Dream:
			def += "Dream missions of this type are too good to be true. They are unguarded targets that the unit must eliminate, and likely without too much trouble. The pay is probably too much for such an easy job, but few mercenaries ask why."
		}
	case mType_Escort:
		def = "This mission involves the unit going from one point to the next while defending a specific target from capture or destruction. "
		switch t.service.genericServiceType {
		case sType_Security:
			def += "Security units assigned to this mission are likely bodyguards or handlers of something or someone worth a great deal to the employer."
		case sType_Warmonger:
			def += "Warmonger mercenaries that are escorting items are also frequently smuggling the item as well, making sure that the target reaches its destination unmolested."
		case sType_Dream:
			def += "Dream escorts are much like any other babysitting role, but pay immensely well due to the extreme importance of the target to the employer."
		}
	case mType_FieldExercise:
		def = "This mission type involves the mercenary unit performing some sort of average task in a potential hostile location. "
		switch t.service.genericServiceType {
		case sType_Security:
			def += "Cadre units often take groups of recruits or employers out into the field to train them in ways that they cannot manage in a gym or classroom."
		case sType_Warmonger:
			def += "Warmonger units are sometimes asked to bring their wares out directly to those who will put them to use."
		case sType_Dream:
			def += "Dream missions of this type send the mercenary unit to a task in a utopian area for them, and they can enjoy their environment while they fulfil their ticket."
		}
	case mType_FirstStrike:
		def = "This mission type is the preliminary attack of any conflict, often starting a greater escalation. "
		switch t.service.genericServiceType {
		case sType_Guerilla:
			def += "Guerrilla mercenaries that sign on for first strike assignments are likely to be making a very public statement about their target at the same time."
		case sType_Commando:
			def += "Commando units that are given the mission to go on first strikes are typically attacking a location with massed firepower and military fervour."
		case sType_Striker:
			def += "Striker units on these missions are tactical offensive groups that hit hard, fast and without pause. Their targets rarely have a chance to defend themselves and employers expect a great deal of momentum and efficiency."
		case sType_Dream:
			def += "Dream missions of this type are rarely difficult for the unit and target unknowing and lightly defended targets that are extremely important to the employer – making them high-paying and lowrisk."
		}
	case mType_Raid:
		def = "This mission type is used to specifically cause financial damage to the target. "
		switch t.service.genericServiceType {
		case sType_Criminal:
			def += "Criminal missions of this type are often planned as thefts, vandalism or arson."
		case sType_Commando:
			def += "Commando units that are given raid missions are sent in to a location to cause as much collateral damage as they can while performing their manoeuvres. They are supposed to go in, inflict mass damage, and then quickly evacuate."
		case sType_Warmonger:
			def += "Warmonger missions of this type are aimed at rivals of the employer, taking assets from them in order to weaken the target’s position against them."
		}
	case mType_Recon:
		def = "This mission involves gathering information on the target. "
		switch t.service.genericServiceType {
		case sType_Guerilla:
			def += "Guerrilla units on recon duty are secretly learning about the target to likely use it as part of an attack or offensive against it later."
		case sType_Cadre:
			def += "Cadre mercenaries that sign up for these missions are in charge of showing others the best way to gather intelligence upon the target."
		case sType_Striker:
			def += "Striker missions of this type are designed to get into hostile territory and learn as much as the mercenaries can, by whatever means necessary, before making a hasty escape with the information."
		case sType_Dream:
			def += "Dream assignments in a recon mission are ‘cushy’ jobs that involve non-hostile targets. They are often personal in nature to the employer, and the mercenaries are being paid for their subtlety instead of their firepower."
		}
	case mType_Retrival:
		def = "This Commando-only mission type involves a heavily armed unit going into enemy territory with guns blazing and engines hot in order to find, obtain and evacuate with the target. These are frequently used in military situations when prisoners are involved, but governmental hands are too politically tied to take action."
	case mType_Sabotage:
		def = "This mission type involves the wilful tampering or even destruction of items or locations belonging to the target. "
		switch t.service.genericServiceType {
		case sType_Criminal:
			def += "Criminal missions of this type are almost always aimed at the local operational authorities, otherwise they would not be considered criminal in nature."
		case sType_Guerilla:
			def += "Guerrilla units signing on to commit acts of sabotage are likely to do so in dramatic and showy way."
		}
	case mType_TechnologicalTest:
		def = "This Dream mission involves the mercenary unit being equipped with a brand new and untested piece of equipment that they are to give field testing for. Whether it is a new type of armour, weapon, or something as mundane as a new type of environ-tent, the unit is paid well to do research into the usefulness of the item. NOTE: This mission grants the mercenary administrator 2d3 additional Ticket Adjustments for use solely in the Post-Ticket Support section of the ticket."
	case mType_Terrorise:
		def = "This Guerrilla-only mission is the application of violence and fear to make some kind of political or social statement on behalf of the employer. This is sometimes considered to be a truly despicable ticket type, but pays very well for a surprisingly low amount of work."
	case mType_Train:
		def = "This Cadre-only mission involves the mercenaries staying on a base, ship, or compound where they will be helping the target learn whatever skills the employer specifies. This is a very safe ticket to undertake, but most mercenaries also find them excruciatingly boring and frustrating."
	case mType_UnlawfulAquisition:
		def = "This Criminal mission type is the basic idea of picking something up for the employer that does not currently belong to them. Whether it is simple theft, hijacking, kidnapping or some other form of ‘acquisition’, the unit must take possession of the target and bring it to the employer."
	}
	return def
}

func exposureTable(i int) string {
	if i <= 2 {
		return "Hidden – Without doing research, no one knows the ticket existed."
	}
	switch i {
	case 3, 4:
		return "Obscure – Only the local public is aware of the ticket’s actions."
	case 5, 6:
		return "Low Profile – Much of the planet is aware of the ticket’s actions, and the local public know the name of the mercenary unit."
	case 7, 8:
		return "Uncommon – The ticket has received some media attention in the local area and the mercenary unit’s involvement is locally public."
	case 9, 10:
		return "Common – The media has spread the mercenary unit’s name throughout the planet and it has spread to neighbouring planets."
	case 11, 12:
		return "Exposed – The mercenary unit’s name is publicly known on a planetary level; even a few specific members’ names are being said."
	}
	return "High Profile – The ticket and the mercenary unit are being talked about throughout the system. At least one member of the unit is being named specifically."
}

func defensiveTargetTypes(i int) (target string, payGradeAdjustment int) {
	switch i {
	default:
		return fmt.Sprintf("Error: index %v cannot be chosen", i), 0
	case 1:
		return "Item", 0
	case 2:
		return "Location", 1
	case 3:
		return "Ally, Individual", 1
	case 4:
		return "Information", -1
	case 5:
		return "Ship", 1
	case 6:
		return "Ally, Group", 2
	}
}

func printDefensiveTargetTypes() {
	fmt.Println("-------------------------------------------")
	fmt.Println("#   Type of Target     Pay Grade Adjustment")
	fmt.Println("1   Item                 --------")
	fmt.Println("2   Location           +1 Increment")
	fmt.Println("3   Ally, Individual   +1 Increment")
	fmt.Println("4   Information        -1 Increment")
	fmt.Println("5   Ship               +1 Increment")
	fmt.Println("6   Ally, Group        +2 Increment")
	fmt.Println("-------------------------------------------")
}

func neutralTargetTypes(i int) (target string, payGradeAdjustment int) {
	switch i {
	default:
		return fmt.Sprintf("Error: index %v cannot be chosen", i), 0
	case 1:
		return "Item", 0
	case 2:
		return "Trade Goods", 0
	case 3:
		return "Individual", -1
	case 4:
		return "Personal Goods", 0
	case 5:
		return "Ship", 2
	case 6:
		return "Activity", 0
	}
}

func printNeutralTargetTypes() {
	fmt.Println("-------------------------------------------")
	fmt.Println("#   Type of Target   Pay Grade Adjustment")
	fmt.Println("1   Item               --------")
	fmt.Println("2   Trade Goods        --------")
	fmt.Println("3   Individual       -1 Increment")
	fmt.Println("4   Personal Goods     --------")
	fmt.Println("5   Ship             +2 Increment")
	fmt.Println("6   Activity           --------")
	fmt.Println("-------------------------------------------")
}

func offensiveTargetTypes(i int) (target string, payGradeAdjustment int) {
	switch i {
	default:
		return fmt.Sprintf("Error: index %v cannot be chosen", i), 0
	case 1:
		return "Individual", 1
	case 2:
		return "Location", 2
	case 3:
		return "Item", 1
	case 4:
		return "Vechicle", 2
	case 5:
		return "Ship", 2
	case 6:
		return "Group", 2
	}
}

func printOffensiveTargetTypes() {
	fmt.Println("-------------------------------------------")
	fmt.Println("#   Type of Target   Pay Grade Adjustment")
	fmt.Println("1   Individual       +1 Increment")
	fmt.Println("2   Location         +2 Increment")
	fmt.Println("3   Item             +1 Increment")
	fmt.Println("4   Vechicle         +2 Increment")
	fmt.Println("5   Ship             +2 Increment")
	fmt.Println("6   Group            +2 Increment")
	fmt.Println("-------------------------------------------")
}

func printPreTicketSupportTypes() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("#     Support Table Used   Pay Grade Adjustment")
	fmt.Println("1-2   Advance Funds        -1 Increment")
	fmt.Println("3-4   Services             -1 Increment")
	fmt.Println("5-6   Equpment             -3 Increment")
	fmt.Println("-----------------------------------------------")
}

func printSupportAdvanceFunds() {
	fmt.Println("-------------------------")
	fmt.Println("#   Advance Funds Offered")
	fmt.Println("1   5,000 Cr")
	fmt.Println("2   10,000 Cr")
	fmt.Println("3   20,000 Cr")
	fmt.Println("4   30,000 Cr")
	fmt.Println("5   40,000 Cr")
	fmt.Println("6   50,000 Cr")
	fmt.Println("-------------------------")
}

func printSupportServices() {
	fmt.Println("--------------------")
	fmt.Println("#   Services Offered")
	fmt.Println("1   Transportation")
	fmt.Println("2   Transportation")
	fmt.Println("3   Eqipment Repairs")
	fmt.Println("4   Rearmement")
	fmt.Println("5   Arms Traiding")
	fmt.Println("6   Medical Process")
	fmt.Println("--------------------")
}

func printSupportEquipment() {
	fmt.Println("---------------------")
	fmt.Println("#   Equipment Offered")
	fmt.Println("1   Basics")
	fmt.Println("2   Armour")
	fmt.Println("3   Weapons")
	fmt.Println("4   Heavy Weapon")
	fmt.Println("5   Transport")
	fmt.Println("6   Specialised Gear")
	fmt.Println("---------------------")
}

func supportAdvanceFundsTable(i int) string {
	arr := []string{
		"5,000 Cr",
		"10,000 Cr",
		"20,000 Cr",
		"30,000 Cr",
		"40,000 Cr",
		"50,000 Cr",
	}
	switch i {
	default:
		return "Error"
	case 1, 2, 3, 4, 5, 6:
		return "Advanced Funds (" + arr[i-1] + ")"
	}
}

func supportServicesTable(i int) string {
	arr := []string{
		"Transportation",
		"Transportation",
		"Eqipment Repairs",
		"Rearmement",
		"Arms Traiding",
		"Medical Process",
	}
	switch i {
	default:
		return "Error"
	case 1, 2, 3, 4, 5, 6:
		return "Service (" + arr[i-1] + ")"
	}
}

func supportEquipmentTable(i int) string {
	arr := []string{
		"Basics",
		"Armour",
		"Weapons",
		"Heavy Weapon",
		"Transport",
		"Specialised Gear",
	}
	switch i {
	default:
		return "Error"
	case 1, 2, 3, 4, 5, 6:
		return "Equpment (" + arr[i-1] + ")"
	}
}
