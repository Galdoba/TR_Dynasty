package supply

import (
	"fmt"
	"testing"
)

func TestSupply(t *testing.T) {
	sc := NewSupplyCache(740, 34)
	sc.AddSupplies(CATEGORY_COMMON, 1040)
	for i := 0; i < 2; i++ {
		fmt.Println(sc.RemainderText())
		if err := sc.ConsumeDaily(50); err != nil {
			t.Errorf("error: %v on try %v", err.Error(), i)
		} else {
			fmt.Println("consumption SUCCESSFUL for", i+1, "day")
			fmt.Println("consumed", daylyConsumption(sc))

		}
		fmt.Println(sc.Remainder())
	}
}
