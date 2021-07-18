package actions

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
	"github.com/urfave/cli"
)

//SpaceEncounter - основная функция для SpaceEncounter
func SpaceEncounter(c *cli.Context) error {
	dailyLoop := []int{}
	uwp, err := uwp.FromString(c.String("uwp"))
	//fmt.Println(err)
	if err != nil {
		return fmt.Errorf("action: SpaceEncounter:\n  %v", err.Error())
	}
	if c.Int("days") < 0 {
		return fmt.Errorf("action: SpaceEncounter:\n  %v", errors.New("flag '-days' can't be negative"))
	}
	if c.Int("days") == 0 {
		dailyLoop = []int{6}
	}
	if c.Int("days") > 0 {
		for i := 0; i < c.Int("days"); i++ {
			dailyLoop = append(dailyLoop, dice.Roll1D())
		}
	}
	dm := 0
	switch {
	default:
		fmt.Println("Location: Border World")
	case isUntravelled(uwp):
		dm = -4
		fmt.Println("Location: Untravelled Space")
	case isWildSpace(uwp):
		dm = -1
		fmt.Println("Location: Wild Space")
	case hasHighport(uwp):
		fmt.Println("Location: Highport")
		dm = 3
	case hasHighTraffic(uwp):
		fmt.Println("Location: High Traffic System")
		dm = 2
	case isSettled(uwp):
		fmt.Println("Location: Settled Space")
		dm = 1
	}
	encHappened := 0
	for i, v := range dailyLoop {
		if v == 6 {
			enc := callEncounterSpace(dm)
			fmt.Println("Day:", i+1)
			fmt.Println("Encounter:", enc)
			encHappened++
		}
	}
	if encHappened == 0 {
		fmt.Println("No space encounter occured in", len(dailyLoop), "days")
	}
	//fmt.Println("Dm =", dm)

	return nil
}

func newKey(dm int) string {
	r1 := dice.Roll("1d6").DM(dm).Sum()
	r2 := dice.Roll("1d6").Sum()
	if r1 < 0 {
		r1 = 0
	}
	r66 := strconv.Itoa(r1) + strconv.Itoa(r2)
	return r66
}

func hasHighport(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Starport().String() == "A" && uwp.Pops().Value() >= 7:
		return true
	case uwp.Starport().String() == "B" && uwp.Pops().Value() >= 8:
		return true
	case uwp.Starport().String() == "C" && uwp.Pops().Value() >= 9:
		return true
	}
}

func hasHighTraffic(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Pops().Value() >= 9:
		return true
	}
}

func isSettled(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Pops().Value() > 6:
		return true
	}
}

func isWildSpace(uwp *uwp.UWP) bool {
	switch {
	default:
		return false
	case uwp.Govr().Value()+uwp.Laws().Value() >= 20:
		return true
	case uwp.Pops().Value() <= 3:
		return true
	}
}

func isUntravelled(uwp *uwp.UWP) bool {
	strprt := uwp.Starport().String()
	switch {
	default:
		return false
	case strprt == "X", strprt == "Y":
		return true
	case uwp.Pops().Value() == 0 || uwp.Govr().Value() == 0 || uwp.Laws().Value() == 0:
		return true
	}
}

func collision() string {
	return "The ship suffers " + dice.Roll("1d6").SumStr() + " damage"
}

func solarFlare() string {
	return fmt.Sprintf("%v rads", dice.Roll1D()*100)
}

func asteroidMining() string {
	switch dice.Roll2D() {
	default:
		return fmt.Sprintf("error: can't define asteroid yield")
	case 2:
		return fmt.Sprintf("%v tons of Precious Metals", dice.Roll2D())
	case 3, 4:
		return fmt.Sprintf("%v tons of Common Ore", dice.Roll2D()*20)
	case 5, 6:
		return fmt.Sprintf("%v tons of Common Ore", dice.Roll2D()*50)
	case 7, 8:
		return fmt.Sprintf("%v tons of Uncommon Ore", dice.Roll2D()*10)
	case 9, 10, 11:
		return fmt.Sprintf("%v tons of Uncommon Ore", dice.Roll2D()*20)
	case 12:
		return fmt.Sprintf("%v tons of Radioactives", dice.Roll1D())
	}
}

func salvage() string {
	dm := (dice.Roll("1d7").Sum() * -1) + 1
	switch dice.Roll("2d6").DM(dm).Sum() {
	default:
		return "Hazard! The ship’s reactor is damaged, the ship is about to break up, there is a virus loose aboard ship, an alien monster killed the crew, pirate's trap..."
	case 4:
		return "No Salvage: Nothing useful can be recovered."
	case 5:
		return "Junk: Minor personal effects, spare parts, trophies and other junk."
	case 6:
		return fmt.Sprintf("Fuel: %v tons of fuel can be extracted from the salvage (not exceeding the derelict’s maximum capacity)", dice.Roll2D()*10)
	case 7:
		return fmt.Sprintf("Equipment: Items such as vacc suits, medical supplies, weapons, with a total value of %v Cr. ", dice.Roll2D()*2000)
	case 8:
		return fmt.Sprintf("Cargo: %v tons of the derelict’s cargo. Roll D66 on the Trade Goods table (page 212) to determine the type.", dice.Roll2D())
	case 9:
		return fmt.Sprintf("Considerable Cargo: %v tons of cargo (up to the derelict’s maximum cargo).", dice.Roll2D()*10)
	case 10:
		return "Interesting Artefact: An alien relic, useful personal data, mail cannister or other adventure hook – or maybe a survivor in a low berth."
	case 11:
		return fmt.Sprintf("Fittings: Weapons turrets, ship’s computers or air/raft, with a total value of %v Cr", dice.Roll2D()*250000)
	case 12:
		return "Ship: The ship is potentially repairable. "
	}

}

