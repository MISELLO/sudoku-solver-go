package main

import (
	"github.com/MISELLO/sudoku-solver-go/solver"

	"flag"
)

var errMsg string

var noFlagArgs []string

var colorsOn bool = true

func main() {

	// Flag management
	ncFlag := flag.Bool("nc", false, "No color, if set disables the colors of the results.")

	flag.Parse()

	if *ncFlag {
		colorsOn = false
	}

	s, e := GetInput()

	if e {
		PrintErrMsg()
		PrintUsage()
	} else {
		board := solver.Load(s)
		stats := solver.Solve(&board)
		s = solver.Unload(board)
		PrintSudoku(s, solver.GivenList(board), solver.Wrong(board))
		PrintSolved(stats.IsSolved())
		PrintIterations(stats.NumIterations())
		PrintNumSolutions(stats.NumSolutions())
		PrintDeduced(stats.Deduced())
	}
}
