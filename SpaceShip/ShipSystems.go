package spaceship

import (
	"strconv"
	"strings"
)

type starship struct {
	projectTL         int
	shipClass         string
	shipType          string
	shipDescription   string
	hull              []systemData
	armor             []systemData
	mDrive            []systemData
	jDrive            []systemData
	powerPlant        []systemData
	fuelTanks         []systemData
	bridge            []systemData
	computer          []systemData
	sensors           []systemData
	weapons           []systemData
	systems           []systemData
	craft             []systemData
	staterooms        []systemData
	software          []systemData
	commonAreas       []systemData
	cargo             []systemData
	powerRequirements []powerRequirementData
	crew              []crewData
}

const (
	shipGazelleClassCloseEscort = "GAZELLE" //-Class Close Escort"
)

type systemData struct {
	descr   string
	tonnage float64
	costMCr float64
}

type powerRequirementData struct {
	systemName       string
	powerRequirement int
}

type crewData struct {
	position string
	num      int
}

func Statistics(class string) string {
	data := ""
	ship := starship{}
	switch strings.ToUpper(class) {
	case shipGazelleClassCloseEscort:
		ship.shipClass = "Gazelle"
		ship.shipType = "Close Escort"
		ship.projectTL = 15
		ship.hull = append(ship.hull, systemData{"400 tons, Standard", 0, 20})
		ship.hull = append(ship.hull, systemData{"Reinforced", 0, 10})
		ship.armor = append(ship.armor, systemData{"Crystaliron, Armor 3", 15, 4.5})
		ship.mDrive = append(ship.mDrive, systemData{"Thrust 6", 24, 48})
		ship.jDrive = append(ship.jDrive, systemData{"Jump-5", 55, 82.5})
		ship.powerPlant = append(ship.powerPlant, systemData{"Fussion, Power 540", 36, 36})
		ship.fuelTanks = append(ship.fuelTanks, systemData{"8 weeks of operation, J-3", 128, 0})
		ship.bridge = append(ship.bridge, systemData{"", 10, 2})
		ship.computer = append(ship.computer, systemData{"Computer 30", 0, 20})
		ship.sensors = append(ship.sensors, systemData{"Military Grade", 2, 4.1})
		ship.weapons = append(ship.weapons, systemData{"Barbette (Particle) x 2", 10, 16})
		ship.weapons = append(ship.weapons, systemData{"Triple Turret (Beam Laser) x 2", 2, 5})
		ship.systems = append(ship.systems, systemData{"Drop Tank Mount (80 tons)", 0.32, 0.16})
		ship.systems = append(ship.systems, systemData{"Fuel Processor (120 tons/day)", 6, 0.3})
		ship.systems = append(ship.systems, systemData{"Armory", 1, 0.25})
		ship.systems = append(ship.systems, systemData{"Fuel Scoops", 0, 1})
		ship.craft = append(ship.craft, systemData{"Docking Space (20 tons)", 22, 5.5})
		ship.craft = append(ship.craft, systemData{"Gig", 0, 6.257})
		ship.staterooms = append(ship.staterooms, systemData{"Standard x 11", 44, 5.5})
		ship.software = append(ship.software, systemData{"Evade/1", 0, 1})
		ship.software = append(ship.software, systemData{"Fire Control/4", 0, 8})
		ship.software = append(ship.software, systemData{"Jump Control/5", 0, 0.5})
		ship.software = append(ship.software, systemData{"Library", 0, 0})
		ship.software = append(ship.software, systemData{"Manoeuvre/0", 0, 0})
		ship.commonAreas = append(ship.commonAreas, systemData{"", 11, 11})
		ship.cargo = append(ship.cargo, systemData{"", 33.68, 0})
		ship.powerRequirements = append(ship.powerRequirements, powerRequirementData{"Basic Ship Systems", 80})
		ship.powerRequirements = append(ship.powerRequirements, powerRequirementData{"Manoeuvre Drive", 240})
		ship.powerRequirements = append(ship.powerRequirements, powerRequirementData{"Jump Drive", 200})
		ship.powerRequirements = append(ship.powerRequirements, powerRequirementData{"Sensors", 2})
		ship.crew = append(ship.crew, crewData{"Captain", 1})
		ship.crew = append(ship.crew, crewData{"Pilot", 3})
		ship.crew = append(ship.crew, crewData{"Engineer", 4})
		ship.crew = append(ship.crew, crewData{"Astrogator", 1})
		ship.crew = append(ship.crew, crewData{"Medic", 1})
		ship.crew = append(ship.crew, crewData{"Gunner", 8})
		ship.crew = append(ship.crew, crewData{"Administrator", 1})
		ship.crew = append(ship.crew, crewData{"Maintainance", 1})
		ship.crew = append(ship.crew, crewData{"Officer", 1})
		data = ship.String()
	}
	return data
}

func (s *starship) String() string {
	str := "Class: " + s.shipClass + "\n"
	str += "Type : " + s.shipType + "\n"
	str += "TL   : " + strconv.Itoa(s.projectTL) + "\n"
	str += "--------------------------------------------------------------------------------\n"
	str += printSystemData("Hull", s.hull)
	str += printSystemData("Armor", s.armor)
	str += printSystemData("M-Drive", s.mDrive)
	str += printSystemData("J-Drive", s.jDrive)
	str += printSystemData("Power Plant", s.powerPlant)
	str += printSystemData("Fuel Tanks", s.fuelTanks)
	str += printSystemData("Bridge", s.bridge)
	str += printSystemData("Computer", s.computer)
	str += printSystemData("Sensor", s.sensors)
	str += printSystemData("Weapons", s.weapons)
	str += printSystemData("Systems", s.systems)
	str += printSystemData("Craft", s.craft)
	str += printSystemData("Staterooms", s.staterooms)
	str += printSystemData("Software", s.software)
	str += printSystemData("Common Areas", s.commonAreas)
	str += printSystemData("Cargo", s.cargo)
	str += "--------------------------------------------------------------------------------\n"
	str += "Power Requirements:\n"
	for _, val := range s.powerRequirements {
		str += "  " + val.systemName + " " + strconv.Itoa(val.powerRequirement) + "\n"
	}
	str += "--------------------------------------------------------------------------------\n"
	str += "Crew:\n"
	crewTotal := 0
	for _, val := range s.crew {
		str += "  " + val.position + " x " + strconv.Itoa(val.num) + "\n"
		crewTotal += val.num
	}
	str += "CREW TOTAL: " + strconv.Itoa(crewTotal) + "\n"

	return str
}

func printSystemData(syst string, sd []systemData) string {
	text := ""
	if len(sd) < 1 {
		return text
	}
	for i, val := range sd {
		if i == 0 {
			text = syst
			for len(text) < 12 {
				text = text + " "
			}
			text += " | "
		} else {
			text += "             | "
		}
		descr := val.descr
		for len(descr) < 40 {
			descr += " "
		}
		tons := strconv.FormatFloat(val.tonnage, 'g', 3, 64) + " tons"
		for len(tons) < 10 {
			tons += " "
		}
		if tons == "0 tons    " {
			tons = "     -    "
		}
		text += descr + " | " + tons + " | " + strconv.FormatFloat(val.costMCr, 'g', 3, 64) + " MCr\n"
	}
	return text
}
