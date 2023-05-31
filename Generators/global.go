package Generators

type Maze struct {
	VerticalWalls   []Wall
	HorizontalWalls []Wall
	Width           uint
	Height          uint
}

type Wall struct {
	IsVertical  bool
	IsPresent   bool
	ID          int
	IsRemovable bool
}

// GenerateNewMaze generate a new maze with the given information.
//
// w is the width, h is the height and algo is the algorithm used
//
// Return an error if a problem occurs and nil if there are no errors
func GenerateNewMaze(w uint, h uint, algo func(*Maze) error) error {
	maze := Maze{Height: h, Width: w}
	return algo(&maze)
}

// CalcID calculate the ID of the Wall
//
// l is the length of the current ID (the width of the maze for a horizontal wall).
// i is the column number.
// j is the row number.
//
// Return the ID
func CalcID(l uint, i uint, j uint) int {
	return int(l*(j-1) + i)
}
