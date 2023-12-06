package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Mouse represents the player's character.
type Mouse struct {
	X, Y int
}

// Maze represents the game maze.
type Maze struct {
	Grid   []string
	Width  int
	Height int
}

func main() {
	mouse := Mouse{X: 1, Y: 1}
	maze := createMaze()

	reader := bufio.NewReader(os.Stdin)
	for {
		clearScreen()
		printMaze(maze, mouse)
		fmt.Print("Move (WASD): ")
		input, _ := reader.ReadString('\n')
		handleInput(strings.TrimSpace(input), &mouse, maze)

		if mouse.X == 1 && mouse.Y == 8 {
			fmt.Println("You've found the cheese! Game over.")
			break
		}
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}



func createMaze() Maze {
	// Simple static maze layout
	return Maze{
		Grid: []string{
			"##########",
			"#        #",
			"# ####### #",
			"# #     # #",
			"# # ### # #",
			"# # # # # #",
			"#   # #   #",
			"####### ###",
			"#C       #",
			"##########",
		},
		Width:  10,
		Height: 10,
	}
}

func printMaze(maze Maze, mouse Mouse) {
	for y, line := range maze.Grid {
		for x, char := range line {
			if mouse.X == x && mouse.Y == y {
				fmt.Print("M")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func handleInput(input string, mouse *Mouse, maze Maze) {
	switch input {
	case "w":
		if maze.Grid[mouse.Y-1][mouse.X] == ' ' || maze.Grid[mouse.Y-1][mouse.X] == 'C' {
			mouse.Y--
		}
	case "a":
		if maze.Grid[mouse.Y][mouse.X-1] == ' ' || maze.Grid[mouse.Y][mouse.X-1] == 'C' {
			mouse.X--
		}
	case "s":
		if maze.Grid[mouse.Y+1][mouse.X] == ' ' || maze.Grid[mouse.Y+1][mouse.X] == 'C' {
			mouse.Y++
		}
	case "d":
		if maze.Grid[mouse.Y][mouse.X+1] == ' ' || maze.Grid[mouse.Y][mouse.X+1] == 'C' {
			mouse.X++
		}
	}
}