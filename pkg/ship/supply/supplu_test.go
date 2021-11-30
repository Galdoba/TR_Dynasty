package supply

import (
	"fmt"
	"testing"
)

func TestSupply(t *testing.T) {
	sc := NewSupplyCache(75000, 488)
	//sc.add(CATEGORY_FOOD, 100)
	sc.add(CATEGORY_MATERIALS_COMMON, 900)
	sc.add(CATEGORY_MATERIALS_RARE, 50)
	fmt.Println(sc.Remainder())
	for i := -5; i < 1500; i++ {
		sc := NewSupplyCache(75000, 488)
		sc.add(CATEGORY_MATERIALS_COMMON, 900)
		cons := sc.ConsumedForDay()
		if cons > 100 || cons < 0 {
			t.Errorf("consumed %v but expected 0-100 (input %v)\n", cons, i)
		}
	}
	fmt.Println(sc.ConsumedForDay(), "% consumed")
	fmt.Println(sc.Remainder())
}
