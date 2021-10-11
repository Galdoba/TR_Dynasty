package calculations

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
)

func Importance(starport, tl, pops, bases, remarks string) int {
	b := 0
	if strings.Contains(bases, "W") {
		b++
	}
	if strings.Contains(bases, "KV") {
		b++
	}
	if strings.Contains(bases, "NS") {
		b++
	}
	r := 0
	if strings.Contains(remarks, "Ag") {
		r++
	}
	if strings.Contains(remarks, "Hi") {
		r++
	}
	if strings.Contains(remarks, "In") {
		r++
	}
	if strings.Contains(remarks, "Ri") {
		r++
	}
	s := 0
	switch starport {
	case "A", "B":
		s++
	case "D", "E", "X", "G", "H", "Y":
		s--
	}
	t := 0
	tlH := ehex.New(tl)
	if tlH.Value() >= 16 {
		t++
	}
	if tlH.Value() >= 10 {
		t++
	}
	if tlH.Value() <= 8 {
		t--
	}
	p := 0
	pH := ehex.New(pops)
	if pH.Value() <= 6 {
		p--
	}
	return b + r + s + t + p
}
