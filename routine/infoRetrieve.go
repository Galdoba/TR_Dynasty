package routine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/utils"
)

func getJumpDrive() int {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "JUMP_DRIVE")
	data := strings.Split(lines[l], ":")
	jd, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return jd
}

func getShipVolume() int {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "SHIP_VOLUME")
	data := strings.Split(lines[l], ":")
	vol, err := strconv.Atoi(data[1])
	if err != nil {
		return -999
	}
	return vol
}

func shipArmed() bool {
	lines := utils.LinesFromTXT("mgt2_traffic.config")
	l := utils.InFileContains("mgt2_traffic.config", "SHIP_VOLUME")
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

	lines := utils.LinesFromTXT(exPath + "mgt2_traffic.config")
	lineNums := utils.InFileContainsN("mgt2_traffic.config", "CARGOENTRY")
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
