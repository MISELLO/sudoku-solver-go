package main

import (
	"os"
	"unicode"
	"unicode/utf8"
)

// GetInput checks several forms of input in order to get the input from the user
// This forms are:
//   - Parameters: 1 single parameter when calling the program that represents the 81 digits for
//     the sudoku, starting from top left and going right, replacing the empty spaces
//     with 0's.
func GetInput() (string, bool) {
	var result string
	var error bool
	if weHaveArgs() {
		if firstArgIsValid() {
			result = os.Args[1]
			error = false
		} else {
			result = ""
			error = true
		}
	} else {
		result = ""
		error = true
	}
	return result, error
}

// weHaveArgs returns true if the the program has been called with at least 1 argument
func weHaveArgs() bool {
	return len(os.Args) >= 2
}

// firstArgIsValid returns true if the first argument is a valid input
func firstArgIsValid() bool {
	s := os.Args[1]
	if len(s) == 81 {
		for i, w := 0, 0; i < 81; i += w {
			rune, width := utf8.DecodeRuneInString(s[i:])
			if !unicode.IsDigit(rune) {
				return false
			}
			w = width
		}
		return true // All are valid digits
	}
	return false
}
