package main

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/ehex"
)

func main() {
	eh := ehex.New().Set("G")
	fmt.Println(eh, &eh)
	eh2 := ehex.New().Set(12)
	fmt.Println(eh2, &eh2)
	fmt.Println(eh2, eh2.Value())
	fmt.Println("//////////")
	s := ehex.New().Set(8)
	d := ehex.New().Set("9")
	e := ehex.New().Set("A")
	i := ehex.New().Set(11)
	ed := ehex.New().Set(12)
	soc := ehex.New().Set("D")

	fmt.Println(ehex.Profile(s, d, e, i, ed, soc))
}
