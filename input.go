package main

/*
#include <stdio.h>
#include <unistd.h>
#include <termios.h>
char getch(){
	char ch = 0;
	struct termios old = {0};
	fflush(stdout);
	if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
	old.c_lflag &= ~ICANON;
	old.c_lflag &= ~ECHO;
	old.c_cc[VMIN] = 1;
	old.c_cc[VTIME] = 0;
	if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
	if( read(0, &ch,1) < 0 ) perror("read()");
	old.c_lflag |= ICANON;
	old.c_lflag |= ECHO;
	if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
	return ch;
}
*/
import "C"

// code above was found on https://stackoverflow.com/questions/14094190/function-similar-to-getchar

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// GetInput checks several forms of input in order to get the input from the user
// This forms are:
//   - File: The route of a file that contains 1 sudoku in different formats.
//   - String: 1 single parameter (not counting flags) when calling the program that represents
//     the 81 digits for the sudoku, starting from top left and going right, replacing the empty
//     spaces with 0's.
//   - Visual: If no parameters are provided a board will be displayed and the user will have to
//     fill it.
func GetInput() (string, bool) {
	var result string
	var error bool

	if weHaveArgs() {
		// We first try to open the first argument
		f, err := os.Open(flag.Arg(0))
		if err == nil {
			// We process the valid input file
			defer f.Close()
			result, error = processFile(f)
		} else if firstArgIsValid() {
			// Input is not a file, but a valid sudoku formatted as a string
			result = flag.Arg(0)
			error = false
		} else {
			// Not a file or valid input
			result = ""
			error = true
		}
	} else { // We do NOT have arguments
		result = visualInput()
		error = false
	}
	return result, error
}

// weHaveArgs returns true if the the program has been called with at least 1 argument
func weHaveArgs() bool {
	return len(flag.Args()) >= 1
}

// firstArgIsValid returns true if the first argument is a valid input
func firstArgIsValid() bool {
	s := flag.Arg(0)
	count := len(s)
	if count > 81 {
		errMsg = "Argument not a file and length is too long"
		return false
	}
	if count < 81 {
		errMsg = "Argument not a file and length is too short"
		return false
	}
	for i := 0; i < count; i++ {
		if s[i] < '0' || s[i] > '9' {
			errMsg = "Argument not a file and digit \"" + string(s[i]) + "\" not recognized"
			return false
		}
	}
	return true // All are valid digits
}

// processFile reads the input file and returns the result and if there has been an error.
func processFile(f *os.File) (string, bool) {
	var result string
	var error bool

	scn := bufio.NewScanner(f)
	scn.Scan()
	s := scn.Text()
	if len(s) == 81 {
		return processFileType1(s)
	}
	if len(s) == 9 {
		return processFileType2(f, scn, s)
	}
	if len(s) == 11 {
		return processFileType3(f, scn, s)
	}
	if len(s) == 22 || len(s) == 23 {
		return processFileType4(f, scn, s)
	}
	errMsg = "Incorrect file content (not supported length)"
	result = ""
	error = true
	return result, error
}

// processFileType1 process the input file when it's type 1 (see samples folder)
func processFileType1(s string) (string, bool) {
	var result string = ""
	var error bool = false

	for i := 0; i < len(s) && !error; i++ {
		if s[i] == ' ' {
			result += "0"
		} else if s[i] >= '0' && s[i] <= '9' {
			result += string(s[i])
		} else {
			errMsg = "Incorrect file content (unrecognized character \"" + string(s[i]) + "\")"
			result = ""
			error = true
		}
	}
	return result, error
}

// processFileType2 process the input file when it's type 2 (see samples folder)
func processFileType2(f *os.File, scn *bufio.Scanner, s string) (string, bool) {
	var result string = ""
	var error bool = false

	for n := 0; n < 9 && !error; n++ {
		if len(s) != 9 {
			errMsg = "Incorrect file content (not supported length)"
			result = ""
			error = true
		}
		for i := 0; i < len(s) && !error; i++ {
			if s[i] == ' ' {
				result += "0"
			} else if s[i] >= '0' && s[i] <= '9' {
				result += string(s[i])
			} else {
				errMsg = "Incorrect file content (unrecognized character \"" + string(s[i]) + "\")"
				result = ""
				error = true
			}
		}
		scn.Scan()
		s = scn.Text()
	}
	return result, error
}

// processFileType3 process the input file when it's type 3 (see samples folder)
func processFileType3(f *os.File, scn *bufio.Scanner, s string) (string, bool) {
	var result string = ""
	var error bool = false

	for n := 0; n < 11 && !error; n++ {
		if n == 3 || n == 7 {
			// We skip the horizontal lines
			scn.Scan()
			s = scn.Text()
			continue
		}
		if len(s) != 11 {
			errMsg = "Incorrect file content (not supported length)"
			result = ""
			error = true
		}
		for i := 0; i < len(s) && !error; i++ {
			if i == 3 || i == 7 {
				// We skip the vertical lines
				continue
			}
			if s[i] == ' ' {
				result += "0"
			} else if s[i] >= '0' && s[i] <= '9' {
				result += string(s[i])
			} else {
				errMsg = "Incorrect file content (unrecognized character \"" + string(s[i]) + "\")"
				result = ""
				error = true
			}
		}
		scn.Scan()
		s = scn.Text()
	}
	return result, error
}

