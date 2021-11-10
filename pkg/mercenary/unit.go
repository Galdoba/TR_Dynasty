package mercenary

import (
	"fmt"

	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

const (
	FORMATION_SQUAD          = 1
	FORMATION_SECTION        = 2
	FORMATION_PLATOON        = 3
	FORMATION_COMPANY        = 4
	FORMATION_BATTALION      = 5
	FORMATION_BRIGADE        = 6
	FORMATION_DIVISION       = 7
	FORMATION_CORPS          = 8
	FORMATION_ARMY           = 9
	CBAS_COMBAT              = "Combat"
	CBAS_BOMBARDMENT         = "Bombardment"
	CBAS_AEROSPACE           = "Aerospace"
	CBAS_SUPPORT             = "Support"
	TRAINING_Untrained       = "Untrained"
	TRAINING_Raw             = "Raw"
	TRAINING_Green           = "Green"
	TRAINING_Trained         = "Trained"
	TRAINING_Effective       = "Effective"
	TRAINING_HighlyEffective = "Highly Effective"
	EQIPMENT_Minimal         = "Minimal"
	EQIPMENT_Sparce          = "Sparce"
	EQIPMENT_Basic           = "Basic"
	EQIPMENT_Standard        = "Standard"
	EQIPMENT_Generous        = "Generous"
	EQIPMENT_Lavish          = "Lavish"
	EQIPMENT_Execive         = "Execive"
	MOBILITY_Static          = "Static"
	MOBILITY_SemiMobile      = "SemiMobile"
	MOBILITY_Infantry        = "Infantry"
	MOBILITY_Mechanised      = "Mechanised"
	MOBILITY_Mounted         = "Mounted"
	MOBILITY_Motorised       = "Motorised"
	MOBILITY_AirMobile       = "AirMobile"
	MOBILITY_GravMobile      = "GravMobile"
	MOBILITY_Aerospace       = "Aerospace"
)

type Force struct {
	Name           string
	ForceType      string
	ForceSize      string
	ForceTL        int
	ForceTraining  string
	ForceEquipment []string
	ActivePersonal int
	CBAS           []int
	Mobility       string
	CEI            int
	ECEI           int
	Morale         int
	Reputation     int
}

func NewForce(name string) (*Force, error) {
	f := Force{}
	f.setupActivePersonel()
	f.ForceSize = f.forceFormationStr()
	f.Name = name
	f.setupForceTL()
	f.ForceType = setForceType()
	f.ForceTraining = setForceTraining()
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Minimal) //
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Minimal) //
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Minimal) //
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Minimal) //
	//f.ForceEquipment[f.primaryCBAS()] = EQIPMENT_Standard         //setForceEqupment()
	f.Mobility = MOBILITY_Infantry
	f.setupCBAS()
	//err := fmt.Errorf("Not implemented")
	return &f, nil
}

func (f *Force) forceFormationStr() string {
	switch forceFormation(f.ActivePersonal) {
	default:
		return fmt.Errorf("Force '%v' cannot be created with %v personal", f.Name, f.ActivePersonal).Error()
	case FORMATION_SQUAD:
		f.ForceSize = "Squad"
	case FORMATION_SECTION:
		f.ForceSize = "Section"
	case FORMATION_PLATOON:
		f.ForceSize = "Platoon"
	case FORMATION_COMPANY:
		f.ForceSize = "Company"
	case FORMATION_BATTALION:
		f.ForceSize = "Battalion"
	case FORMATION_BRIGADE:
		f.ForceSize = "Brigade"
	case FORMATION_DIVISION:
		f.ForceSize = "Division"
	case FORMATION_CORPS:
		f.ForceSize = "Corps"
	case FORMATION_ARMY:
		f.ForceSize = "Army"
	}
	return f.ForceSize
}

