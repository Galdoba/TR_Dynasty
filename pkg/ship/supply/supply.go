package supply

import "fmt"

const (
	CATEGORY_MATERIALS_COMMON = iota
	CATEGORY_MATERIALS_RARE
	CATEGORY_BIOLOGICALS_RARE
	CATEGORY_MATERIALS_EXOTIC
	CONSUMPTION_RATE_Lavish = 101
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

func (sc *supplyCache) SetConsumptionRate(rate, adminEffect int) {
	rtFl := float64(rate) * (float64(adminEffect)*0.05 + 1.0)
	rt := int(rtFl) //надо округлить вверх
	fmt.Println("New cr = ", rt)
	sc.consumptionRate = rt
}

func roundUp(f float64) float64 {
	i := int(f)
	if int(f)-int(i) == 0 {

	}
}

func (sc *supplyCache) Remainder() []int {
	list := []int{CATEGORY_MATERIALS_COMMON, CATEGORY_MATERIALS_RARE, CATEGORY_BIOLOGICALS_RARE, CATEGORY_MATERIALS_EXOTIC}
	remainder := []int{}
	total := 0
	for _, val := range list {
		total += sc.byType[val].capacity
		remainder = append(remainder, sc.byType[val].capacity)
	}
	remainder = append(remainder, total)
	return remainder
}

func (sc *supplyCache) RemainderText() string {
	list := []int{CATEGORY_MATERIALS_COMMON, CATEGORY_MATERIALS_RARE, CATEGORY_BIOLOGICALS_RARE, CATEGORY_MATERIALS_EXOTIC}
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
	r += fmt.Sprintf("               Total: %v/%v SU\n", total, sc.cacheVolume)
	if total > sc.cacheVolume {
		r += fmt.Sprintf("          Cargo used: %v tons\n", (total-sc.cacheVolume)/100+1)
	}
	r += "=================================="
	return r
}

func (sc *supplyCache) ConsumedForDay() int {
	perShip := sc.shipTonnage / 100
	if sc.shipTonnage < 100 {
		perShip++
	}
	perCrew := (sc.activeCrew) / 3
	if sc.activeCrew%3 > 0 {
		perCrew++
	}
	fmt.Println(perShip, perCrew)
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

func (sc *supplyCache) ConsumeDaily(days int) error {
	if days < 1 {
		return fmt.Errorf("days = '%v' (can't be less than 1)", days)
	}
	perShip := sc.shipTonnage / 100
	if sc.shipTonnage < 100 {
		perShip++
	}
	perCrew := (sc.activeCrew) / 3
	if sc.activeCrew%3 > 0 {
		perCrew++
	}
	toConsume := perShip + perCrew
	toConsume = (toConsume * sc.consumptionRate) / 100
	switch (toConsume * days) <= sc.byType[CATEGORY_MATERIALS_COMMON].capacity {
	case true:
		sc.byType[CATEGORY_MATERIALS_COMMON].capacity = sc.byType[CATEGORY_MATERIALS_COMMON].capacity - (toConsume * days)
		return nil
	case false:
		return fmt.Errorf("not enough supplies in the cache for %v days", days)
	}
	return fmt.Errorf("unspecified error")
}

/*
Skander
4000 SU

40 - ship
17 - crew

5700 - day
159600 - 4 weeks

Система
3D-(test10)*40   - от 0 до 960 su (0 - 96000 cr) за 4 - 24 Часа (1 день)

Структура (HP или Armor)
1D%    - 40 - 240 SU
восстанавливает 1D*effectTest8 за 2D часов (2 попытки в день)

Overhaul
200% тонажа * effectAdmin10 за 12+2D дней (требуется раз в год)
800000 Cr



Норка

8 = 24000Cr/month
0 - 24*8 = 0 - 19200 cr

8 - 48 == 800 - 4800 Cr

1480 SU == 148000 Cr




*/
