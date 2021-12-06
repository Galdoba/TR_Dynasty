package maintenance

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

const (
	ALL_SYSTEMS = -1
	DEFECT      = iota
	BREAKDOWN
	FAILURE
	SYSTEM_STRUCTURE_HULL_minor
	SYSTEM_STRUCTURE_HULL_major
	SYSTEM_STRUCTURE_ARMOR ///5
	SYSTEM_STRUCTURE_CARGO
	SYSTEM_ELECTRONICS_Navigation_Systems
	SYSTEM_ELECTRONICS_COMBAT_SENSORS
	SYSTEM_ELECTRONICS_MISSION_SENSORS
	SYSTEM_ELECTRONICS_COMPUTER ///10
	SYSTEM_ELECTRONICS_BRIDGE_COMMAND
	SYSTEM_ELECTRONICS_BRIDGE_ALL
	SYSTEM_POWER_Powerplant_minor
	SYSTEM_POWER_Powerplant_major
	SYSTEM_DRIVES_Jump_Drive_minor /// 15
	SYSTEM_DRIVES_Jump_Drive_major
	SYSTEM_DRIVES_Manoeuvre_Drive_minor
	SYSTEM_DRIVES_Manoeuvre_Drive_major
	SYSTEM_WEAPON_Spinal_Weapon
	SYSTEM_WEAPON_Secondary_Weapon /// 20
	SYSTEM_DEFENSIVE_Defensive_Systems
	SYSTEM_DEFENSIVE_Craft_Bays_or_Drones
	SYSTEM_GENERAL_Life_support
	SYSTEM_GENERAL_Internal_Gravity
	SYSTEM_GENERAL_Small_Craft /// 25
	SYSTEM_GENERAL_Fuel_Processor
	SYSTEM_GENERAL_Mission_System //laboratory, observatory or similar
	SYSTEM_GENERAL_Systems_special
	SYSTEM_DISTRIBUTION_ERROR
	SCALE_MINOR
	SCALE_MAJOR
	SCALE_STRUCTURAL
)

type damageControlManager struct {
	problemWith map[int]problem //dcm.problemWith[SYSTEM_GENERAL_Life_Support]
	shipTonnage int
}

type DamageControlManager interface {
	Assign(...int) error
	RepairCost() (int, int)
}

func NewDamageControlManager(tonnage int) *damageControlManager {
	dcm := damageControlManager{}
	dcm.problemWith = make(map[int]problem)
	dcm.shipTonnage = tonnage
	return &dcm
}

type problem struct {
	defects    int //defect/breakdown/failure
	breakdowns int
	failure    int
	system     int
	scale      int
	repairCost []int
}

//Assign - распределяет и скалирует проблемы с содержанием систем корабля
func (dcm *damageControlManager) Assign(issues ...int) error {
	for _, issue := range issues {
		system := randomSystem()
		if system == SYSTEM_DISTRIBUTION_ERROR {
			return fmt.Errorf("system distribution failed")
		}
		problem := dcm.problemWith[system]
		if problem.isVoid() {
			problem = newProblem(system)
		}

		if err := problem.evaluate(issue); err != nil {
			return fmt.Errorf("problem (%v) evaluate issue (%v): %v", problem, issue, err.Error())
		}
		problem.rollRepairCost(dcm.shipTonnage)
		dcm.problemWith[system] = problem

	}
	return nil
}

func structuralList() []int {
	return []int{
		SYSTEM_STRUCTURE_HULL_minor,
		SYSTEM_STRUCTURE_HULL_major,
		SYSTEM_STRUCTURE_ARMOR,
		SYSTEM_STRUCTURE_CARGO,
	}

}

func majorSystemList() []int {
	return []int{
		SYSTEM_POWER_Powerplant_minor,
		SYSTEM_POWER_Powerplant_major,
		SYSTEM_DRIVES_Jump_Drive_minor,
		SYSTEM_DRIVES_Jump_Drive_major,
		SYSTEM_DRIVES_Manoeuvre_Drive_minor,
		SYSTEM_DRIVES_Manoeuvre_Drive_major,
		SYSTEM_WEAPON_Spinal_Weapon,
		SYSTEM_WEAPON_Secondary_Weapon,
		SYSTEM_DEFENSIVE_Defensive_Systems,
		SYSTEM_DEFENSIVE_Craft_Bays_or_Drones,
	}
}

