package weapons

import (
	"github.com/Galdoba/TR_Dynasty/pkg/core/qrebs"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type weapon struct {
	wType        string
	wSubType     string
	wDescriptor  string
	wBurden      string
	wStage       string
	wUser        string
	wPortability string
	wCode        [7]string
	qrebsStruct  qrebs.EvaluationData
	shortName    string
	longName     string
	tl           int
	wRange       int
	mass         float64
	damageDice   int
	cost         float64
	traits       string
	description  string
}

type WPC struct {
	autoMod bool
	dp      *dice.Dicepool
}

func NewWeaponConstructor() WPC {
	wpc := WPC{}
	wpc.dp = dice.New()
	return wpc
}

func (wpc *WPC) SetAutoMod(am bool) {
	wpc.autoMod = am
}

func (wpc *WPC) SetSeed(seed int) {
	wpc.dp.SetSeed(seed)
}

func New(wpc WPC) weapon {
	wp := weapon{}
	wp.selectType(wpc)
	return wp
}

func (wp *weapon) selectType(wpc WPC) {
	//categoryOptions := []string{}

}
