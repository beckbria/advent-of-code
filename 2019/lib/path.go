package lib

import "math"

// Point represents a point in 2D space
type Point struct {
	X int64
	Y int64
}

// ManhattanDistance returns the manhattan distance between tow points
func (a *Point) ManhattanDistance(b *Point) int64 {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

// SlopeTo finds the slope of a line from one Point to another
func (a *Point) SlopeTo(to *Point) Fraction {
	run := to.X - a.X
	rise := to.Y - a.Y
	return NewFraction(rise, run)
}

// Equals returns true if two points are equal
func (a *Point) Equals(b *Point) bool {
	return a.X == b.X && a.Y == b.Y
}

// Neighbors returns a list of cells immediately adjacent to a point
func (pt *Point) Neighbors() []Point {
	return []Point{
		Point{X: pt.X - 1, Y: pt.Y},
		Point{X: pt.X + 1, Y: pt.Y},
		Point{X: pt.X, Y: pt.Y - 1},
		Point{X: pt.X, Y: pt.Y + 1},
	}
}

// Point3 represents a point in 3D space
type Point3 struct {
	X int64
	Y int64
	Z int64
}

// ManhattanDistance returns the manhattan distance between tow points
func (a *Point3) ManhattanDistance(b *Point3) int64 {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y) + Abs(a.Z-b.Z)
}

// Rectangle represents a 2D rectangle
type Rectangle struct {
	Left   int64
	Top    int64
	Right  int64 // inclusive
	Bottom int64 // inclusive
}

// IsEmpty returns true if a rectangle has no area
func (rc *Rectangle) IsEmpty() bool {
	return (rc.Right <= rc.Left) || (rc.Bottom <= rc.Top)
}

// Intersection returns the intersection of two rectangles
func (a *Rectangle) Intersection(b *Rectangle) Rectangle {
	i := Rectangle{
		Left:   Max(a.Left, b.Left),
		Top:    Max(a.Top, b.Top),
		Right:  Min(a.Right, b.Right),
		Bottom: Min(a.Bottom, b.Bottom)}
	return i
}

// Intersects returns true if a rectangle intersects another
func (a *Rectangle) Intersects(b *Rectangle) bool {
	i := a.Intersection(b)
	return !i.IsEmpty()
}

// SameSlope returns true if the two slopes (in rise/run format) are equal
func SameSlope(a, b *Fraction) bool {
	if a.Numerator == 0 && b.Numerator == 0 {
		// Horizontal
		return (a.Denominator < 0) == (b.Denominator < 0)
	}

	if a.Denominator == 0 && b.Denominator == 0 {
		// Vertical
		return (a.Numerator < 0) == (b.Numerator < 0)
	}

	return a.Equals(b)
}

// SlopeToRadians converts a slope to the number of radians corresponding to that angle.
// As is traditional in math, 0 degrees is to the right, Pi/2 is straight up, Pi is to the left, and 270 is straight down
func SlopeToRadians(s *Fraction) float64 {
	if s.Denominator == 0 {
		if s.Numerator >= 0 {
			return PiOver2
		} else {
			return 3 * PiOver2
		}
	}

	if s.Numerator == 0 {
		if s.Denominator >= 0 {
			return 2 * math.Pi
		} else {
			return math.Pi
		}
	}

	// Now that we've got the pesky division by zero risks out of the way, this becomes a simple arctangent
	return math.Atan2(float64(s.Numerator), float64(s.Denominator))
}
