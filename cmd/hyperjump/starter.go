package hyperjump

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/devtools/cli/user"
)

func StartJumpEvent(w wrld.World) {
	diametr := 1600 * w.GetСharacteristic(constant.PrSize).Value()
	effA := userInputInt("Astrogation Check effect: ")
	effE := userInputInt("Engineering Check effect: ")
	hj := New(effA, effE, diametr)
	fmt.Println(hj.Report())
	fmt.Println(hj.Outcome())
}

func userInputInt(msg ...string) int {
	str := userInputStr(msg...)
	i, err := strconv.Atoi(str)
	for err != nil {
		fmt.Println(err.Error() + "\n")
		str = userInputStr(msg...)
		i, err = strconv.Atoi(str)
	}
	return i
}

func userInputStr(msg ...string) string {
	for i := range msg {
		fmt.Println(msg[i])
	}
	str, err := user.InputStr()
	if err != nil {
		fmt.Println(err.Error())
		return err.Error() + "\n"
	}
	return str
}
