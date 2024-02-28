package main

import (
	"github.com/MISELLO/sudoku-solver-go/solver"
)

var errMsg string

func main() {
	solver.Print()

	s, e := GetInput()

	if e {
		PrintErrMsg()
		PrintUsage()
	} else {
		solver.Load(s)
		s = solver.Unload()
		PrintSudoku(s)
	}
}
