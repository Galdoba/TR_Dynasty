package dice

import (
	"math"
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
	boon       bool
	bane       bool
	src        rand.Source
	rand       rand.Rand
}

// func main() {
// 	dp := Roll("15d16").RerollEach(1).ResultSum()
// 	fmt.Println(dp)

// 	/*
// 		Example:
// 			result := roll.Dice("2d6").RerollSum(12)
// 			roll.Dice("2d6").RerollEach(1)

// 			каждая функция действия должна принимать dp и отдавать его измененным

//   dice.New(SeedFromStr("Planet Name")).RollSdd("2d6")

// 	*/
// }

func New(seed int64) *Dicepool {
	dp := Dicepool{}
	dp.seed = seed
	if dp.seed == 0 {
		dp.seed = time.Now().UTC().UnixNano()
	}
	dp.src = rand.NewSource(dp.seed)
	dp.rand = *rand.New(dp.src)
	return &dp
}

//Roll - создает и возвращает структуру из которой можно брать результат,
//манипулировать.
func Roll(code string) *Dicepool {
	dp := Dicepool{}
	time.Sleep(time.Millisecond)
	if dp.seed == 0 {
		dp.seed = time.Now().UTC().UnixNano()
		dp.src = rand.NewSource(dp.seed)
		dp.rand = *rand.New(dp.src)
	}
	dp.dice, dp.edges = decodeDiceCode(code)
	//rand.Seed(dp.seed)

	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
	}
	return &dp
}

func (dp *Dicepool) RollNext(code string) *Dicepool {
	dp.result = nil
	dp.dice, dp.edges = decodeDiceCode(code)
	dp.modPerDie = 0
	dp.modTotal = 0
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
	}
	return dp
}

func decodeDiceCode(code string) (int, int) {
	code = strings.ToUpper(code)
	data := strings.Split(code, "D")
	var dice int
	dice, _ = strconv.Atoi(data[0])
	if data[0] == "" {
		dice = 1
	}
	edges, err := strconv.Atoi(data[1])
	if err != nil {
		return 0, 0
	}
	return dice, edges
}

func encodeDiceCode(dice, edges int) string {
	return strconv.Itoa(dice) + "D" + strconv.Itoa(edges)
}

//////////////////////////////////////////////////////////
//Results:

//Result - возвращает слайс с результатами броска дайспула
func (dp *Dicepool) Result() []int {
	return dp.result
}

//Sum - возвращает сумму очков броска
func (dp *Dicepool) Sum() int {
	sum := 0
	for i := 0; i < len(dp.result); i++ {
		sum = sum + (dp.result[i] + dp.modPerDie)
	}
	sum = sum + dp.modTotal
	return sum
}

//SumStr - возвращает сумму очков броска в виде стринга
func (dp *Dicepool) SumStr() string {
	return strconv.Itoa(dp.Sum())
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
	if dp.Sum() < tn {
		return false
	}
	return true
}

//////////////////////////////////////////////////////////
//Actions:

//Boon - фиксирует результат броска
func (dp *Dicepool) Boon() *Dicepool {
	lowest := 0
	targetVal := dp.edges
	for i, val := range dp.result {
		if val < targetVal {
			targetVal = val
			lowest = i
		}

	}
	d1 := rand.Intn(dp.edges) + 1
	if d1 > targetVal {
		dp.result[lowest] = d1
	}
	return dp
}

//Bane - фиксирует результат броска
func (dp *Dicepool) Bane() *Dicepool {
	highest := 0
	targetVal := 0
	for i, val := range dp.result {
		if val > targetVal {
			targetVal = val
			highest = i
		}

	}
	d1 := rand.Intn(dp.edges) + 1
	if d1 < targetVal {
		dp.result[highest] = d1
	}
	return dp
}

//DM - фиксирует результат броска
func (dp *Dicepool) DM(s int) *Dicepool {
	dp.modTotal = s
	return dp
}

//ModPerDie - фиксирует результат броска
func (dp *Dicepool) ModPerDie(s int) *Dicepool {
	dp.modPerDie = s
	return dp
}

