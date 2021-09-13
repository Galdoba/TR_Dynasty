package tasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

const (
	TASK_COMMENT_Cooperative      = "Cooperative"
	TASK_COMMENT_Uncertain        = "Uncertain"
	TASK_COMMENT_Opposed          = "Opposed"
	TASK_COMMENT_Hasty            = "Hasty"
	TASK_COMMENT_ExtraHasty       = "Extra Hasty"
	TASK_COMMENT_Cautious         = "Cautious"
	TASK_COMMENT_ThisIsHard       = "This is Hard"
	TASK_DURATION_Unimportant     = 0
	TASK_DURATION_Unimportant_str = ""
	TASK_DURATION_Minutes         = 1
	TASK_DURATION_Minutes_str     = "About 10 minutes"
	TASK_DURATION_Hour            = 2
	TASK_DURATION_Hour_str        = "An Hour"
	TASK_DURATION_Day             = 3
	TASK_DURATION_Day_str         = "All Day"
	TASK_DURATION_Week            = 4
	TASK_DURATION_Week_str        = "A Week"
	TASK_DURATION_Month           = 5
	TASK_DURATION_Month_str       = "A Month"
)

type Task struct {
	dice           *dice.Dicepool
	skill          TaskAsset
	characteristic TaskAsset
	mods           []TaskAsset
	duration       int
	difficulty     int
	purpose        string
	comments       []string
	resolution     string
	spectacular    string
	completed      bool
}

type TaskAsset interface {
	Name() string
	Value() int
	Code() int
}

func Create() *Task {
	t := Task{}
	t.resolution = "Unresolved"
	return &t
}

func (t *Task) SetupEnviroment(purpose string, dif int, timeframe int, comments ...string) error {
	if ContainsAll(comments, TASK_COMMENT_Hasty, TASK_COMMENT_Cautious) {
		return fmt.Errorf("test can not be Hasty and Cautious at the same time")
	}
	if ContainsAll(comments, TASK_COMMENT_ExtraHasty, TASK_COMMENT_Cautious) {
		return fmt.Errorf("test can not be Extra Hasty and Cautious at the same time")
	}
	///
	t.purpose = purpose
	t.difficulty = dif
	t.duration = timeframe
	for _, val := range comments {
		t.comments = append(t.comments, val)
		if val == TASK_COMMENT_Hasty {
			t.difficulty++
		}
		if val == TASK_COMMENT_ExtraHasty {
			t.difficulty += 2
		}
		if val == TASK_COMMENT_Cautious {
			t.difficulty--
		}
	}
	return nil
}

func (t *Task) SetupAssets(asset ...TaskAsset) error {
	for _, val := range asset {
		t.AddAsset(val)
	}
	return nil
}

func ContainsAll(sl []string, elements ...string) bool {
	for _, val := range elements {
		if !utils.ListContains(sl, val) {
			return false
		}
	}
	return true
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
	switch sum <= tn {
	case true:
		t.resolution = "Successful"
		t.completed = true
	case false:
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

func (t *Task) Completed() bool {
	return t.completed
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
	if t.duration != 0 {
		res += fmt.Sprintf("Task took about %v\n", determineDuration(t.duration, t.comments...))
	}
	switch sum > tn {
	case true:
		res += "Task Failed\n"
	case false:
		res += "Task Successful\n"
	}
	if t.spectacular != "" {
		res += t.spectacular + "\n"
	}
	res = strings.TrimSuffix(res, "\n")
	return res
}

func (t *Task) TaskPhrase() string {
	ph := "===TASK PHRASE==================================================================\n"
	ph += "To " + t.purpose
	if t.duration != 0 {
		tmap := make(map[int]string)
		tmap[TASK_DURATION_Minutes] = TASK_DURATION_Minutes_str
		tmap[TASK_DURATION_Hour] = TASK_DURATION_Hour_str
		tmap[TASK_DURATION_Day] = TASK_DURATION_Day_str
		tmap[TASK_DURATION_Week] = TASK_DURATION_Week_str
		tmap[TASK_DURATION_Month] = TASK_DURATION_Month_str
		ph += " [ " + tmap[t.duration] + " ]"
	}
	ph += "\n" + t.difficultyStr() + " <= "
	ph += t.listAssets() + " " + t.listAssetsValues() + "\n"
	c := len(t.comments)
	if c > 1 {
		for i := c; i > 0; c-- {
			ph += t.comments[c] + "\n"
		}
	}
	ph += "================================================================================"
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

type InstructionProcess interface {
	SetDifficulty(int) *Task
}

type instrucionS struct {
	code int
}

func determineDuration(dFactor int, comments ...string) string {
	time := 0
	result := ""
	switch dFactor {
	case 1:
		time = 10 + dice.Flux()
		result = "minutes"
	case 2:
		time = 60 + dice.Flux()*10
		result = "minutes"
	case 3:
		time = 10 + dice.Flux()
		result = "hours"
	case 4:
		time = 6 + dice.Flux()
		result = "days"
	case 5:
		time = 6 + dice.Flux()
		result = "weeks"
	}
	for _, val := range comments {
		if val == TASK_COMMENT_Cautious {
			time = time * 2
		}
		if val == TASK_COMMENT_Hasty {
			time = time / 2
		}
		if val == TASK_COMMENT_ExtraHasty {
			time = time / 4
		}
	}
	return fmt.Sprintf("%d %v", time, result)
}

/*
tm := task.NewTaskMaker()
task := task.Create(
	tm.SetDifficulty(3),
	tm.SetPurpose("Apply to university"),
)

newTask := task.NewTask()
newTask = task.Construct(
	newTask.SetDifficulty(3),
	newTask.SetPurpose("Apply to university"),
)

newTask := task.Task{
	Difficulty: 3
	Purpose: "Apply to university"
}


*/
