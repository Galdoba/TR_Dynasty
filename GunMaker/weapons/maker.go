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
	dCost      int
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

func callDesignTask(key string) designTask {
	dt := designTask{"", "", "", 0, 0, 0.0, 0, "", 0, "", 0, "", 0, 0}
	designMap := make(map[string]designTask)
	designMap["Short Blades_Knife"] = designTask{"Short Blades", "Knife", "K", 1, 0, 0.5, 0, "Cuts", 2, "", 0, "", 0, 50}
	designMap["Short Blades_Dagger"] = designTask{"Short Blades", "Dagger", "D", 2, 0, 0.5, 0, "Cuts", 2, "", 0, "", 0, 50}
	designMap["Short Blades_Trench Knife"] = designTask{"Short Blades", "Trench Knife", "TK", 4, 0, 1.0, 0, "Cuts", 2, "Blow", 1, "", 0, 100}
	designMap["Short Blades_Big Knife"] = designTask{"Short Blades", "Big Knife", "BK", 5, 0, 3.0, 0, "Cuts", 2, "Pen", 2, "", 0, 200}
	designMap["Short Blades_Great Big Knife"] = designTask{"Short Blades", "Great Big Knife", "GBK", 6, 1, 6.0, 0, "Cuts", 2, "Pen", 2, "", 0, 900}
	designMap["Medium Blades_Sword"] = designTask{"Medium Blades", "Sword", "S", 3, 1, 2.0, 0, "Cuts", 2, "", 0, "", 0, 300}
	designMap["Medium Blades_Short Sword"] = designTask{"Medium Blades", "Short Sword", "sS", 3, 1, 1.0, -1, "Cuts", 2, "", 0, "", 0, 300}
	designMap["Medium Blades_Broadsword"] = designTask{"Medium Blades", "Broadsword", "bS", 4, 1, 3.0, 0, "Cuts", 3, "", 0, "", 0, 700}
	designMap["Medium Blades_Cutlass"] = designTask{"Medium Blades", "Cutlass", "C", 3, 1, 2.0, 0, "Cuts", 2, "", 0, "", 0, 200}
	designMap["Medium Blades_Officers Cutlass"] = designTask{"Medium Blades", "Officers Cutlass", "OC", 5, 1, 1.0, 0, "Cuts", 2, "", 0, "", 0, 400}
	designMap["Long Blades_Pike"] = designTask{"Long Blades", "Pike", "P", 1, 1, 2.0, 3, "Cuts", 2, "", 0, "", 0, 50}
	designMap["Special Blades_Axe"] = designTask{"Special Blades", "Axe", "Ax", 2, 0, 2.0, 0, "Cuts", 3, "", 0, "", 0, 60}
	designMap["Special Blades_Space Axe"] = designTask{"Special Blades", "Space Axe", "A", 9, 1, 2.0, 0, "Cuts", 2, "Pen", 2, "", 0, 100}
	designMap["Special Blades_Vibro-Blade"] = designTask{"Special Blades", "Vibro-Blade", "V", 10, 1, 0.5, 0, "Cuts", 2, "", 0, "", 0, 900}
	designMap["Special Blades_Mace"] = designTask{"Special Blades", "Mace", " ", 2, 1, 4.0, 0, "Cuts", 1, "Blow", 2, "", 0, 100}
	designMap["Special Blades_Club"] = designTask{"Special Blades", "Club", " ", 1, 1, 2.0, 0, "Blow", 1, "", 0, "", 0, 10}
	//
	designMap["Body Weapons_Fists"] = designTask{"Body Weapons", "Fists", "Fi", 0, 0, 0.0, 0, "Blow", 1, "", 0, "", 0, 0}
	designMap["Body Weapons_Tentacle"] = designTask{"Body Weapons", "Tentacle", "Te", 0, 0, 0.0, 0, "Hit", 1, "Suff", 1, "", 0, 0}
	designMap["Body Weapons_Horns"] = designTask{"Body Weapons", "Horns", "Ho", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	designMap["Body Weapons_Tusks"] = designTask{"Body Weapons", "Tusks", "Tu", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	designMap["Body Weapons_Fangs"] = designTask{"Body Weapons", "Fangs", "Fa", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	designMap["Body Weapons_Teeth"] = designTask{"Body Weapons", "Teeth", "T", 0, 0, 0.0, 0, "Cut", 1, "", 0, "", 0, 0}
	designMap["Body Weapons_Claws"] = designTask{"Body Weapons", "Claws", "Cl", 0, 0, 0.0, 0, "Cut", 1, "", 0, "", 0, 0}
	designMap["Body Weapons_Dew Claws"] = designTask{"Body Weapons", "Dew Claws", "Dc", 0, 0, 0.0, 0, "Cut", 2, "", 0, "", 0, 0}
	designMap["Body Weapons_Hooves"] = designTask{"Body Weapons", "Hooves", "H", 0, 0, 0.0, 0, "Blow", 2, "", 0, "", 0, 0}
	designMap["Body Weapons_Spikes"] = designTask{"Body Weapons", "Spikes", "Sp", 0, 0, 0.0, 0, "Pen", 2, "", 0, "", 0, 0}
	designMap["Body Weapons_Sting"] = designTask{"Body Weapons", "Sting", "St", 0, 0, 0.0, 0, "Pen", 1, "Poison", 2, "", 0, 0}
	//
	designMap["Artilery_Gun"] = designTask{"Artilery", "Gun", "G", 6, 4, 9.0, 1, "*", 2, "", 0, "", 0, 5000}
	designMap["Artilery_Gatling"] = designTask{"Artilery", "Gatling", "Ga", 7, 4, 40.0, 2, "*", 3, "", 0, "", 0, 8000}
	designMap["Artilery_Cannon"] = designTask{"Artilery", "Cannon", "C", 6, 6, 200.0, 4, "*", 4, "", 0, "", 0, 10000}
	designMap["Artilery_Autocannon"] = designTask{"Artilery", "Autocannon", "aC", 8, 6, 300.0, 4, "*", 5, "", 0, "", 0, 30000}
	designMap["Long Guns_Rifle"] = designTask{"Long Guns", "Rifle", "R", 5, 5, 4.0, 0, "Bullet", 2, "", 0, "", 0, 500}
	designMap["Long Guns_Carbine"] = designTask{"Long Guns", "Carbine", "C", 5, 4, 3.0, -1, "Bullet", 1, "", 0, "", 0, 400}
	designMap["Hand Guns_Pistol"] = designTask{"Hand Guns", "Pistol", "P", 5, 2, 1.1, 0, "Bullet", 1, "", 0, "", 0, 150}
	designMap["Hand Guns_Revolver"] = designTask{"Hand Guns", "Revolver", "R", 4, 2, 1.25, 0, "Bullet", 1, "", 0, "", 0, 100}
	designMap["Shotguns_Shotgun"] = designTask{"Shotguns", "Shotgun", "S", 4, 2, 4.0, 0, "Frag", 2, "", 0, "", 0, 300}
	designMap["Machineguns_Machinegun"] = designTask{"Machineguns", "Machinegun", "Mg", 6, 5, 8.0, 1, "Bullet", 4, "", 0, "", 0, 3000}
	designMap["Projectors_Projector"] = designTask{"Projectors", "Projector", "Pj", 9, 0, 1.0, 0, "*", 1, "", 0, "", 0, 300}
	designMap["Designators_Designator"] = designTask{"Designators", "Designator", "D", 7, 5, 10.0, 1, "*", 1, "", 0, "", 0, 2000}
	designMap["Launchers_Launcher"] = designTask{"Launchers", "Launcher", "L", 6, 3, 10.0, 1, "*", 1, "", 0, "", 0, 1000}
	designMap["Launchers_Multi-Launcher"] = designTask{"Launchers", "Multi-Launcher", "mL", 8, 5, 8.0, 1, "*", 1, "", 0, "", 0, 3000}
	if val, ok := designMap[key]; ok {
		return val
	}
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
