package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/core/skill"
)

func main() {
	for id := 0; id < 38; id++ {
		s, err := skill.New(id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
		fmt.Println("")
	}
}
