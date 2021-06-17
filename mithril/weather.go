package mithril

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	weatherWARM        = "The temperature is 4-8 C"
	weatherCOOL        = "The temperature is 0-4 C"
	weatherCOLD        = "The temperature is -20 C - 0 C"
	weatherVERYCOLD    = "The temperature is -40 C - -20 C"
	weatherEXTREMECOLD = "The temperature is -50 C and lower"
)

type WeatherState struct {
	dicepool       *dice.Dicepool
	conditions     string
	daysSinceStorm int
}

func NewWeather() WeatherState {
	ws := WeatherState{}
	ws.dicepool = dice.New().SetSeed("Mithril2")
	return ws
}

func TestWeather() {
	ws := NewWeather()
	for i := 0; i < 100; i++ {
		fmt.Println("---------", "day", i+1, "---------")
		fmt.Print("Weather    :  ")
		ws.RollWeatherConditions()
		ws.RollStorm()
	}

}

func (ws *WeatherState) RollStorm() {
	r := ws.dicepool.RollNext("2d6").DM(ws.daysSinceStorm).Sum()
	precipation := "None"
	ws.daysSinceStorm++
	switch {

	case ws.dicepool.ResultIs("13+"):
		precipation = "serious storm"
		if ws.dicepool.RollNext("1d6").Sum() == 1 {
			precipation += " (accompanied by spectacular lightning)"
		}
		ws.daysSinceStorm = 0
	case ws.dicepool.ResultIs("7+"):
		precipation = "significant precipitation that day"
	}
	fmt.Println("Precipation: ", r, "	|", precipation)
	fmt.Println("Days since Storm:", ws.daysSinceStorm)
}

func (ws *WeatherState) RollWeatherConditions() {
	dms := 0
	if ws.daysSinceStorm == 0 {
		dms = -4
	}
	//r := ws.dicepool.RollNext("2d6").DM(dms).Sum()
	ws.dicepool.RollNext("2d6").DM(dms).Sum()
	weather := "UNDEFINED (error)"
	switch {
	case ws.dicepool.ResultIs("2-"):
		weather = weatherEXTREMECOLD
	case ws.dicepool.ResultIs("3 5"):
		weather = weatherVERYCOLD
	case ws.dicepool.ResultIs("6 9"):
		weather = weatherCOLD
	case ws.dicepool.ResultIs("10 11"):
		weather = weatherCOLD
	case ws.dicepool.ResultIs("12+"):
		weather = weatherWARM
	}

	fmt.Println(ws.dicepool.Sum(), "	|", weather)
}

/*
2-
-2-
10-11
10 11
12+
12+
-2+
-4--2
-4 -2

1. Проверка наличия '+' или '-' в конце предложения, означает "результат или больше/меньше" - условие1
2. Очищаем от условия1 (если да то игнорируем шаги 3 и 4)
3. Проверка наличия '-' в середине предложения, если да то ожидаем слайс из 2 позиция каждая из которых переводится в Int
4. Создаем временный слайс с Int от миннимума до максимума (проверка правильности слайса)
5. сверяем r с числами из временного слайса.

2D		Result
2-		Extremely Cold
3-5		Very Cold
6-9		Cold
10-11	Cool
12+		Warm


*/
