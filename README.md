# Maze Generator CLI

This is a CLI application generating a random maze.

## Usage

This application takes two parameters:
<!-- - `-h uint` - Set the height of the maze-->
<!-- - `-w uint` - Set the width of the maze-->
- `-s uint` - Size of one side of the maze (the maze is a square)

It has also another parameters:
- `-o string` - Set the output file (if it does not give, no output in file)
- `-d (0,1,2)` - Set the difficulty or 0 by default (0 = easy, 1 = hard, 2 = hardcore)
- `-g uint` - Number of random gates or 0 by default
- `-i uint` - Add "a hole" in the center when generating the maze
- `-help` - Show the help

Example:
```bash
./maze-generator -h 10 -w 10 -o maze.txt -d 1
```
It creates a hard 10x10 maze in the file called `maze.txt`.

## Technologies

- Go 1.20
- ~~(the anhgelus' brain)~~
