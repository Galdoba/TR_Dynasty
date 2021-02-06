package tab

import (
	"fmt"
	"strings"
)

type TabSeparatedTable struct {
	lines []string
}

func NewTST(lines []string) TabSeparatedTable {
	return TabSeparatedTable{lines}
}

func (tst *TabSeparatedTable) PrintTable() {
	//rows := len(tst.lines)
	width := []int{}
	head := strings.Split(tst.lines[0], "	")
	for _, val := range head {
		width = append(width, len(val))
	}
	for _, neck := range tst.lines {
		line := strings.Split(neck, "	")
		for c, val := range line {
			if width[c] < len(val) {
				width[c] = len(val)
			}

		}
	}
	for _, line := range tst.lines {
		body := strings.Split(line, "	")
		for r, val := range body {
			for len(val) < width[r] {
				val += " "
			}
			fmt.Print("| " + val + " ")
		}
		fmt.Print("|\n")
	}
}

func (tst *TabSeparatedTable) AddLine(line string) {
	tst.lines = append(tst.lines, line)
}
