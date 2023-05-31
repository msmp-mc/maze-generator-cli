package Generators

import "fmt"

type Maze struct {
	Walls  []Wall
	Width  uint
	Height uint
	Cells  []Cell
}

type Wall struct {
	IsVertical bool
	IsPresent  bool
	ID         int
	CellsNear  []*Cell
}

type Cell struct {
	ID         int
	WallTop    *Wall
	WallBottom *Wall
	WallLeft   *Wall
	WallRight  *Wall
	MergedCell []*Cell
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

func (m *Maze) generateCells() []Cell {
	c := m.Width * m.Height
	cells := make([]Cell, c)
	for i := uint(0); i < m.Height; i++ {
		for j := uint(0); j < m.Width; j++ {
			jWall := m.GenJForLeftWallFromJOfCell(j)
			var wTop *Wall
			var wBottom *Wall
			var wLeft *Wall
			var wRight *Wall
			if i != 0 {
				wTop = &m.Walls[m.GenIDFromIJForWall(i-1, uint(jWall+1))]
			}
			if i != m.Height-1 {
				wBottom = &m.Walls[m.GenIDFromIJForWall(i, uint(jWall+1))]
			}
			if j != 0 {
				wLeft = &m.Walls[m.GenIDFromIJForWall(i, j)]
			}
			if j != m.Width-1 {
				wRight = &m.Walls[m.GenIDFromIJForWall(i, uint(jWall+2))]
			}
			cell := Cell{
				ID:         int(m.GenIDFromIJForCell(i, j)),
				WallTop:    wTop,
				WallBottom: wBottom,
				WallLeft:   wLeft,
				WallRight:  wRight,
				MergedCell: []*Cell{},
			}
			if wTop != nil {
				wTop.CellsNear = append(wTop.CellsNear, &cell)
			}
			if wBottom != nil {
				wBottom.CellsNear = append(wBottom.CellsNear, &cell)
			}
			if wLeft != nil {
				wLeft.CellsNear = append(wLeft.CellsNear, &cell)
			}
			if wRight != nil {
				wRight.CellsNear = append(wRight.CellsNear, &cell)
			}
			cells[m.GenIDFromIJForCell(i, j)] = cell
		}
	}
	m.Cells = cells
	return cells
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
			wall := m.Walls[m.GenIDFromIJForWall(i, j)]
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

// GenIDFromIJForWall generate the ID of the wall from it's coords representation (IJ)
//
// i is the number of rows
// j is the number of columns
//
// Return the id
func (m *Maze) GenIDFromIJForWall(i uint, j uint) uint {
	if j%2 == 0 {
		return m.Height*i + j/2 - 1*i
	}
	return m.GetHorizontalWallsNumber() + m.Height*i + (j-j%2)/2 - 1*i
}

// GenIDFromIJForCell generate the ID of the cell from it's coords representation (IJ)
//
// i is the number of rows
// j is the number of columns
//
// Return the id
func (m *Maze) GenIDFromIJForCell(i uint, j uint) uint {
	return m.Height*i + j
}

// GenJForLeftWallFromJOfCell generate the J coords for the wall at the left of the cell
//
// j is the number of columns
//
// Return the j (number of columns) for the left wall of the cell or -1 if j = 0
func (m *Maze) GenJForLeftWallFromJOfCell(j uint) int {
	return int(2*(j-1) + 1)
}

func (m *Maze) RenderWalls() {
	s := m.ToScheme()
	text := s.GenerateText()
	println(text)
}

func (s *Scheme) GenerateText() string {
	var l string
	for i := 0; i <= len(s.Contents[0]); i++ {
		if i%2 == 0 {
			l += " "
			continue
		}
		l += "_"
	}
	f := fmt.Sprintf("%s\n", l)
	for _, i := range s.Contents {
		f += fmt.Sprintf("|%s|\n", i)
	}
	return f
}

// mergeCells merge two cells and their group
//
// big is the cell with the biggest group.
// small is the cell with the smallest group.
func mergeCells(big *Cell, small *Cell) {
	for _, c := range small.MergedCell {
		c.ID = big.ID
	}
	small.ID = big.ID
	if len(small.MergedCell) == 0 {
		big.MergedCell = append(big.MergedCell, small.MergedCell...)
	}
	println(len(big.MergedCell))
	big.MergedCell = append(big.MergedCell, small)
}
