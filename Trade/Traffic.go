package trade

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/TrvCore"

	"github.com/Galdoba/TR_Dynasty/world"

	"github.com/Galdoba/devtools/cli/user"
)

/*

TZ:
1. определения фактора P (Passenger)
4 броска 2Д по таблице Passenger Traffic (CRB p207)

Влияющие модификаторы:
команда (вводится пользователем):
The Effect of a Broker, Carouse or Streetwise check

Chief Steward DM+ highest Steward skill on ship

Планета (идет от UWP планеты)
Rolling for High Passengers DM-4
Rolling for Low Passengers DM+1

World Population 1 or less DM-4
World Population 6-7 DM+1
World Population 8 or more DM+3

Starport A DM+2
Starport B DM+1
Starport E DM-1
Starport X DM-3

Amber Zone DM+1
Red Zone DM-4
//////////////////////////////////

2. определение Фактора F (Freight)
3 броска 2Д по таблице Freight Traffic (CRB p208)

Влияющие модификаторы:
команда (вводится пользователем):
The Effect of a Broker or Streetwise check

Планета (идет от UWP планеты)
Rolling for Major Cargo DM-4
Rolling for Incidental Cargo DM+2

World Population 1 or less DM-4
World Population 6-7 DM+2
World Population 8 or more DM+4

Starport A DM+2
Starport B DM+1
Starport E DM-1
Starport X DM-3
Tech Level 6 or less DM-1
Tech Level 9 or more DM+2

Amber Zone DM-2
Red Zone DM-6

//////////////////////////////////////
3. определение фактора М (Mail)
1 бросок 2D + М-фактор

Влияющие модификаторы:

Планета (идет от Предыдущих вычислений)
Freight Traffic DM-10 or less: DM-2
Freight Traffic DM-9 to -5: DM-1
Freight Traffic DM-4 to +4: DM+0
Freight Traffic DM 5 to 9: DM+1
Freight Traffic DM 10 or more: DM+2

команда (вводится пользователем или забирается из конфига):
Travellers’ ship is armed: DM+2
+ Travellers’ highest Naval or Scout rank
+ Travellers’ highest Social Standing DM

Планета (идет от UWP планеты)
World is TL of 5 or less: DM-4

////////////////////////////////////////////
Флаги:
-v (verbose) - выводить информацию о бросках и вычислениях
-fv (full verbose) - тоже что и -v но еще и информация по сиду, времени броска и прочее
-ref (referee) - запрашивать дополнительные модификаторы
-h (help) - справка
-log - создать лог-файл и записать туда информацию от программы

tools:
нужен Логгер!!!


*/

var msDelay time.Duration
var dp *dice.Dicepool

func userInputStr(msg string) string {
	done := false
	fmt.Print(msg)
	for !done {
		uwp, err := user.InputStr()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return uwp
	}
	return "Must not happen !!"
}

func userInputInt(msg string) int {
	done := false
	fmt.Print(msg)
	for !done {
		i, err := user.InputInt()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return i
	}
	return -999
}

func RunTraffic() {
	fmt.Println("mgt2_Traffic  v1.0.0")
	seedStr := userInputStr("Введите дату (формат ДДД-ГГГГ): ")
	swUWP := userInputStr("Введите UWP текущего мира: ")
	sourceWorld := world.NewWorld("Source World").SetUWP(swUWP)
	dp = dice.New(utils.SeedFromString(seedStr + swUWP))
	sp := sourceWorld.StarPort()
	pops := sourceWorld.PlanetaryData("Pops")
	tl := sourceWorld.PlanetaryData("Tech")
	//law := sourceWorld.PlanetaryData("Law")
	//fmt.Println("////////////////////////////")
	fmt.Println("Поиск пассажиров...")
	pFactor := userInputInt("Введите эффект проверки Broker, Carouse или Streetwise: ")

	pFactor += pFactorPops(pops)
	pFactor += pfFactorSp(sp)
	//fmt.Println(sourceWorld, sp, pops, pFactor)

	lowPass := rollPassengerTrafficTable(pFactor + 1)
	midPass := rollPassengerTrafficTable(pFactor)
	basicPass := rollPassengerTrafficTable(pFactor)
	highPass := rollPassengerTrafficTable(pFactor - 4)
	//fmt.Println("////////////////////////////")
	fmt.Println("Поиск грузов...")
	fFactor := userInputInt("Введите эффект проверки Broker или Streetwise: ")
	fFactor += fFactorPops(pops)
	fFactor += pfFactorSp(sp)
	fFactor += fFactorTL(tl)
	inLots := rollFreightTrafficTable(pFactor + 2)
	minorLots := rollFreightTrafficTable(pFactor)
	majorLots := rollFreightTrafficTable(pFactor - 4)

	fmt.Println("Поиск почтовых лотов...")
	mFactor := mFactorFre(fFactor) + 7
	mFactor += mFactorTL(tl)
	mFactor += dp.RollNext("2d6").Sum()
	//fmt.Println("mFactor:", mFactor)
	mContainers := 0
	if mFactor >= 12 {
		mContainers = dp.RollNext("1d6").Sum()
	}
	fmt.Println("////////////////////////////")
	fmt.Println("ПАССАЖИРЫ:")
	fmt.Println("Первый класс:", highPass)
	fmt.Println("Бизнес класс:", midPass)
	fmt.Println("Эконом класс:", basicPass)
	fmt.Println("Спящий класс:", lowPass)
	fmt.Println("ГРУЗОПЕРЕВОЗКИ:")
	fmt.Println("  Больших лотов:", majorLots)
	fmt.Println("    Малых лотов:", minorLots)
	fmt.Println("Случайных лотов:", inLots)
	fmt.Println("ПОЧТА:")
	fmt.Println("доступно", mContainers, "контейнеров с почтой")
	fmt.Println("ДЕТАЛИ ЛОТОВ:")
	distributeLots(majorLots, minorLots, inLots)

}

