package aoc

// Point represents a point in 2D space
type Point struct {
	X int64
	Y int64
}

// ManhattanDistance returns the manhattan distance between tow points
func (a *Point) ManhattanDistance(b Point) int64 {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

// Point3 represents a point in 3D space
type Point3 struct {
	X int64
	Y int64
	Z int64
}

// ManhattanDistance returns the manhattan distance between tow points
func (a *Point3) ManhattanDistance(b Point3) int64 {
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
	//return !((a.Left > b.Right) || (a.Right < b.Left) || (a.Top > b.Bottom) || (a.Bottom < b.Top))
	i := a.Intersection(b)
	return !i.IsEmpty()
}
