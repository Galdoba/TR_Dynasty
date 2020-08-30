package otu

import (
	"errors"
	"strconv"
	"strings"
)

type Info struct {
	Info string
}

var trData []string

func init() {
	trData = TrojanReachData()
}

type InfoRetriver interface {
	Sector() string
	SubSector() string
	HexData() string
	Name() string
	UWP() string
	Bases() []string
	Remarks() []string
	Zone() string
	PBG() string
	Allegiance() string
	Stars() []string
	Iextention() string
	Eextention() string
	Cextention() string
	Nobility() string
	Worlds() string
	RU() string
}

func (oi Info) Sector() string {
	data := strings.Split(oi.Info, "	")
	return data[0]
}

func (oi Info) SubSector() string {
	data := strings.Split(oi.Info, "	")
	return data[1]
}

func (oi Info) Hex() string {
	data := strings.Split(oi.Info, "	")
	return data[2]
}

func (oi Info) Name() string {
	data := strings.Split(oi.Info, "	")
	return data[3]
}
func (oi Info) UWP() string {
	data := strings.Split(oi.Info, "	")
	return data[4]
}
func (oi Info) Bases() []string {
	data := strings.Split(oi.Info, "	")
	bases := parseBasesT5(data[5])
	return bases
}

func parseBasesT5(data string) []string {
	bases := strings.Split(data, "")
	return bases
}

func (oi Info) Remarks() []string {
	data := strings.Split(oi.Info, "	")
	rem := strings.Split(data[6], " ")
	return rem
}
func (oi Info) Zone() string {
	data := strings.Split(oi.Info, "	")
	return data[7]
}
func (oi Info) PBG() string {
	data := strings.Split(oi.Info, "	")
	return data[8]
}
func (oi Info) Allegiance() string {
	data := strings.Split(oi.Info, "	")
	return data[9]
}
func (oi Info) Stars() string {
	data := strings.Split(oi.Info, "	")
	return data[10]
}
func (oi Info) Iextention() string {
	data := strings.Split(oi.Info, "	")
	return data[11]
}
func (oi Info) Eextention() string {
	data := strings.Split(oi.Info, "	")
	return data[12]
}
func (oi Info) Cextention() string {
	data := strings.Split(oi.Info, "	")
	return data[13]
}
func (oi Info) Nobility() string {
	data := strings.Split(oi.Info, "	")
	return data[14]
}
func (oi Info) Worlds() string {
	data := strings.Split(oi.Info, "	")
	return data[15]
}
func (oi Info) RU() string {
	data := strings.Split(oi.Info, "	")
	return data[16]
}

func MapDataByHex(data []string) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range data {
		data := strings.Split(v, "	")
		dataMap[data[2]] = v
	}
	return dataMap
}

func MapDataByName(data []string) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range data {
		data := strings.Split(v, "	")
		dataMap[data[3]] = v
	}
	return dataMap
}

func MapDataByUWP(data []string) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range data {
		data := strings.Split(v, "	")
		dataMap[data[4]] = v
	}
	return dataMap
}

func subSectorOffset(ss string) (int, int) {
	xOffset := 0
	yOffset := 0
	switch ss {
	case "A":
		xOffset = 0
		yOffset = 0
	case "B":
		xOffset = 8
		yOffset = 0
	case "C":
		xOffset = 16
		yOffset = 0
	case "D":
		xOffset = 24
		yOffset = 0
	case "E":
		xOffset = 0
		yOffset = 10
	case "F":
		xOffset = 8
		yOffset = 10
	case "G":
		xOffset = 16
		yOffset = 10
	case "H":
		xOffset = 24
		yOffset = 10
	case "I":
		xOffset = 0
		yOffset = 20
	case "J":
		xOffset = 8
		yOffset = 20
	case "K":
		xOffset = 16
		yOffset = 20
	case "L":
		xOffset = 24
		yOffset = 20
	case "M":
		xOffset = 0
		yOffset = 30
	case "N":
		xOffset = 8
		yOffset = 30
	case "O":
		xOffset = 16
		yOffset = 30
	case "P":
		xOffset = 24
		yOffset = 30
	}
	return xOffset, yOffset
}

func hex5ToHex4(hex5 string) string {
	if len(hex5) != 5 {
		return "Wrong Format"
	}
	hexParts := strings.Split(hex5, "")
	ss := hexParts[0]
	xOffset, yOffset := subSectorOffset(ss)
	x, _ := strconv.Atoi(hexParts[1] + hexParts[2])
	y, _ := strconv.Atoi(hexParts[3] + hexParts[4])
	x += xOffset
	y += yOffset
	res := ""
	if x < 10 {
		res += "0"
	}
	res += strconv.Itoa(x)
	if y < 10 {
		res += "0"
	}
	res += strconv.Itoa(y)
	return res
}

func GetDataOn(input string) (Info, error) {
	if val, ok := MapDataByHex(trData)[input]; ok {
		return Info{val}, nil
	}
	if val, ok := MapDataByHex(trData)[hex5ToHex4(input)]; ok {
		return Info{val}, nil
	}
	if val, ok := MapDataByName(trData)[input]; ok {
		return Info{val}, nil
	}
	if val, ok := MapDataByUWP(trData)[input]; ok {
		return Info{val}, nil
	}
	return Info{}, errors.New("No Data on '" + input + "'")
}
