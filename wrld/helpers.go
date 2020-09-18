package wrld

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
)

func (w World) Population() int {
	pops := w.GetСharacteristic(constant.PrPops).Value()
	if pops == 0 {
		return 0
	}
	pDigit, err := strconv.Atoi(string(byte(w.PBG()[0])))
	if err != nil {
		panic(0)
	}
	for i := 1; i <= pops; i++ {
		pDigit = pDigit * 10
	}
	return pDigit
}
