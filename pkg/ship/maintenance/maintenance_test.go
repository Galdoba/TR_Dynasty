package maintenance

import (
	"fmt"
	"testing"
)

func TestReporter(t *testing.T) {
	screen()
	dcm := NewDamageControlManager(75000)
	fmt.Printf("%v\n", dcm)
	if dcm == nil {
		t.Errorf("dcm cannot ne NIL")
	}
	if err := dcm.Assign(DetermineIssues(2)...); err != nil {
		t.Errorf("Assign error: %v\n", err.Error())
	} else {
		fmt.Printf("defect assigned...\n")
	}

	fmt.Println(dcm)
	for _, v := range dcm.problemWith {
		fmt.Println(v)
	}
}

func screen() {
	fmt.Println("DEFECT =", DEFECT)
	fmt.Println("BREAKDOWN =", BREAKDOWN)
	fmt.Println("FAILURE =", FAILURE)
	fmt.Println("SYSTEM_STRUCTURE_HULL_minor =", SYSTEM_STRUCTURE_HULL_minor)
	fmt.Println("SYSTEM_STRUCTURE_HULL_major =", SYSTEM_STRUCTURE_HULL_major)
	fmt.Println("SYSTEM_STRUCTURE_ARMOR =", SYSTEM_STRUCTURE_ARMOR)
	fmt.Println("SYSTEM_STRUCTURE_CARGO =", SYSTEM_STRUCTURE_CARGO)
	fmt.Println("SYSTEM_ELECTRONICS_Navigation_Systems =", SYSTEM_ELECTRONICS_Navigation_Systems)
	fmt.Println("SYSTEM_ELECTRONICS_COMBAT_SENSORS =", SYSTEM_ELECTRONICS_COMBAT_SENSORS)
	fmt.Println("SYSTEM_ELECTRONICS_MISSION_SENSORS =", SYSTEM_ELECTRONICS_MISSION_SENSORS)
	fmt.Println("SYSTEM_ELECTRONICS_COMPUTER =", SYSTEM_ELECTRONICS_COMPUTER)
	fmt.Println("SYSTEM_ELECTRONICS_BRIDGE_COMMAND =", SYSTEM_ELECTRONICS_BRIDGE_COMMAND)
	fmt.Println("SYSTEM_ELECTRONICS_BRIDGE_ALL =", SYSTEM_ELECTRONICS_BRIDGE_ALL)
	fmt.Println("SYSTEM_POWER_Powerplant_minor =", SYSTEM_POWER_Powerplant_minor)
	fmt.Println("SYSTEM_POWER_Powerplant_major =", SYSTEM_POWER_Powerplant_major)
	fmt.Println("SYSTEM_DRIVES_Jump_Drive_minor =", SYSTEM_DRIVES_Jump_Drive_minor)
	fmt.Println("SYSTEM_DRIVES_Jump_Drive_major =", SYSTEM_DRIVES_Jump_Drive_major)
	fmt.Println("SYSTEM_DRIVES_Manoeuvre_Drive_minor =", SYSTEM_DRIVES_Manoeuvre_Drive_minor)
	fmt.Println("SYSTEM_DRIVES_Manoeuvre_Drive_major =", SYSTEM_DRIVES_Manoeuvre_Drive_major)
	fmt.Println("SYSTEM_WEAPON_Spinal_Weapon =", SYSTEM_WEAPON_Spinal_Weapon)
	fmt.Println("SYSTEM_WEAPON_Secondary_Weapon =", SYSTEM_WEAPON_Secondary_Weapon)
	fmt.Println("SYSTEM_DEFENSIVE_Defensive_Systems =", SYSTEM_DEFENSIVE_Defensive_Systems)
	fmt.Println("SYSTEM_DEFENSIVE_Craft_Bays_or_Drones =", SYSTEM_DEFENSIVE_Craft_Bays_or_Drones)
	fmt.Println("SYSTEM_GENERAL_Life_support =", SYSTEM_GENERAL_Life_support)
	fmt.Println("SYSTEM_GENERAL_Internal_Gravity =", SYSTEM_GENERAL_Internal_Gravity)
	fmt.Println("SYSTEM_GENERAL_Small_Craft =", SYSTEM_GENERAL_Small_Craft)
	fmt.Println("SYSTEM_GENERAL_Fuel_Processor =", SYSTEM_GENERAL_Fuel_Processor)
	fmt.Println("SYSTEM_GENERAL_Mission_System =", SYSTEM_GENERAL_Mission_System)
	fmt.Println("SYSTEM_GENERAL_Systems_special =", SYSTEM_GENERAL_Systems_special)
	fmt.Println("SYSTEM_DISTRIBUTION_ERROR =", SYSTEM_DISTRIBUTION_ERROR)
	fmt.Println("SCALE_MINOR =", SCALE_MINOR)
	fmt.Println("SCALE_MAJOR =", SCALE_MAJOR)
	fmt.Println("SCALE_STRUCTURAL =", SCALE_STRUCTURAL)
}
