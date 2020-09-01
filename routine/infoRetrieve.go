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
