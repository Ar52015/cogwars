package input

type Direction int

const (
	_ Direction = iota
	North
	South
	West
	East
	NorthWest
	NorthEast
	SouthWest
	SouthEast
	Invalid
)

func Input(inp byte) Direction {
	switch inp {
	case 'w':
		return North
	case 'a':
		return West
	case 's':
		return South
	case 'd':
		return East
	default: // when something else is pressed
		return Invalid
	}
}
