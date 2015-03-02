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

type Attacker interface {
	// Returns damage
	Attack() int
}

type Defender interface {
	Defend(a Attacker)
}

type unit struct {
	Name   string
	Path   Waypoint
	Target *Tile
	Health int
}

func (u *unit) Defend(w weapon) {
	u.Health -= w.Attack()
}

type weapon struct {
	// TODO: We could have different kinds of weapons, like ranged/melee, blunt/sharp, etc.
	Damage int
}

func (w weapon) Attack() int {
	// TODO: Random roll?
	return w.Damage
}

type Merchant struct {
	*unit
	// Number of goods the merchant is storing
	Burden int
	// How much the stored goods are worth
	Value int
}

type Guard struct {
	*unit
	*weapon
}

type Thief struct {
	*unit
	*weapon
}
