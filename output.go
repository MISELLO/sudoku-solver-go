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
	fmt.Println()
}

// PrintErrMsg prints the error message declared on main.go and updated on input.go
func PrintErrMsg() {
	fmt.Println(errMsg)
}

// PrintSolved prints if the sudoku could be solved or not.
func PrintSolved(s bool) {
	// Save cursor position
	fmt.Printf("\033[s")

	// Up (11 times)
	fmt.Printf("\033[11A")

	// Right (25 times)
	fmt.Printf("\033[25C")

	// Print "Solved:" in bold
	if colorsOn {
		fmt.Printf("\033[1m Solved: \033[0m")
	} else {
		fmt.Printf(" Solved: ")
	}
	fmt.Printf("    ")

	// Print YES in green background or NO in red background
	if colorsOn {
		if s {
			fmt.Printf("\033[42m YES \033[0m\n")
		} else {
			fmt.Printf("\033[41m NO \033[0m\n")
		}
	} else {
		if s {
			fmt.Printf(" YES \n")
		} else {
			fmt.Printf(" NO \n")
		}
	}

	// Restore cursor position
	fmt.Printf("\033[u")
}

// PrintNumSolutions prints the number of solutions found
func PrintNumSolutions(n int) {
	// Save cursor position
	fmt.Printf("\033[s")

	// Up (10 times)
	fmt.Printf("\033[10A")

	// Right (25 times)
	fmt.Printf("\033[25C")

	// Print "Solutions:" in bold
	if colorsOn {
		fmt.Printf("\033[1m Solutions: \033[0m")
	} else {
		fmt.Printf(" Solutions: ")
	}
	fmt.Printf(" ")

	// Print the number of iterations
	fmt.Printf("%d\n", n)

	// Restore cursor position
	fmt.Printf("\033[u")
}

// PrintIterations prints the number of iterations done
// in order to solve the sudoku
func PrintIterations(n int) {
	// Save cursor position
	fmt.Printf("\033[s")

	// Up (9 times)
	fmt.Printf("\033[9A")

	// Right (25 times)
	fmt.Printf("\033[25C")

	// Print "Iterations:" in bold
	if colorsOn {
		fmt.Printf("\033[1m Iterations: \033[0m")
	} else {
		fmt.Printf(" Iterations: ")
	}

	// Print the number of iterations
	fmt.Printf("%d\n", n)

	// Restore cursor position
	fmt.Printf("\033[u")
}

// PrintDeduced prints the number of deductions done
// in order to solve the sudoku
func PrintDeduced(n int) {
	// Save cursor position
	fmt.Printf("\033[s")

	// Up (8 times)
	fmt.Printf("\033[8A")

	// Right (25 times)
	fmt.Printf("\033[25C")

	// Print "Deduced:" in bold
	if colorsOn {
		fmt.Printf("\033[1m Deduced: \033[0m")
	} else {
		fmt.Printf(" Deduced: ")
	}
	fmt.Printf("   ")

	// Print the number of deductions
	fmt.Printf("%d\n", n)

	// Restore cursor position
	fmt.Printf("\033[u")
}
