package main

func (d *dynasty) rollEvent(eventCode string) {
	//return success bool, duration int
}

func Survived(d dynasty) bool {
	for _, val := range listValues() {
		if d.values[val] < 1 {
			//fmt.Println("Dynasty have crumbled and is no more...")
			return false
		}
	}
	zeroTraits := []string{}
	for _, val := range listTraits() {
		if d.traits[val] < 1 {
			zeroTraits = append(zeroTraits, val)
		}
	}
	if len(zeroTraits) > 1 {
		//fmt.Println("Dynasty is too weak to defend itself from the normal dangers it would face and is swiftly torn asunder by rivals...")
		return false
	}
	vitalChars := []string{Lty, Pop, Tra}
	for _, val := range vitalChars {
		if d.characteristics[val] < 1 {
			//	fmt.Println("Dynasty members riot and rise up from within, destroying the Dynasty’s power base until it cannot stand on its own...")
			return false
		}
	}
	return true
}
