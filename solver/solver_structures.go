package solver

// tCell represents the type of each cell of the board
// num: The number to display on that cell
// avl[0]: The first position is set to true if the number is known from the start
// avl[1-9]: Each position represents if that number is possible on that cell
type tCell struct {
	num uint8
	avl [10]bool
}

// tBoard represents all 81 cells of the board
type tBoard [81]tCell

// tStats is a type that contains result information
type tStats struct {
	solved bool
}
