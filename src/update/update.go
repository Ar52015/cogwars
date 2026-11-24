package update

import (
	"github.com/Ar52015/cogwars/src/input"
	"github.com/Ar52015/cogwars/src/types"
)

func Update(move input.Direction, currPos *types.PosVector) {
	switch move {
	case input.North:
		if currPos.Y == 0 {
			break
		}
		currPos.Y--
	case input.East:
		if currPos.X == 10 {
			break
		}
		currPos.X++
	case input.West:
		if currPos.X == 0 {
			break
		}
		currPos.X--
	case input.South:
		if currPos.Y == 10 {
			break
		}
		currPos.Y++
	case input.Invalid:
		return
	}
}
