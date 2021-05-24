package weapons

import "errors"

const (
	MeleeWeapons = iota
	BodyWeapons
	RangeWeapons
	WeaponsDiscriptors
	WeaponsBurden
	WeaponsStage
	WeaponsUsers
	WeaponsPortability
)

func DataFrom(table, row, col int) (string, error) {
	switch table {
	default:
		return "", errors.New("unknown table adressed")
	case MeleeWeapons,
		BodyWeapons,
		RangeWeapons,
		WeaponsDiscriptors,
		WeaponsBurden,
		WeaponsStage,
		WeaponsUsers,
		WeaponsPortability:

	}
	return "0"
}

func meleeWeapons() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "TL", "Range", "Mass", "qreBs", "H1", "D1", "H2", "D2", "Hits(v1)", "Cr Base"},
		{"Short Blades", "K", "Knife", "1", "R", "0.5", "", "Cuts", "2", "", "", "2", "50"},
		{"Short Blades", "D", "Dagger", "2", "R", "0.5", "", "Cuts", "2", "", "", "2", "50"},
		{"Short Blades", "TK", "Trench Knife", "4", "R", "1", "", "Cuts", "2", "Blow", "1", "2", "100"},
		{"Short Blades", "BK", "Big Knife", "5", "T", "3", "", "Cuts", "2", "Pen", "2", "2", "200"},
		{"Short Blades", "GBK", "Great Big Knife", "6", "1", "6", "", "Cuts", "2", "Pen", "2", "2", "900"},
		{"Medium Blades", "S", "Sword", "3", "1", "2", "", "Cuts", "2", "", "", "2", "300"},
		{"Medium Blades", "sS", "Short Sword", "3", "1", "1", "B= -1", "Cuts", "2", "", "", "2", "300"},
		{"Medium Blades", "bS", "Broadsword", "4", "1", "3", "", "Cuts", "3", "", "", "3", "700"},
		{"Medium Blades", "C", "Cutlass", "3", "1", "2", "", "Cuts", "2", "", "", "2", "200"},
		{"Medium Blades", "OC", "Officers Cutlass", "5", "1", "1", "", "Cuts", "2", "", "", "2", "400"},
		{"Long Blades", "P", "Spear/Pike", "1", "1", "2", "B= +3", "Cuts", "2", "", "", "2", "50"},
		{"Special Blades", "Ax", "Axe", "2", "T", "2", "", "Cuts", "3", "", "", "3", "60"},
		{"Special Blades", "A", "Space Axe", "9", "1", "2", "", "Cuts", "2", "Pen", "2", "2", "500"},
		{"Special Blades", "V", "Vibro-Blade", "10", "1", "0.5", "", "Cuts", "2", "", "", "2", "900"},
		{"Special Blades", "m", "Mace", "2", "1", "4", "", "Cuts", "1", "Blow", "2", "2", "100"},
		{"Special Blades", "c", "Club", "1", "1", "2", "", "Blow", "=C1", "", "", "1", "10"},
	}
	return data
}

func bodyWeapons() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "Range", "H1", "D1", "H2", "D2", "Hits(v1)"},
		{"Body Weapons", "Fi", "Fists", "R", "Blow", "=C1", "", "", "1"},
		{"Body Weapons", "Te", "Tentacle", "0", "Hit", "=C1", "Suff", "1", "1"},
		{"Body Weapons", "Ho", "Horns", "R", "Pen", "=C1", "", "", "2"},
		{"Body Weapons", "Tu", "Tusks", "R", "Pen", "=C1", "", "", "2"},
		{"Body Weapons", "Fa", "Fangs", "R", "Pen", "=C1", "", "", "2"},
		{"Body Weapons", "T", "Teeth", "R", "Cuts", "=C1", "", "", "1"},
		{"Body Weapons", "Cl", "Claws", "R", "Cuts", "=C1", "", "", "1"},
		{"Body Weapons", "Dc", "Dew Claw", "R", "Cuts", "=C1", "", "", "2"},
		{"Body Weapons", "H", "Hooves", "R", "Blow", "=C1", "", "", "2"},
		{"Body Weapons", "Sp", "Spikes", "0", "Pen", "=C1", "", "", "2"},
		{"Body Weapons", "St", "Sting", "R", "Pen", "=C1", "Poison", "2D", "3"},
	}
	return data
}

