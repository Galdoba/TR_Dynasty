package Trade

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/devtools/cli/features"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

func Run() {
	tst := user.InputStr()
	fmt.Println(tst)

	return
	features.TypingSlowly("Initiating Trade Terminal...", 15)

	uwpNotValid := true
	uwp := ""

	for uwpNotValid {
		features.TypingSlowly("\nInput World UWP: ", 15)
		uwp = utils.InputString("")
		if world.UWPisValid(uwp) {
			uwpNotValid = false
		}
		switch uwpNotValid {
		default:
			features.TypingSlowly("ERROR: Input unparseable", 15)
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
	features.TypingSlowly("World Trade Codes: ", 15)
	features.TypingSlowly(tcs+"\n", 15)
	features.TypingSlowly("Connecting to Local Trade Network...OK\n", 15)
	switch actionSelector() {
	case 0:
		features.TypingSlowly("Disconnecting...\n", 15)
		features.TypingSlowly("Have a nice day!\n", 15)
	case 1:
		features.TypingSlowly("Enter Negotiation Effect: ", 15)
		neg := 0
		neg = utils.InputInt("")
		features.TypingSlowly("Compiling Datasheet\n", 15)
		merch := NewMerchant().SetLocalUWP(uwp).SetLocalTC(tc).SetMType(constant.MerchantTypeTrade).SetTraderDice(neg).DetermineGoodsAvailable()
		features.TypingSlowly("Gathering data", 15)
		merch.PurchaseList()
	case 2:
		features.TypingSlowly("Enter Negotiation Effect: ", 15)
		neg := 0
		neg = utils.InputInt("")

		features.TypingSlowly("Enter Trade Good code and volume '[tgCode] [X]': ", 15)
		//err := true
		codeVol := ""
		codeVol = utils.InputString("")
		fmt.Println(neg, codeVol)

	}
	// features.TypingSlowly("Enter Negotiation Effect: ", 15)
	// neg := 0
	// neg = utils.InputInt("")
	// features.TypingSlowly("Compiling Datasheet", 15)
	// merch := NewMerchant().SetLocalUWP(sw.UWP()).SetLocalTC(sw.TradeCodes()).SetMType(constant.MerchantTypeTrade).SetTraderDice(neg).DetermineGoodsAvailable()
	// //fmt.Println(merch)
	// merch.PurchaseList()
	// //merch.ListPrices()
}

func actionSelector() int {
	parsed := false
	var action int
	features.TypingSlowly("Select procesure:\n", 15)
	features.TypingSlowly("[0] - Disconnect\n", 15)
	features.TypingSlowly("[1] - Show available Trade Goods\n", 15)
	features.TypingSlowly("[2] - Sell Trade Goods\n", 15)
	for !parsed {

		action = utils.InputInt("User Input>")
		switch action {
		default:
			features.TypingSlowly("Action "+strconv.Itoa(action)+" Incorrect", 15)
		case 0, 1, 2:
			features.TypingSlowly("Initiating procesure "+strconv.Itoa(action)+"\n", 15)
			parsed = true
		}
	}
	return action
}

func testInput(tst string) string { //TODO: Переписать утилсовые инпуты
	return utils.InputString(tst)
}
