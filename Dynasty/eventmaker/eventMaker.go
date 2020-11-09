package eventmaker

import (
	"time"

	dynasty "github.com/Galdoba/TR_Dynasty/Dynasty"
	"github.com/Galdoba/TR_Dynasty/dice"
)

/*
событие:		// умеет менять данные
время старта	// данные
длительность	// данные
субъект			// умеет Выбирать объект и Решать
объект			// данные
имя				// данные
описание		// данные
метод решения	// данные - скорее функция testFormula

исход			// данные



*/

//event - [SET DESCRIPTION OF OBJECT HERE]
type event struct {
	outcome          int
	object           *dynasty.Dynasty
	startDate        int
	name             string
	description      string
	resolutionEffect int
	subject          *dynasty.Dynasty
	lenght           int
	outcomeDescr     string
	resolutionTest   []string
}

//Subject - returns e.subject as a *dynasty.Dynasty
func (e *event) Subject() *dynasty.Dynasty {
	return e.subject
}

//Lenght - returns e.lenght as a int
func (e *event) Lenght() int {
	return e.lenght
}

//OutcomeDescr - returns e.outcomeDescr as a string
func (e *event) OutcomeDescr() string {
	return e.outcomeDescr
}

//ResolutionTest - returns e.resolutionTest as a []string
func (e *event) ResolutionTest() []string {
	return e.resolutionTest
}

//ResolutionEffect - returns e.resolutionEffect as a int
func (e *event) ResolutionEffect() int {
	return e.resolutionEffect
}

//Object - returns e.object as a *dynasty.Dynasty
func (e *event) Object() *dynasty.Dynasty {
	return e.object
}

//StartDate - returns e.startDate as a int
func (e *event) StartDate() int {
	return e.startDate
}

//Name - returns e.name as a string
func (e *event) Name() string {
	return e.name
}

//Description - returns e.description as a string
func (e *event) Description() string {
	return e.description
}

//Outcome - returns e.outcome as a int
func (e *event) Outcome() int {
	return e.outcome
}

//SetLenght - sets int value for e.lenght
func (e *event) SetLenght(data int) {
	e.lenght = data
}

//SetOutcomeDescr - sets string value for e.outcomeDescr
func (e *event) SetOutcomeDescr(data string) {
	e.outcomeDescr = data
}

//SetResolutionTest - sets []string value for e.resolutionTest
func (e *event) SetResolutionTest(data []string) {
	e.resolutionTest = data
}

//SetResolutionEffect - sets int value for e.resolutionEffect
func (e *event) SetResolutionEffect(data int) {
	e.resolutionEffect = data
}

//SetSubject - sets *dynasty.Dynasty value for e.subject
func (e *event) SetSubject(data *dynasty.Dynasty) {
	e.subject = data
}

//SetStartDate - sets int value for e.startDate
func (e *event) SetStartDate(data int) {
	e.startDate = data
}

//SetName - sets string value for e.name
func (e *event) SetName(data string) {
	e.name = data
}

//SetDescription - sets string value for e.description
func (e *event) SetDescription(data string) {
	e.description = data
}

//SetOutcome - sets int value for e.outcome
func (e *event) SetOutcome(data int) {
	e.outcome = data
}

//SetObject - sets *dynasty.Dynasty value for e.object
func (e *event) SetObject(data *dynasty.Dynasty) {
	e.object = data
}

///////////////////////////////////////////////
//RESOLUTION TEST

type Referee interface {
	Resolve() int
}

type conflict struct {
	testType    string //Opposed/Simple
	subjDM      int
	opposedDM   int
	difficulty  int
	diceFormula string
	engine      *dice.Dicepool
	timeFactor  int
	effect      int
	resolved    bool
	log         string
}

func SetupConflict(ev event) conflict {
	c := conflict{}

	return c
}

func (c *conflict) Resolve() {
	if c.resolved {
		return
	}
	if c.engine == nil {
		c.engine = dice.New(time.Now().UnixNano())
	}
	dm := c.subjDM - c.opposedDM
	r := c.engine.RollNext(c.diceFormula).DM(dm).Sum()
	c.effect = r - c.difficulty
	c.timeFactor = c.engine.RollNext("1d6").DM(c.effect / 2).Sum()
	if c.timeFactor < 1 {
		c.timeFactor = 1
	}
	c.resolved = true
}

//Invade another Dynasty’s territory and claim it: Militarism, 8–48 Months, Difficult (–2), Opposed by Tenacity.
/*
name		Invade another Dynasty’s territory and claim it	//string - u
apttitude	Conquest										//string - k
subjDM		Militarism										//string - k
time		8 Months										//string - p
difficulty	Difficult										//string - k
opposedDM 	Tenacity										//string - k

Public Relations  Militarism  10-60 Months + 70  VeryDifficult  Militarism  Invade another Dynasty’s territory and claim it

*/

type aptCheck struct {
	name        string
	sourceDM    string
	timeFrame   string
	difficulty  string
	oppDM       string
	gameEffects string
}
