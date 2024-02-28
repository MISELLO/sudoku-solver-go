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

// PrintSudoku prints a representation o a sudoku on the screen
func PrintSudoku(s string) {
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
			fmt.Print(" ", s[i:i+1])
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
