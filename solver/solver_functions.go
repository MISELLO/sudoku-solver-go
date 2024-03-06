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

func Unload(board tBoard) string {
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
		for i, cell := range *board {
			if cell.num != 0 {
				//fmt.Println(i)
				mark(board, cell.num, getRow(i))
				mark(board, cell.num, getCol(i))
				mark(board, cell.num, getBlk(i))
			}
		}
		changesDone = setUnique(board)
	}
}

// setUnique checks all cells where the number is not yet known and updates it if
// there is only one number available left.
func setUnique(board *tBoard) bool {
	var changesDone bool = false
	for i := 0; i < len(*board); i++ {
		if (*board)[i].num == 0 {
			(*board)[i].num = whatNumber((*board)[i].avl)
			if (*board)[i].num != 0 {
				changesDone = true
			}
		}
	}
	return changesDone
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

// mark modifies the board (b).
// The list of indexes to modify is on the list (l)
// By marking we set the number (n) as not available.
func mark(b *tBoard, n uint8, l []int) {
	for _, v := range l {
		(*b)[v].avl[n] = false
		//fmt.Println(" Marking number", n, "at position", v, "as false.")
	}
}

// getRow returns the whole row (n) belongs
// in the form of a list of tBoard indexes
func getRow(n int) []int {
	var list []int
	start, end := (n/9)*9, ((n/9)*9)+9
	for i := start; i < end; i++ {
		list = append(list, i)
	}
	return list
}

// getCol returns the whole column (n) belongs
// in the form of a list of tBoard indexes
func getCol(n int) []int {
	var list []int
	start, end := n%9, 9*9
	for i := start; i < end; i += 9 {
		list = append(list, i)
	}
	return list
}

// getBlk returns the whole block (n) belongs
// in the form of a list of tBoard indexes
func getBlk(n int) []int {
	var list []int
	r1 := 9*(n/9-(n/9)%3) + (n%9 - n%3)
	r2 := 9*(n/9-(n/9)%3) + (n%9 - n%3) + 9
	r3 := 9*(n/9-(n/9)%3) + (n%9 - n%3) + 18
	list = append(list, r1, r1+1, r1+2)
	list = append(list, r2, r2+1, r2+2)
	list = append(list, r3, r3+1, r3+2)
	return list
}

// PrintAvailable prints for each position the available numbers (test function)
func PrintAvailable(board tBoard) {
	fmt.Println()
	for i, c := range board {
		if c.num != 0 {
			fmt.Printf("Position %2d contains %d\n", i, c.num)
		} else {
			fmt.Printf("Position %2d has available:", i)
			for j, v := range c.avl {
				if v {
					fmt.Printf(" %d", j)
				}
			}
			fmt.Println()
		}
	}
}