func minorSystemList() []int {
	return []int{
		SYSTEM_ELECTRONICS_Navigation_Systems,
		SYSTEM_ELECTRONICS_COMBAT_SENSORS,
		SYSTEM_ELECTRONICS_MISSION_SENSORS,
		SYSTEM_ELECTRONICS_COMPUTER,
		SYSTEM_ELECTRONICS_BRIDGE_COMMAND,
		SYSTEM_ELECTRONICS_BRIDGE_ALL,
		SYSTEM_GENERAL_Life_support,
		SYSTEM_GENERAL_Internal_Gravity,
		SYSTEM_GENERAL_Small_Craft,
		SYSTEM_GENERAL_Fuel_Processor,
		SYSTEM_GENERAL_Mission_System,
		SYSTEM_GENERAL_Systems_special,
	}
}

func codeMatchScale(code int) int {
	for _, val := range minorSystemList() {
		if val == code {
			return SCALE_MINOR
		}
	}
	for _, val := range majorSystemList() {
		if val == code {
			return SCALE_MAJOR
		}
	}
	for _, val := range structuralList() {
		if val == code {
			return SCALE_STRUCTURAL
		}
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func newProblem(systemCode int) problem {
	pr := problem{
		system:     systemCode,
		defects:    0,
		breakdowns: 0,
		failure:    0,
	}
	pr.scale = codeMatchScale(systemCode)
	return pr
}

func (pr problem) isVoid() bool {
	return pr.breakdowns+pr.defects+pr.system+pr.failure == 0
}

func (pr *problem) evaluate(issue int) error {
	if pr.system == DEFECT || pr.system == BREAKDOWN || pr.system == FAILURE {
		return fmt.Errorf("system code invalid")
	}
	switch issue {
	case DEFECT:
		if pr.defects > 6 {
			pr.breakdowns++
			return nil
		}
		pr.defects++
		return nil
	case BREAKDOWN:
		pr.breakdowns++
		return nil
	case FAILURE:
		pr.failure++
		return nil
	}
	return fmt.Errorf("evaluate func unexpected error")
}

func randomSystem() int {
	switch dice.Roll1D() {
	case 1, 2:
		return rollStructure()
	case 3:
		return rollSensorsAndElectronics()
	case 4:
		return rollDrivesAndPower()
	case 5:
		return rollWeaponAndDefence()
	case 6:
		return rollGeneral()
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func rollStructure() int {
	switch dice.Roll1D() {
	case 1, 2, 3:
		return SYSTEM_STRUCTURE_HULL_minor
	case 4:
		return SYSTEM_STRUCTURE_HULL_major
	case 5:
		return SYSTEM_STRUCTURE_ARMOR
	case 6:
		return SYSTEM_STRUCTURE_CARGO
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func rollSensorsAndElectronics() int {
	switch dice.Roll1D() {
	case 1:
		return SYSTEM_ELECTRONICS_Navigation_Systems
	case 2:
		return SYSTEM_ELECTRONICS_COMBAT_SENSORS
	case 3:
		return SYSTEM_ELECTRONICS_MISSION_SENSORS
	case 4:
		return SYSTEM_ELECTRONICS_COMPUTER
	case 5:
		return SYSTEM_ELECTRONICS_BRIDGE_COMMAND
	case 6:
		return SYSTEM_ELECTRONICS_BRIDGE_ALL
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func rollDrivesAndPower() int {
	switch dice.Roll1D() {
	case 1:
		return SYSTEM_DRIVES_Jump_Drive_minor
	case 2:
		return SYSTEM_DRIVES_Jump_Drive_major
	case 3:
		return SYSTEM_POWER_Powerplant_minor
	case 4:
		return SYSTEM_POWER_Powerplant_major
	case 5:
		return SYSTEM_DRIVES_Manoeuvre_Drive_minor
	case 6:
		return SYSTEM_DRIVES_Manoeuvre_Drive_major
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func rollWeaponAndDefence() int {
	switch dice.Roll1D() {
	case 1:
		return SYSTEM_WEAPON_Spinal_Weapon
	case 2, 3:
		return SYSTEM_WEAPON_Secondary_Weapon
	case 4:
		return SYSTEM_DEFENSIVE_Defensive_Systems
	case 5:
		return SYSTEM_ELECTRONICS_COMBAT_SENSORS
	case 6:
		return SYSTEM_DEFENSIVE_Craft_Bays_or_Drones
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func rollGeneral() int {
	switch dice.Roll1D() {
	case 1:
		return SYSTEM_GENERAL_Life_support
	case 2:
		return SYSTEM_GENERAL_Internal_Gravity
	case 3:
		return SYSTEM_GENERAL_Small_Craft
	case 4:
		return SYSTEM_GENERAL_Fuel_Processor
	case 5:
		return SYSTEM_GENERAL_Mission_System
	case 6:
		return SYSTEM_GENERAL_Systems_special
	}
	return SYSTEM_DISTRIBUTION_ERROR
}

func (pr problem) String() string {
	issue := ""
	if pr.defects > 0 {
		issue += "DEFECT x " + fmt.Sprintf("%v ", pr.defects)
	}
	if pr.breakdowns > 0 {
		issue += "BREAKDOWNS x " + fmt.Sprintf("%v", pr.breakdowns)
	}
	if pr.failure > 0 {
		issue = "System Failed"
	}
	rep := fmt.Sprintf("Repair Cost: SU=%v/RM=%v/EM=%v/RT=%v", pr.repairCost[0], pr.repairCost[1], pr.repairCost[2], pr.repairCost[3])
	return fmt.Sprintf("[Problem {Code: %v = %v; %v}]", pr.system, issue, rep)
}

func DetermineIssues(mods ...int) []int {
	dm := 0
	for _, v := range mods {
		dm += v
	}
	r := dice.Roll2D(dm)
	var issues []int
	switch r {
	default:
		if r < 1 {
			return []int{0, 0, 0}
		}
		issues = append(issues, []int{0, 0, 9}...)
	case 1, 2, 3:
		issues = append(issues, []int{0, 0, 0}...)
	case 4, 5, 6:
		issues = append(issues, []int{1, 0, 0}...)
	case 7, 8, 9:
		issues = append(issues, []int{2, 0, 0}...)
	case 10, 11, 12:
		issues = append(issues, []int{3, 0, 0}...)
	case 13, 14, 15:
		issues = append(issues, []int{1, 1, 0}...)
	case 16, 17, 18:
		issues = append(issues, []int{2, 1, 0}...)
	case 19, 20, 21:
		issues = append(issues, []int{3, 1, 0}...)
	case 22, 23, 24:
		issues = append(issues, []int{1, 2, 1}...)
	case 25, 26, 27:
		issues = append(issues, []int{2, 2, 1}...)
	case 28, 29, 30:
		issues = append(issues, []int{3, 2, 1}...)
	case 31, 32, 33:
		issues = append(issues, []int{1, 3, 2}...)
	case 34, 35, 36:
		issues = append(issues, []int{2, 3, 2}...)
	case 37, 38, 39:
		issues = append(issues, []int{9, 3, 2}...)
	case 40, 41, 42:
		issues = append(issues, []int{9, 9, 2}...)
	}
	res := []int{}
	for i, v := range issues {
		switch i {
		case 0:
			for d := 0; d < v; d++ {
				res = append(res, DEFECT)
			}
		case 1:
			for b := 0; b < v; b++ {
				res = append(res, BREAKDOWN)
			}
		case 2:
			for f := 0; f < v; f++ {
				res = append(res, FAILURE)
			}
		}
	}
	//s := "SU	Defect (minor)    : 2d * тоннаж/750 	Defect (major)    : 2d * тоннаж/150 	Breakdown (minor) : 5d * тоннаж/750 	Breakdown (major) : 5d * тоннаж/150 	Failure (minor)   : 8d * тоннаж/750 	Failure (major)   : 8d * тоннаж/150 	1 armor           : 2d * тоннаж/15 	1 HP              : 2d * 7 	hours 	из книги"
	//fmt.Print(s)
	return res
}

func (pr *problem) rollRepairCost(tons int) {
	su := 0
	rm := 0
	em := 0
	rt := 0
	hpLim := tons * 40 / 100
	v750 := utils.Max(tons/750, 1)
	v150 := utils.Max(tons/150, 1)
	v15 := utils.Max(tons/15, 1)
	for i := 0; i < pr.defects; i++ {
		switch pr.scale {
		case SCALE_MINOR:
			su += dice.Roll2D() * v750
			rt = dice.Roll2D() * v750
		case SCALE_MAJOR:
			su += dice.Roll2D() * v150
			rt = dice.Roll2D() * v150
		}
	}
	for i := 0; i < pr.breakdowns; i++ {
		switch pr.scale {
		case SCALE_MINOR:
			su += dice.Roll5D() * v750
			rt = dice.New().RollNext("6d6").Sum() * v750
			switch dice.Roll1D() {
			case 1, 2, 3:
				rm += dice.Roll3D()
			case 4, 5, 6:
				em += dice.Roll1D()
			}
		case SCALE_MAJOR:
			su += dice.Roll5D() * v150
			rt = dice.New().RollNext("20d6").Sum() * v150
			switch dice.Roll1D() {
			case 1, 2, 3:
				rm += dice.New().RollNext("6d6").Sum()
			case 4, 5, 6:
				em += dice.Roll3D()
			}
		}
	}
	for i := 0; i < pr.failure; i++ {
		switch pr.scale {
		case SCALE_MINOR:
			su += dice.New().RollNext("8d6").Sum() * v750
			switch dice.Roll1D() {
			case 1, 2, 3:
				rm += dice.Roll4D()
			case 4, 5, 6:
				em += dice.Roll2D()
			}
		case SCALE_MAJOR:
			su += dice.New().RollNext("8d6").Sum() * v150
			switch dice.Roll1D() {
			case 1, 2, 3:
				rm += dice.New().RollNext("8d6").Sum()
			case 4, 5, 6:
				em += dice.Roll4D()
			}
		}
	}
	switch pr.system {
	case SYSTEM_STRUCTURE_ARMOR, SYSTEM_STRUCTURE_CARGO:
		d := pr.defects
		b := dice.New().RollNext(strconv.Itoa(pr.breakdowns) + "d3").Sum()
		f := dice.New().RollNext(strconv.Itoa(pr.failure) + "d6").Sum()
		val := d + b + f
		su += val * dice.Roll2D() * v15
		switch dice.Roll1D() {
		case 1, 2, 3:
			rm += dice.Roll2D() * v750 * val
		case 4, 5, 6:
			em += dice.Roll1D() * v750 * val
		}
		rt = dice.Roll2D() * 50
	case SYSTEM_STRUCTURE_HULL_minor, SYSTEM_STRUCTURE_HULL_major:
		hp := 0
		for m := 0; m < pr.defects; m++ {
			hp += utils.Max(hpLim/100, 1)
		}
		for m := 0; m < pr.breakdowns; m++ {
			hp += dice.New().RollNext("1d3").Sum() * hpLim / 100
		}
		for m := 0; m < pr.failure; m++ {
			hp += dice.Roll1D() * hpLim / 100
		}
		if pr.system == SYSTEM_STRUCTURE_HULL_major {
			hp = hp * 10
		}
		for j := 0; j < hp; j++ {
			su += (dice.Roll2D() * 7)
		}
		rt = dice.Roll2D() * hp
	}
	pr.repairCost = append([]int{}, su, rm, em, rt)
}

/*
SU
Defect (minor)    : 2d * тоннаж/750
Defect (major)    : 2d * тоннаж/150
Breakdown (minor) : 5d * тоннаж/750
Breakdown (major) : 5d * тоннаж/150
Failure (minor)   : 8d * тоннаж/750
Failure (major)   : 8d * тоннаж/150
1 armor           : 2d * тоннаж/15
1 HP              : 2d * 7
hours
из книги
*/
