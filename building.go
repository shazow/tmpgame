package tmpgame

type Enclosure interface {
	Entity
	Enter(Unit) error
	Exit(Unit) error
}

//
// Implementations:
//

type Outpost struct {
	Enclosure
	Defender
}

type Town struct {
	Enclosure
}

type Bridge struct {
	Enclosure
}
