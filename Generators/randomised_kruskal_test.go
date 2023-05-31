package Generators

import (
	"github.com/msmp-core/maze-generator-cli/utils"
	"strconv"
	"testing"
)

func TestWallsRandomisedKruskal(t *testing.T) {
	maze, err := GenerateNewMaze(5, 5, NewRandomisedKruskal)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(maze.Walls) != 45 {
		t.Error(utils.FormatTestError("bad length of total walls", strconv.Itoa(45),
			strconv.Itoa(len(maze.Walls))))
	}
}
