package actions

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/mercenary"
)

func NewCombatExtended() error {
	trvForce, err := mercenary.NewForce("Test")
	fmt.Println(trvForce.CapabilityRecord())
	fmt.Println(err)
	return nil
}
