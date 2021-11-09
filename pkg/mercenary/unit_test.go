package mercenary

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

func TestForce(t *testing.T) {
	f, err := NewForce("Test", 28, 14)

	f.ForceType = dice.New().RollFromList([]string{CBAS_COMBAT, CBAS_BOMBARDMENT, CBAS_AEROSPACE, CBAS_SUPPORT})
	f.ForceTraining = dice.New().RollFromList([]string{TRAINING_Untrained, TRAINING_Raw, TRAINING_Green, TRAINING_Trained, TRAINING_Effective, TRAINING_HighlyEffective})
	f.ForceEquipment = dice.New().RollFromList([]string{EQIPMENT_Minimal, EQIPMENT_Sparce, EQIPMENT_Basic, EQIPMENT_Standard, EQIPMENT_Generous, EQIPMENT_Lavish, EQIPMENT_Execive})
	f.setupCBAS()
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}
	fmt.Println(f)
	//fmt.Println(f) test

}
