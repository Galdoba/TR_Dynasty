package wrld

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
)

// //TCer - механика контроля данных торговых кодов
// type TCer interface {
// 	TradeClassifications() string
// 	TradeCodes() []string
// }

// type tc struct {
// 	data string
// }

//TradeCodes - возвращает торговые коды без коментариев с добавлением Lt и Ht
func (w World) TradeCodes() []string {
	data := strings.Split(w.TradeClassifications(), " ")
	tc := []string{}
	for i := range data {
		if len(data[i]) != 2 {
			continue
		}
		tc = append(tc, data[i])
	}
	for _, val := range checkLtHtTradeCodes(w) {
		tc = append(tc, val)
	}
	return tc
}

func checkLtHtTradeCodes(w World) []string {
	add := []string{}
	if w.GetСharacteristic(constant.PrTL).Value() <= TrvCore.EhexToDigit("5") {
		add = append(add, constant.TradeCodeLowTech)
	}
	if w.GetСharacteristic(constant.PrTL).Value() >= TrvCore.EhexToDigit("C") {
		add = append(add, constant.TradeCodeHighTech)
	}
	return add
}
