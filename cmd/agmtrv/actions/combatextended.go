package actions

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/mercenary"
)

func NewCombatExtended() error {
	trvForce, errT := mercenary.NewForce()
	enemyForce, errE := mercenary.NewForce()
	fmt.Println(trvForce.CapabilityRecord())
	fmt.Println(enemyForce.CapabilityRecord())
	fmt.Println(errT)
	fmt.Println(errE)
	return nil
}
