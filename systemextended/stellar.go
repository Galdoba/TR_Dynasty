package systemextended

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/pkg/core/astronomical"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/utils"
)

type SysyemMapConstructor struct {
	Dicepool *dice.Dicepool
}

type starSystemMap struct {
}

func NewSystemMapConstructor(w wrld.World) SysyemMapConstructor {
	//testOnData()
	smc := SysyemMapConstructor{}
	smc.Dicepool = w.DicePool()
	stars := parceStars(w.Stellar())
	starMap := generateSystemStarmap(w.DicePool(), stars)
	for i := range stars {
		fmt.Println(stars[i], starMap[i])
	}

	return smc
}

func generateSystemStarmap(dp *dice.Dicepool, stars []string) []string {
	starNumMatch := false
	foundStars := []string{}
	for !starNumMatch {
		foundStars = []string{"P"}
		if dp.FluxNext() >= 4 {
			foundStars = append(foundStars, "Pc")
		}
		for _, val := range []string{"C", "N", "F"} {
			if dp.FluxNext() >= 4 {
				foundStars = append(foundStars, val)
				if dp.FluxNext() >= 4 {
					foundStars = append(foundStars, val+"c")
				}
			}
		}
		starNumMatch = (len(stars) == len(foundStars))
	}

	return foundStars
}

func parceStars(stellar string) (stars []string) {
	data := strings.Split(stellar, " ")
	starDetected := ""
	for i := 0; i < len(data); i++ {
		switch data[i] {
		default:
			if starDetected != "" {
				stars = append(stars, starDetected)
			}
			starDetected = data[i]
		case "Ia", "Ib", "II", "III", "IV", "V", "VI", "VII":
			starDetected = starDetected + " " + data[i]
			stars = append(stars, starDetected)
			starDetected = ""
		}
	}
	if starDetected != "" {
		stars = append(stars, starDetected)
	}
	return
}

func testOnData() {
	varMap := make(map[string]int)
	varKeys := []string{}
	// for _, data := range utils.LinesFromTXT("d:\\golang\\src\\github.com\\Galdoba\\TR_Dynasty\\data.txt") {
	// 	parts := strings.Split(data, "	")
	// 	stellarData := parts[10]
	// 	stelParts := strings.Split(stellarData, " ")
	// 	for _, sp := range stelParts {
	// 		varMap[sp]++
	// 		varKeys = utils.AppendUniqueStr(varKeys, sp)
	// 	}
	// }
	for _, data := range utils.LinesFromTXT("c:\\Users\\pemaltynov\\go\\src\\github.com\\Galdoba\\TR_Dynasty\\data.txt") {
		parts := strings.Split(data, "	")
		stellarData := parts[10]
		stars := parceStars(stellarData)
		fmt.Print("Try '", stellarData, "\n")
		for _, val := range stars {
			fmt.Print("'", val, "'\n")
		}
		if len(stars) > 1 {
			time.Sleep(time.Second)
		}
		stelParts := strings.Split(stellarData, " ")
		for _, sp := range stelParts {
			varMap[sp]++
			varKeys = utils.AppendUniqueStr(varKeys, sp)
		}
	}
	for i, key := range varKeys {
		fmt.Print(i, " '", key, "' ", varMap[key], "\n")
	}
	os.Exit(0)
}

type systemCoordinates struct {
	starPosition  int
	planetOrbit   int
	sateliteOrbit int
}

// W Total Worlds In System= MW+GG+Belts+2D
// P Mainworld Placement.
// P Gas Giant Placement
// P Planetoid Belt Placement
// P Create other Worlds

func newCoordinates(str, plnt, sat int) systemCoordinates {
	return systemCoordinates{
		starPosition:  str,
		planetOrbit:   plnt,
		sateliteOrbit: sat,
	}
}

func defineHZ(w wrld.World) int {
	primary := astronomical.NewStellarData(w.Stellar())
	fmt.Println(primary)
	fmt.Println("'" + w.Stellar() + "'")
	fmt.Println(parseStellarData(w))
	fmt.Println(astronomical.DecodeStellar(w.Stellar()))
	return 0
}

func parseStellarData(w wrld.World) []string {
	spectralData := strings.Split(w.Stellar(), " ")
	if w.Stellar() == "" {
		panic("\n\n---------------------------------------------------------\nTODO: решить что делать если не хватет официальных данных\n---------------------------------------------------------\n")
	}
	stars := []string{}
	for i, val := range spectralData {
		switch val {
		default:
			if i >= len(spectralData) {
				fmt.Println("Error: i >= len(spectralData)", i, len(spectralData))
				continue
			}
			stars = append(stars, spectralData[i]+" "+spectralData[i+1])
		case "BD":
			stars = append(stars, val)
		case "Ia", "Ib", "II", "III", "IV", "V", "VI", "D":
			continue
		}
	}
	return stars
}
