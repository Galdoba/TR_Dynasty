package weapons

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type designTask struct {
	category    string
	descriptor  string
	dModel      string //code
	dTL         int    //
	dRange      string //
	dMass       float64
	dBurden     int
	dH1         string
	dD1         int
	dH2         string
	dD2         int
	dH3         string
	dD3         int
	misc        string
	dCost       float64
	designStage int
}

type weapon struct {
	wType        string
	wSubType     string
	wDescriptor  string
	wBurden      string
	wStage       string
	wUser        string
	wPortability string
	wModel       [7]string
	qrebsStruct  qrebs
	tl           int
	wRange       int
	mass         float64
	damageDice   int
	cost         float64
	traits       string
	description  string
}

func TestGunMaker() {
	gMaker := NewGunMaker()
	wp := gMaker.MakeWeapon(
		gMaker.Apply("Handguns", "Pistol"),
		//gMaker.Apply("Shotguns", "Assault"),
		gMaker.Apply("Burden", "Snub"),
		//gMaker.Apply("Stage", "Advanced"),
		gMaker.Apply("Users", "(blank)"),
	)
	fmt.Println(wp, gMaker)
}

func (gmkr *gunMaker) MakeWeapon(dT ...designTask) weapon {
	wp := weapon{}
NextTask:
	for _, designTask := range dT {
		for _, concluded := range gmkr.designsConcluded {
			if concluded == designTask.designStage {
				continue NextTask
			}
		}
		wp.absorb(designTask)
		gmkr.designsConcluded = append(gmkr.designsConcluded, designTask.designStage)
	}
	wp.setPortability()
	wp.cleanData()
	return wp
}

func (wp *weapon) absorb(dt designTask) {
	err := errors.New("No Error")

	switch dt.designStage {
	case 0:
		if wp.wSubType != "" {
			return
		}
		fmt.Println("DEBUG: initial")
		wp.setInitialStats(dt)
	case 1:
		fmt.Println("DEBUG: Descriptor")
		wp.setDescriptor(dt)
	case 2:
		fmt.Println("DEBUG: Burden")
		wp.setBurden(dt)
	case 3:
		fmt.Println("DEBUG: Stage")
		wp.setStage(dt)
	case 4:
		fmt.Println("DEBUG: Users")
		wp.setUsers(dt)
	}

	if dt.category == "" {
		//apply only after checking dt.Misc
		return
	}
	//apply only if wp.wType == dt.category
	if wp.wType == dt.category {
		//modyfy stats
		//wp.modify(dt)
	}
	fmt.Println(err)
	fmt.Println("End 'absorb'")
}

func (wp *weapon) setInitialStats(dt designTask) {
	wp.wType = dt.category
	wp.wSubType = dt.descriptor
	wp.wRange = recalculateRange(wp.wRange, dt.dRange)
	wp.damageDice = dt.dD1 + dt.dD2 + dt.dD3
	wp.traits = dt.dH1 + dt.dH2 + dt.dH3
	wp.mass = dt.dMass
	wp.cost = dt.dCost
	wp.tl = dt.dTL
	wp.qrebsStruct = qrebs{0, 0, 0, dt.dBurden, 0}
	wp.wModel[3] = dt.dModel
}

func (wp *weapon) setDescriptor(dt designTask) {
	if wp.wDescriptor != "" {
		return
	}
	wp.wDescriptor = dt.descriptor
	wp.tl = wp.tl + dt.dTL
	wp.wRange = recalculateRange(wp.wRange, dt.dRange)
	wp.mass = wp.mass * dt.dMass
	wp.qrebsStruct.burden = wp.qrebsStruct.burden + dt.dBurden
	if dt.misc == "B+Flux" {
		wp.qrebsStruct.burden = wp.qrebsStruct.burden + dice.Flux()
	}
	if dt.descriptor == "Laser" {
		if strings.Contains(wp.traits, "Bullet") {
			wp.traits = ""
			wp.damageDice = 0
		}
	}
	wp.damageDice = wp.damageDice + dt.dD1 + dt.dD2 + dt.dD3
	wp.traits = wp.traits + dt.dH1 + dt.dH2 + dt.dH3
	wp.cost = wp.cost * dt.dCost
	wp.wModel[2] = dt.dModel
}

