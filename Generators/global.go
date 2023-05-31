package Generators

import "fmt"

type Maze struct {
	Walls  []Wall
	Width  uint
	Height uint
}

type Wall struct {
	IsVertical bool
	IsPresent  bool
	ID         int
}

type Scheme struct {
	Contents []string
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
	xH := m.GetHorizontalWallsNumber()
	x := xH + m.GetVerticalWallsNumber()
	walls := make([]Wall, x)
	for i := uint(0); i < x; i++ {
		walls[i] = Wall{IsVertical: i >= xH, IsPresent: true, ID: int(i)}
	}
	m.Walls = walls
	return walls
}

func (m *Maze) GetHorizontalWallsNumber() uint {
	return m.Width * m.Height
}

func (m *Maze) GetVerticalWallsNumber() uint {
	return (m.Width - 1) * m.Height
}

func (m *Maze) GetVerticalWalls() []Wall {
	return m.Walls[m.GetHorizontalWallsNumber():]
}

func (m *Maze) GetHorizontalWalls() []Wall {
	return m.Walls[:m.GetHorizontalWallsNumber()-1]
}

func (m *Maze) ToScheme() Scheme {
	contents := make([]string, m.Height)
	for i := uint(0); i < m.Height; i++ {
		str := ""
		// size max = width+(width-1) because we must not forget the vertical walls
		for j := uint(0); j < m.Width+(m.Width-1); j++ {
			wall := m.Walls[m.GenIDFromIJ(i, j)]
			if !wall.IsPresent {
				str += " "
				continue
			}
			if wall.IsVertical {
				str += "|"
			} else {
				str += "_"
			}
		}
		contents[i] = str
	}
	return Scheme{Contents: contents}
}

// GenIDFromIJ generate the ID of the wall from it's coords representation (IJ)
//
// i is the number of rows
// j is the number of columns
//
// Return the id
func (m *Maze) GenIDFromIJ(i uint, j uint) uint {
	if j%2 == 0 {
		return m.Height*i + j/2 - 1*i
	}
	return m.GetHorizontalWallsNumber() + m.Height*i + (j-j%2)/2 - 1*i
}

func (m *Maze) RenderWalls() {
	s := m.ToScheme()
	text := s.GenerateText()
	println(text)
}

func (s *Scheme) GenerateText() string {
	l := " "
	for i := 0; i < len(s.Contents[0]); i++ {
		l += "_"
	}
	f := fmt.Sprintf("%s\n", l)
	for _, i := range s.Contents {
		f += fmt.Sprintf("|%s|\n", i)
	}
	f += fmt.Sprintf("%s\n", l)
	return f
}
