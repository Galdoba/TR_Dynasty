package encounters

func EncounterMapCRB(key string) string {
	tag := ""
	switch key {
	case "01":
		tag = "Alien derelict (possible salvage) "
	case "02":
		tag = "Solar flare Solar ﬂare (1d6  100 rads) "
	case "03":
		tag = "Asteroid (empty rock) "
	case "04":
		tag = "Ore-bearing asteroid (possible mining) "
	case "05":
		tag = "Alien vessel (on a mission) "
	case "06":
		tag = "Rock hermit (inhabited rock) "
	case "11":
		tag = "Pirate Pirate "
	case "12":
		tag = "Derelict vessel (possible salvage) "
	case "13":
		tag = "Space station (1–4: derelict) "
	case "14":
		tag = "Comet (may be ancient derelict at its core) "
	case "15":
		tag = "Ore-bearing asteroid (possible mining) "
	case "16":
		tag = "Ship in distress "
	case "21":
		tag = "Pirate "
	case "22":
		tag = "Free trader "
	case "23":
		tag = "Micrometeorite storm Micrometeorite storm (collision!) "
	case "24":
		tag = "Hostile vessel Hostile vessel (roll again for type) "
	case "25":
		tag = "Mining ship "
	case "26":
		tag = "Scout ship "
	case "31":
		tag = "Alien vessel (1–3: trader, 4–6: explorer, 6: spy) "
	case "32":
		tag = "Space junk (possible salvage) "
	case "33":
		tag = "Far trader "
	case "34":
		tag = "Derelict (possible salvage) "
	case "35":
		tag = "Safari or science vessel "
	case "36":
		tag = "Escape pod "
	case "41":
		tag = "Passenger liner "
	case "42":
		tag = "Ship in distress "
	case "43":
		tag = "Colony ship or passenger liner "
	case "44":
		tag = "Scout ship "
	case "45":
		tag = "Space station "
	case "46":
		tag = "X-boat courier "
	case "51":
		tag = "Hostile vesselHostile vessel (roll again for type) "
	case "52":
		tag = "Garbage ejected from a ship "
	case "53":
		tag = "Medical ship or hospital "
	case "54":
		tag = "Lab ship or scout "
	case "55":
		tag = "Patron Patron (roll on the patron table, page 81) "
	case "56":
		tag = "Police ship "
	case "61":
		tag = "Unusually daring pirate "
	case "62":
		tag = "Noble yacht "
	case "63":
		tag = "Warship "
	case "64":
		tag = "Cargo vessel "
	case "65":
		tag = "Navigational buoy or beacon "
	case "66":
		tag = "Unusual ship "
	case "71":
		tag = "Collision with space junk Collision with space junk (collision!) "
	case "72":
		tag = "Automated vessel "
	case "73":
		tag = "Free trader "
	case "74":
		tag = "Dumped cargo pod (roll on random trade goods) "
	case "75":
		tag = "Police vessel "
	case "76":
		tag = "Cargo hauler "
	case "81":
		tag = "Passenger liner "
	case "82":
		tag = "Orbital factory (roll on random trade goods) "
	case "83":
		tag = "Orbital habitat "
	case "84":
		tag = "Orbital habitat "
	case "85":
		tag = "Communications satellite "
	case "86":
		tag = "Defence satellite "
	case "91":
		tag = "Pleasure craft "
	case "92":
		tag = "Space station "
	case "93":
		tag = "Police vessel "
	case "94":
		tag = "Cargo hauler "
	case "95":
		tag = "System defence boat "
	case "96":
		tag = "Grand ﬂeet warship"
	}
	return tag
}
