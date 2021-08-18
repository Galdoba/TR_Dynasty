package ehex

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	expected := ehex{
		value:   0,
		code:    "*",
		comment: ANY_VALUE,
	}
	ehex := New()
	if ehex.code != expected.code {
		t.Errorf("expected ehex.code == %v (have %v)", ehex.code, expected.code)
	}
	if ehex.value != expected.value {
		t.Errorf("expected ehex.code == %v (have %v)", ehex.value, expected.value)
	}
	if ehex.comment != expected.comment {
		t.Errorf("expected ehex.code == %v (have %v)", ehex.comment, expected.comment)
	}
}

type testTableSetCodes struct {
	inputData string
	expected  ehex
}

func TestSetCodes(t *testing.T) {
	ehexObj := New()
	testTable := []testTableSetCodes{
		{
			inputData: "a",
			expected:  ehex{value: 10, code: "A"},
		},
		{
			inputData: "A",
			expected:  ehex{value: 10, code: "A"},
		},
		{
			inputData: "ф",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "%",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "AB",
			expected:  ehex{value: -1, code: ""},
		},
		{
			inputData: "",
			expected:  ehex{value: -1, code: ""},
		},
	}
	////////////
	for i, val := range testTable {
		ehexObj.Set(val.inputData)
		if !match(*ehexObj, val.expected) {
			t.Errorf("input %v: (%v) No Match: \nexecuted: %v\nexpected: %v\n", i+1, val.inputData, ehexObj.printStruct(), val.expected.printStruct())
		}
	}
}

func (ehex *ehex) printStruct() string {
	return fmt.Sprintf("ehex{code: '%v', value: '%v', comment: '%v'}", ehex.code, ehex.value, ehex.comment)
}

func match(ehex1, ehex2 ehex) bool {
	if ehex1.printStruct() == ehex2.printStruct() {
		return true
	}
	return false
}
