package hyperjump

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Galdoba/devtools/cli/features"
	"github.com/Galdoba/devtools/cli/user"
)

const (
	typingDelay = "3ms"
)

var delay time.Duration
var emmersiveMode bool

func StartJumpEvent() {
	emmersiveMode = true
	del, err := time.ParseDuration(typingDelay)
	if err != nil {
		fmt.Println(err.Error())
	}
	delay = del
	effA := userInputInt("Astrogation Check effect: ")
	effE := userInputInt("Engineering Check effect: ")
	hj := New(effA, effE)
	printSlow(hj.Report())
	printSlow(hj.Outcome())
}

func userInputInt(msg ...string) int {
	str := userInputStr(msg...)
	i, err := strconv.Atoi(str)
	for err != nil {
		printSlow(err.Error() + "\n")
		str = userInputStr(msg...)
		i, err = strconv.Atoi(str)
	}
	return i
}

func userInputStr(msg ...string) string {
	for i := range msg {
		printSlow(msg[i])
	}
	str, err := user.InputStr()
	if err != nil {
		printSlow(err.Error())
		return err.Error() + "\n"
	}
	return str
}

func printSlow(text string) {
	if emmersiveMode {
		features.TypingSlowly(text, delay)
	} else {
		fmt.Print(text)
	}
}
