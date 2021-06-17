package mithril

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	weatherWARM        = "The temperature is 4-8 C"
	weatherCOOL        = "The temperature is 0-4 C"
	weatherCOLD        = "The temperature is -20 - 0 C"
	weatherVERYCOLD    = "The temperature is -20 - -40 C"
	weatherEXTREMECOLD = "The temperature is -50 and lower"
)

type WeatherState struct {
	dice           *dice.Dicepool
	conditions     string
	daysSinceStorm int
}

func NewWeather() WeatherState {
	ws := WeatherState{}
	ws.dice = dice.New()
	return ws
}

func (ws *WeatherState) RollStorm() {
	r := ws.dice.RollNext("2d6").DM(ws.daysSinceStorm).Sum()
	precipation := "None"
	if r < 7 {
		fmt.Println("Precipation:", precipation)
		return
	}

}

func ResultIs(r int, expect string) bool {
	last := string(expect[len(expect)-1:])
	data := ""
	compareWith := []int{}
	switch last {
	case "+", "-":
		data = strings.TrimSuffix(expect, last)
		pts, err := strconv.Atoi(data)
		if err != nil {
			panic(err)
		}
		compareWith = append(compareWith, pts)
		if r >= pts {
			return true
		}
	default:

	}
	return true

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
