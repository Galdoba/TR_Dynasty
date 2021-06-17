package mithril

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/devtools/cli/user"
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
	ws.dicepool = dice.New().SetSeed("Mithril")
	return ws
}

func TestWeather() {
	ws := NewWeather()
	for i := 0; i < 1000; i++ {
		fmt.Println("---------", "day", i+1, "---------")

		ws.RollWeatherConditions()
		ws.RollStorm()
		ws.rollIncident()
		if _, err := user.Confirm("Press Enter for next day"); err != nil {
			fmt.Println("\r                                ")
		}
	}

}

func (ws *WeatherState) rollIncident() {
	fmt.Println("Incedent Roll:", ws.dicepool.RollNext("2d6").Sum(), "(more info on page 20)")
	fmt.Println(" ")
	fmt.Println("Ice                    Difficult (10+)")
	fmt.Println("Broken Ice             Average (8+)")
	fmt.Println("Plains                 Very Difficult (12+)")
	fmt.Println("Snow Plains            Difficult (10+)")
	fmt.Println("Broken Terrain         Routine (6+)")
	fmt.Println("Mountains              Easy (4+)")
	fmt.Println("Seacoast or Lake       Difficult (10+)")
	fmt.Println("Open Water             Average (8+)")
	fmt.Println("Apply the following DMs:")
	fmt.Println("	• Driver’s Drive (wheel) ................................. -skill level")
	fmt.Println("	• Highest Recon or Navigation skill ...................... -skill level")
	fmt.Println("	• For every hour past two a driver remains at the controls +1")
	fmt.Println("	• For every hour past five in the")
	fmt.Println("	  last twenty the driver remains at the controls ......... +1")
	fmt.Println("	• Heavy snow falling ..................................... +2")
	fmt.Println("	• Driving at night ....................................... +2")
	fmt.Println("	• Travelling fast ........................................ +2 ")
	fmt.Println("	• Driving recklessly ..................................... +5")
	fmt.Println("If an incident occurs, the referee should roll 1D for the ")
	fmt.Println("nature of the hazard and consult the Incident table.  (p.20)")

}

func (ws *WeatherState) RollStorm() {
	r := ws.dicepool.RollNext("2d6").DM(ws.daysSinceStorm).Sum()
	precipation := "None"
	ws.daysSinceStorm++
	switch {

	case ws.dicepool.ResultIs("13+"):
		precipation = "SERIOUS STORM!!"
		if ws.dicepool.RollNext("1d6").Sum() == 1 {
			precipation += " (accompanied by spectacular lightning)"
		}
		ws.daysSinceStorm = 0
	case ws.dicepool.ResultIs("7+"):
		precipation = "Significant"
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
	fmt.Print("Weather    :  ")
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
