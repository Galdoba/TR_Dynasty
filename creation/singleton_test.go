package creation

import "testing"

func TestGetInstance(t *testing.T) {
	count1 := GetInstance()
	if count1 == nil {
		t.Error("expected pointer to singleton after calling GetInstance(), not nil")
	}
	expected := count1
	currentCount := count1.AddOne()
	if currentCount != 1 {
		t.Error("After calling for the first time to count, the count must be 1 but it's %d\n", currentCount)
	}
	count2 := GetInstance()
	if count2 != expected {
		t.Error("Expected same instance as counter2, but it's different")
	}
}