func (f *Force) setupActivePersonel() {
	tl := -1
	err := fmt.Errorf("Setup Active Personel: ")
	for err != nil {
		fmt.Printf(err.Error() + "\n")
		tl, err = user.InputInt()
		if tl < 2 {
			err = fmt.Errorf("Active Personel cannot be less than 2.\nSetup Active Personel: ")
		}
	}
	f.ActivePersonal = tl
}

func (f *Force) setupForceTL() {
	tl := -1
	err := fmt.Errorf("Setup Force TL: ")
	for err != nil {
		fmt.Printf(err.Error() + "\n")
		tl, err = user.InputInt()
		if tl < 0 {
			err = fmt.Errorf("TL cannot be negative.\nSetup Force TL: ")
		}
		if tl > 33 {
			err = fmt.Errorf("TL%v is imposible.\nSetup Force TL: ", tl)
		}
	}
	f.ForceTL = tl
}

func (f *Force) setupCBAS() {
	switch f.ForceType {
	case CBAS_COMBAT:
		f.CBAS = append(f.CBAS, f.ForceTL/2)
		f.CBAS = append(f.CBAS, 0)
		f.CBAS = append(f.CBAS, 0)
		f.CBAS = append(f.CBAS, utils.Max(1, 0))
	case CBAS_BOMBARDMENT:
		f.CBAS = append(f.CBAS, f.ForceTL/3)
		f.CBAS = append(f.CBAS, f.ForceTL/2)
		f.CBAS = append(f.CBAS, 0)
		f.CBAS = append(f.CBAS, 1)
	case CBAS_AEROSPACE:
		f.CBAS = append(f.CBAS, f.ForceTL/3)
		f.CBAS = append(f.CBAS, 0)
		f.CBAS = append(f.CBAS, f.ForceTL/2)
		f.CBAS = append(f.CBAS, 1)
	case CBAS_SUPPORT:
		f.CBAS = append(f.CBAS, f.ForceTL/3)
		f.CBAS = append(f.CBAS, 0)
		f.CBAS = append(f.CBAS, 0)
		f.CBAS = append(f.CBAS, utils.Max(1, f.ForceTL/2))
	}
	f.equipmentProvision()
	f.applyTrainingEffect()

}

func (f *Force) equipmentProvision() {
	cbas := []string{CBAS_COMBAT, CBAS_BOMBARDMENT, CBAS_AEROSPACE, CBAS_SUPPORT}
	for i := range cbas {
		switch f.ForceEquipment[i] {
		case EQIPMENT_Minimal:
			if f.CBAS[i] < 1 {
				f.CBAS[i] = 1
			}
		case EQIPMENT_Sparce:
			f.CBAS[i] = f.CBAS[i] + 1
		case EQIPMENT_Basic:
			f.CBAS[i] = f.CBAS[i] + 2
		case EQIPMENT_Standard:
			f.CBAS[i] = f.CBAS[i] + 3
		case EQIPMENT_Generous:
			f.CBAS[i] = f.CBAS[i] + 4
		case EQIPMENT_Lavish:
			f.CBAS[i] = f.CBAS[i] + 5
		case EQIPMENT_Execive:
			f.CBAS[i] = f.CBAS[i] + 6
		}
	}
}

func forceFormation(p int) int {
	switch {
	case p >= 2 && p <= 6:
		return FORMATION_SQUAD
	case p >= 2*2 && p <= 2*6:
		return FORMATION_SECTION
	case p >= 3*2*2 && p <= 5*2*5:
		return FORMATION_PLATOON
	case p > 3*3*2*2 && p < 5*5*2*4:
		return FORMATION_COMPANY
	case p > 3*3*3*2*2 && p < 5*5*5*2*3:
		return FORMATION_BATTALION
	case p > 2*3*3*3*2*2 && p < 4*5*5*5*2*3:
		return FORMATION_BRIGADE
	case p > 2*2*3*3*3*2*2 && p < 4*4*5*5*5*2*3:
		return FORMATION_DIVISION
	case p > 2*2*2*3*3*3*2*2 && p < 5*4*4*5*5*5*2*3:
		return FORMATION_CORPS
	case p > 2*2*2*2*3*3*3*2*2 && p < 5*5*4*4*5*5*5*2*3:
		return FORMATION_ARMY
	}
	return -1
}