//SetSeed - фиксирует результат броска
func (dp *Dicepool) SetSeed(s int64) *Dicepool {
	dp.seed = s
	dp.result = nil
	dp.src = rand.NewSource(dp.seed)
	dp.rand = *rand.New(dp.src)
	//rand.Seed(dp.seed)
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
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

//ReRoll - меняет значение броска
// func (dp *Dicepool) ReRoll() *Dicepool {
// 	code := encodeDiceCode(dp.dice, dp.edges)
// 	dpNew := Roll(code)
// 	return dpNew
// }

//////////////////////////////////////////////////////////
//Probe:

//Probe - перечисляет возможные варианты
func Probe(code string, tn int) map[int][]int {
	dice, edges := decodeDiceCode(code)
	resMap := rollCombinations(edges, dice)
	return resMap
}

//ProbeTN - оценивает вероятность достичь tn
func ProbeTN(code string, tn int) float64 {
	dp := new(Dicepool)
	dice, edges := decodeDiceCode(code)
	dp.dice = dice
	dp.edges = edges
	var positiveOutcome int
	for i := 0; i < dice; i++ {
		dp.result = append(dp.result, 1)
	}
	totalRange := int(math.Pow(float64(dp.edges), float64(dp.dice)))
	resMap := rollCombinations(dp.edges, dp.dice)
	for _, v := range resMap {
		tdp := Dicepool{}
		tdp.result = v
		if tdp.ResultTN(tn) {
			positiveOutcome++
		}
	}
	res := float64(positiveOutcome) / float64(totalRange)
	return res
}

func rollCombinations(max, len int) map[int][]int {
	resMap := make(map[int][]int)
	var sl []int
	for i := 0; i < len; i++ {
		sl = append(sl, 1)
	}
	totalRange := int(math.Pow(float64(max), float64(len)))
	resMap[0] = sl
	for i := 0; i < totalRange; i++ {
		activeDie := 0
		for activeDie < len {
			if sl[activeDie] < max {
				sl[activeDie]++
				for j := range sl {
					resMap[i+1] = append(resMap[i+1], sl[j])
				}
				break
			} else {
				sl[activeDie] = 1
				activeDie++
			}
			if activeDie > len {
				break
			}
		}
	}
	return resMap
}

//////////////////////////////////////////////////////////
//QuickRolls:

//RollD66 - Возвращает результат 2d6 в виде string
func RollD66() string {
	return Roll("2d6").ResultString()

}

//Roll1D -
func Roll1D(dm ...int) int {
	mod := 0
	if len(dm) > 0 {
		mod = dm[0]
	}
	return Roll("1d6").DM(mod).Sum()
}

//Roll2D -
func Roll2D(dm ...int) int {
	mod := 0
	if len(dm) > 0 {
		mod = dm[0]
	}
	return Roll("2d6").DM(mod).Sum()
}

//Roll3D -
func Roll3D(dm ...int) int {
	mod := 0
	if len(dm) > 0 {
		mod = dm[0]
	}
	return Roll("3d6").DM(mod).Sum()
}

//Roll4D -
func Roll4D(dm ...int) int {
	mod := 0
	if len(dm) > 0 {
		mod = dm[0]
	}
	return Roll("4d6").DM(mod).Sum()
}

//Roll5D -
func Roll5D(dm ...int) int {
	mod := 0
	if len(dm) > 0 {
		mod = dm[0]
	}
	return Roll("5d6").DM(mod).Sum()
}

func Flux() int {
	d1 := Roll1D()
	d2 := Roll1D()
	return d1 - d2
}

func FluxGOOD() int {
	d1 := Roll1D()
	d2 := Roll1D()
	if d1 >= d2 {
		return d1 - d2
	}
	return d2 - d1
}

func FluxBAD() int {
	d1 := Roll1D()
	d2 := Roll1D()
	if d1 <= d2 {
		return d1 - d2
	}
	return d2 - d1
}
