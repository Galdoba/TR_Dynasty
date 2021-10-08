package survey

import (
	"strconv"
	"strings"
)

type SecondSurveyData struct {
	CoordX        int
	CoordY        int
	Sector        string
	Hex           string
	MW_Name       string
	MW_UWP        string
	MW_Remarks    string
	MW_Importance string
	MW_Economic   string
	MW_Cultural   string
	MW_Nobility   string
	Bases         string
	TravelZone    string
	PBG           string
	Worlds        int
	Allegiance    string
	Stellar       string
	RU            int
	input         string //temp
	errors        []error
}

func Parse(input string) *SecondSurveyData {
	ssd := SecondSurveyData{}
	//ssd.input = input
	data := strings.Split(input, "|")
	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}
	xCoord, errXcoord := strconv.Atoi(data[19])
	ssd.errors = append(ssd.errors, errXcoord)
	yCoord, errYcoord := strconv.Atoi(data[20])
	ssd.errors = append(ssd.errors, errYcoord)
	ssd.CoordX = xCoord
	ssd.CoordY = yCoord
	ssd.Sector = data[23]
	ssd.Hex = data[2]
	ssd.MW_Name = data[1]
	ssd.MW_UWP = data[3]
	ssd.MW_Remarks = data[21]
	ssd.MW_Importance = data[10]
	ssd.MW_Economic = data[12]
	ssd.MW_Cultural = data[13]
	ssd.MW_Nobility = data[14]
	ssd.Bases = data[6]
	ssd.TravelZone = data[5]
	ssd.PBG = data[4]
	worlds, errWorlds := strconv.Atoi(data[5])
	ssd.errors = append(ssd.errors, errWorlds)
	ssd.Worlds = worlds
	ssd.Allegiance = data[5]
	ssd.Stellar = data[5]
	ru, errRu := strconv.Atoi(data[5])
	ssd.errors = append(ssd.errors, errRu)
	ssd.RU = ru

	return &ssd
}
