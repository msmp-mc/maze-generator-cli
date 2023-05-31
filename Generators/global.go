package Generators

type Maze struct {
	Walls  []Wall
	Width  uint
	Height uint
}

type Wall struct {
	IsVertical bool
	ID         int
}

// GenerateNewMaze generate a new maze with the given information.
//
// w is the width, h is the height and algo is the algorithm used
//
// Return an error if a problem occurs and nil if there are no errors
func GenerateNewMaze(w uint, h uint, algo func(*Maze) error) (Maze, error) {
	maze := Maze{Height: h, Width: w}
	err := algo(&maze)
	return maze, err
}

// generateWalls generate the default walls
func (m *Maze) generateWalls() []Wall {
	xV := (m.Width - 1) * m.Height
	xH := m.Width * m.Height
	x := xH + xV
	walls := make([]Wall, x)
	for i := uint(0); i < x; i++ {
		walls[i] = Wall{IsVertical: i >= xH, ID: int(i)}
	}
	m.Walls = walls
	return walls
}

func (m *Maze) RenderWalls() {

}
