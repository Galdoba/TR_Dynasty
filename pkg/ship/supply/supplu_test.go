package supply

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func TestSupply(t *testing.T) {
	//for err := cache.Consume(days); err != nil {
	//	handle(err)
	//}

	sc := NewSupplyCache(75000, 488)

	sc.add(CATEGORY_MATERIALS_COMMON, 75000)
	sc.SetConsumptionRate(110, 1)
	// cons := sc.ConsumedForDay()
	// if cons > 100 || cons < 0 {
	// t.Errorf("consumed %v but expected 0-100\n", cons)
	// }
	for i := 0; i < 20; i++ {
		if err := sc.ConsumeDaily(dice.Roll1D()); err != nil {
			t.Errorf("error: %v on try %v", err.Error(), 20-i)
		} else {
			fmt.Println("consumption SUCCESSFUL")
		}
		fmt.Println(sc.Remainder())
	}

	//fmt.Println(sc.ConsumedForDay(), "% consumed")

}