func rangeWeapons() [][]string {
	data := [][]string{
		{"Category", "Code", "Type", "TL", "Range", "Mass", "qreBs", "H1", "D1", "Misc", "Hits(v1)", "Cr Base"},
		{"Artillery", "G", "Gun", "6", "4", "9", "+1", "*", "2", "", "2", "5000"},
		{"Artillery", "Ga", "Gatling", "7", "4", "40", "+2", "*", "3", "", "2", "8000"},
		{"Artillery", "C", "Cannon", "6", "6", "200", "+4", "*", "4", "", "2", "10000"},
		{"Artillery", "aC", "Autocannon", "8", "6", "300", "+4", "*", "5", "", "3", "30000"},
		{"Long Guns", "R", "Rifle", "5", "5", "4", "0", "Bullet", "2", "Not Bullet if Laser", "2", "500"},
		{"Long Guns", "C", "Carbine", "5", "4", "3", "-1", "Bullet", "1", "Not Bullet if Laser", "1", "400"},
		{"Handguns", "P", "Pistol", "5", "2", "1.1", "0", "Bullet", "1", "Not Bullet if Laser", "1", "150"},
		{"Handguns", "R", "Revolver", "4", "2", "1.25", "0", "Bullet", "1", "Not Bullet if Laser", "1", "100"},
		{"Shotguns", "S", "Shotgun", "4", "2", "4", "0", "Frag", "2", "", "2", "300"},
		{"Machineguns", "Mg", "Machinegun", "6", "5", "8", "+1", "Bullet", "4", "", "4", "3000"},
		{"Projectors", "Pj", "Projector", "9", "0", "1", "0", "*", "1", "", "1", "300"},
		{"Designators", "D", "Designator", "7", "5", "10", "+1", "*", "1", "", "1", "2000"},
		{"Launchers", "L", "Launcher", "6", "3", "10", "+1", "*", "1", "", "0", "1000"},
		{"Launchers", "mL", "Multi-Launcher", "8", "5", "8", "+1", "*", "1", "", "0", "3000"},
	}
	return data
}