func setForceType() string {
	err := fmt.Errorf("Initial")
	chosen := "None"
	for err != nil {
		if chosen, err = user.ChooseOneStr("Select Force Type:", []string{CBAS_COMBAT, CBAS_BOMBARDMENT, CBAS_AEROSPACE, CBAS_SUPPORT}); err == nil {
			return chosen
		}
		fmt.Println(err.Error())
	}
	return chosen
}

func setForceTraining() string {
	err := fmt.Errorf("Initial")
	chosen := "None"
	for err != nil {
		if chosen, err = user.ChooseOneStr("Select Force Training:", []string{"Untrained", "Raw", "Green", "Trained", "Effective", "Highly Effective"}); err == nil {
			return chosen
		}
		fmt.Println(err.Error())
	}
	return chosen
}

func setForceEquipment() string {
	err := fmt.Errorf("Initial")
	chosen := "None"
	for err != nil {
		if chosen, err = user.ChooseOneStr("Select Force Type:", []string{"Minimal", "Sparce", "Basic", "Standard", "Generous", "Lavish", "Execive"}); err == nil {
			return chosen
		}
		fmt.Println(err.Error())
	}
	return chosen
}

func (f *Force) trainingEffect() int {
	switch f.ForceTraining {
	case TRAINING_Untrained:
		return -3
	case TRAINING_Raw:
		return -2
	case TRAINING_Green:
		return -1
	case TRAINING_Trained:
		return 0
	case TRAINING_Effective:
		return 1
	case TRAINING_HighlyEffective:
		return 2
	}
	return -999
}

func (f *Force) primaryCBAS() int {
	switch f.ForceType {
	case CBAS_COMBAT:
		return 0
	case CBAS_BOMBARDMENT:
		return 1
	case CBAS_AEROSPACE:
		return 2
	case CBAS_SUPPORT:
		return 3
	}
	return -1
}

func (f *Force) applyTrainingEffect() {
	cbas := []string{CBAS_COMBAT, CBAS_BOMBARDMENT, CBAS_AEROSPACE, CBAS_SUPPORT}
	switch f.ForceTraining {
	case TRAINING_Untrained:
		for i := range cbas {
			if i == f.primaryCBAS() {
				f.CBAS[i] = utils.Max(f.CBAS[i]-2, 1)
				continue
			}
			f.CBAS[i] = utils.Max(f.CBAS[i]-3, 0)
		}
	case TRAINING_Raw:
		for i := range cbas {
			if i == f.primaryCBAS() {
				f.CBAS[i] = utils.Max(f.CBAS[i]-1, 1)
				continue
			}
			f.CBAS[i] = utils.Max(f.CBAS[i]-2, 0)
		}
	case TRAINING_Green:
		for i := range cbas {
			if i == f.primaryCBAS() {
				f.CBAS[i] = utils.Max(f.CBAS[i], 1)
				continue
			}
			f.CBAS[i] = utils.Max(f.CBAS[i]-1, 0)
		}
	case TRAINING_Trained:
		for i := range cbas {
			if i == f.primaryCBAS() {
				f.CBAS[i] = utils.Max(f.CBAS[i]+1, 1)
				continue
			}
			f.CBAS[i] = utils.Max(f.CBAS[i], 0)
		}
	case TRAINING_Effective:
		for i := range cbas {
			if i == f.primaryCBAS() {
				f.CBAS[i] = utils.Max(f.CBAS[i]+2, 1)
				continue
			}
			f.CBAS[i] = utils.Max(f.CBAS[i]+1, 0)
		}
	case TRAINING_HighlyEffective:
		for i := range cbas {
			if i == f.primaryCBAS() {
				f.CBAS[i] = utils.Max(f.CBAS[i]+3, 1)
				continue
			}
			f.CBAS[i] = utils.Max(f.CBAS[i]+2, 0)
		}
	}
}

