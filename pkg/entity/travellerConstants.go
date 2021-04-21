package entity

const (
	STR = "STR"
	DEX = "DEX"
	END = "END"
	INT = "INT"
	EDU = "EDU"
	SOC = "SOC"
)

func listCharacteristics() []string {
	return []string{
		STR,
		DEX,
		END,
		INT,
		EDU,
		SOC,
	}
}
