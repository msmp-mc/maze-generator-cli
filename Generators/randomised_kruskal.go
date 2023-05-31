package Generators

import (
	"fmt"
	"github.com/msmp-core/maze-generator-cli/utils"
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
	m.RenderWalls()
	println("Base generated!")
	println("Merging...")
	i := 0
	for !m.isFinished() {
		err := m.mergeRandomly()
		if err != nil {
			return err
		}
		if i > 30 {
			break
		}
		i++
	}
	return nil
}

func (m *kruskal) mergeRandomly() error {
	id := utils.RandMax(uint(len(m.Cells) - 1))
	cell := m.Cells[id]
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
	wall.IsPresent = false
	var mergeWith *Cell
	if &cell == wall.CellsNear[0] {
		mergeWith = wall.CellsNear[1]
	} else {
		mergeWith = wall.CellsNear[0]
	}
	println(len(cell.MergedCell), len(mergeWith.MergedCell))
	if len(cell.MergedCell) >= len(mergeWith.MergedCell) {
		mergeCells(&cell, mergeWith)
		return nil
	}
	mergeCells(mergeWith, &cell)
	return nil
}

func (m *kruskal) isFinished() bool {
	for _, cell := range m.Cells {
		if uint(len(cell.MergedCell)) == m.Width*m.Height-1 {
			return true
		}
	}
	return false
}
