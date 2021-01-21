package wrld

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
)

//Population - Возвращает примерное население
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

//PopDigit - Возвращает множитель насиления
func (w World) PopDigit() int {
	pd, err := strconv.Atoi(string(byte(w.PBG()[0])))
	if err != nil {
		panic(0)
	}
	return pd
}

//Belts - возвращает кол-во астеройдных поясов
func (w World) Belts() int {
	be, err := strconv.Atoi(string(byte(w.PBG()[1])))
	if err != nil {
		panic(0)
	}
	return be
}

//GasGigants  -возвращает кол-во газовых гигантов
func (w World) GasGigants() int {
	gg, err := strconv.Atoi(string(byte(w.PBG()[2])))
	if err != nil {
		panic(0)
	}
	return gg
}
