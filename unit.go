package tmpgame

type Attacker interface {
	// Returns damage
	Attack() int
	// Position of the attacker
	Position() Position
}

type Defender interface {
	Alive() bool
	Defend(Attacker)
}

type Unit interface {
	Entity
	Defender
}

//
// Implementations:
//

type Merchant struct {
	Unit
	// Number of goods the merchant is storing
	Burden int
	// How much the stored goods are worth
	Value int
}

type Guard struct {
	Unit
	Attacker
}

type Thief struct {
	Unit
	Attacker
}

// unit implements Unit, Defender
type unit struct {
	name     string
	path     Waypoint
	position *Position
	target   *Position
	health   int
}

func (u unit) Name() string {
	return u.name
}

func (u unit) Position() Position {
	return u.position
}

func (u unit) Alive() bool {
	return u.health > 0
}

func (u *unit) Defend(w weapon) {
	u.health -= w.Attack()
}

// weapon implements Attacker
type weapon struct {
	// TODO: We could have different kinds of weapons, like ranged/melee, blunt/sharp, etc.
	damage int
}

func (w weapon) Attack() int {
	// TODO: Random roll?
	return w.damage
}
