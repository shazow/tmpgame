package tmpgame

import "testing"

func TestGridIdx(t *testing.T) {
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

	pos := g.xy(idx)
	if pos.x != x || pos.y != y {
		t.Errorf("pos is wrong: %s", pos)
	}
}

func TestGirdMove(t *testing.T) {
	g := NewGrid(10, 10)

	idx := g.index(3, 3)
	idx2 := g.move(idx, Undirected)
	pos := g.xy(idx2)
	if pos.x != 3 || pos.y != 3 {
		t.Errorf("xy is wrong: %s", pos)
	}

	pos = g.xy(g.move(idx, North))
	if pos.x != 3 || pos.y != 4 {
		t.Errorf("xy is wrong: %s", pos)
	}

	pos = g.xy(g.move(idx, South))
	if pos.x != 3 || pos.y != 2 {
		t.Errorf("xy is wrong: %s", pos)
	}

	pos = g.xy(g.move(idx, West))
	if pos.x != 2 || pos.y != 3 {
		t.Errorf("xy is wrong: %s", pos)
	}

	pos = g.xy(g.move(idx, East))
	if pos.x != 4 || pos.y != 3 {
		t.Errorf("xy is wrong: %s", pos)
	}
}
