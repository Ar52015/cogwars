package render

import (
	"fmt"

	"github.com/Ar52015/cogwars/src/input"
	"github.com/Ar52015/cogwars/src/types"
)

var (
	upCursor    = "˄"
	downCursor  = "v"
	leftCursor  = "<"
	rightCursor = ">"
)

func Render(grid *[][]string, move input.Direction, currPos types.PosVector, prevPos types.PosVector) {
	fmt.Print("\033[2J") // ANSI escape code to clear screen

	if move != input.Invalid {
		(*grid)[prevPos.Y][prevPos.X] = "•" // change previous position back to dot
	}

	switch move { // update wherever the cursor is now
	case input.North:
		(*grid)[currPos.Y][currPos.X] = upCursor
	case input.South:
		(*grid)[currPos.Y][currPos.X] = downCursor
	case input.West:
		(*grid)[currPos.Y][currPos.X] = leftCursor
	case input.East:
		(*grid)[currPos.Y][currPos.X] = rightCursor
	}

	printGrid(grid)
}

func printGrid(grid *[][]string) {
	for _, i := range *grid {
		for _, j := range i {
			fmt.Printf(" %s ", j)
		}
		fmt.Print("\n")
	}
}
