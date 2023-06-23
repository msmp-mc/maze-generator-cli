package Generators

import (
	"github.com/msmp-core/maze-generator-cli/utils"
	"strconv"
	"testing"
)

func TestWallsRandomisedKruskal(t *testing.T) {
	m, err := GenerateNewMaze(5, 5, 1, 0, 0, NewRandomisedKruskal)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(m.Walls) != 45 {
		t.Error(utils.FormatTestError("bad length of total walls", strconv.Itoa(45),
			strconv.Itoa(len(m.Walls))))
	}

	vertical := m.GetVerticalWalls()
	for _, wall := range vertical {
		if !wall.IsVertical {
			t.Error(utils.FormatTestError("non vertical walls in GetVerticalWalls", "true", "false"))
			break
		}
	}

	horizontal := m.GetHorizontalWalls()
	for _, wall := range horizontal {
		if wall.IsVertical {
			t.Error(utils.FormatTestError("vertical walls in GetHorizontalWalls", "false", "true"))
			break
		}
	}

	s := m.ToScheme()
	if len(s.Contents) != 5 {
		t.Error(utils.FormatTestError("scheme too long", "5", strconv.Itoa(len(s.Contents))))
		return
	}
}
