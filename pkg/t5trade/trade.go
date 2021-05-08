package t5trade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

func Test() {
	cargo := NewCargo(12, []string{"Ri", "Pa", "An", "Cp", "Ph"})
	fmt.Println("Sourceworld Cargo ID:", cargo.CargoID)
	fmt.Println("Marketworld Cargo price:", SellPrice(cargo, 10, []string{"Ri", "Ph", "Pa", "Pz"}))
}

/*

Lot.
Freight.
Cargo.
Speculator.
Market World.
Cost.
Price.
Selling Price.
Delivery. - two common forms of shipment are
OTO. Orbit To Orbit.
STS. Surface To Surface.
*/

type Cargo struct {
	CargoID string //A cargo is identified by stating its source world’s Tech Level, Trade Classifications, and Cost. Tech Level is derived directly from the source world UWP. All trade classifications possible are determined and then listed together
	Cost    int    //Purchase Cost at Sourceworld
}

func NewCargo(swTechLvl int, swTradeCodes []string) Cargo {
	c := Cargo{}
	c.CargoID = determineID(swTechLvl, swTradeCodes)
	c.Cost = determineCost(c.CargoID)
	c.CargoID = c.CargoID + "Cr" + strconv.Itoa(c.Cost)
	return c
}

func SellPrice(c Cargo, mwTL int, mwTC []string) int {
	tc := strings.Split(c.CargoID, " ")
	sell := 5000
	for _, cargoCode := range tc {
		switch cargoCode {
		case "Ag":
			if slicesOverlap([]string{"Ag", "As", "De", "Hi", "In", "Ri", "Va"}, mwTC) {
				sell += 1000
			}
		case "As":
			if slicesOverlap([]string{"As", "In", "Ri", "Va"}, mwTC) {
				sell += 1000
			}
		case "Ba":
			if slicesOverlap([]string{"In"}, mwTC) {
				sell += 1000
			}
		case "Fl":
			if slicesOverlap([]string{"Fl", "In"}, mwTC) {
				sell += 1000
			}
		case "Hi":
			if slicesOverlap([]string{"Hi"}, mwTC) {
				sell += 1000
			}
		case "In":
			if slicesOverlap([]string{"Ag", "As", "De", "Fl", "Hi", "In", "Ri", "Va"}, mwTC) {
				sell += 1000
			}
		case "Na":
			if slicesOverlap([]string{"As", "De", "Va"}, mwTC) {
				sell += 1000
			}
		case "Po":
			if slicesOverlap([]string{"Ag", "Hi", "In", "Ri"}, mwTC) {
				sell += 1000
			}
		case "Ri":
			if slicesOverlap([]string{"Ag", "De", "Hi", "In", "Ri"}, mwTC) {
				sell += 1000
			}
		case "Va":
			if slicesOverlap([]string{"As", "In", "Va"}, mwTC) {
				sell += 1000
			}
		}
	}
	tlEffect := sell * ((ehex.New(tc[0]).Value() - ehex.New(mwTL).Value()) * 10) / 100
	sell += tlEffect
	return sell
}

func slicesOverlap(sl1, sl2 []string) bool {
	for _, v1 := range sl1 {
		for _, v2 := range sl2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func determineID(tl int, tc []string) string {
	id := ehex.New(tl).String() + " - "
	for _, val := range tc {
		for _, tcx := range applicableTradeClassifications() {
			if val == tcx {
				id += val + " "
			}
		}
	}
	return id
}

func applicableTradeClassifications() []string {
	return []string{"Ag", "As", "Ba", "De", "Fl", "Hi", "In", "Lo", "Ni", "Po", "Ri", "Va"}
}

func determineCost(cargoID string) int {
	cost := 3000
	tc := strings.Split(cargoID, " ")
	for _, code := range tc {
		switch code {
		case "Ag", "As", "Hi", "In", "Po":
			cost = cost - 1000
		case "Ba", "De", "Fl", "Lo", "Ni", "Ri", "Va":
			cost = cost + 1000
		}
	}
	cost = cost + (ehex.New(tc[0]).Value() * 100)
	return cost
}

func randomTradeGoods(cargoID string) string {
	descr := ""
	data := strings.Split(cargoID, " ")
	chart := []int{}
	dp := dice.New()
	if utils.ListContains(data, "Ag") {
		ag := dp.RollNext("1d2").Sum()
		chart = append(chart, ag)
	}
	if utils.ListContains(data, "Ga") {
		chart = append(chart, 1)
	}
	if utils.ListContains(data, "Fa") {
		chart = append(chart, 2)
	}
	if utils.ListContains(data, "As") {
		chart = append(chart, 3)
	}
	if utils.ListContains(data, "De") {
		chart = append(chart, 4)
	}
	if utils.ListContains(data, "Fl") {
		chart = append(chart, 5)
	}
	if utils.ListContains(data, "Ic") {
		chart = append(chart, 6)
	}
	if utils.ListContains(data, "Na") {
		chart = append(chart, 7)
	}
	if utils.ListContains(data, "In") {
		chart = append(chart, 8)
	}
	if utils.ListContains(data, "Po") {
		chart = append(chart, 9)
	}
	if utils.ListContains(data, "Ri") {
		chart = append(chart, 10)
	}
	if utils.ListContains(data, "Va") {
		chart = append(chart, 11)
	}
	if utils.ListContains(data, "Cp") || utils.ListContains(data, "Cs") || utils.ListContains(data, "Cx") {
		chart = append(chart, 12)
	}
	if len(chart) == 0 {
		chart = append(chart, 12)
	}

	return descr
}
