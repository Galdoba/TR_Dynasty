package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func main() {
	currentYear := 1106
	dp := dice.New().SetSeed(4253622)
	dm := 0
	damageMap := make(map[int]int)
	for currentYear < 1201 {
		for m := 0; m < 12; m++ {
			fmt.Printf("Date: %v-%v\n", m+1, currentYear)
			r := dp.RollNext("2d6").DM(dm).Sum()
			dm++
			if r <= 8 {
				dm = 0
				r2 := dp.RollNext("2d6").Sum()
				if damageMap[r2] < 6 {

				}
				fmt.Printf("Ship degraded\n")
			}
		}
		currentYear++
	}
	for i := 2; i < 13; i++ {
		fmt.Printf("damageMap[%v] = %v\n", i, damageMap[i])
	}
}
