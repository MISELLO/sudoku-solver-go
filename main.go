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

// var allSolutions bool = false
var maxSol uint
var bruteFTime uint

// CMaxSol is the default value of maxSol
const CMaxSol uint = 10

// CBruteFTime is the default value of bruteFTime
const CBruteFTime uint = 60

func main() {

	// Flag management
	ncFlag := flag.Bool("nc", false, "No color, if set disables the colors of the results.")
	sFlag := flag.Bool("s", false, "String format output, displays only the result as a string of 81 digits.")
	bFlag := flag.Bool("b", false, "Brute-force, it uses brute-force when necessary.")
	flag.UintVar(&maxSol, "ms", CMaxSol, "Max solutions, defines the maximum number of solutions the algorithm will "+
		"calculate. If this number is reached a \"(+)\" will appear next to the Solutions amount output."+
		"(automatically activates brute-force [-b])")
	flag.UintVar(&bruteFTime, "bt", CBruteFTime, "Brute-force time, time in seconds the brute-force algorithm is "+
		"allowed to run. If this time is reached a \"(+)\" will appear next to the Solutions amount output. "+
		"(automatically activates brute-force [-b])")

	flag.Parse()

	if *ncFlag {
		colorsOn = false
	}

	if *sFlag {
		strFmt = true
	}

	if *bFlag || maxSol != CMaxSol || bruteFTime != CBruteFTime {
		bruteForce = true
	}

	s, e := GetInput()

	if e {
		PrintErrMsg()
		PrintUsage()
	} else {
		board := solver.Load(s)
		done := make(chan bool)
		go Processing(done)
		stats := solver.Solve(&board, bruteForce, maxSol, bruteFTime)
		done <- true
		s = solver.Unload(board)
		PrintSudoku(s, solver.GivenList(board), solver.Wrong(board))
		PrintSolved(stats.IsSolved())
		PrintUnknown(solver.CountUnknown(board))
		PrintNumSolutions(stats.NumSolutions(), stats.Interrupted())
		PrintIterations(stats.NumIterations())
		PrintDeduced(stats.Deduced())
		PrintBruteForce(stats.BruteForce())
	}
}
