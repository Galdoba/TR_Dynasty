package weapons

func DataFrom(table, row, col int) string {
	return 0
}

func meleeWeapons() [][]string {
	data := [][]string{
		[]string{"Category       | Code | Descriptor       | TL | Range | Mass | qreBs  | H1   | D1  | H2   | D2 | Hits(v1) | Cr"},
		[]string{"Short Blades   | K    | Knife            | 1  | R     | 0.5  |        | Cuts | 2   |      |    | 2D       | 50"},
		[]string{"Short Blades   | D    | Dagger           | 2  | R     | 0.5  |        | Cuts | 2   |      |    | 2D       | 50"},
		[]string{"Short Blades   | TK   | Trench Knife     | 4  | R     | 1    |        | Cuts | 2   | Blow | 1  | 2D       | 100"},
		[]string{"Short Blades   | BK   | Big Knife        | 5  | T     | 3    |        | Cuts | 2   | Pen  | 2  | 2D       | 200"},
		[]string{"Short Blades   | GBK  | Great Big Knife  | 6  | 1     | 6    |        | Cuts | 2   | Pen  | 2  | 2D       | 900"},
		[]string{"Medium Blades  | S    | Sword            | 3  | 1     | 2    |        | Cuts | 2   |      |    | 2D       | 300"},
		[]string{"Medium Blades  | sS   | Short Sword      | 3  | 1     | 1    | B= - 1 | Cuts | 2   |      |    | 2D       | 300"},
		[]string{"Medium Blades  | bS   | Broadsword       | 4  | 1     | 3    |        | Cuts | 3   |      |    | 3D       | 700"},
		[]string{"Medium Blades  | C    | Cutlass          | 3  | 1     | 2    |        | Cuts | 2   |      |    | 2D       | 200"},
		[]string{"Medium Blades  | OC   | Officers Cutlass | 5  | 1     | 1    |        | Cuts | 2   |      |    | 2D       | 400"},
		[]string{"Long Blades    | P    | Spear/Pike       | 1  | 1     | 2    | B= +3  | Cuts | 2   |      |    | 2D       | 50"},
		[]string{"Special Blades | Ax   | Axe              | 2  | T     | 2    |        | Cuts | 3   |      |    | 3D       | 60"},
		[]string{"Special Blades | A    | Space Axe        | 9  | 1     | 2    |        | Cuts | 2   | Pen  | 2  | 2D       | 500"},
		[]string{"Special Blades | V    | Vibro-Blade      | 10 | 1     | 0.5  |        | Cuts | 2   |      |    | 2D       | 900"},
		[]string{"Special Blades | m    | Mace             | 2  | 1     | 4    |        | Cuts | 1   | Blow | 2  | 2D       | 100"},
		[]string{"Special Blades | c    | Club             | 1  | 1     | 2    |        | Blow | =C1 |      |    | 1D       | 10"},
	}
	return data
}

func bodyWeapons() [][]string {
	data := [][]string{
		[]string{"Category       | Code | Descriptor | Range | H1   | D1  | H2     | D2 | Hits(v1)"},
		[]string{"Body Weapons   | Fi   | Fists      | R     | Blow | =C1 |        |    | 1D"},
		[]string{"Body Weapons   | Te   | Tentacle   | 0     | Hit  | =C1 | Suff   | 1  | 1D"},
		[]string{"Body Weapons   | Ho   | Horns      | R     | Pen  | =C1 |        |    | 2D"},
		[]string{"Body Weapons   | Tu   | Tusks      | R     | Pen  | =C1 |        |    | 2D"},
		[]string{"Body Weapons   | Fa   | Fangs      | R     | Pen  | =C1 |        |    | 2D"},
		[]string{"Body Weapons   | T    | Teeth      | R     | Cuts | =C1 |        |    | 1D"},
		[]string{"Body Weapons   | Cl   | Claws      | R     | Cuts | =C1 |        |    | 1D"},
		[]string{"Body Weapons   | Dc   | Dew Claw   | R     | Cuts | =C1 |        |    | 2D"},
		[]string{"Body Weapons   | H    | Hooves     | R     | Blow | =C1 |        |    | 2D"},
		[]string{"Body Weapons   | Sp   | Spikes     | 0     | Pen  | =C1 |        |    | 2D"},
		[]string{"Body Weapons   | St   | Sting      | R     | Pen  | =C1 | Poison | 2D | 3D"},
	}
	return data
}