func callEncounterSpace(dm int) string {
	key := newKey(dm)
	encMap := make(map[string]string)
	encMap["01"] = "Alien derelict (" + salvage() + ")"
	encMap["02"] = "Solar flare (" + solarFlare() + ")"
	encMap["03"] = "Asteroid (empty rock)"
	encMap["04"] = "Ore-bearing asteroid (" + asteroidMining() + ")"
	encMap["05"] = "Alien vessel (on a mission)"
	encMap["06"] = "Rock hermit (inhabited rock)"
	encMap["11"] = "Pirate"
	encMap["12"] = "Derelict vessel (" + salvage() + ")"
	encMap["13"] = "Space station " + dice.New().RollFromList([]string{"(derelict)", "(derelict)", "(derelict)", "(derelict)", "", ""})
	encMap["14"] = "Comet (may be ancient derelict at its core)"
	encMap["15"] = "Ore-bearing asteroid (" + asteroidMining() + ")"
	encMap["16"] = fmt.Sprintf("Ship in distress (%v)", rollForShipType(dm))
	encMap["21"] = "Pirate"
	encMap["22"] = "Free trader"
	encMap["23"] = "Micrometeorite Storm (" + collision() + ")"
	encMap["24"] = "Hostile vessel (" + rollForShipType(dm) + ")"
	encMap["25"] = "Mining ship"
	encMap["26"] = "Scout Ship"
	encMap["31"] = "Alien vessel (" + dice.New().RollFromList([]string{"trader", "trader", "trader", "explorer", "explorer", "spy"}) + ")"
	encMap["32"] = "Space junk (" + salvage() + ")"
	encMap["33"] = "Far Trader"
	encMap["34"] = "Derelict (" + salvage() + ")"
	encMap["35"] = dice.New().RollFromList([]string{"Safari ship", "Science vessel"})
	encMap["36"] = "Escape pod"
	encMap["41"] = "Passenger liner"
	encMap["42"] = fmt.Sprintf("Ship in distress (%v)", rollForShipType(dm))
	encMap["43"] = dice.New().RollFromList([]string{"Colony ship", "Passenger liner"})
	encMap["44"] = "Scout ship"
	encMap["45"] = "Space station"
	encMap["46"] = "X-Boat Courier"
	encMap["51"] = fmt.Sprintf("Hostile vessel (%v)", rollForShipType(dm))
	encMap["52"] = "Garbage ejected from a ship"
	encMap["53"] = dice.New().RollFromList([]string{"Medical ship", "Hospital ship"})
	encMap["54"] = dice.New().RollFromList([]string{"Lab ship", "Scout"})
	encMap["55"] = "Patron"
	encMap["56"] = "Police ship"
	encMap["61"] = "Unusually daring pirate"
	encMap["62"] = "Noble yacht"
	encMap["63"] = "Warship"
	encMap["64"] = "Cargo vessel"
	encMap["65"] = dice.New().RollFromList([]string{"Navigational Buoy", "Navigational Beacon"})
	encMap["66"] = "Unusual ship"
	encMap["71"] = "Collision with space junk (" + collision() + ")"
	encMap["72"] = "Automated vessel"
	encMap["73"] = "Free Trader"
	encMap["74"] = "Dumped cargo pod (roll on random trade goods)"
	encMap["75"] = "Police vessel"
	encMap["76"] = "Cargo hauler"
	encMap["81"] = "Passenger liner"
	encMap["82"] = "Orbital factory (roll on random trade goods)"
	encMap["83"] = "Orbital habitat"
	encMap["84"] = "Orbital habitat"
	encMap["85"] = "Communications Satellite"
	encMap["86"] = "Defence Satellite"
	encMap["91"] = "Pleasure craft"
	encMap["92"] = "Space station "
	encMap["93"] = "Police vessel"
	encMap["94"] = "Cargo hauler "
	encMap["95"] = "System Defence Boat"
	encMap["96"] = "Grand Fleet warship"
	return encMap[key]
}

func rollForShipType(dm int) string {
	var key2 string
	var shipType string
	encMap2 := make(map[string]string)
	encMap2["22"] = "Free trader"
	encMap2["25"] = "Mining ship"
	encMap2["26"] = "Scout Ship"
	encMap2["33"] = "Far Trader"
	encMap2["35"] = dice.New().RollFromList([]string{"Safari ship", "Science vessel"})
	encMap2["41"] = "Passenger liner"
	encMap2["43"] = dice.New().RollFromList([]string{"Colony ship", "Passenger liner"})
	encMap2["44"] = "Scout ship"
	encMap2["46"] = "X-Boat Courier"
	encMap2["53"] = dice.New().RollFromList([]string{"Medical ship", "Hospital ship"})
	encMap2["54"] = dice.New().RollFromList([]string{"Lab ship", "Scout"})
	encMap2["62"] = "Yacht"
	encMap2["63"] = "Warship"
	encMap2["64"] = "Cargo vessel"
	encMap2["72"] = "Automated vessel"
	encMap2["73"] = "Free Trader"
	encMap2["76"] = "Cargo hauler"
	encMap2["81"] = "Passenger liner"
	encMap2["91"] = "Pleasure craft"
	encMap2["94"] = "Cargo hauler "
	encMap2["95"] = "System Defence Boat"
	encMap2["96"] = "Grand Fleet warship"
	for shipType == "" {
		key2 = newKey(dm)
		shipType = encMap2[key2]
	}
	return shipType
}
