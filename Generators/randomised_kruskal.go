package Generators

import (
	"fmt"
	"github.com/msmp-core/maze-generator-cli/utils"
	"math"
)

type kruskal struct {
	*Maze
}

// NewRandomisedKruskal is a func generating a new maze with the Randomized Kruskal's Algorithm
func NewRandomisedKruskal(b *Maze) error {
	m := kruskal{b}
	println("Generating base...")
	m.generateWalls()
	m.generateCells()
	println("Base generated!")
	println("Merging...")
	i := 0
	for !m.isFinished() {
		if i%10000 == 0 {
			m.RenderWalls()
		}
		err := m.mergeRandomly()
		if err != nil {
			return err
		}
		i++
	}
	if len(m.Cells) != int(m.Width*m.Height) {
		return fmt.Errorf("bad length of cells: %d", len(m.Cells))
	}
	println("Merging finished!")
	println("Starting the difficulty generation...")
	if m.Difficulty > 0 {
		size := len(m.Walls)
		for n := 0; n < int(math.Floor(float64(size/4)*(float64(m.Difficulty)*0.05))); n++ {
			valid := false
			var wall Wall
			for !valid {
				id := uint(utils.RandMax(uint(size - 1)))
				i, j := m.GenIJFromIDOfWall(id)
				if i == 0 || i == m.Width-1 || j == m.Height-1 {
					continue
				}
				wall = m.Walls[id]
				if wall.IsVertical {
					if !(wall.CellsNear[0].Disabled || wall.CellsNear[1].Disabled) {
						valid = true
					}
				} else {
					if !wall.CellsNear[0].Disabled {
						valid = true
					}
				}
			}
			wall.IsPresent = false
		}
	}
	println("Merging done!")
	return nil
}

// mergeRandomly merge two random cells
func (m *kruskal) mergeRandomly() error {
	valid := false
	var cell *Cell
	for !valid {
		id := utils.RandMax(uint(len(m.Cells) - 1))
		cell = m.Cells[id]
		valid = !cell.Disabled
	}
	var walls []*Wall
	if cell.WallLeft != nil {
		walls = append(walls, cell.WallLeft)
	}
	if cell.WallRight != nil {
		walls = append(walls, cell.WallRight)
	}
	if cell.WallTop != nil {
		walls = append(walls, cell.WallTop)
	}
	if cell.WallBottom != nil {
		walls = append(walls, cell.WallBottom)
	}
	if len(walls) < 2 {
		return fmt.Errorf("only %d wall was set for the cell with the id %d", len(walls), cell.ID)
	}
	wId := utils.RandMax(uint(len(walls) - 1))
	wall := walls[wId]
	if wall == nil {
		return fmt.Errorf("the wall with the id %d is nil", wId)
	}
	if !wall.IsPresent {
		return nil
	}
	var mergeWith *Cell
	if len(wall.CellsNear) != 2 {
		return fmt.Errorf("the wall with the id %d have only %d cells near", wall.ID, len(wall.CellsNear))
	}
	if cell.ID != wall.CellsNear[0].ID {
		mergeWith = wall.CellsNear[0]
	} else if cell.ID != wall.CellsNear[1].ID {
		mergeWith = wall.CellsNear[1]
	} else {
		return nil
	}
	if mergeWith.Disabled {
		return nil
	}
	wall.IsPresent = false
	if mergeWith.ID == cell.ID {
		return fmt.Errorf("the cell with the id %d is the same as the cell with the id %d", mergeWith.ID, cell.ID)
	}
	if len(*cell.MergedRef.MergedCell) >= len(*mergeWith.MergedRef.MergedCell) {
		mergeCells(cell, mergeWith)
		m.RenderWalls()
		return nil
	}
	mergeCells(mergeWith, cell)
	m.RenderWalls()
	return nil
}

// isFinished check if the maze is finished
//
// Return true if the maze is finished, false otherwise
func (m *kruskal) isFinished() bool {
	return uint(len(*m.Cells[0].MergedRef.MergedCell)) == m.Width*m.Height-m.Inner*m.Inner
}
