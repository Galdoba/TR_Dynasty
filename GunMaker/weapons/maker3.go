package weapons

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/Galdoba/TR_Dynasty/pkg/core/qrebs"
// 	"github.com/Galdoba/devtools/cli/user"
// )

// const (
// 	LongGuns    = "Long Guns"
// 	Handguns    = "Handguns"
// 	Pistol      = "Pistol"
// 	Revolver    = "Revolver"
// 	Rifle       = "Rifle"
// 	Carbine     = "Carbine"
// 	Accelerator = "Accelerator"
// 	Assault     = "Assault"
// 	Battle      = "Battle"
// 	Combat      = "Combat"
// 	Gauss       = "Gauss"
// 	Hunting     = "Hunting"
// 	Laser       = "Laser"
// 	Splat       = "Splat"
// 	Survival    = "Survival"
// )

// type Weapon struct {
// 	Category   string
// 	Descriptor string
// 	Burden     string
// 	Stage      string
// 	Type       string
// 	TL         int
// 	Range      int
// 	Mass       float64
// 	QREBS      qrebs.EvaluationData
// 	HitDice    int
// 	DiceMod    int

// 	codes  []string
// 	Traits []string
// 	Cr     float64
// }

// func ConstructTest() Weapon {
// 	wp := WeaponBase(Handguns + Pistol)
// 	chose, _ := user.ChooseOne("Select Descriptor:", ValidDescr())
// 	wp.ApplyDescriptor(ValidDescr()[chose])
// 	wp.ApplyBurden("Snub")
// 	//wp.ApplyStage("Advanced")
// 	wp.codes[7] = strconv.Itoa(wp.TL)
// 	fmt.Println(wp)
// 	return wp
// }

// func (wp *Weapon) Stats() string {
// 	str := ""
// 	fullname := wp.Stage + " " + wp.Burden + " " + wp.Category
// 	rnge := 0
// 	switch wp.Range {
// 	case 1:
// 		rnge = 5
// 	case 2:
// 		rnge = 50
// 	case 3:
// 		rnge = 150
// 	case 4:
// 		rnge = 500
// 	case 5:
// 		rnge = 1000
// 	case 6:
// 		rnge = 5000
// 	case 7:
// 		rnge = 50000
// 	}
// 	rnge = rnge * (100 + (wp.QREBS.QualityI() * 10)) / 100
// 	str += fullname + " | " + "TL " + strconv.Itoa(wp.TL) + " | " + strconv.Itoa(rnge) + " | "

// 	return str
// }

// func ValidDescr() []string {
// 	return []string{
// 		"",
// 		Accelerator,
// 		Assault,
// 		Battle,
// 		Combat,
// 		Gauss,
// 		Hunting,
// 		Laser,
// 		Splat,
// 		Survival,
// 	}

// }

// func WeaponBase(category string) Weapon {
// 	wp := Weapon{}

// 	switch category {
// 	default:
// 		fmt.Println(category, "not inplemented")
// 	case LongGuns + Rifle:
// 		wp = Weapon{
// 			Category: LongGuns,
// 			Type:     Rifle,
// 			TL:       5,
// 			Range:    5,
// 			Mass:     4.0,
// 			QREBS:    qrebs.Standard(),
// 			HitDice:  2,
// 			codes:    []string{"", "", "", "R", "", "", "-", ""},
// 			Cr:       500,
// 		}
// 	case LongGuns + Carbine:
// 		wp = Weapon{
// 			Category: LongGuns,
// 			Type:     Carbine,
// 			TL:       5,
// 			Range:    4,
// 			Mass:     3.0,
// 			QREBS:    qrebs.Custom(5, 0, 0, -1, 0),
// 			HitDice:  1,
// 			codes:    []string{"", "", "", "C", "", "", "-", ""},
// 			Cr:       500,
// 		}
// 	case Handguns + Pistol:
// 		wp = Weapon{
// 			Category: Handguns,
// 			Type:     Pistol,
// 			TL:       5,
// 			Range:    2,
// 			Mass:     1.1,
// 			QREBS:    qrebs.Standard(),
// 			HitDice:  1,
// 			codes:    []string{"", "", "", "P", "", "", "-", ""},
// 			Cr:       150,
// 		}
// 	case Handguns + Revolver:
// 		wp = Weapon{
// 			Category: Handguns,
// 			Type:     Revolver,
// 			TL:       4,
// 			Range:    2,
// 			Mass:     1.25,
// 			QREBS:    qrebs.Standard(),
// 			HitDice:  1,
// 			codes:    []string{"", "", "", "R", "", "", "-", ""},
// 			Cr:       100,
// 		}
// 	}

