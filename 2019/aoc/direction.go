package aoc

// Direction represents a cardinal direction
type Direction int

// The four directions (and their compass equivalents)
const (
	Right Direction = 0
	Up    Direction = 90
	Left  Direction = 180
	Down  Direction = 270
	East  Direction = Right
	North Direction = Up
	West  Direction = Left
	South Direction = Down
)

// Ccw returns a direction rotated 90 degrees counterclockwise
func (d Direction) Ccw() Direction {
	return Direction((d + 90) % 360)
}

// Cw returns a direction rotated 90 degrees clockwise
func (d Direction) Cw() Direction {
	return Direction((d + 270) % 360)
}

// Inverse turns a direction around 180 degrees
func (d Direction) Inverse() Direction {
	return Direction((d + 180) % 360)
}

// DeltaX indicates the change in the X coordinate when you move in a direction (left is negative)
func (d Direction) DeltaX() int64 {
	switch d {
	case Left:
		return -1
	case Right:
		return 1
	default:
		return 0
	}
}

// DeltaY indicates the change in the Y coordinate when you move in a direction (up is negative)
func (d Direction) DeltaY() int64 {
	switch d {
	case Up:
		return -1
	case Down:
		return 1
	default:
		return 0
	}
}
