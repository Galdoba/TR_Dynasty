package dataparcer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

func readData() []string {
	lines := utils.LinesFromTXT("c:\\Users\\pemaltynov\\go\\src\\github.com\\Galdoba\\TR_Dynasty\\otu\\otuPuller\\data.txt")
	return lines
}

type parcedData struct {
	Name                 string //W
	Hex                  string //S
	UWP                  string //W
	PBG                  string //S
	Zone                 string //S
	Bases                string //S
	Allegiance           string //S
	Stellar              string //S
	SS                   string //S SubSector Code
	Ix                   string //W
	CalculatedImportance int    //W
	Ex                   string //W
	Cx                   string //W
	Nobility             string //W
	Worlds               int    //S
	ResourceUnits        int    //W
	Subsector            int    //S
	Quadrant             int    //S
	WorldX               int    //S
	WorldY               int    //S
	Remarks              string //W
	LegacyBaseCode       string //S
	Sector               string //S
	SubsectorName        string //S
	SectorAbbreviation   string //S
	AllegianceName       string //S
	OrbitCode            string //W
	Star                 string //W
	err                  error
}

func ParceWorldData(worlddata string) *parcedData {
	pd := parcedData{}
	if worlddata == "{''Worlds'':[]}" {
		pd.err = fmt.Errorf("worlddata is blank")
		return &pd
	}
	pd.Name = "?"
	pd.Hex = "?"
	pd.UWP = "?"
	pd.PBG = "?"
	pd.Zone = "?"
	pd.Bases = "?"
	pd.Allegiance = "?"
	pd.Stellar = "?"
	pd.SS = "?"
	pd.Ix = "?"
	pd.CalculatedImportance = 0
	pd.Ex = "?"
	pd.Cx = "?"
	pd.Nobility = "?"
	pd.Worlds = 0
	pd.ResourceUnits = 0
	pd.Subsector = 0
	pd.Quadrant = 0
	pd.WorldX = 0
	pd.WorldY = 0
	pd.Remarks = "?"
	pd.LegacyBaseCode = "?"
	pd.Sector = "?"
	pd.SubsectorName = "?"
	pd.SectorAbbreviation = "?"
	pd.AllegianceName = "?"
	//////////
	worlddata = strings.TrimPrefix(worlddata, "{''Worlds'':[{''")
	worlddata = strings.TrimSuffix(worlddata, "}]}")
	data := strings.Split(worlddata, ",''")
	for d, field := range data {
		fData := strings.Split(field, "'':''")
		if len(fData) == 1 {
			pd.assignIntData(fData[0])
			if pd.err != nil {
				return &pd
			}
			continue
		}
		fData[1] = strings.TrimSuffix(fData[1], "''")
		switch fData[0] {
		default:
			pd.err = fmt.Errorf("Unknown field met '%v'\n%v %v", fData[0], worlddata, d)
			return &pd
		case "Name":
			pd.Name = fData[1]
		case "Hex":
			pd.Hex = fData[1]
		case "UWP":
			pd.UWP = fData[1]
		case "PBG":
			pd.PBG = fData[1]
		case "Zone":
			pd.Zone = fData[1]
		case "Bases":
			pd.Bases = fData[1]
		case "Allegiance":
			pd.Allegiance = fData[1]
		case "Stellar":
			pd.Stellar = fData[1]
		case "SS":
			pd.SS = fData[1]
		case "Ix":
			pd.Ix = fData[1]
		case "Ex":
			pd.Ex = fData[1]
		case "Cx":
			pd.Cx = fData[1]
		case "Nobility":
			pd.Nobility = fData[1]
		case "Remarks":
			pd.Remarks = fData[1]
		case "LegacyBaseCode":
			pd.LegacyBaseCode = fData[1]
		case "Sector":
			pd.Sector = fData[1]
		case "SubsectorName":
			pd.SubsectorName = fData[1]
		case "SectorAbbreviation":
			pd.SectorAbbreviation = fData[1]
		case "AllegianceName":
			pd.AllegianceName = fData[1]
			//case "CalculatedImportance", "Worlds", "ResourceUnits", "Subsector", "Quadrant", "WorldX", "WorldY":
		}
	}
	if pd.Ix == "?" {
		pd.Ix = "{+?}"
	}
	if pd.Ex == "?" {
		pd.Ex = "(???+?)"
	}
	if pd.Cx == "?" {
		pd.Cx = "[????]"
	}
	if pd.Nobility == "?" {
		pd.Nobility = "Nobl?"
	}
	return &pd
}

func (pd *parcedData) assignIntData(field string) {
	fDataInt := strings.Split(field, "'':")
	switch fDataInt[0] {
	default:
		pd.err = fmt.Errorf("Unknown field met '%v'", field)
	case "CalculatedImportance":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.CalculatedImportance = val
	case "ResourceUnits":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.ResourceUnits = val
	case "Subsector":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.Subsector = val
	case "Quadrant":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.Quadrant = val
	case "WorldX":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.WorldX = val
	case "WorldY":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.WorldY = val
	case "Worlds":
		val, err := strconv.Atoi(fDataInt[1])
		if err != nil {
			pd.err = fmt.Errorf("Unknown value met '%v' in case %v\n", val, field)
		}
		pd.Worlds = val
	}

}
