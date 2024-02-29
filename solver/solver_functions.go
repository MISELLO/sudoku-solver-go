package solver

import "fmt"

func Print() {
	fmt.Println(w)
}

// Load takes a valid string and it converts it to our sudoku structure
func Load(s string) tBoard {
	var board tBoard
	for i := 0; i < len(s); i++ {
		board[i].num = uint8(s[i] - '0')
		makeAvailable(&board[i].avl, s[i])
	}
	return board
}

func Unload(board *tBoard) string {
	var s string
	for i := 0; i < len(board); i++ {
		s += string(board[i].num + '0')
	}
	return s
}

// makeAvailable sets the array of booleans to the proper values depending on what is placed
// on the board cell.
func makeAvailable(a *[10]bool, b byte) {
	switch b {
	case '0':
		*a = [10]bool{false, true, true, true, true, true, true, true, true, true}
	case '1':
		*a = [10]bool{true, true, false, false, false, false, false, false, false, false}
	case '2':
		*a = [10]bool{true, false, true, false, false, false, false, false, false, false}
	case '3':
		*a = [10]bool{true, false, false, true, false, false, false, false, false, false}
	case '4':
		*a = [10]bool{true, false, false, false, true, false, false, false, false, false}
	case '5':
		*a = [10]bool{true, false, false, false, false, true, false, false, false, false}
	case '6':
		*a = [10]bool{true, false, false, false, false, false, true, false, false, false}
	case '7':
		*a = [10]bool{true, false, false, false, false, false, false, true, false, false}
	case '8':
		*a = [10]bool{true, false, false, false, false, false, false, false, true, false}
	case '9':
		*a = [10]bool{true, false, false, false, false, false, false, false, false, true}
	}
}

// Solve tryes to solve the sudoku puzzle
func Solve(board *tBoard) {
	var changesDone bool = true
	for changesDone {
		changesDone = false
		c := markRows(board)
		changesDone = changesDone || c
		setUnique(board)
	}
}

// setUnique checks all cells where the number is not yet known and updates it if
// there is only one number available left.
func setUnique(board *tBoard) {
	for i := 0; i < len(*board); i++ {
		if (*board)[i].num == 0 {
			(*board)[i].num = whatNumber(board[i].avl)
		}
	}
}

// whatNumber checks the array of booleans from indexes 1 to 9. If only one of
// this positions is true returns this number, returns 0 otherwise.
func whatNumber(a [10]bool) uint8 {
	var n uint8 = 0
	for i := 1; i < len(a); i++ {
		if n != 0 && a[i] {
			return 0
		}
		if n == 0 && a[i] {
			n = uint8(i)
		}
	}
	return n
}

// markRows updates the numbers available for every row.
func markRows(board *tBoard) bool {
	return false
}