// 	return wp
// }

// func (wp *Weapon) ApplyDescriptor(descr string) {
// 	switch wp.Category {
// 	default:
// 		fmt.Println(wp.Category, "Description not implemented")
// 	case "Long Guns":
// 		wp.applyLongGunsDescriptor(descr)
// 	case Handguns:
// 		wp.applyHandgunsDescriptor(descr)
// 	}
// 	wp.codes[7] = strconv.Itoa(wp.TL)
// }

// func (wp *Weapon) applyHandgunsDescriptor(descr string) {
// 	switch descr {
// 	default:
// 		fmt.Println(wp.Category, "Description not implemented 1")
// 	case "":
// 		wp.HitDice++
// 	case "Accelerator":
// 		wp.codes[2] = "Ac"
// 		wp.TL = wp.TL + 4
// 		wp.Mass = wp.Mass * 0.6
// 		wp.HitDice += 2
// 		wp.Cr = wp.Cr * 3.0
// 	case "Laser":
// 		wp.codes[2] = "L"
// 		wp.TL = wp.TL + 5
// 		wp.Mass = wp.Mass * 1.2
// 		wp.HitDice += 2
// 		wp.Cr = wp.Cr * 2.0
// 		wp.Traits = append(wp.Traits, "Zero-G")
// 	}
// }

// func (wp *Weapon) applyLongGunsDescriptor(descr string) {
// 	switch descr {
// 	default:
// 		fmt.Println(wp.Category, "Description not implemented 1")
// 	case "":
// 		wp.HitDice++
// 	case "Accelerator":
// 		wp.codes[2] = "Ac"
// 		wp.TL = wp.TL + 4
// 		wp.Mass = wp.Mass * 0.6
// 		wp.HitDice = wp.HitDice + 2
// 		wp.Cr = wp.Cr * 3
// 		wp.Traits = append(wp.Traits, "Zero-G")
// 	case "Assault":
// 		wp.codes[2] = "A"
// 		wp.TL = wp.TL + 2
// 		wp.Mass = wp.Mass * 0.8
// 		wp.Range = 4
// 		wp.HitDice = wp.HitDice + 2
// 		wp.Cr = wp.Cr * 1.5
// 		wp.Traits = append(wp.Traits, "Auto "+strconv.Itoa((wp.HitDice+1)/2))
// 	case "Battle":
// 		wp.codes[2] = "B"
// 		wp.TL = wp.TL + 1
// 		wp.Mass = wp.Mass * 1
// 		wp.Range = 5
// 		wp.QREBS.Change(qrebs.BULK, -1)
// 		wp.HitDice = wp.HitDice + 1
// 		wp.Cr = wp.Cr * 0.8
// 	case "Combat":
// 		wp.codes[2] = "C"
// 		wp.TL = wp.TL + 2
// 		wp.Mass = wp.Mass * 0.9
// 		wp.Range = 3
// 		wp.HitDice = wp.HitDice + 2
// 		wp.Cr = wp.Cr * 1.5
// 		wp.Traits = append(wp.Traits, "Auto "+strconv.Itoa(wp.HitDice))
// 	case "Gauss":
// 		wp.codes[2] = "G"
// 		wp.TL = wp.TL + 7
// 		wp.Mass = wp.Mass * 0.9
// 		wp.HitDice = wp.HitDice + 3
// 		wp.Cr = wp.Cr * 2
// 		wp.Traits = append(wp.Traits, "AP "+strconv.Itoa(wp.HitDice))
// 	case "Hunting":
// 		wp.codes[2] = "H"
// 		wp.Range = 3
// 		wp.Mass = wp.Mass * 0.9
// 		wp.QREBS.Change(qrebs.BULK, -1)
// 		wp.HitDice = wp.HitDice + 1
// 		wp.Cr = wp.Cr * 1.2
// 	case "Laser":
// 		wp.codes[2] = "L"
// 		wp.TL = wp.TL + 5
// 		wp.Mass = wp.Mass * 1.2
// 		wp.HitDice = wp.HitDice + 3
// 		wp.Cr = wp.Cr * 6
// 		wp.Traits = append(wp.Traits, "Zero-G")
// 	case "Splat":
// 		wp.codes[2] = "Sp"
// 		wp.TL = wp.TL + 2
// 		wp.Range = 4
// 		wp.Mass = wp.Mass * 1.3
// 		wp.QREBS.Change(qrebs.BULK, 1)
// 		wp.HitDice = wp.HitDice + 1
// 		wp.Cr = wp.Cr * 2.4
// 	case "Survival":
// 		wp.codes[2] = "S"
// 		wp.Range = 2
// 		wp.Mass = wp.Mass * 0.5
// 		wp.HitDice = wp.HitDice + 1
// 		wp.Cr = wp.Cr * 1.2
// 	}
// 	wp.Descriptor = descr
// }