// processFileType4 process the input file when it's type 4 (see samples folder)
func processFileType4(f *os.File, scn *bufio.Scanner, s string) (string, bool) {
	var result string = ""
	var error bool = false

	for n := 0; n < 11 && !error; n++ {
		if n == 3 || n == 7 {
			// We skip the horizontal lines
			scn.Scan()
			s = scn.Text()
			continue
		}
		if len(s) < 22 || len(s) > 23 {
			errMsg = "Incorrect file content (not supported length)"
			result = ""
			error = true
		}
		for i := 1; i < len(s) && !error; i += 2 {
			if i == 7 || i == 15 {
				// We skip the vertical lines
				continue
			}
			if s[i] == ' ' {
				result += "0"
			} else if s[i] >= '0' && s[i] <= '9' {
				result += string(s[i])
			} else {
				errMsg = "Incorrect file content (unrecognized character \"" + string(s[i]) + "\")"
				result = ""
				error = true
			}
		}
		scn.Scan()
		s = scn.Text()
	}
	return result, error
}

// visualInput makes the user introduce the input on a drawn board on the screen.
func visualInput() string {

	input := "000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	PrintBoard()

	var posX, posY int

	// Save cursor position
	fmt.Printf("\033[s")

	// Go to first position
	fmt.Print("\033[12A\033[1C")

	k := C.getch()
	for k != 10 {
		// Arrow key
		if k == '\033' && C.getch() == '[' { // (27 & 91)
			movement(&posX, &posY)
		}

		// Number key (or tab or space or backspace)
		if (k >= '0' && k <= '9') || k == '\t' || k == ' ' || k == 127 {
			number(byte(k), &posX, &posY, &input)
		}

		k = C.getch()
	}

	// Restore cursor position
	fmt.Printf("\033[u")

	RemoveBoard()
	return input
}

// movement manages the arrow movement inside the visualInput function
func movement(x, y *int) {
	k := C.getch()
	switch k {
	case 'A': // Up (65)
		if *y > 0 {
			if *y%3 == 0 {
				fmt.Print("\033[A")
			}
			fmt.Print("\033[A")
			*y--
		}
	case 'B': // Down (66)
		if *y < 8 {
			*y++
			if *y%3 == 0 {
				fmt.Print("\033[B")
			}
			fmt.Print("\033[B")
		}
	case 'C': // Right (67)
		if *x < 8 {
			*x++
			if *x%3 == 0 {
				fmt.Print("\033[2C")
			}
			fmt.Print("\033[2C")
		}
	case 'D': // Left (68)
		if *x > 0 {
			if *x%3 == 0 {
				fmt.Print("\033[2D")
			}
			fmt.Print("\033[2D")
			*x--
		}
	}
}

// number manages the number input inside the visualInput function
func number(k byte, x, y *int, s *string) {
	numberBackSpManagement(k, x, y, s)

	if k == 127 { // Backspace already managed
		return
	}

	if k >= '1' && k <= '9' {
		fmt.Print(string(k))
		i := (*y)*9 + (*x)
		(*s) = (*s)[:i] + string(k) + (*s)[i+1:]
		fmt.Print("\b")
	} else if k == ' ' || k == '0' {
		i := (*y)*9 + (*x)
		(*s) = (*s)[:i] + "0" + (*s)[i+1:]
		fmt.Print(" \b")
	}
	if *x < 8 {
		*x++
		if *x%3 == 0 {
			fmt.Print("\033[2C")
		}
		fmt.Print("\033[2C")
	} else if *x == 8 && *y < 8 {
		*x = 0
		fmt.Print("\033[20D")
		*y++
		if *y%3 == 0 {
			fmt.Print("\033[B")
		}
		fmt.Print("\033[B")
	}
}

// numberBackSpManagement or "number backspace management" is a function that
// handles the use of backspace inside the number function that is used on
// the visualInput function.
func numberBackSpManagement(k byte, x, y *int, s *string) {
	if k != 127 { // We only work if it is a backspace
		return
	}

	if *x == 0 && *y == 0 {
		// We are at the start, do nothing
		return
	} else if *x == 8 && *y == 8 && (*s)[80] != '0' { // Last position filled
		(*s) = (*s)[:80] + "0"
		fmt.Print(" \b")
		return
	} else if *x == 0 {
		if *y%3 == 0 {
			fmt.Print("\033[A")
		}
		fmt.Print("\033[A")
		*y--
		*x = 8
		fmt.Print("\033[20C")
	} else {
		if *x%3 == 0 {
			fmt.Print("\033[2D")
		}
		fmt.Print("\033[2D")
		*x--
	}
	i := (*y)*9 + (*x)
	(*s) = (*s)[:i] + "0" + (*s)[i+1:]
	fmt.Print(" \b")
	return
}
