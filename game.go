package tmpgame

type Query interface {
	Nearby(Unit) []Unit
}

type game struct {
	Map
	units []Unit
}

func (g *game) Tick() {
	// TODO: ...
}
