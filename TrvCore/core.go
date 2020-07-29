package TrvCore

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"
)

//Roll1D -
func Roll1D(dm ...int) int {
	r := dice.Roll("d6")
	if len(dm) > 0 {
		r.DM(dm[0])
	}
	return r.Sum()
}

//Roll2D -
func Roll2D(mods ...int) int {
	r := utils.RollDice("2d6")
	for i := range mods {
		r = r + mods[i]
	}
	return r
}

//Roll3D -
func Roll3D(mods ...int) int {
	r := utils.RollDice("3d6")
	for i := range mods {
		r = r + mods[i]
	}
	return r
}

func RollD66() string {
	return dice.Roll("2d6").ResultString()
}

// func RerollIf(roll func(...int) int, notWant ...int) int {
// 	i := 0
// 	for {
// 		suggest := roll()
// 		met := false
// 		for i := range notWant {
// 			if suggest == notWant[i] {
// 				met = true
// 			}
// 		}
// 		if !met {
// 			return suggest
// 		}
// 		i++
// 		if i > 100000 {
// 			return suggest
// 		}
// 	}
// }

//Flux -
func Flux() int {
	d1 := Roll1D()
	d2 := Roll1D()
	return d1 - d2
}

//FluxGood -
func FluxGood() int {
	d1 := Roll1D()
	d2 := Roll1D()
	flux := 0
	if d1 > d2 {
		flux = d1 - d2
	} else {
		flux = d2 - d1
	}
	return flux
}

//FluxBad -
func FluxBad() int {
	d1 := Roll1D()
	d2 := Roll1D()
	flux := 0
	if d1 > d2 {
		flux = d2 - d1
	} else {
		flux = d1 - d2
	}
	return flux
}

//DigitToEhex -
func DigitToEhex(num int) string {
	ehex := ""
	switch num {
	default:
		return "_"
	case 0:
		ehex = "0"
	case 1:
		ehex = "1"
	case 2:
		ehex = "2"
	case 3:
		ehex = "3"
	case 4:
		ehex = "4"
	case 5:
		ehex = "5"
	case 6:
		ehex = "6"
	case 7:
		ehex = "7"
	case 8:
		ehex = "8"
	case 9:
		ehex = "9"
	case 10:
		ehex = "A"
	case 11:
		ehex = "B"
	case 12:
		ehex = "C"
	case 13:
		ehex = "D"
	case 14:
		ehex = "E"
	case 15:
		ehex = "F"
	case 16:
		ehex = "G"
	case 17:
		ehex = "H"
	case 18:
		ehex = "J"
	case 19:
		ehex = "K"
	case 20:
		ehex = "L"
	case 21:
		ehex = "M"
	case 22:
		ehex = "N"
	case 23:
		ehex = "P"
	case 24:
		ehex = "Q"
	case 25:
		ehex = "R"
	case 26:
		ehex = "S"
	case 27:
		ehex = "T"
	case 28:
		ehex = "U"
	case 29:
		ehex = "V"
	case 30:
		ehex = "W"
	case 31:
		ehex = "X"
	case 32:
		ehex = "Y"
	case 33:
		ehex = "Z"
	}
	return ehex
}

//EhexToDigit -
func EhexToDigit(lit string) int {
	num := -999
	if len(lit) != 1 {
		return -999
	}
	switch lit {
	case "0":
		num = 0
	case "1":
		num = 1
	case "2":
		num = 2
	case "3":
		num = 3
	case "4":
		num = 4
	case "5":
		num = 5
	case "6":
		num = 6
	case "7":
		num = 7
	case "8":
		num = 8
	case "9":
		num = 9
	case "A":
		num = 10
	case "B":
		num = 11
	case "C":
		num = 12
	case "D":
		num = 13
	case "E":
		num = 14
	case "F":
		num = 15
	case "G":
		num = 16
	case "H":
		num = 17
	case "J":
		num = 18
	case "K":
		num = 19
	case "L":
		num = 20
	case "M":
		num = 21
	case "N":
		num = 22
	case "P":
		num = 23
	case "Q":
		num = 24
	case "R":
		num = 25
	case "S":
		num = 26
	case "T":
		num = 27
	case "U":
		num = 28
	case "V":
		num = 29
	case "W":
		num = 30
	case "X":
		num = 31
	case "Y":
		num = 32
	case "Z":
		num = 33
	case "_":
		num = 100
	}
	return num
}

func EhexIsMore(ehex1, ehex2 string) bool {
	if EhexToDigit(ehex1) > EhexToDigit(ehex2) {
		return true
	}
	return false
}

