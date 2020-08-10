package Trade

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

func Run() {

	fmt.Println("Initiating Trade Terminal...")
	uwpNotValid := true
	uwp := ""

	for uwpNotValid {
		fmt.Println("\nInput World UWP: ")
		uwp = utils.InputString("")
		if world.UWPisValid(uwp) {
			uwpNotValid = false
		}
		switch uwpNotValid {
		default:
			fmt.Println("ERROR: Input unparseable")
		case false:
		}

	}
	//sw := world.FromUWP(uwp).UpdateTC()
	//tc := sw.TradeCodes()
	tc := profile.TradeCodes(uwp)
	tcs := ""
	for i := range tc {
		tcs += tc[i] + " "
	}
	fmt.Println(TradeDMs("15", tc))
	fmt.Println("World Trade Codes: ")
	fmt.Println(tcs + "\n")
	fmt.Println("Connecting to Local Trade Network...OK\n")
	switch actionSelector() {
	case 0:
		fmt.Println("Disconnecting...\n")
		fmt.Println("Have a nice day!\n")
	case 1:
		fmt.Println("Enter Negotiation Effect: ")
		neg := 0
		neg = utils.InputInt("")
		fmt.Println("Compiling Datasheet\n")
		merch := NewMerchant().SetLocalUWP(uwp).SetLocalTC(tc).SetMType(constant.MerchantTypeTrade).SetTraderDice(neg).DetermineGoodsAvailable()
		fmt.Println("Gathering data")
		merch.PurchaseList()
	case 2:
		neg := gatherInt("Roll Broker(SOC) 8 Check and enter Effect: ")
		code := gatherInt("Enter Trade Good code: ")
		vol := gatherInt("Enter Trade Good volume: ")
		fmt.Println("neg, code, vol")
		fmt.Println(neg, code, vol)
		fmt.Println("--------------")

	}
	// fmt.Println("Enter Negotiation Effect: ")
	// neg := 0
	// neg = utils.InputInt("")
	// fmt.Println("Compiling Datasheet")
	// merch := NewMerchant().SetLocalUWP(sw.UWP()).SetLocalTC(sw.TradeCodes()).SetMType(constant.MerchantTypeTrade).SetTraderDice(neg).DetermineGoodsAvailable()
	// //fmt.Println(merch)
	// merch.PurchaseList()
	// //merch.ListPrices()
}

func gatherInt(descr string) int {
	res := 0
	err := errors.New("No Data")
	for err != nil {
		fmt.Println(descr)
		res, err = user.InputInt()
		if err == nil {
			break
		}
		fmt.Println(err.Error() + "\n")
	}
	return res
}

func actionSelector() int {
	parsed := false
	var action int
	fmt.Println("Select procesure:\n")
	fmt.Println("[0] - Disconnect\n")
	fmt.Println("[1] - Show available Trade Goods\n")
	fmt.Println("[2] - Sell Trade Goods\n")
	for !parsed {
		fmt.Println("User Input>")
		action, _ = user.InputInt()
		switch action {
		default:
			fmt.Println("Action " + strconv.Itoa(action) + " Incorrect")
		case 0, 1, 2:
			fmt.Println("Initiating procesure " + strconv.Itoa(action) + "\n")
			parsed = true
		}
	}
	return action
}

func testInput(tst string) string { //TODO: Переписать утилсовые инпуты
	return utils.InputString(tst)
}
