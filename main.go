package main

import (
	"github.com/MISELLO/sudoku-solver-go/solver"

	"flag"
)

var errMsg string

var noFlagArgs []string

// Flags
var colorsOn bool = true
var strFmt bool = false
var bruteForce bool = false

func main() {

	// Flag management
	ncFlag := flag.Bool("nc", false, "No color, if set disables the colors of the results.")
	sFlag := flag.Bool("s", false, "String format output, displays only the result as a string of 81 digits.")
	bFlag := flag.Bool("b", false, "Brute-force, it uses brute-force when necessary [Under development]")

	flag.Parse()

	if *ncFlag {
		colorsOn = false
	}

	if *sFlag {
		strFmt = true
	}

	if *bFlag {
		bruteForce = true
	}

	s, e := GetInput()

	if e {
		PrintErrMsg()
		PrintUsage()
	} else {
		board := solver.Load(s)
		stats := solver.Solve(&board, bruteForce)
		s = solver.Unload(board)
		PrintSudoku(s, solver.GivenList(board), solver.Wrong(board))
		PrintSolved(stats.IsSolved())
		PrintIterations(stats.NumIterations())
		PrintNumSolutions(stats.NumSolutions())
		PrintDeduced(stats.Deduced())
		PrintBruteForce(stats.BruteForce())
	}
}
