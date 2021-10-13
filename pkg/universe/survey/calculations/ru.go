package calculations

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
)

func RU(econ string) int {
	hex := strings.Split(econ, "")
	r := ehex.New(hex[1])
	l := ehex.New(hex[2])
	i := ehex.New(hex[3])
	e, _ := strconv.Atoi(hex[4] + hex[5])
	return r.Value() * l.Value() * i.Value() * e
}
