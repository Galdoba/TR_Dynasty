package wrld

import (
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

type task struct {
	key1, key2 string
}

type worldMaker struct {
}

type WorldMaker interface {
	Apply(string, string)
}

func CustomWorld(name string, tSlice ...task) World {
	w := World{}
	w.data = make(map[string]string)
	w.dicepool = *dice.New(utils.SeedFromString(name + name))
	return w
}

type gunMaker struct {
	dp               *dice.Dicepool //дайспул который используется для бросков
	designsConcluded []int          //список совершенных типов операций
}

func (w *World) Apply(key1, key2 string) task {
	return task{key1, key2}
}
