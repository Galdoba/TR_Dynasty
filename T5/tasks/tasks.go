package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type Task struct {
	dice           *dice.Dicepool
	skill          TaskAsset
	characteristic TaskAsset
	mods           []TaskAsset
	duration       string
	difficulty     int
	purpose        string
	comments       []string
	resolution     string
	spectacular    string
}

type TaskAsset interface {
	Name() string
	Value() int
	Code() int
}

func Create(f ...func(*Task)) *Task {
	t := Task{}
	t.resolution = "Unresolved"

	return &t
}

func (t *Task) ApplyInstrucuctions(funcList ...func(vals ...interface{}) *Task) {
	for i, _ := range funcList {
		//TODO: найти способ скармливать кучу разных функций в поле и вызывать их исполнение
	}
}

type Mod struct {
	name  string
	value int
	code  int
}

func NewMod(name string, value int) *Mod {
	return &Mod{name, value, -99}
}

func (m *Mod) Name() string {
	return m.name
}

func (m *Mod) Value() int {
	return m.value
}

func (m *Mod) Code() int {
	return -99
}

func (t *Task) AddAsset(ta TaskAsset) {
	for _, code := range assets.ListCharacteristics() {
		if code == ta.Code() {
			t.characteristic = ta
			return
		}
	}
	for _, code := range assets.ListSkills() {
		if code == ta.Code() {
			t.skill = ta
			return
		}
	}
	t.mods = append(t.mods, ta)
}

func (t *Task) SetDice(d *dice.Dicepool) {
	t.dice = d
}

func (t *Task) SetPurpose(p string) {
	t.purpose = p
}

func (t *Task) SetDifficulty(dif int) {
	t.difficulty = dif
}

func (t *Task) TargetNumber() int {
	tn := 0
	if t.characteristic != nil {
		tn += t.characteristic.Value()
	}
	if t.skill != nil {
		tn += t.skill.Value()
	}
	for _, mod := range t.mods {
		tn += mod.Value()
	}
	return tn
}

func (t *Task) Resolve() string {
	dif := strconv.Itoa(t.difficulty) + "d6"
	if t.dice == nil {
		t.dice = dice.New()
	}
	result := t.dice.RollNext(dif).Result()
	ones := 0
	sixes := 0
	for _, v := range result {
		if v == 1 {
			ones++
		}
	}
	for _, v := range result {
		if v == 6 {
			sixes++
		}
	}
	tn := t.TargetNumber()
	sum := 0
	for _, v := range result {
		sum += v
	}
	if sum <= tn {
		t.resolution = "Successful"
	} else {
		t.resolution = "Failed"
	}

	if ones >= 3 {
		t.spectacular = "Spectacular Success"
	}
	if sixes >= 3 {
		t.spectacular = "Spectacular Failure"
	}
	if ones >= 3 && sixes >= 3 {
		t.spectacular = "Spectacularly Interesting"
	}
	t.resolution = t.Outcome()

	return t.resolution
}

func (t *Task) Outcome() string {

	res := ""
	result := t.dice.Result()
	tn := t.TargetNumber()
	sum := 0
	for _, v := range result {
		sum += v
	}

	res = fmt.Sprint("Rolled ", result, " = ", sum, " against ", tn, "\n")
	return res
}

func (t *Task) TaskPhrase() string {
	ph := "===TASK PHRASE==================================================================\n"
	ph += "To " + t.purpose
	if t.duration != "" {
		ph += "[" + t.duration + "]"
	}
	ph += "\n" + t.difficultyStr() + " <= "
	ph += t.listAssets() + " " + t.listAssetsValues() + "\n"
	for c := len(t.comments); c > 0; c-- {
		ph += t.comments[c] + "\n"
	}
	ph += "================================================================================\n"
	return ph
}

func (t *Task) listAssets() string {
	la := ""
	if t.characteristic != nil {
		la += t.characteristic.Name() + " + "
	}
	if t.skill != nil {
		la += t.skill.Name() + " + "
	}
	for _, mod := range t.mods {
		la += mod.Name() + " + "
	}
	return strings.TrimSuffix(la, " + ")
}

func (t *Task) listAssetsValues() string {
	la := "["
	if t.characteristic != nil {
		la += fmt.Sprint(t.characteristic.Value(), " ")
	}
	if t.skill != nil {
		la += fmt.Sprint(t.skill.Value(), " ")
	}
	for _, mod := range t.mods {
		la += fmt.Sprint(mod.Value(), " ")
	}
	la = strings.TrimSuffix(la, " ")
	return la + "]"
}

func (t *Task) difficultyStr() string {
	switch t.difficulty {
	default:
		return "Unknown (?D)"
	case 1:
		return "Easy (1D)"
	case 2:
		return "Average (2D)"
	case 3:
		return "Difficult (3D)"
	case 4:
		return "Formidable (4D)"
	case 5:
		return "Staggering (5D)"
	case 6:
		return "Hopeless (6D)"
	case 7:
		return "Impossible (7D)"
	case 8:
		return "Beyond Impossible (8D)"

	}
}
