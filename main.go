package main

import (
	"fmt"
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
	output0 := regexp.MustCompile(`-o [0-9a-zA-Z/.\-_]+`)
	unWidth := widthO.FindString(cli)
	unHeight := height0.FindString(cli)
	unOutput := output0.FindString(cli)
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
	if unOutput != "" {
		output := strings.ReplaceAll(unOutput, "-o ", "")
		err = m.Output(output)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully outputted at %s\n", output)
	}
}

func help() {
	println("------------------------------")
	println("HELP OF THE MAZE GENERATOR CLI")
	println("------------------------------")
	println("-w uint -> Width of the maze")
	println("-h uint -> Height of the maze")
}
