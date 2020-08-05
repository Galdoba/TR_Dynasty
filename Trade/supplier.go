package Trade

import (
	"fmt"
	"os"
	"sort"

	"github.com/Galdoba/TR_Dynasty/cli/prettytable"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

var tgDB map[string][]string

func init() {
	tgDB = TradeGoodRData()
}

type Merchant struct {
	mType            string
	availableTGcodes []string
	localTC          []string
	localUWP         string
	tradeDice        int
	prices           map[string]int
	volume           map[string]int
}

func NewMerchant() Merchant {
	m := Merchant{}
	m.tradeDice = 0
	m.prices = make(map[string]int)
	m.volume = make(map[string]int)
	return m
}

func (m Merchant) SetLocalUWP(uwp string) Merchant {
	m.localUWP = uwp
	return m
}

func (m Merchant) SetLocalTC(tc []string) Merchant {
	m.localTC = tc
	return m
}

func (m Merchant) SetMType(mType string) Merchant {
	mTypeErr := true
	for _, val := range []string{supplierTypeCommon, supplierTypeTrade, supplierTypeNeutral, supplierTypeCommon} {
		if mType == val {
			mTypeErr = false
			break
		}
	}
	if mTypeErr {
		return m
	}
	m.mType = mType
	return m
}

func (m Merchant) SetTraderDice(tDice int) Merchant {
	m.tradeDice = tDice
	return m
}

func (m Merchant) DetermineGoodsAvailable() Merchant {
	var avGoodsCodes []string
	availableCategories := m.availableTGcodes
	switch m.mType {
	default:
		for i := 0; i < dice.Roll1D(); i++ {
			roll := dice.RollD66()
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeCommon:
		availableCategories = []string{"11", "12", "13", "14", "15", "16"}
		add := utils.RollDiceRandom("d6")
		for i := 1; i < add; i++ {
			roll := TrvCore.RollD66()
			if !utils.ListContains([]string{"11", "12", "13", "14", "15", "16"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeTrade:
		availableCategories = append([]string{"11", "12", "13", "14", "15", "16"}, availableCategories...)
		add := utils.RollDiceRandom("d6")
		for i := 0; i < add; i++ {
			roll := dice.RollD66()
			if utils.ListContains([]string{"61", "62", "63", "64", "65", "66"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
		ExcludeFromSliceStr(availableCategories, "61")
		ExcludeFromSliceStr(availableCategories, "62")
		ExcludeFromSliceStr(availableCategories, "63")
		ExcludeFromSliceStr(availableCategories, "64")
		ExcludeFromSliceStr(availableCategories, "65")
		//ExcludeFromSliceStr(availableCategories, "66")
	case supplierTypeNeutral:
		add := dice.Roll1D()
		for i := 0; i < add; i++ {
			roll := dice.RollD66()
			if utils.ListContains([]string{"66"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	case supplierTypeIlligal:
		availableCategories = []string{"61", "62", "63", "64", "65"}
		add := dice.Roll1D()
		for i := 0; i < add; i++ {
			roll := dice.RollD66()
			if !utils.ListContains([]string{"61", "62", "63", "64", "65"}, roll) {
				i--
				continue
			}
			availableCategories = append(availableCategories, roll)
		}
	}
	sort.Strings(availableCategories)
	//fmt.Println(availableCategories)
	for c := range availableCategories {

		key := availableCategories[c]
		m.prices[key] = dice.Roll3D()
		//key = key + dice.Roll("2d6").SumStr()
		m.volume[categoryOf(key)] = m.volume[categoryOf(key)] + RollMaximumForCategory(key)
		avGoodsCodes = append(avGoodsCodes, key)

	}
	sort.Strings(avGoodsCodes)

	m.availableTGcodes = avGoodsCodes
	//fmt.Println(m.availableTGcodes)
	return m
}

func categoryOf(code string) string {
	return string([]byte(code)[0]) + string([]byte(code)[1])
}

func (m Merchant) AvailableTradeGoods() []string {
	return m.availableTGcodes
}

type Seller interface {
	ProposeSell(string) Contract
}

//ProposeSell - Информация о передаче товара от игрока к купцу
func (m Merchant) ProposeSell(code string) Contract {
	//dealDice := 10 //dice.Roll3D()
	dealDice := m.tradeDice
	saleDMmap := getSaleDMmap(code)
	for i := range m.localTC {
		if val, ok := saleDMmap[m.localTC[i]]; ok {
			dealDice = dealDice + val
		}
	}
	return NewContract(1, code, dealDice+m.prices[categoryOf(code)])
}

type Buyer interface {
	ProposeBuy(string) Contract
}

//ProposeBuy - Информация о передаче товара от купца к игроку
func (m Merchant) ProposeBuy(code string) Contract {
	dealDice := m.tradeDice
	purchaseDMmap := getPurchaseDMmap(code) // + dice.Roll("2d6").SumStr())
	for i := range m.localTC {
		if val, ok := purchaseDMmap[m.localTC[i]]; ok {
			dealDice = dealDice + val
		}
	}
	return NewContract(2, code, dealDice+m.prices[categoryOf(code)])
}

func (m Merchant) EncodeContract(code string, cType int) string {
	//#[tgCode]-[die][cType][ta][te][TC]#
	cCode := "#"
	ta := string([]byte(m.localUWP)[5])
	te := string([]byte(m.localUWP)[6])
	tl := string([]byte(m.localUWP)[8])
	cCode += code + "-" + TrvCore.DigitToEhex(m.tradeDice) + TrvCore.DigitToEhex(cType) + ta + te + tl
	for i := range m.localTC {
		cCode += m.localTC[i]
	}
	return cCode
}

func (m Merchant) MakeOffer(code string, operation int) []Contract {

	var allCont []Contract
	maxTons := RollMaximumForCategory(code)
	fmt.Println("maxTons", maxTons)
	exactVolume := make(map[string]int)
	for i := 0; i < maxTons; i++ {
		exactVolume[code+dice.Roll("2d6").SumStr()]++
	}
	table := prettytable.New()
	table.AddRow([]string{"Category", "Operation", "Base Price", "Lot", "Price", "Trade Dice"})
	for k, v := range exactVolume {
		c := Contract{}
		c.lotCode = k
		c.volume = v
		c.cType = operation
		c.contractDice = m.tradeDice
		c.taxingAgent = string([]byte(m.localUWP)[5])
		c.taxingEnviroment = string([]byte(m.localUWP)[6])
		c.lotDescription = getDescription(k)
		c.category = getCategory(k)
		table.AddRow(c.ShowShort())
		allCont = append(allCont, c)
	}
	table = prettytable.InsertSeparatorRow(table, 1)
	table.PTPrint()
	fmt.Println("terminal.GetSize(int(os.Stdout.Fd()))")
	fmt.Println(terminal.GetSize(int(os.Stdout.Fd())))
	return allCont
}
