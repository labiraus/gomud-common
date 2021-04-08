package core

//Direction is a list of cardianl directions
type Direction uint

const (
	ether Direction = iota
	north
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
	up
	down
	in
	out
)

//DirectionCount is the total number of directions
const DirectionCount = 13

//Reverse gets bac
func (dir Direction) Reverse() Direction {
	switch dir {
	case north:
		return south
	case northEast:
		return southWest
	case east:
		return west
	case southEast:
		return northWest
	case south:
		return north
	case southWest:
		return northEast
	case west:
		return east
	case northWest:
		return southEast
	case in:
		return out
	case out:
		return in
	case up:
		return down
	case down:
		return up
	default:
		return ether
	}
}
