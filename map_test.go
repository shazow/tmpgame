package tmpgame

import "testing"

func TestGrid(t *testing.T) {
	g := NewGrid(10, 10)

	idx := g.index(0, 0)
	if idx != 0 {
		t.Errorf("idx is wrong: %d", idx)
	}

	idx = g.index(1, 0)
	if idx != 1 {
		t.Errorf("idx is wrong: %d", idx)
	}

	idx = g.index(0, 1)
	if idx != 10 {
		t.Errorf("idx is wrong: %d", idx)
	}

	x, y := 4, 4
	idx = g.index(x, y)
	if idx != 44 {
		t.Errorf("idx is wrong: %d", idx)
	}

	xx, yy := g.xy(idx)
	if xx != x || yy != y {
		t.Errorf("pos is wrong: %d, %d", xx, yy)
	}
}
