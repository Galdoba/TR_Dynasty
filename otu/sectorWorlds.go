package otu

import "fmt"

var sectorWorlds [][]string

func init() {
	sectorWorlds = append(sectorWorlds, []string{"Adhara", "A0609", "", "B57A687-8", "Wa Ni", "", "", "G"})
	sectorWorlds = append(sectorWorlds, []string{"Allemange", "A0503", "", "X688000-0", "Ba", "R", "Strend Cluster", ""})
	sectorWorlds = append(sectorWorlds, []string{"Armada", "A0608", "", "A540244-A", "De Lo Ni Po", "", "", ""})
	sectorWorlds = append(sectorWorlds, []string{"Bilke", "A0110", "", "D987341-7", "Ni Lo Ga", "", "Florian League", "G"})
}

type otuWorld interface {
	Name() string
	Hex() string
	Bases() string
	UWP() string
	TradeCodes() string
	TravelZone() string
	Allegiance() string
	GasGigants() string
}

func GlobalhexToLocalhex(ghex string) string {
	if len(ghex) != 5 {
		return "ERROR"
	}
	sector := string([]byte(ghex)[0])
	if !checkSubSectorCode(sector) {
		return "ERROR"
	}
	return "good"
}

func checkSubSectorCode(ssCode string) bool {
	switch ssCode {
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P":
		return true
	}
	return false
}

func checkCoords(coords string) bool {
	if len(coords) != 2 {
		return false
	}
	for i, v := range coords {
		fmt.Println(i, v)
	}
	return true
}