func (f *Force) DM(field string) int {
	dm := 0
	switch field {
	case "TL":
		switch {
		case f.ForceTL < 1:
			return -3
		case f.ForceTL == 1 || f.ForceTL == 2:
			return -2
		default:
			return f.ForceTL/3 - 2
		}
	case "Combat":
		switch {
		case f.CBAS[0] < 1:
			dm = -3
		case f.CBAS[0] == 1 || f.CBAS[0] == 2:
			dm = -2
		default:
			dm = f.CBAS[0]/3 - 2
		}
		if f.ForceType == "Combat" {
			dm = dm * 3
		}
		return dm
	case "Bombardment":
		switch {
		case f.CBAS[1] < 1:
			dm = -3
		case f.CBAS[1] == 1 || f.CBAS[1] == 2:
			dm = -2
		default:
			dm = f.CBAS[1]/3 - 2
		}
		if f.ForceType == "Bombardment" {
			dm = dm * 3
		}
		return dm
	case "Aerospace":
		switch {
		case f.CBAS[2] < 1:
			dm = -3
		case f.CBAS[2] == 1 || f.CBAS[2] == 2:
			dm = -2
		default:
			dm = f.CBAS[2]/3 - 2
		}
		if f.ForceType == "Aerospace" {
			dm = dm * 3
		}
		return dm
	case "Support":
		switch {
		case f.CBAS[3] < 1:
			dm = -3
		case f.CBAS[3] == 1 || f.CBAS[3] == 2:
			dm = -2
		default:
			dm = f.CBAS[3]/3 - 2
		}
		if f.ForceType == "Support" {
			dm = dm * 3
		}
		return dm
	case "Mobility":
		switch f.Mobility {
		case MOBILITY_Static:
			return -12
		case MOBILITY_SemiMobile:
			return -6
		case MOBILITY_Infantry:
			return -3
		case MOBILITY_Mechanised:
			return 0
		case MOBILITY_Mounted:
			return 2
		case MOBILITY_Motorised:
			return 4
		case MOBILITY_AirMobile:
			return 6
		case MOBILITY_GravMobile:
			return 8
		case MOBILITY_Aerospace:
			return 12
		}
	}
	return -999
}

func (f *Force) CapabilityRecord() string {
	fcr := ""
	fcr += fmt.Sprintf("Force Name   : %v\n", f.Name)
	fcr += fmt.Sprintf("Unit Size    : %v\n", f.ForceSize)
	fcr += fmt.Sprintf("Unit Type    : %v\n", f.ForceType)
	fcr += fmt.Sprintf("Unit TL      : %v (%v)\n", f.ForceTL, f.DM("TL"))
	fcr += fmt.Sprintf("Mobility Type: %v (%v)\n", f.Mobility, f.DM("Mobility"))
	fcr += fmt.Sprintf("Combat       : %v (%v)\n", f.CBAS[0], f.DM("Combat"))
	fcr += fmt.Sprintf("Bombardment  : %v (%v)\n", f.CBAS[1], f.DM("Bombardment"))
	fcr += fmt.Sprintf("Aerospace    : %v (%v)\n", f.CBAS[2], f.DM("Aerospace"))
	fcr += fmt.Sprintf("Support      : %v (%v)\n", f.CBAS[3], f.DM("Support"))
	fcr += fmt.Sprintf("CEI or DEI   : %v\n", f.CEI)
	fcr += fmt.Sprintf("Morale       : %v\n", f.Morale)
	fcr += fmt.Sprintf("Reputation   : %v\n", f.Reputation)
	return fcr
}
