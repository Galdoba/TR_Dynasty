package dataparcer

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestParcer(t *testing.T) {
	lines := readData()
	problemLines := 0
	haveSolutions := 0
	blankSystems := 0
	systemsSurveed := 0
	//dataMap := make(map[int]int)
	// f, err := os.Create("formattedData.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	for i, line := range lines {
		systemsSurveed++
		// if line == "{''Worlds'':[]}" {
		// 	continue
		// }

		if i == 200 {
			//os.Exit(200)
		}

		pd := ParceWorldData(line)

		switch {
		case pd.err != nil:
			if pd.err.Error() == "worlddata is blank" {
				blankSystems++
				continue
			}
			t.Errorf("parced data on line %v contains error \n%v\n%v", i+1, pd, pd.err.Error())
			panic(1)
		case pd.WorldX == 0 && pd.WorldY == 0 && pd.Name != "Reference":
			t.Errorf("parced data (line %v) does not have World Coordinates\n%v\n%v", i, line, pd)
		case pd.Name == "?":
			t.Errorf("parced data (line %v) does not contain Name field\n%v", i, line)
			panic(1)
		case pd.Hex == "?":
			t.Errorf("parced data (line %v) does not contain Hex field\n%v", i, line)
			panic(1)
		case pd.UWP == "?":
			t.Errorf("parced data (line %v) does not contain UWP field\n%v", i, line)
			panic(1)
		case pd.PBG == "?":
			t.Errorf("parced data (line %v) does not contain PBG field\n%v", i, line)
			panic(1)
		case pd.Zone == "?":
			t.Errorf("parced data (line %v) does not contain Zone field\n%v", i, line)
			panic(1)
		case pd.Bases == "?":
			t.Errorf("parced data (line %v) does not contain Bases field\n%v", i, line)
			panic(1)
		case pd.Allegiance == "?":
			t.Errorf("parced data (line %v) does not contain Allegiance field\n%v", i, line)
			panic(1)
		case pd.Stellar == "?":
			t.Errorf("parced data (line %v) does not contain Stellar field\n%v", i, line)
			panic(1)
		case pd.SS == "?":
			t.Errorf("parced data (line %v) does not contain SS field\n%v", i, line)
			panic(1)
		case pd.Ix == "?":
			t.Errorf("parced data (line %v) does not contain Ix field\n%v", i, line)
			panic(1)
		// case pd.CalculatedImportance == 0:
		// 	t.Errorf("parced data (line %v) does not contain CalculatedImportance field\n%v", i, line)
		// 	panic(1)
		case pd.Ex == "?":
			t.Errorf("parced data (line %v) does not contain Ex field\n%v", i, line)
			panic(1)
		case pd.Cx == "?":
			t.Errorf("parced data (line %v) does not contain Cx field\n%v", i, line)
			panic(1)
		case pd.Nobility == "?":
			t.Errorf("parced data (line %v) does not contain Nobility field\n%v", i, line)
			panic(1)
		case pd.Worlds == 0 && pd.Name != "":
			//t.Errorf("Line %v \nWorlds Expected to be > 0\n%v\nSOLUTION: Recalculate Worlds T5 CB3 p29\n", i, line)
			problemLines++
			haveSolutions++
			// case pd.ResourceUnits == 0:
			// t.Errorf("parced data (line %v) does not contain ResourceUnits field\n%v", i, line)
			// panic(1)
		case pd.Subsector < 0 || pd.Subsector > 15:
			//t.Errorf("parced data (line %v) Subsector field expected to be between 0-15 (have %v)\n%v", i, pd.Subsector, line)
			problemLines++
		case ssExpected(pd.Subsector) != pd.SS:
			//t.Errorf("parced data (line %v) Subsector & SS fields expected to be %v and %v (have %v )\n%v\n", i, pd.Subsector, ssExpected(pd.Subsector), pd.SS, line)
			problemLines++
		case quadrantExpected(pd.Subsector) != pd.Quadrant:
			t.Errorf("parced data (line %v) Subsector & Quadrant fields expected to be %v and %v (have %v )\n%v\n", i, pd.Subsector, quadrantExpected(pd.Subsector), pd.Quadrant, line)
			problemLines++
		// panic(1)
		// case pd.Quadrant == 0:
		// t.Errorf("parced data (line %v) does not contain Quadrant field\n%v", i, line)
		// panic(1)
		// case pd.WorldX == 0:
		// t.Errorf("parced data (line %v) does not contain WorldX field\n%v", i, line)
		// panic(1)
		// case pd.WorldY == 0:
		// t.Errorf("parced data (line %v) does not contain WorldY field\n%v", i, line)
		// panic(1)
		case pd.Remarks == "?":
			t.Errorf("parced data (line %v) does not contain Remarks field\n%v", i, line)
			panic(1)
		case pd.LegacyBaseCode == "?":
			t.Errorf("parced data (line %v) does not contain LegacyBaseCode field\n%v", i, line)
			panic(1)
		case pd.Sector == "?":
			t.Errorf("parced data (line %v) does not contain Sector field\n%v", i, line)
			panic(1)
		case pd.SubsectorName == "?":
			t.Errorf("parced data (line %v) does not contain SubsectorName field\n%v", i, line)
			panic(1)
		case pd.SectorAbbreviation == "?":
			t.Errorf("parced data (line %v) does not contain SectorAbbreviation field\n%v", i, line)
			panic(1)
		case pd.AllegianceName == "?":
			t.Errorf("parced data (line %v) does not contain AllegianceName field\n%v", i, line)
			panic(1)
		case pd.PBG == "":
			t.Errorf("parced data (line %v) does not contain PBG data\n%v", i, line)
			panic(1)
		case strings.Contains(pd.UWP, "?"):
			//t.Errorf("parced data (line %v) does not contain UWP data\n%v", i, line)
			problemLines++
		}
		//dataMap = updateDataTemplate(pd, dataMap)
		// _, errW := f.WriteString(pd.StringF() + "\n")
		// if errW != nil {
		// 	fmt.Println(err)
		// 	f.Close()
		// }
		fmt.Print("Parsed ", systemsSurveed, "/", len(lines), "\r")
	}
	fmt.Println("")
	fmt.Println("------------------------------")
	fmt.Println("Total Systems Parsed:", systemsSurveed)
	fmt.Println("Problems discovered:", problemLines)
	fmt.Println("Solutions have:", haveSolutions)
	fmt.Println("Blank entryes:", blankSystems)
	fmt.Println("------------------------------")
	// for i := 0; i < 30; i++ {
	// 	fmt.Println(i, "=", dataMap[i])
	// }

}

