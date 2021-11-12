package mercenary

const (
	DM_Surprise_Technological = "Technological Surprise"
	DM_Surprise_Tactical      = "Tactical Surprise"
	DM_Surprise_Strategic     = "Strategic Surprise"
)

type CombatEmulation struct {
	travellers           *Force
	enemy                *Force
	combatDM             map[string]int
	reconissanceDetected bool
}

func (ce *CombatEmulation) Resolve() {
	//Phase 1: Reconisance
	// Decide if happens
	// 	Yes:
	//		Reconissance Detection
	//		Reconissance Outcome
	// Intelligence Event
	// CEI check 6,8,or 10
	//Phase 2: Preparation
	//Phase 3: Initial Combat
	//Phase 4: Subsequant Combat
	//Phase 5: Resolution
}

func IntelligenceEvent(effect int) int {
	switch effect {
	case -5, -4:
		return -3
	case -3, -2, -1:
		return -2
	case 0:
		return -1
	case 1, 2, 3:
		return 1
	case 4, 5:
		return 2
	}
	if effect <= -6 {
		return -4
	}
	return 3
}
