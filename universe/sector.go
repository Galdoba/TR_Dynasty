package universe

import "github.com/Galdoba/utils"

type Sector struct {
	name    string
	HexList []*stSyst
}

type stSyst struct {
	name string
	star bool
}

func NewSector(name string) *Sector {
	sec := Sector{}
	sec.name = name
	sec.populate()
	return &sec
}

func (sec *Sector) populate() {
	for _, coords := range coordsList() {
		sec.HexList = append(sec.HexList, NewStarSyst(sec.name, coords))
	}
}

func NewStarSyst(secName, coords string) *stSyst {
	stSys := stSyst{}
	if utils.RandomBool() {
		stSys.star = true
	}
	stSys.name = secName + "  " + coords
	return &stSys
}

func coordsList() []string {
	return []string{
		"0000",
		"0001",
		"0002",
		"0003",
		"0100",
		"0101",
		"0102",
		"0103",
		"0204",
		"0200",
		"0201",
		"0202",
		"0303",
		"0304",
		"0300",
		"0301",
	}
}
