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
		board := solver.Load(s)
		solver.Solve(&board)
		s = solver.Unload(board)
		PrintSudoku(s)
		//solver.PrintAvailable(board)
	}
}