func (wp *weapon) setBurden(dt designTask) {
	if wp.wBurden != "" {
		return
	}
	switch dt.misc {
	case "Q=-2":
		wp.qrebsStruct.quality = -2
	case "Not Pistol, Not Shotgun":
		if wp.wSubType == "Pistol" || wp.wSubType == "Shotgun" {
			return
		}
	case "Only Pistol":
		if wp.wSubType != "Pistol" {
			return
		}
	case "Not Laser":
		if wp.wDescriptor == "Laser" {
			return
		}
	case "Only Gun, Only Machinegun":
		if wp.wSubType != "Gun" && wp.wSubType != "Machinegun" {
			return
		}
	}
	fmt.Println("DEBUG: Burden Approved")
	wp.wBurden = dt.descriptor
	wp.tl = wp.tl + dt.dTL
	wp.wRange = recalculateRange(wp.wRange, dt.dRange)
	wp.mass = wp.mass * dt.dMass
	wp.qrebsStruct.burden = wp.qrebsStruct.burden + dt.dBurden
	wp.damageDice = wp.damageDice + dt.dD1 + dt.dD2 + dt.dD3
	wp.traits = wp.traits + dt.dH1 + dt.dH2 + dt.dH3
	wp.cost = wp.cost * dt.dCost
	wp.wModel[1] = dt.dModel
}

func (wp *weapon) setStage(dt designTask) {
	if wp.wStage != "" {
		return
	}
	switch dt.misc {
	case "EOU-1":
		wp.qrebsStruct.easeOfUse = wp.qrebsStruct.easeOfUse - 1
	case "EOU+1":
		wp.qrebsStruct.easeOfUse = wp.qrebsStruct.easeOfUse + 1
	case "R=+1":
		wp.qrebsStruct.reliability = wp.qrebsStruct.reliability + 1
	case "R=-2":
		wp.qrebsStruct.reliability = wp.qrebsStruct.reliability - 2
	case "Only Designator":
		if wp.wSubType != "Designator" {
			return
		}
	case "Not Pistol":
		if wp.wSubType == "Pistol" {
			return
		}
	case "Only Rifle, Q=+2":
		if wp.wSubType != "Rifle" {
			return
		}
		wp.qrebsStruct.quality = wp.qrebsStruct.quality + 2
	case "Only Rifle, Only Pistol, Q=+2":
		if wp.wSubType != "Rifle" && wp.wSubType != "Pistol" {
			return
		}
		wp.qrebsStruct.quality = 2
	case "R=+4":
		wp.qrebsStruct.reliability = wp.qrebsStruct.reliability + 4
	}
	fmt.Println("DEBUG: Stage Approved")
	wp.wStage = dt.descriptor
	wp.tl = wp.tl + dt.dTL
	wp.wRange = recalculateRange(wp.wRange, dt.dRange)
	wp.mass = wp.mass * dt.dMass
	wp.qrebsStruct.burden = wp.qrebsStruct.burden + dt.dBurden
	wp.damageDice = wp.damageDice + dt.dD1 + dt.dD2 + dt.dD3
	wp.traits = wp.traits + dt.dH1 + dt.dH2 + dt.dH3
	wp.cost = wp.cost * dt.dCost
	wp.wModel[0] = dt.dModel
}

func (wp *weapon) setUsers(dt designTask) {
	if wp.wUser != "" {
		return
	}
	switch dt.misc {
	case "EOU-2":
		wp.qrebsStruct.easeOfUse = wp.qrebsStruct.easeOfUse - 2
	case "EOU+2":
		wp.qrebsStruct.easeOfUse = wp.qrebsStruct.easeOfUse + 2
	case "EOU-1":
		wp.qrebsStruct.easeOfUse = wp.qrebsStruct.easeOfUse - 1
	case "EOU+1":
		wp.qrebsStruct.easeOfUse = wp.qrebsStruct.easeOfUse + 1
	}
	fmt.Println("DEBUG: Users Approved")
	wp.wUser = dt.descriptor
	wp.tl = wp.tl + dt.dTL
	wp.wRange = recalculateRange(wp.wRange, dt.dRange)
	wp.mass = wp.mass * dt.dMass
	wp.qrebsStruct.burden = wp.qrebsStruct.burden + dt.dBurden
	wp.damageDice = wp.damageDice + dt.dD1 + dt.dD2 + dt.dD3
	wp.traits = wp.traits + dt.dH1 + dt.dH2 + dt.dH3
	wp.cost = wp.cost * dt.dCost
	wp.wModel[4] = dt.dModel
}

func (wp *weapon) setPortability() {
	key := "Portability_(blank)"
	if wp.mass >= 20 {
		key = "Portability_Portable"
	}
	if wp.mass > 40 {
		key = "Portability_Crewed"
	}
	if wp.mass > 200 {
		key = "Portability_Turret"
	}
	if wp.mass > 500 {
		key = "Portability_Vechicle Mount"
	}
	if wp.mass > 1000 {
		key = "Portability_Fixed"
	}
	dt := callDesignTask(key)
	fmt.Println("DEBUG:", key)
	wp.wPortability = dt.descriptor
	wp.tl = wp.tl + dt.dTL
	wp.wRange = recalculateRange(wp.wRange, dt.dRange)
	wp.mass = wp.mass * dt.dMass
	wp.qrebsStruct.burden = wp.qrebsStruct.burden + dt.dBurden
	wp.damageDice = wp.damageDice + dt.dD1 + dt.dD2 + dt.dD3
	wp.traits = wp.traits + dt.dH1 + dt.dH2 + dt.dH3
	wp.cost = wp.cost * dt.dCost
	wp.wModel[5] = dt.dModel
}

