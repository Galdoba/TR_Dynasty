package qrebs

import (
	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

//EQUIPMENT EVALUATION SYSTEM
// 5 0 0 0 0
// 5 0-1 2 0
// 50A20
// -00CF

const (
	QUALITY     = 1
	RELIABILITY = 2
	EASEOFUSE   = 3
	BURDEN      = 4
	BULK        = 4
	SAFETY      = 5
)

type EvaluationData struct {
	data map[int]int
}

func New() EvaluationData {
	ed := EvaluationData{}
	ed.data = make(map[int]int)
	return ed
}

func Standard() EvaluationData {
	ed := EvaluationData{}
	ed.data = make(map[int]int)
	ed.data[QUALITY] = 5
	return ed
}

func Custom(q, r, e, b, s int) EvaluationData {
	ed := EvaluationData{}
	ed.data = make(map[int]int)
	ed.data[QUALITY] = q
	ed.data[RELIABILITY] = r
	ed.data[EASEOFUSE] = e
	ed.data[BURDEN] = b
	ed.data[SAFETY] = s
	return ed
}

func (ed *EvaluationData) Change(dataID, value int) {
	ed.data[dataID] = ed.data[dataID] + value
}

//By Stat//////////////////////////////////

func (ed *EvaluationData) Quality() string {
	str := ehex.New(ed.data[1]).String()
	if ed.data[1] > 0 {
		return str
	}
	return "-"
}

func (ed *EvaluationData) QualityI() int {
	return ed.data[1]
}

func (ed *EvaluationData) Reliability() string {
	return AsMod(ehex.New(ed.data[2]))
}

func (ed *EvaluationData) EaseOfUse() string {
	return AsMod(ehex.New(ed.data[3]))
}

func (ed *EvaluationData) Burden() string {
	return AsMod(ehex.New(ed.data[4]))
}

func (ed *EvaluationData) Safety() string {
	return AsMod(ehex.New(ed.data[5]))
}

func (ed *EvaluationData) String() string {
	str := ""
	str += ed.Quality()
	str += ed.Reliability()
	str += ed.EaseOfUse()
	str += ed.Burden()
	str += ed.Safety()
	return str
}

func AsMod(e ehex.DataRetriver) string {
	val := e.Value()
	if val < 0 {
		val = (val * -1) + 10
		return ehex.New(val).String()
	}
	if val > 9 {
		return "?"
	}
	return ehex.New(val).String()
}

func (ed *EvaluationData) Random() {
	ed.data[1] = dice.Roll2D() - 2
	ed.data[2] = dice.Flux()
	ed.data[3] = dice.Flux()
	ed.data[4] = dice.Flux()
	ed.data[5] = dice.Flux()

}