func updateDataTemplate(pd *parcedData, dataMap map[int]int) map[int]int {
	i := 0
	if len(pd.Name) > dataMap[i] {
		dataMap[i] = len(pd.Name)
	}
	i++
	if len(pd.Hex) > dataMap[i] {
		dataMap[i] = len(pd.Hex)
	}
	i++
	if len(pd.UWP) > dataMap[i] {
		dataMap[i] = len(pd.UWP)
	}
	i++
	if len(pd.PBG) > dataMap[i] {
		dataMap[i] = len(pd.PBG)
	}
	i++
	if len(pd.Zone) > dataMap[i] {
		dataMap[i] = len(pd.Zone)
	}
	i++
	if len(pd.Bases) > dataMap[i] {
		dataMap[i] = len(pd.Bases)
	}
	i++
	if len(pd.Allegiance) > dataMap[i] {
		dataMap[i] = len(pd.Allegiance)
	}
	i++
	if len(pd.Stellar) > dataMap[i] {
		dataMap[i] = len(pd.Stellar)
	}
	i++
	if len(pd.SS) > dataMap[i] {
		dataMap[i] = len(pd.SS)
	}
	i++
	if len(pd.Ix) > dataMap[i] {
		dataMap[i] = len(pd.Ix)
	}
	i++
	if len(strconv.Itoa(pd.CalculatedImportance)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.CalculatedImportance))
	}
	i++
	if len(pd.Ex) > dataMap[i] {
		dataMap[i] = len(pd.Ex)
	}
	i++
	if len(pd.Cx) > dataMap[i] {
		dataMap[i] = len(pd.Cx)
	}
	i++
	if len(pd.Nobility) > dataMap[i] {
		dataMap[i] = len(pd.Nobility)
	}
	i++
	if len(strconv.Itoa(pd.Worlds)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.Worlds))
	}
	i++
	if len(strconv.Itoa(pd.ResourceUnits)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.ResourceUnits))
	}
	i++
	if len(strconv.Itoa(pd.Subsector)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.Subsector))
	}
	i++
	if len(strconv.Itoa(pd.Quadrant)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.Quadrant))
	}
	i++
	if len(strconv.Itoa(pd.WorldX)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.WorldX))
	}
	i++
	if len(strconv.Itoa(pd.WorldY)) > dataMap[i] {
		dataMap[i] = len(strconv.Itoa(pd.WorldY))
	}
	i++
	if len(pd.Remarks) > dataMap[i] {
		dataMap[i] = len(pd.Remarks)
	}
	i++
	if len(pd.LegacyBaseCode) > dataMap[i] {
		dataMap[i] = len(pd.LegacyBaseCode)
	}
	i++
	if len(pd.Sector) > dataMap[i] {
		dataMap[i] = len(pd.Sector)
	}
	i++
	if len(pd.SubsectorName) > dataMap[i] {
		dataMap[i] = len(pd.SubsectorName)
	}
	i++
	if len(pd.SectorAbbreviation) > dataMap[i] {
		dataMap[i] = len(pd.SectorAbbreviation)
	}
	i++
	if len(pd.AllegianceName) > dataMap[i] {
		dataMap[i] = len(pd.AllegianceName)
	}
	return dataMap
}

func ssExpected(sub int) string {
	ssAr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
	switch sub {
	default:
		return "unexpected"
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		return ssAr[sub]
	}
}

func quadrantExpected(sub int) int {
	switch sub {
	case 0, 1, 4, 5:
		return 0
	case 2, 3, 6, 7:
		return 1
	case 8, 9, 12, 13:
		return 4 //хз почему так
	case 10, 11, 14, 15:
		return 5 //хз почему так
	}
	return -1
}