func weaponsDiscriptors() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "TL", "R=", "Mass", "qreBs", "H2", "D2", "H3", "D3", "Hits(v1)", "Cr"},
		{"Artillery", "aF", "Anti-Flyer", "+4", "=6", "x6.0", "", "Frag", "1", "Blast", "3", "4", "x3.0"},
		{"Artillery", "aT", "Anti-Tank", "", "=5", "x8.0", "", "Pen", "3", "Blast", "3", "6", "x2.0"},
		{"Artillery", "A", "Assault", "+2", "=4", "x0.8", "", "Bang", "1", "Blast", "2", "3", "x1.5"},
		{"Artillery", "F", "Fusion", "+7", "=4", "x2.3", "", "Pen", "4", "Burn", "4", "8", "x6.0"},
		{"Artillery", "G", "Gauss", "+7", "=4", "x0.9", "", "Bullet", "3", "", "", "3", "x2.0"},
		{"Artillery", "P", "Plasma", "+5", "=4", "x2.5", "", "Pen", "3", "Burn", "3", "6", "x2.0"},
		{"Long Guns", "", "(blank)", "", "", "", "", "", "", "", "", "", "x1.0"},
		{"Long Guns", "Ac", "Accelerator", "+4", "", "x0.6", "", "Bullet", "2", "", "", "2", "x3.0"},
		{"Long Guns", "A", "Assault", "+2", "=4", "x0.8", "", "Blast", "2", "Bang", "1", "3", "x1.5"},
		{"Long Guns", "B", "Battle", "+1", "=5", "x1.0", "+1", "Bullet", "1", "", "", "1", "x0.8"},
		{"Long Guns", "C", "Combat", "+2", "=3", "x0.9", "", "Frag", "2", "", "", "2", "x1.5"},
		{"Long Guns", "D", "Dart", "+1", "=4", "x0.6", "", "Tranq", "1-2-3", "", "", "1-2-3", "x0.9"},
		{"Long Guns", "P", "Poison Dart", "+1", "=4", "x1.0", "", "Poison", "1-2-3", "", "", "1-2-3", "x0.9"},
		{"Long Guns", "G", "Gauss", "+7", "", "x0.9", "", "Bullet", "3", "", "", "3", "x2.0"},
		{"Long Guns", "H", "Hunting", "", "=3", "x0.9", "-1", "Bullet", "1", "", "", "1", "x1.2"},
		{"Long Guns", "L", "Laser", "+5", "", "x1.2", "", "Burn", "2", "Pen", "2", "4", "x6.0"},
		{"Long Guns", "Sp", "Splat", "+2", "=4", "x1.3", "+1", "Bullet", "1", "", "", "1", "x2.4"},
		{"Long Guns", "S", "Survival", "", "=2", "x0.5", "", "Bullet", "1", "", "", "1", "x1.2"},
		{"Handguns", "", "(blank)", "", "", "", "", "", "", "", "", "", "x1.0"},
		{"Handguns", "Ac", "Accelerator", "+4", "", "x0.6", "", "Bullet", "2", "", "", "2", "x3.0"},
		{"Handguns", "L", "Laser", "+5", "", "x1.2", "", "Burn", "2", "Pen", "2", "4", "x2.0"},
		{"Handguns", "M", "Machine", "", "=2", "x1.2", "", "Bullet", "2", "", "", "2", "x1.5"},
		{"Shotguns", "", "(blank)", "", "", "", "", "", "", "", "", "", "x1.0"},
		{"Shotguns", "A", "Assault", "+2", "=4", "x0.8", "", "Bang", "1", "Blast", "2", "3", "x2.0"},
		{"Shotguns", "H", "Hunting", "", "=3", "x0.9", "", "Bullet", "1", "", "", "1", "x1.2"},
		{"Machineguns", "", "(blank)", "", "", "", "", "", "", "", "", "", "x1.0"},
		{"Machineguns", "aF", "Anti-Flyer", "+4", "=6", "x6.0", "", "Frag", "1", "Blast", "3", "4", "x3.0"},
		{"Machineguns", "A", "Assault", "+2", "=4", "x0.8", "", "Bang", "1", "Blast", "2", "3", "x1.5"},
		{"Machineguns", "S", "Sub", "+1", "=3", "x0.3", "", "Bullet", "-1", "", "", "-1", "x0.9"},
		{"Spray", "A", "Acid", "", "=3", "x1.0", "+1", "Corrode", "2", "Pen", "1-2-3", "4", "x3.0"},
		{"Spray", "H", "Fire", "", "=1", "x0.9", "", "Burn", "1-2-3", "Pen", "1-2-3", "2-4-6", "x2.0"},
		{"Spray", "P", "Poison Gas", "", "=2", "x1.0", "", "Gas", "1-2-3", "Poison", "1-2-3", "2-4-6", "x3.0"},
		{"Spray", "S", "Stench", "+3", "=2", "x0.4", "", "Stench", "1-2-3", "", "", "1-2-3", "x1.2"},
		{"Launchers", "aF", "AF Missile", "+4", "=7", "x4.0", "", "Frag", "2", "Blast", "3", "5", "x3.0"},
		{"Launchers", "aT", "AT Missile", "+3", "=4", "x1.0", "+1", "Frag", "2", "Pen", "3", "5", "x2.0"},
		{"Launchers", "Gr", "Grenade", "+1", "=4", "x0.8", "", "Frag", "2", "Blast", "2", "4", "x1.0"},
		{"Launchers", "M", "Missile", "+1", "=6", "x2.2", "", "Frag", "2", "Pen", "2", "4", "x5.0"},
		{"Launchers", "RAM", "RAM Grenade", "+2", "=6", "x1.0", "", "Frag", "2", "Blast", "2", "4", "x3.0"},
		{"Launchers", "R", "Rocket", "-1", "=5", "x3.0", "", "Frag", "2", "Pen", "2", "4", "x1.0"},
	}
	return data
}

func weaponBurden() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "TL", "R=", "Mass", "qreBs", "Misc", "D2", "Comment", "Cr"},
		{"Burden", "", "(blank)", "0", "0", "x1.0", "0", "", "0", "", "x1.0"},
		{"Burden", "aD", "Anti-Designator", "+3", "+1", "x3.0", "+3", "", "1", "Not Pistols. Shotguns.", "x3.0"},
		{"Burden", "B", "Body", "+2", "=1", "x0.5", "-4", "", "-1", "Only Pistols.", "x3.0"},
		{"Burden", "D", "Disposable", "+3", "0", "x0.9", "-1", "Q= -2", "0", "", "x0.5"},
		{"Burden", "H", "Heavy", "0", "+1", "x1.3", "+3", "", "1", "Not Laser", "x2.0"},
		{"Burden", "Lt", "Light", "0", "-1", "x0.7", "-1", "", "-1", "Not Laser", "x1.1"},
		{"Burden", "M", "Magnum", "+1", "+1", "x1.1", "+1", "", "1", "Only Pistols.", "x1.1"},
		{"Burden", "M", "Medium", "0", "0", "x1.0", "0", "", "0", "Not Pistols.", "x1.0"},
		{"Burden", "R", "Recoilless", "+1", "-1", "x1.2", "0", "", "1", "", "x3.0"},
		{"Burden", "Sn", "Snub", "+1", "=2", "x0.7", "-3", "", "1", "", "x1.5"},
		{"Burden", "Vh", "Vheavy", "0", "+5", "x4.0", "+4", "", "5", "", "x5.0"},
		{"Burden", "Vl", "Vlight", "+1", "-2", "x0.6", "-2", "", "-1", "", "x2.0"},
		{"Burden", "Vrf", "VRF", "+2", "0", "x14.0", "+5", "", "1", "Only Guns and MGs", "x9.0"},
	}
	return data
}

