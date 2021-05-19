package weapons

// import (
// 	"fmt"

// 	"github.com/Galdoba/TR_Dynasty/pkg/core/qrebs"
// )

// func ConstructTest() {
// 	w03 := tableWeapons03()
// 	for i := 0; i < len(w03)+2; i++ {
// 		fmt.Println(w03[i])
// 	}
// 	wp := New()
// 	wp.ApplyW03(14)
// 	fmt.Println(wp)
// }

// type Weapon struct {
// 	Model      string
// 	Type       string
// 	Category   string
// 	Descriptor string
// 	LongName   string
// 	TL         int
// 	Range      int
// 	Mass       float64
// 	QREBS      qrebs.EvaluationData
// 	H1         string
// 	D1         int
// 	H2         string
// 	D2         int
// 	H3         string
// 	D3         int
// 	HitsV1     int
// 	Cr         float64
// }

// func New() Weapon {
// 	wp := Weapon{}
// 	wp.Mass = 1
// 	wp.Cr = 1
// 	return wp
// }

// type tableWeapons03Entry struct {
// 	Category string
// 	Code     string
// 	Type     string
// 	TL       int
// 	Range    int
// 	Mass     float64
// 	QREBS    int
// 	H1       string
// 	D1       int
// 	Misc     string
// 	HitsV1   int
// 	Cr       float64
// }

// func tableWeapons03() map[int]tableWeapons03Entry {
// 	table := make(map[int]tableWeapons03Entry)
// 	table[1] = tableWeapons03Entry{"Artillery", "G", "Gun", 6, 4, 9.0, 1, "*", 2, "", 2, 5000.0}
// 	table[2] = tableWeapons03Entry{"Artillery", "Ga", "Gatling", 7, 4, 40.0, 2, "*", 3, "", 2, 8000.0}
// 	table[3] = tableWeapons03Entry{"Artillery", "C", "Cannon", 6, 6, 200.0, 4, "*", 4, "", 2, 10000.0}
// 	table[4] = tableWeapons03Entry{"Artillery", "aC", "Autocannon", 8, 6, 300.0, 4, "*", 5, "", 3, 30000.0}
// 	table[5] = tableWeapons03Entry{"Long Guns", "R", "Rifle", 5, 5, 4.0, 0, "Bullet", 2, "", 2, 500.0}
// 	table[6] = tableWeapons03Entry{"Long Guns", "R", "Rifle", 5, 5, 4.0, 0, "Laser", 2, "", 2, 500.0}
// 	table[7] = tableWeapons03Entry{"Long Guns", "C", "Carbine", 5, 4, 3.0, -1, "Bullet", 1, "", 1, 400.0}
// 	table[8] = tableWeapons03Entry{"Long Guns", "C", "Carbine", 5, 4, 3.0, -1, "Laser", 1, "", 1, 400.0}
// 	table[9] = tableWeapons03Entry{"Handguns", "P", "Pistol", 5, 2, 1.1, 0, "Bullet", 1, "", 1, 150}
// 	table[10] = tableWeapons03Entry{"Handguns", "P", "Pistol", 5, 2, 1.1, 0, "Laser", 1, "", 1, 150}
// 	table[11] = tableWeapons03Entry{"Handguns", "R", "Revolver", 4, 2, 1.25, 0, "Bullet", 1, "", 1, 100}
// 	table[12] = tableWeapons03Entry{"Handguns", "R", "Revolver", 4, 2, 1.25, 0, "Laser", 1, "", 1, 100}
// 	table[13] = tableWeapons03Entry{"Shotguns", "S", "Shotgun", 4, 2, 4, 0, "Frag", 2, "", 2, 300}
// 	table[14] = tableWeapons03Entry{"Machineguns", "Mg", "Machinegun", 6, 5, 8, +1, "Bullet", 4, "", 4, 3000}
// 	table[15] = tableWeapons03Entry{"Projectors", "Pj", "Projector", 9, 0, 1, 0, "*", 1, "", 1, 300}
// 	table[16] = tableWeapons03Entry{"Designators", "D", "Designator", 7, 5, 10, +1, "*", 1, "", 1, 2000}
// 	table[17] = tableWeapons03Entry{"Launchers", "L", "Launcher", 6, 3, 10, +1, "*", 1, "", 0, 1000}
// 	table[18] = tableWeapons03Entry{"Launchers", "mL", "Multi-Launcher", 8, 5, 8, +1, "*", 1, "", 0, 3000}
// 	return table
// }

// func (wp *Weapon) ApplyW03(i int) {
// 	w03 := tableWeapons03()
// 	wp.Category = w03[i].Category
// 	wp.Type = w03[i].Type
// 	wp.TL = w03[i].TL
// 	wp.Mass = wp.Mass * w03[i].Mass
// 	wp.Cr = wp.Cr * w03[i].Cr
// 	wp.H1 = w03[i].H1
// 	wp.D1 = w03[i].D1
// 	wp.HitsV1 = w03[i].HitsV1
// }

// type tableWeapons04Entry struct {
// 	Category   string
// 	Code       string
// 	Descriptor string
// 	TL         int
// 	Range      int
// 	Mass       float64
// 	QREBS      int
// 	H2         string
// 	D2         int
// 	H3         string
// 	D3         int
// 	HitsV1     int
// 	Cr         float64
// }

