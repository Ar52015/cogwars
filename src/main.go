package main

import (
	"os"
	"os/exec"

	"github.com/Ar52015/cogwars/src/input"
	"github.com/Ar52015/cogwars/src/render"
	"github.com/Ar52015/cogwars/src/types"
	"github.com/Ar52015/cogwars/src/update"
)

var (
	grid   = [][]string{{"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "˄", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}, {"•", "•", "•", "•", "•", "•", "•", "•", "•", "•", "•"}}
	curPos = types.PosVector{X: 5, Y: 5} // hardcoded start position
)

func loop() {
	// MAIN GAME LOOP

	// take input
	var inpByte []byte = make([]byte, 1)
	os.Stdin.Read(inpByte)
	inp := inpByte[0]

	// 1. process input
	move := input.Invalid
	move = input.Input(inp)

	// 2. update the grid and the current position
	prevPos := curPos
	update.Update(move, &curPos)

	// 3. render on terminal
	render.Render(&grid, move, curPos, prevPos)
}

func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	for {
		loop()
	}
}
