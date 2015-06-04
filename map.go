package tmpgame

type Direction int

const (
	Undirected Direction = iota
	North
	South
	East
	West
)

type Position interface {
}

type Map interface {
	// From -> To positions
	Path(Position, Position) Waypoint
}

type Waypoint interface {
	Position() Position
	Next() Waypoint
}

func NewWaypoint(positions ...Position) Waypoint {
	waypoints := make([]waypoint, len(positions))
	head := &waypoints[0]
	tail := &waypoints[0]
	for i, w := range waypoints {
		w.pos = positions[i]
		w.next = tail
		tail = &w
	}
	return head
}

// waypoint is an implementation of the Waypoint interface
type waypoint struct {
	pos  Position
	next *waypoint
}

// Position returns the current position
func (w waypoint) Position() Position {
	return w.pos
}

// Next returns the next waypoint
func (w waypoint) Next() Waypoint {
	return w.next
}

func NewGrid(width, height int) *grid {
	return &grid{
		width:  width,
		height: height,
		tiles:  make([]tile, width*height, width*height),
	}
}

type tile bool

type idx int

// grid implements Map
type grid struct {
	width  int
	height int
	tiles  []tile
}

func (g grid) index(x int, y int) idx {
	return idx(x + (y * g.width))
}

func (g grid) xy(p idx) (int, int) {
	x := int(p) % g.width
	y := (int(p) - x) / g.height
	return x, y
}

func (g grid) move(p idx, dir Direction) idx {
	switch dir {
	case North:
		r := int(p) + g.width
		if r > len(g.tiles) {
			return p
		}
		return idx(r)
	case South:
		r := int(p) - g.width
		if r < 0 {
			return p
		}
		return idx(r)
	case East:
		r := int(p) + 1
		if r%g.width == 0 {
			return p
		}
		return idx(r)
	case West:
		if int(p)%g.width == 0 {
			return p
		}
		return idx(p - 1)
	}
	return p
}

// Path returns a Waypoint from a position to another position
func (g grid) Path(from Position, to Position) Waypoint {
	from = from.(*idx)
	to = to.(*idx)
	return nil
}
