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