func rangeWeapons() [][]string {
	data := [][]string{
		[]string{"Category    | Code | Type           | TL | Range | Mass | qreBs | H1     | D1 | Misc                | Hits(v1) | Cr"},
		[]string{"Artillery   | G    | Gun            | 6  | 4     | 9    | +1    | *      | 2  |                     | 2        | 5,000"},
		[]string{"Artillery   | Ga   | Gatling        | 7  | 4     | 40   | +2    | *      | 3  |                     | 2        | 8,000"},
		[]string{"Artillery   | C    | Cannon         | 6  | 6     | 200  | +4    | *      | 4  |                     | 2        | 10,000"},
		[]string{"Artillery   | aC   | Autocannon     | 8  | 6     | 300  | +4    | *      | 5  |                     | 3        | 30,000"},
		[]string{"Long Guns   | R    | Rifle          | 5  | 5     | 4    | 0     | Bullet | 2  | Not Bullet if Laser | 2        | 500"},
		[]string{"Long Guns   | C    | Carbine        | 5  | 4     | 3    | -1    | Bullet | 1  | Not Bullet if Laser | 1        | 400"},
		[]string{"Handguns    | P    | Pistol         | 5  | 2     | 1.1  | 0     | Bullet | 1  | Not Bullet if Laser | 1        | 150"},
		[]string{"Handguns    | R    | Revolver       | 4  | 2     | 1.25 | 0     | Bullet | 1  | Not Bullet if Laser | 1        | 100"},
		[]string{"Shotguns    | S    | Shotgun        | 4  | 2     | 4    | 0     | Frag   | 2  |                     | 2        | 300"},
		[]string{"Machineguns | Mg   | Machinegun     | 6  | 5     | 8    | +1    | Bullet | 4  |                     | 4        | 3,000"},
		[]string{"Projectors  | Pj   | Projector      | 9  | 0     | 1    | 0     | *      | 1  |                     | 1        | 300"},
		[]string{"Designators | D    | Designator     | 7  | 5     | 10   | +1    | *      | 1  |                     | 1        | 2,000"},
		[]string{"Launchers   | L    | Launcher       | 6  | 3     | 10   | +1    | *      | 1  |                     | 0        | 1,000"},
		[]string{"Launchers   | mL   | Multi-Launcher | 8  | 5     | 8    | +1    | *      | 1  |                     | 0        | 3,000"},
	}
	return data
}

