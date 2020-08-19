package tab

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/Galdoba/devtools/cli/prettytable"
)

type Table struct {
	path     string
	lines    []string
	columns  []string
	colWidth []int
}

func NewTable(path string) (Table, error) {
	t := Table{}
	file, err := os.Open(path)
	if err != nil {
		return t, err
	}
	defer file.Close()
	colNum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t.lines = append(t.lines, scanner.Text())
		if colNum == 0 {
			cols := strings.Split(scanner.Text(), "	")
			colNum = len(cols)
		} else {
			if colNum != len(strings.Split(scanner.Text(), "	")) {
				return Table{}, errors.New("Table NOT formatted properly: " + scanner.Text())
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return t, err
	}
	t.updateColWidth()
	return t, nil
}

//CellValue - возвращает значение в ячейки
func (t Table) CellValue(r, c int) string {
	line := t.lines[r]
	if r >= len(t.lines) {
		return "{Error}"
	}
	cols := strings.Split(line, "	")
	if c >= len(cols) {
		return "{Error}"
	}
	return cols[c]
}

//Line - возвращает строку
func (t Table) Line(num int) string {
	return t.lines[num]
}

//ColWidths - Возвращает слайс длинн колонок таблицы
func (t Table) ColWidths() []int {
	return t.colWidth
}

func (t *Table) updateColWidth() {
	line0 := t.lines[0]
	col := strings.Split(line0, "	")
	for _, val := range col {
		t.colWidth = append(t.colWidth, len(val))
	}
	for _, val := range t.lines {
		col := strings.Split(val, "	")
		for c := range col {
			if t.colWidth[c] < len(col[c]) {
				t.colWidth[c] = len(col[c])
			}
		}
	}
}

func (t Table) PTPrint() {
	var ptSl [][]string
	for i := range t.lines {
		sl := strings.Split(t.lines[i], "	")
		ptSl = append(ptSl, sl)
	}
	pt := prettytable.From(ptSl)
	pt.PTPrint()
}
