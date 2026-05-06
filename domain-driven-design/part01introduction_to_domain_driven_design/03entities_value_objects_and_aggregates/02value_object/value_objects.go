package valueobject

type Point struct {
	x int
	y int
}

func NewPoints(x, y int) Point {
	return Point{x: x, y: y}
}

const (
	DirectionUnknown = iota
	DirectionNorth
	DirectionSouth
	DirectionEast
	DirectionWest
)

func TrackPlayer() {
	currLocation := NewPoints(3, 4)
	currLocation = move(currLocation, DirectionNorth)
}

func move(currLocation Point, direction int) Point {
	switch direction {
	case DirectionNorth:
		return NewPoints(currLocation.x, currLocation.y+1)
	case DirectionSouth:
		return NewPoints(currLocation.x, currLocation.y-1)
	case DirectionEast:
		return NewPoints(currLocation.x+1, currLocation.y)
	case DirectionWest:
		return NewPoints(currLocation.x-1, currLocation.y)
	default:
		return currLocation
	}
}
