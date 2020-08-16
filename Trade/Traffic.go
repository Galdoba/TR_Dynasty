package trade

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Galdoba/TR_Dynasty/dice"

	"github.com/Galdoba/TR_Dynasty/TrvCore"

	"github.com/Galdoba/TR_Dynasty/world"

	"github.com/Galdoba/devtools/cli/features"
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

func userInputUWP() string {
	done := false
	fmt.Println("Введите коррдинаты и UWP текущего мира (пример: '0631 В555555-5'):")
	for !done {
		uwp, err := user.InputStr()
		if err != nil {
			features.TypingSlowly(err.Error(), 15)
			continue
		}
		return uwp
	}
	return "Must not happen !!"
}

func RunTraffic() {
	dp = dice.New(1105018)
	sourceWorld := world.NewWorld("Source World").SetUWP(userInputUWP())
	sp := sourceWorld.StarPort()
	pops := sourceWorld.PlanetaryData("Pops")
	pFactor := 0
	pFactor += pFactorPops(pops)
	pFactor += pFactorSp(sp)
	fmt.Println(sourceWorld, sp, pops, pFactor)
	fmt.Println("Low Pass", rollPassengerTrafficTable(pFactor+1))
	fmt.Println("Mid Pass", rollPassengerTrafficTable(pFactor))
	fmt.Println("Basic Pass", rollPassengerTrafficTable(pFactor))
	fmt.Println("High Pass", rollPassengerTrafficTable(pFactor-4))

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

func pFactorSp(sp string) int {
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
	return dp.RollNext(strconv.Itoa(d) + "d6").Sum()
}
