package weapons

import "github.com/Galdoba/TR_Dynasty/dice"

type FillForm struct {
	drs []weaponData
	// Type
	// SubType
	// Descriptor
	// Burden
	// Stage
	// User
	// Notes
	// Options
	// Controls
	// Portability
	// QREBS
	// Totals
	//
	// Model
	// TL
	// Range
	// Mass
	// Burden
	// H1
	// H2
	// D1
	// D2
	// H3
	// D3
	// Cost
}

type weaponData struct {
	dModel  string //code
	dTL     int    //
	dRange  int    //
	dMass   float64
	dBurden int
	dH1     string
	dH2     string
	dD1     int
	dD2     int
	dH3     string
	dD3     int
	dCost   int
}

type weapon struct{}

type designTask struct{}

type gunMaker struct {
	dp *dice.Dicepool //дайспул который используется для бросков
}

type Maker interface { //формирует функции задания для дизайна продукта
	Apply(string, string) designTask
	Roll(string) designTask
	ConcludeDesignProcess() designTask
}

/*

lib maker
SAMPLE CODE:

gunMkr := maker.NewGunMaker()
wp := gunMkr.MakeWeapon(
	gunMkr.Apply(item1, descr1)
	gunMkr.Apply(item2, descr2)
	gunMkr.Roll(item3, item4 ...)
	gunMkr.ConcludeDesignProcess()
)
wp.Model()
wp.LongName()
wp.Fire(effectDM)
wp.DoStuff(args...)


*/
