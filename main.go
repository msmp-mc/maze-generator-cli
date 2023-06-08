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
	//widthO := regexp.MustCompile(`-w [0-9]+`)
	//height0 := regexp.MustCompile(`-h [0-9]+`)
	size0 := regexp.MustCompile(`-s [0-9]+`)
	output0 := regexp.MustCompile(`-o [0-9a-zA-Z/.\-_]+`)
	difficulty0 := regexp.MustCompile(`-d [0-9]+`)
	help0 := regexp.MustCompile(`-help`)
	//unWidth := widthO.FindString(cli)
	//unHeight := height0.FindString(cli)
	unSize := size0.FindString(cli)
	unOutput := output0.FindString(cli)
	unDifficulty := difficulty0.FindString(cli)
	t := help0.FindString(cli)
	if t != "" {
		help()
		return
	}
	//if unHeight == "" || unWidth == "" {
	//	help()
	//	return
	//}
	if unSize == "" {
		help()
		return
	}
	//strWidth := strings.ReplaceAll(unWidth, "-w ", "")
	//strHeight := strings.ReplaceAll(unHeight, "-h ", "")
	strSize := strings.ReplaceAll(unSize, "-s ", "")

	//w, err := strconv.Atoi(strWidth)
	//if err != nil {
	//	panic(err)
	//}
	//h, err := strconv.Atoi(strHeight)
	//if err != nil {
	//	panic(err)
	//}
	s, err := strconv.Atoi(strSize)
	if err != nil {
		panic(err)
	}
	d := 0
	if unDifficulty != "" {
		strDifficulty := strings.ReplaceAll(unDifficulty, "-d ", "")
		d, err = strconv.Atoi(strDifficulty)
		if err != nil {
			panic(err)
		}
	}
	m, err := Generators.GenerateNewMaze(uint(s), uint(s), uint(d), Generators.NewRandomisedKruskal)
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
	println("Required arguments:")
	//println("  -w uint -> Width of the maze")
	//println("  -h uint -> Height of the maze\n")
	println("  -s uint -> Size of one side of the maze\n")
	println("Optional arguments:")
	println("  -d uint -> Difficulty of the maze")
	println("  -o string -> Output file of the new maze")

}
