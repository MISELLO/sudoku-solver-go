package main

import (
	"fmt"
	"time"
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
// w --> an array where each position tells if that position
// is wrong (repeated number) or not.
func PrintSudoku(s string, g [81]bool, w [81]bool) {
	if strFmt {
		printSudokuStringFormat(s, g, w)
	} else {
		printSudokuRegular(s, g, w)
	}
}

func printSudokuStringFormat(s string, g, w [81]bool) {
	for i := 0; i < len(s); i++ {
		if colorsOn && w[i] {
			fmt.Print("\033[41m\033[90m", s[i:i+1], "\033[0m")
		} else if colorsOn && g[i] {
			fmt.Print("\033[90m", s[i:i+1], "\033[0m")
		} else {
			fmt.Print(s[i : i+1])
		}
	}
	fmt.Println()
}

func printSudokuRegular(s string, g, w [81]bool) {
	// Make sure to clear the processing message
	fmt.Println("\033[1K                    ")
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
			printDigit(s[i:i+1], g[i], w[i])
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Println()
	fmt.Println()
}

func printDigit(d string, gi, wi bool) {
	if colorsOn && wi && gi {
		fmt.Print(" \033[41m\033[90m", d, "\033[0m")
	} else if colorsOn && wi {
		fmt.Print(" \033[41m", d, "\033[0m")
	} else if colorsOn && gi {
		fmt.Print(" \033[90m", d, "\033[0m")
	} else {
		fmt.Print(" ", d)
	}
}

// PrintBoard prints an empty board and a message instructing the user to fill it
func PrintBoard() {

	fmt.Println()
	fmt.Println(" Please, fill all known cells and press ENTER when done.")
	fmt.Println()

	for i := 0; i < 81; i++ {
		if i != 0 && i%9 == 0 {
			fmt.Println()
		} else if i%9 != 0 && i%3 == 0 {
			fmt.Print(" |")
		}
		if i != 0 && i%27 == 0 {
			fmt.Println("-------+-------+-------")
		}
		fmt.Print("  ")
	}

	fmt.Println()
	fmt.Println()
}

// RemoveBoard removes all that was drawn by PrintBoard() including the message to the user
func RemoveBoard() {
	for i := 0; i < 15; i++ {
		fmt.Print("\033[1F\033[2K")
	}
}

// PrintErrMsg prints the error message declared on main.go and updated on input.go
func PrintErrMsg() {
	fmt.Println(errMsg)
}

// PrintSolved prints if the sudoku could be solved or not.
func PrintSolved(s bool) {

	// String format enabled: We don't print anything else
	if strFmt {
		return
	}

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
	fmt.Printf("     ")

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
func PrintNumSolutions(n int, interrupt bool) {

	// String format enabled: We don't print anything else
	if strFmt {
		return
	}

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
	fmt.Printf("  ")

	// Print the number of solutions
	if interrupt {
		fmt.Printf("%d(+)\n", n)
	} else {
		fmt.Printf("%d\n", n)
	}

	// Restore cursor position
	fmt.Printf("\033[u")
}

// PrintIterations prints the number of iterations done
// in order to solve the sudoku
func PrintIterations(n int) {

	// String format enabled: We don't print anything else
	if strFmt {
		return
	}

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
	fmt.Printf(" ")

	// Print the number of iterations
	fmt.Printf("%d\n", n)

	// Restore cursor position
	fmt.Printf("\033[u")
}

// PrintDeduced prints the number of deductions done
// in order to solve the sudoku
func PrintDeduced(n int) {

	// String format enabled: We don't print anything else
	if strFmt {
		return
	}

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
	fmt.Printf("    ")

	// Print the number of deductions
	fmt.Printf("%d\n", n)

	// Restore cursor position
	fmt.Printf("\033[u")
}

// PrintBruteForce prints if we have used brute-force or not to solve the sudoku
func PrintBruteForce(b bool) {

	// String format enabled: We don't print anything else
	if strFmt {
		return
	}

	// Save cursor position
	fmt.Printf("\033[s")

	// Up (7 times)
	fmt.Printf("\033[7A")

	// Right (25 times)
	fmt.Printf("\033[25C")

	// Print "Brute-force:" in bold
	if colorsOn {
		fmt.Printf("\033[1m Brute-force: \033[0m")
	} else {
		fmt.Printf(" Brute-force: ")
	}
	fmt.Printf("")

	// Print if we used brute-force or not
	// Print YES in green background or NO in red background
	if colorsOn {
		if b {
			fmt.Printf("\033[42m YES \033[0m\n")
		} else {
			fmt.Printf("\033[41m NO \033[0m\n")
		}
	} else {
		if b {
			fmt.Printf(" YES \n")
		} else {
			fmt.Printf(" NO \n")
		}
	}

	// Restore cursor position
	fmt.Printf("\033[u")
}

// Processing prints something that moves on the screen so the user can see it is still working
func Processing(c <-chan bool) {
	text := "Processing ... |"
	fmt.Printf("%s", text)
	ani := []rune("/—\\|")
	i := 0
	t := time.Now()
	for {
		select {
		case <-c:
			fmt.Printf("\033[1K")
			return
		default:
			if time.Since(t) >= 500*time.Millisecond {
				t = time.Now()
				fmt.Printf("\b")
				fmt.Printf("%c", ani[i])
				i = (i + 1) % len(ani)
			}
		}
	}
}
