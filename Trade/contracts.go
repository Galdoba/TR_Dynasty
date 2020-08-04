package Trade

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
)

//Contract - описываемые свойства структуры с информацией о договоре с участием денег
type Contract struct {
	cType            int
	lotCode          string
	lotDescription   string
	contractDice     int
	volume           int
	taxingAgent      string
	taxingEnviroment string
	timeLimit        string
}

/*
Contract Code
#[tgCode]-[volume]-[die][cType][ta][te][TC]#
*/

//NewContract - создает контракт (структура с информацией о договоре с участием денег)
func NewContract(cType int, lotCode string, cntrctDie int) Contract {
	c := Contract{}
	c.cType = cType
	c.lotCode = lotCode // + dice.Roll("2d6").SumStr()
	c.contractDice = cntrctDie
	c.taxingAgent = "0"
	return c
}

func encodeContractCode(tgCode string, cntrctDie int, tgInfo []string) (cntrctCode string) {
	cntrctCode += tgCode + "-" + TrvCore.DigitToEhex(cntrctDie)
	if len(tgInfo) > 0 {
		cntrctCode += "-"
	}
	for i := range tgInfo {
		cntrctCode += tgInfo[i]
	}
	return cntrctCode + "-cc"
}

func randomEHEX() string {
	return TrvCore.DigitToEhex(dice.Roll("1d33").Sum())
}

func tradeMark() string {
	str := ""
	for i := -4; i < dice.Roll("1d8").Sum(); i++ {
		str = str + randomEHEX()
	}
	return str
}

func (c Contract) Negotiate(effect int) Contract {
	c.contractDice = c.contractDice + effect
	return c
}

func (c Contract) SetTaxingAgent(ta string) Contract {
	return c
}

func (c Contract) ShowShort() string {
	short := c.lotCode + "	"
	price := 0
	//vol := c.volume
	switch c.cType {
	default:
		short += "TODO: cType " + strconv.Itoa(c.cType)
	case 1:
		short += "SELL	"
		price = modifyPriceSale(getBasePrice(c.lotCode), c.contractDice)
	case 2:
		short += "BUY	"
		price = modifyPricePurchase(getBasePrice(c.lotCode), c.contractDice)
	}
	//short += strconv.Itoa(vol) + " tons 	"
	short += strconv.Itoa(price) + "	"
	short += getDescription(c.lotCode)

	return short
}

func (c Contract) String() string {
	cInfo := ""
	cInfo += "Contract type: " + cTypeToString(c.cType) + "\n"
	cInfo += "Lot code: #" + c.lotCode + "-" + TrvCore.DigitToEhex(c.contractDice) + tradeMark() + "\n"
	cInfo += "Lot description: " + getDescription(c.lotCode) + " (" + strconv.Itoa(getBasePrice(c.lotCode)) + ")\n"
	cInfo += "Lot price: "
	price := 0
	tax := 0
	switch c.cType {
	case 1:
		price = modifyPriceSale(getBasePrice(c.lotCode), c.contractDice)
		cInfo += strconv.Itoa(price)
		tax = taxingAmount(c.volume*(price-getBasePrice(c.lotCode)), "4")
	case 2:
		price = modifyPricePurchase(getBasePrice(c.lotCode), c.contractDice)
		cInfo += strconv.Itoa(price)
	}
	if c.volume > 0 {
		cInfo += " Cr per Unit\n"
		cInfo += "Lot volume: " + strconv.Itoa(c.volume) + " Units\n"
		cInfo += "Total Price: " + strconv.Itoa(c.volume*price) + " Cr"
	}
	cInfo += "\n"
	if tax > 0 {
		cInfo += "Tax: " + strconv.Itoa(tax) + " Cr\n"
	}
	cInfo += "----------------------\n"
	cInfo += "Projected Profit: " + strconv.Itoa(price-tax) + " Cr\n"
	return cInfo
}

func (c Contract) SetVolume(newVolume int) Contract {
	c.volume = newVolume
	return c
}

func cTypeToString(i int) string {
	ct := "*UNDEFINED*"
	switch i {
	case 1:
		ct = "SELL"
	case 2:
		ct = "BUY"
	}
	return ct
}

func taxingAmount(profit int, ta string) int {
	taxinGrade := -1
	if profit <= 0 {
		return 0
	}
	for _, val := range []int{0, 1000, 5000, 10000, 25000, 50000, 75000, 100000, 250000, 1000000} {
		if profit > val {
			taxinGrade++
		}
	}
	taxingMap := mapTaxRate()
	taxShare := taxingMap[ta][taxinGrade]
	if taxShare == -1 {
		taxShare = 0
	}
	return taxedFrom(profit, mapTaxRate()[ta][taxinGrade])
}

func taxedFrom(base int, proc int) int {
	return int(float64(base) * (float64(proc) / 100.0))
}

func mapTaxRate() map[string][]int {
	trMap := make(map[string][]int)
	trMap["0"] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	trMap["1"] = []int{6, 8, 10, 12, 12, 12, 15, 15, 22, 25}
	trMap["2"] = []int{3, 3, 5, 5, 10, 10, 10, 12, 12, 18}
	trMap["3"] = []int{8, 8, 10, 10, 10, 12, 12, 12, 14, 14}
	trMap["4"] = []int{5, 6, 8, 10, 12, 14, 18, 20, 22, 25}
	trMap["5"] = []int{4, 4, 6, 8, 8, 10, 10, 12, 14, 16}
	trMap["6"] = []int{8, 10, 12, 14, 16, 20, 20, 20, 22, 25}
	trMap["7"] = []int{dice.Roll1D(), dice.Roll1D(), dice.Roll2D(), dice.Roll2D(), dice.Roll3D(), dice.Roll3D(), dice.Roll3D(), dice.Roll4D(), dice.Roll4D(), dice.Roll5D()}
	trMap["8"] = []int{6, 6, 6, 8, 8, 10, 10, 12, 14, 16}
	trMap["9"] = []int{5, 5, 8, 8, 10, 10, 12, 14, 15, 18}
	trMap["A"] = []int{12, 12, 12, 12, 12, 12, 12, 12, 12, 12}
	trMap["B"] = []int{10, 10, 10, 20, 20, 20, 20, 30, 30, 30}
	trMap["C"] = []int{3, 5, 5, 8, 8, 8, 10, 10, 10, 12}
	trMap["D"] = []int{10, 10, 12, 12, 14, 14, 15, 15, 15, 18}
	trMap["E"] = []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}  //Not in book (Merchant Prince)
	trMap["F"] = []int{13, 15, 17, 19, 21, 23, 25, 27, 29, 35}  //Not in book (Merchant Prince)
	trMap["CM"] = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1} //Ключ для вызова функции CrimeTax()

	return trMap
}
