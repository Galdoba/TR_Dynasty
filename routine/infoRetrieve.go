package routine

import (
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
