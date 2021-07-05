package sai

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

const (
	BULK   = "Bulk"
	LARGE  = "Large"
	MEDIUM = "Medium"
	SMALL  = "Small"
	MINOR  = "Minor"
)

type ShippingActivityData struct {
	dice            *dice.Dicepool
	portClass       string
	populationScore int
	inportanceScore int
	portDices       int
	portDM          int
	portMult        int
	shipPresent     map[string]int
	shipsTotal      int
}

func NewSAI(w wrld.World) ShippingActivityData {
	sa := ShippingActivityData{}
	sa.dice = dice.New().SetSeed(w.Name())
	sa.shipPresent = make(map[string]int)
	portUWP := uwp.New(w.UWP())
	sa.portClass = portUWP.Starport().String()
	sa.portMult = 1
	switch sa.portClass {
	case "A":
		sa.portMult = 3
	case "B":
		sa.portMult = 2
	case "D":
		sa.portDices = -2
		sa.portDM = -1
	case "E":
		sa.portDices = -4
		sa.portDM = -2
	}
	sa.portDices += w.ImportanceVal()
	sa.evaluateShipNumber()
	fmt.Println(sa)
	return sa
}

func (sa *ShippingActivityData) evaluateShipNumber() {
	if sa.portDices < 1 {
		return
	}
	ships := (sa.dice.RollNext(strconv.Itoa(sa.portDices)+"d6").Sum() + (sa.portDices * sa.portDM)) * sa.portMult
	for i := 0; i < ships; i++ {
		switch sa.portClass {
		case "A":
			sa.dice.RollNext("1d100").Sum()
			switch {
			case sa.dice.ResultIs("1 5"):
				sa.shipPresent[BULK]++
			case sa.dice.ResultIs("6 15"):
				sa.shipPresent[LARGE]++
			case sa.dice.ResultIs("16 35"):
				sa.shipPresent[MEDIUM]++
			case sa.dice.ResultIs("36 65"):
				sa.shipPresent[SMALL]++
			case sa.dice.ResultIs("66 100"):
				sa.shipPresent[MINOR]++
			default:
				panic("Case A")
			}
		case "B":
			sa.dice.RollNext("1d100").Sum()
			switch {
			case sa.dice.ResultIs("1 5"):
				sa.shipPresent[LARGE]++
			case sa.dice.ResultIs("6 15"):
				sa.shipPresent[MEDIUM]++
			case sa.dice.ResultIs("16 35"):
				sa.shipPresent[SMALL]++
			case sa.dice.ResultIs("36 100"):
				sa.shipPresent[MINOR]++
			default:
				panic("Case B")
			}
		case "C", "D":
			sa.dice.RollNext("1d100").Sum()
			switch {
			case sa.dice.ResultIs("1 5"):
				sa.shipPresent[MEDIUM]++
			case sa.dice.ResultIs("6 15"):
				sa.shipPresent[SMALL]++
			case sa.dice.ResultIs("16 100"):
				sa.shipPresent[MINOR]++
			default:
				panic("Case C and D")
			}
		case "E":
			sa.dice.RollNext("1d100").Sum()
			switch {
			case sa.dice.ResultIs("1 5"):
				sa.shipPresent[SMALL]++
			case sa.dice.ResultIs("6 100"):
				sa.shipPresent[MINOR]++
			default:
				panic("Case E")
			}
		}
	}

}

func TestSAI() {
	w := wrld.PickWorld()
	fmt.Println(w.SecondSurvey())
	wUWP := uwp.New(w.UWP())
	arriveFactor := 12
	switch wUWP.Starport().String() {
	case "E":
		arriveFactor -= 10
	case "D":
		arriveFactor -= 9
	case "C":
		arriveFactor += 8
	case "B":
		arriveFactor += 7
	case "A":
		arriveFactor += 6
	}
	fmt.Println("arriveFactor =", arriveFactor)
	die := dice.New().SetSeed(w.Name())

	r := die.RollNext("2d6").DM(arriveFactor * -1).Sum()
	fmt.Println(r)

	accum := 0
	for i := 0; i < 100; i++ {
		accum += r
		if accum > 100 {
			accum -= 1000
			fmt.Println("day", i, "Ship arrived")
		}
	}

	return
	fmt.Println("+---Shipping Activity-------+--------+-------+-------+")
	fmt.Println("| Port Class | Bulk | Large | Medium | Small | Minor |")
	fmt.Println("| A          | 2    | 1     | 10     | 20    | 65    |")
	fmt.Println("+------------+------+-------+--------+-------+-------+")
	sa := NewSAI(wrld.PickWorld())
	fmt.Println(sa)
	fmt.Print("Roll: ", sa.portDices, " dices with dm = ", sa.portDM, " and multyply x", sa.portMult)
	fmt.Print("\n")
}
