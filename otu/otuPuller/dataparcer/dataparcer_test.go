package dataparcer

import (
	"fmt"
	"strings"
	"testing"
)

func TestParcer(t *testing.T) {
	lines := readData()
	problemLines := 0
	haveSolutions := 0
	for i, line := range lines {
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
				continue
			}
			t.Errorf("parced data on line %v contains error \n%v\n%v", i+1, pd, pd.err.Error())
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
			t.Errorf("parced data (line %v) Subsector field expected to be between 0-15 (have %v)\n%v", i, pd.Subsector, line)
			problemLines++
		case ssExpected(pd.Subsector) != pd.SS:
			t.Errorf("parced data (line %v) Subsector & SS fields expected to be %v and %v (have %v )\n%v\n", i, pd.Subsector, ssExpected(pd.Subsector), pd.SS, line)
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
			t.Errorf("parced data (line %v) does not contain UWP data\n%v", i, line)
			problemLines++
		}

	}
	fmt.Println("Problems discovered:", problemLines, "| Solutions have:", haveSolutions)

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