func weaponsDiscriptors() [][]string {
	data := [][]string{
		[]string{"Category    | Code | Descriptor   | TL | R= | Mass | qreBs | H2      | D2    | H3     | D3    | Hits(v1) | Cr"},
		[]string{"Artillery   | aF   | Anti-Flyer   | +4 | =6 | x6.0 |       | Frag    | 1     | Blast  | 3     | 4        | x3.0"},
		[]string{"Artillery   | aT   | Anti-Tank    |    | =5 | x8.0 |       | Pen     | 3     | Blast  | 3     | 6        | x2.0"},
		[]string{"Artillery   | A    | Assault      | +2 | =4 | x0.8 |       | Bang    | 1     | Blast  | 2     | 3        | x1.5"},
		[]string{"Artillery   | F    | Fusion       | +7 | =4 | x2.3 |       | Pen     | 4     | Burn   | 4     | 8        | x6.0"},
		[]string{"Artillery   | G    | Gauss        | +7 | =4 | x0.9 |       | Bullet  | 3     |        |       | 3        | x2.0"},
		[]string{"Artillery   | P    | Plasma       | +5 | =4 | x2.5 |       | Pen     | 3     | Burn   | 3     | 6        | x2.0"},
		[]string{"Long Guns   |      | (blank)      |    |    |      |       |         |       |        |       |          | x1.0"},
		[]string{"Long Guns   | Ac   | Accelerator  | +4 |    | x0.6 |       | Bullet  | 2     |        |       | 2        | x3.0"},
		[]string{"Long Guns   | A    | Assault      | +2 | =4 | x0.8 |       | Blast   | 2     | Bang   | 1     | 3        | x1.5"},
		[]string{"Long Guns   | B    | Battle       | +1 | =5 | x1.0 | +1    | Bullet  | 1     |        |       | 1        | x0.8"},
		[]string{"Long Guns   | C    | Combat       | +2 | =3 | x0.9 |       | Frag    | 2     |        |       | 2        | x1.5"},
		[]string{"Long Guns   | D    | Dart         | +1 | =4 | x0.6 |       | Tranq   | 1-2-3 |        |       | 1-2-3    | x0.9"},
		[]string{"Long Guns   | P    | Poison Dart  | +1 | =4 | x1.0 |       | Poison  | 1-2-3 |        |       | 1-2-3    | x0.9"},
		[]string{"Long Guns   | G    | Gauss        | +7 |    | x0.9 |       | Bullet  | 3     |        |       | 3        | x2.0"},
		[]string{"Long Guns   | H    | Hunting      |    | =3 | x0.9 | -1    | Bullet  | 1     |        |       | 1        | x1.2"},
		[]string{"Long Guns   | L    | Laser        | +5 |    | x1.2 |       | Burn    | 2     | Pen    | 2     | 4        | x6.0"},
		[]string{"Long Guns   | Sp   | Splat        | +2 | =4 | x1.3 | +1    | Bullet  | 1     |        |       | 1        | x2.4"},
		[]string{"Long Guns   | S    | Survival     |    | =2 | x0.5 |       | Bullet  | 1     |        |       | 1        | x1.2"},
		[]string{"Handguns    |      | (blank)      |    |    |      |       |         |       |        |       |          | x1.0"},
		[]string{"Handguns    | Ac   | Accelerator  | +4 |    | x0.6 |       | Bullet  | 2     |        |       | 2        | x3.0"},
		[]string{"Handguns    | L    | Laser        | +5 |    | x1.2 |       | Burn    | 2     | Pen    | 2     | 4        | x2.0"},
		[]string{"Handguns    | M    | Machine      |    | =2 | x1.2 |       | Bullet  | 2     |        |       | 2        | x1.5"},
		[]string{"Shotguns    |      | (blank)      |    |    |      |       |         |       |        |       |          | x1.0"},
		[]string{"Shotguns    | A    | Assault      | +2 | =4 | x0.8 |       | Bang    | 1     | Blast  | 2     | 3        | x2.0"},
		[]string{"Shotguns    | H    | Hunting      |    | =3 | x0.9 |       | Bullet  | 1     |        |       | 1        | x1.2"},
		[]string{"Machineguns |      | (blank)      |    |    |      |       |         |       |        |       |          | x1.0"},
		[]string{"Machineguns | aF   | Anti-Flyer  +| 4  | =6 | x6.0 |       | Frag    | 1     | Blast  | 3     | 4        | x3.0"},
		[]string{"Machineguns | A    | Assault     +| 2  | =4 | x0.8 |       | Bang    | 1     | Blast  | 2     | 3        | x1.5"},
		[]string{"Machineguns | S    | Sub         -| 1  | =3 | x0.3 |       | Bullet  | -1    |        |       | -1       | x0.9"},
		[]string{"Spray       | A    | Acid         |    | =3 | x1.0 | +1    | Corrode | 2     | Pen    | 1-2-3 | 4        | x3.0"},
		[]string{"Spray       | H    | Fire         |    | =1 | x0.9 |       | Burn    | 1-2-3 | Pen    | 1-2-3 | 2-4-6    | x2.0"},
		[]string{"Spray       | P    | Poison Gas   |    | =2 | x1.0 |       | Gas     | 1-2-3 | Poison | 1-2-3 | 2-4-6    | x3.0"},
		[]string{"Spray       | S    | Stench      +| 3  | =2 | x0.4 |       | Stench  | 1-2-3 |        |       | 1-2-3    | x1.2"},
		[]string{"Launchers   | aF   | AF Missile  +| 4  | =7 | x4.0 |       | Frag    | 2     | Blast  | 3     | 5        | x3.0"},
		[]string{"Launchers   | aT   | AT Missile  +| 3  | =4 | x1.0 | +1    | Frag    | 2     | Pen    | 3     | 5        | x2.0"},
		[]string{"Launchers   | Gr   | Grenade     +| 1  | =4 | x0.8 |       | Frag    | 2     | Blast  | 2     | 4        | x1.0"},
		[]string{"Launchers   | M    | Missile     +| 1  | =6 | x2.2 |       | Frag    | 2     | Pen    | 2     | 4        | x5.0"},
		[]string{"Launchers   | RAM  | RAM Grenade +| 2  | =6 | x1.0 |       | Frag    | 2     | Blast  | 2     | 4        | x3.0"},
		[]string{"Launchers   | R    | Rocket      -| 1  | =5 | x3.0 |       | Frag    | 2     | Pen    | 2     | 4        | x1.0"},
	}
	return data
}

