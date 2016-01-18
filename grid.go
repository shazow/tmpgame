package tmpgame

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// NewWaypoint returns a linked-list implementation of a Waypoint
func NewWaypoint(positions ...Position) Waypoint {
	if len(positions) == 1 {
		return &waypoint{pos: positions[0]}
	}
	if len(positions) == 0 {
		return nil
	}
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

// Add injects a position into the linked-list of waypoints.
// TODO: Use a pool for waypoints to reduce allocations?
func (w *waypoint) Add(pos Position) *waypoint {
	w.next = &waypoint{pos: pos, next: w.next}
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

type xy struct {
	x int
	y int
}

// grid implements Map
type grid struct {
	width  int
	height int
	tiles  []tile
}

func (g grid) in(pos xy) bool {
	return 0 <= pos.x && pos.x < g.width && 0 <= pos.y && pos.y < g.height
}

func (g grid) index(x int, y int) idx {
	return idx(x + (y * g.width))
}

func (g grid) xy(p idx) xy {
	x := int(p) % g.width
	y := (int(p) - x) / g.height
	return xy{x, y}
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

// Naive manhattan distance function
func (g grid) Distance(from Position, to Position) int {
	a := from.(xy)
	b := to.(xy)
	return abs(a.x-b.x) + abs(a.y-b.y)
}

// Path returns a Waypoint from a position to another position
// TODO: Replace naive implementation with actual pathfinding.
func (g grid) Path(from Position, to Position) Waypoint {
	a := from.(xy)
	b := to.(xy)

	if a == b {
		return nil
	}

	head := &waypoint{pos: a}
	tail := head

	for {
		if a.x < b.x {
			a.x += 1
		} else if a.x > b.x {
			a.x -= 1
		}
		if a.y < b.y {
			a.y += 1
		} else if a.y > b.y {
			a.y -= 1
		} else if a.x == b.x {
			break
		}
		head = head.Add(a)
	}

	return tail
}
