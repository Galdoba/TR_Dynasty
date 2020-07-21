package dice

type Roller interface {
	Roll()
}

type Dicepool struct {
	dice      int
	edges     int
	modPerDie int
	modTotal  int
	seed      int64
	counter   int
	result    []int
}

func NewDicepool(qty, edges int) *Dicepool {
	dp := Dicepool{}
	dp.dice = qty
	dp.edges = edges
	return &dp
}

func (dp *Dicepool) SetSeed(newSeed int64) {
	dp.seed = newSeed
}
