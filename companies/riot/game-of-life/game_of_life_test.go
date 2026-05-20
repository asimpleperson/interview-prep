package main

import (
	"bufio"
	"strings"
	"testing"
)

func makeGame(positions ...Position) *Game {
	board := make(map[Position]Cell, len(positions))
	for _, pos := range positions {
		board[pos] = Cell{Position: pos, Alive: true}
	}
	return &Game{Board: board, Generation: 0}
}

func boardEqual(a, b map[Position]Cell) bool {
	if len(a) != len(b) {
		return false
	}
	for pos := range a {
		if _, ok := b[pos]; !ok {
			return false
		}
	}
	return true
}

func TestTickBlinker(t *testing.T) {
	game := makeGame(Position{0, -1}, Position{0, 0}, Position{0, 1})

	game.Tick()
	want := makeGame(Position{-1, 0}, Position{0, 0}, Position{1, 0})
	if !boardEqual(game.Board, want.Board) {
		t.Errorf("tick 1: got %v, want %v", game.Board, want.Board)
	}

	game.Tick()
	back := makeGame(Position{0, -1}, Position{0, 0}, Position{0, 1})
	if !boardEqual(game.Board, back.Board) {
		t.Errorf("tick 2: got %v, want %v", game.Board, back.Board)
	}
}

func TestTickBlock(t *testing.T) {
	game := makeGame(
		Position{0, 0}, Position{0, 1},
		Position{1, 0}, Position{1, 1},
	)
	snap := makeGame(
		Position{0, 0}, Position{0, 1},
		Position{1, 0}, Position{1, 1},
	)

	for i := 0; i < 5; i++ {
		game.Tick()
		if !boardEqual(game.Board, snap.Board) {
			t.Fatalf("block changed at gen %d", game.Generation)
		}
	}
}

func TestTickGlider(t *testing.T) {
	game := makeGame(
		Position{0, 1},
		Position{1, 2},
		Position{2, 0}, Position{2, 1}, Position{2, 2},
	)

	for i := 0; i < 4; i++ {
		game.Tick()
	}

	// glider shifts by (1,1) every 4 gens
	want := makeGame(
		Position{1, 2},
		Position{2, 3},
		Position{3, 1}, Position{3, 2}, Position{3, 3},
	)
	if !boardEqual(game.Board, want.Board) {
		t.Errorf("got %v, want %v", game.Board, want.Board)
	}
}

func TestTickLargeCoords(t *testing.T) {
	b := int64(2_000_000_000_000)
	game := makeGame(Position{b, b - 1}, Position{b, b}, Position{b, b + 1})
	want := makeGame(Position{b - 1, b}, Position{b, b}, Position{b + 1, b})

	game.Tick()
	if !boardEqual(game.Board, want.Board) {
		t.Errorf("got %v, want %v", game.Board, want.Board)
	}
}

func TestTickEmpty(t *testing.T) {
	game := &Game{Board: make(map[Position]Cell), Generation: 0}
	game.Tick()
	if len(game.Board) != 0 {
		t.Errorf("empty board produced cells: %v", game.Board)
	}
}

func TestTickSingleCellDies(t *testing.T) {
	game := makeGame(Position{0, 0})
	game.Tick()
	if len(game.Board) != 0 {
		t.Errorf("single cell survived: %v", game.Board)
	}
}

func TestTickOverflow(t *testing.T) {
	max := int64(9223372036854775807)
	min := int64(-9223372036854775808)

	// shouldn't panic when neighbors would overflow
	g1 := makeGame(Position{0, max - 1}, Position{0, max}, Position{1, max})
	g1.Tick()

	g2 := makeGame(Position{min, 0}, Position{min, 1}, Position{min + 1, 0})
	g2.Tick()
}

func TestSafeAdd(t *testing.T) {
	max := int64(9223372036854775807)
	min := int64(-9223372036854775808)

	tests := []struct {
		a, b   int64
		wantOk bool
	}{
		{100, 200, true},
		{-100, -200, true},
		{max, 1, false},
		{min, -1, false},
		{max, 0, true},
		{min, 0, true},
	}

	for _, tt := range tests {
		_, ok := safeAdd(tt.a, tt.b)
		if ok != tt.wantOk {
			t.Errorf("safeAdd(%d, %d): ok=%v, want %v", tt.a, tt.b, ok, tt.wantOk)
		}
	}
}

func TestNewGame(t *testing.T) {
	input := `#Life 1.06
0 1
1 2
2 0
2 1
2 2
-2000000000000 -2000000000000
-2000000000001 -2000000000001
-2000000000000 -2000000000001
`
	game, err := NewGame(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Fatal(err)
	}
	if len(game.Board) != 8 {
		t.Fatalf("got %d cells, want 8", len(game.Board))
	}
	if game.Generation != 0 {
		t.Errorf("gen=%d, want 0", game.Generation)
	}
}

func TestNewGameSkipsComments(t *testing.T) {
	input := `#Life 1.06
# comment
5 5

# another
10 10
`
	game, err := NewGame(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Fatal(err)
	}
	if len(game.Board) != 2 {
		t.Fatalf("got %d cells, want 2", len(game.Board))
	}
}

func TestNewGameBadInput(t *testing.T) {
	input := "#Life 1.06\nnope 2\n"
	_, err := NewGame(bufio.NewScanner(strings.NewReader(input)))
	if err == nil {
		t.Fatal("expected error")
	}
}
