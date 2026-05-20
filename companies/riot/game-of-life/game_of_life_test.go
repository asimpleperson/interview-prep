package main

import (
	"bufio"
	"strings"
	"testing"
)

// makeGame is a test helper to create a Game from a list of positions.
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

// TestTickBlinker verifies the period-2 blinker oscillator.
//
//	Generation 0 (vertical):    Generation 1 (horizontal):
//	    . X .                       . . .
//	    . X .                       X X X
//	    . X .                       . . .
func TestTickBlinker(t *testing.T) {
	game := makeGame(
		Position{0, -1}, Position{0, 0}, Position{0, 1},
	)

	game.Tick()
	expected := makeGame(Position{-1, 0}, Position{0, 0}, Position{1, 0})
	if !boardEqual(game.Board, expected.Board) {
		t.Errorf("blinker tick 1: got %v, want %v", game.Board, expected.Board)
	}
	if game.Generation != 1 {
		t.Errorf("expected generation 1, got %d", game.Generation)
	}

	game.Tick()
	reverted := makeGame(Position{0, -1}, Position{0, 0}, Position{0, 1})
	if !boardEqual(game.Board, reverted.Board) {
		t.Errorf("blinker tick 2: got %v, want %v", game.Board, reverted.Board)
	}
	if game.Generation != 2 {
		t.Errorf("expected generation 2, got %d", game.Generation)
	}
}

// TestTickBlock verifies the 2x2 block still life (should never change).
func TestTickBlock(t *testing.T) {
	game := makeGame(
		Position{0, 0}, Position{0, 1},
		Position{1, 0}, Position{1, 1},
	)
	snapshot := makeGame(
		Position{0, 0}, Position{0, 1},
		Position{1, 0}, Position{1, 1},
	)

	for i := 0; i < 5; i++ {
		game.Tick()
		if !boardEqual(game.Board, snapshot.Board) {
			t.Fatalf("block changed at generation %d: %v", game.Generation, game.Board)
		}
	}
}

// TestTickGlider verifies the glider translates by (1, 1) every 4 generations.
func TestTickGlider(t *testing.T) {
	game := makeGame(
		Position{0, 1},
		Position{1, 2},
		Position{2, 0}, Position{2, 1}, Position{2, 2},
	)

	for i := 0; i < 4; i++ {
		game.Tick()
	}

	expected := makeGame(
		Position{1, 2},
		Position{2, 3},
		Position{3, 1}, Position{3, 2}, Position{3, 3},
	)
	if !boardEqual(game.Board, expected.Board) {
		t.Errorf("glider after 4 ticks:\ngot  %v\nwant %v", game.Board, expected.Board)
	}
	if game.Generation != 4 {
		t.Errorf("expected generation 4, got %d", game.Generation)
	}
}

// TestTickLargeCoordinates verifies simulation correctness at extreme int64 positions.
func TestTickLargeCoordinates(t *testing.T) {
	base := int64(2_000_000_000_000)
	game := makeGame(
		Position{base, base - 1}, Position{base, base}, Position{base, base + 1},
	)
	expected := makeGame(
		Position{base - 1, base}, Position{base, base}, Position{base + 1, base},
	)

	game.Tick()
	if !boardEqual(game.Board, expected.Board) {
		t.Errorf("large-coord blinker tick 1: got %v, want %v", game.Board, expected.Board)
	}
}

// TestTickEmptyBoard verifies that an empty board stays empty.
func TestTickEmptyBoard(t *testing.T) {
	game := &Game{Board: make(map[Position]Cell), Generation: 0}
	game.Tick()
	if len(game.Board) != 0 {
		t.Errorf("empty board produced cells: %v", game.Board)
	}
}

// TestTickSingleCell verifies a lone cell dies (underpopulation).
func TestTickSingleCell(t *testing.T) {
	game := makeGame(Position{0, 0})
	game.Tick()
	if len(game.Board) != 0 {
		t.Errorf("single cell survived: %v", game.Board)
	}
}

