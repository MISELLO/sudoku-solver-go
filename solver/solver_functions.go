package solver

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
func Solve(board *tBoard) tStats {
	var changesDone bool = true
	var iter int = 0
	for changesDone {
		changesDone = false
		iter++

		// First strategy: Remove the imposible.
		// If one number candidate is left, this is the number.
		for i, cell := range *board {
			if cell.num != 0 {
				mark(board, cell.num, getRowPos(i))
				mark(board, cell.num, getColPos(i))
				mark(board, cell.num, getBlkPos(i))
			}
		}
		changesDone = setUnique(board)

		// Second strategy: Every row, column and block must have each number.
		// If a number can only be placed in one spot, then this number must go to that spot.
		if !changesDone && !isSolved(*board) {
			for i := 0; i < 9; i++ {
				setUniqueFromList(board, getRowNum(i), &changesDone)
				setUniqueFromList(board, getColNum(i), &changesDone)
				setUniqueFromList(board, getBlkNum(i), &changesDone)
			}
		}
	}
	var stats tStats
	stats.solved = isSolved(*board)
	stats.iterations = iter
	return stats
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
	}
}

// getRowPos returns the whole row (n) belongs
// in the form of a list of tBoard indexes
func getRowPos(n int) []int {
	var list []int
	start, end := (n/9)*9, ((n/9)*9)+9
	for i := start; i < end; i++ {
		list = append(list, i)
	}
	return list
}

// getRowNum returns the row number (n)
// being (n) a number from 0 to 8
func getRowNum(n int) []int {
	var list []int
	start, end := n*9, (n*9)+9
	for i := start; i < end; i++ {
		list = append(list, i)
	}
	return list
}

// getColPos returns the whole column (n) belongs
// in the form of a list of tBoard indexes
func getColPos(n int) []int {
	var list []int
	start, end := n%9, 9*9
	for i := start; i < end; i += 9 {
		list = append(list, i)
	}
	return list
}

// getColNum returns the column number (n)
// being (n) a number from 0 to 8
func getColNum(n int) []int {
	var list []int
	for i := 0; i < 9; i++ {
		list = append(list, n+(9*i))
	}
	return list
}

// getBlkPos returns the whole block (n) belongs
// in the form of a list of tBoard indexes
func getBlkPos(n int) []int {
	var list []int
	r1 := 9*(n/9-(n/9)%3) + (n%9 - n%3)
	r2 := 9*(n/9-(n/9)%3) + (n%9 - n%3) + 9
	r3 := 9*(n/9-(n/9)%3) + (n%9 - n%3) + 18
	list = append(list, r1, r1+1, r1+2)
	list = append(list, r2, r2+1, r2+2)
	list = append(list, r3, r3+1, r3+2)
	return list
}

// getBlkNum returns the block number (n)
// being (n) a number from 0 to 8
func getBlkNum(n int) []int {
	var list []int
	r1 := ((n % 3) * 3) + ((n / 3) * 27)
	r2 := ((n % 3) * 3) + ((n / 3) * 27) + 9
	r3 := ((n % 3) * 3) + ((n / 3) * 27) + 18
	list = append(list, r1, r1+1, r1+2)
	list = append(list, r2, r2+1, r2+2)
	list = append(list, r3, r3+1, r3+2)
	return list
}

// PrintAvailable prints for each position the available numbers (test function)
/*
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
*/

// GivenList returns an array of 81 bools that tells if a position
// was given (true) or calculated (false)
func GivenList(board tBoard) [81]bool {
	var g [81]bool
	for i, v := range board {
		g[i] = v.avl[0]
	}
	return g
}

// isSolved returns true if the sudoku has been solved
func isSolved(board tBoard) bool {
	for _, v := range board {
		if v.num == 0 {
			return false
		}
	}
	return true
}

// setUniqueFromList checks that the given list (l) of indexes reference
// all numbers from 1 to 9. If one number can only be in one place then
// it is placed there and the check variable (c) is set to true.
func setUniqueFromList(b *tBoard, l []int, c *bool) {
	for _, i := range l {
		if (*b)[i].num == 0 {
			for n := 1; n <= 9; n++ {
				if (*b)[i].avl[n] && !hasA(*b, uint8(n), l) && countAvailable(*b, uint8(n), l) == 1 {
					set(b, i, uint8(n))
					*c = true
				}
			}
		}
	}
}

// hasA returns true if the list of indexes (l) already reference
// the number (x).
func hasA(b tBoard, x uint8, l []int) bool {
	for _, i := range l {
		if b[i].num == x {
			return true
		}
	}
	return false
}

// countAvailable returns the number of indexes from the list (l) where
// the board (b) accepts the given number (x).
func countAvailable(b tBoard, x uint8, l []int) int {
	var count int
	for _, i := range l {
		if b[i].avl[x] {
			count++
		}
	}
	return count
}

// set sets the number (x) on the position (i) from the board (b)
func set(b *tBoard, i int, x uint8) {
	(*b)[i].num = x
	for j := 1; j <= 9; j++ {
		(*b)[i].avl[j] = false
	}
	(*b)[i].avl[x] = true
}

// IsSolved tStats method that returns true if the sudoku has been solved
func (s *tStats) IsSolved() bool {
	return s.solved
}

// NumIterations tStats method that returns the number of iterations done
// in order to solve the sudoku
func (s *tStats) NumIterations() int {
	return s.iterations
}
