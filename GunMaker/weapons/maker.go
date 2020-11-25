package weapons

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/dice"
)

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
	dRange     string //
	dMass      float64
	dBurden    int
	dH1        string
	dD1        int
	dH2        string
	dD2        int
	dH3        string
	dD3        int
	misc       string
	dCost      float64
}

type weapon struct {
	wType        string
	wSubType     string
	wDescriptor  string
	wBurden      string
	wStage       string
	wUser        string
	wPortability string
	qrebs        string
	tl           int
	wRange       int
	mass         float64
	damageDice   string
	cost         float64
	traits       string
	description  string
}

func TestGunMaker() {
	gMaker := NewGunMaker()
	wp := gMaker.MakeWeapon(
		gMaker.Apply("Handguns", "Pistol"),
		gMaker.Apply("Handgun", "Laser"),
	)
	fmt.Println(wp, gMaker)
}

func (gmkr *gunMaker) MakeWeapon(dT ...designTask) weapon {
	wp := weapon{}
	for _, designTask := range dT {
		wp.absorb(designTask)
	}
	return wp
}

func (wp *weapon) absorb(dt designTask) {
	fmt.Println("Run 'absorb'")
	err := errors.New("No Error")
	if dt.misc == "Init" {
		if wp.wSubType != "" {
			return
		}
		fmt.Println("Do initial")
		wp.wType = dt.category
		wp.wSubType = dt.descriptor
		wp.wRange = recalculateRange(wp.wRange, dt.dRange)
		wp.damageDice = strconv.Itoa(dt.dD1+dt.dD2+dt.dD3) + "D"
		wp.traits = dt.dH1 + dt.dH2 + dt.dH3
		wp.mass = dt.dMass
		wp.cost = dt.dCost
	}
	fmt.Println(err)
	fmt.Println("End 'absorb'")
}

func NewGunMaker(seed ...int64) gunMaker {
	gMaker := gunMaker{}
	if len(seed) == 0 {
		gMaker.dp = dice.New(time.Now().UnixNano())
	} else {
		gMaker.dp = dice.New(seed[0])
	}
	return gMaker
}

type gunMaker struct {
	dp *dice.Dicepool //дайспул который используется для бросков
}

type Maker interface { //формирует функции задания для дизайна продукта
	Apply(string, string) designTask // key2 = Roll - означает выбрать рандомно
	//Roll(string) designTask //
	//ConcludeDesignProcess() designTask // не уверен что оно нужно
}

func (gmkr *gunMaker) Apply(key1, key2 string) designTask {

	key := key1 + "_" + key2
	return callDesignTask(key)
}

