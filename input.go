package main

import (
	"os"
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
		errMsg = "No arguments on program call"
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
	count := len(s)
	if count > 81 {
		errMsg = "Argument length is too long"
		return false
	}
	if count < 81 {
		errMsg = "Argument length is too short"
		return false
	}
	for i := 0; i < count; i++ {
		if s[i] < '0' || s[i] > '9' {
			errMsg = "Digit \"" + string(s[i]) + "\" not recognized"
			return false
		}
	}
	return true // All are valid digits
}