// func (wp *Weapon) ApplyBurden(burden string) {
// 	switch burden {
// 	default:
// 		fmt.Println(wp.Category, "Description not implemented 2")
// 	case "":
// 	case "Body":
// 		if wp.Type != "Pistol" {
// 			return
// 		}
// 		wp.codes[1] = "B"
// 		wp.TL = wp.TL + 2
// 		wp.Range = 1
// 		wp.Mass = wp.Mass * 0.5
// 		wp.QREBS.Change(qrebs.BURDEN, -4)
// 		wp.Cr = wp.Cr * 3
// 	case "Heavy":
// 		wp.codes[1] = "H"
// 		wp.Range++
// 		wp.HitDice++
// 		wp.Mass = wp.Mass * 1.3
// 		wp.QREBS.Change(qrebs.BURDEN, 3)
// 		wp.Cr = wp.Cr * 2
// 	case "Light":
// 		wp.codes[1] = "Lt"
// 		wp.Range--
// 		wp.HitDice--
// 		wp.Mass = wp.Mass * 0.7
// 		wp.QREBS.Change(qrebs.BURDEN, -1)
// 		wp.Cr = wp.Cr * 2
// 	case "Magnum":
// 		wp.codes[1] = "M"
// 		wp.TL = wp.TL + 2
// 		wp.Range++
// 		wp.HitDice++
// 		wp.Mass = wp.Mass * 1.1
// 		wp.QREBS.Change(qrebs.BURDEN, 1)
// 		wp.Cr = wp.Cr * 1.1
// 	case "Recoilles":
// 		wp.codes[1] = "R"
// 		wp.TL = wp.TL + 1
// 		wp.Range--
// 		wp.HitDice++
// 		wp.Mass = wp.Mass * 1.2
// 		wp.QREBS.Change(qrebs.BURDEN, 0)
// 		wp.Cr = wp.Cr * 3
// 	case "Snub":
// 		wp.codes[1] = "Sn"
// 		wp.TL = wp.TL + 1
// 		wp.Range = 2
// 		wp.HitDice++
// 		wp.Mass = wp.Mass * 0.7
// 		wp.QREBS.Change(qrebs.BURDEN, -3)
// 		wp.Cr = wp.Cr * 1.5
// 	}
// 	wp.Burden = burden
// }

// func (wp *Weapon) ApplyStage(stage string) {
// 	switch stage {
// 	default:
// 		fmt.Println(wp.Category, "Description not implemented 2")
// 	case "":
// 	case "Advanced":
// 		wp.TL = wp.TL + 3
// 		wp.Mass = wp.Mass * 0.8
// 		wp.QREBS.Change(qrebs.BULK, -3)
// 		wp.HitDice += 2
// 		wp.Cr = wp.Cr * 2
// 		wp.codes[0] = "A"
// 	}
// 	wp.Stage = stage
// }

// /*bulett:
// 1 - 6 )5.0 mm
// BFP = 2
// L = 1
// W = 0.5
// D = 0.5
// Dencity = 4.0
// Container (Value 0.2)
// V = 0.25
// //Value =

// */