func distributeLots(majorLots, minorLots, inLots int) {
	lotMap := make(map[int]int)
	dp.DM(0)
	for i := 0; i < majorLots; i++ {
		size := dp.RollNext("1d6").Sum() * 10
		lotMap[size]++
	}
	for i := 0; i < minorLots; i++ {
		size := dp.RollNext("1d6").Sum() * 5
		lotMap[size]++
	}
	for i := 0; i < inLots; i++ {
		size := dp.RollNext("1d6").Sum() * 1
		lotMap[size]++
	}
	sizes := []int{60, 50, 40, 30, 25, 20, 15, 10, 6, 5, 4, 3, 2, 1}
	for _, val := range sizes {
		if lotMap[val] != 0 {
			fmt.Println(val, "tons -", lotMap[val], "lots")
		}
	}
}

func pFactorPops(pops string) int {
	popInt := TrvCore.EhexToDigit(pops)
	switch popInt {
	case 0, 1:
		return -4
	case 6, 7:
		return 1
	default:
		if popInt >= 8 {
			return 3
		}
	}
	return 0
}

func pfFactorSp(sp string) int {
	switch sp {
	case "A":
		return 2
	case "B":
		return 1
	case "E":
		return -1
	case "X":
		return -3
	}
	return 0
}

func rollPassengerTrafficTable(pF int) int {
	pass := dp.RollNext("2d6").DM(pF).Sum()
	d := 0
	switch pass {
	default:
		if pass <= 1 {
			return 0
		}
		if pass >= 20 {
			d = 10
		}
	case 2, 3:
		d = 1
	case 4, 5, 6:
		d = 2
	case 7, 8, 9, 10:
		d = 3
	case 11, 12, 13:
		d = 4
	case 14, 15:
		d = 5
	case 16:
		d = 6
	case 17:
		d = 7
	case 18:
		d = 8
	case 19:
		d = 9
	}
	ps := dp.RollNext(strconv.Itoa(d) + "d6").Sum()
	if ps < 0 {
		ps = 0
	}
	return ps
}

func fFactorPops(pops string) int {
	popsInt := TrvCore.EhexToDigit(pops)
	switch popsInt {
	case 0, 1:
		return -4
	case 6, 7:
		return 2
	default:
		if popsInt >= 8 {
			return 4
		}
	}
	return 0
}

func fFactorTL(tl string) int {
	tlInt := TrvCore.EhexToDigit(tl)
	if tlInt <= 6 {
		return -1
	}
	if tlInt >= 9 {
		return 2
	}
	return 0
}

func rollFreightTrafficTable(fF int) int {
	fre := dp.RollNext("2d6").DM(fF).Sum()
	d := 0
	switch fre {
	default:
		if fre <= 1 {
			return 0
		}
		if fre >= 20 {
			d = 10
		}
	case 2, 3:
		d = 1
	case 4, 5:
		d = 2
	case 6, 7, 8:
		d = 3
	case 9, 10, 11:
		d = 4
	case 12, 13, 14:
		d = 5
	case 15, 16:
		d = 6
	case 17:
		d = 7
	case 18:
		d = 8
	case 19:
		d = 9
	}
	fr := dp.RollNext(strconv.Itoa(d) + "d6").Sum()
	if fr < 0 {
		fr = 0
	}
	return fr
}

func mFactorFre(fFactor int) int {
	if fFactor <= -10 {
		return -2
	}
	if fFactor <= -5 {
		return -1
	}
	if fFactor <= 4 {
		return 0
	}
	if fFactor <= 9 {
		return 1
	}
	return 2
}

func mFactorTL(tl string) int {
	tlInt := TrvCore.EhexToDigit(tl)
	if tlInt <= 5 {
		return -4
	}
	return 0
}
