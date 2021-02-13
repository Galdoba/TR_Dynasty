package task

import (
	"fmt"
	"time"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	//DifficultyVeryEasy -
	DifficultyVeryEasy = 2
	//DifficultyEasy -
	DifficultyEasy = 4
	//DifficultyRoutine -
	DifficultyRoutine = 6
	//DifficultyAverage -
	DifficultyAverage = 8
	//DifficultyDifficult -
	DifficultyDifficult = 10
	//DifficultyVeryDifficult -
	DifficultyVeryDifficult = 12
	//DifficultyFormidable -
	DifficultyFormidable = 14
	//DifficultyImpossible -
	DifficultyImpossible = 16
	//BOON - Плюшка
	BOON = "BOON"
)

type mgt2Task struct {
	phrase      string     //общее описание задания
	modifiers   []modifier //модификаторы для теста
	difficulty  int        //сложность (если не назначена то 8)
	boonAndBane int
}

type modifier struct {
	val     int
	comment string
}

type statement struct {
	value   int
	comment string
}

//NewTask - создает новое задание
func NewTask(phrase string, st ...statement) mgt2Task {
	t := mgt2Task{}
	t.phrase = phrase
	t.difficulty = 8
	for _, sttmnt := range st {
		t.handleStatement(sttmnt)
	}
	fmt.Println(t)
	return t
}

func (t *mgt2Task) handleStatement(sttmnt statement) {
	switch sttmnt.comment {
	default:
		t.modifiers = append(t.modifiers, modifier{sttmnt.value, sttmnt.comment})
	case "DIFFICULTY":
		t.difficulty = -1 * sttmnt.value
	case BOON:
		t.boonAndBane += sttmnt.value
	}
}

//Resolve -
func (t *mgt2Task) Resolve(dicepool ...*dice.Dicepool) (int, int) {
	dp := dice.Dicepool{}
	if len(dicepool) > 0 {
		dp = *dicepool[0]
	} else {
		dp = *dice.New().SetSeed(time.Now().UnixNano())
	}
	mod := 0
	for _, val := range t.modifiers {
		mod += val.val
	}
	mod -= t.difficulty
	dp = *dp.RollNext("2d6").DM(mod)
	if t.boonAndBane > 0 {
		dp = *dp.Boon()
	}
	if t.boonAndBane < 0 {
		dp = *dp.Bane()
	}
	eff := dp.Sum()
	tf := dp.RollNext("1d6").Sum()
	return eff, tf
}

func (t *mgt2Task) Add(st statement) {
	t.handleStatement(st)
}

// type resolver struct {
// 	dp dice.Dicepool
// }

// //Resolver - решает стандартные тесты
// type Resolver interface {
// 	Resolve(mgt2Task) (int, int)
// }

// //NewResolver - берет существующий или создает новый дайспул для решения заданий
// func NewResolver(dp ...dice.Dicepool) resolver {
// 	d := dice.New(time.Now().UnixNano())
// 	if len(dp) < 1 {
// 		return resolver{*d}
// 	}
// 	return resolver{dp[0]}
// }

// func (r *resolver) Resolve(task mgt2Task) (int, int) {
// 	mod := 0
// 	for _, val := range task.modifiers {
// 		mod += val.val
// 	}
// 	mod -= task.difficulty
// 	eff := r.dp.RollNext("2d6").DM(mod).Sum()
// 	tf := r.dp.RollNext("1d6").Sum()
// 	return eff, tf
// }

func Modifier(v int, comment ...string) statement {
	comm := "[NO COMMENT]"
	if len(comment) > 0 {
		comm = comment[0]
	}
	return statement{v, comm}
}

func Difficulty(v int) statement {
	return statement{-v, "DIFFICULTY"}
}

func Boon(v ...int) statement {
	if len(v) > 0 {
		return statement{v[0], BOON}
	}
	return statement{1, BOON}
}

func Bane(v ...int) statement {
	if len(v) > 0 {
		return statement{-1 * v[0], BOON}
	}
	return statement{-1, BOON}
}
