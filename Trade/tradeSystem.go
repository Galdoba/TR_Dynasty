package trade

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Galdoba/devtools/cli/features"
	"github.com/Galdoba/devtools/cli/prettytable"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	profile "github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/devtools/cli/user"
)

//Run -
func Run() {
	features.TypingSlowly("Initiating Trade Terminal...", 15)
	uwpNotValid := true
	uwp := ""
	for uwpNotValid {
		features.TypingSlowly("\nInput World UWP: ", 15)
		uwp, err := user.InputStr()
		if err != nil {
			features.TypingSlowly(err.Error(), 15)
			continue
		}
		if world.UWPisValid(uwp) {
			uwpNotValid = false
		}
		switch uwpNotValid {
		default:
			features.TypingSlowly("ERROR: Input unparseable", 15)
		case false:
		}

	}
	tc := profile.CalculateTradeCodes(uwp)
	tcs := ""
	for i := range tc {
		tcs += tc[i] + " "
	}
	features.TypingSlowly("World Trade Codes: ", 15)
	features.TypingSlowly(tcs+"\n", 15)
	features.TypingSlowly("Connecting to Local Trade Network...OK\n", 15)
	merch := NewMerchant().SetLocalUWP(uwp).SetLocalTC(tc).SetMType(constant.MerchantTypeTrade).SetTraderDice(0).DetermineGoodsAvailable()
	neg := 0
	done := false
	for !done {
		switch actionSelector() {
		case 0:
			features.TypingSlowly("Disconnecting...\n", 15)
			features.TypingSlowly("Have a nice day!\n", 15)
			done = true
		case 1:
			merch.SetTraderDice(neg)
			features.TypingSlowly("Gathering data:\n", 15)
			merch.PurchaseList()
		case 2:
			volCode := gatherIntSl("Enter Trade Good code and amount: ")
			code := volCode[0]
			vol, _ := strconv.Atoi(volCode[1])
			merch.SetTraderDice(neg)
			features.TypingSlowly(merch.SaleProposalLegal(code, vol), 15)
		case 3:
			neg = gatherInt("Roll Broker(SOC) 8 Check and enter Effect: ")
		}
	}
}

func gatherInt(descr string) int {
	res := 0
	err := errors.New("No Data")
	for err != nil {
		features.TypingSlowly(descr, 15)
		res, err = user.InputInt()
		if err == nil {
			break
		}
		features.TypingSlowly(err.Error()+"\n", 15)
	}
	return res
}

func gatherIntSl(descr string) []string {
	var res []int
	err := errors.New("No Data")
	for err != nil {
		features.TypingSlowly(descr, 15)
		res, err = user.InputSliceInt()
		if err == nil && len(res) == 2 {
			break
		}
		err = errors.New("Arguments found: " + strconv.Itoa(len(res)) + " (expecting 2)")
		features.TypingSlowly(err.Error()+"\n", 15)
	}
	var resS []string
	for i := range res {
		resS = append(resS, strconv.Itoa(res[i]))
	}
	return resS
}

func actionSelector() int {
	parsed := false
	var action int
	features.TypingSlowly("Select procesure:\n", 15)
	features.TypingSlowly("[0] - Disconnect\n", 15)
	features.TypingSlowly("[1] - Show available Trade Goods\n", 15)
	features.TypingSlowly("[2] - Sell Trade Goods\n", 15)
	features.TypingSlowly("[3] - Renegotiate prices (Roleplay)\n", 15)
	for !parsed {
		features.TypingSlowly("User Input>", 15)
		action, _ = user.InputInt()
		switch action {
		default:
			features.TypingSlowly("\rAction "+strconv.Itoa(action)+" Incorrect", 15)
		case 0, 1, 2, 3:
			features.TypingSlowly("\rInitiating procesure "+strconv.Itoa(action)+"\n", 15)
			parsed = true
		}
	}
	return action
}

//SaleProposalLegal -
func (m Merchant) SaleProposalLegal(code string, amount int) string {
	basePrice := GetBasePrice(code)
	salePrice := m.CostSale(code)
	profit := (salePrice - basePrice) * amount
	if profit < 0 {
		profit = 0
	}
	tax := taxingAmount(profit, string([]byte(m.localUWP)[5]))
	proposal := ""
	proposal += "Trade Lot: " + strconv.Itoa(amount) + " x " + GetDescription(code) + "\n"
	proposal += " Proposal: " + strconv.Itoa(salePrice) + " x " + strconv.Itoa(amount) + " (" + strconv.Itoa(salePrice*amount) + " Cr)" + "\n"
	proposal += "      Tax: " + strconv.Itoa(tax) + " Cr" + "\n"
	proposal += "---------------------" + "\n"
	proposal += "   Profit: " + strconv.Itoa((salePrice-tax)*amount) + " Cr" + "\n"
	return proposal
}

//PurchaseList -
func (m Merchant) PurchaseList() {
	tb := prettytable.New()
	tb.AddRow([]string{"Item", "Maximum Tons", "Tons per Defined Trade Good (Base Price)", "Cost", "Purchase DM"})
	aC := allCategories()
	for i := range aC {
		l := len(aC)
		p := ((i + 1) * 100 / l)
		fmt.Print(strconv.Itoa(p) + "% done\r")
		catList := listCategory(m, aC[i])
		for l := range catList {
			tb.AddRow(catList[l])
		}
	}
	tb = prettytable.InsertSeparatorRow(tb, 1)
	tb.PTPrintSlow(0)
}

func listCategory(m Merchant, code string) [][]string {
	if code == "66" {
		return [][]string{}
	}
	maxTons := RollMaximumForCategory(code) * countElement(code, m.availableTGcodes)
	purchaseDM, _ := PurchSaleDMs(code, m.localTC)
	var dataSheet [][]string
	exactVolume := make(map[string]int)
	oppCount := 1
	for i := 0; i < maxTons; i++ {
		oppCount++
		descr := dice.Roll("2d6").SumStr()
		switch descr {
		case "3", "4", "5":
			descr = "4"
		case "6", "7", "8":
			descr = "7"
		case "9", "10", "11":
			descr = "10"
		}
		increment := IncreseTG(code + descr)
		exactVolume[code+descr] = exactVolume[code+descr] + increment
		i = i + increment - 1
		if sumMap(exactVolume) > maxTons {
			for sumMap(exactVolume) > maxTons {
				exactVolume[code+descr]--
			}
			break
		}
	}
	madeCat := false
	for _, descr := range []string{"2", "4", "7", "10", "12"} {
		if exactVolume[code+descr] == 0 {
			continue
		}
		dataline := make([]string, 5)
		if !madeCat {
			dataline[0] = GetCategory(code + descr)
			dataline[1] = strconv.Itoa(maxTons)
			dataline[4] = strconv.Itoa(purchaseDM)
			if purchaseDM >= 0 {
				dataline[4] = "+" + dataline[4]
			}

			madeCat = true
		}
		basePrice := GetBasePrice(code + descr)
		dataline[2] = strconv.Itoa(exactVolume[code+descr]) + " x " + GetDescription(code+descr) + " (" + strconv.Itoa(basePrice) + ")"
		costP := m.CostPurchase(code + descr)
		if exactVolume[code+descr] > 0 {
			dataline[3] = strconv.Itoa(costP) // + " (" + strconv.Itoa(basePrice) + ")"
		}
		dataSheet = append(dataSheet, dataline)
	}
	return dataSheet
}
