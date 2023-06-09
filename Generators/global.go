package Generators

import (
	"fmt"
	"github.com/msmp-core/maze-generator-cli/utils"
	"os"
)

type Maze struct {
	Walls      []Wall
	Width      uint
	Height     uint
	Cells      []*Cell
	Difficulty uint
	Gates      uint
	GatesLoc   []GateLoc
	Inner uint
}

type GateLoc struct {
	I int
	J int
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
	MergedCell *[]*Cell
	// MergedRef refers to the parent cell containing the full MergedCell
	MergedRef *Cell
	Disabled bool
}

type Scheme struct {
	Contents []string
}

// GenerateNewMaze generate a new maze with the given information.
//
// w is the width, h is the height, difficulty is the difficulty (0 for easy, 1 for hard and 2 for hardcore), gates is
// the number of gates present and algo is the algorithm used
//
// Return an error if a problem occurs and nil if there are no errors
func GenerateNewMaze(w uint, h uint, difficulty uint, gates uint, inner uint, algo func(*Maze) error) (Maze, error) {
	maze := Maze{Height: h, Width: w, Difficulty: difficulty, Gates: gates, Inner: inner}
	err := algo(&maze)
	maze.handleGates()
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

// generateCells generate the default cells
func (m *Maze) generateCells() []*Cell {
	c := m.Width * m.Height
	cells := make([]*Cell, c)
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
				wLeft = &m.Walls[m.GenIDFromIJForWall(i, uint(jWall))]
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
			}
			cell.MergedCell = &[]*Cell{&cell}
			cell.MergedRef = &cell
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
			cells[m.GenIDFromIJForCell(i, j)] = &cell
		}
	}
	m.Cells = cells
	m.innerHole()
	return cells
}

func (m *Maze) innerHole() {
	for _, c := range m.Cells {
		if !m.isInHole(m.GenIJFromIDForCell(uint(c.ID))) {
			continue
		}
		c.Disabled = true
	}
}

func (m *Maze) isInHole(i uint, j uint) bool {
	cI := m.Width/2
	cJ := m.Height/2
	mid := m.Inner/2
	return i >= cI - mid && i < cI + mid && j >= cJ - mid && j < cJ + mid
}

// GetHorizontalWallsNumber return the number of horizontal walls
func (m *Maze) GetHorizontalWallsNumber() uint {
	return m.Width * m.Height
}

// GetVerticalWallsNumber return the number of vertical walls
func (m *Maze) GetVerticalWallsNumber() uint {
	return (m.Width - 1) * m.Height
}

// GetVerticalWalls return the vertical walls
func (m *Maze) GetVerticalWalls() []Wall {
	return m.Walls[m.GetHorizontalWallsNumber():]
}

// GetHorizontalWalls return the horizontal walls
func (m *Maze) GetHorizontalWalls() []Wall {
	return m.Walls[:m.GetHorizontalWallsNumber()-1]
}

// ToScheme turn a maze into the scheme representing the maze
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
				if len(wall.CellsNear) == 2 && (wall.CellsNear[0].Disabled || wall.CellsNear[1].Disabled) {
					str += "-"
					continue
				}
				str += "|"
			} else {
				if len(wall.CellsNear) == 2 && wall.CellsNear[0].Disabled {
					str += "X"
					continue
				}
				str += "_"
			}
		}
		contents[i] = str
	}
	return Scheme{Contents: contents}
}

// GenIDFromIJForWall generate the ID of the wall from it's coords representation (IJ)
//
// i is the number of rows,
// j is the number of columns
//
// Return the id
func (m *Maze) GenIDFromIJForWall(i uint, j uint) uint {
	if j%2 == 0 {
		return m.Height*i + j/2
	}
	return m.GetHorizontalWallsNumber() + m.Height*i + (j-j%2)/2 - i*1
}

// GenIDFromIJForCell generate the ID of the cell from it's coords representation (IJ)
//
// i is the number of rows,
// j is the number of columns
//
// Return the id
func (m *Maze) GenIDFromIJForCell(i uint, j uint) uint {
	return m.Height*i + j
}

func (m *Maze) GenIJFromIDForCell(id uint) (uint, uint) {
	// id = m.Height*i + j
	j := id%m.Height
	return (id-j)/m.Height, j
}

