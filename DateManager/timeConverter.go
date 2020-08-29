package DateManager

import (
	"strconv"

	"github.com/Galdoba/utils"
)

func HoursToDaysStr(hours float64) string {
	days := int(hours / 24)
	hoursLeft := hours - float64(days*24)
	hoursLeft = utils.RoundFloat64(hoursLeft, 1)
	str := ""
	if days > 0 {
		str += strconv.Itoa(days) + " day"
		if days > 1 {
			str += "s"
		}
	}
	if hoursLeft != 0 {
		str += " " + float64ToString(hoursLeft) + " hours"
	}
	return str
}

func float64ToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 1, 64)
}

func HoursToDays(hours float64) float64 {
	days := hours / 24.0
	days = utils.RoundFloat64(days, 1)
	return days
}
