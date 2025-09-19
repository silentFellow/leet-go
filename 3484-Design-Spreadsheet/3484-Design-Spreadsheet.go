package leetcode

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	spreadSheet map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		spreadSheet: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.spreadSheet[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.spreadSheet, cell)
}

func (this *Spreadsheet) GetValue(formula string) int {
	splitRef := strings.Split(formula[1:], "+")
	f, s := splitRef[0], splitRef[1]

	fv, err1 := strconv.Atoi(f)
	sv, err2 := strconv.Atoi(s)

	if err1 != nil {
		fv = this.spreadSheet[f]
	}
	if err2 != nil {
		sv = this.spreadSheet[s]
	}

	return fv + sv
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
