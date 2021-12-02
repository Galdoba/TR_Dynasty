package supply

import (
	"fmt"

	"github.com/Galdoba/utils"
)

const (
	CATEGORY_COMMON = iota
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
	sc.byType[CATEGORY_COMMON] = &material{CATEGORY_COMMON, 0, 0, 0}
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
	case CATEGORY_COMMON:
		return "Supply Units (SU)"
	case CATEGORY_MATERIALS_RARE:
		return "Materials (Rare)  "
	case CATEGORY_BIOLOGICALS_RARE:
		return "Biologicals (Rare)"
	case CATEGORY_MATERIALS_EXOTIC:
		return "Materials (Exotic)"
	}
	return "Unknown"
}

//AddSupplies - добавляет в кэш соответсвующий ресурс
func (sc *supplyCache) AddSupplies(category, volume int) {
	sc.byType[category].capacity = sc.byType[category].capacity + volume
}

//SetConsumptionRate - задает уровень расхода припасов
func (sc *supplyCache) SetConsumptionRate(rate int) {
	sc.consumptionRate = rate
}

//SetLogisticsEfect - задает эффект проверки Admin(8) модифицирующей расход SU на 2,5% на период
func (sc *supplyCache) SetLogisticsEfect(eff int) {
	sc.logisticsEffect = eff
}

func roundUp(f float64) float64 {
	i := int(f)
	if f-float64(i) > 0 {
		i++
	}
	return utils.RoundFloat64(float64(i), 0)
}

//RemainderText - возвращает слайс с содержимым кэша:
//rem[0] = текущие SU
//rem[1] = текущие Rare Materials
//rem[2] = текущие Rare Biologicals
//rem[3] = текущие Exotic Materials
//rem[4] = текущие Общий объем кэша
//rem[5] = текущие примерный запас на колличество дней при текущем расходе
func (sc *supplyCache) Remainder() []int {
	list := []int{CATEGORY_COMMON, CATEGORY_MATERIALS_RARE, CATEGORY_BIOLOGICALS_RARE, CATEGORY_MATERIALS_EXOTIC}
	remainder := []int{}
	total := 0
	for _, val := range list {
		total += sc.byType[val].capacity
		remainder = append(remainder, sc.byType[val].capacity)
	}
	remainder = append(remainder, total)
	days := remainder[0] / daylyConsumption(sc)
	remainder = append(remainder, days)
	return remainder
}

func suppliesRemained(sc *supplyCache) int {
	return sc.Remainder()[0]
}

//RemainderText - возвращает рапорт для игроков и рефери
func (sc *supplyCache) RemainderText() string {
	rem := sc.Remainder()
	r := "===SUPPLY REPORT==================\n"
	r += fmt.Sprintf("%v : %v/%v\n", str(CATEGORY_COMMON), rem[0], sc.cacheVolume)
	if rem[1] > 0 {
		r += fmt.Sprintf("%v: %v\n", str(CATEGORY_MATERIALS_RARE), rem[1])
	}
	if rem[2] > 0 {
		r += fmt.Sprintf("%v: %v\n", str(CATEGORY_BIOLOGICALS_RARE), rem[2])
	}
	if rem[3] > 0 {
		r += fmt.Sprintf("%v: %v\n", str(CATEGORY_MATERIALS_EXOTIC), rem[3])
	}
	if rem[4] > sc.cacheVolume {
		r += fmt.Sprintf("        Cargo used: %v tons\n", (rem[4]-sc.cacheVolume)/100+1)
	}
	///////////////////////////////
	r += "==================================\n"
	r += fmt.Sprintf("At this level of consumption supplies expected to be last for %d days", rem[5])

	return r
}

func daylyConsumption(sc *supplyCache) int {
	perShip := sc.shipTonnage / 100
	if sc.shipTonnage < 100 {
		perShip++
	}
	perCrew := (sc.activeCrew) / 3
	if sc.activeCrew%3 > 0 {
		perCrew++
	}
	toConsume := perShip + perCrew
	toConsume = (toConsume * sc.consumptionRate) / 100           //modify by Consumption rate
	toConsume = (toConsume * (200 - sc.logisticsEffect*5)) / 200 //modify by Logistic Effect
	return toConsume
}

func (sc *supplyCache) ConsumeDaily(days int) error {
	if days < 1 {
		return fmt.Errorf("days = '%v' (can't be less than 1)", days)
	}
	toBeConsumed := daylyConsumption(sc) * days
	if toBeConsumed > suppliesRemained(sc) {
		return fmt.Errorf("can't consume %v SU (available %v SU)", toBeConsumed, suppliesRemained(sc))
	}
	sc.byType[CATEGORY_COMMON].capacity = sc.byType[CATEGORY_COMMON].capacity - toBeConsumed
	return nil
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
