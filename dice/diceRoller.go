package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//Dicepool -
type Dicepool struct {
	dice       int
	edges      int
	modPerDie  int
	modTotal   int
	seed       int64
	result     []int
	trueRandom bool
}

// func main() {
// 	dp := Roll("15d16").RerollEach(1).ResultSum()
// 	fmt.Println(dp)

// 	/*
// 		Example:
// 			result := roll.Dice("2d6").RerollSum(12)
// 			roll.Dice("2d6").RerollEach(1)

// 			каждая функция действия должна принимать dp и отдавать его измененным

// 	*/
// }

//Roll - создает и возвращает структуру из которой можно брать результат,
//манипулировать.
func Roll(code string) *Dicepool {
	dp := Dicepool{}
	dp.seed = time.Now().UTC().UnixNano()
	dp.dice, dp.edges = decodeDiceCode(code)
	rand.Seed(dp.seed)
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, rand.Intn(dp.edges-1)+1)
	}
	return &dp
}

func decodeDiceCode(code string) (int, int) {
	code = strings.ToUpper(code)
	data := strings.Split(code, "D")
	dice, err := strconv.Atoi(data[0])
	if err != nil {
		return 0, 0
	}
	edges, err := strconv.Atoi(data[1])
	if err != nil {
		return 0, 0
	}
	return dice, edges
}

//////////////////////////////////////////////////////////
//Results:

//Result - возвращает слайс с результатами броска дайспула
func (dp *Dicepool) Result() []int {
	return dp.result
}

//ResultSum - возвращает сумму очков броска
func (dp *Dicepool) ResultSum() int {
	sum := 0
	for i := 0; i < len(dp.result); i++ {
		sum = sum + dp.result[i]
	}
	return sum
}

//ResultString - возвращает результата в виде стринга
func (dp *Dicepool) ResultString() string {
	res := ""
	for i := 0; i < len(dp.result); i++ {
		res = res + strconv.Itoa(dp.result[i])
	}
	return res
}

//ResultTN - возвращает true если сумма броска больше/равна tn
func (dp *Dicepool) ResultTN(tn int) bool {
	if dp.ResultSum() < tn {
		return false
	}
	return true
}

//////////////////////////////////////////////////////////
//Actions:

//SetSeed - фиксирует результат броска
func (dp *Dicepool) SetSeed(s int64) *Dicepool {
	dp.seed = s
	dp.result = nil
	rand.Seed(dp.seed)
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, rand.Intn(dp.edges-1)+1)
	}
	return dp
}

//RerollEach - перебрасывает все unwanted
//TODO: исключить вечную петлю
//TODO: нужен вариант с множеством unwanted
func (dp *Dicepool) RerollEach(unwanted int) *Dicepool {

	for i, val := range dp.result {
		for val == unwanted {
			val = rand.Intn(dp.edges-1) + 1
			dp.result[i] = val
		}
	}
	return dp
}

//ReplaceEach - заменяет все unwanted на wanted
func (dp *Dicepool) ReplaceEach(unwanted, wanted int) *Dicepool {
	for i := range dp.result {
		for dp.result[i] == unwanted {
			dp.result[i] = wanted
		}
	}
	return dp
}

//ReplaceOne - меняет значение конкретного дайса
func (dp *Dicepool) ReplaceOne(die, newVal int) *Dicepool {
	if len(dp.result) < die {
		return dp
	}
	dp.result[die] = newVal
	return dp
}

//////////////////////////////////////////////////////////
//Probe:

//Probe - оценивает вероятность достичь tn
func Probe(code string, tn int) float64 {
	dp := new(Dicepool)
	dice, edges := decodeDiceCode(code)
	dp.dice = dice
	dp.edges = edges
	var d int
	var positiveOutcome int
	fmt.Println(dp)
	for i := 0; i < dice; i++ {
		dp.result = append(dp.result, 1)
	}
	totalRange := dp.dice * dp.edges
	fmt.Println(dp, 1, totalRange)
	for i := 0; i < totalRange; i++ {
		dp.result[d]++
		if dp.result[d] > edges {
			dp.result[d]--
			d++
		}
		fmt.Println(dp, 2)
		if dp.ResultTN(tn) {
			positiveOutcome++
		}
		fmt.Println(dp, 3)
	}
	fmt.Println(positiveOutcome, totalRange)
	res := float64(positiveOutcome) / float64(totalRange)
	return res
}
