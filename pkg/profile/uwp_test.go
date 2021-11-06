package profile

import (
	"fmt"
	"testing"
)

func testFeed() []string {
	return []string{
		"A123456-7",
		"A123456-7 ",
		"A1234567",
		"A 1 2 3 4 5 6 - 7 ",
		"_123___-_",
		"Akkk3456-7",
		"A1256-7",
	}
}

func TestUWP(t *testing.T) {
	for _, feed := range testFeed() {
		uwp := NewUWP(feed)
		fmt.Printf("Test feed %v completed: %v\n", feed, uwp)
		fmt.Printf("Error: %v\n", uwp.err.Error())
	}
}
