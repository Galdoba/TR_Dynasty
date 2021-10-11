package calculations

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
)

func UWPvalid(uwp string) bool {
	hex := strings.Split(uwp, "")
	stpt := ehex.New(hex[0])
	switch stpt.String() {
	default:
		return false
	case "A", "B", "C", "D", "E", "X", "Y", "F", "G", "H", "?":
	}
	switch ehex.New(hex[1]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "?":
	}
	switch ehex.New(hex[2]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "?":
	}
	switch ehex.New(hex[3]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "?":
	}
	switch ehex.New(hex[4]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "?":
	}
	switch ehex.New(hex[5]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "W", "X", "?":
	}
	switch ehex.New(hex[6]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "S", "?":
	}
	switch ehex.New(hex[8]).String() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "?":
	}
	////////////
	return true
}
