package tmpgame

type Building interface {
	Entity
	Enter(Unit) error
}

//
// Implementations:
//

type Outpost struct {
	Building
	Defender
}

type Town struct {
	Building
}

type Bridge struct {
	Building
}
