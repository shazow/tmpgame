package tmpgame

type Query interface {
	Path(Tile) Waypoint
	Nearby(Unit) []Unit
}

type game struct {
	Map
	units []Unit
}

func (g *game) Tick() {
	// TODO: ...
}
