package solver

var w string = "Working!"

// tCell represents the type of each cell of the board
type tCell struct {
	num     uint8
	fixed   bool
	options [9]bool
}

// tBoard represents all 81 cells of the board
type tBoard [81]tCell
