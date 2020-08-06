package Trade

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/cli/features"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

func Run() {

	features.TypingSlowly("Initiating Trade Terminal...", 30)

	uwpNotValid := true
	uwp := ""

	for uwpNotValid {
		features.TypingSlowly("\nInput World UWP: ", 30)
		uwp = utils.InputString("")
		if world.UWPisValid(uwp) {
			uwpNotValid = false
		}
		switch uwpNotValid {
		default:
			features.TypingSlowly("ERROR: Input unparseable", 30)
		case false:
		}

	}
	sw := world.FromUWP(uwp).UpdateTC()
	tc := sw.TradeCodes()
	tcs := ""
	for i := range tc {
		tcs += tc[i] + " "
	}
	features.TypingSlowly("World Trade Codes: ", 30)
	features.TypingSlowly(tcs+"\n", 30)
	features.TypingSlowly("Connecting to Local Trade Network...OK\n", 30)
	switch actionSelector() {
	case 0:
		features.TypingSlowly("Disconnecting...\n", 30)
		features.TypingSlowly("Have a nice day!\n", 30)
	case 1:
		features.TypingSlowly("Enter Negotiation Effect: ", 30)
		neg := 0
		neg = utils.InputInt("")
		features.TypingSlowly("Compiling Datasheet", 30)
		merch := NewMerchant().SetLocalUWP(sw.UWP()).SetLocalTC(sw.TradeCodes()).SetMType(constant.MerchantTypeTrade).SetTraderDice(neg).DetermineGoodsAvailable()

		merch.PurchaseList()
	}
	// features.TypingSlowly("Enter Negotiation Effect: ", 30)
	// neg := 0
	// neg = utils.InputInt("")
	// features.TypingSlowly("Compiling Datasheet", 30)
	// merch := NewMerchant().SetLocalUWP(sw.UWP()).SetLocalTC(sw.TradeCodes()).SetMType(constant.MerchantTypeTrade).SetTraderDice(neg).DetermineGoodsAvailable()
	// //fmt.Println(merch)
	// merch.PurchaseList()
	// //merch.ListPrices()
}

func actionSelector() int {
	parsed := false
	var action int
	features.TypingSlowly("Select procesure:\n", 30)
	features.TypingSlowly("[0] - Disconnect\n", 30)
	features.TypingSlowly("[1] - Show available Trade Goods\n", 30)
	for !parsed {

		action = utils.InputInt("GO: ")
		switch action {
		default:
			features.TypingSlowly("Action "+strconv.Itoa(action)+" Incorrect", 30)
		case 0, 1:
			features.TypingSlowly("Initiating procesure "+strconv.Itoa(action)+"\n", 30)
			parsed = true
		}
	}
	return action
}
