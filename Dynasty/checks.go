package dynasty

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/task"
)

func (d *Dynasty) rollCheckAptitude(chars string, apt string, difficulty int) int {
	//r := d.dicepool.RollNext("2d6").Sum()
	//eff := r + d.aptitudeValue(apt) + d.characteristicDM(chars) - difficulty

	t := task.NewTask("Test Aptitude Check",
		task.Modifier(DM(d.aptitudeValue(chars)), "from "+chars),
		task.Modifier(d.characteristicDM(apt), "from "+apt),
		task.Modifier(-5, "test Control"),
		task.Difficulty(difficulty),
	)
	t.Add(task.Modifier(-5, "test Add2"))
	eff, _ := t.Resolve()

	return eff
}

func (d *Dynasty) probeCheckAptitude(chars string, apt string, difficulty, etn int) float64 {

	eff := d.aptitudeValue(apt) + d.characteristicDM(chars) - difficulty
	fmt.Println(d.aptitudeValue(apt))
	fmt.Println(d.characteristicDM(chars))
	fmt.Println(-difficulty)
	rArray := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	rpArray := []float64{100.000, 97.222, 91.667, 83.333, 72.222, 58.333, 41.667, 27.778, 16.667, 8.333, 2.778}
	for i, rv := range rArray {
		fmt.Println("test rv =", rv, "|", eff+rv, "|", etn, "|", rpArray[i])
		if eff+rv >= etn {
			return rpArray[i]
		}
	}
	return 0.0
}