// TestTickOverflowBoundary verifies that cells at int64 boundaries don't panic
// and that out-of-bound neighbors are safely skipped.
func TestTickOverflowBoundary(t *testing.T) {
	maxInt64 := int64(9223372036854775807)
	minInt64 := int64(-9223372036854775808)

	// A blinker at MaxInt64 Y edge — top neighbor would overflow.
	game := makeGame(
		Position{0, maxInt64 - 1}, Position{0, maxInt64}, Position{1, maxInt64},
	)

	// Should not panic.
	game.Tick()

	if game.Generation != 1 {
		t.Errorf("expected generation 1, got %d", game.Generation)
	}

	// A blinker at MinInt64 X edge — left neighbor would underflow.
	game2 := makeGame(
		Position{minInt64, 0}, Position{minInt64, 1}, Position{minInt64 + 1, 0},
	)

	// Should not panic.
	game2.Tick()

	if game2.Generation != 1 {
		t.Errorf("expected generation 1, got %d", game2.Generation)
	}
}

// TestSafeAdd verifies overflow detection for int64 addition.
func TestSafeAdd(t *testing.T) {
	maxInt64 := int64(9223372036854775807)
	minInt64 := int64(-9223372036854775808)

	tests := []struct {
		name   string
		a, b   int64
		wantOk bool
	}{
		{"no overflow positive", 100, 200, true},
		{"no overflow negative", -100, -200, true},
		{"no overflow mixed", -100, 200, true},
		{"zero", 0, 0, true},
		{"max + 1 overflows", maxInt64, 1, false},
		{"min - 1 underflows", minInt64, -1, false},
		{"max + 0 ok", maxInt64, 0, true},
		{"min + 0 ok", minInt64, 0, true},
		{"large positive sum overflows", maxInt64 / 2 + 1, maxInt64 / 2 + 1, false},
		{"large negative sum underflows", minInt64 / 2 - 1, minInt64 / 2 - 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := safeAdd(tt.a, tt.b)
			if ok != tt.wantOk {
				t.Errorf("safeAdd(%d, %d): got ok=%v, want ok=%v", tt.a, tt.b, ok, tt.wantOk)
			}
			if ok && result != tt.a+tt.b {
				t.Errorf("safeAdd(%d, %d): got %d, want %d", tt.a, tt.b, result, tt.a+tt.b)
			}
		})
	}
}

// TestNewGame verifies creating a game from Life 1.06 input.
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
	scanner := bufio.NewScanner(strings.NewReader(input))
	game, err := NewGame(scanner)
	if err != nil {
		t.Fatalf("NewGame error: %v", err)
	}

	if len(game.Board) != 8 {
		t.Fatalf("expected 8 live cells, got %d", len(game.Board))
	}
	if game.Generation != 0 {
		t.Errorf("expected generation 0, got %d", game.Generation)
	}

	checks := []Position{
		{0, 1},
		{2, 2},
		{-2000000000000, -2000000000000},
	}
	for _, p := range checks {
		cell, ok := game.Board[p]
		if !ok {
			t.Errorf("expected cell at %v to be present", p)
		}
		if !cell.Alive {
			t.Errorf("expected cell at %v to be alive", p)
		}
	}
}

// TestNewGameSkipsBlanksAndComments verifies that blank lines and comment lines are ignored.
func TestNewGameSkipsBlanksAndComments(t *testing.T) {
	input := `#Life 1.06
# this is a comment

5 5

# another comment
10 10
`
	scanner := bufio.NewScanner(strings.NewReader(input))
	game, err := NewGame(scanner)
	if err != nil {
		t.Fatalf("NewGame error: %v", err)
	}

	if len(game.Board) != 2 {
		t.Fatalf("expected 2 live cells, got %d", len(game.Board))
	}
}

// TestNewGameInvalidInput verifies that malformed input returns an error.
func TestNewGameInvalidInput(t *testing.T) {
	input := `#Life 1.06
0 1
not_a_number 2
`
	scanner := bufio.NewScanner(strings.NewReader(input))
	_, err := NewGame(scanner)
	if err == nil {
		t.Fatal("expected error for invalid input, got nil")
	}
}
