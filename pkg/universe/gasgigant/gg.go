package gasgigant

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

type gasgig struct {
	gigantType   string    //SGG / LGG / IG
	gigantSize   ehex.Ehex //
	gigantDiam   int       // miles
	gigantDiamMM float64   // Mm
	gigantG      float64   // тяготение на глубине забора топлива
}

func (gg *gasgig) String() string {
	r := fmt.Sprintf("Type: %v\n", gg.gigantType)
	r += fmt.Sprintf("Size: %v (%v)\n", gg.gigantSize.Code(), gg.gigantSize.Value())
	r += fmt.Sprintf("Diam: %v miles\n", gg.gigantDiam)
	r += fmt.Sprintf("Diam: %v MegaMeters\n", gg.gigantDiamMM)
	r += fmt.Sprintf("G: %vg\n", gg.gigantG)
	return r
}

func New(seed ...string) *gasgig {
	gg := gasgig{}
	dp := dice.New()
	if len(seed) > 0 {
		dp.SetSeed(seed[0])
	}
	gg.gigantSize = ehex.New().Set(dp.RollNext("2d6").DM(19).Sum())
	switch gg.gigantSize.Value() {
	case 21:
		gg.gigantDiam = randomizeInt(30000, dp)
	case 22:
		gg.gigantDiam = randomizeInt(40000, dp)
	case 23:
		gg.gigantDiam = randomizeInt(50000, dp)
	case 24:
		gg.gigantDiam = randomizeInt(60000, dp)
	case 25:
		gg.gigantDiam = randomizeInt(70000, dp)
	case 26:
		gg.gigantDiam = randomizeInt(80000, dp)
	case 27:
		gg.gigantDiam = randomizeInt(90000, dp)
	case 28:
		gg.gigantDiam = randomizeInt(125000, dp)
	case 29:
		gg.gigantDiam = randomizeInt(180000, dp)
	}
	gg.gigantType = "LGG"
	if gg.gigantDiam < 45000 {
		gg.gigantType = "SGG"
		if dp.RollNext("1d2").Sum() == 2 {
			gg.gigantType = "IG"
		}
	}
	gg.gigantG = utils.RoundFloat64(float64(gg.gigantDiam)*0.00001, 2)
	gg.gigantDiamMM = utils.RoundFloat64(float64(gg.gigantDiam)*0.016, 3)
	return &gg
}

func (gg *gasgig) GetType() string {
	return gg.gigantType
}

func (gg *gasgig) GetSizeCode() string {
	return gg.gigantSize.Code()
}

func (gg *gasgig) GetSizeInt() int {
	return gg.gigantSize.Value()
}

func (gg *gasgig) Diam() float64 {
	return gg.gigantDiamMM
}

func (gg *gasgig) DiamMiles() int {
	return gg.gigantDiam
}

func (gg *gasgig) Gravity() string {
	return fmt.Sprintf("%vg", gg.gigantG)
}

func randomizeInt(i int, dp ...*dice.Dicepool) int {
	die := dice.New()
	if len(dp) != 0 {
		die = dp[0]
	}
	resultSl := []int{}
	for i > 0 {
		lowI := i / 10
		rest := i % 10
		resultSl = append(resultSl, rest)
		i = lowI
	}
	skip := false
	for i, val := range resultSl {
		if val == 0 && !skip {
			resultSl[i] = die.FluxNext()
		} else {
			skip = true
		}
	}
	r := 0
	for i, val := range resultSl {
		n := 1
		for s := 1; s < i+1; s++ {
			n = n * 10
		}
		a := n * val
		r += a
	}
	return r
}
