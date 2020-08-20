package trademgt2

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/devtools/cli/user"
)

var sectorMapByHex map[string]string
var sectorMapByName map[string]string
var sectorMapByUWP map[string]string

func init() {
	sectorData := otu.TrojanReachData()
	sectorMapByHex = otu.MapDataByHex(sectorData)
	sectorMapByName = otu.MapDataByName(sectorData)
	sectorMapByUWP = otu.MapDataByUWP(sectorData)

}

func RunMerchantPrince() {
	sourceWorld := LoadWorld("Current World (Name, Hex or UWP): ")
	fmt.Println(sourceWorld)
	targetWorld := LoadWorld("Target World (Name, Hex or UWP): ")
	fmt.Println(targetWorld)
}

func LoadWorld(msg string) world.World {
	done := false
	key := ""
	otuData := ""
	for !done {
		key = userInputStr(msg)
		if val, ok := sectorMapByHex[key]; ok {
			otuData = val
			done = true
			continue
		}
		if val, ok := sectorMapByName[key]; ok {
			otuData = val
			done = true
			continue
		}
		key2 := strings.ToUpper(key)
		if val, ok := sectorMapByUWP[key2]; ok {
			otuData = val
			done = true
			continue
		}
		fmt.Println("No data by key '" + key + "'")
	}
	fmt.Println("Loading world data:")
	fmt.Println("Trojan Reach "+otu.Info{otuData}.Hex(), "-", otu.Info{otuData}.Name(), "("+otu.Info{otuData}.UWP()+")")
	w, err := world.FromOTUdata(otuData)
	if err != nil {
		panic(err)
	}
	return w
}

func userInputStr(msg string) string {
	done := false
	fmt.Print(msg)
	for !done {
		uwp, err := user.InputStr()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return uwp
	}
	return "Must not happen !!"
}

func userInputInt(msg string) int {
	done := false
	fmt.Print(msg)
	for !done {
		i, err := user.InputInt()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return i
	}
	return -999
}