// GenJForLeftWallFromJOfCell generate the J coords for the wall at the left of the cell
//
// j is the number of columns
//
// Return the j (number of columns) for the left wall of the cell or -1 if j = 0
func (m *Maze) GenJForLeftWallFromJOfCell(j uint) int {
	return int(2*j - 1)
}

func (m *Maze) GenIJFromIDOfWall(id uint) (uint, uint) {
	i := (id - id%m.Height) / m.Height
	return i, id - (i)*m.Height
}

// RenderWalls print the walls of the maze
func (m *Maze) RenderWalls() {
	s := m.ToScheme()
	text := s.GenerateText(m)
	println(text)
}

// GenerateText return the text representing a scheme of the maze
func (s *Scheme) GenerateText(m *Maze) string {
	var l string
	for i := 0; i <= len(s.Contents[0]); i++ {
		if i%2 == 0 {
			l += " "
			continue
		}
		done := false
		for _, g := range m.GatesLoc {
			if done {
				continue
			}
			if g.I == -1 && g.J == i {
				l += " "
				done = true
			}
		}
		if done {
			continue
		}
		l += "_"
	}
	f := fmt.Sprintf("%s\n", l)
	for id, i := range s.Contents {
		left := true
		right := true
		for _, g := range m.GatesLoc {
			if g.I != id {
				continue
			}
			if g.J == -1 {
				left = false
			} else if g.J == int(2*m.Width) {
				right = false
			}
		}
		if right && left {
			f += fmt.Sprintf("|%s|\n", i)
		} else if right {
			f += fmt.Sprintf(" %s|\n", i)
		} else if left {
			f += fmt.Sprintf("|%s \n", i)
		}
	}
	return f
}

func (m *Maze) Output(path string) error {
	scheme := m.ToScheme()
	return os.WriteFile(path, []byte(scheme.GenerateText(m)), 0664)
}

// mergeCells merge two cells and their group
//
// big is the cell with the biggest group.
// small is the cell with the smallest group.
func mergeCells(big *Cell, small *Cell) {
	merged := append(*big.MergedRef.MergedCell, *small.MergedRef.MergedCell...)
	big.MergedRef.MergedCell = &merged
	small.MergedRef = big.MergedRef
	big.MergedRef.updateCells()
	small.updateMergedRef(big.MergedRef)
}

// updateCells update the cells after a merge
func (m *Cell) updateCells() {
	for _, c := range *m.MergedCell {
		if c.ID == m.ID {
			continue
		}
		c.ID = m.ID
		c.MergedCell = m.MergedCell
	}
}

// updateMergedRef update the merged reference after a merge
func (m *Cell) updateMergedRef(newRef *Cell) {
	for _, c := range *m.MergedRef.MergedCell {
		c.MergedRef = newRef
	}
	m.MergedRef.MergedRef = newRef
}

// handleGates handle the gates of the maze
func (m *Maze) handleGates() {
	if m.Gates == 0 {
		println("No gates to handle")
		return
	}
	for i := uint(0); i < m.Gates; i++ {
		rand := utils.RandMax(3)
		var r int
		switch rand {
		case 0:
			// top
			r = utils.RandMax(m.Width - 1)
			if r%2 == 0 {
				if r == int(m.Width-1) {
					r--
				} else {
					r++
				}
			}
			m.GatesLoc = append(m.GatesLoc, GateLoc{I: -1, J: r})
		case 1:
			// bottom
			r = utils.RandMax(m.Width - 1)
			if r%2 == 0 {
				if r == int(m.Width-1) {
					r--
				} else {
					r++
				}
			}
			m.GatesLoc = append(m.GatesLoc, GateLoc{I: int(m.Height - 1), J: r})
			wall := &m.Walls[m.GenIDFromIJForCell(m.Height-1, uint(r))]
			wall.IsPresent = false
		case 2:
			// left
			r = utils.RandMax(m.Height - 1)
			m.GatesLoc = append(m.GatesLoc, GateLoc{I: r, J: -1})
		case 3:
			// right
			r = utils.RandMax(m.Height - 1)
			w := 2 * m.Width
			m.GatesLoc = append(m.GatesLoc, GateLoc{I: r, J: int(w)})
		}
	}
}
