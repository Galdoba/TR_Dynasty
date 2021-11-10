package mercenary

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func TestForce(t *testing.T) {
	//f, err := NewForce("Test", 28, 14)
	f := Force{}
	f.Name = "Test"
	f.ActivePersonal = 30 //dice.New().RollNext("1d200").Sum()
	f.ForceTL = 14        //dice.New().RollNext("1d16").DM(-1).Sum()
	f.ForceSize = f.forceFormationStr()
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Generous)
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Minimal)
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Minimal)
	f.ForceEquipment = append(f.ForceEquipment, EQIPMENT_Basic)
	f.ForceType = CBAS_COMBAT //dice.New().RollFromList([]string{CBAS_COMBAT, CBAS_BOMBARDMENT, CBAS_AEROSPACE, CBAS_SUPPORT})
	f.ForceTraining = dice.New().RollFromList([]string{TRAINING_Untrained, TRAINING_Raw, TRAINING_Green, TRAINING_Trained, TRAINING_Effective, TRAINING_HighlyEffective})
	//f.ForceEquipment[f.primaryCBAS()] = dice.New().RollFromList([]string{EQIPMENT_Minimal, EQIPMENT_Sparce, EQIPMENT_Basic, EQIPMENT_Standard, EQIPMENT_Generous, EQIPMENT_Lavish, EQIPMENT_Execive})
	f.Mobility = MOBILITY_Infantry //dice.New().RollFromList([]string{MOBILITY_Static, MOBILITY_SemiMobile, MOBILITY_Infantry, MOBILITY_Mechanised, MOBILITY_Mounted, MOBILITY_Motorised, MOBILITY_AirMobile, MOBILITY_GravMobile, MOBILITY_Aerospace})
	f.setupCBAS()
	// if err != nil {
	// 	t.Errorf("error: %v", err.Error())
	// }
	fmt.Println(f.CapabilityRecord())

	fmt.Println(f)

}