func ValidEhexs() []string {
	return []string{
		"_",
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"J",
		"K",
		"L",
		"M",
		"N",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z"}
}

func GreekToNum(str string) int {
	switch str {
	default:
		return -1
	case "Alpha":
		return 0
	case "Beta":
		return 1
	case "Gamma":
		return 2
	case "Delta":
		return 3
	case "Epsilon":
		return 4
	case "Zeta":
		return 5
	case "Eta":
		return 6
	case "Theta":
		return 7
	case "Iota":
		return 8
	case "Kappa":
		return 9
	case "Lambda":
		return 10
	case "Mu":
		return 11
	case "Nu":
		return 12
	case "Xi":
		return 13
	case "Omicron":
		return 14
	case "Pi":
		return 15
	case "Rho":
		return 16
	case "Sigma":
		return 17
	case "Tau":
		return 18
	case "Upsilon":
		return 19
	case "Phi":
		return 20
	case "Chi":
		return 21
	case "Psi":
		return 22
	case "Omega":
		return 23
	}
}

func NumToGreek(i int) string {
	switch i {
	default:
		return "UNDEFINED"
	case 0:
		return "Alpha"
	case 1:
		return "Beta"
	case 2:
		return "Gamma"
	case 3:
		return "Delta"
	case 4:
		return "Epsilon"
	case 5:
		return "Zeta"
	case 6:
		return "Eta"
	case 7:
		return "Theta"
	case 8:
		return "Iota"
	case 9:
		return "Kappa"
	case 10:
		return "Lambda"
	case 11:
		return "Mu"
	case 12:
		return "Nu"
	case 13:
		return "Xi"
	case 14:
		return "Omicron"
	case 15:
		return "Pi"
	case 16:
		return "Rho"
	case 17:
		return "Sigma"
	case 18:
		return "Tau"
	case 19:
		return "Upsilon"
	case 20:
		return "Phi"
	case 21:
		return "Chi"
	case 22:
		return "Psi"
	case 23:
		return "Omega"
	}
}

func AnglicToNum(str string) int {
	switch str {
	default:
		return -1
	case "Ay":
		return 0
	case "Bee":
		return 1
	case "Cee":
		return 2
	case "Dee":
		return 3
	case "Ee":
		return 4
	case "Eff":
		return 5
	case "Gee":
		return 6
	case "Aitch":
		return 7
	case "Eye":
		return 8
	case "Jay":
		return 9
	case "Kay":
		return 10
	case "Ell":
		return 11
	case "Em":
		return 12
	case "Oh":
		return 13
	case "Pee":
		return 14
	case "Cue":
		return 15
	case "Arr":
		return 16
	case "Ess":
		return 17
	case "Tee":
		return 18
	case "You":
		return 19
	case "Vee":
		return 20
	case "Double-You":
		return 22
	case "Eks":
		return 23
	case "Wye":
		return 24
	case "Zee":
		return 25
	}
}

func NumToAnglic(i int) string {
	switch i {
	default:
		return "UNDEFINED"
	case 0:
		return "Ay"
	case 1:
		return "Bee"
	case 2:
		return "Cee"
	case 3:
		return "Dee"
	case 4:
		return "Ee"
	case 5:
		return "Eff"
	case 6:
		return "Gee"
	case 7:
		return "Aitch"
	case 8:
		return "Eye"
	case 9:
		return "Jay"
	case 10:
		return "Kay"
	case 11:
		return "Ell"
	case 12:
		return "Em"
	case 13:
		return "En"
	case 14:
		return "Oh"
	case 15:
		return "Pee"
	case 16:
		return "Cue"
	case 17:
		return "Arr"
	case 18:
		return "Ess"
	case 19:
		return "Tee"
	case 20:
		return "You"
	case 21:
		return "Vee"
	case 22:
		return "Double-You"
	case 23:
		return "Eks"
	case 24:
		return "Wye"
	case 25:
		return "Zee"
	}
}

type Ehex interface {
	Value() int
	Glyph() string
}

type ehex struct {
	val   int
	glyph string
}

func (ex *ehex) Value() int {
	return ex.val
}

func (ex *ehex) Glyph() string {
	return ex.glyph
}

func EHex(data interface{}) Ehex {
	var ehex ehex
	switch data.(type) {
	default:
		fmt.Println(data)
		return nil
	case byte:
		s := string(data.(byte))
		ehex.glyph = s
		ehex.val = EhexToDigit(s)
	case string:
		s := data.(string)
		fmt.Println(s)
		ehex.glyph = data.(string)
		ehex.val = EhexToDigit(data.(string))
	case int:
		ehex.val = data.(int)
		ehex.glyph = DigitToEhex(data.(int))
	}
	return &ehex
}
