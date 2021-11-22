package mission

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/cei"
)

func Test_Mission(t *testing.T) {
	crew := cei.NewTeam("Crew", 7)
	task := NewOperation("Eat this stuff", NewModifier("Hungry", -1))
	task.AssignResolver(crew)
	if err := task.AbstractResolve(); err != nil {
		t.Errorf("task not resolved: %v (not expected)\n", err.Error())
	}
	fmt.Println(task)
}
