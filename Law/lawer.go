package law

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
)

const (
	lawCheck       = "Check"
	lawInvestigate = "Investigate"
	lawApprehend   = "Apprehend"
	lawSentence    = "Sentence"
)

//RelationsMgt2Core - каркас из корника второй части
type RelationsMgt2Core struct {
	localUWP string
	status   int
	crimes   []crime // или int
	dm       int
}

type crime struct {
	situation string
	dm        int
	responce  string
}

func newCrime(sit, resp string, dm int) crime {
	c := crime{}
	c.situation = sit
	c.dm = dm
	c.responce = resp
	return c
}

func Sentense(uwp string, cr crime) {
	fmt.Println("law TN", lawTN(uwp))
	s := sent(dice.Roll("2d6").DM(cr.dm).Sum())
	fmt.Println(s)

}

func sent(i int) string {
	sent := ""
	switch i {
	default:
		if i <= 0 {
			sent = "Dismissed or trivial"
		}
		if i >= 15 {
			sent = "Death"
		}
	case 1, 2:
		sent = "Fine of " + strconv.Itoa(dice.Roll1D()*1000) + " (per ton of cargo)"
	case 3, 4:
		sent = "Fine of " + strconv.Itoa(dice.Roll2D()*5000) + " (per ton of cargo)"
	case 5, 6:
		sent = "Exile or a fine of " + strconv.Itoa(dice.Roll2D()*10000) + " (per ton of cargo)"
	case 7, 8:
		sent = "Imprisonment for " + dice.Roll("1d6").SumStr() + " months or exile or fine of " + strconv.Itoa(dice.Roll2D()*20000) + " (per ton of cargo)"
	case 9, 10:
		sent = "Imprisonment for " + dice.Roll("1d6").SumStr() + " years or exile"
	case 11, 12:
		sent = "Imprisonment for " + dice.Roll("2d6").SumStr() + " years or exile"
	case 13, 14:
		sent = "Life imprisonment"

	}
	return sent
}

//Lawer - оператор
type Lawer interface {
	Interact() string
}

//NewRelations - новый каркас
func NewRelations(uwp string) RelationsMgt2Core {
	lr := RelationsMgt2Core{}
	lr.localUWP = uwp
	lr = lr.Accuse(newCrime("First Aproach", lawCheck, 0))
	return lr
}

func (lr RelationsMgt2Core) сhangeDM(add int) RelationsMgt2Core {
	lr.dm = lr.dm + add
	return lr
}

//Accuse - добавляет к новое Преступление для обработки
func (lr RelationsMgt2Core) Accuse(c crime) RelationsMgt2Core {
	lr.crimes = append(lr.crimes, c)
	return lr
}

//helpers
func localLaw(uwp string) string {
	return string([]byte(uwp)[6])
}

func lawTN(uwp string) int {
	return TrvCore.EhexToDigit(localLaw(uwp))
}

func (lr RelationsMgt2Core) Check(dm ...int) bool {
	fullDm := 0
	for i := range dm {
		fullDm += dm[i]
	}
	r := dice.Roll2D()
	lawTN := lawTN(lr.localUWP)
	if r <= lawTN {
		fmt.Println("Check Successfull: Beat Admin or Streetwise challenge with tn =", 6+(lawTN-r), "to avoid Investigation")
		return true
	}
	return false
}
