package routine

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Galdoba/devtools/cli/user"
)

func userConfirm(q string) bool {
	err := errors.New("No answer")
	answer := ""
	for err != nil {
		fmt.Print(q + "(y/n): ")
		answer, err = user.InputStr()
		answer = strings.ToUpper(answer)
		switch answer {
		default:
			err = errors.New("Answer not clear")
		case "Y", "YES", "1", "Н":
			return true
		case "N", "NO", "0", "Т":
			return false
		}
		reportErr(err)
	}
	return false
}

func reportErr(err error) {
	if err != nil {
		fmt.Print(err.Error() + " \n")
	}
}