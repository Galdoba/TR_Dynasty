package entity

const (
	STR          = "STR"
	DEX          = "DEX"
	END          = "END"
	INT          = "INT"
	EDU          = "EDU"
	SOC          = "SOC"
	PSI          = "PSI"
	SpeciesHuman = "Human"
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

func listSpeciesTraits(species string) []string {
	traits := []string{}
	switch species {
	default:
		traits = append(traits, "NONE")
	case "Aslan":
		traits = append(traits, "Death Claw")
	case "Vargr":
		traits = append(traits, "Bite")
		traits = append(traits, "Heightened Senses")
	}
	return traits
}