func callDesignTask(key string) designTask {
	dt := designTask{"", "", "", 0, "0", 0.0, 0, "", 0, "", 0, "", 0, "", 1}
	designTaskMap := make(map[string]designTask)
	designTaskMap["Short Blades_Knife"] = designTask{"Short Blades", "Knife", "K", 1, "R", 0.5, 0, "Cuts", 2, "", 0, "", 0, "Init", 50}
	designTaskMap["Short Blades_Dagger"] = designTask{"Short Blades", "Dagger", "D", 2, "R", 0.5, 0, "Cuts", 2, "", 0, "", 0, "Init", 50}
	designTaskMap["Short Blades_Trench Knife"] = designTask{"Short Blades", "Trench Knife", "TK", 4, "R", 1.0, 0, "Cuts", 2, "Blow", 1, "", 0, "Init", 100}
	designTaskMap["Short Blades_Big Knife"] = designTask{"Short Blades", "Big Knife", "BK", 5, "T", 3.0, 0, "Cuts", 2, "Pen", 2, "", 0, "Init", 200}
	designTaskMap["Short Blades_Great Big Knife"] = designTask{"Short Blades", "Great Big Knife", "GBK", 6, "1", 6.0, 0, "Cuts", 2, "Pen", 2, "", 0, "Init", 900}
	designTaskMap["Medium Blades_Sword"] = designTask{"Medium Blades", "Sword", "S", 3, "1", 2.0, 0, "Cuts", 2, "", 0, "", 0, "Init", 300}
	designTaskMap["Medium Blades_Short Sword"] = designTask{"Medium Blades", "Short Sword", "sS", 3, "1", 1.0, -1, "Cuts", 2, "", 0, "", 0, "Init", 300}
	designTaskMap["Medium Blades_Broadsword"] = designTask{"Medium Blades", "Broadsword", "bS", 4, "1", 3.0, 0, "Cuts", 3, "", 0, "", 0, "Init", 700}
	designTaskMap["Medium Blades_Cutlass"] = designTask{"Medium Blades", "Cutlass", "C", 3, "1", 2.0, 0, "Cuts", 2, "", 0, "", 0, "2D", 200}
	designTaskMap["Medium Blades_Officers Cutlass"] = designTask{"Medium Blades", "Officers Cutlass", "OC", 5, "1", 1.0, 0, "Cuts", 2, "", 0, "", 0, "Init", 400}
	designTaskMap["Long Blades_Pike"] = designTask{"Long Blades", "Pike", "P", 1, "1", 2.0, 3, "Cuts", 2, "", 0, "", 0, "Init", 50}
	designTaskMap["Special Blades_Axe"] = designTask{"Special Blades", "Axe", "Ax", 2, "T", 2.0, 0, "Cuts", 3, "", 0, "", 0, "Init", 60}
	designTaskMap["Special Blades_Space Axe"] = designTask{"Special Blades", "Space Axe", "A", 9, "1", 2.0, 0, "Cuts", 2, "Pen", 2, "", 0, "Init", 100}
	designTaskMap["Special Blades_Vibro-Blade"] = designTask{"Special Blades", "Vibro-Blade", "V", 10, "1", 0.5, 0, "Cuts", 2, "", 0, "", 0, "Init", 900}
	designTaskMap["Special Blades_Mace"] = designTask{"Special Blades", "Mace", " ", 2, "1", 4.0, 0, "Cuts", 1, "Blow", 2, "", 0, "Init", 100}
	designTaskMap["Special Blades_Club"] = designTask{"Special Blades", "Club", " ", 1, "1", 2.0, 0, "Blow", 1, "", 0, "", 0, "Init", 10}
	designTaskMap["Body Weapons_Fists"] = designTask{"Body Weapons", "Fists", "Fi", 0, "R", 0.0, 0, "Blow", 1, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Tentacle"] = designTask{"Body Weapons", "Tentacle", "Te", 0, "0", 0.0, 0, "Hit", 1, "Suff", 1, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Horns"] = designTask{"Body Weapons", "Horns", "Ho", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Tusks"] = designTask{"Body Weapons", "Tusks", "Tu", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Fangs"] = designTask{"Body Weapons", "Fangs", "Fa", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Teeth"] = designTask{"Body Weapons", "Teeth", "T", 0, "R", 0.0, 0, "Cut", 1, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Claws"] = designTask{"Body Weapons", "Claws", "Cl", 0, "R", 0.0, 0, "Cut", 1, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Dew Claws"] = designTask{"Body Weapons", "Dew Claws", "Dc", 0, "R", 0.0, 0, "Cut", 2, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Hooves"] = designTask{"Body Weapons", "Hooves", "H", 0, "R", 0.0, 0, "Blow", 2, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Spikes"] = designTask{"Body Weapons", "Spikes", "Sp", 0, "0", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0}
	designTaskMap["Body Weapons_Sting"] = designTask{"Body Weapons", "Sting", "St", 0, "R", 0.0, 0, "Pen", 1, "Poison", 2, "", 0, "Init", 0}
	designTaskMap["Artilery_Gun"] = designTask{"Artilery", "Gun", "G", 6, "4", 9.0, 1, "*", 2, "", 0, "", 0, "Init", 5000}
	designTaskMap["Artilery_Gatling"] = designTask{"Artilery", "Gatling", "Ga", 7, "4", 40.0, 2, "*", 3, "", 0, "", 0, "Init", 8000}
	designTaskMap["Artilery_Cannon"] = designTask{"Artilery", "Cannon", "C", 6, "6", 200.0, 4, "*", 4, "", 0, "", 0, "Init", 10000}
	designTaskMap["Artilery_Autocannon"] = designTask{"Artilery", "Autocannon", "aC", 8, "6", 300.0, 4, "*", 5, "", 0, "", 0, "Init", 30000}
	designTaskMap["Long Guns_Rifle"] = designTask{"Long Guns", "Rifle", "R", 5, "5", 4.0, 0, "Bullet", 2, "", 0, "", 0, "Init", 500}
	designTaskMap["Long Guns_Carbine"] = designTask{"Long Guns", "Carbine", "C", 5, "4", 3.0, -1, "Bullet", 1, "", 0, "", 0, "Init", 400}
	designTaskMap["Handguns_Pistol"] = designTask{"Handguns", "Pistol", "P", 5, "2", 1.1, 0, "Bullet", 1, "", 0, "", 0, "Init", 150}
	designTaskMap["Handguns_Revolver"] = designTask{"Handguns", "Revolver", "R", 4, "2", 1.25, 0, "Bullet", 1, "", 0, "", 0, "Init", 100}
	designTaskMap["Shotguns_Shotgun"] = designTask{"Shotguns", "Shotgun", "S", 4, "2", 4.0, 0, "Frag", 2, "", 0, "", 0, "Init", 300}
	designTaskMap["Machineguns_Machinegun"] = designTask{"Machineguns", "Machinegun", "Mg", 6, "5", 8.0, 1, "Bullet", 4, "", 0, "", 0, "Init", 3000}
	designTaskMap["Projectors_Projector"] = designTask{"Projectors", "Projector", "Pj", 9, "0", 1.0, 0, "*", 1, "", 0, "", 0, "Init", 300}
	designTaskMap["Designators_Designator"] = designTask{"Designators", "Designator", "D", 7, "5", 10.0, 1, "*", 1, "", 0, "", 0, "Init", 2000}
	designTaskMap["Launchers_Launcher"] = designTask{"Launchers", "Launcher", "L", 6, "3", 10.0, 1, "*", 1, "", 0, "", 0, "Init", 1000}
	designTaskMap["Launchers_Multi-Launcher"] = designTask{"Launchers", "Multi-Launcher", "mL", 8, "5", 8.0, 1, "*", 1, "", 0, "", 0, "Init", 3000}
	designTaskMap["Artilery_(blank)"] = designTask{"Artilery", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
	designTaskMap["Artilery_Anti-Flyer"] = designTask{"Artilery", "Anti-Flyer", "aF", 4, "=6", 6.0, 0, "", 0, "Frag", 1, "Blast", 3, "4D", 3}
	designTaskMap["Artilery_Anti-Tank"] = designTask{"Artilery", "Anti-Tank", "aT", 0, "=5", 8.0, 0, "", 0, "Pen", 3, "Blast", 3, "6D", 2}
	designTaskMap["Artilery_Assault"] = designTask{"Artilery", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "3D", 1.5}
	designTaskMap["Artilery_Fusion"] = designTask{"Artilery", "Fusion", "F", 7, "=4", 2.3, 0, "", 0, "Pen", 4, "Burn", 4, "8D", 6}
	designTaskMap["Artilery_Gauss"] = designTask{"Artilery", "Gauss", "G", 7, "=4", 0.9, 0, "", 0, "Bullet", 3, "", 0, "3D", 2}
	designTaskMap["Artilery_Plasma"] = designTask{"Artilery", "Plasma", "P", 5, "=4", 2.5, 0, "", 0, "Pen", 3, "Burn", 3, "6D", 2}
	designTaskMap["Long Guns_(blank)"] = designTask{"Long Guns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
	designTaskMap["Long Guns_Accelerator"] = designTask{"Long Guns", "Accelerator", "Ac", 4, "", 0.6, 0, "", 0, "Bullet", 2, "", 0, "2D", 3.0}
	designTaskMap["Long Guns_Assault"] = designTask{"Long Guns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Blast", 2, "Bang", 1, "3D", 1.5}
	designTaskMap["Long Guns_Battle"] = designTask{"Long Guns", "Battle", "B", 1, "=5", 1.0, 1, "", 0, "Bullet", 1, "", 0, "1D", 0.8}
	designTaskMap["Long Guns_Combat"] = designTask{"Long Guns", "Combat", "C", 2, "=3", 0.9, 0, "", 0, "Frag", 2, "", 0, "2D", 1.5}
	designTaskMap["Long Guns_Dart-1"] = designTask{"Long Guns", "Dart-1", "D1", 1, "=4", 0.6, 0, "", 0, "Tranq", 1, "", 0, "1D", 0.9}
	designTaskMap["Long Guns_Dart-2"] = designTask{"Long Guns", "Dart-2", "D2", 1, "=4", 0.6, 0, "", 0, "Tranq", 2, "", 0, "2D", 0.9}
	designTaskMap["Long Guns_Dart-3"] = designTask{"Long Guns", "Dart-3", "D3", 1, "=4", 0.6, 0, "", 0, "Tranq", 3, "", 0, "3D", 0.9}
	designTaskMap["Long Guns_Poison Dart-1"] = designTask{"Long Guns", "Poison Dart-1", "P1", 1, "=4", 1.0, 0, "", 0, "Poison", 1, "", 0, "1D", 0.9}
	designTaskMap["Long Guns_Poison Dart-2"] = designTask{"Long Guns", "Poison Dart-2", "P2", 1, "=4", 1.0, 0, "", 0, "Poison", 2, "", 0, "2D", 0.9}
	designTaskMap["Long Guns_Poison Dart-3"] = designTask{"Long Guns", "Poison Dart-3", "P3", 1, "=4", 1.0, 0, "", 0, "Poison", 3, "", 0, "3D", 0.9}
	designTaskMap["Long Guns_Gauss"] = designTask{"Long Guns", "Gauss", "G", 7, "", 0.9, 0, "", 0, "Bullet", 3, "", 0, "3D", 2.0}
	designTaskMap["Long Guns_Hunting"] = designTask{"Long Guns", "Hunting", "H", 0, "=3", 0.9, -1, "", 0, "Bullet", 1, "", 0, "1D", 1.2}
	designTaskMap["Long Guns_Laser"] = designTask{"Long Guns", "Laser", "L", 5, "", 1.2, 0, "", 0, "Burn", 2, "Pen", 2, "4D", 6.0}
	designTaskMap["Long Guns_Splat"] = designTask{"Long Guns", "Splat", "Sp", 2, "=4", 1.3, 1, "", 0, "Bullet", 1, "", 0, "1D", 2.4}
	designTaskMap["Long Guns_Survival"] = designTask{"Long Guns", "Survival", "S", 0, "=2", 0.5, 0, "", 0, "Bullet", 1, "", 0, "1D", 1.2}
	designTaskMap["Handguns_(blank)"] = designTask{"Handguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
	designTaskMap["Handguns_Accelerator"] = designTask{"Handguns", "Accelerator", "Ac", 4, "", 0.6, 0, "", 0, "Bullet", 2, "", 0, "2D", 3.0}
	designTaskMap["Handguns_Laser"] = designTask{"Handguns", "Laser", "L", 5, "", 1.2, 0, "", 0, "Burn", 2, "Pen", 2, "4D", 2.0}
	designTaskMap["Handguns_Machine"] = designTask{"Handguns", "Machine", "M", 0, "=2", 0.5, 0, "", 0, "Bullet", 1, "", 0, "1D", 1.2}
	designTaskMap["Shotguns_(blank)"] = designTask{"Shotguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
	designTaskMap["Shotguns_Assault"] = designTask{"Shotguns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "3D", 2.0}
	designTaskMap["Shotguns_Hunting"] = designTask{"Shotguns", "Hunting", "H", 0, "=3", 0.9, 0, "", 0, "Bullet", 1, "", 0, "1D", 2.0}
	designTaskMap["Machineguns_(blank)"] = designTask{"Machineguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
	designTaskMap["Machineguns_Anti-Flyer"] = designTask{"Machineguns", "Anti-Flyer", "aF", 4, "=6", 6.0, 0, "", 0, "Frag", 1, "Blast", 3, "4D", 3.0}
	designTaskMap["Machineguns_Assault"] = designTask{"Machineguns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "3D", 1.5}
	designTaskMap["Machineguns_Sub"] = designTask{"Machineguns", "Sub", "S", -1, "=3", 0.3, 0, "", 0, "Bullet", -1, "", 0, "-1D", 0.9}
	designTaskMap["Launchers_AF Missile"] = designTask{"Launcher", "AF Missile", "aF", 4, "=7", 4.0, 0, "", 0, "Frag", 2, "Blast", 3, "5D", 3.0}
	designTaskMap["Launchers_AT Missile"] = designTask{"Launcher", "AT Missile", "aT", 3, "=4", 1.0, 1, "", 0, "Frag", 2, "Pen", 3, "5D", 2.0}
	designTaskMap["Launchers_Grenade"] = designTask{"Launcher", "Grenade", "Gr", 1, "=4", 0.8, 0, "", 0, "Frag", 2, "Blast", 2, "4D", 1.0}
	designTaskMap["Launchers_Missile"] = designTask{"Launcher", "Missile", "M", 1, "=6", 2.2, 0, "", 0, "Frag", 2, "Pen", 2, "4D", 5.0}
	designTaskMap["Launchers_RAM Grenade"] = designTask{"Launcher", "RAM Grenade", "RAM", 2, "=6", 1.0, 0, "", 0, "Frag", 2, "Blast", 2, "4D", 3.0}
	designTaskMap["Launchers_Rocket"] = designTask{"Launcher", "Rocket", "R", -1, "=5", 3.0, 0, "", 0, "Frag", 2, "Pen", 2, "4D", 1.0}
	designTaskMap["Burden_(blank)"] = designTask{"Burden", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Burden_Anti-Designator"] = designTask{"Burden", "Anti-Designator", "aD", 3, "+1", 3.0, 3, "", 0, "", 1, "", 0, "Not Pistol. Not Shotgun", 3.0}
	designTaskMap["Burden_Body"] = designTask{"Burden", "Body", "B", 2, "=1", 0.5, -4, "", 0, "", -1, "", 0, "Only Pistol", 3.0}
	designTaskMap["Burden_Disposable"] = designTask{"Burden", "Disposable", "D", 3, "+0", 0.9, -1, "", 0, "", 0, "", 0, "Q=-2", 0.5}
	designTaskMap["Burden_Heavy"] = designTask{"Burden", "Heavy", "H", 0, "+1", 1.3, 3, "", 0, "", 1, "", 0, "Not Laser", 2.0}
	designTaskMap["Burden_Light"] = designTask{"Burden", "Light", "Lt", 0, "-1", 0.7, -1, "", 0, "", -1, "", 0, "Not Laser", 1.1}
	designTaskMap["Burden_Magnum"] = designTask{"Burden", "Magnum", "M", 1, "+1", 1.1, 1, "", 0, "", 1, "", 0, "Only Pistol", 1.1}
	designTaskMap["Burden_Medium"] = designTask{"Burden", "Medium", "M", 0, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Burden_Recoilless"] = designTask{"Burden", "Recoilless", "R", 1, "+0", 1.2, 0, "", 0, "", 1, "", 0, "", 3.0} // TODO: Check Errata for Range (was -+1, made +0)
	designTaskMap["Burden_Snub"] = designTask{"Burden", "Snub", "Sn", 1, "=2", 0.7, -3, "", 0, "", 1, "", 0, "", 1.5}
	designTaskMap["Burden_Vheavy"] = designTask{"Burden", "Vheavy", "Vh", 0, "+5", 4.0, 4, "", 0, "", 5, "", 0, "", 5.0}
	designTaskMap["Burden_Vlight"] = designTask{"Burden", "Vlight", "Vl", 1, "-2", 0.6, -2, "", 0, "", -1, "", 0, "", 2.0}
	designTaskMap["Burden_VRF"] = designTask{"Burden", "VRF", "Vrf", 2, "+0", 14.0, 5, "", 0, "", 1, "", 0, "Only Gun, Only Machnegun", 9.0}
	designTaskMap["Stage_(blank)"] = designTask{"Stage", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Stage_Advanced"] = designTask{"Stage", "Advanced", "A", 3, "+0", 0.8, -3, "", 0, "", 2, "", 0, "", 2.0}
	designTaskMap["Stage_Alternate"] = designTask{"Stage", "Alternate", "Alt", 0, "+1", 1.1, 0, "", 0, "", 2, "", 0, "B+Flux", 1.1}
	designTaskMap["Stage_Basic"] = designTask{"Stage", "Basic", "B", 0, "+0", 1.3, 1, "", 0, "", 0, "", 0, "", 0.7}
	designTaskMap["Stage_Early"] = designTask{"Stage", "Early", "E", -1, "-1", 1.7, 1, "", 0, "", 0, "", 0, "EOU-1", 0.7}
	designTaskMap["Stage_Experimental"] = designTask{"Stage", "Experimental", "Exp", -3, "-1", 2.0, 3, "", 0, "", 0, "", 0, "R=-2", 4.0}
	designTaskMap["Stage_Generic"] = designTask{"Stage", "Generic", "Gen", 1, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 0.5}
	designTaskMap["Stage_Improved"] = designTask{"Stage", "Improved", "Im", 1, "+0", 1.0, -1, "", 0, "", 1, "", 0, "EOU+1", 1.1}
	designTaskMap["Stage_Modified"] = designTask{"Stage", "Modified", "Mod", 2, "+0", 0.9, 0, "", 0, "", 1, "", 0, "", 1.2}
	designTaskMap["Stage_Precision"] = designTask{"Stage", "Precision", "Pr", 6, "+3", 4.0, 2, "", 0, "", 0, "", 0, "Only Designator", 5.0}
	designTaskMap["Stage_Prototype"] = designTask{"Stage", "Prototype", "P", -2, "-1", 1.9, 2, "", 0, "", 0, "", 0, "", 3.0}
	designTaskMap["Stage_Remote"] = designTask{"Stage", "Remote", "R", 1, "+0", 1.0, 0, "", 0, "", 0, "", 0, "Not Pistol", 7.0}
	designTaskMap["Stage_Sniper"] = designTask{"Stage", "Sniper", "Sn", 1, "+1", 1.1, 1, "", 0, "", 1, "", 0, "Only Rifle, Q=+2", 2.0} //TODO: Check Errata for +1 on D2 for stantard (makes no sense if it is in standard and not Sniper)
	designTaskMap["Stage_Standard"] = designTask{"Stage", "Standard", "St", 0, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}             //TODO: Check Errata for +1 on D2 for stantard (makes no sense if it is in standard and not Sniper)
	designTaskMap["Stage_Target"] = designTask{"Stage", "Target", "T", 0, "+0", 1.1, 1, "", 0, "", 0, "", 0, "Only Rifle, Only Pistol, Q=+2", 1.5}
	designTaskMap["Stage_Ultimate"] = designTask{"Stage", "Ultimate", "Ul", 4, "+0", 0.7, -4, "", 0, "", 2, "", 0, "R=+4", 1.4}
	designTaskMap["Users_(blank)"] = designTask{"Users", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU+0", 1.0}
	designTaskMap["Users_Universal"] = designTask{"Users", "Universal", "U", 0, "", 1.1, 1, "", 0, "", 0, "", 0, "EOU-1", 1.0}
	designTaskMap["Users_Man"] = designTask{"Users", "Man", "M", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU+0", 1.0}
	designTaskMap["Users_Vargr"] = designTask{"Users", "Vargr", "V", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0}
	designTaskMap["Users_K'kree"] = designTask{"Users", "K'kree", "K", 0, "", 1.3, 2, "", 0, "", 0, "", 0, "EOU+0", 1.0}
	designTaskMap["Users_Grasper (Hiver)"] = designTask{"Users", "Grasper (Hiver)", "H", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0}
	designTaskMap["Users_Paw (Aslan)"] = designTask{"Users", "Paw (Aslan)", "P", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0}
	designTaskMap["Users_Gripper"] = designTask{"Users", "Gripper", "G", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0}
	designTaskMap["Users_Tentacle (Vegan)"] = designTask{"Users", "Tentacle (Vegan)", "T", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0}
	designTaskMap["Users_Socket"] = designTask{"Users", "Socket", "S", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0}
	designTaskMap["Portability_(blank)"] = designTask{"Portability", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Portability_Crewed"] = designTask{"Portability", "Crewed", "C", 0, "", 1.0, 1, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Portability_Fixed"] = designTask{"Portability", "Fixed", "F", 0, "+1", 1.0, 4, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Portability_Portable"] = designTask{"Portability", "Portable", "P", 0, "+1", 1.0, -2, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Portability_Vechicle Mount"] = designTask{"Portability", "Vechicle Mount", "V", 0, "+1", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
	designTaskMap["Portability_Turret"] = designTask{"Portability", "Turret", "T", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
	if val, ok := designTaskMap[key]; ok {
		dt = val
	}
	return dt
}

/*

weaponTypeMap["Long Guns_Rifle"] = designTask{"Long Guns", "Rifle", "R", 5, 5, 4.0, 0, "Bullet", 2, "", 0, "", 0, "2D", 500}
weaponDescriptorMap["Long Guns_Poison Dart-2"] = designTask{"Long Guns", "Poison Dart-2", "P2", 1, 4, 1.0, 0, "", 0, "Poison", 2, "", 0, "2D", 0.9}
																		Sn Snub  TL+1 R=2 Massx0.7 Burden-3  D2 = 1    Cost x 1.5
Improved		TL+1	R+0	Mасса x 1.0	Burden -1 Цена x 1.1 простота обращения "+1", Надежность "+1"

цена = 750 Cr
ТЛ = 8
урон = 2D
Боекомплект: 3 выстрела
трэйты = Zero-G
масса = 2,8 кг
QREBS = 011-30
Better than some, Easy to carry
Модель ImSnRP-8

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

weaponTypeMap := make(map[string]designTask)
weaponTypeMap["Short Blades_Knife"] = designTask{"Short Blades", "Knife", "K", 1, "R", 0.5, 0, "Cuts", 2, "", 0, "", 0, "2D", 50}
weaponTypeMap["Short Blades_Dagger"] = designTask{"Short Blades", "Dagger", "D", 2, "R", 0.5, 0, "Cuts", 2, "", 0, "", 0, "2D", 50}
weaponTypeMap["Short Blades_Trench Knife"] = designTask{"Short Blades", "Trench Knife", "TK", 4, "R", 1.0, 0, "Cuts", 2, "Blow", 1, "", 0, "2D", 100}
weaponTypeMap["Short Blades_Big Knife"] = designTask{"Short Blades", "Big Knife", "BK", 5, "T", 3.0, 0, "Cuts", 2, "Pen", 2, "", 0, "2D", 200}
weaponTypeMap["Short Blades_Great Big Knife"] = designTask{"Short Blades", "Great Big Knife", "GBK", 6, "1", 6.0, 0, "Cuts", 2, "Pen", 2, "", 0, "2D", 900}
weaponTypeMap["Medium Blades_Sword"] = designTask{"Medium Blades", "Sword", "S", 3, "1", 2.0, 0, "Cuts", 2, "", 0, "", 0, "2D", 300}
weaponTypeMap["Medium Blades_Short Sword"] = designTask{"Medium Blades", "Short Sword", "sS", 3, "1", 1.0, -1, "Cuts", 2, "", 0, "", 0, "2D", 300}
weaponTypeMap["Medium Blades_Broadsword"] = designTask{"Medium Blades", "Broadsword", "bS", 4, "1", 3.0, 0, "Cuts", 3, "", 0, "", 0, "3D", 700}
weaponTypeMap["Medium Blades_Cutlass"] = designTask{"Medium Blades", "Cutlass", "C", 3, "1", 2.0, 0, "Cuts", 2, "", 0, "", 0, "2D", 200}
weaponTypeMap["Medium Blades_Officers Cutlass"] = designTask{"Medium Blades", "Officers Cutlass", "OC", 5, "1", 1.0, 0, "Cuts", 2, "", 0, "", 0, "2D", 400}
weaponTypeMap["Long Blades_Pike"] = designTask{"Long Blades", "Pike", "P", 1, "1", 2.0, 3, "Cuts", 2, "", 0, "", 0, "2D", 50}
weaponTypeMap["Special Blades_Axe"] = designTask{"Special Blades", "Axe", "Ax", 2, "T", 2.0, 0, "Cuts", 3, "", 0, "", 0, "2D", 60}
weaponTypeMap["Special Blades_Space Axe"] = designTask{"Special Blades", "Space Axe", "A", 9, "1", 2.0, 0, "Cuts", 2, "Pen", 2, "", 0, "2D", 100}
weaponTypeMap["Special Blades_Vibro-Blade"] = designTask{"Special Blades", "Vibro-Blade", "V", 10, "1", 0.5, 0, "Cuts", 2, "", 0, "", 0, "2D", 900}
weaponTypeMap["Special Blades_Mace"] = designTask{"Special Blades", "Mace", " ", 2, "1", 4.0, 0, "Cuts", 1, "Blow", 2, "", 0, "2D", 100}
weaponTypeMap["Special Blades_Club"] = designTask{"Special Blades", "Club", " ", 1, "1", 2.0, 0, "Blow", 1, "", 0, "", 0, "1D", 10}
//
weaponTypeMap["Body Weapons_Fists"] = designTask{"Body Weapons", "Fists", "Fi", 0, "R", 0.0, 0, "Blow", 1, "", 0, "", 0, "1D", 0}
weaponTypeMap["Body Weapons_Tentacle"] = designTask{"Body Weapons", "Tentacle", "Te", 0, "0", 0.0, 0, "Hit", 1, "Suff", 1, "", 0, "1D", 0}
weaponTypeMap["Body Weapons_Horns"] = designTask{"Body Weapons", "Horns", "Ho", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "2D", 0}
weaponTypeMap["Body Weapons_Tusks"] = designTask{"Body Weapons", "Tusks", "Tu", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "2D", 0}
weaponTypeMap["Body Weapons_Fangs"] = designTask{"Body Weapons", "Fangs", "Fa", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "2D", 0}
weaponTypeMap["Body Weapons_Teeth"] = designTask{"Body Weapons", "Teeth", "T", 0, "R", 0.0, 0, "Cut", 1, "", 0, "", 0, "1D", 0}
weaponTypeMap["Body Weapons_Claws"] = designTask{"Body Weapons", "Claws", "Cl", 0, "R", 0.0, 0, "Cut", 1, "", 0, "", 0, "1D", 0}
weaponTypeMap["Body Weapons_Dew Claws"] = designTask{"Body Weapons", "Dew Claws", "Dc", 0, "R", 0.0, 0, "Cut", 2, "", 0, "", 0, "2D", 0}
weaponTypeMap["Body Weapons_Hooves"] = designTask{"Body Weapons", "Hooves", "H", 0, "R", 0.0, 0, "Blow", 2, "", 0, "", 0, "2D", 0}
weaponTypeMap["Body Weapons_Spikes"] = designTask{"Body Weapons", "Spikes", "Sp", 0, "0", 0.0, 0, "Pen", 2, "", 0, "", 0, "2D", 0}
weaponTypeMap["Body Weapons_Sting"] = designTask{"Body Weapons", "Sting", "St", 0, "R", 0.0, 0, "Pen", 1, "Poison", 2, "", 0, "3D", 0}
//
weaponTypeMap["Artilery_Gun"] = designTask{"Artilery", "Gun", "G", 6, "4", 9.0, 1, "*", 2, "", 0, "", 0, "2D", 5000}
weaponTypeMap["Artilery_Gatling"] = designTask{"Artilery", "Gatling", "Ga", 7, "4", 40.0, 2, "*", 3, "", 0, "", 0, "2D", 8000}
weaponTypeMap["Artilery_Cannon"] = designTask{"Artilery", "Cannon", "C", 6, "6", 200.0, 4, "*", 4, "", 0, "", 0, "2D", 10000}
weaponTypeMap["Artilery_Autocannon"] = designTask{"Artilery", "Autocannon", "aC", 8, "6", 300.0, 4, "*", 5, "", 0, "", 0, "3D", 30000}
weaponTypeMap["Long Guns_Rifle"] = designTask{"Long Guns", "Rifle", "R", 5, "5", 4.0, 0, "Bullet", 2, "", 0, "", 0, "2D", 500}
weaponTypeMap["Long Guns_Carbine"] = designTask{"Long Guns", "Carbine", "C", 5, "4", 3.0, -1, "Bullet", 1, "", 0, "", 0, "1D", 400}
weaponTypeMap["Hand Guns_Pistol"] = designTask{"Hand Guns", "Pistol", "P", 5, "2", 1.1, 0, "Bullet", 1, "", 0, "", 0, "1D", 150}
weaponTypeMap["Hand Guns_Revolver"] = designTask{"Hand Guns", "Revolver", "R", 4, "2", 1.25, 0, "Bullet", 1, "", 0, "", 0, "1D", 100}
weaponTypeMap["Shotguns_Shotgun"] = designTask{"Shotguns", "Shotgun", "S", 4, "2", 4.0, 0, "Frag", 2, "", 0, "", 0, "2D", 300}
weaponTypeMap["Machineguns_Machinegun"] = designTask{"Machineguns", "Machinegun", "Mg", 6, "5", 8.0, 1, "Bullet", 4, "", 0, "", 0, "4D", 3000}
weaponTypeMap["Projectors_Projector"] = designTask{"Projectors", "Projector", "Pj", 9, "0", 1.0, 0, "*", 1, "", 0, "", 0, "1D", 300}
weaponTypeMap["Designators_Designator"] = designTask{"Designators", "Designator", "D", 7, "5", 10.0, 1, "*", 1, "", 0, "", 0, "1D", 2000}
weaponTypeMap["Launchers_Launcher"] = designTask{"Launchers", "Launcher", "L", 6, "3", 10.0, 1, "*", 1, "", 0, "", 0, "0D", 1000}
weaponTypeMap["Launchers_Multi-Launcher"] = designTask{"Launchers", "Multi-Launcher", "mL", 8, "5", 8.0, 1, "*", 1, "", 0, "", 0, "0D", 3000}
if val, ok := weaponTypeMap[key]; ok {
	return val
}
weaponDescriptorMap := make(map[string]designTask)
weaponDescriptorMap["Artilery_(blank)"] = designTask{"Artilery", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
weaponDescriptorMap["Artilery_Anti-Flyer"] = designTask{"Artilery", "Anti-Flyer", "aF", 4, "=6", 6.0, 0, "", 0, "Frag", 1, "Blast", 3, "4D", 3}
weaponDescriptorMap["Artilery_Anti-Tank"] = designTask{"Artilery", "Anti-Tank", "aT", 0, "=5", 8.0, 0, "", 0, "Pen", 3, "Blast", 3, "6D", 2}
weaponDescriptorMap["Artilery_Assault"] = designTask{"Artilery", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "3D", 1.5}
weaponDescriptorMap["Artilery_Fusion"] = designTask{"Artilery", "Fusion", "F", 7, "=4", 2.3, 0, "", 0, "Pen", 4, "Burn", 4, "8D", 6}
weaponDescriptorMap["Artilery_Gauss"] = designTask{"Artilery", "Gauss", "G", 7, "=4", 0.9, 0, "", 0, "Bullet", 3, "", 0, "3D", 2}
weaponDescriptorMap["Artilery_Plasma"] = designTask{"Artilery", "Plasma", "P", 5, "=4", 2.5, 0, "", 0, "Pen", 3, "Burn", 3, "6D", 2}
weaponDescriptorMap["Long Guns_(blank)"] = designTask{"Long Guns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
weaponDescriptorMap["Long Guns_Accelerator"] = designTask{"Long Guns", "Accelerator", "Ac", 4, "", 0.6, 0, "", 0, "Bullet", 2, "", 0, "2D", 3.0}
weaponDescriptorMap["Long Guns_Assault"] = designTask{"Long Guns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Blast", 2, "Bang", 1, "3D", 1.5}
weaponDescriptorMap["Long Guns_Battle"] = designTask{"Long Guns", "Battle", "B", 1, "=5", 1.0, 1, "", 0, "Bullet", 1, "", 0, "1D", 0.8}
weaponDescriptorMap["Long Guns_Combat"] = designTask{"Long Guns", "Combat", "C", 2, "=3", 0.9, 0, "", 0, "Frag", 2, "", 0, "2D", 1.5}
weaponDescriptorMap["Long Guns_Dart-1"] = designTask{"Long Guns", "Dart-1", "D1", 1, "=4", 0.6, 0, "", 0, "Tranq", 1, "", 0, "1D", 0.9}
weaponDescriptorMap["Long Guns_Dart-2"] = designTask{"Long Guns", "Dart-2", "D2", 1, "=4", 0.6, 0, "", 0, "Tranq", 2, "", 0, "2D", 0.9}
weaponDescriptorMap["Long Guns_Dart-3"] = designTask{"Long Guns", "Dart-3", "D3", 1, "=4", 0.6, 0, "", 0, "Tranq", 3, "", 0, "3D", 0.9}
weaponDescriptorMap["Long Guns_Poison Dart-1"] = designTask{"Long Guns", "Poison Dart-1", "P1", 1, "=4", 1.0, 0, "", 0, "Poison", 1, "", 0, "1D", 0.9}
weaponDescriptorMap["Long Guns_Poison Dart-2"] = designTask{"Long Guns", "Poison Dart-2", "P2", 1, "=4", 1.0, 0, "", 0, "Poison", 2, "", 0, "2D", 0.9}
weaponDescriptorMap["Long Guns_Poison Dart-3"] = designTask{"Long Guns", "Poison Dart-3", "P3", 1, "=4", 1.0, 0, "", 0, "Poison", 3, "", 0, "3D", 0.9}
weaponDescriptorMap["Long Guns_Gauss"] = designTask{"Long Guns", "Gauss", "G", 7, "", 0.9, 0, "", 0, "Bullet", 3, "", 0, "3D", 2.0}
weaponDescriptorMap["Long Guns_Hunting"] = designTask{"Long Guns", "Hunting", "H", 0, "=3", 0.9, -1, "", 0, "Bullet", 1, "", 0, "1D", 1.2}
weaponDescriptorMap["Long Guns_Laser"] = designTask{"Long Guns", "Laser", "L", 5, "", 1.2, 0, "", 0, "Burn", 2, "Pen", 2, "4D", 6.0}
weaponDescriptorMap["Long Guns_Splat"] = designTask{"Long Guns", "Splat", "Sp", 2, "=4", 1.3, 1, "", 0, "Bullet", 1, "", 0, "1D", 2.4}
weaponDescriptorMap["Long Guns_Survival"] = designTask{"Long Guns", "Survival", "S", 0, "=2", 0.5, 0, "", 0, "Bullet", 1, "", 0, "1D", 1.2}
weaponDescriptorMap["Handguns_(blank)"] = designTask{"Handguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
weaponDescriptorMap["Handguns_Accelerator"] = designTask{"Handguns", "Accelerator", "Ac", 4, "", 0.6, 0, "", 0, "Bullet", 2, "", 0, "2D", 3.0}
weaponDescriptorMap["Handguns_Laser"] = designTask{"Handguns", "Laser", "L", 5, "", 1.2, 0, "", 0, "Burn", 2, "Pen", 2, "4D", 2.0}
weaponDescriptorMap["Handguns_Machine"] = designTask{"Handguns", "Machine", "M", 0, "=2", 0.5, 0, "", 0, "Bullet", 1, "", 0, "1D", 1.2}
weaponDescriptorMap["Shotguns_(blank)"] = designTask{"Shotguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
weaponDescriptorMap["Shotguns_Assault"] = designTask{"Shotguns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "3D", 2.0}
weaponDescriptorMap["Shotguns_Hunting"] = designTask{"Shotguns", "Hunting", "H", 0, "=3", 0.9, 0, "", 0, "Bullet", 1, "", 0, "1D", 2.0}
weaponDescriptorMap["Machineguns_(blank)"] = designTask{"Machineguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "0D", 1.0}
weaponDescriptorMap["Machineguns_Anti-Flyer"] = designTask{"Machineguns", "Anti-Flyer", "aF", 4, "=6", 6.0, 0, "", 0, "Frag", 1, "Blast", 3, "4D", 3.0}
weaponDescriptorMap["Machineguns_Assault"] = designTask{"Machineguns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "3D", 1.5}
weaponDescriptorMap["Machineguns_Sub"] = designTask{"Machineguns", "Sub", "S", -1, "=3", 0.3, 0, "", 0, "Bullet", -1, "", 0, "-1D", 0.9}
weaponDescriptorMap["Launchers_AF Missile"] = designTask{"Launchers", "AF Missile", "aF", 4, "=7", 4.0, 0, "", 0, "Frag", 2, "Blast", 3, "5D", 3.0}
weaponDescriptorMap["Launchers_AT Missile"] = designTask{"Launchers", "AT Missile", "aT", 3, "=4", 1.0, 1, "", 0, "Frag", 2, "Pen", 3, "5D", 2.0}
weaponDescriptorMap["Launchers_Grenade"] = designTask{"Launchers", "Grenade", "Gr", 1, "=4", 0.8, 0, "", 0, "Frag", 2, "Blast", 2, "4D", 1.0}
weaponDescriptorMap["Launchers_Missile"] = designTask{"Launchers", "Missile", "M", 1, "=6", 2.2, 0, "", 0, "Frag", 2, "Pen", 2, "4D", 5.0}
weaponDescriptorMap["Launchers_RAM Grenade"] = designTask{"Launchers", "RAM Grenade", "RAM", 2, "=6", 1.0, 0, "", 0, "Frag", 2, "Blast", 2, "4D", 3.0}
weaponDescriptorMap["Launchers_Rocket"] = designTask{"Launchers", "Rocket", "R", -1, "=5", 3.0, 0, "", 0, "Frag", 2, "Pen", 2, "4D", 1.0}
burdenDescriptorMap := make(map[string]designTask)
burdenDescriptorMap["Burden_(blank)"] = designTask{"Burden", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
burdenDescriptorMap["Burden_Anti-Designator"] = designTask{"Burden", "Anti-Designator", "aD", 3, "+1", 3.0, 3, "", 0, "", 1, "", 0, "Not Pistol. Not Shotgun", 3.0}
burdenDescriptorMap["Burden_Body"] = designTask{"Burden", "Body", "B", 2, "=1", 0.5, -4, "", 0, "", -1, "", 0, "Only Pistol", 3.0}
burdenDescriptorMap["Burden_Disposable"] = designTask{"Burden", "Disposable", "D", 3, "+0", 0.9, -1, "", 0, "", 0, "", 0, "Q=-2", 0.5}
burdenDescriptorMap["Burden_Heavy"] = designTask{"Burden", "Heavy", "H", 0, "+1", 1.3, 3, "", 0, "", 1, "", 0, "Not Laser", 2.0}
burdenDescriptorMap["Burden_Light"] = designTask{"Burden", "Light", "Lt", 0, "-1", 0.7, -1, "", 0, "", -1, "", 0, "Not Laser", 1.1}
burdenDescriptorMap["Burden_Magnum"] = designTask{"Burden", "Magnum", "M", 1, "+1", 1.1, 1, "", 0, "", 1, "", 0, "Only Pistol", 1.1}
burdenDescriptorMap["Burden_Medium"] = designTask{"Burden", "Medium", "M", 0, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
burdenDescriptorMap["Burden_Recoilless"] = designTask{"Burden", "Recoilless", "R", 1, "+0", 1.2, 0, "", 0, "", 1, "", 0, "", 3.0} // TODO: Check Errata for Range (was -+1, made +0)
burdenDescriptorMap["Burden_Snub"] = designTask{"Burden", "Snub", "Sn", 1, "=2", 0.7, -3, "", 0, "", 1, "", 0, "", 1.5}
burdenDescriptorMap["Burden_Vheavy"] = designTask{"Burden", "Vheavy", "Vh", 0, "+5", 4.0, 4, "", 0, "", 5, "", 0, "", 5.0}
burdenDescriptorMap["Burden_Vlight"] = designTask{"Burden", "Vlight", "Vl", 1, "-2", 0.6, -2, "", 0, "", -1, "", 0, "", 2.0}
burdenDescriptorMap["Burden_VRF"] = designTask{"Burden", "VRF", "Vrf", 2, "+0", 14.0, 5, "", 0, "", 1, "", 0, "Only Gun, Only Machnegun", 9.0}
stageDescriptorMap := make(map[string]designTask)
stageDescriptorMap["Stage_(blank)"] = designTask{"Stage", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
stageDescriptorMap["Stage_Advanced"] = designTask{"Stage", "Advanced", "A", 3, "+0", 0.8, -3, "", 0, "", 2, "", 0, "", 2.0}
stageDescriptorMap["Stage_Alternate"] = designTask{"Stage", "Alternate", "Alt", 0, "+1", 1.1, 0, "", 0, "", 2, "", 0, "B+Flux", 1.1}
stageDescriptorMap["Stage_Basic"] = designTask{"Stage", "Basic", "B", 0, "+0", 1.3, 1, "", 0, "", 0, "", 0, "", 0.7}
stageDescriptorMap["Stage_Early"] = designTask{"Stage", "Early", "E", -1, "-1", 1.7, 1, "", 0, "", 0, "", 0, "EOU-1", 0.7}
stageDescriptorMap["Stage_Experimental"] = designTask{"Stage", "Experimental", "Exp", -3, "-1", 2.0, 3, "", 0, "", 0, "", 0, "R=-2", 4.0}
stageDescriptorMap["Stage_Generic"] = designTask{"Stage", "Generic", "Gen", 1, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 0.5}
stageDescriptorMap["Stage_Improved"] = designTask{"Stage", "Improved", "Im", 1, "+0", 1.0, -1, "", 0, "", 1, "", 0, "EOU+1", 1.1}
stageDescriptorMap["Stage_Modified"] = designTask{"Stage", "Modified", "Mod", 2, "+0", 0.9, 0, "", 0, "", 1, "", 0, "", 1.2}
stageDescriptorMap["Stage_Precision"] = designTask{"Stage", "Precision", "Pr", 6, "+3", 4.0, 2, "", 0, "", 0, "", 0, "Only Designator", 5.0}
stageDescriptorMap["Stage_Prototype"] = designTask{"Stage", "Prototype", "P", -2, "-1", 1.9, 2, "", 0, "", 0, "", 0, "", 3.0}
stageDescriptorMap["Stage_Remote"] = designTask{"Stage", "Remote", "R", 1, "+0", 1.0, 0, "", 0, "", 0, "", 0, "Not Pistol", 7.0}
stageDescriptorMap["Stage_Sniper"] = designTask{"Stage", "Sniper", "Sn", 1, "+1", 1.1, 1, "", 0, "", 1, "", 0, "Only Rifle, Q=+2", 2.0} //TODO: Check Errata for +1 on D2 for stantard (makes no sense if it is in standard and not Sniper)
stageDescriptorMap["Stage_Standard"] = designTask{"Stage", "Standard", "St", 0, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}             //TODO: Check Errata for +1 on D2 for stantard (makes no sense if it is in standard and not Sniper)
stageDescriptorMap["Stage_Target"] = designTask{"Stage", "Target", "T", 0, "+0", 1.1, 1, "", 0, "", 0, "", 0, "Only Rifle, Only Pistol, Q=+2", 1.5}
stageDescriptorMap["Stage_Ultimate"] = designTask{"Stage", "Ultimate", "Ul", 4, "+0", 0.7, -4, "", 0, "", 2, "", 0, "R=+4", 1.4}
usersDescriptorMap := make(map[string]designTask)
usersDescriptorMap["Users_(blank)"] = designTask{"Users", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU+0", 1.0}
usersDescriptorMap["Users_Universal"] = designTask{"Users", "Universal", "U", 0, "", 1.1, 1, "", 0, "", 0, "", 0, "EOU-1", 1.0}
usersDescriptorMap["Users_Man"] = designTask{"Users", "Man", "M", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU+0", 1.0}
usersDescriptorMap["Users_Vargr"] = designTask{"Users", "Vargr", "V", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0}
usersDescriptorMap["Users_K'kree"] = designTask{"Users", "K'kree", "K", 0, "", 1.3, 2, "", 0, "", 0, "", 0, "EOU+0", 1.0}
usersDescriptorMap["Users_Grasper (Hiver)"] = designTask{"Users", "Grasper (Hiver)", "H", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0}
usersDescriptorMap["Users_Paw (Aslan)"] = designTask{"Users", "Paw (Aslan)", "P", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0}
usersDescriptorMap["Users_Gripper"] = designTask{"Users", "Gripper", "G", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0}
usersDescriptorMap["Users_Tentacle (Vegan)"] = designTask{"Users", "Tentacle (Vegan)", "T", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0}
usersDescriptorMap["Users_Socket"] = designTask{"Users", "Socket", "S", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0}
portabilityDescriptorMap := make(map[string]designTask)
portabilityDescriptorMap["Portability_(blank)"] = designTask{"Portability", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
portabilityDescriptorMap["Portability_Crewed"] = designTask{"Portability", "Crewed", "C", 0, "", 1.0, 1, "", 0, "", 0, "", 0, "", 1.0}
portabilityDescriptorMap["Portability_Fixed"] = designTask{"Portability", "Fixed", "F", 0, "+1", 1.0, 4, "", 0, "", 0, "", 0, "", 1.0}
portabilityDescriptorMap["Portability_Portable"] = designTask{"Portability", "Portable", "P", 0, "+1", 1.0, -2, "", 0, "", 0, "", 0, "", 1.0}
portabilityDescriptorMap["Portability_Vechicle Mount"] = designTask{"Portability", "Vechicle Mount", "V", 0, "+1", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}
portabilityDescriptorMap["Portability_Turret"] = designTask{"Portability", "Turret", "T", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0}




*/

func recalculateRange(initial int, new string) int {
	if new == "R" || new == "T" {
		return 1
	}
	if strings.Contains(new, "=") {
		new = strings.TrimPrefix(new, "=")
		newInt, err := strconv.Atoi(new)
		if err != nil {
			panic("recalculateRange: " + err.Error())
		}
		return newInt
	}
	newInt, err := strconv.Atoi(new)
	if err != nil {
		panic("recalculateRange: " + err.Error())
	}
	return initial + newInt
}