func weaponBurden() [][]string {
	data := [][]string{
		[]string{"Category | Code | Descriptor      | TL  | R= | Mass  | qreBs | Misc  | D2 | Comment                | Cr"},
		[]string{"Burden   |      | (blank)         | 0   | 0  | x1.0  | 0     |       | 0  |                        | x1.0"},
		[]string{"Burden   | aD   | Anti-Designator | +3  | +1 | x3.0  | +3    |       | 1  | Not Pistols. Shotguns. | x3.0"},
		[]string{"Burden   | B    | Body            | +2  | =1 | x0.5  | -4    |       | -1 | Only Pistols.          | x3.0"},
		[]string{"Burden   | D    | Disposable      | +3  | 0  | x0.9  | -1    | Q= -2 | 0  |                        | x0.5"},
		[]string{"Burden   | H    | Heavy           | 0   | +1 | x1.3  | +3    |       | 1  | Not Laser              | x2.0"},
		[]string{"Burden   | Lt   | Light           | 0   | -1 | x0.7  | -1    |       | -1 | Not Laser              | x1.1"},
		[]string{"Burden   | M    | Magnum          | +1  | +1 | x1.1  | +1    |       | 1  | Only Pistols.          | x1.1"},
		[]string{"Burden   | M    | Medium          | 0   | 0  | x1.0  | 0     |       | 0  | Not Pistols.           | x1.0"},
		[]string{"Burden   | R    | Recoilless      | +1  | -1 | x1.2  | 0     |       | 1  |                        | x3.0"},
		[]string{"Burden   | Sn   | Snub            | +1  | =2 | x0.7  | -3    |       | 1  |                        | x1.5"},
		[]string{"Burden   | Vh   | Vheavy          | 0   | +5 | x4.0  | +4    |       | 5  |                        | x5.0"},
		[]string{"Burden   | Vl   | Vlight          | +1  | -2 | x0.6  | -2    |       | -1 |                        | x2.0"},
		[]string{"Burden   | Vrf  | VRF             | +2  | 0  | x14.0 | +5    |       | 1  | Only Guns and MGs      | x9.0"},
	}
	return data
}

func weaponStage() [][]string {
	data := [][]string{
		[]string{"Category | Code | Descriptor      | TL  | R= | Mass  | qreBs | Misc  | D2 | Comment                | Cr"},
		[]string{"Stage             (blank)           0     0    x1.0    0               0                             x1.0"},
		[]string{"Stage      A      Advanced          +3    0    x0.8    -3              2                             x2.0"},
		[]string{"Stage      Alt    Alternate         0     +1   x1.1    F               2                             x1.1"},
		[]string{"Stage B Basic 0 +0 x1.3 +1 0 x 0.7"},
		[]string{"Stage E Early -1 -1 x1.7 +1 0 EOU - 1 x 1.2"},
		[]string{"Stage Exp Experimental -3 -1 x2.0 +3 R= - 2 0 x 4.0"},
		[]string{"Stage Gen Generic +1 0 x1.0 0 0 x 0.5"},
		[]string{"Stage Im Improved +1 0 x1.0 -1 R= +1 1 EOU + 1 x 1.1"},
		[]string{"Stage Mod Modified +2 0 x0.9 0 1 x 1.2"},
		[]string{"Stage Pr Precision +6 +3 x4.0 +2 0 Only Designators. x 5.0"},
		[]string{"Stage P Prototype -2 -1 x1.9 +2 0 x 3.0"},
		[]string{"Stage R Remote +1 0 x1.0 0 0 Not Pistols. x 7.0"},
		[]string{"Stage Sn Sniper +1 +1 x1.1 +1 Q= +2 0 Only Rifles. X 2.0"},
		[]string{"Stage St Standard 0 0 x1.0 0 1 x 1.0"},
		[]string{"Stage T Target 0 0 x1.1 +1 Q= +2 0 Only Rifles and Pistols. x 1.5"},
		[]string{"Stage Ul Ultimate +4 0 x0.7 -4 R= +4 2 x 1.4"},
	}
	return data
}
