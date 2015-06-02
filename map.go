package tmpgame

type Tile interface {
}

type Map interface {
	// From Tile -> To Tile
	Path(Tile, Tile) Waypoint
}

type Waypoint interface {
	Tile() Tile
	Next() Waypoint
}