func weaponStage() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "TL", "R=", "Mass", "qreBs", "Misc", "D2", "Comment", "Cr"},
		{"Stage", "", "(blank)", "0", "0", "x1.0", "0", "", "0", "", "x1.0"},
		{"Stage", "A", "Advanced", "+3", "0", "x0.8", "-3", "", "2", "", "x2.0"},
		{"Stage", "Alt", "Alternate", "0", "+1", "x1.1", "F", "", "2", "", "x1.1"},
		{"Stage", "B", "Basic", "0", "+0", "x1.3", "+1", "", "0", "", "x0.7"},
		{"Stage", "E", "Early", "-1", "-1", "x1.7", "+1", "", "0", "EOU -1", "x1.2"},
		{"Stage", "Exp", "Experimenta", "-3", "-1", "x2.0", "+3", "R= -2", "0", "", "x4.0"},
		{"Stage", "Gen", "Generic", "+1", "0", "x1.0", "0", "", "0", "", "x0.5"},
		{"Stage", "Im", "Improved", "+1", "0", "x1.0", "-1", "R= +1", "1", "EOU +1", "x1.1"},
		{"Stage", "Mod", "Modified", "+2", "0", "x0.9", "0", "", "1", "", "x1.2"},
		{"Stage", "Pr", "Precision", "+6", "+3", "x4.0", "+2", "", "0", "Only Designators.", "x5.0"},
		{"Stage", "P", "Prototype", "-2", "-1", "x1.9", "+2", "", "0", "", "x3.0"},
		{"Stage", "R", "Remote", "+1", "0", "x1.0", "0", "", "0", "Not Pistols.", "x7.0"},
		{"Stage", "Sn", "Sniper", "+1", "+1", "x1.1", "+1", "Q= +2", "0", "Only Rifles.", "x2.0"},
		{"Stage", "St", "Standard", "0", "0", "x1.0", "0", "", "1", "", "x1.0"},
		{"Stage", "T", "Target", "0", "0", "x1.1", "+1", "Q= +2", "0", "Only Rifles and Pistols.", "x1.5"},
		{"Stage", "Ul", "Ultimate", "+4", "0", "x0.7", "-4", "R= +4", "2", "", "x1.4"},
	}
	return data
}

func weaponUsers() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "Mass", "qreBs", "Misc", "Comment"},
		{"Users", "", "(blank)", "x1.0", "0", "", ""},
		{"Users", "U", "Universal", "x1.1", "+1", "EOU= -1", "Usable by ANY manipulator."},
		{"Users", "M", "Man", "x1.0", "0", "EOU= 0", ""},
		{"Users", "V", "Vargr", "x1.0", "0", "EOU= -1", ""},
		{"Users", "K", "K’kree", "x1.3", "+2", "EOU= 0", ""},
		{"Users", "H", "Grasper (Hiver)", "x1.0", "0", "EOU= -1", ""},
		{"Users", "P", "Paw (Aslan)", "x1.0", "0", "EOU= -1", ""},
		{"Users", "G", "Gripper", "x1.0", "0", "EOU= -2", ""},
		{"Users", "T", "Tentacle (Vegan)", "x1.0", "0", "EOU= -2", ""},
		{"Users", "S", "Socket", "x1.0", "0", "EOU=- 2", ""},
	}
	return data
}

func weaponPortability() [][]string {
	data := [][]string{
		{"Category", "Code", "Descriptor", "R=", "Mass", "qreBs"},
		{"Portability", "", "(blank)", "0", "x1.0", "0"},
		{"Portability", "C", "Crewed", "0", "x1.0", "+1"},
		{"Portability", "F", "Fixed", "+1", "x1.0", "+4"},
		{"Portability", "P", "Portable", "+1", "x1.0", "-2"},
		{"Portability", "V", "Vehicle Mount", "+1", "x1.0", "0"},
		{"Portability", "T", "Turret", "0", "x1.0", "0"},
	}
	return data
}
