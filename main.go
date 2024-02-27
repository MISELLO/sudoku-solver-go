package main

import (
	"github.com/MISELLO/sudoku-solver-go/solver"
)

func main() {
	solver.Print()

	s, e := GetInput()

	if e {
		PrintUsage()
	} else {
		PrintSudoku(s)
	}
}
