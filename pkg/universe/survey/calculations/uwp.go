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

func Fix(uwp string) string {
	switch uwp {
	case "X233000-X":
		return "X233000-0"
	case "X420000-X":
		return "X420000-0"
	case "X400000-X":
		return "X400000-0"
	case "X100000-X":
		return "X100000-0"
	case "X7A6000-X":
		return "X7A6000-0"
	case "X424000-X":
		return "X424000-0"
	case "X411000-X":
		return "X411000-0"
	case "X110000-X":
		return "X110000-0"
	case "X000000-X":
		return "X000000-0"
	case "X439000-X":
		return "X439000-0"
	case "X000XXX-X":
		return "X000000-0"
	case "B453889-X":
		return "B453889-0"
	case "X200000-X":
		return "X200000-0"
	case "X484XXX-X":
		return "X484000-0"
	case "C857360-N":
		return "C857360-L"
	case "A6VV997-D":
		return "A655997-D"
	case "XXXXXXX-X":
		return "???????-?"
	}
	return uwp
}
