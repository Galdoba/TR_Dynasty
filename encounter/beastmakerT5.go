package encounter

func table00A(flux int) (terrain string) {
	switch flux {
	default:
		terrain = "Unknown"
	case -5:
		terrain = "Mountain"
	case -4:
		terrain = "Desert"
	case -3:
		terrain = "Exotic"
	case -2:
		terrain = "Rough Wood"
	case -1:
		terrain = "Rough"
	case 0:
		terrain = "Clear"
	case 1:
		terrain = "Forest"
	case 2:
		terrain = "Wetlands"
	case 3:
		terrain = "Wetlands Woods"
	case 4:
		terrain = "Ocean"
	case 5:
		terrain = "Ocean Depths"
	}
	return terrain
}

func table00B(flux int, roll int, planetUWP string) {

}
