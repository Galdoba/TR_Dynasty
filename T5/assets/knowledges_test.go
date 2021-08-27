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
			for j := 0; j < 6; j++ {
				if err := kn.Train(); err != nil {
					t.Errorf("training error: %v", err.Error())
				}
			}
		}
	}
}
