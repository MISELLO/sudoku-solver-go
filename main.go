package main

import (
	"github.com/MISELLO/sudoku-solver-go/solver"
)

var errMsg string

var colorsOn bool = true

func main() {
	s, e := GetInput()

	if e {
		PrintErrMsg()
		PrintUsage()
	} else {
		board := solver.Load(s)
		stats := solver.Solve(&board)
		s = solver.Unload(board)
		PrintSudoku(s, solver.GivenList(board))
		PrintSolved(stats.IsSolved())
	}
}
