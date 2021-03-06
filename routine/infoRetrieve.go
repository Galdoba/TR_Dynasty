package routine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	starport "github.com/Galdoba/TR_Dynasty/Starport"
	trade "github.com/Galdoba/TR_Dynasty/Trade"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/utils"
)

func getShipData(query string) int {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", query)
	data := strings.Split(lines[l], ":")
	jd, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return jd
}

func getShipDataStr(query string) string {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", query)
	data := strings.Split(lines[l], ":")
	return data[1]
}

func getPassengerData(query string) int {
	lines := utils.LinesFromTXT(passengerfile)
	l := utils.InFileContains(passengerfile, query)
	data := strings.Split(lines[l], ":")
	jd, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return jd
}

func getPassengerDataStr(query string) string {
	lines := utils.LinesFromTXT(passengerfile)
	l := utils.InFileContains(passengerfile, query)
	data := strings.Split(lines[l], ":")
	return data[1]
}

// func getJumpDrive() int {
// 	lines := utils.LinesFromTXT("mgt2_traffic.config")
// 	l := utils.InFileContains("mgt2_traffic.config", "JUMP_DRIVE")
// 	data := strings.Split(lines[l], ":")
// 	jd, err := strconv.Atoi(data[1])
// 	if err != nil {
// 		return -999
// 	}
// 	return jd
// }

// func getShipVolume() int {
// 	lines := utils.LinesFromTXT("mgt2_traffic.config")
// 	l := utils.InFileContains("mgt2_traffic.config", "SHIP_VOLUME")
// 	data := strings.Split(lines[l], ":")
// 	vol, err := strconv.Atoi(data[1])
// 	if err != nil {
// 		return -999
// 	}
// 	return vol
// }

// func getShipCargoVolume() int {
// 	lines := utils.LinesFromTXT("mgt2_traffic.config")
// 	l := utils.InFileContains("mgt2_traffic.config", "SHIP_CARGO_VOLUME")
// 	data := strings.Split(lines[l], ":")
// 	vol, err := strconv.Atoi(data[1])
// 	if err != nil {
// 		return -999
// 	}
// 	return vol
// }

func shipArmed() bool {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "SHIP_ARMED")
	data := strings.Split(lines[l], ":")
	switch data[1] {
	default:
		return false
	case "TRUE":
		return true
	}
}

func getCrewSOCdm() int {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "MAX_SOC_DM")
	data := strings.Split(lines[l], ":")
	vol, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return vol
}

func getYear() int {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "CURRENT_YEAR")
	data := strings.Split(lines[l], ":")
	vol, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return vol
}

func getCrewNavyScoutMerchantRank() int {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "MAX_NAVY_SCOUT_MERCAHNT_RANK")
	data := strings.Split(lines[l], ":")
	vol, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return vol
}

func mutateTestResultsByTime(effect, time, timeLimit int) (int, int, bool) {
	time -= effect / 2
	abort := false
	if timeLimit < 1 {
		timeLimit = 99999999
	}
	if time > timeLimit {
		i, _ := menu("Time limits is "+strconv.Itoa(timeLimit)+" days, but operation will take more...", "Do not rush and take your time", "Give result as is", "Abort action")
		switch i {
		case 0:
			timeLimit = time
		case 1:
			dif := time - timeLimit
			effect = effect - (dif * 2)
			time = timeLimit
		case 2:
			effect = 0
			time = timeLimit / 2
			abort = true
		}

	}
	if time < 1 || autoMod == true {
		time = 1
	}
	return effect, time, abort
}

func getCargo() []string {

	lines := utils.LinesFromTXT(cargofile)

	lineNums := utils.InFileContainsN(cargofile, "CARGOENTRY")
	cargo := []string{}

	for _, i := range lineNums {
		currentLine := lines[i]
		data := strings.Split(currentLine, ":")
		dataParts := strings.Split(data[1], "_")
		if len(dataParts) != 14 {
			for e := 0; e < len(dataParts); e++ {
				fmt.Println(e, dataParts[e])
			}
			panic(errors.New("Data Corrupted: " + data[1]))
		}
		cargo = append(cargo, data[1])
	}
	return cargo
}

