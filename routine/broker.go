package routine

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/name"
)

var localBroker broker

type broker struct {
	skill int
	cut   float64
}

type Broker interface {
	CutFrom(int) int
	DM() int
}

func NewBroker(skill int) broker {
	br := broker{}
	br.skill = skill
	switch skill {
	default:
		br.cut = 0
		br.skill = 0
	case 0:
		br.cut = 0.5
	case 1:
		br.cut = 1.0
	case 2:
		br.cut = 2.0
	case 3:
		br.cut = 5.0
	case 4:
		br.cut = 7.0
	case 5:
		br.cut = 10.0
	case 6:
		br.cut = 15.0
	}
	return br
}

func (br broker) CutFrom(sum int) int {
	return int((float64(sum) / 100) * br.cut)
}

func (br broker) DM() int {
	return br.skill
}

func chooseBroker() {
	printSlow("Please, choose your Broker:\n")
	printSlow(" [0] - DM: +0	Percentage: 0.5 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [1] - DM: +1	Percentage:   1 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [2] - DM: +2	Percentage:   2 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [3] - DM: +3	Percentage:   5 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [4] - DM: +4	Percentage:   7 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [5] - DM: +5	Percentage:  10 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [6] - DM: +6	Percentage:  15 %	Name: " + name.RandomNew() + "\n")
	printSlow(" [7] - Refuse local broker's services\n")
	valid := false
	for !valid {
		skill := userInputInt("Select: ")
		switch skill {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			localBroker = NewBroker(skill)
			valid = true
			if skill == 7 {
				printSlow("Services refused\n")
			} else {
				printSlow("Local broker hired\n")
			}
		default:
			printSlow("WARNING: can't parse '" + strconv.Itoa(skill) + "'")
		}
	}
}