// func tableWeapons04() map[int]tableWeapons04Entry {
// 	table := make(map[int]tableWeapons04Entry)
// 	//Artillery (blank) x1.0
// 	table[1] = tableWeapons04Entry{"Artillery", "", "", 0, 0, 1.0, 0, "*", -1, "", -1, -1, 1.0}
// 	table[2] = tableWeapons04Entry{"Artillery", "aF", "Anti-Flyer", +4, 6, 6.0, 0, "Frag", 1, "Blast", 3, 4, 3.0}
// 	table[3] = tableWeapons04Entry{"Artillery", "aT", "Anti-Tank", 0, 5, 8.0, 0, "Pen", 3, "Blast", 3, 6, 2.0}
// 	// Artillery A Assault +2 =4 x0.8 Bang 1 Blast 2 3 x 1.5
// 	// Artillery F Fusion +7 =4 x2.3 Pen 4 Burn 4 8 x 6.0
// 	// Artillery G Gauss +7 =4 x0.9 Bullet 3 3 x 2.0
// 	// Artillery P Plasma +5 =4 x2.5 Pen 3 Burn 3 6 x 2.0
// 	// Long Guns (blank) x1.0
// 	// Long Guns Ac Accelerator +4 x0.6 Bullet 2 2 x 3.0
// 	// Long Guns A Assault +2 =4 x0.8 Blast 2 Bang 1 3 x 1.5
// 	// Long Guns B Battle +1 =5 x1.0 +1 Bullet 1 1 x 0.8
// 	// Long Guns C Combat +2 =3 x0.9 Frag 2 2 x 1.5
// 	// Long Guns D Dart +1 =4 x0.6 Tranq 1-2-3 1-2-3 x 0.9
// 	// Long Guns P Poison Dart +1 =4 x1.0 Poison 1-2-3 1-2-3 x 0.9
// 	// Long Guns G Gauss +7 x0.9 Bullet 3 3 x 2.0
// 	// Long Guns H Hunting =3 x0.9 -1 Bullet 1 1 x 1.2
// 	// Long Guns L Laser +5 x1.2 Burn 2 Pen 2 4 x 6.0
// 	// Long Guns Sp Splat +2 =4 x1.3 +1 Bullet 1 1 x 2.4
// 	// Long Guns S Survival =2 x0.5 Bullet 1 1 x 1.2
// 	// Handguns (blank) x1.0
// 	// Handguns Ac Accelerator +4 x0.6 Bullet 2 2 x 3.0
// 	// Handguns L Laser +5 x1.2 Burn 2 Pen 2 4 x 2.0
// 	// Handguns M Machine =2 x1.2 Bullet 2 x 1.5
// 	// Shot Shotguns (blank) x1.0
// 	// Shot Shotguns A Assault +2 =4 x0.8 Bang 1 Blast 2 3 x 2.0
// 	// Shot Shotguns H Hunting =3 x0.9 Bullet 1 1 x 1.2
// 	// Machineguns (blank) x1.0
// 	// Machineguns aF Anti-Flyer +4 =6 x6.0 Frag 1 Blast 3 4 x 3.0
// 	// Machineguns A Assault +2 =4 x0.8 Bang 1 Blast 2 3 x 1.5
// 	// Machineguns S Sub -1 =3 x0.3 Bullet -1 -1 x 0.9
// 	// Spray Designators A Acid =3 x1.0 +1 Corrode 2 Pen 1-2-3 4 x 3.0
// 	// Spray Designators H Fire =1 x0.9 Burn 1-2-3 Pen 1-2-3 2-4-6 x 2.0
// 	// Spray Designators P Poison Gas =2 x1.0 Gas 1-2-3 Poison 1-2-3 2-4-6 x 3.0
// 	// Spray Designators S Stench +3 =2 x0.4 Stench 1-2-3 1-2-3 x 1.2

// 	// Exotic Emp EMP +1 =3 x1.0 EMP 1-2-3 1 x 4.0
// 	// Designators F Flash -1 =2 x0.5 Flash 1-2-3 2 x 1.5
// 	// And Projectors C Freeze +1 =3 x1.0 +1 Cold 1-2-3 2 x 3.0
// 	// G Grav +5 =2 x3.0 Grav 1-2-3 3 x 20.0
// 	// L Laser +5 x1.2 Burn 1-2-3 Pen 1-2-3 2-4-6 x 6.0
// 	// M Mag +4 =1 x2.0 EMP 1-2-3 Mag 1-2-3 2-4-6 x 15.0
// 	// Psi Psi Amp +4 =2 x1.0 Psi 1-2-3 1-2-3 x 9.0
// 	// R Rad +1 =4 x1.0 +2 Rad 1-2-3 1-2-3 x 8.0
// 	// Sh Shock =2 x0.5 Elec 1-2-3 Pain 1-2-3 2-4-6 x 2.0
// 	// S Sonic +3 =2 x0.6 Sound 1-2-3 Bang 1-2-3 2-4-6 x 1.1
// 	// Launcher
// 	// Launchers aF AF Missile +4 =7 x4.0 Frag 2 Blast 3 5 x 3.0
// 	// aT AT Missile +3 =4 x1.0 +1 Frag 2 Pen 3 5 x 2.0
// 	// Gr Grenade +1 =4 x0.8 Frag 2 Blast 2 4 x 1.0
// 	// M Missile +1 =6 x2.2 Frag 2 Pen 2 4 x 5.0
// 	// RAM RAM Grenade +2 =6 x1.0 Frag 2 Blast 2 4 x 3.0
// 	// R Rocket -1 =5 x3.0 Frag 2 Pen 2 4 x 1.0

// 	return table
// }

// func (wp *Weapon) FinalizeDesign() {

// }
