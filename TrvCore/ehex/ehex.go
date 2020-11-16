package ehex

import (
	"fmt"
	"strings"
)

type ehexStruct struct {
	bt byte
}

//DataRetriver -
type DataRetriver interface {
	Value() int
	String() string
}

//TestEhex -
func TestEhex() {

	eh := New("O")
	fmt.Println(eh, eh.Value(), eh.String())

}

//New - Создает Ehex принимая int, string или byte
func New(i interface{}) DataRetriver {
	e := ehexStruct{}
	switch i.(type) {
	default:
		return &ehexStruct{}
	case byte:
		e.bt = i.(byte)
	case string:
		s := i.(string)
		if len(s) < 1 {
			e.bt = byte(130) // Undefined type
			return &e
		}
		if len(s) > 1 {
			s = string([]byte(s)[0])
		}
		s = strings.ToUpper(s)
		e.bt = []byte(s)[0]
	case int:
		if i.(int) > -1 && i.(int) < 34 {
			e.bt = int2Digit(i.(int))
		} else {
			// d := i.(int) % 256
			// e.bt = byte(d)
			return &ehexStruct{}
		}
	}
	e.validate()
	return &e
}

func (e *ehexStruct) validate() {
	switch e.bt {
	default:
		e.bt = 63
	case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57: //digits
	case 42, 45, 63: //special symbols
	case 65, 66, 67, 68, 69, 70, 71, 72, 74, 75, 76, 77, 78, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90: //letters (73 and 79 omited to avoid confusion with 48 and 49)
	}
}

func int2Digit(i int) byte {
	switch i {
	default:
		return byte(0)
	case 0:
		return byte(48)
	case 1:
		return byte(49)
	case 2:
		return byte(50)
	case 3:
		return byte(51)
	case 4:
		return byte(52)
	case 5:
		return byte(53)
	case 6:
		return byte(54)
	case 7:
		return byte(55)
	case 8:
		return byte(56)
	case 9:
		return byte(57)
	case 10:
		return byte(65)
	case 11:
		return byte(66)
	case 12:
		return byte(67)
	case 13:
		return byte(68)
	case 14:
		return byte(69)
	case 15:
		return byte(70)
	case 16:
		return byte(71)
	case 17:
		return byte(72)
	case 18:
		return byte(74)
	case 19:
		return byte(75)
	case 20:
		return byte(76)
	case 21:
		return byte(77)
	case 22:
		return byte(78)
	case 23:
		return byte(80)
	case 24:
		return byte(81)
	case 25:
		return byte(82)
	case 26:
		return byte(83)
	case 27:
		return byte(84)
	case 28:
		return byte(85)
	case 29:
		return byte(86)
	case 30:
		return byte(87)
	case 31:
		return byte(88)
	case 32:
		return byte(89)
	case 33:
		return byte(90)
	}
}

func (e *ehexStruct) String() string {
	return string(e.bt)
}

func (e *ehexStruct) Value() int {
	v := -1
	switch e.bt {
	default:
		return v
	case 42:
		return -2
	case 45:
		return -3
	case 63:
		return -4
	case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
		return int(e.bt) - 48
	case 65, 66, 67, 68, 69, 70, 71, 72:
		return int(e.bt) - 55
	case 74, 75, 76, 77, 78:
		return int(e.bt) - 56
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90:
		return int(e.bt) - 57
	}
}

//Difference - Возвращает радницу между ehexStruct (например dm для TL)
func Difference(ehex1, ehex2 DataRetriver) int {
	return ehex1.Value() - ehex2.Value()
}

// //DigitToEhex -
// func DigitToEhex(num int) string {
// 	ehexStruct := ""
// 	switch num {
// 	default:
// 		return "_"
// 	case 0:
// 		ehexStruct = "0"
// 	case 1:
// 		ehexStruct = "1"
// 	case 2:
// 		ehexStruct = "2"
// 	case 3:
// 		ehexStruct = "3"
// 	case 4:
// 		ehexStruct = "4"
// 	case 5:
// 		ehexStruct = "5"
// 	case 6:
// 		ehexStruct = "6"
// 	case 7:
// 		ehexStruct = "7"
// 	case 8:
// 		ehexStruct = "8"
// 	case 9:
// 		ehexStruct = "9"
// 	case 10:
// 		ehexStruct = "A"
// 	case 11:
// 		ehexStruct = "B"
// 	case 12:
// 		ehexStruct = "C"
// 	case 13:
// 		ehexStruct = "D"
// 	case 14:
// 		ehexStruct = "E"
// 	case 15:
// 		ehexStruct = "F"
// 	case 16:
// 		ehexStruct = "G"
// 	case 17:
// 		ehexStruct = "H"
// 	case 18:
// 		ehexStruct = "J"
// 	case 19:
// 		ehexStruct = "K"
// 	case 20:
// 		ehexStruct = "L"
// 	case 21:
// 		ehexStruct = "M"
// 	case 22:
// 		ehexStruct = "N"
// 	case 23:
// 		ehexStruct = "P"
// 	case 24:
// 		ehexStruct = "Q"
// 	case 25:
// 		ehexStruct = "R"
// 	case 26:
// 		ehexStruct = "S"
// 	case 27:
// 		ehexStruct = "T"
// 	case 28:
// 		ehexStruct = "U"
// 	case 29:
// 		ehexStruct = "V"
// 	case 30:
// 		ehexStruct = "W"
// 	case 31:
// 		ehexStruct = "X"
// 	case 32:
// 		ehexStruct = "Y"
// 	case 33:
// 		ehexStruct = "Z"
// 	}
// 	return ehexStruct
// }

// //EhexToDigit -
// func EhexToDigit(lit string) int {
// 	num := -999
// 	if len(lit) != 1 {
// 		return -999
// 	}
// 	switch lit {
// 	case "0":
// 		num = 0
// 	case "1":
// 		num = 1
// 	case "2":
// 		num = 2
// 	case "3":
// 		num = 3
// 	case "4":
// 		num = 4
// 	case "5":
// 		num = 5
// 	case "6":
// 		num = 6
// 	case "7":
// 		num = 7
// 	case "8":
// 		num = 8
// 	case "9":
// 		num = 9
// 	case "A":
// 		num = 10
// 	case "B":
// 		num = 11
// 	case "C":
// 		num = 12
// 	case "D":
// 		num = 13
// 	case "E":
// 		num = 14
// 	case "F":
// 		num = 15
// 	case "G":
// 		num = 16
// 	case "H":
// 		num = 17
// 	case "J":
// 		num = 18
// 	case "K":
// 		num = 19
// 	case "L":
// 		num = 20
// 	case "M":
// 		num = 21
// 	case "N":
// 		num = 22
// 	case "P":
// 		num = 23
// 	case "Q":
// 		num = 24
// 	case "R":
// 		num = 25
// 	case "S":
// 		num = 26
// 	case "T":
// 		num = 27
// 	case "U":
// 		num = 28
// 	case "V":
// 		num = 29
// 	case "W":
// 		num = 30
// 	case "X":
// 		num = 31
// 	case "Y":
// 		num = 32
// 	case "Z":
// 		num = 33
// 	case "_":
// 		num = 100
// 	}
// 	return num
// }
