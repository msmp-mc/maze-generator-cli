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
	println("Base generated!")
	println("Merging...")
	i := 0
	for !m.isFinished() {
		if i%1000 == 0 {
			m.RenderWalls()
		}
		err := m.mergeRandomly()
		if err != nil {
			return err
		}
		i++
	}
	println("Merging done!")
	if len(m.Cells) != int(m.Width*m.Height) {
		return fmt.Errorf("bad length of cells: %d", len(m.Cells))
	}
	var f string
	for i, cell := range m.Cells {
		if i != 0 && uint(i)%m.Height == 0 {
			println(f)
			f = ""
		}
		f += fmt.Sprintf("%d ", cell.ID)
	}
	println(f)
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
	if mergeWith.ID == cell.ID {
		return fmt.Errorf("the cell with the id %d is the same as the cell with the id %d", mergeWith.ID, cell.ID)
	}
	if len(*cell.MergedRef.MergedCell) >= len(*mergeWith.MergedRef.MergedCell) {
		mergeCells(cell, mergeWith)

		var f string
		for i, c := range m.Cells {
			if i != 0 && uint(i)%m.Height == 0 {
				println(f)
				f = ""
			}
			f += fmt.Sprintf("%d ", c.ID)
		}
		println(f)

		return nil
	}
	mergeCells(mergeWith, cell)

	var f string
	for i, c := range m.Cells {
		if i != 0 && uint(i)%m.Height == 0 {
			println(f)
			f = ""
		}
		f += fmt.Sprintf("%d ", c.ID)
	}
	println(f)

	return nil
}

func (m *kruskal) isFinished() bool {
	return uint(len(*m.Cells[0].MergedRef.MergedCell)) == m.Width*m.Height
}
