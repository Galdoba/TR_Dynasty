package weapons

import "github.com/Galdoba/TR_Dynasty/dice"

type FillForm struct {
	drs []designTask
	// Type
	// SubType
	// Descriptor
	// Burden
	// Stage
	// User
	// Notes
	// Options
	// Controls
	// Portability
	// QREBS
	// Totals
	//
	// Model
	// TL
	// Range
	// Mass
	// Burden
	// H1
	// H2
	// D1
	// D2
	// H3
	// D3
	// Cost
}

type designTask struct {
	category   string
	descriptor string
	dModel     string //code
	dTL        int    //
	dRange     int    //
	dMass      float64
	dBurden    int
	dH1        string
	dD1        int
	dH2        string
	dD2        int
	dH3        string
	dD3        int
	dCost      float64
}

type weapon struct{}

type gunMaker struct {
	dp *dice.Dicepool //дайспул который используется для бросков
}

type Maker interface { //формирует функции задания для дизайна продукта
	Apply(string, string) designTask
	Roll(string) designTask
	ConcludeDesignProcess() designTask
}

func callWeaponType(key string) designTask {
	dt := designTask{"", "", "", 0, 0, 0.0, 0, "", 0, "", 0, "", 0, 1}
	weaponTypeMap := make(map[string]designTask)
	weaponTypeMap["Short Blades_Knife"] = designTask{"Short Blades", "Knife", "K", 1, 0, 0.5, 0, "Cuts", 2, "", 0, "", 0, 50}
	weaponTypeMap["Short Blades_Dagger"] = designTask{"Short Blades", "Dagger", "D", 2, 0, 0.5, 0, "Cuts", 2, "", 0, "", 0, 50}
	weaponTypeMap["Short Blades_Trench Knife"] = designTask{"Short Blades", "Trench Knife", "TK", 4, 0, 1.0, 0, "Cuts", 2, "Blow", 1, "", 0, 100}
	weaponTypeMap["Short Blades_Big Knife"] = designTask{"Short Blades", "Big Knife", "BK", 5, 0, 3.0, 0, "Cuts", 2, "Pen", 2, "", 0, 200}
	weaponTypeMap["Short Blades_Great Big Knife"] = designTask{"Short Blades", "Great Big Knife", "GBK", 6, 1, 6.0, 0, "Cuts", 2, "Pen", 2, "", 0, 900}
	weaponTypeMap["Medium Blades_Sword"] = designTask{"Medium Blades", "Sword", "S", 3, 1, 2.0, 0, "Cuts", 2, "", 0, "", 0, 300}
	weaponTypeMap["Medium Blades_Short Sword"] = designTask{"Medium Blades", "Short Sword", "sS", 3, 1, 1.0, -1, "Cuts", 2, "", 0, "", 0, 300}
	weaponTypeMap["Medium Blades_Broadsword"] = designTask{"Medium Blades", "Broadsword", "bS", 4, 1, 3.0, 0, "Cuts", 3, "", 0, "", 0, 700}
	weaponTypeMap["Medium Blades_Cutlass"] = designTask{"Medium Blades", "Cutlass", "C", 3, 1, 2.0, 0, "Cuts", 2, "", 0, "", 0, 200}
	weaponTypeMap["Medium Blades_Officers Cutlass"] = designTask{"Medium Blades", "Officers Cutlass", "OC", 5, 1, 1.0, 0, "Cuts", 2, "", 0, "", 0, 400}
	weaponTypeMap["Long Blades_Pike"] = designTask{"Long Blades", "Pike", "P", 1, 1, 2.0, 3, "Cuts", 2, "", 0, "", 0, 50}
	weaponTypeMap["Special Blades_Axe"] = designTask{"Special Blades", "Axe", "Ax", 2, 0, 2.0, 0, "Cuts", 3, "", 0, "", 0, 60}
	weaponTypeMap["Special Blades_Space Axe"] = designTask{"Special Blades", "Space Axe", "A", 9, 1, 2.0, 0, "Cuts", 2, "Pen", 2, "", 0, 100}
	weaponTypeMap["Special Blades_Vibro-Blade"] = designTask{"Special Blades", "Vibro-Blade", "V", 10, 1, 0.5, 0, "Cuts", 2, "", 0, "", 0, 900}
	weaponTypeMap["Special Blades_Mace"] = designTask{"Special Blades", "Mace", " ", 2, 1, 4.0, 0, "Cuts", 1, "Blow", 2, "", 0, 100}
	weaponTypeMap["Special Blades_Club"] = designTask{"Special Blades", "Club", " ", 1, 1, 2.0, 0, "Blow", 1, "", 0, "", 0, 10}
	//
	weaponTypeMap["Body Weapons_Fists"] = designTask{"Body Weapons", "Fists", "Fi", 0, 0, 0.0, 0, "Blow", 1, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Tentacle"] = designTask{"Body Weapons", "Tentacle", "Te", 0, 0, 0.0, 0, "Hit", 1, "Suff", 1, "", 0, 0}
	weaponTypeMap["Body Weapons_Horns"] = designTask{"Body Weapons", "Horns", "Ho", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Tusks"] = designTask{"Body Weapons", "Tusks", "Tu", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Fangs"] = designTask{"Body Weapons", "Fangs", "Fa", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Teeth"] = designTask{"Body Weapons", "Teeth", "T", 0, 0, 0.0, 0, "Cut", 1, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Claws"] = designTask{"Body Weapons", "Claws", "Cl", 0, 0, 0.0, 0, "Cut", 1, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Dew Claws"] = designTask{"Body Weapons", "Dew Claws", "Dc", 0, 0, 0.0, 0, "Cut", 2, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Hooves"] = designTask{"Body Weapons", "Hooves", "H", 0, 0, 0.0, 0, "Blow", 2, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Spikes"] = designTask{"Body Weapons", "Spikes", "Sp", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	weaponTypeMap["Body Weapons_Sting"] = designTask{"Body Weapons", "Sting", "St", 0, 0, 0.0, 0, "Pen", 1, "Poison", 2, "", 0, 0}
	//
	weaponTypeMap["Artilery_Gun"] = designTask{"Artilery", "Gun", "G", 6, 4, 9.0, 1, "*", 2, "", 0, "", 0, 5000}
	weaponTypeMap["Artilery_Gatling"] = designTask{"Artilery", "Gatling", "Ga", 7, 4, 40.0, 2, "*", 3, "", 0, "", 0, 8000}
	weaponTypeMap["Artilery_Cannon"] = designTask{"Artilery", "Cannon", "C", 6, 6, 200.0, 4, "*", 4, "", 0, "", 0, 10000}
	weaponTypeMap["Artilery_Autocannon"] = designTask{"Artilery", "Autocannon", "aC", 8, 6, 300.0, 4, "*", 5, "", 0, "", 0, 30000}
	weaponTypeMap["Long Guns_Rifle"] = designTask{"Long Guns", "Rifle", "R", 5, 5, 4.0, 0, "Bullet", 2, "", 0, "", 0, 500}
	weaponTypeMap["Long Guns_Carbine"] = designTask{"Long Guns", "Carbine", "C", 5, 4, 3.0, -1, "Bullet", 1, "", 0, "", 0, 400}
	weaponTypeMap["Hand Guns_Pistol"] = designTask{"Hand Guns", "Pistol", "P", 5, 2, 1.1, 0, "Bullet", 1, "", 0, "", 0, 150}
	weaponTypeMap["Hand Guns_Revolver"] = designTask{"Hand Guns", "Revolver", "R", 4, 2, 1.25, 0, "Bullet", 1, "", 0, "", 0, 100}
	weaponTypeMap["Shotguns_Shotgun"] = designTask{"Shotguns", "Shotgun", "S", 4, 2, 4.0, 0, "Frag", 2, "", 0, "", 0, 300}
	weaponTypeMap["Machineguns_Machinegun"] = designTask{"Machineguns", "Machinegun", "Mg", 6, 5, 8.0, 1, "Bullet", 4, "", 0, "", 0, 3000}
	weaponTypeMap["Projectors_Projector"] = designTask{"Projectors", "Projector", "Pj", 9, 0, 1.0, 0, "*", 1, "", 0, "", 0, 300}
	weaponTypeMap["Designators_Designator"] = designTask{"Designators", "Designator", "D", 7, 5, 10.0, 1, "*", 1, "", 0, "", 0, 2000}
	weaponTypeMap["Launchers_Launcher"] = designTask{"Launchers", "Launcher", "L", 6, 3, 10.0, 1, "*", 1, "", 0, "", 0, 1000}
	weaponTypeMap["Launchers_Multi-Launcher"] = designTask{"Launchers", "Multi-Launcher", "mL", 8, 5, 8.0, 1, "*", 1, "", 0, "", 0, 3000}
	if val, ok := weaponTypeMap[key]; ok {
		return val
	}
	weaponDescriptorMap := make(map[string]designTask)
	weaponDescriptorMap["Artilery_(blank)"] = designTask{"Artilery", "(blank)", "", 0, 0, 1.0, 0, "", 0, "", 0, "", 0, 1.0}
	weaponDescriptorMap["Artilery_Anti-Flyer"] = designTask{"Artilery", "Anti-Flyer", "aF", 4, 6, 6.0, 0, "", 0, "Frag", 1, "Blast", 3, 3}
	weaponDescriptorMap["Artilery_Anti-Tank"] = designTask{"Artilery", "Anti-Tank", "aT", 0, 5, 8.0, 0, "", 0, "Pen", 3, "Blast", 3, 2}
	weaponDescriptorMap["Artilery_Assault"] = designTask{"Artilery", "Assault", "A", 2, 4, 0.8, 0, "", 0, "Bang", 1, "Blast", 2, 1.5}
	weaponDescriptorMap["Artilery_Fusion"] = designTask{"Artilery", "Fusion", "F", 7, 4, 2.3, 0, "", 0, "Pen", 4, "Burn", 4, 6}
	weaponDescriptorMap["Artilery_Gauss"] = designTask{"Artilery", "Gauss", "G", 7, 4, 0.9, 0, "", 0, "Bullet", 3, "", 0, 2}
	weaponDescriptorMap["Artilery_Plasma"] = designTask{"Artilery", "Plasma", "P", 5, 4, 2.5, 0, "", 0, "Pen", 3, "Burn", 3, 2}
	weaponDescriptorMap["Long Guns_(blank)"] = designTask{"Long Guns", "(blank)", "", 0, 0, 1.0, 0, "", 0, "", 0, "", 0, 1.0}
	weaponDescriptorMap["Long Guns_Accelerator"] = designTask{"Long Guns", "Accelerator", "Ac", 4, -1, 0.6, 0, "", 0, "Bullet", 2, "", 0, 3.0}
	weaponDescriptorMap["Long Guns_Accelerator"] = designTask{"Long Guns", "Accelerator", "Ac", 4, -1, 0.6, 0, "", 0, "Bullet", 2, "", 0, 3.0}

	return dt
}

/*

lib maker
SAMPLE CODE:

gunMkr := maker.NewGunMaker()
wp := gunMkr.MakeWeapon(
	gunMkr.Apply(item1, descr1)
	gunMkr.Apply(item2, descr2)
	gunMkr.Roll(item3, item4 ...)
	gunMkr.ConcludeDesignProcess()
)
wp.Model()
wp.LongName()
wp.Fire(effectDM)
wp.DoStuff(args...)


*/
