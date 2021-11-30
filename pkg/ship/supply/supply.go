package supply

const (
	CATEGORY_FOOD = 0 iota
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
	capacity      int
	currentConsumption  int
	expectedConsumption int
}


type supplyCache struct {
	byType 			map[int]*material
	cacheVolume 	int
	cargoConsumed 	int
	logisticsEffect int
}

func NewSupplyCache(volume int) *supplyCache {
	sc := supplyCache{}
	sc.byType = make(map[int]*material)
	sc.byType[CATEGORY_FOOD] = &material{CATEGORY_FOOD}
	sc.byType[CATEGORY_MATERIALS_COMMON] = &material{CATEGORY_MATERIALS_COMMON}
	sc.byType[CATEGORY_MATERIALS_RARE] = &material{CATEGORY_MATERIALS_RARE}
	sc.byType[CATEGORY_BIOLOGICALS_RARE] = &material{CATEGORY_BIOLOGICALS_RARE}
	sc.byType[CATEGORY_MATERIALS_EXOTIC] = &material{CATEGORY_MATERIALS_EXOTIC}
	sc.cacheVolume = volume
}

func (sc *supplyCache) Remainder() string {
	list := []int{CATEGORY_FOOD,CATEGORY_MATERIALS_COMMON,CATEGORY_MATERIALS_RARE,CATEGORY_BIOLOGICALS_RARE,CATEGORY_MATERIALS_EXOTIC}
	total := 0
	for _, val := range sc.byType {
		total += val
	}
	r := "===SUPPLY REPORT============\n"
	for i, val := range list {
		switch i {
		case CATEGORY_FOOD:
			r += fmt.Sprintf("  Food: %v Units")
		}
	}
	return ""
}