func shipInfo() string {
	str := "SHIP DATA:\n"
	str += "   Name: " + getShipDataStr("SHIP_NAME") + "\n"
	str += "  Class: " + getShipDataStr("SHIP_CLASS") + " (Type-" + getShipDataStr("SHIP_TYPE") + ")\n"
	str += "Tonnage: " + getShipDataStr("SHIP_VOLUME") + " tons\n"
	str += "J-Drive: " + "Jump-" + getShipDataStr("JUMP_DRIVE") + "\n"
	str += "M-Drive: " + "Thrust " + getShipDataStr("MANUVER_DRIVE") + "\n"
	str += "   Crew: " + getShipDataStr("CURRENT_CREW") + "\n"
	str += "  Cargo: " + getShipDataStr("SHIP_CARGO_VOLUME") + " tons (" + strconv.Itoa(freeCargoVolume()) + " available)\n"
	str += "--------------------------------------------------------------------------------\n"
	cm := loadCargoManifest()
	str += "CURRENT CARGO:\n"
	if len(cm.entry) == 0 {
		str += "NONE"
	}
	for i := range cm.entry {
		info := cm.entry[i].GetTGCode() + "	" + strconv.Itoa(cm.entry[i].GetVolume()) + " tons " + cm.entry[i].GetDescr()
		if cm.entry[i].GetDestination() != "[NO DATA]" {
			planet, _ := otu.GetDataOn(cm.entry[i].GetDestination())
			info += "	(Freight to " + planet.Name() + "  ETA:" + etaDate(cm.entry[i]) + ")"

		}
		str += info + "\n"
	}
	str += "--------------------------------------------------------------------------------\n"
	str += "CURRENT PASSENGERS:\n"
	pm := loadPassengerManifest()
	if len(pm.entry) == 0 {
		str += "NONE"
	}
	for i := range pm.entry {
		str += pm.entry[i].Category() + ": " + pm.entry[i].Name() + "\n"
	}
	return str
}

func getAllStaterooms() (int, int, int, int) {
	lsr, hsr, ssr, lb := 0, 0, 0, 0
	lsr = getShipData("SHIP_STATEROOMS_LUXURY")
	hsr = getShipData("SHIP_STATEROOMS_HIGH")
	ssr = getShipData("SHIP_STATEROOMS_STANDARD")
	lb = getShipData("SHIP_LOWBIRTHS")
	return lsr, hsr, ssr, lb
}

func portInfo() string {
	str := ""
	clrScrn()
	i, val := menu("Select data:", "Services", "Security", "Tariffs", "Market")
	menuPosition += " > " + val
	clrScrn()
	sp, err := starport.From(sourceWorld)
	if err != nil {
		fmt.Println(err)
	}
	switch i {
	default:
	case 0:
		str += sp.Info() + "\n"
	case 1:
		sec := sp.Security()
		str += sec.String() + "\n"
	case 2:
		str += "TODO: Tariffs Info\n"
	case 3:
		str += importExportInfo()
	}
	return str
}

func importExportInfo() string {
	str := "\nPlanet Imports: \n"
	merch := trade.NewMerchant().SetLocalUWP(sourceWorld.UWP()).SetLocalTC(sourceWorld.TradeCodes()).SetMType(constant.MerchantTypeTrade)
	impL, expL := merch.ListImportExport()
	if len(impL) == 0 {
		str += "   Nothing Significant" + "\n"
	}
	for _, val := range impL {
		str += val + "\n"
	}
	str += "\n"
	str += "Planet Exports: \n"
	if len(expL) == 0 {
		str += "   Nothing Significant" + "\n"
	}
	for _, val := range expL {
		str += val + "\n"
	}
	return str
}