func (wp *weapon) cleanData() {
	if wp.wDescriptor == "(blank)" {
		wp.wDescriptor = ""
	}
	if wp.wBurden == "(blank)" {
		wp.wBurden = ""
	}
	if wp.wStage == "(blank)" {
		wp.wStage = ""
	}
	if wp.wPortability == "(blank)" {
		wp.wPortability = ""
	}
	wp.mass = math.Round(wp.mass*1000) / 1000
	wp.cost = math.Round(wp.cost)
}

func NewGunMaker(seed ...int64) gunMaker {
	gMaker := gunMaker{}
	if len(seed) == 0 {
		gMaker.dp = dice.New().SetSeed(time.Now().UnixNano())
	} else {
		gMaker.dp = dice.New().SetSeed(seed[0])
	}
	return gMaker
}

type gunMaker struct {
	dp               *dice.Dicepool //дайспул который используется для бросков
	designsConcluded []int          //список совершенных типов операций
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
	dt := designTask{"", "", "", 0, "0", 0.0, 0, "", 0, "", 0, "", 0, "", 1, -1}
	designTaskMap := make(map[string]designTask)
	designTaskMap["Short Blades_Knife"] = designTask{"Short Blades", "Knife", "K", 1, "R", 0.5, 0, "Cuts", 2, "", 0, "", 0, "Init", 50, 0}
	designTaskMap["Short Blades_Dagger"] = designTask{"Short Blades", "Dagger", "D", 2, "R", 0.5, 0, "Cuts", 2, "", 0, "", 0, "Init", 50, 0}
	designTaskMap["Short Blades_Trench Knife"] = designTask{"Short Blades", "Trench Knife", "TK", 4, "R", 1.0, 0, "Cuts", 2, "Blow", 1, "", 0, "Init", 100, 0}
	designTaskMap["Short Blades_Big Knife"] = designTask{"Short Blades", "Big Knife", "BK", 5, "T", 3.0, 0, "Cuts", 2, "Pen", 2, "", 0, "Init", 200, 0}
	designTaskMap["Short Blades_Great Big Knife"] = designTask{"Short Blades", "Great Big Knife", "GBK", 6, "1", 6.0, 0, "Cuts", 2, "Pen", 2, "", 0, "Init", 900, 0}
	designTaskMap["Medium Blades_Sword"] = designTask{"Medium Blades", "Sword", "S", 3, "1", 2.0, 0, "Cuts", 2, "", 0, "", 0, "Init", 300, 0}
	designTaskMap["Medium Blades_Short Sword"] = designTask{"Medium Blades", "Short Sword", "sS", 3, "1", 1.0, -1, "Cuts", 2, "", 0, "", 0, "Init", 300, 0}
	designTaskMap["Medium Blades_Broadsword"] = designTask{"Medium Blades", "Broadsword", "bS", 4, "1", 3.0, 0, "Cuts", 3, "", 0, "", 0, "Init", 700, 0}
	designTaskMap["Medium Blades_Cutlass"] = designTask{"Medium Blades", "Cutlass", "C", 3, "1", 2.0, 0, "Cuts", 2, "", 0, "", 0, "2D", 200, 0}
	designTaskMap["Medium Blades_Officers Cutlass"] = designTask{"Medium Blades", "Officers Cutlass", "OC", 5, "1", 1.0, 0, "Cuts", 2, "", 0, "", 0, "Init", 400, 0}
	designTaskMap["Long Blades_Pike"] = designTask{"Long Blades", "Pike", "P", 1, "1", 2.0, 3, "Cuts", 2, "", 0, "", 0, "Init", 50, 0}
	designTaskMap["Special Blades_Axe"] = designTask{"Special Blades", "Axe", "Ax", 2, "T", 2.0, 0, "Cuts", 3, "", 0, "", 0, "Init", 60, 0}
	designTaskMap["Special Blades_Space Axe"] = designTask{"Special Blades", "Space Axe", "A", 9, "1", 2.0, 0, "Cuts", 2, "Pen", 2, "", 0, "Init", 100, 0}
	designTaskMap["Special Blades_Vibro-Blade"] = designTask{"Special Blades", "Vibro-Blade", "V", 10, "1", 0.5, 0, "Cuts", 2, "", 0, "", 0, "Init", 900, 0}
	designTaskMap["Special Blades_Mace"] = designTask{"Special Blades", "Mace", " ", 2, "1", 4.0, 0, "Cuts", 1, "Blow", 2, "", 0, "Init", 100, 0}
	designTaskMap["Special Blades_Club"] = designTask{"Special Blades", "Club", " ", 1, "1", 2.0, 0, "Blow", 1, "", 0, "", 0, "Init", 10, 0}
	designTaskMap["Body Weapons_Fists"] = designTask{"Body Weapons", "Fists", "Fi", 0, "R", 0.0, 0, "Blow", 1, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Tentacle"] = designTask{"Body Weapons", "Tentacle", "Te", 0, "0", 0.0, 0, "Hit", 1, "Suff", 1, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Horns"] = designTask{"Body Weapons", "Horns", "Ho", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Tusks"] = designTask{"Body Weapons", "Tusks", "Tu", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Fangs"] = designTask{"Body Weapons", "Fangs", "Fa", 0, "R", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Teeth"] = designTask{"Body Weapons", "Teeth", "T", 0, "R", 0.0, 0, "Cut", 1, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Claws"] = designTask{"Body Weapons", "Claws", "Cl", 0, "R", 0.0, 0, "Cut", 1, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Dew Claws"] = designTask{"Body Weapons", "Dew Claws", "Dc", 0, "R", 0.0, 0, "Cut", 2, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Hooves"] = designTask{"Body Weapons", "Hooves", "H", 0, "R", 0.0, 0, "Blow", 2, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Spikes"] = designTask{"Body Weapons", "Spikes", "Sp", 0, "0", 0.0, 0, "Pen", 2, "", 0, "", 0, "Init", 0, 0}
	designTaskMap["Body Weapons_Sting"] = designTask{"Body Weapons", "Sting", "St", 0, "R", 0.0, 0, "Pen", 1, "Poison", 2, "", 0, "Init", 0, 0}
	designTaskMap["Artilery_Gun"] = designTask{"Artilery", "Gun", "G", 6, "4", 9.0, 1, "*", 2, "", 0, "", 0, "Init", 5000, 0}
	designTaskMap["Artilery_Gatling"] = designTask{"Artilery", "Gatling", "Ga", 7, "4", 40.0, 2, "*", 3, "", 0, "", 0, "Init", 8000, 0}
	designTaskMap["Artilery_Cannon"] = designTask{"Artilery", "Cannon", "C", 6, "6", 200.0, 4, "*", 4, "", 0, "", 0, "Init", 10000, 0}
	designTaskMap["Artilery_Autocannon"] = designTask{"Artilery", "Autocannon", "aC", 8, "6", 300.0, 4, "*", 5, "", 0, "", 0, "Init", 30000, 0}
	designTaskMap["Long Guns_Rifle"] = designTask{"Long Guns", "Rifle", "R", 5, "5", 4.0, 0, "Bullet", 2, "", 0, "", 0, "Init", 500, 0}
	designTaskMap["Long Guns_Carbine"] = designTask{"Long Guns", "Carbine", "C", 5, "4", 3.0, -1, "Bullet", 1, "", 0, "", 0, "Init", 400, 0}
	designTaskMap["Handguns_Pistol"] = designTask{"Handguns", "Pistol", "P", 5, "2", 1.1, 0, "Bullet", 1, "", 0, "", 0, "Init", 150, 0}
	designTaskMap["Handguns_Revolver"] = designTask{"Handguns", "Revolver", "R", 4, "2", 1.25, 0, "Bullet", 1, "", 0, "", 0, "Init", 100, 0}
	designTaskMap["Shotguns_Shotgun"] = designTask{"Shotguns", "Shotgun", "S", 4, "2", 4.0, 0, "Frag", 2, "", 0, "", 0, "Init", 300, 0}
	designTaskMap["Machineguns_Machinegun"] = designTask{"Machineguns", "Machinegun", "Mg", 6, "5", 8.0, 1, "Bullet", 4, "", 0, "", 0, "Init", 3000, 0}
	designTaskMap["Projectors_Projector"] = designTask{"Projectors", "Projector", "Pj", 9, "0", 1.0, 0, "*", 1, "", 0, "", 0, "Init", 300, 0}
	designTaskMap["Designators_Designator"] = designTask{"Designators", "Designator", "D", 7, "5", 10.0, 1, "*", 1, "", 0, "", 0, "Init", 2000, 0}
	designTaskMap["Launchers_Launcher"] = designTask{"Launchers", "Launcher", "L", 6, "3", 10.0, 1, "*", 1, "", 0, "", 0, "Init", 1000, 0}
	designTaskMap["Launchers_Multi-Launcher"] = designTask{"Launchers", "Multi-Launcher", "mL", 8, "5", 8.0, 1, "*", 1, "", 0, "", 0, "Init", 3000, 0}
	designTaskMap["Artilery_(blank)"] = designTask{"Artilery", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "Descr", 1.0, 1}
	designTaskMap["Artilery_Anti-Flyer"] = designTask{"Artilery", "Anti-Flyer", "aF", 4, "=6", 6.0, 0, "", 0, "Frag", 1, "Blast", 3, "Descr", 3, 1}
	designTaskMap["Artilery_Anti-Tank"] = designTask{"Artilery", "Anti-Tank", "aT", 0, "=5", 8.0, 0, "", 0, "Pen", 3, "Blast", 3, "Descr", 2, 1}
	designTaskMap["Artilery_Assault"] = designTask{"Artilery", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "Descr", 1.5, 1}
	designTaskMap["Artilery_Fusion"] = designTask{"Artilery", "Fusion", "F", 7, "=4", 2.3, 0, "", 0, "Pen", 4, "Burn", 4, "Descr", 6, 1}
	designTaskMap["Artilery_Gauss"] = designTask{"Artilery", "Gauss", "G", 7, "=4", 0.9, 0, "", 0, "Bullet", 3, "", 0, "Descr", 2, 1}
	designTaskMap["Artilery_Plasma"] = designTask{"Artilery", "Plasma", "P", 5, "=4", 2.5, 0, "", 0, "Pen", 3, "Burn", 3, "Descr", 2, 1}
	designTaskMap["Long Guns_(blank)"] = designTask{"Long Guns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "Descr", 1.0, 1}
	designTaskMap["Long Guns_Accelerator"] = designTask{"Long Guns", "Accelerator", "Ac", 4, "", 0.6, 0, "", 0, "Bullet", 2, "", 0, "Descr", 3.0, 1}
	designTaskMap["Long Guns_Assault"] = designTask{"Long Guns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Blast", 2, "Bang", 1, "Descr", 1.5, 1}
	designTaskMap["Long Guns_Battle"] = designTask{"Long Guns", "Battle", "B", 1, "=5", 1.0, 1, "", 0, "Bullet", 1, "", 0, "Descr", 0.8, 1}
	designTaskMap["Long Guns_Combat"] = designTask{"Long Guns", "Combat", "C", 2, "=3", 0.9, 0, "", 0, "Frag", 2, "", 0, "Descr", 1.5, 1}
	designTaskMap["Long Guns_Dart-1"] = designTask{"Long Guns", "Dart-1", "D1", 1, "=4", 0.6, 0, "", 0, "Tranq", 1, "", 0, "Descr", 0.9, 1}
	designTaskMap["Long Guns_Dart-2"] = designTask{"Long Guns", "Dart-2", "D2", 1, "=4", 0.6, 0, "", 0, "Tranq", 2, "", 0, "Descr", 0.9, 1}
	designTaskMap["Long Guns_Dart-3"] = designTask{"Long Guns", "Dart-3", "D3", 1, "=4", 0.6, 0, "", 0, "Tranq", 3, "", 0, "Descr", 0.9, 1}
	designTaskMap["Long Guns_Poison Dart-1"] = designTask{"Long Guns", "Poison Dart-1", "P1", 1, "=4", 1.0, 0, "", 0, "Poison", 1, "", 0, "Descr", 0.9, 1}
	designTaskMap["Long Guns_Poison Dart-2"] = designTask{"Long Guns", "Poison Dart-2", "P2", 1, "=4", 1.0, 0, "", 0, "Poison", 2, "", 0, "Descr", 0.9, 1}
	designTaskMap["Long Guns_Poison Dart-3"] = designTask{"Long Guns", "Poison Dart-3", "P3", 1, "=4", 1.0, 0, "", 0, "Poison", 3, "", 0, "Descr", 0.9, 1}
	designTaskMap["Long Guns_Gauss"] = designTask{"Long Guns", "Gauss", "G", 7, "", 0.9, 0, "", 0, "Bullet", 3, "", 0, "Descr", 2.0, 1}
	designTaskMap["Long Guns_Hunting"] = designTask{"Long Guns", "Hunting", "H", 0, "=3", 0.9, -1, "", 0, "Bullet", 1, "", 0, "Descr", 1.2, 1}
	designTaskMap["Long Guns_Laser"] = designTask{"Long Guns", "Laser", "L", 5, "", 1.2, 0, "", 0, "Burn", 2, "Pen", 2, "Descr", 6.0, 1}
	designTaskMap["Long Guns_Splat"] = designTask{"Long Guns", "Splat", "Sp", 2, "=4", 1.3, 1, "", 0, "Bullet", 1, "", 0, "Descr", 2.4, 1}
	designTaskMap["Long Guns_Survival"] = designTask{"Long Guns", "Survival", "S", 0, "=2", 0.5, 0, "", 0, "Bullet", 1, "", 0, "Descr", 1.2, 1}
	designTaskMap["Handguns_(blank)"] = designTask{"Handguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "Descr", 1.0, 1}
	designTaskMap["Handguns_Accelerator"] = designTask{"Handguns", "Accelerator", "Ac", 4, "", 0.6, 0, "", 0, "Bullet", 2, "", 0, "Descr", 3.0, 1}
	designTaskMap["Handguns_Laser"] = designTask{"Handguns", "Laser", "L", 5, "", 1.2, 0, "", 0, "Burn", 2, "Pen", 2, "Descr", 2.0, 1}
	designTaskMap["Handguns_Machine"] = designTask{"Handguns", "Machine", "M", 0, "=2", 0.5, 0, "", 0, "Bullet", 1, "", 0, "Descr", 1.2, 1}
	designTaskMap["Shotguns_(blank)"] = designTask{"Shotguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "Descr", 1.0, 1}
	designTaskMap["Shotguns_Assault"] = designTask{"Shotguns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "Descr", 2.0, 1}
	designTaskMap["Shotguns_Hunting"] = designTask{"Shotguns", "Hunting", "H", 0, "=3", 0.9, 0, "", 0, "Bullet", 1, "", 0, "Descr", 2.0, 1}
	designTaskMap["Machineguns_(blank)"] = designTask{"Machineguns", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "Descr", 1.0, 1}
	designTaskMap["Machineguns_Anti-Flyer"] = designTask{"Machineguns", "Anti-Flyer", "aF", 4, "=6", 6.0, 0, "", 0, "Frag", 1, "Blast", 3, "Descr", 3.0, 1}
	designTaskMap["Machineguns_Assault"] = designTask{"Machineguns", "Assault", "A", 2, "=4", 0.8, 0, "", 0, "Bang", 1, "Blast", 2, "Descr", 1.5, 1}
	designTaskMap["Machineguns_Sub"] = designTask{"Machineguns", "Sub", "S", -1, "=3", 0.3, 0, "", 0, "Bullet", -1, "", 0, "Descr", 0.9, 1}
	designTaskMap["Launchers_AF Missile"] = designTask{"Launchers", "AF Missile", "aF", 4, "=7", 4.0, 0, "", 0, "Frag", 2, "Blast", 3, "Descr", 3.0, 1}
	designTaskMap["Launchers_AT Missile"] = designTask{"Launchers", "AT Missile", "aT", 3, "=4", 1.0, 1, "", 0, "Frag", 2, "Pen", 3, "Descr", 2.0, 1}
	designTaskMap["Launchers_Grenade"] = designTask{"Launchers", "Grenade", "Gr", 1, "=4", 0.8, 0, "", 0, "Frag", 2, "Blast", 2, "Descr", 1.0, 1}
	designTaskMap["Launchers_Missile"] = designTask{"Launchers", "Missile", "M", 1, "=6", 2.2, 0, "", 0, "Frag", 2, "Pen", 2, "Descr", 5.0, 1}
	designTaskMap["Launchers_RAM Grenade"] = designTask{"Launchers", "RAM Grenade", "RAM", 2, "=6", 1.0, 0, "", 0, "Frag", 2, "Blast", 2, "Descr", 3.0, 1}
	designTaskMap["Launchers_Rocket"] = designTask{"Launchers", "Rocket", "R", -1, "=5", 3.0, 0, "", 0, "Frag", 2, "Pen", 2, "Descr", 1.0, 1}
	designTaskMap["Burden_(blank)"] = designTask{"Burden", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 2}
	designTaskMap["Burden_Anti-Designator"] = designTask{"Burden", "Anti-Designator", "aD", 3, "+1", 3.0, 3, "", 0, "", 1, "", 0, "Not Pistol. Not Shotgun", 3.0, 2}
	designTaskMap["Burden_Body"] = designTask{"Burden", "Body", "B", 2, "=1", 0.5, -4, "", 0, "", -1, "", 0, "Only Pistol", 3.0, 2}
	designTaskMap["Burden_Disposable"] = designTask{"Burden", "Disposable", "D", 3, "+0", 0.9, -1, "", 0, "", 0, "", 0, "Q=-2", 0.5, 2}
	designTaskMap["Burden_Heavy"] = designTask{"Burden", "Heavy", "H", 0, "+1", 1.3, 3, "", 0, "", 1, "", 0, "Not Laser", 2.0, 2}
	designTaskMap["Burden_Light"] = designTask{"Burden", "Light", "Lt", 0, "-1", 0.7, -1, "", 0, "", -1, "", 0, "Not Laser", 1.1, 2}
	designTaskMap["Burden_Magnum"] = designTask{"Burden", "Magnum", "M", 1, "+1", 1.1, 1, "", 0, "", 1, "", 0, "Only Pistol", 1.1, 2}
	designTaskMap["Burden_Medium"] = designTask{"Burden", "Medium", "M", 0, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 2}
	designTaskMap["Burden_Recoilless"] = designTask{"Burden", "Recoilless", "R", 1, "+0", 1.2, 0, "", 0, "", 1, "", 0, "", 3.0, 2} // TODO: Check Errata for Range (was -+1, made +0,2}
	designTaskMap["Burden_Snub"] = designTask{"Burden", "Snub", "Sn", 1, "=2", 0.7, -3, "", 0, "", -1, "", 0, "", 1.5, 2}
	designTaskMap["Burden_Vheavy"] = designTask{"Burden", "Vheavy", "Vh", 0, "+5", 4.0, 4, "", 0, "", 5, "", 0, "", 5.0, 2}
	designTaskMap["Burden_Vlight"] = designTask{"Burden", "Vlight", "Vl", 1, "-2", 0.6, -2, "", 0, "", -1, "", 0, "", 2.0, 2}
	designTaskMap["Burden_VRF"] = designTask{"Burden", "VRF", "Vrf", 2, "+0", 14.0, 5, "", 0, "", 1, "", 0, "Only Gun, Only Machinegun", 9.0, 2}
	designTaskMap["Stage_(blank)"] = designTask{"Stage", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 3}
	designTaskMap["Stage_Advanced"] = designTask{"Stage", "Advanced", "A", 3, "+0", 0.8, -3, "", 0, "", 2, "", 0, "", 2.0, 3}
	designTaskMap["Stage_Alternate"] = designTask{"Stage", "Alternate", "Alt", 0, "+1", 1.1, 0, "", 0, "", 2, "", 0, "B+Flux", 1.1, 3}
	designTaskMap["Stage_Basic"] = designTask{"Stage", "Basic", "B", 0, "+0", 1.3, 1, "", 0, "", 0, "", 0, "", 0.7, 3}
	designTaskMap["Stage_Early"] = designTask{"Stage", "Early", "E", -1, "-1", 1.7, 1, "", 0, "", 0, "", 0, "EOU-1", 0.7, 3}
	designTaskMap["Stage_Experimental"] = designTask{"Stage", "Experimental", "Exp", -3, "-1", 2.0, 3, "", 0, "", 0, "", 0, "R=-2", 4.0, 3}
	designTaskMap["Stage_Generic"] = designTask{"Stage", "Generic", "Gen", 1, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 0.5, 3}
	designTaskMap["Stage_Improved"] = designTask{"Stage", "Improved", "Im", 1, "+0", 1.0, -1, "", 0, "", 1, "", 0, "EOU+1", 1.1, 3}
	designTaskMap["Stage_Modified"] = designTask{"Stage", "Modified", "Mod", 2, "+0", 0.9, 0, "", 0, "", 1, "", 0, "", 1.2, 3}
	designTaskMap["Stage_Precision"] = designTask{"Stage", "Precision", "Pr", 6, "+3", 4.0, 2, "", 0, "", 0, "", 0, "Only Designator", 5.0, 3}
	designTaskMap["Stage_Prototype"] = designTask{"Stage", "Prototype", "P", -2, "-1", 1.9, 2, "", 0, "", 0, "", 0, "", 3.0, 3}
	designTaskMap["Stage_Remote"] = designTask{"Stage", "Remote", "R", 1, "+0", 1.0, 0, "", 0, "", 0, "", 0, "Not Pistol", 7.0, 3}
	designTaskMap["Stage_Sniper"] = designTask{"Stage", "Sniper", "Sn", 1, "+1", 1.1, 1, "", 0, "", 1, "", 0, "Only Rifle, Q=+2", 2.0, 3} //TODO: Check Errata for +1 on D2 for stantard (makes no sense if it is in standard and not Sniper,3}
	designTaskMap["Stage_Standard"] = designTask{"Stage", "Standard", "St", 0, "+0", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 3}             //TODO: Check Errata for +1 on D2 for stantard (makes no sense if it is in standard and not Sniper,3}
	designTaskMap["Stage_Target"] = designTask{"Stage", "Target", "T", 0, "+0", 1.1, 1, "", 0, "", 0, "", 0, "Only Rifle, Only Pistol, Q=+2", 1.5, 3}
	designTaskMap["Stage_Ultimate"] = designTask{"Stage", "Ultimate", "Ul", 4, "+0", 0.7, -4, "", 0, "", 2, "", 0, "R=+4", 1.4, 3}
	designTaskMap["Users_(blank)"] = designTask{"Users", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU+0", 1.0, 4}
	designTaskMap["Users_Universal"] = designTask{"Users", "Universal", "U", 0, "", 1.1, 1, "", 0, "", 0, "", 0, "EOU-1", 1.0, 4}
	designTaskMap["Users_Man"] = designTask{"Users", "Man", "M", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU+0", 1.0, 4}
	designTaskMap["Users_Vargr"] = designTask{"Users", "Vargr", "V", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0, 4}
	designTaskMap["Users_K'kree"] = designTask{"Users", "K'kree", "K", 0, "", 1.3, 2, "", 0, "", 0, "", 0, "EOU+0", 1.0, 4}
	designTaskMap["Users_Grasper (Hiver)"] = designTask{"Users", "Grasper (Hiver)", "H", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0, 4}
	designTaskMap["Users_Paw (Aslan)"] = designTask{"Users", "Paw (Aslan)", "P", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-1", 1.0, 4}
	designTaskMap["Users_Gripper"] = designTask{"Users", "Gripper", "G", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0, 4}
	designTaskMap["Users_Tentacle (Vegan)"] = designTask{"Users", "Tentacle (Vegan)", "T", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0, 4}
	designTaskMap["Users_Socket"] = designTask{"Users", "Socket", "S", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "EOU-2", 1.0, 4}
	designTaskMap["Portability_(blank)"] = designTask{"Portability", "(blank)", "", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 5}
	designTaskMap["Portability_Crewed"] = designTask{"Portability", "Crewed", "C", 0, "", 1.0, 1, "", 0, "", 0, "", 0, "", 1.0, 5}
	designTaskMap["Portability_Fixed"] = designTask{"Portability", "Fixed", "F", 0, "+1", 1.0, 4, "", 0, "", 0, "", 0, "", 1.0, 5}
	designTaskMap["Portability_Portable"] = designTask{"Portability", "Portable", "P", 0, "+1", 1.0, -2, "", 0, "", 0, "", 0, "", 1.0, 5}
	designTaskMap["Portability_Vechicle Mount"] = designTask{"Portability", "Vechicle Mount", "V", 0, "+1", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 5}
	designTaskMap["Portability_Turret"] = designTask{"Portability", "Turret", "T", 0, "", 1.0, 0, "", 0, "", 0, "", 0, "", 1.0, 5}
	if val, ok := designTaskMap[key]; ok {
		dt = val
	}
	return dt
}

type qrebs struct {
	quality     int
	reliability int
	easeOfUse   int
	burden      int
	safety      int
}

func (q *qrebs) String() string {
	str := ""
	str = qrebsConv(q.quality) + qrebsConv(q.reliability) + qrebsConv(q.easeOfUse) + qrebsConv(q.burden) + qrebsConv(q.safety)
	return str
}

func qrebsConv(i int) string {
	if i >= 0 {
		return strconv.Itoa(i)
	}
	i = 9 - i
	return TrvCore.DigitToEhex(i)
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
burdenDescriptorMap["Burden_VRF"] = designTask{"Burden", "VRF", "Vrf", 2, "+0", 14.0, 5, "", 0, "", 1, "", 0, "Only Gun, Only Machinegun", 9.0}
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

	bt := []byte(new)
	if len(bt) == 0 {
		return initial
	}
	if len(bt) == 1 {
		newInt, err := strconv.Atoi(new)
		if err != nil {
			panic("recalculateRange: " + err.Error())
		}
		return newInt // + initial
	}
	switch string(bt[0]) {
	case "=":
		new = strings.TrimPrefix(new, "=")
		newInt, err := strconv.Atoi(new)
		if err != nil {
			panic("recalculateRange: " + err.Error())
		}
		return newInt
	case "+", "-", "":
		newInt, err := strconv.Atoi(new)
		if err != nil {
			panic("recalculateRange: " + err.Error())
		}
		return newInt + initial

	default:
		panic("Unknown case '" + string(bt[0]) + "'")
	}

	//return initial + newInt
}

/*
урон - бросок - среднее значение
1 - 1d6   //3.5
2 - 1d6+2 //5.5
3 - 2d6   //7
4 - 2d6+3 //10
5 - 3d6   //10.5
6 - 4d6-2 //12
7 - 4d6   //14
8 - 5d6-2 //15.5
9 - 5d6   //17.5
10- 6d6   // 21
11- 1DD   // 35
12- 2DD   // 70

1 2 3 4 5 6
*/

/*
Protector
15
T
2 mm
Disk
1.1 Dense
250 Cr base
socket
proxy.digitalresistance.dog

port
443

cred
d41d8cd98f00b204e9800998ecf8427e

*/
