package core

//Directions provides a lookup to get directions from strings
var Directions map[string]Direction = map[string]Direction{
	"north":     north,
	"n":         north,
	"south":     south,
	"s":         south,
	"east":      east,
	"e":         east,
	"west":      west,
	"w":         west,
	"northwest": northWest,
	"nw":        northWest,
	"northeast": northEast,
	"ne":        northEast,
	"southeast": southEast,
	"se":        southEast,
	"southwest": southWest,
	"sw":        southWest,
	"up":        up,
	"u":         up,
	"down":      down,
	"d":         down,
	"in":        in,
	"out":       out,
}

//String returns the name of the direction
func (dir Direction) String() string {
	switch dir {
	case north:
		return "north"
	case northEast:
		return "north east"
	case east:
		return "east"
	case southEast:
		return "south east"
	case south:
		return "south"
	case southWest:
		return "south west"
	case west:
		return "west"
	case northWest:
		return "north west"
	case in:
		return "in"
	case out:
		return "out"
	case up:
		return "up"
	case down:
		return "down"
	case ether:
		return "the ether"
	default:
		return "nowhere"
	}
}

//LeaveString returns the string of when someone leaves in a direction
func (dir Direction) LeaveString() string {
	switch dir {
	case north:
		return "to the north"
	case northEast:
		return "to the north east"
	case east:
		return "to the east"
	case southEast:
		return "to the south east"
	case south:
		return "to the south"
	case southWest:
		return "to the south west"
	case west:
		return "to the west"
	case northWest:
		return "to the north west"
	case in:
		return "inside"
	case out:
		return "outside"
	case up:
		return "going up"
	case down:
		return "going down"
	case ether:
		return "to the ether"
	default:
		return "to nowhere"
	}
}

//EnterString returns the string of when someone enters from a direction
func (dir Direction) EnterString() string {
	switch dir {
	case north:
		return "from the north"
	case northEast:
		return "from the north east"
	case east:
		return "from the east"
	case southEast:
		return "from the south east"
	case south:
		return "from the south"
	case southWest:
		return "from the south west"
	case west:
		return "from the west"
	case northWest:
		return "from the north west"
	case in:
		return "from inside"
	case out:
		return "from outside"
	case up:
		return "from below"
	case down:
		return "from above"
	case ether:
		return "from the ether"
	default:
		return "from nowhere"
	}
}
