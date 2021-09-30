package dataparcer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

func readData() []string {
	//lines := utils.LinesFromTXT("c:\\Users\\pemaltynov\\go\\src\\github.com\\Galdoba\\TR_Dynasty\\otu\\otuPuller\\data.txt")

	lines := utils.LinesFromTXT("d:\\golang\\src\\github.com\\Galdoba\\TR_Dynasty\\otu\\otuPuller\\data.txt")

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
	//OrbitCode            string //W
	//Star                 string //W
	err error
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
	worlddata = strings.Replace(worlddata, ",'',''", ",''", -1)
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
	if pd.SectorAbbreviation == "?" {
		pd.SectorAbbreviation = "****"
	}
	if pd.Name == "?" {
		pd.Name = ""
	}

	return &pd
}

func (pd *parcedData) String() string {
	return fmt.Sprint(
		pd.Name, "|",
		pd.Hex, "|",
		pd.UWP, "|",
		pd.PBG, "|",
		pd.Zone, "|",
		pd.Bases, "|",
		pd.Allegiance, "|",
		pd.Stellar, "|",
		pd.SS, "|",
		pd.Ix, "|",
		pd.CalculatedImportance, "|",
		pd.Ex, "|",
		pd.Cx, "|",
		pd.Nobility, "|",
		pd.Worlds, "|",
		pd.ResourceUnits, "|",
		pd.Subsector, "|",
		pd.Quadrant, "|",
		pd.WorldX, "|",
		pd.WorldY, "|",
		pd.Remarks, "|",
		pd.LegacyBaseCode, "|",
		pd.Sector, "|",
		pd.SubsectorName, "|",
		pd.SectorAbbreviation, "|",
		pd.AllegianceName, "|",
	)
}

func (pd *parcedData) StringF() string {
	str := ""
	Name := pd.Name
	for len(Name) < 28 {
		Name += " "
	}
	str += "|"
	str += Name
	Hex := pd.Hex
	for len(Hex) < 4 {
		Hex += " "
	}
	str += "|"
	str += Hex
	UWP := pd.UWP
	for len(UWP) < 9 {
		UWP += " "
	}
	str += "|"
	str += UWP
	PBG := pd.PBG
	for len(PBG) < 3 {
		PBG += " "
	}
	str += "|"
	str += PBG
	Zone := pd.Zone
	for len(Zone) < 1 {
		Zone += " "
	}
	str += "|"
	str += Zone
	Bases := pd.Bases
	for len(Bases) < 2 {
		Bases += " "
	}
	str += "|"
	str += Bases
	Allegiance := pd.Allegiance
	for len(Allegiance) < 4 {
		Allegiance += " "
	}
	str += "|"
	str += Allegiance
	Stellar := pd.Stellar
	for len(Stellar) < 29 {
		Stellar += " "
	}
	str += "|"
	str += Stellar
	SS := pd.SS
	for len(SS) < 1 {
		SS += " "
	}
	str += "|"
	str += SS
	Ix := pd.Ix
	for len(Ix) < 11 {
		Ix += " "
	}
	str += "|"
	str += Ix
	CalculatedImportance := strconv.Itoa(pd.CalculatedImportance)
	for len(CalculatedImportance) < 2 {
		CalculatedImportance += " "
	}
	str += "|"
	str += CalculatedImportance
	Ex := pd.Ex
	for len(Ex) < 7 {
		Ex += " "
	}
	str += "|"
	str += Ex
	Cx := pd.Cx
	for len(Cx) < 6 {
		Cx += " "
	}
	str += "|"
	str += Cx
	Nobility := pd.Nobility
	for len(Nobility) < 5 {
		Nobility += " "
	}
	str += "|"
	str += Nobility
	Worlds := strconv.Itoa(pd.Worlds)
	for len(Worlds) < 2 {
		Worlds += " "
	}
	str += "|"
	str += Worlds
	ResourceUnits := strconv.Itoa(pd.ResourceUnits)
	for len(ResourceUnits) < 5 {
		ResourceUnits += " "
	}
	str += "|"
	str += ResourceUnits
	Subsector := strconv.Itoa(pd.Subsector)
	for len(Subsector) < 2 {
		Subsector += " "
	}
	str += "|"
	str += Subsector
	Quadrant := strconv.Itoa(pd.Quadrant)
	for len(Quadrant) < 1 {
		Quadrant += " "
	}
	str += "|"
	str += Quadrant
	WorldX := strconv.Itoa(pd.WorldX)
	for len(WorldX) < 5 {
		WorldX += " "
	}
	str += "|"
	str += WorldX
	WorldY := strconv.Itoa(pd.WorldY)
	for len(WorldY) < 5 {
		WorldY += " "
	}
	str += "|"
	str += WorldY
	Remarks := pd.Remarks
	for len(Remarks) < 44 {
		Remarks += " "
	}
	str += "|"
	str += Remarks
	LegacyBaseCode := pd.LegacyBaseCode
	for len(LegacyBaseCode) < 2 {
		LegacyBaseCode += " "
	}
	str += "|"
	str += LegacyBaseCode
	Sector := pd.Sector
	for len(Sector) < 33 {
		Sector += " "
	}
	str += "|"
	str += Sector
	SubsectorName := pd.SubsectorName
	for len(SubsectorName) < 24 {
		SubsectorName += " "
	}
	str += "|"
	str += SubsectorName
	SectorAbbreviation := pd.SectorAbbreviation
	for len(SectorAbbreviation) < 4 {
		SectorAbbreviation += " "
	}
	str += "|"
	str += SectorAbbreviation
	AllegianceName := pd.AllegianceName
	for len(AllegianceName) < 93 {
		AllegianceName += " "
	}
	str += "|"
	str += AllegianceName
	return str
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
