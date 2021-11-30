package supply

import "fmt"

const (
	CATEGORY_FOOD = iota
	CATEGORY_MATERIALS_COMMON
	CATEGORY_MATERIALS_RARE
	CATEGORY_BIOLOGICALS_RARE
	CATEGORY_MATERIALS_EXOTIC
)

/*
потребление припасов на команду:
1 SU на 3 человек в день
потребление корабля:
1% тоннажа в день

*/

type material struct {
	category            int
	capacity            int
	currentConsumption  int
	expectedConsumption int
}

type supplyCache struct {
	byType          map[int]*material
	cacheVolume     int
	cargoConsumed   int
	logisticsEffect int
	shipTonnage     int
	activeCrew      int
	consumptionRate int
}

func NewSupplyCache(tonnage, crew int) *supplyCache {
	sc := supplyCache{}
	sc.byType = make(map[int]*material)
	sc.byType[CATEGORY_FOOD] = &material{CATEGORY_FOOD, 0, 0, 0}
	sc.byType[CATEGORY_MATERIALS_COMMON] = &material{CATEGORY_MATERIALS_COMMON, 0, 0, 0}
	sc.byType[CATEGORY_MATERIALS_RARE] = &material{CATEGORY_MATERIALS_RARE, 0, 0, 0}
	sc.byType[CATEGORY_BIOLOGICALS_RARE] = &material{CATEGORY_BIOLOGICALS_RARE, 0, 0, 0}
	sc.byType[CATEGORY_MATERIALS_EXOTIC] = &material{CATEGORY_MATERIALS_EXOTIC, 0, 0, 0}
	sc.shipTonnage = tonnage
	sc.cacheVolume = tonnage
	sc.activeCrew = crew
	sc.consumptionRate = 100

	return &sc
}

func str(category int) string {
	switch category {
	case CATEGORY_FOOD:
		return "Food              "
	case CATEGORY_MATERIALS_COMMON:
		return "Materials (Common)"
	case CATEGORY_MATERIALS_RARE:
		return "Materials (Rare)  "
	case CATEGORY_BIOLOGICALS_RARE:
		return "Biologicals (Rare)"
	case CATEGORY_MATERIALS_EXOTIC:
		return "Materials (Exotic)"
	}
	return "Unknown"
}

func (sc *supplyCache) add(category, volume int) {
	sc.byType[category].capacity = sc.byType[category].capacity + volume
}

func (sc *supplyCache) Remainder() string {
	list := []int{CATEGORY_FOOD, CATEGORY_MATERIALS_COMMON, CATEGORY_MATERIALS_RARE, CATEGORY_BIOLOGICALS_RARE, CATEGORY_MATERIALS_EXOTIC}
	total := 0
	for _, val := range sc.byType {
		total += val.capacity
	}
	r := "===SUPPLY REPORT==================\n"
	for _, category := range list {
		switch sc.byType[category].capacity > 0 {
		case true:
			r += fmt.Sprintf("  %v: %v Units\n", str(category), sc.byType[category].capacity)
		}
	}
	r += fmt.Sprintf("              Total: %v/%v SU\n", total, sc.cacheVolume)
	r += "=================================="
	return r
}

func (sc *supplyCache) ConsumedForDay() int {
	perShip := sc.shipTonnage / 100
	if sc.shipTonnage%100 > 0 {
		perShip++
	}
	perCrew := sc.activeCrew/3 + 1
	if sc.activeCrew%3 > 0 {
		perCrew++
	}
	toConsume := perShip + perCrew
	toConsume = (toConsume * sc.consumptionRate) / 100
	switch toConsume <= sc.byType[CATEGORY_MATERIALS_COMMON].capacity {
	case true:
		sc.byType[CATEGORY_MATERIALS_COMMON].capacity = sc.byType[CATEGORY_MATERIALS_COMMON].capacity - toConsume
		return sc.consumptionRate
	case false:
		canConsume := sc.byType[CATEGORY_MATERIALS_COMMON].capacity
		sc.byType[CATEGORY_MATERIALS_COMMON].capacity = 0
		return canConsume * 100 / toConsume
	}
	return 0
}
