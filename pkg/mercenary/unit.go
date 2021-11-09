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
	ForceEquipment string
	ActivePersonal int
	CBAS           []int
	Mobility       string
	CEI            int
	ECEI           int
	Morale         int
	Reputation     int
}

func NewForce(name string, personel, tl int) (*Force, error) {
	f := Force{}
	f.ActivePersonal = personel
	switch forceFormation(personel) {
	default:
		return &f, fmt.Errorf("Force '%v' cannot be created with %v personal", name, personel)
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
	f.Name = name + " " + f.ForceSize
	f.ForceTL = tl
	f.ForceType = CBAS_COMBAT            //setForceType()
	f.ForceTraining = TRAINING_Trained   //setForceTraining()
	f.ForceEquipment = EQIPMENT_Generous //setForceEqupment()
	//f.setupCBAS()
	err := fmt.Errorf("Not implemented")
	return &f, err
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
	switch f.ForceTraining {
	case TRAINING_Untrained:
		for i := range f.CBAS {
			f.CBAS[i] = f.CBAS[i] + f.trainingEffect()
		}
	case TRAINING_Raw:
		for i := range f.CBAS {
			f.CBAS[i] = f.CBAS[i] + f.trainingEffect()
		}
	case TRAINING_Green:
		for i := range f.CBAS {
			f.CBAS[i] = f.CBAS[i] + f.trainingEffect()
		}
	case TRAINING_Trained:
		for i := range f.CBAS {
			f.CBAS[i] = f.CBAS[i] + f.trainingEffect()
		}
	case TRAINING_Effective:
		for i := range f.CBAS {
			f.CBAS[i] = f.CBAS[i] + f.trainingEffect()
		}
	case TRAINING_HighlyEffective:
		for i := range f.CBAS {
			f.CBAS[i] = f.CBAS[i] + f.trainingEffect()
		}
	}
	switch f.ForceType {
	case CBAS_COMBAT:
		f.CBAS[0] = f.CBAS[0] + f.equipmentEffect()
	case CBAS_BOMBARDMENT:
		f.CBAS[1] = f.CBAS[1] + f.equipmentEffect()
	case CBAS_AEROSPACE:
		f.CBAS[2] = f.CBAS[2] + f.equipmentEffect()
	case CBAS_SUPPORT:
		f.CBAS[3] = f.CBAS[3] + f.equipmentEffect()
	}
	for i := range f.CBAS {
		if f.CBAS[i] < 0 {
			f.CBAS[i] = 0
		}
	}
}

func (f *Force) equipmentEffect() int {
	switch f.ForceEquipment {
	case EQIPMENT_Minimal:
		return 0
	case EQIPMENT_Sparce:
		return 1
	case EQIPMENT_Basic:
		return 2
	case EQIPMENT_Standard:
		return 3
	case EQIPMENT_Generous:
		return 4
	case EQIPMENT_Lavish:
		return 5
	case EQIPMENT_Execive:
		return 6
	}
	return -999
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
		if chosen, err = user.ChooseOneStr("Select Force Type:", []string{"Untrained", "Raw", "Green", "Trained", "Effective", "Highly Effective"}); err == nil {
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

func (f *Force) DM(field string) int {
	switch field {
	case "TL":
		switch {
		case f.ForceTL < 1:
			return -3
		case f.ForceTL == 1 || f.ForceTL == 2:
			return -2
		}
		return f.ForceTL/3 - 2
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
