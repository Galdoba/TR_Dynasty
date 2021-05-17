package weapons

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/core/qrebs"
)

func ConstructTest() {
	w03 := tableWeapons03()
	for i := 0; i < len(w03)+2; i++ {
		fmt.Println(w03[i])
	}
}

type Weapon struct {
	Model    string
	LongName string
	TL       int
	Range    int
	Mass     float64
	QREBS    qrebs.EvaluationData
	H1       int
	D1       string
	H2       int
	D2       string
	H3       int
	D3       string
	HitsV1   int
	Cr       float64
}

func New() Weapon {
	wp := Weapon{}
	return wp
}

type tableWeapons03Entry struct {
	Category string
	Code     string
	Type     string
	TL       int
	Range    int
	Mass     float64
	QREBS    int
	H1       string
	D1       int
	Misc     string
	HitsV1   int
	Cr       float64
}

func tableWeapons03() map[int]tableWeapons03Entry {
	table := make(map[int]tableWeapons03Entry)
	table[1] = tableWeapons03Entry{"Artillery", "G", "Gun", 6, 4, 9.0, 1, "*", 2, "", 2, 5000.0}
	table[2] = tableWeapons03Entry{"Artillery", "Ga", "Gatling", 7, 4, 40.0, 2, "*", 3, "", 2, 8000.0}
	table[3] = tableWeapons03Entry{"Artillery", "C", "Cannon", 6, 6, 200.0, 4, "*", 4, "", 2, 10000.0}
	table[4] = tableWeapons03Entry{"Artillery", "aC", "Autocannon", 8, 6, 300.0, 4, "*", 5, "", 3, 30000.0}
	table[5] = tableWeapons03Entry{"Long Guns", "R", "Rifle", 5, 5, 4.0, 0, "Bullet", 2, "", 2, 500.0}
	table[6] = tableWeapons03Entry{"Long Guns", "R", "Rifle", 5, 5, 4.0, 0, "Laser", 2, "", 2, 500.0}
	table[7] = tableWeapons03Entry{"Long Guns", "C", "Carbine", 5, 4, 3.0, -1, "Bullet", 1, "", 1, 400.0}
	table[8] = tableWeapons03Entry{"Long Guns", "C", "Carbine", 5, 4, 3.0, -1, "Laser", 1, "", 1, 400.0}
	table[9] = tableWeapons03Entry{"Handguns", "P", "Pistol", 5, 2, 1.1, 0, "Bullet", 1, "", 1, 150}
	table[10] = tableWeapons03Entry{"Handguns", "P", "Pistol", 5, 2, 1.1, 0, "Laser", 1, "", 1, 150}
	table[11] = tableWeapons03Entry{"Handguns", "R", "Revolver", 4, 2, 1.25, 0, "Bullet", 1, "", 1, 100}
	table[12] = tableWeapons03Entry{"Handguns", "R", "Revolver", 4, 2, 1.25, 0, "Laser", 1, "", 1, 100}
	table[13] = tableWeapons03Entry{"Shotguns", "S", "Shotgun", 4, 2, 4, 0, "Frag", 2, "", 2, 300}
	table[14] = tableWeapons03Entry{"Machineguns", "Mg", "Machinegun", 6, 5, 8, +1, "Bullet", 4, "", 4, 3000}
	table[15] = tableWeapons03Entry{"Projectors", "Pj", "Projector", 9, 0, 1, 0, "*", 1, "", 1, 300}
	table[16] = tableWeapons03Entry{"Designators", "D", "Designator", 7, 5, 10, +1, "*", 1, "", 1, 2000}
	table[17] = tableWeapons03Entry{"Launchers", "L", "Launcher", 6, 3, 10, +1, "*", 1, "", 0, 1000}
	table[18] = tableWeapons03Entry{"Launchers", "mL", "Multi-Launcher", 8, 5, 8, +1, "*", 1, "", 0, 3000}
	return table
}
