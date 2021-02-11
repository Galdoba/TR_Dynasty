package wrld

import (
	"errors"
	"fmt"
	"time"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type task struct {
	key1, key2 string
}

type worldMaker struct {
	dp       *dice.Dicepool
	eventLog []string
}

type WorldMaker interface {
	Apply(string, string) task
	AddEntry(string)
}

func NewWorldMaker(seed int64) worldMaker {
	wm := worldMaker{}
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	wm.dp = dice.New().SetSeed(seed)
	return wm
}

func (wm *worldMaker) MakeCustomWorld(name string, taskSlice ...task) World {
	w := World{}
	w.data = make(map[string]string)
	w.dicepool = *dice.New().SetSeed(name + name)
	w.data[worldNAME] = name
	for i := range taskSlice {
		fmt.Println(1)
		if err := w.absorb(taskSlice[i]); err != nil {
			wm.AddEntry(err.Error())
			fmt.Println(2)
			continue
		}
		wm.AddEntry(taskSlice[i].key1 + " absorbsion successfull")
		fmt.Println(3)
	}
	return w
}

func (w *World) absorb(t task) error {
	if _, ok := w.data[t.key1]; ok {
		return errors.New("Data on " + t.key1 + " already exists")
	}
	w.data[t.key1] = t.key2
	return nil
}

func (wm *worldMaker) Apply(key1, key2 string) task {
	return task{key1, key2}
}

func (wm *worldMaker) AddEntry(entry string) {
	wm.eventLog = append(wm.eventLog, entry)
}

func (wm *worldMaker) ShowLog() {
	for i := range wm.eventLog {
		fmt.Print(wm.eventLog[i] + "\n")
	}
}
