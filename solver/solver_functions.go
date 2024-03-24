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

// Solve tries to solve the sudoku puzzle
// board --> a previously loaded board game
// bf    --> brute-force allowed
// as    --> all solutions
func Solve(board *tBoard, bf bool, as bool) tStats {
	var stats tStats
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
			missingCount := countMissing(*board)
			for i := 0; i < 9; i++ {
				setUniqueFromList(board, getRowNum(i), &changesDone)
				setUniqueFromList(board, getColNum(i), &changesDone)
				setUniqueFromList(board, getBlkNum(i), &changesDone)
			}
			missingCount -= countMissing(*board)
			stats.deduced += missingCount
		}
	}

	if bf && !isSolved(*board) && !anyWrong(*board) {
		// We apply brute-force by using backtracking
		stats.bruteForce = true
		ck := make(map[string]bool)
		solveBckTck(*board, &stats, &ck, as)
	}

	stats.solved = isSolved(*board) && !anyWrong(*board)
	if stats.solved {
		stats.solutions = append(stats.solutions, Unload(*board))
	} else if len(stats.solutions) > 0 {
		// No solutions at first, but we got solutions by using brute-force
		stats.solved = true
		for i, v := range stats.solutions[0] {
			// We load the first solution because the original board has missing numbers
			if (*board)[i].num == 0 {
				(*board)[i].num = uint8(byte(v) - '0')
			}
		}
	}

	stats.iterations = iter
	return stats
}

// solveBckTck tries to solve the sudoku puzzle using backtracking (plus the strategies
// defined at "Solve") and stores the different results into (sol)
func solveBckTck(board tBoard, stats *tStats, ck *map[string]bool, allSol bool) {

	//fmt.Println(Unload(board))

	var changesDone bool = true
	for changesDone {
		changesDone = false

		// First strategy: Remove the imposible.
		// If one number candidate is left, this is the number.
		for i, cell := range board {
			if cell.num != 0 {
				mark(&board, cell.num, getRowPos(i))
				mark(&board, cell.num, getColPos(i))
				mark(&board, cell.num, getBlkPos(i))
			}
		}
		changesDone = setUnique(&board)

		// Second strategy: Every row, column and block must have each number.
		// If a number can only be placed in one spot, then this number must go to that spot.
		if !changesDone && !isSolved(board) {
			for i := 0; i < 9; i++ {
				setUniqueFromList(&board, getRowNum(i), &changesDone)
				setUniqueFromList(&board, getColNum(i), &changesDone)
				setUniqueFromList(&board, getBlkNum(i), &changesDone)
			}
		}
	}

	if (*ck)[Unload(board)] {
		return
	}
	(*ck)[Unload(board)] = true

	if isSolved(board) {
		(*stats).solutions = append((*stats).solutions, Unload(board))
		//fmt.Println("We have", len(*stats.solutions), "solutions")
		//fmt.Println(Unload(board))
		return
	}

	for i, v := range board {
		if v.num != 0 {
			continue
		}
		for j := 1; j < len(v.avl); j++ {
			if !allSol && len((*stats).solutions) >= 10 {
				// We don't need more solutions
				(*stats).interrupted = true
				return
			}
			if v.avl[j] {
				solveBckTck(modBoard(board, i, j), stats, ck, allSol)
			}
		}
	}

}

// modBoard returns a modified board where where we have the number (b)
// at position (a)
func modBoard(board tBoard, a, b int) tBoard {
	board[a].num = uint8(b)
	return board
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

// Wrong returns an array of 81 bools that tells if a position
// is wrong (repeated in a row, column or block)
func Wrong(board tBoard) [81]bool {
	var w [81]bool
	for i, v := range board {

		// If we don't know the value it can't be wrong
		if v.num == 0 {
			w[i] = false
			continue
		}

		var count int
		r := getRowPos(i)
		c := getColPos(i)
		b := getBlkPos(i)
		for _, j := range r {
			if v.num == board[j].num {
				count++
			}
		}
		for _, j := range c {
			if v.num == board[j].num {
				count++
			}
		}
		for _, j := range b {
			if v.num == board[j].num {
				count++
			}
		}

		// Number should appear 3 times (once per row, column and block)
		w[i] = count != 3
	}
	return w
}

// anyWrong returns true if there is at least one duplicated number on
// a row, column or block.
func anyWrong(board tBoard) bool {
	l := Wrong(board)
	for _, v := range l {
		if v {
			return true
		}
	}
	return false
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

// countMissing counts the number of incognitas remaining on the board
func countMissing(b tBoard) int {
	count := 0
	for _, v := range b {
		if v.num == 0 {
			count++
		}
	}
	return count
}

// Deduced tStats method that returns the number of iterations done
// in order to solve the sudoku
func (s *tStats) Deduced() int {
	return s.deduced
}

// NumSolutions tStats method that returns the number of solutions found
func (s *tStats) NumSolutions() int {
	return len(s.solutions)
}

// BruteForce tStats method that tells if we applied brute-force or not
func (s *tStats) BruteForce() bool {
	return s.bruteForce
}

// Interrupted tStats method that tells if the calculation of solutions have been truncated.
// (There might be more solutions if yes)
func (s *tStats) Interrupted() bool {
	return s.interrupted
}
