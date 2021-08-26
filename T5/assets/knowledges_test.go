package assets

import (
	"testing"
)

func TestKnowledge(t *testing.T) {
	for i := KNOWLEDGE_Rider; i <= KNOWLEDGE_Specialized; i++ {
		kn := NewKnowledge(i)
		if kn.err != nil {
			t.Errorf("creation error: %v", kn.err.Error())
		} else {

		}
	}
}
