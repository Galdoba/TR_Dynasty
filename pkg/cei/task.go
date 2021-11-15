package cei

type Event struct {
	descr       string
	eventType   string
	eventTests  []int
	eventMorale []int
}

func Task(descr string, diff int) {}
