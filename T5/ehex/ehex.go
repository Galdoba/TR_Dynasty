package ehex

const (
	UNKNOWN   = "unknown"
	SPECIAL   = "special"
	ULTIMATE  = "ultimate"
	ANY_VALUE = "any value"
)

type ehex struct {
	value   int
	code    string
	comment string
}

func New() *ehex {
	return &ehex{value: 0, code: "*", comment: ANY_VALUE}
}

func (eh *ehex) Set(data interface{}) *ehex {
	eh.comment = ""
	eh.code = ""
	switch data.(type) {
	default:
		eh.code = "?"
		eh.comment = UNKNOWN
	case string:
		if hashCode(data.(string)) != -1 {
			eh.code = data.(string)
		}
	case int:
		eh.code = hashValue(data.(int))
	}
	eh.value = hashCode(eh.code)
	switch eh.code {
	case "X", "?":
		eh.comment = UNKNOWN
	case "Y":
		eh.comment = SPECIAL
	case "Z":
		eh.comment = ULTIMATE
	case "*":
		eh.comment = ANY_VALUE
	default:
	}
	return eh
}

///////INTERFACE
type Ehex interface {
	Value() int
	Code() string
	Meaning() string
}

func (e *ehex) Value() int {
	return e.value
}

func (e *ehex) Code() string {
	return e.code
}

func (e *ehex) Meaning() string {
	return e.comment
}

func (e *ehex) String() string {
	return e.code
}

///////HASH

func hashValue(value int) string {
	codeMap := make(map[int]string)
	codeMap[0] = "0"
	codeMap[1] = "1"
	codeMap[2] = "2"
	codeMap[3] = "3"
	codeMap[4] = "4"
	codeMap[5] = "5"
	codeMap[6] = "6"
	codeMap[7] = "7"
	codeMap[8] = "8"
	codeMap[9] = "9"
	codeMap[10] = "A"
	codeMap[11] = "B"
	codeMap[12] = "C"
	codeMap[13] = "D"
	codeMap[14] = "E"
	codeMap[15] = "F"
	codeMap[16] = "G"
	codeMap[17] = "H"
	codeMap[18] = "J"
	codeMap[19] = "K"
	codeMap[20] = "L"
	codeMap[21] = "M"
	codeMap[22] = "N"
	codeMap[23] = "P"
	codeMap[24] = "Q"
	codeMap[25] = "R"
	codeMap[26] = "S"
	codeMap[27] = "T"
	codeMap[28] = "U"
	codeMap[29] = "V"
	codeMap[30] = "W"
	codeMap[31] = "X"
	codeMap[32] = "Y"
	codeMap[33] = "Z"
	if val, ok := codeMap[value]; ok {
		return val
	}
	return "?"
}

func hashCode(code string) int {
	valMap := make(map[string]int)
	valMap["0"] = 0
	valMap["1"] = 1
	valMap["2"] = 2
	valMap["3"] = 3
	valMap["4"] = 4
	valMap["5"] = 5
	valMap["6"] = 6
	valMap["7"] = 7
	valMap["8"] = 8
	valMap["9"] = 9
	valMap["A"] = 10
	valMap["B"] = 11
	valMap["C"] = 12
	valMap["D"] = 13
	valMap["E"] = 14
	valMap["F"] = 15
	valMap["G"] = 16
	valMap["H"] = 17
	valMap["J"] = 18
	valMap["K"] = 19
	valMap["L"] = 20
	valMap["M"] = 21
	valMap["N"] = 22
	valMap["P"] = 23
	valMap["Q"] = 24
	valMap["R"] = 25
	valMap["S"] = 26
	valMap["T"] = 27
	valMap["U"] = 28
	valMap["V"] = 29
	valMap["W"] = 30
	valMap["X"] = 31
	valMap["Y"] = 32
	valMap["Z"] = 33
	valMap["?"] = 0
	valMap["*"] = 0
	if val, ok := valMap[code]; ok {
		return val
	}
	return -1
}

/////UNIVERSAL PROFILE

func Profile(ehexSlice ...Ehex) string {
	str := ""
	for _, val := range ehexSlice {
		str += val.Code()
	}
	return str
}
