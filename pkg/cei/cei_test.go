package cei

import (
	"fmt"
	"testing"
)

func TestCEI(t *testing.T) {
	cei := New(9)
	fmt.Println("Mission: Search planet side location:")
	fmt.Println(cei.ResolveMission("Operation: Land shuttle"))
	fmt.Println(cei.ResolveMission("Operation: Find source and investigate at ground level"))
	fmt.Println(cei.ResolveMission("Operation: Retrive Artefect"))
	fmt.Println(cei.ResolveMission("Operation: Obtain any additional information"))
}
