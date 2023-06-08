package main

import (
	"fmt"
	"github.com/msmp-core/maze-generator-cli/CLI"
	"github.com/msmp-core/maze-generator-cli/Generators"
	"os"
)

func main() {
	cli := ""
	for i, a := range os.Args {
		if i == 0 {
			continue
		}
		cli += a + " "
	}
	options := []*CLI.Option{
		{
			ID: "s", Help: "Size of one side of the maze", ArgsRegex: `[0-9]+`, Required: true, IsInt: true, Disabled: false,
		},
		{
			ID: "w", Help: "Width of the maze", ArgsRegex: `[0-9]+`, Required: true, IsInt: true, Disabled: true,
		},
		{
			ID: "h", Help: "Height of the maze", ArgsRegex: `[0-9]+`, Required: true, IsInt: true, Disabled: true,
		},
		{
			ID: "o", Help: "Output file", ArgsRegex: `[0-9a-zA-Z/.\-_]+`, Required: false, IsInt: false, Disabled: false,
		},
		{
			ID: "d", Help: "Difficulty of the maze", ArgsRegex: `[0-9]+`, Required: false, IsInt: true, Disabled: false,
		},
		{
			ID: "help", Help: "Show this help", ArgsRegex: ``, Required: false, IsInt: false, Disabled: false,
		},
	}
	app := CLI.CLI{Options: options}
	if len(os.Args) == 1 {
		app.Help()
		return
	}
	got, err := app.Parse(cli)
	if err != nil {
		println(err.Error())
		app.Help()
		return
	}
	var s int
	var out string
	d := 0
	for _, g := range got {
		switch g.ID {
		case "s":
			s = g.Value.(int)
		case "d":
			d = g.Value.(int)
		case "o":
			out = g.Value.(string)
		case "help":
			app.Help()
			return
		}
	}

	m, err := Generators.GenerateNewMaze(uint(s), uint(s), uint(d), Generators.NewRandomisedKruskal)
	if err != nil {
		panic(err)
	}
	m.RenderWalls()
	if out != "" {
		err = m.Output(out)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully outputted at %s\n", out)
	}
}
