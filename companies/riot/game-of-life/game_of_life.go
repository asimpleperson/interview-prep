package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

// Game is the controller for a Conway's Game of Life simulation.
// It holds the current board state and tracks the generation number.
type Game struct {
	Board      map[Position]Cell
	Generation int
}

// NewGame creates a Game by parsing live cells from a Life 1.06 input.
func NewGame(scanner *bufio.Scanner) (*Game, error) {
	board, err := parseBoard(scanner)
	if err != nil {
		return nil, err
	}
	return &Game{Board: board, Generation: 0}, nil
}

// Tick advances the simulation by one generation using Conway's rules:
//   - A live cell with fewer than 2 or more than 3 alive neighbors dies.
//   - A dead cell with exactly 3 alive neighbors becomes alive.
func (g *Game) Tick() {
	// Round 1: Build candidate cells with their live neighbor counts.
	// Every position adjacent to a live cell becomes a candidate.
	candidates := make(map[Position]Cell, len(g.Board)*4)

	for pos := range g.Board {
		for _, off := range NeighborOffsets {
			// Guard against int64 overflow at the boundaries.
			neighborX, okX := safeAdd(pos.X, off.X)
			neighborY, okY := safeAdd(pos.Y, off.Y)
			if !okX || !okY {
				continue
			}

			neighborPos := Position{X: neighborX, Y: neighborY}
			cell, exists := candidates[neighborPos]
			if !exists {
				// Create a new dead candidate cell.
				cell = Cell{Position: neighborPos, Alive: false, LiveNeighbors: 0}
				// Carry forward alive status from current board.
				if _, alive := g.Board[neighborPos]; alive {
					cell.Alive = true
				}
			}
			cell.LiveNeighbors++
			candidates[neighborPos] = cell
		}
	}

	// Round 2: Keep only cells that stay alive or become alive based on LiveNeighbors.
	next := make(map[Position]Cell, len(g.Board))
	for pos, cell := range candidates {
		if cell.LiveNeighbors == 3 || (cell.LiveNeighbors == 2 && cell.Alive) {
			cell.Alive = true
			next[pos] = cell
		}
	}

	g.Board = next
	g.Generation++
}

// safeAdd returns a + b and true if the result fits in int64,
// or 0 and false if the addition would overflow.
func safeAdd(a, b int64) (int64, bool) {
	result := a + b
	// Overflow occurs when signs of a and b match but the result's sign differs.
	if (b > 0 && result < a) || (b < 0 && result > a) {
		return 0, false
	}
	return result, true
}

// BoardStatusLife106 returns the current board state formatted in Life 1.06 format.
// Output is sorted by position for deterministic results.
func (g *Game) BoardStatusLife106() string {
	positions := make([]Position, 0, len(g.Board))
	for pos := range g.Board {
		positions = append(positions, pos)
	}

	sort.Slice(positions, func(i, j int) bool {
		if positions[i].X != positions[j].X {
			return positions[i].X < positions[j].X
		}
		return positions[i].Y < positions[j].Y
	})

	var sb strings.Builder
	sb.WriteString("#Life 1.06\n")
	for _, pos := range positions {
		sb.WriteString(FormatCell(g.Board[pos]))
		sb.WriteByte('\n')
	}
	return sb.String()
}

// parseBoard reads a Life 1.06 formatted input and returns the set of live cells.
func parseBoard(scanner *bufio.Scanner) (map[Position]Cell, error) {
	board := make(map[Position]Cell)
	lineNum := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++

		// Strip UTF-8 BOM if present (common on Windows).
		line = strings.TrimPrefix(line, "\xef\xbb\xbf")

		// Skip header and blank lines.
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		cell, err := ParseCell(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum, err)
		}
		board[cell.Position] = cell
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}
	return board, nil
}
