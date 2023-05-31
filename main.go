package main

import (
	"github.com/msmp-core/maze-generator-cli/Generators"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}
	cli := ""
	for i, a := range os.Args {
		if i == 0 {
			continue
		}
		cli += a + " "
	}
	widthO := regexp.MustCompile(`-w [0-9]+`)
	height0 := regexp.MustCompile(`-h [0-9]+`)
	unWidth := widthO.FindString(cli)
	unHeight := height0.FindString(cli)
	if unHeight == "" || unWidth == "" {
		help()
		return
	}
	strWidth := strings.ReplaceAll(unWidth, "-w ", "")
	strHeight := strings.ReplaceAll(unHeight, "-h ", "")

	w, err := strconv.Atoi(strWidth)
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(strHeight)
	if err != nil {
		panic(err)
	}
	m, err := Generators.GenerateNewMaze(uint(w), uint(h), Generators.NewRandomisedKruskal)
	if err != nil {
		panic(err)
	}
	m.RenderWalls()
}

func help() {
	println("------------------------------")
	println("HELP OF THE MAZE GENERATOR CLI")
	println("------------------------------")
	println("-w uint -> Width of the maze")
	println("-h uint -> Height of the maze")
}
