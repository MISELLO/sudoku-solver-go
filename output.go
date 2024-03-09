package main

import (
	"fmt"
)

// PrintUsage prints how to use this command.
func PrintUsage() {
	fmt.Println(" USAGE:")
	fmt.Println("  ./sudoku-solver.exe [sudoku digits]")
	fmt.Println()
}

// PrintSudoku prints a representation of a sudoku on the screen
// s --> a string of 81 digits that represent the sudoku board
// g --> an array where each position tells if that position
// was given (true) or calculated (false)
func PrintSudoku(s string, g [81]bool) {
	fmt.Println()
	for i := 0; i < len(s); i++ {
		if i != 0 && i%9 == 0 {
			fmt.Println()
		} else if i%9 != 0 && i%3 == 0 {
			fmt.Print(" |")
		}
		if i != 0 && i%27 == 0 {
			fmt.Println("-------+-------+-------")
		}

		if s[i:i+1] != "0" {
			if colorsOn && g[i] {
				fmt.Print(" \033[90m", s[i:i+1], "\033[0m")
			} else {
				fmt.Print(" ", s[i:i+1])
			}
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Println()
}

// PrintErrMsg prints the error message declared on main.go and updated on input.go
func PrintErrMsg() {
	fmt.Println(errMsg)
}

// PrintSolved prints if the sudoku could be solved or not.
func PrintSolved(s bool) {
	fmt.Printf("\033[1m  Solved: \033[0m")
	if s {
		fmt.Printf("\033[42m YES \033[0m\n")
	} else {
		fmt.Printf("\033[41m NO \033[0m\n")
	}
}
