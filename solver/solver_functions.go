package solver

import "fmt"

func Print() {
	fmt.Println(w)
}

// Load takes a valid string and it converts it to our sudoku structure
func Load(s string) {
	for i := 0; i < len(s); i++ {
		board[i].num = uint8(s[i] - '0')
		makeAvailable(&board[i].avl, s[i])
	}
}

func Unload() string {
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
