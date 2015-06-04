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
