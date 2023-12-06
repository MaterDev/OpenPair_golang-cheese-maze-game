package main

import "fmt"

// Represents the game maze.
type Maze struct {
	// The actual Grid itself
	Grid   []string
	Width  int
	Height int
}

// Mouse represents the player character's coordinates
type Mouse struct {
	X, Y int
}

// This is the main entrypoint for the game to  run when the terminal command is executed. `go run .`
func main() {
	// intialize a new mouse
	mouse := Mouse{X: 1, Y: 1}
	// intialize a new maze
	maze := createMaze()

	// print maze
	printMaze(maze, mouse)

}

// ! Function to print maze
func printMaze(inc_maze Maze, inc_mouse Mouse) {
	// TODO: Add the logic for managing the position of the mouse in the grid (If/Else)
	// * The mouse will be added to the Grid at print-time
	for y, line := range inc_maze.Grid {
		for x, char := range line {
			// Will print mouse `M` to replace character at given coordinate
			if inc_mouse.X == x && inc_mouse.Y == y {
				fmt.Print("M")
			} else {
				// When looping over a string, the characters are converted into unicode decimals that represent the character.
					// ? Unicode: https://en.wikipedia.org/wiki/List_of_Unicode_characters
					// string() will turn unicode numbers back into the appropriate glyph
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

// ! Function that will return the maze structure.
func createMaze() Maze {
	// Simple static maze layout
	layout := Maze{
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
	return layout
}
